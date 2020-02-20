package ui

import (
	// "crypto/sha256"
	// "fmt"
	// "io"
	// "log"
	// "os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

type LoginContext struct {
	core.QObject

	_ string `property:"clefPath"`
	_ string `property:"binaryHash"`
	_ string `property:"error"`

	_ func()         `slot:"validate"`
	_ func(b string) `slot:"checkPath"`
}

func (t *LoginContext) checkPath(pathToClef string) {
	// TODO!
	// Remember where the binary was 'last' time, and reuse that location
}

func NewLogin(rootWidget *quick.QQuickWidget, name string) *LoginContext {
	if name == "" {
		name = "LoginContext"
	}

	c := NewLoginContext(nil)
	c.ConnectClefPathChanged(c.checkPath)

	rootWidget.RootContext().SetContextProperty(name, c)

	return c
}
