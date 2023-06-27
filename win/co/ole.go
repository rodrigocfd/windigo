//go:build windows

package co

// A COM [class ID], represented as a string.
//
// [class ID]: https://learn.microsoft.com/en-us/windows/win32/com/clsid-key-hklm
type CLSID string

// A COM [interface ID], represented as a string.
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
type IID string
