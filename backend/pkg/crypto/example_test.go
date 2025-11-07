package crypto_test

import (
	"fmt"
	"log"

	"kratos-admin/pkg/crypto"
)

func ExampleEncryptor_Encrypt() {
	// Create an encryptor with a secret key
	encryptor, err := crypto.NewEncryptor("my-secret-key")
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt sensitive data
	plaintext := `{"username":"admin","password":"secret123"}`
	encrypted, err := encryptor.Encrypt(plaintext)
	if err != nil {
		log.Fatal(err)
	}

	// Encrypted data has the "enc:" prefix
	fmt.Println("Data encrypted:", crypto.IsEncrypted(encrypted))
	// Output: Data encrypted: true
}

func ExampleEncryptor_Decrypt() {
	// Create an encryptor with a secret key
	encryptor, err := crypto.NewEncryptor("my-secret-key")
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt some data
	plaintext := `{"username":"admin","password":"secret123"}`
	encrypted, _ := encryptor.Encrypt(plaintext)

	// Decrypt it back
	decrypted, err := encryptor.Decrypt(encrypted)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decryption successful:", decrypted == plaintext)
	// Output: Decryption successful: true
}

func ExampleEncryptor_Decrypt_backwardCompatibility() {
	// Create an encryptor
	encryptor, err := crypto.NewEncryptor("my-secret-key")
	if err != nil {
		log.Fatal(err)
	}

	// Legacy unencrypted data (no "enc:" prefix)
	legacyData := `{"username":"admin","password":"old-data"}`

	// Decrypt handles unencrypted data gracefully
	decrypted, err := encryptor.Decrypt(legacyData)
	if err != nil {
		log.Fatal(err)
	}

	// Returns the data as-is for backward compatibility
	fmt.Println("Legacy data handled:", decrypted == legacyData)
	// Output: Legacy data handled: true
}

func ExampleIsEncrypted() {
	encryptedData := "enc:SGVsbG8gV29ybGQh"
	plaintextData := "Hello, World!"

	fmt.Println("Encrypted:", crypto.IsEncrypted(encryptedData))
	fmt.Println("Plaintext:", crypto.IsEncrypted(plaintextData))
	// Output:
	// Encrypted: true
	// Plaintext: false
}

func ExampleEncryptIfNeeded() {
	// Initialize global encryptor
	crypto.InitGlobalEncryptor("global-secret-key", true)

	// Use global encryptor for convenience
	encrypted, err := crypto.EncryptIfNeeded("sensitive data")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted:", crypto.IsEncrypted(encrypted))
	// Output: Encrypted: true
}

func ExampleDecryptIfNeeded() {
	// Initialize global encryptor
	crypto.InitGlobalEncryptor("global-secret-key", true)

	// Encrypt some data
	encrypted, _ := crypto.EncryptIfNeeded("sensitive data")

	// Decrypt using global encryptor
	decrypted, err := crypto.DecryptIfNeeded(encrypted)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decrypted:", decrypted == "sensitive data")
	// Output: Decrypted: true
}

// Example: Real-world usage with task configuration
func Example_taskConfiguration() {
	// Initialize encryption (typically done at app startup)
	crypto.InitGlobalEncryptor("production-secret-key-min-32-chars", true)

	// Simulate task configuration with sensitive data
	taskConfig := `{
		"host": "imap.gmail.com",
		"port": 993,
		"username": "user@example.com",
		"password": "super-secret-password",
		"tls": true
	}`

	// Before saving to database, encrypt the configuration
	encryptedConfig, err := crypto.EncryptIfNeeded(taskConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Store in database (simulated)
	fmt.Println("Stored in DB:", crypto.IsEncrypted(encryptedConfig))

	// When reading from database, decrypt the configuration
	decryptedConfig, err := crypto.DecryptIfNeeded(encryptedConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Use the decrypted configuration
	fmt.Println("Can use config:", decryptedConfig == taskConfig)

	// Output:
	// Stored in DB: true
	// Can use config: true
}
