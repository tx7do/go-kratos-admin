package crypto

import (
	"strings"
	"testing"
)

func TestNewEncryptor(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		wantErr bool
	}{
		{
			name:    "valid key",
			key:     "my-secret-key",
			wantErr: false,
		},
		{
			name:    "empty key",
			key:     "",
			wantErr: true,
		},
		{
			name:    "long key",
			key:     "this-is-a-very-long-encryption-key-that-should-work-fine",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewEncryptor(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEncryptor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncryptor_EncryptDecrypt(t *testing.T) {
	encryptor, err := NewEncryptor("test-encryption-key-123")
	if err != nil {
		t.Fatalf("Failed to create encryptor: %v", err)
	}

	tests := []struct {
		name      string
		plaintext string
	}{
		{
			name:      "simple text",
			plaintext: "Hello, World!",
		},
		{
			name:      "JSON data",
			plaintext: `{"username":"admin","password":"secret123","host":"mail.example.com"}`,
		},
		{
			name:      "empty string",
			plaintext: "",
		},
		{
			name:      "special characters",
			plaintext: `!@#$%^&*()_+-=[]{}|;:'",.<>?/~`,
		},
		{
			name:      "unicode characters",
			plaintext: "こんにちは世界 🔐 مرحبا بالعالم",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encrypt
			encrypted, err := encryptor.Encrypt(tt.plaintext)
			if err != nil {
				t.Fatalf("Encrypt() error = %v", err)
			}

			// Check encrypted format
			if tt.plaintext != "" && !strings.HasPrefix(encrypted, EncryptedPrefix) {
				t.Errorf("Encrypted data should have prefix %q, got %q", EncryptedPrefix, encrypted)
			}

			// Verify encrypted data is different from plaintext
			if tt.plaintext != "" && encrypted == tt.plaintext {
				t.Errorf("Encrypted data should be different from plaintext")
			}

			// Decrypt
			decrypted, err := encryptor.Decrypt(encrypted)
			if err != nil {
				t.Fatalf("Decrypt() error = %v", err)
			}

			// Verify decrypted matches original
			if decrypted != tt.plaintext {
				t.Errorf("Decrypted data = %q, want %q", decrypted, tt.plaintext)
			}
		})
	}
}

func TestEncryptor_DecryptUnencrypted(t *testing.T) {
	encryptor, err := NewEncryptor("test-key")
	if err != nil {
		t.Fatalf("Failed to create encryptor: %v", err)
	}

	// Test backward compatibility - unencrypted data should pass through
	plaintext := `{"username":"admin","password":"secret"}`
	decrypted, err := encryptor.Decrypt(plaintext)
	if err != nil {
		t.Fatalf("Decrypt() should handle unencrypted data, error = %v", err)
	}

	if decrypted != plaintext {
		t.Errorf("Unencrypted data should pass through unchanged, got %q, want %q", decrypted, plaintext)
	}
}

func TestEncryptor_InvalidCiphertext(t *testing.T) {
	encryptor, err := NewEncryptor("test-key")
	if err != nil {
		t.Fatalf("Failed to create encryptor: %v", err)
	}

	tests := []struct {
		name       string
		ciphertext string
	}{
		{
			name:       "invalid base64",
			ciphertext: "enc:!!!invalid-base64!!!",
		},
		{
			name:       "too short",
			ciphertext: "enc:YWJj", // "abc" in base64, too short for valid ciphertext
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := encryptor.Decrypt(tt.ciphertext)
			if err == nil {
				t.Error("Decrypt() should return error for invalid ciphertext")
			}
		})
	}
}

func TestIsEncrypted(t *testing.T) {
	tests := []struct {
		name string
		data string
		want bool
	}{
		{
			name: "encrypted data",
			data: "enc:some-encrypted-data",
			want: true,
		},
		{
			name: "plaintext",
			data: "plain text data",
			want: false,
		},
		{
			name: "empty string",
			data: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEncrypted(tt.data); got != tt.want {
				t.Errorf("IsEncrypted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptor_DifferentKeys(t *testing.T) {
	// Test that different keys produce different results
	encryptor1, _ := NewEncryptor("key1")
	encryptor2, _ := NewEncryptor("key2")

	plaintext := "secret data"

	// Encrypt with first key
	encrypted1, err := encryptor1.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encrypt() error = %v", err)
	}

	// Try to decrypt with second key (should fail)
	_, err = encryptor2.Decrypt(encrypted1)
	if err == nil {
		t.Error("Decrypt() with different key should fail")
	}

	// Verify correct key works
	decrypted, err := encryptor1.Decrypt(encrypted1)
	if err != nil {
		t.Fatalf("Decrypt() with correct key error = %v", err)
	}
	if decrypted != plaintext {
		t.Errorf("Decrypted data = %q, want %q", decrypted, plaintext)
	}
}

func BenchmarkEncryptor_Encrypt(b *testing.B) {
	encryptor, _ := NewEncryptor("benchmark-key")
	plaintext := `{"username":"admin","password":"secret123","host":"mail.example.com","port":993}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encryptor.Encrypt(plaintext)
	}
}

func BenchmarkEncryptor_Decrypt(b *testing.B) {
	encryptor, _ := NewEncryptor("benchmark-key")
	plaintext := `{"username":"admin","password":"secret123","host":"mail.example.com","port":993}`
	encrypted, _ := encryptor.Encrypt(plaintext)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encryptor.Decrypt(encrypted)
	}
}
