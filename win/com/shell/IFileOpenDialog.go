package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/errco"
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

// Calls CoCreateInstance(). Usually context is CLSCTX_INPROC_SERVER.
//
// ⚠️ You must defer IFileOpenDialog.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func NewIFileOpenDialog(context co.CLSCTX) IFileOpenDialog {
	iUnk := win.CoCreateInstance(
		shellco.CLSID_FileOpenDialog, nil, context,
		shellco.IID_IFileOpenDialog)
	return IFileOpenDialog{
		IFileDialog{
			IModalWindow{IUnknown: iUnk},
		},
	}
}

// Prefer using IFileOpenDialog.GetResultsDisplayNames().
//
// ⚠️ You must defer IShellItemArray.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getresults
func (me *IFileOpenDialog) GetResults() IShellItemArray {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IFileOpenDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetResults, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IShellItemArray{
			win.IUnknown{Ppv: ppvQueried},
		}
	} else {
		panic(hr)
	}
}

// Calls IFileOpenDialog.GetResults() and IShellItemArray.GetDisplayNames(),
// returning the files selected by the user.
func (me *IFileOpenDialog) GetResultsDisplayNames(
	sigdnName shellco.SIGDN) []string {

	isha := me.GetResults()
	defer isha.Release()

	return isha.GetDisplayNames(sigdnName)
}

// ⚠️ You must defer IShellItemArray.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getselecteditems
func (me *IFileOpenDialog) GetSelectedItems() IShellItemArray {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IFileOpenDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetSelectedItems, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IShellItemArray{
			win.IUnknown{Ppv: ppvQueried},
		}
	} else {
		panic(hr)
	}
}
