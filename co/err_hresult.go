//go:build windows

package co

import (
	"fmt"
	"syscall"
)

// [HRESULT] error codes.
//
// We can't simply use syscall.Errno because it's an uintptr (8 bytes), thus a
// native struct with such a field type would be wrong. However, the Unwrap()
// method will return the syscall.Errno value.
//
// [HRESULT]: https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref/0642cb2f-2075-4469-918c-4441e69c548a
type HRESULT uint32

// Implements error interface.
func (hr HRESULT) Error() string {
	return hr.String()
}

// Returns the contained syscall.Errno.
func (hr HRESULT) Unwrap() error {
	return syscall.Errno(hr)
}

// Calls [FormatMessage] and returns a full error description.
//
// [FormatMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-formatmessagew
func (hr HRESULT) String() string {
	return fmt.Sprintf("[%d 0x%02x] %s",
		uint32(hr), uint32(hr), hr.Unwrap().Error())
}

// Returns the HRESULT code.
func (hr HRESULT) Code() uint16 {
	return uint16(uint32(hr) & 0xffff)
}

// Returns the HRESULT facility.
func (hr HRESULT) Facility() FACILITY {
	return FACILITY((uint32(hr) >> 16) & 0x1fff)
}

// [FAILED] macro.
//
// [FAILED]: https://learn.microsoft.com/en-us/windows/win32/api/winerror/nf-winerror-failed
func (hr HRESULT) Failed() bool {
	return int32(hr) < 0
}

// Returns the HRESULT severity.
func (hr HRESULT) Severity() SEVERITY {
	return SEVERITY((uint32(hr) >> 31) & 0x1)
}

// [SUCCEEDED] macro.
//
// [SUCCEEDED]: https://learn.microsoft.com/en-us/windows/win32/api/winerror/nf-winerror-succeeded
func (hr HRESULT) Succeeded() bool {
	return int32(hr) >= 0
}

