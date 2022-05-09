//go:build windows

package automco

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// Automation COM IIDs.
const (
	IID_IDispatch    co.IID = "00020400-0000-0000-c000-000000000046"
	IID_IErrorLog    co.IID = "3127ca40-446e-11ce-8135-00aa004bb851"
	IID_IPropertyBag co.IID = "55272a00-42cb-11ce-8135-00aa004bb851"
	IID_ITypeInfo    co.IID = "00020401-0000-0000-c000-000000000046"
)
