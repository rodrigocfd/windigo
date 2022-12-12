//go:build windows

package win

import (
	"io"

	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to HFILE, providing several operations.
//
// Implements the following standard io interfaces:
//
//		io.ByteReader
//		io.ByteWriter
//		io.Closer
//		io.Reader
//		io.Seeker
//		io.StringWriter
//		io.Writer
//
// Created with FileOpen().
type File struct {
	hFile HFILE
}

// Opens a file, returning a new high-level File object.
//
// ⚠️ You must defer File.Close().
func FileOpen(filePath string, desiredAccess co.FILE_OPEN) (*File, error) {
	var access co.GENERIC
	var share co.FILE_SHARE
	var disposition co.DISPOSITION

	switch desiredAccess {
	case co.FILE_OPEN_READ_EXISTING:
		access = co.GENERIC_READ
		share = co.FILE_SHARE_READ
		disposition = co.DISPOSITION_OPEN_EXISTING
	case co.FILE_OPEN_RW_EXISTING:
		access = co.GENERIC_READ | co.GENERIC_WRITE
		share = co.FILE_SHARE_NONE
		disposition = co.DISPOSITION_OPEN_EXISTING
	case co.FILE_OPEN_RW_OPEN_OR_CREATE:
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

	return &File{hFile: hFile}, nil
}

// Implements io.Closer.
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

// Implements io.Reader.
func (me *File) Read(p []byte) (n int, err error) {
	numRead, err := me.hFile.ReadFile(p)
	if err != nil {
		return 0, err
	}

	if numRead < uint32(len(p)) { // surely there's no more to read
		return int(numRead), io.EOF
	} else if numRead == 0 { // EOF found
		return 0, io.EOF
	} else {
		return int(numRead), nil
	}
}

// Rewinds the internal file pointer and reads all contents at once, then
// rewinds the pointer again.
func (me *File) ReadAll() ([]byte, error) {
	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	fileSize := me.Size()
	buf := make([]byte, fileSize)
	if _, err := me.Read(buf); err != nil {
		return nil, err
	}

	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	return buf, nil
}

// Implements io.ByteReader.
func (me *File) ReadByte() (byte, error) {
	var buf [1]byte
	_, err := me.Read(buf[:])
	return buf[0], err
}

// Truncates or expands the file, according to the new size. Zero will empty the
// file. The internal file pointer will rewind.
func (me *File) Resize(numBytes int) error {
	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return err
	}

	// Simply go beyond file limits.
	if _, err := me.Seek(int64(numBytes), io.SeekStart); err != nil {
		return err
	}

	if err := me.hFile.SetEndOfFile(); err != nil {
		return err
	}

	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return err
	}

	return nil
}

// Implements io.Seeker.
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

	newOff, err := me.hFile.SetFilePointerEx(offset, moveMethod)
	return int64(newOff), err
}

// Retrieves the file size. This value is not cached.
func (me *File) Size() int {
	sz, err := me.hFile.GetFileSizeEx()
	if err != nil {
		panic(err)
	}
	return int(sz)
}

// Implements io.Writer.
func (me *File) Write(p []byte) (n int, err error) {
	written, err := me.hFile.WriteFile(p)
	return int(written), err
}

// Implements io.ByteWriter.
func (me *File) WriteByte(c byte) error {
	_, err := me.Write([]byte{c})
	return err
}

// Implements io.StringWriter.
func (me *File) WriteString(s string) (int, error) {
	serialized := []byte(s)
	return me.Write(serialized)
}
