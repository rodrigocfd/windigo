//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded kernel procedures.
var kernelCache [119]*syscall.Proc

// Loads kernel procedures.
func Kernel(procId PROC_KERNEL) uintptr {
	return LoadProc(SYSDLL_kernel32, kernelCache[:], kernelProcStr, uint64(procId)).Addr()
}

type PROC_KERNEL uint64 // Procedure identifiers for kernel.

// Auto-generated kernel procedure identifier: cache index | str start | str past-end.
const (
	PROC_VerifyVersionInfoW              PROC_KERNEL = 0 | (12 << 16) | (30 << 32)
	PROC_VerSetConditionMask             PROC_KERNEL = 1 | (31 << 16) | (50 << 32)
	PROC_CopyFileW                       PROC_KERNEL = 2 | (60 << 16) | (69 << 32)
	PROC_CreateDirectoryW                PROC_KERNEL = 3 | (70 << 16) | (86 << 32)
	PROC_CreateProcessW                  PROC_KERNEL = 4 | (87 << 16) | (101 << 32)
	PROC_DeleteFileW                     PROC_KERNEL = 5 | (102 << 16) | (113 << 32)
	PROC_ExitProcess                     PROC_KERNEL = 6 | (114 << 16) | (125 << 32)
	PROC_ExpandEnvironmentStringsW       PROC_KERNEL = 7 | (126 << 16) | (151 << 32)
	PROC_FileTimeToSystemTime            PROC_KERNEL = 8 | (152 << 16) | (172 << 32)
	PROC_GetCommandLineW                 PROC_KERNEL = 9 | (173 << 16) | (188 << 32)
	PROC_GetCurrentProcessId             PROC_KERNEL = 10 | (189 << 16) | (208 << 32)
	PROC_GetCurrentThreadId              PROC_KERNEL = 11 | (209 << 16) | (227 << 32)
	PROC_FreeEnvironmentStringsW         PROC_KERNEL = 12 | (228 << 16) | (251 << 32)
	PROC_GetEnvironmentStringsW          PROC_KERNEL = 13 | (252 << 16) | (274 << 32)
	PROC_GetFileAttributesW              PROC_KERNEL = 14 | (275 << 16) | (293 << 32)
	PROC_GetLocalTime                    PROC_KERNEL = 15 | (294 << 16) | (306 << 32)
	PROC_GetStartupInfoW                 PROC_KERNEL = 16 | (307 << 16) | (322 << 32)
	PROC_GetTimeZoneInformation          PROC_KERNEL = 17 | (323 << 16) | (345 << 32)
	PROC_GetSystemInfo                   PROC_KERNEL = 18 | (346 << 16) | (359 << 32)
	PROC_SetConsoleTitleW                PROC_KERNEL = 19 | (360 << 16) | (376 << 32)
	PROC_SetCurrentDirectoryW            PROC_KERNEL = 20 | (377 << 16) | (397 << 32)
	PROC_SystemTimeToFileTime            PROC_KERNEL = 21 | (398 << 16) | (418 << 32)
	PROC_SystemTimeToTzSpecificLocalTime PROC_KERNEL = 22 | (419 << 16) | (450 << 32)
	PROC_CloseHandle                     PROC_KERNEL = 23 | (461 << 16) | (472 << 32)
	PROC_SetFilePointer                  PROC_KERNEL = 24 | (485 << 16) | (499 << 32)
	PROC_SetFilePointerEx                PROC_KERNEL = 25 | (511 << 16) | (527 << 32)
	PROC_CreateFileW                     PROC_KERNEL = 26 | (537 << 16) | (548 << 32)
	PROC_GetFileSizeEx                   PROC_KERNEL = 27 | (549 << 16) | (562 << 32)
	PROC_CreateFileMappingFromApp        PROC_KERNEL = 28 | (563 << 16) | (587 << 32)
	PROC_GetFileTime                     PROC_KERNEL = 29 | (588 << 16) | (599 << 32)
	PROC_LockFile                        PROC_KERNEL = 30 | (600 << 16) | (608 << 32)
	PROC_LockFileEx                      PROC_KERNEL = 31 | (609 << 16) | (619 << 32)
	PROC_ReadFile                        PROC_KERNEL = 32 | (620 << 16) | (628 << 32)
	PROC_SetEndOfFile                    PROC_KERNEL = 33 | (629 << 16) | (641 << 32)
	PROC_UnlockFile                      PROC_KERNEL = 34 | (642 << 16) | (652 << 32)
	PROC_UnlockFileEx                    PROC_KERNEL = 35 | (653 << 16) | (665 << 32)
	PROC_WriteFile                       PROC_KERNEL = 36 | (666 << 16) | (675 << 32)
	PROC_MapViewOfFileFromApp            PROC_KERNEL = 37 | (688 << 16) | (708 << 32)
	PROC_FlushViewOfFile                 PROC_KERNEL = 38 | (709 << 16) | (724 << 32)
	PROC_UnmapViewOfFile                 PROC_KERNEL = 39 | (725 << 16) | (740 << 32)
	PROC_FindFirstFileW                  PROC_KERNEL = 40 | (750 << 16) | (764 << 32)
	PROC_FindClose                       PROC_KERNEL = 41 | (765 << 16) | (774 << 32)
	PROC_FindNextFileW                   PROC_KERNEL = 42 | (775 << 16) | (788 << 32)
	PROC_GlobalAlloc                     PROC_KERNEL = 43 | (800 << 16) | (811 << 32)
	PROC_GlobalFlags                     PROC_KERNEL = 44 | (812 << 16) | (823 << 32)
	PROC_GlobalFree                      PROC_KERNEL = 45 | (824 << 16) | (834 << 32)
	PROC_GlobalLock                      PROC_KERNEL = 46 | (835 << 16) | (845 << 32)
	PROC_GlobalReAlloc                   PROC_KERNEL = 47 | (846 << 16) | (859 << 32)
	PROC_GlobalSize                      PROC_KERNEL = 48 | (860 << 16) | (870 << 32)
	PROC_GlobalUnlock                    PROC_KERNEL = 49 | (871 << 16) | (883 << 32)
	PROC_GetProcessHeap                  PROC_KERNEL = 50 | (893 << 16) | (907 << 32)
	PROC_HeapCreate                      PROC_KERNEL = 51 | (908 << 16) | (918 << 32)
	PROC_HeapAlloc                       PROC_KERNEL = 52 | (919 << 16) | (928 << 32)
	PROC_HeapCompact                     PROC_KERNEL = 53 | (929 << 16) | (940 << 32)
	PROC_HeapDestroy                     PROC_KERNEL = 54 | (941 << 16) | (952 << 32)
	PROC_HeapFree                        PROC_KERNEL = 55 | (953 << 16) | (961 << 32)
	PROC_HeapReAlloc                     PROC_KERNEL = 56 | (962 << 16) | (973 << 32)
	PROC_HeapSize                        PROC_KERNEL = 57 | (974 << 16) | (982 << 32)
	PROC_HeapValidate                    PROC_KERNEL = 58 | (983 << 16) | (995 << 32)
	PROC_GetModuleHandleW                PROC_KERNEL = 59 | (1009 << 16) | (1025 << 32)
	PROC_LoadLibraryW                    PROC_KERNEL = 60 | (1026 << 16) | (1038 << 32)
	PROC_FreeLibrary                     PROC_KERNEL = 61 | (1039 << 16) | (1050 << 32)
	PROC_GetModuleFileNameW              PROC_KERNEL = 62 | (1051 << 16) | (1069 << 32)
	PROC_CreateNamedPipeW                PROC_KERNEL = 63 | (1079 << 16) | (1095 << 32)
	PROC_ConnectNamedPipe                PROC_KERNEL = 64 | (1096 << 16) | (1112 << 32)
	PROC_DisconnectNamedPipe             PROC_KERNEL = 65 | (1113 << 16) | (1132 << 32)
	PROC_GetNamedPipeInfo                PROC_KERNEL = 66 | (1133 << 16) | (1149 << 32)
	PROC_PeekNamedPipe                   PROC_KERNEL = 67 | (1150 << 16) | (1163 << 32)
	PROC_GetCurrentProcess               PROC_KERNEL = 68 | (1176 << 16) | (1193 << 32)
	PROC_OpenProcess                     PROC_KERNEL = 69 | (1194 << 16) | (1205 << 32)
	PROC_GetExitCodeProcess              PROC_KERNEL = 70 | (1206 << 16) | (1224 << 32)
	PROC_GetPriorityClass                PROC_KERNEL = 71 | (1225 << 16) | (1241 << 32)
	PROC_GetProcessHandleCount           PROC_KERNEL = 72 | (1242 << 16) | (1263 << 32)
	PROC_GetProcessId                    PROC_KERNEL = 73 | (1264 << 16) | (1276 << 32)
	PROC_GetProcessPriorityBoost         PROC_KERNEL = 74 | (1277 << 16) | (1300 << 32)
	PROC_GetProcessShutdownParameters    PROC_KERNEL = 75 | (1301 << 16) | (1329 << 32)
	PROC_GetProcessTimes                 PROC_KERNEL = 76 | (1330 << 16) | (1345 << 32)
	PROC_GetProcessVersion               PROC_KERNEL = 77 | (1346 << 16) | (1363 << 32)
	PROC_IsProcessCritical               PROC_KERNEL = 78 | (1364 << 16) | (1381 << 32)
	PROC_IsWow64Process                  PROC_KERNEL = 79 | (1382 << 16) | (1396 << 32)
	PROC_QueryFullProcessImageNameW      PROC_KERNEL = 80 | (1397 << 16) | (1423 << 32)
	PROC_QueryProcessAffinityUpdateMode  PROC_KERNEL = 81 | (1424 << 16) | (1454 << 32)
	PROC_QueryProcessCycleTime           PROC_KERNEL = 82 | (1455 << 16) | (1476 << 32)
	PROC_ReadProcessMemory               PROC_KERNEL = 83 | (1477 << 16) | (1494 << 32)
	PROC_SetPriorityClass                PROC_KERNEL = 84 | (1495 << 16) | (1511 << 32)
	PROC_SetProcessAffinityUpdateMode    PROC_KERNEL = 85 | (1512 << 16) | (1540 << 32)
	PROC_TerminateProcess                PROC_KERNEL = 86 | (1541 << 16) | (1557 << 32)
	PROC_VirtualQueryEx                  PROC_KERNEL = 87 | (1558 << 16) | (1572 << 32)
	PROC_WriteProcessMemory              PROC_KERNEL = 88 | (1573 << 16) | (1591 << 32)
	PROC_CreateToolhelp32Snapshot        PROC_KERNEL = 89 | (1605 << 16) | (1629 << 32)
	PROC_Module32FirstW                  PROC_KERNEL = 90 | (1630 << 16) | (1644 << 32)
	PROC_Module32NextW                   PROC_KERNEL = 91 | (1645 << 16) | (1658 << 32)
	PROC_Process32FirstW                 PROC_KERNEL = 92 | (1659 << 16) | (1674 << 32)
	PROC_Process32NextW                  PROC_KERNEL = 93 | (1675 << 16) | (1689 << 32)
	PROC_Thread32First                   PROC_KERNEL = 94 | (1690 << 16) | (1703 << 32)
	PROC_Thread32Next                    PROC_KERNEL = 95 | (1704 << 16) | (1716 << 32)
	PROC_GetStdHandle                    PROC_KERNEL = 96 | (1725 << 16) | (1737 << 32)
	PROC_GetCurrentConsoleFont           PROC_KERNEL = 97 | (1738 << 16) | (1759 << 32)
	PROC_ReadConsoleW                    PROC_KERNEL = 98 | (1760 << 16) | (1772 << 32)
	PROC_SetConsoleCursorInfo            PROC_KERNEL = 99 | (1773 << 16) | (1793 << 32)
	PROC_SetConsoleCursorPosition        PROC_KERNEL = 100 | (1794 << 16) | (1818 << 32)
	PROC_SetConsoleDisplayMode           PROC_KERNEL = 101 | (1819 << 16) | (1840 << 32)
	PROC_SetConsoleMode                  PROC_KERNEL = 102 | (1841 << 16) | (1855 << 32)
	PROC_SetConsoleScreenBufferSize      PROC_KERNEL = 103 | (1856 << 16) | (1882 << 32)
	PROC_SetConsoleTextAttribute         PROC_KERNEL = 104 | (1883 << 16) | (1906 << 32)
	PROC_WriteConsoleW                   PROC_KERNEL = 105 | (1907 << 16) | (1920 << 32)
	PROC_GetCurrentThread                PROC_KERNEL = 106 | (1932 << 16) | (1948 << 32)
	PROC_GetExitCodeThread               PROC_KERNEL = 107 | (1949 << 16) | (1966 << 32)
	PROC_GetProcessIdOfThread            PROC_KERNEL = 108 | (1967 << 16) | (1987 << 32)
	PROC_GetThreadId                     PROC_KERNEL = 109 | (1988 << 16) | (1999 << 32)
	PROC_GetThreadIdealProcessorEx       PROC_KERNEL = 110 | (2000 << 16) | (2025 << 32)
	PROC_GetThreadIOPendingFlag          PROC_KERNEL = 111 | (2026 << 16) | (2048 << 32)
	PROC_GetThreadPriority               PROC_KERNEL = 112 | (2049 << 16) | (2066 << 32)
	PROC_GetThreadPriorityBoost          PROC_KERNEL = 113 | (2067 << 16) | (2089 << 32)
	PROC_GetThreadTimes                  PROC_KERNEL = 114 | (2090 << 16) | (2104 << 32)
	PROC_ResumeThread                    PROC_KERNEL = 115 | (2105 << 16) | (2117 << 32)
	PROC_TerminateThread                 PROC_KERNEL = 116 | (2118 << 16) | (2133 << 32)
	PROC_SuspendThread                   PROC_KERNEL = 117 | (2134 << 16) | (2147 << 32)
	PROC_WaitForSingleObject             PROC_KERNEL = 118 | (2148 << 16) | (2167 << 32)
)

