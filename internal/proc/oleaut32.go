//go:build windows

package proc

import (
	"syscall"
)

var (
	oleaut32 = syscall.NewLazyDLL("oleaut32.dll")

	OleLoadPicture          = oleaut32.NewProc("OleLoadPicture")
	OleLoadPicturePath      = oleaut32.NewProc("OleLoadPicturePath")
	SysAllocString          = oleaut32.NewProc("SysAllocString")
	SysFreeString           = oleaut32.NewProc("SysFreeString")
	SysReAllocString        = oleaut32.NewProc("SysReAllocString")
	SystemTimeToVariantTime = oleaut32.NewProc("SystemTimeToVariantTime")
	VariantClear            = oleaut32.NewProc("VariantClear")
	VariantInit             = oleaut32.NewProc("VariantInit")
	VariantTimeToSystemTime = oleaut32.NewProc("VariantTimeToSystemTime")
)
