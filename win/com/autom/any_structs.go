package autom

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-arraydesc
type ARRAYDESC struct {
	TdescElem TYPEDESC
	CDims     uint16
	Rgbounds  [1]SAFEARRAYBOUND
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-dispparams
type DISPPARAMS struct {
	rgvarg            *VARIANT
	rgdispidNamedArgs *MEMBERID
	cArgs             uint32
	cNamedArgs        uint32
}

func (dp *DISPPARAMS) SetArgs(v []VARIANT) {
	dp.cArgs = uint32(len(v))
	dp.rgvarg = &v[0]
}

func (dp *DISPPARAMS) SetNamedArgs(v []MEMBERID) {
	dp.cNamedArgs = uint32(len(v))
	dp.rgdispidNamedArgs = &v[0]
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-elemdesc-r1
type ELEMDESC struct {
	TDesc TYPEDESC
	union [16]byte
}

func (ed *ELEMDESC) IdlDesc() *IDLDESC    { return (*IDLDESC)(unsafe.Pointer(&ed.union[0])) }
func (ed *ELEMDESC) ParmDesc() *PARAMDESC { return (*PARAMDESC)(unsafe.Pointer(&ed.union[0])) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
type EXCEPINFO struct {
	WCode             uint16
	wReserved         uint16
	BstrSource        *uint16
	BstrDescription   *uint16
	BstrHelpFile      *uint16
	DwHelpContext     uint16
	pvReserved        uintptr
	PfnDeferredFillIn uintptr
	Scode             int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-funcdesc
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

// Composes various structs. Apparently undocumented.
type IDLDESC struct {
	dwReserved uintptr
	WIDLFlags  automco.IDLFLAG
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-paramdesc
type PARAMDESC struct {
	Pparamdescex *PARAMDESCEX
	WParamFlags  automco.PARAMFLAG
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-paramdescex
type PARAMDESCEX struct {
	CBytes          uint32
	VarDefaultValue VARIANT
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-safearraybound
type SAFEARRAYBOUND struct {
	CElements uint32
	LLbound   int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-typeattr
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-typedesc
type TYPEDESC struct {
	union uintptr
	Vt    automco.VT
}

func (td *TYPEDESC) TypeDesc() *TYPEDESC   { return (*TYPEDESC)(unsafe.Pointer(td.union)) }
func (td *TYPEDESC) ArrayDesc() *ARRAYDESC { return (*ARRAYDESC)(unsafe.Pointer(td.union)) }
func (td *TYPEDESC) HRefType() uint32      { return uint32(td.union) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-vardesc
type VARDESC struct {
	Memid       MEMBERID
	LpstrSchema *uint16
	LpvarValue  *VARIANT // union ULONG | *VARIANT
	ElemdescVar ELEMDESC
	WVarFlags   automco.VARFLAG
	Varkind     automco.VARKIND
}
