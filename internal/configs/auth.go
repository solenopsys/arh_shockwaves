package configs

import (
	"sync"
)

type AuthManager struct {
	PublicKey string
	Nickname  string
}

var authInstance *AuthManager
var authOnce sync.Once

func GetAuthManager() *AuthManager {
	authOnce.Do(func() {
		authInstance = &AuthManager{
			PublicKey: "alexstorm",
			Nickname:  "solenopsys",
		}
	})
	return authInstance
}
