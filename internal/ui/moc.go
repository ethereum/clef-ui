package ui

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
	std_widgets "github.com/therecipe/qt/widgets"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

type CustomLabel_ITF interface {
	std_widgets.QLabel_ITF
	CustomLabel_PTR() *CustomLabel
}

func (ptr *CustomLabel) CustomLabel_PTR() *CustomLabel {
	return ptr
}

func (ptr *CustomLabel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLabel_PTR().Pointer()
	}
	return nil
}

func (ptr *CustomLabel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLabel_PTR().SetPointer(p)
	}
}

func PointerFromCustomLabel(ptr CustomLabel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CustomLabel_PTR().Pointer()
	}
	return nil
}

func NewCustomLabelFromPointer(ptr unsafe.Pointer) (n *CustomLabel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(CustomLabel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *CustomLabel:
			n = deduced

		case *std_widgets.QLabel:
			n = &CustomLabel{QLabel: *deduced}

		default:
			n = new(CustomLabel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackCustomLabel721036_Constructor
func callbackCustomLabel721036_Constructor(ptr unsafe.Pointer) {
	this := NewCustomLabelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectUpdateText(this.QLabel.SetText)
}

//export callbackCustomLabel721036_UpdateText
func callbackCustomLabel721036_UpdateText(ptr unsafe.Pointer, v0 C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "updateText"); signal != nil {
		signal.(func(string))(cGoUnpackString(v0))
	}

}

func (ptr *CustomLabel) ConnectUpdateText(f func(v0 string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateText") {
			C.CustomLabel721036_ConnectUpdateText(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateText"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateText", func(v0 string) {
				signal.(func(string))(v0)
				f(v0)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateText", f)
		}
	}
}

func (ptr *CustomLabel) DisconnectUpdateText() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_DisconnectUpdateText(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateText")
	}
}

func (ptr *CustomLabel) UpdateText(v0 string) {
	if ptr.Pointer() != nil {
		var v0C *C.char
		if v0 != "" {
			v0C = C.CString(v0)
			defer C.free(unsafe.Pointer(v0C))
		}
		C.CustomLabel721036_UpdateText(ptr.Pointer(), C.struct_Moc_PackedString{data: v0C, len: C.longlong(len(v0))})
	}
}

func CustomLabel_QRegisterMetaType() int {
	return int(int32(C.CustomLabel721036_CustomLabel721036_QRegisterMetaType()))
}

func (ptr *CustomLabel) QRegisterMetaType() int {
	return int(int32(C.CustomLabel721036_CustomLabel721036_QRegisterMetaType()))
}

func CustomLabel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomLabel721036_CustomLabel721036_QRegisterMetaType2(typeNameC)))
}

func (ptr *CustomLabel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CustomLabel721036_CustomLabel721036_QRegisterMetaType2(typeNameC)))
}

func CustomLabel_QmlRegisterType() int {
	return int(int32(C.CustomLabel721036_CustomLabel721036_QmlRegisterType()))
}

func (ptr *CustomLabel) QmlRegisterType() int {
	return int(int32(C.CustomLabel721036_CustomLabel721036_QmlRegisterType()))
}

func CustomLabel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.CustomLabel721036_CustomLabel721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomLabel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.CustomLabel721036_CustomLabel721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CustomLabel) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.CustomLabel721036___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomLabel) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *CustomLabel) __addActions_actions_newList() unsafe.Pointer {
	return C.CustomLabel721036___addActions_actions_newList(ptr.Pointer())
}

func (ptr *CustomLabel) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.CustomLabel721036___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomLabel) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *CustomLabel) __insertActions_actions_newList() unsafe.Pointer {
	return C.CustomLabel721036___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *CustomLabel) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.CustomLabel721036___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomLabel) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *CustomLabel) __actions_newList() unsafe.Pointer {
	return C.CustomLabel721036___actions_newList(ptr.Pointer())
}

