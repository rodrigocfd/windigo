//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/wstr"
)

// Handle to an [internal drop structure].
//
// [internal drop structure]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdrop
type HDROP HANDLE

// [DragFinish] function.
//
// If you're using [RegisterDragDrop], don't call this function.
//
// [DragFinish]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragfinish
// [RegisterDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-registerdragdrop
func (hDrop HDROP) DragFinish() {
	_, _, _ = syscall.SyscallN(
		dll.Shell.Load(&_shell_DragFinish, "DragFinish"),
		uintptr(hDrop))
}

var _shell_DragFinish *syscall.Proc

// [DragQueryFile] function. Called internally several times until all files are
// retrieved, then the full paths are returned.
//
// ⚠️ If this HDROP comes from an operation from [co.WS_EX_ACCEPTFILES], you
// must defer [HDROP.DragFinish]. If it comes from [RegisterDragDrop], don't
// call it.
//
// Example:
//
//	var hDrop win.HDROP // initialized somewhere
//
//	// defer hDrop.DragFinish() // only if you're not using RegisterDragDrop()
//
//	files, _ := hDrop.DragQueryFile()
//	for _, file := range files {
//		println(file)
//	}
//
// [DragQueryFile]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragqueryfilew
// [RegisterDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-registerdragdrop
func (hDrop HDROP) DragQueryFile() ([]string, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_DragQueryFileW, "DragQueryFileW"),
		uintptr(hDrop),
		uintptr(0xffff_ffff), 0, 0)
	if ret == 0 {
		return nil, co.ERROR_INVALID_PARAMETER
	}

	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	count := uint32(ret)
	paths := make([]string, 0, count)

	for i := uint32(0); i < count; i++ {
		wBuf.Zero()
		ret, _, _ = syscall.SyscallN(
			dll.Shell.Load(&_shell_DragQueryFileW, "DragQueryFileW"),
			uintptr(hDrop),
			uintptr(i),
			uintptr(wBuf.Ptr()),
			uintptr(uint32(wBuf.Len())))
		if ret == 0 {
			return nil, co.ERROR_INVALID_PARAMETER
		}
		paths = append(paths, wBuf.String())
	}

	return paths, nil
}

var _shell_DragQueryFileW *syscall.Proc

// [DragQueryPoint] function.
//
// Returns true if dropped within client area.
//
// [DragQueryPoint]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragquerypoint
func (hDrop HDROP) DragQueryPoint() (POINT, bool) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&_shell_DragQueryPoint, "DragQueryPoint"),
		uintptr(hDrop),
		uintptr(unsafe.Pointer(&pt)))
	return pt, ret != 0
}

var _shell_DragQueryPoint *syscall.Proc
