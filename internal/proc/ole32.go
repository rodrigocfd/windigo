//go:build windows

package proc

import (
	"syscall"
)

var (
	ole32 = syscall.NewLazyDLL("ole32.dll")

	CLSIDFromProgID  = ole32.NewProc("CLSIDFromProgID")
	CoCreateGuid     = ole32.NewProc("CoCreateGuid")
	CoCreateInstance = ole32.NewProc("CoCreateInstance")
	CoInitializeEx   = ole32.NewProc("CoInitializeEx")
	CoTaskMemAlloc   = ole32.NewProc("CoTaskMemAlloc")
	CoTaskMemFree    = ole32.NewProc("CoTaskMemFree")
	CoTaskMemRealloc = ole32.NewProc("CoTaskMemRealloc")
	CoUninitialize   = ole32.NewProc("CoUninitialize")
	OleInitialize    = ole32.NewProc("OleInitialize")
	OleUninitialize  = ole32.NewProc("OleUninitialize")
	RegisterDragDrop = ole32.NewProc("RegisterDragDrop")
	RevokeDragDrop   = ole32.NewProc("RevokeDragDrop")
)
