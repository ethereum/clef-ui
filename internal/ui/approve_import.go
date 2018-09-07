package ui

import (
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"log"
)

type ApproveImport struct {
	UI 					*quick.QQuickWidget
	ContextObject		*ApproveImportCtx
}


type ApproveImportCtx struct {
	core.QObject

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"password"`
	_ string `property:"confirmPassword"`

	_ func(b int) `signal:"clicked,auto"`
	_ func(b string) `signal:"passwordEdited,auto"`
	_ func(b string) `signal:"confirmPasswordEdited,auto"`
	answer 				int
}

func (t *ApproveImportCtx) clicked(b int) {
	t.answer = b
}

func (t *ApproveImportCtx) Reset() {
	t.answer = 0
	t.SetPassword("")
	t.SetConfirmPassword("")
}

func (t *ApproveImportCtx) passwordEdited(pw string) {
	t.SetPassword(pw)
}

func (t *ApproveImportCtx) confirmPasswordEdited(pw string) {
	t.SetConfirmPassword(pw)
}

func (t *ApproveImportCtx) ClickResponse(reply *params.ApproveImportResponse, res chan bool) {
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

func NewApproveImportUI() *ApproveImport {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_import.qml", 0))
	c := NewApproveImportCtx(nil)
	v := &ApproveImport{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}