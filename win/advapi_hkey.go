//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// Handle to a [registry key].
//
// [registry key]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hkey
type HKEY HANDLE

// [Predefined] registry key.
//
// [Predefined]: https://learn.microsoft.com/en-us/windows/win32/sysinfo/predefined-keys
const (
	HKEY_CLASSES_ROOT        HKEY = 0x8000_0000
	HKEY_CURRENT_USER        HKEY = 0x8000_0001
	HKEY_LOCAL_MACHINE       HKEY = 0x8000_0002
	HKEY_USERS               HKEY = 0x8000_0003
	HKEY_PERFORMANCE_DATA    HKEY = 0x8000_0004
	HKEY_PERFORMANCE_TEXT    HKEY = 0x8000_0050
	HKEY_PERFORMANCE_NLSTEXT HKEY = 0x8000_0060
	HKEY_CURRENT_CONFIG      HKEY = 0x8000_0005
)

// [RegConnectRegistry] function.
//
// Panics if predef_key is different from:
//   - [HKEY_LOCAL_MACHINE];
//   - [HKEY_PERFORMANCE_DATA];
//   - [HKEY_USERS].
//
// ⚠️ You must defer [HKEY.RegCloseKey].
//
// # Example
//
//	hKey, _ := win.RegConnectRegistry(
//		"\\computername",
//		win.HKEY_CURRENT_USER,
//	)
//	defer hKey.RegCloseKey()
//
// [RegConnectRegistry]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regconnectregistryw
func RegConnectRegistry(machineName string, hKey HKEY) (HKEY, error) {
	if hKey != HKEY_LOCAL_MACHINE && hKey != HKEY_PERFORMANCE_DATA && hKey != HKEY_USERS {
		panic("Invalid HKEY.")
	}

	machineName16 := wstr.NewBufWith[wstr.Stack20](machineName, wstr.EMPTY_IS_NIL)
	var openedKey HKEY

	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegConnectRegistryW),
		uintptr(machineName16.UnsafePtr()),
		uintptr(hKey),
		uintptr(unsafe.Pointer(&openedKey)))

	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

// [RegOpenCurrentUser] function.
//
// ⚠️ You must defer [HKEY.RegCloseKey].
//
// [RegOpenCurrentUser]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regopencurrentuser
func RegOpenCurrentUser(accessRights co.KEY) (HKEY, error) {
	var openedKey HKEY
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegOpenCurrentUser),
		uintptr(accessRights),
		uintptr(unsafe.Pointer(&openedKey)))
	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

// [RegCloseKey] function.
//
// [RegCloseKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func (hKey HKEY) RegCloseKey() error {
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegCloseKey),
		uintptr(hKey))
	return utl.ZeroAsSysError(ret)
}

// [RegCopyTree] function.
//
// [RegCopyTree]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regcopytreew
func (hKey HKEY) RegCopyTree(subKey string, dest HKEY) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.EMPTY_IS_NIL)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegCopyTreeW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()))
	return utl.ZeroAsSysError(ret)
}

// [RegCreateKeyEx] function.
//
// ⚠️ You must defer [HKEY.RegCloseKey].
//
// [RegCreateKeyEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regcreatekeyexw
func (hKey HKEY) RegCreateKeyEx(
	subKey string,
	options co.REG_OPTION,
	accessRights co.KEY,
	securityAttributes *SECURITY_ATTRIBUTES,
) (HKEY, error) {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.EMPTY_IS_NIL)
	var openedKey HKEY

	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegCreateKeyExW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()),
		0, 0,
		uintptr(options),
		uintptr(accessRights),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(unsafe.Pointer(&openedKey)))

	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

// [RegDeleteKey] function.
//
// [RegDeleteKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyw
func (hKey HKEY) RegDeleteKey(subKey string) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegDeleteKeyW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()))
	return utl.ZeroAsSysError(ret)
}

