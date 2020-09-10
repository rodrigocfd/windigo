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
