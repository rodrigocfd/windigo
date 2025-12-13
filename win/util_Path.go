//go:build windows

package win

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rodrigocfd/windigo/co"
)

// Returns a new []string with all files and folders within searchPath.
//
// If fileExtension isn't empty, brings only the files and folders with this
// extension.
//
// Does not search recursively. For a recursive search, use [PathEnumDeep].
//
// Calls:
//   - [FindFirstFile]
//   - [HFIND.FindNextFile]
//   - [HFIND.FindClose]
//
// Example:
//
//	paths := win.PathEnum("C:\\Temp", "")
//	mp3s := win.PathEnum("C:\\Temp", "mp3")
func PathEnum(searchPath, fileExtension string) ([]string, error) {
	if strings.Contains(searchPath, "*") {
		return nil, fmt.Errorf("invalid path: %s", searchPath)
	} else if strings.Contains(fileExtension, "*") {
		return nil, fmt.Errorf("invalid file extension: %s", fileExtension)
	}
	searchPath = strings.TrimSpace(searchPath)
	searchPath = strings.TrimSuffix(searchPath, "\\")
	basePath := searchPath
	searchPath += "\\*"
	fileExtension = strings.TrimSpace(fileExtension)
	if fileExtension != "" {
		searchPath += "." + fileExtension
	}

	var wfd WIN32_FIND_DATA
	hFind, found, err := FindFirstFile(searchPath, &wfd)
	if err != nil {
		return nil, fmt.Errorf("PathEnum FindFirstFile: %w", err)
	} else if !found {
		return []string{}, nil // empty, not an error
	}
	defer hFind.FindClose()

	files := make([]string, 0, 20) // arbitrary
	for found {
		fileNameFound := wfd.CFileName()
		if fileNameFound != ".." && fileNameFound != "." {
			files = append(files, basePath+"\\"+fileNameFound)
		}

		if found, err = hFind.FindNextFile(&wfd); err != nil {
			return nil, fmt.Errorf("PathEnum HFIND.FindNextFile: %w", err)
		}
	}
	PathSort(files)
	return files, nil
}

// Returns a new []string with all files within searchPath.
//
// If fileExtension isn't empty, brings only the files with this extension.
//
// Searches recursively. For a non-recursive search, use [PathEnum].
//
// Calls:
//   - [FindFirstFile]
//   - [HFIND.FindNextFile]
//   - [HFIND.FindClose]
//
// Example:
//
//	paths := win.PathEnumDeep("C:\\Temp", "")
//	mp3s := win.PathEnumDeep("C:\\Temp", "mp3")
func PathEnumDeep(searchPath, fileExtension string) ([]string, error) {
	foundFiles, err := PathEnum(searchPath, "") // if we pass extension, subfolders will be skipped
	if err != nil {
		return nil, fmt.Errorf("PathEnumDeep: %w", err)
	}
	if len(foundFiles) == 0 {
		return []string{}, nil
	}

	files := make([]string, 0, len(foundFiles)+20) // arbitrary
	for _, f := range foundFiles {
		if !PathIsFolder(f) {
			if fileExtension == "" || PathHasExtension(f, fileExtension) { // manual extension filter
				files = append(files, f)
			}
		} else {
			nestedFiles, err := PathEnumDeep(f, fileExtension) // recursively
			if err != nil {
				return nil, err // don't wrap to avoid recursion repetition
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
// Example:
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
