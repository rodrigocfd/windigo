//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to an [internal drop structure].
//
// [internal drop structure]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdrop
type HDROP HANDLE

// [DragFinish] function.
//
// This function is rather tricky. Prefer using HDROP.ListFilesAndFinish().
//
// [DragFinish]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragfinish
func (hDrop HDROP) DragFinish() {
	syscall.SyscallN(proc.DragFinish.Addr(),
		uintptr(hDrop))
}

// [DragQueryFile] function.
//
// This function is rather tricky. Prefer using HDROP.ListFilesAndFinish().
//
// [DragQueryFile]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragqueryfilew
func (hDrop HDROP) DragQueryFile(
	iFile uint32,
	lpszFile *uint16,
	cch uint32) uint32 {

	ret, _, err := syscall.SyscallN(proc.DragQueryFile.Addr(),
		uintptr(hDrop), uintptr(iFile), uintptr(unsafe.Pointer(lpszFile)),
		uintptr(cch))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// [DragQueryPoint] function.
//
// Returns true if dropped within client area.
//
// [DragQueryPoint]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragquerypoint
func (hDrop HDROP) DragQueryPoint() (POINT, bool) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(proc.DragQueryPoint.Addr(),
		uintptr(hDrop), uintptr(unsafe.Pointer(&pt)))
	return pt, ret != 0
}
