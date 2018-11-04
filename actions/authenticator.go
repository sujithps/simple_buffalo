package actions

import (
	"errors"
	"fmt"

	"github.com/patrickmn/go-cache"
)

type ClientAuthentication struct {
	cache *cache.Cache
}

func NewClientAuthentication() *ClientAuthentication {
	return &ClientAuthentication{
		cache: cache.New(cache.NoExpiration, 0),
	}
}

func (ca *ClientAuthentication) Authenticate(clientID, passKey string) error {
	if len(clientID) <= 0 || len(passKey) <= 0 {
		return fmt.Errorf("either of Client-ID or Pass-Key is missing, Client-Id: %s", clientID)
	}

	cachedPassKey, found := ca.cache.Get(clientID)
	if !found {
		authorizedApplication, err := ca.getData(clientID)
		if err != nil {
			return fmt.Errorf("failed to query the database: %s", err)
		}

		ca.cache.Set(authorizedApplication.ClientID, authorizedApplication.PassKey, cache.NoExpiration)
		cachedPassKey = authorizedApplication.PassKey
	}

	if cachedPassKey != passKey {
		return errors.New("PassKey is not valid")
	}

	return nil
}

func (ca *ClientAuthentication) getData(_ string) (AuthData, error) {
	return AuthData{
		ClientID: "A",
		PassKey:  "B",
	}, nil
}

type AuthData struct {
	ClientID string
	PassKey  string
}
