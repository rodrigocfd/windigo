//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishelllinkw
type IShellLink interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getarguments
	GetArguments() string

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getdescription
	GetDescription() string

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-geticonlocation
	GetIconLocation() (path string, index int32)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getpath
	GetPath(fd *win.WIN32_FIND_DATA, flags shellco.SLGP) string

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getshowcmd
	GetShowCmd() co.SW

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getworkingdirectory
	GetWorkingDirectory() string

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-resolve
	Resolve(hWnd win.HWND, flags shellco.SLR)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setarguments
	SetArguments(args string)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setdescription
	SetDescription(descr string)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-seticonlocation
	SetIconLocation(path string, index int32)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setpath
	SetPath(path string)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setrelativepath
	SetRelativePath(path string)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setshowcmd
	SetShowCmd(cmd co.SW)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setworkingdirectory
	SetWorkingDirectory(path string)
}

type _IShellLink struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IShellLink.Release().
//
// Example:
//
//	lnk := shell.NewIShellLink(
//		com.CoCreateInstance(
//			shellco.CLSID_ShellLink, nil,
//			comco.CLSCTX_INPROC_SERVER,
//			shellco.IID_IShellLink),
//	)
//	defer lnk.Release()
func NewIShellLink(base com.IUnknown) IShellLink {
	return &_IShellLink{IUnknown: base}
}

func (me *_IShellLink) GetArguments() string {
	buf := make([]uint16, 1024) // arbitrary
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetArguments,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf)
	} else {
		panic(hr)
	}
}

func (me *_IShellLink) GetDescription() string {
	buf := make([]uint16, 1024) // arbitrary
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetDescription,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf)
	} else {
		panic(hr)
	}
}

func (me *_IShellLink) GetIconLocation() (path string, index int32) {
	buf := make([]uint16, 256) // arbitrary
	iconIndex := int32(0)

	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetIconLocation,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1),
		uintptr(unsafe.Pointer(&iconIndex)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf), iconIndex
	} else {
		panic(hr)
	}
}

func (me *_IShellLink) GetPath(
	fd *win.WIN32_FIND_DATA, flags shellco.SLGP) string {

	buf := make([]uint16, 256) // arbitrary
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetPath,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1),
		uintptr(unsafe.Pointer(fd)), uintptr(flags))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf)
	} else {
		panic(hr)
	}
}

func (me *_IShellLink) GetShowCmd() co.SW {
	cmd := co.SW(0)
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetShowCmd,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&cmd)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return cmd
	} else {
		panic(hr)
	}
}

func (me *_IShellLink) GetWorkingDirectory() string {
	buf := make([]uint16, 256) // arbitrary
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).GetWorkingDirectory,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return win.Str.FromNativeSlice(buf)
	} else {
		panic(hr)
	}
}

func (me *_IShellLink) Resolve(hWnd win.HWND, flags shellco.SLR) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).Resolve,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(flags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IShellLink) SetArguments(args string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetArguments,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(args))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IShellLink) SetDescription(descr string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetDescription,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(descr))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IShellLink) SetIconLocation(path string, index int32) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetDescription,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(path))),
		uintptr(index))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IShellLink) SetPath(path string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetPath,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(path))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IShellLink) SetRelativePath(path string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetRelativePath,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(path))), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IShellLink) SetShowCmd(cmd co.SW) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetShowCmd,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(cmd))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IShellLink) SetWorkingDirectory(path string) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellLink)(unsafe.Pointer(*me.Ptr())).SetWorkingDirectory,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(path))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
