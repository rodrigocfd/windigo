//go:build windows

package windxgi

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [CreateDXGIFactory] function.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	factory, _ := windxgi.CreateDXGIFactory(rel)
//
// [CreateDXGIFactory]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-createdxgifactory
func CreateDXGIFactory(releaser *win.OleReleaser) (*IDXGIFactory, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		dll.Dxgi.Load(&_dxgi_CreateDXGIFactory, "CreateDXGIFactory"),
		uintptr(unsafe.Pointer(&codxgi.IID_IDXGIFactory)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IDXGIFactory](ret, ppvtQueried, releaser)
}

var _dxgi_CreateDXGIFactory *syscall.Proc

// [CreateDXGIFactory1] function.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	factory, _ := windxgi.CreateDXGIFactory1(rel)
//
// [CreateDXGIFactory1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-createdxgifactory1
func CreateDXGIFactory1(releaser *win.OleReleaser) (*IDXGIFactory1, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		dll.Dxgi.Load(&_dxgi_CreateDXGIFactory1, "CreateDXGIFactory1"),
		uintptr(unsafe.Pointer(&codxgi.IID_IDXGIFactory1)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IDXGIFactory1](ret, ppvtQueried, releaser)
}

var _dxgi_CreateDXGIFactory1 *syscall.Proc
