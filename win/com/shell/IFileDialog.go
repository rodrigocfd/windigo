package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialog
type IFileDialog struct{ IModalWindow }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IFileDialog.Release().
func NewIFileDialog(base com.IUnknown) IFileDialog {
	return IFileDialog{IModalWindow: NewIModalWindow(base)}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-clearclientdata
func (me *IFileDialog) ClearClientData() {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).ClearClientData, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-close
func (me *IFileDialog) Close(hr errco.ERROR) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).Close, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hr), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// ⚠️ You must defer IShellItem.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getcurrentselection
func (me *IFileDialog) GetCurrentSelection() IShellItem {
	var ppvQueried com.IUnknown
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetCurrentSelection, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(ppvQueried)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfilename
func (me *IFileDialog) GetFileName() string {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetFileName, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		defer win.HTASKMEM(pv).CoTaskMemFree()
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		return name
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfiletypeindex
func (me *IFileDialog) GetFileTypeIndex() int {
	var idx uint32
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetFileTypeIndex, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&idx)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(idx)
	} else {
		panic(hr)
	}
}

// ⚠️ You must defer IShellItem.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfolder
func (me *IFileDialog) GetFolder() IShellItem {
	var ppvQueried com.IUnknown
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetFolder, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(ppvQueried)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getoptions
func (me *IFileDialog) GetOptions() shellco.FOS {
	var fos shellco.FOS
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetOptions, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&fos)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return fos
	} else {
		panic(hr)
	}
}

// Prefer using IFileDialog.GetResultDisplayName().
//
// ⚠️ You must defer IShellItem.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getresult
func (me *IFileDialog) GetResult() IShellItem {
	var ppvQueried com.IUnknown
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetResult, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(ppvQueried)
	} else {
		panic(hr)
	}
}

// Calls IFileDialog.GetResult() and IShellItem.GetDisplayName(), returning the
// file selected by the user.
//
// Example:
//
//  var fd shell.IFileDialog // initialized somewhere
//
//  chosenPath := fd.GetResultDisplayName(shellco.SIGDN_FILESYSPATH)
func (me *IFileDialog) GetResultDisplayName(sigdnName shellco.SIGDN) string {
	ish := me.GetResult()
	defer ish.Release()

	return ish.GetDisplayName(sigdnName)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setclientguid
func (me *IFileDialog) SetClientGuid(guid *win.GUID) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetClientGuid, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(guid)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilename
func (me *IFileDialog) SetFileName(name string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFileName, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(name))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilenamelabel
func (me *IFileDialog) SetFileNameLabel(label string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFileNameLabel, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(label))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// The index is one-based.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypeindex
func (me *IFileDialog) SetFileTypeIndex(index int) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFileTypeIndex, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(index), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// Example:
//
//  var fd shell.IFileDialog // initialized somewhere
//
//  fd.SetFileTypes([]shell.FilterSpec{
//      {Name: "MP3 audio files", Spec: "*.mp3"},
//      {Name: "All files", Spec: "*.*"},
//  })
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypes
func (me *IFileDialog) SetFileTypes(filterSpec []FilterSpec) {
	comdlgFiSp := make([]COMDLG_FILTERSPEC, 0, len(filterSpec))
	for _, fiSp := range filterSpec {
		comdlgFiSp = append(comdlgFiSp, // convert FilterSpec to COMDLG_FILTERSPEC
			COMDLG_FILTERSPEC{
				PszName: win.Str.ToNativePtr(fiSp.Name),
				PszSpec: win.Str.ToNativePtr(fiSp.Spec),
			})
	}

	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFileTypes, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(len(comdlgFiSp)),
		uintptr(unsafe.Pointer(&comdlgFiSp[0])))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfolder
func (me *IFileDialog) SetFolder(si *IShellItem) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFolder, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(si.Ptr())), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setokbuttonlabel
func (me *IFileDialog) SetOkButtonLabel(text string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetOkButtonLabel, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(text))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// Example:
//
//  var fd shell.IFileDialog // initialized somewhere
//
//  fd.SetOptions(
//      fd.GetOptions() |
//      shellco.FOS_FORCEFILESYSTEM |
//      shellco.FOS_FILEMUSTEXIST)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setoptions
func (me *IFileDialog) SetOptions(fos shellco.FOS) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetOptions, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(fos), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-settitle
func (me *IFileDialog) SetTitle(title string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetTitle, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(title))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
