//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Handle to a [theme].
//
// [theme]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/
type HTHEME HANDLE

// [CloseThemeData] function.
//
// Paired with [HWND.OpenThemeData].
//
// [CloseThemeData]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-closethemedata
func (hTheme HTHEME) CloseThemeData() {
	syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_CloseThemeData, "CloseThemeData"),
		uintptr(hTheme))
}

var _uxtheme_CloseThemeData *syscall.Proc

// [DrawThemeBackground] function.
//
// [DrawThemeBackground]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-drawthemebackground
func (hTheme HTHEME) DrawThemeBackground(hdc HDC, partStateId co.VS, rc *RECT, clipRc *RECT) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_DrawThemeBackground, "DrawThemeBackground"),
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()),
		uintptr(unsafe.Pointer(rc)),
		uintptr(unsafe.Pointer(clipRc)))
	return utl.ErrorAsHResult(ret)
}

var _uxtheme_DrawThemeBackground *syscall.Proc

// [GetThemeColor] function.
//
// [GetThemeColor]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemecolor
func (hTheme HTHEME) GetThemeColor(partStateId co.VS, propId co.TMT) (COLORREF, error) {
	var color COLORREF
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemeColor, "GetThemeColor"),
		uintptr(hTheme),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()),
		uintptr(propId),
		uintptr(unsafe.Pointer(&color)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return COLORREF(0), hr
	}
	return color, nil
}

var _uxtheme_GetThemeColor *syscall.Proc

// [GetThemeInt] function.
//
// [GetThemeInt]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemeint
func (hTheme HTHEME) GetThemeInt(partStateId co.VS, propId co.TMT) (int, error) {
	var intVal int32
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemeInt, "GetThemeInt"),
		uintptr(hTheme),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()),
		uintptr(propId),
		uintptr(unsafe.Pointer(&intVal)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(intVal), nil
}

var _uxtheme_GetThemeInt *syscall.Proc

// [GetThemeMetric] function.
//
// [GetThemeMetric]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthememetric
func (hTheme HTHEME) GetThemeMetric(hdc HDC, partStateId co.VS, propId co.TMT) (int, error) {
	var intVal int32
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemeMetric, "GetThemeMetric"),
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()),
		uintptr(propId),
		uintptr(unsafe.Pointer(&intVal)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(intVal), nil
}

var _uxtheme_GetThemeMetric *syscall.Proc

// [GetThemePosition] function.
//
// [GetThemePosition]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemeposition
func (hTheme HTHEME) GetThemePosition(partStateId co.VS, propId co.TMT) (POINT, error) {
	var pt POINT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemePosition, "GetThemePosition"),
		uintptr(hTheme),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()),
		uintptr(propId),
		uintptr(unsafe.Pointer(&pt)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return POINT{}, hr
	}
	return pt, nil
}

var _uxtheme_GetThemePosition *syscall.Proc

// [GetThemePropertyOrigin] function.
//
// [GetThemePropertyOrigin]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemepropertyorigin
func (hTheme HTHEME) GetThemePropertyOrigin(partStateId co.VS, propId co.TMT) (co.PROPERTYORIGIN, error) {
	var origin co.PROPERTYORIGIN
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemePropertyOrigin, "GetThemePropertyOrigin"),
		uintptr(hTheme),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()),
		uintptr(propId),
		uintptr(unsafe.Pointer(&origin)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return co.PROPERTYORIGIN(0), hr
	}
	return origin, nil
}

var _uxtheme_GetThemePropertyOrigin *syscall.Proc

// [GetThemeRect] function.
//
// [GetThemeRect]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemerect
func (hTheme HTHEME) GetThemeRect(partStateId co.VS, propId co.TMT) (RECT, error) {
	var rc RECT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemeRect, "GetThemeRect"),
		uintptr(hTheme),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()),
		uintptr(propId),
		uintptr(unsafe.Pointer(&rc)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return RECT{}, hr
	}
	return rc, nil
}

var _uxtheme_GetThemeRect *syscall.Proc

// [GetThemeString] function.
//
// [GetThemeString]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemestring
func (hTheme HTHEME) GetThemeString(partStateId co.VS, propId co.TMT) (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemeString, "GetThemeString"),
		uintptr(hTheme),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()),
		uintptr(propId),
		uintptr(wBuf.Ptr()),
		uintptr(int32(wBuf.Len())))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}
	return wBuf.String(), nil
}

var _uxtheme_GetThemeString *syscall.Proc

// [GetThemeSysColorBrush] function.
//
// ⚠️ You must defer [HBRUSH.DeleteObject].
//
// [GetThemeSysColorBrush]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesyscolorbrush
func (hTheme HTHEME) GetThemeSysColorBrush(colorId co.TMT) (HBRUSH, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemeSysColorBrush, "GetThemeSysColorBrush"),
		uintptr(hTheme),
		uintptr(colorId))
	if hr := co.HRESULT(ret); ret == 0 && hr != co.HRESULT_S_OK {
		return HBRUSH(0), co.HRESULT_E_FAIL // no info is given on error
	}
	return HBRUSH(ret), nil
}

var _uxtheme_GetThemeSysColorBrush *syscall.Proc

// [GetThemeSysFont] function.
//
// [GetThemeSysFont]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesysfont
func (hTheme HTHEME) GetThemeSysFont(fontId co.TMT) (LOGFONT, error) {
	var lf LOGFONT
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemeSysFont, "GetThemeSysFont"),
		uintptr(hTheme),
		uintptr(fontId),
		uintptr(unsafe.Pointer(&lf)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return LOGFONT{}, hr
	}
	return lf, nil
}

var _uxtheme_GetThemeSysFont *syscall.Proc

// [GetThemeTextMetrics] function.
//
// [GetThemeTextMetrics]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemetextmetrics
func (hTheme HTHEME) GetThemeTextMetrics(hdc HDC, partStateId co.VS) (TEXTMETRIC, error) {
	var tm TEXTMETRIC
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_GetThemeTextMetrics, "GetThemeTextMetrics"),
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()),
		uintptr(unsafe.Pointer(&tm)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return TEXTMETRIC{}, hr
	}
	return tm, nil
}

var _uxtheme_GetThemeTextMetrics *syscall.Proc

// [IsThemeBackgroundPartiallyTransparent] function.
//
// [IsThemeBackgroundPartiallyTransparent]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemebackgroundpartiallytransparent
func (hTheme HTHEME) IsThemeBackgroundPartiallyTransparent(partStateId co.VS) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_IsThemeBackgroundPartiallyTransparent, "IsThemeBackgroundPartiallyTransparent"),
		uintptr(hTheme),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()))
	return ret != 0
}

var _uxtheme_IsThemeBackgroundPartiallyTransparent *syscall.Proc

// [IsThemePartDefined] function.
//
// [IsThemePartDefined]: https://learn.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemepartdefined
func (hTheme HTHEME) IsThemePartDefined(partStateId co.VS) bool {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.UXTHEME, &_uxtheme_IsThemePartDefined, "IsThemePartDefined"),
		uintptr(hTheme),
		uintptr(partStateId.Part()),
		uintptr(partStateId.State()))
	return ret != 0
}

var _uxtheme_IsThemePartDefined *syscall.Proc
