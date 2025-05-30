//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [CommandLineToArgv] function.
//
// Typically used with [GetCommandLine].
//
// [CommandLineToArgv]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-commandlinetoargvw
func CommandLineToArgv(cmdLine string) ([]string, error) {
	cmdLine16 := wstr.NewBufWith[wstr.Stack20](cmdLine, wstr.EMPTY_IS_NIL)
	var pNumArgs int32

	ret, _, err := syscall.SyscallN(_CommandLineToArgvW.Addr(),
		uintptr(cmdLine16.UnsafePtr()),
		uintptr(unsafe.Pointer(&pNumArgs)))
	if ret == 0 {
		return nil, co.ERROR(err)
	}

	lpPtrs := unsafe.Slice((**uint16)(unsafe.Pointer(ret)), pNumArgs) // []*uint16
	strs := make([]string, 0, pNumArgs)

	for _, lpPtr := range lpPtrs {
		strs = append(strs, wstr.WstrPtrToStr(lpPtr))
	}
	return strs, nil
}

var _CommandLineToArgvW = dll.Shell32.NewProc("CommandLineToArgvW")

// [Shell_NotifyIcon] function.
//
// [Shell_NotifyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
func Shell_NotifyIcon(message co.NIM, data *NOTIFYICONDATA) error {
	ret, _, _ := syscall.SyscallN(_Shell_NotifyIconW.Addr(),
		uintptr(message),
		uintptr(unsafe.Pointer(data)))
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}
	return nil
}

var _Shell_NotifyIconW = dll.Shell32.NewProc("Shell_NotifyIconW")

// [Shell_NotifyIconGetRect] function.
//
// [Shell_NotifyIconGetRect]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyicongetrect
func Shell_NotifyIconGetRect(identifier *NOTIFYICONIDENTIFIER) (RECT, error) {
	var rc RECT
	ret, _, _ := syscall.SyscallN(_Shell_NotifyIconGetRect.Addr(),
		uintptr(unsafe.Pointer(identifier)),
		uintptr(unsafe.Pointer(&rc)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return RECT{}, hr
	}
	return rc, nil
}

var _Shell_NotifyIconGetRect = dll.Shell32.NewProc("Shell_NotifyIconGetRect")

// [SHGetFileInfo] function.
//
// ⚠️ You must defer [HICON.DestroyIcon] on the HIcon member of the returned
// struct.
//
// [SHGetFileInfo]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
func SHGetFileInfo(path string, fileAttrs co.FILE_ATTRIBUTE, flags co.SHGFI) (SHFILEINFO, error) {
	path16 := wstr.NewBufWith[wstr.Stack20](path, wstr.ALLOW_EMPTY)
	var sfi SHFILEINFO

	ret, _, _ := syscall.SyscallN(_SHGetFileInfoW.Addr(),
		uintptr(path16.UnsafePtr()),
		uintptr(fileAttrs),
		uintptr(unsafe.Pointer(&sfi)),
		unsafe.Sizeof(sfi),
		uintptr(flags))

	if (flags&co.SHGFI_EXETYPE) == 0 || (flags&co.SHGFI_SYSICONINDEX) == 0 {
		if ret == 0 {
			return SHFILEINFO{}, co.ERROR_INVALID_PARAMETER
		}
	}
	if (flags & co.SHGFI_EXETYPE) != 0 {
		if ret == 0 {
			return SHFILEINFO{}, co.ERROR_INVALID_PARAMETER
		}
	}

	return sfi, nil
}

var _SHGetFileInfoW = dll.Shell32.NewProc("SHGetFileInfoW")
