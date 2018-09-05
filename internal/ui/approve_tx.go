package ui

import (
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

type ApproveTxUI struct {
	UI 					*quick.QQuickWidget
	ContextObject		*ApproveTxCtx
}

type ApproveTxCtx struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"data"`
	_ string `property:"from"`
	_ string `property:"to"`
	_ string `property:"gas"`
	_ string `property:"gasPrice"`
	_ string `property:"nonce"`
	_ string `property:"value"`

	_ func(b int) `signal:"clicked,auto"`

	answer 		int
	formData 	params.Transaction
}

func (t *ApproveTxCtx) init() {
	t.formData = params.Transaction{}
}

func (t *ApproveTxCtx) clicked(b int) {
	t.answer = b
}

func (t *ApproveTxCtx) Reset() {
	t.answer = 0
	t.formData = params.Transaction{}
}

func (t *ApproveTxCtx) ClickResponse(res chan params.ApproveTxResponse) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				r := params.ApproveTxResponse{}
				if t.answer == 1 {
					r.Approved = false
					res <- r
				} else if t.answer == 2 {
					r.Transaction = t.formData
					r.Approved = true
					r.Password = "asdfasdf"
					res <- r
				}
				t.Reset()
			}
		}
	}()
}

func NewApproveTxUI() *ApproveTxUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_tx.qml", 0))
	c := NewApproveTxCtx(nil)
	v := &ApproveTxUI{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}
