//go:build windows

package co

// A [GUID] struct, represented as a string.
//
// [GUID]: https://learn.microsoft.com/en-us/windows/win32/api/guiddef/ns-guiddef-guid
type GUID string

// A COM [class ID], represented as a string.
//
// [class ID]: https://learn.microsoft.com/en-us/windows/win32/com/clsid-key-hklm
type CLSID GUID

// A COM [interface ID], represented as a string.
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
type IID GUID

// Ole GUID identifier.
const (
	IID_IBindCtx          IID = "0000000e-0000-0000-c000-000000000046"
	IID_IDataObject       IID = "0000010e-0000-0000-c000-000000000046"
	IID_IDropTarget       IID = "00000122-0000-0000-c000-000000000046"
	IID_IEnumString       IID = "00000101-0000-0000-c000-000000000046"
	IID_ISequentialStream IID = "0c733a30-2a1c-11ce-ade5-00aa0044773d"
	IID_IStream           IID = "0000000c-0000-0000-c000-000000000046"
	IID_IUnknown          IID = "00000000-0000-0000-c000-000000000046"
	IID_NULL              IID = "00000000-0000-0000-0000-000000000000"
)

// Oleaut GUID identifier.
const (
	IID_IDispatch      IID = "00020400-0000-0000-c000-000000000046"
	IID_IPicture       IID = "7bf80980-bf32-101a-8bbb-00aa00300cab"
	IID_IPropertyStore IID = "886d8eeb-8cf2-4446-8d02-cdba1dbdcf99"
	IID_ITypeInfo      IID = "00020401-0000-0000-c000-000000000046"
	IID_ITypeLib       IID = "00020402-0000-0000-c000-000000000046"
)

// Shell GUID identifier.
const (
	CLSID_FileOpenDialog CLSID = "dc1c5a9c-e88a-4dde-a5a1-60f82a20aef7"
	CLSID_FileOperation  CLSID = "3ad05575-8857-4850-9277-11b85bdb8e09"
	CLSID_FileSaveDialog CLSID = "c0b4e2f3-ba21-4773-8dba-335ec946eb8b"
	CLSID_ShellLink      CLSID = "00021401-0000-0000-c000-000000000046"
	CLSID_TaskbarList    CLSID = "56fdf344-fd6d-11d0-958a-006097c9a090"

	IID_IEnumIDList                IID = "000214f2-0000-0000-c000-000000000046"
	IID_IEnumShellItems            IID = "70629033-e363-4a28-a567-0db78006e6d7"
	IID_IFileDialog                IID = "42f85136-db7e-439c-85f1-e4075d135fc8"
	IID_IFileDialogEvents          IID = "973510db-7d7f-452b-8975-74a85828d354"
	IID_IFileOpenDialog            IID = "d57c7288-d4ad-4768-be02-9d969532d960"
	IID_IFileOperation             IID = "947aab5f-0a5c-4c13-b4d6-4bf7836fc9f8"
	IID_IFileOperationProgressSink IID = "04b0f1a7-9490-44bc-96e1-4296a31252e2"
	IID_IFileSaveDialog            IID = "84bccd23-5fde-4cdb-aea4-af64b83d78ab"
	IID_IModalWindow               IID = "b4db1657-70d7-485e-8e3e-6fcb5a5c1802"
	IID_IOleWindow                 IID = "00000114-0000-0000-c000-000000000046"
	IID_IShellFolder               IID = "000214e6-0000-0000-c000-000000000046"
	IID_IShellItem                 IID = "43826d1e-e718-42ee-bc55-a1e261c37bfe"
	IID_IShellItem2                IID = "7e9fb0d3-919f-4307-ab2e-9b1860310c93"
	IID_IShellItemArray            IID = "b63ea76d-1f85-456f-a19c-48159efa858b"
	IID_IShellItemFilter           IID = "2659b475-eeb8-48b7-8f07-b378810f48cf"
	IID_IShellLink                 IID = "000214f9-0000-0000-c000-000000000046"
	IID_IShellView                 IID = "000214e3-0000-0000-c000-000000000046"
	IID_ITaskbarList               IID = "56fdf342-fd6d-11d0-958a-006097c9a090"
	IID_ITaskbarList2              IID = "602d4995-b13a-429b-a66e-1935e44f4317"
	IID_ITaskbarList3              IID = "ea1afb91-9e28-4b86-90e9-9e9f8a5eefaf"
	IID_ITaskbarList4              IID = "c43dc798-95d1-4bea-9030-bb99e2983a1a"
)

