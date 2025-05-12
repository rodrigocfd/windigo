//go:build windows

package ole

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [DVTARGETDEVICE] struct.
//
// ⚠️ You must call [DVTARGETDEVICE.SetTdSize] to initialize the struct.
//
// # Example
//
//	var dvt ole.DVTARGETDEVICE
//	dvt.SetTdSize()
//
// [DVTARGETDEVICE]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-dvtargetdevice
type DVTARGETDEVICE struct {
	tdSize             uint32
	tdDriverNameOffset uint16
	tdDeviceNameOffset uint16
	tdPortNameOffset   uint16
	tdExtDevmodeOffset uint16
	tdData             [1]byte
}

// Sets the tdSize field to the size of the struct, correctly initializing it.
func (dvt *DVTARGETDEVICE) SetTdSize() {
	dvt.tdSize = uint32(unsafe.Sizeof(*dvt))
}

func (dvt *DVTARGETDEVICE) DriverName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdDriverNameOffset)
	return wstr.WstrPtrToStr((*uint16)(ptr))
}

func (dvt *DVTARGETDEVICE) DeviceName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdDeviceNameOffset)
	return wstr.WstrPtrToStr((*uint16)(ptr))
}

func (dvt *DVTARGETDEVICE) PortName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdPortNameOffset)
	return wstr.WstrPtrToStr((*uint16)(ptr))
}

// [FORMATETC] struct.
//
// [FORMATETC]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-formatetc
type FORMATETC struct {
	CfFormat co.CF
	Ptd      *DVTARGETDEVICE
	Aspect   co.DVASPECT
	Lindex   int32
	Tymed    co.TYMED
}

// [STATSTG] struct.
//
// [STATSTG]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-statstg
type STATSTG struct {
	PwcsName          *uint16
	Type              co.STGTY
	CbSize            uint64
	MTime             win.FILETIME
	CTime             win.FILETIME
	ATime             win.FILETIME
	GrfMode           uint32
	GrfLocksSupported co.LOCKTYPE
	ClsId             win.GUID
	GrfStateBits      uint32
	reserved          uint32
}

// [STGMEDIUM] struct.
//
// If you received this struct from a COM call, you'll have to free the memory
// with [ReleaseStgMedium].
//
// [STGMEDIUM]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-ustgmedium-r1
type STGMEDIUM struct {
	tymed          co.TYMED
	data           uintptr // union
	PUnkForRelease IUnknown
}

func (stg *STGMEDIUM) Tymed() co.TYMED {
	return stg.tymed
}

func (stg *STGMEDIUM) HBitmap() (win.HBITMAP, bool) {
	if stg.tymed == co.TYMED_GDI {
		return win.HBITMAP(stg.data), true
	}
	return win.HBITMAP(0), false
}

func (stg *STGMEDIUM) HGlobal() (win.HGLOBAL, bool) {
	if stg.tymed == co.TYMED_HGLOBAL {
		return win.HGLOBAL(stg.data), true
	}
	return win.HGLOBAL(0), false
}

func (stg *STGMEDIUM) FileName() (string, bool) {
	if stg.tymed == co.TYMED_FILE {
		return wstr.WstrPtrToStr((*uint16)(unsafe.Pointer(stg.data))), true
	}
	return "", false
}

func (stg *STGMEDIUM) IStream(releaser *Releaser) (*IStream, bool) {
	if stg.tymed == co.TYMED_ISTREAM {
		ppvt := (**IUnknownVt)(unsafe.Pointer(stg.data))
		pObj := ComObj[IStream](ppvt)
		pCloned := AddRef(pObj, releaser) // clone, because we'll release it independently
		return pCloned, true
	}
	return nil, false
}
