package ui

import (
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"log"
)

type ApproveListingUI struct {
	UI 					*quick.QQuickWidget
	ContextObject		*ApproveListingCtx
	AccountListModel 	*CustomListModel
}

type ApproveListingCtx struct {
	core.QObject

	_ string `property:"remote"`
	_ string `property:"transport"`
	_ string `property:"endpoint"`
	_ string `property:"from"`
	_ string `property:"message"`
	_ string `property:"rawData"`
	_ string `property:"hash"`

	answer 		int
}

func init() {CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomListModel")}

const (
	Address = int(core.Qt__UserRole) + 1<<iota
	Checked
)

type AccountListItem struct {
	address string
	selected bool
}

type CustomListModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ func()                                 	`signal:"clear,auto"`
	_ func(account params.ApproveListingAccount)        `signal:"add,auto"`

	modelData []params.ApproveListingAccount
	answer 	int
	_ func(b int) `signal:"clicked,auto"`

}

func (m *CustomListModel) init() {
	m.modelData = []params.ApproveListingAccount{}

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
		Checked:  core.NewQByteArray2("selected", -1),
	}
}

func (m *CustomListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]

	if role == int(Address) {
		return core.NewQVariant14(item.Address)
	}

	if role == int(Checked) {
		return core.NewQVariant11(true)
	}

	return core.NewQVariant()
}

func (m*CustomListModel) clear() {
	m.BeginResetModel()
	m.modelData = []params.ApproveListingAccount{}
	m.EndResetModel()
}

func (m *CustomListModel) add(account params.ApproveListingAccount) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, account)
	m.EndInsertRows()
}

func (t *CustomListModel) clicked(b int) {
	log.Println(b)
	t.answer = b
}

func (t *CustomListModel) Reset() {
	t.answer = 0
}

func (t *CustomListModel) ClickResponse(res chan []params.ApproveListingAccount) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				if t.answer == 1 {
					res <- []params.ApproveListingAccount{}
				} else if t.answer == 2 {
					res <- t.modelData
				}
				t.Reset()
			}
		}
	}()
}


func NewApproveListingUI() *ApproveListingUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_listing.qml", 0))
	c := NewApproveListingCtx(nil)
	m := NewCustomListModel(nil)
	v := &ApproveListingUI{
		UI: widget,
		ContextObject: c,
		AccountListModel: m,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.RootContext().SetContextProperty(	"myModel", m)
	widget.SetResizeMode(quick.QQuickWidget__SizeViewToRootObject)
	widget.Hide()
	return v
}

