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
	_ func(s string, v string) `signal:"edited,auto"`

	answer 		int
	formData 	params.Transaction
}

func (t *ApproveTxCtx) init() {
	t.formData = params.Transaction{}
}

func (t *ApproveTxCtx) SetTransaction(tx params.Transaction) {
	t.SetData(tx.Data)
	t.SetNonce(tx.Nonce)
	t.SetValue(tx.Value)
	t.SetGas(tx.Gas)
	t.SetGasPrice(tx.GasPrice)
	t.SetFrom(tx.From)
	t.SetTo(tx.To)
}

func (t *ApproveTxCtx) clicked(b int) {
	t.answer = b
}

func (t *ApproveTxCtx) Reset() {
	t.answer = 0
	t.SetData("")
	t.SetNonce("")
	t.SetValue("")
	t.SetGas("")
	t.SetGasPrice("")
	t.SetFrom("")
	t.SetTo("")
}

func (t *ApproveTxCtx) edited(name string, value string) {
	switch name {
	case "data":
		t.SetData(value)
	case "nonce":
		t.SetNonce(value)
	case "from":
		t.SetFrom(value)
	case "to":
		t.SetTo(value)
	case "gas":
		t.SetGas(value)
	case "gasPrice":
		t.SetGasPrice(value)
	case "value":
		t.SetValue(value)
	}
}

func (t *ApproveTxCtx) ClickResponse(reply *params.ApproveTxResponse, response chan bool) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				reply.Transaction = params.Transaction{
					Data: t.Data(),
					Nonce: t.Nonce(),
					Value: t.Value(),
					Gas: t.Gas(),
					GasPrice: t.GasPrice(),
					From: t.From(),
					To: t.To(),
				}

				if t.answer == 1 {
					reply.Approved = false
					response <- true
				} else if t.answer == 2 {
					reply.Approved = true
					reply.Password = "asdfasdf"
					response <- true
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
