package proc

import (
	"syscall"
)

var (
	ole32 = syscall.NewLazyDLL("ole32.dll")

	CoCreateInstance = ole32.NewProc("CoCreateInstance")
	CoInitializeEx   = ole32.NewProc("CoInitializeEx")
	CoTaskMemAlloc   = ole32.NewProc("CoTaskMemAlloc")
	CoTaskMemFree    = ole32.NewProc("CoTaskMemFree")
	CoTaskMemRealloc = ole32.NewProc("CoTaskMemRealloc")
	CoUninitialize   = ole32.NewProc("CoUninitialize")
)
