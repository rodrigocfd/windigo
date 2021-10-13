package autom

import (
	"github.com/rodrigocfd/windigo/win/com/autom/automco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-elemdesc-r1
type ELEMDESC struct {
	TDesc     TYPEDESC
	Paramdesc PARAMDESC // union IDLDESC | PARAMDESC
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-funcdesc
type FUNCDESC struct {
	Memid             int32
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-typedesc
type TYPEDESC struct {
	Lptdesc *TYPEDESC // union *TYPEDESC | *ARRAYDESC | HREFTYPE
	Vt      VARIANT
}