func (ptr *CustomLabel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CustomLabel721036___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CustomLabel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CustomLabel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CustomLabel721036___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CustomLabel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomLabel721036___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomLabel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomLabel) __findChildren_newList2() unsafe.Pointer {
	return C.CustomLabel721036___findChildren_newList2(ptr.Pointer())
}

func (ptr *CustomLabel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomLabel721036___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomLabel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomLabel) __findChildren_newList3() unsafe.Pointer {
	return C.CustomLabel721036___findChildren_newList3(ptr.Pointer())
}

func (ptr *CustomLabel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomLabel721036___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomLabel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomLabel) __findChildren_newList() unsafe.Pointer {
	return C.CustomLabel721036___findChildren_newList(ptr.Pointer())
}

func (ptr *CustomLabel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CustomLabel721036___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CustomLabel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CustomLabel) __children_newList() unsafe.Pointer {
	return C.CustomLabel721036___children_newList(ptr.Pointer())
}

func NewCustomLabel(parent std_widgets.QWidget_ITF, fo std_core.Qt__WindowType) *CustomLabel {
	tmpValue := NewCustomLabelFromPointer(C.CustomLabel721036_NewCustomLabel(std_widgets.PointerFromQWidget(parent), C.longlong(fo)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewCustomLabel2(text string, parent std_widgets.QWidget_ITF, fo std_core.Qt__WindowType) *CustomLabel {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewCustomLabelFromPointer(C.CustomLabel721036_NewCustomLabel2(C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))}, std_widgets.PointerFromQWidget(parent), C.longlong(fo)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCustomLabel721036_DestroyCustomLabel
func callbackCustomLabel721036_DestroyCustomLabel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~CustomLabel"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).DestroyCustomLabelDefault()
	}
}

func (ptr *CustomLabel) ConnectDestroyCustomLabel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~CustomLabel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~CustomLabel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~CustomLabel", f)
		}
	}
}

func (ptr *CustomLabel) DisconnectDestroyCustomLabel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~CustomLabel")
	}
}

func (ptr *CustomLabel) DestroyCustomLabel() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_DestroyCustomLabel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CustomLabel) DestroyCustomLabelDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_DestroyCustomLabelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomLabel721036_Event
func callbackCustomLabel721036_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomLabelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CustomLabel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomLabel721036_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackCustomLabel721036_FocusNextPrevChild
func callbackCustomLabel721036_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomLabelFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *CustomLabel) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.CustomLabel721036_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackCustomLabel721036_ChangeEvent
func callbackCustomLabel721036_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(ev))
	} else {
		NewCustomLabelFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(ev))
	}
}

func (ptr *CustomLabel) ChangeEventDefault(ev std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))
	}
}

//export callbackCustomLabel721036_Clear
func callbackCustomLabel721036_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *CustomLabel) ClearDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ClearDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_ContextMenuEvent
func callbackCustomLabel721036_ContextMenuEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*std_gui.QContextMenuEvent))(std_gui.NewQContextMenuEventFromPointer(ev))
	} else {
		NewCustomLabelFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(ev))
	}
}

func (ptr *CustomLabel) ContextMenuEventDefault(ev std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(ev))
	}
}

//export callbackCustomLabel721036_FocusInEvent
func callbackCustomLabel721036_FocusInEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(ev))
	} else {
		NewCustomLabelFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *CustomLabel) FocusInEventDefault(ev std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackCustomLabel721036_FocusOutEvent
func callbackCustomLabel721036_FocusOutEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(ev))
	} else {
		NewCustomLabelFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *CustomLabel) FocusOutEventDefault(ev std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackCustomLabel721036_KeyPressEvent
func callbackCustomLabel721036_KeyPressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(ev))
	} else {
		NewCustomLabelFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *CustomLabel) KeyPressEventDefault(ev std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackCustomLabel721036_LinkActivated
func callbackCustomLabel721036_LinkActivated(ptr unsafe.Pointer, link C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "linkActivated"); signal != nil {
		signal.(func(string))(cGoUnpackString(link))
	}

}

