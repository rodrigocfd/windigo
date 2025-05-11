//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IShellLink] COM interface.
//
// [IShellLink]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishelllinkw
type IShellLink struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellLink) IID() co.IID {
	return co.IID_IShellLink
}

// [GetArguments] method.
//
// [GetArguments]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getarguments
func (me *IShellLink) GetArguments() (string, error) {
	buf := make([]uint16, wutil.INFOTIPSIZE) // arbitrary
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetArguments,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.WstrSliceToStr(buf), nil
	} else {
		return "", hr
	}
}

// [GetDescription] method.
//
// [GetDescription]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getdescription
func (me *IShellLink) GetDescription() (string, error) {
	buf := make([]uint16, wutil.INFOTIPSIZE) // arbitrary
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetDescription,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.WstrSliceToStr(buf), nil
	} else {
		return "", hr
	}
}

// [GetHotkey] method.
//
// [GetHotkey]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-gethotkey
func (me *IShellLink) GetHotkey() (co.HOTKEYF, error) {
	var key uint16
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetHotkey,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&key)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.HOTKEYF(key), nil
	} else {
		return co.HOTKEYF(0), hr
	}
}

// [GetIconLocation] method.
//
// [GetIconLocation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-geticonlocation
func (me *IShellLink) GetIconLocation() (path string, index int, hr error) {
	var buf [wutil.MAX_PATH]uint16 // arbitrary
	var iconIndex uint16

	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetIconLocation,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1),
		uintptr(unsafe.Pointer(&iconIndex)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.WstrSliceToStr(buf[:]), int(iconIndex), nil
	} else {
		return "", 0, hr
	}
}

// [GetPath] method.
//
// [GetPath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getpath
func (me *IShellItem) GetPath(fd *win.WIN32_FIND_DATA, flags co.SLGP) (string, error) {
	var buf [wutil.MAX_PATH]uint16 // arbitrary
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetPath,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1),
		uintptr(unsafe.Pointer(fd)), uintptr(flags))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.WstrSliceToStr(buf[:]), nil
	} else {
		return "", hr
	}
}

// [GetShowCmd] method.
//
// [GetShowCmd]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getshowcmd
func (me *IShellLink) GetShowCmd() (co.SW, error) {
	var cmd co.SW
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetShowCmd,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cmd)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return cmd, nil
	} else {
		return co.SW(0), hr
	}
}

// [GetWorkingDirectory] method.
//
// [GetWorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getworkingdirectory
func (me *IShellLink) GetWorkingDirectory() (string, error) {
	var buf [wutil.MAX_PATH]uint16 // arbitrary
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).GetWorkingDirectory,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)-1))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wstr.WstrSliceToStr(buf[:]), nil
	} else {
		return "", hr
	}
}

// [Resolve] method.
//
// [Resolve]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-resolve
func (me *IShellLink) Resolve(hWnd win.HWND, flags co.SLR) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).Resolve,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hWnd), uintptr(flags))
	return wutil.ErrorAsHResult(ret)
}

// [SetArguments] method.
//
// [SetArguments]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setarguments
func (me *IShellLink) SetArguments(args string) error {
	args16 := wstr.NewBufWith[wstr.Stack20](args, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetArguments,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(args16.UnsafePtr()))
	return wutil.ErrorAsHResult(ret)
}

// [SetDescription] method.
//
// [SetDescription]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setdescription
func (me *IShellLink) SetDescription(descr string) error {
	descr16 := wstr.NewBufWith[wstr.Stack20](descr, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetDescription,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(descr16.UnsafePtr()))
	return wutil.ErrorAsHResult(ret)
}

// [SetHotkey] method.
//
// [SetHotkey]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-sethotkey
func (me *IShellLink) SetHotkey(hotkey co.HOTKEYF) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetHotkey,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hotkey))
	return wutil.ErrorAsHResult(ret)
}

// [SetIconLocation] method.
//
// [SetIconLocation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-seticonlocation
func (me *IShellLink) SetIconLocation(path string, index int) error {
	path16 := wstr.NewBufWith[wstr.Stack20](path, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetIconLocation,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(path16.UnsafePtr()), uintptr(index))
	return wutil.ErrorAsHResult(ret)
}

// [SetPath] method.
//
// [SetPath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setpath
func (me *IShellLink) SetPath(path string) error {
	path16 := wstr.NewBufWith[wstr.Stack20](path, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetPath,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(path16.UnsafePtr()))
	return wutil.ErrorAsHResult(ret)
}

// [SetRelativePath] method.
//
// [SetRelativePath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setrelativepath
func (me *IShellLink) SetRelativePath(path string) error {
	path16 := wstr.NewBufWith[wstr.Stack20](path, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetRelativePath,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(path16.UnsafePtr()))
	return wutil.ErrorAsHResult(ret)
}

// [SetShowCmd] method.
//
// [SetShowCmd]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setshowcmd
func (me *IShellLink) SetShowCmd(cmd co.SW) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetShowCmd,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(cmd))
	return wutil.ErrorAsHResult(ret)
}

// [SetWorkingDirectory] method.
//
// [SetWorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setworkingdirectory
func (me *IShellLink) SetWorkingDirectory(path string) error {
	path16 := wstr.NewBufWith[wstr.Stack20](path, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(
		(*_IShellLinkVt)(unsafe.Pointer(*me.Ppvt())).SetWorkingDirectory,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(path16.UnsafePtr()))
	return wutil.ErrorAsHResult(ret)
}

type _IShellLinkVt struct {
	ole.IUnknownVt
	GetPath             uintptr
	GetIDList           uintptr
	SetIDList           uintptr
	GetDescription      uintptr
	SetDescription      uintptr
	GetWorkingDirectory uintptr
	SetWorkingDirectory uintptr
	GetArguments        uintptr
	SetArguments        uintptr
	GetHotkey           uintptr
	SetHotkey           uintptr
	GetShowCmd          uintptr
	SetShowCmd          uintptr
	GetIconLocation     uintptr
	SetIconLocation     uintptr
	SetRelativePath     uintptr
	Resolve             uintptr
	SetPath             uintptr
}
