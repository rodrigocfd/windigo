//go:build windows

package windxgi

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGIKeyedMutex] COM interface.
//
// [IDXGIKeyedMutex]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgikeyedmutex
type IDXGIKeyedMutex struct{ IDXGIDeviceSubObject }

type _IDXGIKeyedMutexVt struct {
	_IDXGIDeviceSubObjectVt
	AcquireSync uintptr
	ReleaseSync uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIKeyedMutex) IID() *co.IID {
	return &codxgi.IID_IDXGIKeyedMutex
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIKeyedMutex) AddRef(releaser *win.OleReleaser) *IDXGIKeyedMutex {
	return utl.OleNewFromAddRef[*IDXGIKeyedMutex](me, releaser)
}

// [AcquireSync] method.
//
// [AcquireSync]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgikeyedmutex-acquiresync
func (me *IDXGIKeyedMutex) AcquireSync(key uint64, milliseconds int) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIKeyedMutexVt](me.Ppvt()).AcquireSync,
		me.Ppvt(),
		uintptr(key),
		uintptr(uint32(milliseconds)))
	return utl.HresultToError(ret)
}

// [ReleaseSync] method.
//
// [ReleaseSync]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgikeyedmutex-releasesync
func (me *IDXGIKeyedMutex) ReleaseSync(key uint64) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIKeyedMutexVt](me.Ppvt()).ReleaseSync,
		me.Ppvt(),
		uintptr(key))
	return utl.HresultToError(ret)
}
