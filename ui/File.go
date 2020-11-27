/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"windigo/co"
	"windigo/win"
)

// Used in OpenFile().
//
// Behavior of the file opening.
type FILE_MODE uint8

const (
	FILE_MODE_R_EXISTING        FILE_MODE = iota // Open an existing file for read only.
	FILE_MODE_RW_EXISTING                        // Open an existing file for read and write.
	FILE_MODE_RW_OPEN_OR_CREATE                  // Open a file or create if it doesn't exist, for read and write.
)

//------------------------------------------------------------------------------

// Manages a file resource.
type File struct {
	hFile win.HFILE
}

// Constructor.
//
// You must defer Close().
func OpenFile(path string, behavior FILE_MODE) (*File, error) {
	var desiredAccess co.GENERIC
	var shareMode co.FILE_SHARE
	var creationDisposition co.FILE_DISPO

	switch behavior {
	case FILE_MODE_R_EXISTING:
		desiredAccess = co.GENERIC_READ
		shareMode = co.FILE_SHARE_READ
		creationDisposition = co.FILE_DISPO_OPEN_EXISTING
	case FILE_MODE_RW_EXISTING:
		desiredAccess = co.GENERIC_READ | co.GENERIC_WRITE
		shareMode = co.FILE_SHARE_NONE
		creationDisposition = co.FILE_DISPO_OPEN_EXISTING
	case FILE_MODE_RW_OPEN_OR_CREATE:
		desiredAccess = co.GENERIC_READ | co.GENERIC_WRITE
		shareMode = co.FILE_SHARE_NONE
		creationDisposition = co.FILE_DISPO_OPEN_ALWAYS
	}

	hFile, err := win.CreateFile(path, desiredAccess, shareMode, nil,
		creationDisposition, co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE,
		co.SECURITY_NONE, 0)
	if err != nil {
		return nil, err
	}
	return &File{hFile: hFile}, nil
}

// Calls CloseHandle() to free the file resource.
func (me *File) Close() {
	if me.hFile != 0 {
		me.hFile.CloseHandle()
		me.hFile = 0
	}
}

// Replaces all file contents, possibly resizing the file.
func (me *File) EraseAndWrite(data []byte) error {
	if err := me.SetSize(len(data)); err != nil {
		return err
	}
	if err := me.hFile.WriteFile(data); err != nil {
		return err
	}
	return me.RewindPointerOffset()
}

// Retrieves the current pointer offset.
func (me *File) PointerOffset() int {
	// https://stackoverflow.com/a/17707021/6923555
	off, err := me.hFile.SetFilePointerEx(0, co.FILE_SETPTR_CURRENT)
	if err != nil {
		panic(err)
	}
	return int(off)
}

// Reads file data at the current internal pointer offset, which then advances.
func (me *File) Read(numBytes uint) ([]byte, error) {
	buf := make([]byte, numBytes)
	if err := me.hFile.ReadFile(buf, uint32(numBytes)); err != nil {
		return nil, err
	}
	return buf, nil
}

// Rewinds the file pointer and reads all the raw file contents.
func (me *File) ReadAll() ([]byte, error) {
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

// Rewinds the internal pointer back to the first byte of the file.
func (me *File) RewindPointerOffset() error {
	_, err := me.hFile.SetFilePointerEx(0, co.FILE_SETPTR_BEGIN)
	return err
}

// Truncates or expands the file, according to the new size.
//
// Zero will empty the file.
func (me *File) SetSize(numBytes int) error {
	// Simply go beyond file limits.
	if _, err := me.hFile.SetFilePointerEx(
		int64(numBytes), co.FILE_SETPTR_BEGIN); err != nil {
		return err
	}

	if err := me.hFile.SetEndOfFile(); err != nil {
		return err
	}

	return me.RewindPointerOffset()
}

// Retrieves the files size. This value is not cached.
func (me *File) Size() int {
	sz, err := me.hFile.GetFileSizeEx()
	if err != nil {
		panic(err)
	}
	return int(sz)
}

// Writes the bytes at current internal pointer offset, which then advances.
func (me *File) Write(data []byte) error {
	return me.hFile.WriteFile(data)
}
