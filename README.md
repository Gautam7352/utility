# Utility

A comprehensive Go utility library providing reusable packages for common development tasks. Built with Go 1.24.5, zero external dependencies, and following Go best practices.

## Packages

### Logger
Structured logging with flexible configuration and multiple output destinations.
- Multiple log levels: Debug, Info, Warning, Error
- Template-based message formatting
- Dual output support (file + stdout)
- Production mode flag
- JSON output mode

### Random
Generate random data for testing and seeding purposes.
- `RandomNumber()` — Random integers in a range
- `RandomString()` — Random alphanumeric strings with prefix/suffix
- `RandomFullName()` — Combine random first and last names

### Utils
Type-safe slice/array operations using Go generics (1.18+).
- Generic `Slicer[T]` wrapper for slices
- Methods: `Get()`, `Set()`, `Append()`, `Prepend()`, `Remove()`, `Clear()`, `Len()`
- Object-oriented interface to slice manipulation

### UUID
RFC 4122 compliant UUID generation and manipulation using cryptographic randomness.
- `NewUUID()` — Generate random v4 UUIDs
- `Parse()` — Parse canonical UUID strings
- `String()` — Format UUID as canonical string
- `Validate()` — Validate UUID format
- JSON marshaling support
- Standard library only, no external dependencies

### Encrypt
AES encryption and decryption utilities for secure string encryption.
- AES-128, AES-192, and AES-256 support
- Secure random IV generation
- CBC mode with PKCS7 padding
- Base64 encoding for safe storage/transmission
- `GenerateKey()` — Create cryptographically secure keys
- `Encrypter.Encrypt()` — Encrypt strings
- `Encrypter.Decrypt()` — Decrypt ciphertexts

## Installation

```bash
go get github.com/good-binary/utility
```

## Quick Start

```go
package main

import (
    "log"
    "github.com/good-binary/utility/encrypt"
    "github.com/good-binary/utility/uuid"
    "github.com/good-binary/utility/logger"
)

func main() {
    // Generate a UUID
    id := uuid.NewUUID()
    
    // Create a logger
    appLogger, _ := logger.NewLogger()
    appLogger.Info("Application started")
    
    // Encrypt data
    key, _ := encrypt.GenerateKey(32)
    enc, _ := encrypt.NewEncrypter(key)
    ciphertext, _ := enc.Encrypt("Sensitive data")
    appLogger.Infof("Encrypted: %s", ciphertext)
}
```

## Features

- **Modular Design** — Each package has single responsibility
- **Zero Dependencies** — Pure Go standard library
- **Well-Tested** — Comprehensive unit tests for all packages
- **Go Generics** — Modern Go 1.18+ generic constraints
- **Interface Implementation** — JSON marshaling/unmarshaling support
- **Security-Focused** — Cryptographic randomness, proper padding, secure keys
- **MIT Licensed** — Free to use in commercial and open-source projects

## License

MIT License © 2024 best-binary
