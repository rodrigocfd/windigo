//go:build windows

package win

import (
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// [CONSOLE_CURSOR_INFO] struct.
//
// [CONSOLE_CURSOR_INFO]: https://learn.microsoft.com/en-us/windows/console/console-cursor-info-str
type CONSOLE_CURSOR_INFO struct {
	DwSize   uint32
	bVisible int32 // BOOL
}

func (cc *CONSOLE_CURSOR_INFO) BVisible() bool       { return cc.bVisible != 0 }
func (cc *CONSOLE_CURSOR_INFO) SetBVisible(val bool) { cc.bVisible = util.BoolToInt32(val) }

// [CONSOLE_FONT_INFO] struct.
//
// [CONSOLE_FONT_INFO]: https://learn.microsoft.com/en-us/windows/console/console-font-info-str
type CONSOLE_FONT_INFO struct {
	NFont      uint32
	DwFontSize COORD
}

// [CONSOLE_READCONSOLE_CONTROL] struct.
//
// ⚠️ You must call SetNLength() to initialize the struct.
//
// [CONSOLE_READCONSOLE_CONTROL]: https://learn.microsoft.com/en-us/windows/console/console-readconsole-control
type CONSOLE_READCONSOLE_CONTROL struct {
	nLength           uint32
	NInitialChars     uint32
	DwCtrlWakeupMask  uint32
	DwControlKeyState co.CKS
}

func (c *CONSOLE_READCONSOLE_CONTROL) SetNLength() { c.nLength = uint32(unsafe.Sizeof(*c)) }

// [COORD] struct.
//
// [COORD]: https://learn.microsoft.com/en-us/windows/console/coord-str
type COORD struct {
	X, Y int16
}

// [DYNAMIC_TIME_ZONE_INFORMATION] struct.
//
// [DYNAMIC_TIME_ZONE_INFORMATION]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/ns-timezoneapi-dynamic_time_zone_information
type DYNAMIC_TIME_ZONE_INFORMATION struct {
	Bias                        int32
	standardName                [32]uint16
	StandardDate                SYSTEMTIME
	StandardBias                int32
	daylightName                [32]uint16
	DaylightDate                SYSTEMTIME
	DaylightBias                int32
	timeZoneKeyName             [128]uint16
	dynamicDaylightTimeDisabled uint8 // BOOLEAN
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) StandardName() string {
	return Str.FromNativeSlice(dtz.standardName[:])
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetStandardName(val string) {
	copy(dtz.standardName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dtz.standardName)-1)))
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) DaylightName() string {
	return Str.FromNativeSlice(dtz.daylightName[:])
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetDaylightName(val string) {
	copy(dtz.daylightName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dtz.daylightName)-1)))
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) TimeZoneKeyName() string {
	return Str.FromNativeSlice(dtz.timeZoneKeyName[:])
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetTimeZoneKeyName(val string) {
	copy(dtz.timeZoneKeyName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dtz.timeZoneKeyName)-1)))
}

func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) DynamicDaylightTimeDisabled() bool {
	return dtz.dynamicDaylightTimeDisabled != 0
}
func (dtz *DYNAMIC_TIME_ZONE_INFORMATION) SetDynamicDaylightTimeDisabled(val bool) {
	dtz.dynamicDaylightTimeDisabled = util.BoolToUint8(val)
}

// [FILETIME] struct.
//
// Can be converted to [SYSTEMTIME] with [FileTimeToSystemTime] function.
//
// [FILETIME]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-filetime
// [SYSTEMTIME]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-systemtime
// [FileTimeToSystemTime]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-filetimetosystemtime
type FILETIME struct {
	dwLowDateTime  uint32
	dwHighDateTime uint32
}

// Returns the internal value converted to epoch in 100-nanoseconds unit.
func (ft *FILETIME) EpochNano100() uint64 { return util.Make64(ft.dwLowDateTime, ft.dwHighDateTime) }

// Replaces the internal value with the given epoch in 100-nanoseconds unit.
func (ft *FILETIME) SetEpochNano100(val uint64) {
	ft.dwLowDateTime, ft.dwHighDateTime = util.Break64(val)
}

// Returns the internal value converted to [time.Time].
func (ft *FILETIME) ToTime() time.Time {
	// https://stackoverflow.com/a/4135003/6923555
	return time.Unix(0, int64(util.Make64(ft.dwLowDateTime, ft.dwHighDateTime)-116_444_736_000_000_000)*100)
}

// Replaces the internal value with the given [time.Time].
func (ft *FILETIME) FromTime(val time.Time) {
	ft.dwLowDateTime, ft.dwHighDateTime = util.Break64(
		uint64(val.UnixNano()/100 + 116_444_736_000_000_000),
	)
}

