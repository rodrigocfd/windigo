//go:build windows

package autom

import (
	"github.com/rodrigocfd/windigo/win/com/autom/automco"
)

// ITypeInfo.ListFunctions() return type.
type FuncDescResume struct {
	MemberId     MEMBERID
	Name         string
	FuncKind     automco.FUNCKIND
	InvokeKind   automco.INVOKEKIND
	NumParams    int
	NumOptParams int
	Flags        automco.FUNCFLAG
}

// ITypeInfo.GetDocumentation() return type.
type TypeDoc struct {
	Name        string
	DocString   string
	HelpContext uint32
	HelpFile    string
}
