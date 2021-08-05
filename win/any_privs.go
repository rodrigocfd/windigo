package win

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// Private constants.
const (
	_CCHILDREN_TITLEBAR   = 5
	_CLR_INVALID          = 0xffff_ffff
	_GDI_ERR              = 0xffff_ffff
	_HGDI_ERROR           = 0xffff_ffff
	_INVALID_HANDLE_VALUE = -1
	_L_MAX_URL_LENGTH     = 2048 + 32 + 4
	_LF_FACESIZE          = 32
	_MAX_LINKID_TEXT      = 48
	_MAX_PATH             = 260
	_UINT_MAX             = 4294967295
)

// Converts interface{}, from a restrict set of types, to uintptr.
// Don't forget to call runtime.KeepAlive() on the original variable.
type _UintptrConv struct {
	val interface{}
}

func (me *_UintptrConv) hbitmapLparamString() uintptr {
	var raw uintptr

	switch v := me.val.(type) {
	case HBITMAP:
		raw = uintptr(v)
	case LPARAM:
		raw = uintptr(v)
	case string:
		raw = uintptr(unsafe.Pointer(Str.ToUint16Ptr(v)))
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(me.val)))
	}

	return raw
}

func (me *_UintptrConv) uint16Hmenu() uintptr {
	var raw uintptr

	switch v := me.val.(type) {
	case uint16:
		raw = uintptr(v)
	case HMENU:
		raw = uintptr(v)
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(me.val)))
	}

	return raw
}

func (me *_UintptrConv) uint16String() uintptr {
	var raw uintptr

	switch v := me.val.(type) {
	case uint16:
		raw = uintptr(v)
	case string:
		raw = uintptr(unsafe.Pointer(Str.ToUint16Ptr(v)))
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(me.val)))
	}

	return raw
}

func (me *_UintptrConv) uint16IdcString() uintptr {
	var raw uintptr

	switch v := me.val.(type) {
	case uint16:
		raw = uintptr(v)
	case co.IDC:
		raw = uintptr(v)
	case string:
		raw = uintptr(unsafe.Pointer(Str.ToUint16Ptr(v)))
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(me.val)))
	}

	return raw
}

func (me *_UintptrConv) uint16IdiString() uintptr {
	var raw uintptr

	switch v := me.val.(type) {
	case uint16:
		raw = uintptr(v)
	case co.IDI:
		raw = uintptr(v)
	case string:
		raw = uintptr(unsafe.Pointer(Str.ToUint16Ptr(v)))
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(me.val)))
	}

	return raw
}
