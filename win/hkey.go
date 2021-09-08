package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a registry key.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hkey
type HKEY HANDLE

// Predefined registry key.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/sysinfo/predefined-keys
const (
	HKEY_CLASSES_ROOT        HKEY = 0x80000000
	HKEY_CURRENT_USER        HKEY = 0x80000001
	HKEY_LOCAL_MACHINE       HKEY = 0x80000002
	HKEY_USERS               HKEY = 0x80000003
	HKEY_PERFORMANCE_DATA    HKEY = 0x80000004
	HKEY_PERFORMANCE_TEXT    HKEY = 0x80000050
	HKEY_PERFORMANCE_NLSTEXT HKEY = 0x80000060
	HKEY_CURRENT_CONFIG      HKEY = 0x80000005
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func (hKey HKEY) CloseKey() error {
	ret, _, _ := syscall.Syscall(proc.RegCloseKey.Addr(), 1,
		uintptr(hKey), 0, 0)
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyw
func (hKey HKEY) DeleteKey(subKey string) error {
	ret, _, _ := syscall.Syscall(proc.RegDeleteKey.Addr(), 2,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(subKey))), 0)

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// samDesired must be KEY_WOW64_32KEY or KEY_WOW64_64KEY.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyexw
func (hKey HKEY) DeleteKeyEx(subKey string, samDesired co.KEY) error {
	ret, _, _ := syscall.Syscall6(proc.RegDeleteKeyEx.Addr(), 4,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(subKey))),
		uintptr(samDesired), 0, 0, 0)
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyvaluew
func (hKey HKEY) DeleteKeyValue(subKey, valueName string) error {
	ret, _, _ := syscall.Syscall(proc.RegDeleteKeyValue.Addr(), 3,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(subKey))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(valueName))))
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletetreew
func (hKey HKEY) DeleteTree(subKey string) error {
	ret, _, _ := syscall.Syscall(proc.RegDeleteTree.Addr(), 2,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(subKey))), 0)
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumkeyexw
func (hKey HKEY) EnumKeyEx() ([]string, error) {
	cSubKeys, cbMaxSubKeyLen := uint32(0), uint32(0)
	err := hKey.QueryInfoKey(nil, &cSubKeys, &cbMaxSubKeyLen,
		nil, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	err = nil
	keyNames := make([]string, 0, cSubKeys)

	keyNameBuf := make([]uint16, cbMaxSubKeyLen+1)
	keyNameBufLen := int(0)

	for i := 0; i < int(cSubKeys); i++ {
		keyNameBufLen = len(keyNameBuf)

		ret, _, _ := syscall.Syscall9(proc.RegEnumKeyEx.Addr(), 8,
			uintptr(hKey), uintptr(i),
			uintptr(unsafe.Pointer(&keyNameBuf[0])),
			uintptr(unsafe.Pointer(&keyNameBufLen)),
			0, 0, 0, 0, 0)

		if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
			return nil, wErr
		}

		keyNames = append(keyNames, Str.FromUint16Slice(keyNameBuf))
	}

	return keyNames, nil
}

// Returned valueNames and valueTypes have the same length.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumvaluew
func (hKey HKEY) EnumValue() (
	valueNames []string, valueTypes []co.REG, err error) {

	cValues, cbMaxValueNameLen := uint32(0), uint32(0)
	err = hKey.QueryInfoKey(nil, nil, nil, nil, &cValues, &cbMaxValueNameLen,
		nil, nil, nil)
	if err != nil {
		return
	}

	err = nil
	valueNames = make([]string, 0, cValues)
	valueTypes = make([]co.REG, 0, cValues)

	valueNameBuf := make([]uint16, cbMaxValueNameLen+2)
	valueNameBufLen := int(0)
	valueTypeBuf := co.REG(0)

	for i := 0; i < int(cValues); i++ {
		valueNameBufLen = len(valueNameBuf)

		ret, _, _ := syscall.Syscall9(proc.RegEnumValue.Addr(), 8,
			uintptr(hKey), uintptr(i),
			uintptr(unsafe.Pointer(&valueNameBuf[0])),
			uintptr(unsafe.Pointer(&valueNameBufLen)),
			0, uintptr(unsafe.Pointer(&valueTypeBuf)), 0, 0, 0)

		if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
			valueNames, valueTypes, err = nil, nil, wErr
			return
		}

		valueNames = append(valueNames, Str.FromUint16Slice(valueNameBuf))
		valueTypes = append(valueTypes, valueTypeBuf)
	}

	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regflushkey