// [RegDeleteKeyEx] function.
//
// samDesired must be [co.KEY_WOW64_32KEY] or [co.KEY_WOW64_64KEY].
//
// [RegDeleteKeyEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyexw
func (hKey HKEY) RegDeleteKeyEx(subKey string, samDesired co.KEY) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegDeleteKeyExW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()),
		uintptr(samDesired&(co.KEY_WOW64_32KEY|co.KEY_WOW64_64KEY)),
		0)
	return utl.ZeroAsSysError(ret)
}

// [RegDeleteKeyValue] function.
//
// [RegDeleteKeyValue]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyvaluew
func (hKey HKEY) RegDeleteKeyValue(subKey, valueName string) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.EMPTY_IS_NIL)
	valueName16 := wstr.NewBufWith[wstr.Stack20](valueName, wstr.EMPTY_IS_NIL)

	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegDeleteKeyValueW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()),
		uintptr(valueName16.UnsafePtr()))
	return utl.ZeroAsSysError(ret)
}

// [RegDeleteTree] function.
//
// [RegDeleteTree]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletetreew
func (hKey HKEY) RegDeleteTree(subKey string) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.EMPTY_IS_NIL)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegDeleteTreeW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()))
	return utl.ZeroAsSysError(ret)
}

// [RegDeleteValue] function.
//
// [RegDeleteValue]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletevaluew
func (hKey HKEY) RegDeleteValue(valueName string) error {
	valueName16 := wstr.NewBufWith[wstr.Stack20](valueName, wstr.EMPTY_IS_NIL)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegDeleteValueW),
		uintptr(hKey),
		uintptr(valueName16.UnsafePtr()))
	return utl.ZeroAsSysError(ret)
}

// [RegDisableReflectionKey] function.
//
// [RegDisableReflectionKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdisablereflectionkey
func (hKey HKEY) RegDisableReflectionKey() error {
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegDisableReflectionKey),
		uintptr(hKey))
	return utl.ZeroAsSysError(ret)
}

// [RegEnableReflectionKey] function.
//
// [RegEnableReflectionKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenablereflectionkey
func (hKey HKEY) RegEnableReflectionKey() error {
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegEnableReflectionKey),
		uintptr(hKey))
	return utl.ZeroAsSysError(ret)
}

// [RegEnumKeyEx] function.
//
// # Example
//
//	hKey, _ := win.HKEY_CURRENT_USER.RegOpenKeyEx(
//		"Control Panel",
//		co.REG_OPTION_NONE,
//		co.KEY_READ)
//	defer hKey.RegCloseKey()
//
//	keys, _ := hKey.RegEnumKeyEx()
//	for _, key := range keys {
//		println(key)
//	}
//
// [RegEnumKeyEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumkeyexw
func (hKey HKEY) RegEnumKeyEx() ([]string, error) {
	nfo, err := hKey.RegQueryInfoKey()
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, nfo.NumSubKeys)
	keyNameBuf := wstr.NewBufSized[wstr.Stack64](nfo.MaxSubKeyNameLen + 1)

	for i := uint(0); i < nfo.NumSubKeys; i++ {
		szKeyNameBuf := uint32(keyNameBuf.Len())

		ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegEnumKeyExW),
			uintptr(hKey),
			uintptr(uint32(i)),
			uintptr(keyNameBuf.UnsafePtr()),
			uintptr(unsafe.Pointer(&szKeyNameBuf)),
			0, 0, 0, 0)

		if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
			return nil, wErr
		}
		keys = append(keys, wstr.WstrSliceToStr(keyNameBuf.HotSlice()))
	}

	return keys, nil
}

// [RegEnumValue] function.
//
// # Example
//
//	hKey, _ := win.HKEY_CURRENT_USER.RegOpenKeyEx(
//		"Control Panel\\Keyboard",
//		co.REG_OPTION_NONE,
//		co.KEY_READ)
//	defer hKey.RegCloseKey()
//
//	names, _ := hKey.RegEnumValue()
//	for name, _ := range names {
//		println(name)
//	}
//
// [RegEnumValue]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumvaluew
func (hKey HKEY) RegEnumValue() ([]string, error) {
	nfo, err := hKey.RegQueryInfoKey()
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, nfo.NumValues)
	valueNameBuf := wstr.NewBufSized[wstr.Stack64](nfo.MaxValueNameLen + 1)

	for i := uint(0); i < nfo.NumValues; i++ {
		szValueNameBuf := uint32(valueNameBuf.Len())

		ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegEnumValueW),
			uintptr(hKey),
			uintptr(uint32(i)),
			uintptr(valueNameBuf.UnsafePtr()),
			uintptr(unsafe.Pointer(&szValueNameBuf)),
			0, 0, 0, 0)

		if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
			return nil, wErr
		}
		names = append(names, wstr.WstrSliceToStr(valueNameBuf.HotSlice()))
	}

	return names, nil
}

