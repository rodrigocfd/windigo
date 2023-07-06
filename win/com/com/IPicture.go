//go:build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com/comco"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [IPicture] COM interface.
//
// [IPicture]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nn-ocidl-ipicture
type IPicture interface {
	IUnknown

	// [Attributes] COM method.
	//
	// [Attributes]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_attributes
	Attributes() comco.PICATTR

	// [CurDC] COM method.
	//
	// [CurDC]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_curdc
	CurDC() win.HDC

	// [Height] COM method.
	//
	// Note that this method returns the height in HIMETRIC units. To convert it
	// to pixels, use HDC.HiMetricToPixel(), or simply call
	// IPicture.SizePixels() method, which already performs the conversion.
	//
	// [Height]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_height
	Height() int32

	// [KeepOriginalFormat] COM method.
	//
	// [KeepOriginalFormat]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_keeporiginalformat
	KeepOriginalFormat() bool

	// [PictureChanged] COM method.
	//
	// [PictureChanged]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-picturechanged
	PictureChanged()

	// [Render] COM method.
	//
	// [Render]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-render
	Render(hdc win.HDC, destOffset win.POINT, destSz win.SIZE,
		srcOffset win.POINT, srcSz win.SIZE) (metafileBounds win.RECT)

	// [SaveAsFile] COM method.
	//
	// [SaveAsFile]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-saveasfile
	SaveAsFile(stream IStream, saveCopy bool) (numBytesWritten int)

	// [SelectPicture] COM method.
	//
	// [SelectPicture]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-selectpicture
	SelectPicture(hdc win.HDC) (win.HDC, win.HBITMAP)

	// [SetKeepOriginalFormat] COM method.
	//
	// [SetKeepOriginalFormat]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-put_keeporiginalformat
	SetKeepOriginalFormat(keep bool)

	// This helper method calls IPicture.Width() and IPicture.Height(), then
	// convers from HIMETRIC units to pixels with HDC.HiMetricToPixel().
	//
	// If hdc is zero, calls win.HWND(0).GetDC() to retrieve the DC for the
	// entire screen.
	SizePixels(hdc win.HDC) win.SIZE

	// [Type] COM method.
	//
	// [Type]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_type
	Type() comco.PICTYPE

	// [Width] COM method.
	//
	// Note that this method returns the width in HIMETRIC units. To convert it
	// to pixels, use HDC.HiMetricToPixel(), or simply call
	// IPicture.SizePixels() method, which already performs the conversion.
	//
	// [Width]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nf-ocidl-ipicture-get_width
	Width() int32
}

type _IPicture struct{ IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IPicture.Release().
func NewIPicture(base IUnknown) IPicture {
	return &_IPicture{IUnknown: base}
}

// [OleLoadPicture] function.
//
// Pass size = 0 to read all the bytes from the stream.
//
// The bytes are copied, so IStream can be released after this function returns.
//
// ⚠️ You must defer IPicture.Release().
//
// # Example
//
//	data := []byte{0x10, 0x11, 0x12}
//	defer runtime.KeepAlive(data)
//
//	stream := SHCreateMemStream(data)
//	defer stream.Release()
//
//	pic := OleLoadPicture(stream, 0, true)
//	defer pic.Release()
//
// [OleLoadPicture]: https://learn.microsoft.com/en-us/windows/win32/api/olectl/nf-olectl-oleloadpicture
func OleLoadPicture(
	stream IStream, size uint32, keepOriginalFormat bool) IPicture {

	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(proc.OleLoadPicture.Addr(),
		uintptr(unsafe.Pointer(stream.Ptr())),
		uintptr(size),
		util.BoolToUintptr(!keepOriginalFormat), // note: reversed
		uintptr(unsafe.Pointer(win.GuidFromIid(comco.IID_IPicture))),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPicture(NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

// [OleLoadPicturePath] function.
//
// The picture must be in BMP (bitmap), JPEG, WMF (metafile), ICO (icon), or GIF
// format.
//
// ⚠️ You must defer IPicture.Release().
//
// [OleLoadPicturePath]: https://learn.microsoft.com/en-us/windows/win32/api/olectl/nf-olectl-oleloadpicturepath
func OleLoadPicturePath(path string, transparentColor win.COLORREF) IPicture {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(proc.OleLoadPicturePath.Addr(),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(path))),
		0, 0, uintptr(transparentColor),
		uintptr(unsafe.Pointer(win.GuidFromIid(comco.IID_IPicture))),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPicture(NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

func (me *_IPicture) Attributes() comco.PICATTR {
	var attr comco.PICATTR
	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).Get_Attributes,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&attr)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return attr
	} else {
		panic(hr)
	}
}

func (me *_IPicture) CurDC() win.HDC {
	var hdc win.HDC
	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).Get_CurDC,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&hdc)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return hdc
	} else {
		panic(hr)
	}
}

