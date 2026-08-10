//go:build windows

package winaut

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/coaut"
)

// [IPicture] COM interface.
//
// [IPicture]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nn-ocidl-ipicture
type IPicture struct{ win.IUnknown }

type _IPictureVt struct {
	utl.IUnknownVt
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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IPicture) IID() *co.IID {
	return &coaut.IID_IPicture
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IPicture) AddRef(releaser *win.OleReleaser) *IPicture {
	return utl.OleNewFromAddRef[*IPicture](me, releaser)
}

// [get_Attributes] method.
//
// [get_Attributes]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_attributes
func (me *IPicture) GetAttributes() (coaut.PICATTR, error) {
	return utl.OleCallReturnStruct[coaut.PICATTR](me,
		utl.Vt[_IPictureVt](me.Ppvt()).Get_Attributes)
}

// [get_CurDC] method.
//
// [get_CurDC]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_curdc
func (me *IPicture) GetCurDC() (win.HDC, error) {
	return utl.OleCallReturnStruct[win.HDC](me,
		utl.Vt[_IPictureVt](me.Ppvt()).Get_CurDC)
}

// [get_Handle] method.
//
// [get_Handle]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_handle
func (me *IPicture) GetHandle() (win.HBITMAP, error) {
	return utl.OleCallReturnStruct[win.HBITMAP](me,
		utl.Vt[_IPictureVt](me.Ppvt()).Get_Handle)
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
func (me *IPicture) GetHeight() (int, error) {
	var cy int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPictureVt](me.Ppvt()).Get_Height,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&cy)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(cy), nil
}

// [get_hPal] method.
//
// [get_hPal]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_hpal
func (me *IPicture) GetHPal() (win.HPALETTE, error) {
	return utl.OleCallReturnStruct[win.HPALETTE](me,
		utl.Vt[_IPictureVt](me.Ppvt()).Get_hPal)
}

// [get_KeepOriginalFormat] method.
//
// [get_KeepOriginalFormat]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_keeporiginalformat
func (me *IPicture) GetKeepOriginalFormat() (bool, error) {
	var keep win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPictureVt](me.Ppvt()).Get_KeepOriginalFormat,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&keep)))
	return utl.HresultToBoolError(int32(keep), ret)
}

// [get_Type] method.
//
// [get_Type]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_type
func (me *IPicture) GetType() (coaut.PICTYPE, error) {
	return utl.OleCallReturnStruct[coaut.PICTYPE](me,
		utl.Vt[_IPictureVt](me.Ppvt()).Get_Type)
}

// [get_Width] method.
//
// Note that this method returns the width in HIMETRIC units. To convert it to
// pixels, use [win.HDC.HiMetricToPixel], or simply call [IPicture.SizePixels]
// method, which already performs the conversion.
//
// [get_Width]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_width
func (me *IPicture) GetWidth() (int, error) {
	var cx int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPictureVt](me.Ppvt()).Get_Width,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&cx)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(cx), nil
}

// [PictureChanged] method.
//
// [PictureChanged]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-picturechanged
func (me *IPicture) PictureChanged() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IPictureVt](me.Ppvt()).PictureChanged)
}

// [put_KeepOriginalFormat] method.
//
// [put_KeepOriginalFormat]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-put_keeporiginalformat
func (me *IPicture) PutKeepOriginalFormat(keep bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPictureVt](me.Ppvt()).Put_KeepOriginalFormat,
		me.Ppvt(),
		utl.BoolToUintptr(keep))
	return utl.HresultToError(ret)
}

// [Render] method.
//
// Example:
//
//	var wnd *ui.Main // initialized somewhere
//	var pic *winaut.IPicture
//
//	wnd.On().WmPaint(func() {
//		var ps win.PAINTSTRUCT
//		hdc, _ := wnd.Hwnd().BeginPaint(&ps)
//		defer wnd.Hwnd().EndPaint(&ps)
//
//		sz, _ := pic.Size()
//		_, _ = pic.Render(hdc,
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
		utl.Vt[_IPictureVt](me.Ppvt()).Render,
		me.Ppvt(),
		uintptr(hdc),
		uintptr(destOffset.X),
		uintptr(destOffset.Y),
		uintptr(destSz.Cx),
		uintptr(destSz.Cy),
		uintptr(srcOffset.X),
		uintptr(srcOffset.Y),
		uintptr(srcSz.Cx),
		uintptr(srcSz.Cy),
		uintptr(unsafe.Pointer(&metafileBounds)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return win.RECT{}, hr
	}
	return metafileBounds, nil
}

// [SaveAsFile] method.
//
// [SaveAsFile]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-saveasfile
func (me *IPicture) SaveAsFile(stream *win.IStream, saveCopy bool) (numBytesWritten int, hr error) {
	var written32 int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPictureVt](me.Ppvt()).SaveAsFile,
		me.Ppvt(),
		stream.Ppvt(),
		utl.BoolToUintptr(saveCopy),
		uintptr(unsafe.Pointer(&written32)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(written32), nil
}

// [SelectPicture] method.
//
// [SelectPicture]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-selectpicture
func (me *IPicture) SelectPicture(hdc win.HDC) (win.HDC, win.HBITMAP, error) {
	var hdcOut win.HDC
	var hBmp win.HBITMAP

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPictureVt](me.Ppvt()).SelectPicture,
		me.Ppvt(),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&hdcOut)),
		uintptr(unsafe.Pointer(&hBmp)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return win.HDC(0), win.HBITMAP(0), hr
	}
	return hdcOut, hBmp, nil
}

// [set_hPal] method.
//
// [set_hPal]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-set_hpal
func (me *IPicture) SetHPal(hPal win.HPALETTE) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPictureVt](me.Ppvt()).Set_hPal,
		me.Ppvt(),
		uintptr(hPal))
	return utl.HresultToError(ret)
}

// Calls [IPicture.GetWidth] and [IPicture.GetHeight] at once.
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
	width, err := me.GetWidth()
	if err != nil {
		return win.SIZE{}, err
	}

	height, err := me.GetHeight()
	if err != nil {
		return win.SIZE{}, err
	}

	return win.SIZE{Cx: int32(width), Cy: int32(height)}, nil
}

// Calls [IPicture.GetWidth] and [IPicture.GetHeight], then converts from
// HIMETRIC units to pixels with [win.HDC.HiMetricToPixel].
//
// If hdc is zero, the method will retrieve the HDC for the whole screen with
// [win.HWND.GetDC].
//
// Example:
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

	himetricX, err := me.GetWidth()
	if err != nil {
		return win.SIZE{}, err
	}
	himetricY, err := me.GetHeight()
	if err != nil {
		return win.SIZE{}, err
	}

	pixelX, pixelY := myHdc.HiMetricToPixel(himetricX, himetricY)
	return win.SIZE{Cx: int32(pixelX), Cy: int32(pixelY)}, nil
}
