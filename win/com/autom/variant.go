//go:build windows

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
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
)

// OLE Automation [VARIANT] type.
//
// Can be created with one of the NewVariant*() functions, and must be freed
// with VariantClear(). Values can be accessed with one of the accessor methods.
//
// [VARIANT]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-variant
type VARIANT struct {
	vt         automco.VT
	wReserved1 uint16
	wReserved2 uint16
	wReserved3 uint16
	data       [16]byte
}

// Frees the internal object of the VARIANT with [VariantClear].
//
// [VariantClear]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantclear
func (vt *VARIANT) VariantClear() {
	syscall.SyscallN(proc.VariantClear.Addr(),
		uintptr(unsafe.Pointer(vt)))
}

// Returns the type of the VARIANT.
func (vt *VARIANT) Type() automco.VT {
	return vt.vt
}

//------------------------------------------------------------------------------

// Creates a new VARIANT object of type VT_EMPTY with [VariantInit].
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantEmpty()
//	defer vari.VariantClear()
//
// [VariantInit]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantinit
func NewVariantEmpty() VARIANT {
	var vt VARIANT
	syscall.SyscallN(proc.VariantInit.Addr(),
		uintptr(unsafe.Pointer(&vt)))
	return vt
}

// Tells whether the VARIANT object has type VT_EMPTY.
func (vt *VARIANT) IsEmpty() bool {
	return vt.vt == automco.VT_EMPTY
}

// Creates a new VARIANT object of type VT_BOOL.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantBool(true)
//	defer vari.VariantClear()
func NewVariantBool(v bool) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_BOOL
	bool16 := util.Iif(v, int16(-1), int16(0)).(int16)
	binary.LittleEndian.PutUint16(vt.data[:], uint16(bool16))
	return vt
}

// If the VARIANT object has type VT_BOOL, returns the value and true.
// Otherwise, returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantBool(true)
//	defer vari.VariantClear()
//
//	if boolVal, ok := vari.Bool(); ok {
//		println(boolVal)
//	}
func (vt *VARIANT) Bool() (actualValue, isBool bool) {
	switch vt.vt {
	case automco.VT_BOOL:
		bool16 := binary.LittleEndian.Uint16(vt.data[:])
		return int16(bool16) != 0, true
	default:
		return false, false
	}
}

// Creates a new VARIANT of type VT_R4.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantFloat32(40.5)
//	defer vari.VariantClear()
func NewVariantFloat32(v float32) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_R4
	binary.LittleEndian.PutUint32(vt.data[:], math.Float32bits(v))
	return vt
}

// If the VARIANT object has type VT_R4, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantFloat32(40.5)
//	defer vari.VariantClear()
//
//	if floatVal, ok := vari.Float32(); ok {
//		println(floatVal)
//	}
func (vt *VARIANT) Float32() (float32, bool) {
	switch vt.vt {
	case automco.VT_R4:
		return math.Float32frombits(binary.LittleEndian.Uint32(vt.data[:])), true
	default:
		return 0, false
	}
}

// Creates a new VARIANT of type VT_R8.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantFloat64(40.5)
//	defer vari.VariantClear()
func NewVariantFloat64(v float64) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_R8
	binary.LittleEndian.PutUint64(vt.data[:], math.Float64bits(v))
	return vt
}

// If the VARIANT object has type VT_R8, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantFloat64(40.5)
//	defer vari.VariantClear()
//
//	if floatVal, ok := vari.Float64(); ok {
//		println(floatVal)
//	}
func (vt *VARIANT) Float64() (float64, bool) {
	switch vt.vt {
	case automco.VT_R8:
		return math.Float64frombits(binary.LittleEndian.Uint64(vt.data[:])), true
	default:
		return 0, false
	}
}

// Creates a new VARIANT of type VT_DISPATCH.
//
// Note that the IDispatch object will be automatically cloned into the VARIANT,
// so you still must call IDispatch.Release() on your source IDispatch.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	var idisp autom.IDispatch // initialized somewhere
//
//	vari := autom.NewVariantIDispatch(idisp)
//	defer vari.VariantClear()
func NewVariantIDispatch(v IDispatch) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_DISPATCH
	cloned := v.AddRef()
	clonedPpv := cloned.Ptr()
	binary.LittleEndian.PutUint64(vt.data[:], uint64(uintptr(unsafe.Pointer(clonedPpv))))
	return vt
}

