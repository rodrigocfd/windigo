//go:build windows

package win

import (
	"fmt"
	"io"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// Reads all the contents of the file at once, immediately. Calls:
//   - [CreateFile]
//   - [HFILE.GetFileSizeEx]
//   - [HFILE.ReadFile]
//   - [HFILE.CloseHandle]
//
// # Example
//
//	contents, _ := win.FileReadNow("C:\\Temp\\foo.txt")
func FileReadNow(filePath string) ([]byte, error) {
	fin, err := FileOpen(filePath, co.FOPEN_READ_EXISTING)
	if err != nil {
		return nil, fmt.Errorf("FileOpen: %w", err)
	}
	defer fin.Close()

	sz, err := fin.Size()
	if err != nil {
		return nil, fmt.Errorf("File.Size: %w", err)
	}

	ret := make([]byte, sz)
	if _, err := fin.Read(ret); err != nil {
		return nil, fmt.Errorf("File.Read: %w", err)
	}

	return ret, nil
}

// Truncates the file, then writes all the contents at once, immediately. Calls:
//   - [CreateFile]
//   - [HFILE.SetEndOfFile]
//   - [HFILE.WriteFile]
//   - [HFILE.CloseHandle]
//
// # Example
//
//	contents := []byte("my text")
//	_ = win.FileWriteNow("C:\\Temp\\foo.txt", contents)
func FileWriteNow(filePath string, contents []byte) error {
	fout, err := FileOpen(filePath, co.FOPEN_RW_OPEN_OR_CREATE)
	if err != nil {
		return fmt.Errorf("FileOpen: %w", err)
	}
	defer fout.Close()

	if err := fout.Hfile().SetEndOfFile(); err != nil {
		return fmt.Errorf("HFILE.SetEndOfFile: %w", err)
	}

	if _, err := fout.Write(contents); err != nil {
		return fmt.Errorf("File.Write: %w", err)
	}

	return nil
}

// High-level abstraction to [HFILE], providing several operations.
//
// If you simply need to read or write the contents, consider using the
// [FileRead] and [FileWrite] functions.
//
// Implements the following standard io interfaces:
//   - [io.ByteReader]
//   - [io.ByteWriter]
//   - [io.Closer]
//   - [io.Reader]
//   - [io.Seeker]
//   - [io.StringWriter]
//   - [io.Writer]
//
// Created with [FileOpen].
//
// # Example
//
//	f, _ := win.FileOpen("C:\\Temp\\foo.txt", co.FOPEN_RW_OPEN_OR_CREATE)
//	defer f.Close()
//
//	fmt.Fprintf(f, "foo")
type File struct {
	hFile HFILE
}

// Opens a file with [CreateFile], returning a new high-level File object.
//
// ⚠️ You must defer [File.Close].
//
// # Example
//
//	f, _ := win.FileOpen("C:\\Temp\\foo.txt", co.FOPEN_RW_OPEN_OR_CREATE)
//	defer f.Close()
func FileOpen(filePath string, desiredAccess co.FOPEN) (*File, error) {
	var access co.GENERIC
	var share co.FILE_SHARE
	var disposition co.DISPOSITION

	switch desiredAccess {
	case co.FOPEN_READ_EXISTING:
		access = co.GENERIC_READ
		share = co.FILE_SHARE_READ
		disposition = co.DISPOSITION_OPEN_EXISTING
	case co.FOPEN_RW_EXISTING:
		access = co.GENERIC_READ | co.GENERIC_WRITE
		share = co.FILE_SHARE_NONE
		disposition = co.DISPOSITION_OPEN_EXISTING
	case co.FOPEN_RW_OPEN_OR_CREATE:
		access = co.GENERIC_READ | co.GENERIC_WRITE
		share = co.FILE_SHARE_NONE
		disposition = co.DISPOSITION_OPEN_ALWAYS
	case co.FOPEN_RW_CREATE:
		access = co.GENERIC_READ | co.GENERIC_WRITE
		share = co.FILE_SHARE_NONE
		disposition = co.DISPOSITION_CREATE_NEW
	}

	hFile, err := CreateFile(filePath, access, share, nil,
		disposition, co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE,
		co.SECURITY_NONE, 0)
	if err != nil {
		return nil, fmt.Errorf("CreateFile: %w", err)
	}

	return &File{hFile: hFile}, nil
}

// Implements [io.Closer].
//
// Calls [HFILE.CloseHandle].
func (me *File) Close() error {
	var e error
	if me.hFile != 0 {
		e = me.hFile.CloseHandle()
		me.hFile = 0
	}
	return e
}

// Returns the underlying handle.
func (me *File) Hfile() HFILE {
	return me.hFile
}

// Implements [io.Reader].
//
// Calls [HFILE.ReadFile] to read the file contents from its current internal
// pointer up to the buffer size.
func (me *File) Read(p []byte) (numBytesRead int, wErr error) {
	numRead, wErr := me.hFile.ReadFile(p, nil)
	if wErr != nil {
		return 0, fmt.Errorf("ReadFile: %w", wErr)
	}

	if numRead < uint(len(p)) { // surely there's no more to read
		return int(numRead), io.EOF
	} else if numRead == 0 { // EOF found
		return 0, io.EOF
	} else {
		return int(numRead), nil
	}
}

// Rewinds the internal file pointer and reads all contents at once, then
// rewinds the pointer again. Returns a []byte with the contents.
//
// Calls [HFILE.SetFilePointerEx] and [HFILE.ReadFile].
func (me *File) ReadAllAsSlice() ([]byte, error) {
	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	fileSize, err := me.Size()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, fileSize)
	if _, err := me.Read(buf); err != nil {
		return nil, err
	}

	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	return buf, nil
}

