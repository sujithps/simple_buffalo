package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/simple_buffalo/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
