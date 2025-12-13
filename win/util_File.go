//go:build windows

package win

import (
	"fmt"
	"io"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Reads all the contents of the file at once, immediately. Calls:
//   - [CreateFile]
//   - [HFILE.GetFileSizeEx]
//   - [HFILE.ReadFile]
//   - [HFILE.CloseHandle]
//
// Example:
//
//	contents, _ := win.FileReadAll("C:\\Temp\\foo.txt")
func FileReadAll(filePath string) ([]byte, error) {
	fin, err := FileOpen(filePath, co.FOPEN_READ_EXISTING)
	if err != nil {
		return nil, fmt.Errorf("FileReadAll: %w", err)
	}
	defer fin.Close()

	sz, err := fin.Size()
	if err != nil {
		return nil, fmt.Errorf("FileReadAll: %w", err)
	}

	ret := make([]byte, sz)
	if _, err := fin.Read(ret); err != nil {
		return nil, fmt.Errorf("FileReadAll: %w", err)
	}

	return ret, nil
}

// Truncates the file, then writes all the contents at once, immediately. Calls:
//   - [CreateFile]
//   - [HFILE.SetEndOfFile]
//   - [HFILE.WriteFile]
//   - [HFILE.CloseHandle]
//
// Example:
//
//	contents := []byte("my text")
//	_ = win.FileWriteAll("C:\\Temp\\foo.txt", contents)
func FileWriteAll(filePath string, contents []byte) error {
	fout, err := FileOpen(filePath, co.FOPEN_RW_OPEN_OR_CREATE)
	if err != nil {
		return fmt.Errorf("FileWriteAll: %w", err)
	}
	defer fout.Close()

	if err := fout.Truncate(); err != nil {
		return fmt.Errorf("FileWriteAll: %w", err)
	}

	if _, err := fout.Write(contents); err != nil {
		return fmt.Errorf("FileWriteAll: %w", err)
	}

	return nil
}

// High-level abstraction to [HFILE], providing several operations.
//
// If you simply need to read or write the contents, consider using the
// [FileReadAll] and [FileWriteAll] functions.
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
// Example:
//
//	f, _ := win.FileOpen("C:\\Temp\\foo.txt", co.FOPEN_RW_OPEN_OR_CREATE)
//	defer f.Close()
//
//	fmt.Fprintf(f, "foo")
type File struct {
	hFile HFILE
}

// Constructs a new [File] by calling [CreateFile].
//
// ⚠️ You must defer [File.Close].
//
// Example:
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
		return nil, fmt.Errorf("FileOpen CreateFile: %w", err)
	}

	return &File{hFile}, nil
}

// Implements [io.Closer].
//
// Calls [HFILE.CloseHandle].
func (me *File) Close() error {
	if me.hFile != 0 {
		err := me.hFile.CloseHandle()
		me.hFile = 0
		if err != nil {
			return fmt.Errorf("File.Close HFILE.CloseHandle: %w", err)
		}
	}
	return nil
}

// Returns the underlying handle.
func (me *File) Hfile() HFILE {
	return me.hFile
}

// Implements [io.Reader].
//
// Calls [HFILE.ReadFile] to read the file contents from its current internal
// pointer up to the buffer size.
func (me *File) Read(p []byte) (int, error) {
	numRead, err := me.hFile.ReadFile(p, nil)
	if err != nil {
		return 0, fmt.Errorf("File.Read HFILE.ReadFile: %w", err)
	}

	if numRead < len(p) { // buffer not completely filled, surely there's no more to read
		return int(numRead), io.EOF
	} else if numRead == 0 { // EOF found
		return 0, io.EOF
	} else { // still more to read
		return int(numRead), nil
	}
}

