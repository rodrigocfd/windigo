//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IFileDialog] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IFileDialog]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialog
type IFileDialog struct{ IModalWindow }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IFileDialog) IID() co.IID {
	return co.IID_IFileDialog
}

// [AddPlace] method.
//
// [AddPlace]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-addplace
func (me *IFileDialog) AddPlace(si *IShellItem, fdap co.FDAP) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).AddPlace,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(si.Ppvt())),
		uintptr(fdap))
	return utl.ErrorAsHResult(ret)
}

// [Advise] method.
//
// Paired with [IFileDialog.Unadvise].
//
// [Advise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-advise
func (me *IFileDialog) Advise(events *IFileDialogEvents) (cookie uint32, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).Advise,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(events.Ppvt())),
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
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).ClearClientData,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [Close] method.
//
// [Close]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-close
func (me *IFileDialog) Close(hr co.ERROR) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).Close,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hr))
	return utl.ErrorAsHResult(ret)
}

// [GetCurrentSelection] method.
//
// [GetCurrentSelection]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getcurrentselection
func (me *IFileDialog) GetCurrentSelection(releaser *OleReleaser) (*IShellItem, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetCurrentSelection,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItem{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetFileName] method.
//
// [GetFileName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfilename
func (me *IFileDialog) GetFileName() (string, error) {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetFileName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&pv)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer HTASKMEM(pv).CoTaskMemFree()
		name := wstr.DecodePtr((*uint16)(unsafe.Pointer(pv)))
		return name, nil
	} else {
		return "", hr
	}
}

// [GetFileTypeIndex] method.
//
// [GetFileTypeIndex]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfiletypeindex
func (me *IFileDialog) GetFileTypeIndex() (uint, error) {
	var idx uint32
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetFileTypeIndex,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&idx)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return uint(idx), nil
	} else {
		return 0, hr
	}
}

// [GetFolder] method.
//
// [GetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfolder
func (me *IFileDialog) GetFolder(releaser *OleReleaser) (*IShellItem, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetFolder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItem{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetOptions] method.
//
// [GetOptions]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getoptions
func (me *IFileDialog) GetOptions() (co.FOS, error) {
	var fos co.FOS
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetOptions,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&fos)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return fos, nil
	} else {
		return co.FOS(0), hr
	}
}

// [GetResult] method.
//
// Returns the selected item after user confirmation, for single-selection
// dialogs – those without [co.FOS_ALLOWMULTISELECT] option.
//
// For multi-selection dialogs, use [IFileOpenDialog.GetResults].
//
// [GetResult]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getresult
func (me *IFileDialog) GetResult(releaser *OleReleaser) (*IShellItem, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).GetResult,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItem{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [SetClientGuid] method.
//
// [SetClientGuid]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setclientguid
func (me *IFileDialog) SetClientGuid(guid *GUID) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetClientGuid,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(guid)))
	return utl.ErrorAsHResult(ret)
}

// [SetDefaultExtension] method.
//
// [SetDefaultExtension]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setdefaultextension
func (me *IFileDialog) SetDefaultExtension(defaultExt string) error {
	var wDefaultExt wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetDefaultExtension,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wDefaultExt.EmptyIsNil(defaultExt)))
	return utl.ErrorAsHResult(ret)
}

// [SetDefaultFolder] method.
//
// [SetDefaultFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setdefaultfolder
func (me *IFileDialog) SetDefaultFolder(si *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetDefaultFolder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(si.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [SetFileName] method.
//
// [SetFileName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilename
func (me *IFileDialog) SetFileName(name string) error {
	var wName wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFileName,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wName.EmptyIsNil(name)))
	return utl.ErrorAsHResult(ret)
}

// [SetFileNameLabel] method.
//
// [SetFileNameLabel]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilenamelabel
func (me *IFileDialog) SetFileNameLabel(label string) error {
	var wLabel wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFileNameLabel,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wLabel.EmptyIsNil(label)))
	return utl.ErrorAsHResult(ret)
}

// [SetFileTypeIndex] method.
//
// The index is one-based.
//
// [SetFileTypeIndex]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypeindex
func (me *IFileDialog) SetFileTypeIndex(index uint) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFileTypeIndex,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(index))
	return utl.ErrorAsHResult(ret)
}

// [SetFileTypes] method.
//
// Example:
//
//	var fd win.IFileDialog // initialized somewhere
//
//	_ = fd.SetFileTypes([]win.COMDLG_FILTERSPEC{
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
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFileTypes,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(len(filterSpec))),
		uintptr(unsafe.Pointer(&nativeFilters[0])))
	return utl.ErrorAsHResult(ret)
}

// [SetFilter] method.
//
// [SetFilter]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilter
func (me *IFileDialog) SetFilter(filter *IShellItemFilter) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFilter,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(filter.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [SetFolder] method.
//
// [SetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfolder
func (me *IFileDialog) SetFolder(si *IShellItem) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetFolder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(si.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [SetOkButtonLabel] method.
//
// [SetOkButtonLabel]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setokbuttonlabel
func (me *IFileDialog) SetOkButtonLabel(text string) error {
	var wText wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetOkButtonLabel,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wText.EmptyIsNil(text)))
	return utl.ErrorAsHResult(ret)
}

// [SetOptions] method.
//
// Example:
//
//	var fd win.IFileDialog // initialized somewhere
//
//	curOpts, _ := fd.GetOptions()
//	_ = fd.SetOptions(curOpts |
//		co.FOS_FORCEFILESYSTEM |
//		co.FOS_FILEMUSTEXIST |
//		co.FOS_ALLOWMULTISELECT,
//	)
//
// [SetOptions]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setoptions
func (me *IFileDialog) SetOptions(fos co.FOS) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetOptions,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(fos))
	return utl.ErrorAsHResult(ret)
}

// [SetTitle] method.
//
// [SetTitle]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-settitle
func (me *IFileDialog) SetTitle(title string) error {
	var wTitle wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).SetTitle,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(wTitle.EmptyIsNil(title)))
	return utl.ErrorAsHResult(ret)
}

// [Unadvise] method.
//
// Paired with [IFileDialog.Advise].
//
// [Unadvise]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-unadvise
func (me *IFileDialog) Unadvise(cookie uint32) error {
	ret, _, _ := syscall.SyscallN(
		(*_IFileDialogVt)(unsafe.Pointer(*me.Ppvt())).Unadvise,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(cookie))
	return utl.ErrorAsHResult(ret)
}

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