//export callbackCustomLabel721036_LinkHovered
func callbackCustomLabel721036_LinkHovered(ptr unsafe.Pointer, link C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "linkHovered"); signal != nil {
		signal.(func(string))(cGoUnpackString(link))
	}

}

//export callbackCustomLabel721036_MouseMoveEvent
func callbackCustomLabel721036_MouseMoveEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewCustomLabelFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *CustomLabel) MouseMoveEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackCustomLabel721036_MousePressEvent
func callbackCustomLabel721036_MousePressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewCustomLabelFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *CustomLabel) MousePressEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackCustomLabel721036_MouseReleaseEvent
func callbackCustomLabel721036_MouseReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(ev))
	} else {
		NewCustomLabelFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *CustomLabel) MouseReleaseEventDefault(ev std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackCustomLabel721036_PaintEvent
func callbackCustomLabel721036_PaintEvent(ptr unsafe.Pointer, vqp unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*std_gui.QPaintEvent))(std_gui.NewQPaintEventFromPointer(vqp))
	} else {
		NewCustomLabelFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(vqp))
	}
}

func (ptr *CustomLabel) PaintEventDefault(vqp std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(vqp))
	}
}

//export callbackCustomLabel721036_SetMovie
func callbackCustomLabel721036_SetMovie(ptr unsafe.Pointer, movie unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setMovie"); signal != nil {
		signal.(func(*std_gui.QMovie))(std_gui.NewQMovieFromPointer(movie))
	} else {
		NewCustomLabelFromPointer(ptr).SetMovieDefault(std_gui.NewQMovieFromPointer(movie))
	}
}

func (ptr *CustomLabel) SetMovieDefault(movie std_gui.QMovie_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetMovieDefault(ptr.Pointer(), std_gui.PointerFromQMovie(movie))
	}
}

//export callbackCustomLabel721036_SetNum2
func callbackCustomLabel721036_SetNum2(ptr unsafe.Pointer, num C.double) {
	if signal := qt.GetSignal(ptr, "setNum2"); signal != nil {
		signal.(func(float64))(float64(num))
	} else {
		NewCustomLabelFromPointer(ptr).SetNum2Default(float64(num))
	}
}

func (ptr *CustomLabel) SetNum2Default(num float64) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetNum2Default(ptr.Pointer(), C.double(num))
	}
}

//export callbackCustomLabel721036_SetNum
func callbackCustomLabel721036_SetNum(ptr unsafe.Pointer, num C.int) {
	if signal := qt.GetSignal(ptr, "setNum"); signal != nil {
		signal.(func(int))(int(int32(num)))
	} else {
		NewCustomLabelFromPointer(ptr).SetNumDefault(int(int32(num)))
	}
}

func (ptr *CustomLabel) SetNumDefault(num int) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetNumDefault(ptr.Pointer(), C.int(int32(num)))
	}
}

//export callbackCustomLabel721036_SetPicture
func callbackCustomLabel721036_SetPicture(ptr unsafe.Pointer, picture unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setPicture"); signal != nil {
		signal.(func(*std_gui.QPicture))(std_gui.NewQPictureFromPointer(picture))
	} else {
		NewCustomLabelFromPointer(ptr).SetPictureDefault(std_gui.NewQPictureFromPointer(picture))
	}
}

func (ptr *CustomLabel) SetPictureDefault(picture std_gui.QPicture_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetPictureDefault(ptr.Pointer(), std_gui.PointerFromQPicture(picture))
	}
}

//export callbackCustomLabel721036_SetPixmap
func callbackCustomLabel721036_SetPixmap(ptr unsafe.Pointer, vqp unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setPixmap"); signal != nil {
		signal.(func(*std_gui.QPixmap))(std_gui.NewQPixmapFromPointer(vqp))
	} else {
		NewCustomLabelFromPointer(ptr).SetPixmapDefault(std_gui.NewQPixmapFromPointer(vqp))
	}
}

func (ptr *CustomLabel) SetPixmapDefault(vqp std_gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetPixmapDefault(ptr.Pointer(), std_gui.PointerFromQPixmap(vqp))
	}
}

