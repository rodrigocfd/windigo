//go:build windows

package win

// [MODULEINFO] struct.
//
// [MODULEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/ns-psapi-moduleinfo
type MODULEINFO struct {
	LpBaseOfDll uintptr
	SizeOfImage uint32
	EntryPoint  uintptr
}
