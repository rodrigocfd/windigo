//go:build windows

package win

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rodrigocfd/windigo/win/co"
)

// Returns all file and folder names on a given directory, filtered to the given
// pattern, by calling [FindFirstFile], [HFIND.FindNextFile] and
// [HFIND.FindClose].
//
// This function is not recursive; to also search nested directories, use
// [EnumFilesDeep].
//
// # Example
//
//	files, _ := win.EnumFiles("C:\\Temp\\*.txt")
//	for _, file := files {
//		println(file)
//	}
func EnumFiles(pathAndPattern string) ([]string, error) {
	var wfd WIN32_FIND_DATA
	hFind, found, err := FindFirstFile(pathAndPattern, &wfd)
	if err != nil {
		return nil, err
	} else if !found {
		return []string{}, nil // empty, not an error
	}
	defer hFind.FindClose()

	dirPath := PathGetPath(pathAndPattern) // path without file name
	files := make([]string, 0)

	for found {
		fileNameFound := wfd.CFileName()
		if fileNameFound != ".." && fileNameFound != "." {
			files = append(files, dirPath+"\\"+fileNameFound)
		}

		if found, err = hFind.FindNextFile(&wfd); err != nil {
			return nil, err
		}
	}

	return files, nil
}

// Returns all files recursively on all folders, by calling [FindFirstFile],
// [HFIND.FindNextFile] and [HFIND.FindClose].
//
// # Example
//
//	files, _ := win.EnumFilesDeep("C:\\Temp")
//	for _, file := files {
//		println(file)
//	}
//
// [FindFirstFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
func EnumFilesDeep(path string) ([]string, error) {
	path, _ = strings.CutSuffix(path, "\\")

	foundFiles, err := EnumFiles(path + "\\*")
	if err != nil {
		return nil, err
	}

	files := make([]string, 0, len(foundFiles))

	for _, file := range foundFiles {
		if !PathIsFolder(file) {
			files = append(files, file)
		} else {
			nestedFiles, err := EnumFilesDeep(file)
			if err != nil {
				return nil, err
			}
			files = append(files, nestedFiles...)
		}
	}

	return files, nil
}

// Returns true if the given file or folder exists. Calls [GetFileAttributes].
//
// Panics on error.
func PathExists(path string) bool {
	attr, err := GetFileAttributes(path)
	if err != nil {
		panic(err)
	}
	return attr != co.FILE_ATTRIBUTE_INVALID
}

// Retrieves the file name of the path.
func PathGetFileName(path string) string {
	if slashIdx := strings.LastIndex(path, "\\"); slashIdx == -1 {
		return path // path contains just the file name
	} else {
		return path[slashIdx+1:]
	}
}

// Retrieves the path without the file name itself, and without trailing
// backslash.
func PathGetPath(path string) string {
	if slashIdx := strings.LastIndex(path, "\\"); slashIdx == -1 {
		return "" // path contains just the file name
	} else {
		return path[0:slashIdx]
	}
}

// Returns whether the path ends with at least one of the given extensions.
//
// # Example
//
//	docPath := "C:\\Temp\\foo.txt"
//	isDocument := win.PathHasExtension(docPath, "txt", "doc")
func PathHasExtension(path string, extensions ...string) bool {
	pathUpper := strings.ToUpper(path)
	for _, extension := range extensions {
		if strings.HasSuffix(pathUpper, strings.ToUpper(extension)) {
			return true
		}
	}
	return false
}

// Returns true if the given path is a folder, and not a file. Calls
// [GetFileAttributes].
//
// Panics on error.
func PathIsFolder(path string) bool {
	attr, err := GetFileAttributes(path)
	if err != nil {
		panic(err)
	}
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_DIRECTORY) != 0
}

// Returns true if the given file or folder is hidden. Calls
// [GetFileAttributes].
//
// Panics on error.
func PathIsHidden(path string) bool {
	attr, err := GetFileAttributes(path)
	if err != nil {
		panic(err)
	}
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_HIDDEN) != 0
}

// Sorts the paths alphabetically, case insensitive, in-place.
func PathSort(paths []string) {
	sort.Slice(paths, func(a, b int) bool {
		return strings.ToUpper(paths[a]) < strings.ToUpper(paths[b])
	})
}

// Replaces the current extension by the new one.
//
// Panics if the path doesn't have a file name.
func PathSwapExtension(path, newExtension string) string {
	if !strings.HasPrefix(newExtension, ".") {
		newExtension = "." + newExtension // must start with a dot
	}

	if strings.HasSuffix(path, "\\") {
		panic(fmt.Sprintf("Path doesn't have a file name: %s", path))
	}

	if idxDot := strings.LastIndex(path, "."); idxDot == -1 {
		return path + newExtension
	} else {
		return path[:idxDot] + newExtension
	}
}
