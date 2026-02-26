package AES_cipher

import (
	"testing"
)

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name      string
		size      int
		wantError bool
	}{
		{"AES-128", 16, false},
		{"AES-192", 24, false},
		{"AES-256", 32, false},
		{"Invalid size", 10, true},
		{"Invalid size", 64, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GenerateKey(tt.size)
			if (err != nil) != tt.wantError {
				t.Errorf("GenerateKey() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if !tt.wantError && len(key) != tt.size {
				t.Errorf("GenerateKey() key size = %d, want %d", len(key), tt.size)
			}
		})
	}
}

func TestNewEncrypter(t *testing.T) {
	tests := []struct {
		name      string
		key       []byte
		wantError bool
	}{
		{"Valid 16 bytes", make([]byte, 16), false},
		{"Valid 24 bytes", make([]byte, 24), false},
		{"Valid 32 bytes", make([]byte, 32), false},
		{"Invalid 10 bytes", make([]byte, 10), true},
		{"Invalid 64 bytes", make([]byte, 64), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewEncrypter(tt.key)
			if (err != nil) != tt.wantError {
				t.Errorf("NewEncrypter() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestEncryptDecrypt(t *testing.T) {
	key, _ := GenerateKey(32)
	encrypter, _ := NewEncrypter(key)

	tests := []struct {
		name      string
		plaintext string
		wantError bool
	}{
		{"Simple text", "Hello, World!", false},
		{"Empty string", "", false},
		{"Long text", "The quick brown fox jumps over the lazy dog. " +
			"This is a longer string to test encryption with multiple blocks.", false},
		{"Special characters", "!@#$%^&*()_+-=[]{}|;':\",./<>?", false},
		{"Unicode", "你好世界🌍", false},
		{"Numbers", "1234567890", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ciphertext, err := encrypter.Encrypt(tt.plaintext)
			if (err != nil) != tt.wantError {
				t.Errorf("Encrypt() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError {
				plaintext, err := encrypter.Decrypt(ciphertext)
				if err != nil {
					t.Errorf("Decrypt() error = %v", err)
					return
				}

				if plaintext != tt.plaintext {
					t.Errorf("Decrypt() plaintext = %q, want %q", plaintext, tt.plaintext)
				}
			}
		})
	}
}

func TestEncryptDifferentResults(t *testing.T) {
	// Ensure that encrypting the same plaintext produces different ciphertexts
	// (due to random IV)
	key, _ := GenerateKey(32)
	encrypter, _ := NewEncrypter(key)

	plaintext := "Test encryption randomness"
	ciphertext1, _ := encrypter.Encrypt(plaintext)
	ciphertext2, _ := encrypter.Encrypt(plaintext)

	if ciphertext1 == ciphertext2 {
		t.Error("Different encryptions of same plaintext should produce different ciphertexts")
	}
}

func TestDecryptInvalidBase64(t *testing.T) {
	key, _ := GenerateKey(32)
	encrypter, _ := NewEncrypter(key)

	_, err := encrypter.Decrypt("not valid base64!!!")
	if err == nil {
		t.Error("Decrypt() should fail with invalid base64")
	}
}

func TestDecryptModifiedCiphertext(t *testing.T) {
	key, _ := GenerateKey(32)
	encrypter, _ := NewEncrypter(key)

	plaintext := "Original message"
	ciphertext, _ := encrypter.Encrypt(plaintext)

	// Modify the ciphertext slightly
	ciphertextBytes := []byte(ciphertext)
	if len(ciphertextBytes) > 4 {
		ciphertextBytes[4] = ciphertextBytes[4] ^ 0xFF // Flip bits
	}
	modifiedCiphertext := string(ciphertextBytes)

	decrypted, err := encrypter.Decrypt(modifiedCiphertext)
	if err == nil && decrypted == plaintext {
		t.Error("Decrypt() should not successfully decrypt tampered ciphertext")
	}
}

func TestEncryptWithDifferentKeys(t *testing.T) {
	key1, _ := GenerateKey(32)
	key2, _ := GenerateKey(32)

	encrypter1, _ := NewEncrypter(key1)
	encrypter2, _ := NewEncrypter(key2)

	plaintext := "Sensitive data"
	ciphertext, _ := encrypter1.Encrypt(plaintext)

	// Try to decrypt with wrong key
	_, err := encrypter2.Decrypt(ciphertext)
	if err == nil {
		t.Error("Decrypt() should fail when using wrong key")
	}
}

func TestPadUnpad(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		blockSize int
	}{
		{"Single byte", []byte{1}, 16},
		{"Partial block", []byte{1, 2, 3}, 16},
		{"Full block", make([]byte, 16), 16},
		{"Multiple blocks", make([]byte, 32), 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			padded := pad(tt.data, tt.blockSize)
			if len(padded)%tt.blockSize != 0 {
				t.Errorf("Padded data is not a multiple of block size")
			}

			unpadded := unpad(padded, tt.blockSize)
			if string(unpadded) != string(tt.data) {
				t.Errorf("Unpadded data doesn't match original")
			}
		})
	}
}
