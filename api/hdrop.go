package api

import (
	"sort"
	"syscall"
	"unsafe"
	"wingows/api/proc"
)

type HDROP HANDLE

func (hdrop HDROP) DragFinish() {
	syscall.Syscall(proc.DragFinish.Addr(), 1,
		uintptr(hdrop), 0, 0)
}

func (hdrop HDROP) DragQueryFile(iFile uint32, lpszFile *uint16,
	cch uint32) uint32 {

	ret, _, _ := syscall.Syscall6(proc.DragQueryFile.Addr(), 4,
		uintptr(hdrop), uintptr(iFile), uintptr(unsafe.Pointer(lpszFile)),
		uintptr(cch), 0, 0)
	if ret == 0 {
		panic("DragQueryFile failed.")
	}
	return uint32(ret)
}

func (hdrop HDROP) DragQueryPoint() (*POINT, bool) {
	pt := &POINT{}
	ret, _, _ := syscall.Syscall(proc.DragQueryPoint.Addr(), 2,
		uintptr(hdrop), uintptr(unsafe.Pointer(pt)), 0)
	return pt, ret != 0 // true if dropped within client area
}

// Calls DragQueryFile successively to retrieve all file names, and releases the
// HDROP handle.
func (hdrop HDROP) GetAllFiles() []string {
	count := hdrop.DragQueryFile(0xFFFFFFFF, nil, 0)
	files := make([]string, 0, count)

	for i := uint32(0); i < count; i++ {
		pathLen := hdrop.DragQueryFile(i, nil, 0) + 1 // room for terminating null
		pathBuf := make([]uint16, pathLen)
		hdrop.DragQueryFile(i, &pathBuf[0], pathLen)
		files = append(files, syscall.UTF16ToString(pathBuf))
	}

	hdrop.DragFinish()
	sort.Strings(files)
	return files
}
