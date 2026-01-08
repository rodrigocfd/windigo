//go:build windows

package co

import (
	"fmt"
	"syscall"
)

// [HRESULT] error codes.
//
// We can't simply use syscall.Errno because it's an uintptr (8 bytes), thus a
// native struct with such a field type would be wrong. However, the
// [HRESULT.Unwrap] method will return the syscall.Errno value.
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

	_HRESULT_CONNECT_E_FIRST        = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_ITF)<<16 | 0x0200
	HRESULT_CONNECT_E_NOCONNECTION  = _HRESULT_CONNECT_E_FIRST + 0 // There is no connection for this connection ID.
	HRESULT_CONNECT_E_ADVISELIMIT   = _HRESULT_CONNECT_E_FIRST + 1 // This implementation's limit for advisory connections has been reached.
	HRESULT_CONNECT_E_CANNOTCONNECT = _HRESULT_CONNECT_E_FIRST + 2 // Connection attempt failed.
	HRESULT_CONNECT_E_OVERRIDDEN    = _HRESULT_CONNECT_E_FIRST + 3 // Must use a derived interface to connect.

	HRESULT_CTL_E_ILLEGALFUNCTIONCALL       = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 5
	HRESULT_CTL_E_OVERFLOW                  = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 6
	HRESULT_CTL_E_OUTOFMEMORY               = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 7
	HRESULT_CTL_E_DIVISIONBYZERO            = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 11
	HRESULT_CTL_E_OUTOFSTRINGSPACE          = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 14
	HRESULT_CTL_E_OUTOFSTACKSPACE           = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 28
	HRESULT_CTL_E_BADFILENAMEORNUMBER       = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 52
	HRESULT_CTL_E_FILENOTFOUND              = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 53
	HRESULT_CTL_E_BADFILEMODE               = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 54
	HRESULT_CTL_E_FILEALREADYOPEN           = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 55
	HRESULT_CTL_E_DEVICEIOERROR             = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 57
	HRESULT_CTL_E_FILEALREADYEXISTS         = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 58
	HRESULT_CTL_E_BADRECORDLENGTH           = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 59
	HRESULT_CTL_E_DISKFULL                  = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 61
	HRESULT_CTL_E_BADRECORDNUMBER           = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 63
	HRESULT_CTL_E_BADFILENAME               = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 64
	HRESULT_CTL_E_TOOMANYFILES              = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 67
	HRESULT_CTL_E_DEVICEUNAVAILABLE         = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 68
	HRESULT_CTL_E_PERMISSIONDENIED          = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 70
	HRESULT_CTL_E_DISKNOTREADY              = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 71
	HRESULT_CTL_E_PATHFILEACCESSERROR       = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 75
	HRESULT_CTL_E_PATHNOTFOUND              = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 76
	HRESULT_CTL_E_INVALIDPATTERNSTRING      = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 93
	HRESULT_CTL_E_INVALIDUSEOFNULL          = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 94
	HRESULT_CTL_E_INVALIDFILEFORMAT         = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 321
	HRESULT_CTL_E_INVALIDPROPERTYVALUE      = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 380
	HRESULT_CTL_E_INVALIDPROPERTYARRAYINDEX = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 381
	HRESULT_CTL_E_SETNOTSUPPORTEDATRUNTIME  = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 382
	HRESULT_CTL_E_SETNOTSUPPORTED           = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 383
	HRESULT_CTL_E_NEEDPROPERTYARRAYINDEX    = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 385
	HRESULT_CTL_E_SETNOTPERMITTED           = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 387
	HRESULT_CTL_E_GETNOTSUPPORTEDATRUNTIME  = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 393
	HRESULT_CTL_E_GETNOTSUPPORTED           = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 394
	HRESULT_CTL_E_PROPERTYNOTFOUND          = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 422
	HRESULT_CTL_E_INVALIDCLIPBOARDFORMAT    = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 460
	HRESULT_CTL_E_INVALIDPICTURE            = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 481
	HRESULT_CTL_E_PRINTERERROR              = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 482
	HRESULT_CTL_E_CANTSAVEFILETOTEMP        = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 735
	HRESULT_CTL_E_SEARCHTEXTNOTFOUND        = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 744
	HRESULT_CTL_E_REPLACEMENTSTOOLONG       = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_CONTROL)<<16 | 746

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

	HRESULT_DV_E_FORMATETC           HRESULT = 0x8004_0064 // Invalid FORMATETC structure.
	HRESULT_DV_E_DVTARGETDEVICE      HRESULT = 0x8004_0065 // Invalid DVTARGETDEVICE structure.
	HRESULT_DV_E_STGMEDIUM           HRESULT = 0x8004_0066 // Invalid STDGMEDIUM structure.
	HRESULT_DV_E_STATDATA            HRESULT = 0x8004_0067 // Invalid STATDATA structure.
	HRESULT_DV_E_LINDEX              HRESULT = 0x8004_0068 // Invalid lindex.
	HRESULT_DV_E_TYMED               HRESULT = 0x8004_0069 // Invalid tymed.
	HRESULT_DV_E_CLIPFORMAT          HRESULT = 0x8004_006a // Invalid clipboard format.
	HRESULT_DV_E_DVASPECT            HRESULT = 0x8004_006b // Invalid aspect(s).
	HRESULT_DV_E_DVTARGETDEVICE_SIZE HRESULT = 0x8004_006c // tdSize parameter of the DVTARGETDEVICE structure is invalid.
	HRESULT_DV_E_NOIVIEWOBJECT       HRESULT = 0x8004_006d // Object doesn't support IViewObject interface.

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

	_HRESULT_PERPROP_E_FIRST          = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_ITF)<<16 | 0x0200
	HRESULT_PERPROP_E_NOPAGEAVAILABLE = _HRESULT_PERPROP_E_FIRST + 0 // No page available for requested property.

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

	_HRESULT_SELFREG_E_FIRST  = HRESULT(SEVERITY_ERROR)<<31 | HRESULT(FACILITY_ITF)<<16 | 0x0200
	HRESULT_SELFREG_E_TYPELIB = _HRESULT_SELFREG_E_FIRST + 0 // Failed to register/unregister type library.
	HRESULT_SELFREG_E_CLASS   = _HRESULT_SELFREG_E_FIRST + 1 // Failed to register/unregister class.

	TYPE_E_BUFFERTOOSMALL          HRESULT = 0x8002_8016 // Buffer too small.
	TYPE_E_FIELDNOTFOUND           HRESULT = 0x8002_8017 // Field name not defined in the record.
	TYPE_E_INVDATAREAD             HRESULT = 0x8002_8018 // Old format or invalid type library.
	TYPE_E_UNSUPFORMAT             HRESULT = 0x8002_8019 // Old format or invalid type library.
	TYPE_E_REGISTRYACCESS          HRESULT = 0x8002_801c // Error accessing the OLE registry.
	TYPE_E_LIBNOTREGISTERED        HRESULT = 0x8002_801d // Library not registered.
	TYPE_E_UNDEFINEDTYPE           HRESULT = 0x8002_8027 // Bound to unknown type.
	TYPE_E_QUALIFIEDNAMEDISALLOWED HRESULT = 0x8002_8028 // Qualified name disallowed.
	TYPE_E_INVALIDSTATE            HRESULT = 0x8002_8029 // Invalid forward reference, or reference to uncompiled type.
	TYPE_E_WRONGTYPEKIND           HRESULT = 0x8002_802a // Type mismatch.
	TYPE_E_ELEMENTNOTFOUND         HRESULT = 0x8002_802b // Element not found.
	TYPE_E_AMBIGUOUSNAME           HRESULT = 0x8002_802c // Ambiguous name.
	TYPE_E_NAMECONFLICT            HRESULT = 0x8002_802d // Name already exists in the library.
	TYPE_E_UNKNOWNLCID             HRESULT = 0x8002_802e // Unknown LCID.
	TYPE_E_DLLFUNCTIONNOTFOUND     HRESULT = 0x8002_802f // Function not defined in specified DLL.
	TYPE_E_BADMODULEKIND           HRESULT = 0x8002_88bd // Wrong module kind for the operation.
	TYPE_E_SIZETOOBIG              HRESULT = 0x8002_88c5 // Size may not exceed 64K.
	TYPE_E_DUPLICATEID             HRESULT = 0x8002_88c6 // Duplicate ID in inheritance hierarchy.
	TYPE_E_INVALIDID               HRESULT = 0x8002_88cf // Incorrect inheritance depth in standard OLE hmember.
	TYPE_E_TYPEMISMATCH            HRESULT = 0x8002_8ca0 // Type mismatch.
	TYPE_E_OUTOFBOUNDS             HRESULT = 0x8002_8ca1 // Invalid number of arguments.
	TYPE_E_IOERROR                 HRESULT = 0x8002_8ca2 // I/O Error.
	TYPE_E_CANTCREATETMPFILE       HRESULT = 0x8002_8ca3 // Error creating unique tmp file.
	TYPE_E_CANTLOADLIBRARY         HRESULT = 0x8002_9c4a // Error loading type library/DLL.
	TYPE_E_INCONSISTENTPROPFUNCS   HRESULT = 0x8002_9c83 // Inconsistent property functions.
	TYPE_E_CIRCULARTYPE            HRESULT = 0x8002_9c84 // Circular dependency between types/modules.

	HRESULT_WINCODEC_ERR_WRONGSTATE                       HRESULT = 0x8898_2f04 // The codec is in the wrong state.
	HRESULT_WINCODEC_ERR_VALUEOUTOFRANGE                  HRESULT = 0x8898_2f05 // The value is out of range.
	HRESULT_WINCODEC_ERR_UNKNOWNIMAGEFORMAT               HRESULT = 0x8898_2f07 // The image format is unknown.
	HRESULT_WINCODEC_ERR_UNSUPPORTEDVERSION               HRESULT = 0x8898_2f0b // The SDK version is unsupported.
	HRESULT_WINCODEC_ERR_NOTINITIALIZED                   HRESULT = 0x8898_2f0c // The component is not initialized.
	HRESULT_WINCODEC_ERR_ALREADYLOCKED                    HRESULT = 0x8898_2f0d // There is already an outstanding read or write lock.
	HRESULT_WINCODEC_ERR_PROPERTYNOTFOUND                 HRESULT = 0x8898_2f40 // The specified bitmap property cannot be found.
	HRESULT_WINCODEC_ERR_PROPERTYNOTSUPPORTED             HRESULT = 0x8898_2f41 // The bitmap codec does not support the bitmap property.
	HRESULT_WINCODEC_ERR_PROPERTYSIZE                     HRESULT = 0x8898_2f42 // The bitmap property size is invalid.
	HRESULT_WINCODEC_ERR_CODECPRESENT                     HRESULT = 0x8898_2f43 // An unknown error has occurred.
	HRESULT_WINCODEC_ERR_CODECNOTHUMBNAIL                 HRESULT = 0x8898_2f44 // The bitmap codec does not support a thumbnail.
	HRESULT_WINCODEC_ERR_PALETTEUNAVAILABLE               HRESULT = 0x8898_2f45 // The bitmap palette is unavailable.
	HRESULT_WINCODEC_ERR_CODECTOOMANYSCANLINES            HRESULT = 0x8898_2f46 // Too many scanlines were requested.
	HRESULT_WINCODEC_ERR_INTERNALERROR                    HRESULT = 0x8898_2f48 // An internal error occurred.
	HRESULT_WINCODEC_ERR_SOURCERECTDOESNOTMATCHDIMENSIONS HRESULT = 0x8898_2f49 // The bitmap bounds do not match the bitmap dimensions.
	HRESULT_WINCODEC_ERR_COMPONENTNOTFOUND                HRESULT = 0x8898_2f50 // The component cannot be found.
	HRESULT_WINCODEC_ERR_IMAGESIZEOUTOFRANGE              HRESULT = 0x8898_2f51 // The bitmap size is outside the valid range.
	HRESULT_WINCODEC_ERR_TOOMUCHMETADATA                  HRESULT = 0x8898_2f52 // There is too much metadata to be written to the bitmap.
	HRESULT_WINCODEC_ERR_BADIMAGE                         HRESULT = 0x8898_2f60 // The image is unrecognized.
	HRESULT_WINCODEC_ERR_BADHEADER                        HRESULT = 0x8898_2f61 // The image header is unrecognized.
	HRESULT_WINCODEC_ERR_FRAMEMISSING                     HRESULT = 0x8898_2f62 // The bitmap frame is missing.
	HRESULT_WINCODEC_ERR_BADMETADATAHEADER                HRESULT = 0x8898_2f63 // The image metadata header is unrecognized.
	HRESULT_WINCODEC_ERR_BADSTREAMDATA                    HRESULT = 0x8898_2f70 // The stream data is unrecognized.
	HRESULT_WINCODEC_ERR_STREAMWRITE                      HRESULT = 0x8898_2f71 // Failed to write to the stream.
	HRESULT_WINCODEC_ERR_STREAMREAD                       HRESULT = 0x8898_2f72 // Failed to read from the stream.
	HRESULT_WINCODEC_ERR_STREAMNOTAVAILABLE               HRESULT = 0x8898_2f73 // The stream is not available.
	HRESULT_WINCODEC_ERR_UNSUPPORTEDPIXELFORMAT           HRESULT = 0x8898_2f80 // The bitmap pixel format is unsupported.
	HRESULT_WINCODEC_ERR_UNSUPPORTEDOPERATION             HRESULT = 0x8898_2f81 // The operation is unsupported.
	HRESULT_WINCODEC_ERR_INVALIDREGISTRATION              HRESULT = 0x8898_2f8a // The component registration is invalid.
	HRESULT_WINCODEC_ERR_COMPONENTINITIALIZEFAILURE       HRESULT = 0x8898_2f8b // The component initialization has failed.
	HRESULT_WINCODEC_ERR_INSUFFICIENTBUFFER               HRESULT = 0x8898_2f8c // The buffer allocated is insufficient.
	HRESULT_WINCODEC_ERR_DUPLICATEMETADATAPRESENT         HRESULT = 0x8898_2f8d // Duplicate metadata is present.
	HRESULT_WINCODEC_ERR_PROPERTYUNEXPECTEDTYPE           HRESULT = 0x8898_2f8e // The bitmap property type is unexpected.
	HRESULT_WINCODEC_ERR_UNEXPECTEDSIZE                   HRESULT = 0x8898_2f8f // The size is unexpected.
	HRESULT_WINCODEC_ERR_INVALIDQUERYREQUEST              HRESULT = 0x8898_2f90 // The property query is invalid.
	HRESULT_WINCODEC_ERR_UNEXPECTEDMETADATATYPE           HRESULT = 0x8898_2f91 // The metadata type is unexpected.
	HRESULT_WINCODEC_ERR_REQUESTONLYVALIDATMETADATAROOT   HRESULT = 0x8898_2f92 // The specified bitmap property is only valid at root level.
	HRESULT_WINCODEC_ERR_INVALIDQUERYCHARACTER            HRESULT = 0x8898_2f93 // The query string contains an invalid character.
	HRESULT_WINCODEC_ERR_WIN32ERROR                       HRESULT = 0x8898_2f94 // Windows Codecs received an error from the Win32 system.
	HRESULT_WINCODEC_ERR_INVALIDPROGRESSIVELEVEL          HRESULT = 0x8898_2f95 // The requested level of detail is not present.
	HRESULT_WINCODEC_ERR_INVALIDJPEGSCANINDEX             HRESULT = 0x8898_2f96 // The scan index is invalid.
	HRESULT_WINCODEC_ERR_UNSUPPORTEDTONEMAPPING           HRESULT = 0x8898_2f97 // The tone mapping mode is not supported.
)
