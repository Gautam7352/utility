package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Encrypter handles AES encryption and decryption operations
type Encrypter struct {
	key []byte
}

// NewEncrypter creates a new Encrypter with the provided key
// Key must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256 respectively
func NewEncrypter(key []byte) (*Encrypter, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("invalid key size: must be 16, 24, or 32 bytes, got %d", len(key))
	}
	return &Encrypter{
		key: key,
	}, nil
}

// GenerateKey creates a random encryption key of the specified size
// size should be 16, 24, or 32 for AES-128, AES-192, or AES-256
func GenerateKey(size int) ([]byte, error) {
	if size != 16 && size != 24 && size != 32 {
		return nil, fmt.Errorf("invalid key size: must be 16, 24, or 32 bytes, got %d", size)
	}
	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}
	return key, nil
}

// Encrypt encrypts a plaintext string using AES-CBC and returns a base64-encoded string
func (e *Encrypter) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	plainBytes := []byte(plaintext)
	plainBytes = pad(plainBytes, block.BlockSize())

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("failed to generate IV: %w", err)
	}

	ciphertext := make([]byte, len(plainBytes))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plainBytes)

	// Prepend IV to ciphertext
	result := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(result), nil
}

// Decrypt decrypts a base64-encoded ciphertext string using AES-CBC
func (e *Encrypter) Decrypt(ciphertextB64 string) (string, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	if len(ciphertextBytes) < block.BlockSize() {
		return "", fmt.Errorf("ciphertext too short")
	}

	// Extract IV from the beginning
	iv := ciphertextBytes[:block.BlockSize()]
	ciphertext := ciphertextBytes[block.BlockSize():]

	if len(ciphertext)%block.BlockSize() != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of block size")
	}

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext = unpad(plaintext, block.BlockSize())
	return string(plaintext), nil
}

// pad adds PKCS7 padding to the plaintext
func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padChar := byte(padding)
	for i := 0; i < padding; i++ {
		data = append(data, padChar)
	}
	return data
}

// unpad removes PKCS7 padding from the plaintext
func unpad(data []byte, blockSize int) []byte {
	if len(data) == 0 {
		return data
	}
	padding := int(data[len(data)-1])
	if padding > blockSize || padding == 0 {
		return data
	}
	return data[:len(data)-padding]
}
