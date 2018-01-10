package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/golang_buffalo/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
