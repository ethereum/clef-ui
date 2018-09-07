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

	custom_params_aaf9d9m "github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

//export callbackApproveListingCtx721036_Constructor
func callbackApproveListingCtx721036_Constructor(ptr unsafe.Pointer) {
	this := NewApproveListingCtxFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackApproveListingCtx721036_Remote
func callbackApproveListingCtx721036_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveListingCtx721036_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx721036_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx721036_SetRemote
func callbackApproveListingCtx721036_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
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
		C.ApproveListingCtx721036_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *ApproveListingCtx) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveListingCtx721036_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveListingCtx721036_RemoteChanged
func callbackApproveListingCtx721036_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *ApproveListingCtx) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.ApproveListingCtx721036_ConnectRemoteChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_DisconnectRemoteChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveListingCtx721036_Transport
func callbackApproveListingCtx721036_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveListingCtx721036_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx721036_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx721036_SetTransport
func callbackApproveListingCtx721036_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
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
		C.ApproveListingCtx721036_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *ApproveListingCtx) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveListingCtx721036_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveListingCtx721036_TransportChanged
func callbackApproveListingCtx721036_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *ApproveListingCtx) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.ApproveListingCtx721036_ConnectTransportChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_DisconnectTransportChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveListingCtx721036_Endpoint
func callbackApproveListingCtx721036_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveListingCtx721036_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx721036_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx721036_SetEndpoint
func callbackApproveListingCtx721036_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
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
		C.ApproveListingCtx721036_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *ApproveListingCtx) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveListingCtx721036_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveListingCtx721036_EndpointChanged
