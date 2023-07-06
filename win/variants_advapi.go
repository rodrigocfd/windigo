//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// Variant type for a Registry value.
//
// # Example
//
//	regVal := RegValSz("Some text")
//
//	if val, ok := regVal.Sz(); ok {
//		println(val)
//	}
type RegVal struct {
	ty  co.REG
	val []byte
}

// Returns the stored co.REG type.
func (me *RegVal) Type() co.REG {
	return me.ty
}

// Creates a new RegVal variant with a co.REG_NONE value.
func RegValNone() RegVal {
	return RegVal{
		ty:  co.REG_NONE,
		val: nil,
	}
}

// Returns true if current value is co.REG_NONE.
func (me *RegVal) IsNone() bool {
	return me.ty == co.REG_NONE
}

// Creates a new RegVal variant with a co.REG_BINARY value.
func RegValBinary(data []byte) RegVal {
	sliceData := make([]byte, len(data))
	copy(sliceData, data)

	return RegVal{
		ty:  co.REG_BINARY,
		val: sliceData,
	}
}

// If the current value is co.REG_BINARY, returns it and true; otherwise nil and
// false.
//
// # Example:
//
//	regVal := win.RegValBinary([]byte{0x10, 0x44})
//
//	if val, ok := regVal.Binary(); ok {
//		println(len(val))
//	}
func (me *RegVal) Binary() ([]byte, bool) {
	if me.ty == co.REG_BINARY {
		data := make([]byte, len(me.val))
		copy(data, me.val)
		return data, true
	} else {
		return nil, false
	}
}

// Creates a new RegVal variant with a co.REG_DWORD value.
func RegValDword(n uint32) RegVal {
	sliceUsr := unsafe.Slice((*byte)(unsafe.Pointer(&n)), unsafe.Sizeof(n))
	sliceData := make([]byte, unsafe.Sizeof(n))
	copy(sliceData, sliceUsr)

	return RegVal{
		ty:  co.REG_DWORD,
		val: sliceData,
	}
}

// If the current value is co.REG_DWORD, returns it and true; otherwise 0 and
// false.
//
// # Example:
//
//	regVal := RegValDword(0x8000_1001)
//
//	if val, ok := regVal.Dword(); ok {
//		println(val)
//	}
func (me *RegVal) Dword() (uint32, bool) {
	if me.ty == co.REG_DWORD {
		data := (*uint32)(unsafe.Pointer(&me.val[0]))
		return *data, true
	} else {
		return 0, false
	}
}

// Creates a new RegVal variant with a co.REG_EXPAND_SZ value.
//
// When the value is retrieved, the environment variables should be expanded
// with [ExpandEnvironmentStrings] function.
//
// [ExpandEnvironmentStrings]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-expandenvironmentstringsw
func RegValExpandSz(s string) RegVal {
	sliceUsr := Str.ToNativeSlice(s)
	sliceData := unsafe.Slice((*byte)(unsafe.Pointer(&sliceUsr[0])), len(sliceUsr)*2)

	return RegVal{
		ty:  co.REG_EXPAND_SZ,
		val: sliceData,
	}
}

// If the current value is co.REG_EXPAND_SZ, returns it and true; otherwise ""
// and false.
//
// Environment variables can be expanded with [ExpandEnvironmentStrings]
// function.
//
// # Example:
//
//	regVal := RegValExpandSz("Some text")
//
//	if val, ok := regVal.ExpandSz(); ok {
//		println(ExpandEnvironmentStrings(val))
//	}
//
// [ExpandEnvironmentStrings]: https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-expandenvironmentstringsw
func (me *RegVal) ExpandSz() (string, bool) {
	if me.ty == co.REG_EXPAND_SZ {
		data := unsafe.Slice((*uint16)(unsafe.Pointer(&me.val[0])), len(me.val)/2)
		return Str.FromNativeSlice(data), true
	} else {
		return "", false
	}
}

// Creates a new RegVal variant with a co.REG_QWORD value.
func RegValQword(n uint64) RegVal {
	sliceUsr := unsafe.Slice((*byte)(unsafe.Pointer(&n)), unsafe.Sizeof(n))
	sliceData := make([]byte, unsafe.Sizeof(n))
	copy(sliceData, sliceUsr)

	return RegVal{
		ty:  co.REG_QWORD,
		val: sliceData,
	}
}

// If the current value is co.REG_QWORD, returns it and true; otherwise 0 and
// false.
//
// # Example:
//
//	regVal := RegValQword(0x8000_3303_0000_1001)
//
//	if val, ok := regVal.Qword(); ok {
//		println(val)
//	}
func (me *RegVal) Qword() (uint64, bool) {
	if me.ty == co.REG_QWORD {
		data := (*uint64)(unsafe.Pointer(&me.val[0]))
		return *data, true
	} else {
		return 0, false
	}
}

// Creates a new RegVal variant with a co.REG_SZ value.
func RegValSz(s string) RegVal {
	sliceUsr := Str.ToNativeSlice(s)
	sliceData := unsafe.Slice((*byte)(unsafe.Pointer(&sliceUsr[0])), len(sliceUsr)*2)

	return RegVal{
		ty:  co.REG_SZ,
		val: sliceData,
	}
}

// If the current value is co.REG_SZ, returns it and true; otherwise "" and
// false.
//
// # Example:
//
//	regVal := RegValSz("Some text")
//
//	if val, ok := regVal.Sz(); ok {
//		println(val)
//	}
func (me *RegVal) Sz() (string, bool) {
	if me.ty == co.REG_SZ {
		data := unsafe.Slice((*uint16)(unsafe.Pointer(&me.val[0])), len(me.val)/2)
		return Str.FromNativeSlice(data), true
	} else {
		return "", false
	}
}
