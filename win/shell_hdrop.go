//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
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
	syscall.SyscallN(_DragFinish.Addr(),
		uintptr(hDrop))
}

var _DragFinish = dll.Shell32.NewProc("DragFinish")

// [DragQueryFile] function. Called internally several times until all files are
// retrieved, then the full paths are returned.
//
// ⚠️ If this HDROP comes from an operation from [WS_EX_ACCEPTFILES], you must
// defer HDROP.DragFinish(). If it comes from [RegisterDragDrop], don't call it.
//
// # Example
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
// [WS_EX_ACCEPTFILES]: https://learn.microsoft.com/en-us/windows/win32/winmsg/extended-window-styles
// [RegisterDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-registerdragdrop
func (hDrop HDROP) DragQueryFile() ([]string, error) {
	ret, _, _ := syscall.SyscallN(_DragQueryFileW.Addr(),
		uintptr(hDrop), uintptr(0xffff_ffff), 0, 0)
	if ret == 0 {
		return nil, co.ERROR_INVALID_PARAMETER
	}

	count := uint32(ret)
	var pathBuf [util.MAX_PATH]uint16 // buffer to receive a path
	paths := make([]string, 0, count)

	for i := uint32(0); i < count; i++ {
		ret, _, _ = syscall.SyscallN(_DragQueryFileW.Addr(),
			uintptr(hDrop), uintptr(i), uintptr(unsafe.Pointer(&pathBuf[0])),
			uintptr(uint32(len(pathBuf))))
		if ret == 0 {
			return nil, co.ERROR_INVALID_PARAMETER
		}
		paths = append(paths, wstr.Utf16SliceToStr(pathBuf[:]))
	}

	return paths, nil
}

var _DragQueryFileW = dll.Shell32.NewProc("DragQueryFileW")

// [DragQueryPoint] function.
//
// Returns true if dropped within client area.
//
// [DragQueryPoint]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-dragquerypoint
func (hDrop HDROP) DragQueryPoint() (POINT, bool) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(_DragQueryPoint.Addr(),
		uintptr(hDrop), uintptr(unsafe.Pointer(&pt)))
	return pt, ret != 0
}

var _DragQueryPoint = dll.Shell32.NewProc("DragQueryPoint")
