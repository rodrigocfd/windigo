//go:build windows

package oleaut

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [IPicture] COM interface.
//
// [IPicture]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nn-ocidl-ipicture
type IPicture struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IPicture) IID() co.IID {
	return co.IID_ITaskbarList
}

// [get_Attributes] method.
//
// [get_Attributes]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_attributes
func (me *IPicture) Attributes() (co.PICATTR, error) {
	var attr co.PICATTR
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Attributes,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&attr)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return attr, nil
	} else {
		return co.PICATTR(0), hr
	}
}

// [get_CurDC] method.
//
// [get_CurDC]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_curdc
func (me *IPicture) CurDC() (win.HDC, error) {
	var hdc win.HDC
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_CurDC,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hdc)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hdc, nil
	} else {
		return win.HDC(0), hr
	}
}

// [get_Handle] method.
//
// [get_Handle]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_handle
func (me *IPicture) Handle() (win.HBITMAP, error) {
	var hBmp win.HBITMAP
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Handle,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hBmp)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hBmp, nil
	} else {
		return win.HBITMAP(0), hr
	}
}

// [get_Height] method.
//
// If you need both width and height, call [IPicture.Size], which returns both.
//
// Note that this method returns the height in HIMETRIC units. To convert it to
// pixels, use [win.HDC.HiMetricToPixel], or simply call [IPicture.SizePixels]
// method, which already performs the conversion.
//
// [get_Height]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_height
func (me *IPicture) Height() (int, error) {
	var cy int32
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Height,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cy)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(cy), nil
	} else {
		return 0, hr
	}
}

// [get_hPal] method.
//
// [get_hPal]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_hpal
func (me *IPicture) HPal() (win.HPALETTE, error) {
	var hPal win.HPALETTE
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_hPal,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hPal)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hPal, nil
	} else {
		return win.HPALETTE(0), hr
	}
}

// [get_KeepOriginalFormat] method.
//
// [get_KeepOriginalFormat]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_keeporiginalformat
func (me *IPicture) KeepOriginalFormat() (bool, error) {
	var keep int32 // BOOL
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_KeepOriginalFormat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&keep)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return keep != 0, nil
	} else {
		return false, hr
	}
}

// [PictureChanged] method.
//
// [PictureChanged]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-picturechanged
func (me *IPicture) PictureChanged() error {
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).PictureChanged,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [Render] method.
//
// # Example
//
//	var wnd ui.Main // initialized somewhere
//	var pic ole.IPicture
//
//	wnd.On().WmPaint(func() {
//		var ps win.PAINTSTRUCT
//		hdc, _ := wnd.Hwnd().BeginPaint(&ps)
//		defer wnd.Hwnd().EndPaint(&ps)
//
//		sz, _ := pic.Size()
//		pic.Render(hdc,
//			win.POINT{},
//			win.SIZE{Cx: ps.RcPaint.Right, Cy: ps.RcPaint.Bottom},
//			win.POINT{X: 0, Y: sz.Cy},
//			win.SIZE{Cx: sz.Cx, Cy: -sz.Cy},
//		)
//	})
//
// [Render]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-render
func (me *IPicture) Render(
	hdc win.HDC,
	destOffset win.POINT,
	destSz win.SIZE,
	srcOffset win.POINT,
	srcSz win.SIZE,
) (metafileBounds win.RECT, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Render,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hdc),
		uintptr(destOffset.X), uintptr(destOffset.Y),
		uintptr(destSz.Cx), uintptr(destSz.Cy),
		uintptr(srcOffset.X), uintptr(srcOffset.Y),
		uintptr(srcSz.Cx), uintptr(srcSz.Cy),
		uintptr(unsafe.Pointer(&metafileBounds)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		hr = nil
	} else {
		metafileBounds = win.RECT{}
	}
	return
}

// [SaveAsFile] method.
//
// [SaveAsFile]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-saveasfile
func (me *IPicture) SaveAsFile(stream *ole.IStream, saveCopy bool) (numBytesWritten uint, hr error) {
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).SaveAsFile,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		utl.BoolToUintptr(saveCopy),
		uintptr(unsafe.Pointer(&numBytesWritten)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		hr = nil
	} else {
		numBytesWritten = 0
	}
	return
}

