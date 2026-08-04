//go:build windows

package win

import (
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [ACTCTX] struct, with C memory layout.
//
// ⚠️ You must call [ACTCTX.SetCbSize] to initialize the struct.
//
// Example:
//
//	var ac win.ACTCTX
//	ac.SetCbSize()
//
// [ACTCTX]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/ns-winbase-actctxw
type ACTCTX struct {
	cbSize                 uint32
	DwFlags                co.ACTCTX
	LpSource               *uint16 // Convert to/from string with [wstr.DecodePtr] and [wstr.EncodeToPtr].
	WProcessorArchitecture co.PROCESSOR_ARCHITECTURE
	WLangId                LANGID
	LpAssemblyDirectory    *uint16 // Convert to/from string with [wstr.DecodePtr] and [wstr.EncodeToPtr].
	LpResourceName         *uint16 // Convert to/from string with [wstr.DecodePtr] and [wstr.EncodeToPtr].
	LpApplicationName      *uint16 // Convert to/from string with [wstr.DecodePtr] and [wstr.EncodeToPtr].
	HModule                HINSTANCE
}

// Sets the internal cbSize field to the size of the struct, correctly
// initializing it.
func (ac *ACTCTX) SetCbSize() {
	ac.cbSize = uint32(unsafe.Sizeof(*ac))
}

// An [atom].
//
// [atom]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#atom
type ATOM uint16

// [BOOL] type, which represents a boolean value as an int32, with false=0 and true=1.
//
// [BOOL]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#bool
type BOOL int32

// Converts the int32 value to bool.
func (b BOOL) Ok() bool {
	return b != 0
}

// Sets the int32 value with a bool.
func (b *BOOL) Set(v bool) {
	if v {
		*b = BOOL(1)
	} else {
		*b = BOOL(0)
	}
}

// [CONSOLE_CURSOR_INFO] struct, with C memory layout.
//
// [CONSOLE_CURSOR_INFO]: https://learn.microsoft.com/en-us/windows/console/console-cursor-info-str
type CONSOLE_CURSOR_INFO struct {
	DwSize   uint32
	BVisible BOOL
}

// [CONSOLE_FONT_INFO] struct, with C memory layout.
//
// [CONSOLE_FONT_INFO]: https://learn.microsoft.com/en-us/windows/console/console-font-info-str
type CONSOLE_FONT_INFO struct {
	NFont      uint32
	DwFontSize COORD
}

// [CONSOLE_FONT_INFOEX] struct, with C memory layout.
//
// ⚠️ You must call [CONSOLE_FONT_INFOEX.SetCbSize] to initialize the struct.
//
// Example:
//
//	var cfix win.CONSOLE_FONT_INFOEX
//	cfix.SetCbSize()
//
// [CONSOLE_FONT_INFOEX]: https://learn.microsoft.com/en-us/windows/console/console-font-infoex
type CONSOLE_FONT_INFOEX struct {
	cbSize     uint32
	NFont      uint32
	DwFontSize COORD
	fontFamily uint32 // combination of co.TMPF and co.FF
	FontWeight uint32
	faceName   [utl.LF_FACESIZE]uint16
}

// Sets the internal cbSize field to the size of the struct, correctly
// initializing it.
func (cfix *CONSOLE_FONT_INFOEX) SetCbSize() {
	cfix.cbSize = uint32(unsafe.Sizeof(*cfix))
}

func (ac *CONSOLE_FONT_INFOEX) Pitch() co.TMPF {
	return co.TMPF(ac.fontFamily & 0b1111)
}
func (ac *CONSOLE_FONT_INFOEX) SetPitch(val co.TMPF) {
	ac.fontFamily &^= 0b1111 // clear bits
	ac.fontFamily |= uint32(val & 0b1111)
}

func (ac *CONSOLE_FONT_INFOEX) Family() co.FF {
	return co.FF(ac.fontFamily & 0b1111_0000)
}
func (ac *CONSOLE_FONT_INFOEX) SetFamily(val co.FF) {
	ac.fontFamily &^= 0b1111_0000 // clear bits
	ac.fontFamily |= uint32(val & 0b1111_0000)
}

func (ac *CONSOLE_FONT_INFOEX) FaceName() string {
	return wstr.DecodeSlice(ac.faceName[:])
}
func (ac *CONSOLE_FONT_INFOEX) SetFaceName(val string) {
	wstr.EncodeToBuf(ac.faceName[:], val)
}

// [CONSOLE_READCONSOLE_CONTROL] struct, with C memory layout.
//
// ⚠️ You must call [CONSOLE_READCONSOLE_CONTROL.SetNLength] to initialize the
// struct.
//
// Example:
//
//	var crc win.CONSOLE_READCONSOLE_CONTROL
//	crc.SetNLength()
//
// [CONSOLE_READCONSOLE_CONTROL]: https://learn.microsoft.com/en-us/windows/console/console-readconsole-control
type CONSOLE_READCONSOLE_CONTROL struct {
	nLength           uint32
	NInitialChars     uint32
	DwCtrlWakeupMask  uint32
	DwControlKeyState co.CKS
}

// Sets the internal nLength field to the size of the struct, correctly
// initializing it.
func (c *CONSOLE_READCONSOLE_CONTROL) SetNLength() {
	c.nLength = uint32(unsafe.Sizeof(*c))
}

// [COORD] struct, with C memory layout.
//
// [COORD]: https://learn.microsoft.com/en-us/windows/console/coord-str
type COORD struct {
	X, Y int16
}

// [FILETIME] struct, with C memory layout.
//
// [FILETIME]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-filetime
type FILETIME struct {
	dwLowDateTime  uint32
	dwHighDateTime uint32
}

// Returns the internal value converted to epoch in 100-nanoseconds unit.
func (ft *FILETIME) ToEpochNano100() int {
	return int(utl.Make64(ft.dwLowDateTime, ft.dwHighDateTime))
}

// Replaces the internal value with the given epoch in 100-nanoseconds unit.
func (ft *FILETIME) SetEpochNano100(val int) {
	ft.dwLowDateTime, ft.dwHighDateTime = utl.Break64(uint64(val))
}

// Returns the internal value converted to [time.Time].
func (ft *FILETIME) ToTime() time.Time {
	// https://stackoverflow.com/a/4135003
	return time.Unix(0, int64(utl.Make64(ft.dwLowDateTime, ft.dwHighDateTime)-116_444_736_000_000_000)*100)
}

// Replaces the internal value with the given [time.Time].
//
// Example:
//
//	var ft win.FILETIME
//	ft.SetTime(time.Now())
func (ft *FILETIME) SetTime(val time.Time) {
	ft.dwLowDateTime, ft.dwHighDateTime = utl.Break64(
		uint64(val.UnixNano()/100 + 116_444_736_000_000_000),
	)
}

// [HEAP_OPTIMIZE_RESOURCES_INFORMATION] struct, with C memory layout.
//
// ⚠️ You must call [HEAP_OPTIMIZE_RESOURCES_INFORMATION.SetVersion] to initialize the struct.
//
// Example:
//
//	var ho win.HEAP_OPTIMIZE_RESOURCES_INFORMATION
//	ho.SetVersion()
//
// [HEAP_OPTIMIZE_RESOURCES_INFORMATION]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-heap_optimize_resources_information
type HEAP_OPTIMIZE_RESOURCES_INFORMATION struct {
	version uint32
	Flags   uint32 // Must be set to zero.
}

// Sets the internal version field to the size of the struct, correctly
// initializing it.
func (ho *HEAP_OPTIMIZE_RESOURCES_INFORMATION) SetVersion() {
	ho.version = utl.HEAP_OPTIMIZE_RESOURCES_CURRENT_VERSION
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

// [MEMORY_BASIC_INFORMATION] struct, with C memory layout.
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

// [MEMORYSTATUSEX] struct, with C memory layout.
//
// ⚠️ You must call [MEMORYSTATUSEX.SetDwLength] to initialize the struct.
//
// Example:
//
//	var ms win.MEMORYSTATUSEX
//	ms.SetDwLength()
//
// [MEMORYSTATUSEX]: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-memorystatusex
type MEMORYSTATUSEX struct {
	dwLength                uint32
	MemoryLoad              uint32 // A number between 0 and 100 that specifies the approximate percentage of physical memory that is in use.
	TotalPhys               uint64 // The amount of actual physical memory, in bytes.
	AvailPhys               uint64 // The amount of physical memory currently available, in bytes. This is the amount of physical memory that can be immediately reused without having to write its contents to disk first. It is the sum of the size of the standby, free, and zero lists.
	TotalPageFile           uint64 // The current committed memory limit for the system or the current process, whichever is smaller, in bytes.
	AvailPageFile           uint64 // The maximum amount of memory the current process can commit, in bytes. This value is equal to or smaller than the system-wide available commit value.
	TotalVirtual            uint64 // The size of the user-mode portion of the virtual address space of the calling process, in bytes. This value depends on the type of process, the type of processor, and the configuration of the operating system.
	AvailVirtual            uint64 // The amount of unreserved and uncommitted memory currently in the user-mode portion of the virtual address space of the calling process, in bytes.
	ullAvailExtendedVirtual uint64
}

// Sets the internal dwLength field to the size of the struct, correctly
// initializing it.
func (ms *MEMORYSTATUSEX) SetDwLength() {
	ms.dwLength = uint32(unsafe.Sizeof(*ms))
}

// [MODULEENTRY32] struct, with C memory layout.
//
// ⚠️ You must call [MODULEENTRY32.SetDwSize] to initialize the struct.
//
// Example:
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

// Sets the internal dwSize field to the size of the struct, correctly
// initializing it.
func (me *MODULEENTRY32) SetDwSize() {
	me.dwSize = uint32(unsafe.Sizeof(*me))
}

func (me *MODULEENTRY32) SzModule() string {
	return wstr.DecodeSlice(me.szModule[:])
}
func (me *MODULEENTRY32) SetSzModule(val string) {
	wstr.EncodeToBuf(me.szModule[:], val)
}

func (me *MODULEENTRY32) SzExePath() string {
	return wstr.DecodeSlice(me.szExePath[:])
}
func (me *MODULEENTRY32) SetSzExePath(val string) {
	wstr.EncodeToBuf(me.szExePath[:], val)
}

// [OSVERSIONINFOEX] struct, with C memory layout.
//
// ⚠️ You must call [OSVERSIONINFOEX.SetDwOsVersionInfoSize] to initialize the
// struct.
//
// Example:
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
	return wstr.DecodeSlice(osv.szCSDVersion[:])
}
func (osv *OSVERSIONINFOEX) SetSzCSDVersion(val string) {
	wstr.EncodeToBuf(osv.szCSDVersion[:], val)
}

// [OVERLAPPED] struct, with C memory layout.
//
// [OVERLAPPED]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-overlapped
type OVERLAPPED struct {
	Internal     uintptr
	InternalHigh uintptr
	Pointer      uintptr
	HEvent       uintptr // HEVENT
}

// [PACKAGE_ID] struct, with C memory layout.
//
// [PACKAGE_ID]: https://learn.microsoft.com/en-us/windows/win32/api/appmodel/ns-appmodel-package_id
type PACKAGE_ID struct {
	reserved              uint32
	ProcessorArchitecture co.PROCESSOR_ARCH_ENUM
	Version               PACKAGE_VERSION
	name                  *uint16
	publisher             *uint16
	resourceId            *uint16
	publisherId           *uint16
}

func (pi *PACKAGE_ID) Name() string {
	return wstr.DecodePtr(pi.name)
}
func (pi *PACKAGE_ID) Publisher() string {
	return wstr.DecodePtr(pi.publisher)
}
func (pi *PACKAGE_ID) ResourceId() string {
	return wstr.DecodePtr(pi.resourceId)
}
func (pi *PACKAGE_ID) PublisherId() string {
	return wstr.DecodePtr(pi.publisherId)
}

// [PACKAGE_VERSION] struct, with C memory layout.
//
// [PACKAGE_VERSION]: https://learn.microsoft.com/en-us/windows/win32/api/appmodel/ns-appmodel-package_version
type PACKAGE_VERSION struct {
	Revision uint16
	Build    uint16
	Minor    uint16
	Major    uint16
}

// [PROCESS_INFORMATION] struct, with C memory layout.
//
// [PROCESS_INFORMATION]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-process_information
type PROCESS_INFORMATION struct {
	HProcess    HPROCESS
	HThread     HTHREAD
	DwProcessId uint32
	DwThreadId  uint32
}

// [PROCESSENTRY32] struct, with C memory layout.
//
// ⚠️ You must call [PROCESSENTRY32.SetDwSize] to initialize the struct.
//
// Example:
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

// Sets the internal dwSize field to the size of the struct, correctly
// initializing it.
func (pe *PROCESSENTRY32) SetDwSize() {
	pe.dwSize = uint32(unsafe.Sizeof(*pe))
}

func (me *PROCESSENTRY32) SzExeFile() string {
	return wstr.DecodeSlice(me.szExeFile[:])
}
func (me *PROCESSENTRY32) SetSzExeFile(val string) {
	wstr.EncodeToBuf(me.szExeFile[:], val)
}

// [PROCESSOR_NUMBER] struct, with C memory layout.
//
// [PROCESSOR_NUMBER]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-processor_number
type PROCESSOR_NUMBER struct {
	Group    uint16
	Number   uint8
	reserved uint8
}

// [SECURITY_ATTRIBUTES] struct, with C memory layout.
//
// ⚠️ You must call [SECURITY_ATTRIBUTES.SetNLength] to initialize the struct.
//
// Example:
//
//	var sa win.SECURITY_ATTRIBUTES
//	sa.SetNLength()
//
// [SECURITY_ATTRIBUTES]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/aa379560(v=vs.85)
type SECURITY_ATTRIBUTES struct {
	nLength              uint32
	LpSecurityDescriptor uintptr // LPVOID
	BInheritHandle       BOOL
}

// Sets the internal nLength field to the size of the struct, correctly
// initializing it.
func (sa *SECURITY_ATTRIBUTES) SetNLength() {
	sa.nLength = uint32(unsafe.Sizeof(*sa))
}

// [STARTUPINFO] struct, with C memory layout.
//
// ⚠️ You must call [STARTUPINFO.SetCb] to initialize the struct.
//
// Example:
//
//	var si win.STARTUPINFO
//	si.SetCb()
//
// [STARTUPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-startupinfow
type STARTUPINFO struct {
	cb              uint32
	lpReserved      *uint16
	LpDesktop       *uint16 // Convert to/from string with [wstr.DecodePtr] and [wstr.EncodeToPtr].
	LpTitle         *uint16 // Convert to/from string with [wstr.DecodePtr] and [wstr.EncodeToPtr].
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

// Sets the internal cb field to the size of the struct, correctly
// initializing it.
func (si *STARTUPINFO) SetCb() {
	si.cb = uint32(unsafe.Sizeof(*si))
}

// [SYSTEM_INFO] struct, with C memory layout.
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

// [SYSTEMTIME] struct, with C memory layout.
//
// [SYSTEMTIME]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-systemtime
type SYSTEMTIME struct {
	Year         uint16 // The valid values for this member are 1601 through 30827.
	Month        uint16 // January=1, December=12.
	DayOfWeek    uint16 // Sunday=0, Saturday=6.
	Day          uint16 // The valid values for this member are 1 through 31.
	Hour         uint16 // The valid values for this member are 0 through 23.
	Minute       uint16 // The valid values for this member are 0 through 59.
	Second       uint16 // The valid values for this member are 0 through 59.
	Milliseconds uint16 // The valid values for this member are 0 through 999.
}

// Decomposes a [time.Duration] into this SYSTEMTIME fields.
//
// Example:
//
//	var st win.SYSTEMTIME
//	st.SetDuration(time.Minute * 3)
func (st *SYSTEMTIME) SetDuration(dur time.Duration) {
	*st = SYSTEMTIME{}
	st.Hour = uint16(dur / time.Hour)
	st.Minute = uint16((dur -
		time.Duration(st.Hour)*time.Hour) / time.Minute)
	st.Second = uint16((dur -
		time.Duration(st.Hour)*time.Hour -
		time.Duration(st.Minute)*time.Minute) / time.Second)
	st.Milliseconds = uint16((dur -
		time.Duration(st.Hour)*time.Hour -
		time.Duration(st.Minute)*time.Minute -
		time.Duration(st.Second)*time.Second) / time.Millisecond)
}

// Converts this SYSTEMTIME to [time.Time].
//
// Example:
//
//	st := win.GetLocalTime()
//	t := st.ToTime()
func (st *SYSTEMTIME) ToTime() time.Time {
	return time.Date(int(st.Year), time.Month(st.Month), int(st.Day),
		int(st.Hour), int(st.Minute), int(st.Second),
		int(st.Milliseconds)*1_000_000,
		time.Local)
}

// Fills this SYSTEMTIME with the value of a [time.Time], in the current
// timezone.
//
// Example:
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

// [THREADENTRY32] struct, with C memory layout.
//
// ⚠️ You must call [THREADENTRY32.SetDwSize] to initialize the struct.
//
// Example:
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

// Sets the internal dwSize field to the size of the struct, correctly
// initializing it.
func (te *THREADENTRY32) SetDwSize() {
	te.dwSize = uint32(unsafe.Sizeof(*te))
}

// [TIME_ZONE_INFORMATION] struct, with C memory layout.
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
	return wstr.DecodeSlice(tzi.standardName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetStandardName(val string) {
	wstr.EncodeToBuf(tzi.standardName[:], val)
}

func (tzi *TIME_ZONE_INFORMATION) DaylightName() string {
	return wstr.DecodeSlice(tzi.daylightName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetDaylightName(val string) {
	wstr.EncodeToBuf(tzi.daylightName[:], val)
}

// [WIN32_FIND_DATA] struct, with C memory layout.
//
// [WIN32_FIND_DATA]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-win32_find_dataw
type WIN32_FIND_DATA struct {
	DwFileAttributes   co.FILE_ATTRIBUTE
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	nFileSizeHigh      uint32
	nFileSizeLow       uint32
	dwReserved0        uint32
	dwReserved1        uint32
	cFileName          [utl.MAX_PATH]uint16
	cAlternateFileName [14]uint16
	dwFileType         uint32
	dwCreatorType      uint32
	wFinderFlags       uint16
}

func (wfd *WIN32_FIND_DATA) CFileName() string {
	return wstr.DecodeSlice(wfd.cFileName[:])
}
func (wfd *WIN32_FIND_DATA) SetCFileName(val string) {
	wstr.EncodeToBuf(wfd.cFileName[:], val)
}

func (wfd *WIN32_FIND_DATA) NFileSize() int {
	return int(utl.Make64(wfd.nFileSizeLow, wfd.nFileSizeHigh))
}
func (wfd *WIN32_FIND_DATA) SetNFileSize(val int) {
	lo, hi := utl.Break64(uint64(val))
	wfd.nFileSizeLow = lo
	wfd.nFileSizeHigh = hi
}

func (wfd *WIN32_FIND_DATA) CAlternateFileName() string {
	return wstr.DecodeSlice(wfd.cAlternateFileName[:])
}
func (wfd *WIN32_FIND_DATA) SetCAlternateFileName(val string) {
	wstr.EncodeToBuf(wfd.cAlternateFileName[:], val)
}
