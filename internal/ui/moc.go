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

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

//export callbackCustomListModel721036_Constructor
func callbackCustomListModel721036_Constructor(ptr unsafe.Pointer) {
	this := NewCustomListModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectRemove(this.remove)
	this.ConnectAdd(this.add)
	this.ConnectEdit(this.edit)
	this.init()
}

//export callbackCustomListModel721036_Remove
func callbackCustomListModel721036_Remove(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "remove"); signal != nil {
		signal.(func())()
	}

}

func (ptr *CustomListModel) ConnectRemove(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remove") {
			C.CustomListModel721036_ConnectRemove(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "remove"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remove", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remove", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectRemove() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_DisconnectRemove(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remove")
	}
}

func (ptr *CustomListModel) Remove() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_Remove(ptr.Pointer())
	}
}

//export callbackCustomListModel721036_Add
func callbackCustomListModel721036_Add(ptr unsafe.Pointer, obj C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "add"); signal != nil {
		signal.(func([]*std_core.QVariant))(func(l C.struct_Moc_PackedList) []*std_core.QVariant {
			out := make([]*std_core.QVariant, int(l.len))
			tmpList := NewCustomListModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__add_obj_atList(i)
			}
			return out
		}(obj))
	}

}

func (ptr *CustomListModel) ConnectAdd(f func(obj []*std_core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "add") {
			C.CustomListModel721036_ConnectAdd(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "add"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "add", func(obj []*std_core.QVariant) {
				signal.(func([]*std_core.QVariant))(obj)
				f(obj)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "add", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectAdd() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_DisconnectAdd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "add")
	}
}

func (ptr *CustomListModel) Add(obj []*std_core.QVariant) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_Add(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__add_obj_newList())
			for _, v := range obj {
				tmpList.__add_obj_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackCustomListModel721036_Edit
func callbackCustomListModel721036_Edit(ptr unsafe.Pointer, address C.struct_Moc_PackedString, checked C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "edit"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(address), cGoUnpackString(checked))
	}

}

func (ptr *CustomListModel) ConnectEdit(f func(address string, checked string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "edit") {
			C.CustomListModel721036_ConnectEdit(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "edit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "edit", func(address string, checked string) {
				signal.(func(string, string))(address, checked)
				f(address, checked)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "edit", f)
		}
	}
}

func (ptr *CustomListModel) DisconnectEdit() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_DisconnectEdit(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "edit")
	}
}

func (ptr *CustomListModel) Edit(address string, checked string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		var checkedC *C.char
		if checked != "" {
			checkedC = C.CString(checked)
			defer C.free(unsafe.Pointer(checkedC))
		}
		C.CustomListModel721036_Edit(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_Moc_PackedString{data: checkedC, len: C.longlong(len(checked))})
	}
}

func CustomListModel_QRegisterMetaType() int {
	return int(int32(C.CustomListModel721036_CustomListModel721036_QRegisterMetaType()))
}

func (ptr *CustomListModel) QRegisterMetaType() int {
	return int(int32(C.CustomListModel721036_CustomListModel721036_QRegisterMetaType()))
}

func CustomListModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomListModel721036_CustomListModel721036_QRegisterMetaType2(typeNameC)))
}

func (ptr *CustomListModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomListModel721036_CustomListModel721036_QRegisterMetaType2(typeNameC)))
}

func CustomListModel_QmlRegisterType() int {
	return int(int32(C.CustomListModel721036_CustomListModel721036_QmlRegisterType()))
}

