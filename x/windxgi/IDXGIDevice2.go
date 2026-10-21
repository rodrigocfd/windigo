//go:build windows

package windxgi

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGIDevice2] COM interface.
//
// [IDXGIDevice2]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nn-dxgi1_2-idxgidevice2
type IDXGIDevice2 struct{ IDXGIDevice1 }

type _IDXGIDevice2Vt struct {
	_IDXGIDevice1Vt
	OfferResources   uintptr
	ReclaimResources uintptr
	EnqueueSetEvent  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIDevice2) IID() *co.IID {
	return &codxgi.IID_IDXGIDevice2
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIDevice2) AddRef(releaser *win.OleReleaser) *IDXGIDevice2 {
	return utl.OleNewFromAddRef[*IDXGIDevice2](me, releaser)
}

// [EnqueueSetEvent] method.
//
// [EnqueueSetEvent]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgidevice2-enqueuesetevent
func (me *IDXGIDevice2) EnqueueSetEvent(hEvent win.HANDLE) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDevice2Vt](me.Ppvt()).EnqueueSetEvent,
		me.Ppvt(),
		uintptr(hEvent))
	return utl.HresultToError(ret)
}

// [OfferResources] method.
//
// [OfferResources]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgidevice2-offerresources
func (me *IDXGIDevice2) OfferResources(
	priority codxgi.DXGI_OFFER_RESOURCE_PRIORITY,
	resources ...*IDXGIResource,
) error {
	resPpvts := make([]uintptr, 0, len(resources))
	for _, r := range resources {
		resPpvts = append(resPpvts, r.Ppvt())
	}

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDevice2Vt](me.Ppvt()).OfferResources,
		me.Ppvt(),
		uintptr(uint32(len(resources))),
		uintptr(unsafe.Pointer(&resPpvts[0])),
		uintptr(priority))
	return utl.HresultToError(ret)
}

// [ReclaimResources] method.
//
// [ReclaimResources]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgidevice2-reclaimresources
func (me *IDXGIDevice2) ReclaimResources(resources ...*IDXGIResource) ([]bool, error) {
	resPpvts := make([]uintptr, 0, len(resources))
	for _, r := range resources {
		resPpvts = append(resPpvts, r.Ppvt())
	}

	boolsBuf := make([]win.BOOL, len(resources))

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDevice2Vt](me.Ppvt()).ReclaimResources,
		me.Ppvt(),
		uintptr(uint32(len(resources))),
		uintptr(unsafe.Pointer(&resPpvts[0])),
		uintptr(unsafe.Pointer(&boolsBuf[0])))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}

	discarded := make([]bool, 0, len(resources))
	for _, b := range boolsBuf {
		discarded = append(discarded, b.Ok())
	}
	return discarded, nil
}
