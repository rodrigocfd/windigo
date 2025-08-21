//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IFileOperation] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var op *win.IFileOperation
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_FileOperation,
//		nil,
//		co.CLSCTX_ALL,
//		&op,
//	)
//
// [IFileOperation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileoperation
type IFileOperation struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileOperation) IID() co.IID {
	return co.IID_IFileOperation
}

// [Advise] method.
//
// Paired with [IFileOperation.Unadvise].
//
// [Advise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-advise
func (me *IFileOperation) Advise(fops *IFileOperationProgressSink) (uint32, error) {
	var cookie uint32
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).Advise,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(fops.Ppvt())),
		uintptr(unsafe.Pointer(&cookie)))

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
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var op *win.IFileOperation
//	_ = win.CoCreateInstance(
//		rel,
//		co.CLSID_FileOperation,
//		nil,
//		co.CLSCTX_ALL,
//		&op,
//	)
//
//	var file, dest *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &file)
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\mydir", &dest)
//
//	_ = op.CopyItem(file, dest, "new name.txt", nil)
//	_ = op.PerformOperations()
//
// [CopyItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-copyitem
func (me *IFileOperation) CopyItem(
	item, destFolder *IShellItem,
	copyName string,
	fops *IFileOperationProgressSink,
) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pCopyName := wbuf.PtrEmptyIsNil(copyName)

	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).CopyItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())),
		uintptr(unsafe.Pointer(destFolder.Ppvt())),
		uintptr(pCopyName),
		uintptr(ppvtOrNil(fops)))
	return utl.ErrorAsHResult(ret)
}

// [DeleteItem] method.
//
// [DeleteItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-deleteitem
func (me *IFileOperation) DeleteItem(item *IShellItem, fops *IFileOperationProgressSink) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).DeleteItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())),
		uintptr(ppvtOrNil(fops)))
	return utl.ErrorAsHResult(ret)
}

// [GetAnyOperationsAborted] method.
//
// [GetAnyOperationsAborted]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-getanyoperationsaborted
func (me *IFileOperation) GetAnyOperationsAborted() (bool, error) {
	var bVal int32 // BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).GetAnyOperationsAborted,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&bVal)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return bVal != 0, nil
	} else {
		return false, hr
	}
}

// [MoveItem] method.
//
// [MoveItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-moveitem
func (me *IFileOperation) MoveItem(
	item, destFolder *IShellItem,
	newName string,
	fops *IFileOperationProgressSink,
) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pNewName := wbuf.PtrEmptyIsNil(newName)

	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).MoveItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())),
		uintptr(unsafe.Pointer(destFolder.Ppvt())),
		uintptr(pNewName),
		uintptr(ppvtOrNil(fops)))
	return utl.ErrorAsHResult(ret)
}

// [NewItem] method.
//
// [NewItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-newitem
func (me *IFileOperation) NewItem(
	destFolder *IShellItem,
	fileAtt co.FILE_ATTRIBUTE,
	name, templateName string,
	fops *IFileOperationProgressSink,
) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pName := wbuf.PtrAllowEmpty(name)
	pTemplateName := wbuf.PtrEmptyIsNil(templateName)

	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).NewItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(destFolder.Ppvt())),
		uintptr(fileAtt),
		uintptr(pName),
		uintptr(pTemplateName),
		uintptr(ppvtOrNil(fops)))
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

// [RenameItem] method.
//
// [RenameItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-renameitem
func (me *IFileOperation) RenameItem(
	item *IShellItem,
	newName string,
	fops *IFileOperationProgressSink,
) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pNewName := wbuf.PtrEmptyIsNil(newName)

	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).RenameItem,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(item.Ppvt())),
		uintptr(pNewName),
		uintptr(ppvtOrNil(fops)))
	return utl.ErrorAsHResult(ret)
}

// [SetOperationFlags] method.
//
// [SetOperationFlags]:
func (me *IFileOperation) SetOperationFlags(flags co.FOF) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).SetOperationFlags,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(flags))
	return utl.ErrorAsHResult(ret)
}

// [SetOwnerWindow] method.
//
// [SetOwnerWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-setownerwindow
func (me *IFileOperation) SetOwnerWindow(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).SetOwnerWindow,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd))
	return utl.ErrorAsHResult(ret)
}

// [SetProgressMessage] method.
//
// [SetProgressMessage]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-setprogressmessage
func (me *IFileOperation) SetProgressMessage(message string) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pMessage := wbuf.PtrEmptyIsNil(message)

	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).SetProgressMessage,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(pMessage))
	return utl.ErrorAsHResult(ret)
}

// [Unadvise] method.
//
// Paired with [IFileOperation.Advise].
//
// [Unadvise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-unadvise
func (me *IFileOperation) Unadvise(cookie uint32) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileOperationVt)(unsafe.Pointer(*me.Ppvt())).Unadvise,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(cookie))
	return utl.ErrorAsHResult(ret)
}

type _IFileOperationVt struct {
	_IUnknownVt
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
