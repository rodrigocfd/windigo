//go:build windows

// This package contains native Win32 functions, structs and handles. They are
// implemented as close as possible to the original C/C++ declarations, so you
// can use the abundant online documentation. In addition to that, each entity
// has a link to its [official docs], so you can lookup the correct usage.
//
// Functions whose first parameter is a handle are implemented as a method of
// that handle. For example, [MessageBox] function's first parameter is a HWND,
// so MessageBox is a method of HWND:
//
//	var hwnd win.HWND // initialized somewhere
//
//	hwnd.MessageBox("Saying hello", "Hello", co.MB_ICONINFORMATION)
//
// All other functions are just regular free functions.
//
// Some structs, like [WNDCLASSEX], require the initialization of a cbSize field
// (or other similar name) with the size of the struct itself. Each one of these
// structs is properly documented, and they have a specific method for this, for
// example:
//
//	var wcx win.WNDCLASSEX
//	wcx.SetCbSize()
//
// All constants are declared in the co package.
//
// [MessageBox]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
// [WNDCLASSEX]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-wndclassexw
// [official docs]: https://learn.microsoft.com/en-us/windows/win32/api/
package win
