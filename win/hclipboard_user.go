//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to the [clipboard]. Actually this handle does not exist, it only
// serves the purpose of logically group the clipboard functions.
//
// This handle is returned by HWND.OpenClipboard().
//
// [clipboard]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/clipboard
type HCLIPBOARD struct{}

// [CloseClipboard] function.
//
// [CloseClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closeclipboard
func (HCLIPBOARD) CloseClipboard() {
	ret, _, err := syscall.SyscallN(proc.CloseClipboard.Addr())
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [CountClipboardFormats] function.
//
// [CountClipboardFormats]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-countclipboardformats
func (HCLIPBOARD) CountClipboardFormats() int32 {
	ret, _, err := syscall.SyscallN(proc.CountClipboardFormats.Addr())
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return int32(ret)
}

// [EmptyClipboard] function.
//
// [EmptyClipboard]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-emptyclipboard
func (HCLIPBOARD) EmptyClipboard() {
	ret, _, err := syscall.SyscallN(proc.EmptyClipboard.Addr())
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [EnumClipboardFormats] function.
//
// [EnumClipboardFormats]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumclipboardformats
func (HCLIPBOARD) EnumClipboardFormats() []co.CF {
	formats := make([]co.CF, 0, 30) // arbitrary
	thisFormat := co.CF(0)

	for {
		ret, _, err := syscall.SyscallN(proc.EnumClipboardFormats.Addr(),
			uintptr(thisFormat))

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

// [GetClipboardSequenceNumber] function.
//
// [GetClipboardSequenceNumber]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardsequencenumber
func (HCLIPBOARD) GetClipboardSequenceNumber() uint32 {
	ret, _, _ := syscall.SyscallN(proc.GetClipboardSequenceNumber.Addr())
	return uint32(ret)
}

// [IsClipboardFormatAvailable] function.
//
// [IsClipboardFormatAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isclipboardformatavailable
func (HCLIPBOARD) IsClipboardFormatAvailable(format co.CF) bool {
	ret, _, err := syscall.SyscallN(proc.IsClipboardFormatAvailable.Addr(),
		uintptr(format))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return ret != 0
}

// [SetClipboardData] function.
//
// ⚠️ hMem will be owned by the clipboard, do not call HGLOBAL.Free() anymore.
//
// Unless you're doing something specific, prefer HCLIPBOARD.WriteBitmap() or
// HCLIPBOARD.WriteString().
//
// [SetClipboardData]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclipboarddata
func (HCLIPBOARD) SetClipboardData(format co.CF, hMem HGLOBAL) {
	ret, _, err := syscall.SyscallN(proc.SetClipboardData.Addr(),
		uintptr(format), uintptr(hMem))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
