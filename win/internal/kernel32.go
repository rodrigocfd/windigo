/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	dllKernel32 = syscall.NewLazyDLL("kernel32.dll")

	CloseHandle                     = dllKernel32.NewProc("CloseHandle")
	CreateDirectory                 = dllKernel32.NewProc("CreateDirectoryW")
	CreateFile                      = dllKernel32.NewProc("CreateFileW")
	CreateFileMapping               = dllKernel32.NewProc("CreateFileMappingW")
	CreateProcess                   = dllKernel32.NewProc("CreateProcessW")
	DeleteFile                      = dllKernel32.NewProc("DeleteFileW")
	ExpandEnvironmentStrings        = dllKernel32.NewProc("ExpandEnvironmentStringsW")
	FileTimeToSystemTime            = dllKernel32.NewProc("FileTimeToSystemTime")
	FindClose                       = dllKernel32.NewProc("FindClose")
	FindFirstFile                   = dllKernel32.NewProc("FindFirstFileW")
	FindNextFile                    = dllKernel32.NewProc("FindNextFileW")
	GetCurrentProcessId             = dllKernel32.NewProc("GetCurrentProcessId")
	GetCurrentThreadId              = dllKernel32.NewProc("GetCurrentThreadId")
	GetFileAttributes               = dllKernel32.NewProc("GetFileAttributesW")
	GetFileSizeEx                   = dllKernel32.NewProc("GetFileSizeEx")
	GetModuleHandle                 = dllKernel32.NewProc("GetModuleHandleW")
	MapViewOfFile                   = dllKernel32.NewProc("MapViewOfFile")
	MulDiv                          = dllKernel32.NewProc("MulDiv")
	ReadFile                        = dllKernel32.NewProc("ReadFile")
	SetEndOfFile                    = dllKernel32.NewProc("SetEndOfFile")
	SetFilePointerEx                = dllKernel32.NewProc("SetFilePointerEx")
	Sleep                           = dllKernel32.NewProc("Sleep")
	SystemTimeToFileTime            = dllKernel32.NewProc("SystemTimeToFileTime")
	SystemTimeToTzSpecificLocalTime = dllKernel32.NewProc("SystemTimeToTzSpecificLocalTime")
	TzSpecificLocalTimeToSystemTime = dllKernel32.NewProc("TzSpecificLocalTimeToSystemTime")
	UnmapViewOfFile                 = dllKernel32.NewProc("UnmapViewOfFile")
	VerifyVersionInfo               = dllKernel32.NewProc("VerifyVersionInfoW")
	VerSetConditionMask             = dllKernel32.NewProc("VerSetConditionMask")
	WriteFile                       = dllKernel32.NewProc("WriteFile")
)
