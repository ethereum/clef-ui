package ui

import (
	"github.com/ethereum/go-ethereum/accounts"
	core2 "github.com/ethereum/go-ethereum/signer/core"
	"github.com/kyokan/clef-ui/internal/utils"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

type ApproveListingUI struct {
	UI               *quick.QQuickWidget
	ContextObject    *ApproveListingCtx
	AccountListModel *CustomListModel
}

type ApproveListingCtx struct {
	core.QObject

	_ func()                    `constructor:"init"`
	_ func(b int)               `signal:"clicked,auto"`
	_ func()                    `signal:"back,auto"`
	_ func(i int, checked bool) `signal:"onCheckStateChanged,auto"`

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

func init() { CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomListModel") }

const (
	Address = int(core.Qt__UserRole) + 1<<iota
	Checked
)

type AccountListItem struct {
	address  string
	selected bool
}

type CustomListModel struct {
	core.QAbstractListModel

	_ func()                         `constructor:"init"`
	_ func()                         `signal:"clear,auto"`
	_ func(account accounts.Account) `signal:"add,auto"`

	modelData  []accounts.Account
	checkState map[int]bool
}

func (m *CustomListModel) init() {
	m.modelData = []accounts.Account{}
	m.checkState = map[int]bool{}

	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomListModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomListModel) columnCount(*core.QModelIndex) int {
	return 2
}

func (m *CustomListModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		Address: core.NewQByteArray2("address", -1),
		Checked: core.NewQByteArray2("selected", -1),
	}
}

func (m *CustomListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]

	if role == int(Address) {
		address, _ := clefutils.ToChecksumAddress(item.Address.String())
		return core.NewQVariant14(address)
	}

	if role == int(Checked) {
		return core.NewQVariant11(true)
	}

	return core.NewQVariant()
}

func (m *CustomListModel) clear() {
	m.BeginResetModel()
	m.modelData = []accounts.Account{}
	m.checkState = map[int]bool{}
	m.EndResetModel()
}

func (m *CustomListModel) add(account accounts.Account) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, account)
	m.EndInsertRows()
}

func (t *ApproveListingCtx) clicked(b int) {
	t.answerCh <- b
}

func (t *ApproveListingCtx) onCheckStateChanged(i int, checked bool) {
	t.accounts.checkState[i] = checked
}

func (t *ApproveListingCtx) back() {
	select {
	case t.answerCh <- -1:
	default:
	}
}

func (t *ApproveListingCtx) ClickResponse(res chan *core2.ListResponse) {
	go func() {
		// Wait for user to complete the form
		answer := <-t.answerCh
		if answer != -1 { // Go back
			reply := new(core2.ListResponse)
			if answer == 2 { // Approve at least some accounts
				accounts := make([]accounts.Account, 0)
				for i, account := range t.accounts.modelData {
					if t.accounts.checkState[i] {
						accounts = append(accounts, account)
					}
				}
				reply.Accounts = accounts
			}
			res <- reply
		}
		t.Reset()
	}()
}

func NewApproveListingUI(clefUi *ClefUI) *ApproveListingUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_listing.qml", 0))
	c := NewApproveListingCtx(nil)
	c.ClefUI = clefUi
	c.answerCh = make(chan int)
	v := &ApproveListingUI{
		UI:            widget,
		ContextObject: c,
	}
	widget.SetStyleSheet("margin: 0;")
	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.RootContext().SetContextProperty("accounts", c.accounts)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}
