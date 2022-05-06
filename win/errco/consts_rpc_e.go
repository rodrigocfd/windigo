package errco

const (
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