//go:build windows

package win

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// An [atom].
//
// [atom]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#atom
type ATOM uint16

// [FILETIME] struct.
//
// [FILETIME]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-filetime
type FILETIME struct {
	dwLowDateTime  uint32
	dwHighDateTime uint32
}

// Returns the internal value converted to epoch in 100-nanoseconds unit.
func (ft *FILETIME) ToEpochNano100() uint {
	return uint(utl.Make64(ft.dwLowDateTime, ft.dwHighDateTime))
}

// Replaces the internal value with the given epoch in 100-nanoseconds unit.
func (ft *FILETIME) SetEpochNano100(val uint) {
	ft.dwLowDateTime, ft.dwHighDateTime = utl.Break64(uint64(val))
}

// Returns the internal value converted to [time.Time].
func (ft *FILETIME) ToTime() time.Time {
	// https://stackoverflow.com/a/4135003/6923555
	return time.Unix(0, int64(utl.Make64(ft.dwLowDateTime, ft.dwHighDateTime)-116_444_736_000_000_000)*100)
}

// Replaces the internal value with the given [time.Time].
//
// # Example
//
//	var ft win.FILETIME
//	ft.SetTime(time.Now())
func (ft *FILETIME) SetTime(val time.Time) {
	ft.dwLowDateTime, ft.dwHighDateTime = utl.Break64(
		uint64(val.UnixNano()/100 + 116_444_736_000_000_000),
	)
}

// [GUID] struct.
//
// Can be created with [GuidFrom].
//
// [GUID]: https://learn.microsoft.com/en-us/windows/win32/api/guiddef/ns-guiddef-guid
type GUID struct {
	data1 uint32
	data2 uint16
	data3 uint16
	data4 uint64
}

// Formats the GUID as a string.
func (g *GUID) String() string {
	data4 := utl.ReverseBytes64(g.data4)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		g.data1, g.data2, g.data3,
		data4>>48, data4&0xffff_ffff_ffff)
}

// Creates a [GUID] from a const string representation.
//
// Panics if malformed.
func GuidFrom[T ~string](strGuid T) GUID {
	strs := strings.SplitN(string(strGuid), "-", 5)
	if len(strs) != 5 {
		panic(fmt.Sprintf("Malformed GUID parts: %s", strGuid))
	}

	num1, e := strconv.ParseUint(strs[0], 16, 32)
	if e != nil {
		panic(e)
	}
	if num1 > 0xffff_ffff {
		panic(fmt.Sprintf("GUID part 1 overflow: %x", num1))
	}

	var nums16 [3]uint16
	for p := 1; p <= 3; p++ {
		num, e := strconv.ParseUint(strs[p], 16, 16)
		if e != nil {
			panic(e)
		}
		if num > 0xffff {
			panic(fmt.Sprintf("GUID part %d overflows: %x", p, num))
		}
		nums16[p-1] = uint16(num)
	}

	num5, e := strconv.ParseUint(strs[4], 16, 64)
	if e != nil {
		panic(e)
	}
	if num5 > 0xffff_ffff_ffff {
		panic(fmt.Sprintf("GUID part 5 overflow: %x", num5))
	}

	return GUID{
		data1: uint32(num1),
		data2: nums16[0],
		data3: nums16[1],
		data4: utl.ReverseBytes64((uint64(nums16[2]) << 48) | num5),
	}
}

// Language and sublanguage [identifier].
//
// Created with [MAKELANGID].
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/language-identifiers
type LANGID uint16

// Predefined language [identifier].
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/language-identifiers
const (
	LANGID_SYSTEM_DEFAULT LANGID = LANGID((uint16(co.SUBLANG_SYS_DEFAULT) << 10) | uint16(co.LANG_NEUTRAL))
	LANGID_USER_DEFAULT   LANGID = LANGID((uint16(co.SUBLANG_DEFAULT) << 10) | uint16(co.LANG_NEUTRAL))
)

// [MAKELANGID] macro.
//
// [MAKELANGID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-makelangid
func MAKELANGID(lang co.LANG, subLang co.SUBLANG) LANGID {
	return LANGID((uint16(subLang) << 10) | uint16(lang))
}

// [PRIMARYLANGID] macro.
//
// [PRIMARYLANGID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-primarylangid
func (lid LANGID) Lang() co.LANG {
	return co.LANG(uint16(lid) & 0x3ff)
}

// [SUBLANGID] macro.
//
// [SUBLANGID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-sublangid
func (lid LANGID) SubLang() co.SUBLANG {
	return co.SUBLANG(uint16(lid) >> 10)
}

// Locale [identifier].
//
// Created with [MAKELCID].
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/locale-identifiers
type LCID uint32

