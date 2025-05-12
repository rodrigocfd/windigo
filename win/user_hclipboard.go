//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// A handle to the [clipboard]. Actually this handle does not exist, it only
// serves the purpose of logically group the clipboard functions.
//
// # Example
//
//	hClip, _ := win.OpenClipboard(win.HWND(0))
//	defer hClip.CloseClipboard()
//
// [clipboard]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/clipboard
type HCLIPBOARD struct{}

// [OpenClipboard] function.
//
// ⚠️ You must defer [HCLIPBOARD.CloseClipboard].
//
// # Example
//
//	hClip, _ := win.OpenClipboard(win.HWND(0))
//	defer hClip.CloseClipboard()
//
// [OpenClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-openclipboard
func OpenClipboard(hWndOwner HWND) (HCLIPBOARD, error) {
	ret, _, err := syscall.SyscallN(_OpenClipboard.Addr(),
		uintptr(hWndOwner))
	if ret == 0 {
		return HCLIPBOARD{}, co.ERROR(err)
	}
	return HCLIPBOARD{}, nil
}

var _OpenClipboard = dll.User32.NewProc("OpenClipboard")

// [CloseClipboard] function.
//
// Paired with [OpenClipboard].
//
// [CloseClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closeclipboard
func (HCLIPBOARD) CloseClipboard() error {
	ret, _, err := syscall.SyscallN(_CloseClipboard.Addr())
	return wutil.ZeroAsGetLastError(ret, err)
}

var _CloseClipboard = dll.User32.NewProc("CloseClipboard")

// [CountClipboardFormats] function.
//
// [CountClipboardFormats]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-countclipboardformats
func (HCLIPBOARD) CountClipboardFormats() (uint, error) {
	ret, _, err := syscall.SyscallN(_CountClipboardFormats.Addr())
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint(ret), nil
}

var _CountClipboardFormats = dll.User32.NewProc("CountClipboardFormats")

// [EmptyClipboard] function.
//
// [EmptyClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-emptyclipboard
func (HCLIPBOARD) EmptyClipboard() error {
	ret, _, err := syscall.SyscallN(_EmptyClipboard.Addr())
	return wutil.ZeroAsGetLastError(ret, err)
}

var _EmptyClipboard = dll.User32.NewProc("EmptyClipboard")

// [EnumClipboardFormats] function.
//
// [EnumClipboardFormats]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumclipboardformats
func (HCLIPBOARD) EnumClipboardFormats() ([]co.CF, error) {
	formats := make([]co.CF, 0)
	thisFormat := co.CF(0)

	for {
		ret, _, err := syscall.SyscallN(_EnumClipboardFormats.Addr(),
			uintptr(thisFormat))

		if ret == 0 {
			if wErr := co.ERROR(err); wErr == co.ERROR_SUCCESS {
				break // no more formats
			} else {
				return nil, wErr
			}
		} else {
			thisFormat = co.CF(ret)
			formats = append(formats, thisFormat)
		}
	}

	return formats, nil
}

var _EnumClipboardFormats = dll.User32.NewProc("EnumClipboardFormats")

// [GetClipboardData] function.
//
// Returns a newly-allocated slice, with a copy of the clipboard contents.
//
// If format is not correct, the function fails weirdly, returning
// [co.ERROR_SUCCESS].
//
// [GetClipboardData]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboarddata
func (HCLIPBOARD) GetClipboardData(format co.CF) ([]byte, error) {
	ret, _, err := syscall.SyscallN(_GetClipboardData.Addr(),
		uintptr(format))
	if ret == 0 {
		return nil, co.ERROR(err)
	}

	hGlobal := HGLOBAL(ret)
	szData, wErr := hGlobal.GlobalSize()
	if wErr != nil {
		return nil, wErr
	}

	ptrData, wErr := hGlobal.GlobalLock()
	if wErr != nil {
		return nil, wErr
	}
	defer hGlobal.GlobalUnlock()

	mem := unsafe.Slice((*byte)(unsafe.Pointer(ptrData)), szData) // map in-memory content
	buf := make([]byte, szData)
	copy(buf, mem)
	return buf, nil
}

var _GetClipboardData = dll.User32.NewProc("GetClipboardData")

// [GetClipboardFormatName] function.
//
// [GetClipboardFormatName]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardformatnamew
func (HCLIPBOARD) GetClipboardFormatName(format co.CF) (string, error) {
	var buf [wutil.MAX_PATH]uint16
	ret, _, err := syscall.SyscallN(_GetClipboardFormatNameW.Addr(),
		uintptr(format))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(buf[:]), nil
}

var _GetClipboardFormatNameW = dll.User32.NewProc("GetClipboardFormatNameW")

// [GetClipboardSequenceNumber] function.
//
// [GetClipboardSequenceNumber]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardsequencenumber
func (HCLIPBOARD) GetClipboardSequenceNumber() int {
	ret, _, _ := syscall.SyscallN(_GetClipboardSequenceNumber.Addr())
	return int(ret)
}

var _GetClipboardSequenceNumber = dll.User32.NewProc("GetClipboardSequenceNumber")

// [IsClipboardFormatAvailable] function.
//
// [IsClipboardFormatAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isclipboardformatavailable
func (HCLIPBOARD) IsClipboardFormatAvailable(format co.CF) (bool, error) {
	ret, _, err := syscall.SyscallN(_IsClipboardFormatAvailable.Addr(),
		uintptr(format))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return false, wErr
	}
	return ret != 0, nil
}

var _IsClipboardFormatAvailable = dll.User32.NewProc("IsClipboardFormatAvailable")

// [SetClipboardData] function.
//
// The data will be copied into the clipboard buffer.
//
// [SetClipboardData]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclipboarddata
func (HCLIPBOARD) SetClipboardData(format co.CF, data []byte) error {
	hGlobal, wErr := GlobalAlloc(co.GMEM_MOVEABLE, uint(len(data)))
	if wErr != nil {
		return wErr
	}

	ptr, wErr := hGlobal.GlobalLock()
	if wErr != nil {
		return wErr
	}
	mem := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), len(data))
	copy(mem, data)
	hGlobal.GlobalUnlock()

	ret, _, err := syscall.SyscallN(_SetClipboardData.Addr(),
		uintptr(format), uintptr(hGlobal)) // HGLOBAL will be owned by the clipboard
	return wutil.ZeroAsGetLastError(ret, err)
}

var _SetClipboardData = dll.User32.NewProc("SetClipboardData")
