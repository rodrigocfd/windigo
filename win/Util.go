/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

type _UtilT struct{}

// Internal win package utilities.
var _Util _UtilT

// Syntactic sugar; converts bool to 0 or 1.
func (_UtilT) BoolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func (_UtilT) BoolToUint32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func (_UtilT) BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// Private constants.
const (
	_CCHDEVICENAME        = 32
	_CLR_INVALID          = 0xFFFF_FFFF
	_HGDI_ERROR           = 0xFFFF_FFFF
	_INVALID_FILE_SIZE    = 0xFFFF_FFFF
	_INVALID_HANDLE_VALUE = -1
	_L_MAX_URL_LENGTH     = 2048 + 32 + 3
	_LF_FACESIZE          = 32
	_MAX_LINKID_TEXT      = 48
	_MAX_PATH             = 260
)
