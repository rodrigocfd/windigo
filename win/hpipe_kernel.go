//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a pipe.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HPIPE HANDLE

// ⚠️ You must defer HPIPE.CloseHandle().
//
// 📑 https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-createnamedpipew
func CreateNamedPipe(
	name string, dwOpenMode co.PIPE_ACCESS, dwPipeMode co.PIPE,
	nMaxInstances uint, nOutBufferSize uint, nInBufferSize uint,
	nDefaultTimeOut uint, securityAttributes *SECURITY_ATTRIBUTES) (HPIPE, error) {

	ret, _, err := syscall.SyscallN(proc.CreateNamedPipe.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(name))),
		uintptr(dwOpenMode),
		uintptr(dwPipeMode),
		uintptr(nMaxInstances),
		uintptr(nOutBufferSize),
		uintptr(nInBufferSize),
		uintptr(nDefaultTimeOut),
		uintptr(unsafe.Pointer(securityAttributes)))

	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return HPIPE(ret), nil
}

// 📑 https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-connectnamedpipe
func (hPipe HPIPE) ConnectNamedPipe() error {
	ret, _, err := syscall.SyscallN(proc.ConnectNamedPipe.Addr(),
		uintptr(hPipe), uintptr(0))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-disconnectnamedpipe
func (hPipe HPIPE) DisconnectNamedPipe() error {
	ret, _, err := syscall.SyscallN(proc.DisconnectNamedPipe.Addr(),
		uintptr(hPipe), uintptr(0))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hPipe HPIPE) CloseHandle() error {
	return HFILE(hPipe).CloseHandle()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hPipe HPIPE) ReadFile(
	buffer []byte, overlapped *OVERLAPPED) (numBytesRead uint32, e error) {

	return HFILE(hPipe).ReadFile(buffer, overlapped)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hPipe HPIPE) WriteFile(
	data []byte, overlapped *OVERLAPPED) (numBytesWritten uint32, e error) {

	return HFILE(hPipe).WriteFile(data, overlapped)
}