//export callbackCustomLabel721036_SetText
func callbackCustomLabel721036_SetText(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setText"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewCustomLabelFromPointer(ptr).SetTextDefault(cGoUnpackString(vqs))
	}
}

func (ptr *CustomLabel) SetTextDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.CustomLabel721036_SetTextDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackCustomLabel721036_MinimumSizeHint
func callbackCustomLabel721036_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewCustomLabelFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *CustomLabel) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.CustomLabel721036_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackCustomLabel721036_SizeHint
func callbackCustomLabel721036_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewCustomLabelFromPointer(ptr).SizeHintDefault())
}

func (ptr *CustomLabel) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.CustomLabel721036_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackCustomLabel721036_HeightForWidth
func callbackCustomLabel721036_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewCustomLabelFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *CustomLabel) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomLabel721036_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackCustomLabel721036_Close
func callbackCustomLabel721036_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomLabelFromPointer(ptr).CloseDefault())))
}

func (ptr *CustomLabel) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.CustomLabel721036_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackCustomLabel721036_ActionEvent
func callbackCustomLabel721036_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*std_gui.QActionEvent))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *CustomLabel) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackCustomLabel721036_CloseEvent
func callbackCustomLabel721036_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*std_gui.QCloseEvent))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *CustomLabel) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackCustomLabel721036_CustomContextMenuRequested
func callbackCustomLabel721036_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*std_core.QPoint))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackCustomLabel721036_DragEnterEvent
func callbackCustomLabel721036_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*std_gui.QDragEnterEvent))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *CustomLabel) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackCustomLabel721036_DragLeaveEvent
func callbackCustomLabel721036_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*std_gui.QDragLeaveEvent))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *CustomLabel) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackCustomLabel721036_DragMoveEvent
func callbackCustomLabel721036_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*std_gui.QDragMoveEvent))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *CustomLabel) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackCustomLabel721036_DropEvent
func callbackCustomLabel721036_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*std_gui.QDropEvent))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *CustomLabel) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackCustomLabel721036_EnterEvent
func callbackCustomLabel721036_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CustomLabel) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCustomLabel721036_Hide
func callbackCustomLabel721036_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).HideDefault()
	}
}

func (ptr *CustomLabel) HideDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_HideDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_HideEvent
func callbackCustomLabel721036_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*std_gui.QHideEvent))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *CustomLabel) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackCustomLabel721036_InputMethodEvent
func callbackCustomLabel721036_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*std_gui.QInputMethodEvent))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *CustomLabel) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackCustomLabel721036_KeyReleaseEvent
func callbackCustomLabel721036_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *CustomLabel) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackCustomLabel721036_LeaveEvent
func callbackCustomLabel721036_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CustomLabel) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCustomLabel721036_Lower
func callbackCustomLabel721036_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).LowerDefault()
	}
}

func (ptr *CustomLabel) LowerDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_LowerDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_MouseDoubleClickEvent
func callbackCustomLabel721036_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *CustomLabel) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackCustomLabel721036_MoveEvent
func callbackCustomLabel721036_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*std_gui.QMoveEvent))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *CustomLabel) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackCustomLabel721036_Raise
func callbackCustomLabel721036_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *CustomLabel) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_RaiseDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_Repaint
func callbackCustomLabel721036_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *CustomLabel) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_RepaintDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_ResizeEvent
func callbackCustomLabel721036_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*std_gui.QResizeEvent))(std_gui.NewQResizeEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *CustomLabel) ResizeEventDefault(event std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(event))
	}
}

