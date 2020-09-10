/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"sort"
	"strings"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

type _SysDlgUtilT struct{}

// System dialogs utilities.
var SysDlgUtil _SysDlgUtilT

// Shows the open file system dialog, choice restricted to 1 file.
//
// Example of filtersWithPipe:
//
// []string{"Text files (*.txt)|*.txt", "All files|*.*"}
func (_SysDlgUtilT) FileOpen(
	owner Window, filtersWithPipe []string) (string, bool) {

	zFilters := filterToUtf16(filtersWithPipe)
	result := [260]uint16{} // MAX_PATH

	ofn := win.OPENFILENAME{
		HwndOwner:   owner.Hwnd(),
		LpstrFilter: &zFilters[0],
		LpstrFile:   &result[0],
		NMaxFile:    uint32(len(result)),
		Flags:       co.OFN_EXPLORER | co.OFN_ENABLESIZING | co.OFN_FILEMUSTEXIST,
	}

	if !win.GetOpenFileName(&ofn) {
		return "", false
	}
	return syscall.UTF16ToString(result[:]), true
}

// Shows the open file system dialog, user can choose multiple files.
//
// Example of filtersWithPipe:
//
// []string{"Text files (*.txt)|*.txt", "All files|*.*"}
func (_SysDlgUtilT) FileOpenMany(
	owner Window, filtersWithPipe []string) ([]string, bool) {

	zFilters := filterToUtf16(filtersWithPipe)
	multiBuf := make([]uint16, 65536) // http://www.askjf.com/?q=2179s http://www.askjf.com/?q=2181s

	ofn := win.OPENFILENAME{
		HwndOwner:   owner.Hwnd(),
		LpstrFilter: &zFilters[0],
		LpstrFile:   &multiBuf[0],
		NMaxFile:    uint32(len(multiBuf)),
		Flags:       co.OFN_EXPLORER | co.OFN_ENABLESIZING | co.OFN_FILEMUSTEXIST | co.OFN_ALLOWMULTISELECT,
	}

	if !win.GetOpenFileName(&ofn) {
		return nil, false
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
		return []string{syscall.UTF16ToString(resultStrs[0][:])}, true
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
	return final, true
}

// Shows the save file system dialog.
//
// Default name can be empty, default extension can be empty, or like "txt".
//
// Example of filtersWithPipe:
//
// []string{"Text files (*.txt)|*.txt", "All files|*.*"}
func (_SysDlgUtilT) FileSave(
	owner Window, defaultName, defaultExt string,
	filtersWithPipe []string) (string, bool) {

	zFilters := filterToUtf16(filtersWithPipe)
	defExt := win.StrToSlice(defaultExt)

	result := [260]uint16{} // MAX_PATH
	if defaultName != "" {
		copy(result[:], win.StrToSlice(defaultName))
	}

	ofn := win.OPENFILENAME{
		HwndOwner:   owner.Hwnd(),
		LpstrFilter: &zFilters[0],
		LpstrFile:   &result[0],
		NMaxFile:    uint32(len(result)),
		Flags:       co.OFN_HIDEREADONLY | co.OFN_OVERWRITEPROMPT,
		// If absent, no default extension is appended.
		// If present, even if empty, default extension is appended; if no default
		// extension, first one of the filter is appended.
		LpstrDefExt: &defExt[0],
	}

	if !win.GetSaveFileName(&ofn) {
		return "", false
	}
	return syscall.UTF16ToString(result[:]), true
}

func filterToUtf16(filtersWithPipe []string) []uint16 {
	// Each filter as []uint16 with terminating null.
	filters16 := make([][]uint16, 0, len(filtersWithPipe))
	charCount := 0
	for _, filter := range filtersWithPipe {
		filters16 = append(filters16, win.StrToSlice(filter))
		charCount += len(filter) + 1 // also count terminating null
	}

	// Concat all filters into one big []uint16, null-separated, double-null-terminated.
	finalBuf := make([]uint16, 0, charCount+1)
	for _, filter16 := range filters16 {
		finalBuf = append(finalBuf, filter16...)
	}
	finalBuf = append(finalBuf, 0) // double terminating null

	for i := range finalBuf {
		if finalBuf[i] == '|' {
			finalBuf[i] = 0 // replace pipes with nulls
		}
	}
	return finalBuf
}

var (
	_globalMsgBoxHook   = win.HHOOK(0)
	_globalMsgBoxParent = win.HWND(0)
)

// Ordinary MessageBox(), but centered at parent.
func (_SysDlgUtilT) MsgBox(
	parent Window, message, caption string, flags co.MB) co.MBID {

	_globalMsgBoxParent = parent.Hwnd()

	_globalMsgBoxHook = win.SetWindowsHookEx(co.WH_CBT,
		func(code int32, wp win.WPARAM, lp win.LPARAM) uintptr {
			// http://www.codeguru.com/cpp/w-p/win32/messagebox/print.php/c4541
			if co.HCBT(code) == co.HCBT_ACTIVATE {
				hMsgBox := win.HWND(wp)

				if hMsgBox != 0 {
					rcMsgBox := hMsgBox.GetWindowRect()
					rcParent := _globalMsgBoxParent.GetWindowRect()

					rcScreen := win.RECT{}
					win.SystemParametersInfo(
						co.SPI_GETWORKAREA, 0, unsafe.Pointer(&rcScreen), 0) // desktop size

					// Adjusted x,y coordinates to message box window.
					pos := win.POINT{
						X: rcParent.Left +
							(rcParent.Right-rcParent.Left)/2 -
							(rcMsgBox.Right-rcMsgBox.Left)/2,
						Y: rcParent.Top +
							(rcParent.Bottom-rcParent.Top)/2 -
							(rcMsgBox.Bottom-rcMsgBox.Top)/2,
					}

					// Screen out-of-bounds corrections.
					if pos.X < 0 {
						pos.X = 0
					} else if pos.X+(rcMsgBox.Right-rcMsgBox.Left) > rcScreen.Right {
						pos.X = rcScreen.Right - (rcMsgBox.Right - rcMsgBox.Left)
					}
					if pos.Y < 0 {
						pos.Y = 0
					} else if pos.Y+(rcMsgBox.Bottom-rcMsgBox.Top) > rcScreen.Bottom {
						pos.Y = rcScreen.Bottom - (rcMsgBox.Bottom - rcMsgBox.Top)
					}

					hMsgBox.MoveWindow(pos.X, pos.Y,
						int32(rcMsgBox.Right-rcMsgBox.Left),
						int32(rcMsgBox.Bottom-rcMsgBox.Top),
						false)
				}
				_globalMsgBoxHook.UnhookWindowsHookEx() // release global hook
				_globalMsgBoxHook = 0
			}
			return win.HHOOK(0).CallNextHookEx(code, wp, lp)
		},
		win.HINSTANCE(0), win.GetCurrentThreadId())

	return parent.Hwnd().MessageBox(message, caption, flags)
}