func (ptr *CustomListModel) QmlRegisterType() int {
	return int(int32(C.CustomListModel721036_CustomListModel721036_QmlRegisterType()))
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
	return int(int32(C.CustomListModel721036_CustomListModel721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.CustomListModel721036_CustomListModel721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomListModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel721036_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.CustomListModel721036_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel721036_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomListModel721036_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel721036_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.CustomListModel721036_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomListModel721036___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomListModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.CustomListModel721036___setItemData_roles_newList(ptr.Pointer())
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
		}(C.CustomListModel721036___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomListModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel721036___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.CustomListModel721036___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel721036___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.CustomListModel721036___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel721036___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.CustomListModel721036___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomListModel721036___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomListModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.CustomListModel721036___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.CustomListModel721036___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *CustomListModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.CustomListModel721036___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomListModel721036___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomListModel) __roleNames_newList() unsafe.Pointer {
	return C.CustomListModel721036___roleNames_newList(ptr.Pointer())
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
		}(C.CustomListModel721036___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomListModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomListModel721036___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomListModel) __itemData_newList() unsafe.Pointer {
	return C.CustomListModel721036___itemData_newList(ptr.Pointer())
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
		}(C.CustomListModel721036___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *CustomListModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel721036___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.CustomListModel721036___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel721036___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __match_newList() unsafe.Pointer {
	return C.CustomListModel721036___match_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel721036___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *CustomListModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.CustomListModel721036___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel721036_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomListModel721036_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel721036_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *CustomListModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *CustomListModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.CustomListModel721036_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomListModel721036___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomListModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CustomListModel721036___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomListModel721036___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomListModel) __findChildren_newList2() unsafe.Pointer {
	return C.CustomListModel721036___findChildren_newList2(ptr.Pointer())
}

func (ptr *CustomListModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomListModel721036___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomListModel) __findChildren_newList3() unsafe.Pointer {
	return C.CustomListModel721036___findChildren_newList3(ptr.Pointer())
}

func (ptr *CustomListModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomListModel721036___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomListModel) __findChildren_newList() unsafe.Pointer {
	return C.CustomListModel721036___findChildren_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomListModel721036___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomListModel) __children_newList() unsafe.Pointer {
	return C.CustomListModel721036___children_newList(ptr.Pointer())
}

func (ptr *CustomListModel) __add_obj_atList(i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomListModel721036___add_obj_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *CustomListModel) __add_obj_setList(i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036___add_obj_setList(ptr.Pointer(), std_core.PointerFromQVariant(i))
	}
}

func (ptr *CustomListModel) __add_obj_newList() unsafe.Pointer {
	return C.CustomListModel721036___add_obj_newList(ptr.Pointer())
}

func NewCustomListModel(parent std_core.QObject_ITF) *CustomListModel {
	tmpValue := NewCustomListModelFromPointer(C.CustomListModel721036_NewCustomListModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCustomListModel721036_DestroyCustomListModel
func callbackCustomListModel721036_DestroyCustomListModel(ptr unsafe.Pointer) {
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
		C.CustomListModel721036_DestroyCustomListModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CustomListModel) DestroyCustomListModelDefault() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_DestroyCustomListModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomListModel721036_DropMimeData
func callbackCustomListModel721036_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackCustomListModel721036_Index
func callbackCustomListModel721036_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewCustomListModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *CustomListModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel721036_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel721036_Sibling
func callbackCustomListModel721036_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewCustomListModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *CustomListModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel721036_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel721036_Flags
func callbackCustomListModel721036_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewCustomListModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomListModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.CustomListModel721036_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackCustomListModel721036_InsertColumns
func callbackCustomListModel721036_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackCustomListModel721036_InsertRows
func callbackCustomListModel721036_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackCustomListModel721036_MoveColumns
func callbackCustomListModel721036_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomListModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackCustomListModel721036_MoveRows
func callbackCustomListModel721036_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *CustomListModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackCustomListModel721036_RemoveColumns
func callbackCustomListModel721036_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackCustomListModel721036_RemoveRows
func callbackCustomListModel721036_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackCustomListModel721036_SetData
func callbackCustomListModel721036_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomListModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackCustomListModel721036_SetHeaderData
func callbackCustomListModel721036_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *CustomListModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackCustomListModel721036_SetItemData
func callbackCustomListModel721036_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
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
		return C.CustomListModel721036_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewCustomListModelFromPointer(NewCustomListModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

//export callbackCustomListModel721036_Submit
func callbackCustomListModel721036_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *CustomListModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackCustomListModel721036_ColumnsAboutToBeInserted
func callbackCustomListModel721036_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel721036_ColumnsAboutToBeMoved
func callbackCustomListModel721036_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackCustomListModel721036_ColumnsAboutToBeRemoved
func callbackCustomListModel721036_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel721036_ColumnsInserted
func callbackCustomListModel721036_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel721036_ColumnsMoved
func callbackCustomListModel721036_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackCustomListModel721036_ColumnsRemoved
func callbackCustomListModel721036_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel721036_DataChanged
func callbackCustomListModel721036_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
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

//export callbackCustomListModel721036_FetchMore
func callbackCustomListModel721036_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewCustomListModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *CustomListModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackCustomListModel721036_HeaderDataChanged
func callbackCustomListModel721036_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel721036_LayoutAboutToBeChanged
func callbackCustomListModel721036_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
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

//export callbackCustomListModel721036_LayoutChanged
func callbackCustomListModel721036_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
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

//export callbackCustomListModel721036_ModelAboutToBeReset
func callbackCustomListModel721036_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackCustomListModel721036_ModelReset
func callbackCustomListModel721036_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackCustomListModel721036_ResetInternalData
func callbackCustomListModel721036_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewCustomListModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *CustomListModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackCustomListModel721036_Revert
func callbackCustomListModel721036_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewCustomListModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *CustomListModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_RevertDefault(ptr.Pointer())
	}
}

//export callbackCustomListModel721036_RowsAboutToBeInserted
func callbackCustomListModel721036_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackCustomListModel721036_RowsAboutToBeMoved
func callbackCustomListModel721036_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackCustomListModel721036_RowsAboutToBeRemoved
func callbackCustomListModel721036_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel721036_RowsInserted
func callbackCustomListModel721036_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel721036_RowsMoved
func callbackCustomListModel721036_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackCustomListModel721036_RowsRemoved
func callbackCustomListModel721036_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackCustomListModel721036_Sort
func callbackCustomListModel721036_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewCustomListModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *CustomListModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackCustomListModel721036_RoleNames
func callbackCustomListModel721036_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
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
		}(C.CustomListModel721036_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackCustomListModel721036_ItemData
func callbackCustomListModel721036_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
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
		}(C.CustomListModel721036_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackCustomListModel721036_MimeData
func callbackCustomListModel721036_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
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
		tmpValue := std_core.NewQMimeDataFromPointer(C.CustomListModel721036_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
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

//export callbackCustomListModel721036_Buddy
func callbackCustomListModel721036_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomListModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomListModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel721036_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel721036_Parent
func callbackCustomListModel721036_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewCustomListModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomListModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.CustomListModel721036_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel721036_Match
func callbackCustomListModel721036_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
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
		}(C.CustomListModel721036_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackCustomListModel721036_Span
func callbackCustomListModel721036_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewCustomListModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *CustomListModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.CustomListModel721036_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel721036_MimeTypes
func callbackCustomListModel721036_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewCustomListModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *CustomListModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.CustomListModel721036_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackCustomListModel721036_Data
func callbackCustomListModel721036_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomListModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *CustomListModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomListModel721036_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel721036_HeaderData
func callbackCustomListModel721036_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewCustomListModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *CustomListModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomListModel721036_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomListModel721036_SupportedDragActions
func callbackCustomListModel721036_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewCustomListModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *CustomListModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomListModel721036_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomListModel721036_SupportedDropActions
func callbackCustomListModel721036_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewCustomListModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *CustomListModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.CustomListModel721036_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCustomListModel721036_CanDropMimeData
func callbackCustomListModel721036_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackCustomListModel721036_CanFetchMore
func callbackCustomListModel721036_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackCustomListModel721036_HasChildren
func callbackCustomListModel721036_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *CustomListModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackCustomListModel721036_ColumnCount
func callbackCustomListModel721036_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomListModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomListModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel721036_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomListModel721036_RowCount
func callbackCustomListModel721036_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewCustomListModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *CustomListModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomListModel721036_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackCustomListModel721036_Event
func callbackCustomListModel721036_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CustomListModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackCustomListModel721036_EventFilter
func callbackCustomListModel721036_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomListModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CustomListModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomListModel721036_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackCustomListModel721036_ChildEvent
func callbackCustomListModel721036_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCustomListModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CustomListModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCustomListModel721036_ConnectNotify
func callbackCustomListModel721036_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomListModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomListModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomListModel721036_CustomEvent
func callbackCustomListModel721036_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewCustomListModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CustomListModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCustomListModel721036_DeleteLater
func callbackCustomListModel721036_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewCustomListModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CustomListModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomListModel721036_Destroyed
func callbackCustomListModel721036_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackCustomListModel721036_DisconnectNotify
func callbackCustomListModel721036_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomListModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomListModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomListModel721036_ObjectNameChanged
func callbackCustomListModel721036_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackCustomListModel721036_TimerEvent
func callbackCustomListModel721036_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCustomListModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CustomListModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type CtxObject_ITF interface {
	std_core.QObject_ITF
	CtxObject_PTR() *CtxObject
}

func (ptr *CtxObject) CtxObject_PTR() *CtxObject {
	return ptr
}

func (ptr *CtxObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *CtxObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromCtxObject(ptr CtxObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CtxObject_PTR().Pointer()
	}
	return nil
}

func NewCtxObjectFromPointer(ptr unsafe.Pointer) (n *CtxObject) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(CtxObject)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *CtxObject:
			n = deduced

		case *std_core.QObject:
			n = &CtxObject{QObject: *deduced}

		default:
			n = new(CtxObject)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackCtxObject721036_Constructor
func callbackCtxObject721036_Constructor(ptr unsafe.Pointer) {
	this := NewCtxObjectFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClicked(this.clicked)
}

//export callbackCtxObject721036_Clicked
func callbackCtxObject721036_Clicked(ptr unsafe.Pointer, b C.int) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(int))(int(int32(b)))
	}

}

func (ptr *CtxObject) ConnectClicked(f func(b int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.CtxObject721036_ConnectClicked(ptr.Pointer())
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

func (ptr *CtxObject) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *CtxObject) Clicked(b int) {
	if ptr.Pointer() != nil {
		C.CtxObject721036_Clicked(ptr.Pointer(), C.int(int32(b)))
	}
}

//export callbackCtxObject721036_Remote
func callbackCtxObject721036_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "remote"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewCtxObjectFromPointer(ptr).RemoteDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CtxObject) ConnectRemote(f func() string) {
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

func (ptr *CtxObject) DisconnectRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "remote")
	}
}

