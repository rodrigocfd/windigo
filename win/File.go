package win

import (
	"io"

	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to HFILE, providing several operations.
//
// Note that it implements several standard io interfaces, being interchangeable
// with functions that accept them.
//
// Created with FileOpen().
type File interface {
	io.ByteReader
	io.ByteWriter
	io.Closer
	io.Reader
	io.Seeker
	io.StringWriter
	io.Writer

	Hfile() HFILE              // Returns the underlying HFILE.
	ReadAll() ([]byte, error)  // Rewinds the internal file pointer and reads all contents at once, then rewinds the pointer again.
	Resize(numBytes int) error // Truncates or expands the file, according to the new size. Zero will empty the file. The internal file pointer will rewind.
	Size() int                 // Retrieves the file size. This value is not cached.
}

//------------------------------------------------------------------------------

type _File struct {
	hFile HFILE
}

// Opens a file, returning a new high-level File object.
//
// ⚠️ You must defer File.Close().
func FileOpen(filePath string, desiredAccess co.FILE_OPEN) (File, error) {
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

	return &_File{hFile: hFile}, nil
}

// Implements io.ByteReader.
func (me *_File) ReadByte() (byte, error) {
	var buf [1]byte
	_, err := me.Read(buf[:])
	return buf[0], err
}

// Implements io.ByteWriter.
func (me *_File) WriteByte(c byte) error {
	_, err := me.Write([]byte{c})
	return err
}

// Implements io.Closer.
func (me *_File) Close() error {
	var e error
	if me.hFile != 0 {
		e = me.hFile.CloseHandle()
		me.hFile = 0
	}
	return e
}

// Implements io.Reader.
func (me *_File) Read(p []byte) (int, error) {
	numRead, err := me.hFile.ReadFile(p, uint32(len(p)))
	if err != nil {
		return 0, err
	}

	if numRead < len(p) { // surely there's no more to read
		return numRead, io.EOF
	} else if numRead == 0 { // EOF found
		return 0, io.EOF
	} else {
		return numRead, nil
	}
}

// Implements io.Seeker.
func (me *_File) Seek(offset int64, whence int) (int64, error) {
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

// Implements io.StringWriter.
func (me *_File) WriteString(s string) (int, error) {
	serialized := []byte(s)
	return me.Write(serialized)
}

// Implements io.Writer.
func (me *_File) Write(p []byte) (int, error) {
	written, err := me.hFile.WriteFile(p)
	return written, err
}

func (me *_File) Hfile() HFILE {
	return me.hFile
}

func (me *_File) ReadAll() ([]byte, error) {
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

func (me *_File) Resize(numBytes int) error {
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

func (me *_File) Size() int {
	sz, err := me.hFile.GetFileSizeEx()
	if err != nil {
		panic(err)
	}
	return int(sz)
}
