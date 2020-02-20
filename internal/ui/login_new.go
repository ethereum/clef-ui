package ui

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ethereum/clef-ui/external"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type LoginContext struct {
	core.QObject

	server external.Server

	_ string `property:"clefPath"`
	_ string `property:"binaryHash"`
	_ string `property:"error"`

	_ func()         `signal:"start,auto"`
	_ func(b string) `signal:"checkPath,auto"`

	// using channels as signals, since it seems like it's not possible to emit signals from golang
	InitiateServerRequest chan string
}

func (t *LoginContext) setServer(server external.Server) {
	t.server = server
}

func (t *LoginContext) start() {
	t.InitiateServerRequest <- t.ClefPath()
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

func NewLogin(rootContext *qml.QQmlContext, name string) *LoginContext {
	if name == "" {
		name = "LoginContext"
	}

	c := NewLoginContext(nil)
	c.SetError("")
	c.InitiateServerRequest = make(chan string)
	c.ConnectClefPathChanged(c.checkPath)

	rootContext.SetContextProperty(name, c)

	return c
}
