package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"time"
)

type WhitelistedClient struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ClientID  string    `json:"client_id" db:"client_id"`
	PassKey   string    `json:"pass_key" db:"pass_key"`
}

// String is not required by pop and may be deleted
func (w WhitelistedClient) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// WhitelistedClients is not required by pop and may be deleted
type WhitelistedClients []WhitelistedClient

// String is not required by pop and may be deleted
func (w WhitelistedClients) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *WhitelistedClient) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: w.ClientID, Name: "ClientID"},
		&validators.StringIsPresent{Field: w.PassKey, Name: "PassKey"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *WhitelistedClient) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *WhitelistedClient) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
