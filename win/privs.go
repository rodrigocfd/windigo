//go:build windows

package win

// Private constants from advapi.
const (
	_MANAGED_APPS_USERAPPLICATIONS  = 0x1
	_MANAGED_APPS_FROMCATEGORY      = 0x2
	_MANAGED_APPS_INFOLEVEL_DEFAULT = 0x1_0000
)

// Private constants from comctl.
const (
	_L_MAX_URL_LENGTH = 2048 + 32 + 4
	_MAX_LINKID_TEXT  = 48
)

// Private constants from gdi.
const (
	_CLR_INVALID  = 0xffff_ffff
	_GDI_ERR      = 0xffff_ffff
	_HGDI_ERROR   = 0xffff_ffff
	_LF_FACESIZE  = 32
	_REGION_ERROR = 0
	_SP_ERROR     = -1
)

// Private constants from kernel.
const (
	_GMEM_INVALID_HANDLE  = 0x8000
	_INVALID_HANDLE_VALUE = -1
	_LMEM_INVALID_HANDLE  = 0x8000
	_MAX_MODULE_NAME32    = 255
	_MAX_PATH             = 260
)

// Private constants from ole.
const (
	_HIMETRIC_PER_INCH = 2540
)

// Private constants form shell.
const (
	_UINT_MAX = 4294967295
)

// Private constants from user.
const (
	_CCHDEVICENAME      = 32
	_CCHILDREN_TITLEBAR = 5
	_CP_WINUNICODE      = 1200
	_TIMEOUT_ASYNC      = 0xffff_ffff
)
