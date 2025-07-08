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

	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pMachineName := wbuf.PtrEmptyIsNil(machineName)

	var openedKey HKEY

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegConnectRegistryW, "RegConnectRegistryW"),
		uintptr(pMachineName),
		uintptr(hKey),
		uintptr(unsafe.Pointer(&openedKey)))

	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

var _RegConnectRegistryW *syscall.Proc

// [RegOpenCurrentUser] function.
//
// ⚠️ You must defer [HKEY.RegCloseKey].
//
// [RegOpenCurrentUser]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regopencurrentuser
func RegOpenCurrentUser(accessRights co.KEY) (HKEY, error) {
	var openedKey HKEY
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegOpenCurrentUser, "RegOpenCurrentUser"),
		uintptr(accessRights),
		uintptr(unsafe.Pointer(&openedKey)))
	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

var _RegOpenCurrentUser *syscall.Proc

// [RegCloseKey] function.
//
// [RegCloseKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func (hKey HKEY) RegCloseKey() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegCloseKey, "RegCloseKey"),
		uintptr(hKey))
	return utl.ZeroAsSysError(ret)
}

var _RegCloseKey *syscall.Proc

// [RegCopyTree] function.
//
// [RegCopyTree]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regcopytreew
func (hKey HKEY) RegCopyTree(subKey string, dest HKEY) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrEmptyIsNil(subKey)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegCopyTreeW, "RegCopyTreeW"),
		uintptr(hKey),
		uintptr(pSubKey))
	return utl.ZeroAsSysError(ret)
}

var _RegCopyTreeW *syscall.Proc

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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrEmptyIsNil(subKey)

	var openedKey HKEY

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegCreateKeyExW, "RegCreateKeyExW"),
		uintptr(hKey),
		uintptr(pSubKey),
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

var _RegCreateKeyExW *syscall.Proc

// [RegDeleteKey] function.
//
// [RegDeleteKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyw
func (hKey HKEY) RegDeleteKey(subKey string) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrAllowEmpty(subKey)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegDeleteKeyW, "RegDeleteKeyW"),
		uintptr(hKey),
		uintptr(pSubKey))
	return utl.ZeroAsSysError(ret)
}

var _RegDeleteKeyW *syscall.Proc

// [RegDeleteKeyEx] function.
//
// samDesired must be [co.KEY_WOW64_32KEY] or [co.KEY_WOW64_64KEY].
//
// [RegDeleteKeyEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyexw
func (hKey HKEY) RegDeleteKeyEx(subKey string, samDesired co.KEY) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrAllowEmpty(subKey)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegDeleteKeyExW, "RegDeleteKeyExW"),
		uintptr(hKey),
		uintptr(pSubKey),
		uintptr(samDesired&(co.KEY_WOW64_32KEY|co.KEY_WOW64_64KEY)),
		0)
	return utl.ZeroAsSysError(ret)
}

var _RegDeleteKeyExW *syscall.Proc

// [RegDeleteKeyValue] function.
//
// [RegDeleteKeyValue]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyvaluew
func (hKey HKEY) RegDeleteKeyValue(subKey, valueName string) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrEmptyIsNil(subKey)
	pValueName := wbuf.PtrEmptyIsNil(valueName)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegDeleteKeyValueW, "RegDeleteKeyValueW"),
		uintptr(hKey),
		uintptr(pSubKey),
		uintptr(pValueName))
	return utl.ZeroAsSysError(ret)
}

var _RegDeleteKeyValueW *syscall.Proc

// [RegDeleteTree] function.
//
// [RegDeleteTree]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletetreew
func (hKey HKEY) RegDeleteTree(subKey string) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrEmptyIsNil(subKey)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegDeleteTreeW, "RegDeleteTreeW"),
		uintptr(hKey),
		uintptr(pSubKey))
	return utl.ZeroAsSysError(ret)
}

var _RegDeleteTreeW *syscall.Proc

// [RegDeleteValue] function.
//
// [RegDeleteValue]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletevaluew
func (hKey HKEY) RegDeleteValue(valueName string) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pValueName := wbuf.PtrEmptyIsNil(valueName)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegDeleteValueW, "RegDeleteValueW"),
		uintptr(hKey),
		uintptr(pValueName))
	return utl.ZeroAsSysError(ret)
}

