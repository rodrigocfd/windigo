/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

import (
	"encoding/binary"
)

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

func makeGuid(d1 uint32, d2, d3 uint16, d4 uint64) GUID {
	guid := GUID{Data1: d1, Data2: d2, Data3: d3}
	binary.BigEndian.PutUint64(guid.Data4[:], d4)
	return guid
}

var (
	Guid_IUnknown = makeGuid(0x00000000, 0x0000, 0x0000, 0xc000000000000046)

	Guid_ITaskbarList  = makeGuid(0x56fdf344, 0xfd6d, 0x11d0, 0x958a006097c9a090)
	Guid_ITaskbarList3 = makeGuid(0xea1afb91, 0x9e28, 0x4b86, 0x90e99e9f8a5eefaf)
)
