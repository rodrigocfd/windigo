//go:build windows

package oleaut

import (
	"encoding/binary"
	"math"
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// OLE Automation [VARIANT] type.
//
// [VARIANT]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-variant
type VARIANT struct {
	tag        co.VT
	wReserved1 uint16
	wReserved2 uint16
	wReserved3 uint16
	data       [16]byte
}

// Calls [VariantClear].
//
// You usually don't need to call this method directly, since every function
// which returns a [COM] object will require a Releaser to manage the object's
// lifetime.
//
// [VariantClear]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantclear
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func (me *VARIANT) Release() {
	syscall.SyscallN(_VariantClear.Addr(),
		uintptr(unsafe.Pointer(me))) // ignore errors
}

var _VariantClear = dll.Oleaut32.NewProc("VariantClear")

// Returns the type of the VARIANT.
func (vt *VARIANT) Type() co.VT {
	return vt.tag
}

// [VariantInit] function.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariantEmpty(rel)
//
// [VariantInit]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantinit
func NewVariantEmpty(releaser *ole.Releaser) *VARIANT {
	v := &VARIANT{}
	syscall.SyscallN(_VariantInit.Addr(),
		uintptr(unsafe.Pointer(v)))
	releaser.Add(v)
	return v
}

var _VariantInit = dll.Oleaut32.NewProc("VariantInit")

// Returns true if current type is VT_EMPTY.
func (v *VARIANT) IsEmpty() bool {
	return v.tag == co.VT_EMPTY
}

// Calls [VariantInit] and sets the type and value.
//
// Allowed [types]:
//   - bool (VT_BOOL)
//   - float32 (VT_R4)
//   - float64 (VT_R8)
//   - *ole.IDispatch (VT_DISPATCH)
//   - int8 (VT_I1)
//   - int16 (VT_I2)
//   - int32 (VT_I4)
//   - int64 (VT_I8)
//   - string (VT_BSTR)
//   - time.Time (VT_DATE)
//   - uint8 (VT_UI1)
//   - uint16 (VT_UI2)
//   - uint32 (VT_UI4)
//   - uint64 (VT_UI8)
//
// Panics if the type of the value is not allowed.
//
// [VariantInit]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-variantinit
// [types]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ne-wtypes-varenum
func NewVariant(releaser *ole.Releaser, value interface{}) *VARIANT {
	v := NewVariantEmpty(releaser)

	switch val := value.(type) {
	case bool:
		v.tag = co.VT_BOOL
		bInt16 := int16(-1)
		if val {
			bInt16 = 1
		}
		binary.LittleEndian.PutUint16(v.data[:], uint16(bInt16))
	case float32:
		v.tag = co.VT_R4
		binary.LittleEndian.PutUint32(v.data[:], math.Float32bits(val))
	case float64:
		v.tag = co.VT_R8
		binary.LittleEndian.PutUint64(v.data[:], math.Float64bits(val))
	case *IDispatch:
		v.tag = co.VT_DISPATCH
		vt.AddRef(val.Ppvt()) // clone, because we'll release it independently
		rawPpvt := uintptr(unsafe.Pointer(val.Ppvt()))
		binary.LittleEndian.PutUint64(v.data[:], uint64(rawPpvt))
	case int8:
		v.tag = co.VT_I1
		v.data[0] = uint8(val)
	case int16:
		v.tag = co.VT_I2
		binary.LittleEndian.PutUint16(v.data[:], uint16(val))
	case int32:
		v.tag = co.VT_I4
		binary.LittleEndian.PutUint32(v.data[:], uint32(val))
	case int64:
		v.tag = co.VT_I8
		binary.LittleEndian.PutUint64(v.data[:], uint64(val))
	case string:
		v.tag = co.VT_BSTR
		bstr, _ := SysAllocString(val) // will be owned by the VARIANT
		binary.LittleEndian.PutUint64(v.data[:], uint64(bstr))
	case time.Time:
		v.tag = co.VT_DATE

		var double float64
		var st win.SYSTEMTIME
		st.SetTime(val)

		ret, _, _ := syscall.SyscallN(_SystemTimeToVariantTime.Addr(),
			uintptr(unsafe.Pointer(&st)), uintptr(unsafe.Pointer(&double)))
		if ret == 0 {
			panic("SystemTimeToVariantTime() failed.") // should never happen, time.Time is always valid
		}
		binary.LittleEndian.PutUint64(v.data[:], math.Float64bits(double))
	case uint8:
		v.tag = co.VT_UI1
		v.data[0] = val
	case uint16:
		v.tag = co.VT_UI2
		binary.LittleEndian.PutUint16(v.data[:], val)
	case uint32:
		v.tag = co.VT_UI4
		binary.LittleEndian.PutUint32(v.data[:], val)
	case uint64:
		v.tag = co.VT_UI8
		binary.LittleEndian.PutUint64(v.data[:], val)
	default:
		panic("Invalid VARIANT value type.")
	}

	return v
}

// If the object has type VT_BOOL, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, true)
//
//	if boolVal, ok := v.Bool(); ok {
//		println(boolVal)
//	}
func (v *VARIANT) Bool() (actualValue, isBool bool) {
	if v.tag == co.VT_BOOL {
		bUint16 := binary.LittleEndian.Uint16(v.data[:])
		return int16(bUint16) != 0, true
	}
	return false, false
}

// If the object has type VT_R4, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, float32(43.5))
//
//	if floatVal, ok := v.Float32(); ok {
//		println(floatVal)
//	}
func (v *VARIANT) Float32() (float32, bool) {
	if v.tag == co.VT_R4 {
		return math.Float32frombits(binary.LittleEndian.Uint32(v.data[:])), true
	}
	return 0, false
}

// If the object has type VT_R8, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, float64(43.5))
//
//	if floatVal, ok := v.Float64(); ok {
//		println(floatVal)
//	}
func (v *VARIANT) Float64() (float64, bool) {
	if v.tag == co.VT_R8 {
		return math.Float64frombits(binary.LittleEndian.Uint64(v.data[:])), true
	}
	return 0, false
}

// If the object has type VT_DISPATCH, returns the value and true. Otherwise,
// returns a default value and false.
//
// The returned object is a clone of the stored one, and must be released.
//
// # Example
//
//	var pDisp IDispatch // initialized somewhere
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, pDisp)
//
//	if pDispVal, ok := v.IDispatch(rel); ok {
//		println(pDisp.Ppvt())
//	}
func (v *VARIANT) IDispatch(releaser *ole.Releaser) (*IDispatch, bool) {
	if v.tag == co.VT_DISPATCH {
		rawPpvt := uintptr(binary.LittleEndian.Uint64(v.data[:]))
		ppvt := (**vt.IUnknown)(unsafe.Pointer(rawPpvt))
		vt.AddRef(ppvt) // clone, so it can be released independently

		pObj := vt.NewObj[IDispatch](ppvt)
		releaser.Add(pObj)
		return pObj, true
	}
	return nil, false
}