var _RegDeleteValueW *syscall.Proc

// [RegDisableReflectionKey] function.
//
// [RegDisableReflectionKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdisablereflectionkey
func (hKey HKEY) RegDisableReflectionKey() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegDisableReflectionKey, "RegDisableReflectionKey"),
		uintptr(hKey))
	return utl.ZeroAsSysError(ret)
}

var _RegDisableReflectionKey *syscall.Proc

// [RegEnableReflectionKey] function.
//
// [RegEnableReflectionKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenablereflectionkey
func (hKey HKEY) RegEnableReflectionKey() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegEnableReflectionKey, "RegEnableReflectionKey"),
		uintptr(hKey))
	return utl.ZeroAsSysError(ret)
}

var _RegEnableReflectionKey *syscall.Proc

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

	if nfo.NumSubKeys == 0 {
		return []string{}, nil // no subkeys
	}

	keys := make([]string, 0, nfo.NumSubKeys) // to be returned

	keyNameBuf := wstr.NewBufDecoder(nfo.MaxSubKeyNameLen + 1)
	defer keyNameBuf.Free()

	for i := uint(0); i < nfo.NumSubKeys; i++ {
		szKeyNameBuf := uint32(keyNameBuf.Len())

		ret, _, _ := syscall.SyscallN(
			dll.Load(dll.ADVAPI32, &_RegEnumKeyExW, "RegEnumKeyExW"),
			uintptr(hKey),
			uintptr(uint32(i)),
			uintptr(keyNameBuf.UnsafePtr()),
			uintptr(unsafe.Pointer(&szKeyNameBuf)),
			0, 0, 0, 0)

		if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
			return nil, wErr
		}
		keys = append(keys, keyNameBuf.String())
	}

	return keys, nil
}

var _RegEnumKeyExW *syscall.Proc

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

	if nfo.NumValues == 0 {
		return []string{}, nil // no values
	}

	names := make([]string, 0, nfo.NumValues) // to be returned

	valueNameBuf := wstr.NewBufDecoder(nfo.MaxValueNameLen + 1)
	defer valueNameBuf.Free()

	for i := uint(0); i < nfo.NumValues; i++ {
		szValueNameBuf := uint32(valueNameBuf.Len())

		ret, _, _ := syscall.SyscallN(
			dll.Load(dll.ADVAPI32, &_RegEnumValueW, "RegEnumValueW"),
			uintptr(hKey),
			uintptr(uint32(i)),
			uintptr(valueNameBuf.UnsafePtr()),
			uintptr(unsafe.Pointer(&szValueNameBuf)),
			0, 0, 0, 0)

		if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
			return nil, wErr
		}
		names = append(names, valueNameBuf.String())
	}

	return names, nil
}

var _RegEnumValueW *syscall.Proc

// [RegFlushKey] function.
//
// [RegFlushKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regflushkey
func (hKey HKEY) RegFlushKey() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegFlushKey, "RegFlushKey"),
		uintptr(hKey))
	return utl.ZeroAsSysError(ret)
}

var _RegFlushKey *syscall.Proc

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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrEmptyIsNil(subKey)
	pValueName := wbuf.PtrEmptyIsNil(valueName)

	var dataBuf []byte // will be copied straight into RegVal

	for {
		var szDataBytes uint32

		ret, _, _ := syscall.SyscallN( // 1st call to retrieve size only
			dll.Load(dll.ADVAPI32, &_RegGetValueW, "RegGetValueW"),
			uintptr(hKey),
			uintptr(pSubKey),
			uintptr(pValueName),
			uintptr(flags),
			0, 0,
			uintptr(unsafe.Pointer(&szDataBytes)))

		if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
			return RegVal{}, wErr
		}

		dataBuf = make([]byte, szDataBytes)
		var dataType uint32

		ret, _, _ = syscall.SyscallN( // 2nd call to retrieve the data
			dll.Load(dll.ADVAPI32, &_RegGetValueW, "RegGetValueW"),
			uintptr(hKey),
			uintptr(pSubKey),
			uintptr(pValueName),
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

var _RegGetValueW *syscall.Proc

// [RegLoadKey] function.
//
// [RegLoadKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regloadkeyw
func (hKey HKEY) RegLoadKey(subKey, file string) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrEmptyIsNil(subKey)
	pFile := wbuf.PtrAllowEmpty(file)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegLoadKeyW, "RegLoadKeyW"),
		uintptr(hKey),
		uintptr(pSubKey),
		uintptr(pFile))
	return utl.ZeroAsSysError(ret)
}

