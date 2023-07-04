//go:build windows

package comco

// [CoCreateInstance] dwClsContext.
//
// [CoCreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
type CLSCTX uint32

const (
	CLSCTX_INPROC_SERVER          CLSCTX = 0x1
	CLSCTX_INPROC_HANDLER         CLSCTX = 0x2
	CLSCTX_LOCAL_SERVER           CLSCTX = 0x4
	CLSCTX_INPROC_SERVER16        CLSCTX = 0x8
	CLSCTX_REMOTE_SERVER          CLSCTX = 0x10
	CLSCTX_INPROC_HANDLER16       CLSCTX = 0x20
	CLSCTX_NO_CODE_DOWNLOAD       CLSCTX = 0x400
	CLSCTX_NO_CUSTOM_MARSHAL      CLSCTX = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD   CLSCTX = 0x2000
	CLSCTX_NO_FAILURE_LOG         CLSCTX = 0x4000
	CLSCTX_DISABLE_AAA            CLSCTX = 0x8000
	CLSCTX_ENABLE_AAA             CLSCTX = 0x1_0000
	CLSCTX_FROM_DEFAULT_CONTEXT   CLSCTX = 0x2_0000
	CLSCTX_ACTIVATE_X86_SERVER    CLSCTX = 0x4_0000
	CLSCTX_ACTIVATE_32_BIT_SERVER CLSCTX = CLSCTX_ACTIVATE_X86_SERVER
	CLSCTX_ACTIVATE_64_BIT_SERVER CLSCTX = 0x8_0000
	CLSCTX_ENABLE_CLOAKING        CLSCTX = 0x10_0000
	CLSCTX_APPCONTAINER           CLSCTX = 0x40_0000
	CLSCTX_ACTIVATE_AAA_AS_IU     CLSCTX = 0x80_0000
	CLSCTX_ACTIVATE_ARM32_SERVER  CLSCTX = 0x200_0000
	CLSCTX_PS_DLL                 CLSCTX = 0x8000_0000
	CLSCTX_ALL                    CLSCTX = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
	CLSCTX_SERVER                 CLSCTX = CLSCTX_INPROC_SERVER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
)

// [CoInitializeEx] dwCoInit.
//
// [CoInitializeEx]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
type COINIT uint32

const (
	COINIT_APARTMENTTHREADED COINIT = 0x2
	COINIT_MULTITHREADED     COINIT = 0x0
	COINIT_DISABLE_OLE1DDE   COINIT = 0x4
	COINIT_SPEED_OVER_MEMORY COINIT = 0x8
)

// [LOCKTYPE] enumeration.
//
// [LOCKTYPE]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ne-objidl-locktype
type LOCKTYPE uint32

const (
	LOCKTYPE_WRITE     LOCKTYPE = 1
	LOCKTYPE_EXCLUSIVE LOCKTYPE = 2
	LOCKTYPE_ONLYONCE  LOCKTYPE = 4
)

// [PICTUREATTRIBUTES] enumeration.
//
// [PICTUREATTRIBUTES]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/ne-ocidl-pictureattributes
type PICATTR uint32

const (
	PICATTR_SCALABLE    PICATTR = 0x01
	PICATTR_TRANSPARENT PICATTR = 0x02
)

// [PICTYPE] constants.
//
// [PICTYPE]: https://learn.microsoft.com/en-us/windows/win32/com/pictype-constants
type PICTYPE int16

const (
	PICTYPE_UNINITIALIZED PICTYPE = -1
	PICTYPE_NONE          PICTYPE = 0
	PICTYPE_BITMAP        PICTYPE = 1
	PICTYPE_METAFILE      PICTYPE = 2
	PICTYPE_ICON          PICTYPE = 3
	PICTYPE_ENHMETAFILE   PICTYPE = 4
)

// [STGC] enumeration.
//
// [STGC]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ne-wtypes-stgc
type STGC uint32

const (
	STGC_DEFAULT                            STGC = 0
	STGC_OVERWRITE                          STGC = 1
	STGC_ONLYIFCURRENT                      STGC = 2
	STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE STGC = 4
	STGC_CONSOLIDATE                        STGC = 8
)

// [STREAM_SEEK] enumeration.
//
// [STREAM_SEEK]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ne-objidl-stream_seek
type STREAM_SEEK uint32

const (
	STREAM_SEEK_SET STREAM_SEEK = 0
	STREAM_SEEK_CUR STREAM_SEEK = 1
	STREAM_SEEK_END STREAM_SEEK = 2
)