func callbackApproveListingCtx721036_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *ApproveListingCtx) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.ApproveListingCtx721036_ConnectEndpointChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_DisconnectEndpointChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveListingCtx721036_From
func callbackApproveListingCtx721036_From(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveListingCtx721036_From(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) FromDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx721036_FromDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx721036_SetFrom
func callbackApproveListingCtx721036_SetFrom(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
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
		C.ApproveListingCtx721036_SetFrom(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

func (ptr *ApproveListingCtx) SetFromDefault(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveListingCtx721036_SetFromDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveListingCtx721036_FromChanged
func callbackApproveListingCtx721036_FromChanged(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	}

}

func (ptr *ApproveListingCtx) ConnectFromChanged(f func(from string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromChanged") {
			C.ApproveListingCtx721036_ConnectFromChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_DisconnectFromChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_FromChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveListingCtx721036_Message
func callbackApproveListingCtx721036_Message(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveListingCtx721036_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) MessageDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx721036_MessageDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx721036_SetMessage
func callbackApproveListingCtx721036_SetMessage(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
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
		C.ApproveListingCtx721036_SetMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func (ptr *ApproveListingCtx) SetMessageDefault(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.ApproveListingCtx721036_SetMessageDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackApproveListingCtx721036_MessageChanged
func callbackApproveListingCtx721036_MessageChanged(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "messageChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	}

}

func (ptr *ApproveListingCtx) ConnectMessageChanged(f func(message string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "messageChanged") {
			C.ApproveListingCtx721036_ConnectMessageChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_DisconnectMessageChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_MessageChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackApproveListingCtx721036_RawData
func callbackApproveListingCtx721036_RawData(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveListingCtx721036_RawData(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) RawDataDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx721036_RawDataDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx721036_SetRawData
func callbackApproveListingCtx721036_SetRawData(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
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
		C.ApproveListingCtx721036_SetRawData(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

func (ptr *ApproveListingCtx) SetRawDataDefault(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.ApproveListingCtx721036_SetRawDataDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackApproveListingCtx721036_RawDataChanged
func callbackApproveListingCtx721036_RawDataChanged(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "rawDataChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(rawData))
	}

}

func (ptr *ApproveListingCtx) ConnectRawDataChanged(f func(rawData string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rawDataChanged") {
			C.ApproveListingCtx721036_ConnectRawDataChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_DisconnectRawDataChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_RawDataChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackApproveListingCtx721036_Hash
func callbackApproveListingCtx721036_Hash(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveListingCtx721036_Hash(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveListingCtx) HashDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveListingCtx721036_HashDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveListingCtx721036_SetHash
func callbackApproveListingCtx721036_SetHash(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
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
		C.ApproveListingCtx721036_SetHash(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

func (ptr *ApproveListingCtx) SetHashDefault(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.ApproveListingCtx721036_SetHashDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

//export callbackApproveListingCtx721036_HashChanged
func callbackApproveListingCtx721036_HashChanged(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hashChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(hash))
	}

}

func (ptr *ApproveListingCtx) ConnectHashChanged(f func(hash string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hashChanged") {
			C.ApproveListingCtx721036_ConnectHashChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_DisconnectHashChanged(ptr.Pointer())
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
		C.ApproveListingCtx721036_HashChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

func ApproveListingCtx_QRegisterMetaType() int {
	return int(int32(C.ApproveListingCtx721036_ApproveListingCtx721036_QRegisterMetaType()))
}

func (ptr *ApproveListingCtx) QRegisterMetaType() int {
	return int(int32(C.ApproveListingCtx721036_ApproveListingCtx721036_QRegisterMetaType()))
}

func ApproveListingCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveListingCtx721036_ApproveListingCtx721036_QRegisterMetaType2(typeNameC)))
}

func (ptr *ApproveListingCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveListingCtx721036_ApproveListingCtx721036_QRegisterMetaType2(typeNameC)))
}

func ApproveListingCtx_QmlRegisterType() int {
	return int(int32(C.ApproveListingCtx721036_ApproveListingCtx721036_QmlRegisterType()))
}

func (ptr *ApproveListingCtx) QmlRegisterType() int {
	return int(int32(C.ApproveListingCtx721036_ApproveListingCtx721036_QmlRegisterType()))
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
	return int(int32(C.ApproveListingCtx721036_ApproveListingCtx721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.ApproveListingCtx721036_ApproveListingCtx721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveListingCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ApproveListingCtx721036___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ApproveListingCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ApproveListingCtx721036___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ApproveListingCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveListingCtx721036___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveListingCtx) __findChildren_newList2() unsafe.Pointer {
	return C.ApproveListingCtx721036___findChildren_newList2(ptr.Pointer())
}

func (ptr *ApproveListingCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveListingCtx721036___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveListingCtx) __findChildren_newList3() unsafe.Pointer {
	return C.ApproveListingCtx721036___findChildren_newList3(ptr.Pointer())
}

func (ptr *ApproveListingCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveListingCtx721036___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveListingCtx) __findChildren_newList() unsafe.Pointer {
	return C.ApproveListingCtx721036___findChildren_newList(ptr.Pointer())
}

func (ptr *ApproveListingCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveListingCtx721036___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveListingCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveListingCtx) __children_newList() unsafe.Pointer {
	return C.ApproveListingCtx721036___children_newList(ptr.Pointer())
}

func NewApproveListingCtx(parent std_core.QObject_ITF) *ApproveListingCtx {
	tmpValue := NewApproveListingCtxFromPointer(C.ApproveListingCtx721036_NewApproveListingCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApproveListingCtx721036_DestroyApproveListingCtx
func callbackApproveListingCtx721036_DestroyApproveListingCtx(ptr unsafe.Pointer) {
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
		C.ApproveListingCtx721036_DestroyApproveListingCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ApproveListingCtx) DestroyApproveListingCtxDefault() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036_DestroyApproveListingCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveListingCtx721036_Event
func callbackApproveListingCtx721036_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveListingCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ApproveListingCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApproveListingCtx721036_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackApproveListingCtx721036_EventFilter
func callbackApproveListingCtx721036_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveListingCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ApproveListingCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApproveListingCtx721036_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackApproveListingCtx721036_ChildEvent
func callbackApproveListingCtx721036_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewApproveListingCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ApproveListingCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackApproveListingCtx721036_ConnectNotify
func callbackApproveListingCtx721036_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveListingCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveListingCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveListingCtx721036_CustomEvent
func callbackApproveListingCtx721036_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewApproveListingCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ApproveListingCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackApproveListingCtx721036_DeleteLater
func callbackApproveListingCtx721036_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewApproveListingCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ApproveListingCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveListingCtx721036_Destroyed
func callbackApproveListingCtx721036_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackApproveListingCtx721036_DisconnectNotify
func callbackApproveListingCtx721036_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveListingCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveListingCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveListingCtx721036_ObjectNameChanged
func callbackApproveListingCtx721036_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApproveListingCtx721036_TimerEvent
func callbackApproveListingCtx721036_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewApproveListingCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ApproveListingCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveListingCtx721036_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
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

//export callbackApproveNewAccountCtx721036_Constructor
func callbackApproveNewAccountCtx721036_Constructor(ptr unsafe.Pointer) {
	this := NewApproveNewAccountCtxFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClicked(this.clicked)
	this.ConnectPasswordEdited(this.passwordEdited)
	this.ConnectConfirmPasswordEdited(this.confirmPasswordEdited)
	this.init()
}

//export callbackApproveNewAccountCtx721036_Clicked
func callbackApproveNewAccountCtx721036_Clicked(ptr unsafe.Pointer, b C.int) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(int))(int(int32(b)))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectClicked(f func(b int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.ApproveNewAccountCtx721036_ConnectClicked(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *ApproveNewAccountCtx) Clicked(b int) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036_Clicked(ptr.Pointer(), C.int(int32(b)))
	}
}

//export callbackApproveNewAccountCtx721036_PasswordEdited
func callbackApproveNewAccountCtx721036_PasswordEdited(ptr unsafe.Pointer, b C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "passwordEdited"); signal != nil {
		signal.(func(string))(cGoUnpackString(b))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectPasswordEdited(f func(b string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "passwordEdited") {
			C.ApproveNewAccountCtx721036_ConnectPasswordEdited(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_DisconnectPasswordEdited(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_PasswordEdited(ptr.Pointer(), C.struct_Moc_PackedString{data: bC, len: C.longlong(len(b))})
	}
}

//export callbackApproveNewAccountCtx721036_ConfirmPasswordEdited
func callbackApproveNewAccountCtx721036_ConfirmPasswordEdited(ptr unsafe.Pointer, b C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "confirmPasswordEdited"); signal != nil {
		signal.(func(string))(cGoUnpackString(b))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectConfirmPasswordEdited(f func(b string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "confirmPasswordEdited") {
			C.ApproveNewAccountCtx721036_ConnectConfirmPasswordEdited(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_DisconnectConfirmPasswordEdited(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_ConfirmPasswordEdited(ptr.Pointer(), C.struct_Moc_PackedString{data: bC, len: C.longlong(len(b))})
	}
}

//export callbackApproveNewAccountCtx721036_Remote
func callbackApproveNewAccountCtx721036_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveNewAccountCtx721036_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx721036_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx721036_SetRemote
func callbackApproveNewAccountCtx721036_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
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
		C.ApproveNewAccountCtx721036_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *ApproveNewAccountCtx) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveNewAccountCtx721036_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveNewAccountCtx721036_RemoteChanged
func callbackApproveNewAccountCtx721036_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.ApproveNewAccountCtx721036_ConnectRemoteChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_DisconnectRemoteChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveNewAccountCtx721036_Transport
func callbackApproveNewAccountCtx721036_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveNewAccountCtx721036_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx721036_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx721036_SetTransport
func callbackApproveNewAccountCtx721036_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
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
		C.ApproveNewAccountCtx721036_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *ApproveNewAccountCtx) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveNewAccountCtx721036_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveNewAccountCtx721036_TransportChanged
func callbackApproveNewAccountCtx721036_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.ApproveNewAccountCtx721036_ConnectTransportChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_DisconnectTransportChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveNewAccountCtx721036_Endpoint
func callbackApproveNewAccountCtx721036_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveNewAccountCtx721036_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx721036_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx721036_SetEndpoint
func callbackApproveNewAccountCtx721036_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
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
		C.ApproveNewAccountCtx721036_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *ApproveNewAccountCtx) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveNewAccountCtx721036_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveNewAccountCtx721036_EndpointChanged
func callbackApproveNewAccountCtx721036_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.ApproveNewAccountCtx721036_ConnectEndpointChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_DisconnectEndpointChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveNewAccountCtx721036_Password
func callbackApproveNewAccountCtx721036_Password(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveNewAccountCtx721036_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) PasswordDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx721036_PasswordDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx721036_SetPassword
func callbackApproveNewAccountCtx721036_SetPassword(ptr unsafe.Pointer, password C.struct_Moc_PackedString) {
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
		C.ApproveNewAccountCtx721036_SetPassword(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

func (ptr *ApproveNewAccountCtx) SetPasswordDefault(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.ApproveNewAccountCtx721036_SetPasswordDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackApproveNewAccountCtx721036_PasswordChanged
func callbackApproveNewAccountCtx721036_PasswordChanged(ptr unsafe.Pointer, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "passwordChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(password))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectPasswordChanged(f func(password string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "passwordChanged") {
			C.ApproveNewAccountCtx721036_ConnectPasswordChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_DisconnectPasswordChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_PasswordChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackApproveNewAccountCtx721036_ConfirmPassword
func callbackApproveNewAccountCtx721036_ConfirmPassword(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveNewAccountCtx721036_ConfirmPassword(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveNewAccountCtx) ConfirmPasswordDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveNewAccountCtx721036_ConfirmPasswordDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveNewAccountCtx721036_SetConfirmPassword
func callbackApproveNewAccountCtx721036_SetConfirmPassword(ptr unsafe.Pointer, confirmPassword C.struct_Moc_PackedString) {
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
		C.ApproveNewAccountCtx721036_SetConfirmPassword(ptr.Pointer(), C.struct_Moc_PackedString{data: confirmPasswordC, len: C.longlong(len(confirmPassword))})
	}
}

func (ptr *ApproveNewAccountCtx) SetConfirmPasswordDefault(confirmPassword string) {
	if ptr.Pointer() != nil {
		var confirmPasswordC *C.char
		if confirmPassword != "" {
			confirmPasswordC = C.CString(confirmPassword)
			defer C.free(unsafe.Pointer(confirmPasswordC))
		}
		C.ApproveNewAccountCtx721036_SetConfirmPasswordDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: confirmPasswordC, len: C.longlong(len(confirmPassword))})
	}
}

//export callbackApproveNewAccountCtx721036_ConfirmPasswordChanged
func callbackApproveNewAccountCtx721036_ConfirmPasswordChanged(ptr unsafe.Pointer, confirmPassword C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "confirmPasswordChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(confirmPassword))
	}

}

func (ptr *ApproveNewAccountCtx) ConnectConfirmPasswordChanged(f func(confirmPassword string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "confirmPasswordChanged") {
			C.ApproveNewAccountCtx721036_ConnectConfirmPasswordChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_DisconnectConfirmPasswordChanged(ptr.Pointer())
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
		C.ApproveNewAccountCtx721036_ConfirmPasswordChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: confirmPasswordC, len: C.longlong(len(confirmPassword))})
	}
}

func ApproveNewAccountCtx_QRegisterMetaType() int {
	return int(int32(C.ApproveNewAccountCtx721036_ApproveNewAccountCtx721036_QRegisterMetaType()))
}

func (ptr *ApproveNewAccountCtx) QRegisterMetaType() int {
	return int(int32(C.ApproveNewAccountCtx721036_ApproveNewAccountCtx721036_QRegisterMetaType()))
}

func ApproveNewAccountCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveNewAccountCtx721036_ApproveNewAccountCtx721036_QRegisterMetaType2(typeNameC)))
}

func (ptr *ApproveNewAccountCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveNewAccountCtx721036_ApproveNewAccountCtx721036_QRegisterMetaType2(typeNameC)))
}

func ApproveNewAccountCtx_QmlRegisterType() int {
	return int(int32(C.ApproveNewAccountCtx721036_ApproveNewAccountCtx721036_QmlRegisterType()))
}

func (ptr *ApproveNewAccountCtx) QmlRegisterType() int {
	return int(int32(C.ApproveNewAccountCtx721036_ApproveNewAccountCtx721036_QmlRegisterType()))
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
	return int(int32(C.ApproveNewAccountCtx721036_ApproveNewAccountCtx721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.ApproveNewAccountCtx721036_ApproveNewAccountCtx721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveNewAccountCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ApproveNewAccountCtx721036___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ApproveNewAccountCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ApproveNewAccountCtx721036___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ApproveNewAccountCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveNewAccountCtx721036___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveNewAccountCtx) __findChildren_newList2() unsafe.Pointer {
	return C.ApproveNewAccountCtx721036___findChildren_newList2(ptr.Pointer())
}

func (ptr *ApproveNewAccountCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveNewAccountCtx721036___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveNewAccountCtx) __findChildren_newList3() unsafe.Pointer {
	return C.ApproveNewAccountCtx721036___findChildren_newList3(ptr.Pointer())
}

func (ptr *ApproveNewAccountCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveNewAccountCtx721036___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveNewAccountCtx) __findChildren_newList() unsafe.Pointer {
	return C.ApproveNewAccountCtx721036___findChildren_newList(ptr.Pointer())
}

func (ptr *ApproveNewAccountCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveNewAccountCtx721036___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveNewAccountCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveNewAccountCtx) __children_newList() unsafe.Pointer {
	return C.ApproveNewAccountCtx721036___children_newList(ptr.Pointer())
}

func NewApproveNewAccountCtx(parent std_core.QObject_ITF) *ApproveNewAccountCtx {
	tmpValue := NewApproveNewAccountCtxFromPointer(C.ApproveNewAccountCtx721036_NewApproveNewAccountCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApproveNewAccountCtx721036_DestroyApproveNewAccountCtx
func callbackApproveNewAccountCtx721036_DestroyApproveNewAccountCtx(ptr unsafe.Pointer) {
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
		C.ApproveNewAccountCtx721036_DestroyApproveNewAccountCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ApproveNewAccountCtx) DestroyApproveNewAccountCtxDefault() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036_DestroyApproveNewAccountCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveNewAccountCtx721036_Event
func callbackApproveNewAccountCtx721036_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveNewAccountCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ApproveNewAccountCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApproveNewAccountCtx721036_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackApproveNewAccountCtx721036_EventFilter
func callbackApproveNewAccountCtx721036_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveNewAccountCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ApproveNewAccountCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApproveNewAccountCtx721036_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackApproveNewAccountCtx721036_ChildEvent
func callbackApproveNewAccountCtx721036_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ApproveNewAccountCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackApproveNewAccountCtx721036_ConnectNotify
func callbackApproveNewAccountCtx721036_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveNewAccountCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveNewAccountCtx721036_CustomEvent
func callbackApproveNewAccountCtx721036_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ApproveNewAccountCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackApproveNewAccountCtx721036_DeleteLater
func callbackApproveNewAccountCtx721036_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ApproveNewAccountCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveNewAccountCtx721036_Destroyed
func callbackApproveNewAccountCtx721036_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackApproveNewAccountCtx721036_DisconnectNotify
func callbackApproveNewAccountCtx721036_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveNewAccountCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveNewAccountCtx721036_ObjectNameChanged
func callbackApproveNewAccountCtx721036_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApproveNewAccountCtx721036_TimerEvent
func callbackApproveNewAccountCtx721036_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewApproveNewAccountCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ApproveNewAccountCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveNewAccountCtx721036_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
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

//export callbackApproveSignDataCtx721036_Constructor
func callbackApproveSignDataCtx721036_Constructor(ptr unsafe.Pointer) {
	this := NewApproveSignDataCtxFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClicked(this.clicked)
}

//export callbackApproveSignDataCtx721036_Clicked
func callbackApproveSignDataCtx721036_Clicked(ptr unsafe.Pointer, b C.int) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(int))(int(int32(b)))
	}

}

func (ptr *ApproveSignDataCtx) ConnectClicked(f func(b int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.ApproveSignDataCtx721036_ConnectClicked(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *ApproveSignDataCtx) Clicked(b int) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036_Clicked(ptr.Pointer(), C.int(int32(b)))
	}
}

//export callbackApproveSignDataCtx721036_Remote
func callbackApproveSignDataCtx721036_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveSignDataCtx721036_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx721036_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx721036_SetRemote
func callbackApproveSignDataCtx721036_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
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
		C.ApproveSignDataCtx721036_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *ApproveSignDataCtx) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveSignDataCtx721036_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveSignDataCtx721036_RemoteChanged
func callbackApproveSignDataCtx721036_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *ApproveSignDataCtx) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.ApproveSignDataCtx721036_ConnectRemoteChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_DisconnectRemoteChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveSignDataCtx721036_Transport
func callbackApproveSignDataCtx721036_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveSignDataCtx721036_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx721036_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx721036_SetTransport
func callbackApproveSignDataCtx721036_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
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
		C.ApproveSignDataCtx721036_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *ApproveSignDataCtx) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveSignDataCtx721036_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveSignDataCtx721036_TransportChanged
func callbackApproveSignDataCtx721036_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *ApproveSignDataCtx) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.ApproveSignDataCtx721036_ConnectTransportChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_DisconnectTransportChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveSignDataCtx721036_Endpoint
func callbackApproveSignDataCtx721036_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveSignDataCtx721036_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx721036_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx721036_SetEndpoint
func callbackApproveSignDataCtx721036_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
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
		C.ApproveSignDataCtx721036_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *ApproveSignDataCtx) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveSignDataCtx721036_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveSignDataCtx721036_EndpointChanged
func callbackApproveSignDataCtx721036_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *ApproveSignDataCtx) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.ApproveSignDataCtx721036_ConnectEndpointChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_DisconnectEndpointChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveSignDataCtx721036_From
func callbackApproveSignDataCtx721036_From(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveSignDataCtx721036_From(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) FromDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx721036_FromDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx721036_SetFrom
func callbackApproveSignDataCtx721036_SetFrom(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
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
		C.ApproveSignDataCtx721036_SetFrom(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

func (ptr *ApproveSignDataCtx) SetFromDefault(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveSignDataCtx721036_SetFromDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveSignDataCtx721036_FromChanged
func callbackApproveSignDataCtx721036_FromChanged(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	}

}

func (ptr *ApproveSignDataCtx) ConnectFromChanged(f func(from string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromChanged") {
			C.ApproveSignDataCtx721036_ConnectFromChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_DisconnectFromChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_FromChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveSignDataCtx721036_Message
func callbackApproveSignDataCtx721036_Message(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveSignDataCtx721036_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) MessageDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx721036_MessageDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx721036_SetMessage
func callbackApproveSignDataCtx721036_SetMessage(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
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
		C.ApproveSignDataCtx721036_SetMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func (ptr *ApproveSignDataCtx) SetMessageDefault(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.ApproveSignDataCtx721036_SetMessageDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackApproveSignDataCtx721036_MessageChanged
func callbackApproveSignDataCtx721036_MessageChanged(ptr unsafe.Pointer, message C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "messageChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	}

}

func (ptr *ApproveSignDataCtx) ConnectMessageChanged(f func(message string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "messageChanged") {
			C.ApproveSignDataCtx721036_ConnectMessageChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_DisconnectMessageChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_MessageChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackApproveSignDataCtx721036_RawData
func callbackApproveSignDataCtx721036_RawData(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveSignDataCtx721036_RawData(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) RawDataDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx721036_RawDataDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx721036_SetRawData
func callbackApproveSignDataCtx721036_SetRawData(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
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
		C.ApproveSignDataCtx721036_SetRawData(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

func (ptr *ApproveSignDataCtx) SetRawDataDefault(rawData string) {
	if ptr.Pointer() != nil {
		var rawDataC *C.char
		if rawData != "" {
			rawDataC = C.CString(rawData)
			defer C.free(unsafe.Pointer(rawDataC))
		}
		C.ApproveSignDataCtx721036_SetRawDataDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackApproveSignDataCtx721036_RawDataChanged
func callbackApproveSignDataCtx721036_RawDataChanged(ptr unsafe.Pointer, rawData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "rawDataChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(rawData))
	}

}

func (ptr *ApproveSignDataCtx) ConnectRawDataChanged(f func(rawData string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rawDataChanged") {
			C.ApproveSignDataCtx721036_ConnectRawDataChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_DisconnectRawDataChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_RawDataChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: rawDataC, len: C.longlong(len(rawData))})
	}
}

//export callbackApproveSignDataCtx721036_Hash
func callbackApproveSignDataCtx721036_Hash(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveSignDataCtx721036_Hash(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveSignDataCtx) HashDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveSignDataCtx721036_HashDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveSignDataCtx721036_SetHash
func callbackApproveSignDataCtx721036_SetHash(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
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
		C.ApproveSignDataCtx721036_SetHash(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

func (ptr *ApproveSignDataCtx) SetHashDefault(hash string) {
	if ptr.Pointer() != nil {
		var hashC *C.char
		if hash != "" {
			hashC = C.CString(hash)
			defer C.free(unsafe.Pointer(hashC))
		}
		C.ApproveSignDataCtx721036_SetHashDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

//export callbackApproveSignDataCtx721036_HashChanged
func callbackApproveSignDataCtx721036_HashChanged(ptr unsafe.Pointer, hash C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hashChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(hash))
	}

}

func (ptr *ApproveSignDataCtx) ConnectHashChanged(f func(hash string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hashChanged") {
			C.ApproveSignDataCtx721036_ConnectHashChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_DisconnectHashChanged(ptr.Pointer())
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
		C.ApproveSignDataCtx721036_HashChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hashC, len: C.longlong(len(hash))})
	}
}

func ApproveSignDataCtx_QRegisterMetaType() int {
	return int(int32(C.ApproveSignDataCtx721036_ApproveSignDataCtx721036_QRegisterMetaType()))
}

func (ptr *ApproveSignDataCtx) QRegisterMetaType() int {
	return int(int32(C.ApproveSignDataCtx721036_ApproveSignDataCtx721036_QRegisterMetaType()))
}

func ApproveSignDataCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveSignDataCtx721036_ApproveSignDataCtx721036_QRegisterMetaType2(typeNameC)))
}

func (ptr *ApproveSignDataCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveSignDataCtx721036_ApproveSignDataCtx721036_QRegisterMetaType2(typeNameC)))
}

func ApproveSignDataCtx_QmlRegisterType() int {
	return int(int32(C.ApproveSignDataCtx721036_ApproveSignDataCtx721036_QmlRegisterType()))
}

func (ptr *ApproveSignDataCtx) QmlRegisterType() int {
	return int(int32(C.ApproveSignDataCtx721036_ApproveSignDataCtx721036_QmlRegisterType()))
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
	return int(int32(C.ApproveSignDataCtx721036_ApproveSignDataCtx721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.ApproveSignDataCtx721036_ApproveSignDataCtx721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveSignDataCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ApproveSignDataCtx721036___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ApproveSignDataCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ApproveSignDataCtx721036___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ApproveSignDataCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveSignDataCtx721036___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveSignDataCtx) __findChildren_newList2() unsafe.Pointer {
	return C.ApproveSignDataCtx721036___findChildren_newList2(ptr.Pointer())
}

func (ptr *ApproveSignDataCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveSignDataCtx721036___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveSignDataCtx) __findChildren_newList3() unsafe.Pointer {
	return C.ApproveSignDataCtx721036___findChildren_newList3(ptr.Pointer())
}

func (ptr *ApproveSignDataCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveSignDataCtx721036___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveSignDataCtx) __findChildren_newList() unsafe.Pointer {
	return C.ApproveSignDataCtx721036___findChildren_newList(ptr.Pointer())
}

func (ptr *ApproveSignDataCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveSignDataCtx721036___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveSignDataCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveSignDataCtx) __children_newList() unsafe.Pointer {
	return C.ApproveSignDataCtx721036___children_newList(ptr.Pointer())
}

func NewApproveSignDataCtx(parent std_core.QObject_ITF) *ApproveSignDataCtx {
	tmpValue := NewApproveSignDataCtxFromPointer(C.ApproveSignDataCtx721036_NewApproveSignDataCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApproveSignDataCtx721036_DestroyApproveSignDataCtx
func callbackApproveSignDataCtx721036_DestroyApproveSignDataCtx(ptr unsafe.Pointer) {
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
		C.ApproveSignDataCtx721036_DestroyApproveSignDataCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ApproveSignDataCtx) DestroyApproveSignDataCtxDefault() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036_DestroyApproveSignDataCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveSignDataCtx721036_Event
func callbackApproveSignDataCtx721036_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveSignDataCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ApproveSignDataCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApproveSignDataCtx721036_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackApproveSignDataCtx721036_EventFilter
func callbackApproveSignDataCtx721036_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveSignDataCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ApproveSignDataCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApproveSignDataCtx721036_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackApproveSignDataCtx721036_ChildEvent
func callbackApproveSignDataCtx721036_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ApproveSignDataCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackApproveSignDataCtx721036_ConnectNotify
func callbackApproveSignDataCtx721036_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveSignDataCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveSignDataCtx721036_CustomEvent
func callbackApproveSignDataCtx721036_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ApproveSignDataCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackApproveSignDataCtx721036_DeleteLater
func callbackApproveSignDataCtx721036_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewApproveSignDataCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ApproveSignDataCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveSignDataCtx721036_Destroyed
func callbackApproveSignDataCtx721036_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackApproveSignDataCtx721036_DisconnectNotify
func callbackApproveSignDataCtx721036_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveSignDataCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveSignDataCtx721036_ObjectNameChanged
func callbackApproveSignDataCtx721036_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApproveSignDataCtx721036_TimerEvent
func callbackApproveSignDataCtx721036_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewApproveSignDataCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ApproveSignDataCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveSignDataCtx721036_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
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

//export callbackApproveTxCtx721036_Constructor
func callbackApproveTxCtx721036_Constructor(ptr unsafe.Pointer) {
	this := NewApproveTxCtxFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClicked(this.clicked)
	this.ConnectEdited(this.edited)
	this.init()
}

//export callbackApproveTxCtx721036_Clicked
func callbackApproveTxCtx721036_Clicked(ptr unsafe.Pointer, b C.int) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(int))(int(int32(b)))
	}

}

func (ptr *ApproveTxCtx) ConnectClicked(f func(b int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.ApproveTxCtx721036_ConnectClicked(ptr.Pointer())
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

func (ptr *ApproveTxCtx) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *ApproveTxCtx) Clicked(b int) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036_Clicked(ptr.Pointer(), C.int(int32(b)))
	}
}

//export callbackApproveTxCtx721036_Edited
func callbackApproveTxCtx721036_Edited(ptr unsafe.Pointer, s C.struct_Moc_PackedString, v C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "edited"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(s), cGoUnpackString(v))
	}

}

func (ptr *ApproveTxCtx) ConnectEdited(f func(s string, v string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "edited") {
			C.ApproveTxCtx721036_ConnectEdited(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectEdited(ptr.Pointer())
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
		C.ApproveTxCtx721036_Edited(ptr.Pointer(), C.struct_Moc_PackedString{data: sC, len: C.longlong(len(s))}, C.struct_Moc_PackedString{data: vC, len: C.longlong(len(v))})
	}
}

//export callbackApproveTxCtx721036_Remote
func callbackApproveTxCtx721036_Remote(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_Remote(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) RemoteDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_RemoteDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetRemote
func callbackApproveTxCtx721036_SetRemote(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetRemote(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

func (ptr *ApproveTxCtx) SetRemoteDefault(remote string) {
	if ptr.Pointer() != nil {
		var remoteC *C.char
		if remote != "" {
			remoteC = C.CString(remote)
			defer C.free(unsafe.Pointer(remoteC))
		}
		C.ApproveTxCtx721036_SetRemoteDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveTxCtx721036_RemoteChanged
func callbackApproveTxCtx721036_RemoteChanged(ptr unsafe.Pointer, remote C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "remoteChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(remote))
	}

}

func (ptr *ApproveTxCtx) ConnectRemoteChanged(f func(remote string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "remoteChanged") {
			C.ApproveTxCtx721036_ConnectRemoteChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectRemoteChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_RemoteChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteC, len: C.longlong(len(remote))})
	}
}

//export callbackApproveTxCtx721036_Transport
func callbackApproveTxCtx721036_Transport(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_Transport(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) TransportDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_TransportDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetTransport
func callbackApproveTxCtx721036_SetTransport(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetTransport(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

func (ptr *ApproveTxCtx) SetTransportDefault(transport string) {
	if ptr.Pointer() != nil {
		var transportC *C.char
		if transport != "" {
			transportC = C.CString(transport)
			defer C.free(unsafe.Pointer(transportC))
		}
		C.ApproveTxCtx721036_SetTransportDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveTxCtx721036_TransportChanged
func callbackApproveTxCtx721036_TransportChanged(ptr unsafe.Pointer, transport C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transportChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(transport))
	}

}

func (ptr *ApproveTxCtx) ConnectTransportChanged(f func(transport string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transportChanged") {
			C.ApproveTxCtx721036_ConnectTransportChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectTransportChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_TransportChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transportC, len: C.longlong(len(transport))})
	}
}

//export callbackApproveTxCtx721036_Endpoint
func callbackApproveTxCtx721036_Endpoint(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_Endpoint(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) EndpointDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_EndpointDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetEndpoint
func callbackApproveTxCtx721036_SetEndpoint(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetEndpoint(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

func (ptr *ApproveTxCtx) SetEndpointDefault(endpoint string) {
	if ptr.Pointer() != nil {
		var endpointC *C.char
		if endpoint != "" {
			endpointC = C.CString(endpoint)
			defer C.free(unsafe.Pointer(endpointC))
		}
		C.ApproveTxCtx721036_SetEndpointDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveTxCtx721036_EndpointChanged
func callbackApproveTxCtx721036_EndpointChanged(ptr unsafe.Pointer, endpoint C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "endpointChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(endpoint))
	}

}

func (ptr *ApproveTxCtx) ConnectEndpointChanged(f func(endpoint string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "endpointChanged") {
			C.ApproveTxCtx721036_ConnectEndpointChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectEndpointChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_EndpointChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: endpointC, len: C.longlong(len(endpoint))})
	}
}

//export callbackApproveTxCtx721036_Data
func callbackApproveTxCtx721036_Data(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) DataDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_DataDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetData
func callbackApproveTxCtx721036_SetData(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetData(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

func (ptr *ApproveTxCtx) SetDataDefault(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.ApproveTxCtx721036_SetDataDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackApproveTxCtx721036_DataChanged
func callbackApproveTxCtx721036_DataChanged(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *ApproveTxCtx) ConnectDataChanged(f func(data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataChanged") {
			C.ApproveTxCtx721036_ConnectDataChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectDataChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DataChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackApproveTxCtx721036_From
func callbackApproveTxCtx721036_From(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_From(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) FromDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_FromDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetFrom
func callbackApproveTxCtx721036_SetFrom(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetFrom(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

func (ptr *ApproveTxCtx) SetFromDefault(from string) {
	if ptr.Pointer() != nil {
		var fromC *C.char
		if from != "" {
			fromC = C.CString(from)
			defer C.free(unsafe.Pointer(fromC))
		}
		C.ApproveTxCtx721036_SetFromDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveTxCtx721036_FromChanged
func callbackApproveTxCtx721036_FromChanged(ptr unsafe.Pointer, from C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fromChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(from))
	}

}

func (ptr *ApproveTxCtx) ConnectFromChanged(f func(from string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fromChanged") {
			C.ApproveTxCtx721036_ConnectFromChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectFromChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_FromChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fromC, len: C.longlong(len(from))})
	}
}

//export callbackApproveTxCtx721036_To
func callbackApproveTxCtx721036_To(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_To(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) ToDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_ToDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetTo
func callbackApproveTxCtx721036_SetTo(ptr unsafe.Pointer, to C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetTo(ptr.Pointer(), C.struct_Moc_PackedString{data: toC, len: C.longlong(len(to))})
	}
}

func (ptr *ApproveTxCtx) SetToDefault(to string) {
	if ptr.Pointer() != nil {
		var toC *C.char
		if to != "" {
			toC = C.CString(to)
			defer C.free(unsafe.Pointer(toC))
		}
		C.ApproveTxCtx721036_SetToDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: toC, len: C.longlong(len(to))})
	}
}

//export callbackApproveTxCtx721036_ToChanged
func callbackApproveTxCtx721036_ToChanged(ptr unsafe.Pointer, to C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "toChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(to))
	}

}

func (ptr *ApproveTxCtx) ConnectToChanged(f func(to string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "toChanged") {
			C.ApproveTxCtx721036_ConnectToChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectToChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_ToChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: toC, len: C.longlong(len(to))})
	}
}

//export callbackApproveTxCtx721036_Gas
func callbackApproveTxCtx721036_Gas(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_Gas(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) GasDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_GasDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetGas
func callbackApproveTxCtx721036_SetGas(ptr unsafe.Pointer, gas C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetGas(ptr.Pointer(), C.struct_Moc_PackedString{data: gasC, len: C.longlong(len(gas))})
	}
}

func (ptr *ApproveTxCtx) SetGasDefault(gas string) {
	if ptr.Pointer() != nil {
		var gasC *C.char
		if gas != "" {
			gasC = C.CString(gas)
			defer C.free(unsafe.Pointer(gasC))
		}
		C.ApproveTxCtx721036_SetGasDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: gasC, len: C.longlong(len(gas))})
	}
}

//export callbackApproveTxCtx721036_GasChanged
func callbackApproveTxCtx721036_GasChanged(ptr unsafe.Pointer, gas C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "gasChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(gas))
	}

}

func (ptr *ApproveTxCtx) ConnectGasChanged(f func(gas string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gasChanged") {
			C.ApproveTxCtx721036_ConnectGasChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectGasChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_GasChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: gasC, len: C.longlong(len(gas))})
	}
}

//export callbackApproveTxCtx721036_GasPrice
func callbackApproveTxCtx721036_GasPrice(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_GasPrice(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) GasPriceDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_GasPriceDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetGasPrice
func callbackApproveTxCtx721036_SetGasPrice(ptr unsafe.Pointer, gasPrice C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetGasPrice(ptr.Pointer(), C.struct_Moc_PackedString{data: gasPriceC, len: C.longlong(len(gasPrice))})
	}
}

func (ptr *ApproveTxCtx) SetGasPriceDefault(gasPrice string) {
	if ptr.Pointer() != nil {
		var gasPriceC *C.char
		if gasPrice != "" {
			gasPriceC = C.CString(gasPrice)
			defer C.free(unsafe.Pointer(gasPriceC))
		}
		C.ApproveTxCtx721036_SetGasPriceDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: gasPriceC, len: C.longlong(len(gasPrice))})
	}
}

//export callbackApproveTxCtx721036_GasPriceChanged
func callbackApproveTxCtx721036_GasPriceChanged(ptr unsafe.Pointer, gasPrice C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "gasPriceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(gasPrice))
	}

}

func (ptr *ApproveTxCtx) ConnectGasPriceChanged(f func(gasPrice string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gasPriceChanged") {
			C.ApproveTxCtx721036_ConnectGasPriceChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectGasPriceChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_GasPriceChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: gasPriceC, len: C.longlong(len(gasPrice))})
	}
}

//export callbackApproveTxCtx721036_Nonce
func callbackApproveTxCtx721036_Nonce(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_Nonce(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) NonceDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_NonceDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetNonce
func callbackApproveTxCtx721036_SetNonce(ptr unsafe.Pointer, nonce C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetNonce(ptr.Pointer(), C.struct_Moc_PackedString{data: nonceC, len: C.longlong(len(nonce))})
	}
}

func (ptr *ApproveTxCtx) SetNonceDefault(nonce string) {
	if ptr.Pointer() != nil {
		var nonceC *C.char
		if nonce != "" {
			nonceC = C.CString(nonce)
			defer C.free(unsafe.Pointer(nonceC))
		}
		C.ApproveTxCtx721036_SetNonceDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: nonceC, len: C.longlong(len(nonce))})
	}
}

//export callbackApproveTxCtx721036_NonceChanged
func callbackApproveTxCtx721036_NonceChanged(ptr unsafe.Pointer, nonce C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "nonceChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(nonce))
	}

}

func (ptr *ApproveTxCtx) ConnectNonceChanged(f func(nonce string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nonceChanged") {
			C.ApproveTxCtx721036_ConnectNonceChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectNonceChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_NonceChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: nonceC, len: C.longlong(len(nonce))})
	}
}

//export callbackApproveTxCtx721036_Value
func callbackApproveTxCtx721036_Value(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.ApproveTxCtx721036_Value(ptr.Pointer()))
	}
	return ""
}

func (ptr *ApproveTxCtx) ValueDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ApproveTxCtx721036_ValueDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackApproveTxCtx721036_SetValue
func callbackApproveTxCtx721036_SetValue(ptr unsafe.Pointer, value C.struct_Moc_PackedString) {
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
		C.ApproveTxCtx721036_SetValue(ptr.Pointer(), C.struct_Moc_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *ApproveTxCtx) SetValueDefault(value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.ApproveTxCtx721036_SetValueDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

//export callbackApproveTxCtx721036_ValueChanged
func callbackApproveTxCtx721036_ValueChanged(ptr unsafe.Pointer, value C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "valueChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	}

}

func (ptr *ApproveTxCtx) ConnectValueChanged(f func(value string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valueChanged") {
			C.ApproveTxCtx721036_ConnectValueChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_DisconnectValueChanged(ptr.Pointer())
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
		C.ApproveTxCtx721036_ValueChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func ApproveTxCtx_QRegisterMetaType() int {
	return int(int32(C.ApproveTxCtx721036_ApproveTxCtx721036_QRegisterMetaType()))
}

func (ptr *ApproveTxCtx) QRegisterMetaType() int {
	return int(int32(C.ApproveTxCtx721036_ApproveTxCtx721036_QRegisterMetaType()))
}

func ApproveTxCtx_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveTxCtx721036_ApproveTxCtx721036_QRegisterMetaType2(typeNameC)))
}

func (ptr *ApproveTxCtx) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApproveTxCtx721036_ApproveTxCtx721036_QRegisterMetaType2(typeNameC)))
}

func ApproveTxCtx_QmlRegisterType() int {
	return int(int32(C.ApproveTxCtx721036_ApproveTxCtx721036_QmlRegisterType()))
}

func (ptr *ApproveTxCtx) QmlRegisterType() int {
	return int(int32(C.ApproveTxCtx721036_ApproveTxCtx721036_QmlRegisterType()))
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
	return int(int32(C.ApproveTxCtx721036_ApproveTxCtx721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.ApproveTxCtx721036_ApproveTxCtx721036_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApproveTxCtx) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ApproveTxCtx721036___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ApproveTxCtx) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ApproveTxCtx721036___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ApproveTxCtx) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveTxCtx721036___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveTxCtx) __findChildren_newList2() unsafe.Pointer {
	return C.ApproveTxCtx721036___findChildren_newList2(ptr.Pointer())
}

