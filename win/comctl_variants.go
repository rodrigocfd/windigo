package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

type (
	// Variant type for an icon identifier for TASKDIALOGCONFIG.
	//
	// Example:
	//
	//  func Foo(i win.TdcIcon) {
	//      // ...
	//  }
	//
	//  Foo(win.TdcIconHicon(win.HICON(0)))
	//  Foo(win.TdcIconInt(301))
	//  Foo(win.TdcIconTdi(co.TD_ICON_ERROR))
	//  Foo(win.TdcIconNone{})
	TdcIcon      interface{ implTdcIcon() }
	TdcIconHicon HICON      // TdcIcon variant: HICON.
	TdcIconInt   uint16     // TdcIcon variant: uint16.
	TdcIconTdi   co.TD_ICON // TdcIcon variant: co.TD_ICON.
	TdcIconNone  struct{}   // TdcIcon variant: no value.
)

func (TdcIconHicon) implTdcIcon() {}
func (TdcIconInt) implTdcIcon()   {}
func (TdcIconTdi) implTdcIcon()   {}
func (TdcIconNone) implTdcIcon()  {}

func variantTdcIcon(v TdcIcon) uintptr {
	switch v := v.(type) {
	case TdcIconHicon:
		return uintptr(v)
	case TdcIconInt:
		return uintptr(v)
	case TdcIconTdi:
		return uintptr(v)
	case TdcIconNone:
		return 0
	default:
		panic("TdcIconNone cannot be nil.")
	}
}
