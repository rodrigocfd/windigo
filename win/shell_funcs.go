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
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pCmdLine := wbuf.PtrEmptyIsNil(cmdLine)

	var pNumArgs int32

	ret, _, err := syscall.SyscallN(
		dll.Shell(&_CommandLineToArgvW, "CommandLineToArgvW"),
		uintptr(pCmdLine),
		uintptr(unsafe.Pointer(&pNumArgs)))
	if ret == 0 {
		return nil, co.ERROR(err)
	}

	lpPtrs := unsafe.Slice((**uint16)(unsafe.Pointer(ret)), pNumArgs) // []*uint16
	strs := make([]string, 0, pNumArgs)

	for _, lpPtr := range lpPtrs {
		strs = append(strs, wstr.WinPtrToGo(lpPtr))
	}
	return strs, nil
}

var _CommandLineToArgvW *syscall.Proc

// [Shell_NotifyIcon] function.
//
// [Shell_NotifyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
func Shell_NotifyIcon(message co.NIM, data *NOTIFYICONDATA) error {
	ret, _, _ := syscall.SyscallN(
		dll.Shell(&_Shell_NotifyIconW, "Shell_NotifyIconW"),
		uintptr(message),
		uintptr(unsafe.Pointer(data)))
	if ret == 0 {
		return co.ERROR_INVALID_PARAMETER
	}
	return nil
}

var _Shell_NotifyIconW *syscall.Proc

// [Shell_NotifyIconGetRect] function.
//
// [Shell_NotifyIconGetRect]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyicongetrect
func Shell_NotifyIconGetRect(identifier *NOTIFYICONIDENTIFIER) (RECT, error) {
	var rc RECT
	ret, _, _ := syscall.SyscallN(
		dll.Shell(&_Shell_NotifyIconGetRect, "Shell_NotifyIconGetRect"),
		uintptr(unsafe.Pointer(identifier)),
		uintptr(unsafe.Pointer(&rc)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return RECT{}, hr
	}
	return rc, nil
}

var _Shell_NotifyIconGetRect *syscall.Proc

// [SHGetFileInfo] function.
//
// ⚠️ You must defer [HICON.DestroyIcon] on the HIcon member of the returned
// struct.
//
// [SHGetFileInfo]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
func SHGetFileInfo(path string, fileAttrs co.FILE_ATTRIBUTE, flags co.SHGFI) (SHFILEINFO, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pPath := wbuf.PtrAllowEmpty(path)

	var sfi SHFILEINFO

	ret, _, _ := syscall.SyscallN(
		dll.Shell(&_SHGetFileInfoW, "SHGetFileInfoW"),
		uintptr(pPath),
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

var _SHGetFileInfoW *syscall.Proc
