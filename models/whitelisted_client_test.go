package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_WhitelistedClient(t *testing.T) {
	clientID := "client-id"
	passKey := "pass-key"
	wc := WhitelistedClient{ClientID: clientID,
		PassKey: passKey,}

	assert.Equal(t, clientID, wc.ClientID)
	assert.Equal(t, passKey, wc.PassKey)
}
