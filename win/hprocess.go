package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Handle to a process.
type HPROCESS HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hProcess HPROCESS) CloseHandle() error {
	ret, _, err := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hProcess), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodeprocess
func (hProcess HPROCESS) GetExitCodeProcess() uint32 {
	lpExitCode := uint32(0)
	ret, _, err := syscall.Syscall(proc.GetExitCodeProcess.Addr(), 2,
		uintptr(hProcess), uintptr(unsafe.Pointer(&lpExitCode)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return lpExitCode
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessid
func (hProcess HPROCESS) GetProcessId() uint32 {
	ret, _, err := syscall.Syscall(proc.GetProcessId.Addr(), 1,
		uintptr(hProcess), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesstimes
func (hProcess HPROCESS) GetProcessTimes(
	lpCreationTime, lpExitTime, lpKernelTime, lpUserTime *FILETIME) {

	ret, _, err := syscall.Syscall6(proc.GetProcessTimes.Addr(), 5,
		uintptr(hProcess), uintptr(unsafe.Pointer(lpCreationTime)),
		uintptr(unsafe.Pointer(lpExitTime)), uintptr(unsafe.Pointer(lpKernelTime)),
		uintptr(unsafe.Pointer(lpUserTime)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hProcess HPROCESS) WaitForSingleObject(dwMilliseconds uint32) co.WAIT {
	ret, _, err := syscall.Syscall(proc.WaitForSingleObject.Addr(), 2,
		uintptr(hProcess), uintptr(dwMilliseconds), 0)
	if co.WAIT(ret) == co.WAIT_FAILED {
		panic(errco.ERROR(err))
	}
	return co.WAIT(ret)
}
