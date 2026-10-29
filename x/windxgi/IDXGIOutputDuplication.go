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

// [IDXGIOutputDuplication] COM interface.
//
// [IDXGIOutputDuplication]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nn-dxgi1_2-idxgioutputduplication
type IDXGIOutputDuplication struct{ IDXGIObject }

type _IDXGIOutputDuplicationVt struct {
	IDXGIObject
	GetDesc              uintptr
	AcquireNextFrame     uintptr
	GetFrameDirtyRects   uintptr
	GetFrameMoveRects    uintptr
	GetFramePointerShape uintptr
	MapDesktopSurface    uintptr
	UnMapDesktopSurface  uintptr
	ReleaseFrame         uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIOutputDuplication) IID() *co.IID {
	return &codxgi.IID_IDXGIOutputDuplication
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIOutputDuplication) AddRef(releaser *win.OleReleaser) *IDXGIOutputDuplication {
	return utl.OleNewFromAddRef[*IDXGIOutputDuplication](me, releaser)
}

// [AcquireNextFrame] method.
//
// [AcquireNextFrame]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgioutputduplication-acquirenextframe
func (me *IDXGIOutputDuplication) AcquireNextFrame(
	releaser *win.OleReleaser,
	msTimeout int,
) (DXGI_OUTDUPL_FRAME_INFO, *IDXGIResource, error) {
	var frameInfo DXGI_OUTDUPL_FRAME_INFO
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIOutputDuplicationVt](me.Ppvt()).AcquireNextFrame,
		me.Ppvt(),
		uintptr(uint32(msTimeout)),
		uintptr(unsafe.Pointer(&frameInfo)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return DXGI_OUTDUPL_FRAME_INFO{}, nil, hr
	}
	pObj := utl.OleNew[*IDXGIResource](ppvtQueried, releaser)
	return frameInfo, pObj, nil
}

// [GetDesc] method.
//
// [GetDesc]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgioutputduplication-getdesc
func (me *IDXGIOutputDuplication) GetDesc() DXGI_OUTDUPL_DESC {
	var desc DXGI_OUTDUPL_DESC
	_, _, _ = syscall.SyscallN(
		utl.Vt[_IDXGIOutputDuplicationVt](me.Ppvt()).GetDesc,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&desc)))
	return desc
}

// [GetFrameDirtyRects] method.
//
// [GetFrameDirtyRects]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgioutputduplication-getframedirtyrects
func (me *IDXGIOutputDuplication) GetFrameDirtyRects() ([]win.RECT, error) {
	const RECT_SZ = uint32(unsafe.Sizeof(win.RECT{}))
	rectsBuf := make([]win.RECT, 1)
	var reqBytes uint32

	for {
		ret, _, _ := syscall.SyscallN(
			utl.Vt[_IDXGIOutputDuplicationVt](me.Ppvt()).GetFrameDirtyRects,
			me.Ppvt(),
			uintptr(len(rectsBuf))*uintptr(RECT_SZ),
			uintptr(unsafe.Pointer(&rectsBuf[0])),
			uintptr(unsafe.Pointer(&reqBytes)))

		if hr := co.HRESULT(ret); hr == codxgi.HRESULT_DXGI_ERROR_MORE_DATA {
			requiredCount := (reqBytes + RECT_SZ - 1) / RECT_SZ
			rectsBuf = make([]win.RECT, requiredCount) // increase buffer size to try again
		} else if hr != co.HRESULT_S_OK {
			return nil, hr
		} else {
			rectsBuf = make([]win.RECT, reqBytes/RECT_SZ)
			return rectsBuf, nil
		}
	}
}

// [GetFrameMoveRects] method.
//
// [GetFrameMoveRects]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi1_2/nf-dxgi1_2-idxgioutputduplication-getframemoverects
func (me *IDXGIOutputDuplication) GetFrameMoveRects() ([]DXGI_OUTDUPL_MOVE_RECT, error) {
	const MR_SZ = uint32(unsafe.Sizeof(DXGI_OUTDUPL_MOVE_RECT{}))
	mrsBuf := make([]DXGI_OUTDUPL_MOVE_RECT, 1)
	var reqBytes uint32

	for {
		ret, _, _ := syscall.SyscallN(
			utl.Vt[_IDXGIOutputDuplicationVt](me.Ppvt()).GetFrameMoveRects,
			me.Ppvt(),
			uintptr(len(mrsBuf))*uintptr(MR_SZ),
			uintptr(unsafe.Pointer(&mrsBuf[0])),
			uintptr(unsafe.Pointer(&reqBytes)))

		if hr := co.HRESULT(ret); hr == codxgi.HRESULT_DXGI_ERROR_MORE_DATA {
			requiredCount := (reqBytes + MR_SZ - 1) / MR_SZ
			mrsBuf = make([]DXGI_OUTDUPL_MOVE_RECT, requiredCount) // increase buffer size to try again
		} else if hr != co.HRESULT_S_OK {
			return nil, hr
		} else {
			mrsBuf = make([]DXGI_OUTDUPL_MOVE_RECT, reqBytes/MR_SZ)
			return mrsBuf, nil
		}
	}
}
