package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

type (
	// An interface which accepts a string or a nil value.
	//
	// Example:
	//
	//  func Foo(s win.StrOrNil) {}
	//
	//  Foo(win.StrVal("some text"))
	//  Foo(nil)
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
	// Example:
	//
	//  func Foo(cn win.ClassName) {}
	//
	//  Foo(ClassNameAtom(ATOM(100)))
	//  Foo(ClassNameStr("MY_CLASS_NAME"))
	//  Foo(nil)
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
	// Example:
	//
	//  func Foo(c win.CursorRes) {}
	//
	//  Foo(win.CursorResIdc(co.IDC_ARROW))
	//  Foo(win.CursorResInt(301))
	//  Foo(win.CursorResStr("MY_CURSOR"))
	CursorRes    interface{ isCursorRes() }
	CursorResIdc co.IDC // A co.IDC cursor resource identifier for CursorResId interface.
	CursorResInt uint16 // A number cursor resource identifier for CursorResId interface.
	CursorResStr string // A string cursor resource identifier for CursorResId interface.
)

func (CursorResIdc) isCursorRes() {}
func (CursorResInt) isCursorRes() {}
func (CursorResStr) isCursorRes() {}

func variantCursorResId(v CursorRes) unsafe.Pointer {
	switch v := v.(type) {
	case CursorResIdc:
		return unsafe.Pointer(uintptr(v))
	case CursorResInt:
		return unsafe.Pointer(uintptr(v))
	case CursorResStr:
		return unsafe.Pointer(Str.ToNativePtr(string(v)))
	default:
		panic("CursorResId does not accept a nil value.")
	}
}

//------------------------------------------------------------------------------

type (
	// An icon resource identifier.
	//
	// Example:
	//
	//  func Foo(i win.IconRes) {}
	//
	//  Foo(win.IconResIdc(co.IDI_ERROR))
	//  Foo(win.IconResInt(201))
	//  Foo(win.IconResStr("MY_ICON"))
	IconRes    interface{ isIconRes() }
	IconResIdc co.IDI // A co.IDI icon resource identifier for IconResId interface.
	IconResInt uint16 // A number icon resource identifier for IconResId interface.
	IconResStr string // A string icon resource identifier for IconResId interface.
)

func (IconResIdc) isIconRes() {}
func (IconResInt) isIconRes() {}
func (IconResStr) isIconRes() {}

func variantIconResId(v IconRes) unsafe.Pointer {
	switch v := v.(type) {
	case IconResIdc:
		return unsafe.Pointer(uintptr(v))
	case IconResInt:
		return unsafe.Pointer(uintptr(v))
	case IconResStr:
		return unsafe.Pointer(Str.ToNativePtr(string(v)))
	default:
		panic("CursorResId does not accept a nil value.")
	}
}

//------------------------------------------------------------------------------

type (
	// A menu item identifier.
	//
	// Example:
	//
	//  func Foo(i win.MenuItem) {}
	//
	//  Foo(win.MenuItemCmd(4001))
	//  Foo(win.MenuItemPos(2))
	MenuItem    interface{ isIdPos() }
	MenuItemCmd uint16 // A command ID for MenuItem interface.
	MenuItemPos uint16 // A zero-based index for MenuItem interface.
)

func (MenuItemCmd) isIdPos() {}
func (MenuItemPos) isIdPos() {}

func variantMenuItem(v MenuItem) (uintptr, co.MF) {
	switch v := v.(type) {
	case MenuItemCmd:
		return uintptr(v), co.MF_BYCOMMAND
	case MenuItemPos:
		return uintptr(v), co.MF_BYPOSITION
	default:
		panic("MenuItem does not accept a nil value.")
	}
}

//------------------------------------------------------------------------------

type (
	// A resource identifier.
	//
	// Example:
	//
	//  func Foo(r win.ResId) {}
	//
	//  Foo(win.ResIdInt(101))
	//  Foo(win.ResIdStr("MY_RESOURCE"))
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
	// Example:
	//
	//  func Foo(i win.TdcIcon) {}
	//
	//  Foo(win.TdcIconHicon(win.HICON(0)))
	//  Foo(win.TdcIconInt(301))
	//  Foo(win.TdcIconTdi(co.TD_ICON_ERROR))
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
