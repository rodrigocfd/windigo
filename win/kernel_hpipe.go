//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/wstr"
)

// A handle to a [pipe].
//
// [pipe]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#handle
type HPIPE HANDLE

// [CreateNamedPipe] function.
//
// Panics if nMaxInstances, nOutBufferSize, nInBufferSize or nDefaultTimeOut is
// negative.
//
// ⚠️ You must defer HPIPE.CloseHandle().
//
// [CreateNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-createnamedpipew
func CreateNamedPipe(
	name string,
	dwOpenMode co.PIPE_ACCESS,
	dwPipeMode co.PIPE,
	nMaxInstances int,
	nOutBufferSize int,
	nInBufferSize int,
	nDefaultTimeOut int,
	securityAttributes *SECURITY_ATTRIBUTES,
) (HPIPE, error) {
	var wName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_CreateNamedPipeW, "CreateNamedPipeW"),
		uintptr(wName.AllowEmpty(name)),
		uintptr(dwOpenMode),
		uintptr(dwPipeMode),
		uintptr(uint32(nMaxInstances)),
		uintptr(uint32(nOutBufferSize)),
		uintptr(uint32(nInBufferSize)),
		uintptr(uint32(nDefaultTimeOut)),
		uintptr(unsafe.Pointer(securityAttributes)))
	if ret == 0 {
		return HPIPE(0), co.ERROR(err)
	}
	return HPIPE(ret), nil
}

var _kernel_CreateNamedPipeW *syscall.Proc

// [CreatePipe] function.
//
// [CreatePipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-createpipe
func CreatePipe(
	securityAttributes *SECURITY_ATTRIBUTES,
	size int,
) (hPipeRead, hPipeWrite HPIPE, wErr error) {
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_CreatePipe, "CreatePipe"),
		uintptr(unsafe.Pointer(&hPipeRead)),
		uintptr(unsafe.Pointer(&hPipeWrite)),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(uint32(size)))
	if ret == 0 {
		return HPIPE(0), HPIPE(0), co.ERROR(err)
	}
	return hPipeRead, hPipeWrite, nil
}

var _kernel_CreatePipe *syscall.Proc

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hPipe HPIPE) CloseHandle() error {
	return HANDLE(hPipe).CloseHandle()
}

// [ConnectNamedPipe] function.
//
// [ConnectNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-connectnamedpipe
func (hPipe HPIPE) ConnectNamedPipe() error {
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_ConnectNamedPipe, "ConnectNamedPipe"),
		uintptr(hPipe),
		uintptr(0))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

var _kernel_ConnectNamedPipe *syscall.Proc

// [DisconnectNamedPipe] function.
//
// [DisconnectNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-disconnectnamedpipe
func (hPipe HPIPE) DisconnectNamedPipe() error {
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_DisconnectNamedPipe, "DisconnectNamedPipe"),
		uintptr(hPipe),
		uintptr(0))
	if ret == 0 {
		return co.ERROR(err)
	}
	return nil
}

var _kernel_DisconnectNamedPipe *syscall.Proc

// [GetNamedPipeInfo] function.
//
// [GetNamedPipeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-getnamedpipeinfo
func (hPipe HPIPE) GetNamedPipeInfo() (HpipeInfo, error) {
	var info HpipeInfo
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_GetNamedPipeInfo, "GetNamedPipeInfo"),
		uintptr(hPipe),
		uintptr(unsafe.Pointer(&info.Flags)),
		uintptr(unsafe.Pointer(&info.OutBuffer)),
		uintptr(unsafe.Pointer(&info.InBuffer)),
		uintptr(unsafe.Pointer(&info.MaxInsts)))
	if ret == 0 {
		return HpipeInfo{}, co.ERROR(err)
	}
	return info, nil
}

var _kernel_GetNamedPipeInfo *syscall.Proc

// Returned by [HPIPE.GetNamedPipeInfo].
type HpipeInfo struct {
	Flags     co.PIPE
	OutBuffer uint32
	InBuffer  uint32
	MaxInsts  uint32
}

// [PeekNamedPipe] function.
//
// [PeekNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-peeknamedpipe
func (hPipe HPIPE) PeekNamedPipe(buffer []byte) (HpipePeek, error) {
	var info HpipePeek
	ret, _, err := syscall.SyscallN(
		dll.Kernel.Load(&_kernel_PeekNamedPipe, "PeekNamedPipe"),
		uintptr(hPipe),
		uintptr(unsafe.Pointer(unsafe.SliceData(buffer))),
		uintptr(uint32(len(buffer))),
		uintptr(unsafe.Pointer(&info.Read)),
		uintptr(unsafe.Pointer(&info.Available)),
		uintptr(unsafe.Pointer(&info.Left)))
	if ret == 0 {
		return HpipePeek{}, co.ERROR(err)
	}
	return info, nil
}

var _kernel_PeekNamedPipe *syscall.Proc

// Returned by [HPIPE.PeekNamedPipe].
type HpipePeek struct {
	Read      uint32
	Available uint32
	Left      uint32
}

// [ReadFile] function.
//
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (hPipe HPIPE) ReadFile(
	buffer []byte,
	overlapped *OVERLAPPED,
) (numBytesRead int, wErr error) {
	return HFILE(hPipe).ReadFile(buffer, overlapped)
}

// [SetHandleInformation] function.
//
// [SetHandleInformation]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-sethandleinformation
func (hPipe HPIPE) SetHandleInformation(mask, flags co.HANDLE_FLAG) error {
	return HFILE(hPipe).SetHandleInformation(mask, flags)
}

// [WriteFile] function.
//
// [WriteFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-writefile
func (hPipe HPIPE) WriteFile(
	data []byte,
	overlapped *OVERLAPPED,
) (numBytesWritten int, wErr error) {
	return HFILE(hPipe).WriteFile(data, overlapped)
}
