//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [IBindCtx] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	bindCtx, _ := win.CreateBindCtx(rel)
//
// [IBindCtx]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ibindctx
type IBindCtx struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IBindCtx) IID() co.IID {
	return co.IID_IBindCtx
}

// [EnumObjectParam] method.
//
// [EnumObjectParam]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-enumobjectparam
func (me *IBindCtx) EnumObjectParam(releaser *OleReleaser) (*IEnumString, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).EnumObjectParam,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IEnumString{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetBindOptions] method.
//
// [GetBindOptions]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-getbindoptions
func (me *IBindCtx) GetBindOptions() (BIND_OPTS3, error) {
	var bo BIND_OPTS3
	bo.SetCbStruct()

	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).GetBindOptions,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&bo)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return bo, nil
	} else {
		return BIND_OPTS3{}, hr
	}
}

// [GetObjectParam] method.
//
// [GetObjectParam]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-getobjectparam
func (me *IBindCtx) GetObjectParam(releaser *OleReleaser, key string, ppOut interface{}) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut) // safety, because pOut will receive the new COM object

	var ppvtQueried **_IUnknownVt
	var wKey wstr.BufEncoder

	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).GetObjectParam,
		uintptr(wKey.AllowEmpty(key)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pOut = utl.OleCreateObj(ppOut, unsafe.Pointer(ppvtQueried)).(OleObj)
		releaser.Add(pOut)
		return nil
	} else {
		return hr
	}
}

// [RegisterObjectBound] method.
//
// [RegisterObjectBound]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-registerobjectbound
func (me *IBindCtx) RegisterObjectBound(obj *IUnknown) error {
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).RegisterObjectBound,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(obj.Ppvt())))
	return utl.HresultToError(ret)
}

// [ReleaseBoundObjects] method.
//
// [ReleaseBoundObjects]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-releaseboundobjects
func (me *IBindCtx) ReleaseBoundObjects() error {
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).ReleaseBoundObjects,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.HresultToError(ret)
}

// [RevokeObjectBound] method.
//
// [RevokeObjectBound]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-revokeobjectbound
func (me *IBindCtx) RevokeObjectBound(obj *IUnknown) error {
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).RevokeObjectBound,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(obj.Ppvt())))
	return utl.HresultToError(ret)
}

// [SetBindOptions] method.
//
// [SetBindOptions]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-setbindoptions
func (me *IBindCtx) SetBindOptions(bindOpts *BIND_OPTS3) error {
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).SetBindOptions,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(bindOpts)))
	return utl.HresultToError(ret)
}

type _IBindCtxVt struct {
	_IUnknownVt
	RegisterObjectBound   uintptr
	RevokeObjectBound     uintptr
	ReleaseBoundObjects   uintptr
	SetBindOptions        uintptr
	GetBindOptions        uintptr
	GetRunningObjectTable uintptr
	RegisterObjectParam   uintptr
	GetObjectParam        uintptr
	EnumObjectParam       uintptr
	RevokeObjectParam     uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IDataObject] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IDataObject]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-idataobject
type IDataObject struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDataObject) IID() co.IID {
	return co.IID_IDataObject
}

// [GetCanonicalFormatEtc] method.
//
// [GetCanonicalFormatEtc]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-idataobject-getcanonicalformatetc
func (me *IDataObject) GetCanonicalFormatEtc(etcIn *FORMATETC) (FORMATETC, error) {
	var etcOut FORMATETC
	ret, _, _ := syscall.SyscallN(
		(*_IDataObjectVt)(unsafe.Pointer(*me.Ppvt())).GetCanonicalFormatEtc,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(etcIn)),
		uintptr(unsafe.Pointer(&etcOut)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return etcOut, nil
	} else {
		return FORMATETC{}, hr
	}
}

