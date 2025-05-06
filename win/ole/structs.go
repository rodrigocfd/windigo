//go:build windows

package ole

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [MEMBERID] identifiers a member in a type description.
//
// [MEMBERID]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/automat/memberid
type MEMBERID int32

// [ARRAYDESC] struct.
//
// [ARRAYDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-arraydesc
type ARRAYDESC struct {
	TdescElem TYPEDESC
	CDims     uint16
	rgbounds  [1]SAFEARRAYBOUND
}

func (ad *ARRAYDESC) Rgbounds(i int) *SAFEARRAYBOUND {
	return &ad.rgbounds[i]
}

// [DISPPARAMS] sruct.
//
// [DISPPARAMS]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-dispparams
type DISPPARAMS struct {
	rgvarg            *VARIANT // in reverse order
	rgdispidNamedArgs *co.DISPID
	cArgs             uint32
	cNamedArgs        uint32
}

func (dp *DISPPARAMS) SetArgs(v []VARIANT) {
	dp.cArgs = uint32(len(v))
	dp.rgvarg = &v[0]
}

func (dp *DISPPARAMS) SetNamedArgs(v ...co.DISPID) {
	dp.cNamedArgs = uint32(len(v))
	dp.rgdispidNamedArgs = &v[0]
}

// [DVTARGETDEVICE] struct.
//
// ⚠️ You must call SetTdSize() to initialize the struct.
//
// # Example
//
//	var dvt ole.DVTARGETDEVICE
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
	return wstr.Utf16PtrToStr((*uint16)(ptr))
}

func (dvt *DVTARGETDEVICE) DeviceName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdDeviceNameOffset)
	return wstr.Utf16PtrToStr((*uint16)(ptr))
}

func (dvt *DVTARGETDEVICE) PortName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdPortNameOffset)
	return wstr.Utf16PtrToStr((*uint16)(ptr))
}

// [ELEMDESC] struct.
//
// [ELEMDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-elemdesc-r1
type ELEMDESC struct {
	TDesc TYPEDESC
	union [16]byte
}

func (ed *ELEMDESC) IdlDesc() *IDLDESC {
	return (*IDLDESC)(unsafe.Pointer(&ed.union[0]))
}

func (ed *ELEMDESC) ParmDesc() *PARAMDESC {
	return (*PARAMDESC)(unsafe.Pointer(&ed.union[0]))
}

// [EXCEPINFO] struct syntactic sugar.
//
// When IDispatch.Invoke() remote call fails, the raw [EXCEPINFO] is converted
// into this one, then returned as the error.
//
// Implements error interface.
//
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
type EXCEPINFO struct {
	Code        int
	Source      string
	Description string
	HelpFile    string
}

// Implements error interface.
func (e *EXCEPINFO) Error() string {
	return e.Source + ": " + e.Description
}

// [EXCEPINFO] struct.
//
// ⚠️ You must call Free() to release the pointers.
//
// # Example
//
//	var e ole._EXCEPINFO
//	defer e.Free()
//
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
type _EXCEPINFO struct {
	WCode             uint16
	wReserved         uint16
	BstrSource        uintptr
	BstrDescription   uintptr
	BstrHelpFile      uintptr
	DwHelpContext     uint16
	pvReserved        uintptr
	PfnDeferredFillIn uintptr
	Scode             int32
}

// Converts the information into the syntactic sugar struct.
func (e *_EXCEPINFO) Serialize() *EXCEPINFO {
	var nfo EXCEPINFO

	nfo.Code = int(e.WCode)
	if nfo.Code == 0 {
		nfo.Code = int(e.Scode)
	}

	if e.BstrSource != 0 {
		bstr := BSTR(e.BstrSource) // don't free
		nfo.Source = bstr.String()
	}
	if e.BstrDescription != 0 {
		bstr := BSTR(e.BstrDescription) // don't free
		nfo.Description = bstr.String()
	}
	if e.BstrHelpFile != 0 {
		bstr := BSTR(e.BstrHelpFile) // don't free
		nfo.HelpFile = bstr.String()
	}

	return &nfo
}

