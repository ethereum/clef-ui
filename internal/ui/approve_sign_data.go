package ui

import (
	core2 "github.com/ethereum/go-ethereum/signer/core"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

type ApproveSignDataUI struct {
	UI            *quick.QQuickWidget
	ContextObject *ApproveSignDataCtx
}

type ApproveSignDataCtx struct {
	core.QObject

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"from"`
	_ string `property:"message"`
	_ string `property:"rawData"`
	_ string `property:"hash"`
	_ string `property:"password"`
	_ string `property:"fromSrc"`

	_ func(b int)                  `signal:"clicked,auto"`
	_ func()                       `signal:"back,auto"`
	_ func(b string, value string) `signal:"edited,auto"`

	answerCh chan int
	ClefUI *ClefUI
}

func (t *ApproveSignDataCtx) clicked(b int) {
	select {
	case t.answerCh <- -b:
	default:
	}
}

func (t *ApproveSignDataCtx) back() {
	select {
	case t.answerCh <- -1:
	default:
	}
}

func (t *ApproveSignDataCtx) edited(name string, value string) {
	switch name {
	case "password":
		t.SetPassword(value)
	}
}

func (t *ApproveSignDataCtx) Reset() {
	t.SetRemote("")
	t.SetTransport("")
	t.SetEndpoint("")
	t.SetFrom("")
	t.SetMessage("")
	t.SetRawData("")
	t.SetHash("")
	t.SetPassword("")
	t.ClefUI.BackToMain <- true
}

func (t *ApproveSignDataCtx) ClickResponse(res chan *core2.SignDataResponse) {
	go func() {
		answer := <-t.answerCh
		if answer == -1 { // Go back
			t.Reset()
			return
		}
		var reply = new(core2.SignDataResponse)
		if answer == 1 {
			reply.Approved = false
		} else if answer == 2 {
			reply.Approved = true
			reply.Password = t.Password()
		}
		res <- reply
		t.Reset()

	}()
}

func NewApproveSignDataUI(clefUi *ClefUI) *ApproveSignDataUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_sign_data.qml", 0))
	c := NewApproveSignDataCtx(nil)
	c.ClefUI = clefUi
	v := &ApproveSignDataUI{
		UI:            widget,
		ContextObject: c,
	}
	c.answerCh = make(chan int)
	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}
