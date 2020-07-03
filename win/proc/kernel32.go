/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	dllKernel32 = syscall.NewLazyDLL("kernel32.dll")

	CloseHandle         = dllKernel32.NewProc("CloseHandle")
	CreateDirectory     = dllKernel32.NewProc("CreateDirectoryW")
	CreateFile          = dllKernel32.NewProc("CreateFileW")
	CreateFileMapping   = dllKernel32.NewProc("CreateFileMappingW")
	DeleteFile          = dllKernel32.NewProc("DeleteFileW")
	GetFileAttributes   = dllKernel32.NewProc("GetFileAttributesW")
	GetFileSize         = dllKernel32.NewProc("GetFileSize")
	GetFileSizeEx       = dllKernel32.NewProc("GetFileSizeEx")
	GetModuleHandle     = dllKernel32.NewProc("GetModuleHandleW")
	MapViewOfFile       = dllKernel32.NewProc("MapViewOfFile")
	MulDiv              = dllKernel32.NewProc("MulDiv")
	ReadFile            = dllKernel32.NewProc("ReadFile")
	SetEndOfFile        = dllKernel32.NewProc("SetEndOfFile")
	SetFilePointer      = dllKernel32.NewProc("SetFilePointer")
	SetFilePointerEx    = dllKernel32.NewProc("SetFilePointerEx")
	UnmapViewOfFile     = dllKernel32.NewProc("UnmapViewOfFile")
	VerifyVersionInfo   = dllKernel32.NewProc("VerifyVersionInfoW")
	VerSetConditionMask = dllKernel32.NewProc("VerSetConditionMask")
	WriteFile           = dllKernel32.NewProc("WriteFile")
)
