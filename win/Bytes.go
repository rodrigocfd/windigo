package win

import (
	"encoding/binary"
)

type _BytesT struct{}

// Byte operations.
var Bytes _BytesT

// Appends an uint16 onto a []byte, returning the newly allocated slice.
func (_BytesT) Append16(
	dest []byte, encoding binary.ByteOrder, val uint16) []byte {

	buf := [2]byte{}
	encoding.PutUint16(buf[:], val)
	return append(dest, buf[:]...)
}

// Appends an uint32 onto a []byte, returning the newly allocated slice.
func (_BytesT) Append32(
	dest []byte, encoding binary.ByteOrder, val uint32) []byte {

	buf := [4]byte{}
	encoding.PutUint32(buf[:], val)
	return append(dest, buf[:]...)
}

// Appends an uint64 onto a []byte, returning the newly allocated slice.
func (_BytesT) Append64(
	dest []byte, encoding binary.ByteOrder, val uint64) []byte {

	buf := [8]byte{}
	encoding.PutUint64(buf[:], val)
	return append(dest, buf[:]...)
}

// Extracts the high-order uint8 from an uint16.
//
// Same as HIBYTE macro.
func (_BytesT) Hi8(value uint16) uint8 {
	return uint8(value >> 8 & 0xff)
}

// Extracts the high-order uint16 from an uint32.
//
// Same as HIWORD macro.
func (_BytesT) Hi16(value uint32) uint16 {
	return uint16(value >> 16 & 0xffff)
}

// Extracts the high-order uint32 from an uint64.
func (_BytesT) Hi32(value uint64) uint32 {
	return uint32(value >> 32 & 0xffff_ffff)
}

// Extracts the low-order uint8 from an uint16.
//
// Same as LOBYTE macro.
func (_BytesT) Lo8(value uint16) uint8 {
	return uint8(value & 0xff)
}

// Extracts the low-order uint16 from an uint32.
//
// Same as LOWORD macro.
func (_BytesT) Lo16(value uint32) uint16 {
	return uint16(value & 0xffff)
}

// Extracts the low-order uint32 from an uint64.
func (_BytesT) Lo32(value uint64) uint32 {
	return uint32(value & 0xffff_ffff)
}

// Assembles an uint16 from two uint8.
//
// Same as MAKEWORD macro.
func (_BytesT) Make16(lo, hi uint8) uint16 {
	return (uint16(lo) & 0xff) | ((uint16(hi) & 0xff) << 8)
}

// Assembles an uint32 from two uint16.
//
// Same as MAKELONG, MAKEWPARAM and MAKELPARAM macros.
func (_BytesT) Make32(lo, hi uint16) uint32 {
	return (uint32(lo) & 0xffff) | ((uint32(hi) & 0xffff) << 16)
}

// Assembles an uint64 from two uint32.
func (_BytesT) Make64(lo, hi uint32) uint64 {
	return (uint64(lo) & 0xffff_ffff) | ((uint64(hi) & 0xffff_ffff) << 32)
}

// Tells whether the number has the nth bit set.
//
// bitPosition must be in the range 0-7.
func (_BytesT) HasBit(number, bitPosition uint8) bool {
	return (number & (1 << bitPosition)) > 0
}

// Returns a new number with the nth bit set or clear.
//
// bitPosition must be in the range 0-7.
func (_BytesT) SetBit(number, bitPosition uint8, doSet bool) uint8 {
	if doSet {
		return number | (1 << bitPosition)
	} else {
		return number &^ (1 << bitPosition)
	}
}
