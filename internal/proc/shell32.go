//go:build windows

package proc

import (
	"syscall"
)

var (
	shell32 = syscall.NewLazyDLL("shell32.dll")

	CommandLineToArgv           = shell32.NewProc("CommandLineToArgvW")
	DragAcceptFiles             = shell32.NewProc("DragAcceptFiles")
	DragFinish                  = shell32.NewProc("DragFinish")
	DragQueryFile               = shell32.NewProc("DragQueryFileW")
	DragQueryPoint              = shell32.NewProc("DragQueryPoint")
	DuplicateIcon               = shell32.NewProc("DuplicateIcon")
	ExtractIconEx               = shell32.NewProc("ExtractIconExW")
	SHCreateItemFromParsingName = shell32.NewProc("SHCreateItemFromParsingName")
	Shell_NotifyIcon            = shell32.NewProc("Shell_NotifyIconW")
	SHGetFileInfo               = shell32.NewProc("SHGetFileInfoW")
)
