package ui

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/kyokan/clef-ui/internal/identicon"
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/kyokan/clef-ui/internal/utils"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"log"
	"math/big"
	"strconv"
	"strings"
)

const (
	WEI  	= 0
	GWEI 	= 1
	ETH		= 2
)

type Diff struct {
	Key 			string
	OriginalValue 	string
	NewValue 		string
}

type ApproveTxUI struct {
	UI 					*quick.QQuickWidget
	ContextObject		*ApproveTxCtx
}

type ApproveTxCtx struct {
	core.QObject

	_ func() 					`constructor:"init"`

	_ float64 					`property:"valueUnit"`
	_ string 					`property:"remote"`
	_ string 					`property:"transport"`
	_ string 					`property:"endpoint"`
	_ string 					`property:"data"`
	_ string 					`property:"from"`
	_ string 					`property:"fromWarning"`
	_ bool 	 					`property:"fromVisible"`
	_ string 					`property:"to"`
	_ string 					`property:"toWarning"`
	_ bool   					`property:"toVisible"`
	_ string 					`property:"gas"`
	_ string 					`property:"gasPrice"`
	_ float64 					`property:"gasPriceUnit"`
	_ string 					`property:"nonce"`
	_ string 					`property:"value"`
	_ string 					`property:"password"`
	_ string 					`property:"fromSrc"`
	_ string 					`property:"toSrc"`
	_ string 					`property:"diff"`

	_ func(b int) 				`signal:"clicked,auto"`
	_ func() 					`signal:"checkTxDiff,auto"`
	_ func() 					`signal:"back,auto"`
	_ func(s string, v string) 	`signal:"edited,auto"`
	_ func(v int) 				`signal:"changeValueUnit,auto"`
	_ func(v int) 				`signal:"changeGasPriceUnit,auto"`

	answer 		int
	formData 	params.Transaction
	ClefUI 		*ClefUI
	rawTx 		params.Transaction
}

func (t *ApproveTxCtx) init() {
	t.formData = params.Transaction{}
	t.SetValueUnit(clefutils.Ether)
	t.SetGasPriceUnit(clefutils.GWei)
}

func (t *ApproveTxCtx) back() {
	t.Reset()
	t.ClefUI.BackToMain <- true
}

func (t *ApproveTxCtx) SetTransaction(tx params.Transaction) {

	value, _ := hexutil.DecodeBig(tx.Value)
	gas, _ := hexutil.DecodeBig(tx.Gas)
	gasPrice, _ := hexutil.DecodeBig(tx.GasPrice)
	nonce, _ := hexutil.DecodeBig(tx.Nonce)
	t.rawTx = tx

	t.SetFromSrc(identicon.ToBase64Img(tx.From))
	t.SetToSrc(identicon.ToBase64Img(tx.To))
	t.SetData(tx.Data)
	t.SetNonce(nonce.String())
	t.changeValue(value.String(), clefutils.Wei, t.ValueUnit())
	t.SetGas(gas.String())
	t.changeGasPrice(gasPrice.String(), clefutils.Wei, t.GasPriceUnit())

	address, _ := clefutils.ToChecksumAddress(tx.From)

	t.SetFrom(address)
	t.SetTo(tx.To)
}

func (t *ApproveTxCtx) changeValue(value string, fromUnit float64, toUnit float64) {
	floatValue, _, err := big.ParseFloat(value, 10, 0, big.ToNearestEven)
	if err != nil {
		log.Printf("Cannot parse float value %v", value)
		return
	}

	valueString := clefutils.ConvertUnitAndGetString(floatValue, fromUnit, toUnit)
	t.SetValue(valueString)
}

func (t *ApproveTxCtx) changeGasPrice(value string, fromUnit float64, toUnit float64) {
	floatValue, _, err := big.ParseFloat(value, 10, 0, big.ToNearestEven)
	if err != nil {
		log.Printf("Cannot parse float value %v", value)
		return
	}

	valueString := clefutils.ConvertUnitAndGetString(floatValue, fromUnit, toUnit)
	t.SetGasPrice(valueString)
}

func (t *ApproveTxCtx) clicked(b int) {
	t.answer = b
}

func (t *ApproveTxCtx) changeValueUnit(unitIndex int) {
	value := t.Value()
	floatValue, _, err := big.ParseFloat(value, 10, 0, big.ToNearestEven)
	previousUnit := t.ValueUnit()

	if err != nil {
		log.Printf("Cannot parse value %v", value)
		return
	}

	switch unitIndex {
	case WEI:
		value = clefutils.ConvertUnitAndGetString(floatValue, previousUnit, clefutils.Wei)
		t.SetValueUnit(clefutils.Wei)
		t.SetValue(value)
		return
	case GWEI:
		value = clefutils.ConvertUnitAndGetString(floatValue, previousUnit, clefutils.GWei)
		t.SetValueUnit(clefutils.GWei)
		t.SetValue(value)
	case ETH:
		value = clefutils.ConvertUnitAndGetString(floatValue, previousUnit, clefutils.Ether)
		t.SetValueUnit(clefutils.Ether)
		t.SetValue(value)
	}
}

