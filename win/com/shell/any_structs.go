package shell

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
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
	DwMask  co.THB
	IId     uint32
	IBitmap uint32
	HIcon   win.HICON
	SzTip   [260]uint16
	DwFlags co.THBF
}
