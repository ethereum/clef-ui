package ui

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type LoginContext struct {
	core.QObject

	_ string `property:"clefPath"`
	_ string `property:"binaryHash"`
	_ string `property:"error"`

	_ func()         `signal:"start,auto"`
	_ func(b string) `signal:"checkPath,auto"`

	answer       int
	ClefUI       *ClefUI
	ReadyToStart chan string
}

func (t *LoginContext) start() {
	if t.Error() != "" {
		return
	}

	t.ReadyToStart <- t.ClefPath()
}

func (t *LoginContext) checkPath(pathToClef string) {
	// todo: remember where the binary was 'last' time, and reuse that location
	log.Println(pathToClef)
	t.SetClefPath(pathToClef)

	f, err := os.Open(pathToClef)
	if err != nil {
		t.SetError(err.Error())
		log.Println(err)
		return
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		t.SetError(err.Error())
		log.Println(err)
		return
	}

	binaryHash := fmt.Sprintf("%x", h.Sum(nil))
	t.SetBinaryHash(binaryHash)
	t.SetError("")
}

func NewLogin(rootContext *qml.QQmlContext, readytoStart chan string) *LoginContext {
	c := NewLoginContext(nil)
	c.SetError("")
	c.ConnectClefPathChanged(c.checkPath)
	c.ReadyToStart = readytoStart

	rootContext.SetContextProperty("LoginContext", c)

	return c
}
