//go:build windows

package win

import (
	"unsafe"
)

// Variant type for an optional string value.
//
// Example:
//
//	realStr := win.StrOptSome("foo")
//
//	if s, ok := realStr.Str(); ok {
//		println(s)
//	}
//
//	fakeStr := win.StrOptNone()
type StrOpt struct {
	isSome bool
	str    string
}

// Creates a new StrOpt with an empty value.
func StrOptNone() StrOpt {
	return StrOpt{}
}

// Creates a new StrOpt with a string value.
func StrOptSome(str string) StrOpt {
	return StrOpt{
		isSome: true,
		str:    str,
	}
}

func (me *StrOpt) IsSome() bool        { return me.isSome }
func (me *StrOpt) IsNone() bool        { return !me.isSome }
func (me *StrOpt) Str() (string, bool) { return me.str, me.isSome }

// Returns the *uint16 of the string converted to a native pointer, or nil.
func (me *StrOpt) Raw() unsafe.Pointer {
	if me.isSome {
		buf := Str.ToNativePtr(me.str)
		return unsafe.Pointer(buf)
	} else {
		return nil
	}
}