func (ptr *ApproveTxCtx) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveTxCtx721036___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveTxCtx) __findChildren_newList3() unsafe.Pointer {
	return C.ApproveTxCtx721036___findChildren_newList3(ptr.Pointer())
}

func (ptr *ApproveTxCtx) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveTxCtx721036___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveTxCtx) __findChildren_newList() unsafe.Pointer {
	return C.ApproveTxCtx721036___findChildren_newList(ptr.Pointer())
}

func (ptr *ApproveTxCtx) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ApproveTxCtx721036___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApproveTxCtx) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApproveTxCtx) __children_newList() unsafe.Pointer {
	return C.ApproveTxCtx721036___children_newList(ptr.Pointer())
}

func NewApproveTxCtx(parent std_core.QObject_ITF) *ApproveTxCtx {
	tmpValue := NewApproveTxCtxFromPointer(C.ApproveTxCtx721036_NewApproveTxCtx(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApproveTxCtx721036_DestroyApproveTxCtx
func callbackApproveTxCtx721036_DestroyApproveTxCtx(ptr unsafe.Pointer) {
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
		C.ApproveTxCtx721036_DestroyApproveTxCtx(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ApproveTxCtx) DestroyApproveTxCtxDefault() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036_DestroyApproveTxCtxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveTxCtx721036_Event
func callbackApproveTxCtx721036_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveTxCtxFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ApproveTxCtx) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApproveTxCtx721036_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackApproveTxCtx721036_EventFilter
func callbackApproveTxCtx721036_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApproveTxCtxFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ApproveTxCtx) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApproveTxCtx721036_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackApproveTxCtx721036_ChildEvent
func callbackApproveTxCtx721036_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewApproveTxCtxFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ApproveTxCtx) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackApproveTxCtx721036_ConnectNotify
func callbackApproveTxCtx721036_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveTxCtxFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveTxCtx) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveTxCtx721036_CustomEvent
func callbackApproveTxCtx721036_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewApproveTxCtxFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ApproveTxCtx) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackApproveTxCtx721036_DeleteLater
func callbackApproveTxCtx721036_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewApproveTxCtxFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ApproveTxCtx) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApproveTxCtx721036_Destroyed
func callbackApproveTxCtx721036_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackApproveTxCtx721036_DisconnectNotify
func callbackApproveTxCtx721036_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApproveTxCtxFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApproveTxCtx) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApproveTxCtx721036_ObjectNameChanged
func callbackApproveTxCtx721036_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApproveTxCtx721036_TimerEvent
func callbackApproveTxCtx721036_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewApproveTxCtxFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ApproveTxCtx) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApproveTxCtx721036_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
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

