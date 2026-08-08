//go:build windows

// This package contains native [Automation] functions, structs and handles.
// They are implemented as close as possible to the original C/C++ declarations,
// so you can use the abundant online documentation. In addition o that, each
// entity has a link to its [official docs], so you can lookup the correct
// usage.
//
// All constants are declared in the cowic package.
//
// [Automation]: https://learn.microsoft.com/en-us/windows/win32/api/_automat/
// [official docs]: https://learn.microsoft.com/en-us/windows/win32/api/
package winaut
