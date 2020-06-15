/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	dllAdvapi32 = syscall.NewLazyDLL("advapi32.dll")

	RegCloseKey = dllAdvapi32.NewProc("RegCloseKey")
	// RegOpenKeyEx = dllAdvapi32.NewProc("RegOpenKeyExW")
)
