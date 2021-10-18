package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to HFILE, providing several operations.
//
// Created with OpenFile().
type File interface {
	// Releases the file resource.
	Close()

	// Replaces all file contents, possibly resizing the file.
	EraseAndWrite(data []byte) error

	// Returns the underlying HFILE.
	Hfile() HFILE

	// Retrieves the current file pointer offset.
	PointerOffset() int

	// Reads data from file at current pointer offset. The pointer will then
	// advance.
	Read(numBytes int) ([]byte, error)

	// Rewinds the file pointer and reads all the raw file contents. Then
	// rewinds the pointer again.
	ReadAll() ([]byte, error)

	// Truncates or expands the file, according to the new size. Zero will empty
	// the file.
	Resize(numBytes int) error

	// Rewinds the internal pointer back to the beginning of the file.
	RewindPointer() error

	// Retrieves the files size. This value is not cached.
	Size() int

	// Writes the bytes at current internal pointer offset, which then advances.
	Write(data []byte) error
}

//------------------------------------------------------------------------------

type _File struct {
	hFile HFILE
}

// Opens a file, returning a new high-level File object.
//
// ⚠️ You must defer File.Close().
func OpenFile(filePath string, desiredAccess co.OPEN_FILE) (File, error) {
	var access co.GENERIC
	var share co.FILE_SHARE
	var disposition co.DISPOSITION

	switch desiredAccess {
	case co.OPEN_FILE_READ_EXISTING:
		access = co.GENERIC_READ
		share = co.FILE_SHARE_READ
		disposition = co.DISPOSITION_OPEN_EXISTING
	case co.OPEN_FILE_RW_EXISTING:
		access = co.GENERIC_READ | co.GENERIC_WRITE
		share = co.FILE_SHARE_NONE
		disposition = co.DISPOSITION_OPEN_EXISTING
	case co.OPEN_FILE_RW_OPEN_OR_CREATE:
		access = co.GENERIC_READ | co.GENERIC_WRITE
		share = co.FILE_SHARE_NONE
		disposition = co.DISPOSITION_OPEN_ALWAYS
	}

	hFile, err := CreateFile(filePath, access, share, nil,
		disposition, co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE,
		co.SECURITY_NONE, 0)
	if err != nil {
		return nil, err
	}

	return &_File{hFile: hFile}, nil
}

func (me *_File) Close() {
	if me.hFile != 0 {
		me.hFile.CloseHandle()
		me.hFile = 0
	}
}

func (me *_File) EraseAndWrite(data []byte) error {
	if err := me.Resize(len(data)); err != nil {
		return err
	}
	if err := me.hFile.WriteFile(data); err != nil {
		return err
	}
	return me.RewindPointer()
}

func (me *_File) Hfile() HFILE {
	return me.hFile
}

func (me *_File) PointerOffset() int {
	off, err := me.hFile.SetFilePointerEx(0, co.FILE_FROM_CURRENT) // https://stackoverflow.com/a/17707021/6923555
	if err != nil {
		panic(err)
	}
	return int(off)
}

func (me *_File) Read(numBytes int) ([]byte, error) {
	buf := make([]byte, numBytes)
	if err := me.hFile.ReadFile(buf, uint32(numBytes)); err != nil {
		return nil, err
	}
	return buf, nil
}

func (me *_File) ReadAll() ([]byte, error) {
	if err := me.RewindPointer(); err != nil {
		return nil, err
	}
	fileSize := me.Size()
	buf := make([]byte, fileSize)

	// Read the contents into our allocated buffer.
	// Will truncate if file data overflows uint32.
	if err := me.hFile.ReadFile(buf, uint32(fileSize)); err != nil {
		return nil, err
	}

	if err := me.RewindPointer(); err != nil {
		return nil, err
	}
	return buf, nil
}

func (me *_File) Resize(numBytes int) error {
	// Simply go beyond file limits.
	if _, err := me.hFile.SetFilePointerEx(
		int64(numBytes), co.FILE_FROM_BEGIN); err != nil {
		return err
	}

	if err := me.hFile.SetEndOfFile(); err != nil {
		return err
	}

	return me.RewindPointer()
}

func (me *_File) RewindPointer() error {
	_, err := me.hFile.SetFilePointerEx(0, co.FILE_FROM_BEGIN)
	return err
}

func (me *_File) Size() int {
	sz, err := me.hFile.GetFileSizeEx()
	if err != nil {
		panic(err)
	}
	return int(sz)
}

func (me *_File) Write(data []byte) error {
	return me.hFile.WriteFile(data)
}