// [RegFlushKey] function.
//
// [RegFlushKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regflushkey
func (hKey HKEY) RegFlushKey() error {
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegFlushKey),
		uintptr(hKey))
	return utl.ZeroAsSysError(ret)
}

// [RegGetValue] function.
//
// # Example
//
//	regVal, _ := win.HKEY_CURRENT_USER.RegGetValue(
//		"Control Panel\\Mouse",
//		"Beep",
//		co.RRF_RT_ANY,
//	)
//
//	str, _ := regVal.Sz()
//	println(str)
//
// [RegGetValue]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-reggetvaluew
func (hKey HKEY) RegGetValue(subKey, valueName string, flags co.RRF) (RegVal, error) {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.EMPTY_IS_NIL)
	valueName16 := wstr.NewBufWith[wstr.Stack20](valueName, wstr.EMPTY_IS_NIL)

	var dataBuf []byte // will be copied straight into RegVal

	for {
		var szDataBytes uint32

		ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegGetValueW), // 1st call to retrieve size only
			uintptr(hKey),
			uintptr(subKey16.UnsafePtr()),
			uintptr(valueName16.UnsafePtr()),
			uintptr(flags),
			0, 0,
			uintptr(unsafe.Pointer(&szDataBytes)))

		if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
			return RegVal{}, wErr
		}

		dataBuf = make([]byte, szDataBytes)
		var dataType uint32

		ret, _, _ = syscall.SyscallN(dll.Advapi(dll.PROC_RegGetValueW), // 2nd call to retrieve the data
			uintptr(hKey),
			uintptr(subKey16.UnsafePtr()),
			uintptr(valueName16.UnsafePtr()),
			uintptr(flags),
			uintptr(unsafe.Pointer(&dataType)),
			uintptr(unsafe.Pointer(&dataBuf[0])),
			uintptr(unsafe.Pointer(&szDataBytes)))

		if wErr := co.ERROR(ret); wErr == co.ERROR_SUCCESS {
			dataBuf = dataBuf[:szDataBytes] // data length may have shrunk
			return regValParse(dataBuf, co.REG(dataType))
		} else if wErr == co.ERROR_MORE_DATA {
			continue // value changed in a concurrent operation; retry
		} else {
			return RegVal{}, wErr
		}
	}
}

// [RegLoadKey] function.
//
// [RegLoadKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regloadkeyw
func (hKey HKEY) RegLoadKey(subKey, file string) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.EMPTY_IS_NIL)
	file16 := wstr.NewBufWith[wstr.Stack20](file, wstr.ALLOW_EMPTY)

	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegLoadKeyW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()),
		uintptr(file16.UnsafePtr()))
	return utl.ZeroAsSysError(ret)
}

// [RegOpenKeyEx] function.
//
// ⚠️ You must defer [HKEY.RegCloseKey].
//
// # Example
//
//	hKey, _ := win.HKEY_CURRENT_USER.RegOpenKeyEx(
//		"Control Panel\\Keyboard",
//		co.REG_OPTION_NONE,
//		co.KEY_READ)
//	defer hKey.RegCloseKey()
//
// [RegOpenKeyEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regopenkeyexw
func (hKey HKEY) RegOpenKeyEx(
	subKey string,
	options co.REG_OPTION,
	accessRights co.KEY,
) (HKEY, error) {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.EMPTY_IS_NIL)
	var openedKey HKEY

	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegOpenKeyExW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()),
		uintptr(options),
		uintptr(accessRights),
		uintptr(unsafe.Pointer(&openedKey)))

	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

