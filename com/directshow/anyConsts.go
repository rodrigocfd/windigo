/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package directshow

import (
	"github.com/rodrigocfd/windigo/co"
)

const (
	ERROR_VFW_E_NOT_FOUND co.ERROR = 0x80040216
)

// https://docs.microsoft.com/en-us/windows/win32/api/evr/ne-evr-mfvideoaspectratiomode
type MFVideoARMode uint32

const (
	MFVideoARMode_None             MFVideoARMode = 0
	MFVideoARMode_PreservePicture  MFVideoARMode = 0x1
	MFVideoARMode_PreservePixel    MFVideoARMode = 0x2
	MFVideoARMode_NonLinearStretch MFVideoARMode = 0x4
	MFVideoARMode_Mask             MFVideoARMode = 0x7
)
