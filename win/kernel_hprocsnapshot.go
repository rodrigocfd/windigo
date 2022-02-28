package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Handle to a process snapshot.
type HPROCSNAPSHOT HANDLE

// ⚠️ You must defer HPROCSNAPSHOT.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
func CreateToolhelp32Snapshot(
	flags co.TH32CS, processId uint32) (HPROCSNAPSHOT, error) {

	ret, _, err := syscall.Syscall(proc.CreateToolhelp32Snapshot.Addr(), 2,
		uintptr(flags), uintptr(processId), 0)
	if int(ret) == _INVALID_HANDLE_VALUE {
		return HPROCSNAPSHOT(0), errco.ERROR(err)
	}
	return HPROCSNAPSHOT(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hProcSnap HPROCSNAPSHOT) CloseHandle() error {
	ret, _, err := syscall.Syscall(proc.CloseHandle.Addr(), 1,
		uintptr(hProcSnap), 0, 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumModules().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32firstw
func (hProcSnap HPROCSNAPSHOT) Module32First(buf *MODULEENTRY32) (bool, error) {
	ret, _, err := syscall.Syscall(proc.Module32First.Addr(), 2,
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)), 0)
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a module was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumModules().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32nextw
func (hProcSnap HPROCSNAPSHOT) Module32Next(buf *MODULEENTRY32) (bool, error) {
	ret, _, err := syscall.Syscall(proc.Module32Next.Addr(), 2,
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)), 0)
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a module was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumProcesses().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
func (hProcSnap HPROCSNAPSHOT) Process32First(
	buf *PROCESSENTRY32) (bool, error) {

	ret, _, err := syscall.Syscall(proc.Process32First.Addr(), 2,
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)), 0)
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a process was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumProcesses().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
func (hProcSnap HPROCSNAPSHOT) Process32Next(
	buf *PROCESSENTRY32) (bool, error) {

	ret, _, err := syscall.Syscall(proc.Process32Next.Addr(), 2,
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)), 0)
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a process was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumThreads().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32first
func (hProcSnap HPROCSNAPSHOT) Thread32First(buf *THREADENTRY32) (bool, error) {
	ret, _, err := syscall.Syscall(proc.Thread32First.Addr(), 2,
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)), 0)
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a thread was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumThreads().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32next
func (hProcSnap HPROCSNAPSHOT) Thread32Next(buf *THREADENTRY32) (bool, error) {
	ret, _, err := syscall.Syscall(proc.Thread32Next.Addr(), 2,
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)), 0)
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a thread was found
}