func (ptr *CtxObject) Remote() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *CtxObject) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackCtxObject721036_SetRemote
func callbackCtxObject721036_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setRemote"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	} else {
		NewCtxObjectFromPointer(ptr).SetRemoteDefault(cGoUnpackString(remote))
	}
}

func (ptr *CtxObject) ConnectSetRemote(f func(remote string)) {
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

func (ptr *CtxObject) DisconnectSetRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRemote")
	}
}

func (ptr *CtxObject) SetRemote(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.CtxObject721036_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *CtxObject) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.CtxObject721036_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackCtxObject721036_RemoteChanged
func callbackCtxObject721036_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *CtxObject) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.CtxObject721036_ConnectRemoteChanged(ptr.Pointer())
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

func (ptr *CtxObject) DisconnectRemoteChanged() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DisconnectRemoteChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "remoteChanged")
	}
}

func (ptr *CtxObject) RemoteChanged(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.CtxObject721036_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackCtxObject721036_Transport
func callbackCtxObject721036_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "transport"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewCtxObjectFromPointer(ptr).TransportDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CtxObject) ConnectTransport(f func() string) {
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

func (ptr *CtxObject) DisconnectTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transport")
	}
}

func (ptr *CtxObject) Transport() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *CtxObject) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackCtxObject721036_SetTransport
func callbackCtxObject721036_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTransport"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	} else {
		NewCtxObjectFromPointer(ptr).SetTransportDefault(cGoUnpackString(transport))
	}
}