//export callbackCustomLabel721036_SetDisabled
func callbackCustomLabel721036_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewCustomLabelFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *CustomLabel) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackCustomLabel721036_SetEnabled
func callbackCustomLabel721036_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewCustomLabelFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *CustomLabel) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackCustomLabel721036_SetFocus2
func callbackCustomLabel721036_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *CustomLabel) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_SetHidden
func callbackCustomLabel721036_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewCustomLabelFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *CustomLabel) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackCustomLabel721036_SetStyleSheet
func callbackCustomLabel721036_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewCustomLabelFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *CustomLabel) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.CustomLabel721036_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackCustomLabel721036_SetVisible
func callbackCustomLabel721036_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewCustomLabelFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *CustomLabel) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackCustomLabel721036_SetWindowModified
func callbackCustomLabel721036_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewCustomLabelFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *CustomLabel) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackCustomLabel721036_SetWindowTitle
func callbackCustomLabel721036_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewCustomLabelFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *CustomLabel) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.CustomLabel721036_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackCustomLabel721036_Show
func callbackCustomLabel721036_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).ShowDefault()
	}
}

func (ptr *CustomLabel) ShowDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ShowDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_ShowEvent
func callbackCustomLabel721036_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*std_gui.QShowEvent))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *CustomLabel) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackCustomLabel721036_ShowFullScreen
func callbackCustomLabel721036_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *CustomLabel) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_ShowMaximized
func callbackCustomLabel721036_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *CustomLabel) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_ShowMinimized
func callbackCustomLabel721036_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *CustomLabel) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_ShowNormal
func callbackCustomLabel721036_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *CustomLabel) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_TabletEvent
func callbackCustomLabel721036_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*std_gui.QTabletEvent))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *CustomLabel) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackCustomLabel721036_Update
func callbackCustomLabel721036_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *CustomLabel) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_UpdateDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_UpdateMicroFocus
func callbackCustomLabel721036_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *CustomLabel) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackCustomLabel721036_WheelEvent
func callbackCustomLabel721036_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*std_gui.QWheelEvent))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *CustomLabel) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackCustomLabel721036_WindowIconChanged
func callbackCustomLabel721036_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*std_gui.QIcon))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackCustomLabel721036_WindowTitleChanged
func callbackCustomLabel721036_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackCustomLabel721036_PaintEngine
func callbackCustomLabel721036_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine(signal.(func() *std_gui.QPaintEngine)())
	}

	return std_gui.PointerFromQPaintEngine(NewCustomLabelFromPointer(ptr).PaintEngineDefault())
}

func (ptr *CustomLabel) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.CustomLabel721036_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackCustomLabel721036_InputMethodQuery
func callbackCustomLabel721036_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewCustomLabelFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *CustomLabel) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.CustomLabel721036_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackCustomLabel721036_HasHeightForWidth
func callbackCustomLabel721036_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomLabelFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *CustomLabel) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.CustomLabel721036_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackCustomLabel721036_Metric
func callbackCustomLabel721036_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(std_gui.QPaintDevice__PaintDeviceMetric) int)(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewCustomLabelFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *CustomLabel) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.CustomLabel721036_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackCustomLabel721036_InitPainter
func callbackCustomLabel721036_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		signal.(func(*std_gui.QPainter))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewCustomLabelFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *CustomLabel) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackCustomLabel721036_EventFilter
func callbackCustomLabel721036_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCustomLabelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CustomLabel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.CustomLabel721036_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackCustomLabel721036_ChildEvent
func callbackCustomLabel721036_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CustomLabel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCustomLabel721036_ConnectNotify
func callbackCustomLabel721036_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomLabelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomLabel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomLabel721036_CustomEvent
func callbackCustomLabel721036_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CustomLabel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCustomLabel721036_DeleteLater
func callbackCustomLabel721036_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewCustomLabelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CustomLabel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCustomLabel721036_Destroyed
func callbackCustomLabel721036_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackCustomLabel721036_DisconnectNotify
func callbackCustomLabel721036_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCustomLabelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CustomLabel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCustomLabel721036_ObjectNameChanged
func callbackCustomLabel721036_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackCustomLabel721036_TimerEvent
func callbackCustomLabel721036_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCustomLabelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CustomLabel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CustomLabel721036_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
