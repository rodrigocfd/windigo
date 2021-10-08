package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

type (
	// An interface which accepts a string or a nil value.
	//
	// You can pass StrVal("text") or nil.
	StrOrNil interface{ isStrOrNil() }
	StrVal   string // A string value for StrOrNil interface.
)

func (StrVal) isStrOrNil() {}

func variantStrOrNil(v StrOrNil) unsafe.Pointer {
	if v != nil {
		s := v.(StrVal)
		return unsafe.Pointer(Str.ToNativePtr(string(s)))
	}
	return nil
}

//------------------------------------------------------------------------------

type (
	// A window class name identifier.
	//
	// You can pass ClassNameAtom(atom), ClassNameStr("NAME") or nil.
	ClassName     interface{ isClassName() }
	ClassNameAtom ATOM   // An atom window class name identifier for ClassName interface.
	ClassNameStr  string // A string window class name identifier for ClassName interface.
)

func (ClassNameAtom) isClassName() {}
func (ClassNameStr) isClassName()  {}

func variantClassName(v ClassName) unsafe.Pointer {
	switch v := v.(type) {
	case ClassNameAtom:
		return unsafe.Pointer(uintptr(v))
	case ClassNameStr:
		return unsafe.Pointer(Str.ToNativePtr(string(v)))
	default:
		return nil
	}
}

//------------------------------------------------------------------------------

type (
	// A cursor resource identifier.
	//
	CursorResId    interface{ isCursorResId() }
	CursorResIdIdc co.IDC // A co.IDC cursor resource identifier for CursorResId interface.
	CursorResIdInt uint16 // A number cursor resource identifier for CursorResId interface.
	CursorResIdStr string // A string cursor resource identifier for CursorResId interface.
)

func (CursorResIdIdc) isCursorResId() {}
func (CursorResIdInt) isCursorResId() {}
func (CursorResIdStr) isCursorResId() {}

func variantCursorResId(v CursorResId) unsafe.Pointer {
	switch v := v.(type) {
	case CursorResIdIdc:
		return unsafe.Pointer(uintptr(v))
	case CursorResIdInt:
		return unsafe.Pointer(uintptr(v))
	case CursorResIdStr:
		return unsafe.Pointer(Str.ToNativePtr(string(v)))
	default:
		panic("CursorResId does not accept a nil value.")
	}
}

//------------------------------------------------------------------------------

type (
	// An icon resource identifier.
	//
	IconResId    interface{ isIconResId() }
	IconResIdIdc co.IDI // A co.IDI icon resource identifier for IconResId interface.
	IconResIdInt uint16 // A number icon resource identifier for IconResId interface.
	IconResIdStr string // A string icon resource identifier for IconResId interface.
)

func (IconResIdIdc) isIconResId() {}
func (IconResIdInt) isIconResId() {}
func (IconResIdStr) isIconResId() {}

func variantIconResId(v IconResId) unsafe.Pointer {
	switch v := v.(type) {
	case IconResIdIdc:
		return unsafe.Pointer(uintptr(v))
	case IconResIdInt:
		return unsafe.Pointer(uintptr(v))
	case IconResIdStr:
		return unsafe.Pointer(Str.ToNativePtr(string(v)))
	default:
		panic("CursorResId does not accept a nil value.")
	}
}

//------------------------------------------------------------------------------

type (
	// A resource identifier.
	//
	// You can pass ResIdStr("ABC") or ResIdInt(100).
	ResId    interface{ isResId() }
	ResIdInt uint16 // A number resource identifier for ResId interface.
	ResIdStr string // A string resource identifier for ResId interface.
)

func (ResIdInt) isResId() {}
func (ResIdStr) isResId() {}

func variantResId(v ResId) unsafe.Pointer {
	switch v := v.(type) {
	case ResIdInt:
		return unsafe.Pointer(uintptr(v))
	case ResIdStr:
		return unsafe.Pointer(Str.ToNativePtr(string(v)))
	default:
		panic("ResId does not accept a nil value.")
	}
}

//------------------------------------------------------------------------------

type (
	// Icon identifier for TASKDIALOGCONFIG.
	//
	// You can pass TdcIconHicon(hicon), TdcIconInt(100),
	// TdcIconTdi(co.TD_ICON_ERROR) or nil.
	TdcIcon      interface{ isTdcIcon() }
	TdcIconHicon HICON      // An HICON identifier for TdcIcon interface.
	TdcIconInt   uint16     // A number identifier for TdcIcon interface.
	TdcIconTdi   co.TD_ICON // A co.TD_ICON identifier for TdcIcon interface.
)

func (TdcIconHicon) isTdcIcon() {}
func (TdcIconInt) isTdcIcon()   {}
func (TdcIconTdi) isTdcIcon()   {}

func variantTdcIcon(v TdcIcon) uintptr {
	switch v := v.(type) {
	case TdcIconHicon:
		return uintptr(v)
	case TdcIconInt:
		return uintptr(v)
	case TdcIconTdi:
		return uintptr(v)
	default:
		return 0
	}
}
