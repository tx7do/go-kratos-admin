package lua

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	lua "github.com/yuin/gopher-lua"

	"kratos-admin/pkg/eventbus"
	"kratos-admin/pkg/lua/api"
	"kratos-admin/pkg/lua/hook"
	"kratos-admin/pkg/lua/internal/convert"
	"kratos-admin/pkg/oss"
)

// CallbackInfo stores information about a registered callback
type CallbackInfo struct {
	L        *lua.LState
	Function *lua.LFunction
	HookName string
}

// Engine manages Lua VM lifecycle and execution
type Engine struct {
	config          *Config
	pool            *vmPool
	logger          *log.Helper
	registry        *hook.Registry
	rdb             *redis.Client         // Redis client for cache operations
	eventbusManager *eventbus.Manager     // EventBus manager
	ossClient       *oss.MinIOClient      // OSS/MinIO client
	callbacks       map[string][]*CallbackInfo // Hook callbacks (hook name -> multiple callbacks)
	dedicatedVMs    map[*lua.LState]bool     // VMs that should not be pooled
	mu              sync.RWMutex
}

// Config defines Lua engine configuration
type Config struct {
	MaxVMs         int           // Maximum concurrent VMs (default: 10)
	VMTimeout      time.Duration // Execution timeout per script (default: 5s)
	MaxMemory      int64         // Memory limit per VM in bytes (default: 50MB)
	EnableDebug    bool          // Enable debug logging
	ScriptDir      string        // Directory for file-based scripts
	AllowedModules []string      // Allowed Lua modules
	PoolSize       int           // VM pool size (default: 5)
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		MaxVMs:         10,
		VMTimeout:      5 * time.Second,
		MaxMemory:      50 * 1024 * 1024, // 50MB
		EnableDebug:    false,
		ScriptDir:      "scripts",
		AllowedModules: []string{},
		PoolSize:       5,
	}
}

// NewEngine creates a new Lua engine
func NewEngine(config *Config, logger log.Logger) *Engine {
	if config == nil {
		config = DefaultConfig()
	}

	l := log.NewHelper(log.With(logger, "module", "lua/engine"))

	engine := &Engine{
		config:       config,
		logger:       l,
		registry:     hook.NewRegistry(),
		callbacks:    make(map[string][]*CallbackInfo),
		dedicatedVMs: make(map[*lua.LState]bool),
	}

	// Initialize VM pool
	engine.pool = newVMPool(config.PoolSize, func() *lua.LState {
		return engine.createVM()
	})

	l.Infof("Lua engine initialized (pool: %d, timeout: %s)", config.PoolSize, config.VMTimeout)

	// Automatically load scripts from ScriptDir if configured
	if config.ScriptDir != "" {
		if err := engine.LoadScriptsFromDir(context.Background(), config.ScriptDir); err != nil {
			l.Errorf("Failed to load scripts from %s: %v", config.ScriptDir, err)
		}
	}

	return engine
}

// createVM creates a new Lua VM with sandbox and API bindings
func (e *Engine) createVM() *lua.LState {
	L := lua.NewState(lua.Options{
		CallStackSize:       120,
		RegistrySize:        120 * 20,
		SkipOpenLibs:        true, // We'll selectively open safe libs
		IncludeGoStackTrace: e.config.EnableDebug,
	})

	// Open safe standard libraries
	e.openSafeLibs(L)

	// Register custom API modules
	e.registerAPIs(L)

	return L
}

