package crypto

import (
	"encoding/json"
	"fmt"
)

const (
	// EncryptedConfigKey is the key used to store encrypted configuration in task payload
	EncryptedConfigKey = "_encrypted_config"
	// IsEncryptedKey indicates if the payload contains encrypted data
	IsEncryptedKey = "_is_encrypted"
)

// EncryptPayload encrypts the entire payload and returns a map with encrypted data
// This is used to store encrypted configuration in Redis/Asynq
func EncryptPayload(payload map[string]interface{}) (map[string]interface{}, error) {
	// Marshal the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Encrypt the JSON
	encrypted, err := EncryptIfNeeded(string(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt payload: %w", err)
	}

	// Return a map with encrypted config and metadata
	result := map[string]interface{}{
		EncryptedConfigKey: encrypted,
		IsEncryptedKey:     true,
	}

	// Preserve non-sensitive fields that might be needed for routing/scheduling
	if taskID, ok := payload["task_id"]; ok {
		result["task_id"] = taskID
	}
	if taskType, ok := payload["task_type"]; ok {
		result["task_type"] = taskType
	}

	return result, nil
}

// DecryptPayload decrypts the payload if it contains encrypted configuration
// Returns the decrypted payload map
func DecryptPayload(payload map[string]interface{}) (map[string]interface{}, error) {
	// Check if payload is encrypted
	isEncrypted, ok := payload[IsEncryptedKey].(bool)
	if !ok || !isEncrypted {
		// Not encrypted, return as-is for backward compatibility
		return payload, nil
	}

	// Get encrypted config
	encryptedConfig, ok := payload[EncryptedConfigKey].(string)
	if !ok {
		return nil, fmt.Errorf("encrypted config not found or invalid type")
	}

	// Decrypt the config
	decrypted, err := DecryptIfNeeded(encryptedConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt payload: %w", err)
	}

	// Unmarshal back to map
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(decrypted), &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal decrypted payload: %w", err)
	}

	return result, nil
}

// MustEncryptPayload encrypts payload and panics on error (useful for testing)
func MustEncryptPayload(payload map[string]interface{}) map[string]interface{} {
	encrypted, err := EncryptPayload(payload)
	if err != nil {
		panic(err)
	}
	return encrypted
}

// MustDecryptPayload decrypts payload and panics on error (useful for testing)
func MustDecryptPayload(payload map[string]interface{}) map[string]interface{} {
	decrypted, err := DecryptPayload(payload)
	if err != nil {
		panic(err)
	}
	return decrypted
}

// HasEncryptedPayload checks if the payload contains encrypted configuration
func HasEncryptedPayload(payload map[string]interface{}) bool {
	isEncrypted, ok := payload[IsEncryptedKey].(bool)
	return ok && isEncrypted
}
