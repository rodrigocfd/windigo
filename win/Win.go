/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"encoding/binary"
)

type _WinT struct{}

// Internal win package utilities.
var _Win _WinT

// Syntactic sugar; converts bool to 0 or 1.
func (_WinT) BoolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func (_WinT) BoolToUint32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func (_WinT) BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// Builds a GUID struct.
func (_WinT) NewGuid(d1 uint32, d2, d3 uint16, d4 uint64) *GUID {
	newGuid := GUID{
		Data1: d1,
		Data2: d2,
		Data3: d3,
		Data4: d4,
	}

	buf64 := [8]byte{}
	binary.BigEndian.PutUint64(buf64[:], newGuid.Data4)
	newGuid.Data4 = binary.LittleEndian.Uint64(buf64[:]) // reverse bytes
	return &newGuid
}
