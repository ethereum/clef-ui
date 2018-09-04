package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

type CustomLabel struct {
	widgets.QLabel

	_ func(string) `signal:"updateText,auto(this.QLabel.setText)"` //TODO: support this.setText as well
}

type ApproveSignDataUI struct {
	UI 					*quick.QQuickWidget
	ContextObject		*CtxObject
}

type CtxObject struct {
	core.QObject

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"from"`
	_ string `property:"message"`
	_ string `property:"rawData"`
	_ string `property:"hash"`

	_ func(b int) `signal:"clicked,auto"`

	answer 		int
}

func (t *CtxObject) clicked(b int) {
	t.answer = b
}

func (t *CtxObject) Reset() {
	t.answer = 0
}

func (t *CtxObject) ClickResponse(res chan map[string]string) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				if t.answer == 1 {
					res <- map[string]string{
						"approved": "false",
						"password": "",
					}
				} else if t.answer == 2 {
					res <- map[string]string{
						"approved": "true",
						"password": "asdfasdf",
					}
				}
				t.Reset()
			}
		}
	}()
}

func NewApproveSignDataUI() *ApproveSignDataUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_sign_data.qml", 0))
	c := NewCtxObject(nil)
	v := &ApproveSignDataUI{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)

	return v
}
