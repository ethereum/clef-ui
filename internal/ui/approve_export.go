package ui

import (
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

type ApproveExport struct {
	UI 					*quick.QQuickWidget
	ContextObject		*ApproveExportCtx
}


type ApproveExportCtx struct {
	core.QObject

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"address"`
	_ string `property:"password"`

	_ func(b int) `signal:"clicked,auto"`
	_ func(b string) `signal:"passwordEdited,auto"`
	answer 				int
}

func (t *ApproveExportCtx) clicked(b int) {
	t.answer = b
}

func (t *ApproveExportCtx) Reset() {
	t.answer = 0
	t.SetPassword("")
	t.SetAddress("")
}

func (t *ApproveExportCtx) passwordEdited(pw string) {
	t.SetPassword(pw)
}

func (t *ApproveExportCtx) ClickResponse(reply *params.ApproveExportResponse, res chan bool) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				if t.answer == 1 {
					reply.Approved = false
					res <- true
				} else if t.answer == 2 {
					//pw := t.Password()
					//
					//if len(pw) == 0 {
					//	log.Println("Password cannot be empty")
					//	done = false
					//	t.Reset()
					//	continue
					//}

					reply.Approved = true
					//reply.Password = pw
					res <- true
				}
				t.Reset()
			}
		}
	}()
}

func NewApproveExportUI() *ApproveExport {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_export.qml", 0))
	c := NewApproveExportCtx(nil)
	v := &ApproveExport{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}