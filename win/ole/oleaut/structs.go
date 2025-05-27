//go:build windows

package oleaut

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

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

// [MEMBERID] identifiers a member in a type description.
//
// [MEMBERID]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/automat/memberid
type MEMBERID int32

// Predefined [MEMBERID].
//
// [MEMBERID]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/automat/memberid
const MEMBERID_NIL = MEMBERID(co.DISPID_UNKNOWN)

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
// When [IDispatch.Invoke] remote call fails, the raw EXCEPINFO is converted
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
// ⚠️ You must call [_EXCEPINFO.Free] to release the pointers.
//
// # Example
//
//	var e oleaut._EXCEPINFO
//	defer e.Free()
//
// [EXCEPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-excepinfo
type _EXCEPINFO struct {
	WCode             uint16
	wReserved         uint16
	BstrSource        BSTR
	BstrDescription   BSTR
	BstrHelpFile      BSTR
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

	// The BSTRs are freed inside IDispatch.Invoke().
	if e.BstrSource != 0 {
		nfo.Source = e.BstrSource.String()
	}
	if e.BstrDescription != 0 {
		nfo.Description = e.BstrDescription.String()
	}
	if e.BstrHelpFile != 0 {
		nfo.HelpFile = e.BstrHelpFile.String()
	}
	return &nfo
}

// Releases all the pointers.
func (e *_EXCEPINFO) Free() {
	BSTR(e.BstrSource).SysFreeString()
	BSTR(e.BstrDescription).SysFreeString()
	BSTR(e.BstrHelpFile).SysFreeString()
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

// [FUNCDESC] struct. Implements [ole.ComResource].
//
// [FUNCDESC]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-funcdesc
type FuncDescData struct {
	*FUNCDESC
	owner *ITypeInfo
}

// [ITypeInfo.ReleaseFuncDesc] method. Implements [ole.ComResource].
//
// You usually don't need to call this method directly, since every function
// which returns a [COM] object will require an [ole.Releaser] to manage the
// object's lifetime.
//
// [ITypeInfo.ReleaseFuncDesc]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypeinfo-releasefuncdesc
func (me *FuncDescData) Release() {
	me.owner._ReleaseFuncDesc(me.FUNCDESC)
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