// Rewinds the internal file pointer and reads all contents at once, then
// rewinds the pointer again. Returns a [Vec] with the contents.
//
// Calls [HFILE.SetFilePointerEx] and [HFILE.ReadFile].
//
// ⚠️ You must defer [Vec.Free] on the returned Vec.
//
// # Example
//
//	f, _ := win.FileOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
//
//	data, _ := f.ReadAllAsVec()
//	defer data.Free()
//
//	txt := string(data.HotSlice())
//	println(txt)
//
// [SetFilePointerEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (me *File) ReadAllAsVec() (Vec[byte], error) {
	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return Vec[byte]{}, err
	}

	fileSize, err := me.Size()
	if err != nil {
		return Vec[byte]{}, err
	}

	heapBuf := NewVecSized[byte](fileSize, 0x00)
	if _, err := me.Read(heapBuf.HotSlice()); err != nil {
		heapBuf.Free()
		return Vec[byte]{}, err
	}

	if _, err := me.Seek(0, io.SeekStart); err != nil {
		heapBuf.Free()
		return Vec[byte]{}, err
	}

	return heapBuf, nil
}

// Implements [io.ByteReader].
//
// Calls [HFILE.ReadFile].
func (me *File) ReadByte() (byte, error) {
	var buf [1]byte
	_, err := me.Read(buf[:])
	return buf[0], err
}

// Truncates or expands the file, according to the new size. Zero will empty the
// file. The internal file pointer will rewind.
//
// Calls [HFILE.SetFilePointerEx] and [HFILE.SetEndOfFile].
func (me *File) Resize(numBytes uint) error {
	// Simply go beyond file limits.
	if _, err := me.Seek(int(numBytes), io.SeekStart); err != nil {
		return err
	}

	if err := me.hFile.SetEndOfFile(); err != nil {
		return fmt.Errorf("SetEndOfFile: %w", err)
	}

	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return err
	}

	return nil
}

// Implements [io.Seeker].
//
// Moves the internal pointer with [HFILE.SetFilePointerEx].
func (me *File) Seek(offset int, whence int) (int64, error) {
	var moveMethod co.FILE_FROM
	switch whence {
	case io.SeekCurrent:
		moveMethod = co.FILE_FROM_CURRENT
	case io.SeekStart:
		moveMethod = co.FILE_FROM_BEGIN
	case io.SeekEnd:
		moveMethod = co.FILE_FROM_END
	}

	newOff, err := me.hFile.SetFilePointerEx(offset, moveMethod)
	if err != nil {
		return 0, fmt.Errorf("SetFilePointerEx: %w", err)
	}
	return int64(newOff), nil
}

// Retrieves the file size with [HFILE.GetFileSizeEx]. This value is not cached.
func (me *File) Size() (uint, error) {
	sz, err := me.hFile.GetFileSizeEx()
	if err != nil {
		return 0, fmt.Errorf("GetFileSizeEx: %w", err)
	}
	return sz, nil
}

// Implements [io.Writer].
//
// Calls [HFILE.WriteFile] to write a slice at current internal pointer
// position.
func (me *File) Write(p []byte) (n int, wErr error) {
	written, wErr := me.hFile.WriteFile(p, nil)
	if wErr != nil {
		return 0, fmt.Errorf("WriteFile: %w", wErr)
	}
	return int(written), nil
}

// Implements [io.ByteWriter].
//
// Calls [HFILE.WriteFile] to write a byte at current internal pointer position.
func (me *File) WriteByte(c byte) error {
	_, err := me.Write([]byte{c})
	return err
}

// Implements [io.StringWriter].
//
// Calls [HFILE.WriteFile] to write a string at current internal pointer
// position.
func (me *File) WriteString(s string) (int, error) {
	serialized := []byte(s)
	return me.Write(serialized)
}

