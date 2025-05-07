//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a thread.
type HTHREAD HANDLE

// [GetCurrentThread] function.
//
// [GetCurrentThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthread
func GetCurrentThread() HTHREAD {
	ret, _, _ := syscall.SyscallN(_GetCurrentThread.Addr())
	return HTHREAD(ret)
}

var _GetCurrentThread = dll.Kernel32.NewProc("GetCurrentThread")

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hThread HTHREAD) CloseHandle() error {
	return HANDLE(hThread).CloseHandle()
}

// [GetExitCodeThread] function.
//
// [GetExitCodeThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodethread
func (hThread HTHREAD) GetExitCodeThread() (uint32, error) {
	var exitCode uint32
	ret, _, err := syscall.SyscallN(_GetExitCodeThread.Addr(),
		uintptr(hThread), uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return exitCode, nil
}

var _GetExitCodeThread = dll.Kernel32.NewProc("GetExitCodeThread")

// [GetProcessIdOfThread] function.
//
// [GetProcessIdOfThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessidofthread
func (hThread HTHREAD) GetProcessIdOfThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(_GetProcessIdOfThread.Addr(),
		uintptr(hThread))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _GetProcessIdOfThread = dll.Kernel32.NewProc("GetProcessIdOfThread")

// [GetThreadId] function.
//
// [GetThreadId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadid
func (hThread HTHREAD) GetThreadId() (uint32, error) {
	ret, _, err := syscall.SyscallN(_GetThreadId.Addr(),
		uintptr(hThread))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _GetThreadId = dll.Kernel32.NewProc("GetThreadId")

// [GetThreadTimes] function.
//
// [GetThreadTimes]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadtimes
func (hThread HTHREAD) GetThreadTimes() (creation, exit, kernel, user time.Time, wErr error) {
	var ftCreation FILETIME
	var ftExit FILETIME
	var ftKernel FILETIME
	var ftUser FILETIME

	ret, _, err := syscall.SyscallN(_GetThreadTimes.Addr(),
		uintptr(hThread), uintptr(unsafe.Pointer(&ftCreation)),
		uintptr(unsafe.Pointer(&ftExit)),
		uintptr(unsafe.Pointer(&ftKernel)),
		uintptr(unsafe.Pointer(&ftUser)))
	if ret == 0 {
		wErr = co.ERROR(err)
	}
	return ftCreation.ToTime(), ftExit.ToTime(), ftKernel.ToTime(), ftUser.ToTime(), nil
}

var _GetThreadTimes = dll.Kernel32.NewProc("GetThreadTimes")

// [ResumeThread] function.
//
// [ResumeThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-resumethread
func (hThread HTHREAD) ResumeThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(_ResumeThread.Addr(),
		uintptr(hThread))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _ResumeThread = dll.Kernel32.NewProc("ResumeThread")

// [TerminateThread] function.
//
// [TerminateThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminatethread
func (hThread HTHREAD) TerminateThread(exitCode uint32) error {
	ret, _, err := syscall.SyscallN(_TerminateThread.Addr(),
		uintptr(hThread), uintptr(exitCode))
	return util.ZeroToGetLastError(ret, err)
}

var _TerminateThread = dll.Kernel32.NewProc("TerminateThread")

// [SuspendThread] function.
//
// [SuspendThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-suspendthread
func (hThread HTHREAD) SuspendThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(_SuspendThread.Addr(),
		uintptr(hThread))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _SuspendThread = dll.Kernel32.NewProc("SuspendThread")

// [WaitForSingleObject] function.
//
// For INFINITE, use HTHREAD.WaitForSingleObjectInfinite().
//
// [WaitForSingleObject]: https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hThread HTHREAD) WaitForSingleObject(milliseconds uint) (co.WAIT, error) {
	ret, _, err := syscall.SyscallN(_WaitForSingleObject.Addr(),
		uintptr(hThread), uintptr(milliseconds))
	if co.WAIT(ret) == co.WAIT_FAILED {
		return co.WAIT_FAILED, co.ERROR(err)
	}
	return co.WAIT(ret), nil
}

var _WaitForSingleObject = dll.Kernel32.NewProc("WaitForSingleObject")

// [WaitForSingleObject] function with INFINITE value.
//
// [WaitForSingleObject]: https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hThread HTHREAD) WaitForSingleObjectInfinite() (co.WAIT, error) {
	return hThread.WaitForSingleObject(util.INFINITE)
}
