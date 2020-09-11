/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
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

// Converts a string into a GUID.
func (_WinT) StrToGuid(s string) GUID {
	parts := strings.Split(s, "-")
	if len(parts) != 5 {
		panic(fmt.Sprintf("Bad GUID with %d parts: %s.", len(parts), s))
	}
	guid := GUID{}

	if n1, err := strconv.ParseUint(parts[0], 16, 32); err != nil {
		panic(fmt.Sprintf("Bad 1st part GUID: %s, %s. %s", parts[0], s, err.Error()))
	} else {
		guid.Data1 = uint32(n1)
	}

	if n2, err := strconv.ParseUint(parts[1], 16, 16); err != nil {
		panic(fmt.Sprintf("Bad 2nd part GUID: %s, %s. %s", parts[1], s, err.Error()))
	} else {
		guid.Data2 = uint16(n2)
	}

	if n3, err := strconv.ParseUint(parts[2], 16, 16); err != nil {
		panic(fmt.Sprintf("Bad 3rd part GUID: %s, %s. %s", parts[2], s, err.Error()))
	} else {
		guid.Data2 = uint16(n3)
	}

	if n4, err := strconv.ParseUint(parts[3]+parts[4], 16, 64); err != nil {
		panic(fmt.Sprintf("Bad 4th part GUID: %s, %s. %s", parts[3]+parts[4], s, err.Error()))
	} else {
		guid.Data4 = n4
	}

	buf64 := [8]byte{}
	binary.BigEndian.PutUint64(buf64[:], guid.Data4)
	guid.Data4 = binary.LittleEndian.Uint64(buf64[:]) // reverse bytes

	return guid
}
