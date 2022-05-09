//go:build windows

package proc

import (
	"syscall"
)

var (
	d2d1 = syscall.NewLazyDLL("d2d1.dll")

	D2D1CreateFactory = d2d1.NewProc("D2D1CreateFactory")
)
