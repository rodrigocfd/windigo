/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Manages a memory-mapped file resource.
type FileMapped struct {
	objFile  File
	hMap     win.HFILEMAP
	pMem     win.HFILEMAP_PTR
	sz       uint
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

// Returns a slice to the memory-mapped bytes. The FileMapped object must remain
// open while the slice is being used.
//
// If you need to close the file right away, use CopyToBuffer() instead.
func (me *FileMapped) HotSlice() []byte {
	// https://stackoverflow.com/a/43592538
	// https://golang.org/pkg/internal/unsafeheader/#Slice
	var sliceMem = struct { // slice memory layout
		addr unsafe.Pointer
		len  int
		cap  int
	}{unsafe.Pointer(me.pMem), int(me.sz), int(me.sz)}

	return *(*[]byte)(unsafe.Pointer(&sliceMem))
}

func (me *FileMapped) OpenExistingForRead(path string) *win.WinError {
	return me.rawOpen(path, true)
}

func (me *FileMapped) OpenExistingForReadWrite(path string) *win.WinError {
	return me.rawOpen(path, false)
}

// Copies all file data into a []byte and returns it.
func (me *FileMapped) CopyAllData() []byte {
	return me.CopyDataChunk(0, me.sz)
}

// Copies file data into a []byte and returns it, starting from offset, with
// given length.
func (me *FileMapped) CopyDataChunk(offset, length uint) []byte {
	hotSlice := me.HotSlice()
	buf := make([]byte, length)
	copy(buf, hotSlice[offset:offset+length])
	return buf
}

// Truncates or expands the file, according to the new size. Zero will empty the
// file.
//
// Internally, the file is unmapped, then remapped back into memory.
func (me *FileMapped) SetSize(numBytes uint) *win.WinError {
	me.pMem.UnmapViewOfFile()
	me.hMap.CloseHandle()
	if err := me.objFile.SetSize(numBytes); err != nil {
		return err
	}
	return me.mapInMemory()
}

// Retrieves the file size. This value is cached.
func (me *FileMapped) Size() uint {
	return me.sz
}

func (me *FileMapped) mapInMemory() *win.WinError {
	// Mapping into memory.
	pageFlags := co.PAGE_READWRITE
	if me.readOnly {
		pageFlags = co.PAGE_READONLY
	}

	var err *win.WinError
	me.hMap, err = me.objFile.hFile.CreateFileMapping(
		nil, pageFlags, co.SEC_NONE, 0, "")
	if err != nil {
		return err
	}

	// Get pointer to data block.
	mapFlags := co.FILE_MAP_WRITE
	if me.readOnly {
		mapFlags = co.FILE_MAP_READ
	}
	me.pMem, err = me.hMap.MapViewOfFile(mapFlags, 0, 0)
	if err != nil {
		return err
	}

	// Cache file size.
	me.sz = me.objFile.Size()

	return nil // file mapped successfully
}

func (me *FileMapped) rawOpen(path string, readOnly bool) *win.WinError {
	me.Close()
	if readOnly {
		if err := me.objFile.OpenExistingForRead(path); err != nil {
			return err
		}
	} else {
		if err := me.objFile.OpenExistingForReadWrite(path); err != nil {
			return err
		}
	}
	me.readOnly = readOnly // keep flag
	return me.mapInMemory()
}
