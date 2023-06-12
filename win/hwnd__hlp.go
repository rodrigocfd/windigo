//go:build windows

package win

import "github.com/rodrigocfd/windigo/win/co"

// This helper method returns the window instance with GetWindowLongPtr().
func (hWnd HWND) Hinstance() HINSTANCE {
	return HINSTANCE(hWnd.GetWindowLongPtr(co.GWLP_HINSTANCE))
}

// Allegedly [undocumented] Win32 function.
//
// [undocumented]: https://stackoverflow.com/a/16975012
func (hWnd HWND) IsTopLevelWindow() bool {
	return hWnd == hWnd.GetAncestor(co.GA_ROOT)
}
