//go:build windows

package winsh

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IFileDialog] COM interface.
//
// [IFileDialog]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialog
type IFileDialog struct{ IModalWindow }

type _IFileDialogVt struct {
	_IModalWindowVt
	SetFileTypes        uintptr
	SetFileTypeIndex    uintptr
	GetFileTypeIndex    uintptr
	Advise              uintptr
	Unadvise            uintptr
	SetOptions          uintptr
	GetOptions          uintptr
	SetDefaultFolder    uintptr
	SetFolder           uintptr
	GetFolder           uintptr
	GetCurrentSelection uintptr
	SetFileName         uintptr
	GetFileName         uintptr
	SetTitle            uintptr
	SetOkButtonLabel    uintptr
	SetFileNameLabel    uintptr
	GetResult           uintptr
	AddPlace            uintptr
	SetDefaultExtension uintptr
	Close               uintptr
	SetClientGuid       uintptr
	ClearClientData     uintptr
	SetFilter           uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileDialog) IID() *co.IID {
	return &cosh.IID_IFileDialog
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IFileDialog) AddRef(releaser *win.OleReleaser) *IFileDialog {
	return utl.OleNewFromAddRef[*IFileDialog](me, releaser)
}

// [AddPlace] method.
//
// [AddPlace]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-addplace
func (me *IFileDialog) AddPlace(si *IShellItem, fdap cosh.FDAP) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).AddPlace,
		me.Ppvt(),
		si.Ppvt(),
		uintptr(fdap))
	return utl.HresultToError(ret)
}

// [Advise] method.
//
// Paired with [IFileDialog.Unadvise].
//
// [Advise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-advise
func (me *IFileDialog) Advise(events *IFileDialogEvents) (cookie uint32, hr error) {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).Advise,
		me.Ppvt(),
		events.Ppvt(),
		uintptr(unsafe.Pointer(&cookie)))
	if hr = co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return cookie, nil
}

// [ClearClientData] method.
//
// [ClearClientData]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-clearclientdata
func (me *IFileDialog) ClearClientData() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IFileDialogVt](me.Ppvt()).ClearClientData)
}

// [Close] method.
//
// [Close]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-close
func (me *IFileDialog) Close(hr co.ERROR) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).Close,
		me.Ppvt(),
		uintptr(hr))
	return utl.HresultToError(ret)
}

// [GetCurrentSelection] method.
//
// [GetCurrentSelection]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getcurrentselection
func (me *IFileDialog) GetCurrentSelection(releaser *win.OleReleaser) (*IShellItem, error) {
	return utl.OleNewFromCallWithoutParms[*IShellItem](me, releaser,
		utl.Vt[_IFileDialogVt](me.Ppvt()).GetCurrentSelection)
}

// [GetFileName] method.
//
// [GetFileName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfilename
func (me *IFileDialog) GetFileName() (string, error) {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).GetFileName,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&pv)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	defer win.HTASKMEM(pv).CoTaskMemFree()
	name := wstr.DecodePtr((*uint16)(unsafe.Pointer(pv)))
	return name, nil
}

// [GetFileTypeIndex] method.
//
// [GetFileTypeIndex]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfiletypeindex
func (me *IFileDialog) GetFileTypeIndex() (int, error) {
	var idx uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).GetFileTypeIndex,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&idx)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(idx), nil
}

// [GetFolder] method.
//
// [GetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfolder
func (me *IFileDialog) GetFolder(releaser *win.OleReleaser) (*IShellItem, error) {
	return utl.OleNewFromCallWithoutParms[*IShellItem](me, releaser,
		utl.Vt[_IFileDialogVt](me.Ppvt()).GetFolder)
}

// [GetOptions] method.
//
// [GetOptions]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getoptions
func (me *IFileDialog) GetOptions() (cosh.FOS, error) {
	return utl.OleCallReturnStruct[cosh.FOS](me,
		utl.Vt[_IFileDialogVt](me.Ppvt()).GetOptions)
}

// [GetResult] method.
//
// Returns the selected item after user confirmation, for single-selection
// dialogs – those without [cosh.FOS_ALLOWMULTISELECT] option.
//
// For multi-selection dialogs, use [IFileOpenDialog.GetResults].
//
// [GetResult]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getresult
func (me *IFileDialog) GetResult(releaser *win.OleReleaser) (*IShellItem, error) {
	return utl.OleNewFromCallWithoutParms[*IShellItem](me, releaser,
		utl.Vt[_IFileDialogVt](me.Ppvt()).GetResult)
}

// [SetClientGuid] method.
//
// [SetClientGuid]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setclientguid
func (me *IFileDialog) SetClientGuid(pGuid *co.GUID) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetClientGuid,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pGuid)))
	return utl.HresultToError(ret)
}

