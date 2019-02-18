package ui

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	custom_accounts_0de74dm "github.com/ethereum/go-ethereum/accounts"
	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type TxListModel_ITF interface {
	std_core.QAbstractListModel_ITF
	TxListModel_PTR() *TxListModel
}

func (ptr *TxListModel) TxListModel_PTR() *TxListModel {
	return ptr
}

func (ptr *TxListModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *TxListModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromTxListModel(ptr TxListModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TxListModel_PTR().Pointer()
	}
	return nil
}

func NewTxListModelFromPointer(ptr unsafe.Pointer) (n *TxListModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TxListModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TxListModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &TxListModel{QAbstractListModel: *deduced}

		default:
			n = new(TxListModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackTxListModel77da62_Constructor
func callbackTxListModel77da62_Constructor(ptr unsafe.Pointer) {
	this := NewTxListModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClear(this.clear)
	this.ConnectAdd(this.add)
	this.ConnectRemove(this.remove)
	this.init()
}

//export callbackTxListModel77da62_Clear
func callbackTxListModel77da62_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		signal.(func())()
	}

}

func (ptr *TxListModel) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clear") {
			C.TxListModel77da62_ConnectClear(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clear", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", f)
		}
	}
}

func (ptr *TxListModel) DisconnectClear() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_DisconnectClear(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *TxListModel) Clear() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_Clear(ptr.Pointer())
	}
}

//export callbackTxListModel77da62_Add
func callbackTxListModel77da62_Add(ptr unsafe.Pointer, tx C.uintptr_t) {
	var txD *IncomingRequestItem
	if txI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(tx))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(tx)))
		txD = txI.(*IncomingRequestItem)
	}
	if signal := qt.GetSignal(ptr, "add"); signal != nil {
		signal.(func(*IncomingRequestItem))(txD)
	}

}

func (ptr *TxListModel) ConnectAdd(f func(tx *IncomingRequestItem)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "add") {
			C.TxListModel77da62_ConnectAdd(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "add"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "add", func(tx *IncomingRequestItem) {
				signal.(func(*IncomingRequestItem))(tx)
				f(tx)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "add", f)
		}
	}
}

func (ptr *TxListModel) DisconnectAdd() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_DisconnectAdd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "add")
	}
}

func (ptr *TxListModel) Add(tx *IncomingRequestItem) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(tx), tx)
		C.TxListModel77da62_Add(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(tx))))
	}
}

//export callbackTxListModel77da62_Remove
func callbackTxListModel77da62_Remove(ptr unsafe.Pointer, i C.int) {
	if signal := qt.GetSignal(ptr, "remove"); signal != nil {
		signal.(func(int))(int(int32(i)))
	}

}

func (ptr *TxListModel) ConnectRemove(f func(i int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remove") {
			C.TxListModel77da62_ConnectRemove(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "remove"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remove", func(i int) {
				signal.(func(int))(i)
				f(i)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remove", f)
		}
	}
}

func (ptr *TxListModel) DisconnectRemove() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_DisconnectRemove(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remove")
	}
}

func (ptr *TxListModel) Remove(i int) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_Remove(ptr.Pointer(), C.int(int32(i)))
	}
}

//export callbackTxListModel77da62_IsEmpty
func callbackTxListModel77da62_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).IsEmptyDefault())))
}

func (ptr *TxListModel) ConnectIsEmpty(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isEmpty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", f)
		}
	}
}

func (ptr *TxListModel) DisconnectIsEmpty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isEmpty")
	}
}

func (ptr *TxListModel) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *TxListModel) IsEmptyDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_IsEmptyDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackTxListModel77da62_SetIsEmpty
func callbackTxListModel77da62_SetIsEmpty(ptr unsafe.Pointer, isEmpty C.char) {
	if signal := qt.GetSignal(ptr, "setIsEmpty"); signal != nil {
		signal.(func(bool))(int8(isEmpty) != 0)
	} else {
		NewTxListModelFromPointer(ptr).SetIsEmptyDefault(int8(isEmpty) != 0)
	}
}

func (ptr *TxListModel) ConnectSetIsEmpty(f func(isEmpty bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setIsEmpty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setIsEmpty", func(isEmpty bool) {
				signal.(func(bool))(isEmpty)
				f(isEmpty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setIsEmpty", f)
		}
	}
}

func (ptr *TxListModel) DisconnectSetIsEmpty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setIsEmpty")
	}
}

func (ptr *TxListModel) SetIsEmpty(isEmpty bool) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_SetIsEmpty(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isEmpty))))
	}
}

func (ptr *TxListModel) SetIsEmptyDefault(isEmpty bool) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_SetIsEmptyDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isEmpty))))
	}
}

//export callbackTxListModel77da62_IsEmptyChanged
func callbackTxListModel77da62_IsEmptyChanged(ptr unsafe.Pointer, isEmpty C.char) {
	if signal := qt.GetSignal(ptr, "isEmptyChanged"); signal != nil {
		signal.(func(bool))(int8(isEmpty) != 0)
	}

}

func (ptr *TxListModel) ConnectIsEmptyChanged(f func(isEmpty bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "isEmptyChanged") {
			C.TxListModel77da62_ConnectIsEmptyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "isEmptyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isEmptyChanged", func(isEmpty bool) {
				signal.(func(bool))(isEmpty)
				f(isEmpty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isEmptyChanged", f)
		}
	}
}

func (ptr *TxListModel) DisconnectIsEmptyChanged() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_DisconnectIsEmptyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "isEmptyChanged")
	}
}

func (ptr *TxListModel) IsEmptyChanged(isEmpty bool) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_IsEmptyChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isEmpty))))
	}
}

func TxListModel_QRegisterMetaType() int {
	return int(int32(C.TxListModel77da62_TxListModel77da62_QRegisterMetaType()))
}

func (ptr *TxListModel) QRegisterMetaType() int {
	return int(int32(C.TxListModel77da62_TxListModel77da62_QRegisterMetaType()))
}

func TxListModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TxListModel77da62_TxListModel77da62_QRegisterMetaType2(typeNameC)))
}

func (ptr *TxListModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TxListModel77da62_TxListModel77da62_QRegisterMetaType2(typeNameC)))
}

func TxListModel_QmlRegisterType() int {
	return int(int32(C.TxListModel77da62_TxListModel77da62_QmlRegisterType()))
}

func (ptr *TxListModel) QmlRegisterType() int {
	return int(int32(C.TxListModel77da62_TxListModel77da62_QmlRegisterType()))
}

func TxListModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TxListModel77da62_TxListModel77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TxListModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TxListModel77da62_TxListModel77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TxListModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListModel77da62_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.TxListModel77da62_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *TxListModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListModel77da62_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.TxListModel77da62_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TxListModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListModel77da62_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.TxListModel77da62_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *TxListModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TxListModel77da62___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TxListModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.TxListModel77da62___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *TxListModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.TxListModel77da62___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TxListModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListModel77da62___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.TxListModel77da62___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *TxListModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListModel77da62___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.TxListModel77da62___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *TxListModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListModel77da62___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.TxListModel77da62___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *TxListModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.TxListModel77da62___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TxListModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.TxListModel77da62___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *TxListModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.TxListModel77da62___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TxListModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.TxListModel77da62___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *TxListModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TxListModel77da62___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TxListModel) __roleNames_newList() unsafe.Pointer {
	return C.TxListModel77da62___roleNames_newList(ptr.Pointer())
}

func (ptr *TxListModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.TxListModel77da62___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TxListModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TxListModel77da62___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TxListModel) __itemData_newList() unsafe.Pointer {
	return C.TxListModel77da62___itemData_newList(ptr.Pointer())
}

func (ptr *TxListModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.TxListModel77da62___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TxListModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListModel77da62___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.TxListModel77da62___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *TxListModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListModel77da62___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListModel) __match_newList() unsafe.Pointer {
	return C.TxListModel77da62___match_newList(ptr.Pointer())
}

func (ptr *TxListModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListModel77da62___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.TxListModel77da62___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *TxListModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListModel77da62_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.TxListModel77da62_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TxListModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListModel77da62_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.TxListModel77da62_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TxListModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TxListModel77da62___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TxListModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TxListModel77da62___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TxListModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListModel77da62___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListModel) __findChildren_newList2() unsafe.Pointer {
	return C.TxListModel77da62___findChildren_newList2(ptr.Pointer())
}

func (ptr *TxListModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListModel77da62___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListModel) __findChildren_newList3() unsafe.Pointer {
	return C.TxListModel77da62___findChildren_newList3(ptr.Pointer())
}

func (ptr *TxListModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListModel77da62___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListModel) __findChildren_newList() unsafe.Pointer {
	return C.TxListModel77da62___findChildren_newList(ptr.Pointer())
}

func (ptr *TxListModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListModel77da62___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListModel) __children_newList() unsafe.Pointer {
	return C.TxListModel77da62___children_newList(ptr.Pointer())
}

func NewTxListModel(parent std_core.QObject_ITF) *TxListModel {
	tmpValue := NewTxListModelFromPointer(C.TxListModel77da62_NewTxListModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTxListModel77da62_DestroyTxListModel
func callbackTxListModel77da62_DestroyTxListModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~TxListModel"); signal != nil {
		signal.(func())()
	} else {
		NewTxListModelFromPointer(ptr).DestroyTxListModelDefault()
	}
}

func (ptr *TxListModel) ConnectDestroyTxListModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~TxListModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~TxListModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~TxListModel", f)
		}
	}
}

func (ptr *TxListModel) DisconnectDestroyTxListModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~TxListModel")
	}
}

func (ptr *TxListModel) DestroyTxListModel() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_DestroyTxListModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TxListModel) DestroyTxListModelDefault() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_DestroyTxListModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTxListModel77da62_DropMimeData
func callbackTxListModel77da62_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListModel77da62_Index
func callbackTxListModel77da62_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewTxListModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *TxListModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListModel77da62_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTxListModel77da62_Sibling
func callbackTxListModel77da62_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewTxListModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *TxListModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListModel77da62_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTxListModel77da62_Flags
func callbackTxListModel77da62_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewTxListModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TxListModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.TxListModel77da62_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackTxListModel77da62_InsertColumns
func callbackTxListModel77da62_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListModel77da62_InsertRows
func callbackTxListModel77da62_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListModel77da62_MoveColumns
func callbackTxListModel77da62_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TxListModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackTxListModel77da62_MoveRows
func callbackTxListModel77da62_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TxListModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackTxListModel77da62_RemoveColumns
func callbackTxListModel77da62_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListModel77da62_RemoveRows
func callbackTxListModel77da62_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListModel77da62_SetData
func callbackTxListModel77da62_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TxListModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackTxListModel77da62_SetHeaderData
func callbackTxListModel77da62_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TxListModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackTxListModel77da62_SetItemData
func callbackTxListModel77da62_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewTxListModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *TxListModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewTxListModelFromPointer(NewTxListModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackTxListModel77da62_Submit
func callbackTxListModel77da62_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *TxListModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackTxListModel77da62_ColumnsAboutToBeInserted
func callbackTxListModel77da62_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListModel77da62_ColumnsAboutToBeMoved
func callbackTxListModel77da62_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackTxListModel77da62_ColumnsAboutToBeRemoved
func callbackTxListModel77da62_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListModel77da62_ColumnsInserted
func callbackTxListModel77da62_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListModel77da62_ColumnsMoved
func callbackTxListModel77da62_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackTxListModel77da62_ColumnsRemoved
func callbackTxListModel77da62_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListModel77da62_DataChanged
func callbackTxListModel77da62_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackTxListModel77da62_FetchMore
func callbackTxListModel77da62_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewTxListModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *TxListModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackTxListModel77da62_HeaderDataChanged
func callbackTxListModel77da62_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListModel77da62_LayoutAboutToBeChanged
func callbackTxListModel77da62_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackTxListModel77da62_LayoutChanged
func callbackTxListModel77da62_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackTxListModel77da62_ModelAboutToBeReset
func callbackTxListModel77da62_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackTxListModel77da62_ModelReset
func callbackTxListModel77da62_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackTxListModel77da62_ResetInternalData
func callbackTxListModel77da62_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewTxListModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *TxListModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackTxListModel77da62_Revert
func callbackTxListModel77da62_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewTxListModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *TxListModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_RevertDefault(ptr.Pointer())
	}
}

//export callbackTxListModel77da62_RowsAboutToBeInserted
func callbackTxListModel77da62_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackTxListModel77da62_RowsAboutToBeMoved
func callbackTxListModel77da62_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackTxListModel77da62_RowsAboutToBeRemoved
func callbackTxListModel77da62_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListModel77da62_RowsInserted
func callbackTxListModel77da62_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListModel77da62_RowsMoved
func callbackTxListModel77da62_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackTxListModel77da62_RowsRemoved
func callbackTxListModel77da62_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListModel77da62_Sort
func callbackTxListModel77da62_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewTxListModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *TxListModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackTxListModel77da62_RoleNames
func callbackTxListModel77da62_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTxListModelFromPointer(NewTxListModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTxListModelFromPointer(NewTxListModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewTxListModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TxListModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.TxListModel77da62_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackTxListModel77da62_ItemData
func callbackTxListModel77da62_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTxListModelFromPointer(NewTxListModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTxListModelFromPointer(NewTxListModelFromPointer(nil).__itemData_newList())
		for k, v := range NewTxListModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TxListModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.TxListModel77da62_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackTxListModel77da62_MimeData
func callbackTxListModel77da62_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewTxListModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewTxListModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *TxListModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.TxListModel77da62_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTxListModelFromPointer(NewTxListModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTxListModel77da62_Buddy
func callbackTxListModel77da62_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTxListModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TxListModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListModel77da62_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTxListModel77da62_Parent
func callbackTxListModel77da62_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTxListModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TxListModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListModel77da62_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTxListModel77da62_Match
func callbackTxListModel77da62_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTxListModelFromPointer(NewTxListModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTxListModelFromPointer(NewTxListModelFromPointer(nil).__match_newList())
		for _, v := range NewTxListModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TxListModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewTxListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.TxListModel77da62_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackTxListModel77da62_Span
func callbackTxListModel77da62_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewTxListModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TxListModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.TxListModel77da62_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackTxListModel77da62_MimeTypes
func callbackTxListModel77da62_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewTxListModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *TxListModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.TxListModel77da62_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackTxListModel77da62_Data
func callbackTxListModel77da62_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTxListModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *TxListModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TxListModel77da62_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTxListModel77da62_HeaderData
func callbackTxListModel77da62_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTxListModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *TxListModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TxListModel77da62_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTxListModel77da62_SupportedDragActions
func callbackTxListModel77da62_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewTxListModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *TxListModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TxListModel77da62_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTxListModel77da62_SupportedDropActions
func callbackTxListModel77da62_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewTxListModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *TxListModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TxListModel77da62_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTxListModel77da62_CanDropMimeData
func callbackTxListModel77da62_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListModel77da62_CanFetchMore
func callbackTxListModel77da62_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListModel77da62_HasChildren
func callbackTxListModel77da62_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListModel77da62_ColumnCount
func callbackTxListModel77da62_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTxListModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TxListModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListModel77da62_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTxListModel77da62_RowCount
func callbackTxListModel77da62_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTxListModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TxListModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListModel77da62_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTxListModel77da62_Event
func callbackTxListModel77da62_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TxListModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTxListModel77da62_EventFilter
func callbackTxListModel77da62_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TxListModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListModel77da62_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTxListModel77da62_ChildEvent
func callbackTxListModel77da62_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTxListModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TxListModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTxListModel77da62_ConnectNotify
func callbackTxListModel77da62_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTxListModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TxListModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTxListModel77da62_CustomEvent
func callbackTxListModel77da62_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewTxListModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TxListModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTxListModel77da62_DeleteLater
func callbackTxListModel77da62_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewTxListModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TxListModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTxListModel77da62_Destroyed
func callbackTxListModel77da62_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackTxListModel77da62_DisconnectNotify
func callbackTxListModel77da62_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTxListModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TxListModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTxListModel77da62_ObjectNameChanged
func callbackTxListModel77da62_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackTxListModel77da62_TimerEvent
func callbackTxListModel77da62_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTxListModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TxListModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TxListModel77da62_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type ApproveListingCtx_ITF interface {
	std_core.QObject_ITF
	ApproveListingCtx_PTR() *ApproveListingCtx
}

func (ptr *ApproveListingCtx) ApproveListingCtx_PTR() *ApproveListingCtx {
	return ptr
}

func (ptr *ApproveListingCtx) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *ApproveListingCtx) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromApproveListingCtx(ptr ApproveListingCtx_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ApproveListingCtx_PTR().Pointer()
	}
	return nil
}

func NewApproveListingCtxFromPointer(ptr unsafe.Pointer) (n *ApproveListingCtx) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ApproveListingCtx)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ApproveListingCtx:
			n = deduced

		case *std_core.QObject:
			n = &ApproveListingCtx{QObject: *deduced}

		default:
			n = new(ApproveListingCtx)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackApproveListingCtx77da62_Constructor
func callbackApproveListingCtx77da62_Constructor(ptr unsafe.Pointer) {
	this := NewApproveListingCtxFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectBack(this.back)
	this.ConnectApprove(this.approve)
	this.ConnectReject(this.reject)
	this.ConnectOnCheckStateChanged(this.onCheckStateChanged)
	this.init()
}

//export callbackApproveListingCtx77da62_Back
func callbackApproveListingCtx77da62_Back(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "back"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveListingCtx) ConnectBack(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "back") {
			C.ApproveListingCtx77da62_ConnectBack(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "back"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "back", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "back", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectBack() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectBack(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "back")
	}
}

func (ptr *ApproveListingCtx) Back() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_Back(ptr.Pointer())
	}
}

//export callbackApproveListingCtx77da62_Approve
func callbackApproveListingCtx77da62_Approve(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "approve"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveListingCtx) ConnectApprove(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "approve") {
			C.ApproveListingCtx77da62_ConnectApprove(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "approve"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "approve", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "approve", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectApprove() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectApprove(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "approve")
	}
}

func (ptr *ApproveListingCtx) Approve() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_Approve(ptr.Pointer())
	}
}

//export callbackApproveListingCtx77da62_Reject
func callbackApproveListingCtx77da62_Reject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reject"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveListingCtx) ConnectReject(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "reject") {
			C.ApproveListingCtx77da62_ConnectReject(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "reject"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "reject", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reject", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectReject() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectReject(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "reject")
	}
}

func (ptr *ApproveListingCtx) Reject() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_Reject(ptr.Pointer())
	}
}

//export callbackApproveListingCtx77da62_OnCheckStateChanged
func callbackApproveListingCtx77da62_OnCheckStateChanged(ptr unsafe.Pointer, i C.int, checked C.char) {
	if signal := qt.GetSignal(ptr, "onCheckStateChanged"); signal != nil {
		signal.(func(int, bool))(int(int32(i)), int8(checked) != 0)
	}

}

func (ptr *ApproveListingCtx) ConnectOnCheckStateChanged(f func(i int, checked bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onCheckStateChanged") {
			C.ApproveListingCtx77da62_ConnectOnCheckStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "onCheckStateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onCheckStateChanged", func(i int, checked bool) {
				signal.(func(int, bool))(i, checked)
				f(i, checked)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onCheckStateChanged", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectOnCheckStateChanged() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectOnCheckStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onCheckStateChanged")
	}
}

func (ptr *ApproveListingCtx) OnCheckStateChanged(i int, checked bool) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_OnCheckStateChanged(ptr.Pointer(), C.int(int32(i)), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

//export callbackApproveListingCtx77da62_TriggerUpdate
func callbackApproveListingCtx77da62_TriggerUpdate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "triggerUpdate"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveListingCtx) ConnectTriggerUpdate(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "triggerUpdate") {
			C.ApproveListingCtx77da62_ConnectTriggerUpdate(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "triggerUpdate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "triggerUpdate", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "triggerUpdate", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectTriggerUpdate() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectTriggerUpdate(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "triggerUpdate")
	}
}

func (ptr *ApproveListingCtx) TriggerUpdate() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_TriggerUpdate(ptr.Pointer())
	}
}

//export callbackApproveListingCtx77da62_Remote
func callbackApproveListingCtx77da62_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "remote"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveListingCtxFromPointer(ptr).RemoteDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveListingCtx) ConnectRemote(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "remote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remote", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remote", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "remote")
	}
}

func (ptr *ApproveListingCtx) Remote() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx77da62_SetRemote
func callbackApproveListingCtx77da62_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setRemote"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	} else {
		NewApproveListingCtxFromPointer(ptr).SetRemoteDefault(cGoUnpackString(remote))
	}
}

func (ptr *ApproveListingCtx) ConnectSetRemote(f func(remote string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRemote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectSetRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRemote")
	}
}

func (ptr *ApproveListingCtx) SetRemote(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveListingCtx77da62_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *ApproveListingCtx) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveListingCtx77da62_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveListingCtx77da62_RemoteChanged
func callbackApproveListingCtx77da62_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *ApproveListingCtx) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.ApproveListingCtx77da62_ConnectRemoteChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "remoteChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectRemoteChanged() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectRemoteChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remoteChanged")
	}
}

func (ptr *ApproveListingCtx) RemoteChanged(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveListingCtx77da62_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveListingCtx77da62_Transport
func callbackApproveListingCtx77da62_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "transport"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveListingCtxFromPointer(ptr).TransportDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveListingCtx) ConnectTransport(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "transport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transport", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transport", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transport")
	}
}

func (ptr *ApproveListingCtx) Transport() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx77da62_SetTransport
func callbackApproveListingCtx77da62_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTransport"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	} else {
		NewApproveListingCtxFromPointer(ptr).SetTransportDefault(cGoUnpackString(transport))
	}
}

func (ptr *ApproveListingCtx) ConnectSetTransport(f func(transport string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTransport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectSetTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransport")
	}
}

func (ptr *ApproveListingCtx) SetTransport(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveListingCtx77da62_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *ApproveListingCtx) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveListingCtx77da62_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveListingCtx77da62_TransportChanged
func callbackApproveListingCtx77da62_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *ApproveListingCtx) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.ApproveListingCtx77da62_ConnectTransportChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transportChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectTransportChanged() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectTransportChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transportChanged")
	}
}

func (ptr *ApproveListingCtx) TransportChanged(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveListingCtx77da62_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveListingCtx77da62_Endpoint
func callbackApproveListingCtx77da62_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "endpoint"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveListingCtxFromPointer(ptr).EndpointDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveListingCtx) ConnectEndpoint(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endpoint")
	}
}

func (ptr *ApproveListingCtx) Endpoint() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx77da62_SetEndpoint
func callbackApproveListingCtx77da62_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setEndpoint"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	} else {
		NewApproveListingCtxFromPointer(ptr).SetEndpointDefault(cGoUnpackString(endpoint))
	}
}

func (ptr *ApproveListingCtx) ConnectSetEndpoint(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEndpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectSetEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEndpoint")
	}
}

func (ptr *ApproveListingCtx) SetEndpoint(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveListingCtx77da62_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *ApproveListingCtx) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveListingCtx77da62_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveListingCtx77da62_EndpointChanged
func callbackApproveListingCtx77da62_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *ApproveListingCtx) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.ApproveListingCtx77da62_ConnectEndpointChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "endpointChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectEndpointChanged() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectEndpointChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "endpointChanged")
	}
}