var _RegLoadKeyW *syscall.Proc

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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrEmptyIsNil(subKey)

	var openedKey HKEY

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegOpenKeyExW, "RegOpenKeyExW"),
		uintptr(hKey),
		uintptr(pSubKey),
		uintptr(options),
		uintptr(accessRights),
		uintptr(unsafe.Pointer(&openedKey)))

	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

var _RegOpenKeyExW *syscall.Proc

// [RegQueryInfoKey] function.
//
// [RegQueryInfoKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryinfokeyw
func (hKey HKEY) RegQueryInfoKey() (HkeyInfo, error) {
	classBuf := wstr.NewBufDecoder(64)
	defer classBuf.Free()

	var (
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

		ret, _, _ := syscall.SyscallN(
			dll.Load(dll.ADVAPI32, &_RegQueryInfoKeyW, "RegQueryInfoKeyW"),
			uintptr(hKey),
			uintptr(classBuf.Len()),
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
		Class:                 classBuf.String(),
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

var _RegQueryInfoKeyW *syscall.Proc

// Returned by [HKEY.RegQueryInfoKey].
type HkeyInfo struct {
	Class                 string // User-defined class of the key.
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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()

	valents := make([]_VALENT, 0, len(valueNames))
	for _, valueName := range valueNames {
		valents = append(valents, _VALENT{
			ValueName: (*uint16)(wbuf.PtrAllowEmpty(valueName)),
		})
	}

	for {
		var szDataBytes uint32
		ret, _, _ := syscall.SyscallN( // 1st call to retrieve size only
			dll.Load(dll.ADVAPI32, &_RegQueryMultipleValuesW, "RegQueryMultipleValuesW"),
			uintptr(hKey),
			uintptr(unsafe.Pointer(&valents[0])),
			uintptr(uint32(len(valueNames))),
			0,
			uintptr(unsafe.Pointer(&szDataBytes)))

		if wErr := co.ERROR(ret); wErr != co.ERROR_MORE_DATA {
			return nil, wErr
		}

		dataBuf := make([]byte, szDataBytes)
		ret, _, _ = syscall.SyscallN( // 2nd call to retrieve the data
			dll.Load(dll.ADVAPI32, &_RegQueryMultipleValuesW, "RegQueryMultipleValuesW"),
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

var _RegQueryMultipleValuesW *syscall.Proc

// [RegQueryReflectionKey] function.
//
// [RegQueryReflectionKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryreflectionkey
func (hKey HKEY) RegQueryReflectionKey() (bool, error) {
	var bVal int32 // BOOL
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegQueryReflectionKey, "RegQueryReflectionKey"),
		uintptr(hKey),
		uintptr(unsafe.Pointer(&bVal)))
	if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
		return false, wErr
	}
	return bVal != 0, nil
}

var _RegQueryReflectionKey *syscall.Proc

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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pValueName := wbuf.PtrEmptyIsNil(valueName)

	for {
		var szDataBytes uint32
		ret, _, _ := syscall.SyscallN( // 1st call to retrieve size only
			dll.Load(dll.ADVAPI32, &_RegQueryValueExW, "RegQueryValueExW"),
			uintptr(hKey),
			uintptr(pValueName),
			0, 0, 0,
			uintptr(unsafe.Pointer(&szDataBytes)))

		if wErr := co.ERROR(ret); wErr != co.ERROR_SUCCESS {
			return RegVal{}, wErr
		}

		dataBuf := make([]byte, szDataBytes) // will be passed straight into RegVal
		var dataType uint32

		ret, _, _ = syscall.SyscallN( // 2nd call to retrieve the data
			dll.Load(dll.ADVAPI32, &_RegQueryValueExW, "RegQueryValueExW"),
			uintptr(hKey),
			uintptr(pValueName),
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

var _RegQueryValueExW *syscall.Proc

// [RegRenameKey] function.
//
// [RegRenameKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regrenamekey
func (hKey HKEY) RegRenameKey(subKey, newName string) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrAllowEmpty(subKey)
	pNewName := wbuf.PtrEmptyIsNil(newName)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegRenameKey, "RegRenameKey"),
		uintptr(hKey),
		uintptr(pSubKey),
		uintptr(pNewName))
	return utl.ZeroAsSysError(ret)
}

var _RegRenameKey *syscall.Proc

// [RegReplaceKey] function
//
// [RegReplaceKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regreplacekeyw
func (hKey HKEY) RegReplaceKey(subKey, srcFile, destBackupFile string) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrEmptyIsNil(subKey)
	pSrcFile := wbuf.PtrEmptyIsNil(srcFile)
	pDestBackupFile := wbuf.PtrEmptyIsNil(destBackupFile)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegReplaceKeyW, "RegReplaceKeyW"),
		uintptr(hKey),
		uintptr(pSubKey),
		uintptr(pSrcFile),
		uintptr(pDestBackupFile))
	return utl.ZeroAsSysError(ret)
}

