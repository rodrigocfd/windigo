/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"sort"
	"unsafe"
	"windigo/co"
	"windigo/com/shell"
	"windigo/win"
)

type _SysDlgT struct{}

// System dialogs.
var SysDlg _SysDlgT

// Shows the open file system dialog, choice restricted to 1 file.
func (_SysDlgT) OpenSingleFile(
	parent Parent, filterSpec []shell.FilterSpec) (string, bool) {

	fileOpenDialog := shell.CoCreateIFileOpenDialog(co.CLSCTX_INPROC_SERVER)
	defer fileOpenDialog.Release()

	flags := fileOpenDialog.GetOptions()
	fileOpenDialog.SetOptions(flags |
		shell.FOS_FORCEFILESYSTEM | shell.FOS_FILEMUSTEXIST)

	fileOpenDialog.SetFileTypes(filterSpec)
	fileOpenDialog.SetFileTypeIndex(0) // first filter chosen by default

	if !fileOpenDialog.Show(parent.Hwnd()) {
		return "", false // user cancelled
	}

	shellItem := fileOpenDialog.GetResult()
	defer shellItem.Release()

	return shellItem.GetDisplayName(shell.SIGDN_FILESYSPATH), true
}

// Shows the open file system dialog, user can choose multiple files.
func (_SysDlgT) OpenMultipleFiles(
	parent Parent, filterSpec []shell.FilterSpec) ([]string, bool) {

	fileOpenDialog := shell.CoCreateIFileOpenDialog(co.CLSCTX_INPROC_SERVER)
	defer fileOpenDialog.Release()

	flags := fileOpenDialog.GetOptions()
	fileOpenDialog.SetOptions(flags |
		shell.FOS_FORCEFILESYSTEM | shell.FOS_FILEMUSTEXIST | shell.FOS_ALLOWMULTISELECT)

	fileOpenDialog.SetFileTypes(filterSpec)
	fileOpenDialog.SetFileTypeIndex(0) // first filter chosen by default

	if !fileOpenDialog.Show(parent.Hwnd()) {
		return nil, false // user cancelled
	}

	shellItemArray := fileOpenDialog.GetResults()
	defer shellItemArray.Release()

	files := shellItemArray.GetDisplayNames()
	sort.Strings(files)
	return files, true
}

// Shows the file save system dialog.
func (_SysDlgT) SaveFile(
	parent Parent, defaultPath, defaultFileName string,
	filterSpec []shell.FilterSpec) (string, bool) {

	fileSaveDialog := shell.CoCreateIFileSaveDialog(co.CLSCTX_INPROC_SERVER)
	defer fileSaveDialog.Release()

	flags := fileSaveDialog.GetOptions()
	fileSaveDialog.SetOptions(flags | shell.FOS_FORCEFILESYSTEM)

	fileSaveDialog.SetFileTypes(filterSpec)
	fileSaveDialog.SetFileTypeIndex(0) // first filter chosen by default

	if defaultPath != "" {
		shellItem := shell.NewShellItem(defaultPath)
		fileSaveDialog.SetFolder(shellItem)
		shellItem.Release()
	}
	if defaultFileName != "" {
		fileSaveDialog.SetFileName(defaultFileName)
	}

	if !fileSaveDialog.Show(parent.Hwnd()) {
		return "", false // user cancelled
	}

	shellItem := fileSaveDialog.GetResult()
	defer shellItem.Release()

	return shellItem.GetDisplayName(shell.SIGDN_FILESYSPATH), true
}

// Shows the choose folder system dialog.
//
// The returned file path won't have a trailing slash.
func (_SysDlgT) ChooseFolder(parent Parent) (string, bool) {
	fileOpenDialog := shell.CoCreateIFileOpenDialog(co.CLSCTX_INPROC_SERVER)
	defer fileOpenDialog.Release()

	flags := fileOpenDialog.GetOptions()
	fileOpenDialog.SetOptions(flags |
		shell.FOS_FORCEFILESYSTEM | shell.FOS_PICKFOLDERS)

	if !fileOpenDialog.Show(parent.Hwnd()) {
		return "", false // user cancelled
	}

	shellItem := fileOpenDialog.GetResult()
	defer shellItem.Release()

	return shellItem.GetDisplayName(shell.SIGDN_FILESYSPATH), true
}

var (
	_globalMsgBoxHook   = win.HHOOK(0)
	_globalMsgBoxParent = win.HWND(0)
)

// Ordinary MessageBox(), but centered at parent.
func (_SysDlgT) MsgBox(
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
