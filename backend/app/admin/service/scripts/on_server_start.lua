-- on_server_start.lua
-- This script runs when the server starts

-- Load required modules
local log = require "kratos_logger"
local hook = require "kratos_hook"

log.info("========================================")
log.info("  Server Startup Script Executing")
log.info("========================================")

-- Register the on_server_start hook and attach callback
hook.register("on_server_start", "Triggered when server starts", function(ctx)
    log.info("✓ Server is starting up...")

    -- Access configuration passed from Go
    local config = ctx.get("config")
    local service = ctx.get("service")

    -- Debug: Log what we actually received
    log.info("DEBUG: config type = " .. type(config))
    log.info("DEBUG: config value = " .. tostring(config))
    log.info("DEBUG: service type = " .. type(service))
    log.info("DEBUG: service value = " .. tostring(service))

    -- Example: Log configuration details
    if config then
        log.info("✓ Configuration loaded")
        if type(config) == "table" then
            log.info("Configuration keys:")
            for key, value in pairs(config) do
                -- Be careful with logging complex values
                if type(value) == "string" or type(value) == "number" or type(value) == "boolean" then
                    log.info("  " .. key .. " = " .. tostring(value))
                else
                    log.info("  " .. key .. " = [" .. type(value) .. "]")
                end
            end
        else
            log.warn("Config is not a table, it's: " .. type(config))
        end
        -- You can access config fields here
        -- Note: Complex nested structures are converted to Lua tables
    else
        log.warn("Config is nil!")
    end

    if service then
        log.info("✓ Service configuration loaded")
        if type(service) == "table" then
            log.info("Service is a table")
        else
            log.warn("Service is not a table, it's: " .. type(service))
        end
    else
        log.warn("Service is nil!")
    end

    -- You can perform startup tasks here:
    -- - Check system health
    -- - Initialize cache
    -- - Load configuration
    -- - Register event handlers
    -- - Set up initial data

    -- Example: Set a startup timestamp in cache (if Redis is configured)
    -- cache.set("server_start_time", os.time(), 0)

    -- Example: Publish server start event (if EventBus is configured)
    -- eventbus.publish("server.started", {
    --     timestamp = os.time(),
    --     version = "1.0.0"
    -- }, "system")

    log.info("✓ Initializing services...")
    log.info("✓ Startup hook executed successfully")

    return true
end)

log.info("✓ on_server_start hook registered")