// High-level abstraction to [HFILEMAP], providing several operations.
//
// Note that memory-mapped files may present issues in x86 architectures; if so,
// just use the ordinary File.
//
// If you simply need to read or write the contents, consider using the
// [FileRead] and [FileWrite] functions.
//
// Created with [FileMapOpen].
type FileMap struct {
	objFile  *File
	hMap     HFILEMAP
	pMem     HFILEMAPVIEW
	sz       uint
	readOnly bool
}

// Opens a memory-mapped file, returning a new high-level FileMap object.
//
// Note that memory-mapped files may present issues in x86 architectures; if so,
// just use the ordinary FileOpen.
//
// ⚠️ You must defer [FileMap.Close].
//
// # Example
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
func FileMapOpen(filePath string, desiredAccess co.FOPEN) (*FileMap, error) {
	objFile, err := FileOpen(filePath, desiredAccess)
	if err != nil {
		return nil, err
	}

	me := &FileMap{
		objFile:  objFile,
		hMap:     HFILEMAP(0),
		pMem:     HFILEMAPVIEW(0),
		sz:       0,
		readOnly: desiredAccess == co.FOPEN_READ_EXISTING,
	}

	if err := me.mapInMemory(); err != nil {
		me.Close()
		return nil, err
	}
	return me, nil
}

// Unmaps and releases the file resource.
func (me *FileMap) Close() error {
	var e1, e2, e3 error
	if me.pMem != 0 {
		e1 = me.pMem.UnmapViewOfFile()
		me.pMem = 0
	}
	if me.hMap != 0 {
		e2 = me.hMap.CloseHandle()
		me.hMap = 0
	}
	e3 = me.objFile.Close()
	me.sz = 0

	if e1 != nil {
		return e1
	} else if e2 != nil {
		return e2
	} else {
		return e3
	}
}

// Returns a slice to the memory-mapped bytes.
//
// The FileMap object must remain open while the slice is being used.
func (me *FileMap) HotSlice() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(me.pMem)), me.sz)
}

// Returns a new []byte with a copy of all data in the file.
func (me *FileMap) ReadAllAsSlice() []byte {
	return me.ReadChunkAsSlice(0, me.sz)
}

// Returns a new [Vec] with a copy of all data in the file.
//
// ⚠️ You must defer [Vec.Free] on the returned Vec.
//
// # Example
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
//
//	data := f.ReadAllAsVec()
//	defer data.Free()
//
//	txt := string(data.HotSlice())
//	println(txt)
func (me *FileMap) ReadAllAsVec() Vec[byte] {
	return me.ReadChunkAsVec(0, me.sz)
}

// Returns a new []byte with a copy of data, start with offset, and with the
// given length.
func (me *FileMap) ReadChunkAsSlice(offset, length uint) []byte {
	hotSlice := me.HotSlice()
	buf := make([]byte, length)
	copy(buf, hotSlice[offset:offset+length])
	return buf
}

// Returns a [Vec] with a copy of the data, start with offset, and with the
// given length.
//
// ⚠️ You must defer [Vec.Free] on the returned Vec.
//
// # Example
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
//
//	data := f.ReadChunkAsVec(0, 30)
//	defer data.Free()
//
//	txt := string(data.HotSlice())
//	println(txt)
func (me *FileMap) ReadChunkAsVec(offset, length uint) Vec[byte] {
	hotSlice := me.HotSlice()
	heapBuf := NewVec[byte]()
	heapBuf.Append(hotSlice[offset : offset+length]...)
	return heapBuf
}

// Truncates or expands the file, according to the new size. Zero will empty the
// file.
//
// Internally, the file is unmapped, then remapped back into memory.
func (me *FileMap) Resize(numBytes uint) error {
	me.pMem.UnmapViewOfFile()
	me.hMap.CloseHandle()
	if err := me.objFile.Resize(numBytes); err != nil {
		return err
	}
	return me.mapInMemory()
}

// Retrieves the file size. This value is cached.
func (me *FileMap) Size() uint {
	return me.sz
}

func (me *FileMap) mapInMemory() error {
	// Mapping into memory.
	pageFlags := co.PAGE_READONLY
	if !me.readOnly {
		pageFlags = co.PAGE_READWRITE
	}

	var err error
	me.hMap, err = me.objFile.Hfile().
		CreateFileMapping(nil, pageFlags, co.SEC_NONE, 0, "")
	if err != nil {
		return fmt.Errorf("CreateFileMapping: %w", err)
	}

	// Get pointer to data block.
	mapFlags := co.FILE_MAP_READ
	if !me.readOnly {
		mapFlags = co.FILE_MAP_WRITE
	}

	if me.pMem, err = me.hMap.MapViewOfFile(mapFlags, 0, 0); err != nil {
		return fmt.Errorf("MapViewOfFile: %w", err)
	}

	// Cache file size.
	me.sz, err = me.objFile.Size()
	if err != nil {
		return fmt.Errorf("Size: %w", err)
	}

	return nil // file mapped successfully
}