func (ptr *ApproveListingCtx) EndpointChanged(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveListingCtx77da62_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveListingCtx77da62_From
func callbackApproveListingCtx77da62_From(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "from"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveListingCtxFromPointer(ptr).FromDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveListingCtx) ConnectFrom(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "from"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "from", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "from", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "from")
	}
}

func (ptr *ApproveListingCtx) From() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_From(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) FromDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_FromDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx77da62_SetFrom
func callbackApproveListingCtx77da62_SetFrom(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFrom"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	} else {
		NewApproveListingCtxFromPointer(ptr).SetFromDefault(cGoUnpackString(from))
	}
}

func (ptr *ApproveListingCtx) ConnectSetFrom(f func(from string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFrom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFrom", func(from string) {
				signal.(func(string))(from)
				f(from)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFrom", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectSetFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFrom")
	}
}

func (ptr *ApproveListingCtx) SetFrom(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveListingCtx77da62_SetFrom(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

func (ptr *ApproveListingCtx) SetFromDefault(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveListingCtx77da62_SetFromDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveListingCtx77da62_FromChanged
func callbackApproveListingCtx77da62_FromChanged(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	}

}

func (ptr *ApproveListingCtx) ConnectFromChanged(f func(from string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromChanged") {
			C.ApproveListingCtx77da62_ConnectFromChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fromChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromChanged", func(from string) {
				signal.(func(string))(from)
				f(from)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromChanged", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectFromChanged() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectFromChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fromChanged")
	}
}

func (ptr *ApproveListingCtx) FromChanged(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveListingCtx77da62_FromChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveListingCtx77da62_Message
func callbackApproveListingCtx77da62_Message(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "message"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveListingCtxFromPointer(ptr).MessageDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveListingCtx) ConnectMessage(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "message"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "message", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "message", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "message")
	}
}

func (ptr *ApproveListingCtx) Message() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) MessageDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_MessageDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx77da62_SetMessage
func callbackApproveListingCtx77da62_SetMessage(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setMessage"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	} else {
		NewApproveListingCtxFromPointer(ptr).SetMessageDefault(cGoUnpackString(message))
	}
}

func (ptr *ApproveListingCtx) ConnectSetMessage(f func(message string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setMessage", func(message string) {
				signal.(func(string))(message)
				f(message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMessage", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectSetMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMessage")
	}
}

func (ptr *ApproveListingCtx) SetMessage(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.ApproveListingCtx77da62_SetMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func (ptr *ApproveListingCtx) SetMessageDefault(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.ApproveListingCtx77da62_SetMessageDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackApproveListingCtx77da62_MessageChanged
func callbackApproveListingCtx77da62_MessageChanged(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "messageChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	}

}

func (ptr *ApproveListingCtx) ConnectMessageChanged(f func(message string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "messageChanged") {
			C.ApproveListingCtx77da62_ConnectMessageChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "messageChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "messageChanged", func(message string) {
				signal.(func(string))(message)
				f(message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "messageChanged", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectMessageChanged() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "messageChanged")
	}
}

func (ptr *ApproveListingCtx) MessageChanged(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.ApproveListingCtx77da62_MessageChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackApproveListingCtx77da62_RawData
func callbackApproveListingCtx77da62_RawData(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "rawData"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveListingCtxFromPointer(ptr).RawDataDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveListingCtx) ConnectRawData(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rawData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rawData", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rawData", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectRawData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rawData")
	}
}

func (ptr *ApproveListingCtx) RawData() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_RawData(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) RawDataDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_RawDataDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx77da62_SetRawData
func callbackApproveListingCtx77da62_SetRawData(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setRawData"); signal != nil {
		signal.(func(string))(cGoUnpackString(rawData))
	} else {
		NewApproveListingCtxFromPointer(ptr).SetRawDataDefault(cGoUnpackString(rawData))
	}
}

func (ptr *ApproveListingCtx) ConnectSetRawData(f func(rawData string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRawData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRawData", func(rawData string) {
				signal.(func(string))(rawData)
				f(rawData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRawData", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectSetRawData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRawData")
	}
}

func (ptr *ApproveListingCtx) SetRawData(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.ApproveListingCtx77da62_SetRawData(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

func (ptr *ApproveListingCtx) SetRawDataDefault(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.ApproveListingCtx77da62_SetRawDataDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackApproveListingCtx77da62_RawDataChanged
func callbackApproveListingCtx77da62_RawDataChanged(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "rawDataChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(rawData))
	}

}

func (ptr *ApproveListingCtx) ConnectRawDataChanged(f func(rawData string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rawDataChanged") {
			C.ApproveListingCtx77da62_ConnectRawDataChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rawDataChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rawDataChanged", func(rawData string) {
				signal.(func(string))(rawData)
				f(rawData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rawDataChanged", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectRawDataChanged() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectRawDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rawDataChanged")
	}
}

func (ptr *ApproveListingCtx) RawDataChanged(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.ApproveListingCtx77da62_RawDataChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackApproveListingCtx77da62_Hash
func callbackApproveListingCtx77da62_Hash(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "hash"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveListingCtxFromPointer(ptr).HashDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveListingCtx) ConnectHash(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hash"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hash", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hash", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectHash() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hash")
	}
}

func (ptr *ApproveListingCtx) Hash() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_Hash(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) HashDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx77da62_HashDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx77da62_SetHash
func callbackApproveListingCtx77da62_SetHash(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setHash"); signal != nil {
		signal.(func(string))(cGoUnpackString(hash))
	} else {
		NewApproveListingCtxFromPointer(ptr).SetHashDefault(cGoUnpackString(hash))
	}
}

func (ptr *ApproveListingCtx) ConnectSetHash(f func(hash string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHash"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setHash", func(hash string) {
				signal.(func(string))(hash)
				f(hash)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHash", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectSetHash() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHash")
	}
}

func (ptr *ApproveListingCtx) SetHash(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.ApproveListingCtx77da62_SetHash(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

func (ptr *ApproveListingCtx) SetHashDefault(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.ApproveListingCtx77da62_SetHashDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

//export callbackApproveListingCtx77da62_HashChanged
func callbackApproveListingCtx77da62_HashChanged(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hashChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(hash))
	}

}

func (ptr *ApproveListingCtx) ConnectHashChanged(f func(hash string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hashChanged") {
			C.ApproveListingCtx77da62_ConnectHashChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hashChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hashChanged", func(hash string) {
				signal.(func(string))(hash)
				f(hash)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hashChanged", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectHashChanged() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectHashChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hashChanged")
	}
}

func (ptr *ApproveListingCtx) HashChanged(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.ApproveListingCtx77da62_HashChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

func ApproveListingCtx_QRegisterMetaType() int {
	return int(int32(C.ApproveListingCtx77da62_ApproveListingCtx77da62_QRegisterMetaType()))
}

func (ptr *ApproveListingCtx) QRegisterMetaType() int {
	return int(int32(C.ApproveListingCtx77da62_ApproveListingCtx77da62_QRegisterMetaType()))
}

func ApproveListingCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveListingCtx77da62_ApproveListingCtx77da62_QRegisterMetaType2(typeNameC)))
}

func (ptr *ApproveListingCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveListingCtx77da62_ApproveListingCtx77da62_QRegisterMetaType2(typeNameC)))
}

func ApproveListingCtx_QmlRegisterType() int {
	return int(int32(C.ApproveListingCtx77da62_ApproveListingCtx77da62_QmlRegisterType()))
}

func (ptr *ApproveListingCtx) QmlRegisterType() int {
	return int(int32(C.ApproveListingCtx77da62_ApproveListingCtx77da62_QmlRegisterType()))
}

func ApproveListingCtx_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApproveListingCtx77da62_ApproveListingCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveListingCtx) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApproveListingCtx77da62_ApproveListingCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveListingCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ApproveListingCtx77da62___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ApproveListingCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ApproveListingCtx77da62___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ApproveListingCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveListingCtx77da62___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveListingCtx) __findChildren_newList2() unsafe.Pointer {
	return C.ApproveListingCtx77da62___findChildren_newList2(ptr.Pointer())
}

func (ptr *ApproveListingCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveListingCtx77da62___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveListingCtx) __findChildren_newList3() unsafe.Pointer {
	return C.ApproveListingCtx77da62___findChildren_newList3(ptr.Pointer())
}

func (ptr *ApproveListingCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveListingCtx77da62___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveListingCtx) __findChildren_newList() unsafe.Pointer {
	return C.ApproveListingCtx77da62___findChildren_newList(ptr.Pointer())
}

func (ptr *ApproveListingCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveListingCtx77da62___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveListingCtx) __children_newList() unsafe.Pointer {
	return C.ApproveListingCtx77da62___children_newList(ptr.Pointer())
}

func NewApproveListingCtx(parent std_core.QObject_ITF) *ApproveListingCtx {
	tmpValue := NewApproveListingCtxFromPointer(C.ApproveListingCtx77da62_NewApproveListingCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApproveListingCtx77da62_DestroyApproveListingCtx
func callbackApproveListingCtx77da62_DestroyApproveListingCtx(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ApproveListingCtx"); signal != nil {
		signal.(func())()
	} else {
		NewApproveListingCtxFromPointer(ptr).DestroyApproveListingCtxDefault()
	}
}

func (ptr *ApproveListingCtx) ConnectDestroyApproveListingCtx(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ApproveListingCtx"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~ApproveListingCtx", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ApproveListingCtx", f)
		}
	}
}

func (ptr *ApproveListingCtx) DisconnectDestroyApproveListingCtx() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ApproveListingCtx")
	}
}

func (ptr *ApproveListingCtx) DestroyApproveListingCtx() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DestroyApproveListingCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ApproveListingCtx) DestroyApproveListingCtxDefault() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DestroyApproveListingCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveListingCtx77da62_Event
func callbackApproveListingCtx77da62_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveListingCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ApproveListingCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveListingCtx77da62_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackApproveListingCtx77da62_EventFilter
func callbackApproveListingCtx77da62_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveListingCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ApproveListingCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveListingCtx77da62_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackApproveListingCtx77da62_ChildEvent
func callbackApproveListingCtx77da62_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewApproveListingCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ApproveListingCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackApproveListingCtx77da62_ConnectNotify
func callbackApproveListingCtx77da62_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveListingCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveListingCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveListingCtx77da62_CustomEvent
func callbackApproveListingCtx77da62_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewApproveListingCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ApproveListingCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackApproveListingCtx77da62_DeleteLater
func callbackApproveListingCtx77da62_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewApproveListingCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ApproveListingCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveListingCtx77da62_Destroyed
func callbackApproveListingCtx77da62_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackApproveListingCtx77da62_DisconnectNotify
func callbackApproveListingCtx77da62_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveListingCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveListingCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveListingCtx77da62_ObjectNameChanged
func callbackApproveListingCtx77da62_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApproveListingCtx77da62_TimerEvent
func callbackApproveListingCtx77da62_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewApproveListingCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ApproveListingCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx77da62_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type ApproveNewAccountCtx_ITF interface {
	std_core.QObject_ITF
	ApproveNewAccountCtx_PTR() *ApproveNewAccountCtx
}

func (ptr *ApproveNewAccountCtx) ApproveNewAccountCtx_PTR() *ApproveNewAccountCtx {
	return ptr
}

func (ptr *ApproveNewAccountCtx) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromApproveNewAccountCtx(ptr ApproveNewAccountCtx_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ApproveNewAccountCtx_PTR().Pointer()
	}
	return nil
}

func NewApproveNewAccountCtxFromPointer(ptr unsafe.Pointer) (n *ApproveNewAccountCtx) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ApproveNewAccountCtx)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ApproveNewAccountCtx:
			n = deduced

		case *std_core.QObject:
			n = &ApproveNewAccountCtx{QObject: *deduced}

		default:
			n = new(ApproveNewAccountCtx)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackApproveNewAccountCtx77da62_Constructor
func callbackApproveNewAccountCtx77da62_Constructor(ptr unsafe.Pointer) {
	this := NewApproveNewAccountCtxFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClicked(this.clicked)
	this.ConnectBack(this.back)
	this.ConnectPasswordEdited(this.passwordEdited)
	this.ConnectConfirmPasswordEdited(this.confirmPasswordEdited)
	this.init()
}

//export callbackApproveNewAccountCtx77da62_Clicked
func callbackApproveNewAccountCtx77da62_Clicked(ptr unsafe.Pointer, b C.int) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(int))(int(int32(b)))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectClicked(f func(b int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.ApproveNewAccountCtx77da62_ConnectClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clicked", func(b int) {
				signal.(func(int))(b)
				f(b)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *ApproveNewAccountCtx) Clicked(b int) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_Clicked(ptr.Pointer(), C.int(int32(b)))
	}
}

//export callbackApproveNewAccountCtx77da62_Back
func callbackApproveNewAccountCtx77da62_Back(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "back"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveNewAccountCtx) ConnectBack(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "back") {
			C.ApproveNewAccountCtx77da62_ConnectBack(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "back"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "back", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "back", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectBack() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectBack(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "back")
	}
}

func (ptr *ApproveNewAccountCtx) Back() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_Back(ptr.Pointer())
	}
}

//export callbackApproveNewAccountCtx77da62_PasswordEdited
func callbackApproveNewAccountCtx77da62_PasswordEdited(ptr unsafe.Pointer, b C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "passwordEdited"); signal != nil {
		signal.(func(string))(cGoUnpackString(b))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectPasswordEdited(f func(b string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "passwordEdited") {
			C.ApproveNewAccountCtx77da62_ConnectPasswordEdited(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "passwordEdited"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "passwordEdited", func(b string) {
				signal.(func(string))(b)
				f(b)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "passwordEdited", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectPasswordEdited() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectPasswordEdited(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "passwordEdited")
	}
}

func (ptr *ApproveNewAccountCtx) PasswordEdited(b string) {
	if ptr.Pointer() != nil {
		var bC *C.char
		if b != "" {
			bC = C.CString(b)
			defer C.free(unsafe.Pointer(bC))
		}
		C.ApproveNewAccountCtx77da62_PasswordEdited(ptr.Pointer(), C.struct_Moc_PackedString{data: bC, len: C.longlong(len(b))})
	}
}

//export callbackApproveNewAccountCtx77da62_ConfirmPasswordEdited
func callbackApproveNewAccountCtx77da62_ConfirmPasswordEdited(ptr unsafe.Pointer, b C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "confirmPasswordEdited"); signal != nil {
		signal.(func(string))(cGoUnpackString(b))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectConfirmPasswordEdited(f func(b string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "confirmPasswordEdited") {
			C.ApproveNewAccountCtx77da62_ConnectConfirmPasswordEdited(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "confirmPasswordEdited"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "confirmPasswordEdited", func(b string) {
				signal.(func(string))(b)
				f(b)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "confirmPasswordEdited", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectConfirmPasswordEdited() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectConfirmPasswordEdited(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "confirmPasswordEdited")
	}
}

func (ptr *ApproveNewAccountCtx) ConfirmPasswordEdited(b string) {
	if ptr.Pointer() != nil {
		var bC *C.char
		if b != "" {
			bC = C.CString(b)
			defer C.free(unsafe.Pointer(bC))
		}
		C.ApproveNewAccountCtx77da62_ConfirmPasswordEdited(ptr.Pointer(), C.struct_Moc_PackedString{data: bC, len: C.longlong(len(b))})
	}
}

//export callbackApproveNewAccountCtx77da62_Remote
func callbackApproveNewAccountCtx77da62_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "remote"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveNewAccountCtxFromPointer(ptr).RemoteDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveNewAccountCtx) ConnectRemote(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "remote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remote", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remote", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "remote")
	}
}

func (ptr *ApproveNewAccountCtx) Remote() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx77da62_SetRemote
func callbackApproveNewAccountCtx77da62_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setRemote"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).SetRemoteDefault(cGoUnpackString(remote))
	}
}

func (ptr *ApproveNewAccountCtx) ConnectSetRemote(f func(remote string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRemote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectSetRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRemote")
	}
}

func (ptr *ApproveNewAccountCtx) SetRemote(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveNewAccountCtx77da62_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *ApproveNewAccountCtx) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveNewAccountCtx77da62_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveNewAccountCtx77da62_RemoteChanged
func callbackApproveNewAccountCtx77da62_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.ApproveNewAccountCtx77da62_ConnectRemoteChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "remoteChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectRemoteChanged() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectRemoteChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remoteChanged")
	}
}

func (ptr *ApproveNewAccountCtx) RemoteChanged(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveNewAccountCtx77da62_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveNewAccountCtx77da62_Transport
func callbackApproveNewAccountCtx77da62_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "transport"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveNewAccountCtxFromPointer(ptr).TransportDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveNewAccountCtx) ConnectTransport(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "transport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transport", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transport", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transport")
	}
}

func (ptr *ApproveNewAccountCtx) Transport() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx77da62_SetTransport
func callbackApproveNewAccountCtx77da62_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTransport"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).SetTransportDefault(cGoUnpackString(transport))
	}
}

func (ptr *ApproveNewAccountCtx) ConnectSetTransport(f func(transport string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTransport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectSetTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransport")
	}
}

func (ptr *ApproveNewAccountCtx) SetTransport(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveNewAccountCtx77da62_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *ApproveNewAccountCtx) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveNewAccountCtx77da62_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveNewAccountCtx77da62_TransportChanged
func callbackApproveNewAccountCtx77da62_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.ApproveNewAccountCtx77da62_ConnectTransportChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transportChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectTransportChanged() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectTransportChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transportChanged")
	}
}

func (ptr *ApproveNewAccountCtx) TransportChanged(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveNewAccountCtx77da62_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveNewAccountCtx77da62_Endpoint
func callbackApproveNewAccountCtx77da62_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "endpoint"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveNewAccountCtxFromPointer(ptr).EndpointDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveNewAccountCtx) ConnectEndpoint(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endpoint")
	}
}

func (ptr *ApproveNewAccountCtx) Endpoint() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx77da62_SetEndpoint
func callbackApproveNewAccountCtx77da62_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setEndpoint"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).SetEndpointDefault(cGoUnpackString(endpoint))
	}
}

func (ptr *ApproveNewAccountCtx) ConnectSetEndpoint(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEndpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectSetEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEndpoint")
	}
}

func (ptr *ApproveNewAccountCtx) SetEndpoint(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveNewAccountCtx77da62_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *ApproveNewAccountCtx) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveNewAccountCtx77da62_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveNewAccountCtx77da62_EndpointChanged
func callbackApproveNewAccountCtx77da62_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.ApproveNewAccountCtx77da62_ConnectEndpointChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "endpointChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectEndpointChanged() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectEndpointChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "endpointChanged")
	}
}

func (ptr *ApproveNewAccountCtx) EndpointChanged(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveNewAccountCtx77da62_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveNewAccountCtx77da62_Password
func callbackApproveNewAccountCtx77da62_Password(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "password"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveNewAccountCtxFromPointer(ptr).PasswordDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveNewAccountCtx) ConnectPassword(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "password"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "password", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "password", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectPassword() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "password")
	}
}

func (ptr *ApproveNewAccountCtx) Password() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) PasswordDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_PasswordDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx77da62_SetPassword
func callbackApproveNewAccountCtx77da62_SetPassword(ptr unsafe.Pointer, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setPassword"); signal != nil {
		signal.(func(string))(cGoUnpackString(password))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).SetPasswordDefault(cGoUnpackString(password))
	}
}

func (ptr *ApproveNewAccountCtx) ConnectSetPassword(f func(password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPassword"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPassword", func(password string) {
				signal.(func(string))(password)
				f(password)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPassword", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectSetPassword() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPassword")
	}
}

func (ptr *ApproveNewAccountCtx) SetPassword(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveNewAccountCtx77da62_SetPassword(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

func (ptr *ApproveNewAccountCtx) SetPasswordDefault(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveNewAccountCtx77da62_SetPasswordDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackApproveNewAccountCtx77da62_PasswordChanged
func callbackApproveNewAccountCtx77da62_PasswordChanged(ptr unsafe.Pointer, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "passwordChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(password))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectPasswordChanged(f func(password string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "passwordChanged") {
			C.ApproveNewAccountCtx77da62_ConnectPasswordChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "passwordChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "passwordChanged", func(password string) {
				signal.(func(string))(password)
				f(password)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "passwordChanged", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectPasswordChanged() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectPasswordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "passwordChanged")
	}
}

func (ptr *ApproveNewAccountCtx) PasswordChanged(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveNewAccountCtx77da62_PasswordChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackApproveNewAccountCtx77da62_ConfirmPassword
func callbackApproveNewAccountCtx77da62_ConfirmPassword(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "confirmPassword"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveNewAccountCtxFromPointer(ptr).ConfirmPasswordDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveNewAccountCtx) ConnectConfirmPassword(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "confirmPassword"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "confirmPassword", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "confirmPassword", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectConfirmPassword() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "confirmPassword")
	}
}

func (ptr *ApproveNewAccountCtx) ConfirmPassword() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_ConfirmPassword(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) ConfirmPasswordDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx77da62_ConfirmPasswordDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx77da62_SetConfirmPassword
func callbackApproveNewAccountCtx77da62_SetConfirmPassword(ptr unsafe.Pointer, confirmPassword C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setConfirmPassword"); signal != nil {
		signal.(func(string))(cGoUnpackString(confirmPassword))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).SetConfirmPasswordDefault(cGoUnpackString(confirmPassword))
	}
}

func (ptr *ApproveNewAccountCtx) ConnectSetConfirmPassword(f func(confirmPassword string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setConfirmPassword"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setConfirmPassword", func(confirmPassword string) {
				signal.(func(string))(confirmPassword)
				f(confirmPassword)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setConfirmPassword", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectSetConfirmPassword() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setConfirmPassword")
	}
}

func (ptr *ApproveNewAccountCtx) SetConfirmPassword(confirmPassword string) {
	if ptr.Pointer() != nil {
		var confirmPasswordC *C.char
		if confirmPassword != "" {
			confirmPasswordC = C.CString(confirmPassword)
			defer C.free(unsafe.Pointer(confirmPasswordC))
		}
		C.ApproveNewAccountCtx77da62_SetConfirmPassword(ptr.Pointer(), C.struct_Moc_PackedString{data: confirmPasswordC, len: C.longlong(len(confirmPassword))})
	}
}

func (ptr *ApproveNewAccountCtx) SetConfirmPasswordDefault(confirmPassword string) {
	if ptr.Pointer() != nil {
		var confirmPasswordC *C.char
		if confirmPassword != "" {
			confirmPasswordC = C.CString(confirmPassword)
			defer C.free(unsafe.Pointer(confirmPasswordC))
		}
		C.ApproveNewAccountCtx77da62_SetConfirmPasswordDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: confirmPasswordC, len: C.longlong(len(confirmPassword))})
	}
}

//export callbackApproveNewAccountCtx77da62_ConfirmPasswordChanged
func callbackApproveNewAccountCtx77da62_ConfirmPasswordChanged(ptr unsafe.Pointer, confirmPassword C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "confirmPasswordChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(confirmPassword))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectConfirmPasswordChanged(f func(confirmPassword string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "confirmPasswordChanged") {
			C.ApproveNewAccountCtx77da62_ConnectConfirmPasswordChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "confirmPasswordChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "confirmPasswordChanged", func(confirmPassword string) {
				signal.(func(string))(confirmPassword)
				f(confirmPassword)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "confirmPasswordChanged", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectConfirmPasswordChanged() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectConfirmPasswordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "confirmPasswordChanged")
	}
}

func (ptr *ApproveNewAccountCtx) ConfirmPasswordChanged(confirmPassword string) {
	if ptr.Pointer() != nil {
		var confirmPasswordC *C.char
		if confirmPassword != "" {
			confirmPasswordC = C.CString(confirmPassword)
			defer C.free(unsafe.Pointer(confirmPasswordC))
		}
		C.ApproveNewAccountCtx77da62_ConfirmPasswordChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: confirmPasswordC, len: C.longlong(len(confirmPassword))})
	}
}

