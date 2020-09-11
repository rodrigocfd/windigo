/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

// COM class ID.
type CLSID string

const (
	CLSID_FilterGraph CLSID = "e436ebb3-524f-11ce-9f53-0020af0ba770"
	CLSID_TaskbarList CLSID = "56fdf344-fd6d-11d0-958a-006097c9a090"
)

// COM interface ID.
type IID string

const (
	IID_IBaseFilter   IID = "56a86895-0ad4-11ce-b03a-0020af0ba770"
	IID_IDispatch     IID = "00020400-0000-0000-c000-000000000046"
	IID_IFilterGraph  IID = "56a8689f-0ad4-11ce-b03a-0020af0ba770"
	IID_IGraphBuilder IID = "56a868a9-0ad4-11ce-b03a-0020af0ba770"
	IID_IMediaFilter  IID = "56a86899-0ad4-11ce-b03a-0020af0ba770"
	IID_IPersist      IID = "0000010c-0000-0000-c000-000000000046"
	IID_ITaskbarList  IID = "56fdf342-fd6d-11d0-958a-006097c9a090"
	IID_ITaskbarList2 IID = "602d4995-b13a-429b-a66e-1935e44f4317"
	IID_ITaskbarList3 IID = "ea1afb91-9e28-4b86-90e9-9e9f8a5eefaf"
	IID_IUnknown      IID = "00000000-0000-0000-c000-000000000046"
)
