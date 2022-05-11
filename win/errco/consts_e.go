//go:build windows

package errco

const (
	E_UNEXPECTED           ERROR = 0x8000_ffff
	E_NOTIMPL              ERROR = 0x8000_4001
	E_OUTOFMEMORY          ERROR = 0x8007_000e
	E_INVALIDARG           ERROR = 0x8007_0057
	E_NOINTERFACE          ERROR = 0x8000_4002
	E_POINTER              ERROR = 0x8000_4003
	E_HANDLE               ERROR = 0x8007_0006
	E_ABORT                ERROR = 0x8000_4004
	E_FAIL                 ERROR = 0x8000_4005
	E_ACCESSDENIED         ERROR = 0x8007_0005
	E_PENDING              ERROR = 0x8000_000a
	E_BOUNDS               ERROR = 0x8000_000b
	E_CHANGED_STATE        ERROR = 0x8000_000c
	E_ILLEGAL_STATE_CHANGE ERROR = 0x8000_000d
	E_ILLEGAL_METHOD_CALL  ERROR = 0x8000_000e

	CO_E_NOTINITIALIZED     ERROR = 0x8004_01f0
	CO_E_ALREADYINITIALIZED ERROR = 0x8004_01f1
	CO_E_CANTDETERMINECLASS ERROR = 0x8004_01f2
	CO_E_CLASSSTRING        ERROR = 0x8004_01f3
	CO_E_IIDSTRING          ERROR = 0x8004_01f4
	CO_E_APPNOTFOUND        ERROR = 0x8004_01f5
	CO_E_APPSINGLEUSE       ERROR = 0x8004_01f6
	CO_E_ERRORINAPP         ERROR = 0x8004_01f7
	CO_E_DLLNOTFOUND        ERROR = 0x8004_01f8
	CO_E_ERRORINDLL         ERROR = 0x8004_01f9
	CO_E_WRONGOSFORAPP      ERROR = 0x8004_01fa
	CO_E_OBJNOTREG          ERROR = 0x8004_01fb
	CO_E_OBJISREG           ERROR = 0x8004_01fc
	CO_E_OBJNOTCONNECTED    ERROR = 0x8004_01fd
	CO_E_APPDIDNTREG        ERROR = 0x8004_01fe
	CO_E_RELEASED           ERROR = 0x8004_01ff

	DISP_E_UNKNOWNINTERFACE ERROR = 0x8002_0001
	DISP_E_MEMBERNOTFOUND   ERROR = 0x8002_0003
	DISP_E_PARAMNOTFOUND    ERROR = 0x8002_0004
	DISP_E_TYPEMISMATCH     ERROR = 0x8002_0005
	DISP_E_UNKNOWNNAME      ERROR = 0x8002_0006
	DISP_E_NONAMEDARGS      ERROR = 0x8002_0007
	DISP_E_BADVARTYPE       ERROR = 0x8002_0008
	DISP_E_EXCEPTION        ERROR = 0x8002_0009
	DISP_E_OVERFLOW         ERROR = 0x8002_000a
	DISP_E_BADINDEX         ERROR = 0x8002_000b
	DISP_E_UNKNOWNLCID      ERROR = 0x8002_000c
	DISP_E_ARRAYISLOCKED    ERROR = 0x8002_000d
	DISP_E_BADPARAMCOUNT    ERROR = 0x8002_000e
	DISP_E_PARAMNOTOPTIONAL ERROR = 0x8002_000f
	DISP_E_BADCALLEE        ERROR = 0x8002_0010
	DISP_E_NOTACOLLECTION   ERROR = 0x8002_0011
	DISP_E_DIVBYZERO        ERROR = 0x8002_0012
	DISP_E_BUFFERTOOSMALL   ERROR = 0x8002_0013

	DRAGDROP_E_NOTREGISTERED             ERROR = 0x8004_0100
	DRAGDROP_E_ALREADYREGISTERED         ERROR = 0x8004_0101
	DRAGDROP_E_INVALIDHWND               ERROR = 0x8004_0102
	DRAGDROP_E_CONCURRENT_DRAG_ATTEMPTED ERROR = 0x8004_0103

	DV_E_FORMATETC           ERROR = 0x8004_0064
	DV_E_DVTARGETDEVICE      ERROR = 0x8004_0065
	DV_E_STGMEDIUM           ERROR = 0x8004_0066
	DV_E_STATDATA            ERROR = 0x8004_0067
	DV_E_LINDEX              ERROR = 0x8004_0068
	DV_E_TYMED               ERROR = 0x8004_0069
	DV_E_CLIPFORMAT          ERROR = 0x8004_006a
	DV_E_DVASPECT            ERROR = 0x8004_006b
	DV_E_DVTARGETDEVICE_SIZE ERROR = 0x8004_006c

	OLE_E_OLEVERB             ERROR = 0x8004_0000
	OLE_E_ADVF                ERROR = 0x8004_0001
	OLE_E_ENUM_NOMORE         ERROR = 0x8004_0002
	OLE_E_ADVISENOTSUPPORTED  ERROR = 0x8004_0003
	OLE_E_NOCONNECTION        ERROR = 0x8004_0004
	OLE_E_NOTRUNNING          ERROR = 0x8004_0005
	OLE_E_NOCACHE             ERROR = 0x8004_0006
	OLE_E_BLANK               ERROR = 0x8004_0007
	OLE_E_CLASSDIFF           ERROR = 0x8004_0008
	OLE_E_CANT_GETMONIKER     ERROR = 0x8004_0009
	OLE_E_CANT_BINDTOSOURCE   ERROR = 0x8004_000a
	OLE_E_STATIC              ERROR = 0x8004_000b
	OLE_E_PROMPTSAVECANCELLED ERROR = 0x8004_000c
	OLE_E_INVALIDRECT         ERROR = 0x8004_000d
	OLE_E_WRONGCOMPOBJ        ERROR = 0x8004_000e
	OLE_E_INVALIDHWND         ERROR = 0x8004_000f
	OLE_E_NOT_INPLACEACTIVE   ERROR = 0x8004_0010
	OLE_E_CANTCONVERT         ERROR = 0x8004_0011
	OLE_E_NOSTORAGE           ERROR = 0x8004_0012

	RPC_E_CALL_REJECTED               ERROR = 0x8001_0001
	RPC_E_CALL_CANCELED               ERROR = 0x8001_0002
	RPC_E_CANTPOST_INSENDCALL         ERROR = 0x8001_0003
	RPC_E_CANTCALLOUT_INASYNCCALL     ERROR = 0x8001_0004
	RPC_E_CANTCALLOUT_INEXTERNALCALL  ERROR = 0x8001_0005
	RPC_E_CONNECTION_TERMINATED       ERROR = 0x8001_0006
	RPC_E_SERVER_DIED                 ERROR = 0x8001_0007
	RPC_E_CLIENT_DIED                 ERROR = 0x8001_0008
	RPC_E_INVALID_DATAPACKET          ERROR = 0x8001_0009
	RPC_E_CANTTRANSMIT_CALL           ERROR = 0x8001_000a
	RPC_E_CLIENT_CANTMARSHAL_DATA     ERROR = 0x8001_000b
	RPC_E_CLIENT_CANTUNMARSHAL_DATA   ERROR = 0x8001_000c
	RPC_E_SERVER_CANTMARSHAL_DATA     ERROR = 0x8001_000d
	RPC_E_SERVER_CANTUNMARSHAL_DATA   ERROR = 0x8001_000e
	RPC_E_INVALID_DATA                ERROR = 0x8001_000f
	RPC_E_INVALID_PARAMETER           ERROR = 0x8001_0010
	RPC_E_CANTCALLOUT_AGAIN           ERROR = 0x8001_0011
	RPC_E_SERVER_DIED_DNE             ERROR = 0x8001_0012
	RPC_E_SYS_CALL_FAILED             ERROR = 0x8001_0100
	RPC_E_OUT_OF_RESOURCES            ERROR = 0x8001_0101
	RPC_E_ATTEMPTED_MULTITHREAD       ERROR = 0x8001_0102
	RPC_E_NOT_REGISTERED              ERROR = 0x8001_0103
	RPC_E_FAULT                       ERROR = 0x8001_0104
	RPC_E_SERVERFAULT                 ERROR = 0x8001_0105
	RPC_E_CHANGED_MODE                ERROR = 0x8001_0106
	RPC_E_INVALIDMETHOD               ERROR = 0x8001_0107
	RPC_E_DISCONNECTED                ERROR = 0x8001_0108
	RPC_E_RETRY                       ERROR = 0x8001_0109
	RPC_E_SERVERCALL_RETRYLATER       ERROR = 0x8001_010a
	RPC_E_SERVERCALL_REJECTED         ERROR = 0x8001_010b
	RPC_E_INVALID_CALLDATA            ERROR = 0x8001_010c
	RPC_E_CANTCALLOUT_ININPUTSYNCCALL ERROR = 0x8001_010d
	RPC_E_WRONG_THREAD                ERROR = 0x8001_010e
	RPC_E_THREAD_NOT_INIT             ERROR = 0x8001_010f
	RPC_E_VERSION_MISMATCH            ERROR = 0x8001_0110
	RPC_E_INVALID_HEADER              ERROR = 0x8001_0111
	RPC_E_INVALID_EXTENSION           ERROR = 0x8001_0112
	RPC_E_INVALID_IPID                ERROR = 0x8001_0113
	RPC_E_INVALID_OBJECT              ERROR = 0x8001_0114
	RPC_E_CALL_COMPLETE               ERROR = 0x8001_0117
	RPC_E_UNSECURE_CALL               ERROR = 0x8001_0118
	RPC_E_TOO_LATE                    ERROR = 0x8001_0119
	RPC_E_NO_GOOD_SECURITY_PACKAGES   ERROR = 0x8001_011a
	RPC_E_ACCESS_DENIED               ERROR = 0x8001_011b
	RPC_E_REMOTE_DISABLED             ERROR = 0x8001_011c
	RPC_E_INVALID_OBJREF              ERROR = 0x8001_011d
	RPC_E_NO_CONTEXT                  ERROR = 0x8001_011e
	RPC_E_TIMEOUT                     ERROR = 0x8001_011f
	RPC_E_NO_SYNC                     ERROR = 0x8001_0120
	RPC_E_FULLSIC_REQUIRED            ERROR = 0x8001_0121
	RPC_E_INVALID_STD_NAME            ERROR = 0x8001_0122
	RPC_E_UNEXPECTED                  ERROR = 0x8001_ffff
)
