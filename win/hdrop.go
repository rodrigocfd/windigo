package win

import (
	"sort"
	"strings"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to an internal drop structure.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdrop
type HDROP HANDLE

// Prefer using HDROP.GetFilesAndFinish().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragfinish
func (hDrop HDROP) DragFinish() {
	syscall.Syscall(proc.DragFinish.Addr(), 1,
		uintptr(hDrop), 0, 0)
}

// Prefer using HDROP.GetFilesAndFinish().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragqueryfilew
func (hDrop HDROP) DragQueryFile(
	iFile uint32, lpszFile *uint16, cch uint32) uint32 {

	ret, _, err := syscall.Syscall6(proc.DragQueryFile.Addr(), 4,
		uintptr(hDrop), uintptr(iFile), uintptr(unsafe.Pointer(lpszFile)),
		uintptr(cch), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint32(ret)
}

// Returns true if dropped within client area.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragquerypoint
func (hDrop HDROP) DragQueryPoint() (POINT, bool) {
	pt := POINT{}
	ret, _, _ := syscall.Syscall(proc.DragQueryPoint.Addr(), 2,
		uintptr(hDrop), uintptr(unsafe.Pointer(&pt)), 0)
	return pt, ret != 0
}

// Retrieves all file names with DragQueryFile() and calls DragFinish().
func (hDrop HDROP) GetFilesAndFinish() []string {
	pathBuf := [_MAX_PATH + 1]uint16{} // buffer to receive all paths
	count := hDrop.DragQueryFile(0xffff_ffff, nil, 0)
	paths := make([]string, 0, count) // paths to be returned

	for i := uint32(0); i < count; i++ {
		pathLen := hDrop.DragQueryFile(i, nil, 0) + 1 // room for terminating null
		hDrop.DragQueryFile(i, &pathBuf[0], pathLen)
		paths = append(paths, Str.FromUint16Slice(pathBuf[:]))
	}
	hDrop.DragFinish()

	sort.Slice(paths, func(a, b int) bool {
		return strings.ToUpper(paths[a]) < strings.ToUpper(paths[b]) // case insensitive
	})
	return paths
}