// Rewinds the internal file pointer and reads all contents at once, then
// rewinds the pointer again. Returns a []byte with the contents.
//
// Calls [HFILE.SetFilePointerEx] and [HFILE.ReadFile].
func (me *File) ReadAllAsSlice() ([]byte, error) {
	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("File.ReadAllAsSlice: %w", err)
	}

	fileSize, err := me.Size()
	if err != nil {
		return nil, fmt.Errorf("File.ReadAllAsSlice: %w", err)
	}

	buf := make([]byte, fileSize)
	if _, err := me.Read(buf); err != nil {
		return nil, fmt.Errorf("File.ReadAllAsSlice: %w", err)
	}

	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("File.ReadAllAsSlice: %w", err)
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
// Example:
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
		return Vec[byte]{}, fmt.Errorf("File.ReadAllAsVec: %w", err)
	}

	fileSize, err := me.Size()
	if err != nil {
		return Vec[byte]{}, fmt.Errorf("File.ReadAllAsVec: %w", err)
	}

	heapBuf := NewVecSized[byte](fileSize, 0x00)
	if _, err := me.Read(heapBuf.HotSlice()); err != nil {
		heapBuf.Free()
		return Vec[byte]{}, fmt.Errorf("File.ReadAllAsVec: %w", err)
	}

	if _, err := me.Seek(0, io.SeekStart); err != nil {
		heapBuf.Free()
		return Vec[byte]{}, fmt.Errorf("File.ReadAllAsVec: %w", err)
	}

	return heapBuf, nil
}

// Implements [io.ByteReader].
//
// Calls [HFILE.ReadFile].
func (me *File) ReadByte() (byte, error) {
	var buf [1]byte
	if _, err := me.Read(buf[:]); err != nil {
		return 0, fmt.Errorf("File.ReadByte: %w", err)
	}
	return buf[0], nil
}

// Truncates or expands the file, according to the new size. Zero will empty the
// file. The internal file pointer will rewind.
//
// Calls [HFILE.SetFilePointerEx] and [HFILE.SetEndOfFile].
//
// For some reason, sometimes the resized file ends up with more space than
// requested. So be careful with this function.
//
// Panics if numBytes is negative.
func (me *File) Resize(numBytes int) error {
	utl.PanicNeg(numBytes)

	// Simply go beyond file limits if needed.
	if _, err := me.Seek(int64(numBytes), io.SeekStart); err != nil {
		return fmt.Errorf("File.Resize: %w", err)
	}

	if err := me.hFile.SetEndOfFile(); err != nil {
		return fmt.Errorf("File.Resize HFILE.SetEndOfFile: %w", err)
	}

	// Rewind pointer.
	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("File.Resize: %w", err)
	}

	return nil
}

// Implements [io.Seeker].
//
// Moves the internal pointer with [HFILE.SetFilePointerEx].
func (me *File) Seek(offset int64, whence int) (int64, error) {
	var moveMethod co.FILE_FROM
	switch whence {
	case io.SeekCurrent:
		moveMethod = co.FILE_FROM_CURRENT
	case io.SeekStart:
		moveMethod = co.FILE_FROM_BEGIN
	case io.SeekEnd:
		moveMethod = co.FILE_FROM_END
	}

	newOff, err := me.hFile.SetFilePointerEx(int(offset), moveMethod)
	if err != nil {
		return 0, fmt.Errorf("File.Seek HFILE.SetFilePointerEx: %w", err)
	}
	return int64(newOff), nil
}

// Retrieves the file size with [HFILE.GetFileSizeEx]. This value is not cached.
func (me *File) Size() (int, error) {
	sz, err := me.hFile.GetFileSizeEx()
	if err != nil {
		return 0, fmt.Errorf("File.Size HFILE.GetFileSizeEx: %w", err)
	}
	return sz, nil
}

// Sets the file size to zero, deleting all its contents.
func (me *File) Truncate() error {
	if err := me.Resize(0); err != nil {
		return fmt.Errorf("File.Truncate: %w", err)
	}
	return nil
}

// Implements [io.Writer].
//
// Calls [HFILE.WriteFile] to write a slice at current internal pointer
// position.
func (me *File) Write(p []byte) (int, error) {
	written, err := me.hFile.WriteFile(p, nil)
	if err != nil {
		return 0, fmt.Errorf("File.Write HFILE.WriteFile: %w", err)
	}
	return int(written), nil
}

// Implements [io.ByteWriter].
//
// Calls [HFILE.WriteFile] to write a byte at current internal pointer position.
func (me *File) WriteByte(c byte) error {
	if _, err := me.Write([]byte{c}); err != nil {
		return fmt.Errorf("File.WriteByte: %w", err)
	}
	return nil
}

// Implements [io.StringWriter].
//
// Calls [HFILE.WriteFile] to write a string at current internal pointer
// position.
func (me *File) WriteString(s string) (int, error) {
	serialized := []byte(s)
	written, err := me.Write(serialized)
	if err != nil {
		return 0, fmt.Errorf("File.WriteString: %w", err)
	}
	return written, nil
}

