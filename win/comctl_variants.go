package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

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
