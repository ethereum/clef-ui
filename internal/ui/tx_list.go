package ui

import (
	"strings"

	"github.com/ethereum/clef-ui/internal/identicon"
	"github.com/ethereum/go-ethereum/common"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

func init() {
	CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "TxListModel")
}

const (
	From = int(core.Qt__UserRole) + 1<<iota
	Method
	FromSrc
	IsUnknown
)

var KnownAccounts = make(map[common.Address]struct{})

type TxListUI struct {
	UI        *quick.QQuickWidget
	CtxObject *TxListCtx
}

type requestInvocation interface {
	handle(ui *ClefUI)
}

type IncomingRequestItem struct {
	From        string
	Description string
	RPC         requestInvocation
	IsUnknown   int
	OnRemove    chan int
	ID          int
}

func (item *IncomingRequestItem) Remove() {
	item.OnRemove <- item.ID
}

type TxListModel struct {
	core.QAbstractListModel

	_ bool `property:"isEmpty"`

	_ func()                       `constructor:"init"`
	_ func()                       `signal:"clear,auto"`
	_ func(tx IncomingRequestItem) `signal:"add,auto"`
	_ func(i int)                  `signal:"remove,auto"`

	modelData    []IncomingRequestItem
	idCounter    int
	removeItemCh chan int
}

func (m *TxListModel) init() {
	m.modelData = []IncomingRequestItem{}
	m.idCounter = 0
	m.removeItemCh = make(chan int)

	go func() {
		for {
			index := <-m.removeItemCh
			m.Remove(index)
		}
	}()

	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *TxListModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *TxListModel) columnCount(*core.QModelIndex) int {
	return 3
}

func (m *TxListModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		From:      core.NewQByteArray2("from", -1),
		Method:    core.NewQByteArray2("method", -1),
		FromSrc:   core.NewQByteArray2("fromSrc", -1),
		IsUnknown: core.NewQByteArray2("isUnknown", -1),
	}
}

func (m *TxListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]

	if role == int(From) {
		if address, err := common.NewMixedcaseAddressFromString(item.From); err != nil {
			return core.NewQVariant15(item.From)
		} else {
			return core.NewQVariant15(address.Original())

		}
	}
	if role == int(Method) {
		return core.NewQVariant15(item.Description)
	}
	if role == int(FromSrc) {
		if addr, err := common.NewMixedcaseAddressFromString(item.From); err != nil {
			return core.NewQVariant15("")
		} else {
			return core.NewQVariant15(identicon.ToBase64Img(addr.Address()))
		}
	}
	if role == int(IsUnknown) {
		return core.NewQVariant5(item.IsUnknown)
	}
	return core.NewQVariant()
}

func (m *TxListModel) clear() {
	m.BeginResetModel()
	m.modelData = []IncomingRequestItem{}
	m.EndResetModel()
}

func (m *TxListModel) add(tx IncomingRequestItem) {
	address := strings.ToLower(tx.From)
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	tx.ID = m.idCounter
	tx.OnRemove = m.removeItemCh

	if address == " - " {
		tx.IsUnknown = 0
	} else if _, exists := KnownAccounts[common.HexToAddress(address)]; exists {
		tx.IsUnknown = 0
	} else {
		tx.IsUnknown = 1
	}

	m.modelData = append(m.modelData, tx)
	m.idCounter++
	m.evalIsEmpty()
	m.EndInsertRows()
}

func (m *TxListModel) evalIsEmpty() {
	m.SetIsEmpty(len(m.modelData) == 0)
}

func (m *TxListModel) remove(id int) {
	items := m.modelData
	//m.Clear()
	for i, tx := range items {
		if tx.ID == id {
			//m.Add(tx)
			m.BeginRemoveRows(core.NewQModelIndex(), i, i)
			m.modelData = append(m.modelData[:i], m.modelData[i+1:]...)
			m.EndRemoveRows()
		}
	}
	m.evalIsEmpty()
}

// Context Object for the view
type TxListCtx struct {
	core.QObject

	_ func()      `constructor:"init"`
	_ func(b int) `signal:"clicked,auto"`

	_ string `property:"shortenAddress"`
	_ string `property:"selectedSrc"`

	selectedIndex int
	transactions  *TxListModel
	accounts      *TxListAccountsModel
	ClefUI        *ClefUI
}

func (c *TxListCtx) init() {
	c.transactions = NewTxListModel(nil)
	c.accounts = NewTxListAccountsModel(nil)
}

func (c *TxListCtx) clicked(index int) {
	request := c.transactions.modelData[index]
	c.ClefUI.operationCh <- request.RPC

}

func NewTxListUI(clefUi *ClefUI) *TxListUI {
	c := NewTxListCtx(nil)
	c.ClefUI = clefUi

	widget := quick.NewQQuickWidget(nil)
	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.RootContext().SetContextProperty("transactions", c.transactions)
	widget.RootContext().SetContextProperty("accounts", c.accounts)

	widget.SetSource(core.NewQUrl3("qrc:/qml/tx_list.qml", 0))
	widget.SetStyleSheet("margin: 0;")
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return &TxListUI{
		UI:        widget,
		CtxObject: c,
	}

}