// High-level abstraction to [HFILEMAP], providing several operations.
//
// Note that memory-mapped files may present issues in x86 architectures; if so,
// just use the ordinary [File].
//
// If you simply need to read or write the contents at once, consider using the
// simpler [FileReadAll] and [FileWriteAll] functions.
//
// Created with [FileMapOpen].
//
// Example:
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
type FileMap struct {
	file *File
	hMap HFILEMAP
	pMem HFILEMAPVIEW
	sz   int
}

// Constructs a new [FileMap] by opening the file and mapping it into memory
// with [HFILE.CreateFileMapping].
//
// Note that memory-mapped files may present issues in x86 architectures; if so,
// just call the ordinary [FileOpen] to work with a non-memory-mapped [File].
//
// ⚠️ You must defer [FileMap.Close].
//
// Example:
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
func FileMapOpen(filePath string, desiredAccess co.FOPEN) (*FileMap, error) {
	file, err := FileOpen(filePath, desiredAccess)
	if err != nil {
		return nil, fmt.Errorf("FileMapOpen: %w", err)
	}

	// Map into memory.
	pageFlags := co.PAGE_READONLY
	if desiredAccess != co.FOPEN_READ_EXISTING {
		pageFlags = co.PAGE_READWRITE
	}

	hMap, err := file.Hfile().CreateFileMapping(nil, pageFlags, co.SEC_NONE, 0, "")
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("FileMapOpen HFILE.CreateFileMapping: %w", err)
	}

	// Get pointer to data block.
	mapFlags := co.FILE_MAP_READ
	if desiredAccess != co.FOPEN_READ_EXISTING {
		mapFlags = co.FILE_MAP_WRITE
	}

	pMem, err := hMap.MapViewOfFile(mapFlags, 0, 0)
	if err != nil {
		hMap.CloseHandle()
		file.Close()
		return nil, fmt.Errorf("FileMapOpen HFILEMAP.MapViewOfFile: %w", err)
	}

	// Cache file size.
	sz, err := file.Size()
	if err != nil {
		pMem.UnmapViewOfFile()
		hMap.CloseHandle()
		file.Close()
		return nil, fmt.Errorf("FileMapOpen: %w", err)
	}

	return &FileMap{file, hMap, pMem, sz}, nil
}

// Unmaps and releases the file resource.
func (me *FileMap) Close() error {
	var errRet error
	if me.pMem != 0 {
		err := me.pMem.UnmapViewOfFile()
		me.pMem = 0
		if err != nil {
			errRet = fmt.Errorf("FileMap.Close HFILEMAPVIEW.UnmapViewOfFile: %w", err)
		}
	}

	if me.hMap != 0 {
		err := me.hMap.CloseHandle()
		me.hMap = 0
		if err != nil && errRet == nil { // only report if pMem.UnmapViewOfFile() succeeded
			errRet = fmt.Errorf("FileMap.Close HFILEMAP.CloseHandle: %w", err)
		}
	}

	err := me.file.Close()
	me.sz = 0
	if err != nil && errRet == nil { // only report if hMap.CloseHandle() succeeded
		errRet = fmt.Errorf("FileMap.Close: %w", err)
	}

	return errRet
}

// Returns a slice to the memory-mapped bytes.
//
// The [FileMap] object must remain open while the slice is being used.
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
// Example:
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
//
// Panics if offset or length is negative.
func (me *FileMap) ReadChunkAsSlice(offset, length int) []byte {
	utl.PanicNeg(offset, length)
	hotSlice := me.HotSlice()
	buf := make([]byte, length)
	copy(buf, hotSlice[offset:offset+length])
	return buf
}

// Returns a [Vec] with a copy of the data, start with offset, and with the
// given length.
//
// Panics if offset or length is negative.
//
// ⚠️ You must defer [Vec.Free] on the returned Vec.
//
// Example:
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
//
//	data := f.ReadChunkAsVec(0, 30)
//	defer data.Free()
//
//	txt := string(data.HotSlice())
//	println(txt)
func (me *FileMap) ReadChunkAsVec(offset, length int) Vec[byte] {
	utl.PanicNeg(offset, length)
	hotSlice := me.HotSlice()
	heapBuf := NewVec[byte]()
	heapBuf.Append(hotSlice[offset : offset+length]...)
	return heapBuf
}

// Retrieves the file size. This value is cached.
func (me *FileMap) Size() int {
	return me.sz
}
