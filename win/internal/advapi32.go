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
	dllAdvapi32 = syscall.NewLazyDLL("advapi32.dll")

	RegCloseKey     = dllAdvapi32.NewProc("RegCloseKey")
	RegEnumValue    = dllAdvapi32.NewProc("RegEnumValueW")
	RegOpenKeyEx    = dllAdvapi32.NewProc("RegOpenKeyExW")
	RegQueryValueEx = dllAdvapi32.NewProc("RegQueryValueExW")
)
