//go:build windows

package coaut

import (
	"github.com/rodrigocfd/windigo/co"
)

const (
	HRESULT_TYPE_E_BUFFERTOOSMALL          co.HRESULT = 0x8002_8016 // Buffer too small.
	HRESULT_TYPE_E_FIELDNOTFOUND           co.HRESULT = 0x8002_8017 // Field name not defined in the record.
	HRESULT_TYPE_E_INVDATAREAD             co.HRESULT = 0x8002_8018 // Old format or invalid type library.
	HRESULT_TYPE_E_UNSUPFORMAT             co.HRESULT = 0x8002_8019 // Old format or invalid type library.
	HRESULT_TYPE_E_REGISTRYACCESS          co.HRESULT = 0x8002_801c // Error accessing the OLE registry.
	HRESULT_TYPE_E_LIBNOTREGISTERED        co.HRESULT = 0x8002_801d // Library not registered.
	HRESULT_TYPE_E_UNDEFINEDTYPE           co.HRESULT = 0x8002_8027 // Bound to unknown type.
	HRESULT_TYPE_E_QUALIFIEDNAMEDISALLOWED co.HRESULT = 0x8002_8028 // Qualified name disallowed.
	HRESULT_TYPE_E_INVALIDSTATE            co.HRESULT = 0x8002_8029 // Invalid forward reference, or reference to uncompiled type.
	HRESULT_TYPE_E_WRONGTYPEKIND           co.HRESULT = 0x8002_802a // Type mismatch.
	HRESULT_TYPE_E_ELEMENTNOTFOUND         co.HRESULT = 0x8002_802b // Element not found.
	HRESULT_TYPE_E_AMBIGUOUSNAME           co.HRESULT = 0x8002_802c // Ambiguous name.
	HRESULT_TYPE_E_NAMECONFLICT            co.HRESULT = 0x8002_802d // Name already exists in the library.
	HRESULT_TYPE_E_UNKNOWNLCID             co.HRESULT = 0x8002_802e // Unknown LCID.
	HRESULT_TYPE_E_DLLFUNCTIONNOTFOUND     co.HRESULT = 0x8002_802f // Function not defined in specified DLL.
	HRESULT_TYPE_E_BADMODULEKIND           co.HRESULT = 0x8002_88bd // Wrong module kind for the operation.
	HRESULT_TYPE_E_SIZETOOBIG              co.HRESULT = 0x8002_88c5 // Size may not exceed 64K.
	HRESULT_TYPE_E_DUPLICATEID             co.HRESULT = 0x8002_88c6 // Duplicate ID in inheritance hierarchy.
	HRESULT_TYPE_E_INVALIDID               co.HRESULT = 0x8002_88cf // Incorrect inheritance depth in standard OLE hmember.
	HRESULT_TYPE_E_TYPEMISMATCH            co.HRESULT = 0x8002_8ca0 // Type mismatch.
	HRESULT_TYPE_E_OUTOFBOUNDS             co.HRESULT = 0x8002_8ca1 // Invalid number of arguments.
	HRESULT_TYPE_E_IOERROR                 co.HRESULT = 0x8002_8ca2 // I/O Error.
	HRESULT_TYPE_E_CANTCREATETMPFILE       co.HRESULT = 0x8002_8ca3 // Error creating unique tmp file.
	HRESULT_TYPE_E_CANTLOADLIBRARY         co.HRESULT = 0x8002_9c4a // Error loading type library/DLL.
	HRESULT_TYPE_E_INCONSISTENTPROPFUNCS   co.HRESULT = 0x8002_9c83 // Inconsistent property functions.
	HRESULT_TYPE_E_CIRCULARTYPE            co.HRESULT = 0x8002_9c84 // Circular dependency between types/modules.
)
