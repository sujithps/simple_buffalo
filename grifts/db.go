package grifts

import (
	"github.com/markbates/grift/grift"
	"github.com/simple_buffalo/models"
)

var _ = grift.Namespace("db", func() {
	grift.Desc("seed:wc", "Seeds a whitelisted_clients entry")
	grift.Add("seed:wc", func(c *grift.Context) error {
		models.DB.TruncateAll()

		wc := &models.WhitelistedClient{
			ClientID: "user001",
			PassKey:  "pass001",
		}

		err := models.DB.Create(wc)
		return err
	})

})
