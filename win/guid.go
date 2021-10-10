package win

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/rodrigocfd/windigo/win/co"
)

// Can be created with NewGuidFromClsid() or NewGuidFromIid().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/guiddef/ns-guiddef-guid
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 uint64
}

// Returns a GUID struct from a CLSID string.
func NewGuidFromClsid(clsid co.CLSID) *GUID {
	return _NewGuidFromStr(string(clsid))
}

// Returns a GUID struct from an IID string.
func NewGuidFromIid(iid co.IID) *GUID {
	return _NewGuidFromStr(string(iid))
}

func _NewGuidFromStr(strGuid string) *GUID {
	if len(strGuid) != 36 {
		panic(fmt.Sprintf("Malformed GUID: %s", strGuid))
	}

	strs := strings.Split(strGuid, "-")
	if len(strs) != 5 {
		panic(fmt.Sprintf("Malformed GUID parts: %s", strGuid))
	}

	num1, e := strconv.ParseUint(strs[0], 16, 32)
	if e != nil {
		panic(e)
	}
	num2, e := strconv.ParseUint(strs[1], 16, 16)
	if e != nil {
		panic(e)
	}
	num3, e := strconv.ParseUint(strs[2], 16, 16)
	if e != nil {
		panic(e)
	}
	num4, e := strconv.ParseUint(strs[3], 16, 16)
	if e != nil {
		panic(e)
	}
	num5, e := strconv.ParseUint(strs[4], 16, 64)
	if e != nil {
		panic(e)
	}

	newGuid := &GUID{
		Data1: uint32(num1),
		Data2: uint16(num2),
		Data3: uint16(num3),
		Data4: (uint64(num4) << 48) | num5,
	}

	buf64 := [8]byte{}
	binary.BigEndian.PutUint64(buf64[:], newGuid.Data4)
	newGuid.Data4 = binary.LittleEndian.Uint64(buf64[:]) // reverse bytes of Data4
	return newGuid
}
