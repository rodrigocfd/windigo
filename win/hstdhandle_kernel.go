//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [standard device] – standard input, standard output, or
// standard error.
//
// [standard device]: https://learn.microsoft.com/en-us/windows/console/getstdhandle
type HSTDHANDLE HANDLE

// [GetStdHandle] function.
//
// [GetStdHandle]: https://learn.microsoft.com/en-us/windows/console/getstdhandle
func GetStdHandle(handle co.STD) (HSTDHANDLE, error) {
	ret, _, err := syscall.SyscallN(proc.GetStdHandle.Addr(),
		uintptr(handle))
	if int(ret) == _INVALID_HANDLE_VALUE {
		return HSTDHANDLE(0), errco.ERROR(err)
	}
	return HSTDHANDLE(ret), nil
}

// [GetCurrentConsoleFont] function.
//
// [GetCurrentConsoleFont]: https://learn.microsoft.com/en-us/windows/console/getcurrentconsolefont
func (hStd HSTDHANDLE) GetCurrentConsoleFont(
	maximumWindow bool,
	info *CONSOLE_FONT_INFO) error {

	ret, _, err := syscall.SyscallN(proc.GetCurrentConsoleFont.Addr(),
		uintptr(hStd), util.BoolToUintptr(maximumWindow),
		uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [ReadConsole] function.
//
// [ReadConsole]: https://learn.microsoft.com/en-us/windows/console/readconsole
func (hStd HSTDHANDLE) ReadConsole(
	maxChars int,
	inputControl *CONSOLE_READCONSOLE_CONTROL) (string, error) {

	buf := make([]uint16, maxChars+1)
	var numCharsRead uint32

	ret, _, err := syscall.SyscallN(proc.ReadConsole.Addr(),
		uintptr(hStd), uintptr(unsafe.Pointer(&buf[0])), uintptr(maxChars),
		uintptr(unsafe.Pointer(&numCharsRead)),
		uintptr(unsafe.Pointer(inputControl)))
	if ret == 0 {
		return "", errco.ERROR(err)
	}
	return Str.FromNativeSlice(buf), nil
}

// [SetConsoleCursorInfo] function.
//
// [SetConsoleCursorInfo]: https://learn.microsoft.com/en-us/windows/console/setconsolecursorinfo
func (hStd HSTDHANDLE) SetConsoleCursorInfo(info *CONSOLE_CURSOR_INFO) error {
	ret, _, err := syscall.SyscallN(proc.SetConsoleCursorInfo.Addr(),
		uintptr(hStd), uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [SetConsoleCursorPosition] function.
//
// [SetConsoleCursorPosition]: https://learn.microsoft.com/en-us/windows/console/coord-str
func (hStd HSTDHANDLE) SetConsoleCursorPosition(x, y int) error {
	ret, _, err := syscall.SyscallN(proc.SetConsoleCursorPosition.Addr(),
		uintptr(hStd), uintptr(x), uintptr(y))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [SetConsoleDisplayMode] function.
//
// [SetConsoleDisplayMode]: https://learn.microsoft.com/en-us/windows/console/setconsoledisplaymode
func (hStd HSTDHANDLE) SetConsoleDisplayMode(mode co.CONSOLE) (SIZE, error) {
	var coord COORD
	ret, _, err := syscall.SyscallN(proc.SetConsoleDisplayMode.Addr(),
		uintptr(hStd), uintptr(mode), uintptr(unsafe.Pointer(&coord)))
	if ret == 0 {
		return SIZE{}, errco.ERROR(err)
	}
	return SIZE{Cx: int32(coord.X), Cy: int32(coord.Y)}, nil
}

// [SetConsoleMode] function.
//
// [SetConsoleMode]: https://learn.microsoft.com/en-us/windows/console/setconsolemode
func (hStd HSTDHANDLE) SetConsoleMode(mode co.ENABLE) error {
	ret, _, err := syscall.SyscallN(proc.SetConsoleMode.Addr(),
		uintptr(hStd), uintptr(mode))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [SetConsoleScreenBufferSize] function.
//
// [SetConsoleScreenBufferSize]: https://learn.microsoft.com/en-us/windows/console/setconsolescreenbuffersize
func (hStd HSTDHANDLE) SetConsoleScreenBufferSize(x, y int) error {
	ret, _, err := syscall.SyscallN(proc.SetConsoleScreenBufferSize.Addr(),
		uintptr(hStd), uintptr(x), uintptr(y))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [WriteConsole] function.
//
// [WriteConsole]: https://learn.microsoft.com/en-us/windows/console/writeconsole
func (hStd HSTDHANDLE) WriteConsole(text string) (numCharsWritten int, e error) {
	ret, _, err := syscall.SyscallN(proc.WriteConsole.Addr(),
		uintptr(hStd),
		uintptr(unsafe.Pointer(Str.ToNativePtr(text))), uintptr(len(text)),
		uintptr(unsafe.Pointer(&numCharsWritten)), 0)
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return numCharsWritten, nil
}
