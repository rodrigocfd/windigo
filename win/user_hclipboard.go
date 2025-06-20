//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
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
	ret, _, err := syscall.SyscallN(
		dll.User(&_OpenClipboard, "OpenClipboard"),
		uintptr(hWndOwner))
	if ret == 0 {
		return HCLIPBOARD{}, co.ERROR(err)
	}
	return HCLIPBOARD{}, nil
}

var _OpenClipboard *syscall.Proc

// [CloseClipboard] function.
//
// Paired with [OpenClipboard].
//
// [CloseClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closeclipboard
func (HCLIPBOARD) CloseClipboard() error {
	ret, _, err := syscall.SyscallN(
		dll.User(&_CloseClipboard, "CloseClipboard"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _CloseClipboard *syscall.Proc

// [CountClipboardFormats] function.
//
// [CountClipboardFormats]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-countclipboardformats
func (HCLIPBOARD) CountClipboardFormats() (uint, error) {
	ret, _, err := syscall.SyscallN(
		dll.User(&_CountClipboardFormats, "CountClipboardFormats"))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return uint(ret), nil
}

var _CountClipboardFormats *syscall.Proc

// [EmptyClipboard] function.
//
// [EmptyClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-emptyclipboard
func (HCLIPBOARD) EmptyClipboard() error {
	ret, _, err := syscall.SyscallN(
		dll.User(&_EmptyClipboard, "EmptyClipboard"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _EmptyClipboard *syscall.Proc

// [EnumClipboardFormats] function.
//
// [EnumClipboardFormats]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumclipboardformats
func (HCLIPBOARD) EnumClipboardFormats() ([]co.CF, error) {
	formats := make([]co.CF, 0)
	thisFormat := co.CF(0)

	for {
		ret, _, err := syscall.SyscallN(
			dll.User(&_EnumClipboardFormats, "EnumClipboardFormats"),
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

var _EnumClipboardFormats *syscall.Proc

// [GetClipboardData] function.
//
// Returns a newly-allocated slice, with a copy of the clipboard contents.
//
// If format is not correct, the function fails weirdly, returning
// [co.ERROR_SUCCESS].
//
// [GetClipboardData]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboarddata
func (HCLIPBOARD) GetClipboardData(format co.CF) ([]byte, error) {
	ret, _, err := syscall.SyscallN(
		dll.User(&_GetClipboardData, "GetClipboardData"),
		uintptr(format))
	if ret == 0 {
		return nil, co.ERROR(err)
	}

	hGlobal := HGLOBAL(ret)
	sliceMem, wErr := hGlobal.GlobalLockSlice() // map in-memory content
	if wErr != nil {
		return nil, wErr
	}
	defer hGlobal.GlobalUnlock()

	buf := make([]byte, len(sliceMem))
	copy(buf, sliceMem)
	return buf, nil
}

var _GetClipboardData *syscall.Proc

// [GetClipboardFormatName] function.
//
// [GetClipboardFormatName]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardformatnamew
func (HCLIPBOARD) GetClipboardFormatName(format co.CF) (string, error) {
	recvBuf := wstr.NewBufReceiver(wstr.BUF_MAX)
	defer recvBuf.Free()

	ret, _, err := syscall.SyscallN(
		dll.User(&_GetClipboardFormatNameW, "GetClipboardFormatNameW"),
		uintptr(format),
		uintptr(recvBuf.UnsafePtr()),
		uintptr(int32(recvBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return recvBuf.String(), nil
}

var _GetClipboardFormatNameW *syscall.Proc

// [GetClipboardSequenceNumber] function.
//
// [GetClipboardSequenceNumber]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardsequencenumber
func (HCLIPBOARD) GetClipboardSequenceNumber() int {
	ret, _, _ := syscall.SyscallN(
		dll.User(&_GetClipboardSequenceNumber, "GetClipboardSequenceNumber"))
	return int(ret)
}

var _GetClipboardSequenceNumber *syscall.Proc

// [IsClipboardFormatAvailable] function.
//
// [IsClipboardFormatAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isclipboardformatavailable
func (HCLIPBOARD) IsClipboardFormatAvailable(format co.CF) (bool, error) {
	ret, _, err := syscall.SyscallN(
		dll.User(&_IsClipboardFormatAvailable, "IsClipboardFormatAvailable"),
		uintptr(format))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return false, wErr
	}
	return ret != 0, nil
}

var _IsClipboardFormatAvailable *syscall.Proc

// [SetClipboardData] function.
//
// The data will be copied into the clipboard buffer.
//
// [SetClipboardData]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclipboarddata
func (HCLIPBOARD) SetClipboardData(format co.CF, data []byte) error {
	hGlobal, wErr := GlobalAlloc(co.GMEM_MOVEABLE, uint(len(data))) // will be owned by the clipboard
	if wErr != nil {
		return wErr
	}

	sliceMem, wErr := hGlobal.GlobalLockSlice()
	if wErr != nil {
		return wErr
	}
	defer hGlobal.GlobalUnlock()
	copy(sliceMem, data)

	ret, _, err := syscall.SyscallN(
		dll.User(&_SetClipboardData, "SetClipboardData"),
		uintptr(format),
		uintptr(hGlobal))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetClipboardData *syscall.Proc
