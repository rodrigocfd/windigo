//go:build windows

package win

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [IBindCtx] COM interface.
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

type _IBindCtxVt struct {
	utl.IUnknownVt
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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IBindCtx) IID() *co.IID {
	return &co.IID_IBindCtx
}

// [EnumObjectParam] method.
//
// [EnumObjectParam]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-enumobjectparam
func (me *IBindCtx) EnumObjectParam(releaser *OleReleaser) (*IEnumString, error) {
	return utl.OleNewFromCallWithoutParms[*IEnumString](me, releaser,
		utl.Vt[_IBindCtxVt](me.ppvt).EnumObjectParam)
}

// [GetBindOptions] method.
//
// [GetBindOptions]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-getbindoptions
func (me *IBindCtx) GetBindOptions() (BIND_OPTS3, error) {
	var bo BIND_OPTS3
	bo.SetCbStruct()

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IBindCtxVt](me.ppvt).GetBindOptions,
		me.ppvt,
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
	utl.OleValidateRelease(ppOut)
	var ppvtQueried uintptr
	var wKey wstr.BufEncoder

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IBindCtxVt](me.ppvt).GetObjectParam,
		me.ppvt,
		uintptr(wKey.AllowEmpty(key)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleInjectIfOk(ret, ppOut, ppvtQueried, releaser)
}

// [RegisterObjectBound] method.
//
// [RegisterObjectBound]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-registerobjectbound
func (me *IBindCtx) RegisterObjectBound(obj *IUnknown) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IBindCtxVt](me.ppvt).RegisterObjectBound,
		me.ppvt,
		obj.ppvt)
	return utl.HresultToError(ret)
}

// [ReleaseBoundObjects] method.
//
// [ReleaseBoundObjects]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-releaseboundobjects
func (me *IBindCtx) ReleaseBoundObjects() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IBindCtxVt](me.ppvt).ReleaseBoundObjects)
}

// [RevokeObjectBound] method.
//
// [RevokeObjectBound]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-revokeobjectbound
func (me *IBindCtx) RevokeObjectBound(obj *IUnknown) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IBindCtxVt](me.ppvt).RevokeObjectBound,
		me.ppvt,
		obj.ppvt)
	return utl.HresultToError(ret)
}

// [SetBindOptions] method.
//
// [SetBindOptions]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-setbindoptions
func (me *IBindCtx) SetBindOptions(pBindOpts *BIND_OPTS3) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IBindCtxVt](me.ppvt).SetBindOptions,
		me.ppvt,
		uintptr(unsafe.Pointer(pBindOpts)))
	return utl.HresultToError(ret)
}

// [IDataObject] COM interface.
//
// [IDataObject]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-idataobject
type IDataObject struct{ IUnknown }

type _IDataObjectVt struct {
	utl.IUnknownVt
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

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDataObject) IID() *co.IID {
	return &co.IID_IDataObject
}

// [GetCanonicalFormatEtc] method.
//
// [GetCanonicalFormatEtc]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-idataobject-getcanonicalformatetc
func (me *IDataObject) GetCanonicalFormatEtc(pEtcIn *FORMATETC) (FORMATETC, error) {
	var etcOut FORMATETC
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDataObjectVt](me.ppvt).GetCanonicalFormatEtc,
		me.ppvt,
		uintptr(unsafe.Pointer(pEtcIn)),
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
func (me *IDataObject) GetData(pEtc *FORMATETC) (STGMEDIUM, error) {
	var stg STGMEDIUM
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDataObjectVt](me.ppvt).GetData,
		me.ppvt,
		uintptr(unsafe.Pointer(pEtc)),
		uintptr(unsafe.Pointer(&stg)))
	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return stg, nil
	} else {
		return STGMEDIUM{}, hr
	}
}

