//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
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
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetStdHandle, "GetStdHandle"),
		uintptr(which))
	if int(ret) == utl.INVALID_HANDLE_VALUE {
		return HSTD(0), co.ERROR(err)
	}
	return HSTD(ret), nil
}

var _kernel_GetStdHandle *syscall.Proc

// [GetCurrentConsoleFont] function.
//
// [GetCurrentConsoleFont]: https://learn.microsoft.com/en-us/windows/console/getcurrentconsolefont
func (hStd HSTD) GetCurrentConsoleFont(maximumWindow bool) (CONSOLE_FONT_INFO, error) {
	var cfi CONSOLE_FONT_INFO
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetCurrentConsoleFont, "GetCurrentConsoleFont"),
		uintptr(hStd),
		utl.BoolToUintptr(maximumWindow),
		uintptr(unsafe.Pointer(&cfi)))
	if ret == 0 {
		return CONSOLE_FONT_INFO{}, co.ERROR(err)
	}
	return cfi, nil
}

var _kernel_GetCurrentConsoleFont *syscall.Proc

// [GetCurrentConsoleFontEx] function.
//
// [GetCurrentConsoleFontEx]: https://learn.microsoft.com/en-us/windows/console/getcurrentconsolefontex
func (hStd HSTD) GetCurrentConsoleFontEx(maximumWindow bool) (CONSOLE_FONT_INFOEX, error) {
	var cfix CONSOLE_FONT_INFOEX
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetCurrentConsoleFontEx, "GetCurrentConsoleFontEx"),
		uintptr(hStd),
		utl.BoolToUintptr(maximumWindow),
		uintptr(unsafe.Pointer(&cfix)))
	if ret == 0 {
		return CONSOLE_FONT_INFOEX{}, co.ERROR(err)
	}
	return cfix, nil
}

var _kernel_GetCurrentConsoleFontEx *syscall.Proc

// [GetLargestConsoleWindowSize] function.
//
// [GetLargestConsoleWindowSize]: https://learn.microsoft.com/en-us/windows/console/getlargestconsolewindowsize
func (hStd HSTD) GetLargestConsoleWindowSize() (COORD, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetLargestConsoleWindowSize, "GetLargestConsoleWindowSize"),
		uintptr(hStd))
	if ret == 0 {
		return COORD{}, co.ERROR(err)
	}
	return COORD{int16(LOWORD(uint32(ret))), int16(HIWORD(uint32(ret)))}, nil
}

var _kernel_GetLargestConsoleWindowSize *syscall.Proc

// [GetNumberOfConsoleInputEvents] function.
//
// [GetNumberOfConsoleInputEvents]: https://learn.microsoft.com/en-us/windows/console/getnumberofconsoleinputevents
func (hStd HSTD) GetNumberOfConsoleInputEvents() (int, error) {
	var numberOfEvents uint32
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_GetNumberOfConsoleInputEvents, "GetNumberOfConsoleInputEvents"),
		uintptr(hStd),
		uintptr(unsafe.Pointer(&numberOfEvents)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return int(numberOfEvents), nil
}

var _kernel_GetNumberOfConsoleInputEvents *syscall.Proc

// [ReadConsole] function.
//
// Panics if maxCharsToRead is negative.
//
// [ReadConsole]: https://learn.microsoft.com/en-us/windows/console/readconsole
func (hStd HSTD) ReadConsole(
	maxCharsToRead int,
	inputControl *CONSOLE_READCONSOLE_CONTROL,
) (text string, numCharsRead int, wErr error) {
	utl.PanicNeg(maxCharsToRead)

	var wBuf wstr.BufDecoder
	wBuf.Alloc(maxCharsToRead + 1)

	var numRead32 uint32

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_ReadConsoleW, "ReadConsoleW"),
		uintptr(hStd),
		uintptr(wBuf.Ptr()),
		uintptr(uint32(maxCharsToRead)),
		uintptr(unsafe.Pointer(&numRead32)),
		uintptr(unsafe.Pointer(inputControl)))
	if ret == 0 {
		return "", 0, co.ERROR(err)
	}
	return wBuf.String(), int(numRead32), nil
}

var _kernel_ReadConsoleW *syscall.Proc

