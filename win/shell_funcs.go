package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Typically used with GetCommandLine().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-commandlinetoargvw
func CommandLineToArgv(cmdLine string) []string {
	var pNumArgs int32
	ret, _, err := syscall.Syscall(proc.CommandLineToArgv.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToNativePtr(cmdLine))),
		uintptr(unsafe.Pointer(&pNumArgs)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}

	lpPtrs := unsafe.Slice((**uint16)(unsafe.Pointer(ret)), pNumArgs) // []*uint16
	strs := make([]string, 0, pNumArgs)

	for _, lpPtr := range lpPtrs {
		strs = append(strs, Str.FromNativePtr(lpPtr))
	}
	return strs
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
func ShellNotifyIcon(message co.NIM, data *NOTIFYICONDATA) error {
	ret, _, err := syscall.Syscall(proc.Shell_NotifyIcon.Addr(), 2,
		uintptr(message), uintptr(unsafe.Pointer(data)), 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// Depends of CoInitializeEx().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
func SHGetFileInfo(
	path string, fileAttributes co.FILE_ATTRIBUTE,
	sfi *SHFILEINFO, flags co.SHGFI) {

	ret, _, err := syscall.Syscall6(proc.SHGetFileInfo.Addr(), 5,
		uintptr(unsafe.Pointer(Str.ToNativePtr(path))),
		uintptr(fileAttributes), uintptr(unsafe.Pointer(sfi)),
		unsafe.Sizeof(*sfi), uintptr(flags), 0)

	if (flags&co.SHGFI_EXETYPE) == 0 || (flags&co.SHGFI_SYSICONINDEX) == 0 {
		if ret == 0 {
			panic(errco.ERROR(err))
		}
	}

	if (flags & co.SHGFI_EXETYPE) != 0 {
		if ret == 0 {
			panic(errco.ERROR(err))
		}
	}
}