func (t *ApproveTxCtx) changeGasPriceUnit(unitIndex int) {
	gasPrice := t.GasPrice()
	floatValue, _, err := big.ParseFloat(gasPrice, 10, 0, big.ToNearestEven)
	previousUnit := t.GasPriceUnit()

	if err != nil {
		log.Printf("Cannot parse value %v", gasPrice)
		return
	}
	switch unitIndex {
	case WEI:
		gasPrice = clefutils.ConvertUnitAndGetString(floatValue, previousUnit, clefutils.Wei)
		t.SetGasPriceUnit(clefutils.Wei)
		t.SetGasPrice(gasPrice)
	case GWEI:
		gasPrice = clefutils.ConvertUnitAndGetString(floatValue, previousUnit, clefutils.GWei)
		t.SetGasPriceUnit(clefutils.GWei)
		t.SetGasPrice(gasPrice)
	case ETH:
		gasPrice = clefutils.ConvertUnitAndGetString(floatValue, previousUnit, clefutils.Ether)
		t.SetGasPriceUnit(clefutils.Ether)
		t.SetGasPrice(gasPrice)
	}
}

func (t *ApproveTxCtx) Reset() {
	t.answer = 0
	t.SetRemote("")
	t.SetTransport("")
	t.SetEndpoint("")
	t.SetData("")
	t.SetNonce("")
	t.SetValue("")
	t.SetValueUnit(clefutils.Ether)
	t.SetGas("")
	t.SetGasPrice("")
	t.SetGasPriceUnit(clefutils.GWei)
	t.SetFrom("")
	t.SetFromSrc("")
	t.SetTo("")
	t.SetToSrc("")
	t.SetPassword("")
	t.SetFromVisible(false)
	t.SetToVisible(false)
	t.SetFromWarning("")
	t.SetToWarning("")
	t.SetDiff("")
	t.rawTx = params.Transaction{}

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
		checksum, err := clefutils.ToChecksumAddress(value)
		if err != nil {
			t.SetFromWarning("Invalid Address")
			t.SetFromVisible(true)
			return
		}
		if checksum != value {
			t.SetFromWarning("Invalid Checksum")
			t.SetFromVisible(true)
			return
		}

		t.SetFromWarning("")
		t.SetFromVisible(false)
	case "to":
		t.SetTo(value)
		checksum, err := clefutils.ToChecksumAddress(value)
		if err != nil {
			t.SetToWarning("Invalid Address")
			t.SetToVisible(true)
			return
		}
		if checksum != value {
			t.SetToWarning("Invalid Checksum")
			t.SetToVisible(true)
			return
		}

		t.SetToWarning("")
		t.SetToVisible(false)
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

func (t *ApproveTxCtx) getNewTx() params.Transaction {
	gasPriceFloat, _, err := big.ParseFloat(t.GasPrice(), 10, 0, big.ToNearestEven)
	if err != nil {
		log.Printf("Cannot parse float value %v", t.GasPrice())
		return params.Transaction{}
	}
	gasPriceString := clefutils.ConvertUnitAndGetString(gasPriceFloat, t.GasPriceUnit(), clefutils.Wei)

	valueFloat, _, err := big.ParseFloat(t.Value(), 10, 0, big.ToNearestEven)
	if err != nil {
		log.Printf("Cannot parse float value %v", t.Value())
		return params.Transaction{}
	}
	valueString := clefutils.ConvertUnitAndGetString(valueFloat, t.ValueUnit(), clefutils.Wei)


	gasPrice, _ := strconv.ParseUint(gasPriceString, 10, 64)
	gas, _ := strconv.ParseUint(t.Gas(), 10, 64)
	v, _ := strconv.ParseUint(valueString, 10, 64)
	nonce, _ := strconv.ParseUint(t.Nonce(), 10, 64)

	return params.Transaction{
		Data: t.Data(),
		Nonce: hexutil.EncodeUint64(nonce),
		Value: hexutil.EncodeUint64(v),
		Gas: hexutil.EncodeUint64(gas),
		GasPrice: hexutil.EncodeUint64(gasPrice),
		From: t.From(),
		To: t.To(),
	}
}

func (t *ApproveTxCtx) checkTxDiff() {
	diffs := compareTransactions(t.rawTx,  t.getNewTx())

	if len(diffs) == 0 {
		t.SetDiff("")
		return
	}

	changes := ""

	for _, diff := range diffs {
		diffString := fmt.Sprintf("[%v]\nFrom:%v\nTo:%v\n\n", strings.ToUpper(diff.Key), diff.OriginalValue, diff.NewValue)
		changes += diffString
	}

	t.SetDiff(changes)
}

func (t *ApproveTxCtx) ClickResponse(reply *params.ApproveTxResponse, response chan bool) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				replyTx := t.getNewTx()

				reply.Transaction = replyTx

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

func maybeAddDiff(diffs *[]Diff, o string, n string, key string) {
	if o == n {
		return
	}

	diff := Diff{
		Key: key,
		OriginalValue: o,
		NewValue: n,
	}

	*diffs = append(*diffs, diff)
}

func compareTransactions(o params.Transaction, n params.Transaction) []Diff {
	var diffs []Diff

	maybeAddDiff(&diffs, o.Data, n.Data, "data")
	maybeAddDiff(&diffs, o.Nonce, n.Nonce, "nonce")
	maybeAddDiff(&diffs, o.Value, n.Value, "value")
	maybeAddDiff(&diffs, o.Gas, n.Gas, "gas")
	maybeAddDiff(&diffs, o.GasPrice, n.GasPrice, "gasPrice")
	maybeAddDiff(&diffs, strings.ToLower(o.From), strings.ToLower(n.From), "from")
	maybeAddDiff(&diffs, strings.ToLower(o.To), strings.ToLower(n.To), "to")

	return diffs
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
