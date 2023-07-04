//go:build windows

package comvt

// [IBindCtx] virtual table.
//
// [IBindCtx]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ibindctx
type IBindCtx struct {
	IUnknown
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

// [IPersist] virtual table.
//
// [IPersist]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
type IPersist struct {
	IUnknown
	GetClassID uintptr
}

// [IPicture] virtual table.
//
// [IPicture]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/nn-ocidl-ipicture
type IPicture struct {
	IUnknown
	Get_Handle             uintptr
	Get_hPal               uintptr
	Get_Type               uintptr
	Get_Width              uintptr
	Get_Height             uintptr
	Render                 uintptr
	Set_hPal               uintptr
	Get_CurDC              uintptr
	SelectPicture          uintptr
	Get_KeepOriginalFormat uintptr
	Put_KeepOriginalFormat uintptr
	PictureChanged         uintptr
	SaveAsFile             uintptr
	Get_Attributes         uintptr
}

// [ISequentialStream] virtual table.
//
// [ISequentialStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-isequentialstream
type ISequentialStream struct {
	IUnknown
	Read  uintptr
	Write uintptr
}

// [IStream] virtual table.
//
// [IStream]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-istream
type IStream struct {
	ISequentialStream
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

// [IUnknown] virtual table, base to all COM virtual tables.
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}
