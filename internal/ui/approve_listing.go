package ui

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

type ApproveListingUI struct {
	UI 					*quick.QQuickWidget
	ContextObject		*CtxObject
}

func init() { CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomListModel") }

type ListItem struct {
	Address string
	checked string
}

type CustomListModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ func()                                  `signal:"remove,auto"`
	_ func(obj []*core.QVariant)              `signal:"add,auto"`
	_ func(address string, checked string) `signal:"edit,auto"`

	modelData []ListItem
}

func (m *CustomListModel) init() {
	m.modelData = []ListItem{{"john", "doe"}, {"john", "bob"}}

	m.ConnectRowCount(m.rowCount)
	m.ConnectData(m.data)
}

func (m *CustomListModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := m.modelData[index.Row()]
	return core.NewQVariant14(fmt.Sprintf("%v %v", item.Address, item.checked))
}

func (m *CustomListModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomListModel) add(item []*core.QVariant) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, ListItem{item[0].ToString(), item[1].ToString()})
	m.EndInsertRows()
}

func (m *CustomListModel) edit(address string, checked string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[len(m.modelData)-1] = ListItem{address, checked}
	m.DataChanged(m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}

func NewApproveListingUI() *ApproveListingUI {
	widget := quick.NewQQuickWidget(nil)
	widget.SetSource(core.NewQUrl3("qrc:/qml/approve_listing.qml", 0))
	c := NewCtxObject(nil)
	v := &ApproveListingUI{
		UI: widget,
		ContextObject: c,
	}

	widget.RootContext().SetContextProperty("ctxObject", c)
	widget.Show()
	return v
}

