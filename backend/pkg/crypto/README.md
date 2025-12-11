# Task Configuration Encryption

This package provides AES-256-GCM encryption/decryption for sensitive task configuration data stored in the database.

## Features

- **AES-256-GCM Encryption**: Industry-standard authenticated encryption
- **Automatic Encryption/Decryption**: Transparent encryption via Ent hooks
- **Backward Compatible**: Gracefully handles unencrypted legacy data
- **Configurable**: Enable/disable via configuration
- **Migration Tool**: Encrypt existing data with zero downtime

## Security

### Encryption Method
- **Algorithm**: AES-256-GCM (Galois/Counter Mode)
- **Key Derivation**: SHA-256 hash of the configured key
- **Authentication**: Built-in message authentication via GCM
- **Nonce**: Random 12-byte nonce for each encryption operation

### Key Management
⚠️ **IMPORTANT**: The encryption key should be:
1. At least 32 characters long
2. Stored as an environment variable in production
3. Never committed to version control
4. Rotated periodically
5. Backed up securely

## Configuration

### Environment Variable (Recommended)
```bash
export ENCRYPTION_KEY="your-secure-random-key-minimum-32-characters-long"
```

### Configuration File
Add to `configs/ai.yaml`:
```yaml
encryption:
  key: "${ENCRYPTION_KEY:change-this-in-production}"
  enabled: true
```

## Usage

### Automatic Encryption/Decryption

The system automatically encrypts/decrypts task payloads:

```go
// Creating a task - payload is automatically encrypted before saving
task := &adminV1.CreateTaskRequest{
    Data: &adminV1.Task{
        TypeName: "email_processor",
        TaskPayload: `{"username":"admin","password":"secret123"}`,
    },
}
result, err := taskService.Create(ctx, task)

// Reading a task - payload is automatically decrypted
task, err := taskService.Get(ctx, &adminV1.GetTaskRequest{Id: 1})
// task.TaskPayload is decrypted and ready to use
```

### Manual Encryption/Decryption

```go
import "go-wind-admin/pkg/crypto"

// Initialize encryptor
encryptor, err := crypto.NewEncryptor("your-secret-key")
if err != nil {
    return err
}

// Encrypt data
encrypted, err := encryptor.Encrypt("sensitive data")
// Result: "enc:base64-encoded-ciphertext"

// Decrypt data
decrypted, err := encryptor.Decrypt(encrypted)
// Result: "sensitive data"

// Check if data is encrypted
if crypto.IsEncrypted(data) {
    // Data is encrypted
}
```

## Migration

### Encrypt Existing Data

To encrypt existing unencrypted task configurations:

```bash
# Dry run (preview changes without applying)
cd backend/app/admin/service/cmd/migrate-encrypt
go run main.go --conf ../../configs --dry-run

# Apply encryption
go run main.go --conf ../../configs
```

The migration tool:
- ✓ Checks each task's payload
- ✓ Skips already encrypted data
- ✓ Skips empty payloads
- ✓ Reports progress and errors
- ✓ Provides detailed summary

### Zero-Downtime Migration

The encryption system supports zero-downtime migration:

1. **Phase 1**: Enable encryption and deploy
   - New tasks are encrypted
   - Old tasks work (backward compatible)

2. **Phase 2**: Run migration tool
   - Encrypts existing data
   - Can be run during normal operation

3. **Phase 3**: Verify
   - All tasks are encrypted
   - System fully operational

## Encrypted Data Format

Encrypted data uses this format:
```
enc:<base64-encoded-ciphertext>
```

Where ciphertext contains:
```
[12-byte nonce][encrypted data][16-byte auth tag]
```

The `enc:` prefix allows the system to:
- Identify encrypted vs plaintext data
- Support backward compatibility
- Safely migrate existing data

## Testing

Run the test suite:
```bash
cd backend/pkg/crypto
go test -v
```

Run benchmarks:
```bash
go test -bench=. -benchmem
```

## Error Handling

### Encryption Disabled
If encryption is disabled, data is stored as plaintext:
```yaml
encryption:
  enabled: false
```

### Missing Encryption Key
If no key is configured, the system logs a warning and operates in plaintext mode.

### Decryption Failures
Decryption failures are logged but don't block operations - the system returns the encrypted data as-is for debugging.

## Best Practices

### Production Deployment

1. **Generate a Strong Key**
   ```bash
   # Generate a random 32-byte key
   openssl rand -base64 32
   ```

2. **Set Environment Variable**
   ```bash
   export ENCRYPTION_KEY="generated-key-from-step-1"
   ```

3. **Verify Configuration**
   ```bash
   # Check that encryption is enabled
   grep -A2 "encryption:" configs/ai.yaml
   ```

4. **Deploy Application**
   - New tasks will be automatically encrypted
   - Old tasks remain readable

5. **Migrate Existing Data**
   ```bash
   # Run migration tool
   ./migrate-encrypt --conf configs
   ```

### Key Rotation

To rotate encryption keys:

1. Deploy new application version with old key
2. Use migration tool to re-encrypt with new key
3. Update configuration with new key
4. Restart application

### Monitoring

Monitor these metrics:
- Encryption/decryption error rates
- Migration progress
- Database query performance

## Performance

Typical performance on modern hardware:
- Encryption: ~500,000 ops/sec
- Decryption: ~500,000 ops/sec
- Database overhead: <1ms per operation

## Security Considerations

### What is Encrypted
- Task payload (task_payload field)
- Any data containing credentials, passwords, API keys

### What is NOT Encrypted
- Task metadata (type, status, timestamps)
- Task IDs and names
- Audit logs

### Threat Model
Protects against:
- ✓ Database dumps
- ✓ SQL injection (payload exfiltration)
- ✓ Backup compromise
- ✓ Insider threats (DBA access)

Does NOT protect against:
- ✗ Application-level compromise
- ✗ Key compromise
- ✗ Memory dumps of running process
- ✗ Side-channel attacks

## Troubleshooting

### "Invalid ciphertext" error
- Encryption key may have changed
- Data may be corrupted
- Check application logs for details

### Tasks not decrypting
- Verify encryption is enabled in config
- Check encryption key is set correctly
- Ensure InitEncryption() is called on startup

### Migration fails
- Check database connectivity
- Verify encryption key is correct
- Review migration tool output for specific errors

## Examples

### Example 1: Email Task Configuration
```json
{
  "host": "imap.gmail.com",
  "port": 993,
  "username": "user@example.com",
  "password": "secret-password",
  "tls": true
}
```

When stored in database:
```
enc:SGVsbG8gV29ybGQhCg==...base64-encoded-encrypted-data
```

### Example 2: API Integration
```json
{
  "api_url": "https://api.example.com",
  "api_key": "sk-1234567890abcdef",
  "api_secret": "super-secret-key"
}
```

After encryption, credentials are protected in the database.

## License

This encryption implementation is part of the go-wind-admin project.
