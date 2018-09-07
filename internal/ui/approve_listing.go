package ui

import (
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
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

	_ func() 										`constructor:"init"`
	_ func()                                 		`signal:"clear,auto"`
	_ func(account params.ApproveListingAccount)    `signal:"add,auto"`
	_ func(b int) 									`signal:"clicked,auto"`
	_ func(i int, checked bool) 					`signal:"onCheckStateChanged,auto"`

	modelData 										[]params.ApproveListingAccount
	checkState										map[int]bool
	answer 											int

}

func (m *CustomListModel) init() {
	m.modelData = []params.ApproveListingAccount{}
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
	m.checkState = map[int]bool{}
	m.answer = 0
	m.EndResetModel()
}

func (m *CustomListModel) add(account params.ApproveListingAccount) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, account)
	m.EndInsertRows()
}

func (t *CustomListModel) clicked(b int) {
	t.answer = b
}

func (t *CustomListModel) onCheckStateChanged(i int, checked bool) {
	t.checkState[i] = checked
}

func (t *CustomListModel) ClickResponse(reply *params.ApproveListingResponse, res chan bool) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				if t.answer == 1 {
					res <- true
				} else if t.answer == 2 {
					accounts := make([]params.ApproveListingAccount, 0)

					for i, account := range t.modelData {
						if t.checkState[i] {
							accounts = append(accounts, account)
						}
					}
					reply.Accounts = accounts
					res <- true
				}
				t.Clear()
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
	widget.SetStyleSheet("margin: 0;")
	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.RootContext().SetContextProperty(	"myModel", m)
	widget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)
	widget.Hide()
	return v
}

