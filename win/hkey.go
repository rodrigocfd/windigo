/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"sort"
	"strings"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HKEY HANDLE

// Returned by RegEnumValue.
type RegistryValueInfo struct {
	DataType co.REG
	Name     string
	Size     uint32 // In uint8, notice wide strings are uint16 chars.
}

func (hKey HKEY) RegCloseKey() {
	ret, _, _ := syscall.Syscall(proc.RegCloseKey.Addr(), 1,
		uintptr(hKey), 0, 0)
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegCloseKey failed: %d %s\n",
			ret, syscall.Errno(ret).Error()))
	}
}

func (hKey HKEY) RegEnumValue() []RegistryValueInfo {
	retVals := make([]RegistryValueInfo, 0)
	dwIndex := uint32(0)
	nameBufSize := 64 // arbitrary
	nameBuf := make([]uint16, nameBufSize)
	dataType := co.REG_NONE
	dataBufSize := uint32(0)

	for {
		ret, _, _ := syscall.Syscall9(proc.RegEnumValue.Addr(), 8,
			uintptr(hKey), uintptr(dwIndex), uintptr(unsafe.Pointer(&nameBuf[0])),
			uintptr(unsafe.Pointer(&nameBufSize)), 0,
			uintptr(unsafe.Pointer(&dataType)),
			0, uintptr(unsafe.Pointer(&dataBufSize)), 0)

		if co.ERROR(ret) == co.ERROR_SUCCESS {
			retVals = append(retVals, RegistryValueInfo{
				DataType: dataType,
				Name:     syscall.UTF16ToString(nameBuf),
				Size:     dataBufSize,
			})
			dwIndex++
		} else if co.ERROR(ret) == co.ERROR_NO_MORE_ITEMS { // we're done
			break
		} else if co.ERROR(ret) == co.ERROR_MORE_DATA { // increase buffer size
			nameBufSize += 64 // arbitrary
			nameBuf = make([]uint16, nameBufSize)
		} else {
			panic(fmt.Sprintf("RegEnumValue failed: %d %s\n",
				ret, syscall.Errno(ret).Error()))
		}
	}

	sort.Slice(retVals, func(i, j int) bool {
		return strings.ToUpper(retVals[i].Name) < strings.ToUpper(retVals[j].Name)
	})
	return retVals
}

func RegOpenKeyEx(hKeyPredef co.HKEY, lpSubKey string, ulOptions co.REG_OPTION,
	samDesired co.KEY) HKEY {

	hKey := HKEY(0)
	ret, _, _ := syscall.Syscall6(proc.RegOpenKeyEx.Addr(), 5,
		uintptr(hKeyPredef), uintptr(unsafe.Pointer(StrToUtf16Ptr(lpSubKey))),
		uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(&hKey)),
		0)
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegOpenKeyEx failed: %d %s\n",
			ret, syscall.Errno(ret).Error()))
	}
	return hKey
}

func (hKey HKEY) RegQueryValueExString(lpValueName string) string {
	dataType := co.REG_NONE
	dataBufSize := uint32(0)

	ret, _, _ := syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(StrToUtf16Ptr(lpValueName))), 0,
		uintptr(unsafe.Pointer(&dataType)),
		0, uintptr(unsafe.Pointer(&dataBufSize)))
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegQueryValueEx failed: %d %s\n",
			ret, syscall.Errno(ret).Error()))
	} else if dataType != co.REG_SZ {
		panic(fmt.Sprintf("Registry data isn't string, type is %d.\n", dataType))
	}

	dataBuf := make([]uint16, dataBufSize/2) // returned size is in bytes, we've got wide chars
	ret, _, _ = syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(StrToUtf16Ptr(lpValueName))), 0,
		uintptr(unsafe.Pointer(&dataType)), uintptr(unsafe.Pointer(&dataBuf[0])),
		uintptr(unsafe.Pointer(&dataBufSize)))
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegQueryValueEx failed: %d %s\n",
			ret, syscall.Errno(ret).Error()))
	}

	return syscall.UTF16ToString(dataBuf)
}

func (hKey HKEY) RegQueryValueExUint32(lpValueName string) uint32 {
	dataType := co.REG_NONE
	dataBufSize := uint32(0)

	ret, _, _ := syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(StrToUtf16Ptr(lpValueName))), 0,
		uintptr(unsafe.Pointer(&dataType)),
		0, uintptr(unsafe.Pointer(&dataBufSize)))
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegQueryValueEx failed: %d %s\n",
			ret, syscall.Errno(ret).Error()))
	} else if dataType != co.REG_DWORD {
		panic(fmt.Sprintf("Registry data isn't uint32, type is %d.\n", dataType))
	}

	dataBuf := uint32(0) // 4 bytes
	ret, _, _ = syscall.Syscall6(proc.RegQueryValueEx.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(StrToUtf16Ptr(lpValueName))), 0,
		uintptr(unsafe.Pointer(&dataType)), uintptr(unsafe.Pointer(&dataBuf)),
		uintptr(unsafe.Pointer(&dataBufSize)))
	if co.ERROR(ret) != co.ERROR_SUCCESS {
		panic(fmt.Sprintf("RegQueryValueEx failed: %d %s\n",
			ret, syscall.Errno(ret).Error()))
	}

	return dataBuf
}
