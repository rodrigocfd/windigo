/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	advapi32Dll = syscall.NewLazyDLL("advapi32.dll")

	RegCloseKey       = advapi32Dll.NewProc("RegCloseKey")
	RegDeleteKeyValue = advapi32Dll.NewProc("RegDeleteKeyValueW")
	RegDeleteValue    = advapi32Dll.NewProc("RegDeleteValueW")
	RegEnumValue      = advapi32Dll.NewProc("RegEnumValueW")
	RegGetValue       = advapi32Dll.NewProc("RegGetValueW")
	RegOpenKeyEx      = advapi32Dll.NewProc("RegOpenKeyExW")
	RegQueryValueEx   = advapi32Dll.NewProc("RegQueryValueExW")
	RegSetKeyValue    = advapi32Dll.NewProc("RegSetKeyValueW")
)
