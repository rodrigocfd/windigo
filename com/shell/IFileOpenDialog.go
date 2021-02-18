/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

type (
	// IFileOpenDialog > IFileDialog > IModalWindow > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileopendialog
	IFileOpenDialog struct{ IFileDialog }

	IFileOpenDialogVtbl struct {
		IFileDialogVtbl
		GetResults       uintptr
		GetSelectedItems uintptr
	}
)

// Typically uses CLSCTX_INPROC_SERVER.
//
// You must defer Release().
func CoCreateIFileOpenDialog(dwClsContext co.CLSCTX) *IFileOpenDialog {
	iUnk, err := win.CoCreateInstance(
		win.NewGuid(0xdc1c5a9c, 0xe88a, 0x4dde, 0xa5a1, 0x60f82a20aef7), // CLSID_FileOpenDialog
		nil,
		dwClsContext,
		win.NewGuid(0xd57c7288, 0xd4ad, 0x4768, 0xbe02, 0x9d969532d960), // IID_IFileOpenDialog
	)
	if err != nil {
		panic(err)
	}
	return &IFileOpenDialog{
		IFileDialog{
			IModalWindow{
				IUnknown: *iUnk,
			},
		},
	}
}

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getresults
func (me *IFileOpenDialog) GetResults() *IShellItemArray {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*IFileOpenDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetResults, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileOpenDialog.GetResults"))
	}
	return &IShellItemArray{
		IUnknown: win.IUnknown{Ppv: ppvQueried},
	}
}

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getselecteditems
func (me *IFileOpenDialog) GetSelectedItems() *IShellItemArray {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*IFileOpenDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetSelectedItems, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileOpenDialog.GetSelectedItems"))
	}
	return &IShellItemArray{
		IUnknown: win.IUnknown{Ppv: ppvQueried},
	}
}
