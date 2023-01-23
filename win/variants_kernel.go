//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// Variant type for a resource type.
//
// Example:
//
//	rsrcTy := win.RsrcTypeRt(co.RT_ACCELERATOR)
type RsrcType struct {
	curType uint8
	rt      co.RT  // 1
	str     string // 2
}

// Creates a new RsrcType variant with a co.RT value.
func RsrcTypeRt(rt co.RT) RsrcType {
	return RsrcType{
		curType: 1,
		rt:      rt,
	}
}

// Creates a new RsrcType variant with a string value.
func RsrcTypeStr(str string) RsrcType {
	return RsrcType{
		curType: 2,
		str:     str,
	}
}

func (me *RsrcType) Rt() (co.RT, bool)   { return me.rt, me.curType == 1 }
func (me *RsrcType) Str() (string, bool) { return me.str, me.curType == 2 }

// Converts the internal value to uintptr; pointer must be kept alive.
func (me *RsrcType) raw() (val uintptr, ptr *uint16) {
	switch me.curType {
	case 1:
		return uintptr(me.rt), nil
	case 2:
		buf := Str.ToNativePtr(me.str)
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("Invalid RsrcType value.")
	}
}

//------------------------------------------------------------------------------

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
