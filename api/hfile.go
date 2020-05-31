/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package api

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/api/proc"
	"wingows/co"
)

type HFILE HANDLE

func (hfile HFILE) GetFileSize() int64 {
	buf := int64(0)
	ret, _, lerr := syscall.Syscall(proc.GetFileSizeEx.Addr(), 1,
		uintptr(hfile), uintptr(unsafe.Pointer(&buf)), 0)
	if ret == 0 && lerr != 0 {
		panic(fmt.Sprintf("GetFileSizeEx failed: %d %s\n",
			lerr, lerr.Error()))
	}
	return buf
}

func CreateFile(fileName string, desiredAccess co.GENERIC,
	shareMode co.FILE_SHARE, securityAttributes uintptr,
	creationDisposition co.FILE_DISP, attributes co.FILE_ATTRIBUTE,
	flags co.FILE_FLAG, security co.SECURITY) HFILE {

	return 0
}
