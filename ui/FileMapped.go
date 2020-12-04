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

// Used in OpenFileMapped().
//
// Behavior of the file opening.
type FILEMAP_MODE uint8

const (
	FILEMAP_MODE_R  FILEMAP_MODE = iota // Open an existing file for read only.
	FILEMAP_MODE_RW                     // Open an existing file for read and write.
)

//------------------------------------------------------------------------------

// Manages a memory-mapped file resource.
type FileMapped struct {
	objFile  *File
	hMap     win.HFILEMAP
	pMem     win.HFILEMAP_PTR
	sz       int
	readOnly bool // necessary for SetSize()
}

// Constructor.
//
// You must defer Close().
func OpenFileMapped(path string, behavior FILEMAP_MODE) (*FileMapped, error) {
	var fBeh FILE_MODE
	var readOnly bool

	switch behavior {
	case FILEMAP_MODE_R:
		fBeh = FILE_MODE_R_EXISTING
		readOnly = true
	case FILEMAP_MODE_RW:
		fBeh = FILE_MODE_RW_EXISTING
		readOnly = false
	}

	objFile, err := OpenFile(path, fBeh)
	if err != nil {
		return nil, err
	}

	me := FileMapped{
		objFile:  objFile,
		readOnly: readOnly,
	}

	if err := me.mapInMemory(); err != nil {
		return nil, err
	}

	return &me, nil
}

// Unmaps and frees the file resource.
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

// Returns a new []byte with a copy of all data in the file.
func (me *FileMapped) ReadAll() []byte {
	return me.ReadChunk(0, me.sz)
}

// Returns a new []byte with a copy of data, start with offset, and with the
// given length.
func (me *FileMapped) ReadChunk(offset, length int) []byte {
	hotSlice := me.HotSlice()
	buf := make([]byte, length)
	copy(buf, hotSlice[offset:offset+length])
	return buf
}

// Truncates or expands the file, according to the new size. Zero will empty the
// file.
//
// Internally, the file is unmapped, then remapped back into memory.
func (me *FileMapped) Resize(numBytes int) error {
	me.pMem.UnmapViewOfFile()
	me.hMap.CloseHandle()
	if err := me.objFile.Resize(numBytes); err != nil {
		return err
	}
	return me.mapInMemory()
}

// Retrieves the file size. This value is cached.
func (me *FileMapped) Size() int {
	return me.sz
}

func (me *FileMapped) mapInMemory() error {
	// Mapping into memory.
	pageFlags := co.PAGE_READWRITE
	if me.readOnly {
		pageFlags = co.PAGE_READONLY
	}

	var err error
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
