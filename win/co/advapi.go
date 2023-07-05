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
	KEY_READ KEY = (KEY(STANDARD_RIGHTS_READ) | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY) & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
	// Combines the STANDARD_RIGHTS_WRITE, KEY_SET_VALUE, and KEY_CREATE_SUB_KEY
	// access rights.
	KEY_WRITE KEY = (KEY(STANDARD_RIGHTS_WRITE) | KEY_SET_VALUE | KEY_CREATE_SUB_KEY) & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
	// Equivalent to KEY_READ.
	KEY_EXECUTE KEY = KEY_READ & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
	// Combines the STANDARD_RIGHTS_REQUIRED, KEY_QUERY_VALUE, KEY_SET_VALUE,
	// KEY_CREATE_SUB_KEY, KEY_ENUMERATE_SUB_KEYS, KEY_NOTIFY, and
	// KEY_CREATE_LINK access rights.
	KEY_ALL_ACCESS KEY = (KEY(STANDARD_RIGHTS_ALL) | KEY_QUERY_VALUE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY | KEY_CREATE_LINK) & ^KEY(STANDARD_RIGHTS_SYNCHRONIZE)
)

// Registry value [types].
//
// [types]: https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-value-types
type REG uint32

const (
	REG_NONE                REG = 0  // No value type.
	REG_SZ                  REG = 1  // Unicode nul terminated string.
	REG_EXPAND_SZ           REG = 2  // Unicode nul terminated string (with environment variable references).
	REG_BINARY              REG = 3  // Free form binary.
	REG_DWORD               REG = 4  // 32-bit number.
	REG_DWORD_LITTLE_ENDIAN REG = 4  // 32-bit number (same as REG_DWORD).
	REG_DWORD_BIG_ENDIAN    REG = 5  // 32-bit number.
	REG_LINK                REG = 6  // Symbolic Link (unicode).
	REG_MULTI_SZ            REG = 7  // Multiple Unicode strings.
	REG_QWORD               REG = 11 // 64-bit number.
	REG_QWORD_LITTLE_ENDIAN REG = 11 // 64-bit number (same as REG_QWORD).
)

// [RegOpenKeyEx] ulOptions.
//
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

// [RegGetValue] dwFlags.
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
	RRF_RT_DWORD         RRF = RRF_RT_REG_BINARY | RRF_RT_REG_DWORD
	RRF_RT_QWORD         RRF = RRF_RT_REG_BINARY | RRF_RT_REG_QWORD
	RRF_RT_ANY           RRF = 0x0000_ffff

	RRF_SUBKEY_WOW6464KEY RRF = 0x0001_0000
	RRF_SUBKEY_WOW6432KEY RRF = 0x0002_0000
	RRF_NOEXPAND          RRF = 0x1000_0000
	RRF_ZEROONFAILURE     RRF = 0x2000_0000
)

// Token [access rights].
//
// [access rights]: https://learn.microsoft.com/en-us/windows/win32/secauthz/access-rights-for-access-token-objects
type TOKEN uint32

const (
	TOKEN_DELETE       TOKEN = TOKEN(STANDARD_RIGHTS_DELETE)
	TOKEN_READ_CONTROL TOKEN = TOKEN(STANDARD_RIGHTS_READ_CONTROL)
	TOKEN_WRITE_DAC    TOKEN = TOKEN(STANDARD_RIGHTS_WRITE_DAC)
	TOKEN_WRITE_OWNER  TOKEN = TOKEN(STANDARD_RIGHTS_WRITE_OWNER)

	TOKEN_ASSIGN_PRIMARY        TOKEN = 0x0001
	TOKEN_DUPLICATE             TOKEN = 0x0002
	TOKEN_IMPERSONATE           TOKEN = 0x0004
	TOKEN_QUERY                 TOKEN = 0x0008
	TOKEN_QUERY_SOURCE          TOKEN = 0x0010
	TOKEN_ADJUST_PRIVILEGES     TOKEN = 0x0020
	TOKEN_ADJUST_GROUPS         TOKEN = 0x0040
	TOKEN_ADJUST_DEFAULT        TOKEN = 0x0080
	TOKEN_ADJUST_SESSIONID      TOKEN = 0x0100
	TOKEN_ALL_ACCESS_P          TOKEN = TOKEN(STANDARD_RIGHTS_REQUIRED) | TOKEN_ASSIGN_PRIMARY | TOKEN_DUPLICATE | TOKEN_IMPERSONATE | TOKEN_QUERY | TOKEN_QUERY_SOURCE | TOKEN_ADJUST_PRIVILEGES | TOKEN_ADJUST_GROUPS | TOKEN_ADJUST_DEFAULT
	TOKEN_ALL_ACCESS            TOKEN = TOKEN_ALL_ACCESS_P | TOKEN_ADJUST_SESSIONID
	TOKEN_READ                  TOKEN = TOKEN(STANDARD_RIGHTS_READ) | TOKEN_QUERY
	TOKEN_WRITE                 TOKEN = TOKEN(STANDARD_RIGHTS_WRITE) | TOKEN_ADJUST_PRIVILEGES | TOKEN_ADJUST_GROUPS | TOKEN_ADJUST_DEFAULT
	TOKEN_EXECUTE               TOKEN = TOKEN(STANDARD_RIGHTS_EXECUTE)
	TOKEN_TRUST_CONSTRAINT_MASK TOKEN = TOKEN(STANDARD_RIGHTS_READ) | TOKEN_QUERY | TOKEN_QUERY_SOURCE
	TOKEN_ACCESS_PSEUDO_HANDLE  TOKEN = TOKEN_QUERY | TOKEN_QUERY_SOURCE
)
