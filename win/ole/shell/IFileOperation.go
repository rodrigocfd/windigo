//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IFileOperation] COM interface.
//
// # Example
//
//	ole.CoInitializeEx(co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer ole.CoUninitialize()
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	op, _ := ole.CoCreateInstance[shell.IFileOperation](
//		rel, co.CLSID_FileOperation, co.CLSCTX_ALL)
//
// [IFileOperation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileoperation
type IFileOperation struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileOperation) IID() co.IID {
	return co.IID_IFileOperation
}

// [Advise] method.
//
// [Advise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-advise
func (me *IFileOperation) Advise(fops *IFileOperationProgressSink) (uint32, error) {
	var cookie uint32

	var pSink unsafe.Pointer
	if fops != nil {
		pSink = unsafe.Pointer(fops.Ppvt())
	}

	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).Advise,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(pSink), uintptr(unsafe.Pointer(&cookie)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return cookie, nil
	} else {
		return 0, hr
	}
}

// [ApplyPropertiesToItem] method.
//
// [ApplyPropertiesToItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-applypropertiestoitem
func (me *IFileOperation) ApplyPropertiesToItem(item *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).ApplyPropertiesToItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [CopyItem] method.
//
// # Example
//
//	ole.CoInitializeEx(co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer ole.CoUninitialize()
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	op, _ := ole.CoCreateInstance[shell.IFileOperation](
//		rel, co.CLSID_FileOperation, co.CLSCTX_ALL)
//
//	file, _ := shell.SHCreateItemFromParsingName[shell.IShellItem](
//		rel, "C:\\Temp\\foo.txt")
//	dest, _ := shell.SHCreateItemFromParsingName[shell.IShellItem](
//		rel, "C:\\Temp\\mydir")
//
//	op.CopyItem(file, dest, "new name.txt", nil)
//	op.PerformOperations()
//
// [CopyItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-copyitem
func (me *IFileOperation) CopyItem(
	item, destFolder *IShellItem,
	copyName string,
	fops *IFileOperationProgressSink,
) error {
	copyName16 := wstr.NewBufWith[wstr.Stack20](copyName, wstr.EMPTY_IS_NIL)

	var pSink unsafe.Pointer
	if fops != nil {
		pSink = unsafe.Pointer(fops.Ppvt())
	}

	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).CopyItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())), uintptr(unsafe.Pointer(destFolder.Ppvt())),
		uintptr(copyName16.UnsafePtr()), uintptr(pSink))
	return utl.ErrorAsHResult(ret)
}

// [DeleteItem] method.
//
// [DeleteItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-deleteitem
func (me *IFileOperation) DeleteItem(item *IShellItem, fops *IFileOperationProgressSink) error {
	var pSink unsafe.Pointer
	if fops != nil {
		pSink = unsafe.Pointer(fops.Ppvt())
	}

	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).DeleteItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())), uintptr(pSink))
	return utl.ErrorAsHResult(ret)
}

// [PerformOperations] method.
//
// [PerformOperations]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-performoperations
func (me *IFileOperation) PerformOperations() error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).PerformOperations,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

type _IFileOperationVt struct {
	ole.IUnknownVt
	Advise                  uintptr
	Unadvise                uintptr
	SetOperationFlags       uintptr
	SetProgressMessage      uintptr
	SetProgressDialog       uintptr
	SetProperties           uintptr
	SetOwnerWindow          uintptr
	ApplyPropertiesToItem   uintptr
	ApplyPropertiesToItems  uintptr
	RenameItem              uintptr
	RenameItems             uintptr
	MoveItem                uintptr
	MoveItems               uintptr
	CopyItem                uintptr
	CopyItems               uintptr
	DeleteItem              uintptr
	DeleteItems             uintptr
	NewItem                 uintptr
	PerformOperations       uintptr
	GetAnyOperationsAborted uintptr
}
