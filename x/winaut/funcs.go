//go:build windows

package winaut

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/coaut"
)

// [OleLoadPicture] function.
//
// Pass size = 0 to read all the bytes from the stream.
//
// The bytes are copied, so [IStream] can be released after this function
// returns.
//
// The picture must be in the following formats:
//   - BMP (bitmap)
//   - JPEG
//   - WMF (metafile)
//   - ICO (icon)
//   - GIF
//
// If the image format is not supported, returns
// [co.HRESULT_CTL_E_INVALIDPICTURE].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	data := []byte{0x10, 0x11, 0x12}
//	defer runtime.KeepAlive(data)
//
//	stream, _ := win.SHCreateMemStream(rel, data)
//	pic, _ := winaut.OleLoadPicture(rel, stream, 0, true)
//
// [OleLoadPicture]: https://learn.microsoft.com/en-us/windows/win32/api/olectl/nf-olectl-oleloadpicture
func OleLoadPicture(
	releaser *win.OleReleaser,
	stream *win.IStream,
	size int,
	keepOriginalFormat bool,
) (*IPicture, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		dll.Oleaut.Load(&_oleaut_OleLoadPicture, "OleLoadPicture"),
		stream.Ppvt(),
		uintptr(int32(size)),
		utl.BoolToUintptr(!keepOriginalFormat), // note: reversed
		uintptr(unsafe.Pointer(&coaut.IID_IPicture)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IPicture](ret, ppvtQueried, releaser)
}

var _oleaut_OleLoadPicture *syscall.Proc

// [OleLoadPicturePath] function.
//
// The picture must be in the following formats:
//   - BMP (bitmap)
//   - JPEG
//   - WMF (metafile)
//   - ICO (icon)
//   - GIF
//
// If the image format is not supported, returns
// [co.HRESULT_CTL_E_INVALIDPICTURE].
//
// [OleLoadPicturePath]: https://learn.microsoft.com/en-us/windows/win32/api/olectl/nf-olectl-oleloadpicturepath
func OleLoadPicturePath(
	releaser *win.OleReleaser,
	path string,
	transparentColor win.COLORREF,
) (*IPicture, error) {
	var wPath wstr.BufEncoder
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		dll.Oleaut.Load(&_oleaut_OleLoadPicturePath, "OleLoadPicturePath"),
		uintptr(wPath.EmptyIsNil(path)),
		0, 0,
		uintptr(transparentColor),
		uintptr(unsafe.Pointer(&coaut.IID_IPicture)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IPicture](ret, ppvtQueried, releaser)
}

var _oleaut_OleLoadPicturePath *syscall.Proc
