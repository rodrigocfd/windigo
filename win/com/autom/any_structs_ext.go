//go:build windows

package autom

import (
	"strconv"
	"strings"

	"github.com/rodrigocfd/windigo/win/com/autom/automco"
)

// Exception information from an IDispatch.Invoke() call.
//
// Implements error interface.
type ExceptionInfo struct {
	Code        int32
	Source      string
	Description string
	HelpFile    string
}

// Implements error interface.
func (e *ExceptionInfo) Error() string {
	var buf strings.Builder
	buf.Grow(len(e.Source) + len(e.Description) + len(e.HelpFile) + 10) // arbitrary

	buf.WriteString(strconv.Itoa(int(e.Code)))
	if e.Source != "" {
		buf.WriteString("; ")
		buf.WriteString(e.Source)
	}
	if e.Description != "" {
		buf.WriteString("; ")
		buf.WriteString(e.Description)
	}
	if e.HelpFile != "" {
		buf.WriteString("; ")
		buf.WriteString(e.HelpFile)
	}
	return buf.String()
}

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