// [GetData] method.
//
// ⚠️ You must defer [ReleaseStgMedium] on the returned object.
//
// [GetData]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-idataobject-getdata
func (me *IDataObject) GetData(etc *FORMATETC) (STGMEDIUM, error) {
	var stg STGMEDIUM
	ret, _, _ := syscall.SyscallN(
		(*_IDataObjectVt)(unsafe.Pointer(*me.Ppvt())).GetData,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(etc)),
		uintptr(unsafe.Pointer(&stg)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return stg, nil
	} else {
		return STGMEDIUM{}, hr
	}
}

// [QueryGetData] method.
//
// [QueryGetData]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-idataobject-querygetdata
func (me *IDataObject) QueryGetData(etc *FORMATETC) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDataObjectVt)(unsafe.Pointer(*me.Ppvt())).QueryGetData,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(etc)))
	return utl.HresultToError(ret)
}

type _IDataObjectVt struct {
	_IUnknownVt
	GetData               uintptr
	GetDataHere           uintptr
	QueryGetData          uintptr
	GetCanonicalFormatEtc uintptr
	SetData               uintptr
	EnumFormatEtc         uintptr
	DAdvise               uintptr
	DUnadvise             uintptr
	EnumDAdvise           uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IEnumString] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IEnumString]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ienumstring
type IEnumString struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumString) IID() co.IID {
	return co.IID_IEnumString
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-clone
func (me *IEnumString) Clone(releaser *OleReleaser) (*IEnumString, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IEnumStringVt)(unsafe.Pointer(*me.Ppvt())).Clone,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IEnumString{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// Returns all string values by calling [IEnumString.Next].
func (me *IEnumString) Enum() ([]string, error) {
	strs := make([]string, 0)
	var s string
	var hr error

	for {
		s, hr = me.Next()
		if hr != nil { // actual error
			return nil, hr
		} else if s == "" { // no more items to fetch
			return strs, nil
		} else { // item fetched
			strs = append(strs, s)
		}
	}
}

// [Next] method.
//
// [Next]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-next
func (me *IEnumString) Next() (string, error) {
	var pv uintptr
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		(*_IEnumStringVt)(unsafe.Pointer(*me.Ppvt())).Next,
		uintptr(unsafe.Pointer(me.Ppvt())),
		1,
		uintptr(unsafe.Pointer(&pv)),
		uintptr(unsafe.Pointer(&numFetched)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer HTASKMEM(pv).CoTaskMemFree()
		name := wstr.DecodePtr((*uint16)(unsafe.Pointer(pv)))
		return name, nil
	} else if hr == co.HRESULT_S_FALSE {
		return "", nil
	} else {
		return "", hr
	}
}

// [Reset] method.
//
// [Reset]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-reset
func (me *IEnumString) Reset() error {
	ret, _, _ := syscall.SyscallN(
		(*_IEnumStringVt)(unsafe.Pointer(*me.Ppvt())).Reset,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.HresultToError(ret)
}

// [Skip] method.
//
// Panics if count is negative.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-skip
func (me *IEnumString) Skip(count int) error {
	utl.PanicNeg(count)
	ret, _, _ := syscall.SyscallN(
		(*_IEnumStringVt)(unsafe.Pointer(*me.Ppvt())).Skip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(count)))
	return utl.HresultToError(ret)
}

type _IEnumStringVt struct {
	_IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [ISequentialStream] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ISequentialStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-isequentialstream
type ISequentialStream struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ISequentialStream) IID() co.IID {
	return co.IID_ISequentialStream
}

// [Read] method.
//
// If returned numBytesRead is lower than requested buffer size, it means
// the end of stream was reached.
//
// [Read]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-isequentialstream-read
func (me *ISequentialStream) Read(destBuf []byte) (numBytesRead int, hr error) {
	var read32 uint32
	ret, _, _ := syscall.SyscallN(
		(*_ISequentialStreamVt)(unsafe.Pointer(*me.Ppvt())).Read,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(unsafe.SliceData(destBuf))),
		uintptr(uint32(len(destBuf))),
		uintptr(unsafe.Pointer(&read32)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(read32), nil
	} else {
		return 0, hr
	}
}

// [Write] method.
//
// [Write]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-isequentialstream-write
func (me *ISequentialStream) Write(data []byte) (numBytesWritten int, hr error) {
	var written32 uint32
	ret, _, _ := syscall.SyscallN(
		(*_ISequentialStreamVt)(unsafe.Pointer(*me.Ppvt())).Write,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(unsafe.SliceData(data))),
		uintptr(uint32(len(data))),
		uintptr(unsafe.Pointer(&written32)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(written32), nil
	} else {
		return 0, hr
	}
}

type _ISequentialStreamVt struct {
	_IUnknownVt
	Read  uintptr
	Write uintptr
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// [IStream] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-istream
type IStream struct{ ISequentialStream }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IStream) IID() co.IID {
	return co.IID_IStream
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-clone
func (me *IStream) Clone(releaser *OleReleaser) (*IStream, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Clone,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IStream{ISequentialStream{IUnknown{ppvtQueried}}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [Commit] method.
//
// [Commit]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-commit
func (me *IStream) Commit(flags co.STGC) error {
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Commit,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(flags))
	return utl.HresultToError(ret)
}

// [CopyTo] method.
//
// Panics if numBytes is negative.
//
// [CopyTo]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-copyto
func (me *IStream) CopyTo(
	dest *IStream,
	numBytes int,
) (numBytesRead, numBytesWritten int, hr error) {
	utl.PanicNeg(numBytes)
	var read64, written64 uint64

	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).CopyTo,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(dest.Ppvt())),
		uintptr(uint64(numBytes)),
		uintptr(unsafe.Pointer(&read64)),
		uintptr(unsafe.Pointer(&written64)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(read64), int(written64), nil
	} else {
		return 0, 0, hr
	}
}

// [LockRegion] method.
//
// Panics if offset or length is negative.
//
// ⚠️ You must defer [IStream.UnlockRegion].
//
// [LockRegion]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-lockregion
func (me *IStream) LockRegion(offset, length int, lockType co.LOCKTYPE) error {
	utl.PanicNeg(offset, length)
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).LockRegion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint64(offset)),
		uintptr(uint64(length)),
		uintptr(lockType))
	return utl.HresultToError(ret)
}

