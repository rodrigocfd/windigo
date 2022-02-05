package win

// Specifies an RGB color.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/colorref
type COLORREF uint32

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rgb
func RGB(red, green, blue uint8) COLORREF {
	return COLORREF(uint32(red) | (uint32(green) << 8) | (uint32(blue) << 16))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getrvalue
func (c COLORREF) Red() uint8 {
	return LOBYTE(LOWORD(uint32(c)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getgvalue
func (c COLORREF) Green() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 8))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbvalue
func (c COLORREF) Blue() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 16))
}

// Converts the COLORREF to an RGBQUAD struct.
func (c COLORREF) ToRgbquad() RGBQUAD {
	rq := RGBQUAD{}
	rq.SetBlue(c.Blue())
	rq.SetGreen(c.Green())
	rq.SetRed(c.Red())
	return rq
}
