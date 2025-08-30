//go:build windows

package co

// [GetWindowLongPtr] and [SetWindowLongPtr] nIndex. Also includes constants
// with GWL prefix.
//
// [GetWindowLongPtr]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
// [SetWindowLongPtr]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowlongptrw
type GWLP int32

const (
	GWLP_WNDPROC    GWLP = -4  // Originally with GWL prefix.
	GWLP_HINSTANCE  GWLP = -6  // Originally with GWL prefix.
	GWLP_HWNDPARENT GWLP = -8  // Originally with GWL prefix.
	GWLP_ID         GWLP = -12 // Originally with GWL prefix.
	GWLP_STYLE      GWLP = -16 // Originally with GWL prefix.
	GWLP_EXSTYLE    GWLP = -20 // Originally with GWL prefix.
	GWLP_USERDATA   GWLP = -21 // Originally with GWL prefix.

	GWLP_DWLP_MSGRESULT GWLP = 0 // Originally with DWL prefix.
	GWLP_DWLP_DLGPROC   GWLP = 4 // Originally with DWL prefix.
	GWLP_DWLP_USER      GWLP = 8 // Originally with DWL prefix.
)