// If the object has type VT_I1, returns the value and true. Otherwise, returns
// a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, int8(50))
//
//	if intVal, ok := v.Int8(); ok {
//		println(intVal)
//	}
func (v *VARIANT) Int8() (int8, bool) {
	if v.tag == co.VT_I1 {
		return int8(v.data[0]), true
	}
	return 0, false
}

// If the VARIANT object has type VT_I2, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, int16(50))
//
//	if intVal, ok := v.Int16(); ok {
//		println(intVal)
//	}
func (v *VARIANT) Int16() (int16, bool) {
	if v.tag == co.VT_I2 {
		return int16(binary.LittleEndian.Uint16(v.data[:])), true
	}
	return 0, false
}

// If the object has type VT_I4, returns the value and true. Otherwise, returns
// a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, int32(50))
//
//	if intVal, ok := v.Int32(); ok {
//		println(intVal)
//	}
func (v *VARIANT) Int32() (int32, bool) {
	if v.tag == co.VT_I4 {
		return int32(binary.LittleEndian.Uint32(v.data[:])), true
	}
	return 0, false
}

// If the object has type VT_I8, returns the value and true. Otherwise, returns
// a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, int64(50))
//
//	if intVal, ok := v.Int64(); ok {
//		println(intVal)
//	}
func (v *VARIANT) Int64() (int64, bool) {
	if v.tag == co.VT_I8 {
		return int64(binary.LittleEndian.Uint64(v.data[:])), true
	}
	return 0, false
}

