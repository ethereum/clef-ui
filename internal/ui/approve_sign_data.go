package ui

import (
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

type ApproveSignDataUI struct {
	UI 					*quick.QQuickWidget
	ContextObject		*ApproveSignDataCtx
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

	_ func(b int) `signal:"clicked,auto"`

	answer 		int
}

func (t *ApproveSignDataCtx) clicked(b int) {
	t.answer = b
}

func (t *ApproveSignDataCtx) Reset() {
	t.answer = 0
	t.SetRemote("")
	t.SetTransport("")
	t.SetEndpoint("")
	t.SetFrom("")
	t.SetMessage("")
	t.SetRawData("")
	t.SetHash("")
}

func (t *ApproveSignDataCtx) ClickResponse(reply *params.ApproveSignDataResponse, res chan bool) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				if t.answer == 1 {
					reply.Approved = false
				} else if t.answer == 2 {
					reply.Approved = true
					reply.Password = "asdfasdf"
				}
				res <- true
				t.Reset()
			}
		}
	}()
}

func NewApproveSignDataUI() *ApproveSignDataUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_sign_data.qml", 0))
	c := NewApproveSignDataCtx(nil)
	v := &ApproveSignDataUI{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}
