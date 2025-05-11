//go:build windows

package wutil

// Internal constants for ole/shell.
const (
	INFOTIPSIZE = 1024
)

// Internal constants for comctl.
const (
	L_MAX_URL_LENGTH = 2048 + 32 + 4
	MAX_LINKID_TEXT  = 48
)

// Internal constants for gdi.
const (
	CCHDEVICENAME     = 32
	CCHFORMNAME       = 32
	CLR_INVALID       = 0xffff_ffff
	DM_SPECVERSION    = 0x0401
	GDI_ERR           = 0xffff_ffff
	HGDI_ERROR        = 0xffff_ffff
	HIMETRIC_PER_INCH = 2540
	LF_FACESIZE       = 32
	REGION_ERROR      = 0
)

// Internal constants for kernel.
const (
	GMEM_INVALID_HANDLE  = 0x8000
	INFINITE             = 0xffff_ffff
	INVALID_HANDLE_VALUE = -1
	MAX_MODULE_NAME32    = 255
	MAX_PATH             = 260
	TIME_ZONE_INVALID    = 0xffff_ffff
)

// Internal constants for user.
const (
	CCHILDREN_TITLEBAR = 5
	WC_DIALOG          = uint16(0x8002)
)

// Internal constants for ole/shell.
const (
	PID_FIRST_USABLE = 0x2
)
