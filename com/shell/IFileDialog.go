/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package shell

import (
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type (
	// IFileDialog > IModalWindow > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialog
	IFileDialog struct{ IModalWindow }

	IFileDialogVtbl struct {
		IModalWindowVtbl
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
)

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-clearclientdata
func (me *IFileDialog) ClearClientData() {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).ClearClientData, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.ClearClientData"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-close
func (me *IFileDialog) Close(hr co.ERROR) {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).Close, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(hr), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.Close"))
	}
}

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getcurrentselection
func (me *IFileDialog) GetCurrentSelection() *IShellItem {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetCurrentSelection, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.GetCurrentSelection"))
	}
	return &IShellItem{
		IUnknown: win.IUnknown{Ppv: ppvQueried},
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfilename
func (me *IFileDialog) GetFileName() string {
	var pv *uint16
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetFileName, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pv)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.GetFileName"))
	}
	name := win.Str.FromUint16Ptr(pv)
	win.CoTaskMemFree(unsafe.Pointer(pv))
	return name
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfiletypeindex
func (me *IFileDialog) GetFileTypeIndex() int {
	idx := uint32(0)
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetFileTypeIndex, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&idx)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.GetFileTypeIndex"))
	}
	return int(idx)
}

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfolder
func (me *IFileDialog) GetFolder() *IShellItem {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetFolder, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.GetFolder"))
	}
	return &IShellItem{
		IUnknown: win.IUnknown{Ppv: ppvQueried},
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getoptions
func (me *IFileDialog) GetOptions() FOS {
	fos := FOS(0)
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetOptions, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&fos)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.GetOptions"))
	}
	return fos
}

// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getresult
func (me *IFileDialog) GetResult() *IShellItem {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).GetResult, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.GetResult"))
	}
	return &IShellItem{
		IUnknown: win.IUnknown{Ppv: ppvQueried},
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setclientguid
func (me *IFileDialog) SetClientGuid(guid *win.GUID) {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetClientGuid, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(guid)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.SetClientGuid"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilename
func (me *IFileDialog) SetFileName(pszName string) {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetFileName, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pszName))), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.SetFileName"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilenamelabel
func (me *IFileDialog) SetFileNameLabel(pszLabel string) {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetFileNameLabel, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pszLabel))), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.SetFileNameLabel"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypeindex
func (me *IFileDialog) SetFileTypeIndex(iFileType int) {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetFileTypeIndex, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(iFileType), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.SetFileTypeIndex"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypes
func (me *IFileDialog) SetFileTypes(rgFilterSpec []FilterSpec) {
	comdlgFiSp := make([]COMDLG_FILTERSPEC, 0, len(rgFilterSpec))
	for _, fiSp := range rgFilterSpec {
		comdlgFiSp = append(comdlgFiSp, // convert FilterSpec to COMDLG_FILTERSPEC
			COMDLG_FILTERSPEC{
				PszName: win.Str.ToUint16Ptr(fiSp.Name),
				PszSpec: win.Str.ToUint16Ptr(fiSp.Spec),
			})
	}

	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetFileTypes, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(len(comdlgFiSp)),
		uintptr(unsafe.Pointer(&comdlgFiSp[0])))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.SetFileTypes"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfolder
func (me *IFileDialog) SetFolder(psi *IShellItem) {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetFolder, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(psi.Ppv)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.SetFolder"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setokbuttonlabel
func (me *IFileDialog) SetOkButtonLabel(pszText string) {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetOkButtonLabel, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pszText))), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.SetOkButtonLabel"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setoptions
func (me *IFileDialog) SetOptions(fos FOS) {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetOptions, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(fos), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.SetOptions"))
	}
}

// https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-settitle
func (me *IFileDialog) SetTitle(pszTitle string) {
	ret, _, _ := syscall.Syscall(
		(*IFileDialogVtbl)(unsafe.Pointer(*me.Ppv)).SetTitle, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(pszTitle))), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		panic(win.NewWinError(lerr, "IFileDialog.SetTitle"))
	}
}