// [RegQueryInfoKey] function.
//
// [RegQueryInfoKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryinfokeyw
func (hKey HKEY) RegQueryInfoKey() (HkeyInfo, error) {
	var (
		classBuf              = wstr.NewBufSized[wstr.Stack64](64)
		numSubKeys            uint32
		maxSubKeyNameLen      uint32
		maxClassNameLen       uint32
		numValues             uint32
		maxValueNameLen       uint32
		maxValueDataLen       uint32
		securityDescriptorLen uint32
		ft                    FILETIME
	)

	for {
		szClassBuf := uint32(classBuf.Len())

		ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegQueryInfoKeyW),
			uintptr(hKey),
			uintptr(classBuf.UnsafePtr()),
			uintptr(unsafe.Pointer(&szClassBuf)),
			0,
			uintptr(unsafe.Pointer(&numSubKeys)),
			uintptr(unsafe.Pointer(&maxSubKeyNameLen)),
			uintptr(unsafe.Pointer(&maxClassNameLen)),
			uintptr(unsafe.Pointer(&numValues)),
			uintptr(unsafe.Pointer(&maxValueNameLen)),
			uintptr(unsafe.Pointer(&maxValueDataLen)),
			uintptr(unsafe.Pointer(&securityDescriptorLen)),
			uintptr(unsafe.Pointer(&ft)))

		if wErr := co.ERROR(ret); wErr == co.ERROR_MORE_DATA {
			classBuf.Resize(classBuf.Len() + 64) // increase buffer size to try again
		} else if wErr != co.ERROR_SUCCESS {
			return HkeyInfo{}, wErr
		} else {
			break
		}
	}

	return HkeyInfo{
		Class:                 wstr.WstrSliceToStr(classBuf.HotSlice()),
		NumSubKeys:            uint(numSubKeys),
		MaxSubKeyNameLen:      uint(maxSubKeyNameLen),
		MaxClassNameLen:       uint(maxClassNameLen),
		NumValues:             uint(numValues),
		MaxValueNameLen:       uint(maxValueNameLen),
		MaxValueDataLen:       uint(maxValueDataLen),
		SecurityDescriptorLen: uint(securityDescriptorLen),
		LastWriteTime:         ft.ToTime(),
	}, nil

}

// Returned by [HKEY.RegQueryInfoKey].
type HkeyInfo struct {
	Class                 string
	NumSubKeys            uint
	MaxSubKeyNameLen      uint
	MaxClassNameLen       uint
	NumValues             uint
	MaxValueNameLen       uint
	MaxValueDataLen       uint
	SecurityDescriptorLen uint
	LastWriteTime         time.Time
}

// [RegQueryMultipleValues] function.
//
// # Example
//
//	hKey, _ := win.HKEY_CURRENT_USER.RegOpenKeyEx(
//		"Control Panel\\Desktop", co.REG_OPTION_NONE, co.KEY_READ)
//	defer hKey.RegCloseKey()
//
//	regVals, _ := hKey.RegQueryMultipleValues("DragWidth", "WallPaper")
//	for _, regVal := range regVals {
//		str, _ := regVal.Sz()
//		println(str)
//	}
//
// [RegQueryMultipleValues]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regquerymultiplevaluesw
func (hKey HKEY) RegQueryMultipleValues(valueNames ...string) ([]RegVal, error) {
	valueNames16 := wstr.NewArray(valueNames...)
	var dataBuf []byte

	valents := make([]_VALENT, 0, len(valueNames))
	for idx := range valueNames {
		valents = append(valents, _VALENT{
			ValueName: valueNames16.PtrOf(uint(idx)),
		})
	}

	for {
		var szDataBytes uint32
		ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegQueryMultipleValuesW), // 1st call to retrieve size only
			uintptr(hKey),
			uintptr(unsafe.Pointer(&valents[0])),
			uintptr(uint32(len(valueNames))),
			0,
			uintptr(unsafe.Pointer(&szDataBytes)))

		if wErr := co.ERROR(ret); wErr != co.ERROR_MORE_DATA {
			return nil, wErr
		}

		dataBuf = make([]byte, szDataBytes)
		ret, _, _ = syscall.SyscallN(dll.Advapi(dll.PROC_RegQueryMultipleValuesW), // 2nd call to retrieve the data
			uintptr(hKey),
			uintptr(unsafe.Pointer(&valents[0])),
			uintptr(uint32(len(valueNames))),
			uintptr(unsafe.Pointer(&dataBuf[0])),
			uintptr(unsafe.Pointer(&szDataBytes)))

		if wErr := co.ERROR(ret); wErr == co.ERROR_SUCCESS {
			dataBuf = dataBuf[:szDataBytes] // data length may have shrunk
			regVals := make([]RegVal, 0, len(valents))
			for _, valent := range valents {
				regVal, err := regValParse(valent.bufProjection(dataBuf), valent.Type)
				if err != nil {
					return nil, err
				}
				regVals = append(regVals, regVal)
			}
			return regVals, nil
		} else if wErr == co.ERROR_MORE_DATA {
			continue // value changed in a concurrent operation; retry
		} else {
			return nil, wErr
		}
	}
}

