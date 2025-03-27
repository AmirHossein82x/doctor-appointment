package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/AmirHossein82x/doctor-appointment/internal/config"
	"github.com/google/uuid"
)

// Generate a 12-byte random nonce
func generateNonce() ([]byte, error) {
	nonce := make([]byte, 12) // AES-GCM requires a 12-byte nonce
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

// Decode Base64 encryption key
func decodeEncryptionKey(encodedKey string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return nil, err
	}
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("invalid AES key size, must be 16, 24, or 32 bytes")
	}
	return key, nil
}

// Encrypt UUID using AES-GCM
func EncryptUUID(userID uuid.UUID) (string, error) {
	cfg := config.LoadConfig()
	key, err := decodeEncryptionKey(cfg.ENCRYPTION_KEY)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce, err := generateNonce()
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(userID.String()), nil)
	finalData := append(nonce, ciphertext...) // Store nonce + encrypted data together

	return base64.URLEncoding.EncodeToString(finalData), nil
}

// Decrypt UUID using AES-GCM
func DecryptUUID(encryptedText string) (string, error) {
	cfg := config.LoadConfig()
	key, err := decodeEncryptionKey(cfg.ENCRYPTION_KEY)
	if err != nil {
		return "", err
	}

	data, err := base64.URLEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := data[:12]         // Extract nonce
	encryptedData := data[12:] // Extract encrypted UUID
	plaintext, err := aesGCM.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
