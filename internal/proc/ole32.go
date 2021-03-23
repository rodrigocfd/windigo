package proc

import (
	"syscall"
)

var (
	ole32 = syscall.NewLazyDLL("ole32.dll")

	CoCreateInstance = ole32.NewProc("CoCreateInstance")
	CoInitializeEx   = ole32.NewProc("CoInitializeEx")
	CoTaskMemFree    = ole32.NewProc("CoTaskMemFree")
	CoUninitialize   = ole32.NewProc("CoUninitialize")
)
