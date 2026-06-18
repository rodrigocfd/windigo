//go:build windows

package win

import (
	"time"

	"github.com/rodrigocfd/windigo/internal/utl"
)

// Tagged union for a timeout value, which can be:
//   - [time.Duration]
//   - infinite
//
// Example:
//
//	timeout := TimeoutDur(4 * time.Second)
//
//	if dur, ok := timeout.Dur(); ok {
//		println(dur)
//	}
type Timeout struct {
	tag _TagTimeout
	ms  uint32
}

type _TagTimeout uint8

const (
	_TagTimeoutMs _TagTimeout = iota
	_TagTimeoutInfinite
)

// Constructs a new [Timeout] with the given [time.Duration].
func TimeoutDur(dur time.Duration) Timeout {
	return Timeout{
		tag: _TagTimeoutMs,
		ms:  uint32(dur.Milliseconds()),
	}
}

// If the value is [time.Duration], returns it and true.
func (me *Timeout) Dur() (time.Duration, bool) {
	if me.tag == _TagTimeoutMs {
		return time.Duration(me.ms) * time.Millisecond, true
	}
	return time.Duration(0), false
}

// Constructs a new [Timeout] with infinite value.
func TimeoutInfinite() Timeout {
	return Timeout{
		tag: _TagTimeoutInfinite,
		ms:  utl.INFINITE,
	}
}

// Returns true if infinite.
func (me *Timeout) IsInfinite() bool {
	return me.tag == _TagTimeoutInfinite
}

// Returns the internal value as uint32.
func (me *Timeout) raw() uint32 {
	return me.ms
}
