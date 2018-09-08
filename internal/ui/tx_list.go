package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"log"
)

func init() {CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "TxListModel")}

const (
	From = int(core.Qt__UserRole) + 1<<iota
	Method
)

type TxListUI struct {
	UI 				*quick.QQuickWidget
	CtxObject 		*TxListCtx
}

type TxListItem struct {
	From   		string
	Method 		string
	RPC    		interface{}
	OnRemove 	chan int
	ID 			int
}

func (item *TxListItem) Remove() {
	item.OnRemove <- item.ID
}

type TxListModel struct {
	core.QAbstractListModel

	_ func() 					`constructor:"init"`
	_ func()                    `signal:"clear,auto"`
	_ func(tx *TxListItem)    	`signal:"add,auto"`
	_ func(i int)    			`signal:"remove,auto"`

	modelData 					[]*TxListItem
	idCounter 					int
	OnRemove 					chan int
}

func (m *TxListModel) init() {
	m.modelData = []*TxListItem{}
	m.idCounter = 0
	m.OnRemove = make(chan int)

	go func() {
		for {
			index := <- m.OnRemove
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
	return 2
}

func (m *TxListModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		From: core.NewQByteArray2("from", -1),
		Method:  core.NewQByteArray2("method", -1),
	}
}

func (m *TxListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]

	if role == int(From) {
		return core.NewQVariant14(item.From)
	}

	if role == int(Method) {
		return core.NewQVariant14(item.Method)
	}

	return core.NewQVariant()
}

func (m *TxListModel) clear() {
	m.BeginResetModel()
	m.modelData = []*TxListItem{}
	m.EndResetModel()
}

func (m *TxListModel) add(tx *TxListItem) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	tx.ID = m.idCounter
	tx.OnRemove = m.OnRemove
	m.modelData = append(m.modelData, tx)
	m.idCounter++
	m.EndInsertRows()
}

func (m *TxListModel) remove(id int) {
	if len(m.modelData) == 0 {
		return
	}

	oldTxList := m.modelData

	m.Clear()

	for _, tx := range oldTxList {
		if tx.ID != id {
			m.Add(tx)
		}
	}

	log.Println(len(m.modelData))
}

type TxListCtx struct {
	core.QObject

	_ func() 		`constructor:"init"`
	_ func(b int) 	`signal:"clicked,auto"`

	transactions 	*TxListModel
	ClefUI 			*ClefUI
}

func (c *TxListCtx) init() {
	c.transactions = NewTxListModel(nil	)
}

func (c *TxListCtx) clicked(index int) {
	request := c.transactions.modelData[index]
	ui := c.ClefUI
	rpc := request.RPC
	switch request.Method {
	case "ApproveListing":
		rpc := rpc.(ApproveListingRequest)
		ui.ApproveListingRequest <- rpc
	case "ApproveSignData":
		rpc := rpc.(ApproveSignDataRequest)
		ui.ApproveSignDataRequest <- rpc
	case "ApproveNewAccount":
		rpc := rpc.(ApproveNewAccountRequest)
		ui.ApproveNewAccountRequest <- rpc
	case "ApproveTx":
		rpc := rpc.(ApproveTxRequest)
		ui.ApproveTxRequest <- rpc
	case "ApproveImport":
		rpc := rpc.(ApproveImportRequest)
		ui.ApproveImportRequest <- rpc
	case "ApproveExport":
		rpc := rpc.(ApproveExportRequest)
		ui.ApproveExportRequest <- rpc
	}

}

func NewTxListUI(clefUi *ClefUI) *TxListUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/tx_list.qml", 0))
	c := NewTxListCtx(nil)
	c.ClefUI = clefUi
	v := &TxListUI{
		UI: widget,
		CtxObject: c,
	}
	widget.SetStyleSheet("margin: 0;")
	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.RootContext().SetContextProperty(	"transactions", c.transactions)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}
