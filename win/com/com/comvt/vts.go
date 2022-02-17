package comvt

// IBindCtx virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ibindctx
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

// IPersist virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ipersist
type IPersist struct {
	IUnknown
	GetClassID uintptr
}

// IUnknown virtual table, base to all COM virtual tables.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}
