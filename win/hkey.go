package win

import (
	"runtime"
	"sort"
	"strings"
	"syscall"
	"time"
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
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))), 0)
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
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(samDesired), 0, 0, 0)
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyvaluew
func (hKey HKEY) DeleteKeyValue(subKey, valueName string) error {
	ret, _, _ := syscall.Syscall(proc.RegDeleteKeyValue.Addr(), 3,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(valueName))))
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletetreew
func (hKey HKEY) DeleteTree(subKey string) error {
	ret, _, _ := syscall.Syscall(proc.RegDeleteTree.Addr(), 2,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))), 0)
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumkeyexw
func (hKey HKEY) EnumKeyEx() ([]string, error) {
	keyInfo, err := hKey.QueryInfoKey()
	if err != nil {
		return nil, err
	}

	keyNames := make([]string, 0, keyInfo.NumSubKeys)        // key names to be returned
	keyNameBuf := make([]uint16, keyInfo.MaxSubKeyNameLen+1) // to receive the names of the keys
	var keyNameBufLen uint32

	for i := 0; i < int(keyInfo.NumSubKeys); i++ {
		keyNameBufLen = uint32(len(keyNameBuf)) // reset available buffer size

		ret, _, _ := syscall.Syscall9(proc.RegEnumKeyEx.Addr(), 8,
			uintptr(hKey), uintptr(i),
			uintptr(unsafe.Pointer(&keyNameBuf[0])),
			uintptr(unsafe.Pointer(&keyNameBufLen)), // receives the number of chars without null
			0, 0, 0, 0, 0)

		if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
			return nil, wErr
		}

		keyNames = append(keyNames, Str.FromNativeSlice(keyNameBuf))
	}

	Path.Sort(keyNames)
	return keyNames, nil
}

type _ValueEnum struct {
	Name string
	Type co.REG
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumvaluew
func (hKey HKEY) EnumValue() ([]_ValueEnum, error) {
	keyInfo, err := hKey.QueryInfoKey()
	if err != nil {
		return nil, err
	}

	values := make([]_ValueEnum, 0, keyInfo.NumValues) // to be returned

	valueNameBuf := make([]uint16, keyInfo.MaxValueNameLen+2) // room to avoid "more data" error
	var valueNameBufLen uint32
	var valueTypeBuf co.REG

	for i := 0; i < int(keyInfo.NumValues); i++ {
		valueNameBufLen = uint32(len(valueNameBuf)) // reset available buffer size

		ret, _, _ := syscall.Syscall9(proc.RegEnumValue.Addr(), 8,
			uintptr(hKey), uintptr(i),
			uintptr(unsafe.Pointer(&valueNameBuf[0])),
			uintptr(unsafe.Pointer(&valueNameBufLen)), // receives the number of chars without null
			0, uintptr(unsafe.Pointer(&valueTypeBuf)), 0, 0, 0)

		if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
			return nil, wErr
		}

		values = append(values, struct {
			Name string
			Type co.REG
		}{
			Name: Str.FromNativeSlice(valueNameBuf),
			Type: valueTypeBuf,
		})
	}

	sort.Slice(values, func(a, b int) bool {
		return strings.ToUpper(values[a].Name) < strings.ToUpper(values[b].Name) // case insensitive
	})
	return values, nil
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
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(value))),
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

	var openedKey HKEY
	ret, _, _ := syscall.Syscall6(proc.RegOpenKeyEx.Addr(), 5,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(ulOptions), uintptr(samDesired),
		uintptr(unsafe.Pointer(&openedKey)), 0)

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

type _KeyInfo struct {
	Class                 string
	NumSubKeys            uint32
	MaxSubKeyNameLen      uint32
	MaxSubKeyClassLen     uint32
	NumValues             uint32
	MaxValueNameLen       uint32
	MaxValueDataLen       uint32
	SecurityDescriptorLen uint32
	LastWriteTime         time.Time
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryinfokeyw
func (hKey HKEY) QueryInfoKey() (_KeyInfo, error) {
	info := _KeyInfo{}
	classBuf := [_MAX_PATH + 1]uint16{} // arbitrary
	classBufLen := uint32(len(classBuf))
	ft := FILETIME{}

	ret, _, _ := syscall.Syscall12(proc.RegQueryInfoKey.Addr(), 12,
		uintptr(hKey),
		uintptr(unsafe.Pointer(&classBuf[0])), uintptr(unsafe.Pointer(&classBufLen)), 0,
		uintptr(unsafe.Pointer(&info.NumSubKeys)),
		uintptr(unsafe.Pointer(&info.MaxSubKeyNameLen)),
		uintptr(unsafe.Pointer(&info.MaxSubKeyClassLen)),
		uintptr(unsafe.Pointer(&info.NumValues)),
		uintptr(unsafe.Pointer(&info.MaxValueNameLen)),
		uintptr(unsafe.Pointer(&info.MaxValueDataLen)),
		uintptr(unsafe.Pointer(&info.SecurityDescriptorLen)),
		uintptr(unsafe.Pointer(&ft)))

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return info, wErr
	}

	info.Class = Str.FromNativeSlice(classBuf[:])
	info.LastWriteTime = ft.ToTime()
	return info, nil
}

// Reads a REG_BINARY value with HKEY.GetValue().
func (hKey HKEY) ReadBinary(subKey, value string) []byte {
	var pDataLen uint32
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
	var pData uint32
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
	var pData uint64
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
	var pDataLen uint32
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

	return Str.FromNativeSlice(pData)
}

// This function is rather tricky. Prefer using HKEY.WriteBinary(),
// HKEY.WriteDword(), HKEY.WriteQword() or HKEY.WriteString().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsetkeyvaluew
func (hKey HKEY) SetKeyValue(
	subKey, valueName string, dwType co.REG,
	pData unsafe.Pointer, dataLen uint32) error {

	ret, _, _ := syscall.Syscall6(proc.RegSetKeyValue.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(valueName))),
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
	lpData16 := Str.ToNativeSlice(data)
	err := hKey.SetKeyValue(subKey, valueName, co.REG_SZ,
		unsafe.Pointer(&lpData16[0]), uint32(len(lpData16)*2)) // pass size in bytes, including terminating null
	runtime.KeepAlive(lpData16)
	if err != nil {
		panic(err)
	}
}
