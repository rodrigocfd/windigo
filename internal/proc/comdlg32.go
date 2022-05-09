//go:build windows

package proc

import (
	"syscall"
)

var (
	comdlg32 = syscall.NewLazyDLL("comdlg32.dll")

	ChooseColor          = comdlg32.NewProc("ChooseColorW")
	CommDlgExtendedError = comdlg32.NewProc("CommDlgExtendedError")
)
