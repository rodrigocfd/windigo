package shell

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shtypes/ns-shtypes-comdlg_filterspec
type COMDLG_FILTERSPEC struct {
	PszName *uint16
	PszSpec *uint16
}

// COMDLG_FILTERSPEC syntactic sugar.
type FilterSpec struct {
	Name string
	Spec string
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ns-shobjidl_core-thumbbutton
type THUMBBUTTON struct {
	DwMask  shellco.THB
	IId     uint32
	IBitmap uint32
	HIcon   win.HICON
	szTip   [260]uint16
	DwFlags shellco.THBF
}

func (tb *THUMBBUTTON) SzTip() string { return win.Str.FromUint16Slice(tb.szTip[:]) }
func (tb *THUMBBUTTON) SetSzTip(val string) {
	copy(tb.szTip[:], win.Str.ToUint16Slice(win.Str.Substr(val, 0, len(tb.szTip)-1)))
}
