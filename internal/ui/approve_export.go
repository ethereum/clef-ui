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

	_ string 			`property:"remote"`
	_ string 			`property:"transport"`
	_ string 			`property:"endpoint"`
	_ string 			`property:"address"`
	_ string 			`property:"password"`
	_ string 			`property:"fromSrc"`

	_ func(b int) 		`signal:"clicked,auto"`
	_ func() 			`signal:"back,auto"`
	_ func(b string) 	`signal:"passwordEdited,auto"`
	answer 				int
	ClefUI 				*ClefUI
}

func (t *ApproveExportCtx) clicked(b int) {
	t.answer = b
}

func (t *ApproveExportCtx) back() {
	t.Reset()
	t.ClefUI.BackToMain <- true
}

func (t *ApproveExportCtx) Reset() {
	t.answer = 0
	t.SetRemote("")
	t.SetTransport("")
	t.SetEndpoint("")
	t.SetPassword("")
	t.SetAddress("")
	t.ClefUI.BackToMain <- true
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
					reply.Approved = true
					res <- true
				}
				t.Reset()
			}
		}
	}()
}

func NewApproveExportUI(clefUi *ClefUI) *ApproveExport {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_export.qml", 0))
	c := NewApproveExportCtx(nil)
	c.ClefUI = clefUi
	v := &ApproveExport{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}