// Predefined locale [identifier].
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/locale-identifiers
const (
	LCID_SYSTEM_DEFAULT LCID = LCID((uint32(co.SORT_DEFAULT) << 16) | uint32(LANGID_SYSTEM_DEFAULT))
	LCID_USER_DEFAULT   LCID = LCID((uint32(co.SORT_DEFAULT) << 16) | uint32(LANGID_USER_DEFAULT))
)

// [MAKELCID] macro.
//
// [MAKELCID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-makelcid
func MAKELCID(langId LANGID, sortId co.SORT) LCID {
	return LCID((uint32(sortId) << 16) | uint32(langId))
}

// [LANGIDFROMLCID] macro.
//
// [LANGIDFROMLCID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-langidfromlcid
func (lcid LCID) LangId() LANGID {
	return LANGID(uint16(lcid))
}

// [SORTIDFROMLCID] macro.
//
// [SORTIDFROMLCID]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-sortidfromlcid
func (lcid LCID) SortId() co.SORT {
	return co.SORT((uint32(lcid) >> 16) & 0xf)
}

// [MEMORY_BASIC_INFORMATION] struct.
//
// [MEMORY_BASIC_INFORMATION]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-memory_basic_information
type MEMORY_BASIC_INFORMATION struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect co.PAGE
	PartitionId       uint16
	RegionSize        uintptr
	State             co.MEM
	Protect           co.PAGE
	Type              co.MEM
}

// [MODULEENTRY32] struct.
//
// ⚠️ You must call [MODULEENTRY32.SetDwSize] to initialize the struct.
//
// # Example
//
//	var me win.MODULEENTRY32
//	me.SetDwSize()
//
// [MODULEENTRY32]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/ns-tlhelp32-moduleentry32w
type MODULEENTRY32 struct {
	dwSize        uint32
	Th32ModuleID  uint32
	Th32ProcessID uint32
	GlblcntUsage  uint32
	ProccntUsage  uint32
	ModBaseAddr   uintptr
	ModBaseSize   uint32
	HModule       HINSTANCE
	szModule      [utl.MAX_MODULE_NAME32 + 1]uint16
	szExePath     [utl.MAX_PATH]uint16
}

// Sets the dwSize field to the size of the struct, correctly initializing it.
func (me *MODULEENTRY32) SetDwSize() {
	me.dwSize = uint32(unsafe.Sizeof(*me))
}

func (me *MODULEENTRY32) SzModule() string {
	return wstr.WstrSliceToStr(me.szModule[:])
}
func (me *MODULEENTRY32) SetSzModule(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(me.szModule)-1)), me.szModule[:])
}

func (me *MODULEENTRY32) SzExePath() string {
	return wstr.WstrSliceToStr(me.szExePath[:])
}
func (me *MODULEENTRY32) SetSzExePath(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(me.szExePath)-1)), me.szExePath[:])
}

// [OSVERSIONINFOEX] struct.
//
// ⚠️ You must call [OSVERSIONINFOEX.SetDwOsVersionInfoSize] to initialize the
// struct.
//
// # Example
//
//	var osv win.OSVERSIONINFOEX
//	osv.SetDwOsVersionInfoSize()
//
// [OSVERSIONINFOEX]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoexw
type OSVERSIONINFOEX struct {
	dwOsVersionInfoSize uint32
	DwMajorVersion      uint32
	DwMinorVersion      uint32
	DwBuildNumber       uint32
	DWPlatformId        uint32
	szCSDVersion        [128]uint16
	WServicePackMajor   uint16
	WServicePackMinor   uint16
	WSuiteMask          co.VER_SUITE
	WProductType        uint8
	wReserved           uint8
}

// Sets the dwOsVersionInfoSize field to the size of the struct, correctly
// initializing it.
func (osv *OSVERSIONINFOEX) SetDwOsVersionInfoSize() {
	osv.dwOsVersionInfoSize = uint32(unsafe.Sizeof(*osv))
}

func (osv *OSVERSIONINFOEX) SzCSDVersion() string {
	return wstr.WstrSliceToStr(osv.szCSDVersion[:])
}
func (osv *OSVERSIONINFOEX) SetSzCSDVersion(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(osv.szCSDVersion)-1)), osv.szCSDVersion[:])
}

