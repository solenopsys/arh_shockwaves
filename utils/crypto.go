package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/cosmos/go-bip39"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"log"
	"math/big"
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

// gen mnemonic
func GenMnemonic() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	return mnemonic
}

func PrivateKeyFromSeed(seedPhrase string) *secp256k1.PrivKeySecp256k1 {
	hash := sha256.Sum256([]byte(seedPhrase))
	privateKey := make([]byte, 32)
	copy(privateKey, hash[:])

	println("HASH", hex.EncodeToString(privateKey))

	k1 := secp256k1.PrivKeySecp256k1([32]byte{})

	r := bytes.NewReader(privateKey)
	_, err := io.ReadFull(r, k1[:])
	if err != nil {
		log.Fatal(err)
	}

	hexPrivate := hex.EncodeToString(k1[:])
	println("Private key:", hexPrivate)
	return &k1
}

func LoadPrivateKeyFromString(keyStr string) (*secp256k1.PrivKeySecp256k1, error) {
	keyBytes, err := hex.DecodeString(keyStr)
	if err != nil {
		return nil, err
	}

	if len(keyBytes) != 32 {
		return nil, errors.New("invalid key size")
	}

	var key [32]byte
	copy(key[:], keyBytes)

	k1 := secp256k1.PrivKeySecp256k1(key)
	return &k1, nil
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
