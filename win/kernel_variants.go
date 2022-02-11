package win

import (
	"unsafe"
)

type (
	// Variant type for an optional string value.
	//
	// Example:
	//
	//  func Foo(maybeName win.StrOpt) {
	//      // ...
	//  }
	//
	//  Foo(win.StrOptVal("some name"))
	//  Foo(win.StrOptNone{})
	StrOpt     interface{ implStrOpt() }
	StrOptVal  string   // StrOpt variant: string value.
	StrOptNone struct{} // StrOpt variant: no value.
)

func (StrOptVal) implStrOpt()  {}
func (StrOptNone) implStrOpt() {}

func variantStrOpt(v StrOpt) unsafe.Pointer {
	switch v := v.(type) {
	case StrOptVal:
		return unsafe.Pointer(Str.ToNativePtr(string(v)))
	case StrOptNone:
		return nil
	default:
		panic("StrOpt cannot be nil.")
	}
}
