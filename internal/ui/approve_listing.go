package ui

import (
	"fmt"

	"github.com/ethereum/clef-ui/internal/utils"
	"github.com/ethereum/go-ethereum/accounts"
	core2 "github.com/ethereum/go-ethereum/signer/core"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

const (
	Address = int(core.Qt__UserRole) + 1<<iota
	Checked
)

type ApproveListingUI struct {
	UI               *quick.QQuickWidget
	ContextObject    *ApproveListingCtx
	AccountListModel *CustomListModel
}

type ApproveListingCtx struct {
	core.QObject

	_ func() `constructor:"init"`
	//_ func(b int)               `signal:"clicked,auto"`
	_ func()                    `signal:"back,auto"`
	_ func()                    `signal:"approve,auto"`
	_ func()                    `signal:"reject,auto"`
	_ func(i int, checked bool) `signal:"onCheckStateChanged,auto"`

	_ func() `signal:"triggerUpdate"`

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"from"`
	_ string `property:"message"`
	_ string `property:"rawData"`
	_ string `property:"hash"`

	accounts *CustomListModel
	ClefUI   *ClefUI
	answerCh chan (int)
}

func (ctx *ApproveListingCtx) init() {
	ctx.accounts = NewCustomListModel(nil)
}

func (ctx *ApproveListingCtx) Reset() {
	ctx.SetRemote("")
	ctx.SetTransport("")
	ctx.SetEndpoint("")
	ctx.accounts.Clear()
	ctx.ClefUI.BackToMain <- true
}

func init() {
	CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomListModel")
}

type AccountListItem struct {
	address  string
	selected bool
}

func (t *ApproveListingCtx) onCheckStateChanged(i int, checked bool) {
	t.accounts.checkState[i] = checked
}

func (ctx *ApproveListingCtx) back() {
	select {
	case ctx.answerCh <- USR_BACK:
	default:
	}
}
func (ctx *ApproveListingCtx) reject() {
	select {
	case ctx.answerCh <- USR_REJECT:
	default:
	}
}
func (ctx *ApproveListingCtx) approve() {
	select {
	case ctx.answerCh <- USR_APPROVE:
	default:
	}
}

// ExternalSetAccounts should be used to update the model from not-gui-thread, e.g. when new data arrives
// from the network.
func (ctx *ApproveListingCtx) ExternalSetAccounts(accounts []accounts.Account) {
	// We can update the model data in-place
	ctx.accounts.modelData = append(ctx.accounts.modelData, accounts...)
	// Then invoke the QML object, which will make a callback to us.
	// Once we're in the callback, we can signal BeginResetModel and EndResetModel
	ctx.TriggerUpdate()
}

func (ctx *ApproveListingCtx) ClickResponse(res chan *core2.ListResponse) {
	go func() {
		// Wait for user to complete the form
		answer := <-ctx.answerCh
		if answer != USR_BACK { // Go back
			reply := new(core2.ListResponse)
			if answer == USR_APPROVE { // Approve at least some accounts
				accounts := make([]accounts.Account, 0)
				for i, account := range ctx.accounts.modelData {
					if ctx.accounts.checkState[i] {
						accounts = append(accounts, account)
					}
				}
				reply.Accounts = accounts
			}
			res <- reply
		}
		ctx.Reset()
	}()
}

// CustomListModel implements a custom list model to handle the list of accounts
type CustomListModel struct {
	core.QAbstractListModel

	_ func()                         `constructor:"init"`
	_ func()                         `signal:"clear,auto"`
	_ func(account accounts.Account) `signal:"add,auto"`

	// Signals to make the UI thread issue beginInsertRow/endInsertRows, if we want to make updates
	// to the model
	_ func() `slot:"callbackFromQml"`

	_ string `property:"updateRequired"`

	modelData  []accounts.Account
	checkState map[int]bool
}

func (model *CustomListModel) init() {
	model.modelData = []accounts.Account{}
	model.checkState = map[int]bool{}

	model.ConnectRoleNames(model.roleNames)
	model.ConnectRowCount(model.rowCount)
	model.ConnectColumnCount(model.columnCount)
	model.ConnectData(model.data)
}

func (model *CustomListModel) rowCount(*core.QModelIndex) int {
	return len(model.modelData)
}

func (model *CustomListModel) columnCount(*core.QModelIndex) int {
	return 2
}

func (model *CustomListModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		Address: core.NewQByteArray2("address", -1),
		Checked: core.NewQByteArray2("selected", -1),
	}
}

func (model *CustomListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := model.modelData[index.Row()]
	if role == int(Address) {
		address, _ := clefutils.ToChecksumAddress(item.Address.String())
		return core.NewQVariant14(address)
	}
	if role == int(Checked) {
		return core.NewQVariant11(true)
	}
	return core.NewQVariant()
}

func (model *CustomListModel) clear() {
	model.BeginResetModel()
	model.modelData = []accounts.Account{}
	model.checkState = map[int]bool{}
	model.EndResetModel()
}

func (model *CustomListModel) add(account accounts.Account) {
	model.BeginInsertRows(core.NewQModelIndex(), len(model.modelData), 1+len(model.modelData))
	model.modelData = append(model.modelData, account)
	model.EndInsertRows()
}

func NewApproveListingUI(clefUi *ClefUI) *ApproveListingUI {
	ctx := NewApproveListingCtx(nil)
	ctx.ClefUI = clefUi
	ctx.answerCh = make(chan int)
	ctx.accounts.ConnectCallbackFromQml(func() {
		fmt.Print("callback from qml\n")
		ctx.accounts.BeginResetModel()
		ctx.accounts.EndResetModel()

	})

	widget := quick.NewQQuickWidget(nil)
	widget.RootContext().SetContextProperty("ctxObject", ctx)
	widget.RootContext().SetContextProperty("accounts", ctx.accounts)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_listing.qml", 0))
	widget.SetStyleSheet("margin: 0;")
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	widget.Hide()

	return &ApproveListingUI{
		UI:            widget,
		ContextObject: ctx,
	}
}