// [MANAGEDAPPLICATION] struct.
//
// [MANAGEDAPPLICATION]: https://learn.microsoft.com/en-us/windows/win32/api/appmgmt/ns-appmgmt-managedapplication
type MANAGEDAPPLICATION struct {
	pszPackageName *uint16
	pszPublisher   *uint16
	DwVersionHi    uint32
	DwVersionLo    uint32
	DwRevision     uint32
	GpoId          GUID
	pszPolicyName  *uint16
	ProductId      GUID
	Language       LANGID
	pszOwner       *uint16
	pszCompany     *uint16
	pszComments    *uint16
	pszContact     *uint16
	pszSupportUrl  *uint16
	DwPathType     uint32
	bInstalled     int32 // BOOL
}

func (ma *MANAGEDAPPLICATION) PszPackageName() string { return Str.FromNativePtr(ma.pszPackageName) }
func (ma *MANAGEDAPPLICATION) PszPublisher() string   { return Str.FromNativePtr(ma.pszPublisher) }
func (ma *MANAGEDAPPLICATION) PszPolicyName() string  { return Str.FromNativePtr(ma.pszPolicyName) }
func (ma *MANAGEDAPPLICATION) PszOwner() string       { return Str.FromNativePtr(ma.pszOwner) }
func (ma *MANAGEDAPPLICATION) PszCompany() string     { return Str.FromNativePtr(ma.pszCompany) }
func (ma *MANAGEDAPPLICATION) PszComments() string    { return Str.FromNativePtr(ma.pszComments) }
func (ma *MANAGEDAPPLICATION) PszContact() string     { return Str.FromNativePtr(ma.pszContact) }
func (ma *MANAGEDAPPLICATION) PszSupportUrl() string  { return Str.FromNativePtr(ma.pszSupportUrl) }
func (ma *MANAGEDAPPLICATION) BInstalled() bool       { return ma.bInstalled != 0 }

// [MODULEENTRY32] struct.
//
// ⚠️ You must call SetDwSize() to initialize the struct.
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
	szModule      [_MAX_MODULE_NAME32 + 1]uint16
	szExePath     [_MAX_PATH]uint16
}

func (me *MODULEENTRY32) SetDwSize() { me.dwSize = uint32(unsafe.Sizeof(*me)) }

func (me *MODULEENTRY32) SzModule() string { return Str.FromNativeSlice(me.szModule[:]) }
func (me *MODULEENTRY32) SetSzModule(val string) {
	copy(me.szModule[:], Str.ToNativeSlice(Str.Substr(val, 0, len(me.szModule)-1)))
}

func (me *MODULEENTRY32) SzExePath() string { return Str.FromNativeSlice(me.szExePath[:]) }
func (me *MODULEENTRY32) SetSzExePath(val string) {
	copy(me.szExePath[:], Str.ToNativeSlice(Str.Substr(val, 0, len(me.szExePath)-1)))
}

// [OSVERSIONINFOEX] struct.
//
// ⚠️ You must call SetDwOsVersionInfoSize() to initialize the struct.
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

func (osv *OSVERSIONINFOEX) SetDwOsVersionInfoSize() {
	osv.dwOsVersionInfoSize = uint32(unsafe.Sizeof(*osv))
}

func (osv *OSVERSIONINFOEX) SzCSDVersion() string { return Str.FromNativeSlice(osv.szCSDVersion[:]) }
func (osv *OSVERSIONINFOEX) SetSzCSDVersion(val string) {
	copy(osv.szCSDVersion[:], Str.ToNativeSlice(Str.Substr(val, 0, len(osv.szCSDVersion)-1)))
}

// [OVERLAPPED] struct.
//
// [OVERLAPPED]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-overlapped
type OVERLAPPED struct {
	Internal     uintptr
	InternalHigh uintptr
	Pointer      uintptr
	HEvent       HEVENT
}

// [PROCESSENTRY32] struct.
//
// ⚠️ You must call SetDwSize() to initialize the struct.
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
	szExeFile           [_MAX_PATH]uint16
}

func (pe *PROCESSENTRY32) SetDwSize() { pe.dwSize = uint32(unsafe.Sizeof(*pe)) }

