//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
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
		uintptr(hThread),
		uintptr(unsafe.Pointer(&exitCode)))
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

// [GetThreadIdealProcessorEx] function.
//
// [GetThreadIdealProcessorEx]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadidealprocessorex
func (hThread HTHREAD) GetThreadIdealProcessorEx() (PROCESSOR_NUMBER, error) {
	var pi PROCESSOR_NUMBER
	ret, _, err := syscall.SyscallN(_GetThreadIdealProcessorEx.Addr(),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&pi)))
	if ret == 0 {
		return PROCESSOR_NUMBER{}, co.ERROR(err)
	}
	return pi, nil
}

var _GetThreadIdealProcessorEx = dll.Kernel32.NewProc("GetThreadIdealProcessorEx")

// [GetThreadIOPendingFlag] function.
//
// [GetThreadIOPendingFlag]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadiopendingflag
func (hThread HTHREAD) GetThreadIOPendingFlag() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(_GetThreadIOPendingFlag.Addr(),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _GetThreadIOPendingFlag = dll.Kernel32.NewProc("GetThreadIOPendingFlag")

// [GetThreadPriority] function.
//
// [GetThreadPriority]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadpriority
func (hThread HTHREAD) GetThreadPriority() (co.THREAD_PRIORITY, error) {
	ret, _, err := syscall.SyscallN(_GetThreadPriority.Addr(),
		uintptr(hThread))
	if ret == utl.THREAD_PRIORITY_ERROR_RETURN {
		return co.THREAD_PRIORITY(0), co.ERROR(err)
	}
	return co.THREAD_PRIORITY(ret), nil
}

var _GetThreadPriority = dll.Kernel32.NewProc("GetThreadPriority")

// [GetThreadPriorityBoost] function.
//
// [GetThreadPriorityBoost]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadpriorityboost
func (hThread HTHREAD) GetThreadPriorityBoost() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(_GetThreadPriorityBoost.Addr(),
		uintptr(hThread),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _GetThreadPriorityBoost = dll.Kernel32.NewProc("GetThreadPriorityBoost")

// [GetThreadTimes] function.
//
// [GetThreadTimes]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadtimes
func (hThread HTHREAD) GetThreadTimes() (HthreadTimes, error) {
	var ftCreation FILETIME
	var ftExit FILETIME
	var ftKernel FILETIME
	var ftUser FILETIME

	ret, _, err := syscall.SyscallN(_GetThreadTimes.Addr(),
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

// Returned by [HTHREAD.GetThreadTimes].
type HthreadTimes struct {
	Creation time.Time
	Exit     time.Time
	Kernel   time.Time
	User     time.Time
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
		uintptr(hThread),
		uintptr(exitCode))
	return utl.ZeroAsGetLastError(ret, err)
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
// For INFINITE, use [HTHREAD.WaitForSingleObjectInfinite].
//
// [WaitForSingleObject]: https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hThread HTHREAD) WaitForSingleObject(milliseconds uint) (co.WAIT, error) {
	ret, _, err := syscall.SyscallN(_WaitForSingleObject.Addr(),
		uintptr(hThread),
		uintptr(milliseconds))
	if co.WAIT(ret) == co.WAIT_FAILED {
		return co.WAIT_FAILED, co.ERROR(err)
	}
	return co.WAIT(ret), nil
}

var _WaitForSingleObject = dll.Kernel32.NewProc("WaitForSingleObject")

// [HTHREAD.WaitForSingleObject] function with INFINITE value.
func (hThread HTHREAD) WaitForSingleObjectInfinite() (co.WAIT, error) {
	return hThread.WaitForSingleObject(utl.INFINITE)
}
