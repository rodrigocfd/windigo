/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

func hiWord(value uint32) uint16 { return uint16(value >> 16 & 0xffff) }
func loWord(value uint32) uint16 { return uint16(value) }
func hiByte(value uint16) uint8  { return uint8(value >> 8 & 0xff) }
func loByte(value uint16) uint8  { return uint8(value) }

// Simple conversion for syscalls.
func boolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// Simple conversion for syscalls.
func boolToUintptr(b bool) uintptr {
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