// [ReadFile] function.
//
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hStd HSTD) ReadFile(
	buffer []byte,
	overlapped *OVERLAPPED,
) (numBytesRead int, wErr error) {
	return HFILE(hStd).ReadFile(buffer, overlapped)
}

// [SetConsoleCursorInfo] function.
//
// [SetConsoleCursorInfo]: https://learn.microsoft.com/en-us/windows/console/setconsolecursorinfo
func (hStd HSTD) SetConsoleCursorInfo(info *CONSOLE_CURSOR_INFO) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_SetConsoleCursorInfo, "SetConsoleCursorInfo"),
		uintptr(hStd),
		uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

var _kernel_SetConsoleCursorInfo *syscall.Proc

// [SetConsoleCursorPosition] function.
//
// [SetConsoleCursorPosition]: https://learn.microsoft.com/en-us/windows/console/coord-str
func (hStd HSTD) SetConsoleCursorPosition(x, y int) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_SetConsoleCursorPosition, "SetConsoleCursorPosition"),
		uintptr(hStd),
		uintptr(int16(x)),
		uintptr(int16(y)))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

var _kernel_SetConsoleCursorPosition *syscall.Proc

// [SetConsoleDisplayMode] function.
//
// [SetConsoleDisplayMode]: https://learn.microsoft.com/en-us/windows/console/setconsoledisplaymode
func (hStd HSTD) SetConsoleDisplayMode(mode co.CONSOLE_MODE) (SIZE, error) {
	var coord COORD
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_SetConsoleDisplayMode, "SetConsoleDisplayMode"),
		uintptr(hStd),
		uintptr(mode),
		uintptr(unsafe.Pointer(&coord)))
	if ret == 0 {
		return SIZE{}, co.ERROR(err)
	}
	return SIZE{Cx: int32(coord.X), Cy: int32(coord.Y)}, nil
}

var _kernel_SetConsoleDisplayMode *syscall.Proc

// [SetConsoleMode] function.
//
// [SetConsoleMode]: https://learn.microsoft.com/en-us/windows/console/setconsolemode
func (hStd HSTD) SetConsoleMode(mode co.ENABLE) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_SetConsoleMode, "SetConsoleMode"),
		uintptr(hStd),
		uintptr(mode))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

var _kernel_SetConsoleMode *syscall.Proc

// [SetConsoleScreenBufferSize] function.
//
// [SetConsoleScreenBufferSize]: https://learn.microsoft.com/en-us/windows/console/setconsolescreenbuffersize
func (hStd HSTD) SetConsoleScreenBufferSize(x, y int) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_SetConsoleScreenBufferSize, "SetConsoleScreenBufferSize"),
		uintptr(hStd),
		uintptr(int16(x)),
		uintptr(int16(y)))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

var _kernel_SetConsoleScreenBufferSize *syscall.Proc

// [SetConsoleTextAttribute] function.
//
// [SetConsoleTextAttribute]: https://learn.microsoft.com/en-us/windows/console/setconsoletextattribute
func (hStd HSTD) SetConsoleTextAttribute(attrs co.CHAR_ATTR) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_SetConsoleTextAttribute, "SetConsoleTextAttribute"),
		uintptr(hStd),
		uintptr(attrs))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

var _kernel_SetConsoleTextAttribute *syscall.Proc

// [WriteConsole] function.
//
// [WriteConsole]: https://learn.microsoft.com/en-us/windows/console/writeconsole
func (hStd HSTD) WriteConsole(text string) (numCharsWritten int, wErr error) {
	var wText wstr.BufEncoder
	var numWritten32 uint32

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_kernel_WriteConsoleW, "WriteConsoleW"),
		uintptr(hStd),
		uintptr(wText.AllowEmpty(text)),
		uintptr(uint32(len(text))),
		uintptr(unsafe.Pointer(&numWritten32)),
		0)
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return int(numWritten32), nil
}

var _kernel_WriteConsoleW *syscall.Proc

// [WriteFile] function.
//
// [WriteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hStd HSTD) WriteFile(
	data []byte,
	overlapped *OVERLAPPED,
) (numBytesWritten int, wErr error) {
	return HFILE(hStd).WriteFile(data, overlapped)
}