var _RegReplaceKeyW *syscall.Proc

// [RegRestoreKey] function.
//
// Paired with [HKEY.RegSaveKey] or [HKEY.RegSaveKeyEx].
//
// [RegRestoreKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regrestorekeyw
func (hKey HKEY) RegRestoreKey(srcFile string, flags co.REG_RESTORE) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSrcFile := wbuf.PtrAllowEmpty(srcFile)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegRestoreKeyW, "RegRestoreKeyW"),
		uintptr(hKey),
		uintptr(pSrcFile),
		uintptr(flags))
	return utl.ZeroAsSysError(ret)
}

var _RegRestoreKeyW *syscall.Proc

// [RegSaveKey] function.
//
// Paired with [HKEY.RegRestoreKey].
//
// [RegSaveKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsavekeyw
func (hKey HKEY) RegSaveKey(destFile string, securityAttributes *SECURITY_ATTRIBUTES) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pDestFile := wbuf.PtrAllowEmpty(destFile)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegSaveKeyW, "RegSaveKeyW"),
		uintptr(hKey),
		uintptr(pDestFile),
		uintptr(unsafe.Pointer(securityAttributes)))
	return utl.ZeroAsSysError(ret)
}

var _RegSaveKeyW *syscall.Proc

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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pDestFile := wbuf.PtrAllowEmpty(destFile)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegSaveKeyExW, "RegSaveKeyExW"),
		uintptr(hKey),
		uintptr(pDestFile),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(flags))
	return utl.ZeroAsSysError(ret)
}

var _RegSaveKeyExW *syscall.Proc

// [RegSetKeyValue] function.
//
// [RegSetKeyValue]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsetkeyvaluew
func (hKey HKEY) RegSetKeyValue(subKey, valueName string, data RegVal) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrAllowEmpty(subKey)
	pValueName := wbuf.PtrEmptyIsNil(valueName)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegSetKeyValueW, "RegSetKeyValueW"),
		uintptr(hKey),
		uintptr(pSubKey),
		uintptr(pValueName),
		uintptr(data.Type()),
		uintptr(unsafe.Pointer(&data.data[0])),
		uintptr(uint32(len(data.data))))
	return utl.ZeroAsSysError(ret)
}

var _RegSetKeyValueW *syscall.Proc

// [RegSetValueEx] function.
//
// [RegSetValueEx]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsetvalueexw
func (hKey HKEY) RegSetValueEx(valueName string, data RegVal) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pValueName := wbuf.PtrEmptyIsNil(valueName)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegSetValueExW, "RegSetValueExW"),
		uintptr(hKey),
		uintptr(pValueName),
		0,
		uintptr(data.Type()),
		uintptr(unsafe.Pointer(&data.data[0])),
		uintptr(uint32(len(data.data))))
	return utl.ZeroAsSysError(ret)
}

var _RegSetValueExW *syscall.Proc

// [RegUnLoadKey] function.
//
// [RegUnLoadKey]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regunloadkeyw
func (hKey HKEY) RegUnLoadKey(subKey string) error {
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pSubKey := wbuf.PtrEmptyIsNil(subKey)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.ADVAPI32, &_RegUnLoadKeyW, "RegUnLoadKeyW"),
		uintptr(hKey),
		uintptr(pSubKey))
	return utl.ZeroAsSysError(ret)
}

var _RegUnLoadKeyW *syscall.Proc
