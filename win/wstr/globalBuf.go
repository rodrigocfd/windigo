//go:build windows

package wstr

import (
	"sync/atomic"
)

const BUF_MAX uint = 260 // Size of the global buffer.

// The global static buffer is the primary place to store Windows UTF-16 chars
// under conversion. It's protected from concurrent usage.
type _GlobalBuf struct {
	emptyBuf  [1]uint16       // Global buffer for a valid, empty string.
	buf       [BUF_MAX]uint16 // Global buffer to store the strings.
	inUse     atomic.Bool     // Locks the global buffer, so 1 instance can use it at once.
	idxNextCh int             // Index of next empty uint16 char within globalBuf.
}

var globalBuf _GlobalBuf // Global static buffer object.

// Attempts to get the lock to the global buffer, returning true if successful.
func (me *_GlobalBuf) tryGet() bool {
	return me.inUse.CompareAndSwap(false, true)
}

// Releases the lock to the global buffer.
func (me *_GlobalBuf) release() {
	me.idxNextCh = 0
	me.inUse.Store(false)
}

// Returns true if the global buffer, at its current position, has enough room.
func (me *_GlobalBuf) canFit(numChars int) bool {
	return len(me.buf[me.idxNextCh:]) >= numChars
}
