package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileopendialog
type IFileOpenDialog interface {
	IFileDialog

	// Prefer using IFileOpenDialog.ListResultDisplayNames(), which retrieves
	// the paths directly.
	//
	// ⚠️ You must defer IShellItemArray.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getresults
	GetResults() IShellItemArray

	// ⚠️ You must defer IShellItemArray.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileopendialog-getselecteditems
	GetSelectedItems() IShellItemArray

	// Calls IFileOpenDialog.GetResults() and
	// IShellItemArray.ListDisplayNames(), returning the files selected by the
	// user.
	//
	// Example:
	//
	//  var fod shell.IFileOpenDialog // initialized somewhere
	//
	//  chosenFiles := fod.ListResultDisplayNames(shellco.SIGDN_FILESYSPATH)
	ListResultDisplayNames(sigdnName shellco.SIGDN) []string
}

type _IFileOpenDialog struct{ IFileDialog }

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
	return &_IFileOpenDialog{IFileDialog: NewIFileDialog(base)}
}

func (me *_IFileOpenDialog) GetResults() IShellItemArray {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileOpenDialog)(unsafe.Pointer(*me.Ptr())).GetResults, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItemArray(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}

func (me *_IFileOpenDialog) GetSelectedItems() IShellItemArray {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileOpenDialog)(unsafe.Pointer(*me.Ptr())).GetSelectedItems, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItemArray(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}

func (me *_IFileOpenDialog) ListResultDisplayNames(
	sigdnName shellco.SIGDN) []string {

	isha := me.GetResults()
	defer isha.Release()

	return isha.ListDisplayNames(sigdnName)
}
