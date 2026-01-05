//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/wstr"
)

// [BIND_OPTS3] struct.
//
// ⚠️ You must call [BIND_OPTS3.SetCbStruct] to initialize the struct.
//
// Example:
//
//	var bo3 win.BIND_OPTS3
//	bo3.SetCbStruct()
//
// [BIND_OPTS3]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-bind_opts3-r1
type BIND_OPTS3 struct {
	cbStruct            uint32
	GrfFlags            co.BIND
	GrfMode             co.STGM
	DwTickCountDeadline uint32
	DwTrackFlags        co.SLR
	DwClassContext      co.CLSCTX
	Locale              LCID
	PServerInfo         *COSERVERINFO
	Hwnd                HWND
}

// Sets the cbStruct field to the size of the struct, correctly initializing it.
func (bo *BIND_OPTS3) SetCbStruct() {
	bo.cbStruct = uint32(unsafe.Sizeof(*bo))
}

// [COAUTHIDENTITY] struct.
//
// [COAUTHIDENTITY]: https://learn.microsoft.com/en-us/windows/win32/api/wtypesbase/ns-wtypesbase-coauthidentity
type COAUTHIDENTITY struct {
	user           *uint16
	userLength     uint32
	domain         *uint16
	domainLength   uint32
	password       *uint16
	passwordLength uint32
	Flags          co.SEC_WINNT_AUTH_IDENTITY
}

func (coai *COAUTHIDENTITY) User() string {
	return wstr.DecodeSlice(unsafe.Slice(coai.user, coai.userLength))
}
func (coai *COAUTHIDENTITY) SetUser(val string) {
	buf := wstr.EncodeToSlice(val)
	coai.user = unsafe.SliceData(buf)
	coai.userLength = uint32(len(buf) - 1) // without terminating null
}

func (coai *COAUTHIDENTITY) Domain() string {
	return wstr.DecodeSlice(unsafe.Slice(coai.domain, coai.domainLength))
}
func (coai *COAUTHIDENTITY) SetDomain(val string) {
	buf := wstr.EncodeToSlice(val)
	coai.domain = unsafe.SliceData(buf)
	coai.domainLength = uint32(len(buf) - 1) // without terminating null
}

func (coai *COAUTHIDENTITY) Password() string {
	return wstr.DecodeSlice(unsafe.Slice(coai.password, coai.passwordLength))
}
func (coai *COAUTHIDENTITY) SetPassword(val string) {
	buf := wstr.EncodeToSlice(val)
	coai.password = unsafe.SliceData(buf)
	coai.passwordLength = uint32(len(buf) - 1) // without terminating null
}

// [COAUTHINFO] struct.
//
// [COAUTHINFO]: https://learn.microsoft.com/en-us/windows/win32/api/wtypesbase/ns-wtypesbase-coauthinfo
type COAUTHINFO struct {
	DwAuthnSvc           co.RPC_C_AUTHN
	DwAuthzSvc           co.RPC_C_AUTHZ
	PwszServerPrincName  *uint16
	DwAuthnLevel         co.RPC_C_AUTHN_LEVEL
	DwImpersonationLevel co.RPC_C_IMP_LEVEL
	PAuthIdentityData    *COAUTHIDENTITY
	DwCapabilities       co.EOAC_QOS
}

// [COSERVERINFO] struct.
//
// [COSERVERINFO]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-coserverinfo
type COSERVERINFO struct {
	dwReserved1 uint32
	PwszName    *uint16
	PAuthInfo   *COAUTHINFO
	dwReserved2 uint32
}

// [DVTARGETDEVICE] struct.
//
// ⚠️ You must call [DVTARGETDEVICE.SetTdSize] to initialize the struct.
//
// Example:
//
//	var dvt win.DVTARGETDEVICE
//	dvt.SetTdSize()
//
// [DVTARGETDEVICE]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-dvtargetdevice
type DVTARGETDEVICE struct {
	tdSize             uint32
	tdDriverNameOffset uint16
	tdDeviceNameOffset uint16
	tdPortNameOffset   uint16
	tdExtDevmodeOffset uint16
	tdData             [1]byte
}

// Sets the tdSize field to the size of the struct, correctly initializing it.
func (dvt *DVTARGETDEVICE) SetTdSize() {
	dvt.tdSize = uint32(unsafe.Sizeof(*dvt))
}

func (dvt *DVTARGETDEVICE) DriverName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdDriverNameOffset)
	return wstr.DecodePtr((*uint16)(ptr))
}

func (dvt *DVTARGETDEVICE) DeviceName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdDeviceNameOffset)
	return wstr.DecodePtr((*uint16)(ptr))
}

func (dvt *DVTARGETDEVICE) PortName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdPortNameOffset)
	return wstr.DecodePtr((*uint16)(ptr))
}

// [FORMATETC] struct.
//
// [FORMATETC]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-formatetc
type FORMATETC struct {
	CfFormat co.CF
	Ptd      *DVTARGETDEVICE
	Aspect   co.DVASPECT
	Lindex   int32
	Tymed    co.TYMED
}

// [STATSTG] struct.
//
// [STATSTG]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-statstg
type STATSTG struct {
	PwcsName          *uint16
	Type              co.STGTY
	CbSize            uint64
	MTime             FILETIME
	CTime             FILETIME
	ATime             FILETIME
	GrfMode           uint32
	GrfLocksSupported co.LOCKTYPE
	ClsId             GUID
	GrfStateBits      uint32
	reserved          uint32
}

// [STGMEDIUM] struct.
//
// If you received this struct from a COM call, you'll have to free the memory
// with [ReleaseStgMedium].
//
// [STGMEDIUM]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-ustgmedium-r1
type STGMEDIUM struct {
	tymed          co.TYMED
	data           uintptr // union
	PUnkForRelease IUnknown
}

// Returns the tymed field.
func (stg *STGMEDIUM) Tymed() co.TYMED {
	return stg.tymed
}

// Attemps to return the [HBITMAP] if tymed == co.TYMED_GDI.
func (stg *STGMEDIUM) HBitmap() (HBITMAP, bool) {
	if stg.tymed == co.TYMED_GDI {
		return HBITMAP(stg.data), true
	}
	return HBITMAP(0), false
}

// Attemps to return the [HGLOBAL] if tymed == co.TYMED_HGLOBAL.
func (stg *STGMEDIUM) HGlobal() (HGLOBAL, bool) {
	if stg.tymed == co.TYMED_HGLOBAL {
		return HGLOBAL(stg.data), true
	}
	return HGLOBAL(0), false
}

// Attemps to return the string if tymed == co.TYMED_FILE.
func (stg *STGMEDIUM) FileName() (string, bool) {
	if stg.tymed == co.TYMED_FILE {
		return wstr.DecodePtr((*uint16)(unsafe.Pointer(stg.data))), true
	}
	return "", false
}

// Attemps to return the [IStream] if tymed == co.TYMED_ISTREAM.
func (stg *STGMEDIUM) IStream(releaser *OleReleaser) (*IStream, bool) {
	if stg.tymed == co.TYMED_ISTREAM {
		ppvt := (**_IUnknownVt)(unsafe.Pointer(stg.data))
		pCurrent := &IStream{ISequentialStream{IUnknown{ppvt}}}

		var pCloned *IStream
		pCurrent.AddRef(releaser, &pCloned) // clone, because we'll release it independently
		return pCloned, true
	}
	return nil, false
}
