package co

// A COM interface class ID, represented as a string.
type CLSID string

// A COM interface ID, represented as a string.
type IID string

const (
	IID_IUnknown  IID = "00000000-0000-0000-c000-000000000046"
	IID_IDispatch IID = "00020400-0000-0000-c000-000000000046"
)
