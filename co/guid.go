/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 uint64
}

var (
	Guid_IUnknown = GUID{0x00000000, 0x0000, 0x0000, 0xc000000000000046}

	Guid_ITaskbarList  = GUID{0x56fdf344, 0xfd6d, 0x11d0, 0x958a006097c9a090}
	Guid_ITaskbarList3 = GUID{0xea1afb91, 0x9e28, 0x4b86, 0x90e99e9f8a5eefaf}
)
