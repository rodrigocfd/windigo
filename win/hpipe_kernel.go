//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a [pipe].
//
// [pipe]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HPIPE HANDLE

// [CreateNamedPipe] function.
//
// ⚠️ You must defer HPIPE.CloseHandle().
//
// [CreateNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-createnamedpipew
func CreateNamedPipe(
	name string,
	dwOpenMode co.PIPE_ACCESS,
	dwPipeMode co.PIPE,
	nMaxInstances uint,
	nOutBufferSize uint,
	nInBufferSize uint,
	nDefaultTimeOut uint,
	securityAttributes *SECURITY_ATTRIBUTES) (HPIPE, error) {

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

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hPipe HPIPE) CloseHandle() error {
	return HFILE(hPipe).CloseHandle()
}

// [ConnectNamedPipe] function.
//
// [ConnectNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-connectnamedpipe
func (hPipe HPIPE) ConnectNamedPipe() error {
	ret, _, err := syscall.SyscallN(proc.ConnectNamedPipe.Addr(),
		uintptr(hPipe), uintptr(0))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [DisconnectNamedPipe] function.
//
// [DisconnectNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-disconnectnamedpipe
func (hPipe HPIPE) DisconnectNamedPipe() error {
	ret, _, err := syscall.SyscallN(proc.DisconnectNamedPipe.Addr(),
		uintptr(hPipe), uintptr(0))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

type _HpipeInfo struct {
	Flags     co.PIPE
	OutBuffer uint32
	InBuffer  uint32
	MaxInsts  uint32
}

// [GetNamedPipeInfo] function.
//
// [GetNamedPipeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-getnamedpipeinfo
func (hPipe HPIPE) GetNamedPipeInfo() (_HpipeInfo, error) {
	var info _HpipeInfo
	ret, _, err := syscall.SyscallN(proc.GetNamedPipeInfo.Addr(),
		uintptr(hPipe), uintptr(unsafe.Pointer(&info.Flags)),
		uintptr(unsafe.Pointer(&info.OutBuffer)),
		uintptr(unsafe.Pointer(&info.InBuffer)),
		uintptr(unsafe.Pointer(&info.MaxInsts)))
	if ret == 0 {
		return _HpipeInfo{}, errco.ERROR(err)
	}
	return info, nil
}

type _HpipePeek struct {
	Read      uint32
	Available uint32
	Left      uint32
}

// [PeekNamedPipe] function.
//
// [PeekNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-peeknamedpipe
func (hPipe HPIPE) PeekNamedPipe(buffer []byte) (_HpipePeek, error) {
	var info _HpipePeek
	ret, _, err := syscall.SyscallN(proc.PeekNamedPipe.Addr(),
		uintptr(hPipe), uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&info.Read)),
		uintptr(unsafe.Pointer(&info.Available)),
		uintptr(unsafe.Pointer(&info.Left)))
	if ret == 0 {
		return _HpipePeek{}, errco.ERROR(err)
	}
	return info, nil
}

// [ReadFile] function.
//
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hPipe HPIPE) ReadFile(
	buffer []byte, overlapped *OVERLAPPED) (numBytesRead uint32, e error) {

	return HFILE(hPipe).ReadFile(buffer, overlapped)
}

// [WriteFile] function.
//
// [WriteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hPipe HPIPE) WriteFile(
	data []byte, overlapped *OVERLAPPED) (numBytesWritten uint32, e error) {

	return HFILE(hPipe).WriteFile(data, overlapped)
}
