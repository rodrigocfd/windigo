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

// ⚠️ You must defer HTHREAD.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthread
func GetCurrentThread() HTHREAD {
	ret, _, _ := syscall.Syscall(proc.GetCurrentThread.Addr(), 0,
		0, 0, 0)
	return HTHREAD(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadid
func GetCurrentThreadId() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetCurrentThreadId.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

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
func (hThread HTHREAD) GetExitCodeThread() (uint32, error) {
	exitCode := uint32(0)
	ret, _, err := syscall.Syscall(proc.GetExitCodeThread.Addr(), 2,
		uintptr(hThread), uintptr(unsafe.Pointer(&exitCode)), 0)
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return exitCode, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessidofthread
func (hThread HTHREAD) GetProcessIdOfThread() (uint32, error) {
	ret, _, err := syscall.Syscall(proc.GetProcessIdOfThread.Addr(), 1,
		uintptr(hThread), 0, 0)
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadid
func (hThread HTHREAD) GetThreadId() (uint32, error) {
	ret, _, err := syscall.Syscall(proc.GetThreadId.Addr(), 1,
		uintptr(hThread), 0, 0)
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadtimes
func (hThread HTHREAD) GetThreadTimes() (
	creationTime, exitTime, kernelTime, userTime FILETIME, e error) {

	ret, _, err := syscall.Syscall6(proc.GetThreadTimes.Addr(), 5,
		uintptr(hThread), uintptr(unsafe.Pointer(&creationTime)),
		uintptr(unsafe.Pointer(&exitTime)), uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)), 0)
	if ret == 0 {
		e = errco.ERROR(err)
	}
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-resumethread
func (hThread HTHREAD) ResumeThread() (uint32, error) {
	ret, _, err := syscall.Syscall(proc.ResumeThread.Addr(), 1,
		uintptr(hThread), 0, 0)
	if int(ret) == -1 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminatethread
func (hThread HTHREAD) TerminateThread(exitCode uint32) error {
	ret, _, err := syscall.Syscall(proc.TerminateThread.Addr(), 2,
		uintptr(hThread), uintptr(exitCode), 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-suspendthread
func (hThread HTHREAD) SuspendThread() (uint32, error) {
	ret, _, err := syscall.Syscall(proc.SuspendThread.Addr(), 1,
		uintptr(hThread), 0, 0)
	if int(ret) == -1 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hThread HTHREAD) WaitForSingleObject(milliseconds uint32) (co.WAIT, error) {
	ret, _, err := syscall.Syscall(proc.WaitForSingleObject.Addr(), 2,
		uintptr(hThread), uintptr(milliseconds), 0)
	if co.WAIT(ret) == co.WAIT_FAILED {
		return co.WAIT_FAILED, errco.ERROR(err)
	}
	return co.WAIT(ret), nil
}