// If the VARIANT object has type VT_DISPATCH, returns the value and true.
// Otherwise, returns a default value and false.
//
// ⚠️ You must defer IDispatch.Release() on the returned object.
//
// # Example
//
//	var idisp autom.IDispatch // initialized somewhere
//
//	vari := autom.NewVariantIDispatch(idisp)
//	defer vari.VariantClear()
//
//	if idispVal, ok := vari.IDispatch(); ok {
//		defer idispVal.Release()
//		println(idispVal.Ptr())
//	}
func (vt *VARIANT) IDispatch() (IDispatch, bool) {
	switch vt.vt {
	case automco.VT_DISPATCH:
		ppvData := uintptr(binary.LittleEndian.Uint64(vt.data[:]))
		ppv := (**comvt.IUnknown)(unsafe.Pointer(ppvData))
		iDisp := NewIDispatch(com.NewIUnknown(ppv))
		return NewIDispatch(iDisp.AddRef()), true
	default:
		return nil, false
	}
}

// Creates a new VARIANT of type VT_I1.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantInt8(50)
//	defer vari.VariantClear()
func NewVariantInt8(v int8) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_I1
	vt.data[0] = uint8(v)
	return vt
}

// If the VARIANT object has type VT_I1, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantInt8(50)
//	defer vari.VariantClear()
//
//	if intVal, ok := vari.Int8(); ok {
//		println(intVal)
//	}
func (vt *VARIANT) Int8() (int8, bool) {
	switch vt.vt {
	case automco.VT_I1:
		return int8(vt.data[0]), true
	default:
		return 0, false
	}
}

// Creates a new VARIANT of type VT_I2.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantInt16(50)
//	defer vari.VariantClear()
func NewVariantInt16(v int16) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_I2
	binary.LittleEndian.PutUint16(vt.data[:], uint16(v))
	return vt
}

// If the VARIANT object has type VT_I2, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantInt16(50)
//	defer vari.VariantClear()
//
//	if intVal, ok := vari.Int16(); ok {
//		println(intVal)
//	}
func (vt *VARIANT) Int16() (int16, bool) {
	switch vt.vt {
	case automco.VT_I2:
		return int16(binary.LittleEndian.Uint16(vt.data[:])), true
	default:
		return 0, false
	}
}

// Creates a new VARIANT of type VT_I4.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantInt32(50)
//	defer vari.VariantClear()
func NewVariantInt32(v int32) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_I4
	binary.LittleEndian.PutUint32(vt.data[:], uint32(v))
	return vt
}

// If the VARIANT object has type VT_I4, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantInt32(50)
//	defer vari.VariantClear()
//
//	if intVal, ok := vari.Int32(); ok {
//		println(intVal)
//	}
func (vt *VARIANT) Int32() (int32, bool) {
	switch vt.vt {
	case automco.VT_I4:
		return int32(binary.LittleEndian.Uint32(vt.data[:])), true
	default:
		return 0, false
	}
}

// Creates a new VARIANT of type VT_I8.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantInt64(50)
//	defer vari.VariantClear()
func NewVariantInt64(v int64) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_I8
	binary.LittleEndian.PutUint64(vt.data[:], uint64(v))
	return vt
}

// If the VARIANT object has type VT_I8, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantInt64(50)
//	defer vari.VariantClear()
//
//	if intVal, ok := vari.Int64(); ok {
//		println(intVal)
//	}
func (vt *VARIANT) Int64() (int64, bool) {
	switch vt.vt {
	case automco.VT_I8:
		return int64(binary.LittleEndian.Uint64(vt.data[:])), true
	default:
		return 0, false
	}
}

// Creates a new VARIANT object of type VT_BSTR.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantStr("foo")
//	defer vari.VariantClear()
func NewVariantStr(v string) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_BSTR
	bstr := SysAllocString(v) // will be owned by the VARIANT
	binary.LittleEndian.PutUint64(vt.data[:], uint64(bstr))
	return vt
}

// If the VARIANT object has type VT_BSTR, returns the value and true.
// Otherwise, returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantStr("foo")
//	defer vari.VariantClear()
//
//	if strVal, ok := vari.Str(); ok {
//		println(strVal)
//	}
func (vt *VARIANT) Str() (string, bool) {
	switch vt.vt {
	case automco.VT_BSTR:
		bstr := BSTR(binary.LittleEndian.Uint64(vt.data[:]))
		return bstr.String(), true
	default:
		return "", false
	}
}

