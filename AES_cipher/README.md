# Encrypt Package

The `encrypt` package provides AES encryption and decryption utilities for securely encrypting strings using the AES algorithm. It uses CBC mode with PKCS7 padding and random IV generation for enhanced security.

## Features

- **AES Encryption**: Supports AES-128, AES-192, and AES-256
- **Secure Random IV**: Each encryption generates a unique IV
- **Base64 Encoding**: Encrypted output is base64-encoded for safe storage and transmission
- **PKCS7 Padding**: Proper padding handling for block cipher operations
- **No Dependencies**: Uses only Go's standard library

## Usage

### Generate a Key

```go
// Generate a 256-bit key (32 bytes) for AES-256
key, err := encrypt.GenerateKey(32)
if err != nil {
    log.Fatal(err)
}

// You can also use 16 bytes for AES-128 or 24 bytes for AES-192
key, err := encrypt.GenerateKey(16) // AES-128
key, err := encrypt.GenerateKey(24) // AES-192
```

### Create an Encrypter

```go
encrypter, err := encrypt.NewEncrypter(key)
if err != nil {
    log.Fatal(err)
}
```

### Encrypt a String

```go
plaintext := "This is sensitive information"
ciphertext, err := encrypter.Encrypt(plaintext)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Encrypted:", ciphertext)
```

### Decrypt a String

```go
decrypted, err := encrypter.Decrypt(ciphertext)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Decrypted:", decrypted) // Output: This is sensitive information
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/good-binary/utility/encrypt"
)

func main() {
    // Generate a secure key
    key, err := encrypt.GenerateKey(32)
    if err != nil {
        log.Fatal(err)
    }

    // Create an encrypter
    encrypter, err := encrypt.NewEncrypter(key)
    if err != nil {
        log.Fatal(err)
    }

    // Encrypt a message
    message := "Hello, World!"
    ciphertext, err := encrypter.Encrypt(message)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Encrypted:", ciphertext)

    // Decrypt the message
    decrypted, err := encrypter.Decrypt(ciphertext)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Decrypted:", decrypted)
}
```

## API Reference

### GenerateKey(size int) ([]byte, error)

Generates a random encryption key of the specified size.

- **Parameters:**
  - `size`: Key size in bytes (16, 24, or 32)
- **Returns:**
  - `[]byte`: The generated key
  - `error`: Error if size is invalid

### NewEncrypter(key []byte) (*Encrypter, error)

Creates a new Encrypter instance with the provided key.

- **Parameters:**
  - `key`: Must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256
- **Returns:**
  - `*Encrypter`: The encrypter instance
  - `error`: Error if key size is invalid

### Encrypter.Encrypt(plaintext string) (string, error)

Encrypts a plaintext string using AES-CBC with random IV.

- **Parameters:**
  - `plaintext`: The string to encrypt
- **Returns:**
  - `string`: Base64-encoded ciphertext with IV prepended
  - `error`: Error during encryption

### Encrypter.Decrypt(ciphertextB64 string) (string, error)

Decrypts a base64-encoded ciphertext string.

- **Parameters:**
  - `ciphertextB64`: Base64-encoded ciphertext (with IV prepended)
- **Returns:**
  - `string`: The decrypted plaintext
  - `error`: Error during decryption

## Security Notes

1. **Key Management**: Store encryption keys securely. Consider using environment variables or key management systems.
2. **Random IV**: Each encryption uses a unique random IV, providing semantic security. This means the same plaintext encrypted multiple times will produce different ciphertexts.
3. **Key Size**: Use AES-256 (32 bytes) for maximum security in most scenarios.
4. **No Authentication**: This package provides encryption but not authentication. For authenticated encryption, consider using additional authentication codes.

## Testing

Run the test suite with:

```bash
go test ./encrypt -v
```

The tests cover:
- Key generation with valid and invalid sizes
- Encrypter creation
- Encryption and decryption round-trips
- Unicode and special character handling
- Random IV generation
- Invalid base64 handling
- Tampered ciphertext detection
- Wrong key usage
- Padding and unpadding operations
