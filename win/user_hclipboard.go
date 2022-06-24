//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to the clipboard. Actually this handle does not exist, it only
// serves the purpose of logically group the clipboard functions.
//
// This handle is returned by HWND.OpenClipboard().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/clipboard
type HCLIPBOARD struct{}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closeclipboard
func (HCLIPBOARD) CloseClipboard() {
	ret, _, err := syscall.Syscall(proc.CloseClipboard.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-countclipboardformats
func (HCLIPBOARD) CountClipboardFormats() int32 {
	ret, _, err := syscall.Syscall(proc.CountClipboardFormats.Addr(), 0,
		0, 0, 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-emptyclipboard
func (HCLIPBOARD) EmptyClipboard() {
	ret, _, err := syscall.Syscall(proc.EmptyClipboard.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumclipboardformats
func (HCLIPBOARD) EnumClipboardFormats() []co.CF {
	formats := make([]co.CF, 0, 30) // arbitrary
	thisFormat := co.CF(0)

	for {
		ret, _, err := syscall.Syscall(proc.EnumClipboardFormats.Addr(), 1,
			uintptr(thisFormat), 0, 0)

		if ret == 0 {
			if wErr := errco.ERROR(err); wErr == errco.SUCCESS {
				return formats
			} else {
				panic(wErr)
			}
		} else {
			thisFormat = co.CF(ret)
			formats = append(formats, thisFormat)
		}
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardsequencenumber
func (HCLIPBOARD) GetClipboardSequenceNumber() uint32 {
	ret, _, _ := syscall.Syscall(proc.GetClipboardSequenceNumber.Addr(), 0,
		0, 0, 0)
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isclipboardformatavailable
func (HCLIPBOARD) IsClipboardFormatAvailable(format co.CF) bool {
	ret, _, err := syscall.Syscall(proc.IsClipboardFormatAvailable.Addr(), 1,
		uintptr(format), 0, 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return ret != 0
}

// ⚠️ hMem will be owned by the clipboard, do not call HGLOBAL.Free() anymore.
//
// Unless you're doing something specific, prefer HCLIPBOARD.WriteBitmap() or
// HCLIPBOARD.WriteString().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclipboarddata
func (HCLIPBOARD) SetClipboardData(format co.CF, hMem HGLOBAL) {
	ret, _, err := syscall.Syscall(proc.SetClipboardData.Addr(), 2,
		uintptr(format), uintptr(hMem), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
