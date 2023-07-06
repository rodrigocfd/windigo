//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// Variant type for an icon identifier for TASKDIALOGCONFIG.
//
// # Example
//
//	ico := TdcIconTdi(co.TD_ICON_ERROR)
//
//	if tdi, ok := ico.Tdi(); ok {
//		println(tdi)
//	}
type TdcIcon struct {
	curType uint8      // 0: none
	hIcon   HICON      // 1
	id      uint16     // 2
	tdIcon  co.TD_ICON // 3
}

// Creates a new TdcIcon variant with an empty value.
func TdcIconNone() TdcIcon {
	return TdcIcon{}
}

// Creates a new TdcIcon variant with an HICON value.
func TdcIconHicon(hIcon HICON) TdcIcon {
	return TdcIcon{
		curType: 1,
		hIcon:   hIcon,
	}
}

// Creates a new TdcIcon variant with an int value.
func TdcIconInt(id int) TdcIcon {
	return TdcIcon{
		curType: 2,
		id:      uint16(id),
	}
}

// Creates a new TdcIcon variant with a co.TD_ICON value.
func TdcIconTdi(tdIcon co.TD_ICON) TdcIcon {
	return TdcIcon{
		curType: 3,
		tdIcon:  tdIcon,
	}
}

func (me *TdcIcon) IsNone() bool            { return me.curType == 0 }
func (me *TdcIcon) HIcon() (HICON, bool)    { return me.hIcon, me.curType == 1 }
func (me *TdcIcon) Id() (int, bool)         { return int(me.id), me.curType == 2 }
func (me *TdcIcon) Tdi() (co.TD_ICON, bool) { return me.tdIcon, me.curType == 3 }

// Converts the internal value to uintptr.
func (me *TdcIcon) raw() uint64 {
	switch me.curType {
	case 0:
		return 0
	case 1:
		return uint64(me.hIcon)
	case 2:
		return uint64(me.id)
	case 3:
		return uint64(me.tdIcon)
	default:
		panic("Invalid TdcIcon value.")
	}
}
