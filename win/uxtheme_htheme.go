//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// Handle to a [theme].
//
// [theme]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/
type HTHEME HANDLE

// [CloseThemeData] function.
//
// Paired with [OpenThemeData].
//
// [CloseThemeData]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-closethemedata
// [OpenThemeData]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-openthemedata
func (hTheme HTHEME) CloseThemeData() {
	syscall.SyscallN(_CloseThemeData.Addr(),
		uintptr(hTheme))
}

var _CloseThemeData = dll.Uxtheme.NewProc("CloseThemeData")

// [DrawThemeBackground] function.
//
// [DrawThemeBackground]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-drawthemebackground
func (hTheme HTHEME) DrawThemeBackground(hdc HDC, partStateId co.VS, rc *RECT, clipRc *RECT) error {
	ret, _, _ := syscall.SyscallN(_DrawThemeBackground.Addr(),
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(clipRc)))
	return wutil.ErrorAsHResult(ret)
}

var _DrawThemeBackground = dll.Uxtheme.NewProc("DrawThemeBackground")

// [GetThemeColor] function.
//
// [GetThemeColor]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemecolor
func (hTheme HTHEME) GetThemeColor(partStateId co.VS, propId co.TMT) (COLORREF, error) {
	var color COLORREF
	ret, _, _ := syscall.SyscallN(_GetThemeColor.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&color)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return COLORREF(0), hr
	}
	return color, nil
}

var _GetThemeColor = dll.Uxtheme.NewProc("GetThemeColor")

// [GetThemeInt] function.
//
// [GetThemeInt]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemeint
func (hTheme HTHEME) GetThemeInt(partStateId co.VS, propId co.TMT) (int, error) {
	var intVal int32
	ret, _, _ := syscall.SyscallN(_GetThemeInt.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&intVal)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(intVal), nil
}

var _GetThemeInt = dll.Uxtheme.NewProc("GetThemeInt")

// [GetThemeMetric] function.
//
// [GetThemeMetric]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthememetric
func (hTheme HTHEME) GetThemeMetric(hdc HDC, partStateId co.VS, propId co.TMT) (int, error) {
	var intVal int32
	ret, _, _ := syscall.SyscallN(_GetThemeMetric.Addr(),
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&intVal)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(intVal), nil
}

var _GetThemeMetric = dll.Uxtheme.NewProc("GetThemeMetric")

// [GetThemePosition] function.
//
// [GetThemePosition]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemeposition
func (hTheme HTHEME) GetThemePosition(partStateId co.VS, propId co.TMT) (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(_GetThemePosition.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&pt)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return POINT{}, hr
	}
	return pt, nil
}

var _GetThemePosition = dll.Uxtheme.NewProc("GetThemePosition")

// [GetThemePropertyOrigin] function.
//
// [GetThemePropertyOrigin]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemepropertyorigin
func (hTheme HTHEME) GetThemePropertyOrigin(partStateId co.VS, propId co.TMT) (co.PROPERTYORIGIN, error) {
	var origin co.PROPERTYORIGIN
	ret, _, _ := syscall.SyscallN(_GetThemePropertyOrigin.Addr(),
		uintptr(hTheme),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&origin)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return co.PROPERTYORIGIN(0), hr
	}
	return origin, nil
}

var _GetThemePropertyOrigin = dll.Uxtheme.NewProc("GetThemePropertyOrigin")

// [GetThemeRect] function.
//
// [GetThemeRect]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemerect
func (hTheme HTHEME) GetThemeRect(partStateId co.VS, propId co.TMT) (RECT, error) {
	var rc RECT
	ret, _, _ := syscall.SyscallN(_GetThemeRect.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&rc)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return RECT{}, hr
	}
	return rc, nil
}

var _GetThemeRect = dll.Uxtheme.NewProc("GetThemeRect")

// [GetThemeString] function.
//
// [GetThemeString]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemestring
func (hTheme HTHEME) GetThemeString(partStateId co.VS, propId co.TMT) (string, error) {
	var buf [wutil.MAX_PATH]uint16 // arbitrary
	ret, _, _ := syscall.SyscallN(_GetThemeString.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}
	return wstr.Utf16SliceToStr(buf[:]), nil
}

var _GetThemeString = dll.Uxtheme.NewProc("GetThemeString")

// [GetThemeSysColorBrush] function.
//
// ⚠️ You must defer HBRUSH.DeleteObject().
//
// [GetThemeSysColorBrush]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesyscolorbrush
func (hTheme HTHEME) GetThemeSysColorBrush(colorId co.TMT) (HBRUSH, error) {
	ret, _, _ := syscall.SyscallN(_GetThemeSysColorBrush.Addr(),
		uintptr(hTheme), uintptr(colorId))
	if hr := co.HRESULT(ret); ret == 0 && hr != co.HRESULT_S_OK {
		return HBRUSH(0), co.HRESULT_E_FAIL // no info is given on error
	}
	return HBRUSH(ret), nil
}

var _GetThemeSysColorBrush = dll.Uxtheme.NewProc("GetThemeSysColorBrush")

// [GetThemeSysFont] function.
//
// [GetThemeSysFont]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesysfont
func (hTheme HTHEME) GetThemeSysFont(fontId co.TMT) (LOGFONT, error) {
	var lf LOGFONT
	ret, _, _ := syscall.SyscallN(_GetThemeSysFont.Addr(),
		uintptr(hTheme), uintptr(fontId), uintptr(unsafe.Pointer(&lf)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return LOGFONT{}, hr
	}
	return lf, nil
}

var _GetThemeSysFont = dll.Uxtheme.NewProc("GetThemeSysFont")

// [GetThemeTextMetrics] function.
//
// [GetThemeTextMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemetextmetrics
func (hTheme HTHEME) GetThemeTextMetrics(hdc HDC, partStateId co.VS) (TEXTMETRIC, error) {
	var tm TEXTMETRIC
	ret, _, _ := syscall.SyscallN(_GetThemeTextMetrics.Addr(),
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(unsafe.Pointer(&tm)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return TEXTMETRIC{}, hr
	}
	return tm, nil
}

var _GetThemeTextMetrics = dll.Uxtheme.NewProc("GetThemeTextMetrics")

// [IsThemeBackgroundPartiallyTransparent] function.
//
// [IsThemeBackgroundPartiallyTransparent]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemebackgroundpartiallytransparent
func (hTheme HTHEME) IsThemeBackgroundPartiallyTransparent(partStateId co.VS) bool {
	ret, _, _ := syscall.SyscallN(_IsThemeBackgroundPartiallyTransparent.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()))
	return ret != 0
}

var _IsThemeBackgroundPartiallyTransparent = dll.Uxtheme.NewProc("IsThemeBackgroundPartiallyTransparent")

// [IsThemePartDefined] function.
//
// [IsThemePartDefined]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemepartdefined
func (hTheme HTHEME) IsThemePartDefined(partStateId co.VS) bool {
	ret, _, _ := syscall.SyscallN(_IsThemePartDefined.Addr(),
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()))
	return ret != 0
}

var _IsThemePartDefined = dll.Uxtheme.NewProc("IsThemePartDefined")
