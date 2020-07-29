/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 uint64
}

type (
	CLSID GUID // COM class ID
	IID   GUID // COM interface ID
)

var (
	CLSID_FilterGraph = CLSID(GUID{0xe436ebb3, 0x524f, 0x11ce, 0x9f53_0020af0ba770})
	CLSID_TaskbarList = CLSID(GUID{0x56fdf344, 0xfd6d, 0x11d0, 0x958a_006097c9a090})
)

var (
	IID_IBaseFilter   = IID(GUID{0x56a86895, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770})
	IID_IDispatch     = IID(GUID{0x00020400, 0x0000, 0x0000, 0xc000_000000000046})
	IID_IFilterGraph  = IID(GUID{0x56a8689f, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770})
	IID_IGraphBuilder = IID(GUID{0x56a868a9, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770})
	IID_IMediaFilter  = IID(GUID{0x56a86899, 0x0ad4, 0x11ce, 0xb03a_0020af0ba770})
	IID_IPersist      = IID(GUID{0x0000010c, 0x0000, 0x0000, 0xc000_000000000046})
	IID_ITaskbarList  = IID(GUID{0x56fdf342, 0xfd6d, 0x11d0, 0x958a_006097c9a090})
	IID_ITaskbarList2 = IID(GUID{0x602d4995, 0xb13a, 0x429b, 0xa66e_1935e44f4317})
	IID_ITaskbarList3 = IID(GUID{0xea1afb91, 0x9e28, 0x4b86, 0x90e9_9e9f8a5eefaf})
	IID_IUnknown      = IID(GUID{0x00000000, 0x0000, 0x0000, 0xc000_000000000046})
)
