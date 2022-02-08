package win

import (
	"unsafe"
)

type (
	// Variant type which accepts string or nil.
	//
	// Example:
	//
	//  func Foo(s win.StrOrNil) {}
	//
	//  Foo(win.StrVal("some text"))
	//  Foo(nil)
	StrOrNil interface{ isStrOrNil() }
	StrVal   string // StrOrNil variant: string.
)

func (StrVal) isStrOrNil() {}

func variantStrOrNil(v StrOrNil) unsafe.Pointer {
	if v != nil {
		s := v.(StrVal)
		return unsafe.Pointer(Str.ToNativePtr(string(s)))
	}
	return nil
}
