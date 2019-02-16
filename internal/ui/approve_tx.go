package ui

import (
	"fmt"
	"github.com/kyokan/clef-ui/internal/identicon"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	core2 "github.com/ethereum/go-ethereum/signer/core"
	"github.com/kyokan/clef-ui/internal/utils"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

const (
	WEI  = 0
	GWEI = 1
	ETH  = 2
)

type diff struct {
	key      string
	original string
	edited   string
}

type ApproveTxUI struct {
	UI            *quick.QQuickWidget
	ContextObject *ApproveTxCtx
}

type ApproveTxCtx struct {
	core.QObject

	_ func() `constructor:"init"`

	_ int    `property:"valueUnit"`
	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"data"`
	_ string `property:"from"`
	_ string `property:"fromWarning"`
	_ bool   `property:"fromVisible"`
	_ string `property:"to"`
	_ string `property:"toWarning"`
	_ bool   `property:"toVisible"`
	_ string `property:"gas"`
	_ string `property:"gasPrice"`
	_ int    `property:"gasPriceUnit"`
	_ string `property:"nonce"`
	_ string `property:"value"`
	_ string `property:"password"`
	_ string `property:"fromSrc"`
	_ string `property:"toSrc"`
	_ string `property:"diff"`

	_ func(b int)              `signal:"clicked,auto"`
	_ func()                   `signal:"checkTxDiff,auto"`
	_ func()                   `signal:"back,auto"`
	_ func(s string, v string) `signal:"edited,auto"`
	_ func(v int)              `signal:"changeValueUnit,auto"`
	_ func(v int)              `signal:"changeGasPriceUnit,auto"`

	answerCh chan (int)
	formData core2.SendTxArgs
	ClefUI   *ClefUI
	rawTx    core2.SendTxArgs
}

func (t *ApproveTxCtx) init() {
	t.formData = core2.SendTxArgs{}
	t.SetValueUnit(GWEI)
	t.SetGasPriceUnit(GWEI)
}

func (t *ApproveTxCtx) SetTransaction(tx core2.SendTxArgs) {

	t.rawTx = tx

	t.SetData(tx.Data.String())
	t.SetNonce(fmt.Sprintf("%d", tx.Nonce))
	{
		bigV := big.Int(tx.Value)
		t.SetValueUnit(GWEI)
		t.SetValue(clefutils.WeiToString(&bigV, clefutils.GWei))
	}
	t.SetGas(fmt.Sprintf("%d", tx.Gas))

	{
		bigGas := big.Int(tx.GasPrice)
		t.SetGasPriceUnit(GWEI)
		t.SetGasPrice(clefutils.WeiToString(&bigGas, clefutils.GWei))
	}
	t.SetFrom(tx.From.Original())
	t.SetFromSrc(identicon.ToBase64Img(tx.From.Address()))

	if tx.To != nil {
		t.SetTo(tx.To.Original())
		t.SetToSrc(identicon.ToBase64Img(tx.To.Address()))
	}
}

func indexToUnit(index int) uint64 {
	switch index {
	case WEI:
		return clefutils.Wei
	case GWEI:
		return clefutils.GWei
	case ETH:
		return clefutils.Ether
	default:
		// TODO Make it possible to signal errors
		fmt.Errorf("Unknown unit index")
		return 0
	}
}
func (t *ApproveTxCtx) changeValueUnit(toIndex int) {
	fromUnit := indexToUnit(t.ValueUnit())
	weiVal, err := clefutils.TextToWei(t.Value(), fromUnit)
	if err != nil {
		log.Printf("Unable to parse value %v", t.Value())
		return
	}
	var toUnit = indexToUnit(toIndex)
	strVal := clefutils.WeiToString(weiVal, toUnit)
	t.SetValue(strVal)
	t.SetValueUnit(toIndex)
}

func (t *ApproveTxCtx) changeGasPriceUnit(toIndex int) {
	fromUnit := indexToUnit(t.GasPriceUnit())
	weiVal, err := clefutils.TextToWei(t.GasPrice(), fromUnit)
	if err != nil {
		log.Printf("Unable to parse value %v", t.GasPrice())
		return
	}
	var toUnit = indexToUnit(toIndex)
	strVal := clefutils.WeiToString(weiVal, toUnit)
	t.SetGasPrice(strVal)
	t.SetGasPriceUnit(toIndex)
}

// Reset clears the form and signals to go back to main
func (t *ApproveTxCtx) Reset() {
	t.SetRemote("")
	t.SetTransport("")
	t.SetEndpoint("")
	t.SetData("")
	t.SetNonce("")
	t.SetValue("")
	t.SetValueUnit(GWEI)
	t.SetGas("")
	t.SetGasPrice("")
	t.SetGasPriceUnit(GWEI)
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
	t.formData = core2.SendTxArgs{}

	t.ClefUI.BackToMain <- true
}

// edited is invoked by the QT framework when the user edits the fields,
// making it possible to react to the events, e.g. updating identicons
func (t *ApproveTxCtx) edited(name string, value string) {

	switch name {
	case "data":
		t.SetData(value)
	case "nonce":
		t.SetNonce(value)
	case "from":
		t.SetFrom(value)
		var (
			warning = ""
			icon    = ""
		)
		addr, err := common.NewMixedcaseAddressFromString(value)
		if err != nil {
			warning = "Invalid Address"
		} else {
			icon = identicon.ToBase64Img(addr.Address())
			if !addr.ValidChecksum() {
				warning = "Invalid checksum"
			}
		}
		t.SetFromWarning(warning)
		t.SetFromSrc(icon)
		t.SetFromVisible(len(warning) > 0)
	case "to":
		t.SetTo(value)
		var (
			warning = ""
			icon    = ""
		)
		if len(value) == 0 {
			warning = "Contract creation"
		} else {
			addr, err := common.NewMixedcaseAddressFromString(value)
			if err != nil {
				warning = "Invalid Address"
			} else {
				icon = identicon.ToBase64Img(addr.Address())
				if !addr.ValidChecksum() {
					warning = "Invalid Checksum"
				}
			}
		}
		t.SetToWarning(warning)
		t.SetToSrc(icon)
		t.SetToVisible(len(warning) > 0)
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

func (t *ApproveTxCtx) getNewTx() (core2.SendTxArgs, error) {
	tx := core2.SendTxArgs{}

	// Gasprice
	if gp, err := clefutils.TextToWei(t.GasPrice(), indexToUnit(t.GasPriceUnit())); err != nil {
		return tx, err
	} else {
		tx.GasPrice = hexutil.Big(*gp)
	}
	// Value
	if v, err := clefutils.TextToWei(t.Value(), indexToUnit(t.ValueUnit())); err != nil {
		return tx, err
	} else {
		tx.Value = hexutil.Big(*v)
	}
	// Gas
	if gas, err := strconv.ParseUint(t.Gas(), 10, 64); err != nil {
		return tx, err
	} else {
		tx.Gas = hexutil.Uint64(gas)
	}
	// Nonce
	if nonce, err := strconv.ParseUint(t.Nonce(), 10, 64); err != nil {
		return tx, err
	} else {
		tx.Nonce = hexutil.Uint64(nonce)
	}
	// Data
	if data, err := hexutil.Decode(t.Data()); err != nil {
		return tx, err
	} else {
		d := hexutil.Bytes(data)
		tx.Data = &d
	}
	if from, err := common.NewMixedcaseAddressFromString(t.From()); err != nil {
		return tx, err
	} else {
		tx.From = *from
	}
	if txtTo := t.To(); len(txtTo) == 0 {
		// No receiver, creates a contract
		tx.To = nil
	} else {
		if to, err := common.NewMixedcaseAddressFromString(txtTo); err != nil {
			return tx, err
		} else {
			tx.To = to
		}
	}

	return tx, nil
}

func (t *ApproveTxCtx) checkTxDiff() {
	t.checkTxDiffWithError()
}
func (t *ApproveTxCtx) checkTxDiffWithError() error {
	modified, err := t.getNewTx()
	if err != nil {
		t.SetDiff(fmt.Sprint("Could not calculate diff: %v", err))
		return err
	}
	diffs := compareTransactions(t.rawTx, modified)
	if len(diffs) == 0 {
		t.SetDiff("")
		return nil
	}
	changes := ""
	for _, diff := range diffs {
		diffString := fmt.Sprintf("[%v]\nFrom:%v\nTo:%v\n\n", strings.ToUpper(diff.key), diff.original, diff.edited)
		changes += diffString
	}
	t.SetDiff(changes)
	return nil
}

// back is invoked when user clicks back
func (t *ApproveTxCtx) back() {
	select {
	case t.answerCh <- -1:
	default:
	}
}

// clicked means user clicked either Approve or Reject
func (t *ApproveTxCtx) clicked(b int) {
	select {
	case t.answerCh <- b:
	default:
	}

}

func (t *ApproveTxCtx) ClickResponse(responseCh chan *core2.SignTxResponse) {
	go func() {
		answer := <-t.answerCh
		if answer == -1 { // Go back
			t.Reset()
			return
		}
		var reply = new(core2.SignTxResponse)
		if answer == 1 { // Reject
			reply.Approved = false
		} else { // Approve
			// TODO handle errors here
			// We might want to validate already in the clicked() method, and pass the
			// replyTx over the channel
			replyTx, _ := t.getNewTx()
			reply.Transaction = replyTx
			reply.Approved = true

			reply.Password = t.Password()
		}
		t.Reset()
		responseCh <- reply
	}()
}

func maybeAddDiff(diffs *[]diff, o string, n string, key string) {
	if o == n {
		return
	}
	*diffs = append(*diffs, diff{
		key:      key,
		original: o,
		edited:   n,
	})
}

func compareTransactions(o core2.SendTxArgs, n core2.SendTxArgs) []diff {
	var diffs []diff
	maybeAddDiff(&diffs, o.Data.String(), n.Data.String(), "data")
	maybeAddDiff(&diffs, o.Nonce.String(), n.Nonce.String(), "nonce")
	maybeAddDiff(&diffs, o.Value.String(), n.Value.String(), "value")
	maybeAddDiff(&diffs, o.Gas.String(), n.Gas.String(), "gas")
	maybeAddDiff(&diffs, o.GasPrice.String(), n.GasPrice.String(), "gasPrice")
	maybeAddDiff(&diffs, o.From.Address().String(), n.From.Address().String(), "from")
	maybeAddDiff(&diffs, o.To.Address().String(), n.To.Address().String(), "to")
	return diffs
}

func NewApproveTxUI(clefUi *ClefUI) *ApproveTxUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_tx.qml", 0))
	c := NewApproveTxCtx(nil)
	c.ClefUI = clefUi
	v := &ApproveTxUI{
		UI:            widget,
		ContextObject: c,
	}
	c.answerCh = make(chan int)
	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}