func ApproveNewAccountCtx_QRegisterMetaType() int {
	return int(int32(C.ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QRegisterMetaType()))
}

func (ptr *ApproveNewAccountCtx) QRegisterMetaType() int {
	return int(int32(C.ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QRegisterMetaType()))
}

func ApproveNewAccountCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QRegisterMetaType2(typeNameC)))
}

func (ptr *ApproveNewAccountCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QRegisterMetaType2(typeNameC)))
}

func ApproveNewAccountCtx_QmlRegisterType() int {
	return int(int32(C.ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QmlRegisterType()))
}

func (ptr *ApproveNewAccountCtx) QmlRegisterType() int {
	return int(int32(C.ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QmlRegisterType()))
}

func ApproveNewAccountCtx_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveNewAccountCtx) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveNewAccountCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ApproveNewAccountCtx77da62___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ApproveNewAccountCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ApproveNewAccountCtx77da62___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ApproveNewAccountCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveNewAccountCtx77da62___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveNewAccountCtx) __findChildren_newList2() unsafe.Pointer {
	return C.ApproveNewAccountCtx77da62___findChildren_newList2(ptr.Pointer())
}

func (ptr *ApproveNewAccountCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveNewAccountCtx77da62___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveNewAccountCtx) __findChildren_newList3() unsafe.Pointer {
	return C.ApproveNewAccountCtx77da62___findChildren_newList3(ptr.Pointer())
}

func (ptr *ApproveNewAccountCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveNewAccountCtx77da62___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveNewAccountCtx) __findChildren_newList() unsafe.Pointer {
	return C.ApproveNewAccountCtx77da62___findChildren_newList(ptr.Pointer())
}

func (ptr *ApproveNewAccountCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveNewAccountCtx77da62___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveNewAccountCtx) __children_newList() unsafe.Pointer {
	return C.ApproveNewAccountCtx77da62___children_newList(ptr.Pointer())
}

func NewApproveNewAccountCtx(parent std_core.QObject_ITF) *ApproveNewAccountCtx {
	tmpValue := NewApproveNewAccountCtxFromPointer(C.ApproveNewAccountCtx77da62_NewApproveNewAccountCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApproveNewAccountCtx77da62_DestroyApproveNewAccountCtx
func callbackApproveNewAccountCtx77da62_DestroyApproveNewAccountCtx(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ApproveNewAccountCtx"); signal != nil {
		signal.(func())()
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).DestroyApproveNewAccountCtxDefault()
	}
}

func (ptr *ApproveNewAccountCtx) ConnectDestroyApproveNewAccountCtx(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ApproveNewAccountCtx"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~ApproveNewAccountCtx", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ApproveNewAccountCtx", f)
		}
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectDestroyApproveNewAccountCtx() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ApproveNewAccountCtx")
	}
}

func (ptr *ApproveNewAccountCtx) DestroyApproveNewAccountCtx() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DestroyApproveNewAccountCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ApproveNewAccountCtx) DestroyApproveNewAccountCtxDefault() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DestroyApproveNewAccountCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveNewAccountCtx77da62_Event
func callbackApproveNewAccountCtx77da62_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveNewAccountCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ApproveNewAccountCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveNewAccountCtx77da62_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackApproveNewAccountCtx77da62_EventFilter
func callbackApproveNewAccountCtx77da62_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveNewAccountCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ApproveNewAccountCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveNewAccountCtx77da62_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackApproveNewAccountCtx77da62_ChildEvent
func callbackApproveNewAccountCtx77da62_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ApproveNewAccountCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackApproveNewAccountCtx77da62_ConnectNotify
func callbackApproveNewAccountCtx77da62_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveNewAccountCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveNewAccountCtx77da62_CustomEvent
func callbackApproveNewAccountCtx77da62_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ApproveNewAccountCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackApproveNewAccountCtx77da62_DeleteLater
func callbackApproveNewAccountCtx77da62_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ApproveNewAccountCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveNewAccountCtx77da62_Destroyed
func callbackApproveNewAccountCtx77da62_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackApproveNewAccountCtx77da62_DisconnectNotify
func callbackApproveNewAccountCtx77da62_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveNewAccountCtx77da62_ObjectNameChanged
func callbackApproveNewAccountCtx77da62_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApproveNewAccountCtx77da62_TimerEvent
func callbackApproveNewAccountCtx77da62_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ApproveNewAccountCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx77da62_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type ApproveSignDataCtx_ITF interface {
	std_core.QObject_ITF
	ApproveSignDataCtx_PTR() *ApproveSignDataCtx
}

func (ptr *ApproveSignDataCtx) ApproveSignDataCtx_PTR() *ApproveSignDataCtx {
	return ptr
}

func (ptr *ApproveSignDataCtx) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *ApproveSignDataCtx) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromApproveSignDataCtx(ptr ApproveSignDataCtx_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ApproveSignDataCtx_PTR().Pointer()
	}
	return nil
}

func NewApproveSignDataCtxFromPointer(ptr unsafe.Pointer) (n *ApproveSignDataCtx) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ApproveSignDataCtx)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ApproveSignDataCtx:
			n = deduced

		case *std_core.QObject:
			n = &ApproveSignDataCtx{QObject: *deduced}

		default:
			n = new(ApproveSignDataCtx)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackApproveSignDataCtx77da62_Constructor
func callbackApproveSignDataCtx77da62_Constructor(ptr unsafe.Pointer) {
	this := NewApproveSignDataCtxFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClicked(this.clicked)
	this.ConnectOnBack(this.onBack)
	this.ConnectOnApprove(this.onApprove)
	this.ConnectOnReject(this.onReject)
	this.ConnectEdited(this.edited)
}

//export callbackApproveSignDataCtx77da62_Clicked
func callbackApproveSignDataCtx77da62_Clicked(ptr unsafe.Pointer, b C.int) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(int))(int(int32(b)))
	}

}

func (ptr *ApproveSignDataCtx) ConnectClicked(f func(b int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.ApproveSignDataCtx77da62_ConnectClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clicked", func(b int) {
				signal.(func(int))(b)
				f(b)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *ApproveSignDataCtx) Clicked(b int) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_Clicked(ptr.Pointer(), C.int(int32(b)))
	}
}

//export callbackApproveSignDataCtx77da62_OnBack
func callbackApproveSignDataCtx77da62_OnBack(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "onBack"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveSignDataCtx) ConnectOnBack(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onBack") {
			C.ApproveSignDataCtx77da62_ConnectOnBack(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "onBack"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onBack", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onBack", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectOnBack() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectOnBack(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onBack")
	}
}

func (ptr *ApproveSignDataCtx) OnBack() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_OnBack(ptr.Pointer())
	}
}

//export callbackApproveSignDataCtx77da62_OnApprove
func callbackApproveSignDataCtx77da62_OnApprove(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "onApprove"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveSignDataCtx) ConnectOnApprove(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onApprove") {
			C.ApproveSignDataCtx77da62_ConnectOnApprove(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "onApprove"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onApprove", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onApprove", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectOnApprove() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectOnApprove(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onApprove")
	}
}

func (ptr *ApproveSignDataCtx) OnApprove() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_OnApprove(ptr.Pointer())
	}
}

//export callbackApproveSignDataCtx77da62_OnReject
func callbackApproveSignDataCtx77da62_OnReject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "onReject"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveSignDataCtx) ConnectOnReject(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onReject") {
			C.ApproveSignDataCtx77da62_ConnectOnReject(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "onReject"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onReject", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onReject", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectOnReject() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectOnReject(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onReject")
	}
}

func (ptr *ApproveSignDataCtx) OnReject() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_OnReject(ptr.Pointer())
	}
}

//export callbackApproveSignDataCtx77da62_Edited
func callbackApproveSignDataCtx77da62_Edited(ptr unsafe.Pointer, b C.struct_Moc_PackedString, value C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "edited"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(b), cGoUnpackString(value))
	}

}

func (ptr *ApproveSignDataCtx) ConnectEdited(f func(b string, value string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "edited") {
			C.ApproveSignDataCtx77da62_ConnectEdited(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "edited"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "edited", func(b string, value string) {
				signal.(func(string, string))(b, value)
				f(b, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "edited", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectEdited() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectEdited(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "edited")
	}
}

func (ptr *ApproveSignDataCtx) Edited(b string, value string) {
	if ptr.Pointer() != nil {
		var bC *C.char
		if b != "" {
			bC = C.CString(b)
			defer C.free(unsafe.Pointer(bC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.ApproveSignDataCtx77da62_Edited(ptr.Pointer(), C.struct_Moc_PackedString{data: bC, len: C.longlong(len(b))}, C.struct_Moc_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

//export callbackApproveSignDataCtx77da62_Remote
func callbackApproveSignDataCtx77da62_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "remote"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveSignDataCtxFromPointer(ptr).RemoteDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveSignDataCtx) ConnectRemote(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "remote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remote", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remote", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "remote")
	}
}

func (ptr *ApproveSignDataCtx) Remote() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx77da62_SetRemote
func callbackApproveSignDataCtx77da62_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setRemote"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).SetRemoteDefault(cGoUnpackString(remote))
	}
}

func (ptr *ApproveSignDataCtx) ConnectSetRemote(f func(remote string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRemote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectSetRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRemote")
	}
}

func (ptr *ApproveSignDataCtx) SetRemote(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveSignDataCtx77da62_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *ApproveSignDataCtx) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveSignDataCtx77da62_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveSignDataCtx77da62_RemoteChanged
func callbackApproveSignDataCtx77da62_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *ApproveSignDataCtx) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.ApproveSignDataCtx77da62_ConnectRemoteChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "remoteChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectRemoteChanged() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectRemoteChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remoteChanged")
	}
}

func (ptr *ApproveSignDataCtx) RemoteChanged(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveSignDataCtx77da62_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveSignDataCtx77da62_Transport
func callbackApproveSignDataCtx77da62_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "transport"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveSignDataCtxFromPointer(ptr).TransportDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveSignDataCtx) ConnectTransport(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "transport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transport", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transport", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transport")
	}
}

func (ptr *ApproveSignDataCtx) Transport() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx77da62_SetTransport
func callbackApproveSignDataCtx77da62_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTransport"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).SetTransportDefault(cGoUnpackString(transport))
	}
}

func (ptr *ApproveSignDataCtx) ConnectSetTransport(f func(transport string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTransport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectSetTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransport")
	}
}

func (ptr *ApproveSignDataCtx) SetTransport(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveSignDataCtx77da62_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *ApproveSignDataCtx) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveSignDataCtx77da62_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveSignDataCtx77da62_TransportChanged
func callbackApproveSignDataCtx77da62_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *ApproveSignDataCtx) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.ApproveSignDataCtx77da62_ConnectTransportChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transportChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectTransportChanged() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectTransportChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transportChanged")
	}
}

func (ptr *ApproveSignDataCtx) TransportChanged(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveSignDataCtx77da62_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveSignDataCtx77da62_Endpoint
func callbackApproveSignDataCtx77da62_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "endpoint"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveSignDataCtxFromPointer(ptr).EndpointDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveSignDataCtx) ConnectEndpoint(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endpoint")
	}
}

func (ptr *ApproveSignDataCtx) Endpoint() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx77da62_SetEndpoint
func callbackApproveSignDataCtx77da62_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setEndpoint"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).SetEndpointDefault(cGoUnpackString(endpoint))
	}
}

func (ptr *ApproveSignDataCtx) ConnectSetEndpoint(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEndpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectSetEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEndpoint")
	}
}

func (ptr *ApproveSignDataCtx) SetEndpoint(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveSignDataCtx77da62_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *ApproveSignDataCtx) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveSignDataCtx77da62_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveSignDataCtx77da62_EndpointChanged
func callbackApproveSignDataCtx77da62_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *ApproveSignDataCtx) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.ApproveSignDataCtx77da62_ConnectEndpointChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "endpointChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectEndpointChanged() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectEndpointChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "endpointChanged")
	}
}

func (ptr *ApproveSignDataCtx) EndpointChanged(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveSignDataCtx77da62_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveSignDataCtx77da62_From
func callbackApproveSignDataCtx77da62_From(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "from"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveSignDataCtxFromPointer(ptr).FromDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveSignDataCtx) ConnectFrom(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "from"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "from", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "from", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "from")
	}
}

func (ptr *ApproveSignDataCtx) From() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_From(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) FromDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_FromDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx77da62_SetFrom
func callbackApproveSignDataCtx77da62_SetFrom(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFrom"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).SetFromDefault(cGoUnpackString(from))
	}
}

func (ptr *ApproveSignDataCtx) ConnectSetFrom(f func(from string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFrom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFrom", func(from string) {
				signal.(func(string))(from)
				f(from)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFrom", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectSetFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFrom")
	}
}

func (ptr *ApproveSignDataCtx) SetFrom(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveSignDataCtx77da62_SetFrom(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

func (ptr *ApproveSignDataCtx) SetFromDefault(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveSignDataCtx77da62_SetFromDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveSignDataCtx77da62_FromChanged
func callbackApproveSignDataCtx77da62_FromChanged(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	}

}

func (ptr *ApproveSignDataCtx) ConnectFromChanged(f func(from string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromChanged") {
			C.ApproveSignDataCtx77da62_ConnectFromChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fromChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromChanged", func(from string) {
				signal.(func(string))(from)
				f(from)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromChanged", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectFromChanged() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectFromChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fromChanged")
	}
}

func (ptr *ApproveSignDataCtx) FromChanged(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveSignDataCtx77da62_FromChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveSignDataCtx77da62_Message
func callbackApproveSignDataCtx77da62_Message(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "message"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveSignDataCtxFromPointer(ptr).MessageDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveSignDataCtx) ConnectMessage(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "message"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "message", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "message", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "message")
	}
}

func (ptr *ApproveSignDataCtx) Message() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) MessageDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_MessageDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx77da62_SetMessage
func callbackApproveSignDataCtx77da62_SetMessage(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setMessage"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).SetMessageDefault(cGoUnpackString(message))
	}
}

func (ptr *ApproveSignDataCtx) ConnectSetMessage(f func(message string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setMessage", func(message string) {
				signal.(func(string))(message)
				f(message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMessage", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectSetMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMessage")
	}
}

func (ptr *ApproveSignDataCtx) SetMessage(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.ApproveSignDataCtx77da62_SetMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func (ptr *ApproveSignDataCtx) SetMessageDefault(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.ApproveSignDataCtx77da62_SetMessageDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackApproveSignDataCtx77da62_MessageChanged
func callbackApproveSignDataCtx77da62_MessageChanged(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "messageChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	}

}

func (ptr *ApproveSignDataCtx) ConnectMessageChanged(f func(message string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "messageChanged") {
			C.ApproveSignDataCtx77da62_ConnectMessageChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "messageChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "messageChanged", func(message string) {
				signal.(func(string))(message)
				f(message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "messageChanged", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectMessageChanged() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "messageChanged")
	}
}

func (ptr *ApproveSignDataCtx) MessageChanged(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.ApproveSignDataCtx77da62_MessageChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackApproveSignDataCtx77da62_RawData
func callbackApproveSignDataCtx77da62_RawData(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "rawData"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveSignDataCtxFromPointer(ptr).RawDataDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveSignDataCtx) ConnectRawData(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rawData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rawData", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rawData", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectRawData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rawData")
	}
}

func (ptr *ApproveSignDataCtx) RawData() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_RawData(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) RawDataDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_RawDataDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx77da62_SetRawData
func callbackApproveSignDataCtx77da62_SetRawData(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setRawData"); signal != nil {
		signal.(func(string))(cGoUnpackString(rawData))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).SetRawDataDefault(cGoUnpackString(rawData))
	}
}

func (ptr *ApproveSignDataCtx) ConnectSetRawData(f func(rawData string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRawData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRawData", func(rawData string) {
				signal.(func(string))(rawData)
				f(rawData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRawData", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectSetRawData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRawData")
	}
}

func (ptr *ApproveSignDataCtx) SetRawData(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.ApproveSignDataCtx77da62_SetRawData(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

func (ptr *ApproveSignDataCtx) SetRawDataDefault(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.ApproveSignDataCtx77da62_SetRawDataDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackApproveSignDataCtx77da62_RawDataChanged
func callbackApproveSignDataCtx77da62_RawDataChanged(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "rawDataChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(rawData))
	}

}

func (ptr *ApproveSignDataCtx) ConnectRawDataChanged(f func(rawData string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rawDataChanged") {
			C.ApproveSignDataCtx77da62_ConnectRawDataChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rawDataChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rawDataChanged", func(rawData string) {
				signal.(func(string))(rawData)
				f(rawData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rawDataChanged", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectRawDataChanged() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectRawDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rawDataChanged")
	}
}

func (ptr *ApproveSignDataCtx) RawDataChanged(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.ApproveSignDataCtx77da62_RawDataChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackApproveSignDataCtx77da62_Hash
func callbackApproveSignDataCtx77da62_Hash(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "hash"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveSignDataCtxFromPointer(ptr).HashDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveSignDataCtx) ConnectHash(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hash"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hash", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hash", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectHash() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hash")
	}
}

func (ptr *ApproveSignDataCtx) Hash() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_Hash(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) HashDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_HashDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx77da62_SetHash
func callbackApproveSignDataCtx77da62_SetHash(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setHash"); signal != nil {
		signal.(func(string))(cGoUnpackString(hash))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).SetHashDefault(cGoUnpackString(hash))
	}
}

func (ptr *ApproveSignDataCtx) ConnectSetHash(f func(hash string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHash"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setHash", func(hash string) {
				signal.(func(string))(hash)
				f(hash)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHash", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectSetHash() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHash")
	}
}

func (ptr *ApproveSignDataCtx) SetHash(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.ApproveSignDataCtx77da62_SetHash(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

func (ptr *ApproveSignDataCtx) SetHashDefault(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.ApproveSignDataCtx77da62_SetHashDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

//export callbackApproveSignDataCtx77da62_HashChanged
func callbackApproveSignDataCtx77da62_HashChanged(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hashChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(hash))
	}

}

func (ptr *ApproveSignDataCtx) ConnectHashChanged(f func(hash string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hashChanged") {
			C.ApproveSignDataCtx77da62_ConnectHashChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hashChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hashChanged", func(hash string) {
				signal.(func(string))(hash)
				f(hash)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hashChanged", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectHashChanged() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectHashChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hashChanged")
	}
}

func (ptr *ApproveSignDataCtx) HashChanged(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.ApproveSignDataCtx77da62_HashChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

//export callbackApproveSignDataCtx77da62_Password
func callbackApproveSignDataCtx77da62_Password(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "password"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveSignDataCtxFromPointer(ptr).PasswordDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveSignDataCtx) ConnectPassword(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "password"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "password", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "password", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectPassword() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "password")
	}
}

func (ptr *ApproveSignDataCtx) Password() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) PasswordDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_PasswordDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx77da62_SetPassword
func callbackApproveSignDataCtx77da62_SetPassword(ptr unsafe.Pointer, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setPassword"); signal != nil {
		signal.(func(string))(cGoUnpackString(password))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).SetPasswordDefault(cGoUnpackString(password))
	}
}

func (ptr *ApproveSignDataCtx) ConnectSetPassword(f func(password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPassword"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPassword", func(password string) {
				signal.(func(string))(password)
				f(password)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPassword", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectSetPassword() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPassword")
	}
}

func (ptr *ApproveSignDataCtx) SetPassword(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveSignDataCtx77da62_SetPassword(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

func (ptr *ApproveSignDataCtx) SetPasswordDefault(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveSignDataCtx77da62_SetPasswordDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackApproveSignDataCtx77da62_PasswordChanged
func callbackApproveSignDataCtx77da62_PasswordChanged(ptr unsafe.Pointer, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "passwordChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(password))
	}

}

func (ptr *ApproveSignDataCtx) ConnectPasswordChanged(f func(password string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "passwordChanged") {
			C.ApproveSignDataCtx77da62_ConnectPasswordChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "passwordChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "passwordChanged", func(password string) {
				signal.(func(string))(password)
				f(password)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "passwordChanged", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectPasswordChanged() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectPasswordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "passwordChanged")
	}
}

func (ptr *ApproveSignDataCtx) PasswordChanged(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveSignDataCtx77da62_PasswordChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackApproveSignDataCtx77da62_FromSrc
func callbackApproveSignDataCtx77da62_FromSrc(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "fromSrc"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveSignDataCtxFromPointer(ptr).FromSrcDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveSignDataCtx) ConnectFromSrc(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fromSrc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromSrc", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromSrc", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectFromSrc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fromSrc")
	}
}

func (ptr *ApproveSignDataCtx) FromSrc() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_FromSrc(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) FromSrcDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx77da62_FromSrcDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx77da62_SetFromSrc
func callbackApproveSignDataCtx77da62_SetFromSrc(ptr unsafe.Pointer, fromSrc C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFromSrc"); signal != nil {
		signal.(func(string))(cGoUnpackString(fromSrc))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).SetFromSrcDefault(cGoUnpackString(fromSrc))
	}
}

func (ptr *ApproveSignDataCtx) ConnectSetFromSrc(f func(fromSrc string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFromSrc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFromSrc", func(fromSrc string) {
				signal.(func(string))(fromSrc)
				f(fromSrc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFromSrc", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectSetFromSrc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFromSrc")
	}
}

func (ptr *ApproveSignDataCtx) SetFromSrc(fromSrc string) {
	if ptr.Pointer() != nil {
		var fromSrcC *C.char
		if fromSrc != "" {
			fromSrcC = C.CString(fromSrc)
			defer C.free(unsafe.Pointer(fromSrcC))
		}
		C.ApproveSignDataCtx77da62_SetFromSrc(ptr.Pointer(), C.struct_Moc_PackedString{data: fromSrcC, len: C.longlong(len(fromSrc))})
	}
}

func (ptr *ApproveSignDataCtx) SetFromSrcDefault(fromSrc string) {
	if ptr.Pointer() != nil {
		var fromSrcC *C.char
		if fromSrc != "" {
			fromSrcC = C.CString(fromSrc)
			defer C.free(unsafe.Pointer(fromSrcC))
		}
		C.ApproveSignDataCtx77da62_SetFromSrcDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromSrcC, len: C.longlong(len(fromSrc))})
	}
}

//export callbackApproveSignDataCtx77da62_FromSrcChanged
func callbackApproveSignDataCtx77da62_FromSrcChanged(ptr unsafe.Pointer, fromSrc C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromSrcChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(fromSrc))
	}

}

func (ptr *ApproveSignDataCtx) ConnectFromSrcChanged(f func(fromSrc string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromSrcChanged") {
			C.ApproveSignDataCtx77da62_ConnectFromSrcChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fromSrcChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromSrcChanged", func(fromSrc string) {
				signal.(func(string))(fromSrc)
				f(fromSrc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromSrcChanged", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectFromSrcChanged() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectFromSrcChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fromSrcChanged")
	}
}

func (ptr *ApproveSignDataCtx) FromSrcChanged(fromSrc string) {
	if ptr.Pointer() != nil {
		var fromSrcC *C.char
		if fromSrc != "" {
			fromSrcC = C.CString(fromSrc)
			defer C.free(unsafe.Pointer(fromSrcC))
		}
		C.ApproveSignDataCtx77da62_FromSrcChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromSrcC, len: C.longlong(len(fromSrc))})
	}
}

func ApproveSignDataCtx_QRegisterMetaType() int {
	return int(int32(C.ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QRegisterMetaType()))
}

func (ptr *ApproveSignDataCtx) QRegisterMetaType() int {
	return int(int32(C.ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QRegisterMetaType()))
}

func ApproveSignDataCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QRegisterMetaType2(typeNameC)))
}

