package win

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Private constants.
const (
	_CCHILDREN_TITLEBAR   = 5
	_CLR_INVALID          = 0xffff_ffff
	_GDI_ERR              = 0xffff_ffff
	_GMEM_INVALID_HANDLE  = 0x8000
	_HGDI_ERROR           = 0xffff_ffff
	_INVALID_HANDLE_VALUE = -1
	_L_MAX_URL_LENGTH     = 2048 + 32 + 4
	_LF_FACESIZE          = 32
	_MAX_LINKID_TEXT      = 48
	_MAX_PATH             = 260
	_UINT_MAX             = 4294967295
)

// Converts val to uint16 or string, or panics.
func _PullUint16String(val interface{}) uintptr {
	switch v := val.(type) {
	case uint16:
		return uintptr(v)
	case string:
		return uintptr(unsafe.Pointer(Str.ToUint16Ptr(v))) // runtime.KeepAlive()
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(val)))
	}
}
