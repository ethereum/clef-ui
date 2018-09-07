package ui

import (
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"log"
)

type ApproveNewAccount struct {
	UI 					*quick.QQuickWidget
	ContextObject		*ApproveNewAccountCtx
}

type ApproveNewAccountCtx struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"password"`
	_ string `property:"confirmPassword"`

	_ func(b int) `signal:"clicked,auto"`
	_ func(b string) `signal:"passwordEdited,auto"`
	_ func(b string) `signal:"confirmPasswordEdited,auto"`
	answer 				int
	ClefUI 				*ClefUI
}

func (t *ApproveNewAccountCtx) init() {
}

func (t *ApproveNewAccountCtx) clicked(b int) {
	t.answer = b
}

func (t *ApproveNewAccountCtx) Reset() {
	t.answer = 0
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

func (t *ApproveNewAccountCtx) ClickResponse(reply *params.ApproveNewAccountResponse, res chan bool) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				if t.answer == 1 {
					reply.Approved = false
					reply.Password = ""
					res <- true
				} else if t.answer == 2 {
					pw := t.Password()
					cpw := t.ConfirmPassword()

					if len(pw) == 0 || len(cpw) == 0 {
						log.Println("Password cannot be empty")
						done = false
						t.Reset()
						continue
					}
					if pw != cpw {
						log.Printf("%v does not match %v", pw, cpw)
						done = false
						t.Reset()
						continue
					}
					reply.Approved = true
					reply.Password = cpw
					res <- true
				}
				t.Reset()
			}
		}
	}()
}

func NewApproveNewAccountUI(clefUi *ClefUI) *ApproveNewAccount {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_new_account.qml", 0))
	c := NewApproveNewAccountCtx(nil)
	c.ClefUI = clefUi
	v := &ApproveNewAccount{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}
