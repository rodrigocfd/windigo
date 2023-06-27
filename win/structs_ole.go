//go:build windows

package win

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// [GUID] struct.
//
// Can be created with NewGuidFromClsid() or NewGuidFromIid().
//
// [GUID]: https://learn.microsoft.com/en-us/windows/win32/api/guiddef/ns-guiddef-guid
type GUID struct {
	data1 uint32
	data2 uint16
	data3 uint16
	data4 uint64
}

// Returns a GUID struct from a CLSID string.
func GuidFromClsid(clsid co.CLSID) *GUID {
	return _NewGuidFromStr(string(clsid))
}

// Returns a GUID struct from an IID string.
func GuidFromIid(iid co.IID) *GUID {
	return _NewGuidFromStr(string(iid))
}

// Formats the GUID as a string.
func (g *GUID) String() string {
	data4 := util.ReverseBytes64(g.data4)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		g.data1, g.data2, g.data3,
		data4>>48, data4&0xffff_ffff_ffff)
}

func _NewGuidFromStr(strGuid string) *GUID {
	strs := strings.Split(strGuid, "-")
	if len(strs) != 5 {
		panic(fmt.Sprintf("Malformed GUID parts: %s", strGuid))
	}

	num1, e := strconv.ParseUint(strs[0], 16, 32)
	if e != nil {
		panic(e)
	}
	if num1 > 0xffff_ffff {
		panic(fmt.Sprintf("GUID part 1 overflow: %x", num1))
	}

	var nums16 [3]uint16
	for p := 1; p <= 3; p++ {
		num, e := strconv.ParseUint(strs[p], 16, 16)
		if e != nil {
			panic(e)
		}
		if num > 0xffff {
			panic(fmt.Sprintf("GUID part %d overflows: %x", p, num))
		}
		nums16[p-1] = uint16(num)
	}

	num5, e := strconv.ParseUint(strs[4], 16, 64)
	if e != nil {
		panic(e)
	}
	if num5 > 0xffff_ffff_ffff {
		panic(fmt.Sprintf("GUID part 5 overflow: %x", num5))
	}

	return &GUID{
		data1: uint32(num1),
		data2: nums16[0],
		data3: nums16[1],
		data4: util.ReverseBytes64((uint64(nums16[2]) << 48) | num5),
	}
}