func (ptr *ApproveSignDataCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QRegisterMetaType2(typeNameC)))
}

func ApproveSignDataCtx_QmlRegisterType() int {
	return int(int32(C.ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QmlRegisterType()))
}

func (ptr *ApproveSignDataCtx) QmlRegisterType() int {
	return int(int32(C.ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QmlRegisterType()))
}

func ApproveSignDataCtx_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveSignDataCtx) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveSignDataCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ApproveSignDataCtx77da62___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ApproveSignDataCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ApproveSignDataCtx77da62___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ApproveSignDataCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveSignDataCtx77da62___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveSignDataCtx) __findChildren_newList2() unsafe.Pointer {
	return C.ApproveSignDataCtx77da62___findChildren_newList2(ptr.Pointer())
}

func (ptr *ApproveSignDataCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveSignDataCtx77da62___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveSignDataCtx) __findChildren_newList3() unsafe.Pointer {
	return C.ApproveSignDataCtx77da62___findChildren_newList3(ptr.Pointer())
}

func (ptr *ApproveSignDataCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveSignDataCtx77da62___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveSignDataCtx) __findChildren_newList() unsafe.Pointer {
	return C.ApproveSignDataCtx77da62___findChildren_newList(ptr.Pointer())
}

func (ptr *ApproveSignDataCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveSignDataCtx77da62___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveSignDataCtx) __children_newList() unsafe.Pointer {
	return C.ApproveSignDataCtx77da62___children_newList(ptr.Pointer())
}

func NewApproveSignDataCtx(parent std_core.QObject_ITF) *ApproveSignDataCtx {
	tmpValue := NewApproveSignDataCtxFromPointer(C.ApproveSignDataCtx77da62_NewApproveSignDataCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApproveSignDataCtx77da62_DestroyApproveSignDataCtx
func callbackApproveSignDataCtx77da62_DestroyApproveSignDataCtx(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ApproveSignDataCtx"); signal != nil {
		signal.(func())()
	} else {
		NewApproveSignDataCtxFromPointer(ptr).DestroyApproveSignDataCtxDefault()
	}
}

func (ptr *ApproveSignDataCtx) ConnectDestroyApproveSignDataCtx(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ApproveSignDataCtx"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~ApproveSignDataCtx", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ApproveSignDataCtx", f)
		}
	}
}

func (ptr *ApproveSignDataCtx) DisconnectDestroyApproveSignDataCtx() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ApproveSignDataCtx")
	}
}

func (ptr *ApproveSignDataCtx) DestroyApproveSignDataCtx() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DestroyApproveSignDataCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ApproveSignDataCtx) DestroyApproveSignDataCtxDefault() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DestroyApproveSignDataCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveSignDataCtx77da62_Event
func callbackApproveSignDataCtx77da62_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveSignDataCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ApproveSignDataCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveSignDataCtx77da62_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackApproveSignDataCtx77da62_EventFilter
func callbackApproveSignDataCtx77da62_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveSignDataCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ApproveSignDataCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveSignDataCtx77da62_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackApproveSignDataCtx77da62_ChildEvent
func callbackApproveSignDataCtx77da62_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ApproveSignDataCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackApproveSignDataCtx77da62_ConnectNotify
func callbackApproveSignDataCtx77da62_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveSignDataCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveSignDataCtx77da62_CustomEvent
func callbackApproveSignDataCtx77da62_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ApproveSignDataCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackApproveSignDataCtx77da62_DeleteLater
func callbackApproveSignDataCtx77da62_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewApproveSignDataCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ApproveSignDataCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveSignDataCtx77da62_Destroyed
func callbackApproveSignDataCtx77da62_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackApproveSignDataCtx77da62_DisconnectNotify
func callbackApproveSignDataCtx77da62_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveSignDataCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveSignDataCtx77da62_ObjectNameChanged
func callbackApproveSignDataCtx77da62_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApproveSignDataCtx77da62_TimerEvent
func callbackApproveSignDataCtx77da62_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ApproveSignDataCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx77da62_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type ApproveTxCtx_ITF interface {
	std_core.QObject_ITF
	ApproveTxCtx_PTR() *ApproveTxCtx
}

func (ptr *ApproveTxCtx) ApproveTxCtx_PTR() *ApproveTxCtx {
	return ptr
}

func (ptr *ApproveTxCtx) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *ApproveTxCtx) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromApproveTxCtx(ptr ApproveTxCtx_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ApproveTxCtx_PTR().Pointer()
	}
	return nil
}

func NewApproveTxCtxFromPointer(ptr unsafe.Pointer) (n *ApproveTxCtx) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ApproveTxCtx)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ApproveTxCtx:
			n = deduced

		case *std_core.QObject:
			n = &ApproveTxCtx{QObject: *deduced}

		default:
			n = new(ApproveTxCtx)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackApproveTxCtx77da62_Constructor
func callbackApproveTxCtx77da62_Constructor(ptr unsafe.Pointer) {
	this := NewApproveTxCtxFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectApprove(this.approve)
	this.ConnectReject(this.reject)
	this.ConnectCheckTxDiff(this.checkTxDiff)
	this.ConnectBack(this.back)
	this.ConnectEdited(this.edited)
	this.ConnectChangeValueUnit(this.changeValueUnit)
	this.ConnectChangeGasPriceUnit(this.changeGasPriceUnit)
	this.init()
}

//export callbackApproveTxCtx77da62_Approve
func callbackApproveTxCtx77da62_Approve(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "approve"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveTxCtx) ConnectApprove(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "approve") {
			C.ApproveTxCtx77da62_ConnectApprove(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "approve"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "approve", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "approve", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectApprove() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectApprove(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "approve")
	}
}

func (ptr *ApproveTxCtx) Approve() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_Approve(ptr.Pointer())
	}
}

//export callbackApproveTxCtx77da62_Reject
func callbackApproveTxCtx77da62_Reject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reject"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveTxCtx) ConnectReject(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "reject") {
			C.ApproveTxCtx77da62_ConnectReject(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "reject"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "reject", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reject", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectReject() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectReject(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "reject")
	}
}

func (ptr *ApproveTxCtx) Reject() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_Reject(ptr.Pointer())
	}
}

//export callbackApproveTxCtx77da62_CheckTxDiff
func callbackApproveTxCtx77da62_CheckTxDiff(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "checkTxDiff"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveTxCtx) ConnectCheckTxDiff(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "checkTxDiff") {
			C.ApproveTxCtx77da62_ConnectCheckTxDiff(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "checkTxDiff"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "checkTxDiff", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "checkTxDiff", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectCheckTxDiff() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectCheckTxDiff(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "checkTxDiff")
	}
}

func (ptr *ApproveTxCtx) CheckTxDiff() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_CheckTxDiff(ptr.Pointer())
	}
}

//export callbackApproveTxCtx77da62_Back
func callbackApproveTxCtx77da62_Back(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "back"); signal != nil {
		signal.(func())()
	}

}

func (ptr *ApproveTxCtx) ConnectBack(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "back") {
			C.ApproveTxCtx77da62_ConnectBack(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "back"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "back", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "back", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectBack() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectBack(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "back")
	}
}

func (ptr *ApproveTxCtx) Back() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_Back(ptr.Pointer())
	}
}

//export callbackApproveTxCtx77da62_Edited
func callbackApproveTxCtx77da62_Edited(ptr unsafe.Pointer, s C.struct_Moc_PackedString, v C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "edited"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(s), cGoUnpackString(v))
	}

}

func (ptr *ApproveTxCtx) ConnectEdited(f func(s string, v string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "edited") {
			C.ApproveTxCtx77da62_ConnectEdited(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "edited"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "edited", func(s string, v string) {
				signal.(func(string, string))(s, v)
				f(s, v)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "edited", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectEdited() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectEdited(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "edited")
	}
}

func (ptr *ApproveTxCtx) Edited(s string, v string) {
	if ptr.Pointer() != nil {
		var sC *C.char
		if s != "" {
			sC = C.CString(s)
			defer C.free(unsafe.Pointer(sC))
		}
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		C.ApproveTxCtx77da62_Edited(ptr.Pointer(), C.struct_Moc_PackedString{data: sC, len: C.longlong(len(s))}, C.struct_Moc_PackedString{data: vC, len: C.longlong(len(v))})
	}
}

//export callbackApproveTxCtx77da62_ChangeValueUnit
func callbackApproveTxCtx77da62_ChangeValueUnit(ptr unsafe.Pointer, v C.int) {
	if signal := qt.GetSignal(ptr, "changeValueUnit"); signal != nil {
		signal.(func(int))(int(int32(v)))
	}

}

func (ptr *ApproveTxCtx) ConnectChangeValueUnit(f func(v int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "changeValueUnit") {
			C.ApproveTxCtx77da62_ConnectChangeValueUnit(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "changeValueUnit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "changeValueUnit", func(v int) {
				signal.(func(int))(v)
				f(v)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "changeValueUnit", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectChangeValueUnit() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectChangeValueUnit(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "changeValueUnit")
	}
}

func (ptr *ApproveTxCtx) ChangeValueUnit(v int) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_ChangeValueUnit(ptr.Pointer(), C.int(int32(v)))
	}
}

//export callbackApproveTxCtx77da62_ChangeGasPriceUnit
func callbackApproveTxCtx77da62_ChangeGasPriceUnit(ptr unsafe.Pointer, v C.int) {
	if signal := qt.GetSignal(ptr, "changeGasPriceUnit"); signal != nil {
		signal.(func(int))(int(int32(v)))
	}

}

func (ptr *ApproveTxCtx) ConnectChangeGasPriceUnit(f func(v int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "changeGasPriceUnit") {
			C.ApproveTxCtx77da62_ConnectChangeGasPriceUnit(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "changeGasPriceUnit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "changeGasPriceUnit", func(v int) {
				signal.(func(int))(v)
				f(v)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "changeGasPriceUnit", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectChangeGasPriceUnit() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectChangeGasPriceUnit(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "changeGasPriceUnit")
	}
}

func (ptr *ApproveTxCtx) ChangeGasPriceUnit(v int) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_ChangeGasPriceUnit(ptr.Pointer(), C.int(int32(v)))
	}
}

//export callbackApproveTxCtx77da62_ValueUnit
func callbackApproveTxCtx77da62_ValueUnit(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "valueUnit"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewApproveTxCtxFromPointer(ptr).ValueUnitDefault()))
}

func (ptr *ApproveTxCtx) ConnectValueUnit(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "valueUnit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "valueUnit", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valueUnit", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectValueUnit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "valueUnit")
	}
}

func (ptr *ApproveTxCtx) ValueUnit() int {
	if ptr.Pointer() != nil {
		return int(int32(C.ApproveTxCtx77da62_ValueUnit(ptr.Pointer())))
	}
	return 0
}

func (ptr *ApproveTxCtx) ValueUnitDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.ApproveTxCtx77da62_ValueUnitDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackApproveTxCtx77da62_SetValueUnit
func callbackApproveTxCtx77da62_SetValueUnit(ptr unsafe.Pointer, valueUnit C.int) {
	if signal := qt.GetSignal(ptr, "setValueUnit"); signal != nil {
		signal.(func(int))(int(int32(valueUnit)))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetValueUnitDefault(int(int32(valueUnit)))
	}
}

func (ptr *ApproveTxCtx) ConnectSetValueUnit(f func(valueUnit int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setValueUnit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setValueUnit", func(valueUnit int) {
				signal.(func(int))(valueUnit)
				f(valueUnit)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setValueUnit", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetValueUnit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setValueUnit")
	}
}

func (ptr *ApproveTxCtx) SetValueUnit(valueUnit int) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_SetValueUnit(ptr.Pointer(), C.int(int32(valueUnit)))
	}
}

func (ptr *ApproveTxCtx) SetValueUnitDefault(valueUnit int) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_SetValueUnitDefault(ptr.Pointer(), C.int(int32(valueUnit)))
	}
}

//export callbackApproveTxCtx77da62_ValueUnitChanged
func callbackApproveTxCtx77da62_ValueUnitChanged(ptr unsafe.Pointer, valueUnit C.int) {
	if signal := qt.GetSignal(ptr, "valueUnitChanged"); signal != nil {
		signal.(func(int))(int(int32(valueUnit)))
	}

}

func (ptr *ApproveTxCtx) ConnectValueUnitChanged(f func(valueUnit int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valueUnitChanged") {
			C.ApproveTxCtx77da62_ConnectValueUnitChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valueUnitChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "valueUnitChanged", func(valueUnit int) {
				signal.(func(int))(valueUnit)
				f(valueUnit)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valueUnitChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectValueUnitChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectValueUnitChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valueUnitChanged")
	}
}

func (ptr *ApproveTxCtx) ValueUnitChanged(valueUnit int) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_ValueUnitChanged(ptr.Pointer(), C.int(int32(valueUnit)))
	}
}

//export callbackApproveTxCtx77da62_Remote
func callbackApproveTxCtx77da62_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "remote"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).RemoteDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectRemote(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "remote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remote", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remote", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "remote")
	}
}

func (ptr *ApproveTxCtx) Remote() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetRemote
func callbackApproveTxCtx77da62_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setRemote"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetRemoteDefault(cGoUnpackString(remote))
	}
}

func (ptr *ApproveTxCtx) ConnectSetRemote(f func(remote string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRemote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRemote")
	}
}

func (ptr *ApproveTxCtx) SetRemote(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveTxCtx77da62_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *ApproveTxCtx) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveTxCtx77da62_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveTxCtx77da62_RemoteChanged
func callbackApproveTxCtx77da62_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *ApproveTxCtx) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.ApproveTxCtx77da62_ConnectRemoteChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "remoteChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectRemoteChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectRemoteChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remoteChanged")
	}
}

func (ptr *ApproveTxCtx) RemoteChanged(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveTxCtx77da62_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveTxCtx77da62_Transport
func callbackApproveTxCtx77da62_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "transport"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).TransportDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectTransport(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "transport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transport", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transport", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transport")
	}
}

func (ptr *ApproveTxCtx) Transport() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetTransport
func callbackApproveTxCtx77da62_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTransport"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetTransportDefault(cGoUnpackString(transport))
	}
}

func (ptr *ApproveTxCtx) ConnectSetTransport(f func(transport string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTransport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransport")
	}
}

func (ptr *ApproveTxCtx) SetTransport(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveTxCtx77da62_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *ApproveTxCtx) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveTxCtx77da62_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveTxCtx77da62_TransportChanged
func callbackApproveTxCtx77da62_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *ApproveTxCtx) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.ApproveTxCtx77da62_ConnectTransportChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transportChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectTransportChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectTransportChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transportChanged")
	}
}

func (ptr *ApproveTxCtx) TransportChanged(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveTxCtx77da62_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveTxCtx77da62_Endpoint
func callbackApproveTxCtx77da62_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "endpoint"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).EndpointDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectEndpoint(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endpoint")
	}
}

func (ptr *ApproveTxCtx) Endpoint() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetEndpoint
func callbackApproveTxCtx77da62_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setEndpoint"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetEndpointDefault(cGoUnpackString(endpoint))
	}
}

func (ptr *ApproveTxCtx) ConnectSetEndpoint(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEndpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEndpoint")
	}
}

func (ptr *ApproveTxCtx) SetEndpoint(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveTxCtx77da62_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *ApproveTxCtx) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveTxCtx77da62_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveTxCtx77da62_EndpointChanged
func callbackApproveTxCtx77da62_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *ApproveTxCtx) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.ApproveTxCtx77da62_ConnectEndpointChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "endpointChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectEndpointChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectEndpointChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "endpointChanged")
	}
}

func (ptr *ApproveTxCtx) EndpointChanged(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveTxCtx77da62_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveTxCtx77da62_Data
func callbackApproveTxCtx77da62_Data(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).DataDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectData(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "data", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *ApproveTxCtx) Data() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) DataDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_DataDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetData
func callbackApproveTxCtx77da62_SetData(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetDataDefault(cGoUnpackString(data))
	}
}

func (ptr *ApproveTxCtx) ConnectSetData(f func(data string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setData", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setData", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setData")
	}
}

func (ptr *ApproveTxCtx) SetData(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.ApproveTxCtx77da62_SetData(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

func (ptr *ApproveTxCtx) SetDataDefault(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.ApproveTxCtx77da62_SetDataDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackApproveTxCtx77da62_DataChanged
func callbackApproveTxCtx77da62_DataChanged(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *ApproveTxCtx) ConnectDataChanged(f func(data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataChanged") {
			C.ApproveTxCtx77da62_ConnectDataChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dataChanged", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectDataChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dataChanged")
	}
}

func (ptr *ApproveTxCtx) DataChanged(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.ApproveTxCtx77da62_DataChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackApproveTxCtx77da62_From
func callbackApproveTxCtx77da62_From(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "from"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).FromDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectFrom(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "from"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "from", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "from", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "from")
	}
}

func (ptr *ApproveTxCtx) From() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_From(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) FromDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_FromDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetFrom
func callbackApproveTxCtx77da62_SetFrom(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFrom"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetFromDefault(cGoUnpackString(from))
	}
}

func (ptr *ApproveTxCtx) ConnectSetFrom(f func(from string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFrom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFrom", func(from string) {
				signal.(func(string))(from)
				f(from)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFrom", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFrom")
	}
}

func (ptr *ApproveTxCtx) SetFrom(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveTxCtx77da62_SetFrom(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

func (ptr *ApproveTxCtx) SetFromDefault(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveTxCtx77da62_SetFromDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveTxCtx77da62_FromChanged
func callbackApproveTxCtx77da62_FromChanged(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	}

}

func (ptr *ApproveTxCtx) ConnectFromChanged(f func(from string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromChanged") {
			C.ApproveTxCtx77da62_ConnectFromChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fromChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromChanged", func(from string) {
				signal.(func(string))(from)
				f(from)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectFromChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectFromChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fromChanged")
	}
}

func (ptr *ApproveTxCtx) FromChanged(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveTxCtx77da62_FromChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveTxCtx77da62_FromWarning
func callbackApproveTxCtx77da62_FromWarning(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "fromWarning"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).FromWarningDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectFromWarning(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fromWarning"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromWarning", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromWarning", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectFromWarning() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fromWarning")
	}
}

func (ptr *ApproveTxCtx) FromWarning() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_FromWarning(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) FromWarningDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_FromWarningDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetFromWarning
func callbackApproveTxCtx77da62_SetFromWarning(ptr unsafe.Pointer, fromWarning C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFromWarning"); signal != nil {
		signal.(func(string))(cGoUnpackString(fromWarning))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetFromWarningDefault(cGoUnpackString(fromWarning))
	}
}

func (ptr *ApproveTxCtx) ConnectSetFromWarning(f func(fromWarning string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFromWarning"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFromWarning", func(fromWarning string) {
				signal.(func(string))(fromWarning)
				f(fromWarning)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFromWarning", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetFromWarning() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFromWarning")
	}
}

func (ptr *ApproveTxCtx) SetFromWarning(fromWarning string) {
	if ptr.Pointer() != nil {
		var fromWarningC *C.char
		if fromWarning != "" {
			fromWarningC = C.CString(fromWarning)
			defer C.free(unsafe.Pointer(fromWarningC))
		}
		C.ApproveTxCtx77da62_SetFromWarning(ptr.Pointer(), C.struct_Moc_PackedString{data: fromWarningC, len: C.longlong(len(fromWarning))})
	}
}

func (ptr *ApproveTxCtx) SetFromWarningDefault(fromWarning string) {
	if ptr.Pointer() != nil {
		var fromWarningC *C.char
		if fromWarning != "" {
			fromWarningC = C.CString(fromWarning)
			defer C.free(unsafe.Pointer(fromWarningC))
		}
		C.ApproveTxCtx77da62_SetFromWarningDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromWarningC, len: C.longlong(len(fromWarning))})
	}
}

//export callbackApproveTxCtx77da62_FromWarningChanged
func callbackApproveTxCtx77da62_FromWarningChanged(ptr unsafe.Pointer, fromWarning C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromWarningChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(fromWarning))
	}

}

func (ptr *ApproveTxCtx) ConnectFromWarningChanged(f func(fromWarning string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromWarningChanged") {
			C.ApproveTxCtx77da62_ConnectFromWarningChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fromWarningChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromWarningChanged", func(fromWarning string) {
				signal.(func(string))(fromWarning)
				f(fromWarning)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromWarningChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectFromWarningChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectFromWarningChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fromWarningChanged")
	}
}

func (ptr *ApproveTxCtx) FromWarningChanged(fromWarning string) {
	if ptr.Pointer() != nil {
		var fromWarningC *C.char
		if fromWarning != "" {
			fromWarningC = C.CString(fromWarning)
			defer C.free(unsafe.Pointer(fromWarningC))
		}
		C.ApproveTxCtx77da62_FromWarningChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromWarningC, len: C.longlong(len(fromWarning))})
	}
}

//export callbackApproveTxCtx77da62_IsFromVisible
func callbackApproveTxCtx77da62_IsFromVisible(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isFromVisible"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveTxCtxFromPointer(ptr).IsFromVisibleDefault())))
}

func (ptr *ApproveTxCtx) ConnectIsFromVisible(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isFromVisible"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isFromVisible", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isFromVisible", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectIsFromVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isFromVisible")
	}
}

func (ptr *ApproveTxCtx) IsFromVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveTxCtx77da62_IsFromVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *ApproveTxCtx) IsFromVisibleDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveTxCtx77da62_IsFromVisibleDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackApproveTxCtx77da62_SetFromVisible
func callbackApproveTxCtx77da62_SetFromVisible(ptr unsafe.Pointer, fromVisible C.char) {
	if signal := qt.GetSignal(ptr, "setFromVisible"); signal != nil {
		signal.(func(bool))(int8(fromVisible) != 0)
	} else {
		NewApproveTxCtxFromPointer(ptr).SetFromVisibleDefault(int8(fromVisible) != 0)
	}
}

func (ptr *ApproveTxCtx) ConnectSetFromVisible(f func(fromVisible bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFromVisible"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFromVisible", func(fromVisible bool) {
				signal.(func(bool))(fromVisible)
				f(fromVisible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFromVisible", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetFromVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFromVisible")
	}
}

func (ptr *ApproveTxCtx) SetFromVisible(fromVisible bool) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_SetFromVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fromVisible))))
	}
}

func (ptr *ApproveTxCtx) SetFromVisibleDefault(fromVisible bool) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_SetFromVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fromVisible))))
	}
}

//export callbackApproveTxCtx77da62_FromVisibleChanged
func callbackApproveTxCtx77da62_FromVisibleChanged(ptr unsafe.Pointer, fromVisible C.char) {
	if signal := qt.GetSignal(ptr, "fromVisibleChanged"); signal != nil {
		signal.(func(bool))(int8(fromVisible) != 0)
	}

}

func (ptr *ApproveTxCtx) ConnectFromVisibleChanged(f func(fromVisible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromVisibleChanged") {
			C.ApproveTxCtx77da62_ConnectFromVisibleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fromVisibleChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromVisibleChanged", func(fromVisible bool) {
				signal.(func(bool))(fromVisible)
				f(fromVisible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromVisibleChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectFromVisibleChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectFromVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fromVisibleChanged")
	}
}

func (ptr *ApproveTxCtx) FromVisibleChanged(fromVisible bool) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_FromVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fromVisible))))
	}
}

//export callbackApproveTxCtx77da62_To
func callbackApproveTxCtx77da62_To(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "to"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).ToDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectTo(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "to"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "to", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "to", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectTo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "to")
	}
}

func (ptr *ApproveTxCtx) To() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_To(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) ToDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_ToDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetTo
func callbackApproveTxCtx77da62_SetTo(ptr unsafe.Pointer, to C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTo"); signal != nil {
		signal.(func(string))(cGoUnpackString(to))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetToDefault(cGoUnpackString(to))
	}
}