// [OVERLAPPED] struct.
//
// [OVERLAPPED]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-overlapped
type OVERLAPPED struct {
	Internal     uintptr
	InternalHigh uintptr
	Pointer      uintptr
	HEvent       uintptr // HEVENT
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

// [PROCESS_INFORMATION] struct.
//
// [PROCESS_INFORMATION]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-process_information
type PROCESS_INFORMATION struct {
	HProcess    HPROCESS
	HThread     HTHREAD
	DwProcessId uint32
	DwThreadId  uint32
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

// [PROCESSENTRY32] struct.
//
// ⚠️ You must call [PROCESSENTRY32.SetDwSize] to initialize the struct.
//
// # Example
//
//	var pe win.PROCESSENTRY32
//	pe.SetDwSize()
//
// [PROCESSENTRY32]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/ns-tlhelp32-processentry32w
type PROCESSENTRY32 struct {
	dwSize              uint32
	cntUsage            uint32
	Th32ProcessID       uint32
	th32DefaultHeapID   uintptr
	th32ModuleID        uint32
	CntThreads          uint32
	Th32ParentProcessID uint32
	PcPriClassBase      int32
	dwFlags             uint32
	szExeFile           [utl.MAX_PATH]uint16
}

// Sets the dwSize field to the size of the struct, correctly initializing it.
func (pe *PROCESSENTRY32) SetDwSize() {
	pe.dwSize = uint32(unsafe.Sizeof(*pe))
}

func (me *PROCESSENTRY32) SzExeFile() string {
	return wstr.WstrSliceToStr(me.szExeFile[:])
}
func (me *PROCESSENTRY32) SetSzExeFile(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(me.szExeFile)-1)), me.szExeFile[:])
}

// [PROCESSOR_NUMBER] struct.
//
// [PROCESSOR_NUMBER]:
type PROCESSOR_NUMBER struct {
	Group    uint16
	Number   uint8
	reserved uint8
}

// [SECURITY_ATTRIBUTES] struct.
//
// ⚠️ You must call [SECURITY_ATTRIBUTES.SetNLength] to initialize the struct.
//
// # Example
//
//	var sa win.SECURITY_ATTRIBUTES
//	sa.SetNLength()
//
// [SECURITY_ATTRIBUTES]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/aa379560(v=vs.85)
type SECURITY_ATTRIBUTES struct {
	nLength              uint32
	LpSecurityDescriptor uintptr // LPVOID
	bInheritHandle       int32   // BOOL
}

// Sets the nLength field to the size of the struct, correctly initializing it.
func (sa *SECURITY_ATTRIBUTES) SetNLength() {
	sa.nLength = uint32(unsafe.Sizeof(*sa))
}

func (sa *SECURITY_ATTRIBUTES) BInheritHandle() bool {
	return sa.bInheritHandle != 0
}
func (sa *SECURITY_ATTRIBUTES) SetBInheritHandle(val bool) {
	sa.bInheritHandle = utl.BoolToInt32(val)
}

// [STARTUPINFO] struct.
//
// ⚠️ You must call [STARTUPINFO.SetCb] to initialize the struct.
//
// # Example
//
//	var si win.STARTUPINFO
//	si.SetCb()
//
// [STARTUPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-startupinfow
type STARTUPINFO struct {
	cb              uint32
	lpReserved      *uint16
	LpDesktop       *uint16
	LpTitle         *uint16
	DwX             uint32
	DwY             uint32
	DwXSize         uint32
	DwYSize         uint32
	DwXCountChars   uint32
	DwYCountChars   uint32
	DwFillAttribute co.STARTFILL
	DwFlags         co.STARTF
	WShowWindow     co.STARTSW
	cbReserved2     uint16
	lpReserved2     *uint8
	HStdInput       HANDLE
	HStdOutput      HANDLE
	HStdError       HANDLE
}

// Sets the cb field to the size of the struct, correctly initializing it.
func (si *STARTUPINFO) SetCb() {
	si.cb = uint32(unsafe.Sizeof(*si))
}

// [SYSTEM_INFO] struct.
//
// [SYSTEM_INFO]: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info
type SYSTEM_INFO struct {
	WProcessorArchitecture      co.PROCESSOR_ARCHITECTURE
	wReserved                   uint16
	DwPageSize                  uint32
	LpMinimumApplicationAddress uintptr
	LpMaximumApplicationAddress uintptr
	DwActiveProcessorMask       uintptr
	DwNumberOfProcessors        uint32
	DwProcessorType             co.PROCESSOR
	DwAllocationGranularity     uint32
	WProcessorLevel             uint16
	WProcessorRevision          uint16
}

// [SYSTEMTIME] struct.
//
// [SYSTEMTIME]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-systemtime
type SYSTEMTIME struct {
	WYear         uint16
	WMonth        uint16
	WDayOfWeek    uint16
	WDay          uint16
	WHour         uint16
	WMinute       uint16
	WSecond       uint16
	WMilliseconds uint16
}

