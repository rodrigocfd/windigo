package co

// Get/SetWindowLongPtr() nIndex. Also includes constants with GWL prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
type GWL int32

const (
	GWL_WNDPROC    GWL = -4  // Originally with GWLP prefix.
	GWL_HINSTANCE  GWL = -6  // Originally with GWLP prefix.
	GWL_HWNDPARENT GWL = -8  // Originally with GWLP prefix.
	GWL_ID         GWL = -12 // Originally with GWLP prefix.
	GWL_STYLE      GWL = -16
	GWL_EXSTYLE    GWL = -20
	GWL_USERDATA   GWL = -21 // Originally with GWLP prefix.

	GWL_DWL_MSGRESULT GWL = 0                     // Originally with DWLP prefix.
	GWL_DWL_DLGPROC   GWL = GWL_DWL_MSGRESULT + 8 // Originally with DWLP prefix.
	GWL_DWL_USER      GWL = GWL_DWL_DLGPROC + 8   // Originally with DWLP prefix.
)
