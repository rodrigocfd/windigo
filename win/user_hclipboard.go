//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Handle to the [clipboard]. Actually this handle does not exist, it only
// serves the purpose of logically group the clipboard functions.
//
// Example:
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
// Example:
//
//	hClip, _ := win.OpenClipboard(win.HWND(0))
//	defer hClip.CloseClipboard()
//
// [OpenClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-openclipboard
func OpenClipboard(hWndOwner HWND) (HCLIPBOARD, error) {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_OpenClipboard, "OpenClipboard"),
		uintptr(hWndOwner))
	if ret == 0 {
		return HCLIPBOARD{}, co.ERROR(err)
	}
	return HCLIPBOARD{}, nil
}

var _user_OpenClipboard *syscall.Proc

// [CloseClipboard] function.
//
// Paired with [OpenClipboard].
//
// [CloseClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closeclipboard
func (HCLIPBOARD) CloseClipboard() error {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_CloseClipboard, "CloseClipboard"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_CloseClipboard *syscall.Proc

// [CountClipboardFormats] function.
//
// [CountClipboardFormats]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-countclipboardformats
func (HCLIPBOARD) CountClipboardFormats() (int, error) {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_CountClipboardFormats, "CountClipboardFormats"))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return 0, wErr
	}
	return int(int32(ret)), nil
}

var _user_CountClipboardFormats *syscall.Proc

// [EmptyClipboard] function.
//
// [EmptyClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-emptyclipboard
func (HCLIPBOARD) EmptyClipboard() error {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_EmptyClipboard, "EmptyClipboard"))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_EmptyClipboard *syscall.Proc

// [EnumClipboardFormats] function.
//
// [EnumClipboardFormats]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumclipboardformats
func (HCLIPBOARD) EnumClipboardFormats() ([]co.CF, error) {
	formats := make([]co.CF, 0)
	thisFormat := co.CF(0)

	for {
		ret, _, err := syscall.SyscallN(
			dll.User.Load(&_user_EnumClipboardFormats, "EnumClipboardFormats"),
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

var _user_EnumClipboardFormats *syscall.Proc

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
		dll.User.Load(&_user_GetClipboardData, "GetClipboardData"),
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

var _user_GetClipboardData *syscall.Proc

// [GetClipboardFormatName] function.
//
// [GetClipboardFormatName]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardformatnamew
func (HCLIPBOARD) GetClipboardFormatName(format co.CF) (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_GetClipboardFormatNameW, "GetClipboardFormatNameW"),
		uintptr(format),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wBuf.String(), nil
}

var _user_GetClipboardFormatNameW *syscall.Proc

// [GetClipboardSequenceNumber] function.
//
// [GetClipboardSequenceNumber]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardsequencenumber
func (HCLIPBOARD) GetClipboardSequenceNumber() int {
	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_GetClipboardSequenceNumber, "GetClipboardSequenceNumber"))
	return int(uint32(ret))
}

var _user_GetClipboardSequenceNumber *syscall.Proc

// [IsClipboardFormatAvailable] function.
//
// [IsClipboardFormatAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isclipboardformatavailable
func (HCLIPBOARD) IsClipboardFormatAvailable(format co.CF) (bool, error) {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_IsClipboardFormatAvailable, "IsClipboardFormatAvailable"),
		uintptr(format))
	if wErr := co.ERROR(err); ret == 0 && wErr != co.ERROR_SUCCESS {
		return false, wErr
	}
	return ret != 0, nil
}

var _user_IsClipboardFormatAvailable *syscall.Proc

// [SetClipboardData] function.
//
// The data will be copied into the clipboard buffer.
//
// [SetClipboardData]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclipboarddata
func (HCLIPBOARD) SetClipboardData(format co.CF, data []byte) error {
	hGlobal, wErr := GlobalAlloc(co.GMEM_MOVEABLE, len(data)) // will be owned by the clipboard
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
		dll.User.Load(&_user_SetClipboardData, "SetClipboardData"),
		uintptr(format),
		uintptr(hGlobal))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_SetClipboardData *syscall.Proc