//export callbackCustomListModel721036_Constructor
func callbackCustomListModel721036_Constructor(ptr unsafe.Pointer) {
	this := NewCustomListModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClear(this.clear)
	this.ConnectAdd(this.add)
	this.ConnectClicked(this.clicked)
	this.ConnectOnCheckStateChanged(this.onCheckStateChanged)
	this.init()
}

//export callbackCustomListModel721036_Clear
func callbackCustomListModel721036_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		signal.(func())()
	}

}

func (ptr *CustomListModel) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clear") {
			C.CustomListModel721036_ConnectClear(ptr.Pointer())
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
		C.CustomListModel721036_DisconnectClear(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *CustomListModel) Clear() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_Clear(ptr.Pointer())
	}
}

//export callbackCustomListModel721036_Add
func callbackCustomListModel721036_Add(ptr unsafe.Pointer, account C.uintptr_t) {
	var accountD custom_params_aaf9d9m.ApproveListingAccount
	if accountI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(account))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(account)))
		accountD = accountI.(custom_params_aaf9d9m.ApproveListingAccount)
	}
	if signal := qt.GetSignal(ptr, "add"); signal != nil {
		signal.(func(custom_params_aaf9d9m.ApproveListingAccount))(accountD)
	}

}

func (ptr *CustomListModel) ConnectAdd(f func(account custom_params_aaf9d9m.ApproveListingAccount)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "add") {
			C.CustomListModel721036_ConnectAdd(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "add"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "add", func(account custom_params_aaf9d9m.ApproveListingAccount) {
				signal.(func(custom_params_aaf9d9m.ApproveListingAccount))(account)
				f(account)
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

func (ptr *CustomListModel) Add(account custom_params_aaf9d9m.ApproveListingAccount) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&account), account)
		C.CustomListModel721036_Add(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&account))))
	}
}

