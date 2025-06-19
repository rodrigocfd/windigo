//go:build windows

package win

import (
	"encoding/binary"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// Tagged union for a Registry value.
//
// # Example
//
//	regVal := win.RegValSz("Some text")
//
//	if val, ok := regVal.Sz(); ok {
//		println(val)
//	}
type RegVal struct {
	tag  co.REG
	data []byte
}

// Returns the type of the RegVal.
func (me *RegVal) Type() co.REG {
	return me.tag
}

// Creates a new [RegVal] with a [co.REG_NONE] value.
func RegValNone() RegVal {
	return RegVal{
		tag: co.REG_NONE,
	}
}

// If the value is [co.REG_NONE], returns true.
func (me *RegVal) IsNone() bool {
	return me.tag == co.REG_NONE
}

// Creates a new [RegVal] with a [co.REG_BINARY] value.
//
// Note that the data content is not copied, the slice pointer is simply stored.
func RegValBinary(data []byte) RegVal {
	return RegVal{
		tag:  co.REG_BINARY,
		data: data,
	}
}

// If the value is [co.REG_BINARY], returns it and true.
func (me *RegVal) Binary() ([]byte, bool) {
	if me.tag == co.REG_BINARY {
		return me.data, true
	}
	return nil, false
}

// Creates a new [RegVal] with a [co.REG_EXPAND_SZ] value.
//
// The environment variables can be expanded with [ExpandEnvironmentStrings].
func RegValExpandSz(s string) RegVal {
	str16, _ := syscall.UTF16FromString(s)
	data := unsafe.Slice((*byte)(unsafe.Pointer(&str16[0])), len(str16)*2)

	return RegVal{
		tag:  co.REG_EXPAND_SZ,
		data: data,
	}
}

// If the value is [co.REG_EXPAND_SZ], returns it and true.
//
// The environment variables can be expanded with [ExpandEnvironmentStrings].
func (me *RegVal) ExpandSz() (string, bool) {
	if me.tag == co.REG_EXPAND_SZ {
		str16 := unsafe.Slice((*uint16)(unsafe.Pointer(&me.data[0])), len(me.data)/2)
		return wstr.WinSliceToGo(str16), true
	}
	return "", false
}

// Creates a new [RegVal] with a [co.REG_DWORD] value.
//
// Same as [co.REG_DWORD_LITTLE_ENDIAN].
func RegValDword(n uint32) RegVal {
	var data [4]byte
	binary.LittleEndian.PutUint32(data[:], n)

	return RegVal{
		tag:  co.REG_DWORD,
		data: data[:],
	}
}

// If the value is [co.REG_DWORD] or [co.REG_DWORD_LITTLE_ENDIAN], returns it
// and true.
func (me *RegVal) Dword() (uint32, bool) {
	if me.tag == co.REG_DWORD {
		return binary.LittleEndian.Uint32(me.data), true
	}
	return 0, false
}

// Creates a new [RegVal] with a [co.REG_DWORD_BIG_ENDIAN] value.
func RegValDwordBigEndian(n uint32) RegVal {
	var data [4]byte
	binary.BigEndian.PutUint32(data[:], n)

	return RegVal{
		tag:  co.REG_DWORD_BIG_ENDIAN,
		data: data[:],
	}
}

// If the value is [co.REG_DWORD_BIG_ENDIAN], returns it and true.
func (me *RegVal) DwordBigEndian() (uint32, bool) {
	if me.tag == co.REG_DWORD_BIG_ENDIAN {
		return binary.BigEndian.Uint32(me.data), true
	}
	return 0, false
}

// Creates a new [RegVal] with a [co.REG_MULTI_SZ] value.
func RegValMultiSz(strs ...string) RegVal {
	buf := wstr.GoArrToWinSlice(strs...)
	data := unsafe.Slice((*byte)(unsafe.Pointer(&buf[0])), len(buf)*2)

	return RegVal{
		tag:  co.REG_MULTI_SZ,
		data: data,
	}
}

// If the value is [co.REG_MULTI_SZ], returns it and true.
func (me *RegVal) MultiSz() ([]string, bool) {
	if me.tag == co.REG_MULTI_SZ {
		pStr16 := (*uint16)(unsafe.Pointer(&me.data[0]))
		return wstr.WinArrPtrToGo(pStr16), true
	}
	return nil, false
}

// Creates a new [RegVal] with a [co.REG_QWORD] value.
//
// Same as [co.REG_QWORD_LITTLE_ENDIAN].
func RegValQword(n uint64) RegVal {
	var data [4]byte
	binary.LittleEndian.PutUint64(data[:], n)

	return RegVal{
		tag:  co.REG_QWORD,
		data: data[:],
	}
}

// If the value is [co.REG_QWORD] or [co.REG_QWORD_LITTLE_ENDIAN], returns it
// and true.
func (me *RegVal) Qword() (uint64, bool) {
	if me.tag == co.REG_QWORD {
		return binary.LittleEndian.Uint64(me.data), true
	}
	return 0, false
}

// Creates a new [RegVal] with a [co.REG_SZ] value.
func RegValSz(s string) RegVal {
	str16, _ := syscall.UTF16FromString(s)
	data := unsafe.Slice((*byte)(unsafe.Pointer(&str16[0])), len(str16)*2)

	return RegVal{
		tag:  co.REG_SZ,
		data: data,
	}
}

// If the value is [co.REG_SZ], returns it and true.
//
// # Example:
//
//	regVal := win.RegValSz("Some text")
//
//	if val, ok := regVal.Sz(); ok {
//		println(val)
//	}
func (me *RegVal) Sz() (string, bool) {
	if me.tag == co.REG_SZ {
		str16 := unsafe.Slice((*uint16)(unsafe.Pointer(&me.data[0])), len(me.data)/2)
		return wstr.WinSliceToGo(str16), true
	}
	return "", false
}

// Builds a [RegVal] from a data block, with a few validations.
func regValParse(data []byte, regType co.REG) (RegVal, error) {
	isDword := regType == co.REG_DWORD || regType == co.REG_DWORD_BIG_ENDIAN
	isQword := regType == co.REG_QWORD
	sz := uintptr(len(data))

	if (isDword && sz != unsafe.Sizeof(uint32(0))) ||
		(isQword && sz != unsafe.Sizeof(uint64(0))) { // validate integer sizes
		return RegVal{}, co.ERROR_INVALID_DATA
	}

	return RegVal{regType, data}, nil
}
