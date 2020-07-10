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

type fileUtilT struct{}

// File utilities.
var FileUtil fileUtilT

// Returns all the file names that match a pattern like "C:\\foo\\*.txt".
func (fileUtilT) ListFilesInFolder(pathAndPattern string) []string {
	retFiles := make([]string, 0)
	dirPath := filepath.Dir(pathAndPattern)

	wfd := win.WIN32_FIND_DATA{}
	hFind, found := win.FindFirstFile(pathAndPattern, &wfd)
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

func (fileUtilT) PathExists(path string) bool {
	attr, _ := win.GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID
}

func (fileUtilT) PathIsFolder(path string) bool {
	attr, _ := win.GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_DIRECTORY) != 0
}

func (fileUtilT) PathIsHidden(path string) bool {
	attr, _ := win.GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_HIDDEN) != 0
}