// Releases all the pointers.
func (e *_EXCEPINFO) Free() {
	BSTR(e.BstrSource).SysFreeString()
	BSTR(e.BstrDescription).SysFreeString()
	BSTR(e.BstrHelpFile).SysFreeString()
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

// [FUNCDESC] struct.
//
// [FUNCDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-funcdesc
type FUNCDESC struct {
	Memid             MEMBERID
	Lprgscode         *int32
	LprgelemdescParam *ELEMDESC
	Funckind          co.FUNCKIND
	Invkind           co.INVOKEKIND
	Callconv          co.CALLCONV
	CParams           int16
	CParamsOpt        int16
	OVft              int16
	CScodes           int16
	ElemdescFunc      ELEMDESC
	WFuncFlags        co.FUNCFLAG
}

// [IDLDESC] struct.
//
// [IDLDESC]: https://learn.microsoft.com/en-us/previous-versions/windows/embedded/aa515591(v=msdn.10)
type IDLDESC struct {
	dwReserved uintptr
	WIDLFlags  co.IDLFLAG
}

// [PARAMDESC] struct.
//
// [PARAMDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-paramdesc
type PARAMDESC struct {
	Pparamdescex *PARAMDESCEX
	WParamFlags  co.PARAMFLAG
}

// [PARAMDESCEX] struct.
//
// [PARAMDESCEX]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-paramdescex
type PARAMDESCEX struct {
	CBytes          uint32
	VarDefaultValue VARIANT
}

// [SAFEARRAYBOUND] struct.
//
// [SAFEARRAYBOUND]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-safearraybound
type SAFEARRAYBOUND struct {
	CElements uint32
	LLbound   int32
}

// [TYPEDESC] struct.
//
// [TYPEDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-typedesc
type TYPEDESC struct {
	union uintptr
	Vt    co.VT
}

func (td *TYPEDESC) TypeDesc() *TYPEDESC {
	return (*TYPEDESC)(unsafe.Pointer(td.union))
}

func (td *TYPEDESC) ArrayDesc() *ARRAYDESC {
	return (*ARRAYDESC)(unsafe.Pointer(td.union))
}

func (td *TYPEDESC) HRefType() uint32 {
	return uint32(td.union)
}

// [STGMEDIUM] struct.
//
// If you received this struct from a COM call, you'll have to free the memory
// with [ReleaseStgMedium].
//
// [STGMEDIUM]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-ustgmedium-r1
// [ReleaseStgMedium]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-releasestgmedium
type STGMEDIUM struct {
	tymed          co.TYMED
	data           uintptr // union
	PUnkForRelease IUnknown
}

func (stg *STGMEDIUM) Tymed() co.TYMED {
	return stg.tymed
}

func (stg *STGMEDIUM) HBitmap() (win.HBITMAP, bool) {
	if stg.tymed == co.TYMED_GDI {
		return win.HBITMAP(stg.data), true
	}
	return win.HBITMAP(0), false
}

func (stg *STGMEDIUM) HGlobal() (win.HGLOBAL, bool) {
	if stg.tymed == co.TYMED_HGLOBAL {
		return win.HGLOBAL(stg.data), true
	}
	return win.HGLOBAL(0), false
}

func (stg *STGMEDIUM) FileName() (string, bool) {
	if stg.tymed == co.TYMED_FILE {
		return wstr.Utf16PtrToStr((*uint16)(unsafe.Pointer(stg.data))), true
	}
	return "", false
}

func (stg *STGMEDIUM) IStream() (*IStream, bool) {
	if stg.tymed == co.TYMED_ISTREAM {
		pObj := vt.NewObj[IStream]((**vt.IUnknown)(unsafe.Pointer(stg.data)))
		return pObj, true // not added to a Releaser, will be freed by ReleaseStgMedium()
	}
	return nil, false
}
