package ui

import (
	"crypto/sha256"
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"io"
	"log"
	"os"
	"path"
)

type LoginUI struct {
	UI 					*quick.QQuickWidget
	ContextObject		*LoginCtx
}

type LoginCtx struct {
	core.QObject

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"gopath"`
	_ string `property:"binaryHash"`
	_ bool 	 `property:"isValid"`

	_ func() 			`signal:"clicked,auto"`
	_ func(b string) 	`signal:"edited,auto"`

	answer 				int
	ClefUI 				*ClefUI
	ReadyToStart 		chan string
}

func (t *LoginCtx) clicked() {
	if !t.IsValid() {
		return
	}
	gopath := t.Gopath()
	log.Println(gopath)
	t.ReadyToStart <- gopath
}


func (t *LoginCtx) edited(dir string) {
	t.SetGopath(dir)

	f, err := os.Open(path.Join(dir, "bin", "clef"))
	if err != nil {
		t.SetBinaryHash(err.Error())
		t.SetIsValid(false)
		log.Println(err)
		return
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		t.SetBinaryHash(err.Error())
		t.SetIsValid(false)
		log.Println(err)
		return
	}

	binaryHash := fmt.Sprintf("%x", h.Sum(nil))
	t.SetBinaryHash(binaryHash)
	t.SetIsValid(len(binaryHash) > 0)
}

func (t *LoginCtx) Reset() {
	t.answer = 0
	t.SetRemote("")
	t.SetTransport("")
	t.SetEndpoint("")
	t.SetGopath("")
	t.SetBinaryHash("")
	t.SetIsValid(false)
	t.ClefUI.BackToMain <- true
}

func NewLoginUI(clefUi *ClefUI, readyToStart chan string) *LoginUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/login.qml", 0))
	c := NewLoginCtx(nil)
	c.ReadyToStart = readyToStart
	c.ClefUI = clefUi
	v := &LoginUI{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Show()
	return v
}
