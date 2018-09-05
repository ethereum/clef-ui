package ui

import (
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

	_ func(b int) `signal:"clicked,auto"`

	answer 		int
}

func (t *ApproveListingCtx) clicked(b int) {
	log.Println(b)
	t.answer = b
}

func (t *ApproveListingCtx) Reset() {
	t.answer = 0
}

func (t *ApproveListingCtx) ClickResponse(res chan map[string]string) {
	go func() {
		done := false
		for !done {
			if t.answer != 0 {
				done = true
				if t.answer == 1 {
					res <- map[string]string{
						"approved": "false",
						"password": "",
					}
				} else if t.answer == 2 {
					res <- map[string]string{
						"approved": "true",
						"password": "asdfasdf",
					}
				}
				t.Reset()
			}
		}
	}()
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
	_ func(address string, selected bool)        `signal:"add,auto"`

	modelData []AccountListItem
}

func (m *CustomListModel) init() {
	m.modelData = []AccountListItem{}

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
		return core.NewQVariant14(item.address)
	}

	if role == int(Checked) {
		return core.NewQVariant11(item.selected)
	}

	return core.NewQVariant()
}

func (m*CustomListModel) clear() {
	m.BeginResetModel()
	m.modelData = []AccountListItem{}
	m.EndResetModel()
}

func (m *CustomListModel) add(address string, selected bool) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, AccountListItem{address, selected})
	m.EndInsertRows()
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

