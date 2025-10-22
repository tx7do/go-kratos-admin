package lua

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

func TestLoadScriptFile(t *testing.T) {
	engine := NewEngine(DefaultConfig(), log.DefaultLogger)
	defer engine.Close()

	// Create a temporary script file
	tmpDir := t.TempDir()
	scriptPath := filepath.Join(tmpDir, "test_script.lua")

	scriptContent := `
		local log = require "kratos_logger"
		local hook = require "kratos_hook"

		log.info("Test script loaded")

		-- Register a test hook with callback
		hook.register("test.load", "Test load hook", function(ctx)
			log.info("Test callback executed")
			ctx.set("loaded", true)
			return true
		end)
	`

	err := os.WriteFile(scriptPath, []byte(scriptContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test script: %v", err)
	}

	// Load the script
	err = engine.LoadScriptFile(context.Background(), scriptPath)
	if err != nil {
		t.Fatalf("Failed to load script: %v", err)
	}

	// Execute the hook to verify callback was registered
	ctx := &Context{
		ID:       "test-load",
		HookName: "test.load",
		Data:     make(map[string]interface{}),
	}

	err = engine.ExecuteHook(context.Background(), "test.load", ctx)
	if err != nil {
		t.Fatalf("Failed to execute hook: %v", err)
	}

	// Verify callback executed
	if loaded, ok := ctx.Data["loaded"].(bool); !ok || !loaded {
		t.Error("Callback was not executed")
	}

	t.Log("✓ Script loading test passed")
}

func TestLoadScriptsFromDir(t *testing.T) {
	engine := NewEngine(DefaultConfig(), log.DefaultLogger)
	defer engine.Close()

	// Create a temporary directory with multiple scripts
	tmpDir := t.TempDir()

	// Script 1: Register hook1
	script1 := `
		local log = require "kratos_logger"
		local hook = require "kratos_hook"

		hook.register("dir.test.hook1", "Hook 1", function(ctx)
			local count = ctx.get("count") or 0
			ctx.set("count", count + 1)
			log.info("Hook 1 executed")
			return true
		end)
	`
	err := os.WriteFile(filepath.Join(tmpDir, "script1.lua"), []byte(script1), 0644)
	if err != nil {
		t.Fatalf("Failed to write script1: %v", err)
	}

	// Script 2: Register hook2
	script2 := `
		local log = require "kratos_logger"
		local hook = require "kratos_hook"

		hook.register("dir.test.hook2", "Hook 2", function(ctx)
			local count = ctx.get("count") or 0
			ctx.set("count", count + 1)
			log.info("Hook 2 executed")
			return true
		end)
	`
	err = os.WriteFile(filepath.Join(tmpDir, "script2.lua"), []byte(script2), 0644)
	if err != nil {
		t.Fatalf("Failed to write script2: %v", err)
	}

	// Non-Lua file (should be ignored)
	err = os.WriteFile(filepath.Join(tmpDir, "readme.txt"), []byte("Not a Lua script"), 0644)
	if err != nil {
		t.Fatalf("Failed to write readme: %v", err)
	}

	// Load all scripts from directory
	err = engine.LoadScriptsFromDir(context.Background(), tmpDir)
	if err != nil {
		t.Fatalf("Failed to load scripts: %v", err)
	}

	// Execute both hooks
	ctx := &Context{
		ID:   "test-dir",
		Data: make(map[string]interface{}),
	}

	ctx.HookName = "dir.test.hook1"
	err = engine.ExecuteHook(context.Background(), "dir.test.hook1", ctx)
	if err != nil {
		t.Fatalf("Failed to execute hook1: %v", err)
	}

	ctx.HookName = "dir.test.hook2"
	err = engine.ExecuteHook(context.Background(), "dir.test.hook2", ctx)
	if err != nil {
		t.Fatalf("Failed to execute hook2: %v", err)
	}

	// Verify both hooks executed
	count, ok := ctx.Data["count"].(float64)
	if !ok {
		t.Error("Count not set")
	} else if int(count) != 2 {
		t.Errorf("Expected count to be 2, got: %d", int(count))
	}

	t.Log("✓ Directory loading test passed")
}

func TestLoadScriptsFromDir_NonExistent(t *testing.T) {
	engine := NewEngine(DefaultConfig(), log.DefaultLogger)
	defer engine.Close()

	// Try to load from non-existent directory (should not error)
	err := engine.LoadScriptsFromDir(context.Background(), "/non/existent/path")
	if err != nil {
		t.Errorf("Expected no error for non-existent directory, got: %v", err)
	}

	t.Log("✓ Non-existent directory test passed")
}

func TestLoadScriptString(t *testing.T) {
	engine := NewEngine(DefaultConfig(), log.DefaultLogger)
	defer engine.Close()

	script := `
		local log = require "kratos_logger"
		local hook = require "kratos_hook"

		hook.register("string.test", "String test hook", function(ctx)
			ctx.set("from_string", true)
			log.info("String script executed")
			return true
		end)
	`

	err := engine.LoadScriptString(context.Background(), "test_string", script)
	if err != nil {
		t.Fatalf("Failed to load script string: %v", err)
	}

	// Execute the hook
	ctx := &Context{
		ID:       "test-string",
		HookName: "string.test",
		Data:     make(map[string]interface{}),
	}

	err = engine.ExecuteHook(context.Background(), "string.test", ctx)
	if err != nil {
		t.Fatalf("Failed to execute hook: %v", err)
	}

	// Verify
	if fromString, ok := ctx.Data["from_string"].(bool); !ok || !fromString {
		t.Error("String script callback was not executed")
	}

	t.Log("✓ String script loading test passed")
}