func (ptr *CtxObject) ConnectSetTransport(f func(transport string)) {
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

func (ptr *CtxObject) DisconnectSetTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransport")
	}
}

func (ptr *CtxObject) SetTransport(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.CtxObject721036_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *CtxObject) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.CtxObject721036_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackCtxObject721036_TransportChanged
func callbackCtxObject721036_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *CtxObject) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.CtxObject721036_ConnectTransportChanged(ptr.Pointer())
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

func (ptr *CtxObject) DisconnectTransportChanged() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DisconnectTransportChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transportChanged")
	}
}

func (ptr *CtxObject) TransportChanged(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.CtxObject721036_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackCtxObject721036_Endpoint
func callbackCtxObject721036_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "endpoint"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewCtxObjectFromPointer(ptr).EndpointDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CtxObject) ConnectEndpoint(f func() string) {
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

func (ptr *CtxObject) DisconnectEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endpoint")
	}
}

func (ptr *CtxObject) Endpoint() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *CtxObject) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackCtxObject721036_SetEndpoint
func callbackCtxObject721036_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setEndpoint"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	} else {
		NewCtxObjectFromPointer(ptr).SetEndpointDefault(cGoUnpackString(endpoint))
	}
}

func (ptr *CtxObject) ConnectSetEndpoint(f func(endpoint string)) {
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

func (ptr *CtxObject) DisconnectSetEndpoint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEndpoint")
	}
}

