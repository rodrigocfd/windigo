package autom

import (
	"encoding/binary"
	"math"
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automco"
)

// OLE Automation VARIANT type.
//
// Can be created with one of the NewVariant*() functions, and must be freed
// with VariantClear(). Values can be accessed with one of the accessor methods,
// which will panic if the data type is different.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-variant
type VARIANT struct {
	vt         automco.VT
	wReserved1 uint16
	wReserved2 uint16
	wReserved3 uint16
	data       [16]byte
}

// Frees the internal object of the VARIANT.
func (vt *VARIANT) VariantClear() {
	syscall.Syscall(proc.VariantClear.Addr(), 1,
		uintptr(unsafe.Pointer(vt)), 0, 0)
}

// Returns the type of the VARIANT.
func (vt *VARIANT) Type() automco.VT {
	return vt.vt
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantEmpty() VARIANT {
	var vt VARIANT
	syscall.Syscall(proc.VariantInit.Addr(), 1,
		uintptr(unsafe.Pointer(&vt)), 0, 0)
	return vt
}

//------------------------------------------------------------------------------

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantBool(v bool) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_BOOL
	bool16 := util.Iif(v, int16(-1), int16(0)).(int16)
	binary.LittleEndian.PutUint16(vt.data[:], uint16(bool16))
	return vt
}
func (vt *VARIANT) Bool() bool {
	if vt.vt != automco.VT_BOOL {
		panic("Variant does not own a VT_BOOL value.")
	}
	bool16 := binary.LittleEndian.Uint16(vt.data[:])
	return int16(bool16) != 0
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantBstr(v BSTR) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_BSTR
	binary.LittleEndian.PutUint64(vt.data[:], uint64(v))
	return vt
}
func (vt *VARIANT) Bstr() BSTR {
	if vt.vt != automco.VT_BSTR {
		panic("Variant does not own a VT_BSTR value.")
	}
	return BSTR(binary.LittleEndian.Uint64(vt.data[:]))
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantDate(v time.Time) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_DATE

	var double float64
	var st win.SYSTEMTIME
	st.FromTime(v)

	ret, _, _ := syscall.Syscall(proc.SystemTimeToVariantTime.Addr(), 2,
		uintptr(unsafe.Pointer(&st)), uintptr(unsafe.Pointer(&double)), 0)
	if ret == 0 {
		panic("SystemTimeToVariantTime() failed.")
	}

	binary.LittleEndian.PutUint64(vt.data[:], math.Float64bits(double))
	return vt
}
func (vt *VARIANT) Date() time.Time {
	if vt.vt != automco.VT_DATE {
		panic("Variant does not own a VT_DATE value.")
	}

	double := math.Float64frombits(binary.LittleEndian.Uint64(vt.data[:]))
	var st win.SYSTEMTIME

	ret, _, _ := syscall.Syscall(proc.VariantTimeToSystemTime.Addr(), 2,
		uintptr(math.Float64bits(double)), uintptr(unsafe.Pointer(&st)), 0)
	if ret == 0 {
		panic("VariantTimeToSystemTime() failed.")
	}

	return st.ToTime()
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantInt16(v int16) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_I2
	binary.LittleEndian.PutUint16(vt.data[:], uint16(v))
	return vt
}
func (vt *VARIANT) Int16() int16 {
	if vt.vt != automco.VT_I2 {
		panic("Variant does not own a VT_I2 value.")
	}
	return int16(binary.LittleEndian.Uint16(vt.data[:]))
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantInt32(v int32) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_I4
	binary.LittleEndian.PutUint32(vt.data[:], uint32(v))
	return vt
}
func (vt *VARIANT) Int32() int32 {
	if vt.vt != automco.VT_I4 {
		panic("Variant does not own a VT_I4 value.")
	}
	return int32(binary.LittleEndian.Uint32(vt.data[:]))
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantInt8(v int8) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_I1
	vt.data[0] = uint8(v)
	return vt
}
func (vt *VARIANT) Int8() int8 {
	if vt.vt != automco.VT_I1 {
		panic("Variant does not own a VT_I1 value.")
	}
	return int8(vt.data[0])
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantReal4(v float32) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_R4
	binary.LittleEndian.PutUint32(vt.data[:], math.Float32bits(v))
	return vt
}
func (vt *VARIANT) Real4() float32 {
	if vt.vt != automco.VT_R4 {
		panic("Variant does not own a VT_R4 value.")
	}
	return math.Float32frombits(binary.LittleEndian.Uint32(vt.data[:]))
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantReal8(v float64) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_R8
	binary.LittleEndian.PutUint64(vt.data[:], math.Float64bits(v))
	return vt
}
func (vt *VARIANT) Real8() float64 {
	if vt.vt != automco.VT_R8 {
		panic("Variant does not own a VT_R8 value.")
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(vt.data[:]))
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantUint16(v uint16) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_UI2
	binary.LittleEndian.PutUint16(vt.data[:], v)
	return vt
}
func (vt *VARIANT) Uint16() uint16 {
	if vt.vt != automco.VT_UI2 {
		panic("Variant does not own a VT_UI2 value.")
	}
	return binary.LittleEndian.Uint16(vt.data[:])
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantUint32(v uint32) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_UI4
	binary.LittleEndian.PutUint32(vt.data[:], v)
	return vt
}
func (vt *VARIANT) Uint32() uint32 {
	if vt.vt != automco.VT_UI4 {
		panic("Variant does not own a VT_UI4 value.")
	}
	return binary.LittleEndian.Uint32(vt.data[:])
}

// ⚠️ You must defer VARIANT.VariantClear().
func NewVariantUint8(v uint8) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_UI1
	vt.data[0] = v
	return vt
}
func (vt *VARIANT) Uint8() uint8 {
	if vt.vt != automco.VT_UI1 {
		panic("Variant does not own a VT_UI1 value.")
	}
	return vt.data[0]
}
