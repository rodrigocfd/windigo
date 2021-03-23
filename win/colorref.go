package win

// Specifies an RGB color.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/colorref
type COLORREF uint32

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getrvalue
func (c COLORREF) GetRValue() uint8 {
	return LOBYTE(LOWORD(uint32(c)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getgvalue
func (c COLORREF) GetGValue() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 8))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbvalue
func (c COLORREF) GetBValue() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 16))
}
