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

// [TOKEN_ELEVATION_TYPE] enumeration.
//
// [TOKEN_ELEVATION_TYPE]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ne-winnt-token_elevation_type
type TOKEN_ELEVATION_TYPE uint32

const (
	TOKEN_ELEVATION_TYPE_Default TOKEN_ELEVATION_TYPE = iota + 1
	TOKEN_ELEVATION_TYPE_Full
	TOKEN_ELEVATION_TYPE_Limited
)

// [TOKEN_INFORMATION_CLASS] enumeration.
//
// [TOKEN_INFORMATION_CLASS]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ne-winnt-token_information_class
type TOKEN_INFO uint32

const (
	TOKEN_INFO_User TOKEN_INFO = iota + 1
	TOKEN_INFO_Groups
	TOKEN_INFO_Privileges
	TOKEN_INFO_Owner
	TOKEN_INFO_PrimaryGroup
	TOKEN_INFO_DefaultDacl
	TOKEN_INFO_Source
	TOKEN_INFO_Type
	TOKEN_INFO_ImpersonationLevel
	TOKEN_INFO_Statistics
	TOKEN_INFO_RestrictedSids
	TOKEN_INFO_SessionId
	TOKEN_INFO_GroupsAndPrivileges
	TOKEN_INFO_SessionReference
	TOKEN_INFO_SandBoxInert
	TOKEN_INFO_AuditPolicy
	TOKEN_INFO_Origin
	TOKEN_INFO_ElevationType
	TOKEN_INFO_Linked
	TOKEN_INFO_Elevation
	TOKEN_INFO_HasRestrictions
	TOKEN_INFO_AccessInformation
	TOKEN_INFO_VirtualizationAllowed
	TOKEN_INFO_VirtualizationEnabled
	TOKEN_INFO_IntegrityLevel
	TOKEN_INFO_UIAccess
	TOKEN_INFO_MandatoryPolicy
	TOKEN_INFO_LogonSid
	TOKEN_INFO_IsAppContainer
	TOKEN_INFO_Capabilities
	TOKEN_INFO_AppContainerSid
	TOKEN_INFO_AppContainerNumber
	TOKEN_INFO_UserClaimAttributes
	TOKEN_INFO_DeviceClaimAttributes
	TOKEN_INFO_RestrictedUserClaimAttributes
	TOKEN_INFO_RestrictedDeviceClaimAttributes
	TOKEN_INFO_DeviceGroups
	TOKEN_INFO_RestrictedDeviceGroups
	TOKEN_INFO_SecurityAttributes
	TOKEN_INFO_IsRestricted
	TOKEN_INFO_ProcessTrustLevel
	TOKEN_INFO_PrivateNameSpace
	TOKEN_INFO_SingletonAttributes
	TOKEN_INFO_BnoIsolation
	TOKEN_INFO_ChildProcessFlags
	TOKEN_INFO_IsLessPrivilegedAppContainer
	TOKEN_INFO_IsSandboxed
	TOKEN_INFO_OriginatingProcessTrustLevel
)

// [TOKEN_MANDATORY_POLICY] policy.
//
// [TOKEN_MANDATORY_POLICY]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-token_mandatory_policy
type TOKEN_POLICY uint32

const (
	TOKEN_POLICY_OFF             TOKEN_POLICY = 0x0
	TOKEN_POLICY_NO_WRITE_UP     TOKEN_POLICY = 0x1
	TOKEN_POLICY_NEW_PROCESS_MIN TOKEN_POLICY = 0x2
	TOKEN_POLICY_VALID_MASK      TOKEN_POLICY = 0x3
)
