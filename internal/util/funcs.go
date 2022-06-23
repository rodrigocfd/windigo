//go:build windows

package util

import (
	"encoding/binary"
	"strings"
	"time"
)

// Tells whether the number has the nth bit set.
//
// bitPosition must be in the range 0-7.
func BitIsSet(number, bitPosition uint8) bool {
	return (number & (1 << bitPosition)) > 0
}

// Returns a new number with the nth bit set or clear.
//
// bitPosition must be in the range 0-7.
func BitSet(number, bitPosition uint8, doSet bool) uint8 {
	if doSet {
		return number | (1 << bitPosition)
	} else {
		return number &^ (1 << bitPosition)
	}
}

// Syntactic sugar; converts bool to 0 or 1.
func BoolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func BoolToUint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// Returns first value if condition is true, otherwise returns second.
//
// Return type must be cast accordingly.
func Iif(cond bool, ifTrue, ifFalse interface{}) interface{} {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

// "&He && she" becomes "He & she".
func RemoveAccelAmpersands(text string) string {
	runes := []rune(text)
	buf := strings.Builder{}
	buf.Grow(len(runes)) // prealloc for performance

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == '&' && runes[i+1] != '&' {
			continue
		}
		buf.WriteRune(runes[i])
	}
	if runes[len(runes)-1] != '&' {
		buf.WriteRune(runes[len(runes)-1])
	}
	return buf.String()
}

// Reverses the bytes, not the bits.
func ReverseBytes64(n uint64) uint64 {
	var buf64 [8]byte
	binary.LittleEndian.PutUint64(buf64[:], n)
	return binary.BigEndian.Uint64(buf64[:])
}

//------------------------------------------------------------------------------

// Assembles an uint16 from two uint8.
func Make16(lo, hi uint8) uint16 {
	return (uint16(lo) & 0xff) | ((uint16(hi) & 0xff) << 8)
}

// Assembles an uint32 from two uint16.
func Make32(lo, hi uint16) uint32 {
	return (uint32(lo) & 0xffff) | ((uint32(hi) & 0xffff) << 16)
}

// Assembles an uint64 from two uint32.
func Make64(lo, hi uint32) uint64 {
	return (uint64(lo) & 0xffff_ffff) | ((uint64(hi) & 0xffff_ffff) << 32)
}

// Breaks an uint16 into low and high uint8.
func Break16(val uint16) (lo, hi uint8) {
	return uint8(val & 0xff), uint8(val >> 8 & 0xff)
}

// Breaks an uint32 into low and high uint16.
func Break32(val uint32) (lo, hi uint16) {
	return uint16(val & 0xffff), uint16(val >> 16 & 0xffff)
}

// Breaks an uint64 into low and high uint32.
func Break64(val uint64) (lo, hi uint32) {
	return uint32(val & 0xffff_ffff), uint32(val >> 32 & 0xffff_ffff)
}

//------------------------------------------------------------------------------

// Converts time.Duration to 100 nanoseconds.
func DurationToNano100(duration time.Duration) int64 {
	return int64(duration) * 10_000 / int64(time.Millisecond)
}

// Converts 100 nanoseconds to time.Duration.
func Nano100ToDuration(nanosec100 int64) time.Duration {
	return time.Duration(nanosec100 / 10_000 * int64(time.Millisecond))
}