func (ptr *CtxObject) SetEndpoint(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.CtxObject721036_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *CtxObject) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.CtxObject721036_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackCtxObject721036_EndpointChanged
func callbackCtxObject721036_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *CtxObject) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.CtxObject721036_ConnectEndpointChanged(ptr.Pointer())
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

func (ptr *CtxObject) DisconnectEndpointChanged() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DisconnectEndpointChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "endpointChanged")
	}
}

func (ptr *CtxObject) EndpointChanged(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.CtxObject721036_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackCtxObject721036_From
func callbackCtxObject721036_From(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "from"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewCtxObjectFromPointer(ptr).FromDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CtxObject) ConnectFrom(f func() string) {
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

func (ptr *CtxObject) DisconnectFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "from")
	}
}

func (ptr *CtxObject) From() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_From(ptr.Pointer()))
	}
	return ""
}

func (ptr *CtxObject) FromDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_FromDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackCtxObject721036_SetFrom
func callbackCtxObject721036_SetFrom(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFrom"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	} else {
		NewCtxObjectFromPointer(ptr).SetFromDefault(cGoUnpackString(from))
	}
}

func (ptr *CtxObject) ConnectSetFrom(f func(from string)) {
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

func (ptr *CtxObject) DisconnectSetFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFrom")
	}
}

func (ptr *CtxObject) SetFrom(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.CtxObject721036_SetFrom(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

func (ptr *CtxObject) SetFromDefault(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.CtxObject721036_SetFromDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackCtxObject721036_FromChanged
func callbackCtxObject721036_FromChanged(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	}

}

func (ptr *CtxObject) ConnectFromChanged(f func(from string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromChanged") {
			C.CtxObject721036_ConnectFromChanged(ptr.Pointer())
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

func (ptr *CtxObject) DisconnectFromChanged() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DisconnectFromChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fromChanged")
	}
}

func (ptr *CtxObject) FromChanged(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.CtxObject721036_FromChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackCtxObject721036_Message
func callbackCtxObject721036_Message(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "message"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewCtxObjectFromPointer(ptr).MessageDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CtxObject) ConnectMessage(f func() string) {
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

func (ptr *CtxObject) DisconnectMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "message")
	}
}

func (ptr *CtxObject) Message() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *CtxObject) MessageDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_MessageDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackCtxObject721036_SetMessage
func callbackCtxObject721036_SetMessage(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setMessage"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	} else {
		NewCtxObjectFromPointer(ptr).SetMessageDefault(cGoUnpackString(message))
	}
}

