//go:build windows

package win

// Handle to a
// [palette](https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hpalette).
type HPALETTE HANDLE

// [DeleteObject] function.
//
// [DeleteObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hPal HPALETTE) DeleteObject() error {
	return HGDIOBJ(hPal).DeleteObject()
}
