//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// A handle to a [standard device] – standard input, standard output, or
// standard error.
//
// [standard device]: https://learn.microsoft.com/en-us/windows/console/getstdhandle
type HSTD HANDLE

// [GetStdHandle] function.
//
// [GetStdHandle]: https://learn.microsoft.com/en-us/windows/console/getstdhandle
func GetStdHandle(which co.STD) (HSTD, error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetStdHandle),
		uintptr(which))
	if int(ret) == utl.INVALID_HANDLE_VALUE {
		return HSTD(0), co.ERROR(err)
	}
	return HSTD(ret), nil
}

// [GetCurrentConsoleFont] function.
//
// [GetCurrentConsoleFont]: https://learn.microsoft.com/en-us/windows/console/getcurrentconsolefont
func (hStd HSTD) GetCurrentConsoleFont(maximumWindow bool) (CONSOLE_FONT_INFO, error) {
	var cfi CONSOLE_FONT_INFO
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetCurrentConsoleFont),
		uintptr(hStd),
		utl.BoolToUintptr(maximumWindow),
		uintptr(unsafe.Pointer(&cfi)))
	if ret == 0 {
		return CONSOLE_FONT_INFO{}, co.ERROR(err)
	}
	return cfi, nil
}

// [ReadConsole] function.
//
// [ReadConsole]: https://learn.microsoft.com/en-us/windows/console/readconsole
func (hStd HSTD) ReadConsole(
	maxCharsToRead uint,
	inputControl *CONSOLE_READCONSOLE_CONTROL,
) (text string, numCharsRead uint, wErr error) {
	recvBuf := wstr.NewBufReceiver(maxCharsToRead + 1)
	defer recvBuf.Free()

	var numRead32 uint32

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_ReadConsoleW),
		uintptr(hStd),
		uintptr(recvBuf.UnsafePtr()),
		uintptr(maxCharsToRead),
		uintptr(unsafe.Pointer(&numRead32)),
		uintptr(unsafe.Pointer(inputControl)))
	if ret == 0 {
		return "", 0, co.ERROR(err)
	}
	return recvBuf.String(), uint(numRead32), nil
}

// [SetConsoleCursorInfo] function.
//
// [SetConsoleCursorInfo]: https://learn.microsoft.com/en-us/windows/console/setconsolecursorinfo
func (hStd HSTD) SetConsoleCursorInfo(info *CONSOLE_CURSOR_INFO) error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_SetConsoleCursorInfo),
		uintptr(hStd),
		uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

// [SetConsoleCursorPosition] function.
//
// [SetConsoleCursorPosition]: https://learn.microsoft.com/en-us/windows/console/coord-str
func (hStd HSTD) SetConsoleCursorPosition(x, y int) error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_SetConsoleCursorPosition),
		uintptr(hStd),
		uintptr(int16(x)),
		uintptr(int16(y)))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

// [SetConsoleDisplayMode] function.
//
// [SetConsoleDisplayMode]: https://learn.microsoft.com/en-us/windows/console/setconsoledisplaymode
func (hStd HSTD) SetConsoleDisplayMode(mode co.CONSOLE_MODE) (SIZE, error) {
	var coord COORD
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_SetConsoleDisplayMode),
		uintptr(hStd),
		uintptr(mode),
		uintptr(unsafe.Pointer(&coord)))
	if ret == 0 {
		return SIZE{}, co.ERROR(err)
	}
	return SIZE{Cx: int32(coord.X), Cy: int32(coord.Y)}, nil
}

// [SetConsoleMode] function.
//
// [SetConsoleMode]: https://learn.microsoft.com/en-us/windows/console/setconsolemode
func (hStd HSTD) SetConsoleMode(mode co.ENABLE) error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_SetConsoleMode),
		uintptr(hStd),
		uintptr(mode))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

// [SetConsoleScreenBufferSize] function.
//
// [SetConsoleScreenBufferSize]: https://learn.microsoft.com/en-us/windows/console/setconsolescreenbuffersize
func (hStd HSTD) SetConsoleScreenBufferSize(x, y int) error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_SetConsoleScreenBufferSize),
		uintptr(hStd),
		uintptr(int16(x)),
		uintptr(int16(y)))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

// [SetConsoleTextAttribute] function.
//
// [SetConsoleTextAttribute]: https://learn.microsoft.com/en-us/windows/console/setconsoletextattribute
func (hStd HSTD) SetConsoleTextAttribute(attrs co.CHAR_ATTR) error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_SetConsoleTextAttribute),
		uintptr(hStd),
		uintptr(attrs))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

// [WriteConsole] function.
//
// [WriteConsole]: https://learn.microsoft.com/en-us/windows/console/writeconsole
func (hStd HSTD) WriteConsole(text string) (numCharsWritten uint, wErr error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pText := wbuf.PtrAllowEmpty(text)

	var numWritten32 uint32

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_WriteConsoleW),
		uintptr(hStd),
		uintptr(pText),
		uintptr(uint32(len(text))),
		uintptr(unsafe.Pointer(&numWritten32)),
		0)
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint(numWritten32), nil
}
