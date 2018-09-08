package ui

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"strconv"
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
	_ string `property:"password"`

	_ func(b int) `signal:"clicked,auto"`
	_ func(s string, v string) `signal:"edited,auto"`

	answer 		int
	formData 	params.Transaction
	ClefUI 		*ClefUI
}

func (t *ApproveTxCtx) init() {
	t.formData = params.Transaction{}
}

func (t *ApproveTxCtx) SetTransaction(tx params.Transaction) {
	value, _ := hexutil.DecodeBig(tx.Value)
	gas, _ := hexutil.DecodeBig(tx.Gas)
	gasPrice, _ := hexutil.DecodeBig(tx.GasPrice)
	nonce, _ := hexutil.DecodeBig(tx.Nonce)

	t.SetData(tx.Data)
	t.SetNonce(nonce.String())
	t.SetValue(value.String())
	t.SetGas(gas.String())
	t.SetGasPrice(gasPrice.String())
	t.SetFrom(tx.From)
	t.SetTo(tx.To)
}

func (t *ApproveTxCtx) clicked(b int) {
	t.answer = b
}

func (t *ApproveTxCtx) Reset() {
	t.answer = 0
	t.SetRemote("")
	t.SetTransport("")
	t.SetEndpoint("")
	t.SetData("")
	t.SetNonce("")
	t.SetValue("")
	t.SetGas("")
	t.SetGasPrice("")
	t.SetFrom("")
	t.SetTo("")
	t.SetPassword("")
	t.ClefUI.BackToMain <- true
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
	case "password":
		t.SetPassword(value)
	}
}

func (t *ApproveTxCtx) ClickResponse(reply *params.ApproveTxResponse, response chan bool) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				gas, _ := strconv.ParseUint(t.Gas(), 10, 64)
				gasPrice, _ := strconv.ParseUint(t.GasPrice(), 10, 64)
				v, _ := strconv.ParseUint(t.Value(), 10, 64)
				nonce, _ := strconv.ParseUint(t.Nonce(), 10, 64)
				reply.Transaction = params.Transaction{
					Data: t.Data(),
					Nonce: hexutil.EncodeUint64(nonce),
					Value: hexutil.EncodeUint64(v),
					Gas: hexutil.EncodeUint64(gas),
					GasPrice: hexutil.EncodeUint64(gasPrice),
					From: t.From(),
					To: t.To(),
				}

				if t.answer == 1 {
					reply.Approved = false
					response <- true
				} else if t.answer == 2 {
					reply.Approved = true
					value := t.Password()
					reply.Password = value
					response <- true
				}

				t.Reset()
			}
		}
	}()
}

func NewApproveTxUI(clefUi *ClefUI) *ApproveTxUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_tx.qml", 0))
	c := NewApproveTxCtx(nil)
	c.ClefUI = clefUi
	v := &ApproveTxUI{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}
