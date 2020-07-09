/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"path/filepath"
	"sort"
	"strings"
	"syscall"
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

func (me *File) HFile() win.HFILE {
	return me.hFile
}

// Returns all the file names that match a pattern like "C:\\foo\\*.txt".
func ListFilesInFolder(pathAndPattern string) []string {
	retFiles := make([]string, 0)
	dirPath := filepath.Dir(pathAndPattern)

	wfd := win.WIN32_FIND_DATA{}
	found, hFind := win.FindFirstFile(pathAndPattern, &wfd)
	defer hFind.FindClose()

	for found {
		fileNameFound := syscall.UTF16ToString(wfd.CFileName[:])
		if fileNameFound != ".." {
			retFiles = append(retFiles, dirPath+"\\"+fileNameFound)
		}

		found = hFind.FindNextFile(&wfd)
	}

	sort.Slice(retFiles, func(i, j int) bool { // case insensitive
		return strings.ToUpper(retFiles[i]) < strings.ToUpper(retFiles[j])
	})
	return retFiles
}

func (me *File) OpenExistingForRead(path string) *File {
	me.hFile = win.CreateFile(path, co.GENERIC_READ,
		co.FILE_SHARE_READ, nil, co.FILE_DISPO_OPEN_EXISTING,
		co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE, co.SECURITY_NONE, 0)
	return me
}

func (me *File) OpenExistingForReadWrite(path string) *File {
	me.hFile = win.CreateFile(path, co.GENERIC_READ|co.GENERIC_WRITE,
		co.FILE_SHARE_NONE, nil, co.FILE_DISPO_OPEN_EXISTING,
		co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE, co.SECURITY_NONE, 0)
	return me
}

func (me *File) OpenOrCreate(path string) *File {
	me.hFile = win.CreateFile(path, co.GENERIC_READ|co.GENERIC_WRITE,
		co.FILE_SHARE_NONE, nil, co.FILE_DISPO_OPEN_ALWAYS,
		co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE, co.SECURITY_NONE, 0)
	return me
}

// Rewinds the file pointer and reads all the raw file contents.
func (me *File) ReadAll() []uint8 {
	me.Rewind()
	sz := me.Size()
	buf := make([]uint8, sz)
	me.hFile.ReadFile(buf, sz)
	me.Rewind()
	return buf
}

// Rewinds the file pointer back to the first byte of file.
func (me *File) Rewind() *File {
	me.hFile.SetFilePointer(0, co.FILE_SETPTR_BEGIN)
	return me
}

// Expands or shrinks the file.
func (me *File) SetSize(numBytes uint32) *File {
	me.hFile.SetFilePointer(int32(numBytes), co.FILE_SETPTR_BEGIN)
	me.hFile.SetEndOfFile()
	me.Rewind()
	return me
}

func (me *File) Size() uint32 {
	return me.hFile.GetFileSize()
}

func PathExists(path string) bool {
	return win.GetFileAttributes(path) != co.FILE_ATTRIBUTE_INVALID
}

func PathIsFolder(path string) bool {
	return (win.GetFileAttributes(path) & co.FILE_ATTRIBUTE_DIRECTORY) != 0
}

func PathIsHidden(path string) bool {
	return (win.GetFileAttributes(path) & co.FILE_ATTRIBUTE_HIDDEN) != 0
}

// Replaces all file contents, possibly resizing the file.
func (me *File) EraseAndWrite(data []uint8) *File {
	me.SetSize(uint32(len(data)))
	me.hFile.WriteFile(data)
	me.Rewind()
	return me
}
