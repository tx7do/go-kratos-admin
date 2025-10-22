package crypto

import (
	"sync"
)

var (
	globalEncryptor *Encryptor
	once            sync.Once
)

// InitGlobalEncryptor initializes the global encryptor instance
func InitGlobalEncryptor(key string, enabled bool) error {
	var err error
	once.Do(func() {
		if !enabled {
			// Create a no-op encryptor if encryption is disabled
			globalEncryptor = &Encryptor{}
			return
		}

		globalEncryptor, err = NewEncryptor(key)
	})
	return err
}

// GetGlobalEncryptor returns the global encryptor instance
func GetGlobalEncryptor() *Encryptor {
	if globalEncryptor == nil {
		// Return a no-op encryptor if not initialized
		return &Encryptor{}
	}
	return globalEncryptor
}

// EncryptIfNeeded encrypts data if the global encryptor is initialized
func EncryptIfNeeded(plaintext string) (string, error) {
	encryptor := GetGlobalEncryptor()
	if encryptor == nil || len(encryptor.key) == 0 {
		// Encryption disabled, return plaintext
		return plaintext, nil
	}
	return encryptor.Encrypt(plaintext)
}

// DecryptIfNeeded decrypts data if it's encrypted
func DecryptIfNeeded(ciphertext string) (string, error) {
	encryptor := GetGlobalEncryptor()
	if encryptor == nil || len(encryptor.key) == 0 {
		// Encryption disabled, return as-is
		return ciphertext, nil
	}
	return encryptor.Decrypt(ciphertext)
}
