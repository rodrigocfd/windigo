//go:build windows

package wstr

import (
	"unsafe"
)

// Buffer to receive Windows UTF-16 strings.
type BufReceiver struct {
	gotGlobal bool // Do we have the lock to the global buffer?
	faceSize  uint // Declared size, which can be smaller than actual buffer size.
	localBuf  []uint16
}

// Constructs a new buffer to receive Windows UTF-16 strings, with the given
// initial length. Tries to use the global buffer; if already in use, allocates
// a dynamic buffer.
func NewBufReceiver(numChars uint) BufReceiver {
	me := BufReceiver{
		gotGlobal: globalBuf.tryGet(),
	}
	me.Resize(numChars)
	return me
}

// Releases the lock for the buffer. No further strings should be added: the
// object must be discarded.
func (me *BufReceiver) Free() {
	if me.gotGlobal {
		globalBuf.release()
	}
	me.localBuf = nil
}

// Returns the size of the receiving buffer.
func (me *BufReceiver) Len() uint {
	return me.faceSize
}

// Resizes the receiving buffer to the given number of chars. Always grows.
func (me *BufReceiver) Resize(numChars uint) {
	if me.gotGlobal {
		if !globalBuf.canFit(int(numChars)) {
			me.localBuf = make([]uint16, numChars)
			copy(me.localBuf, globalBuf.buf[:])
			globalBuf.release() // we won't use the global buffer anymore
			me.gotGlobal = false
		}
	} else {
		if numChars > uint(len(me.localBuf)) { // requesting a buffer even larger
			newLocalBuf := make([]uint16, numChars)
			copy(newLocalBuf, me.localBuf)
			me.localBuf = newLocalBuf
		}
	}
	me.faceSize = numChars
}

// Converts the receiving buffer content to a Go string.
func (me *BufReceiver) String() string {
	return WinSliceToGo(me.HotSlice())
}

// Returns a slice over the block.
func (me *BufReceiver) HotSlice() []uint16 {
	if me.gotGlobal {
		return globalBuf.buf[:me.faceSize]
	} else {
		return me.localBuf[:me.faceSize]
	}
}

// Returns the *uint16 to the beginning of the block.
func (me *BufReceiver) UnsafePtr() unsafe.Pointer {
	if me.gotGlobal {
		return unsafe.Pointer(&globalBuf.buf[0])
	} else {
		return unsafe.Pointer(&me.localBuf[0])
	}
}
