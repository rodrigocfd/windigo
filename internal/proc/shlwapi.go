package proc

import (
	"syscall"
)

var (
	shlwapi = syscall.NewLazyDLL("shlwapi")

	SHCreateMemStream = shlwapi.NewProc("SHCreateMemStream")
)
