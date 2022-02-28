package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileopendialog
type IFileOpenDialog struct{ IFileDialog }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IFileOpenDialog.Release().
//
// Example:
//
//  fod := shell.NewIFileOpenDialog(
//      com.CoCreateInstance(
//          shellco.CLSID_FileOpenDialog, nil,
//          comco.CLSCTX_INPROC_SERVER,
//          shellco.IID_IFileOpenDialog),
//  )
//  defer fod.Release()
func NewIFileOpenDialog(base com.IUnknown) IFileOpenDialog {
	return IFileOpenDialog{IFileDialog: NewIFileDialog(base)}
}

// Prefer using IFileOpenDialog.ListResultDisplayNames(), which retrieves the
// paths directly.
//
// ⚠️ You must defer IShellItemArray.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getresults
func (me *IFileOpenDialog) GetResults() IShellItemArray {
	var ppvQueried com.IUnknown
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

// ⚠️ You must defer IShellItemArray.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getselecteditems
func (me *IFileOpenDialog) GetSelectedItems() IShellItemArray {
	var ppvQueried com.IUnknown
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

// Calls IFileOpenDialog.GetResults() and IShellItemArray.ListDisplayNames(),
// returning the files selected by the user.
//
// Example:
//
//  var fod shell.IFileOpenDialog // initialized somewhere
//
//  chosenFiles := fod.ListResultDisplayNames(shellco.SIGDN_FILESYSPATH)
func (me *IFileOpenDialog) ListResultDisplayNames(
	sigdnName shellco.SIGDN) []string {

	isha := me.GetResults()
	defer isha.Release()

	return isha.ListDisplayNames(sigdnName)
}
