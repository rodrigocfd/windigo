//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a thread.
type HTHREAD HANDLE

// [GetCurrentThread] function.
//
// [GetCurrentThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthread
func GetCurrentThread() HTHREAD {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetCurrentThread, "GetCurrentThread"))
	return HTHREAD(ret)
}

var _GetCurrentThread *syscall.Proc

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
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetExitCodeThread, "GetExitCodeThread"),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return exitCode, nil
}

var _GetExitCodeThread *syscall.Proc

// [GetProcessIdOfThread] function.
//
// [GetProcessIdOfThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessidofthread
func (hThread HTHREAD) GetProcessIdOfThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetProcessIdOfThread, "GetProcessIdOfThread"),
		uintptr(hThread))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _GetProcessIdOfThread *syscall.Proc

// [GetThreadId] function.
//
// [GetThreadId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadid
func (hThread HTHREAD) GetThreadId() (uint32, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetThreadId, "GetThreadId"),
		uintptr(hThread))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _GetThreadId *syscall.Proc

// [GetThreadIdealProcessorEx] function.
//
// [GetThreadIdealProcessorEx]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadidealprocessorex
func (hThread HTHREAD) GetThreadIdealProcessorEx() (PROCESSOR_NUMBER, error) {
	var pi PROCESSOR_NUMBER
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetThreadIdealProcessorEx, "GetThreadIdealProcessorEx"),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&pi)))
	if ret == 0 {
		return PROCESSOR_NUMBER{}, co.ERROR(err)
	}
	return pi, nil
}

var _GetThreadIdealProcessorEx *syscall.Proc

// [GetThreadIOPendingFlag] function.
//
// [GetThreadIOPendingFlag]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadiopendingflag
func (hThread HTHREAD) GetThreadIOPendingFlag() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetThreadIOPendingFlag, "GetThreadIOPendingFlag"),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _GetThreadIOPendingFlag *syscall.Proc

// [GetThreadPriority] function.
//
// [GetThreadPriority]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadpriority
func (hThread HTHREAD) GetThreadPriority() (co.THREAD_PRIORITY, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetThreadPriority, "GetThreadPriority"),
		uintptr(hThread))
	if ret == utl.THREAD_PRIORITY_ERROR_RETURN {
		return co.THREAD_PRIORITY(0), co.ERROR(err)
	}
	return co.THREAD_PRIORITY(ret), nil
}

var _GetThreadPriority *syscall.Proc

// [GetThreadPriorityBoost] function.
//
// [GetThreadPriorityBoost]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadpriorityboost
func (hThread HTHREAD) GetThreadPriorityBoost() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetThreadPriorityBoost, "GetThreadPriorityBoost"),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _GetThreadPriorityBoost *syscall.Proc

// [GetThreadTimes] function.
//
// [GetThreadTimes]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadtimes
func (hThread HTHREAD) GetThreadTimes() (HthreadTimes, error) {
	var ftCreation FILETIME
	var ftExit FILETIME
	var ftKernel FILETIME
	var ftUser FILETIME

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetThreadTimes, "GetThreadTimes"),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&ftCreation)),
		uintptr(unsafe.Pointer(&ftExit)),
		uintptr(unsafe.Pointer(&ftKernel)),
		uintptr(unsafe.Pointer(&ftUser)))
	if ret == 0 {
		return HthreadTimes{}, co.ERROR(err)
	}

	return HthreadTimes{
		Creation: ftCreation.ToTime(),
		Exit:     ftExit.ToTime(),
		Kernel:   ftKernel.ToTime(),
		User:     ftUser.ToTime(),
	}, nil
}

var _GetThreadTimes *syscall.Proc

// Returned by [HTHREAD.GetThreadTimes].
type HthreadTimes struct {
	Creation time.Time
	Exit     time.Time
	Kernel   time.Time
	User     time.Time
}

// [ResumeThread] function.
//
// [ResumeThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-resumethread
func (hThread HTHREAD) ResumeThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_ResumeThread, "ResumeThread"),
		uintptr(hThread))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _ResumeThread *syscall.Proc

// [TerminateThread] function.
//
// [TerminateThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminatethread
func (hThread HTHREAD) TerminateThread(exitCode uint32) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_TerminateThread, "TerminateThread"),
		uintptr(hThread),
		uintptr(exitCode))
	return utl.ZeroAsGetLastError(ret, err)
}

var _TerminateThread *syscall.Proc

// [SuspendThread] function.
//
// [SuspendThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-suspendthread
func (hThread HTHREAD) SuspendThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_SuspendThread, "SuspendThread"),
		uintptr(hThread))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _SuspendThread *syscall.Proc

// [WaitForSingleObject] function.
//
// For INFINITE, use [HTHREAD.WaitForSingleObjectInfinite].
//
// Panics if milliseconds is negative.
//
// [WaitForSingleObject]: https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hThread HTHREAD) WaitForSingleObject(milliseconds int) (co.WAIT, error) {
	utl.PanicNeg(milliseconds)
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_WaitForSingleObject, "WaitForSingleObject"),
		uintptr(hThread),
		uintptr(uint32(milliseconds)))
	if co.WAIT(ret) == co.WAIT_FAILED {
		return co.WAIT_FAILED, co.ERROR(err)
	}
	return co.WAIT(ret), nil
}

var _WaitForSingleObject *syscall.Proc

// [HTHREAD.WaitForSingleObject] function with INFINITE value.
func (hThread HTHREAD) WaitForSingleObjectInfinite() (co.WAIT, error) {
	return hThread.WaitForSingleObject(utl.INFINITE)
}
