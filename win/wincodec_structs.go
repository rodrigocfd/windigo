//go:build windows

package win

// [WICRect] struct.
//
// [WICRect]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ns-wincodec-wicrect
type WICRect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}
