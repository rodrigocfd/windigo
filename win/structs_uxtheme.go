//go:build windows

package win

// [MARGINS] struct.
//
// [MARGINS]: https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/ns-uxtheme-margins
type MARGINS struct {
	CxLeftWidth    int32
	CxRightWidth   int32
	CyTopHeight    int32
	CyBottomHeight int32
}
