package tools

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"xs/pkg/io"
)

func CreateToken(payload map[string]string, privateKey *ecdsa.PrivateKey) string {
	body, err := json.Marshal(payload)
	if err != nil {
		io.Panic(err)
	}

	hash := sha256.Sum256(body)
	signature, err := privateKey.Sign(rand.Reader, hash[:], crypto.SHA256)
	if err != nil {
		io.Panic(err)
	}

	token := make([]byte, hex.EncodedLen(len(signature))+hex.EncodedLen(len(body)))
	hex.Encode(token[:hex.EncodedLen(len(signature))], signature)
	hex.Encode(token[hex.EncodedLen(len(signature)):], body)

	return string(token)
}

func ReadToken(token string, publicKey *ecdsa.PublicKey) (bool, map[string]string) {
	fromHexSignature, _ := hex.DecodeString(token[:128])
	fromHexBody, _ := hex.DecodeString(token[128:])

	hash := sha256.Sum256(fromHexBody)

	verified := ecdsa.Verify(publicKey, hash[:],
		new(big.Int).SetBytes(fromHexSignature[:len(fromHexSignature)/2]),
		new(big.Int).SetBytes(fromHexSignature[len(fromHexSignature)/2:]))

	if verified {
		var payload map[string]string
		err := json.Unmarshal(fromHexBody, &payload)
		if err != nil {
			io.Panic(err)
		}
		return true, payload
	} else {
		return false, nil
	}
}