//export callbackCustomListModel721036_Clicked
func callbackCustomListModel721036_Clicked(ptr unsafe.Pointer, b C.int) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		signal.(func(int))(int(int32(b)))
	}

}

func (ptr *CustomListModel) ConnectClicked(f func(b int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.CustomListModel721036_ConnectClicked(ptr.Pointer())
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

func (ptr *CustomListModel) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *CustomListModel) Clicked(b int) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_Clicked(ptr.Pointer(), C.int(int32(b)))
	}
}

//export callbackCustomListModel721036_OnCheckStateChanged
func callbackCustomListModel721036_OnCheckStateChanged(ptr unsafe.Pointer, i C.int, checked C.char) {
	if signal := qt.GetSignal(ptr, "onCheckStateChanged"); signal != nil {
		signal.(func(int, bool))(int(int32(i)), int8(checked) != 0)
	}

}

func (ptr *CustomListModel) ConnectOnCheckStateChanged(f func(i int, checked bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "onCheckStateChanged") {
			C.CustomListModel721036_ConnectOnCheckStateChanged(ptr.Pointer())
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

func (ptr *CustomListModel) DisconnectOnCheckStateChanged() {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_DisconnectOnCheckStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "onCheckStateChanged")
	}
}

func (ptr *CustomListModel) OnCheckStateChanged(i int, checked bool) {
	if ptr.Pointer() != nil {
		C.CustomListModel721036_OnCheckStateChanged(ptr.Pointer(), C.int(int32(i)), C.char(int8(qt.GoBoolToInt(checked))))
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
