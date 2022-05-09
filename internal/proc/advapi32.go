//go:build windows

package proc

import (
	"syscall"
)

var (
	advapi32 = syscall.NewLazyDLL("advapi32.dll")

	RegCloseKey       = advapi32.NewProc("RegCloseKey")
	RegDeleteKey      = advapi32.NewProc("RegDeleteKeyW")
	RegDeleteKeyEx    = advapi32.NewProc("RegDeleteKeyExW")
	RegDeleteKeyValue = advapi32.NewProc("RegDeleteKeyValueW")
	RegDeleteTree     = advapi32.NewProc("RegDeleteTreeW")
	RegEnumKeyEx      = advapi32.NewProc("RegEnumKeyExW")
	RegEnumValue      = advapi32.NewProc("RegEnumValueW")
	RegFlushKey       = advapi32.NewProc("RegFlushKey")
	RegGetValue       = advapi32.NewProc("RegGetValueW")
	RegOpenKeyEx      = advapi32.NewProc("RegOpenKeyExW")
	RegQueryInfoKey   = advapi32.NewProc("RegQueryInfoKeyW")
	RegSetKeyValue    = advapi32.NewProc("RegSetKeyValueW")
)
