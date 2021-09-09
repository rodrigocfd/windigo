package win

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rodrigocfd/windigo/win/co"
)

type _PathT struct{}

// File and folder path utilities.
var Path _PathT

// Returns whether the path ends with at least one of the given extensions.
func (_PathT) HasExtension(path string, extensions ...string) bool {
	for _, extension := range extensions {
		if strings.HasSuffix(strings.ToUpper(path), strings.ToUpper(extension)) {
			return true
		}
	}
	return false
}

// Returns all the file names that match a pattern like "C:\\foo\\*.txt".
func (_PathT) ListFilesInFolder(pathAndPattern string) ([]string, error) {
	wfd := WIN32_FIND_DATA{}
	hFind, found, err := FindFirstFile(pathAndPattern, &wfd)
	if err != nil {
		return nil, err
	} else if !found {
		return []string{}, nil // empty array, no error
	}
	defer hFind.FindClose()

	retFiles := make([]string, 0, 5)        // arbitrary
	dirPath := Path.GetPath(pathAndPattern) // path without file name

	for found {
		fileNameFound := wfd.CFileName()
		if fileNameFound != ".." {
			retFiles = append(retFiles, dirPath+"\\"+fileNameFound)
		}

		if found, err = hFind.FindNextFile(&wfd); err != nil {
			return nil, err
		}
	}

	sort.Slice(retFiles, func(a, b int) bool { // case insensitive
		return strings.ToUpper(retFiles[a]) < strings.ToUpper(retFiles[b])
	})
	return retFiles, nil // search finished successfully
}

// Tells if a given file or folder exists.
func (_PathT) Exists(path string) bool {
	attr, _ := GetFileAttributes(path)
	return attr != co.FILE_ATTRIBUTE_INVALID
}

// Retrieves the file name of the path.
func (_PathT) GetFileName(path string) string {
	slashIdx := strings.LastIndex(path, "\\")
	return path[slashIdx+1:]
}

// Retrieves the path without the file name itself, and without trailing slash.
func (_PathT) GetPath(path string) string {
	slashIdx := strings.LastIndex(path, "\\")
	return path[0:slashIdx]
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

// Sorts the paths alphabetically, case insensitive.
func (_PathT) Sort(paths []string) {
	sort.Slice(paths, func(a, b int) bool {
		return strings.ToUpper(paths[a]) < strings.ToUpper(paths[b])
	})
}

// Replaces the current extension by the new one, which must start with a dot.
func (_PathT) SwapExtension(path, newExtension string) string {
	if !strings.HasPrefix(newExtension, ".") {
		panic(fmt.Sprintf("New extension must start with a dot: \"%s\"", newExtension))
	}

	idxDot := strings.LastIndex(path, ".")
	if idxDot == -1 {
		panic(fmt.Sprintf("Path has no extension to be swapped: \"%s\"", path))
	}

	return path[:idxDot] + newExtension
}
