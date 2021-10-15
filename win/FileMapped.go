package win

import (
	"strings"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to HFILEMAP, providing several operations.
//
// Created with OpenFileMapped().
type FileMapped interface {
	// Unmaps and releases the file resource.
	Close()

	// Returns a slice to the memory-mapped bytes. The FileMapped object must
	// remain open while the slice is being used.
	//
	// If you need to close the file right away, use CopyToBuffer() instead.
	HotSlice() []byte

	// Returns a new []byte with a copy of all data in the file.
	ReadAll() []byte

	// Returns a new []byte with a copy of data, start with offset, and with the
	// given length.
	ReadChunk(offset, length int) []byte

	// Parses the file content as text and returns the lines.
	ReadLines() []string

	// Truncates or expands the file, according to the new size. Zero will empty
	// the file.
	//
	// Internally, the file is unmapped, then remapped back into memory.
	Resize(numBytes int) error

	// Retrieves the file size. This value is cached.
	Size() int
}

//------------------------------------------------------------------------------

type _FileMapped struct {
	objFile  File
	hMap     HFILEMAP
	pMem     HFILEMAPVIEW
	sz       int
	readOnly bool
}

// Opens a memory-mapped file, returning a new high-level FileMapped object.
//
// ⚠️ You must defer FileMapped.Close().
func OpenFileMapped(
	filePath string, desiredAccess co.OPEN_FILE) (FileMapped, error) {

	objFile, err := OpenFile(filePath, desiredAccess)
	if err != nil {
		return nil, err
	}

	me := &_FileMapped{
		objFile:  objFile,
		hMap:     HFILEMAP(0),
		pMem:     HFILEMAPVIEW(0),
		sz:       0,
		readOnly: desiredAccess == co.OPEN_FILE_READ_EXISTING,
	}

	if err := me.mapInMemory(); err != nil {
		me.Close()
		return nil, err
	}
	return me, nil
}

func (me *_FileMapped) Close() {
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

func (me *_FileMapped) HotSlice() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(me.pMem)), me.sz)
}

func (me *_FileMapped) ReadAll() []byte {
	return me.ReadChunk(0, me.sz)
}

func (me *_FileMapped) ReadChunk(offset, length int) []byte {
	hotSlice := me.HotSlice()
	buf := make([]byte, length)
	copy(buf, hotSlice[offset:offset+length])
	return buf
}

func (me *_FileMapped) ReadLines() []string {
	allText := string(me.HotSlice())
	lines := strings.Split(allText, "\n")

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return lines
}

func (me *_FileMapped) Resize(numBytes int) error {
	me.pMem.UnmapViewOfFile()
	me.hMap.CloseHandle()
	if err := me.objFile.Resize(numBytes); err != nil {
		return err
	}
	return me.mapInMemory()
}

func (me *_FileMapped) Size() int {
	return me.sz
}

func (me *_FileMapped) mapInMemory() error {
	// Mapping into memory.
	pageFlags := util.Iif(me.readOnly,
		co.PAGE_READONLY, co.PAGE_READWRITE).(co.PAGE)

	var err error
	me.hMap, err = me.objFile.Hfile().
		CreateFileMapping(nil, pageFlags, co.SEC_NONE, 0, nil)
	if err != nil {
		return err
	}

	// Get pointer to data block.
	mapFlags := util.Iif(me.readOnly,
		co.FILE_MAP_READ, co.FILE_MAP_WRITE).(co.FILE_MAP)

	if me.pMem, err = me.hMap.MapViewOfFile(mapFlags, 0, 0); err != nil {
		return err
	}

	// Cache file size.
	me.sz = me.objFile.Size()

	return nil // file mapped successfully
}
