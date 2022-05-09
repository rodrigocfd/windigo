//go:build windows

package co

// Get/SetWindowLongPtr() nIndex. Also includes constants with GWL prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
type GWL int32

const (
	GWL_WNDPROC    GWL = -4
	GWL_HINSTANCE  GWL = -6
	GWL_HWNDPARENT GWL = -8
	GWL_ID         GWL = -12
	GWL_STYLE      GWL = -16
	GWL_EXSTYLE    GWL = -20
	GWL_USERDATA   GWL = -21

	GWL_DWL_MSGRESULT GWL = 0 // Originally with DWL prefix.
	GWL_DWL_DLGPROC   GWL = 4 // Originally with DWL prefix.
	GWL_DWL_USER      GWL = 8 // Originally with DWL prefix.
)
