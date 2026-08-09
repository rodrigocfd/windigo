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

// [IShellLink] COM interface.
//
// [IShellLink]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishelllinkw
type IShellLink struct{ win.IUnknown }

type _IShellLinkVt struct {
	utl.IUnknownVt
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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellLink) IID() *co.IID {
	return &cosh.IID_IShellLink
}

// [GetArguments] method.
//
// [GetArguments]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getarguments
func (me *IShellLink) GetArguments() (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(utl.INFOTIPSIZE) // arbitrary

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).GetArguments,
		me.Ppvt(),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len())))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), nil
	} else {
		return "", hr
	}
}

// [GetDescription] method.
//
// [GetDescription]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getdescription
func (me *IShellLink) GetDescription() (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(utl.INFOTIPSIZE) // arbitrary

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).GetDescription,
		me.Ppvt(),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len())))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), nil
	} else {
		return "", hr
	}
}

// [GetHotkey] method.
//
// [GetHotkey]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-gethotkey
func (me *IShellLink) GetHotkey() (cosh.HOTKEYF, error) {
	return utl.OleCallReturnStruct[cosh.HOTKEYF](me,
		utl.Vt[_IShellLinkVt](me.Ppvt()).GetHotkey)
}

// [GetIconLocation] method.
//
// [GetIconLocation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-geticonlocation
func (me *IShellLink) GetIconLocation() (path string, index int, hr error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)
	var iconIndex uint16

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).GetIconLocation,
		me.Ppvt(),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len()-1)),
		uintptr(unsafe.Pointer(&iconIndex)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), int(iconIndex), nil
	} else {
		return "", 0, hr
	}
}

// [GetPath] method.
//
// [GetPath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getpath
func (me *IShellLink) GetPath(pWfd *win.WIN32_FIND_DATA, flags cosh.SLGP) (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).GetPath,
		me.Ppvt(),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len()-1)),
		uintptr(unsafe.Pointer(pWfd)),
		uintptr(flags))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), nil
	} else {
		return "", hr
	}
}

// [GetShowCmd] method.
//
// [GetShowCmd]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getshowcmd
func (me *IShellLink) GetShowCmd() (co.SW, error) {
	return utl.OleCallReturnStruct[co.SW](me,
		utl.Vt[_IShellLinkVt](me.Ppvt()).GetShowCmd)
}

// [GetWorkingDirectory] method.
//
// [GetWorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getworkingdirectory
func (me *IShellLink) GetWorkingDirectory() (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).GetWorkingDirectory,
		me.Ppvt(),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len()-1)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return wBuf.String(), nil
	} else {
		return "", hr
	}
}

// [Resolve] method.
//
// [Resolve]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-resolve
func (me *IShellLink) Resolve(hWnd win.HWND, flags co.SLR) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).Resolve,
		me.Ppvt(),
		uintptr(hWnd),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [SetArguments] method.
//
// [SetArguments]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setarguments
func (me *IShellLink) SetArguments(args string) error {
	var wArgs wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).SetArguments,
		me.Ppvt(),
		uintptr(wArgs.AllowEmpty(args)))
	return utl.HresultToError(ret)
}

// [SetDescription] method.
//
// [SetDescription]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setdescription
func (me *IShellLink) SetDescription(descr string) error {
	var wDescr wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).SetDescription,
		me.Ppvt(),
		uintptr(wDescr.AllowEmpty(descr)))
	return utl.HresultToError(ret)
}

// [SetHotkey] method.
//
// [SetHotkey]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-sethotkey
func (me *IShellLink) SetHotkey(hotkey cosh.HOTKEYF) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).SetHotkey,
		me.Ppvt(),
		uintptr(hotkey))
	return utl.HresultToError(ret)
}

// [SetIconLocation] method.
//
// [SetIconLocation]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-seticonlocation
func (me *IShellLink) SetIconLocation(path string, index int) error {
	var wPath wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).SetIconLocation,
		me.Ppvt(),
		uintptr(wPath.AllowEmpty(path)),
		uintptr(int32(index)))
	return utl.HresultToError(ret)
}

// [SetPath] method.
//
// [SetPath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setpath
func (me *IShellLink) SetPath(path string) error {
	var wPath wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).SetPath,
		me.Ppvt(),
		uintptr(wPath.AllowEmpty(path)))
	return utl.HresultToError(ret)
}

// [SetRelativePath] method.
//
// [SetRelativePath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setrelativepath
func (me *IShellLink) SetRelativePath(path string) error {
	var wPath wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).SetRelativePath,
		me.Ppvt(),
		uintptr(wPath.AllowEmpty(path)))
	return utl.HresultToError(ret)
}

// [SetShowCmd] method.
//
// [SetShowCmd]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setshowcmd
func (me *IShellLink) SetShowCmd(cmd co.SW) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).SetShowCmd,
		me.Ppvt(),
		uintptr(cmd))
	return utl.HresultToError(ret)
}

// [SetWorkingDirectory] method.
//
// [SetWorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-setworkingdirectory
func (me *IShellLink) SetWorkingDirectory(path string) error {
	var wPath wstr.BufEncoder
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellLinkVt](me.Ppvt()).SetWorkingDirectory,
		me.Ppvt(),
		uintptr(wPath.AllowEmpty(path)))
	return utl.HresultToError(ret)
}