const (
	HRESULT_S_OK    HRESULT = 0 // The operation completed successfully.
	HRESULT_S_FALSE HRESULT = 1 // The operation completed successfully returning FALSE.

	HRESULT_CLASS_E_NOAGGREGATION     HRESULT = 0x8004_0110 // Class does not support aggregation (or class object is remote).
	HRESULT_CLASS_E_CLASSNOTAVAILABLE HRESULT = 0x8004_0111 // ClassFactory cannot supply requested class.
	HRESULT_CLASS_E_NOTLICENSED       HRESULT = 0x8004_0112 // Class is not licensed for use.

	HRESULT_CO_E_NOTINITIALIZED     HRESULT = 0x8004_01f0 // CoInitialize has not been called.
	HRESULT_CO_E_ALREADYINITIALIZED HRESULT = 0x8004_01f1 // CoInitialize has already been called.
	HRESULT_CO_E_CANTDETERMINECLASS HRESULT = 0x8004_01f2 // Class of object cannot be determined.
	HRESULT_CO_E_CLASSSTRING        HRESULT = 0x8004_01f3 // Invalid class string.
	HRESULT_CO_E_IIDSTRING          HRESULT = 0x8004_01f4 // Invalid interface string.
	HRESULT_CO_E_APPNOTFOUND        HRESULT = 0x8004_01f5 // Application not found.
	HRESULT_CO_E_APPSINGLEUSE       HRESULT = 0x8004_01f6 // Application cannot be run more than once.
	HRESULT_CO_E_ERRORINAPP         HRESULT = 0x8004_01f7 // Some error in application program.
	HRESULT_CO_E_DLLNOTFOUND        HRESULT = 0x8004_01f8 // DLL for class not found.
	HRESULT_CO_E_ERRORINDLL         HRESULT = 0x8004_01f9 // Error in the DLL.
	HRESULT_CO_E_WRONGOSFORAPP      HRESULT = 0x8004_01fa // Wrong OS or OS version for application.
	HRESULT_CO_E_OBJNOTREG          HRESULT = 0x8004_01fb // Object is not registered.
	HRESULT_CO_E_OBJISREG           HRESULT = 0x8004_01fc // Object is already registered.
	HRESULT_CO_E_OBJNOTCONNECTED    HRESULT = 0x8004_01fd // Object is not connected to server.
	HRESULT_CO_E_APPDIDNTREG        HRESULT = 0x8004_01fe // Application was launched but it didn't register a class factory.
	HRESULT_CO_E_RELEASED           HRESULT = 0x8004_01ff // Object has been released.

	HRESULT_DISP_E_UNKNOWNINTERFACE HRESULT = 0x8002_0001 // Unknown interface.
	HRESULT_DISP_E_MEMBERNOTFOUND   HRESULT = 0x8002_0003 // Member not found.
	HRESULT_DISP_E_PARAMNOTFOUND    HRESULT = 0x8002_0004 // Parameter not found.
	HRESULT_DISP_E_TYPEMISMATCH     HRESULT = 0x8002_0005 // Type mismatch.
	HRESULT_DISP_E_UNKNOWNNAME      HRESULT = 0x8002_0006 // Unknown name.
	HRESULT_DISP_E_NONAMEDARGS      HRESULT = 0x8002_0007 // No named arguments.
	HRESULT_DISP_E_BADVARTYPE       HRESULT = 0x8002_0008 // Bad variable type.
	HRESULT_DISP_E_EXCEPTION        HRESULT = 0x8002_0009 // Exception occurred.
	HRESULT_DISP_E_OVERFLOW         HRESULT = 0x8002_000a // Out of present range.
	HRESULT_DISP_E_BADINDEX         HRESULT = 0x8002_000b // Invalid index.
	HRESULT_DISP_E_UNKNOWNLCID      HRESULT = 0x8002_000c // Unknown language.
	HRESULT_DISP_E_ARRAYISLOCKED    HRESULT = 0x8002_000d // Memory is locked.
	HRESULT_DISP_E_BADPARAMCOUNT    HRESULT = 0x8002_000e // Invalid number of parameters.
	HRESULT_DISP_E_PARAMNOTOPTIONAL HRESULT = 0x8002_000f // Parameter not optional.
	HRESULT_DISP_E_BADCALLEE        HRESULT = 0x8002_0010 // Invalid callee.
	HRESULT_DISP_E_NOTACOLLECTION   HRESULT = 0x8002_0011 // Does not support a collection.
	HRESULT_DISP_E_DIVBYZERO        HRESULT = 0x8002_0012 // Division by zero.
	HRESULT_DISP_E_BUFFERTOOSMALL   HRESULT = 0x8002_0013 // Buffer too small.

	HRESULT_DRAGDROP_E_NOTREGISTERED             HRESULT = 0x8004_0100 // Trying to revoke a drop target that has not been registered.
	HRESULT_DRAGDROP_E_ALREADYREGISTERED         HRESULT = 0x8004_0101 // This window has already been registered as a drop target.
	HRESULT_DRAGDROP_E_INVALIDHWND               HRESULT = 0x8004_0102 // Invalid window handle.
	HRESULT_DRAGDROP_E_CONCURRENT_DRAG_ATTEMPTED HRESULT = 0x8004_0103 // A drag operation is already in progress.

	HRESULT_DV_E_FORMATETC           HRESULT = 0x80040064 // Invalid FORMATETC structure.
	HRESULT_DV_E_DVTARGETDEVICE      HRESULT = 0x80040065 // Invalid DVTARGETDEVICE structure.
	HRESULT_DV_E_STGMEDIUM           HRESULT = 0x80040066 // Invalid STDGMEDIUM structure.
	HRESULT_DV_E_STATDATA            HRESULT = 0x80040067 // Invalid STATDATA structure.
	HRESULT_DV_E_LINDEX              HRESULT = 0x80040068 // Invalid lindex.
	HRESULT_DV_E_TYMED               HRESULT = 0x80040069 // Invalid tymed.
	HRESULT_DV_E_CLIPFORMAT          HRESULT = 0x8004006a // Invalid clipboard format.
	HRESULT_DV_E_DVASPECT            HRESULT = 0x8004006b // Invalid aspect(s).
	HRESULT_DV_E_DVTARGETDEVICE_SIZE HRESULT = 0x8004006c // tdSize parameter of the DVTARGETDEVICE structure is invalid.
	HRESULT_DV_E_NOIVIEWOBJECT       HRESULT = 0x8004006d // Object doesn't support IViewObject interface.

	HRESULT_E_UNEXPECTED           HRESULT = 0x8000_ffff // Catastrophic failure.
	HRESULT_E_NOTIMPL              HRESULT = 0x8000_4001 // Not implemented.
	HRESULT_E_OUTOFMEMORY          HRESULT = 0x8007_000e // Not enough memory resources are available to complete this operation.
	HRESULT_E_INVALIDARG           HRESULT = 0x8007_0057 // The parameter is incorrect.
	HRESULT_E_NOINTERFACE          HRESULT = 0x8000_4002 // No such interface supported.
	HRESULT_E_POINTER              HRESULT = 0x8000_4003 // Invalid pointer.
	HRESULT_E_HANDLE               HRESULT = 0x8007_0006 // The handle is invalid.
	HRESULT_E_ABORT                HRESULT = 0x8000_4004 // Operation aborted.
	HRESULT_E_FAIL                 HRESULT = 0x8000_4005 // Unspecified error.
	HRESULT_E_ACCESSDENIED         HRESULT = 0x8007_0005 // Access is denied.
	HRESULT_E_PENDING              HRESULT = 0x8000_000a // The data necessary to complete this operation is not yet available.
	HRESULT_E_BOUNDS               HRESULT = 0x8000_000b // The operation attempted to access data outside the valid range.
	HRESULT_E_CHANGED_STATE        HRESULT = 0x8000_000c // A concurrent or interleaved operation changed the state of the object, invalidating this operation.
	HRESULT_E_ILLEGAL_STATE_CHANGE HRESULT = 0x8000_000d // An illegal state change was requested.
	HRESULT_E_ILLEGAL_METHOD_CALL  HRESULT = 0x8000_000e // A method was called at an unexpected time.

	HRESULT_OLE_E_OLEVERB             HRESULT = 0x8004_0000 // Invalid OLEVERB structure.
	HRESULT_OLE_E_ADVF                HRESULT = 0x8004_0001 // Invalid advise flags.
	HRESULT_OLE_E_ENUM_NOMORE         HRESULT = 0x8004_0002 // Can't enumerate any more, because the associated data is missing.
	HRESULT_OLE_E_ADVISENOTSUPPORTED  HRESULT = 0x8004_0003 // This implementation doesn't take advises.
	HRESULT_OLE_E_NOCONNECTION        HRESULT = 0x8004_0004 // There is no connection for this connection ID.
	HRESULT_OLE_E_NOTRUNNING          HRESULT = 0x8004_0005 // Need to run the object to perform this operation.
	HRESULT_OLE_E_NOCACHE             HRESULT = 0x8004_0006 // There is no cache to operate on.
	HRESULT_OLE_E_BLANK               HRESULT = 0x8004_0007 // Uninitialized object.
	HRESULT_OLE_E_CLASSDIFF           HRESULT = 0x8004_0008 // Linked object's source class has changed.
	HRESULT_OLE_E_CANT_GETMONIKER     HRESULT = 0x8004_0009 // Not able to get the moniker of the object.
	HRESULT_OLE_E_CANT_BINDTOSOURCE   HRESULT = 0x8004_000a // Not able to bind to the source.
	HRESULT_OLE_E_STATIC              HRESULT = 0x8004_000b // Object is static; operation not allowed.
	HRESULT_OLE_E_PROMPTSAVECANCELLED HRESULT = 0x8004_000c // User canceled out of save dialog.
	HRESULT_OLE_E_INVALIDRECT         HRESULT = 0x8004_000d // Invalid rectangle.
	HRESULT_OLE_E_WRONGCOMPOBJ        HRESULT = 0x8004_000e // compobj.dll is too old for the ole2.dll initialized.
	HRESULT_OLE_E_INVALIDHWND         HRESULT = 0x8004_000f // Invalid window handle.
	HRESULT_OLE_E_NOT_INPLACEACTIVE   HRESULT = 0x8004_0010 // Object is not in any of the inplace active states.
	HRESULT_OLE_E_CANTCONVERT         HRESULT = 0x8004_0011 // Not able to convert object.
	HRESULT_OLE_E_NOSTORAGE           HRESULT = 0x8004_0012 // Not able to perform the operation because object is not given storage yet.

	HRESULT_REGDB_E_READREGDB         HRESULT = 0x8004_0150 // Could not read key from registry.
	HRESULT_REGDB_E_WRITEREGDB        HRESULT = 0x8004_0151 // Could not write key to registry.
	HRESULT_REGDB_E_KEYMISSING        HRESULT = 0x8004_0152 // Could not find the key in the registry.
	HRESULT_REGDB_E_INVALIDVALUE      HRESULT = 0x8004_0153 // Invalid value for registry.
	HRESULT_REGDB_E_CLASSNOTREG       HRESULT = 0x8004_0154 // Class not registered.
	HRESULT_REGDB_E_IIDNOTREG         HRESULT = 0x8004_0155 // Interface not registered.
	HRESULT_REGDB_E_BADTHREADINGMODEL HRESULT = 0x8004_0156 // Threading model entry is not valid.

	HRESULT_RPC_E_CALL_REJECTED               HRESULT = 0x8001_0001 // Call was rejected by callee.
	HRESULT_RPC_E_CALL_CANCELED               HRESULT = 0x8001_0002 // Call was canceled by the message filter.
	HRESULT_RPC_E_CANTPOST_INSENDCALL         HRESULT = 0x8001_0003 // The caller is dispatching an intertask SendMessage call and cannot call out via PostMessage.
	HRESULT_RPC_E_CANTCALLOUT_INASYNCCALL     HRESULT = 0x8001_0004 // The caller is dispatching an asynchronous call and cannot make an outgoing call on behalf of this call.
	HRESULT_RPC_E_CANTCALLOUT_INEXTERNALCALL  HRESULT = 0x8001_0005 // It is illegal to call out while inside message filter.
	HRESULT_RPC_E_CONNECTION_TERMINATED       HRESULT = 0x8001_0006 // The connection terminated or is in a bogus state and cannot be used any more. Other connections are still valid.
	HRESULT_RPC_E_SERVER_DIED                 HRESULT = 0x8001_0007 // The callee (server [not server application]) is not available and disappeared; all connections are invalid. The call may have executed.
	HRESULT_RPC_E_CLIENT_DIED                 HRESULT = 0x8001_0008 // The caller (client) disappeared while the callee (server) was processing a call.
	HRESULT_RPC_E_INVALID_DATAPACKET          HRESULT = 0x8001_0009 // The data packet with the marshalled parameter data is incorrect.
	HRESULT_RPC_E_CANTTRANSMIT_CALL           HRESULT = 0x8001_000a // The call was not transmitted properly; the message queue was full and was not emptied after yielding.
	HRESULT_RPC_E_CLIENT_CANTMARSHAL_DATA     HRESULT = 0x8001_000b // The client (caller) cannot marshall the parameter data - low memory, etc.
	HRESULT_RPC_E_CLIENT_CANTUNMARSHAL_DATA   HRESULT = 0x8001_000c // The client (caller) cannot unmarshall the return data - low memory, etc.
	HRESULT_RPC_E_SERVER_CANTMARSHAL_DATA     HRESULT = 0x8001_000d // The server (callee) cannot marshall the return data - low memory, etc.
	HRESULT_RPC_E_SERVER_CANTUNMARSHAL_DATA   HRESULT = 0x8001_000e // The server (callee) cannot unmarshall the parameter data - low memory, etc.
	HRESULT_RPC_E_INVALID_DATA                HRESULT = 0x8001_000f // Received data is invalid; could be server or client data.
	HRESULT_RPC_E_INVALID_PARAMETER           HRESULT = 0x8001_0010 // A particular parameter is invalid and cannot be (un)marshalled.
	HRESULT_RPC_E_CANTCALLOUT_AGAIN           HRESULT = 0x8001_0011 // There is no second outgoing call on same channel in DDE conversation.
	HRESULT_RPC_E_SERVER_DIED_DNE             HRESULT = 0x8001_0012 // The callee (server [not server application]) is not available and disappeared; all connections are invalid. The call did not execute.
	HRESULT_RPC_E_SYS_CALL_FAILED             HRESULT = 0x8001_0100 // System call failed.
	HRESULT_RPC_E_OUT_OF_RESOURCES            HRESULT = 0x8001_0101 // Could not allocate some required resource (memory, events, ...)
	HRESULT_RPC_E_ATTEMPTED_MULTITHREAD       HRESULT = 0x8001_0102 // Attempted to make calls on more than one thread in single threaded mode.
	HRESULT_RPC_E_NOT_REGISTERED              HRESULT = 0x8001_0103 // The requested interface is not registered on the server object.
	HRESULT_RPC_E_FAULT                       HRESULT = 0x8001_0104 // RPC could not call the server or could not return the results of calling the server.
	HRESULT_RPC_E_SERVERFAULT                 HRESULT = 0x8001_0105 // The server threw an exception.
	HRESULT_RPC_E_CHANGED_MODE                HRESULT = 0x8001_0106 // Cannot change thread mode after it is set.
	HRESULT_RPC_E_INVALIDMETHOD               HRESULT = 0x8001_0107 // The method called does not exist on the server.
	HRESULT_RPC_E_DISCONNECTED                HRESULT = 0x8001_0108 // The object invoked has disconnected from its clients.
	HRESULT_RPC_E_RETRY                       HRESULT = 0x8001_0109 // The object invoked chose not to process the call now. Try again later.
	HRESULT_RPC_E_SERVERCALL_RETRYLATER       HRESULT = 0x8001_010a // The message filter indicated that the application is busy.
	HRESULT_RPC_E_SERVERCALL_REJECTED         HRESULT = 0x8001_010b // The message filter rejected the call.
	HRESULT_RPC_E_INVALID_CALLDATA            HRESULT = 0x8001_010c // A call control interfaces was called with invalid data.
	HRESULT_RPC_E_CANTCALLOUT_ININPUTSYNCCALL HRESULT = 0x8001_010d // An outgoing call cannot be made since the application is dispatching an input-synchronous call.
	HRESULT_RPC_E_WRONG_THREAD                HRESULT = 0x8001_010e // The application called an interface that was marshalled for a different thread.
	HRESULT_RPC_E_THREAD_NOT_INIT             HRESULT = 0x8001_010f // CoInitialize has not been called on the current thread.
	HRESULT_RPC_E_VERSION_MISMATCH            HRESULT = 0x8001_0110 // The version of OLE on the client and server machines does not match.
	HRESULT_RPC_E_INVALID_HEADER              HRESULT = 0x8001_0111 // OLE received a packet with an invalid header.
	HRESULT_RPC_E_INVALID_EXTENSION           HRESULT = 0x8001_0112 // OLE received a packet with an invalid extension.
	HRESULT_RPC_E_INVALID_IPID                HRESULT = 0x8001_0113 // The requested object or interface does not exist.
	HRESULT_RPC_E_INVALID_OBJECT              HRESULT = 0x8001_0114 // The requested object does not exist.
	HRESULT_RPC_S_CALLPENDING                 HRESULT = 0x8001_0115 // OLE has sent a request and is waiting for a reply.
	HRESULT_RPC_S_WAITONTIMER                 HRESULT = 0x8001_0116 // OLE is waiting before retrying a request.
	HRESULT_RPC_E_CALL_COMPLETE               HRESULT = 0x8001_0117 // Call context cannot be accessed after call completed.
	HRESULT_RPC_E_UNSECURE_CALL               HRESULT = 0x8001_0118 // Impersonate on unsecure calls is not supported.
	HRESULT_RPC_E_TOO_LATE                    HRESULT = 0x8001_0119 // Security must be initialized before any interfaces are marshalled or unmarshalled. It cannot be changed once initialized.
	HRESULT_RPC_E_NO_GOOD_SECURITY_PACKAGES   HRESULT = 0x8001_011a // No security packages are installed on this machine or the user is not logged on or there are no compatible security packages between the client and server.
	HRESULT_RPC_E_ACCESS_DENIED               HRESULT = 0x8001_011b // Access is denied.
	HRESULT_RPC_E_REMOTE_DISABLED             HRESULT = 0x8001_011c // Remote calls are not allowed for this process.
	HRESULT_RPC_E_INVALID_OBJREF              HRESULT = 0x8001_011d // The marshaled interface data packet (OBJREF) has an invalid or unknown format.
	HRESULT_RPC_E_NO_CONTEXT                  HRESULT = 0x8001_011e // No context is associated with this call. This happens for some custom marshalled calls and on the client side of the call.
	HRESULT_RPC_E_TIMEOUT                     HRESULT = 0x8001_011f // This operation returned because the timeout period expired.
	HRESULT_RPC_E_NO_SYNC                     HRESULT = 0x8001_0120 // There are no synchronize objects to wait on.
	HRESULT_RPC_E_FULLSIC_REQUIRED            HRESULT = 0x8001_0121 // Full subject issuer chain SSL principal name expected from the server.
	HRESULT_RPC_E_INVALID_STD_NAME            HRESULT = 0x8001_0122 // Principal name is not a valid MSSTD name.
)
