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

func (me *File) OpenExistingForRead(path string) *File {
	return me.rawOpen(path, co.GENERIC_READ,
		co.FILE_SHARE_READ, co.FILE_DISPO_OPEN_EXISTING)
}

func (me *File) OpenExistingForReadWrite(path string) *File {
	return me.rawOpen(path, co.GENERIC_READ|co.GENERIC_WRITE,
		co.FILE_SHARE_NONE, co.FILE_DISPO_OPEN_EXISTING)
}

func (me *File) OpenOrCreate(path string) *File {
	return me.rawOpen(path, co.GENERIC_READ|co.GENERIC_WRITE,
		co.FILE_SHARE_NONE, co.FILE_DISPO_OPEN_ALWAYS)
}

// Rewinds the file pointer and reads all the raw file contents.
func (me *File) ReadAll() []byte {
	me.Rewind()
	sz := me.Size()
	buf := make([]byte, sz)
	me.hFile.ReadFile(buf, uint32(sz)) // will truncate if actual size overflows uint32
	me.Rewind()
	return buf
}

// Rewinds the file pointer back to the first byte of file.
func (me *File) Rewind() *File {
	me.hFile.SetFilePointer(0, co.FILE_SETPTR_BEGIN)
	return me
}

// Truncates or expands the file, according to the new size.
// Zero will empty the file.
func (me *File) SetSize(numBytes uint64) *File {
	me.hFile.SetFilePointerEx(int64(numBytes), co.FILE_SETPTR_BEGIN) // simply go beyond
	me.hFile.SetEndOfFile()
	me.Rewind()
	return me
}

// Retrieves the files size. This value is not cached.
func (me *File) Size() uint64 {
	return uint64(me.hFile.GetFileSizeEx()) // no reason to return an unsigned
}

// Replaces all file contents, possibly resizing the file.
func (me *File) EraseAndWrite(data []byte) *File {
	me.SetSize(uint64(len(data)))
	me.hFile.WriteFile(data)
	me.Rewind()
	return me
}

func (me *File) rawOpen(path string, desiredAccess co.GENERIC,
	shareMode co.FILE_SHARE, creationDisposition co.FILE_DISPO) *File {

	me.Close()
	me.hFile = win.CreateFile(path, desiredAccess, shareMode, nil,
		creationDisposition, co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE,
		co.SECURITY_NONE, 0)
	return me
}
