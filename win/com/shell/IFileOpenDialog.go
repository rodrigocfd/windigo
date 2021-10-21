package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileopendialog
type IFileOpenDialog struct{ IFileDialog }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer IFileOpenDialog.Release().
//
// Example:
//
//  fod := shell.NewIFileOpenDialog(
//      win.CoCreateInstance(
//          shellco.CLSID_FileOpenDialog, nil,
//          co.CLSCTX_INPROC_SERVER,
//          shellco.IID_IFileOpenDialog),
//  )
//  defer fod.Release()
func NewIFileOpenDialog(ptr win.IUnknownPtr) IFileOpenDialog {
	return IFileOpenDialog{
		IFileDialog: NewIFileDialog(ptr),
	}
}

// Prefer using IFileOpenDialog.GetResultsDisplayNames().
//
// ⚠️ You must defer IShellItemArray.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getresults
func (me *IFileOpenDialog) GetResults() IShellItemArray {
	var ppvQueried win.IUnknownPtr
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileOpenDialog)(unsafe.Pointer(*me.Ptr())).GetResults, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItemArray(ppvQueried)
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
	var ppvQueried win.IUnknownPtr
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileOpenDialog)(unsafe.Pointer(*me.Ptr())).GetSelectedItems, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItemArray(ppvQueried)
	} else {
		panic(hr)
	}
}
