/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"sort"
	"strings"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Any child control with HWND and ID.
type Control interface {
	Window
	Id() int32
}

// Any window with a HWND handle.
type Window interface {
	Hwnd() win.HWND
}

//------------------------------------------------------------------------------

// Enables or disables many controls at once.
func EnableControls(enabled bool, ctrls []Control) {
	for _, ctrl := range ctrls {
		ctrl.Hwnd().EnableWindow(enabled)
	}
}

// Shows the open file system dialog, choice restricted to 1 file.
// Example filtersWithPipe:
// []string{"Text files (*.txt)|*.txt", "All files|*.*"}
func ShowFileOpen(owner Window, filtersWithPipe []string) (bool, string) {
	zFilters := filterToUtf16(filtersWithPipe)
	result := [260]uint16{} // MAX_PATH

	ofn := win.OPENFILENAME{
		HwndOwner:   owner.Hwnd(),
		LpstrFilter: uintptr(unsafe.Pointer(&zFilters[0])),
		LpstrFile:   uintptr(unsafe.Pointer(&result[0])),
		NMaxFile:    uint32(len(result)),
		Flags:       co.OFN_EXPLORER | co.OFN_ENABLESIZING | co.OFN_FILEMUSTEXIST,
	}

	if !ofn.GetOpenFileName() {
		return false, ""
	}
	return true, syscall.UTF16ToString(result[:])
}

// Shows the open file system dialog, user can choose multiple files.
// Example filtersWithPipe:
// []string{"Text files (*.txt)|*.txt", "All files|*.*"}
func ShowFileOpenMany(owner Window, filtersWithPipe []string) (bool, []string) {
	zFilters := filterToUtf16(filtersWithPipe)
	multiBuf := make([]uint16, 65536) // http://www.askjf.com/?q=2179s http://www.askjf.com/?q=2181s

	ofn := win.OPENFILENAME{
		HwndOwner:   owner.Hwnd(),
		LpstrFilter: uintptr(unsafe.Pointer(&zFilters[0])),
		LpstrFile:   uintptr(unsafe.Pointer(&multiBuf[0])),
		NMaxFile:    uint32(len(multiBuf)),
		Flags:       co.OFN_EXPLORER | co.OFN_ENABLESIZING | co.OFN_FILEMUSTEXIST | co.OFN_ALLOWMULTISELECT,
	}

	if !ofn.GetOpenFileName() {
		return false, []string{}
	}

	resultStrs := make([][]uint16, 0)
	beginIdx := 0
	for i := 0; i < len(multiBuf)-1; i++ {
		if multiBuf[i] == 0 { // found end of a string
			resultStrs = append(resultStrs, multiBuf[beginIdx:i+1]) // includes terminating null
			if multiBuf[i+1] == 0 {
				break // double terminating null: end
			}
			beginIdx = i + 1
		}
	}

	// User selected only 1 file, this string is the full path, and that's all.
	if len(resultStrs) == 1 {
		return true, []string{syscall.UTF16ToString(resultStrs[0][:])}
	}

	// User selected 2 or more files.
	// 1st string is the base path, the others are the filenames.
	final := make([]string, 0, len(resultStrs)-1)
	basePath := syscall.UTF16ToString(resultStrs[0]) + "\\"
	for i := 1; i < len(resultStrs); i++ {
		final = append(final, basePath+syscall.UTF16ToString(resultStrs[i]))
	}
	sort.Slice(final, func(i, j int) bool { // case insensitive
		return strings.ToUpper(final[i]) < strings.ToUpper(final[j])
	})
	return true, final
}

// Shows the save file system dialog.
// Default name can be empty.
// Default extension can be empty, or like "txt".
// Example filtersWithPipe:
// []string{"Text files (*.txt)|*.txt", "All files|*.*"}
func ShowFileSave(owner Window, defaultName, defaultExt string,
	filtersWithPipe []string) (bool, string) {

	zFilters := filterToUtf16(filtersWithPipe)
	defExt := win.StrToSlice(defaultExt)

	result := [260]uint16{} // MAX_PATH
	if defaultName != "" {
		copy(result[:], win.StrToSlice(defaultName))
	}

	ofn := win.OPENFILENAME{
		HwndOwner:   owner.Hwnd(),
		LpstrFilter: uintptr(unsafe.Pointer(&zFilters[0])),
		LpstrFile:   uintptr(unsafe.Pointer(&result[0])),
		NMaxFile:    uint32(len(result)),
		Flags:       co.OFN_HIDEREADONLY | co.OFN_OVERWRITEPROMPT,
		// If absent, no default extension is appended.
		// If present, even if empty, default extension is appended; if no default
		// extension, first one of the filter is appended.
		LpstrDefExt: uintptr(unsafe.Pointer(&defExt[0])),
	}

	if !ofn.GetSaveFileName() {
		return false, ""
	}
	return true, syscall.UTF16ToString(result[:])
}

func filterToUtf16(filtersWithPipe []string) []uint16 {
	filters16 := make([][]uint16, 0, len(filtersWithPipe)) // each filter as []uint16 with terminating null
	charCount := 0
	for _, filter := range filtersWithPipe {
		filters16 = append(filters16, win.StrToSlice(filter))
		charCount += len(filter) + 1 // also count terminating null
	}

	finalBuf := make([]uint16, 0, charCount+1) // double terminating null
	for _, filter16 := range filters16 {
		finalBuf = append(finalBuf, filter16...) // concat all filters into one big slice
	}
	finalBuf = append(finalBuf, 0) // double terminating null

	for i := range finalBuf {
		if finalBuf[i] == '|' {
			finalBuf[i] = 0 // replace pipes with nulls
		}
	}
	return finalBuf
}
