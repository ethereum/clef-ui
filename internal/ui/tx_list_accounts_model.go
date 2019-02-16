package ui

import (
	"github.com/therecipe/qt/core"
)

func init() { CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "TxListAccountsModel") }

const (
	Text = int(core.Qt__UserRole) + 1<<iota
)

type TxListAccountsModel struct {
	core.QAbstractListModel

	_ func()          `constructor:"init"`
	_ func(tx string) `signal:"add,auto"`

	modelData []string
}

func (m *TxListAccountsModel) init() {
	m.modelData = []string{}

	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *TxListAccountsModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *TxListAccountsModel) columnCount(*core.QModelIndex) int {
	return 1
}

func (m *TxListAccountsModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		Text: core.NewQByteArray2("text", -1),
	}
}

func (m *TxListAccountsModel) data(index *core.QModelIndex, role int) *core.QVariant {
	row := index.Row()
	item := m.modelData[row]

	if row == 0 {
		return core.NewQVariant14(item)
	}

	if role == int(Text) {
		return core.NewQVariant14(item)
	}

	return core.NewQVariant()
}

func (m *TxListAccountsModel) add(address string) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, address)
	m.EndInsertRows()
}