func (ptr *ApproveTxCtx) ConnectSetTo(f func(to string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTo"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTo", func(to string) {
				signal.(func(string))(to)
				f(to)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTo", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetTo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTo")
	}
}

func (ptr *ApproveTxCtx) SetTo(to string) {
	if ptr.Pointer() != nil {
		var toC *C.char
		if to != "" {
			toC = C.CString(to)
			defer C.free(unsafe.Pointer(toC))
		}
		C.ApproveTxCtx77da62_SetTo(ptr.Pointer(), C.struct_Moc_PackedString{data: toC, len: C.longlong(len(to))})
	}
}

func (ptr *ApproveTxCtx) SetToDefault(to string) {
	if ptr.Pointer() != nil {
		var toC *C.char
		if to != "" {
			toC = C.CString(to)
			defer C.free(unsafe.Pointer(toC))
		}
		C.ApproveTxCtx77da62_SetToDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: toC, len: C.longlong(len(to))})
	}
}

//export callbackApproveTxCtx77da62_ToChanged
func callbackApproveTxCtx77da62_ToChanged(ptr unsafe.Pointer, to C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "toChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(to))
	}

}

func (ptr *ApproveTxCtx) ConnectToChanged(f func(to string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "toChanged") {
			C.ApproveTxCtx77da62_ConnectToChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "toChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "toChanged", func(to string) {
				signal.(func(string))(to)
				f(to)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "toChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectToChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectToChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "toChanged")
	}
}

func (ptr *ApproveTxCtx) ToChanged(to string) {
	if ptr.Pointer() != nil {
		var toC *C.char
		if to != "" {
			toC = C.CString(to)
			defer C.free(unsafe.Pointer(toC))
		}
		C.ApproveTxCtx77da62_ToChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: toC, len: C.longlong(len(to))})
	}
}

//export callbackApproveTxCtx77da62_ToWarning
func callbackApproveTxCtx77da62_ToWarning(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "toWarning"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).ToWarningDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectToWarning(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "toWarning"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "toWarning", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "toWarning", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectToWarning() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "toWarning")
	}
}

func (ptr *ApproveTxCtx) ToWarning() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_ToWarning(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) ToWarningDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_ToWarningDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetToWarning
func callbackApproveTxCtx77da62_SetToWarning(ptr unsafe.Pointer, toWarning C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setToWarning"); signal != nil {
		signal.(func(string))(cGoUnpackString(toWarning))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetToWarningDefault(cGoUnpackString(toWarning))
	}
}

func (ptr *ApproveTxCtx) ConnectSetToWarning(f func(toWarning string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setToWarning"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setToWarning", func(toWarning string) {
				signal.(func(string))(toWarning)
				f(toWarning)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setToWarning", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetToWarning() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setToWarning")
	}
}

func (ptr *ApproveTxCtx) SetToWarning(toWarning string) {
	if ptr.Pointer() != nil {
		var toWarningC *C.char
		if toWarning != "" {
			toWarningC = C.CString(toWarning)
			defer C.free(unsafe.Pointer(toWarningC))
		}
		C.ApproveTxCtx77da62_SetToWarning(ptr.Pointer(), C.struct_Moc_PackedString{data: toWarningC, len: C.longlong(len(toWarning))})
	}
}

func (ptr *ApproveTxCtx) SetToWarningDefault(toWarning string) {
	if ptr.Pointer() != nil {
		var toWarningC *C.char
		if toWarning != "" {
			toWarningC = C.CString(toWarning)
			defer C.free(unsafe.Pointer(toWarningC))
		}
		C.ApproveTxCtx77da62_SetToWarningDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: toWarningC, len: C.longlong(len(toWarning))})
	}
}

//export callbackApproveTxCtx77da62_ToWarningChanged
func callbackApproveTxCtx77da62_ToWarningChanged(ptr unsafe.Pointer, toWarning C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "toWarningChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(toWarning))
	}

}

func (ptr *ApproveTxCtx) ConnectToWarningChanged(f func(toWarning string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "toWarningChanged") {
			C.ApproveTxCtx77da62_ConnectToWarningChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "toWarningChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "toWarningChanged", func(toWarning string) {
				signal.(func(string))(toWarning)
				f(toWarning)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "toWarningChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectToWarningChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectToWarningChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "toWarningChanged")
	}
}

func (ptr *ApproveTxCtx) ToWarningChanged(toWarning string) {
	if ptr.Pointer() != nil {
		var toWarningC *C.char
		if toWarning != "" {
			toWarningC = C.CString(toWarning)
			defer C.free(unsafe.Pointer(toWarningC))
		}
		C.ApproveTxCtx77da62_ToWarningChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: toWarningC, len: C.longlong(len(toWarning))})
	}
}

//export callbackApproveTxCtx77da62_IsToVisible
func callbackApproveTxCtx77da62_IsToVisible(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isToVisible"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveTxCtxFromPointer(ptr).IsToVisibleDefault())))
}

func (ptr *ApproveTxCtx) ConnectIsToVisible(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isToVisible"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isToVisible", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isToVisible", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectIsToVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isToVisible")
	}
}

func (ptr *ApproveTxCtx) IsToVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveTxCtx77da62_IsToVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *ApproveTxCtx) IsToVisibleDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveTxCtx77da62_IsToVisibleDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackApproveTxCtx77da62_SetToVisible
func callbackApproveTxCtx77da62_SetToVisible(ptr unsafe.Pointer, toVisible C.char) {
	if signal := qt.GetSignal(ptr, "setToVisible"); signal != nil {
		signal.(func(bool))(int8(toVisible) != 0)
	} else {
		NewApproveTxCtxFromPointer(ptr).SetToVisibleDefault(int8(toVisible) != 0)
	}
}

func (ptr *ApproveTxCtx) ConnectSetToVisible(f func(toVisible bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setToVisible"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setToVisible", func(toVisible bool) {
				signal.(func(bool))(toVisible)
				f(toVisible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setToVisible", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetToVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setToVisible")
	}
}

func (ptr *ApproveTxCtx) SetToVisible(toVisible bool) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_SetToVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(toVisible))))
	}
}

func (ptr *ApproveTxCtx) SetToVisibleDefault(toVisible bool) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_SetToVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(toVisible))))
	}
}

//export callbackApproveTxCtx77da62_ToVisibleChanged
func callbackApproveTxCtx77da62_ToVisibleChanged(ptr unsafe.Pointer, toVisible C.char) {
	if signal := qt.GetSignal(ptr, "toVisibleChanged"); signal != nil {
		signal.(func(bool))(int8(toVisible) != 0)
	}

}

func (ptr *ApproveTxCtx) ConnectToVisibleChanged(f func(toVisible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "toVisibleChanged") {
			C.ApproveTxCtx77da62_ConnectToVisibleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "toVisibleChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "toVisibleChanged", func(toVisible bool) {
				signal.(func(bool))(toVisible)
				f(toVisible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "toVisibleChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectToVisibleChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectToVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "toVisibleChanged")
	}
}

func (ptr *ApproveTxCtx) ToVisibleChanged(toVisible bool) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_ToVisibleChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(toVisible))))
	}
}

//export callbackApproveTxCtx77da62_Gas
func callbackApproveTxCtx77da62_Gas(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "gas"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).GasDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectGas(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "gas"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gas", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gas", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectGas() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "gas")
	}
}

func (ptr *ApproveTxCtx) Gas() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_Gas(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) GasDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_GasDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetGas
func callbackApproveTxCtx77da62_SetGas(ptr unsafe.Pointer, gas C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setGas"); signal != nil {
		signal.(func(string))(cGoUnpackString(gas))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetGasDefault(cGoUnpackString(gas))
	}
}

func (ptr *ApproveTxCtx) ConnectSetGas(f func(gas string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGas"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setGas", func(gas string) {
				signal.(func(string))(gas)
				f(gas)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGas", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetGas() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGas")
	}
}

func (ptr *ApproveTxCtx) SetGas(gas string) {
	if ptr.Pointer() != nil {
		var gasC *C.char
		if gas != "" {
			gasC = C.CString(gas)
			defer C.free(unsafe.Pointer(gasC))
		}
		C.ApproveTxCtx77da62_SetGas(ptr.Pointer(), C.struct_Moc_PackedString{data: gasC, len: C.longlong(len(gas))})
	}
}

func (ptr *ApproveTxCtx) SetGasDefault(gas string) {
	if ptr.Pointer() != nil {
		var gasC *C.char
		if gas != "" {
			gasC = C.CString(gas)
			defer C.free(unsafe.Pointer(gasC))
		}
		C.ApproveTxCtx77da62_SetGasDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: gasC, len: C.longlong(len(gas))})
	}
}

//export callbackApproveTxCtx77da62_GasChanged
func callbackApproveTxCtx77da62_GasChanged(ptr unsafe.Pointer, gas C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "gasChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(gas))
	}

}

func (ptr *ApproveTxCtx) ConnectGasChanged(f func(gas string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gasChanged") {
			C.ApproveTxCtx77da62_ConnectGasChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gasChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gasChanged", func(gas string) {
				signal.(func(string))(gas)
				f(gas)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gasChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectGasChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectGasChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gasChanged")
	}
}

func (ptr *ApproveTxCtx) GasChanged(gas string) {
	if ptr.Pointer() != nil {
		var gasC *C.char
		if gas != "" {
			gasC = C.CString(gas)
			defer C.free(unsafe.Pointer(gasC))
		}
		C.ApproveTxCtx77da62_GasChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: gasC, len: C.longlong(len(gas))})
	}
}

//export callbackApproveTxCtx77da62_GasPrice
func callbackApproveTxCtx77da62_GasPrice(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "gasPrice"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).GasPriceDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectGasPrice(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "gasPrice"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gasPrice", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gasPrice", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectGasPrice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "gasPrice")
	}
}

func (ptr *ApproveTxCtx) GasPrice() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_GasPrice(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) GasPriceDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_GasPriceDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetGasPrice
func callbackApproveTxCtx77da62_SetGasPrice(ptr unsafe.Pointer, gasPrice C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setGasPrice"); signal != nil {
		signal.(func(string))(cGoUnpackString(gasPrice))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetGasPriceDefault(cGoUnpackString(gasPrice))
	}
}

func (ptr *ApproveTxCtx) ConnectSetGasPrice(f func(gasPrice string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGasPrice"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setGasPrice", func(gasPrice string) {
				signal.(func(string))(gasPrice)
				f(gasPrice)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGasPrice", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetGasPrice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGasPrice")
	}
}

func (ptr *ApproveTxCtx) SetGasPrice(gasPrice string) {
	if ptr.Pointer() != nil {
		var gasPriceC *C.char
		if gasPrice != "" {
			gasPriceC = C.CString(gasPrice)
			defer C.free(unsafe.Pointer(gasPriceC))
		}
		C.ApproveTxCtx77da62_SetGasPrice(ptr.Pointer(), C.struct_Moc_PackedString{data: gasPriceC, len: C.longlong(len(gasPrice))})
	}
}

func (ptr *ApproveTxCtx) SetGasPriceDefault(gasPrice string) {
	if ptr.Pointer() != nil {
		var gasPriceC *C.char
		if gasPrice != "" {
			gasPriceC = C.CString(gasPrice)
			defer C.free(unsafe.Pointer(gasPriceC))
		}
		C.ApproveTxCtx77da62_SetGasPriceDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: gasPriceC, len: C.longlong(len(gasPrice))})
	}
}

//export callbackApproveTxCtx77da62_GasPriceChanged
func callbackApproveTxCtx77da62_GasPriceChanged(ptr unsafe.Pointer, gasPrice C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "gasPriceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(gasPrice))
	}

}

func (ptr *ApproveTxCtx) ConnectGasPriceChanged(f func(gasPrice string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gasPriceChanged") {
			C.ApproveTxCtx77da62_ConnectGasPriceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gasPriceChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gasPriceChanged", func(gasPrice string) {
				signal.(func(string))(gasPrice)
				f(gasPrice)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gasPriceChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectGasPriceChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectGasPriceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gasPriceChanged")
	}
}

func (ptr *ApproveTxCtx) GasPriceChanged(gasPrice string) {
	if ptr.Pointer() != nil {
		var gasPriceC *C.char
		if gasPrice != "" {
			gasPriceC = C.CString(gasPrice)
			defer C.free(unsafe.Pointer(gasPriceC))
		}
		C.ApproveTxCtx77da62_GasPriceChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: gasPriceC, len: C.longlong(len(gasPrice))})
	}
}

//export callbackApproveTxCtx77da62_GasPriceUnit
func callbackApproveTxCtx77da62_GasPriceUnit(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "gasPriceUnit"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewApproveTxCtxFromPointer(ptr).GasPriceUnitDefault()))
}

func (ptr *ApproveTxCtx) ConnectGasPriceUnit(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "gasPriceUnit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gasPriceUnit", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gasPriceUnit", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectGasPriceUnit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "gasPriceUnit")
	}
}

func (ptr *ApproveTxCtx) GasPriceUnit() int {
	if ptr.Pointer() != nil {
		return int(int32(C.ApproveTxCtx77da62_GasPriceUnit(ptr.Pointer())))
	}
	return 0
}

func (ptr *ApproveTxCtx) GasPriceUnitDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.ApproveTxCtx77da62_GasPriceUnitDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackApproveTxCtx77da62_SetGasPriceUnit
func callbackApproveTxCtx77da62_SetGasPriceUnit(ptr unsafe.Pointer, gasPriceUnit C.int) {
	if signal := qt.GetSignal(ptr, "setGasPriceUnit"); signal != nil {
		signal.(func(int))(int(int32(gasPriceUnit)))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetGasPriceUnitDefault(int(int32(gasPriceUnit)))
	}
}

func (ptr *ApproveTxCtx) ConnectSetGasPriceUnit(f func(gasPriceUnit int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGasPriceUnit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setGasPriceUnit", func(gasPriceUnit int) {
				signal.(func(int))(gasPriceUnit)
				f(gasPriceUnit)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGasPriceUnit", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetGasPriceUnit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGasPriceUnit")
	}
}

func (ptr *ApproveTxCtx) SetGasPriceUnit(gasPriceUnit int) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_SetGasPriceUnit(ptr.Pointer(), C.int(int32(gasPriceUnit)))
	}
}

func (ptr *ApproveTxCtx) SetGasPriceUnitDefault(gasPriceUnit int) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_SetGasPriceUnitDefault(ptr.Pointer(), C.int(int32(gasPriceUnit)))
	}
}

//export callbackApproveTxCtx77da62_GasPriceUnitChanged
func callbackApproveTxCtx77da62_GasPriceUnitChanged(ptr unsafe.Pointer, gasPriceUnit C.int) {
	if signal := qt.GetSignal(ptr, "gasPriceUnitChanged"); signal != nil {
		signal.(func(int))(int(int32(gasPriceUnit)))
	}

}

func (ptr *ApproveTxCtx) ConnectGasPriceUnitChanged(f func(gasPriceUnit int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gasPriceUnitChanged") {
			C.ApproveTxCtx77da62_ConnectGasPriceUnitChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gasPriceUnitChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gasPriceUnitChanged", func(gasPriceUnit int) {
				signal.(func(int))(gasPriceUnit)
				f(gasPriceUnit)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gasPriceUnitChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectGasPriceUnitChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectGasPriceUnitChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gasPriceUnitChanged")
	}
}

func (ptr *ApproveTxCtx) GasPriceUnitChanged(gasPriceUnit int) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_GasPriceUnitChanged(ptr.Pointer(), C.int(int32(gasPriceUnit)))
	}
}

//export callbackApproveTxCtx77da62_Nonce
func callbackApproveTxCtx77da62_Nonce(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "nonce"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).NonceDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectNonce(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nonce"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "nonce", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nonce", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectNonce() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nonce")
	}
}

func (ptr *ApproveTxCtx) Nonce() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_Nonce(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) NonceDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_NonceDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetNonce
func callbackApproveTxCtx77da62_SetNonce(ptr unsafe.Pointer, nonce C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setNonce"); signal != nil {
		signal.(func(string))(cGoUnpackString(nonce))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetNonceDefault(cGoUnpackString(nonce))
	}
}

func (ptr *ApproveTxCtx) ConnectSetNonce(f func(nonce string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setNonce"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setNonce", func(nonce string) {
				signal.(func(string))(nonce)
				f(nonce)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setNonce", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetNonce() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setNonce")
	}
}

func (ptr *ApproveTxCtx) SetNonce(nonce string) {
	if ptr.Pointer() != nil {
		var nonceC *C.char
		if nonce != "" {
			nonceC = C.CString(nonce)
			defer C.free(unsafe.Pointer(nonceC))
		}
		C.ApproveTxCtx77da62_SetNonce(ptr.Pointer(), C.struct_Moc_PackedString{data: nonceC, len: C.longlong(len(nonce))})
	}
}

func (ptr *ApproveTxCtx) SetNonceDefault(nonce string) {
	if ptr.Pointer() != nil {
		var nonceC *C.char
		if nonce != "" {
			nonceC = C.CString(nonce)
			defer C.free(unsafe.Pointer(nonceC))
		}
		C.ApproveTxCtx77da62_SetNonceDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: nonceC, len: C.longlong(len(nonce))})
	}
}

//export callbackApproveTxCtx77da62_NonceChanged
func callbackApproveTxCtx77da62_NonceChanged(ptr unsafe.Pointer, nonce C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "nonceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(nonce))
	}

}

func (ptr *ApproveTxCtx) ConnectNonceChanged(f func(nonce string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nonceChanged") {
			C.ApproveTxCtx77da62_ConnectNonceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "nonceChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "nonceChanged", func(nonce string) {
				signal.(func(string))(nonce)
				f(nonce)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nonceChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectNonceChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectNonceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "nonceChanged")
	}
}

func (ptr *ApproveTxCtx) NonceChanged(nonce string) {
	if ptr.Pointer() != nil {
		var nonceC *C.char
		if nonce != "" {
			nonceC = C.CString(nonce)
			defer C.free(unsafe.Pointer(nonceC))
		}
		C.ApproveTxCtx77da62_NonceChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: nonceC, len: C.longlong(len(nonce))})
	}
}

//export callbackApproveTxCtx77da62_Value
func callbackApproveTxCtx77da62_Value(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "value"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).ValueDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectValue(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "value"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "value", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "value", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "value")
	}
}

func (ptr *ApproveTxCtx) Value() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_Value(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) ValueDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_ValueDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetValue
func callbackApproveTxCtx77da62_SetValue(ptr unsafe.Pointer, value C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setValue"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetValueDefault(cGoUnpackString(value))
	}
}

func (ptr *ApproveTxCtx) ConnectSetValue(f func(value string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setValue", func(value string) {
				signal.(func(string))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setValue", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setValue")
	}
}

func (ptr *ApproveTxCtx) SetValue(value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.ApproveTxCtx77da62_SetValue(ptr.Pointer(), C.struct_Moc_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *ApproveTxCtx) SetValueDefault(value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.ApproveTxCtx77da62_SetValueDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

//export callbackApproveTxCtx77da62_ValueChanged
func callbackApproveTxCtx77da62_ValueChanged(ptr unsafe.Pointer, value C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "valueChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	}

}

func (ptr *ApproveTxCtx) ConnectValueChanged(f func(value string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valueChanged") {
			C.ApproveTxCtx77da62_ConnectValueChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valueChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", func(value string) {
				signal.(func(string))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valueChanged")
	}
}

func (ptr *ApproveTxCtx) ValueChanged(value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.ApproveTxCtx77da62_ValueChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

//export callbackApproveTxCtx77da62_Password
func callbackApproveTxCtx77da62_Password(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "password"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).PasswordDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectPassword(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "password"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "password", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "password", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectPassword() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "password")
	}
}

func (ptr *ApproveTxCtx) Password() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) PasswordDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_PasswordDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetPassword
func callbackApproveTxCtx77da62_SetPassword(ptr unsafe.Pointer, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setPassword"); signal != nil {
		signal.(func(string))(cGoUnpackString(password))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetPasswordDefault(cGoUnpackString(password))
	}
}

func (ptr *ApproveTxCtx) ConnectSetPassword(f func(password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPassword"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPassword", func(password string) {
				signal.(func(string))(password)
				f(password)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPassword", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetPassword() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPassword")
	}
}

func (ptr *ApproveTxCtx) SetPassword(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveTxCtx77da62_SetPassword(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

func (ptr *ApproveTxCtx) SetPasswordDefault(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveTxCtx77da62_SetPasswordDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackApproveTxCtx77da62_PasswordChanged
func callbackApproveTxCtx77da62_PasswordChanged(ptr unsafe.Pointer, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "passwordChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(password))
	}

}

func (ptr *ApproveTxCtx) ConnectPasswordChanged(f func(password string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "passwordChanged") {
			C.ApproveTxCtx77da62_ConnectPasswordChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "passwordChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "passwordChanged", func(password string) {
				signal.(func(string))(password)
				f(password)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "passwordChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectPasswordChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectPasswordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "passwordChanged")
	}
}

func (ptr *ApproveTxCtx) PasswordChanged(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveTxCtx77da62_PasswordChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackApproveTxCtx77da62_FromSrc
func callbackApproveTxCtx77da62_FromSrc(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "fromSrc"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).FromSrcDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectFromSrc(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fromSrc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromSrc", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromSrc", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectFromSrc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fromSrc")
	}
}

func (ptr *ApproveTxCtx) FromSrc() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_FromSrc(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) FromSrcDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_FromSrcDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetFromSrc
func callbackApproveTxCtx77da62_SetFromSrc(ptr unsafe.Pointer, fromSrc C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFromSrc"); signal != nil {
		signal.(func(string))(cGoUnpackString(fromSrc))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetFromSrcDefault(cGoUnpackString(fromSrc))
	}
}

func (ptr *ApproveTxCtx) ConnectSetFromSrc(f func(fromSrc string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFromSrc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFromSrc", func(fromSrc string) {
				signal.(func(string))(fromSrc)
				f(fromSrc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFromSrc", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetFromSrc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFromSrc")
	}
}

func (ptr *ApproveTxCtx) SetFromSrc(fromSrc string) {
	if ptr.Pointer() != nil {
		var fromSrcC *C.char
		if fromSrc != "" {
			fromSrcC = C.CString(fromSrc)
			defer C.free(unsafe.Pointer(fromSrcC))
		}
		C.ApproveTxCtx77da62_SetFromSrc(ptr.Pointer(), C.struct_Moc_PackedString{data: fromSrcC, len: C.longlong(len(fromSrc))})
	}
}

func (ptr *ApproveTxCtx) SetFromSrcDefault(fromSrc string) {
	if ptr.Pointer() != nil {
		var fromSrcC *C.char
		if fromSrc != "" {
			fromSrcC = C.CString(fromSrc)
			defer C.free(unsafe.Pointer(fromSrcC))
		}
		C.ApproveTxCtx77da62_SetFromSrcDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromSrcC, len: C.longlong(len(fromSrc))})
	}
}

//export callbackApproveTxCtx77da62_FromSrcChanged
func callbackApproveTxCtx77da62_FromSrcChanged(ptr unsafe.Pointer, fromSrc C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromSrcChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(fromSrc))
	}

}

func (ptr *ApproveTxCtx) ConnectFromSrcChanged(f func(fromSrc string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromSrcChanged") {
			C.ApproveTxCtx77da62_ConnectFromSrcChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fromSrcChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromSrcChanged", func(fromSrc string) {
				signal.(func(string))(fromSrc)
				f(fromSrc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromSrcChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectFromSrcChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectFromSrcChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fromSrcChanged")
	}
}

func (ptr *ApproveTxCtx) FromSrcChanged(fromSrc string) {
	if ptr.Pointer() != nil {
		var fromSrcC *C.char
		if fromSrc != "" {
			fromSrcC = C.CString(fromSrc)
			defer C.free(unsafe.Pointer(fromSrcC))
		}
		C.ApproveTxCtx77da62_FromSrcChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromSrcC, len: C.longlong(len(fromSrc))})
	}
}

