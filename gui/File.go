/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"wingows/co"
	"wingows/win"
)

// Manages a file resource.
type File struct {
	hFile win.HFILE
}

func (me *File) Close() {
	if me.hFile != 0 {
		me.hFile.CloseHandle()
		me.hFile = 0
	}
}

// Replaces all file contents, possibly resizing the file.
func (me *File) EraseAndWrite(data []byte) *win.WinError {
	if err := me.SetSize(uint64(len(data))); err != nil {
		return err
	}
	if err := me.hFile.WriteFile(data); err != nil {
		return err
	}
	return me.Rewind()
}

func (me *File) OpenExistingForRead(path string) *win.WinError {
	return me.rawOpen(path, co.GENERIC_READ,
		co.FILE_SHARE_READ, co.FILE_DISPO_OPEN_EXISTING)
}

func (me *File) OpenExistingForReadWrite(path string) *win.WinError {
	return me.rawOpen(path, co.GENERIC_READ|co.GENERIC_WRITE,
		co.FILE_SHARE_NONE, co.FILE_DISPO_OPEN_EXISTING)
}

func (me *File) OpenOrCreate(path string) *win.WinError {
	return me.rawOpen(path, co.GENERIC_READ|co.GENERIC_WRITE,
		co.FILE_SHARE_NONE, co.FILE_DISPO_OPEN_ALWAYS)
}

// Rewinds the file pointer and reads all the raw file contents.
func (me *File) ReadAll() ([]byte, *win.WinError) {
	if err := me.Rewind(); err != nil {
		return nil, err
	}
	fileSize := me.Size()
	buf := make([]byte, fileSize)

	// Read the contents into our allocated buffer.
	// Will truncate if file data overflows uint32.
	if err := me.hFile.ReadFile(buf, uint32(fileSize)); err != nil {
		return nil, err
	}

	if err := me.Rewind(); err != nil {
		return nil, err
	}
	return buf, nil
}

// Rewinds the file pointer back to the first byte of file.
func (me *File) Rewind() *win.WinError {
	return me.hFile.SetFilePointerEx(0, co.FILE_SETPTR_BEGIN)
}

// Truncates or expands the file, according to the new size.
//
// Zero will empty the file.
func (me *File) SetSize(numBytes uint64) *win.WinError {
	// Simply go beyond file limits.
	if err := me.hFile.SetFilePointerEx(int64(numBytes), co.FILE_SETPTR_BEGIN); err != nil {
		return err
	}
	if err := me.hFile.SetEndOfFile(); err != nil {
		return err
	}
	return me.Rewind()
}

// Retrieves the files size. This value is not cached.
func (me *File) Size() uint64 {
	sz, err := me.hFile.GetFileSizeEx()
	if err != nil {
		panic(err.Error())
	}
	return uint64(sz) // no reason to return an unsigned
}

func (me *File) rawOpen(
	path string, desiredAccess co.GENERIC, shareMode co.FILE_SHARE,
	creationDisposition co.FILE_DISPO) *win.WinError {

	me.Close()
	var err *win.WinError
	me.hFile, err = win.CreateFile(path, desiredAccess, shareMode, nil,
		creationDisposition, co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE,
		co.SECURITY_NONE, 0)
	return err
}
