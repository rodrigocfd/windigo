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
	kernel32Dll = syscall.NewLazyDLL("kernel32.dll")

	CloseHandle                     = kernel32Dll.NewProc("CloseHandle")
	CreateDirectory                 = kernel32Dll.NewProc("CreateDirectoryW")
	CreateFile                      = kernel32Dll.NewProc("CreateFileW")
	CreateFileMapping               = kernel32Dll.NewProc("CreateFileMappingW")
	CreateProcess                   = kernel32Dll.NewProc("CreateProcessW")
	DeleteFile                      = kernel32Dll.NewProc("DeleteFileW")
	ExpandEnvironmentStrings        = kernel32Dll.NewProc("ExpandEnvironmentStringsW")
	FileTimeToSystemTime            = kernel32Dll.NewProc("FileTimeToSystemTime")
	FindClose                       = kernel32Dll.NewProc("FindClose")
	FindFirstFile                   = kernel32Dll.NewProc("FindFirstFileW")
	FindNextFile                    = kernel32Dll.NewProc("FindNextFileW")
	GetCurrentProcessId             = kernel32Dll.NewProc("GetCurrentProcessId")
	GetCurrentThreadId              = kernel32Dll.NewProc("GetCurrentThreadId")
	GetFileAttributes               = kernel32Dll.NewProc("GetFileAttributesW")
	GetFileSizeEx                   = kernel32Dll.NewProc("GetFileSizeEx")
	GetModuleHandle                 = kernel32Dll.NewProc("GetModuleHandleW")
	MapViewOfFile                   = kernel32Dll.NewProc("MapViewOfFile")
	MulDiv                          = kernel32Dll.NewProc("MulDiv")
	ReadFile                        = kernel32Dll.NewProc("ReadFile")
	SetEndOfFile                    = kernel32Dll.NewProc("SetEndOfFile")
	SetFilePointerEx                = kernel32Dll.NewProc("SetFilePointerEx")
	Sleep                           = kernel32Dll.NewProc("Sleep")
	SystemTimeToFileTime            = kernel32Dll.NewProc("SystemTimeToFileTime")
	SystemTimeToTzSpecificLocalTime = kernel32Dll.NewProc("SystemTimeToTzSpecificLocalTime")
	TzSpecificLocalTimeToSystemTime = kernel32Dll.NewProc("TzSpecificLocalTimeToSystemTime")
	UnmapViewOfFile                 = kernel32Dll.NewProc("UnmapViewOfFile")
	VerifyVersionInfo               = kernel32Dll.NewProc("VerifyVersionInfoW")
	VerSetConditionMask             = kernel32Dll.NewProc("VerSetConditionMask")
	WriteFile                       = kernel32Dll.NewProc("WriteFile")
)