// Wincodec GUID identifier.
const (
	_CLSID_WICImagingFactory2         CLSID = "317d06e8-5f24-433d-bdf7-79ce68d8abc2"
	_CLSID_WICPngDecoder2             CLSID = "e018945b-aa86-4008-9bd4-6777a1e40c11"
	CLSID_WICBmpDecoder               CLSID = "6b462062-7cbf-400d-9fdb-813dd10f2778"
	CLSID_WICDdsDecoder               CLSID = "9053699f-a341-429d-9e90-ee437cf80c73"
	CLSID_WICGifDecoder               CLSID = "381dda3c-9ce9-4834-a23e-1f98f8fc52be"
	CLSID_WICIcoDecoder               CLSID = "c61bfcdf-2e0f-4aad-a8d7-e06bafebcdfe"
	CLSID_WICImagingFactory                 = _CLSID_WICImagingFactory2
	CLSID_WICJpegDecoder              CLSID = "9456a480-e88b-43ea-9e73-0b2d9b71b1ca"
	CLSID_WICPngDecoder                     = _CLSID_WICPngDecoder2
	CLSID_WICTiffDecoder              CLSID = "b54e85d9-fe23-499f-8b88-6acea713752b"
	CLSID_WICWmpDecoder               CLSID = "a26cec36-234c-4950-ae16-e34aace71d0d"
	CLSID_WICBmpEncoder               CLSID = "69be8bb4-d66d-47c8-865a-ed1589433782"
	CLSID_WICPngEncoder               CLSID = "27949969-876a-41d7-9447-568f6a35a4dc"
	CLSID_WICJpegEncoder              CLSID = "1a34f5c1-4a5a-46dc-b644-1f4567e7a676"
	CLSID_WICGifEncoder               CLSID = "114f5598-0b22-40a0-86a1-c83ea495adbd"
	CLSID_WICTiffEncoder              CLSID = "0131be10-2001-4c5f-a9b0-cc88fab64ce8"
	CLSID_WICWmpEncoder               CLSID = "ac4ce3cb-e1c1-44cd-8215-5a1665509ec2"
	CLSID_WICDdsEncode                CLSID = "a61dde94-66ce-4ac1-881b-71680588895e"
	CLSID_WICAdngDecoder              CLSID = "981d9411-909e-42a7-8f5d-a747ff052edb"
	CLSID_WICJpegQualcommPhoneEncoder CLSID = "68ed5c62-f534-4979-b2b3-686a12b2b34c"
	CLSID_WICHeifDecoder              CLSID = "e9a4a80a-44fe-4de4-8971-7150b10a5199"
	CLSID_WICHeifEncoder              CLSID = "0dbecec1-9eb3-4860-9c6f-ddbe86634575"
	CLSID_WICWebpDecoder              CLSID = "7693e886-51c9-4070-8419-9f70738ec8fa"
	CLSID_WICRAWDecoder               CLSID = "41945702-8302-44a6-9445-ac98e8afa086"
	CLSID_WICJpegXLDecoder            CLSID = "fc6ceece-aef5-4a23-96ec-5984ffb486d9"
	CLSID_WICJpegXLEncoder            CLSID = "0e4ecd3b-1ba6-4636-8198-56c73040964a"

	IID_IWICBitmap            IID = "00000121-a8f2-4877-ba0a-fd2b6645fb94"
	IID_IWICBitmapCodecInfo   IID = "e87a44c4-b76e-4c47-8b09-298eb12a2714"
	IID_IWICBitmapDecoder     IID = "9edde9e7-8dee-47ea-99df-e6faf2ed44bf"
	IID_IWICBitmapDecoderInfo IID = "d8cd007f-d08f-4191-9bfc-236ea7f0e4b5"
	IID_IWICBitmapEncoder     IID = "00000103-a8f2-4877-ba0a-fd2b6645fb94"
	IID_IWICBitmapEncoderInfo IID = "94c9b4ee-a09f-4f92-8a1e-4a9bce7e76fb"
	IID_IWICBitmapFrameDecode IID = "3b16811b-6a43-4ec9-a813-3d930c13b940"
	IID_IWICBitmapLock        IID = "00000123-a8f2-4877-ba0a-fd2b6645fb94"
	IID_IWICBitmapSource      IID = "00000120-a8f2-4877-ba0a-fd2b6645fb94"
	IID_IWICComponentInfo     IID = "23bc3f0a-698b-4357-886b-f24d50671334"
	IID_IWICFormatConverter   IID = "00000301-a8f2-4877-ba0a-fd2b6645fb94"
	IID_IWICImagingFactory    IID = "ec5ec8a9-c395-4314-9c77-54d7a935ff70"
	IID_IWICPalette           IID = "00000040-a8f2-4877-ba0a-fd2b6645fb94"
	IID_IWICStream            IID = "135ff860-22b7-4ddf-b0f6-218f4f299a43"
)
