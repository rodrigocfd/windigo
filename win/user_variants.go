package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

type (
	// Variant type for a class name identifier.
	//
	// Example:
	//
	//  func Foo(cn win.ClassName) {
	//      // ...
	//  }
	//
	//  Foo(win.ClassNameAtom(ATOM(100)))
	//  Foo(win.ClassNameStr("MY_CLASS_NAME"))
	//  Foo(win.ClassNameNone{})
	ClassName     interface{ implClassName() }
	ClassNameAtom ATOM     // ClassName variant: ATOM.
	ClassNameStr  string   // ClassName variant: string.
	ClassNameNone struct{} // ClassName variant: no value.
)

func (ClassNameAtom) implClassName() {}
func (ClassNameStr) implClassName()  {}
func (ClassNameNone) implClassName() {}

func variantClassName(v ClassName) (uintptr, *uint16) { // pointer must be kept alive
	switch v := v.(type) {
	case ClassNameAtom:
		return uintptr(v), nil
	case ClassNameStr:
		buf := Str.ToNativePtr(string(v))
		return uintptr(unsafe.Pointer(buf)), buf
	case ClassNameNone:
		return 0, nil
	default:
		panic("ClassName cannot be nil.")
	}
}

type (
	// Variant type for a resource identifier.
	//
	// Example:
	//
	//  func Foo(c win.CursorRes) {
	//      // ...
	//  }
	//
	//  Foo(win.CursorResIdc(co.IDC_ARROW))
	//  Foo(win.CursorResInt(301))
	//  Foo(win.CursorResStr("MY_CURSOR"))
	CursorRes    interface{ implCursorRes() }
	CursorResIdc co.IDC // CursorRes variant: co.IDC.
	CursorResInt uint16 // CursorRes variant: uint16.
	CursorResStr string // CursorRes variant: string.
)

func (CursorResIdc) implCursorRes() {}
func (CursorResInt) implCursorRes() {}
func (CursorResStr) implCursorRes() {}

func variantCursorResId(v CursorRes) (uintptr, *uint16) { // pointer must be kept alive
	switch v := v.(type) {
	case CursorResIdc:
		return uintptr(v), nil
	case CursorResInt:
		return uintptr(v), nil
	case CursorResStr:
		buf := Str.ToNativePtr(string(v))
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("CursorResId cannot be nil.")
	}
}

type (
	// Variant type for an icon resource identifier.
	//
	// Example:
	//
	//  func Foo(i win.IconRes) {
	//      // ...
	//  }
	//
	//  Foo(win.IconResIdc(co.IDI_ERROR))
	//  Foo(win.IconResInt(201))
	//  Foo(win.IconResStr("MY_ICON"))
	IconRes    interface{ implIconRes() }
	IconResIdc co.IDI // IconRes variant: co.IDI.
	IconResInt uint16 // IconRes variant: uint16.
	IconResStr string // IconRes variant: string.
)

func (IconResIdc) implIconRes() {}
func (IconResInt) implIconRes() {}
func (IconResStr) implIconRes() {}

func variantIconResId(v IconRes) (uintptr, *uint16) { // pointer must be kept alive
	switch v := v.(type) {
	case IconResIdc:
		return uintptr(v), nil
	case IconResInt:
		return uintptr(v), nil
	case IconResStr:
		buf := Str.ToNativePtr(string(v))
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("CursorResId cannot be nil.")
	}
}

type (
	// Variant type for a menu item identifier.
	//
	// Example:
	//
	//  func Foo(i win.MenuItem) {
	//      // ...
	//  }
	//
	//  Foo(win.MenuItemCmd(4001))
	//  Foo(win.MenuItemPos(2))
	MenuItem    interface{ implMenuItem() }
	MenuItemCmd uint16 // MenuItem variant: uint16 command ID.
	MenuItemPos uint16 // MenuItem variant: uint16 zero-based item index.
)

func (MenuItemCmd) implMenuItem() {}
func (MenuItemPos) implMenuItem() {}

func variantMenuItem(v MenuItem) (uintptr, co.MF) {
	switch v := v.(type) {
	case MenuItemCmd:
		return uintptr(v), co.MF_BYCOMMAND
	case MenuItemPos:
		return uintptr(v), co.MF_BYPOSITION
	default:
		panic("MenuItem cannot be nil.")
	}
}

type (
	// Variant type for a resource identifier.
	//
	// Example:
	//
	//  func Foo(r win.ResId) {
	//      // ...
	//  }
	//
	//  Foo(win.ResIdInt(101))
	//  Foo(win.ResIdStr("MY_RESOURCE"))
	ResId    interface{ implResId() }
	ResIdInt uint16 // ResId variant: uint16.
	ResIdStr string // ResId variant: string.
)

func (ResIdInt) implResId() {}
func (ResIdStr) implResId() {}

func variantResId(v ResId) (uintptr, *uint16) { // pointer must be kept alive
	switch v := v.(type) {
	case ResIdInt:
		return uintptr(v), nil
	case ResIdStr:
		buf := Str.ToNativePtr(string(v))
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("ResId cannot be nil.")
	}
}
