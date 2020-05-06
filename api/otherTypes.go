package api

import "syscall"

type (
	ATOM     uint16
	COLORREF uint32
	HANDLE   syscall.Handle
	HBITMAP  HGDIOBJ
	HCURSOR  HANDLE
	HGDIOBJ  HANDLE
	HICON    HANDLE
	HMENU    HANDLE
	HMONITOR HANDLE
	HRGN     HGDIOBJ
)

type (
	LPARAM uintptr
	WPARAM uintptr
)
