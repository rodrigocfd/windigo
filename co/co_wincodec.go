//go:build windows

package co

// Wincodec CLSID identifier.
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
)

// Wincodec IID identifier.
const (
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

// [WICComponentSigning] enumeration.
//
// [WICComponentSigning]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ne-wincodec-wiccomponentsigning
type WIC_COMPONENTSIGN uint32

const (
	WIC_COMPONENTSIGN_Signed   WIC_COMPONENTSIGN = 0x1
	WIC_COMPONENTSIGN_Unsigned WIC_COMPONENTSIGN = 0x2
	WIC_COMPONENTSIGN_Safe     WIC_COMPONENTSIGN = 0x4
	WIC_COMPONENTSIGN_Disabled WIC_COMPONENTSIGN = 0x8000_0000
)

// WIC [container format] [GUID], represented as a string.
//
// [container format]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/nf-wincodec-iwicimagingfactory-createencoder
type WIC_CONTAINER GUID

const (
	WIC_CONTAINER_Bmp    WIC_CONTAINER = "0af1d87e-fcfe-4188-bdeb-a7906471cbe3"
	WIC_CONTAINER_Png    WIC_CONTAINER = "1b7cfaf4-713f-473c-bbcd-6137425faeaf"
	WIC_CONTAINER_Ico    WIC_CONTAINER = "a3a860c4-338f-4c17-919a-fba4b5628f21"
	WIC_CONTAINER_Jpeg   WIC_CONTAINER = "19e4a5aa-5662-4fc5-a0c0-1758028e1057"
	WIC_CONTAINER_Tiff   WIC_CONTAINER = "163bcc30-e2e9-4f0b-961d-a3e9fdb788a3"
	WIC_CONTAINER_Gif    WIC_CONTAINER = "1f8a5601-7d4d-4cbd-9c82-1bc8d4eeb9a5"
	WIC_CONTAINER_Wmp    WIC_CONTAINER = "57a37caa-367a-4540-916b-f183c5093a4b"
	WIC_CONTAINER_Dds    WIC_CONTAINER = "9967cb95-2e85-4ac8-8ca2-83d7ccd425c9"
	WIC_CONTAINER_Adng   WIC_CONTAINER = "f3ff6d0d-38c0-41c4-b1fe-1f3824f17b84"
	WIC_CONTAINER_Heif   WIC_CONTAINER = "e1e62521-6787-405b-a339-500715b5763f"
	WIC_CONTAINER_Webp   WIC_CONTAINER = "e094b0e2-67f2-45b3-b0ea-115337ca7cf3"
	WIC_CONTAINER_Raw    WIC_CONTAINER = "fe99ce60-f19c-433c-a3ae-00acefa9ca21"
	WIC_CONTAINER_JpegXL WIC_CONTAINER = "fec14e3f-427a-4736-aae6-27ed84f69322"
)

// REFWICPixelFormatGUID, the WIC pixel format [GUID], represented as a string.
type WIC_PIXELFORMAT GUID

const (
	WIC_PIXELFORMAT_DontCare                        WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc900"
	WIC_PIXELFORMAT_1bppIndexed                     WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc901"
	WIC_PIXELFORMAT_2bppIndexed                     WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc902"
	WIC_PIXELFORMAT_4bppIndexed                     WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc903"
	WIC_PIXELFORMAT_8bppIndexed                     WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc904"
	WIC_PIXELFORMAT_BlackWhite                      WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc905"
	WIC_PIXELFORMAT_2bppGray                        WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc906"
	WIC_PIXELFORMAT_4bppGray                        WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc907"
	WIC_PIXELFORMAT_8bppGray                        WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc908"
	WIC_PIXELFORMAT_8bppAlpha                       WIC_PIXELFORMAT = "e6cd0116-eeba-4161-aa85-27dd9fb3a895"
	WIC_PIXELFORMAT_8bppDepth                       WIC_PIXELFORMAT = "4c9c9f45-1d89-4e31-9bc7-69343a0dca69"
	WIC_PIXELFORMAT_8bppGain                        WIC_PIXELFORMAT = "a884022a-af13-4c16-b746-619bf618b878"
	WIC_PIXELFORMAT_24bppRGBGain                    WIC_PIXELFORMAT = "a5022b24-7109-443b-9948-25b6ed8f39fd"
	WIC_PIXELFORMAT_32bppBGRGain                    WIC_PIXELFORMAT = "837d6738-208a-43e0-8995-79ab74407402"
	WIC_PIXELFORMAT_16bppBGR555                     WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc909"
	WIC_PIXELFORMAT_16bppBGR565                     WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc90a"
	WIC_PIXELFORMAT_16bppBGRA5551                   WIC_PIXELFORMAT = "05ec7c2b-f1e6-4961-ad46-e1cc810a87d2"
	WIC_PIXELFORMAT_16bppGray                       WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc90b"
	WIC_PIXELFORMAT_24bppBGR                        WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc90c"
	WIC_PIXELFORMAT_24bppRGB                        WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc90d"
	WIC_PIXELFORMAT_32bppBGR                        WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc90e"
	WIC_PIXELFORMAT_32bppBGRA                       WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc90f"
	WIC_PIXELFORMAT_32bppPBGRA                      WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc910"
	WIC_PIXELFORMAT_32bppGrayFloat                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc911"
	WIC_PIXELFORMAT_32bppRGB                        WIC_PIXELFORMAT = "d98c6b95-3efe-47d6-bb25-eb1748ab0cf1"
	WIC_PIXELFORMAT_32bppRGBA                       WIC_PIXELFORMAT = "f5c7ad2d-6a8d-43dd-a7a8-a29935261ae9"
	WIC_PIXELFORMAT_32bppPRGBA                      WIC_PIXELFORMAT = "3cc4a650-a527-4d37-a916-3142c7ebedba"
	WIC_PIXELFORMAT_48bppRGB                        WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc915"
	WIC_PIXELFORMAT_48bppBGR                        WIC_PIXELFORMAT = "e605a384-b468-46ce-bb2e-36f180e64313"
	WIC_PIXELFORMAT_64bppRGB                        WIC_PIXELFORMAT = "a1182111-186d-4d42-bc6a-9c8303a8dff9"
	WIC_PIXELFORMAT_64bppRGBA                       WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc916"
	WIC_PIXELFORMAT_64bppBGRA                       WIC_PIXELFORMAT = "1562ff7c-d352-46f9-979e-42976b792246"
	WIC_PIXELFORMAT_64bppPRGBA                      WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc917"
	WIC_PIXELFORMAT_64bppPBGRA                      WIC_PIXELFORMAT = "8c518e8e-a4ec-468b-ae70-c9a35a9c5530"
	WIC_PIXELFORMAT_16bppGrayFixedPoint             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc913"
	WIC_PIXELFORMAT_32bppBGR101010                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc914"
	WIC_PIXELFORMAT_48bppRGBFixedPoint              WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc912"
	WIC_PIXELFORMAT_48bppBGRFixedPoint              WIC_PIXELFORMAT = "49ca140e-cab6-493b-9ddf-60187c37532a"
	WIC_PIXELFORMAT_96bppRGBFixedPoint              WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc918"
	WIC_PIXELFORMAT_96bppRGBFloat                   WIC_PIXELFORMAT = "e3fed78f-e8db-4acf-84c1-e97f6136b327"
	WIC_PIXELFORMAT_128bppRGBAFloat                 WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc919"
	WIC_PIXELFORMAT_128bppPRGBAFloat                WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc91a"
	WIC_PIXELFORMAT_128bppRGBFloat                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc91b"
	WIC_PIXELFORMAT_32bppCMYK                       WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc91c"
	WIC_PIXELFORMAT_64bppRGBAFixedPoint             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc91d"
	WIC_PIXELFORMAT_64bppBGRAFixedPoint             WIC_PIXELFORMAT = "356de33c-54d2-4a23-bb4-9b7bf9b1d42d"
	WIC_PIXELFORMAT_64bppRGBFixedPoint              WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc940"
	WIC_PIXELFORMAT_128bppRGBAFixedPoint            WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc91e"
	WIC_PIXELFORMAT_128bppRGBFixedPoint             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc941"
	WIC_PIXELFORMAT_64bppRGBAHalf                   WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc93a"
	WIC_PIXELFORMAT_64bppPRGBAHalf                  WIC_PIXELFORMAT = "58ad26c2-c623-4d9d-b320-387e49f8c442"
	WIC_PIXELFORMAT_64bppRGBHalf                    WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc942"
	WIC_PIXELFORMAT_48bppRGBHalf                    WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc93b"
	WIC_PIXELFORMAT_32bppRGBE                       WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc93d"
	WIC_PIXELFORMAT_16bppGrayHalf                   WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc93e"
	WIC_PIXELFORMAT_32bppGrayFixedPoint             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc93f"
	WIC_PIXELFORMAT_32bppRGBA1010102                WIC_PIXELFORMAT = "25238d72-fcf9-4522-b514-5578e5ad55e0"
	WIC_PIXELFORMAT_32bppRGBA1010102XR              WIC_PIXELFORMAT = "00de6b9a-c101-434b-b502-d0165ee1122c"
	WIC_PIXELFORMAT_32bppR10G10B10A2                WIC_PIXELFORMAT = "604e1bb5-8a3c-4b65-b11c-bc0b8dd75b7f"
	WIC_PIXELFORMAT_32bppR10G10B10A2HDR10           WIC_PIXELFORMAT = "9c215c5d-1acc-4f0e-a4bc-70fb3ae8fd28"
	WIC_PIXELFORMAT_64bppCMYK                       WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc91f"
	WIC_PIXELFORMAT_24bpp3Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc920"
	WIC_PIXELFORMAT_32bpp4Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc921"
	WIC_PIXELFORMAT_40bpp5Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc922"
	WIC_PIXELFORMAT_48bpp6Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc923"
	WIC_PIXELFORMAT_56bpp7Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc924"
	WIC_PIXELFORMAT_64bpp8Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc925"
	WIC_PIXELFORMAT_48bpp3Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc926"
	WIC_PIXELFORMAT_64bpp4Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc927"
	WIC_PIXELFORMAT_80bpp5Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc928"
	WIC_PIXELFORMAT_96bpp6Channels                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc929"
	WIC_PIXELFORMAT_112bpp7Channels                 WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc92a"
	WIC_PIXELFORMAT_128bpp8Channels                 WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc92b"
	WIC_PIXELFORMAT_40bppCMYKAlpha                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc92c"
	WIC_PIXELFORMAT_80bppCMYKAlpha                  WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc92d"
	WIC_PIXELFORMAT_32bpp3ChannelsAlpha             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc92e"
	WIC_PIXELFORMAT_40bpp4ChannelsAlpha             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc92f"
	WIC_PIXELFORMAT_48bpp5ChannelsAlpha             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc930"
	WIC_PIXELFORMAT_56bpp6ChannelsAlpha             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc931"
	WIC_PIXELFORMAT_64bpp7ChannelsAlpha             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc932"
	WIC_PIXELFORMAT_72bpp8ChannelsAlpha             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc933"
	WIC_PIXELFORMAT_64bpp3ChannelsAlpha             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc934"
	WIC_PIXELFORMAT_80bpp4ChannelsAlpha             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc935"
	WIC_PIXELFORMAT_96bpp5ChannelsAlpha             WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc936"
	WIC_PIXELFORMAT_112bpp6ChannelsAlpha            WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc937"
	WIC_PIXELFORMAT_128bpp7ChannelsAlpha            WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc938"
	WIC_PIXELFORMAT_144bpp8ChannelsAlpha            WIC_PIXELFORMAT = "6fddc324-4e03-4bfe-b185-3d77768dc939"
	WIC_PIXELFORMAT_8bppY                           WIC_PIXELFORMAT = "91b4db54-2df9-42f0-b449-2909bb3df88e"
	WIC_PIXELFORMAT_8bppCb                          WIC_PIXELFORMAT = "1339f224-6bfe-4c3e-9302-e4f3a6d0ca2a"
	WIC_PIXELFORMAT_8bppCr                          WIC_PIXELFORMAT = "b8145053-2116-49f0-8835-ed844b205c51"
	WIC_PIXELFORMAT_16bppCbCr                       WIC_PIXELFORMAT = "ff95ba6e-11e0-4263-bb45-01721f3460a4"
	WIC_PIXELFORMAT_16bppYQuantizedDctCoefficients  WIC_PIXELFORMAT = "a355f433-48e8-4a42-84d8-e2aa26ca80a4"
	WIC_PIXELFORMAT_16bppCbQuantizedDctCoefficients WIC_PIXELFORMAT = "d2c4ff61-56a5-49c2-8b5c-4c1925964837"
	WIC_PIXELFORMAT_16bppCrQuantizedDctCoefficients WIC_PIXELFORMAT = "2fe354f0-1680-42d8-9231-e73c0565bfc1"
)

// [WICBitmapAlphaChannelOption] enumeration.
//
// [WICBitmapAlphaChannelOption]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ne-wincodec-wicbitmapalphachanneloption
type WICBMP_ALPHACH uint32

const (
	WICBMP_ALPHACH_UseAlpha              WICBMP_ALPHACH = 0
	WICBMP_ALPHACH_UsePremultipliedAlpha WICBMP_ALPHACH = 0x1
	WICBMP_ALPHACH_IgnoreAlpha           WICBMP_ALPHACH = 0x2
)

// [WICBitmapCreateCacheOption] enumeration.
//
// [WICBitmapCreateCacheOption]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ne-wincodec-wicbitmapcreatecacheoption
type WICBMP_CACHE uint32

const (
	WICBMP_CACHE_No       WICBMP_CACHE = 0
	WICBMP_CACHE_OnDemand WICBMP_CACHE = 0x1
	WICBMP_CACHE_OnLoad   WICBMP_CACHE = 0x2
)

// [WICBitmapDitherType] enumeration.
//
// [WICBitmapDitherType]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ne-wincodec-wicbitmapdithertype
type WICBMP_DITHER uint32

const (
	WICBMP_DITHER_None           WICBMP_DITHER = 0
	WICBMP_DITHER_Solid          WICBMP_DITHER = 0
	WICBMP_DITHER_Ordered4x4     WICBMP_DITHER = 0x1
	WICBMP_DITHER_Ordered8x8     WICBMP_DITHER = 0x2
	WICBMP_DITHER_Ordered16x16   WICBMP_DITHER = 0x3
	WICBMP_DITHER_Spiral4x4      WICBMP_DITHER = 0x4
	WICBMP_DITHER_Spiral8x8      WICBMP_DITHER = 0x5
	WICBMP_DITHER_DualSpiral4x4  WICBMP_DITHER = 0x6
	WICBMP_DITHER_DualSpiral8x8  WICBMP_DITHER = 0x7
	WICBMP_DITHER_ErrorDiffusion WICBMP_DITHER = 0x8
)

// [WICBitmapLockFlags] enumeration.
//
// [WICBitmapLockFlags]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ne-wincodec-wicbitmaplockflags
type WICBMP_LOCK uint32

const (
	WICBMP_LOCK_Read  WICBMP_LOCK = 0x1
	WICBMP_LOCK_Write WICBMP_LOCK = 0x2
)

// [WICBitmapPaletteType] enumeration.
//
// [WICBitmapPaletteType]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ne-wincodec-wicbitmappalettetype
type WICBMP_PAL uint32

const (
	WICBMP_PAL_Custom           WICBMP_PAL = 0
	WICBMP_PAL_MedianCut        WICBMP_PAL = 0x1
	WICBMP_PAL_FixedBW          WICBMP_PAL = 0x2
	WICBMP_PAL_FixedHalftone8   WICBMP_PAL = 0x3
	WICBMP_PAL_FixedHalftone27  WICBMP_PAL = 0x4
	WICBMP_PAL_FixedHalftone64  WICBMP_PAL = 0x5
	WICBMP_PAL_FixedHalftone125 WICBMP_PAL = 0x6
	WICBMP_PAL_FixedHalftone216 WICBMP_PAL = 0x7
	WICBMP_PAL_FixedWebPalette             = WICBMP_PAL_FixedHalftone216
	WICBMP_PAL_FixedHalftone252 WICBMP_PAL = 0x8
	WICBMP_PAL_FixedHalftone256 WICBMP_PAL = 0x9
	WICBMP_PAL_FixedGray4       WICBMP_PAL = 0xa
	WICBMP_PAL_FixedGray16      WICBMP_PAL = 0xb
	WICBMP_PAL_FixedGray256     WICBMP_PAL = 0xc
)

// [WICDecodeOptions] enumeration.
//
// [WICDecodeOptions]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ne-wincodec-wicdecodeoptions
type WICDEC_METADATACACHE uint32

const (
	WICDEC_METADATACACHE_OnDemand WICDEC_METADATACACHE = 0
	WICDEC_METADATACACHE_OnLoad   WICDEC_METADATACACHE = 0x1
)

// [WICBitmapDecoderCapabilities] enumeration.
//
// [WICBitmapDecoderCapabilities]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ne-wincodec-wicbitmapdecodercapabilities
type WICDEC_CAP uint32

const (
	WICDEC_CAP_SameEncoder          WICDEC_CAP = 0x1
	WICDEC_CAP_CanDecodeAllImages   WICDEC_CAP = 0x2
	WICDEC_CAP_CanDecodeSomeImages  WICDEC_CAP = 0x4
	WICDEC_CAP_CanEnumerateMetadata WICDEC_CAP = 0x8
	WICDEC_CAP_CanDecodeThumbnail   WICDEC_CAP = 0x10
)

// [WICBitmapEncoderCacheOption] enumeration.
//
// [WICBitmapEncoderCacheOption]: https://learn.microsoft.com/en-us/windows/win32/api/wincodec/ne-wincodec-wicbitmapencodercacheoption
type WICENC_CACHE uint32

const (
	WICENC_CACHE_InMemory WICENC_CACHE = 0
	WICENC_CACHE_TempFile WICENC_CACHE = 0x1
	WICENC_CACHE_No       WICENC_CACHE = 0x2
)