//export callbackApproveTxCtx77da62_ToSrc
func callbackApproveTxCtx77da62_ToSrc(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "toSrc"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).ToSrcDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectToSrc(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "toSrc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "toSrc", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "toSrc", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectToSrc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "toSrc")
	}
}

func (ptr *ApproveTxCtx) ToSrc() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_ToSrc(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) ToSrcDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_ToSrcDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetToSrc
func callbackApproveTxCtx77da62_SetToSrc(ptr unsafe.Pointer, toSrc C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setToSrc"); signal != nil {
		signal.(func(string))(cGoUnpackString(toSrc))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetToSrcDefault(cGoUnpackString(toSrc))
	}
}

func (ptr *ApproveTxCtx) ConnectSetToSrc(f func(toSrc string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setToSrc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setToSrc", func(toSrc string) {
				signal.(func(string))(toSrc)
				f(toSrc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setToSrc", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetToSrc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setToSrc")
	}
}

func (ptr *ApproveTxCtx) SetToSrc(toSrc string) {
	if ptr.Pointer() != nil {
		var toSrcC *C.char
		if toSrc != "" {
			toSrcC = C.CString(toSrc)
			defer C.free(unsafe.Pointer(toSrcC))
		}
		C.ApproveTxCtx77da62_SetToSrc(ptr.Pointer(), C.struct_Moc_PackedString{data: toSrcC, len: C.longlong(len(toSrc))})
	}
}

func (ptr *ApproveTxCtx) SetToSrcDefault(toSrc string) {
	if ptr.Pointer() != nil {
		var toSrcC *C.char
		if toSrc != "" {
			toSrcC = C.CString(toSrc)
			defer C.free(unsafe.Pointer(toSrcC))
		}
		C.ApproveTxCtx77da62_SetToSrcDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: toSrcC, len: C.longlong(len(toSrc))})
	}
}

//export callbackApproveTxCtx77da62_ToSrcChanged
func callbackApproveTxCtx77da62_ToSrcChanged(ptr unsafe.Pointer, toSrc C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "toSrcChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(toSrc))
	}

}

func (ptr *ApproveTxCtx) ConnectToSrcChanged(f func(toSrc string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "toSrcChanged") {
			C.ApproveTxCtx77da62_ConnectToSrcChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "toSrcChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "toSrcChanged", func(toSrc string) {
				signal.(func(string))(toSrc)
				f(toSrc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "toSrcChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectToSrcChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectToSrcChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "toSrcChanged")
	}
}

func (ptr *ApproveTxCtx) ToSrcChanged(toSrc string) {
	if ptr.Pointer() != nil {
		var toSrcC *C.char
		if toSrc != "" {
			toSrcC = C.CString(toSrc)
			defer C.free(unsafe.Pointer(toSrcC))
		}
		C.ApproveTxCtx77da62_ToSrcChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: toSrcC, len: C.longlong(len(toSrc))})
	}
}

//export callbackApproveTxCtx77da62_Diff
func callbackApproveTxCtx77da62_Diff(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "diff"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewApproveTxCtxFromPointer(ptr).DiffDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ApproveTxCtx) ConnectDiff(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "diff"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "diff", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "diff", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectDiff() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "diff")
	}
}

func (ptr *ApproveTxCtx) Diff() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_Diff(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) DiffDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx77da62_DiffDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx77da62_SetDiff
func callbackApproveTxCtx77da62_SetDiff(ptr unsafe.Pointer, diff C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setDiff"); signal != nil {
		signal.(func(string))(cGoUnpackString(diff))
	} else {
		NewApproveTxCtxFromPointer(ptr).SetDiffDefault(cGoUnpackString(diff))
	}
}

func (ptr *ApproveTxCtx) ConnectSetDiff(f func(diff string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDiff"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDiff", func(diff string) {
				signal.(func(string))(diff)
				f(diff)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDiff", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectSetDiff() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDiff")
	}
}

func (ptr *ApproveTxCtx) SetDiff(diff string) {
	if ptr.Pointer() != nil {
		var diffC *C.char
		if diff != "" {
			diffC = C.CString(diff)
			defer C.free(unsafe.Pointer(diffC))
		}
		C.ApproveTxCtx77da62_SetDiff(ptr.Pointer(), C.struct_Moc_PackedString{data: diffC, len: C.longlong(len(diff))})
	}
}

func (ptr *ApproveTxCtx) SetDiffDefault(diff string) {
	if ptr.Pointer() != nil {
		var diffC *C.char
		if diff != "" {
			diffC = C.CString(diff)
			defer C.free(unsafe.Pointer(diffC))
		}
		C.ApproveTxCtx77da62_SetDiffDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: diffC, len: C.longlong(len(diff))})
	}
}

//export callbackApproveTxCtx77da62_DiffChanged
func callbackApproveTxCtx77da62_DiffChanged(ptr unsafe.Pointer, diff C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "diffChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(diff))
	}

}

func (ptr *ApproveTxCtx) ConnectDiffChanged(f func(diff string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "diffChanged") {
			C.ApproveTxCtx77da62_ConnectDiffChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "diffChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "diffChanged", func(diff string) {
				signal.(func(string))(diff)
				f(diff)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "diffChanged", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectDiffChanged() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectDiffChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "diffChanged")
	}
}

func (ptr *ApproveTxCtx) DiffChanged(diff string) {
	if ptr.Pointer() != nil {
		var diffC *C.char
		if diff != "" {
			diffC = C.CString(diff)
			defer C.free(unsafe.Pointer(diffC))
		}
		C.ApproveTxCtx77da62_DiffChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: diffC, len: C.longlong(len(diff))})
	}
}

func ApproveTxCtx_QRegisterMetaType() int {
	return int(int32(C.ApproveTxCtx77da62_ApproveTxCtx77da62_QRegisterMetaType()))
}

func (ptr *ApproveTxCtx) QRegisterMetaType() int {
	return int(int32(C.ApproveTxCtx77da62_ApproveTxCtx77da62_QRegisterMetaType()))
}

func ApproveTxCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveTxCtx77da62_ApproveTxCtx77da62_QRegisterMetaType2(typeNameC)))
}

func (ptr *ApproveTxCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveTxCtx77da62_ApproveTxCtx77da62_QRegisterMetaType2(typeNameC)))
}

func ApproveTxCtx_QmlRegisterType() int {
	return int(int32(C.ApproveTxCtx77da62_ApproveTxCtx77da62_QmlRegisterType()))
}

func (ptr *ApproveTxCtx) QmlRegisterType() int {
	return int(int32(C.ApproveTxCtx77da62_ApproveTxCtx77da62_QmlRegisterType()))
}

func ApproveTxCtx_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApproveTxCtx77da62_ApproveTxCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveTxCtx) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApproveTxCtx77da62_ApproveTxCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveTxCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ApproveTxCtx77da62___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ApproveTxCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ApproveTxCtx77da62___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ApproveTxCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveTxCtx77da62___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveTxCtx) __findChildren_newList2() unsafe.Pointer {
	return C.ApproveTxCtx77da62___findChildren_newList2(ptr.Pointer())
}

func (ptr *ApproveTxCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveTxCtx77da62___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveTxCtx) __findChildren_newList3() unsafe.Pointer {
	return C.ApproveTxCtx77da62___findChildren_newList3(ptr.Pointer())
}

func (ptr *ApproveTxCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveTxCtx77da62___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveTxCtx) __findChildren_newList() unsafe.Pointer {
	return C.ApproveTxCtx77da62___findChildren_newList(ptr.Pointer())
}

func (ptr *ApproveTxCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveTxCtx77da62___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveTxCtx) __children_newList() unsafe.Pointer {
	return C.ApproveTxCtx77da62___children_newList(ptr.Pointer())
}

func NewApproveTxCtx(parent std_core.QObject_ITF) *ApproveTxCtx {
	tmpValue := NewApproveTxCtxFromPointer(C.ApproveTxCtx77da62_NewApproveTxCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApproveTxCtx77da62_DestroyApproveTxCtx
func callbackApproveTxCtx77da62_DestroyApproveTxCtx(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ApproveTxCtx"); signal != nil {
		signal.(func())()
	} else {
		NewApproveTxCtxFromPointer(ptr).DestroyApproveTxCtxDefault()
	}
}

func (ptr *ApproveTxCtx) ConnectDestroyApproveTxCtx(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ApproveTxCtx"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~ApproveTxCtx", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ApproveTxCtx", f)
		}
	}
}

func (ptr *ApproveTxCtx) DisconnectDestroyApproveTxCtx() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ApproveTxCtx")
	}
}

func (ptr *ApproveTxCtx) DestroyApproveTxCtx() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DestroyApproveTxCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ApproveTxCtx) DestroyApproveTxCtxDefault() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DestroyApproveTxCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveTxCtx77da62_Event
func callbackApproveTxCtx77da62_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveTxCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ApproveTxCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveTxCtx77da62_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackApproveTxCtx77da62_EventFilter
func callbackApproveTxCtx77da62_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveTxCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ApproveTxCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ApproveTxCtx77da62_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackApproveTxCtx77da62_ChildEvent
func callbackApproveTxCtx77da62_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewApproveTxCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ApproveTxCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackApproveTxCtx77da62_ConnectNotify
func callbackApproveTxCtx77da62_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveTxCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveTxCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveTxCtx77da62_CustomEvent
func callbackApproveTxCtx77da62_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewApproveTxCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ApproveTxCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackApproveTxCtx77da62_DeleteLater
func callbackApproveTxCtx77da62_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewApproveTxCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ApproveTxCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveTxCtx77da62_Destroyed
func callbackApproveTxCtx77da62_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackApproveTxCtx77da62_DisconnectNotify
func callbackApproveTxCtx77da62_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveTxCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveTxCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveTxCtx77da62_ObjectNameChanged
func callbackApproveTxCtx77da62_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApproveTxCtx77da62_TimerEvent
func callbackApproveTxCtx77da62_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewApproveTxCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ApproveTxCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx77da62_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type CustomListModel_ITF interface {
	std_core.QAbstractListModel_ITF
	CustomListModel_PTR() *CustomListModel
}

func (ptr *CustomListModel) CustomListModel_PTR() *CustomListModel {
	return ptr
}

func (ptr *CustomListModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *CustomListModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromCustomListModel(ptr CustomListModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CustomListModel_PTR().Pointer()
	}
	return nil
}

func NewCustomListModelFromPointer(ptr unsafe.Pointer) (n *CustomListModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(CustomListModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *CustomListModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &CustomListModel{QAbstractListModel: *deduced}

		default:
			n = new(CustomListModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackCustomListModel77da62_Constructor
func callbackCustomListModel77da62_Constructor(ptr unsafe.Pointer) {
	this := NewCustomListModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClear(this.clear)
	this.ConnectAdd(this.add)
	this.init()
}

//export callbackCustomListModel77da62_Clear
func callbackCustomListModel77da62_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		signal.(func())()
	}

}

func (ptr *CustomListModel) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clear") {
			C.CustomListModel77da62_ConnectClear(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clear", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectClear() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_DisconnectClear(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *CustomListModel) Clear() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_Clear(ptr.Pointer())
	}
}

//export callbackCustomListModel77da62_Add
func callbackCustomListModel77da62_Add(ptr unsafe.Pointer, account C.uintptr_t) {
	var accountD custom_accounts_0de74dm.Account
	if accountI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(account))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(account)))
		accountD = accountI.(custom_accounts_0de74dm.Account)
	}
	if signal := qt.GetSignal(ptr, "add"); signal != nil {
		signal.(func(custom_accounts_0de74dm.Account))(accountD)
	}

}

func (ptr *CustomListModel) ConnectAdd(f func(account custom_accounts_0de74dm.Account)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "add") {
			C.CustomListModel77da62_ConnectAdd(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "add"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "add", func(account custom_accounts_0de74dm.Account) {
				signal.(func(custom_accounts_0de74dm.Account))(account)
				f(account)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "add", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectAdd() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_DisconnectAdd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "add")
	}
}

func (ptr *CustomListModel) Add(account custom_accounts_0de74dm.Account) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&account), account)
		C.CustomListModel77da62_Add(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&account))))
	}
}

//export callbackCustomListModel77da62_CallbackFromQml
func callbackCustomListModel77da62_CallbackFromQml(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "callbackFromQml"); signal != nil {
		signal.(func())()
	}

}

func (ptr *CustomListModel) ConnectCallbackFromQml(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "callbackFromQml"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "callbackFromQml", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "callbackFromQml", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectCallbackFromQml() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "callbackFromQml")
	}
}

func (ptr *CustomListModel) CallbackFromQml() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_CallbackFromQml(ptr.Pointer())
	}
}

//export callbackCustomListModel77da62_UpdateRequired
func callbackCustomListModel77da62_UpdateRequired(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "updateRequired"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewCustomListModelFromPointer(ptr).UpdateRequiredDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CustomListModel) ConnectUpdateRequired(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateRequired"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateRequired", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateRequired", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectUpdateRequired() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateRequired")
	}
}

func (ptr *CustomListModel) UpdateRequired() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CustomListModel77da62_UpdateRequired(ptr.Pointer()))
	}
	return ""
}

func (ptr *CustomListModel) UpdateRequiredDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CustomListModel77da62_UpdateRequiredDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackCustomListModel77da62_SetUpdateRequired
func callbackCustomListModel77da62_SetUpdateRequired(ptr unsafe.Pointer, updateRequired C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setUpdateRequired"); signal != nil {
		signal.(func(string))(cGoUnpackString(updateRequired))
	} else {
		NewCustomListModelFromPointer(ptr).SetUpdateRequiredDefault(cGoUnpackString(updateRequired))
	}
}

func (ptr *CustomListModel) ConnectSetUpdateRequired(f func(updateRequired string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setUpdateRequired"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setUpdateRequired", func(updateRequired string) {
				signal.(func(string))(updateRequired)
				f(updateRequired)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setUpdateRequired", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectSetUpdateRequired() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setUpdateRequired")
	}
}

func (ptr *CustomListModel) SetUpdateRequired(updateRequired string) {
	if ptr.Pointer() != nil {
		var updateRequiredC *C.char
		if updateRequired != "" {
			updateRequiredC = C.CString(updateRequired)
			defer C.free(unsafe.Pointer(updateRequiredC))
		}
		C.CustomListModel77da62_SetUpdateRequired(ptr.Pointer(), C.struct_Moc_PackedString{data: updateRequiredC, len: C.longlong(len(updateRequired))})
	}
}

func (ptr *CustomListModel) SetUpdateRequiredDefault(updateRequired string) {
	if ptr.Pointer() != nil {
		var updateRequiredC *C.char
		if updateRequired != "" {
			updateRequiredC = C.CString(updateRequired)
			defer C.free(unsafe.Pointer(updateRequiredC))
		}
		C.CustomListModel77da62_SetUpdateRequiredDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: updateRequiredC, len: C.longlong(len(updateRequired))})
	}
}

//export callbackCustomListModel77da62_UpdateRequiredChanged
func callbackCustomListModel77da62_UpdateRequiredChanged(ptr unsafe.Pointer, updateRequired C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "updateRequiredChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(updateRequired))
	}

}

func (ptr *CustomListModel) ConnectUpdateRequiredChanged(f func(updateRequired string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateRequiredChanged") {
			C.CustomListModel77da62_ConnectUpdateRequiredChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateRequiredChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateRequiredChanged", func(updateRequired string) {
				signal.(func(string))(updateRequired)
				f(updateRequired)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateRequiredChanged", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectUpdateRequiredChanged() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_DisconnectUpdateRequiredChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateRequiredChanged")
	}
}

func (ptr *CustomListModel) UpdateRequiredChanged(updateRequired string) {
	if ptr.Pointer() != nil {
		var updateRequiredC *C.char
		if updateRequired != "" {
			updateRequiredC = C.CString(updateRequired)
			defer C.free(unsafe.Pointer(updateRequiredC))
		}
		C.CustomListModel77da62_UpdateRequiredChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: updateRequiredC, len: C.longlong(len(updateRequired))})
	}
}

func CustomListModel_QRegisterMetaType() int {
	return int(int32(C.CustomListModel77da62_CustomListModel77da62_QRegisterMetaType()))
}

func (ptr *CustomListModel) QRegisterMetaType() int {
	return int(int32(C.CustomListModel77da62_CustomListModel77da62_QRegisterMetaType()))
}

func CustomListModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomListModel77da62_CustomListModel77da62_QRegisterMetaType2(typeNameC)))
}

func (ptr *CustomListModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomListModel77da62_CustomListModel77da62_QRegisterMetaType2(typeNameC)))
}

func CustomListModel_QmlRegisterType() int {
	return int(int32(C.CustomListModel77da62_CustomListModel77da62_QmlRegisterType()))
}

func (ptr *CustomListModel) QmlRegisterType() int {
	return int(int32(C.CustomListModel77da62_CustomListModel77da62_QmlRegisterType()))
}

func CustomListModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.CustomListModel77da62_CustomListModel77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomListModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.CustomListModel77da62_CustomListModel77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomListModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel77da62_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.CustomListModel77da62_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel77da62_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomListModel77da62_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel77da62_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.CustomListModel77da62_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomListModel77da62___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomListModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.CustomListModel77da62___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.CustomListModel77da62___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomListModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel77da62___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.CustomListModel77da62___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel77da62___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.CustomListModel77da62___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel77da62___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.CustomListModel77da62___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomListModel77da62___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomListModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.CustomListModel77da62___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomListModel77da62___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomListModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.CustomListModel77da62___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomListModel77da62___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomListModel) __roleNames_newList() unsafe.Pointer {
	return C.CustomListModel77da62___roleNames_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.CustomListModel77da62___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomListModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomListModel77da62___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomListModel) __itemData_newList() unsafe.Pointer {
	return C.CustomListModel77da62___itemData_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.CustomListModel77da62___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomListModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel77da62___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.CustomListModel77da62___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel77da62___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __match_newList() unsafe.Pointer {
	return C.CustomListModel77da62___match_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel77da62___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.CustomListModel77da62___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel77da62_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomListModel77da62_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel77da62_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomListModel77da62_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomListModel77da62___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomListModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CustomListModel77da62___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomListModel77da62___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomListModel) __findChildren_newList2() unsafe.Pointer {
	return C.CustomListModel77da62___findChildren_newList2(ptr.Pointer())
}

func (ptr *CustomListModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomListModel77da62___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomListModel) __findChildren_newList3() unsafe.Pointer {
	return C.CustomListModel77da62___findChildren_newList3(ptr.Pointer())
}

func (ptr *CustomListModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomListModel77da62___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomListModel) __findChildren_newList() unsafe.Pointer {
	return C.CustomListModel77da62___findChildren_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomListModel77da62___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomListModel) __children_newList() unsafe.Pointer {
	return C.CustomListModel77da62___children_newList(ptr.Pointer())
}

func NewCustomListModel(parent std_core.QObject_ITF) *CustomListModel {
	tmpValue := NewCustomListModelFromPointer(C.CustomListModel77da62_NewCustomListModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCustomListModel77da62_DestroyCustomListModel
func callbackCustomListModel77da62_DestroyCustomListModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~CustomListModel"); signal != nil {
		signal.(func())()
	} else {
		NewCustomListModelFromPointer(ptr).DestroyCustomListModelDefault()
	}
}

func (ptr *CustomListModel) ConnectDestroyCustomListModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~CustomListModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~CustomListModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~CustomListModel", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectDestroyCustomListModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~CustomListModel")
	}
}

func (ptr *CustomListModel) DestroyCustomListModel() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_DestroyCustomListModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CustomListModel) DestroyCustomListModelDefault() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_DestroyCustomListModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomListModel77da62_DropMimeData
func callbackCustomListModel77da62_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_Index
func callbackCustomListModel77da62_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewCustomListModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *CustomListModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel77da62_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel77da62_Sibling
func callbackCustomListModel77da62_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewCustomListModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *CustomListModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel77da62_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel77da62_Flags
func callbackCustomListModel77da62_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewCustomListModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomListModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.CustomListModel77da62_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackCustomListModel77da62_InsertColumns
func callbackCustomListModel77da62_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_InsertRows
func callbackCustomListModel77da62_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_MoveColumns
func callbackCustomListModel77da62_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomListModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_MoveRows
func callbackCustomListModel77da62_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomListModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_RemoveColumns
func callbackCustomListModel77da62_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_RemoveRows
func callbackCustomListModel77da62_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_SetData
func callbackCustomListModel77da62_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomListModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_SetHeaderData
func callbackCustomListModel77da62_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomListModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_SetItemData
func callbackCustomListModel77da62_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewCustomListModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *CustomListModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackCustomListModel77da62_Submit
func callbackCustomListModel77da62_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *CustomListModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackCustomListModel77da62_ColumnsAboutToBeInserted
func callbackCustomListModel77da62_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel77da62_ColumnsAboutToBeMoved
func callbackCustomListModel77da62_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackCustomListModel77da62_ColumnsAboutToBeRemoved
func callbackCustomListModel77da62_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel77da62_ColumnsInserted
func callbackCustomListModel77da62_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel77da62_ColumnsMoved
func callbackCustomListModel77da62_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackCustomListModel77da62_ColumnsRemoved
func callbackCustomListModel77da62_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel77da62_DataChanged
func callbackCustomListModel77da62_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackCustomListModel77da62_FetchMore
func callbackCustomListModel77da62_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewCustomListModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *CustomListModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackCustomListModel77da62_HeaderDataChanged
func callbackCustomListModel77da62_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel77da62_LayoutAboutToBeChanged
func callbackCustomListModel77da62_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackCustomListModel77da62_LayoutChanged
func callbackCustomListModel77da62_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackCustomListModel77da62_ModelAboutToBeReset
func callbackCustomListModel77da62_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackCustomListModel77da62_ModelReset
func callbackCustomListModel77da62_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackCustomListModel77da62_ResetInternalData
func callbackCustomListModel77da62_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewCustomListModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *CustomListModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackCustomListModel77da62_Revert
func callbackCustomListModel77da62_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewCustomListModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *CustomListModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_RevertDefault(ptr.Pointer())
	}
}

//export callbackCustomListModel77da62_RowsAboutToBeInserted
func callbackCustomListModel77da62_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackCustomListModel77da62_RowsAboutToBeMoved
func callbackCustomListModel77da62_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackCustomListModel77da62_RowsAboutToBeRemoved
func callbackCustomListModel77da62_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel77da62_RowsInserted
func callbackCustomListModel77da62_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel77da62_RowsMoved
func callbackCustomListModel77da62_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackCustomListModel77da62_RowsRemoved
func callbackCustomListModel77da62_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel77da62_Sort
func callbackCustomListModel77da62_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewCustomListModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *CustomListModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackCustomListModel77da62_RoleNames
func callbackCustomListModel77da62_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewCustomListModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *CustomListModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.CustomListModel77da62_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackCustomListModel77da62_ItemData
func callbackCustomListModel77da62_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__itemData_newList())
		for k, v := range NewCustomListModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *CustomListModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.CustomListModel77da62_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackCustomListModel77da62_MimeData
func callbackCustomListModel77da62_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewCustomListModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewCustomListModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *CustomListModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.CustomListModel77da62_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel77da62_Buddy
func callbackCustomListModel77da62_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomListModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomListModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel77da62_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel77da62_Parent
func callbackCustomListModel77da62_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomListModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomListModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel77da62_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel77da62_Match
func callbackCustomListModel77da62_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__match_newList())
		for _, v := range NewCustomListModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *CustomListModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.CustomListModel77da62_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackCustomListModel77da62_Span