// [RegQueryReflectionKey] function.
//
// [RegQueryReflectionKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryreflectionkey
func (hKey HKEY) RegQueryReflectionKey() (bool, error) {
	var bVal int32 // BOOL
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegQueryReflectionKey),
		uintptr(hKey),
		uintptr(unsafe.Pointer(&bVal)))
	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return false, wErr
	}
	return bVal != 0, nil
}

// [RegQueryValueEx] function.
//
// # Example
//
//	hKey, _ := win.HKEY_CURRENT_USER.RegOpenKeyEx(
//		"Control Panel\\Mouse",
//		co.REG_OPTION_NONE,
//		co.KEY_READ)
//	defer hKey.RegCloseKey()
//
//	regVal, _ := hKey.RegQueryValueEx("Beep")
//
//	str, _ := regVal.Sz()
//	println(str)
//
// [RegQueryValueEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryvalueexw
func (hKey HKEY) RegQueryValueEx(valueName string) (RegVal, error) {
	valueName16 := wstr.NewBufWith[wstr.Stack20](valueName, wstr.EMPTY_IS_NIL)
	var dataBuf []byte // will be copied straight into RegVal

	for {
		var szDataBytes uint32
		ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegQueryValueExW), // 1st call to retrieve size only
			uintptr(hKey),
			uintptr(valueName16.UnsafePtr()),
			0, 0, 0,
			uintptr(unsafe.Pointer(&szDataBytes)))

		if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
			return RegVal{}, wErr
		}

		dataBuf = make([]byte, szDataBytes)
		var dataType uint32

		ret, _, _ = syscall.SyscallN(dll.Advapi(dll.PROC_RegQueryValueExW), // 2nd call to retrieve the data
			uintptr(hKey),
			uintptr(valueName16.UnsafePtr()),
			0,
			uintptr(unsafe.Pointer(&dataType)),
			uintptr(unsafe.Pointer(&dataBuf[0])),
			uintptr(unsafe.Pointer(&szDataBytes)))

		if wErr := co.ERROR(ret); wErr == co.ERROR_SUCCESS {
			dataBuf = dataBuf[:szDataBytes] // data length may have shrunk
			return regValParse(dataBuf, co.REG(dataType))
		} else if wErr == co.ERROR_MORE_DATA {
			continue // value changed in a concurrent operation; retry
		} else {
			return RegVal{}, wErr
		}
	}
}

// [RegRenameKey] function.
//
// [RegRenameKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regrenamekey
func (hKey HKEY) RegRenameKey(subKey, newName string) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.ALLOW_EMPTY)
	newName16 := wstr.NewBufWith[wstr.Stack20](newName, wstr.EMPTY_IS_NIL)

	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegRenameKey),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()),
		uintptr(newName16.UnsafePtr()))
	return utl.ZeroAsSysError(ret)
}

