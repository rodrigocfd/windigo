/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"wingows/co"
	"wingows/win"
)

type _PathUtilT struct{}

// File and folder path utilities.
var PathUtil _PathUtilT

// Returns all the file names that match a pattern like "C:\\foo\\*.txt".
func (_PathUtilT) ListFilesInFolder(
	pathAndPattern string) ([]string, *win.WinError) {

	retFiles := make([]string, 0)
	dirPath := filepath.Dir(pathAndPattern) // path without file name

	wfd := win.WIN32_FIND_DATA{}
	hFind, found, err := win.FindFirstFile(pathAndPattern, &wfd)
	if err != nil {
		return nil, err
	}
	defer hFind.FindClose()

	for found {
		fileNameFound := syscall.UTF16ToString(wfd.CFileName[:])
		if fileNameFound != ".." {
			retFiles = append(retFiles, dirPath+"\\"+fileNameFound)
		}

		found, err = hFind.FindNextFile(&wfd)
		if err != nil {
			return nil, err
		}
	}

	sort.Slice(retFiles, func(i, j int) bool { // case insensitive
		return strings.ToUpper(retFiles[i]) < strings.ToUpper(retFiles[j])
	})
	return retFiles, nil // search finished successfully
}

// Tells if a given file or folder exists.
func (_PathUtilT) PathExists(path string) bool {
	attr, _ := win.GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID
}

// Tells if a given path is a folder, and not a file.
func (_PathUtilT) PathIsFolder(path string) bool {
	attr, _ := win.GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_DIRECTORY) != 0
}

// Tells if the given file or folder is hidden.
func (_PathUtilT) PathIsHidden(path string) bool {
	attr, _ := win.GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_HIDDEN) != 0
}