// Declaration of kernel procedure names.
const kernelProcStr = `
--funcs386
VerifyVersionInfoW
VerSetConditionMask

--funcs
CopyFileW
CreateDirectoryW
CreateProcessW
DeleteFileW
ExitProcess
ExpandEnvironmentStringsW
FileTimeToSystemTime
GetCommandLineW
GetCurrentProcessId
GetCurrentThreadId
FreeEnvironmentStringsW
GetEnvironmentStringsW
GetFileAttributesW
GetLocalTime
GetStartupInfoW
GetTimeZoneInformation
GetSystemInfo
SetConsoleTitleW
SetCurrentDirectoryW
SystemTimeToFileTime
SystemTimeToTzSpecificLocalTime

--handle
CloseHandle

--hfile386
SetFilePointer

--hfile64
SetFilePointerEx

--hfile
CreateFileW
GetFileSizeEx
CreateFileMappingFromApp
GetFileTime
LockFile
LockFileEx
ReadFile
SetEndOfFile
UnlockFile
UnlockFileEx
WriteFile

--hfilemap
MapViewOfFileFromApp
FlushViewOfFile
UnmapViewOfFile

--hfind
FindFirstFileW
FindClose
FindNextFileW

--hglobal
GlobalAlloc
GlobalFlags
GlobalFree
GlobalLock
GlobalReAlloc
GlobalSize
GlobalUnlock

--hheap
GetProcessHeap
HeapCreate
HeapAlloc
HeapCompact
HeapDestroy
HeapFree
HeapReAlloc
HeapSize
HeapValidate

--hinstance
GetModuleHandleW
LoadLibraryW
FreeLibrary
GetModuleFileNameW

--hpipe
CreateNamedPipeW
ConnectNamedPipe
DisconnectNamedPipe
GetNamedPipeInfo
PeekNamedPipe

--hprocess
GetCurrentProcess
OpenProcess
GetExitCodeProcess
GetPriorityClass
GetProcessHandleCount
GetProcessId
GetProcessPriorityBoost
GetProcessShutdownParameters
GetProcessTimes
GetProcessVersion
IsProcessCritical
IsWow64Process
QueryFullProcessImageNameW
QueryProcessAffinityUpdateMode
QueryProcessCycleTime
ReadProcessMemory
SetPriorityClass
SetProcessAffinityUpdateMode
TerminateProcess
VirtualQueryEx
WriteProcessMemory

--hprocsnap
CreateToolhelp32Snapshot
Module32FirstW
Module32NextW
Process32FirstW
Process32NextW
Thread32First
Thread32Next

--hstd
GetStdHandle
GetCurrentConsoleFont
ReadConsoleW
SetConsoleCursorInfo
SetConsoleCursorPosition
SetConsoleDisplayMode
SetConsoleMode
SetConsoleScreenBufferSize
SetConsoleTextAttribute
WriteConsoleW

--hthread
GetCurrentThread
GetExitCodeThread
GetProcessIdOfThread
GetThreadId
GetThreadIdealProcessorEx
GetThreadIOPendingFlag
GetThreadPriority
GetThreadPriorityBoost
GetThreadTimes
ResumeThread
TerminateThread
SuspendThread
WaitForSingleObject
`