// Creates a new VARIANT object with type VT_DATE.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantTime(time.Now())
//	defer vari.VariantClear()
func NewVariantTime(v time.Time) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_DATE

	var double float64
	var st win.SYSTEMTIME
	st.FromTime(v)

	ret, _, _ := syscall.SyscallN(proc.SystemTimeToVariantTime.Addr(),
		uintptr(unsafe.Pointer(&st)), uintptr(unsafe.Pointer(&double)))
	if ret == 0 {
		panic("SystemTimeToVariantTime() failed.")
	}

	binary.LittleEndian.PutUint64(vt.data[:], math.Float64bits(double))
	return vt
}

// If the object contains a value of type time.Time, returns it and true.
// Otherwise, returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantTime(true)
//	defer vari.VariantClear()
//
//	if timeVal, ok := vari.Time(); ok {
//		println(timeVal.Format(time.ANSIC))
//	}
func (vt *VARIANT) Time() (time.Time, bool) {
	switch vt.vt {
	case automco.VT_DATE:
		double := math.Float64frombits(binary.LittleEndian.Uint64(vt.data[:]))
		var st win.SYSTEMTIME

		ret, _, _ := syscall.SyscallN(proc.VariantTimeToSystemTime.Addr(),
			uintptr(math.Float64bits(double)), uintptr(unsafe.Pointer(&st)))
		if ret == 0 {
			panic("VariantTimeToSystemTime() failed.")
		}
		return st.ToTime(), true

	default:
		return time.Time{}, false
	}
}

// Creates a new VARIANT of type VT_UI1.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantUint8(50)
//	defer vari.VariantClear()
func NewVariantUint8(v uint8) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_UI1
	vt.data[0] = v
	return vt
}

// If the VARIANT object has type VT_UI1, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantUint8(50)
//	defer vari.VariantClear()
//
//	if intVal, ok := vari.Uint8(); ok {
//		println(intVal)
//	}
func (vt *VARIANT) Uint8() (uint8, bool) {
	switch vt.vt {
	case automco.VT_UI1:
		return vt.data[0], true
	default:
		return 0, false
	}
}

// Creates a new VARIANT of type VT_UI2.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantUint16(50)
//	defer vari.VariantClear()
func NewVariantUint16(v uint16) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_UI2
	binary.LittleEndian.PutUint16(vt.data[:], v)
	return vt
}

// If the VARIANT object has type VT_UI2, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantUint16(50)
//	defer vari.VariantClear()
//
//	if intVal, ok := vari.Uint16(); ok {
//		println(intVal)
//	}
func (vt *VARIANT) Uint16() (uint16, bool) {
	switch vt.vt {
	case automco.VT_UI2:
		return binary.LittleEndian.Uint16(vt.data[:]), true
	default:
		return 0, false
	}
}

// Creates a new VARIANT of type VT_UI4.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantUint32(50)
//	defer vari.VariantClear()
func NewVariantUint32(v uint32) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_UI4
	binary.LittleEndian.PutUint32(vt.data[:], v)
	return vt
}

// If the VARIANT object has type VT_UI4, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantUint32(50)
//	defer vari.VariantClear()
//
//	if intVal, ok := vari.Uint32(); ok {
//		println(intVal)
//	}
func (vt *VARIANT) Uint32() (uint32, bool) {
	switch vt.vt {
	case automco.VT_UI4:
		return binary.LittleEndian.Uint32(vt.data[:]), true
	default:
		return 0, false
	}
}

// Creates a new VARIANT of type VT_UI8.
//
// ⚠️ You must defer VARIANT.VariantClear().
//
// # Example
//
//	vari := autom.NewVariantUint64(50)
//	defer vari.VariantClear()
func NewVariantUint64(v uint64) VARIANT {
	vt := NewVariantEmpty()
	vt.vt = automco.VT_UI8
	binary.LittleEndian.PutUint64(vt.data[:], v)
	return vt
}

// If the VARIANT object has type VT_UI8, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	vari := autom.NewVariantUint64(50)
//	defer vari.VariantClear()
//
//	if intVal, ok := vari.Uint64(); ok {
//		println(intVal)
//	}
func (vt *VARIANT) Uint64() (uint64, bool) {
	switch vt.vt {
	case automco.VT_UI8:
		return binary.LittleEndian.Uint64(vt.data[:]), true
	default:
		return 0, false
	}
}
