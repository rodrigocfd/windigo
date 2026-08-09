//go:build windows

package winsh

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IFileOperation] COM interface.
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
//	var op *winsh.IFileOperation
//	_ = win.CoCreateInstance(
//		rel,
//		&cosh.CLSID_FileOperation,
//		nil,
//		co.CLSCTX_ALL,
//		&op,
//	)
//
// [IFileOperation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileoperation
type IFileOperation struct{ win.IUnknown }

type _IFileOperationVt struct {
	utl.IUnknownVt
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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileOperation) IID() *co.IID {
	return &cosh.IID_IFileOperation
}

// [Advise] method.
//
// Paired with [IFileOperation.Unadvise].
//
// [Advise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-advise
func (me *IFileOperation) Advise(fops *IFileOperationProgressSink) (uint32, error) {
	var cookie uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOpenDialogVt](me.Ppvt()).Advise,
		me.Ppvt(),
		fops.Ppvt(),
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
		utl.Vt[_IFileOperationVt](me.Ppvt()).ApplyPropertiesToItem,
		me.Ppvt(),
		item.Ppvt())
	return utl.HresultToError(ret)
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
//	var op *winsh.IFileOperation
//	_ = win.CoCreateInstance(
//		rel,
//		&cosh.CLSID_FileOperation,
//		nil,
//		co.CLSCTX_ALL,
//		&op,
//	)
//
//	var file, dest *winsh.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &file)
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp\\mydir", &dest)
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
	var wCopyName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).CopyItem,
		me.Ppvt(),
		item.Ppvt(),
		destFolder.Ppvt(),
		uintptr(wCopyName.EmptyIsNil(copyName)),
		utl.OlePpvtOrNil(fops))
	return utl.HresultToError(ret)
}

// [DeleteItem] method.
//
// [DeleteItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-deleteitem
func (me *IFileOperation) DeleteItem(item *IShellItem, fops *IFileOperationProgressSink) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).DeleteItem,
		me.Ppvt(),
		item.Ppvt(),
		utl.OlePpvtOrNil(fops))
	return utl.HresultToError(ret)
}

// [GetAnyOperationsAborted] method.
//
// [GetAnyOperationsAborted]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-getanyoperationsaborted
func (me *IFileOperation) GetAnyOperationsAborted() (bool, error) {
	var bVal win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).GetAnyOperationsAborted,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&bVal)))
	return utl.HresultToBoolError(int32(bVal), ret)
}

// [MoveItem] method.
//
// [MoveItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-moveitem
func (me *IFileOperation) MoveItem(
	item, destFolder *IShellItem,
	newName string,
	fops *IFileOperationProgressSink,
) error {
	var wNewName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).MoveItem,
		me.Ppvt(),
		item.Ppvt(),
		destFolder.Ppvt(),
		uintptr(wNewName.AllowEmpty(newName)),
		utl.OlePpvtOrNil(fops))
	return utl.HresultToError(ret)
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
	var wName, wTemplateName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).NewItem,
		me.Ppvt(),
		destFolder.Ppvt(),
		uintptr(fileAtt),
		uintptr(wName.AllowEmpty(name)),
		uintptr(wTemplateName.EmptyIsNil(templateName)),
		utl.OlePpvtOrNil(fops))
	return utl.HresultToError(ret)
}

// [PerformOperations] method.
//
// [PerformOperations]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-performoperations
func (me *IFileOperation) PerformOperations() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IFileOperationVt](me.Ppvt()).PerformOperations)
}

// [RenameItem] method.
//
// [RenameItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-renameitem
func (me *IFileOperation) RenameItem(
	item *IShellItem,
	newName string,
	fops *IFileOperationProgressSink,
) error {
	var wNewName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).RenameItem,
		me.Ppvt(),
		item.Ppvt(),
		uintptr(wNewName.EmptyIsNil(newName)),
		utl.OlePpvtOrNil(fops))
	return utl.HresultToError(ret)
}

// [SetOperationFlags] method.
//
// [SetOperationFlags]:
func (me *IFileOperation) SetOperationFlags(flags cosh.FOF) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).SetOperationFlags,
		me.Ppvt(),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [SetOwnerWindow] method.
//
// [SetOwnerWindow]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-setownerwindow
func (me *IFileOperation) SetOwnerWindow(hWnd win.HWND) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).SetOwnerWindow,
		me.Ppvt(),
		uintptr(hWnd))
	return utl.HresultToError(ret)
}

// [SetProgressMessage] method.
//
// [SetProgressMessage]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-setprogressmessage
func (me *IFileOperation) SetProgressMessage(message string) error {
	var wMessage wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).SetProgressMessage,
		me.Ppvt(),
		uintptr(wMessage.EmptyIsNil(message)))
	return utl.HresultToError(ret)
}

// [Unadvise] method.
//
// Paired with [IFileOperation.Advise].
//
// [Unadvise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifileoperation-unadvise
func (me *IFileOperation) Unadvise(cookie uint32) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileOperationVt](me.Ppvt()).Unadvise,
		me.Ppvt(),
		uintptr(cookie))
	return utl.HresultToError(ret)
}