// [Revert] method.
//
// [Revert]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-revert
func (me *IStream) Revert() error {
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Revert,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.HresultToError(ret)
}

// [Seek] method.
//
// [Seek]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-seek
func (me *IStream) Seek(displacement int, origin co.STREAM_SEEK) (newOffset int, hr error) {
	var newOff64 uint64
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Seek,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(int64(displacement)),
		uintptr(origin),
		uintptr(unsafe.Pointer(&newOff64)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(newOff64), nil
	} else {
		return 0, hr
	}
}

// [SetSize] method.
//
// Panics if newSize is negative.
//
// [SetSize]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-setsize
func (me *IStream) SetSize(newSize int) error {
	utl.PanicNeg(newSize)
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).SetSize,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint64(newSize)))
	return utl.HresultToError(ret)
}

// [Stat] method.
//
// [Stat]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-stat
func (me *IStream) Stat(flag co.STATFLAG) (STATSTG, error) {
	var stg STATSTG
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).Stat,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&stg)),
		uintptr(flag))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return stg, nil
	} else {
		return STATSTG{}, hr
	}
}

// [UnlockRegion] method.
//
// Paired with [IStream.LockRegion].
//
// Panics if offset or length is negative.
//
// [UnlockRegion]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-unlockregion
func (me *IStream) UnlockRegion(offset, length int, lockType co.LOCKTYPE) error {
	utl.PanicNeg(offset, length)
	ret, _, _ := syscall.SyscallN(
		(*_IStreamVt)(unsafe.Pointer(*me.Ppvt())).UnlockRegion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint64(offset)),
		uintptr(uint64(length)),
		uintptr(lockType))
	return utl.HresultToError(ret)
}

type _IStreamVt struct {
	_ISequentialStreamVt
	Seek         uintptr
	SetSize      uintptr
	CopyTo       uintptr
	Commit       uintptr
	Revert       uintptr
	LockRegion   uintptr
	UnlockRegion uintptr
	Stat         uintptr
	Clone        uintptr
}