func callbackCustomListModel77da62_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewCustomListModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomListModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.CustomListModel77da62_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel77da62_MimeTypes
func callbackCustomListModel77da62_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewCustomListModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *CustomListModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.CustomListModel77da62_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackCustomListModel77da62_Data
func callbackCustomListModel77da62_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomListModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *CustomListModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomListModel77da62_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel77da62_HeaderData
func callbackCustomListModel77da62_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomListModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *CustomListModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomListModel77da62_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel77da62_SupportedDragActions
func callbackCustomListModel77da62_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewCustomListModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *CustomListModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomListModel77da62_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomListModel77da62_SupportedDropActions
func callbackCustomListModel77da62_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewCustomListModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *CustomListModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomListModel77da62_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomListModel77da62_CanDropMimeData
func callbackCustomListModel77da62_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_CanFetchMore
func callbackCustomListModel77da62_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_HasChildren
func callbackCustomListModel77da62_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_ColumnCount
func callbackCustomListModel77da62_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomListModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomListModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel77da62_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomListModel77da62_RowCount
func callbackCustomListModel77da62_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomListModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomListModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel77da62_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomListModel77da62_Event
func callbackCustomListModel77da62_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CustomListModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_EventFilter
func callbackCustomListModel77da62_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CustomListModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CustomListModel77da62_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackCustomListModel77da62_ChildEvent
func callbackCustomListModel77da62_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCustomListModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CustomListModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCustomListModel77da62_ConnectNotify
func callbackCustomListModel77da62_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomListModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomListModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomListModel77da62_CustomEvent
func callbackCustomListModel77da62_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewCustomListModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CustomListModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCustomListModel77da62_DeleteLater
func callbackCustomListModel77da62_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewCustomListModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CustomListModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomListModel77da62_Destroyed
func callbackCustomListModel77da62_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackCustomListModel77da62_DisconnectNotify
func callbackCustomListModel77da62_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomListModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomListModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomListModel77da62_ObjectNameChanged
func callbackCustomListModel77da62_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackCustomListModel77da62_TimerEvent
func callbackCustomListModel77da62_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCustomListModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CustomListModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel77da62_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type TxListAccountsModel_ITF interface {
	std_core.QAbstractListModel_ITF
	TxListAccountsModel_PTR() *TxListAccountsModel
}

func (ptr *TxListAccountsModel) TxListAccountsModel_PTR() *TxListAccountsModel {
	return ptr
}

func (ptr *TxListAccountsModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *TxListAccountsModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromTxListAccountsModel(ptr TxListAccountsModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TxListAccountsModel_PTR().Pointer()
	}
	return nil
}

func NewTxListAccountsModelFromPointer(ptr unsafe.Pointer) (n *TxListAccountsModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TxListAccountsModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TxListAccountsModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &TxListAccountsModel{QAbstractListModel: *deduced}

		default:
			n = new(TxListAccountsModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackTxListAccountsModel77da62_Constructor
func callbackTxListAccountsModel77da62_Constructor(ptr unsafe.Pointer) {
	this := NewTxListAccountsModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectAdd(this.add)
	this.init()
}

//export callbackTxListAccountsModel77da62_Add
func callbackTxListAccountsModel77da62_Add(ptr unsafe.Pointer, tx C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "add"); signal != nil {
		signal.(func(string))(cGoUnpackString(tx))
	}

}

func (ptr *TxListAccountsModel) ConnectAdd(f func(tx string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "add") {
			C.TxListAccountsModel77da62_ConnectAdd(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "add"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "add", func(tx string) {
				signal.(func(string))(tx)
				f(tx)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "add", f)
		}
	}
}

func (ptr *TxListAccountsModel) DisconnectAdd() {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_DisconnectAdd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "add")
	}
}

func (ptr *TxListAccountsModel) Add(tx string) {
	if ptr.Pointer() != nil {
		var txC *C.char
		if tx != "" {
			txC = C.CString(tx)
			defer C.free(unsafe.Pointer(txC))
		}
		C.TxListAccountsModel77da62_Add(ptr.Pointer(), C.struct_Moc_PackedString{data: txC, len: C.longlong(len(tx))})
	}
}

func TxListAccountsModel_QRegisterMetaType() int {
	return int(int32(C.TxListAccountsModel77da62_TxListAccountsModel77da62_QRegisterMetaType()))
}

func (ptr *TxListAccountsModel) QRegisterMetaType() int {
	return int(int32(C.TxListAccountsModel77da62_TxListAccountsModel77da62_QRegisterMetaType()))
}

func TxListAccountsModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TxListAccountsModel77da62_TxListAccountsModel77da62_QRegisterMetaType2(typeNameC)))
}

func (ptr *TxListAccountsModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TxListAccountsModel77da62_TxListAccountsModel77da62_QRegisterMetaType2(typeNameC)))
}

func TxListAccountsModel_QmlRegisterType() int {
	return int(int32(C.TxListAccountsModel77da62_TxListAccountsModel77da62_QmlRegisterType()))
}

func (ptr *TxListAccountsModel) QmlRegisterType() int {
	return int(int32(C.TxListAccountsModel77da62_TxListAccountsModel77da62_QmlRegisterType()))
}

func TxListAccountsModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TxListAccountsModel77da62_TxListAccountsModel77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TxListAccountsModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TxListAccountsModel77da62_TxListAccountsModel77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TxListAccountsModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListAccountsModel77da62_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListAccountsModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListAccountsModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListAccountsModel77da62_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListAccountsModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListAccountsModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListAccountsModel77da62_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListAccountsModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListAccountsModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TxListAccountsModel77da62___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TxListAccountsModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.TxListAccountsModel77da62___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TxListAccountsModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListAccountsModel77da62___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListAccountsModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListAccountsModel77da62___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListAccountsModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListAccountsModel77da62___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListAccountsModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListAccountsModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.TxListAccountsModel77da62___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TxListAccountsModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.TxListAccountsModel77da62___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TxListAccountsModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TxListAccountsModel77da62___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TxListAccountsModel) __roleNames_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___roleNames_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.TxListAccountsModel77da62___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TxListAccountsModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TxListAccountsModel77da62___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TxListAccountsModel) __itemData_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___itemData_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.TxListAccountsModel77da62___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TxListAccountsModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListAccountsModel77da62___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListAccountsModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListAccountsModel77da62___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListAccountsModel) __match_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___match_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListAccountsModel77da62___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TxListAccountsModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListAccountsModel77da62_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListAccountsModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListAccountsModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListAccountsModel77da62_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TxListAccountsModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TxListAccountsModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TxListAccountsModel77da62___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TxListAccountsModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListAccountsModel77da62___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListAccountsModel) __findChildren_newList2() unsafe.Pointer {
	return C.TxListAccountsModel77da62___findChildren_newList2(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListAccountsModel77da62___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListAccountsModel) __findChildren_newList3() unsafe.Pointer {
	return C.TxListAccountsModel77da62___findChildren_newList3(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListAccountsModel77da62___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListAccountsModel) __findChildren_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___findChildren_newList(ptr.Pointer())
}

func (ptr *TxListAccountsModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListAccountsModel77da62___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListAccountsModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListAccountsModel) __children_newList() unsafe.Pointer {
	return C.TxListAccountsModel77da62___children_newList(ptr.Pointer())
}

func NewTxListAccountsModel(parent std_core.QObject_ITF) *TxListAccountsModel {
	tmpValue := NewTxListAccountsModelFromPointer(C.TxListAccountsModel77da62_NewTxListAccountsModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTxListAccountsModel77da62_DestroyTxListAccountsModel
func callbackTxListAccountsModel77da62_DestroyTxListAccountsModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~TxListAccountsModel"); signal != nil {
		signal.(func())()
	} else {
		NewTxListAccountsModelFromPointer(ptr).DestroyTxListAccountsModelDefault()
	}
}

func (ptr *TxListAccountsModel) ConnectDestroyTxListAccountsModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~TxListAccountsModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~TxListAccountsModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~TxListAccountsModel", f)
		}
	}
}

func (ptr *TxListAccountsModel) DisconnectDestroyTxListAccountsModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~TxListAccountsModel")
	}
}

func (ptr *TxListAccountsModel) DestroyTxListAccountsModel() {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_DestroyTxListAccountsModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TxListAccountsModel) DestroyTxListAccountsModelDefault() {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_DestroyTxListAccountsModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTxListAccountsModel77da62_DropMimeData
func callbackTxListAccountsModel77da62_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListAccountsModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_Index
func callbackTxListAccountsModel77da62_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewTxListAccountsModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *TxListAccountsModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListAccountsModel77da62_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTxListAccountsModel77da62_Sibling
func callbackTxListAccountsModel77da62_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewTxListAccountsModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *TxListAccountsModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListAccountsModel77da62_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTxListAccountsModel77da62_Flags
func callbackTxListAccountsModel77da62_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewTxListAccountsModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TxListAccountsModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.TxListAccountsModel77da62_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackTxListAccountsModel77da62_InsertColumns
func callbackTxListAccountsModel77da62_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListAccountsModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_InsertRows
func callbackTxListAccountsModel77da62_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListAccountsModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_MoveColumns
func callbackTxListAccountsModel77da62_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TxListAccountsModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_MoveRows
func callbackTxListAccountsModel77da62_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TxListAccountsModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_RemoveColumns
func callbackTxListAccountsModel77da62_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListAccountsModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_RemoveRows
func callbackTxListAccountsModel77da62_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListAccountsModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_SetData
func callbackTxListAccountsModel77da62_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TxListAccountsModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_SetHeaderData
func callbackTxListAccountsModel77da62_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TxListAccountsModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_SetItemData
func callbackTxListAccountsModel77da62_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewTxListAccountsModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *TxListAccountsModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewTxListAccountsModelFromPointer(NewTxListAccountsModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_Submit
func callbackTxListAccountsModel77da62_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *TxListAccountsModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_ColumnsAboutToBeInserted
func callbackTxListAccountsModel77da62_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListAccountsModel77da62_ColumnsAboutToBeMoved
func callbackTxListAccountsModel77da62_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackTxListAccountsModel77da62_ColumnsAboutToBeRemoved
func callbackTxListAccountsModel77da62_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListAccountsModel77da62_ColumnsInserted
func callbackTxListAccountsModel77da62_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListAccountsModel77da62_ColumnsMoved
func callbackTxListAccountsModel77da62_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackTxListAccountsModel77da62_ColumnsRemoved
func callbackTxListAccountsModel77da62_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListAccountsModel77da62_DataChanged
func callbackTxListAccountsModel77da62_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackTxListAccountsModel77da62_FetchMore
func callbackTxListAccountsModel77da62_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewTxListAccountsModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *TxListAccountsModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackTxListAccountsModel77da62_HeaderDataChanged
func callbackTxListAccountsModel77da62_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListAccountsModel77da62_LayoutAboutToBeChanged
func callbackTxListAccountsModel77da62_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackTxListAccountsModel77da62_LayoutChanged
func callbackTxListAccountsModel77da62_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackTxListAccountsModel77da62_ModelAboutToBeReset
func callbackTxListAccountsModel77da62_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackTxListAccountsModel77da62_ModelReset
func callbackTxListAccountsModel77da62_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackTxListAccountsModel77da62_ResetInternalData
func callbackTxListAccountsModel77da62_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewTxListAccountsModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *TxListAccountsModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackTxListAccountsModel77da62_Revert
func callbackTxListAccountsModel77da62_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewTxListAccountsModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *TxListAccountsModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_RevertDefault(ptr.Pointer())
	}
}

//export callbackTxListAccountsModel77da62_RowsAboutToBeInserted
func callbackTxListAccountsModel77da62_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackTxListAccountsModel77da62_RowsAboutToBeMoved
func callbackTxListAccountsModel77da62_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackTxListAccountsModel77da62_RowsAboutToBeRemoved
func callbackTxListAccountsModel77da62_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListAccountsModel77da62_RowsInserted
func callbackTxListAccountsModel77da62_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListAccountsModel77da62_RowsMoved
func callbackTxListAccountsModel77da62_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackTxListAccountsModel77da62_RowsRemoved
func callbackTxListAccountsModel77da62_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTxListAccountsModel77da62_Sort
func callbackTxListAccountsModel77da62_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewTxListAccountsModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *TxListAccountsModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackTxListAccountsModel77da62_RoleNames
func callbackTxListAccountsModel77da62_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTxListAccountsModelFromPointer(NewTxListAccountsModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTxListAccountsModelFromPointer(NewTxListAccountsModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewTxListAccountsModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TxListAccountsModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.TxListAccountsModel77da62_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackTxListAccountsModel77da62_ItemData
func callbackTxListAccountsModel77da62_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTxListAccountsModelFromPointer(NewTxListAccountsModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTxListAccountsModelFromPointer(NewTxListAccountsModelFromPointer(nil).__itemData_newList())
		for k, v := range NewTxListAccountsModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TxListAccountsModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.TxListAccountsModel77da62_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackTxListAccountsModel77da62_MimeData
func callbackTxListAccountsModel77da62_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewTxListAccountsModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewTxListAccountsModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *TxListAccountsModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.TxListAccountsModel77da62_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTxListAccountsModelFromPointer(NewTxListAccountsModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTxListAccountsModel77da62_Buddy
func callbackTxListAccountsModel77da62_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTxListAccountsModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TxListAccountsModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListAccountsModel77da62_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTxListAccountsModel77da62_Parent
func callbackTxListAccountsModel77da62_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTxListAccountsModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TxListAccountsModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TxListAccountsModel77da62_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTxListAccountsModel77da62_Match
func callbackTxListAccountsModel77da62_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTxListAccountsModelFromPointer(NewTxListAccountsModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTxListAccountsModelFromPointer(NewTxListAccountsModelFromPointer(nil).__match_newList())
		for _, v := range NewTxListAccountsModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TxListAccountsModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewTxListAccountsModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.TxListAccountsModel77da62_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackTxListAccountsModel77da62_Span
func callbackTxListAccountsModel77da62_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewTxListAccountsModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TxListAccountsModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.TxListAccountsModel77da62_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackTxListAccountsModel77da62_MimeTypes
func callbackTxListAccountsModel77da62_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewTxListAccountsModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *TxListAccountsModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.TxListAccountsModel77da62_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackTxListAccountsModel77da62_Data
func callbackTxListAccountsModel77da62_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTxListAccountsModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *TxListAccountsModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TxListAccountsModel77da62_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTxListAccountsModel77da62_HeaderData
func callbackTxListAccountsModel77da62_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTxListAccountsModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *TxListAccountsModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TxListAccountsModel77da62_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTxListAccountsModel77da62_SupportedDragActions
func callbackTxListAccountsModel77da62_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewTxListAccountsModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *TxListAccountsModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TxListAccountsModel77da62_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTxListAccountsModel77da62_SupportedDropActions
func callbackTxListAccountsModel77da62_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewTxListAccountsModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *TxListAccountsModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TxListAccountsModel77da62_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTxListAccountsModel77da62_CanDropMimeData
func callbackTxListAccountsModel77da62_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListAccountsModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_CanFetchMore
func callbackTxListAccountsModel77da62_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListAccountsModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_HasChildren
func callbackTxListAccountsModel77da62_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TxListAccountsModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_ColumnCount
func callbackTxListAccountsModel77da62_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTxListAccountsModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TxListAccountsModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListAccountsModel77da62_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTxListAccountsModel77da62_RowCount
func callbackTxListAccountsModel77da62_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTxListAccountsModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TxListAccountsModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TxListAccountsModel77da62_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTxListAccountsModel77da62_Event
func callbackTxListAccountsModel77da62_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TxListAccountsModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_EventFilter
func callbackTxListAccountsModel77da62_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListAccountsModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TxListAccountsModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListAccountsModel77da62_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTxListAccountsModel77da62_ChildEvent
func callbackTxListAccountsModel77da62_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTxListAccountsModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TxListAccountsModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTxListAccountsModel77da62_ConnectNotify
func callbackTxListAccountsModel77da62_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTxListAccountsModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TxListAccountsModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTxListAccountsModel77da62_CustomEvent
func callbackTxListAccountsModel77da62_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewTxListAccountsModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TxListAccountsModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTxListAccountsModel77da62_DeleteLater
func callbackTxListAccountsModel77da62_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewTxListAccountsModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TxListAccountsModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTxListAccountsModel77da62_Destroyed
func callbackTxListAccountsModel77da62_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackTxListAccountsModel77da62_DisconnectNotify
func callbackTxListAccountsModel77da62_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTxListAccountsModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TxListAccountsModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTxListAccountsModel77da62_ObjectNameChanged
func callbackTxListAccountsModel77da62_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackTxListAccountsModel77da62_TimerEvent
func callbackTxListAccountsModel77da62_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTxListAccountsModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TxListAccountsModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TxListAccountsModel77da62_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type LoginCtx_ITF interface {
	std_core.QObject_ITF
	LoginCtx_PTR() *LoginCtx
}

func (ptr *LoginCtx) LoginCtx_PTR() *LoginCtx {
	return ptr
}

func (ptr *LoginCtx) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *LoginCtx) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromLoginCtx(ptr LoginCtx_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.LoginCtx_PTR().Pointer()
	}
	return nil
}

func NewLoginCtxFromPointer(ptr unsafe.Pointer) (n *LoginCtx) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(LoginCtx)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *LoginCtx:
			n = deduced

		case *std_core.QObject:
			n = &LoginCtx{QObject: *deduced}

		default:
			n = new(LoginCtx)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackLoginCtx77da62_Constructor
func callbackLoginCtx77da62_Constructor(ptr unsafe.Pointer) {
	this := NewLoginCtxFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClicked(this.clicked)
	this.ConnectEdited(this.edited)
}

//export callbackLoginCtx77da62_Clicked
func callbackLoginCtx77da62_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *LoginCtx) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.LoginCtx77da62_ConnectClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clicked", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *LoginCtx) Clicked() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_Clicked(ptr.Pointer())
	}
}

//export callbackLoginCtx77da62_Edited
func callbackLoginCtx77da62_Edited(ptr unsafe.Pointer, b C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "edited"); signal != nil {
		signal.(func(string))(cGoUnpackString(b))
	}

}

func (ptr *LoginCtx) ConnectEdited(f func(b string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "edited") {
			C.LoginCtx77da62_ConnectEdited(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "edited"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "edited", func(b string) {
				signal.(func(string))(b)
				f(b)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "edited", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectEdited() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DisconnectEdited(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "edited")
	}
}

func (ptr *LoginCtx) Edited(b string) {
	if ptr.Pointer() != nil {
		var bC *C.char
		if b != "" {
			bC = C.CString(b)
			defer C.free(unsafe.Pointer(bC))
		}
		C.LoginCtx77da62_Edited(ptr.Pointer(), C.struct_Moc_PackedString{data: bC, len: C.longlong(len(b))})
	}
}

//export callbackLoginCtx77da62_Remote
func callbackLoginCtx77da62_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "remote"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewLoginCtxFromPointer(ptr).RemoteDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *LoginCtx) ConnectRemote(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "remote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remote", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remote", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "remote")
	}
}

func (ptr *LoginCtx) Remote() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *LoginCtx) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackLoginCtx77da62_SetRemote
func callbackLoginCtx77da62_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setRemote"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	} else {
		NewLoginCtxFromPointer(ptr).SetRemoteDefault(cGoUnpackString(remote))
	}
}

func (ptr *LoginCtx) ConnectSetRemote(f func(remote string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRemote"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRemote", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectSetRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRemote")
	}
}

func (ptr *LoginCtx) SetRemote(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.LoginCtx77da62_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *LoginCtx) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.LoginCtx77da62_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackLoginCtx77da62_RemoteChanged
func callbackLoginCtx77da62_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *LoginCtx) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.LoginCtx77da62_ConnectRemoteChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "remoteChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", func(remote string) {
				signal.(func(string))(remote)
				f(remote)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remoteChanged", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectRemoteChanged() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DisconnectRemoteChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remoteChanged")
	}
}

func (ptr *LoginCtx) RemoteChanged(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.LoginCtx77da62_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackLoginCtx77da62_Transport
func callbackLoginCtx77da62_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "transport"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewLoginCtxFromPointer(ptr).TransportDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *LoginCtx) ConnectTransport(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "transport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transport", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transport", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transport")
	}
}

func (ptr *LoginCtx) Transport() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *LoginCtx) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackLoginCtx77da62_SetTransport
func callbackLoginCtx77da62_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTransport"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	} else {
		NewLoginCtxFromPointer(ptr).SetTransportDefault(cGoUnpackString(transport))
	}
}

func (ptr *LoginCtx) ConnectSetTransport(f func(transport string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTransport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTransport", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectSetTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransport")
	}
}

func (ptr *LoginCtx) SetTransport(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.LoginCtx77da62_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *LoginCtx) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.LoginCtx77da62_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackLoginCtx77da62_TransportChanged
func callbackLoginCtx77da62_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *LoginCtx) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.LoginCtx77da62_ConnectTransportChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transportChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", func(transport string) {
				signal.(func(string))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transportChanged", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectTransportChanged() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DisconnectTransportChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transportChanged")
	}
}

func (ptr *LoginCtx) TransportChanged(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.LoginCtx77da62_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackLoginCtx77da62_Endpoint
func callbackLoginCtx77da62_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "endpoint"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewLoginCtxFromPointer(ptr).EndpointDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *LoginCtx) ConnectEndpoint(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpoint", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endpoint")
	}
}

func (ptr *LoginCtx) Endpoint() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *LoginCtx) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackLoginCtx77da62_SetEndpoint
func callbackLoginCtx77da62_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setEndpoint"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	} else {
		NewLoginCtxFromPointer(ptr).SetEndpointDefault(cGoUnpackString(endpoint))
	}
}

func (ptr *LoginCtx) ConnectSetEndpoint(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEndpoint"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEndpoint", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectSetEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEndpoint")
	}
}

func (ptr *LoginCtx) SetEndpoint(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.LoginCtx77da62_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *LoginCtx) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.LoginCtx77da62_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackLoginCtx77da62_EndpointChanged
func callbackLoginCtx77da62_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *LoginCtx) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.LoginCtx77da62_ConnectEndpointChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "endpointChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", func(endpoint string) {
				signal.(func(string))(endpoint)
				f(endpoint)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endpointChanged", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectEndpointChanged() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DisconnectEndpointChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "endpointChanged")
	}
}

func (ptr *LoginCtx) EndpointChanged(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.LoginCtx77da62_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackLoginCtx77da62_Gopath
func callbackLoginCtx77da62_Gopath(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "gopath"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewLoginCtxFromPointer(ptr).GopathDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *LoginCtx) ConnectGopath(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "gopath"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gopath", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gopath", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectGopath() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "gopath")
	}
}

func (ptr *LoginCtx) Gopath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_Gopath(ptr.Pointer()))
	}
	return ""
}

func (ptr *LoginCtx) GopathDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_GopathDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackLoginCtx77da62_SetGopath
func callbackLoginCtx77da62_SetGopath(ptr unsafe.Pointer, gopath C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setGopath"); signal != nil {
		signal.(func(string))(cGoUnpackString(gopath))
	} else {
		NewLoginCtxFromPointer(ptr).SetGopathDefault(cGoUnpackString(gopath))
	}
}

