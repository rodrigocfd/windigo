package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to HFILE, providing several operations.
//
// Created with OpenFile().
type File interface {
	Close()                            // Releases the file resource.
	CurrentPointerOffset() int         // Retrieves the current file pointer offset.
	EraseAndWrite(data []byte) error   // Replaces all file contents, possibly resizing the file.
	Read(numBytes int) ([]byte, error) // Reads data from file at current pointer offset. The pointer will then advance.
	ReadAll() ([]byte, error)          // Rewinds the file pointer and reads all the raw file contents.
	Resize(numBytes int) error         // Truncates or expands the file, according to the new size. Zero will empty the file.
	RewindPointerOffset() error        // Rewinds the internal pointer back to the beginning of the file.
	Size() int                         // Retrieves the files size. This value is not cached.
	Write(data []byte) error           // Writes the bytes at current internal pointer offset, which then advances.
}

//------------------------------------------------------------------------------

type _File struct {
	hFile HFILE
}

// Opens a file, returning a new high-level File object.
//
// ⚠️ You must defer Close().
func OpenFile(filePath string, behavior co.OPEN_FILE) (File, error) {
	me := &_File{}
	if err := me.openFile(filePath, behavior); err != nil {
		return nil, err
	}
	return me, nil
}

func (me *_File) openFile(filePath string, behavior co.OPEN_FILE) error {
	var desiredAccess co.GENERIC
	var shareMode co.FILE_SHARE
	var creationDisposition co.DISPOSITION

	switch behavior {
	case co.OPEN_FILE_READ_EXISTING:
		desiredAccess = co.GENERIC_READ
		shareMode = co.FILE_SHARE_READ
		creationDisposition = co.DISPOSITION_OPEN_EXISTING
	case co.OPEN_FILE_RW_EXISTING:
		desiredAccess = co.GENERIC_READ | co.GENERIC_WRITE
		shareMode = co.FILE_SHARE_NONE
		creationDisposition = co.DISPOSITION_OPEN_EXISTING
	case co.OPEN_FILE_RW_OPEN_OR_CREATE:
		desiredAccess = co.GENERIC_READ | co.GENERIC_WRITE
		shareMode = co.FILE_SHARE_NONE
		creationDisposition = co.DISPOSITION_OPEN_ALWAYS
	}

	hFile, err := CreateFile(filePath, desiredAccess, shareMode, nil,
		creationDisposition, co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE,
		co.SECURITY_NONE, 0)
	if err != nil {
		return err
	}

	me.hFile = hFile
	return nil
}

func (me *_File) Close() {
	if me.hFile != 0 {
		me.hFile.CloseHandle()
		me.hFile = 0
	}
}

func (me *_File) CurrentPointerOffset() int {
	off, err := me.hFile.SetFilePointerEx(0, co.FILE_FROM_CURRENT) // https://stackoverflow.com/a/17707021/6923555
	if err != nil {
		panic(err)
	}
	return int(off)
}

func (me *_File) EraseAndWrite(data []byte) error {
	if err := me.Resize(len(data)); err != nil {
		return err
	}
	if err := me.hFile.WriteFile(data); err != nil {
		return err
	}
	return me.RewindPointerOffset()
}

func (me *_File) Read(numBytes int) ([]byte, error) {
	buf := make([]byte, numBytes)
	if err := me.hFile.ReadFile(buf, uint32(numBytes)); err != nil {
		return nil, err
	}
	return buf, nil
}

func (me *_File) ReadAll() ([]byte, error) {
	if err := me.RewindPointerOffset(); err != nil {
		return nil, err
	}
	fileSize := me.Size()
	buf := make([]byte, fileSize)

	// Read the contents into our allocated buffer.
	// Will truncate if file data overflows uint32.
	if err := me.hFile.ReadFile(buf, uint32(fileSize)); err != nil {
		return nil, err
	}

	if err := me.RewindPointerOffset(); err != nil {
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

	return me.RewindPointerOffset()
}

func (me *_File) RewindPointerOffset() error {
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
