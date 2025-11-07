package lua

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LoadScriptsFromDir loads all .lua files from a directory
func (e *Engine) LoadScriptsFromDir(ctx context.Context, dir string) error {
	e.logger.Infof("📂 Loading Lua scripts from directory: %s", dir)

	// Check if directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		e.logger.Warnf("Scripts directory does not exist: %s", dir)
		return nil // Not an error, just no scripts to load
	}

	// Walk through directory
	var loadedCount int
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-.lua files
		if info.IsDir() || !strings.HasSuffix(path, ".lua") {
			return nil
		}

		// Load the script
		e.logger.Infof("📄 Loading script: %s", path)
		if err := e.LoadScriptFile(ctx, path); err != nil {
			e.logger.Errorf("❌ Failed to load script %s: %v", path, err)
			// Continue loading other scripts even if one fails
			return nil
		}

		e.logger.Infof("✅ Successfully loaded script: %s", path)
		loadedCount++
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk scripts directory: %w", err)
	}

	e.logger.Infof("Loaded %d Lua scripts from %s", loadedCount, dir)
	return nil
}

// LoadScriptFile loads and executes a single Lua script file
func (e *Engine) LoadScriptFile(ctx context.Context, filePath string) error {
	e.logger.Debugf("Loading script file: %s", filePath)

	// Read file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read script file: %w", err)
	}

	// Get a VM from pool
	L := e.pool.Get()

	// Defer: check if VM is dedicated before returning to pool
	defer func() {
		e.mu.RLock()
		isDedicated := e.dedicatedVMs[L]
		e.mu.RUnlock()

		if !isDedicated {
			// Return to pool only if not marked as dedicated
			e.pool.Put(L)
		} else {
			e.logger.Debugf("VM marked as dedicated, not returning to pool")
		}
	}()

	// Set context
	L.SetContext(ctx)

	// Execute the script
	// This allows the script to call hook.register() and hook.add_script()
	if err := L.DoString(string(content)); err != nil {
		return fmt.Errorf("failed to execute script: %w", err)
	}

	e.logger.Debugf("Successfully loaded script: %s", filePath)
	return nil
}

// LoadScriptString loads and executes a Lua script from a string
func (e *Engine) LoadScriptString(ctx context.Context, scriptName, source string) error {
	e.logger.Debugf("Loading script: %s", scriptName)

	// Get a VM from pool
	L := e.pool.Get()

	// Defer: check if VM is dedicated before returning to pool
	defer func() {
		e.mu.RLock()
		isDedicated := e.dedicatedVMs[L]
		e.mu.RUnlock()

		if !isDedicated {
			// Return to pool only if not marked as dedicated
			e.pool.Put(L)
		} else {
			e.logger.Debugf("VM marked as dedicated, not returning to pool")
		}
	}()

	// Set context
	L.SetContext(ctx)

	// Execute the script
	if err := L.DoString(source); err != nil {
		return fmt.Errorf("failed to execute script %s: %w", scriptName, err)
	}

	e.logger.Debugf("Successfully loaded script: %s", scriptName)
	return nil
}