func (hKey HKEY) FlushKey() error {
	ret, _, _ := syscall.Syscall(proc.RegFlushKey.Addr(), 1,
		uintptr(hKey), 0, 0)

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// This function is rather tricky. Prefer using HKEY.ReadBinary(),
// HKEY.ReadDword(), HKEY.ReadQword() or HKEY.ReadString().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-reggetvaluew
func (hKey HKEY) GetValue(
	subKey, value string, flags co.RRF, pdwType *co.REG,
	pData unsafe.Pointer, pDataLen *uint32) error {

	ret, _, _ := syscall.Syscall9(proc.RegGetValue.Addr(), 7,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(subKey))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(value))),
		uintptr(flags), uintptr(unsafe.Pointer(pdwType)),
		uintptr(pData), uintptr(unsafe.Pointer(pDataLen)), 0, 0)

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// ⚠️ You must defer HKEY.CloseKey().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regopenkeyexw
func (hKey HKEY) OpenKeyEx(
	subKey string, ulOptions co.REG_OPTION, samDesired co.KEY) (HKEY, error) {

	openedKey := HKEY(0)
	ret, _, _ := syscall.Syscall6(proc.RegOpenKeyEx.Addr(), 5,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(subKey))),
		uintptr(ulOptions), uintptr(samDesired),
		uintptr(unsafe.Pointer(&openedKey)), 0)

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

// Pass pointers for the values you want to retrieve, pass the others as nil.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryinfokeyw
func (hKey HKEY) QueryInfoKey(
	lpClass *string,
	cSubKeys, cbMaxSubKeyLen, cbMaxClassLen, cValues,
	cbMaxValueNameLen, cbMaxValueLen, cbSecurityDescriptor *uint32,
	ftLastWriteTime *FILETIME) error {

	classBuf := []uint16{}
	cchClassBuf := uint32(0)

	var ( // all retrievable values, nil by default
		classP                *uint16
		cchClassP             *uint32
		cSubKeysP             *uint32
		cbMaxSubKeyLenP       *uint32
		cbMaxClassLenP        *uint32
		cValuesP              *uint32
		cbMaxValueNameLenP    *uint32
		cbMaxValueLenP        *uint32
		cbSecurityDescriptorP *uint32
		ftLastWriteTimeP      *FILETIME
	)
	if lpClass != nil {
		classBuf = make([]uint16, 255+1) // arbitrary
		classP = &classBuf[0]
		cchClassBuf = uint32(len(classBuf))
		cchClassP = &cchClassBuf
	}
	if cSubKeys != nil {
		cSubKeysP = cSubKeys
	}
	if cbMaxSubKeyLen != nil {
		cbMaxSubKeyLenP = cbMaxSubKeyLen
	}
	if cbMaxClassLen != nil {
		cbMaxClassLenP = cbMaxClassLen
	}
	if cValues != nil {
		cValuesP = cValues
	}
	if cbMaxValueNameLen != nil {
		cbMaxValueNameLenP = cbMaxValueNameLen
	}
	if cbMaxValueLen != nil {
		cbMaxValueLenP = cbMaxValueLen
	}
	if cbSecurityDescriptor != nil {
		cbSecurityDescriptorP = cbSecurityDescriptor
	}
	if ftLastWriteTime != nil {
		ftLastWriteTimeP = ftLastWriteTime
	}

	ret, _, _ := syscall.Syscall12(proc.RegQueryInfoKey.Addr(), 12,
		uintptr(hKey),
		uintptr(unsafe.Pointer(classP)), uintptr(unsafe.Pointer(cchClassP)), 0,
		uintptr(unsafe.Pointer(cSubKeysP)), uintptr(unsafe.Pointer(cbMaxSubKeyLenP)),
		uintptr(unsafe.Pointer(cbMaxClassLenP)), uintptr(unsafe.Pointer(cValuesP)),
		uintptr(unsafe.Pointer(cbMaxValueNameLenP)), uintptr(unsafe.Pointer(cbMaxValueLenP)),
		uintptr(unsafe.Pointer(cbSecurityDescriptorP)),
		uintptr(unsafe.Pointer(ftLastWriteTimeP)))

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}

	if lpClass != nil {
		*lpClass = Str.FromUint16Slice(classBuf[:])
	}
	return nil
}

