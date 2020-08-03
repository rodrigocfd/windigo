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
