//go:build windows

package cowic

import (
	"github.com/rodrigocfd/windigo/co"
)

const (
	HRESULT_WINCODEC_ERR_WRONGSTATE                       co.HRESULT = 0x8898_2f04 // The codec is in the wrong state.
	HRESULT_WINCODEC_ERR_VALUEOUTOFRANGE                  co.HRESULT = 0x8898_2f05 // The value is out of range.
	HRESULT_WINCODEC_ERR_UNKNOWNIMAGEFORMAT               co.HRESULT = 0x8898_2f07 // The image format is unknown.
	HRESULT_WINCODEC_ERR_UNSUPPORTEDVERSION               co.HRESULT = 0x8898_2f0b // The SDK version is unsupported.
	HRESULT_WINCODEC_ERR_NOTINITIALIZED                   co.HRESULT = 0x8898_2f0c // The component is not initialized.
	HRESULT_WINCODEC_ERR_ALREADYLOCKED                    co.HRESULT = 0x8898_2f0d // There is already an outstanding read or write lock.
	HRESULT_WINCODEC_ERR_PROPERTYNOTFOUND                 co.HRESULT = 0x8898_2f40 // The specified bitmap property cannot be found.
	HRESULT_WINCODEC_ERR_PROPERTYNOTSUPPORTED             co.HRESULT = 0x8898_2f41 // The bitmap codec does not support the bitmap property.
	HRESULT_WINCODEC_ERR_PROPERTYSIZE                     co.HRESULT = 0x8898_2f42 // The bitmap property size is invalid.
	HRESULT_WINCODEC_ERR_CODECPRESENT                     co.HRESULT = 0x8898_2f43 // An unknown error has occurred.
	HRESULT_WINCODEC_ERR_CODECNOTHUMBNAIL                 co.HRESULT = 0x8898_2f44 // The bitmap codec does not support a thumbnail.
	HRESULT_WINCODEC_ERR_PALETTEUNAVAILABLE               co.HRESULT = 0x8898_2f45 // The bitmap palette is unavailable.
	HRESULT_WINCODEC_ERR_CODECTOOMANYSCANLINES            co.HRESULT = 0x8898_2f46 // Too many scanlines were requested.
	HRESULT_WINCODEC_ERR_INTERNALERROR                    co.HRESULT = 0x8898_2f48 // An internal error occurred.
	HRESULT_WINCODEC_ERR_SOURCERECTDOESNOTMATCHDIMENSIONS co.HRESULT = 0x8898_2f49 // The bitmap bounds do not match the bitmap dimensions.
	HRESULT_WINCODEC_ERR_COMPONENTNOTFOUND                co.HRESULT = 0x8898_2f50 // The component cannot be found.
	HRESULT_WINCODEC_ERR_IMAGESIZEOUTOFRANGE              co.HRESULT = 0x8898_2f51 // The bitmap size is outside the valid range.
	HRESULT_WINCODEC_ERR_TOOMUCHMETADATA                  co.HRESULT = 0x8898_2f52 // There is too much metadata to be written to the bitmap.
	HRESULT_WINCODEC_ERR_BADIMAGE                         co.HRESULT = 0x8898_2f60 // The image is unrecognized.
	HRESULT_WINCODEC_ERR_BADHEADER                        co.HRESULT = 0x8898_2f61 // The image header is unrecognized.
	HRESULT_WINCODEC_ERR_FRAMEMISSING                     co.HRESULT = 0x8898_2f62 // The bitmap frame is missing.
	HRESULT_WINCODEC_ERR_BADMETADATAHEADER                co.HRESULT = 0x8898_2f63 // The image metadata header is unrecognized.
	HRESULT_WINCODEC_ERR_BADSTREAMDATA                    co.HRESULT = 0x8898_2f70 // The stream data is unrecognized.
	HRESULT_WINCODEC_ERR_STREAMWRITE                      co.HRESULT = 0x8898_2f71 // Failed to write to the stream.
	HRESULT_WINCODEC_ERR_STREAMREAD                       co.HRESULT = 0x8898_2f72 // Failed to read from the stream.
	HRESULT_WINCODEC_ERR_STREAMNOTAVAILABLE               co.HRESULT = 0x8898_2f73 // The stream is not available.
	HRESULT_WINCODEC_ERR_UNSUPPORTEDPIXELFORMAT           co.HRESULT = 0x8898_2f80 // The bitmap pixel format is unsupported.
	HRESULT_WINCODEC_ERR_UNSUPPORTEDOPERATION             co.HRESULT = 0x8898_2f81 // The operation is unsupported.
	HRESULT_WINCODEC_ERR_INVALIDREGISTRATION              co.HRESULT = 0x8898_2f8a // The component registration is invalid.
	HRESULT_WINCODEC_ERR_COMPONENTINITIALIZEFAILURE       co.HRESULT = 0x8898_2f8b // The component initialization has failed.
	HRESULT_WINCODEC_ERR_INSUFFICIENTBUFFER               co.HRESULT = 0x8898_2f8c // The buffer allocated is insufficient.
	HRESULT_WINCODEC_ERR_DUPLICATEMETADATAPRESENT         co.HRESULT = 0x8898_2f8d // Duplicate metadata is present.
	HRESULT_WINCODEC_ERR_PROPERTYUNEXPECTEDTYPE           co.HRESULT = 0x8898_2f8e // The bitmap property type is unexpected.
	HRESULT_WINCODEC_ERR_UNEXPECTEDSIZE                   co.HRESULT = 0x8898_2f8f // The size is unexpected.
	HRESULT_WINCODEC_ERR_INVALIDQUERYREQUEST              co.HRESULT = 0x8898_2f90 // The property query is invalid.
	HRESULT_WINCODEC_ERR_UNEXPECTEDMETADATATYPE           co.HRESULT = 0x8898_2f91 // The metadata type is unexpected.
	HRESULT_WINCODEC_ERR_REQUESTONLYVALIDATMETADATAROOT   co.HRESULT = 0x8898_2f92 // The specified bitmap property is only valid at root level.
	HRESULT_WINCODEC_ERR_INVALIDQUERYCHARACTER            co.HRESULT = 0x8898_2f93 // The query string contains an invalid character.
	HRESULT_WINCODEC_ERR_WIN32ERROR                       co.HRESULT = 0x8898_2f94 // Windows Codecs received an error from the Win32 system.
	HRESULT_WINCODEC_ERR_INVALIDPROGRESSIVELEVEL          co.HRESULT = 0x8898_2f95 // The requested level of detail is not present.
	HRESULT_WINCODEC_ERR_INVALIDJPEGSCANINDEX             co.HRESULT = 0x8898_2f96 // The scan index is invalid.
	HRESULT_WINCODEC_ERR_UNSUPPORTEDTONEMAPPING           co.HRESULT = 0x8898_2f97 // The tone mapping mode is not supported.
)
