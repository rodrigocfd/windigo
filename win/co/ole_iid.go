//go:build windows

package co

// A COM [interface ID], represented as a string.
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
type IID string

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
