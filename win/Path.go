package win

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/rodrigocfd/windigo/win/co"
)

type _PathT struct{}

// File and folder path functions.
var Path _PathT

// Returns all the file names that match a pattern like "C:\\foo\\*.txt".
func (_PathT) ListFilesInFolder(pathAndPattern string) ([]string, error) {
	retFiles := make([]string, 0, 5)        // arbitrary
	dirPath := filepath.Dir(pathAndPattern) // path without file name

	wfd := WIN32_FIND_DATA{}
	hFind, found, lerr := FindFirstFile(pathAndPattern, &wfd)
	if lerr != nil {
		return nil, lerr
	}
	defer hFind.FindClose()

	for found {
		fileNameFound := Str.FromUint16Slice(wfd.CFileName[:])
		if fileNameFound != ".." {
			retFiles = append(retFiles, dirPath+"\\"+fileNameFound)
		}

		if found, lerr = hFind.FindNextFile(&wfd); lerr != nil {
			return nil, lerr
		}
	}

	sort.Slice(retFiles, func(i, j int) bool { // case insensitive
		return strings.ToUpper(retFiles[i]) < strings.ToUpper(retFiles[j])
	})
	return retFiles, nil // search finished successfully
}

// Tells if a given file or folder exists.
func (_PathT) Exists(path string) bool {
	attr, _ := GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID
}

// Tells if a given path is a folder, and not a file.
func (_PathT) IsFolder(path string) bool {
	attr, _ := GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_DIRECTORY) != 0
}

// Tells if the given file or folder is hidden.
func (_PathT) IsHidden(path string) bool {
	attr, _ := GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_HIDDEN) != 0
}