func (me *PROCESSENTRY32) SzExeFile() string { return Str.FromNativeSlice(me.szExeFile[:]) }
func (me *PROCESSENTRY32) SetSzExeFile(val string) {
	copy(me.szExeFile[:], Str.ToNativeSlice(Str.Substr(val, 0, len(me.szExeFile)-1)))
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

// [SECURITY_ATTRIBUTES] struct.
//
// ⚠️ You must call SetNLength() to initialize the struct.
//
// [SECURITY_ATTRIBUTES]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/aa379560(v=vs.85)
type SECURITY_ATTRIBUTES struct {
	nLength              uint32
	LpSecurityDescriptor uintptr // LPVOID
	bInheritHandle       int32   // BOOL
}

func (sa *SECURITY_ATTRIBUTES) SetNLength() { sa.nLength = uint32(unsafe.Sizeof(*sa)) }

func (sa *SECURITY_ATTRIBUTES) BInheritHandle() bool       { return sa.bInheritHandle != 0 }
func (sa *SECURITY_ATTRIBUTES) SetBInheritHandle(val bool) { sa.bInheritHandle = util.BoolToInt32(val) }

// [STARTUPINFO] struct.
//
// ⚠️ You must call SetCb() to initialize the struct.
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
	DwFillAttribute uint32
	DwFlags         co.STARTF
	WShowWindow     uint16 // co.SW, should be uint16.
	cbReserved2     uint16
	lpReserved2     *uint8
	HStdInput       uintptr
	HStdOutput      uintptr
	HStdError       uintptr
}

func (si *STARTUPINFO) SetCb() { si.cb = uint32(unsafe.Sizeof(*si)) }

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
// Can be converted to [FILETIME] with [SystemTimeToFileTime] function.
//
// [SYSTEMTIME]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-systemtime
// [FILETIME]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-filetime
// [SystemTimeToFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-systemtimetofiletime
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

// Decomposes a time.Duration into this SYSTEMTIME fields.
func (st *SYSTEMTIME) FromDuration(dur time.Duration) {
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

// Fills this SYSTEMTIME with the value of a time.Time.
func (st *SYSTEMTIME) FromTime(val time.Time) {
	var ft FILETIME
	ft.FromTime(val)

	var stUtc SYSTEMTIME
	FileTimeToSystemTime(&ft, &stUtc)

	// When converted, SYSTEMTIME will receive UTC values, so we need to convert
	// the fields to current timezone.
	SystemTimeToTzSpecificLocalTime(nil, &stUtc, st)
}

// Converts this SYSTEMTIME to time.Time.
func (st *SYSTEMTIME) ToTime() time.Time {
	return time.Date(int(st.WYear),
		time.Month(st.WMonth), int(st.WDay),
		int(st.WHour), int(st.WMinute), int(st.WSecond),
		int(st.WMilliseconds)*1_000_000,
		time.Local)
}

// [THREADENTRY32] struct.
//
// ⚠️ You must call SetDwSize() to initialize the struct.
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

func (te *THREADENTRY32) SetDwSize() { te.dwSize = uint32(unsafe.Sizeof(*te)) }

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
	return Str.FromNativeSlice(tzi.standardName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetStandardName(val string) {
	copy(tzi.standardName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(tzi.standardName)-1)))
}

func (tzi *TIME_ZONE_INFORMATION) DaylightName() string {
	return Str.FromNativeSlice(tzi.daylightName[:])
}
func (tzi *TIME_ZONE_INFORMATION) SetDaylightName(val string) {
	copy(tzi.daylightName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(tzi.daylightName)-1)))
}

// [WIN32_FIND_DATA] struct.
//
// [WIN32_FIND_DATA]: https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-win32_find_dataw
type WIN32_FIND_DATA struct {
	DwFileAttributes    co.FILE_ATTRIBUTE
	FtCreationTime      FILETIME
	FtLastAccessTime    FILETIME
	FtLastWriteTime     FILETIME
	NFileSizeHigh       uint32
	NFileSizeLow        uint32
	dwReserved0         uint32
	dwReserved1         uint32
	cFileName           [_MAX_PATH]uint16
	cCAlternateFileName [14]uint16
	DwFileType          uint32
	DwCreatorType       uint32
	WFinderFlags        uint16
}

func (wfd *WIN32_FIND_DATA) CFileName() string { return Str.FromNativeSlice(wfd.cFileName[:]) }
func (wfd *WIN32_FIND_DATA) SetCFileName(val string) {
	copy(wfd.cFileName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(wfd.cFileName)-1)))
}

func (wfd *WIN32_FIND_DATA) CAlternateFileName() string {
	return Str.FromNativeSlice(wfd.cCAlternateFileName[:])
}
func (wfd *WIN32_FIND_DATA) SetCAlternateFileName(val string) {
	copy(wfd.cCAlternateFileName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(wfd.cCAlternateFileName)-1)))
}
