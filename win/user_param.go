//go:build windows

package win

// First message parameter.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#wparam
type WPARAM uintptr

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makewparam
func MAKEWPARAM(lo, hi uint16) WPARAM {
	return WPARAM(MAKELONG(lo, hi))
}

func (wp WPARAM) LoWord() uint16 { return LOWORD(uint32(wp)) }
func (wp WPARAM) HiWord() uint16 { return HIWORD(uint32(wp)) }

// Second message parameter.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#lparam
type LPARAM uintptr

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-makelparam
func MAKELPARAM(lo, hi uint16) LPARAM {
	return LPARAM(MAKELONG(lo, hi))
}

func (lp LPARAM) LoWord() uint16 { return LOWORD(uint32(lp)) }
func (lp LPARAM) HiWord() uint16 { return HIWORD(uint32(lp)) }

func (lp LPARAM) MakePoint() POINT {
	return POINT{
		X: int32(lp.LoWord()),
		Y: int32(lp.HiWord()),
	}
}

func (lp LPARAM) MakeSize() SIZE {
	return SIZE{
		Cx: int32(lp.LoWord()),
		Cy: int32(lp.HiWord()),
	}
}
