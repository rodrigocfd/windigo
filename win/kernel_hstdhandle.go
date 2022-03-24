package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a standard device (standard input, standard output, or standard
// error).
//
// 📑 https://docs.microsoft.com/en-us/windows/console/getstdhandle
type HSTDHANDLE HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/console/getstdhandle
func GetStdHandle(handle co.STD) HSTDHANDLE {
	ret, _, err := syscall.Syscall(proc.GetStdHandle.Addr(), 1,
		uintptr(handle), 0, 0)
	if int(ret) == _INVALID_HANDLE_VALUE {
		panic(errco.ERROR(err))
	}
	return HSTDHANDLE(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/console/readconsole
func (hStd HSTDHANDLE) ReadConsole(
	maxChars int,
	inputControl *CONSOLE_READCONSOLE_CONTROL) (string, error) {

	buf := make([]uint16, maxChars+1)
	var numCharsRead uint32

	ret, _, err := syscall.Syscall6(proc.ReadConsole.Addr(), 5,
		uintptr(hStd), uintptr(unsafe.Pointer(&buf[0])), uintptr(maxChars),
		uintptr(unsafe.Pointer(&numCharsRead)),
		uintptr(unsafe.Pointer(inputControl)), 0)
	if ret == 0 {
		return "", errco.ERROR(err)
	}
	return Str.FromNativeSlice(buf), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/console/setconsolemode
func (hStd HSTDHANDLE) SetConsoleMode(mode co.ENABLE) error {
	ret, _, err := syscall.Syscall(proc.SetConsoleMode.Addr(), 2,
		uintptr(hStd), uintptr(mode), 0)
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/console/writeconsole
func (hStd HSTDHANDLE) WriteConsole(text string) (numCharsWritten int, e error) {
	ret, _, err := syscall.Syscall6(proc.WriteConsole.Addr(), 5,
		uintptr(hStd),
		uintptr(unsafe.Pointer(Str.ToNativePtr(text))), uintptr(len(text)),
		uintptr(unsafe.Pointer(&numCharsWritten)), 0, 0)
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return numCharsWritten, nil
}