// openSafeLibs opens only safe Lua standard libraries
func (e *Engine) openSafeLibs(L *lua.LState) {
	// Safe libraries
	lua.OpenBase(L)
	lua.OpenTable(L)
	lua.OpenString(L)
	lua.OpenMath(L)

	// Remove dangerous functions from base
	L.SetGlobal("dofile", lua.LNil)
	L.SetGlobal("loadfile", lua.LNil)
	L.SetGlobal("load", lua.LNil)
	L.SetGlobal("loadstring", lua.LNil)

	// Don't open full package library (unsafe), but create package.preload table
	// for our module system to work
	packageTable := L.NewTable()
	preloadTable := L.NewTable()
	packageTable.RawSetString("preload", preloadTable)
	L.SetGlobal("package", packageTable)

	// Create minimal require() function
	L.SetGlobal("require", L.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)

		// Check if module is already loaded in package.loaded
		pkg := L.GetGlobal("package")
		if pkg == lua.LNil {
			L.RaiseError("package table not found")
			return 0
		}
		pkgTable := pkg.(*lua.LTable)

		loaded := pkgTable.RawGetString("loaded")
		if loaded == lua.LNil {
			loaded = L.NewTable()
			pkgTable.RawSetString("loaded", loaded)
		}
		loadedTable := loaded.(*lua.LTable)

		// Return cached module if already loaded
		cached := loadedTable.RawGetString(name)
		if cached != lua.LNil {
			L.Push(cached)
			return 1
		}

		// Get loader from package.preload
		preload := pkgTable.RawGetString("preload")
		if preload == lua.LNil {
			L.RaiseError("module '%s' not found in package.preload", name)
			return 0
		}
		preloadTable := preload.(*lua.LTable)

		loader := preloadTable.RawGetString(name)
		if loader == lua.LNil {
			L.RaiseError("module '%s' not found", name)
			return 0
		}

		// Call the loader function
		L.Push(loader)
		L.Call(0, 1)

		// Cache the result in package.loaded
		result := L.Get(-1)
		loadedTable.RawSetString(name, result)

		L.Push(result)
		return 1
	}))
}

// registerAPIs registers custom Go APIs for Lua
func (e *Engine) registerAPIs(L *lua.LState) {
	// Register logger API from api package
	api.RegisterLogger(L, e.logger)

	// Register cache API if Redis client is available
	if e.rdb != nil {
		api.RegisterCache(L, e.rdb, e.logger)
	}

	// Register eventbus API if manager is available
	if e.eventbusManager != nil {
		api.RegisterEventBus(L, e.eventbusManager, e.logger)
	}

	// Register OSS API if client is available
	if e.ossClient != nil {
		api.RegisterOSS(L, e.ossClient, e.logger)
	}

	// Register Crypto API (always available - uses global encryptor)
	api.RegisterCrypto(L, e.logger)

	// Register hook API for self-registration
	api.RegisterHookAPI(L, e, e.logger)

	// Register task API for task handler registration
	api.RegisterTask(L, e, e.logger)

	// Register utility API (sleep, time, etc.)
	api.RegisterUtilAPI(L, e.logger)

	// Set up module aliases and global convenience variables
	// This allows scripts to use both require('logger') and log.info() directly
	preloadTable := L.GetGlobal("package").(*lua.LTable).RawGetString("preload").(*lua.LTable)

	// Create shorter aliases for modules in package.preload
	// So scripts can use require('logger') instead of require('kratos_logger')
	if logLoader := preloadTable.RawGetString("kratos_logger"); logLoader != lua.LNil {
		preloadTable.RawSetString("logger", logLoader)
	}
	if hookLoader := preloadTable.RawGetString("kratos_hook"); hookLoader != lua.LNil {
		preloadTable.RawSetString("hook", hookLoader)
	}
	if utilLoader := preloadTable.RawGetString("kratos_util"); utilLoader != lua.LNil {
		preloadTable.RawSetString("util", utilLoader)
	}

	// Load logger and set as global 'log' for convenience
	if logLoader, ok := preloadTable.RawGetString("logger").(*lua.LFunction); ok {
		L.Push(logLoader)
		if err := L.PCall(0, 1, nil); err == nil {
			L.SetGlobal("log", L.Get(-1))
			L.Pop(1)
		}
	}

	// Load hook and set as global 'hook' for convenience
	if hookLoader, ok := preloadTable.RawGetString("hook").(*lua.LFunction); ok {
		L.Push(hookLoader)
		if err := L.PCall(0, 1, nil); err == nil {
			L.SetGlobal("hook", L.Get(-1))
			L.Pop(1)
		}
	}

	// Future APIs will be registered here:
	// api.RegisterHTTP(L, httpClient)
	// api.RegisterDatabase(L, db)
	// etc.
}

