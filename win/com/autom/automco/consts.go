package automco

// FUNCDESC callconv.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-funcdesc
type CALLCONV uint32

const (
	CALLCONV_FASTCALL CALLCONV = iota
	CALLCONV_CDECL
	CALLCONV_MSCPASCAL
	CALLCONV_PASCAL
	CALLCONV_MACPASCAL
	CALLCONV_STDCALL
	CALLCONV_FPFASTCALL
	CALLCONV_SYSCALL
	CALLCONV_MPWCDECL
	CALLCONV_MPWPASCAL
	CALLCONV_MAX
)

// FUNCDESC funckind.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-funcdesc
type FUNCKIND uint32

const (
	FUNCKIND_VIRTUAL FUNCKIND = iota
	FUNCKIND_PUREVIRTUAL
	FUNCKIND_NONVIRTUAL
	FUNCKIND_STATIC
	FUNCKIND_DISPATCH
)

// FUNCDESC invkind.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-funcdesc
type INVOKEKIND uint32

const (
	INVOKEKIND_FUNC           INVOKEKIND = 1
	INVOKEKIND_PROPERTYGET    INVOKEKIND = 2
	INVOKEKIND_PROPERTYPUT    INVOKEKIND = 4
	INVOKEKIND_PROPERTYPUTREF INVOKEKIND = 8
)

// VARIANT types.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wtypes/ne-wtypes-varenum
type VT uint16

const (
	VT_EMPTY            VT = 0      // Nothing.
	VT_NULL             VT = 1      // SQL style NULL.
	VT_I2               VT = 2      // 2 byte signed int.
	VT_I4               VT = 3      // 4 byte signed int.
	VT_R4               VT = 4      // 4 byte real.
	VT_R8               VT = 5      // 8 byte real.
	VT_CY               VT = 6      // Currency.
	VT_DATE             VT = 7      // Date.
	VT_BSTR             VT = 8      // OLE Automation string.
	VT_DISPATCH         VT = 9      // IDispatch pointer.
	VT_ERROR            VT = 10     // SCODE.
	VT_BOOL             VT = 11     // True = -1, False = 0.
	VT_VARIANT          VT = 12     // VARIANT pointer.
	VT_UNKNOWN          VT = 13     // IUnknown pointer.
	VT_DECIMAL          VT = 14     // 16 byte fixed point.
	VT_I1               VT = 16     // Signed char.
	VT_UI1              VT = 17     // Unsigned char.
	VT_UI2              VT = 18     // Unsigned short.
	VT_UI4              VT = 19     // ULONG.
	VT_I8               VT = 20     // Signed 64-bit int.
	VT_UI8              VT = 21     // Unsigned 64-bit int.
	VT_INT              VT = 22     // Signed machine int.
	VT_UINT             VT = 23     // Unsigned machine int.
	VT_VOID             VT = 24     // C style void.
	VT_HRESULT          VT = 25     // Standard return type.
	VT_PTR              VT = 26     // Pointer type.
	VT_SAFEARRAY        VT = 27     // Use VT_ARRAY in VARIANT.
	VT_CARRAY           VT = 28     // C style array.
	VT_USERDEFINED      VT = 29     // User defined type.
	VT_LPSTR            VT = 30     // Null terminated string.
	VT_LPWSTR           VT = 31     // Wide null terminated string.
	VT_RECORD           VT = 36     // User defined type.
	VT_INT_PTR          VT = 37     // Signed machine register size width.
	VT_UINT_PTR         VT = 38     // Unsigned machine register size width.
	VT_FILETIME         VT = 64     // FILETIME.
	VT_BLOB             VT = 65     // Length of prefixed bytes.
	VT_STREAM           VT = 66     // Name of the stream follows.
	VT_STORAGE          VT = 67     // Name of the storage follows.
	VT_STREAMED_OBJECT  VT = 68     // Stream contains an object.
	VT_STORED_OBJECT    VT = 69     // Storage contains an object.
	VT_BLOB_OBJECT      VT = 70     // Blob contains an object.
	VT_CF               VT = 71     // Clipboard format.
	VT_CLSID            VT = 72     // A class ID.
	VT_VERSIONED_STREAM VT = 73     // Stream with a GUID version.
	VT_BSTR_BLOB        VT = 0xfff  // Reserved for system use.
	VT_VECTOR           VT = 0x1000 // Simple counted array.
	VT_ARRAY            VT = 0x2000 // SAFEARRAY pointer.
	VT_BYREF            VT = 0x4000 // Void pointer for local use.
	VT_RESERVED         VT = 0x8000
	VT_ILLEGAL          VT = 0xffff
	VT_ILLEGALMASKED    VT = 0xfff
	VT_TYPEMASK         VT = 0xfff
)
