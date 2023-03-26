package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/pbkdf2"
	"time"
)

func GenJwt(user string, access string, secret []byte) string {
	claims := jwt.MapClaims{
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"issuer":   "urn:solenopsys:issuer",
		"audience": "urn:solenopsys:audience",
		"user":     user,
		"access":   access,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secret)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return ""
	}

	return signedToken
}

func DecryptKeyData(encryptedText string, password string) ([]byte, error) {
	// algorithm := "AES-256-CBC"
	key := make([]byte, 32)
	iv := make([]byte, 16)
	salt := make([]byte, 8)

	// Generate key and IV from password and salt using PBKDF2
	derivedKey := pbkdf2.Key([]byte(password), salt, 100000, len(key)+len(iv), sha256.New)
	copy(key, derivedKey[:len(key)])
	copy(iv, derivedKey[len(key):])

	// Decode encrypted data from hex string
	encryptedData, err := hex.DecodeString(encryptedText)
	if err != nil {
		return nil, err
	}

	// Decrypt data using AES-CBC mode
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptedData, encryptedData)

	// Remove padding
	paddingLen := int(encryptedData[len(encryptedData)-1])
	return encryptedData[:len(encryptedData)-paddingLen], nil
}
