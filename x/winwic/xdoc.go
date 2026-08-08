//go:build windows

// This package contains native [Windows Imaging Component] functions, structs
// and handles. They are implemented as close as possible to the original C/C++
// declarations, so you can use the abundant online documentation. In addition
// o that, each entity has a link to its [official docs], so you can lookup the
// correct usage.
//
// All constants are declared in the cowic package.
//
// [Windows Imaging Component]: https://learn.microsoft.com/en-us/windows/win32/wic/-wic-lh
// [official docs]: https://learn.microsoft.com/en-us/windows/win32/api/
package winwic
