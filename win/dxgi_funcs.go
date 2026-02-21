//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
)

// [CreateDXGIFactory] function.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	factory, _ := win.CreateDXGIFactory(rel)
//
// [CreateDXGIFactory]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-createdxgifactory
func CreateDXGIFactory(releaser *OleReleaser) (*IDXGIFactory, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		dll.Dxgi.Load(&_dxgi_CreateDXGIFactory, "CreateDXGIFactory"),
		uintptr(unsafe.Pointer(&co.IID_IDXGIFactory)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*IDXGIFactory](ret, ppvtQueried, releaser)
}

var _dxgi_CreateDXGIFactory *syscall.Proc