func (ptr *CtxObject) ConnectSetMessage(f func(message string)) {
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

func (ptr *CtxObject) DisconnectSetMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMessage")
	}
}

func (ptr *CtxObject) SetMessage(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.CtxObject721036_SetMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func (ptr *CtxObject) SetMessageDefault(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.CtxObject721036_SetMessageDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackCtxObject721036_MessageChanged
func callbackCtxObject721036_MessageChanged(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "messageChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	}

}

func (ptr *CtxObject) ConnectMessageChanged(f func(message string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "messageChanged") {
			C.CtxObject721036_ConnectMessageChanged(ptr.Pointer())
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

func (ptr *CtxObject) DisconnectMessageChanged() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "messageChanged")
	}
}

func (ptr *CtxObject) MessageChanged(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.CtxObject721036_MessageChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackCtxObject721036_RawData
func callbackCtxObject721036_RawData(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "rawData"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewCtxObjectFromPointer(ptr).RawDataDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CtxObject) ConnectRawData(f func() string) {
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

func (ptr *CtxObject) DisconnectRawData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rawData")
	}
}

func (ptr *CtxObject) RawData() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_RawData(ptr.Pointer()))
	}
	return ""
}

func (ptr *CtxObject) RawDataDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_RawDataDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackCtxObject721036_SetRawData
func callbackCtxObject721036_SetRawData(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setRawData"); signal != nil {
		signal.(func(string))(cGoUnpackString(rawData))
	} else {
		NewCtxObjectFromPointer(ptr).SetRawDataDefault(cGoUnpackString(rawData))
	}
}

func (ptr *CtxObject) ConnectSetRawData(f func(rawData string)) {
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

func (ptr *CtxObject) DisconnectSetRawData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRawData")
	}
}

func (ptr *CtxObject) SetRawData(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.CtxObject721036_SetRawData(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

func (ptr *CtxObject) SetRawDataDefault(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.CtxObject721036_SetRawDataDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackCtxObject721036_RawDataChanged
func callbackCtxObject721036_RawDataChanged(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "rawDataChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(rawData))
	}

}

func (ptr *CtxObject) ConnectRawDataChanged(f func(rawData string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rawDataChanged") {
			C.CtxObject721036_ConnectRawDataChanged(ptr.Pointer())
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

func (ptr *CtxObject) DisconnectRawDataChanged() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DisconnectRawDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rawDataChanged")
	}
}

func (ptr *CtxObject) RawDataChanged(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.CtxObject721036_RawDataChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackCtxObject721036_Hash
func callbackCtxObject721036_Hash(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "hash"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewCtxObjectFromPointer(ptr).HashDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CtxObject) ConnectHash(f func() string) {
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

func (ptr *CtxObject) DisconnectHash() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hash")
	}
}

func (ptr *CtxObject) Hash() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_Hash(ptr.Pointer()))
	}
	return ""
}

func (ptr *CtxObject) HashDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject721036_HashDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackCtxObject721036_SetHash
func callbackCtxObject721036_SetHash(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setHash"); signal != nil {
		signal.(func(string))(cGoUnpackString(hash))
	} else {
		NewCtxObjectFromPointer(ptr).SetHashDefault(cGoUnpackString(hash))
	}
}

func (ptr *CtxObject) ConnectSetHash(f func(hash string)) {
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

func (ptr *CtxObject) DisconnectSetHash() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHash")
	}
}

