package ui

import (
	"fmt"

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
	_ func()                       `signal:"onBack,auto"`
	_ func()                       `signal:"onApprove,auto"`
	_ func()                       `signal:"onReject,auto"`
	_ func(b string, value string) `signal:"edited,auto"`

	answerCh chan int
	ClefUI   *ClefUI
}

func (t *ApproveSignDataCtx) clicked(b int) {
	select {
	case t.answerCh <- -b:
	default:
	}
}
func (t *ApproveSignDataCtx) onApprove() {
	fmt.Printf("onApprove called")
}
func (t *ApproveSignDataCtx) onReject() {
	fmt.Printf("onReject called")
}

func (t *ApproveSignDataCtx) onBack() {
	fmt.Printf("Back() called")
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
		}
		res <- reply
		t.Reset()

	}()
}

func NewApproveSignDataUI(clefUi *ClefUI) *ApproveSignDataUI {
	c := NewApproveSignDataCtx(nil)
	c.ClefUI = clefUi
	c.answerCh = make(chan int)

	widget := quick.NewQQuickWidget(nil)

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_sign_data.qml", 0))
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()

	return &ApproveSignDataUI{
		UI:            widget,
		ContextObject: c,
	}
}
