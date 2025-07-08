//go:build windows

package wstr

import (
	"unsafe"
)

// Buffer to convert Go strings into null-terminated Windows UTF-16 strings.
type BufConverter struct {
	gotGlobal bool // Do we have the lock to the global buffer?
	localBufs [][]uint16
}

// Constructs a new buffer to convert Go strings into null-terminated Windows
// UTF-16 strings. Tries to use the global buffer; if already in use, allocates
// a dynamic buffer.
func NewBufConverter() BufConverter {
	return BufConverter{
		gotGlobal: globalBuf.tryGet(),
	}
}

// Releases the lock for the buffer. No further strings should be added: the
// object must be discarded.
func (me *BufConverter) Free() {
	if me.gotGlobal {
		globalBuf.release()
	}
	me.localBufs = nil
}

// Converts a Go string into a null-terminated Windows UTF-16 string, storing it
// in the internal buffer, returning its []uint16.
//
// If the string is empty, allocates an 1-char buffer with just the terminating
// null, returning a pointer to it.
func (me *BufConverter) SliceAllowEmpty(s string) []uint16 {
	if s == "" {
		return globalBuf.emptyBuf[:] // for empty strings, always return the same buffer
	}
	return unsafe.Slice(me.add(s))
}

// Converts a Go string into a null-terminated Windows UTF-16 string, storing it
// in the internal buffer, returning its *uint16.
//
// If the string is empty, allocates an 1-char buffer with just the terminating
// null, returning a pointer to it.
func (me *BufConverter) PtrAllowEmpty(s string) unsafe.Pointer {
	if s == "" {
		return unsafe.Pointer(&globalBuf.emptyBuf[0]) // for empty strings, always return the same buffer
	}
	ptr, _ := me.add(s)
	return unsafe.Pointer(ptr)
}

// Converts a Go string into a null-terminated Windows UTF-16 string, storing it
// in the internal buffer, returning its *uint16.
//
// If the string is empty, returns a nil pointer.
func (me *BufConverter) PtrEmptyIsNil(s string) unsafe.Pointer {
	if s == "" {
		return nil
	}
	ptr, _ := me.add(s)
	return unsafe.Pointer(ptr)
}

func (me *BufConverter) add(s string) (*uint16, int) {
	numChars := len([]rune(s)) + 1 // plus terminating null
	if !me.gotGlobal || !globalBuf.canFit(numChars) {
		// This string won't fit the global buffer, create a local one for it.
		// Keep the lock to global buffer, because we may already have strings
		// in it, and future strings may also fit in it.
		newLocalBuf := make([]uint16, numChars)
		GoToWinBuf(s, newLocalBuf)
		me.localBufs = append(me.localBufs, newLocalBuf)
		return &newLocalBuf[0], numChars
	} else {
		// Put the string in the global buffer; this is the optimal case.
		idx0 := globalBuf.idxNextCh
		GoToWinBuf(s, globalBuf.buf[globalBuf.idxNextCh:])
		globalBuf.idxNextCh += numChars
		return &globalBuf.buf[idx0], numChars
	}
}

// Converts the Go strings into multiple null-terminated strings, sequentially,
// followed by another terminating null. Stores it in the internal buffer, and
// returns a *uint16 to the beginning of the block.
//
// If no strings, returns a nil pointer.
func (me *BufConverter) PtrMulti(ss ...string) unsafe.Pointer {
	if len(ss) == 0 {
		return nil
	}

	numChars := 1 // double terminating null
	for _, s := range ss {
		numChars += len([]rune(s)) + 1 // plus terminating null
	}

	if !me.gotGlobal || !globalBuf.canFit(numChars) {
		// This string won't fit the global buffer, create a local one for it.
		// Keep the lock to global buffer, because we may already have strings
		// in it, and future strings may also fit in it.
		newLocalBuf := make([]uint16, numChars)
		idxLocal := 0
		for _, s := range ss {
			GoToWinBuf(s, newLocalBuf[idxLocal:])
			idxLocal += len([]rune(s)) + 1
		}
		me.localBufs = append(me.localBufs, newLocalBuf)
		return unsafe.Pointer(&newLocalBuf[0])
	} else {
		// Put the block in the global buffer; this is the optimal case.
		idx0 := globalBuf.idxNextCh
		for _, s := range ss {
			GoToWinBuf(s, globalBuf.buf[globalBuf.idxNextCh:])
			globalBuf.idxNextCh += len([]rune(s)) + 1
		}
		globalBuf.buf[globalBuf.idxNextCh] = 0x0000 // double terminating null
		globalBuf.idxNextCh += 1
		return unsafe.Pointer(&globalBuf.buf[idx0])
	}
}

// Removes all strings in the buffer, potentially invalidating all pointers.
func (me *BufConverter) Clear() {
	if me.gotGlobal {
		globalBuf.idxNextCh = 0
	}
	me.localBufs = nil
}
