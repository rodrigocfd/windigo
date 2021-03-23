package proc

import (
	"syscall"
)

var (
	advapi32 = syscall.NewLazyDLL("advapi32.dll")

	RegCloseKey    = advapi32.NewProc("RegCloseKey")
	RegGetValue    = advapi32.NewProc("RegGetValueW")
	RegSetKeyValue = advapi32.NewProc("RegSetKeyValueW")
)