// Calls [IDataObject.QueryGetData] to check whether it contains [HDROP]
// contents. If so, calls [IDataObject.GetData] to retrieve the strings.
//
// This method is intended to be used with [IDropTarget.DragEnter] and
// [IDropTarget.Drop], where files are dropped onto a window.
func (me *IDataObject) GetDataHDrop() ([]string, error) {
	fetc := FORMATETC{
		CfFormat: co.CF_HDROP,
		Aspect:   co.DVASPECT_CONTENT,
		Lindex:   -1,
		Tymed:    co.TYMED_HGLOBAL,
	}
	if err := me.QueryGetData(&fetc); err != nil {
		return nil, err
	}

	stg, err := me.GetData(&fetc)
	if err != nil {
		return nil, err
	}
	defer ReleaseStgMedium(&stg)

	hGlobal, ok := stg.HGlobal()
	if !ok {
		return nil, fmt.Errorf("STGMEDIUM didn't have HGLOBAL")
	}

	hMem, _ := hGlobal.GlobalLock()
	defer hGlobal.GlobalUnlock()

	hDrop := HDROP(hMem) // DragFinish() crashes ReleaseStgMedium(), don't call
	files, err := hDrop.DragQueryFile()
	if err != nil {
		return nil, err
	}

	return files, nil
}

// [QueryGetData] method.
//
// [QueryGetData]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-idataobject-querygetdata
func (me *IDataObject) QueryGetData(pEtc *FORMATETC) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDataObjectVt](me.ppvt).QueryGetData,
		me.ppvt,
		uintptr(unsafe.Pointer(pEtc)))
	return utl.HresultToError(ret)
}

// [IEnumString] COM interface.
//
// [IEnumString]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ienumstring
type IEnumString struct{ IUnknown }

type _IEnumStringVt struct {
	utl.IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumString) IID() *co.IID {
	return &co.IID_IEnumString
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-clone
func (me *IEnumString) Clone(releaser *OleReleaser) (*IEnumString, error) {
	return utl.OleNewFromCallWithoutParms[*IEnumString](me, releaser,
		utl.Vt[_IEnumStringVt](me.ppvt).Clone)
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
		utl.Vt[_IEnumStringVt](me.ppvt).Next,
		me.ppvt,
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
	return utl.OleCallWithoutParms(me, utl.Vt[_IEnumStringVt](me.ppvt).Reset)
}

// [Skip] method.
//
// Panics if count is negative.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-skip
func (me *IEnumString) Skip(count int) error {
	utl.PanicNeg(count)
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IEnumStringVt](me.ppvt).Skip,
		me.ppvt,
		uintptr(uint32(count)))
	return utl.HresultToError(ret)
}

// [IEnumUnknown] COM interface.
//
// [IEnumUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ienumunknown
type IEnumUnknown struct{ IUnknown }

type _IEnumUnknownVt struct {
	utl.IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumUnknown) IID() *co.IID {
	return &co.IID_IEnumUnknown
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumunknown-clone
func (me *IEnumUnknown) Clone(releaser *OleReleaser) (*IEnumUnknown, error) {
	return utl.OleNewFromCallWithoutParms[*IEnumUnknown](me, releaser,
		utl.Vt[_IEnumUnknownVt](me.ppvt).Clone)
}

// Returns all [IUnknown] values by calling [IEnumUnknown.Next].
func (me *IEnumUnknown) Enum(releaser *OleReleaser) ([]*IUnknown, error) {
	objs := make([]*IUnknown, 0)
	var pObj *IUnknown
	var hr error

	for {
		pObj, hr = me.Next(releaser)
		if hr != nil { // actual error
			return nil, hr
		} else if pObj == nil { // no more items to fetch
			return objs, nil
		} else { // item fetched
			objs = append(objs, pObj)
		}
	}
}

// [Next] method.
//
// [Next]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumunknown-next
func (me *IEnumUnknown) Next(releaser *OleReleaser) (*IUnknown, error) {
	var ppvtQueried uintptr
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IEnumUnknownVt](me.ppvt).Next,
		me.ppvt,
		1,
		uintptr(unsafe.Pointer(&ppvtQueried)),
		uintptr(unsafe.Pointer(&numFetched)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := utl.OleNew[*IUnknown](ppvtQueried, releaser)
		return pObj, nil
	} else if hr == co.HRESULT_S_FALSE {
		return nil, nil
	} else {
		return nil, hr
	}
}

