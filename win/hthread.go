package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Handle to a thread.
type HTHREAD HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hThread HTHREAD) CloseHandle() error {
	ret, _, err := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hThread), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodethread
func (hThread HTHREAD) GetExitCodeThread() uint32 {
	exitCode := uint32(0)
	ret, _, err := syscall.Syscall(proc.GetExitCodeThread.Addr(), 2,
		uintptr(hThread), uintptr(unsafe.Pointer(&exitCode)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return exitCode
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessidofthread
func (hThread HTHREAD) GetProcessIdOfThread() uint32 {
	ret, _, err := syscall.Syscall(proc.GetProcessIdOfThread.Addr(), 1,
		uintptr(hThread), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadid
func (hThread HTHREAD) GetThreadId() uint32 {
	ret, _, err := syscall.Syscall(proc.GetThreadId.Addr(), 1,
		uintptr(hThread), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadtimes
func (hThread HTHREAD) GetThreadTimes(
	creationTime, exitTime, kernelTime, userTime *FILETIME) {

	ret, _, err := syscall.Syscall6(proc.GetThreadTimes.Addr(), 5,
		uintptr(hThread), uintptr(unsafe.Pointer(creationTime)),
		uintptr(unsafe.Pointer(exitTime)), uintptr(unsafe.Pointer(kernelTime)),
		uintptr(unsafe.Pointer(userTime)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hThread HTHREAD) WaitForSingleObject(milliseconds uint32) co.WAIT {
	ret, _, err := syscall.Syscall(proc.WaitForSingleObject.Addr(), 2,
		uintptr(hThread), uintptr(milliseconds), 0)
	if co.WAIT(ret) == co.WAIT_FAILED {
		panic(errco.ERROR(err))
	}
	return co.WAIT(ret)
}
