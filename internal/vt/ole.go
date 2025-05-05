//go:build windows

package vt

type IDataObject struct {
	IUnknown
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

type IDispatch struct {
	IUnknown
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

type IDropTarget struct {
	IUnknown
	DragEnter uintptr
	DragOver  uintptr
	DragLeave uintptr
	Drop      uintptr
}

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

type ISequentialStream struct {
	IUnknown
	Read  uintptr
	Write uintptr
}

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

type ITypeInfo struct {
	IUnknown
	GetTypeAttr          uintptr
	GetTypeComp          uintptr
	GetFuncDesc          uintptr
	GetVarDesc           uintptr
	GetNames             uintptr
	GetRefTypeOfImplType uintptr
	GetImplTypeFlags     uintptr
	GetIDsOfNames        uintptr
	Invoke               uintptr
	GetDocumentation     uintptr
	GetDllEntry          uintptr
	GetRefTypeInfo       uintptr
	AddressOfMember      uintptr
	CreateInstance       uintptr
	GetMops              uintptr
	GetContainingTypeLib uintptr
	ReleaseTypeAttr      uintptr
	ReleaseFuncDesc      uintptr
	ReleaseVarDesc       uintptr
}

type IUnknown struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}
