//go:build windows

package shell

import (
	"strings"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [COMDLG_FILTERSPEC] struct syntactic sugar.
//
// When the native syscall is made, this struct is converted into the raw
// struct.
//
// [COMDLG_FILTERSPEC]: https://learn.microsoft.com/en-us/windows/win32/api/shtypes/ns-shtypes-comdlg_filterspec
type COMDLG_FILTERSPEC struct {
	Name string
	Spec string
}

// [COMDLG_FILTERSPEC] struct.
//
// [COMDLG_FILTERSPEC]: https://learn.microsoft.com/en-us/windows/win32/api/shtypes/ns-shtypes-comdlg_filterspec
type _COMDLG_FILTERSPEC struct {
	PszName *uint16
	PszSpec *uint16
}

// [ITEMIDLIST] struct.
//
// [ITEMIDLIST]: https://learn.microsoft.com/en-us/windows/win32/api/shtypes/ns-shtypes-itemidlist
type ITEMIDLIST uintptr

// Calls [ole.HTASKMEM.CoTaskMemFree].
//
// You usually don't need to call this method directly, since every function
// which returns a [COM] object will require a Releaser to manage the object's
// lifetime.
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func (il *ITEMIDLIST) Release() {
	if *il != 0 {
		ole.HTASKMEM(*il).CoTaskMemFree()
		*il = 0
	}
}

// [PROPERTYKEY] struct.
//
// [PROPERTYKEY]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ns-wtypes-propertykey
type PROPERTYKEY struct {
	data [20]byte // packed
}

// Creates a [PROPERTYKEY] from a string representation.
func PropertykeyFrom(pkey co.PKEY) PROPERTYKEY {
	parts := strings.SplitN(string(pkey), " ", 2)
	fmtId := win.GuidFrom(parts[0])
	pId := wstr.ParseUint(parts[1])

	var out PROPERTYKEY
	out.SetFmdId(fmtId)
	out.SetPId(uint32(pId))
	return out
}

func (pk *PROPERTYKEY) FmtId() win.GUID {
	return *(*win.GUID)(unsafe.Pointer(&pk.data[0]))
}
func (pk *PROPERTYKEY) SetFmdId(fmtId win.GUID) {
	*(*win.GUID)(unsafe.Pointer(&pk.data[0])) = fmtId
}

func (pk *PROPERTYKEY) PId() uint32 {
	return *(*uint32)(unsafe.Pointer(&pk.data[16]))
}
func (pk *PROPERTYKEY) SetPId(pId uint32) {
	*(*uint32)(unsafe.Pointer(&pk.data[16])) = pId
}

// [THUMBBUTTON] struct.
//
// [THUMBBUTTON]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ns-shobjidl_core-thumbbutton
type THUMBBUTTON struct {
	DwMask  co.THB
	IId     uint32
	IBitmap uint32
	HIcon   win.HICON
	szTip   [260]uint16
	DwFlags co.THBF
}

func (tb *THUMBBUTTON) SzTip() string {
	return wstr.WstrSliceToStr(tb.szTip[:])
}
func (tb *THUMBBUTTON) SetSzTip(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(tb.szTip)-1)), tb.szTip[:])
}
