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
		dll.Load(dll.KERNEL32, &_kernel_GetCurrentThread, "GetCurrentThread"))
	return HTHREAD(ret)
}

var _kernel_GetCurrentThread *syscall.Proc

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
		dll.Load(dll.KERNEL32, &_kernel_GetExitCodeThread, "GetExitCodeThread"),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return exitCode, nil
}

var _kernel_GetExitCodeThread *syscall.Proc

// [GetProcessIdOfThread] function.
//
// [GetProcessIdOfThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessidofthread
func (hThread HTHREAD) GetProcessIdOfThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetProcessIdOfThread, "GetProcessIdOfThread"),
		uintptr(hThread))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _kernel_GetProcessIdOfThread *syscall.Proc

// [GetThreadId] function.
//
// [GetThreadId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadid
func (hThread HTHREAD) GetThreadId() (uint32, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetThreadId, "GetThreadId"),
		uintptr(hThread))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _kernel_GetThreadId *syscall.Proc

// [GetThreadIdealProcessorEx] function.
//
// [GetThreadIdealProcessorEx]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadidealprocessorex
func (hThread HTHREAD) GetThreadIdealProcessorEx() (PROCESSOR_NUMBER, error) {
	var pi PROCESSOR_NUMBER
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetThreadIdealProcessorEx, "GetThreadIdealProcessorEx"),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&pi)))
	if ret == 0 {
		return PROCESSOR_NUMBER{}, co.ERROR(err)
	}
	return pi, nil
}

var _kernel_GetThreadIdealProcessorEx *syscall.Proc

// [GetThreadIOPendingFlag] function.
//
// [GetThreadIOPendingFlag]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadiopendingflag
func (hThread HTHREAD) GetThreadIOPendingFlag() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetThreadIOPendingFlag, "GetThreadIOPendingFlag"),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _kernel_GetThreadIOPendingFlag *syscall.Proc

// [GetThreadPriority] function.
//
// [GetThreadPriority]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadpriority
func (hThread HTHREAD) GetThreadPriority() (co.THREAD_PRIORITY, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetThreadPriority, "GetThreadPriority"),
		uintptr(hThread))
	if ret == utl.THREAD_PRIORITY_ERROR_RETURN {
		return co.THREAD_PRIORITY(0), co.ERROR(err)
	}
	return co.THREAD_PRIORITY(ret), nil
}

var _kernel_GetThreadPriority *syscall.Proc

// [GetThreadPriorityBoost] function.
//
// [GetThreadPriorityBoost]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadpriorityboost
func (hThread HTHREAD) GetThreadPriorityBoost() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetThreadPriorityBoost, "GetThreadPriorityBoost"),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _kernel_GetThreadPriorityBoost *syscall.Proc

// [GetThreadTimes] function.
//
// [GetThreadTimes]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadtimes
func (hThread HTHREAD) GetThreadTimes() (HthreadTimes, error) {
	var ftCreation FILETIME
	var ftExit FILETIME
	var ftKernel FILETIME
	var ftUser FILETIME

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetThreadTimes, "GetThreadTimes"),
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

var _kernel_GetThreadTimes *syscall.Proc

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
		dll.Load(dll.KERNEL32, &_kernel_ResumeThread, "ResumeThread"),
		uintptr(hThread))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _kernel_ResumeThread *syscall.Proc

// [TerminateThread] function.
//
// [TerminateThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminatethread
func (hThread HTHREAD) TerminateThread(exitCode uint32) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_TerminateThread, "TerminateThread"),
		uintptr(hThread),
		uintptr(exitCode))
	return utl.ZeroAsGetLastError(ret, err)
}

var _kernel_TerminateThread *syscall.Proc

// [SuspendThread] function.
//
// [SuspendThread]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-suspendthread
func (hThread HTHREAD) SuspendThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_SuspendThread, "SuspendThread"),
		uintptr(hThread))
	if int32(ret) == -1 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _kernel_SuspendThread *syscall.Proc

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
		dll.Load(dll.KERNEL32, &_kernel_WaitForSingleObject, "WaitForSingleObject"),
		uintptr(hThread),
		uintptr(uint32(milliseconds)))
	if co.WAIT(ret) == co.WAIT_FAILED {
		return co.WAIT_FAILED, co.ERROR(err)
	}
	return co.WAIT(ret), nil
}

var _kernel_WaitForSingleObject *syscall.Proc

// [HTHREAD.WaitForSingleObject] function with INFINITE value.
func (hThread HTHREAD) WaitForSingleObjectInfinite() (co.WAIT, error) {
	return hThread.WaitForSingleObject(utl.INFINITE)
}