// [Reset] method.
//
// [Reset]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumunknown-reset
func (me *IEnumUnknown) Reset() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IEnumUnknownVt](me.ppvt).Reset)
}

// [Skip] method.
//
// Panics if count is negative.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumunknown-skip
func (me *IEnumUnknown) Skip(count int) error {
	utl.PanicNeg(count)
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IEnumUnknownVt](me.ppvt).Skip,
		me.ppvt,
		uintptr(uint32(count)))
	return utl.HresultToError(ret)
}

// [ISequentialStream] COM interface.
//
// [ISequentialStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-isequentialstream
type ISequentialStream struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ISequentialStream) IID() *co.IID {
	return &co.IID_ISequentialStream
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
		utl.Vt[utl.ISequentialStreamVt](me.ppvt).Read,
		me.ppvt,
		uintptr(unsafe.Pointer(&destBuf[0])),
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
		utl.Vt[utl.ISequentialStreamVt](me.ppvt).Write,
		me.ppvt,
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(uint32(len(data))),
		uintptr(unsafe.Pointer(&written32)))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(written32), nil
	} else {
		return 0, hr
	}
}

// [IStream] COM interface.
//
// [IStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-istream
type IStream struct{ ISequentialStream }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IStream) IID() *co.IID {
	return &co.IID_IStream
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-clone
func (me *IStream) Clone(releaser *OleReleaser) (*IStream, error) {
	return utl.OleNewFromCallWithoutParms[*IStream](me, releaser,
		utl.Vt[utl.IStreamVt](me.ppvt).Clone)
}

// [Commit] method.
//
// [Commit]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-commit
func (me *IStream) Commit(flags co.STGC) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[utl.IStreamVt](me.ppvt).Commit,
		me.ppvt,
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
		utl.Vt[utl.IStreamVt](me.ppvt).CopyTo,
		me.ppvt,
		dest.ppvt,
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
		utl.Vt[utl.IStreamVt](me.ppvt).LockRegion,
		me.ppvt,
		uintptr(uint64(offset)),
		uintptr(uint64(length)),
		uintptr(lockType))
	return utl.HresultToError(ret)
}

// [Revert] method.
//
// [Revert]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-revert
func (me *IStream) Revert() error {
	return utl.OleCallWithoutParms(me, utl.Vt[utl.IStreamVt](me.ppvt).Revert)
}

// [Seek] method.
//
// [Seek]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-seek
func (me *IStream) Seek(displacement int, origin co.STREAM_SEEK) (newOffset int, hr error) {
	var newOff64 uint64
	ret, _, _ := syscall.SyscallN(
		utl.Vt[utl.IStreamVt](me.ppvt).Seek,
		me.ppvt,
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
		utl.Vt[utl.IStreamVt](me.ppvt).SetSize,
		me.ppvt,
		uintptr(uint64(newSize)))
	return utl.HresultToError(ret)
}

// [Stat] method.
//
// [Stat]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-istream-stat
func (me *IStream) Stat(flag co.STATFLAG) (STATSTG, error) {
	var stg STATSTG
	ret, _, _ := syscall.SyscallN(
		utl.Vt[utl.IStreamVt](me.ppvt).Stat,
		me.ppvt,
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
		utl.Vt[utl.IStreamVt](me.ppvt).UnlockRegion,
		me.ppvt,
		uintptr(uint64(offset)),
		uintptr(uint64(length)),
		uintptr(lockType))
	return utl.HresultToError(ret)
}
