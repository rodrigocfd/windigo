//go:build windows

package win

import (
	"runtime"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// This helper method reads a REG_BINARY key value with HKEY.GetValue().
func (hKey HKEY) GetBinary(subKey, value string) []byte {
	var pDataLen uint32
	pdwType := co.REG_BINARY

	err := hKey.RegGetValue(subKey, value, co.RRF_RT_REG_BINARY, // retrieve length
		&pdwType, nil, &pDataLen)
	if err != nil {
		panic(err)
	}

	pData := make([]byte, pDataLen)

	err = hKey.RegGetValue(subKey, value, co.RRF_RT_REG_SZ, // retrieve string
		&pdwType, unsafe.Pointer(&pData[0]), &pDataLen)
	if err != nil {
		panic(err)
	}

	return pData
}

// This helper method reads a REG_DWORD key value with HKEY.GetValue().
func (hKey HKEY) GetDword(subKey, value string) uint32 {
	var pData uint32
	pDataLen := uint32(unsafe.Sizeof(pData))
	pdwType := co.REG_DWORD

	err := hKey.RegGetValue(subKey, value, co.RRF_RT_REG_DWORD,
		&pdwType, unsafe.Pointer(&pData), &pDataLen)
	if err != nil {
		panic(err)
	}
	return pData
}

// This helper method reads a REG_QWORD key value with HKEY.GetValue().
func (hKey HKEY) GetQword(subKey, value string) uint64 {
	var pData uint64
	pDataLen := uint32(unsafe.Sizeof(pData))
	pdwType := co.REG_QWORD

	err := hKey.RegGetValue(subKey, value, co.RRF_RT_REG_QWORD,
		&pdwType, unsafe.Pointer(&pData), &pDataLen)
	if err != nil {
		panic(err)
	}
	return pData
}

// This helper method reads a REG_SZ key value with HKEY.GetValue().
func (hKey HKEY) GetString(subKey, value string) string {
	var pDataLen uint32
	pdwType := co.REG_SZ

	err := hKey.RegGetValue(subKey, value, co.RRF_RT_REG_SZ, // retrieve length
		&pdwType, nil, &pDataLen)
	if err != nil {
		panic(err)
	}

	pData := make([]uint16, pDataLen/2) // pcbData is in bytes; terminating null included

	err = hKey.RegGetValue(subKey, value, co.RRF_RT_REG_SZ, // retrieve string
		&pdwType, unsafe.Pointer(&pData[0]), &pDataLen)
	if err != nil {
		panic(err)
	}

	return Str.FromNativeSlice(pData)
}

// This helper method writes a REG_BINARY key value with HKEY.SetKeyValue().
func (hKey HKEY) PutBinary(subKey, valueName string, data []byte) {
	err := hKey.RegSetKeyValue(subKey, valueName, co.REG_BINARY,
		unsafe.Pointer(&data[0]), uint32(len(data)))
	if err != nil {
		panic(err)
	}
}

// This helper method writes a REG_DWORD key value with HKEY.SetKeyValue().
func (hKey HKEY) PutDword(subKey, valueName string, data uint32) {
	err := hKey.RegSetKeyValue(subKey, valueName, co.REG_DWORD,
		unsafe.Pointer(&data), uint32(unsafe.Sizeof(data)))
	if err != nil {
		panic(err)
	}
}

// This helper method writes a REG_QWORD key value with HKEY.SetKeyValue().
func (hKey HKEY) PutQword(subKey, valueName string, data uint64) {
	err := hKey.RegSetKeyValue(subKey, valueName, co.REG_QWORD,
		unsafe.Pointer(&data), uint32(unsafe.Sizeof(data)))
	if err != nil {
		panic(err)
	}
}

// This helper method writes a REG_SZ key value with HKEY.SetKeyValue().
func (hKey HKEY) PutString(subKey, valueName string, data string) {
	lpData16 := Str.ToNativeSlice(data)
	err := hKey.RegSetKeyValue(subKey, valueName, co.REG_SZ,
		unsafe.Pointer(&lpData16[0]), uint32(len(lpData16)*2)) // pass size in bytes, including terminating null
	runtime.KeepAlive(lpData16)
	if err != nil {
		panic(err)
	}
}
