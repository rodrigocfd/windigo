package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

type (
	// Variant type which accepts string or a nil.
	//
	// Example:
	//
	//  func Foo(s win.StrOrNil) {}
	//
	//  Foo(win.StrVal("some text"))
	//  Foo(nil)
	StrOrNil interface{ isStrOrNil() }
	StrVal   string // StrOrNil variant: string.
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
	// Variant type for a class name identifier.
	//
	// Example:
	//
	//  func Foo(cn win.ClassName) {}
	//
	//  Foo(ClassNameAtom(ATOM(100)))
	//  Foo(ClassNameStr("MY_CLASS_NAME"))
	//  Foo(nil)
	ClassName     interface{ isClassName() }
	ClassNameAtom ATOM   // ClassName variant: ATOM.
	ClassNameStr  string // ClassName variant: string.
)

func (ClassNameAtom) isClassName() {}
func (ClassNameStr) isClassName()  {}

func variantClassName(v ClassName) (uintptr, *uint16) { // pointer must be kept alive
	var buf *uint16
	switch v := v.(type) {
	case ClassNameAtom:
		return uintptr(v), nil
	case ClassNameStr:
		buf = Str.ToNativePtr(string(v))
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		return 0, nil
	}
}

//------------------------------------------------------------------------------

type (
	// Variant type for a resource identifier.
	//
	// Example:
	//
	//  func Foo(c win.CursorRes) {}
	//
	//  Foo(win.CursorResIdc(co.IDC_ARROW))
	//  Foo(win.CursorResInt(301))
	//  Foo(win.CursorResStr("MY_CURSOR"))
	CursorRes    interface{ isCursorRes() }
	CursorResIdc co.IDC // CursorRes variant: co.IDC.
	CursorResInt uint16 // CursorRes variant: uint16.
	CursorResStr string // CursorRes variant: string.
)

func (CursorResIdc) isCursorRes() {}
func (CursorResInt) isCursorRes() {}
func (CursorResStr) isCursorRes() {}

func variantCursorResId(v CursorRes) (uintptr, *uint16) { // pointer must be kept alive
	var buf *uint16
	switch v := v.(type) {
	case CursorResIdc:
		return uintptr(v), nil
	case CursorResInt:
		return uintptr(v), nil
	case CursorResStr:
		buf = Str.ToNativePtr(string(v))
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("CursorResId does not accept a nil value.")
	}
}

//------------------------------------------------------------------------------

type (
	// Variant type for an icon resource identifier.
	//
	// Example:
	//
	//  func Foo(i win.IconRes) {}
	//
	//  Foo(win.IconResIdc(co.IDI_ERROR))
	//  Foo(win.IconResInt(201))
	//  Foo(win.IconResStr("MY_ICON"))
	IconRes    interface{ isIconRes() }
	IconResIdc co.IDI // IconRes variant: co.IDI.
	IconResInt uint16 // IconRes variant: uint16.
	IconResStr string // IconRes variant: string.
)

func (IconResIdc) isIconRes() {}
func (IconResInt) isIconRes() {}
func (IconResStr) isIconRes() {}

func variantIconResId(v IconRes) (uintptr, *uint16) { // pointer must be kept alive
	var buf *uint16
	switch v := v.(type) {
	case IconResIdc:
		return uintptr(v), nil
	case IconResInt:
		return uintptr(v), nil
	case IconResStr:
		buf = Str.ToNativePtr(string(v))
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("CursorResId does not accept a nil value.")
	}
}

//------------------------------------------------------------------------------

type (
	// Variant type for a menu item identifier.
	//
	// Example:
	//
	//  func Foo(i win.MenuItem) {}
	//
	//  Foo(win.MenuItemCmd(4001))
	//  Foo(win.MenuItemPos(2))
	MenuItem    interface{ isIdPos() }
	MenuItemCmd uint16 // MenuItem variant: uint16 command ID.
	MenuItemPos uint16 // MenuItem variant: uint16 zero-based item index.
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
	// Variant type for a resource identifier.
	//
	// Example:
	//
	//  func Foo(r win.ResId) {}
	//
	//  Foo(win.ResIdInt(101))
	//  Foo(win.ResIdStr("MY_RESOURCE"))
	ResId    interface{ isResId() }
	ResIdInt uint16 // ResId variant: uint16.
	ResIdStr string // ResId variant: string.
)

func (ResIdInt) isResId() {}
func (ResIdStr) isResId() {}

func variantResId(v ResId) (uintptr, *uint16) { // pointer must be kept alive
	var buf *uint16
	switch v := v.(type) {
	case ResIdInt:
		return uintptr(v), nil
	case ResIdStr:
		buf = Str.ToNativePtr(string(v))
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("ResId does not accept a nil value.")
	}
}

//------------------------------------------------------------------------------

type (
	// Variant type for an icon identifier for TASKDIALOGCONFIG.
	//
	// Example:
	//
	//  func Foo(i win.TdcIcon) {}
	//
	//  Foo(win.TdcIconHicon(win.HICON(0)))
	//  Foo(win.TdcIconInt(301))
	//  Foo(win.TdcIconTdi(co.TD_ICON_ERROR))
	TdcIcon      interface{ isTdcIcon() }
	TdcIconHicon HICON      // TdcIcon variant: HICON.
	TdcIconInt   uint16     // TdcIcon variant: uint16.
	TdcIconTdi   co.TD_ICON // TdcIcon variant: co.TD_ICON.
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