// [SelectPicture] method.
//
// [SelectPicture]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-selectpicture
func (me *IPicture) SelectPicture(hdc win.HDC) (win.HDC, win.HBITMAP, error) {
	var hdcOut win.HDC
	var hBmp win.HBITMAP

	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).SelectPicture,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&hdcOut)),
		uintptr(unsafe.Pointer(&hBmp)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return hdcOut, hBmp, nil
	} else {
		return win.HDC(0), win.HBITMAP(0), hr
	}
}

// [put_KeepOriginalFormat] method.
//
// [put_KeepOriginalFormat]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-put_keeporiginalformat
func (me *IPicture) SetKeepOriginalFormat(keep bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Put_KeepOriginalFormat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(keep))
	return utl.ErrorAsHResult(ret)
}

// [set_hPal] method.
//
// [set_hPal]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-set_hpal
func (me *IPicture) SetHPal(hPal win.HPALETTE) error {
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Set_hPal,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(hPal))
	return utl.ErrorAsHResult(ret)
}

// Calls [IPicture.Width] and [IPicture.Height] at once.
//
// If you need both width and height, call [IPicture.Size], which returns both.
//
// Note that this method returns the size in HIMETRIC units. To convert it to
// pixels, use [win.HDC.HiMetricToPixel], or simply call [IPicture.SizePixels]
// method, which already performs the conversion.
//
// [IPicture.Width]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_width
// [IPicture.Height]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_height
func (me *IPicture) Size() (win.SIZE, error) {
	width, err := me.Width()
	if err != nil {
		return win.SIZE{}, err
	}

	height, err := me.Height()
	if err != nil {
		return win.SIZE{}, err
	}

	return win.SIZE{Cx: int32(width), Cy: int32(height)}, nil
}

// Calls [IPicture.Width] and [IPicture.Height], then convers from HIMETRIC
// units to pixels with [win.HDC.HiMetricToPixel].
//
// If hdc is zero, the method will retrieve the HDC for the whole screen with
// [win.HWND.GetDC].
//
// # Example
//
//	hdcScreen, _ := win.HWND(0).GetDC()
//	defer win.HWND(0).ReleaseDC(hdcScreen)
//
//	sz, _ := pic.SizePixels(hdcScreen)
func (me *IPicture) SizePixels(hdc win.HDC) (win.SIZE, error) {
	myHdc := hdc
	if myHdc == 0 {
		myHdc, err := win.HWND(0).GetDC() // DC of the entire screen
		if err != nil {
			return win.SIZE{}, err
		}
		defer win.HWND(0).ReleaseDC(myHdc)
	}

	himetricX, err := me.Width()
	if err != nil {
		return win.SIZE{}, err
	}
	himetricY, err := me.Height()
	if err != nil {
		return win.SIZE{}, err
	}

	pixelX, pixelY := myHdc.HiMetricToPixel(himetricX, himetricY)
	return win.SIZE{Cx: int32(pixelX), Cy: int32(pixelY)}, nil
}

// [get_Type] method.
//
// [get_Type]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_type
func (me *IPicture) Type() (co.PICTYPE, error) {
	var picty co.PICTYPE
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Type,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&picty)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return picty, nil
	} else {
		return co.PICTYPE(0), hr
	}
}

// [get_Width] method.
//
// Note that this method returns the width in HIMETRIC units. To convert it to
// pixels, use [win.HDC.HiMetricToPixel], or simply call [IPicture.SizePixels]
// method, which already performs the conversion.
//
// [get_Width]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_width
func (me *IPicture) Width() (int, error) {
	var cx int32
	ret, _, _ := syscall.SyscallN(
		(*_IPictureVt)(unsafe.Pointer(*me.Ppvt())).Get_Width,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cx)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(cx), nil
	} else {
		return 0, hr
	}
}

type _IPictureVt struct {
	ole.IUnknownVt
	Get_Handle             uintptr
	Get_hPal               uintptr
	Get_Type               uintptr
	Get_Width              uintptr
	Get_Height             uintptr
	Render                 uintptr
	Set_hPal               uintptr
	Get_CurDC              uintptr
	SelectPicture          uintptr
	Get_KeepOriginalFormat uintptr
	Put_KeepOriginalFormat uintptr
	PictureChanged         uintptr
	SaveAsFile             uintptr
	Get_Attributes         uintptr
}
