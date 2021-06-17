package errco

import (
	"fmt"
	"syscall"
)

// Returned by GetLastError(), also an HRESULT.
//
// We can't simply use syscall.Errno because it's an uintptr (8 bytes), thus a
// native struct with such a field type would be wrong. However, the Unwrap()
// method will return the syscall.Errno value.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/debug/system-error-codes
type ERROR uint32

// Calls FormatMessage() and returns a full error description.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-formatmessagew
func (err ERROR) Error() string {
	return fmt.Sprintf("[%d 0x%02x] %s",
		uint32(err), uint32(err), err.Unwrap().Error())
}

// Returns the contained syscall.Errno.
func (err ERROR) Unwrap() error {
	return syscall.Errno(err)
}
