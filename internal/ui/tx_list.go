package ui

import (
	"bytes"
	"encoding/json"
	"github.com/kyokan/clef-ui/internal/identicon"
	"github.com/kyokan/clef-ui/internal/utils"
	"strings"

	//"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"log"
	"net/http"
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

var KnownAccounts = make(map[string]string)

type TxListUI struct {
	UI 				*quick.QQuickWidget
	CtxObject 		*TxListCtx
}

type TxListItem struct {
	From   		string
	Method 		string
	RPC    		interface{}
	IsUnknown   int
	OnRemove 	chan int
	ID 			int
}

func (item *TxListItem) Remove() {
	item.OnRemove <- item.ID
}

type TxListModel struct {
	core.QAbstractListModel

	_ bool 							`property:"isEmpty"`

	_ func() 						`constructor:"init"`
	_ func()                    	`signal:"clear,auto"`
	_ func(tx *TxListItem)    		`signal:"add,auto"`
	_ func(i int)					`signal:"remove,auto"`

	modelData 						[]*TxListItem
	idCounter 						int
	OnRemove 						chan int
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
	return 3
}

func (m *TxListModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		From: core.NewQByteArray2("from", -1),
		Method:  core.NewQByteArray2("method", -1),
		FromSrc: core.NewQByteArray2("fromSrc", -1),
		IsUnknown: core.NewQByteArray2("isUnknown", -1),
	}
}

func (m *TxListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]

	if role == int(From) {
		address, _ := clefutils.ToChecksumAddress(item.From)
		return core.NewQVariant14(address)
	}

	if role == int(Method) {
		return core.NewQVariant14(item.Method)
	}

	if role == int(FromSrc) {
		return core.NewQVariant14(identicon.ToBase64Img(item.From))
	}

	if role == int(IsUnknown) {
		return core.NewQVariant7(item.IsUnknown)
	}

	return core.NewQVariant()
}

func (m *TxListModel) clear() {
	m.BeginResetModel()
	m.modelData = []*TxListItem{}
	m.EndResetModel()
}

func (m *TxListModel) add(tx *TxListItem) {
	address := strings.ToLower(tx.From)
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	tx.ID = m.idCounter
	tx.OnRemove = m.OnRemove

	if address == " - " || KnownAccounts[address] == address {
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
	transactions := m.modelData
	m.SetIsEmpty(len(transactions) == 0)
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

	m.evalIsEmpty()
}

// Context Object for the view
type TxListCtx struct {
	core.QObject

	_ func() 		`constructor:"init"`
	_ func(b int) 	`signal:"clicked,auto"`

	_ string `property:"shortenAddress"`
	_ string `property:"selectedSrc"`

	selectedIndex 	int
	transactions 	*TxListModel
	accounts 		*TxListAccountsModel
	ClefUI 			*ClefUI
}

var shouldRequest = true
func (c *TxListCtx) RequestAccounts() {
	if !shouldRequest {
		return
	}
	shouldRequest = false
	url := "http://localhost:8550"

	var jsonStr = []byte(`{"jsonrpc":"2.0","method":"account_list","id":0}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	var data struct {
		Jsonrpc 	string `json:"jsonrpc"`
		Id 			int `json:"id"`
		Result 		[]string `json:"result"`
	}
	json.NewDecoder(resp.Body).Decode(&data)
	c.accounts = NewTxListAccountsModel(nil)
	for _, account := range data.Result {
		address := strings.ToLower(account)
		KnownAccounts[address] = address
	}

}

func (c *TxListCtx) init() {
	c.transactions = NewTxListModel(nil	)
	c.accounts = NewTxListAccountsModel(nil)
	//c.SetShortenAddress(ALL_ACCOUNTS)
	//c.accounts.Add(ALL_ACCOUNTS)
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
	widget.RootContext().SetContextProperty(	"accounts", c.accounts)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}
