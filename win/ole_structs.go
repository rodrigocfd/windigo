//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [DVTARGETDEVICE] struct.
//
// ⚠️ You must call [DVTARGETDEVICE.SetTdSize] to initialize the struct.
//
// # Example
//
//	var dvt win.DVTARGETDEVICE
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
	return wstr.WinPtrToGo((*uint16)(ptr))
}

func (dvt *DVTARGETDEVICE) DeviceName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdDeviceNameOffset)
	return wstr.WinPtrToGo((*uint16)(ptr))
}

func (dvt *DVTARGETDEVICE) PortName() string {
	ptr := unsafe.Pointer(dvt)
	ptr = unsafe.Add(ptr, dvt.tdPortNameOffset)
	return wstr.WinPtrToGo((*uint16)(ptr))
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
	MTime             FILETIME
	CTime             FILETIME
	ATime             FILETIME
	GrfMode           uint32
	GrfLocksSupported co.LOCKTYPE
	ClsId             GUID
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

// Returns the tymed field.
func (stg *STGMEDIUM) Tymed() co.TYMED {
	return stg.tymed
}

// Attemps to return the [HBITMAP] if tymed == co.TYMED_GDI.
func (stg *STGMEDIUM) HBitmap() (HBITMAP, bool) {
	if stg.tymed == co.TYMED_GDI {
		return HBITMAP(stg.data), true
	}
	return HBITMAP(0), false
}

// Attemps to return the [HGLOBAL] if tymed == co.TYMED_HGLOBAL.
func (stg *STGMEDIUM) HGlobal() (HGLOBAL, bool) {
	if stg.tymed == co.TYMED_HGLOBAL {
		return HGLOBAL(stg.data), true
	}
	return HGLOBAL(0), false
}

// Attemps to return the string if tymed == co.TYMED_FILE.
func (stg *STGMEDIUM) FileName() (string, bool) {
	if stg.tymed == co.TYMED_FILE {
		return wstr.WinPtrToGo((*uint16)(unsafe.Pointer(stg.data))), true
	}
	return "", false
}

// Attemps to return the [IStream] if tymed == co.TYMED_ISTREAM.
func (stg *STGMEDIUM) IStream(releaser *OleReleaser) (*IStream, bool) {
	if stg.tymed == co.TYMED_ISTREAM {
		ppvt := (**_IUnknownVt)(unsafe.Pointer(stg.data))
		pCurrent := &IStream{ISequentialStream{IUnknown{ppvt}}}

		var pCloned *IStream
		pCurrent.AddRef(releaser, &pCloned) // clone, because we'll release it independently
		return pCloned, true
	}
	return nil, false
}
