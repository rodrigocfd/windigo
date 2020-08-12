/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Manages a memory-mapped file resource.
type FileMapped struct {
	objFile  File
	hMap     win.HFILEMAP
	pMem     win.HFILEMAP_PTR
	sz       uint64
	readOnly bool // necessary for SetSize()
}

func (me *FileMapped) Close() {
	if me.pMem != 0 {
		me.pMem.UnmapViewOfFile()
		me.pMem = 0
	}
	if me.hMap != 0 {
		me.hMap.CloseHandle()
		me.hMap = 0
	}
	me.objFile.Close()
	me.sz = 0
}

// Allocates a byte buffer and copies all the data into it.
func (me *FileMapped) CopyAllToBuffer() []byte {
	return me.CopyToBuffer(0, me.sz)
}

// Allocates a byte buffer and copies data into it.
func (me *FileMapped) CopyToBuffer(offset, length uint64) []byte {
	hotSlice := me.HotSlice()
	buf := make([]byte, length)
	copy(buf, hotSlice[offset:offset+length])
	return buf
}

// Returns a slice to the memory-mapped bytes. The FileMapped object must remain
// open while the slice is being used.
//
// To close the file and still work on the data, use CopyToBuffer().
func (me *FileMapped) HotSlice() []byte {
	// https://golang.org/src/syscall/syscall_unix.go#L52
	var sliceMem = struct {
		addr uintptr
		len  int
		cap  int
	}{uintptr(me.pMem), int(me.sz), int(me.sz)}

	return *(*[]byte)(unsafe.Pointer(&sliceMem))
}

func (me *FileMapped) OpenExistingForRead(path string) *FileMapped {
	return me.rawOpen(path, true)
}

func (me *FileMapped) OpenExistingForReadWrite(path string) *FileMapped {
	return me.rawOpen(path, false)
}

// Truncates or expands the file, according to the new size. Zero will empty the
// file.
//
// Internally, the file is unmapped, then remapped back into memory.
func (me *FileMapped) SetSize(numBytes uint64) *FileMapped {
	me.pMem.UnmapViewOfFile()
	me.hMap.CloseHandle()
	me.objFile.SetSize(numBytes)
	return me.mapInMemory()
}

// Retrieves the file size. This value is cached.
func (me *FileMapped) Size() uint64 {
	return me.sz
}

func (me *FileMapped) mapInMemory() *FileMapped {
	// Mapping into memory.
	pageFlags := co.PAGE_READWRITE
	if me.readOnly {
		pageFlags = co.PAGE_READONLY
	}
	me.hMap = me.objFile.hFile.CreateFileMapping(
		nil, pageFlags, co.SEC_NONE, 0, "")

	// Get pointer to data block.
	mapFlags := co.FILE_MAP_WRITE
	if me.readOnly {
		mapFlags = co.FILE_MAP_READ
	}
	me.pMem = me.hMap.MapViewOfFile(mapFlags, 0, 0)

	// Cache file size.
	me.sz = me.objFile.Size()

	return me
}

func (me *FileMapped) rawOpen(path string, readOnly bool) *FileMapped {
	me.Close()
	if readOnly {
		me.objFile.OpenExistingForRead(path)
	} else {
		me.objFile.OpenExistingForReadWrite(path)
	}
	me.readOnly = readOnly
	return me.mapInMemory()
}
