//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/win/co"
)

// [VerifyVersionInfo] function.
//
// [VerifyVersionInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
func VerifyVersionInfo(ovi *OSVERSIONINFOEX, typeMask co.VER, conditionMask uint64) (bool, error) {
	ovi.SetDwOsVersionInfoSize() // safety

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_VerifyVersionInfoW),
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask),
		uintptr(conditionMask))

	if wErr := co.ERROR(err); ret == 0 && wErr == co.ERROR_OLD_WIN_VERSION {
		return false, nil
	} else if ret == 0 {
		return false, wErr // actual error
	} else {
		return true, nil
	}
}

// [VerSetConditionMask] function.
//
// [VerSetConditionMask]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-versetconditionmask
func VerSetConditionMask(conditionMask uint64, typeMask co.VER, condition co.VER_COND) uint64 {
	ret, _, _ := syscall.SyscallN(dll.Kernel(dll.PROC_VerSetConditionMask),
		uintptr(conditionMask),
		uintptr(typeMask),
		uintptr(condition))
	return uint64(ret)
}
