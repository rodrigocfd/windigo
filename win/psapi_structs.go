//go:build windows

package win

import (
	"unsafe"
)

// [MODULEINFO] struct.
//
// [MODULEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/ns-psapi-moduleinfo
type MODULEINFO struct {
	LpBaseOfDll uintptr
	SizeOfImage uint32
	EntryPoint  uintptr
}

// [PERFORMANCE_INFORMATION] struct.
//
// ⚠️ You must call [PERFORMANCE_INFORMATION.SetCb] to initialize the struct.
//
// # Example
//
//	var pi win.PERFORMANCE_INFORMATION
//	pi.SetCb()
//
// [PERFORMANCE_INFORMATION]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/ns-psapi-performance_information
type PERFORMANCE_INFORMATION struct {
	cb                uint32
	CommitTotal       uintptr
	CommitLimit       uintptr
	CommitPeak        uintptr
	PhysicalTotal     uintptr
	PhysicalAvailable uintptr
	SystemCache       uintptr
	KernelTotal       uintptr
	KernelPaged       uintptr
	KernelNonpaged    uintptr
	PageSize          uintptr
	HandleCount       uint32
	ProcessCount      uint32
	ThreadCount       uint32
}

// Sets the cb field to the size of the struct, correctly initializing it.
func (pi *PERFORMANCE_INFORMATION) SetCb() {
	pi.cb = uint32(unsafe.Sizeof(*pi))
}

// [PROCESS_MEMORY_COUNTERS_EX] struct.
//
// ⚠️ You must call [PROCESS_MEMORY_COUNTERS_EX.SetCb] to initialize the struct.
//
// # Example
//
//	var pi win.PERFORMANCE_INFORMATION
//	pi.SetCb()
//
// [PROCESS_MEMORY_COUNTERS_EX]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/ns-psapi-process_memory_counters_ex
type PROCESS_MEMORY_COUNTERS_EX struct {
	cb                         uint32
	PageFaultCount             uint32
	PeakWorkingSetSize         uintptr
	WorkingSetSize             uintptr
	QuotaPeakPagedPoolUsage    uintptr
	QuotaPagedPoolUsage        uintptr
	QuotaPeakNonPagedPoolUsage uintptr
	QuotaNonPagedPoolUsage     uintptr
	PagefileUsage              uintptr
	PeakPagefileUsage          uintptr
	PrivateUsage               uintptr
}

// Sets the cb field to the size of the struct, correctly initializing it.
func (pmc *PROCESS_MEMORY_COUNTERS_EX) SetCb() {
	pmc.cb = uint32(unsafe.Sizeof(*pmc))
}
