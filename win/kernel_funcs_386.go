//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [VerifyVersionInfo] function.
//
// [VerifyVersionInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
func VerifyVersionInfo(ovi *OSVERSIONINFOEX, typeMask co.VER, conditionMask uint64) (bool, error) {
	ovi.SetDwOsVersionInfoSize() // safety
	cMaskLo, cMaskHi := utl.Break64(conditionMask)

	ret, _, err := syscall.SyscallN(
		dll.Kernel(&_VerifyVersionInfoW, "VerifyVersionInfoW"),
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask),
		uintptr(cMaskLo),
		uintptr(cMaskHi))

	if wErr := co.ERROR(err); ret == 0 && wErr == co.ERROR_OLD_WIN_VERSION {
		return false, nil
	} else if ret == 0 {
		return false, wErr // actual error
	} else {
		return true, nil
	}
}

var _VerifyVersionInfoW *syscall.Proc

// [VerSetConditionMask] function.
//
// [VerSetConditionMask]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-versetconditionmask
func VerSetConditionMask(conditionMask uint64, typeMask co.VER, condition co.VER_COND) uint64 {
	cMaskLo, cMaskHi := utl.Break64(conditionMask)
	retLo, retHi, _ := syscall.SyscallN(
		dll.Kernel(&_VerSetConditionMask, "VerSetConditionMask"),
		uintptr(cMaskLo),
		uintptr(cMaskHi),
		uintptr(typeMask),
		uintptr(condition))
	return utl.Make64(uint32(retLo), uint32(retHi))
}

var _VerSetConditionMask *syscall.Proc
