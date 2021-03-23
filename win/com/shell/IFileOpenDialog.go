package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IFileOpenDialogVtbl struct {
	_IFileDialogVtbl
	GetResults       uintptr
	GetSelectedItems uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileopendialog
type IFileOpenDialog struct {
	IFileDialog // Base IFileDialog > IModalWindow > IUnknown.
}

// Typically uses CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateIFileOpenDialog(dwClsContext co.CLSCTX) IFileOpenDialog {
	clsidFileOpenDialog := win.NewGuid(0xdc1c5a9c, 0xe88a, 0x4dde, 0xa5a1, 0x60f82a20aef7)
	iidIFileOpenDialog := win.NewGuid(0xd57c7288, 0xd4ad, 0x4768, 0xbe02, 0x9d969532d960)

	iUnk, err := win.CoCreateInstance(
		clsidFileOpenDialog, nil, dwClsContext, iidIFileOpenDialog)
	if err != nil {
		panic(err)
	}
	return IFileOpenDialog{
		IFileDialog{
			IModalWindow{IUnknown: iUnk},
		},
	}
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getresults
func (me *IFileOpenDialog) GetResults() IShellItemArray {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IFileOpenDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetResults, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return IShellItemArray{
		win.IUnknown{Ppv: ppvQueried},
	}
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getselecteditems
func (me *IFileOpenDialog) GetSelectedItems() IShellItemArray {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IFileOpenDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetSelectedItems, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return IShellItemArray{
		win.IUnknown{Ppv: ppvQueried},
	}
}
