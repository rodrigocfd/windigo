package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// Variant type for a class name identifier.
type ClassName struct {
	curType uint8  // 0: none
	atom    ATOM   // 1
	str     string // 2
}

// Creates a new ClassName variant with an empty value.
func ClassNameNone() ClassName {
	return ClassName{}
}

// Creates a new ClassName variant with an ATOM value.
func ClassNameAtom(atom ATOM) ClassName {
	return ClassName{
		curType: 1,
		atom:    atom,
	}
}

// Creates a new ClassName variant with a string value.
func ClassNameStr(str string) ClassName {
	return ClassName{
		curType: 2,
		str:     str,
	}
}

func (me *ClassName) IsNone() bool        { return me.curType == 0 }
func (me *ClassName) Atom() (ATOM, bool)  { return me.atom, me.curType == 1 }
func (me *ClassName) Str() (string, bool) { return me.str, me.curType == 2 }

func (me *ClassName) raw() (uintptr, *uint16) { // pointer must be kept alive
	switch me.curType {
	case 0:
		return 0, nil
	case 1:
		return uintptr(me.atom), nil
	case 2:
		buf := Str.ToNativePtr(me.str)
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("Invalid ClassName value.")
	}
}

//------------------------------------------------------------------------------

// Variant type for a cursor resource identifier.
type CursorRes struct {
	curType uint8
	idc     co.IDC // 1
	id      uint16 // 2
	str     string // 3
}

// Creates a new CursorRes variant with a co.IDC value.
func CursorResIdc(idc co.IDC) CursorRes {
	return CursorRes{
		curType: 1,
		idc:     idc,
	}
}

// Creates a new CursorRes variant with an int value.
func CursorResInt(id int) CursorRes {
	return CursorRes{
		curType: 2,
		id:      uint16(id),
	}
}

// Creates a new CursorRes variant with a string value.
func CursorResStr(str string) CursorRes {
	return CursorRes{
		curType: 3,
		str:     str,
	}
}

func (me *CursorRes) Idc() (co.IDC, bool) { return me.idc, me.curType == 1 }
func (me *CursorRes) Id() (int, bool)     { return int(me.id), me.curType == 2 }
func (me *CursorRes) Str() (string, bool) { return me.str, me.curType == 3 }

func (me *CursorRes) raw() (uintptr, *uint16) { // pointer must be kept alive
	switch me.curType {
	case 1:
		return uintptr(me.idc), nil
	case 2:
		return uintptr(me.id), nil
	case 3:
		buf := Str.ToNativePtr(me.str)
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("Invalid CursorRes value.")
	}
}

//------------------------------------------------------------------------------

// Variant type for an icon resource identifier.
type IconRes struct {
	curType uint8
	idi     co.IDI // 1
	id      uint16 // 2
	str     string // 3
}

// Creates a new IconRes variant with a co.IDI value.
func IconResIdi(idi co.IDI) IconRes {
	return IconRes{
		curType: 1,
		idi:     idi,
	}
}

// Creates a new IconRes variant with an int value.
func IconResInt(id int) IconRes {
	return IconRes{
		curType: 2,
		id:      uint16(id),
	}
}

// Creates a new IconRes variant with a string value.
func IconResStr(str string) IconRes {
	return IconRes{
		curType: 3,
		str:     str,
	}
}

func (me *IconRes) Idi() (co.IDI, bool) { return me.idi, me.curType == 1 }
func (me *IconRes) Id() (int, bool)     { return int(me.id), me.curType == 2 }
func (me *IconRes) Str() (string, bool) { return me.str, me.curType == 3 }

func (me *IconRes) raw() (uintptr, *uint16) { // pointer must be kept alive
	switch me.curType {
	case 1:
		return uintptr(me.idi), nil
	case 2:
		return uintptr(me.id), nil
	case 3:
		buf := Str.ToNativePtr(me.str)
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("Invalid IconRes value.")
	}
}

//------------------------------------------------------------------------------

// Variant type for a menu item identifier.
type MenuItem struct {
	curType uint8
	n       uint16 // 1: cmd, 2: pos
}

// Creates a new MenuItem variant with a command ID.
func MenuItemCmd(cmdId int) MenuItem {
	return MenuItem{
		curType: 1,
		n:       uint16(cmdId),
	}
}

// Creates a new MenuItem variant with a zero-based item index.
func MenuItemPos(pos int) MenuItem {
	return MenuItem{
		curType: 2,
		n:       uint16(pos),
	}
}

func (me *MenuItem) Cmd() (int, bool) { return int(me.n), me.curType == 1 }
func (me *MenuItem) Pos() (int, bool) { return int(me.n), me.curType == 2 }
func (me *MenuItem) Flag() co.MF {
	return util.Iif(me.curType == 1, co.MF_BYCOMMAND, co.MF_BYPOSITION).(co.MF)
}

func (me *MenuItem) raw() (uintptr, co.MF) {
	switch me.curType {
	case 1:
		return uintptr(me.n), co.MF_BYCOMMAND
	case 2:
		return uintptr(me.n), co.MF_BYPOSITION
	default:
		panic("Invalid MenuItem value.")
	}
}

//------------------------------------------------------------------------------

// Variant type for a resource identifier.
type ResId struct {
	curType uint8
	id      uint16 // 1
	str     string // 2
}

// Creates a new ResId variant with an int value.
func ResIdInt(id int) ResId {
	return ResId{
		curType: 1,
		id:      uint16(id),
	}
}

// Creates a new ResId variant with a string value.
func ResIdStr(str string) ResId {
	return ResId{
		curType: 2,
		str:     str,
	}
}

func (me *ResId) Id() (int, bool)     { return int(me.id), me.curType == 1 }
func (me *ResId) Str() (string, bool) { return me.str, me.curType == 2 }

func (me *ResId) raw() (uintptr, *uint16) { // pointer must be kept alive
	switch me.curType {
	case 1:
		return uintptr(me.id), nil
	case 2:
		buf := Str.ToNativePtr(me.str)
		return uintptr(unsafe.Pointer(buf)), buf
	default:
		panic("Invalid ResId value.")
	}
}