// Decomposes a [time.Duration] into this SYSTEMTIME fields.
//
// # Example
//
//	var st win.SYSTEMTIME
//	st.SetDuration(time.Minute * 3)
func (st *SYSTEMTIME) SetDuration(dur time.Duration) {
	*st = SYSTEMTIME{}
	st.WHour = uint16(dur / time.Hour)
	st.WMinute = uint16((dur -
		time.Duration(st.WHour)*time.Hour) / time.Minute)
	st.WSecond = uint16((dur -
		time.Duration(st.WHour)*time.Hour -
		time.Duration(st.WMinute)*time.Minute) / time.Second)
	st.WMilliseconds = uint16((dur -
		time.Duration(st.WHour)*time.Hour -
		time.Duration(st.WMinute)*time.Minute -
		time.Duration(st.WSecond)*time.Second) / time.Millisecond)
}

// Converts this SYSTEMTIME to [time.Time].
func (st *SYSTEMTIME) ToTime() time.Time {
	return time.Date(int(st.WYear), time.Month(st.WMonth), int(st.WDay),
		int(st.WHour), int(st.WMinute), int(st.WSecond),
		int(st.WMilliseconds)*1_000_000,
		time.Local)
}

// Fills this SYSTEMTIME with the value of a [time.Time], in the current
// timezone.
//
// # Example
//
//	var st win.SYSTEMTIME
//	st.SetTime(time.Now())
func (st *SYSTEMTIME) SetTime(val time.Time) error {
	var ft FILETIME
	ft.SetTime(val)

	stUtc, err := FileTimeToSystemTime(&ft)
	if err != nil {
		return err
	}

	// When converted, SYSTEMTIME will receive UTC values, so we need to convert
	// the fields to current timezone.
	*st, err = SystemTimeToTzSpecificLocalTime(nil, &stUtc)
	return err
}

// [THREADENTRY32] struct.
//
// ⚠️ You must call [THREADENTRY32.SetDwSize] to initialize the struct.
//
// # Example
//
//	var te win.THREADENTRY32
//	te.SetDwSize()
//
// [THREADENTRY32]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/ns-tlhelp32-threadentry32
type THREADENTRY32 struct {
	dwSize             uint32
	cntUsage           uint32
	Th32ThreadID       uint32
	Th32OwnerProcessID uint32
	TpBasePri          int32
	tpDeltaPri         int32
	dwFlags            uint32
}

// Sets the dwSize field to the size of the struct, correctly initializing it.
func (te *THREADENTRY32) SetDwSize() {
	te.dwSize = uint32(unsafe.Sizeof(*te))
}

// [TIME_ZONE_INFORMATION] struct.
//
// [TIME_ZONE_INFORMATION]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/ns-timezoneapi-time_zone_information
type TIME_ZONE_INFORMATION struct {
	Bias         int32
	standardName [32]uint16
	StandardDate SYSTEMTIME
	StandardBias int32
	daylightName [32]uint16
	DaylightDate SYSTEMTIME
	DaylightBias int32
}

func (tzi *TIME_ZONE_INFORMATION) StandardName() string {
	return wstr.WstrSliceToStr(tzi.standardName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetStandardName(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(tzi.standardName)-1)), tzi.standardName[:])
}

func (tzi *TIME_ZONE_INFORMATION) DaylightName() string {
	return wstr.WstrSliceToStr(tzi.daylightName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetDaylightName(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(tzi.daylightName)-1)), tzi.daylightName[:])
}

// [WIN32_FIND_DATA] struct.
//
// [WIN32_FIND_DATA]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-win32_find_dataw
type WIN32_FIND_DATA struct {
	DwFileAttributes   co.FILE_ATTRIBUTE
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	NFileSizeHigh      uint32
	NFileSizeLow       uint32
	dwReserved0        uint32
	dwReserved1        uint32
	cFileName          [utl.MAX_PATH]uint16
	cAlternateFileName [14]uint16
	DwFileType         uint32
	DwCreatorType      uint32
	WFinderFlags       uint16
}

func (wfd *WIN32_FIND_DATA) CFileName() string {
	return wstr.WstrSliceToStr(wfd.cFileName[:])
}
func (wfd *WIN32_FIND_DATA) SetCFileName(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(wfd.cFileName)-1)), wfd.cFileName[:])
}

func (wfd *WIN32_FIND_DATA) CAlternateFileName() string {
	return wstr.WstrSliceToStr(wfd.cAlternateFileName[:])
}
func (wfd *WIN32_FIND_DATA) SetCAlternateFileName(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(wfd.cAlternateFileName)-1)), wfd.cAlternateFileName[:])
}