// Execute executes a Lua script with given context
func (e *Engine) Execute(ctx context.Context, script *Script, execCtx *Context) error {
	// Get VM from pool
	L := e.pool.Get()
	defer e.pool.Put(L)

	// Set execution context
	if err := e.setContext(L, execCtx); err != nil {
		return fmt.Errorf("failed to set context: %w", err)
	}

	// Create timeout context
	timeoutCtx, cancel := context.WithTimeout(ctx, e.config.VMTimeout)
	defer cancel()

	// Execute script with timeout
	errChan := make(chan error, 1)
	go func() {
		// Set context in VM
		L.SetContext(timeoutCtx)

		// Load and execute script
		if err := L.DoString(script.Source); err != nil {
			errChan <- fmt.Errorf("script execution error: %w", err)
			return
		}

		// Call execute function if exists
		if fn := L.GetGlobal("execute"); fn.Type() == lua.LTFunction {
			// Push function and arguments
			L.Push(fn)
			L.Push(e.contextToLuaTable(L, execCtx))

			// Call function (1 argument, 1 return value)
			if err := L.PCall(1, 1, nil); err != nil {
				errChan <- fmt.Errorf("execute function error: %w", err)
				return
			}

			// Get result
			ret := L.Get(-1)
			L.Pop(1)

			// Check if script returned false (abort)
			if ret.Type() == lua.LTBool && !lua.LVAsBool(ret) {
				errChan <- fmt.Errorf("script returned false")
				return
			}
		}

		errChan <- nil
	}()

	// Wait for completion or timeout
	select {
	case err := <-errChan:
		return err
	case <-timeoutCtx.Done():
		// Force stop the VM
		L.Close()
		// Create new VM for pool
		newL := e.createVM()
		e.pool.Replace(L, newL)
		return fmt.Errorf("script execution timeout after %s", e.config.VMTimeout)
	}
}

// ExecuteHook executes all scripts registered for a hook
func (e *Engine) ExecuteHook(ctx context.Context, hookName string, execCtx *Context) error {
	// Check for callbacks
	e.mu.RLock()
	callbacks := e.callbacks[hookName]
	e.mu.RUnlock()

	// Execute all callbacks if registered
	for i, callback := range callbacks {
		start := time.Now()
		err := e.executeCallback(ctx, callback, execCtx)
		duration := time.Since(start)

		if err != nil {
			e.logger.Errorf("Callback %d failed (hook: %s, duration: %s): %v",
				i+1, hookName, duration, err)
			return fmt.Errorf("callback %d failed: %w", i+1, err)
		}

		e.logger.Debugf("Callback %d completed (hook: %s, duration: %s)",
			i+1, hookName, duration)
	}

	// Execute registered scripts
	hookScripts := e.registry.GetScripts(hookName)
	if len(hookScripts) == 0 && len(callbacks) == 0 {
		e.logger.Debugf("No scripts or callbacks registered for hook: %s", hookName)
		return nil
	}

	e.logger.Debugf("Executing %d scripts for hook: %s", len(hookScripts), hookName)

	for _, hookScript := range hookScripts {
		if !hookScript.Enabled {
			continue
		}

		// Convert hook.Script to lua.Script
		script := &Script{
			ID:          hookScript.ID,
			Name:        hookScript.Name,
			Hook:        hookScript.Hook,
			Source:      hookScript.Source,
			Enabled:     hookScript.Enabled,
			Priority:    hookScript.Priority,
			Description: hookScript.Description,
			Version:     hookScript.Version,
			Author:      hookScript.Author,
			Critical:    hookScript.Critical,
		}

		start := time.Now()
		err := e.Execute(ctx, script, execCtx)
		duration := time.Since(start)

		if err != nil {
			e.logger.Errorf("Script '%s' failed (hook: %s, duration: %s): %v",
				script.Name, hookName, duration, err)
			// Always return error if script fails (including returning false)
			// This allows hooks to abort processing
			return fmt.Errorf("script '%s' failed: %w", script.Name, err)
		}

		e.logger.Debugf("Script '%s' completed (hook: %s, duration: %s)",
			script.Name, hookName, duration)
	}

	return nil
}

