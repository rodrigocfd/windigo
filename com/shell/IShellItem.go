/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package shell

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// IShellItem > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
	IShellItem struct{ win.IUnknown }

	IShellItemVtbl struct {
		win.IUnknownVtbl
		BindToHandler  uintptr
		GetParent      uintptr
		GetDisplayName uintptr
		GetAttributes  uintptr
		Compare        uintptr
	}
)

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
func NewShellItem(thePath string) *IShellItem {
	iUnk := win.SHCreateItemFromParsingName(thePath, 0,
		win.NewGuid(0x43826d1e, 0xe718, 0x42ee, 0xbc55, 0xa1e261c37bfe)) // IID_IShellItem
	return &IShellItem{
		IUnknown: *iUnk,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getdisplayname
func (me *IShellItem) GetDisplayName(sigdnName SIGDN) string {
	var pv *uint16
	ret, _, _ := syscall.Syscall(
		(*IShellItemVtbl)(unsafe.Pointer(*me.Ppv)).GetDisplayName, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(sigdnName), uintptr(unsafe.Pointer(&pv)))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IShellItem.GetDisplayName"))
	}
	name := win.Str.FromUint16Ptr(pv)
	win.CoTaskMemFree(unsafe.Pointer(pv))
	return name
}
