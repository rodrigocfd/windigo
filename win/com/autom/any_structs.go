//go:build windows

package autom

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automco"
)

// [ARRAYDESC] struct.
//
// [ARRAYDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-arraydesc
type ARRAYDESC struct {
	TdescElem TYPEDESC
	CDims     uint16
	rgbounds  [1]SAFEARRAYBOUND
}

func (ad *ARRAYDESC) Rgbounds(i int) *SAFEARRAYBOUND { return &ad.rgbounds[i] }

// [DISPPARAMS] sruct.
//
// [DISPPARAMS]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-dispparams
type DISPPARAMS struct {
	rgvarg            *VARIANT
	rgdispidNamedArgs *automco.DISPID
	cArgs             uint32
	cNamedArgs        uint32
}

func (dp *DISPPARAMS) SetArgs(v ...VARIANT) {
	dp.cArgs = uint32(len(v))
	dp.rgvarg = &v[0]
}

func (dp *DISPPARAMS) SetNamedArgs(v ...automco.DISPID) {
	dp.cNamedArgs = uint32(len(v))
	dp.rgdispidNamedArgs = &v[0]
}

// [ELEMDESC] struct.
//
// [ELEMDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-elemdesc-r1
type ELEMDESC struct {
	TDesc TYPEDESC
	union [16]byte
}

func (ed *ELEMDESC) IdlDesc() *IDLDESC    { return (*IDLDESC)(unsafe.Pointer(&ed.union[0])) }
func (ed *ELEMDESC) ParmDesc() *PARAMDESC { return (*PARAMDESC)(unsafe.Pointer(&ed.union[0])) }

// [EXCEPINFO] struct.
//
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
type EXCEPINFO struct {
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

func (e *EXCEPINFO) ReleaseStrings() (string, string, string) {
	s0, s1, s2 := "", "", ""
	if e.BstrSource != 0 {
		bstr := BSTR(e.BstrSource)
		defer bstr.SysFreeString()
		s0 = bstr.String()
	}
	if e.BstrDescription != 0 {
		bstr := BSTR(e.BstrDescription)
		defer bstr.SysFreeString()
		s1 = bstr.String()
	}
	if e.BstrHelpFile != 0 {
		bstr := BSTR(e.BstrHelpFile)
		defer bstr.SysFreeString()
		s2 = bstr.String()
	}
	return s0, s1, s2
}

// [FUNCDESC] struct.
//
// [FUNCDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-funcdesc
type FUNCDESC struct {
	Memid             MEMBERID
	Lprgscode         *int32
	LprgelemdescParam *ELEMDESC
	Funckind          automco.FUNCKIND
	Invkind           automco.INVOKEKIND
	Callconv          automco.CALLCONV
	CParams           int16
	CParamsOpt        int16
	OVft              int16
	CScodes           int16
	ElemdescFunc      ELEMDESC
	WFuncFlags        automco.FUNCFLAG
}

// [IDLDESC] struct.
//
// [IDLDESC]: https://learn.microsoft.com/en-us/previous-versions/windows/embedded/aa515591(v=msdn.10)
type IDLDESC struct {
	dwReserved uintptr
	WIDLFlags  automco.IDLFLAG
}

// [PARAMDESC] struct.
//
// [PARAMDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-paramdesc
type PARAMDESC struct {
	Pparamdescex *PARAMDESCEX
	WParamFlags  automco.PARAMFLAG
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

// [TYPEATTR] struct.
//
// [TYPEATTR]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-typeattr
type TYPEATTR struct {
	Guid             win.GUID
	Lcid             win.LCID
	dwReserved       uint32
	MemidConstructor MEMBERID
	MemidDestructor  MEMBERID
	lpstrSchema      *uint16
	CbSizeInstance   uint32
	Typekind         automco.TYPEKIND
	CFuncs           uint16
	CVars            uint16
	CImplTypes       uint16
	CbSizeVft        uint16
	CbAlignment      uint16
	WTypeFlags       automco.TYPEFLAG
	WMajorVerNum     uint16
	WMinorVerNum     uint16
	TdescAlias       TYPEDESC
	IdldescType      IDLDESC
}

// [TYPEDESC] struct.
//
// [TYPEDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-typedesc
type TYPEDESC struct {
	union uintptr
	Vt    automco.VT
}

func (td *TYPEDESC) TypeDesc() *TYPEDESC   { return (*TYPEDESC)(unsafe.Pointer(td.union)) }
func (td *TYPEDESC) ArrayDesc() *ARRAYDESC { return (*ARRAYDESC)(unsafe.Pointer(td.union)) }
func (td *TYPEDESC) HRefType() uint32      { return uint32(td.union) }

// [VARDESC] struct.
//
// [VARDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-vardesc
type VARDESC struct {
	Memid       MEMBERID
	LpstrSchema *uint16
	LpvarValue  *VARIANT // union ULONG | *VARIANT
	ElemdescVar ELEMDESC
	WVarFlags   automco.VARFLAG
	Varkind     automco.VARKIND
}
