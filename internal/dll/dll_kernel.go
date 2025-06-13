//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded kernel procedures.
var kernelCache [108]*syscall.Proc

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
	PROC_SetCurrentDirectoryW            PROC_KERNEL = 19 | (360 << 16) | (380 << 32)
	PROC_SystemTimeToFileTime            PROC_KERNEL = 20 | (381 << 16) | (401 << 32)
	PROC_SystemTimeToTzSpecificLocalTime PROC_KERNEL = 21 | (402 << 16) | (433 << 32)
	PROC_CloseHandle                     PROC_KERNEL = 22 | (444 << 16) | (455 << 32)
	PROC_SetFilePointer                  PROC_KERNEL = 23 | (468 << 16) | (482 << 32)
	PROC_SetFilePointerEx                PROC_KERNEL = 24 | (494 << 16) | (510 << 32)
	PROC_CreateFileW                     PROC_KERNEL = 25 | (520 << 16) | (531 << 32)
	PROC_GetFileSizeEx                   PROC_KERNEL = 26 | (532 << 16) | (545 << 32)
	PROC_CreateFileMappingFromApp        PROC_KERNEL = 27 | (546 << 16) | (570 << 32)
	PROC_GetFileTime                     PROC_KERNEL = 28 | (571 << 16) | (582 << 32)
	PROC_LockFile                        PROC_KERNEL = 29 | (583 << 16) | (591 << 32)
	PROC_LockFileEx                      PROC_KERNEL = 30 | (592 << 16) | (602 << 32)
	PROC_ReadFile                        PROC_KERNEL = 31 | (603 << 16) | (611 << 32)
	PROC_SetEndOfFile                    PROC_KERNEL = 32 | (612 << 16) | (624 << 32)
	PROC_UnlockFile                      PROC_KERNEL = 33 | (625 << 16) | (635 << 32)
	PROC_UnlockFileEx                    PROC_KERNEL = 34 | (636 << 16) | (648 << 32)
	PROC_WriteFile                       PROC_KERNEL = 35 | (649 << 16) | (658 << 32)
	PROC_MapViewOfFileFromApp            PROC_KERNEL = 36 | (671 << 16) | (691 << 32)
	PROC_FlushViewOfFile                 PROC_KERNEL = 37 | (692 << 16) | (707 << 32)
	PROC_UnmapViewOfFile                 PROC_KERNEL = 38 | (708 << 16) | (723 << 32)
	PROC_FindFirstFileW                  PROC_KERNEL = 39 | (733 << 16) | (747 << 32)
	PROC_FindClose                       PROC_KERNEL = 40 | (748 << 16) | (757 << 32)
	PROC_FindNextFileW                   PROC_KERNEL = 41 | (758 << 16) | (771 << 32)
	PROC_GlobalAlloc                     PROC_KERNEL = 42 | (783 << 16) | (794 << 32)
	PROC_GlobalFlags                     PROC_KERNEL = 43 | (795 << 16) | (806 << 32)
	PROC_GlobalFree                      PROC_KERNEL = 44 | (807 << 16) | (817 << 32)
	PROC_GlobalLock                      PROC_KERNEL = 45 | (818 << 16) | (828 << 32)
	PROC_GlobalReAlloc                   PROC_KERNEL = 46 | (829 << 16) | (842 << 32)
	PROC_GlobalSize                      PROC_KERNEL = 47 | (843 << 16) | (853 << 32)
	PROC_GlobalUnlock                    PROC_KERNEL = 48 | (854 << 16) | (866 << 32)
	PROC_GetProcessHeap                  PROC_KERNEL = 49 | (876 << 16) | (890 << 32)
	PROC_HeapCreate                      PROC_KERNEL = 50 | (891 << 16) | (901 << 32)
	PROC_HeapAlloc                       PROC_KERNEL = 51 | (902 << 16) | (911 << 32)
	PROC_HeapCompact                     PROC_KERNEL = 52 | (912 << 16) | (923 << 32)
	PROC_HeapDestroy                     PROC_KERNEL = 53 | (924 << 16) | (935 << 32)
	PROC_HeapFree                        PROC_KERNEL = 54 | (936 << 16) | (944 << 32)
	PROC_HeapReAlloc                     PROC_KERNEL = 55 | (945 << 16) | (956 << 32)
	PROC_HeapSize                        PROC_KERNEL = 56 | (957 << 16) | (965 << 32)
	PROC_HeapValidate                    PROC_KERNEL = 57 | (966 << 16) | (978 << 32)
	PROC_GetModuleHandleW                PROC_KERNEL = 58 | (992 << 16) | (1008 << 32)
	PROC_LoadLibraryW                    PROC_KERNEL = 59 | (1009 << 16) | (1021 << 32)
	PROC_FreeLibrary                     PROC_KERNEL = 60 | (1022 << 16) | (1033 << 32)
	PROC_GetModuleFileNameW              PROC_KERNEL = 61 | (1034 << 16) | (1052 << 32)
	PROC_CreateNamedPipeW                PROC_KERNEL = 62 | (1062 << 16) | (1078 << 32)
	PROC_ConnectNamedPipe                PROC_KERNEL = 63 | (1079 << 16) | (1095 << 32)
	PROC_DisconnectNamedPipe             PROC_KERNEL = 64 | (1096 << 16) | (1115 << 32)
	PROC_GetNamedPipeInfo                PROC_KERNEL = 65 | (1116 << 16) | (1132 << 32)
	PROC_PeekNamedPipe                   PROC_KERNEL = 66 | (1133 << 16) | (1146 << 32)
	PROC_GetCurrentProcess               PROC_KERNEL = 67 | (1159 << 16) | (1176 << 32)
	PROC_OpenProcess                     PROC_KERNEL = 68 | (1177 << 16) | (1188 << 32)
	PROC_GetExitCodeProcess              PROC_KERNEL = 69 | (1189 << 16) | (1207 << 32)
	PROC_GetPriorityClass                PROC_KERNEL = 70 | (1208 << 16) | (1224 << 32)
	PROC_GetProcessHandleCount           PROC_KERNEL = 71 | (1225 << 16) | (1246 << 32)
	PROC_GetProcessId                    PROC_KERNEL = 72 | (1247 << 16) | (1259 << 32)
	PROC_GetProcessPriorityBoost         PROC_KERNEL = 73 | (1260 << 16) | (1283 << 32)
	PROC_GetProcessShutdownParameters    PROC_KERNEL = 74 | (1284 << 16) | (1312 << 32)
	PROC_GetProcessTimes                 PROC_KERNEL = 75 | (1313 << 16) | (1328 << 32)
	PROC_GetProcessVersion               PROC_KERNEL = 76 | (1329 << 16) | (1346 << 32)
	PROC_IsProcessCritical               PROC_KERNEL = 77 | (1347 << 16) | (1364 << 32)
	PROC_IsWow64Process                  PROC_KERNEL = 78 | (1365 << 16) | (1379 << 32)
	PROC_QueryFullProcessImageNameW      PROC_KERNEL = 79 | (1380 << 16) | (1406 << 32)
	PROC_QueryProcessAffinityUpdateMode  PROC_KERNEL = 80 | (1407 << 16) | (1437 << 32)
	PROC_QueryProcessCycleTime           PROC_KERNEL = 81 | (1438 << 16) | (1459 << 32)
	PROC_ReadProcessMemory               PROC_KERNEL = 82 | (1460 << 16) | (1477 << 32)
	PROC_SetPriorityClass                PROC_KERNEL = 83 | (1478 << 16) | (1494 << 32)
	PROC_SetProcessAffinityUpdateMode    PROC_KERNEL = 84 | (1495 << 16) | (1523 << 32)
	PROC_TerminateProcess                PROC_KERNEL = 85 | (1524 << 16) | (1540 << 32)
	PROC_VirtualQueryEx                  PROC_KERNEL = 86 | (1541 << 16) | (1555 << 32)
	PROC_WriteProcessMemory              PROC_KERNEL = 87 | (1556 << 16) | (1574 << 32)
	PROC_CreateToolhelp32Snapshot        PROC_KERNEL = 88 | (1588 << 16) | (1612 << 32)
	PROC_Module32FirstW                  PROC_KERNEL = 89 | (1613 << 16) | (1627 << 32)
	PROC_Module32NextW                   PROC_KERNEL = 90 | (1628 << 16) | (1641 << 32)
	PROC_Process32FirstW                 PROC_KERNEL = 91 | (1642 << 16) | (1657 << 32)
	PROC_Process32NextW                  PROC_KERNEL = 92 | (1658 << 16) | (1672 << 32)
	PROC_Thread32First                   PROC_KERNEL = 93 | (1673 << 16) | (1686 << 32)
	PROC_Thread32Next                    PROC_KERNEL = 94 | (1687 << 16) | (1699 << 32)
	PROC_GetCurrentThread                PROC_KERNEL = 95 | (1711 << 16) | (1727 << 32)
	PROC_GetExitCodeThread               PROC_KERNEL = 96 | (1728 << 16) | (1745 << 32)
	PROC_GetProcessIdOfThread            PROC_KERNEL = 97 | (1746 << 16) | (1766 << 32)
	PROC_GetThreadId                     PROC_KERNEL = 98 | (1767 << 16) | (1778 << 32)
	PROC_GetThreadIdealProcessorEx       PROC_KERNEL = 99 | (1779 << 16) | (1804 << 32)
	PROC_GetThreadIOPendingFlag          PROC_KERNEL = 100 | (1805 << 16) | (1827 << 32)
	PROC_GetThreadPriority               PROC_KERNEL = 101 | (1828 << 16) | (1845 << 32)
	PROC_GetThreadPriorityBoost          PROC_KERNEL = 102 | (1846 << 16) | (1868 << 32)
	PROC_GetThreadTimes                  PROC_KERNEL = 103 | (1869 << 16) | (1883 << 32)
	PROC_ResumeThread                    PROC_KERNEL = 104 | (1884 << 16) | (1896 << 32)
	PROC_TerminateThread                 PROC_KERNEL = 105 | (1897 << 16) | (1912 << 32)
	PROC_SuspendThread                   PROC_KERNEL = 106 | (1913 << 16) | (1926 << 32)
	PROC_WaitForSingleObject             PROC_KERNEL = 107 | (1927 << 16) | (1946 << 32)
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
