//go:build windows

package co

// [Registry key] security and access rights.
//
// [Registry key]: https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-key-security-and-access-rights
type KEY uint32

const (
	// Required to query the values of a registry key.
	KEY_QUERY_VALUE KEY = 0x0001
	// Required to create, delete, or set a registry value.
	KEY_SET_VALUE KEY = 0x0002
	// Required to create a subkey of a registry key.
	KEY_CREATE_SUB_KEY KEY = 0x0004
	// Required to enumerate the subkeys of a registry key.
	KEY_ENUMERATE_SUB_KEYS KEY = 0x0008
	// Required to request change notifications for a registry key or for
	// subkeys of a registry key.
	KEY_NOTIFY KEY = 0x0010
	// Reserved for system use.
	KEY_CREATE_LINK KEY = 0x0020
	// Indicates that an application on 64-bit Windows should operate on the
	// 32-bit registry view.
	KEY_WOW64_32KEY KEY = 0x0200
	// Indicates that an application on 64-bit Windows should operate on the
	// 64-bit registry view.
	KEY_WOW64_64KEY KEY = 0x0100
	// Undocumented flag.
	KEY_WOW64_RES KEY = 0x0300

	// Combines the STANDARD_RIGHTS_READ, KEY_QUERY_VALUE,
	// KEY_ENUMERATE_SUB_KEYS, and KEY_NOTIFY values.
	KEY_READ = (KEY(STANDARD_RIGHTS_READ) | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY) & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
	// Combines the STANDARD_RIGHTS_WRITE, KEY_SET_VALUE, and KEY_CREATE_SUB_KEY
	// access rights.
	KEY_WRITE = (KEY(STANDARD_RIGHTS_WRITE) | KEY_SET_VALUE | KEY_CREATE_SUB_KEY) & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
	// Equivalent to KEY_READ.
	KEY_EXECUTE = KEY_READ & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
	// Combines the STANDARD_RIGHTS_REQUIRED, KEY_QUERY_VALUE, KEY_SET_VALUE,
	// KEY_CREATE_SUB_KEY, KEY_ENUMERATE_SUB_KEYS, KEY_NOTIFY, and
	// KEY_CREATE_LINK access rights.
	KEY_ALL_ACCESS = (KEY(STANDARD_RIGHTS_ALL) | KEY_QUERY_VALUE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY | KEY_CREATE_LINK) & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
)

// Registry value [types].
//
// [types]: https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-value-types
type REG uint32

const (
	REG_NONE                       REG = 0
	REG_SZ                         REG = 1
	REG_EXPAND_SZ                  REG = 2
	REG_BINARY                     REG = 3
	REG_DWORD                      REG = 4
	REG_DWORD_LITTLE_ENDIAN        REG = 4
	REG_DWORD_BIG_ENDIAN           REG = 5
	REG_LINK                       REG = 6
	REG_MULTI_SZ                   REG = 7
	REG_RESOURCE_LIST              REG = 8
	REG_FULL_RESOURCE_DESCRIPTOR   REG = 9
	REG_RESOURCE_REQUIREMENTS_LIST REG = 10
	REG_QWORD                      REG = 11
	REG_QWORD_LITTLE_ENDIAN        REG = 11
)

// [RegCreateKeyEx] and [RegOpenKeyEx] options.
//
// [RegCreateKeyEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regcreatekeyexw
// [RegOpenKeyEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regopenkeyexw
type REG_OPTION uint32

const (
	REG_OPTION_NONE            REG_OPTION = 0
	REG_OPTION_RESERVED        REG_OPTION = 0x0000_0000
	REG_OPTION_NON_VOLATILE    REG_OPTION = 0x0000_0000
	REG_OPTION_VOLATILE        REG_OPTION = 0x0000_0001
	REG_OPTION_CREATE_LINK     REG_OPTION = 0x0000_0002
	REG_OPTION_BACKUP_RESTORE  REG_OPTION = 0x0000_0004
	REG_OPTION_OPEN_LINK       REG_OPTION = 0x0000_0008
	REG_OPTION_DONT_VIRTUALIZE REG_OPTION = 0x0000_0010
)

// [RegRestoreKey] flags. Originally has REG prefix.
//
// [RegRestoreKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regrestorekeyw
type REG_RESTORE uint32

const (
	REG_RESTORE_FORCE               REG_RESTORE = 0x0000_0008
	REG_RESTORE_WHOLE_HIVE_VOLATILE REG_RESTORE = 0x0000_0001
)

// [RegSaveKeyEx] flags. Originally has REG prefix.
//
// [RegSaveKeyEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsavekeyexw
type REG_SAVE uint32

const (
	REG_SAVE_STANDARD_FORMAT REG_SAVE = 1
	REG_SAVE_LATEST_FORMAT   REG_SAVE = 2
	REG_SAVE_NO_COMPRESSION  REG_SAVE = 4
)

// [RegGetValue] flags.
//
// [RegGetValue]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-reggetvaluew
type RRF uint32

const (
	RRF_RT_REG_NONE      RRF = 0x0000_0001
	RRF_RT_REG_SZ        RRF = 0x0000_0002
	RRF_RT_REG_EXPAND_SZ RRF = 0x0000_0004
	RRF_RT_REG_BINARY    RRF = 0x0000_0008
	RRF_RT_REG_DWORD     RRF = 0x0000_0010
	RRF_RT_REG_MULTI_SZ  RRF = 0x0000_0020
	RRF_RT_REG_QWORD     RRF = 0x0000_0040
	RRF_RT_DWORD             = RRF_RT_REG_BINARY | RRF_RT_REG_DWORD
	RRF_RT_QWORD             = RRF_RT_REG_BINARY | RRF_RT_REG_QWORD
	RRF_RT_ANY           RRF = 0x0000_ffff

	RRF_SUBKEY_WOW6464KEY RRF = 0x0001_0000
	RRF_SUBKEY_WOW6432KEY RRF = 0x0002_0000
	RRF_WOW64_MASK        RRF = 0x0003_0000

	RRF_NOEXPAND      RRF = 0x1000_0000
	RRF_ZEROONFAILURE RRF = 0x2000_0000
)