func (me *_IPicture) Height() int32 {
	var cy int32
	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).Get_Height,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&cy)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return cy
	} else {
		panic(hr)
	}
}

func (me *_IPicture) KeepOriginalFormat() bool {
	var keep int32 // BOOL
	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).Get_KeepOriginalFormat,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&keep)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return keep != 0
	} else {
		panic(hr)
	}
}

func (me *_IPicture) PictureChanged() {
	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).PictureChanged,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPicture) Render(
	hdc win.HDC,
	destOffset win.POINT, destSz win.SIZE,
	srcOffset win.POINT, srcSz win.SIZE) (metafileBounds win.RECT) {

	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).Render,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hdc),
		uintptr(destOffset.X), uintptr(destOffset.Y),
		uintptr(destSz.Cx), uintptr(destSz.Cy),
		uintptr(srcOffset.X), uintptr(srcOffset.Y),
		uintptr(srcSz.Cx), uintptr(srcSz.Cy),
		uintptr(unsafe.Pointer(&metafileBounds)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_IPicture) SaveAsFile(
	stream IStream, saveCopy bool) (numBytesWritten int) {

	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).SaveAsFile,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(stream.Ptr())),
		util.BoolToUintptr(saveCopy),
		uintptr(unsafe.Pointer(&numBytesWritten)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return
	} else {
		panic(hr)
	}
}

func (me *_IPicture) SelectPicture(hdc win.HDC) (win.HDC, win.HBITMAP) {
	var hdcOut win.HDC
	var hBmp win.HBITMAP

	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).SelectPicture,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hdc),
		uintptr(unsafe.Pointer(&hdcOut)),
		uintptr(unsafe.Pointer(&hBmp)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return hdcOut, hBmp
	} else {
		panic(hr)
	}
}

func (me *_IPicture) SetKeepOriginalFormat(keep bool) {
	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).Put_KeepOriginalFormat,
		uintptr(unsafe.Pointer(me.Ptr())),
		util.BoolToUintptr(keep))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPicture) SizePixels(hdc win.HDC) win.SIZE {
	myHdc := hdc
	if myHdc == win.HDC(0) {
		myHdc = win.HWND(0).GetDC() // DC of the entire screen
		defer win.HWND(0).ReleaseDC(myHdc)
	}

	himetricX, himetricY := me.Width(), me.Height()
	pixelX, pixelY := myHdc.HiMetricToPixel(himetricX, himetricY)
	return win.SIZE{Cx: pixelX, Cy: pixelY}
}

func (me *_IPicture) Type() comco.PICTYPE {
	var picty comco.PICTYPE
	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).Get_Type,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&picty)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return picty
	} else {
		panic(hr)
	}
}

func (me *_IPicture) Width() int32 {
	var cx int32
	ret, _, _ := syscall.SyscallN(
		(*comvt.IPicture)(unsafe.Pointer(*me.Ptr())).Get_Width,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&cx)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return cx
	} else {
		panic(hr)
	}
}