// executeCallback executes a registered callback function
func (e *Engine) executeCallback(ctx context.Context, callback *CallbackInfo, execCtx *Context) error {
	L := callback.L

	// Create timeout context
	timeoutCtx, cancel := context.WithTimeout(ctx, e.config.VMTimeout)
	defer cancel()

	// Execute callback with timeout
	errChan := make(chan error, 1)
	go func() {
		// Set context in VM
		L.SetContext(timeoutCtx)

		// Push function and context argument
		L.Push(callback.Function)
		L.Push(e.contextToLuaTable(L, execCtx))

		// Call function (1 argument, 1 return value)
		if err := L.PCall(1, 1, nil); err != nil {
			errChan <- fmt.Errorf("callback execution error: %w", err)
			return
		}

		// Get result
		ret := L.Get(-1)
		L.Pop(1)

		// Check if callback returned false (abort)
		if ret.Type() == lua.LTBool && !lua.LVAsBool(ret) {
			errChan <- fmt.Errorf("callback returned false")
			return
		}

		errChan <- nil
	}()

	// Wait for completion or timeout
	select {
	case err := <-errChan:
		return err
	case <-timeoutCtx.Done():
		return fmt.Errorf("callback execution timeout after %s", e.config.VMTimeout)
	}
}

// RegisterHook registers a hook point
func (e *Engine) RegisterHook(name, description string) error {
	return e.registry.RegisterHook(name, description)
}

// AddScript adds a script to a hook
// Accepts *Script or any struct with compatible fields
func (e *Engine) AddScript(hookName string, script interface{}) error {
	var hookScript *hook.Script

	switch s := script.(type) {
	case *Script:
		// Convert lua.Script to hook.Script
		hookScript = &hook.Script{
			ID:          s.ID,
			Name:        s.Name,
			Hook:        s.Hook,
			Source:      s.Source,
			Enabled:     s.Enabled,
			Priority:    s.Priority,
			Description: s.Description,
			Version:     s.Version,
			Author:      s.Author,
			Critical:    s.Critical,
		}
	default:
		// Try to extract fields using reflection for api.Script or similar types
		// This avoids circular dependency
		if scriptStruct, ok := extractScriptFields(script); ok {
			hookScript = scriptStruct
		} else {
			return fmt.Errorf("invalid script type: %T", script)
		}
	}

	return e.registry.AddScript(hookName, hookScript)
}

// extractScriptFields uses reflection to extract script fields from any compatible struct
func extractScriptFields(script interface{}) (*hook.Script, bool) {
	v := reflect.ValueOf(script)

	// Dereference pointer if needed
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, false
	}

	// Helper to safely get string field
	getStringField := func(name string) string {
		field := v.FieldByName(name)
		if field.IsValid() && field.Kind() == reflect.String {
			return field.String()
		}
		return ""
	}

	// Helper to safely get bool field
	getBoolField := func(name string) bool {
		field := v.FieldByName(name)
		if field.IsValid() && field.Kind() == reflect.Bool {
			return field.Bool()
		}
		return false
	}

	// Helper to safely get int field
	getIntField := func(name string) int {
		field := v.FieldByName(name)
		if field.IsValid() && field.Kind() == reflect.Int {
			return int(field.Int())
		}
		return 0
	}

	// Extract fields
	hookScript := &hook.Script{
		ID:          0, // auto-generated (hook.Script uses uint32)
		Name:        getStringField("Name"),
		Hook:        getStringField("Hook"),
		Source:      getStringField("Source"),
		Enabled:     getBoolField("Enabled"),
		Priority:    getIntField("Priority"),
		Description: getStringField("Description"),
		Version:     0, // default version (hook.Script uses int)
		Author:      getStringField("Author"),
		Critical:    getBoolField("Critical"),
	}

	// Validate that we got at least the required fields
	if hookScript.Name == "" || hookScript.Source == "" {
		return nil, false
	}

	return hookScript, true
}

// RemoveScript removes a script from a hook
func (e *Engine) RemoveScript(hookName, scriptName string) error {
	return e.registry.RemoveScript(hookName, scriptName)
}

// ListHooks returns all registered hooks
func (e *Engine) ListHooks() []string {
	return e.registry.ListHooks()
}

// RegisterCallback registers a Lua callback function for a hook
func (e *Engine) RegisterCallback(hookName string, L *lua.LState, fn *lua.LFunction) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Append callback to the list for this hook
	e.callbacks[hookName] = append(e.callbacks[hookName], &CallbackInfo{
		L:        L,
		Function: fn,
		HookName: hookName,
	})

	// Mark this VM as dedicated (should not be pooled)
	e.dedicatedVMs[L] = true

	callbackCount := len(e.callbacks[hookName])
	e.logger.Infof("Registered callback for hook: %s (total: %d callbacks)", hookName, callbackCount)
}

