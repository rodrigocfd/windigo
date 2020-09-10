/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"windigo/win/proc"
)

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdrop
type HDROP HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragfinish
func (hDrop HDROP) DragFinish() {
	syscall.Syscall(proc.DragFinish.Addr(), 1,
		uintptr(hDrop), 0, 0)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragqueryfilew
func (hDrop HDROP) DragQueryFile(iFile uint32, lpszFile *uint16,
	cch uint32) uint32 {

	ret, _, _ := syscall.Syscall6(proc.DragQueryFile.Addr(), 4,
		uintptr(hDrop), uintptr(iFile), uintptr(unsafe.Pointer(lpszFile)),
		uintptr(cch), 0, 0)
	if ret == 0 {
		hDrop.DragFinish() // free resource
		panic("DragQueryFile failed.")
	}
	return uint32(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragquerypoint
func (hDrop HDROP) DragQueryPoint() (*POINT, bool) {
	pt := &POINT{}
	ret, _, _ := syscall.Syscall(proc.DragQueryPoint.Addr(), 2,
		uintptr(hDrop), uintptr(unsafe.Pointer(pt)), 0)
	return pt, ret != 0 // true if dropped within client area
}
