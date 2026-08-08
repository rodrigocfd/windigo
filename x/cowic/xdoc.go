//go:build windows

// This package contains native [Windows Imaging Component] constants. All
// constants have their own type, in order to avoid improper mixing.
//
// They are named as close as possible to the original C/C++ declarations, so
// you can use the abundant online documentation. In addition to that, each
// entity has a link to its [official docs], so you can lookup the correct
// usage.
//
// [Windows Imaging Component]: https://learn.microsoft.com/en-us/windows/win32/wic/-wic-lh
// [official docs]: https://learn.microsoft.com/en-us/windows/win32/api/
package cowic
