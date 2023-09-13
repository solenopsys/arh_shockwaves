package funcs

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"xs/pkg/io"
)

type RegisterData struct {
	Transport    string `json:"transport"`
	Login        string `json:"login"`
	EncryptedKey string `json:"encryptedKey"`
	PublicKey    string `json:"publicKey"`
	Hash         string `json:"hash"`
}

func GenHash(password string, login string) string {
	hash := sha256.Sum256([]byte(password + login))
	return hex.EncodeToString(hash[:])
}

func postRequest(url string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//

func LoadKey(password string, login string) []byte {

	hash := GenHash(password, login)
	response, err := postRequest("http://127.0.0.1:8899/api/key", []byte(hash))

	defer response.Body.Close()
	// response to RegisterData

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		io.Panic(err)
	}

	return bodyBytes

}

func UnMarshal(bytes []byte) *RegisterData {
	var register RegisterData = RegisterData{}
	err := json.Unmarshal(bytes, &register)

	if err != nil {
		io.Println(err)
	}

	return &register
}