// [RegReplaceKey] function
//
// [RegReplaceKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regreplacekeyw
func (hKey HKEY) RegReplaceKey(subKey, srcFile, destBackupFile string) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.EMPTY_IS_NIL)
	srcFile16 := wstr.NewBufWith[wstr.Stack20](srcFile, wstr.EMPTY_IS_NIL)
	destBackupFile16 := wstr.NewBufWith[wstr.Stack20](destBackupFile, wstr.EMPTY_IS_NIL)

	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegReplaceKeyW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()),
		uintptr(srcFile16.UnsafePtr()),
		uintptr(destBackupFile16.UnsafePtr()))
	return utl.ZeroAsSysError(ret)
}

// [RegRestoreKey] function.
//
// Paired with [HKEY.RegSaveKey] or [HKEY.RegSaveKeyEx].
//
// [RegRestoreKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regrestorekeyw
func (hKey HKEY) RegRestoreKey(srcFile string, flags co.REG_RESTORE) error {
	srcFile16 := wstr.NewBufWith[wstr.Stack20](srcFile, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegRestoreKeyW),
		uintptr(hKey),
		uintptr(srcFile16.UnsafePtr()),
		uintptr(flags))
	return utl.ZeroAsSysError(ret)
}

// [RegSaveKey] function.
//
// Paired with [HKEY.RegRestoreKey].
//
// [RegSaveKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsavekeyw
func (hKey HKEY) RegSaveKey(destFile string, securityAttributes *SECURITY_ATTRIBUTES) error {
	destFile16 := wstr.NewBufWith[wstr.Stack20](destFile, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegSaveKeyW),
		uintptr(hKey),
		uintptr(destFile16.UnsafePtr()),
		uintptr(unsafe.Pointer(securityAttributes)))
	return utl.ZeroAsSysError(ret)
}

// [RegSaveKeyEx] function.
//
// Paired with [HKEY.RegRestoreKey].
//
// [RegSaveKeyEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsavekeyexw
func (hKey HKEY) RegSaveKeyEx(
	destFile string,
	securityAttributes *SECURITY_ATTRIBUTES,
	flags co.REG_SAVE,
) error {
	destFile16 := wstr.NewBufWith[wstr.Stack20](destFile, wstr.ALLOW_EMPTY)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegSaveKeyExW),
		uintptr(hKey),
		uintptr(destFile16.UnsafePtr()),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(flags))
	return utl.ZeroAsSysError(ret)
}

// [RegSetKeyValue] function.
//
// [RegSetKeyValue]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsetkeyvaluew
func (hKey HKEY) RegSetKeyValue(subKey, valueName string, data RegVal) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.ALLOW_EMPTY)
	valueName16 := wstr.NewBufWith[wstr.Stack20](valueName, wstr.EMPTY_IS_NIL)

	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegSetKeyValueW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()),
		uintptr(valueName16.UnsafePtr()),
		uintptr(data.Type()),
		uintptr(unsafe.Pointer(&data.data[0])),
		uintptr(uint32(len(data.data))))
	return utl.ZeroAsSysError(ret)
}

// [RegSetValueEx] function.
//
// [RegSetValueEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsetvalueexw
func (hKey HKEY) RegSetValueEx(valueName string, data RegVal) error {
	valueName16 := wstr.NewBufWith[wstr.Stack20](valueName, wstr.EMPTY_IS_NIL)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegSetValueExW),
		uintptr(hKey),
		uintptr(valueName16.UnsafePtr()),
		0,
		uintptr(data.Type()),
		uintptr(unsafe.Pointer(&data.data[0])),
		uintptr(uint32(len(data.data))))
	return utl.ZeroAsSysError(ret)
}

// [RegUnLoadKey] function.
//
// [RegUnLoadKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regunloadkeyw
func (hKey HKEY) RegUnLoadKey(subKey string) error {
	subKey16 := wstr.NewBufWith[wstr.Stack20](subKey, wstr.EMPTY_IS_NIL)
	ret, _, _ := syscall.SyscallN(dll.Advapi(dll.PROC_RegUnLoadKeyW),
		uintptr(hKey),
		uintptr(subKey16.UnsafePtr()))
	return utl.ZeroAsSysError(ret)
}