func (ptr *LoginCtx) ConnectSetGopath(f func(gopath string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGopath"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setGopath", func(gopath string) {
				signal.(func(string))(gopath)
				f(gopath)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGopath", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectSetGopath() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGopath")
	}
}

func (ptr *LoginCtx) SetGopath(gopath string) {
	if ptr.Pointer() != nil {
		var gopathC *C.char
		if gopath != "" {
			gopathC = C.CString(gopath)
			defer C.free(unsafe.Pointer(gopathC))
		}
		C.LoginCtx77da62_SetGopath(ptr.Pointer(), C.struct_Moc_PackedString{data: gopathC, len: C.longlong(len(gopath))})
	}
}

func (ptr *LoginCtx) SetGopathDefault(gopath string) {
	if ptr.Pointer() != nil {
		var gopathC *C.char
		if gopath != "" {
			gopathC = C.CString(gopath)
			defer C.free(unsafe.Pointer(gopathC))
		}
		C.LoginCtx77da62_SetGopathDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: gopathC, len: C.longlong(len(gopath))})
	}
}

//export callbackLoginCtx77da62_GopathChanged
func callbackLoginCtx77da62_GopathChanged(ptr unsafe.Pointer, gopath C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "gopathChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(gopath))
	}

}

func (ptr *LoginCtx) ConnectGopathChanged(f func(gopath string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gopathChanged") {
			C.LoginCtx77da62_ConnectGopathChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gopathChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gopathChanged", func(gopath string) {
				signal.(func(string))(gopath)
				f(gopath)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gopathChanged", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectGopathChanged() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DisconnectGopathChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gopathChanged")
	}
}

func (ptr *LoginCtx) GopathChanged(gopath string) {
	if ptr.Pointer() != nil {
		var gopathC *C.char
		if gopath != "" {
			gopathC = C.CString(gopath)
			defer C.free(unsafe.Pointer(gopathC))
		}
		C.LoginCtx77da62_GopathChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: gopathC, len: C.longlong(len(gopath))})
	}
}

//export callbackLoginCtx77da62_BinaryHash
func callbackLoginCtx77da62_BinaryHash(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "binaryHash"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewLoginCtxFromPointer(ptr).BinaryHashDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *LoginCtx) ConnectBinaryHash(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "binaryHash"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "binaryHash", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "binaryHash", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectBinaryHash() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "binaryHash")
	}
}

func (ptr *LoginCtx) BinaryHash() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_BinaryHash(ptr.Pointer()))
	}
	return ""
}

func (ptr *LoginCtx) BinaryHashDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.LoginCtx77da62_BinaryHashDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackLoginCtx77da62_SetBinaryHash
func callbackLoginCtx77da62_SetBinaryHash(ptr unsafe.Pointer, binaryHash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setBinaryHash"); signal != nil {
		signal.(func(string))(cGoUnpackString(binaryHash))
	} else {
		NewLoginCtxFromPointer(ptr).SetBinaryHashDefault(cGoUnpackString(binaryHash))
	}
}

func (ptr *LoginCtx) ConnectSetBinaryHash(f func(binaryHash string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setBinaryHash"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setBinaryHash", func(binaryHash string) {
				signal.(func(string))(binaryHash)
				f(binaryHash)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setBinaryHash", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectSetBinaryHash() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setBinaryHash")
	}
}

func (ptr *LoginCtx) SetBinaryHash(binaryHash string) {
	if ptr.Pointer() != nil {
		var binaryHashC *C.char
		if binaryHash != "" {
			binaryHashC = C.CString(binaryHash)
			defer C.free(unsafe.Pointer(binaryHashC))
		}
		C.LoginCtx77da62_SetBinaryHash(ptr.Pointer(), C.struct_Moc_PackedString{data: binaryHashC, len: C.longlong(len(binaryHash))})
	}
}

func (ptr *LoginCtx) SetBinaryHashDefault(binaryHash string) {
	if ptr.Pointer() != nil {
		var binaryHashC *C.char
		if binaryHash != "" {
			binaryHashC = C.CString(binaryHash)
			defer C.free(unsafe.Pointer(binaryHashC))
		}
		C.LoginCtx77da62_SetBinaryHashDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: binaryHashC, len: C.longlong(len(binaryHash))})
	}
}

//export callbackLoginCtx77da62_BinaryHashChanged
func callbackLoginCtx77da62_BinaryHashChanged(ptr unsafe.Pointer, binaryHash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "binaryHashChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(binaryHash))
	}

}

func (ptr *LoginCtx) ConnectBinaryHashChanged(f func(binaryHash string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "binaryHashChanged") {
			C.LoginCtx77da62_ConnectBinaryHashChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "binaryHashChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "binaryHashChanged", func(binaryHash string) {
				signal.(func(string))(binaryHash)
				f(binaryHash)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "binaryHashChanged", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectBinaryHashChanged() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DisconnectBinaryHashChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "binaryHashChanged")
	}
}

func (ptr *LoginCtx) BinaryHashChanged(binaryHash string) {
	if ptr.Pointer() != nil {
		var binaryHashC *C.char
		if binaryHash != "" {
			binaryHashC = C.CString(binaryHash)
			defer C.free(unsafe.Pointer(binaryHashC))
		}
		C.LoginCtx77da62_BinaryHashChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: binaryHashC, len: C.longlong(len(binaryHash))})
	}
}

//export callbackLoginCtx77da62_IsValid
func callbackLoginCtx77da62_IsValid(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isValid"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewLoginCtxFromPointer(ptr).IsValidDefault())))
}

func (ptr *LoginCtx) ConnectIsValid(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isValid"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isValid", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isValid", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectIsValid() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isValid")
	}
}

func (ptr *LoginCtx) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.LoginCtx77da62_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *LoginCtx) IsValidDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.LoginCtx77da62_IsValidDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackLoginCtx77da62_SetIsValid
func callbackLoginCtx77da62_SetIsValid(ptr unsafe.Pointer, isValid C.char) {
	if signal := qt.GetSignal(ptr, "setIsValid"); signal != nil {
		signal.(func(bool))(int8(isValid) != 0)
	} else {
		NewLoginCtxFromPointer(ptr).SetIsValidDefault(int8(isValid) != 0)
	}
}

func (ptr *LoginCtx) ConnectSetIsValid(f func(isValid bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setIsValid"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setIsValid", func(isValid bool) {
				signal.(func(bool))(isValid)
				f(isValid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setIsValid", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectSetIsValid() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setIsValid")
	}
}

func (ptr *LoginCtx) SetIsValid(isValid bool) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_SetIsValid(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isValid))))
	}
}

func (ptr *LoginCtx) SetIsValidDefault(isValid bool) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_SetIsValidDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isValid))))
	}
}

//export callbackLoginCtx77da62_IsValidChanged
func callbackLoginCtx77da62_IsValidChanged(ptr unsafe.Pointer, isValid C.char) {
	if signal := qt.GetSignal(ptr, "isValidChanged"); signal != nil {
		signal.(func(bool))(int8(isValid) != 0)
	}

}

func (ptr *LoginCtx) ConnectIsValidChanged(f func(isValid bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "isValidChanged") {
			C.LoginCtx77da62_ConnectIsValidChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "isValidChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isValidChanged", func(isValid bool) {
				signal.(func(bool))(isValid)
				f(isValid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isValidChanged", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectIsValidChanged() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DisconnectIsValidChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "isValidChanged")
	}
}

func (ptr *LoginCtx) IsValidChanged(isValid bool) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_IsValidChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isValid))))
	}
}

func LoginCtx_QRegisterMetaType() int {
	return int(int32(C.LoginCtx77da62_LoginCtx77da62_QRegisterMetaType()))
}

func (ptr *LoginCtx) QRegisterMetaType() int {
	return int(int32(C.LoginCtx77da62_LoginCtx77da62_QRegisterMetaType()))
}

func LoginCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.LoginCtx77da62_LoginCtx77da62_QRegisterMetaType2(typeNameC)))
}

func (ptr *LoginCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.LoginCtx77da62_LoginCtx77da62_QRegisterMetaType2(typeNameC)))
}

func LoginCtx_QmlRegisterType() int {
	return int(int32(C.LoginCtx77da62_LoginCtx77da62_QmlRegisterType()))
}

func (ptr *LoginCtx) QmlRegisterType() int {
	return int(int32(C.LoginCtx77da62_LoginCtx77da62_QmlRegisterType()))
}

func LoginCtx_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.LoginCtx77da62_LoginCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *LoginCtx) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.LoginCtx77da62_LoginCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *LoginCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.LoginCtx77da62___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *LoginCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *LoginCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.LoginCtx77da62___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *LoginCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.LoginCtx77da62___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *LoginCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *LoginCtx) __findChildren_newList2() unsafe.Pointer {
	return C.LoginCtx77da62___findChildren_newList2(ptr.Pointer())
}

func (ptr *LoginCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.LoginCtx77da62___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *LoginCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *LoginCtx) __findChildren_newList3() unsafe.Pointer {
	return C.LoginCtx77da62___findChildren_newList3(ptr.Pointer())
}

func (ptr *LoginCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.LoginCtx77da62___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *LoginCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *LoginCtx) __findChildren_newList() unsafe.Pointer {
	return C.LoginCtx77da62___findChildren_newList(ptr.Pointer())
}

func (ptr *LoginCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.LoginCtx77da62___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *LoginCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *LoginCtx) __children_newList() unsafe.Pointer {
	return C.LoginCtx77da62___children_newList(ptr.Pointer())
}

func NewLoginCtx(parent std_core.QObject_ITF) *LoginCtx {
	tmpValue := NewLoginCtxFromPointer(C.LoginCtx77da62_NewLoginCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackLoginCtx77da62_DestroyLoginCtx
func callbackLoginCtx77da62_DestroyLoginCtx(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~LoginCtx"); signal != nil {
		signal.(func())()
	} else {
		NewLoginCtxFromPointer(ptr).DestroyLoginCtxDefault()
	}
}

func (ptr *LoginCtx) ConnectDestroyLoginCtx(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~LoginCtx"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~LoginCtx", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~LoginCtx", f)
		}
	}
}

func (ptr *LoginCtx) DisconnectDestroyLoginCtx() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~LoginCtx")
	}
}

func (ptr *LoginCtx) DestroyLoginCtx() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DestroyLoginCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *LoginCtx) DestroyLoginCtxDefault() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DestroyLoginCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackLoginCtx77da62_Event
func callbackLoginCtx77da62_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewLoginCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *LoginCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.LoginCtx77da62_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackLoginCtx77da62_EventFilter
func callbackLoginCtx77da62_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewLoginCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *LoginCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.LoginCtx77da62_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackLoginCtx77da62_ChildEvent
func callbackLoginCtx77da62_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewLoginCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *LoginCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackLoginCtx77da62_ConnectNotify
func callbackLoginCtx77da62_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewLoginCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *LoginCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackLoginCtx77da62_CustomEvent
func callbackLoginCtx77da62_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewLoginCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *LoginCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackLoginCtx77da62_DeleteLater
func callbackLoginCtx77da62_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewLoginCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *LoginCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackLoginCtx77da62_Destroyed
func callbackLoginCtx77da62_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackLoginCtx77da62_DisconnectNotify
func callbackLoginCtx77da62_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewLoginCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *LoginCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackLoginCtx77da62_ObjectNameChanged
func callbackLoginCtx77da62_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackLoginCtx77da62_TimerEvent
func callbackLoginCtx77da62_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewLoginCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *LoginCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.LoginCtx77da62_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type TxListCtx_ITF interface {
	std_core.QObject_ITF
	TxListCtx_PTR() *TxListCtx
}

func (ptr *TxListCtx) TxListCtx_PTR() *TxListCtx {
	return ptr
}

func (ptr *TxListCtx) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *TxListCtx) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromTxListCtx(ptr TxListCtx_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TxListCtx_PTR().Pointer()
	}
	return nil
}

func NewTxListCtxFromPointer(ptr unsafe.Pointer) (n *TxListCtx) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TxListCtx)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TxListCtx:
			n = deduced

		case *std_core.QObject:
			n = &TxListCtx{QObject: *deduced}

		default:
			n = new(TxListCtx)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackTxListCtx77da62_Constructor
func callbackTxListCtx77da62_Constructor(ptr unsafe.Pointer) {
	this := NewTxListCtxFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClicked(this.clicked)
	this.init()
}

//export callbackTxListCtx77da62_Clicked
func callbackTxListCtx77da62_Clicked(ptr unsafe.Pointer, b C.int) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(int))(int(int32(b)))
	}

}

func (ptr *TxListCtx) ConnectClicked(f func(b int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.TxListCtx77da62_ConnectClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clicked", func(b int) {
				signal.(func(int))(b)
				f(b)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", f)
		}
	}
}

func (ptr *TxListCtx) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *TxListCtx) Clicked(b int) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_Clicked(ptr.Pointer(), C.int(int32(b)))
	}
}

//export callbackTxListCtx77da62_ShortenAddress
func callbackTxListCtx77da62_ShortenAddress(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "shortenAddress"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTxListCtxFromPointer(ptr).ShortenAddressDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TxListCtx) ConnectShortenAddress(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "shortenAddress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "shortenAddress", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shortenAddress", f)
		}
	}
}

func (ptr *TxListCtx) DisconnectShortenAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "shortenAddress")
	}
}

func (ptr *TxListCtx) ShortenAddress() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TxListCtx77da62_ShortenAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *TxListCtx) ShortenAddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TxListCtx77da62_ShortenAddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTxListCtx77da62_SetShortenAddress
func callbackTxListCtx77da62_SetShortenAddress(ptr unsafe.Pointer, shortenAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setShortenAddress"); signal != nil {
		signal.(func(string))(cGoUnpackString(shortenAddress))
	} else {
		NewTxListCtxFromPointer(ptr).SetShortenAddressDefault(cGoUnpackString(shortenAddress))
	}
}

func (ptr *TxListCtx) ConnectSetShortenAddress(f func(shortenAddress string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setShortenAddress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setShortenAddress", func(shortenAddress string) {
				signal.(func(string))(shortenAddress)
				f(shortenAddress)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setShortenAddress", f)
		}
	}
}

func (ptr *TxListCtx) DisconnectSetShortenAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setShortenAddress")
	}
}

func (ptr *TxListCtx) SetShortenAddress(shortenAddress string) {
	if ptr.Pointer() != nil {
		var shortenAddressC *C.char
		if shortenAddress != "" {
			shortenAddressC = C.CString(shortenAddress)
			defer C.free(unsafe.Pointer(shortenAddressC))
		}
		C.TxListCtx77da62_SetShortenAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: shortenAddressC, len: C.longlong(len(shortenAddress))})
	}
}

func (ptr *TxListCtx) SetShortenAddressDefault(shortenAddress string) {
	if ptr.Pointer() != nil {
		var shortenAddressC *C.char
		if shortenAddress != "" {
			shortenAddressC = C.CString(shortenAddress)
			defer C.free(unsafe.Pointer(shortenAddressC))
		}
		C.TxListCtx77da62_SetShortenAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: shortenAddressC, len: C.longlong(len(shortenAddress))})
	}
}

//export callbackTxListCtx77da62_ShortenAddressChanged
func callbackTxListCtx77da62_ShortenAddressChanged(ptr unsafe.Pointer, shortenAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "shortenAddressChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(shortenAddress))
	}

}

func (ptr *TxListCtx) ConnectShortenAddressChanged(f func(shortenAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shortenAddressChanged") {
			C.TxListCtx77da62_ConnectShortenAddressChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shortenAddressChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "shortenAddressChanged", func(shortenAddress string) {
				signal.(func(string))(shortenAddress)
				f(shortenAddress)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shortenAddressChanged", f)
		}
	}
}

func (ptr *TxListCtx) DisconnectShortenAddressChanged() {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_DisconnectShortenAddressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shortenAddressChanged")
	}
}

func (ptr *TxListCtx) ShortenAddressChanged(shortenAddress string) {
	if ptr.Pointer() != nil {
		var shortenAddressC *C.char
		if shortenAddress != "" {
			shortenAddressC = C.CString(shortenAddress)
			defer C.free(unsafe.Pointer(shortenAddressC))
		}
		C.TxListCtx77da62_ShortenAddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: shortenAddressC, len: C.longlong(len(shortenAddress))})
	}
}

//export callbackTxListCtx77da62_SelectedSrc
func callbackTxListCtx77da62_SelectedSrc(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "selectedSrc"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTxListCtxFromPointer(ptr).SelectedSrcDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TxListCtx) ConnectSelectedSrc(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectedSrc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectedSrc", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectedSrc", f)
		}
	}
}

func (ptr *TxListCtx) DisconnectSelectedSrc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectedSrc")
	}
}

func (ptr *TxListCtx) SelectedSrc() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TxListCtx77da62_SelectedSrc(ptr.Pointer()))
	}
	return ""
}

func (ptr *TxListCtx) SelectedSrcDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TxListCtx77da62_SelectedSrcDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTxListCtx77da62_SetSelectedSrc
func callbackTxListCtx77da62_SetSelectedSrc(ptr unsafe.Pointer, selectedSrc C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setSelectedSrc"); signal != nil {
		signal.(func(string))(cGoUnpackString(selectedSrc))
	} else {
		NewTxListCtxFromPointer(ptr).SetSelectedSrcDefault(cGoUnpackString(selectedSrc))
	}
}

func (ptr *TxListCtx) ConnectSetSelectedSrc(f func(selectedSrc string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSelectedSrc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setSelectedSrc", func(selectedSrc string) {
				signal.(func(string))(selectedSrc)
				f(selectedSrc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSelectedSrc", f)
		}
	}
}

func (ptr *TxListCtx) DisconnectSetSelectedSrc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSelectedSrc")
	}
}

func (ptr *TxListCtx) SetSelectedSrc(selectedSrc string) {
	if ptr.Pointer() != nil {
		var selectedSrcC *C.char
		if selectedSrc != "" {
			selectedSrcC = C.CString(selectedSrc)
			defer C.free(unsafe.Pointer(selectedSrcC))
		}
		C.TxListCtx77da62_SetSelectedSrc(ptr.Pointer(), C.struct_Moc_PackedString{data: selectedSrcC, len: C.longlong(len(selectedSrc))})
	}
}

func (ptr *TxListCtx) SetSelectedSrcDefault(selectedSrc string) {
	if ptr.Pointer() != nil {
		var selectedSrcC *C.char
		if selectedSrc != "" {
			selectedSrcC = C.CString(selectedSrc)
			defer C.free(unsafe.Pointer(selectedSrcC))
		}
		C.TxListCtx77da62_SetSelectedSrcDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: selectedSrcC, len: C.longlong(len(selectedSrc))})
	}
}

//export callbackTxListCtx77da62_SelectedSrcChanged
func callbackTxListCtx77da62_SelectedSrcChanged(ptr unsafe.Pointer, selectedSrc C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "selectedSrcChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(selectedSrc))
	}

}

func (ptr *TxListCtx) ConnectSelectedSrcChanged(f func(selectedSrc string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectedSrcChanged") {
			C.TxListCtx77da62_ConnectSelectedSrcChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectedSrcChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectedSrcChanged", func(selectedSrc string) {
				signal.(func(string))(selectedSrc)
				f(selectedSrc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectedSrcChanged", f)
		}
	}
}

func (ptr *TxListCtx) DisconnectSelectedSrcChanged() {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_DisconnectSelectedSrcChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectedSrcChanged")
	}
}

func (ptr *TxListCtx) SelectedSrcChanged(selectedSrc string) {
	if ptr.Pointer() != nil {
		var selectedSrcC *C.char
		if selectedSrc != "" {
			selectedSrcC = C.CString(selectedSrc)
			defer C.free(unsafe.Pointer(selectedSrcC))
		}
		C.TxListCtx77da62_SelectedSrcChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: selectedSrcC, len: C.longlong(len(selectedSrc))})
	}
}

func TxListCtx_QRegisterMetaType() int {
	return int(int32(C.TxListCtx77da62_TxListCtx77da62_QRegisterMetaType()))
}

func (ptr *TxListCtx) QRegisterMetaType() int {
	return int(int32(C.TxListCtx77da62_TxListCtx77da62_QRegisterMetaType()))
}

func TxListCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TxListCtx77da62_TxListCtx77da62_QRegisterMetaType2(typeNameC)))
}

func (ptr *TxListCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TxListCtx77da62_TxListCtx77da62_QRegisterMetaType2(typeNameC)))
}

func TxListCtx_QmlRegisterType() int {
	return int(int32(C.TxListCtx77da62_TxListCtx77da62_QmlRegisterType()))
}

func (ptr *TxListCtx) QmlRegisterType() int {
	return int(int32(C.TxListCtx77da62_TxListCtx77da62_QmlRegisterType()))
}

func TxListCtx_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TxListCtx77da62_TxListCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TxListCtx) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TxListCtx77da62_TxListCtx77da62_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TxListCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TxListCtx77da62___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TxListCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TxListCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TxListCtx77da62___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TxListCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListCtx77da62___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListCtx) __findChildren_newList2() unsafe.Pointer {
	return C.TxListCtx77da62___findChildren_newList2(ptr.Pointer())
}

func (ptr *TxListCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListCtx77da62___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListCtx) __findChildren_newList3() unsafe.Pointer {
	return C.TxListCtx77da62___findChildren_newList3(ptr.Pointer())
}

func (ptr *TxListCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListCtx77da62___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListCtx) __findChildren_newList() unsafe.Pointer {
	return C.TxListCtx77da62___findChildren_newList(ptr.Pointer())
}

func (ptr *TxListCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TxListCtx77da62___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TxListCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TxListCtx) __children_newList() unsafe.Pointer {
	return C.TxListCtx77da62___children_newList(ptr.Pointer())
}

func NewTxListCtx(parent std_core.QObject_ITF) *TxListCtx {
	tmpValue := NewTxListCtxFromPointer(C.TxListCtx77da62_NewTxListCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTxListCtx77da62_DestroyTxListCtx
func callbackTxListCtx77da62_DestroyTxListCtx(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~TxListCtx"); signal != nil {
		signal.(func())()
	} else {
		NewTxListCtxFromPointer(ptr).DestroyTxListCtxDefault()
	}
}

func (ptr *TxListCtx) ConnectDestroyTxListCtx(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~TxListCtx"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~TxListCtx", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~TxListCtx", f)
		}
	}
}

func (ptr *TxListCtx) DisconnectDestroyTxListCtx() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~TxListCtx")
	}
}

func (ptr *TxListCtx) DestroyTxListCtx() {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_DestroyTxListCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TxListCtx) DestroyTxListCtxDefault() {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_DestroyTxListCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTxListCtx77da62_Event
func callbackTxListCtx77da62_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TxListCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListCtx77da62_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTxListCtx77da62_EventFilter
func callbackTxListCtx77da62_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTxListCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TxListCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TxListCtx77da62_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTxListCtx77da62_ChildEvent
func callbackTxListCtx77da62_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTxListCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TxListCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTxListCtx77da62_ConnectNotify
func callbackTxListCtx77da62_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTxListCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TxListCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTxListCtx77da62_CustomEvent
func callbackTxListCtx77da62_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewTxListCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TxListCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTxListCtx77da62_DeleteLater
func callbackTxListCtx77da62_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewTxListCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TxListCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTxListCtx77da62_Destroyed
func callbackTxListCtx77da62_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackTxListCtx77da62_DisconnectNotify
func callbackTxListCtx77da62_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTxListCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TxListCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTxListCtx77da62_ObjectNameChanged
func callbackTxListCtx77da62_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackTxListCtx77da62_TimerEvent
func callbackTxListCtx77da62_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTxListCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TxListCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TxListCtx77da62_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
