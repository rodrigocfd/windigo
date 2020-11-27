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
	// IShellItemArray > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemarray
	IShellItemArray struct{ win.IUnknown }

	IShellItemArrayVtbl struct {
		win.IUnknownVtbl
		BindToHandler              uintptr
		GetPropertyStore           uintptr
		GetPropertyDescriptionList uintptr
		GetAttributes              uintptr
		GetCount                   uintptr
		GetItemAt                  uintptr
		EnumItems                  uintptr
	}
)

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getcount
func (me *IShellItemArray) GetCount() int {
	count := uint32(0)
	ret, _, _ := syscall.Syscall(
		(*IShellItemArrayVtbl)(unsafe.Pointer(*me.Ppv)).GetCount, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&count)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IShellItemArray.GetCount"))
	}
	return int(count)
}

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getitemat
func (me *IShellItemArray) GetItemAt(index int) *IShellItem {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*IShellItemArrayVtbl)(unsafe.Pointer(*me.Ppv)).GetItemAt, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(index), uintptr(unsafe.Pointer(&ppvQueried)))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IShellItemArray.GetItemAt"))
	}
	return &IShellItem{
		IUnknown: win.IUnknown{Ppv: ppvQueried},
	}
}

// Syntactic sugar, calls GetDisplayName() on each IShellItem.
func (me *IShellItemArray) GetDisplayNames() []string {
	count := me.GetCount()
	files := make([]string, 0, count)

	for i := 0; i < count; i++ {
		shellItem := me.GetItemAt(i)
		files = append(files, shellItem.GetDisplayName(SIGDN_FILESYSPATH))
		shellItem.Release()
	}

	return files
}
