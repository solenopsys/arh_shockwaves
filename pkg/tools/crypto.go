package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/pbkdf2"
	"math/big"
	"xs/pkg/io"
)

// decrypt data for decode private key
func DecryptKeyData(encryptedText string, password string) ([]byte, error) {
	// algorithm := "AES-256-CBC"
	key := make([]byte, 32)
	iv := make([]byte, 16)
	salt := make([]byte, 0)

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

func LoadPrivateKeyFromString(keyStr string) (**ecdsa.PrivateKey, error) {

	importedPrivKey, err := crypto.HexToECDSA(keyStr)
	if err != nil {
		io.Fatal(err)
	}
	return &importedPrivKey, nil
}

func LoadPublicKeyFromString(keyStr string) (*ecdsa.PublicKey, error) {
	keyBytes, err := hex.DecodeString(keyStr)
	if err != nil {
		return nil, err
	}

	curve := elliptic.P256() // secp256k1 curve
	if len(keyBytes) != 2*curve.Params().BitSize/8+1 {
		return nil, errors.New("invalid key size")
	}

	x := new(big.Int).SetBytes(keyBytes[1 : curve.Params().BitSize/8+1])
	y := new(big.Int).SetBytes(keyBytes[curve.Params().BitSize/8+1:])

	pubKey := new(ecdsa.PublicKey)
	pubKey.Curve = curve
	pubKey.X = x
	pubKey.Y = y

	return pubKey, nil
}