func (ptr *CtxObject) SetHash(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.CtxObject721036_SetHash(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

func (ptr *CtxObject) SetHashDefault(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.CtxObject721036_SetHashDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

//export callbackCtxObject721036_HashChanged
func callbackCtxObject721036_HashChanged(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hashChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(hash))
	}

}

func (ptr *CtxObject) ConnectHashChanged(f func(hash string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hashChanged") {
			C.CtxObject721036_ConnectHashChanged(ptr.Pointer())
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

func (ptr *CtxObject) DisconnectHashChanged() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DisconnectHashChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hashChanged")
	}
}

func (ptr *CtxObject) HashChanged(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.CtxObject721036_HashChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

func CtxObject_QRegisterMetaType() int {
	return int(int32(C.CtxObject721036_CtxObject721036_QRegisterMetaType()))
}

func (ptr *CtxObject) QRegisterMetaType() int {
	return int(int32(C.CtxObject721036_CtxObject721036_QRegisterMetaType()))
}

func CtxObject_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CtxObject721036_CtxObject721036_QRegisterMetaType2(typeNameC)))
}

func (ptr *CtxObject) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CtxObject721036_CtxObject721036_QRegisterMetaType2(typeNameC)))
}

func CtxObject_QmlRegisterType() int {
	return int(int32(C.CtxObject721036_CtxObject721036_QmlRegisterType()))
}

func (ptr *CtxObject) QmlRegisterType() int {
	return int(int32(C.CtxObject721036_CtxObject721036_QmlRegisterType()))
}

func CtxObject_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.CtxObject721036_CtxObject721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CtxObject) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.CtxObject721036_CtxObject721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CtxObject) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CtxObject721036___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CtxObject) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CtxObject721036___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CtxObject) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject721036___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __findChildren_newList2() unsafe.Pointer {
	return C.CtxObject721036___findChildren_newList2(ptr.Pointer())
}

func (ptr *CtxObject) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject721036___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __findChildren_newList3() unsafe.Pointer {
	return C.CtxObject721036___findChildren_newList3(ptr.Pointer())
}

func (ptr *CtxObject) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject721036___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __findChildren_newList() unsafe.Pointer {
	return C.CtxObject721036___findChildren_newList(ptr.Pointer())
}

func (ptr *CtxObject) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject721036___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __children_newList() unsafe.Pointer {
	return C.CtxObject721036___children_newList(ptr.Pointer())
}

func NewCtxObject(parent std_core.QObject_ITF) *CtxObject {
	tmpValue := NewCtxObjectFromPointer(C.CtxObject721036_NewCtxObject(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCtxObject721036_DestroyCtxObject
func callbackCtxObject721036_DestroyCtxObject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~CtxObject"); signal != nil {
		signal.(func())()
	} else {
		NewCtxObjectFromPointer(ptr).DestroyCtxObjectDefault()
	}
}

func (ptr *CtxObject) ConnectDestroyCtxObject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~CtxObject"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~CtxObject", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~CtxObject", f)
		}
	}
}

func (ptr *CtxObject) DisconnectDestroyCtxObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~CtxObject")
	}
}

func (ptr *CtxObject) DestroyCtxObject() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DestroyCtxObject(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CtxObject) DestroyCtxObjectDefault() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DestroyCtxObjectDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCtxObject721036_Event
func callbackCtxObject721036_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCtxObjectFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CtxObject) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CtxObject721036_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackCtxObject721036_EventFilter
func callbackCtxObject721036_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCtxObjectFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CtxObject) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CtxObject721036_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackCtxObject721036_ChildEvent
func callbackCtxObject721036_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCtxObjectFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CtxObject) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCtxObject721036_ConnectNotify
func callbackCtxObject721036_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCtxObjectFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CtxObject) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCtxObject721036_CustomEvent
func callbackCtxObject721036_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewCtxObjectFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CtxObject) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCtxObject721036_DeleteLater
func callbackCtxObject721036_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewCtxObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CtxObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCtxObject721036_Destroyed
func callbackCtxObject721036_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackCtxObject721036_DisconnectNotify
func callbackCtxObject721036_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCtxObjectFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CtxObject) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCtxObject721036_ObjectNameChanged
func callbackCtxObject721036_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackCtxObject721036_TimerEvent
func callbackCtxObject721036_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCtxObjectFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CtxObject) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject721036_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
