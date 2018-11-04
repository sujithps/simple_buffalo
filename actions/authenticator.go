package actions

import (
	"errors"
	"fmt"

	"github.com/patrickmn/go-cache"
	"github.com/simple_buffalo/models"
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
		whitelistedClient, err := ca.getDBData(clientID)
		if err != nil {
			return fmt.Errorf("failed to query the database: %s", err)
		}

		ca.cache.Set(whitelistedClient.ClientID, whitelistedClient.PassKey, cache.NoExpiration)
		cachedPassKey = whitelistedClient.PassKey
	}

	if cachedPassKey != passKey {
		return errors.New("PassKey is not valid")
	}

	return nil
}

func (ca *ClientAuthentication) getDBData(clientID string) (models.WhitelistedClient, error) {

	whitelistedClients := []models.WhitelistedClient{}
	query := models.DB.Where("client_id = ?", clientID)
	err := query.All(&whitelistedClients)

	if len(whitelistedClients) == 0 {
		return models.WhitelistedClient{}, errors.New("could not a matching clientId")
	}

	return whitelistedClients[0], err
}