// If the VARIANT has type VT_BSTR, returns the value and true. Otherwise,
// returns a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, "foo")
//
//	if strVal, ok := v.Str(); ok {
//		println(strVal)
//	}
func (v *VARIANT) Str() (string, bool) {
	if v.tag == co.VT_BSTR {
		bstr := BSTR(binary.LittleEndian.Uint64(v.data[:])) // retrieve pointer, but don't free
		return bstr.String(), true
	}
	return "", false
}

var _SystemTimeToVariantTime = dll.Oleaut32.NewProc("SystemTimeToVariantTime")

// If the object contains a value of type time.Time, returns it and true,
// calling [VariantTimeToSystemTime]. Otherwise, returns a default value and
// false.
//
// Panics on error.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, time.Now())
//
//	if dateVal, ok := v.Date(); ok {
//		println(dateVal.Format(time.ANSIC))
//	}
//
// [VariantTimeToSystemTime]: https://learn.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-varianttimetosystemtime
func (v *VARIANT) Date() (time.Time, bool) {
	if v.tag == co.VT_DATE {
		double := math.Float64frombits(binary.LittleEndian.Uint64(v.data[:]))
		var st win.SYSTEMTIME

		ret, _, _ := syscall.SyscallN(_VariantTimeToSystemTime.Addr(),
			uintptr(math.Float64bits(double)), uintptr(unsafe.Pointer(&st)))
		if ret == 0 {
			panic("VariantTimeToSystemTime() failed.") // should never happen, time.Time is always valid
		}
		return st.ToTime(), true
	}
	return time.Time{}, false
}

var _VariantTimeToSystemTime = dll.Oleaut32.NewProc("VariantTimeToSystemTime")

// If the object has type VT_UI1, returns the value and true. Otherwise, returns
// a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, uint8(50))
//
//	if intVal, ok := v.Uint8(); ok {
//		println(intVal)
//	}
func (v *VARIANT) Uint8() (uint8, bool) {
	if v.tag == co.VT_UI1 {
		return v.data[0], true
	}
	return 0, false
}

// If the object has type VT_UI2, returns the value and true. Otherwise, returns
// a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, uint16(50))
//
//	if intVal, ok := v.Uint16(); ok {
//		println(intVal)
//	}
func (v *VARIANT) Uint16() (uint16, bool) {
	if v.tag == co.VT_UI2 {
		return binary.LittleEndian.Uint16(v.data[:]), true
	}
	return 0, false
}

// If the object has type VT_UI4, returns the value and true. Otherwise, returns
// a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, uint32(50))
//
//	if intVal, ok := v.Uint32(); ok {
//		println(intVal)
//	}
func (v *VARIANT) Uint32() (uint32, bool) {
	if v.tag == co.VT_UI4 {
		return binary.LittleEndian.Uint32(v.data[:]), true
	}
	return 0, false
}

// If the object has type VT_UI8, returns the value and true. Otherwise, returns
// a default value and false.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	v := oleaut.NewVariant(rel, uint64(50))
//
//	if intVal, ok := v.Uint64(); ok {
//		println(intVal)
//	}
func (v *VARIANT) Uint64() (uint64, bool) {
	if v.tag == co.VT_UI8 {
		return binary.LittleEndian.Uint64(v.data[:]), true
	}
	return 0, false
}
