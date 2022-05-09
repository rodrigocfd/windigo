//go:build windows

package win

import (
	"runtime"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// This helper method reads a REG_BINARY key value with HKEY.GetValue().
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

// This helper method reads a REG_DWORD key value with HKEY.GetValue().
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

// This helper method reads a REG_QWORD key value with HKEY.GetValue().
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

// This helper method reads a REG_SZ key value with HKEY.GetValue().
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

// This helper method writes a REG_BINARY key value with HKEY.SetKeyValue().
func (hKey HKEY) WriteBinary(subKey, valueName string, data []byte) {
	err := hKey.SetKeyValue(subKey, valueName, co.REG_BINARY,
		unsafe.Pointer(&data[0]), uint32(len(data)))
	if err != nil {
		panic(err)
	}
}

// This helper method writes a REG_DWORD key value with HKEY.SetKeyValue().
func (hKey HKEY) WriteDword(subKey, valueName string, data uint32) {
	err := hKey.SetKeyValue(subKey, valueName, co.REG_DWORD,
		unsafe.Pointer(&data), uint32(unsafe.Sizeof(data)))
	if err != nil {
		panic(err)
	}
}

// This helper method writes a REG_QWORD key value with HKEY.SetKeyValue().
func (hKey HKEY) WriteQword(subKey, valueName string, data uint64) {
	err := hKey.SetKeyValue(subKey, valueName, co.REG_QWORD,
		unsafe.Pointer(&data), uint32(unsafe.Sizeof(data)))
	if err != nil {
		panic(err)
	}
}

// This helper method writes a REG_SZ key value with HKEY.SetKeyValue().
func (hKey HKEY) WriteString(subKey, valueName string, data string) {
	lpData16 := Str.ToNativeSlice(data)
	err := hKey.SetKeyValue(subKey, valueName, co.REG_SZ,
		unsafe.Pointer(&lpData16[0]), uint32(len(lpData16)*2)) // pass size in bytes, including terminating null
	runtime.KeepAlive(lpData16)
	if err != nil {
		panic(err)
	}
}