// [SetDefaultExtension] method.
//
// [SetDefaultExtension]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setdefaultextension
func (me *IFileDialog) SetDefaultExtension(defaultExt string) error {
	var wDefaultExt wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetDefaultExtension,
		me.Ppvt(),
		uintptr(wDefaultExt.EmptyIsNil(defaultExt)))
	return utl.HresultToError(ret)
}

// [SetDefaultFolder] method.
//
// [SetDefaultFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setdefaultfolder
func (me *IFileDialog) SetDefaultFolder(si *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetDefaultFolder,
		me.Ppvt(),
		si.Ppvt())
	return utl.HresultToError(ret)
}

// [SetFileName] method.
//
// [SetFileName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilename
func (me *IFileDialog) SetFileName(name string) error {
	var wName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetFileName,
		me.Ppvt(),
		uintptr(wName.EmptyIsNil(name)))
	return utl.HresultToError(ret)
}

// [SetFileNameLabel] method.
//
// [SetFileNameLabel]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilenamelabel
func (me *IFileDialog) SetFileNameLabel(label string) error {
	var wLabel wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetFileNameLabel,
		me.Ppvt(),
		uintptr(wLabel.EmptyIsNil(label)))
	return utl.HresultToError(ret)
}

// [SetFileTypeIndex] method.
//
// The index is one-based.
//
// [SetFileTypeIndex]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypeindex
func (me *IFileDialog) SetFileTypeIndex(index int) error {
	if index < 1 {
		panic(fmt.Sprintf("Index is one-based: %d.", index))
	}

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetFileTypeIndex,
		me.Ppvt(),
		uintptr(uint32(index)))
	return utl.HresultToError(ret)
}

// [SetFileTypes] method.
//
// Example:
//
//	var fd *winsh.IFileDialog // initialized somewhere
//
//	_ = fd.SetFileTypes([]winsh.COMDLG_FILTERSPEC{
//		{Name: "MP3 audio files", Spec: "*.mp3"},
//		{Name: "All files", Spec: "*.*"},
//	})
//
// [SetFileTypes]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypes
func (me *IFileDialog) SetFileTypes(filterSpec []COMDLG_FILTERSPEC) error {
	nativeFilters := make([]_COMDLG_FILTERSPEC, 0, len(filterSpec))
	for _, fs := range filterSpec {
		nativeFilters = append(nativeFilters, _COMDLG_FILTERSPEC{
			PszName: (*uint16)(wstr.EncodeToPtr(fs.Name)),
			PszSpec: (*uint16)(wstr.EncodeToPtr(fs.Spec)),
		})
	}

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetFileTypes,
		me.Ppvt(),
		uintptr(uint32(len(filterSpec))),
		uintptr(unsafe.Pointer(&nativeFilters[0])))
	return utl.HresultToError(ret)
}

// [SetFilter] method.
//
// [SetFilter]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilter
func (me *IFileDialog) SetFilter(filter *IShellItemFilter) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetFilter,
		me.Ppvt(),
		filter.Ppvt())
	return utl.HresultToError(ret)
}

// [SetFolder] method.
//
// [SetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfolder
func (me *IFileDialog) SetFolder(si *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetFolder,
		me.Ppvt(),
		si.Ppvt())
	return utl.HresultToError(ret)
}

// [SetOkButtonLabel] method.
//
// [SetOkButtonLabel]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setokbuttonlabel
func (me *IFileDialog) SetOkButtonLabel(text string) error {
	var wText wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetOkButtonLabel,
		me.Ppvt(),
		uintptr(wText.EmptyIsNil(text)))
	return utl.HresultToError(ret)
}

// [SetOptions] method.
//
// Example:
//
//	var fd *win.IFileDialog // initialized somewhere
//
//	curOpts, _ := fd.GetOptions()
//	_ = fd.SetOptions(curOpts |
//		co.FOS_FORCEFILESYSTEM |
//		co.FOS_FILEMUSTEXIST |
//		co.FOS_ALLOWMULTISELECT,
//	)
//
// [SetOptions]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setoptions
func (me *IFileDialog) SetOptions(fos cosh.FOS) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetOptions,
		me.Ppvt(),
		uintptr(fos))
	return utl.HresultToError(ret)
}

// [SetTitle] method.
//
// [SetTitle]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-settitle
func (me *IFileDialog) SetTitle(title string) error {
	var wTitle wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).SetTitle,
		me.Ppvt(),
		uintptr(wTitle.EmptyIsNil(title)))
	return utl.HresultToError(ret)
}

// [Unadvise] method.
//
// Paired with [IFileDialog.Advise].
//
// [Unadvise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-unadvise
func (me *IFileDialog) Unadvise(cookie uint32) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IFileDialogVt](me.Ppvt()).Unadvise,
		me.Ppvt(),
		uintptr(cookie))
	return utl.HresultToError(ret)
}
