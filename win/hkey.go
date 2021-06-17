package win

import (
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

// Reads a REG_BINARY value with RegGetValue().
func (hKey HKEY) ReadBinary(lpSubKey, lpValue string) []byte {
	pdwType := co.REG_BINARY
	pcbData := uint32(0)

	err := hKey.RegGetValue(lpSubKey, lpValue, co.RRF_RT_REG_BINARY, // retrieve length
		&pdwType, nil, &pcbData)
	if err != nil {
		panic(err)
	}

	pvData := make([]byte, pcbData)

	err = hKey.RegGetValue(lpSubKey, lpValue, co.RRF_RT_REG_SZ, // retrieve string
		&pdwType, unsafe.Pointer(&pvData[0]), &pcbData)
	if err != nil {
		panic(err)
	}

	return pvData
}

// Reads a REG_DWORD value with RegGetValue().
func (hKey HKEY) ReadDword(lpSubKey, lpValue string) uint32 {
	pdwType := co.REG_DWORD
	pvData := uint32(0)
	pcbData := uint32(unsafe.Sizeof(pvData))

	err := hKey.RegGetValue(lpSubKey, lpValue, co.RRF_RT_REG_DWORD,
		&pdwType, unsafe.Pointer(&pvData), &pcbData)
	if err != nil {
		panic(err)
	}
	return pvData
}

// Reads a REG_SZ value with RegGetValue().
func (hKey HKEY) ReadString(lpSubKey, lpValue string) string {
	pdwType := co.REG_SZ
	pcbData := uint32(0)

	err := hKey.RegGetValue(lpSubKey, lpValue, co.RRF_RT_REG_SZ, // retrieve length
		&pdwType, nil, &pcbData)
	if err != nil {
		panic(err)
	}

	pvData := make([]uint16, pcbData/2) // pcbData is in bytes; terminating null included

	err = hKey.RegGetValue(lpSubKey, lpValue, co.RRF_RT_REG_SZ, // retrieve string
		&pdwType, unsafe.Pointer(&pvData[0]), &pcbData)
	if err != nil {
		panic(err)
	}

	return Str.FromUint16Slice(pvData)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func (hKey HKEY) RegCloseKey() error {
	ret, _, _ := syscall.Syscall(proc.RegCloseKey.Addr(), 1,
		uintptr(hKey), 0, 0)
	if err := errco.ERROR(ret); err != errco.SUCCESS {
		return err
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-reggetvaluew
func (hKey HKEY) RegGetValue(
	lpSubKey, lpValue string, dwFlags co.RRF, pdwType *co.REG,
	pvData unsafe.Pointer, pcbData *uint32) error {

	ret, _, _ := syscall.Syscall9(proc.RegGetValue.Addr(), 7,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSubKey))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValue))),
		uintptr(dwFlags), uintptr(unsafe.Pointer(pdwType)),
		uintptr(pvData), uintptr(unsafe.Pointer(pcbData)), 0, 0)

	if err := errco.ERROR(ret); err != errco.SUCCESS {
		return err
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsetkeyvaluew
func (hKey HKEY) RegSetKeyValue(
	lpSubKey, lpValueName string, dwType co.REG,
	lpData unsafe.Pointer, cbData uint32) error {

	ret, _, _ := syscall.Syscall6(proc.RegSetKeyValue.Addr(), 6,
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpSubKey))),
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpValueName))),
		uintptr(dwType), uintptr(lpData), uintptr(cbData))

	if err := errco.ERROR(ret); err != errco.SUCCESS {
		return err
	}
	return nil
}

// Writes a REG_BINARY value with RegSetKeyValue().
func (hKey HKEY) WriteBinary(lpSubKey, lpValueName string, lpData []byte) {
	err := hKey.RegSetKeyValue(lpSubKey, lpValueName, co.REG_BINARY,
		unsafe.Pointer(&lpData[0]), uint32(len(lpData)))
	if err != nil {
		panic(err)
	}
}

// Writes a REG_DWORD value with RegSetKeyValue().
func (hKey HKEY) WriteDword(lpSubKey, lpValueName string, lpData uint32) {
	err := hKey.RegSetKeyValue(lpSubKey, lpValueName, co.REG_DWORD,
		unsafe.Pointer(&lpData), uint32(unsafe.Sizeof(lpData)))
	if err != nil {
		panic(err)
	}
}

// Writes a REG_SZ value with RegSetKeyValue().
func (hKey HKEY) WriteString(lpSubKey, lpValueName string, lpData string) {
	slice := Str.ToUint16Slice(lpData)
	err := hKey.RegSetKeyValue(lpSubKey, lpValueName, co.REG_SZ,
		unsafe.Pointer(&slice[0]), uint32(len(slice)*2)) // pass size in bytes, including terminating null
	if err != nil {
		panic(err)
	}
}
