//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialog
type IFileDialog interface {
	IModalWindow

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-clearclientdata
	ClearClientData()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-close
	Close(hr errco.ERROR)

	// ⚠️ You must defer IShellItem.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getcurrentselection
	GetCurrentSelection() IShellItem

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfilename
	GetFileName() string

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfiletypeindex
	GetFileTypeIndex() int

	// ⚠️ You must defer IShellItem.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getfolder
	GetFolder() IShellItem

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getoptions
	GetOptions() shellco.FOS

	// Prefer using IFileDialog.GetResultDisplayName(), which retrieves the path
	// directly.
	//
	// ⚠️ You must defer IShellItem.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-getresult
	GetResult() IShellItem

	// This helper method calls IFileDialog.GetResult() and
	// IShellItem.GetDisplayName(), returning the single file selected by the
	// user.
	//
	// Example:
	//
	//	var fd shell.IFileDialog // initialized somewhere
	//
	//	chosenPath := fd.GetResultDisplayName(shellco.SIGDN_FILESYSPATH)
	GetResultDisplayName(sigdnName shellco.SIGDN) string

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setclientguid
	SetClientGuid(guid *win.GUID)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilename
	SetFileName(name string)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfilenamelabel
	SetFileNameLabel(label string)

	// The index is one-based.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypeindex
	SetFileTypeIndex(index int)

	// Example:
	//
	//	var fd shell.IFileDialog // initialized somewhere
	//
	//	fd.SetFileTypes([]shell.FilterSpec{
	//		{Name: "MP3 audio files", Spec: "*.mp3"},
	//		{Name: "All files", Spec: "*.*"},
	//	})
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfiletypes
	SetFileTypes(filterSpec []FilterSpec)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setfolder
	SetFolder(si IShellItem)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setokbuttonlabel
	SetOkButtonLabel(text string)

	// Example:
	//
	//	var fd shell.IFileDialog // initialized somewhere
	//
	//	fd.SetOptions(
	//		fd.GetOptions() |
	//		shellco.FOS_FORCEFILESYSTEM |
	//		shellco.FOS_FILEMUSTEXIST)
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-setoptions
	SetOptions(fos shellco.FOS)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ifiledialog-settitle
	SetTitle(title string)
}

type _IFileDialog struct{ IModalWindow }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IFileDialog.Release().
func NewIFileDialog(base com.IUnknown) IFileDialog {
	return &_IFileDialog{IModalWindow: NewIModalWindow(base)}
}

func (me *_IFileDialog) ClearClientData() {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).ClearClientData,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) Close(hr errco.ERROR) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).Close,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hr))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) GetCurrentSelection() IShellItem {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetCurrentSelection,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}

func (me *_IFileDialog) GetFileName() string {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetFileName,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		defer win.HTASKMEM(pv).CoTaskMemFree()
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		return name
	} else {
		panic(hr)
	}
}

func (me *_IFileDialog) GetFileTypeIndex() int {
	var idx uint32
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetFileTypeIndex,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&idx)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(idx)
	} else {
		panic(hr)
	}
}

func (me *_IFileDialog) GetFolder() IShellItem {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetFolder,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}

func (me *_IFileDialog) GetOptions() shellco.FOS {
	var fos shellco.FOS
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetOptions,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&fos)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return fos
	} else {
		panic(hr)
	}
}

func (me *_IFileDialog) GetResult() IShellItem {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).GetResult,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}

func (me *_IFileDialog) GetResultDisplayName(sigdnName shellco.SIGDN) string {
	ish := me.GetResult()
	defer ish.Release()

	return ish.GetDisplayName(sigdnName)
}

func (me *_IFileDialog) SetClientGuid(guid *win.GUID) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetClientGuid,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(guid)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) SetFileName(name string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFileName,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(name))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) SetFileNameLabel(label string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFileNameLabel,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(label))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) SetFileTypeIndex(index int) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFileTypeIndex,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(index))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) SetFileTypes(filterSpec []FilterSpec) {
	comdlgFiSp := make([]COMDLG_FILTERSPEC, 0, len(filterSpec))
	for _, fiSp := range filterSpec {
		comdlgFiSp = append(comdlgFiSp, // convert FilterSpec to COMDLG_FILTERSPEC
			COMDLG_FILTERSPEC{
				PszName: win.Str.ToNativePtr(fiSp.Name),
				PszSpec: win.Str.ToNativePtr(fiSp.Spec),
			})
	}

	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFileTypes,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(len(comdlgFiSp)),
		uintptr(unsafe.Pointer(&comdlgFiSp[0])))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) SetFolder(si IShellItem) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetFolder,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(si.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) SetOkButtonLabel(text string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetOkButtonLabel,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(text))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) SetOptions(fos shellco.FOS) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetOptions,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(fos))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFileDialog) SetTitle(title string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IFileDialog)(unsafe.Pointer(*me.Ptr())).SetTitle,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(title))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
