package ui

import (
	core2 "github.com/ethereum/go-ethereum/signer/core"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"log"
)

type ApproveNewAccount struct {
	UI            *quick.QQuickWidget
	ContextObject *ApproveNewAccountCtx
}

type ApproveNewAccountCtx struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"password"`
	_ string `property:"confirmPassword"`

	_        func(b int)    `signal:"clicked,auto"`
	_        func()         `signal:"back,auto"`
	_        func(b string) `signal:"passwordEdited,auto"`
	_        func(b string) `signal:"confirmPasswordEdited,auto"`
	answerCh chan int
	ClefUI   *ClefUI
}

func (t *ApproveNewAccountCtx) init() {
}

func (t *ApproveNewAccountCtx) clicked(b int) {
	select {
	case t.answerCh <- -b:
	default:
	}
}

// back is invoked when user clicks back
func (t *ApproveNewAccountCtx) back() {
	select {
	case t.answerCh <- -1:
	default:
	}
}

func (t *ApproveNewAccountCtx) Reset() {
	t.SetRemote("")
	t.SetTransport("")
	t.SetEndpoint("")
	t.SetPassword("")
	t.SetConfirmPassword("")
	t.ClefUI.BackToMain <- true
}

func (t *ApproveNewAccountCtx) passwordEdited(pw string) {
	t.SetPassword(pw)
}

func (t *ApproveNewAccountCtx) confirmPasswordEdited(pw string) {
	t.SetConfirmPassword(pw)
}

func (t *ApproveNewAccountCtx) ClickResponse(replyCh chan *core2.NewAccountResponse) {
	go func() {
		for {
			answer := <-t.answerCh
			if answer == -1 { // Go back
				t.Reset()
				return
			}
			var reply = new(core2.NewAccountResponse)
			if answer == 1 { // Reject
				reply.Approved = false
				t.Reset()
				replyCh <- reply
				return
			}
			// Approval, check that passwords match
			// TODO! Remove passwords from here
			pw := t.Password()
			cpw := t.ConfirmPassword()
			if len(pw) == 0 || len(cpw) == 0 {
				// TODO show error better than just a log output
				log.Println("Password cannot be empty")
			} else if pw != cpw {
				log.Println("Passwords do not match")
			} else {
				reply.Approved = true
				t.Reset()
				replyCh <- reply
				return
			}
		}
	}()
}

func NewApproveNewAccountUI(clefUi *ClefUI) *ApproveNewAccount {
	c := NewApproveNewAccountCtx(nil)
	c.ClefUI = clefUi
	c.answerCh = make(chan int)

	widget := quick.NewQQuickWidget(nil)
	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_new_account.qml", 0))
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return &ApproveNewAccount{
		UI:            widget,
		ContextObject: c,
	}
}
