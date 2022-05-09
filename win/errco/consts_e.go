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
)