// MarkVMDedicated marks a VM as dedicated (won't be returned to pool)
func (e *Engine) MarkVMDedicated(L *lua.LState) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.dedicatedVMs[L] = true
	e.logger.Debugf("VM marked as dedicated")
}

// SetRedis sets the Redis client for cache operations
func (e *Engine) SetRedis(rdb *redis.Client) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.rdb = rdb
	e.logger.Info("Redis client configured for Lua cache API")
}

// SetEventBus sets the EventBus manager for event operations
func (e *Engine) SetEventBus(manager *eventbus.Manager) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.eventbusManager = manager
	e.logger.Info("EventBus manager configured for Lua eventbus API")
}

// SetOSS sets the OSS client for object storage operations
func (e *Engine) SetOSS(client *oss.MinIOClient) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ossClient = client
	e.logger.Info("OSS client configured for Lua OSS API")
}

// setContext sets the execution context in the VM
func (e *Engine) setContext(L *lua.LState, ctx *Context) error {
	// Store context as upvalue for API functions
	L.SetGlobal("_context", lua.LString(ctx.ID))
	return nil
}

// contextToLuaTable converts Context to Lua table
func (e *Engine) contextToLuaTable(L *lua.LState, ctx *Context) *lua.LTable {
	table := L.NewTable()

	// Add context methods
	table.RawSetString("get", L.NewFunction(func(L *lua.LState) int {
		key := L.CheckString(1)
		if val, ok := ctx.Data[key]; ok {
			L.Push(convert.ToLuaValue(L, val))
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	table.RawSetString("set", L.NewFunction(func(L *lua.LState) int {
		key := L.CheckString(1)
		val := L.Get(2)
		ctx.Data[key] = convert.ToGoValue(val)
		return 0
	}))

	table.RawSetString("stop", L.NewFunction(func(L *lua.LState) int {
		reason := L.CheckString(1)
		ctx.Stopped = true
		ctx.StopReason = reason
		return 0
	}))

	return table
}

// Close closes the engine and all VMs
func (e *Engine) Close() error {
	e.logger.Info("Closing Lua engine...")

	// Close dedicated VMs first
	e.mu.Lock()
	for vm := range e.dedicatedVMs {
		if vm != nil {
			// Safely close the VM
			func() {
				defer func() {
					if r := recover(); r != nil {
						e.logger.Warnf("Recovered from panic while closing VM: %v", r)
					}
				}()
				vm.Close()
			}()
		}
	}
	e.dedicatedVMs = make(map[*lua.LState]bool)
	e.mu.Unlock()

	// Close pooled VMs
	e.pool.Close()

	return nil
}

// vmPool manages a pool of Lua VMs for reuse
type vmPool struct {
	vms     chan *lua.LState
	factory func() *lua.LState
	size    int
	closed  bool
	mu      sync.Mutex
}

func newVMPool(size int, factory func() *lua.LState) *vmPool {
	pool := &vmPool{
		vms:     make(chan *lua.LState, size),
		factory: factory,
		size:    size,
	}

	// Pre-create VMs
	for i := 0; i < size; i++ {
		pool.vms <- factory()
	}

	return pool
}

func (p *vmPool) Get() *lua.LState {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return p.factory()
	}
	p.mu.Unlock()

	select {
	case vm := <-p.vms:
		return vm
	default:
		// Pool empty, create new VM
		return p.factory()
	}
}

func (p *vmPool) Put(vm *lua.LState) {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		vm.Close()
		return
	}
	p.mu.Unlock()

	// Reset VM state
	vm.SetTop(0)

	select {
	case p.vms <- vm:
		// Returned to pool
	default:
		// Pool full, close VM
		vm.Close()
	}
}

func (p *vmPool) Replace(old, new *lua.LState) {
	old.Close()
	p.Put(new)
}

func (p *vmPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return
	}
	p.closed = true

	close(p.vms)
	for vm := range p.vms {
		vm.Close()
	}
}