// Reads a REG_BINARY value with HKEY.GetValue().
func (hKey HKEY) ReadBinary(subKey, value string) []byte {
	pDataLen := uint32(0)
	pdwType := co.REG_BINARY

	err := hKey.GetValue(subKey, value, co.RRF_RT_REG_BINARY, // retrieve length
		&pdwType, nil, &pDataLen)
	if err != nil {
		panic(err)
	}

	pData := make([]byte, pDataLen)

	err = hKey.GetValue(subKey, value, co.RRF_RT_REG_SZ, // retrieve string
		&pdwType, unsafe.Pointer(&pData[0]), &pDataLen)
	if err != nil {
		panic(err)
	}

	return pData
}

// Reads a REG_DWORD value with HKEY.GetValue().
func (hKey HKEY) ReadDword(subKey, value string) uint32 {
	pData := uint32(0)
	pDataLen := uint32(unsafe.Sizeof(pData))
	pdwType := co.REG_DWORD

	err := hKey.GetValue(subKey, value, co.RRF_RT_REG_DWORD,
		&pdwType, unsafe.Pointer(&pData), &pDataLen)
	if err != nil {
		panic(err)
	}
	return pData
}

// Reads a REG_QWORD value with HKEY.GetValue().
func (hKey HKEY) ReadQword(subKey, value string) uint64 {
	pData := uint64(0)
	pDataLen := uint32(unsafe.Sizeof(pData))
	pdwType := co.REG_QWORD

	err := hKey.GetValue(subKey, value, co.RRF_RT_REG_QWORD,
		&pdwType, unsafe.Pointer(&pData), &pDataLen)
	if err != nil {
		panic(err)
	}
	return pData
}

// Reads a REG_SZ value with HKEY.GetValue().
func (hKey HKEY) ReadString(subKey, value string) string {
	pDataLen := uint32(0)
	pdwType := co.REG_SZ

	err := hKey.GetValue(subKey, value, co.RRF_RT_REG_SZ, // retrieve length
		&pdwType, nil, &pDataLen)
	if err != nil {
		panic(err)
	}

	pData := make([]uint16, pDataLen/2) // pcbData is in bytes; terminating null included

	err = hKey.GetValue(subKey, value, co.RRF_RT_REG_SZ, // retrieve string
		&pdwType, unsafe.Pointer(&pData[0]), &pDataLen)
	if err != nil {
		panic(err)
	}

	return Str.FromUint16Slice(pData)
}

// This function is rather tricky. Prefer using HKEY.WriteBinary(),
// HKEY.WriteDword(), HKEY.WriteQword() or HKEY.WriteString().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsetkeyvaluew
func (hKey HKEY) SetKeyValue(
	subKey, valueName string, dwType co.REG,
	pData unsafe.Pointer, dataLen uint32) error {

	ret, _, _ := syscall.Syscall6(proc.RegSetKeyValue.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(subKey))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(valueName))),
		uintptr(dwType), uintptr(pData), uintptr(dataLen))

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// Writes a REG_BINARY value with HKEY.SetKeyValue().
func (hKey HKEY) WriteBinary(subKey, valueName string, data []byte) {
	err := hKey.SetKeyValue(subKey, valueName, co.REG_BINARY,
		unsafe.Pointer(&data[0]), uint32(len(data)))
	if err != nil {
		panic(err)
	}
}

// Writes a REG_DWORD value with HKEY.SetKeyValue().
func (hKey HKEY) WriteDword(subKey, valueName string, data uint32) {
	err := hKey.SetKeyValue(subKey, valueName, co.REG_DWORD,
		unsafe.Pointer(&data), uint32(unsafe.Sizeof(data)))
	if err != nil {
		panic(err)
	}
}

// Writes a REG_QWORD value with HKEY.SetKeyValue().
func (hKey HKEY) WriteQword(subKey, valueName string, data uint64) {
	err := hKey.SetKeyValue(subKey, valueName, co.REG_QWORD,
		unsafe.Pointer(&data), uint32(unsafe.Sizeof(data)))
	if err != nil {
		panic(err)
	}
}

// Writes a REG_SZ value with HKEY.SetKeyValue().
func (hKey HKEY) WriteString(subKey, valueName string, data string) {
	lpData16 := Str.ToUint16Slice(data)
	err := hKey.SetKeyValue(subKey, valueName, co.REG_SZ,
		unsafe.Pointer(&lpData16[0]), uint32(len(lpData16)*2)) // pass size in bytes, including terminating null
	runtime.KeepAlive(lpData16)
	if err != nil {
		panic(err)
	}
}
