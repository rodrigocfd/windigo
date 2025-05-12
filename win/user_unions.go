//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// Tagged union for a class name identifier, which can be:
//   - none
//   - [ATOM]
//   - string
//
// # Example
//
//	clsName := win.ClassNameStr("FOO")
//
//	if s, ok := clsName.Str(); ok {
//		println(s)
//	}
type ClassName struct {
	tag  _ClassNameTag
	atom ATOM
	str  string
}

type _ClassNameTag uint8

const (
	_ClassNameTag_none _ClassNameTag = 0x0
	_ClassNameTag_atom _ClassNameTag = 0x1
	_ClassNameTag_str  _ClassNameTag = 0x2
)

// Creates a new [ClassName] with an empty value.
func ClassNameNone() ClassName {
	return ClassName{
		tag: _ClassNameTag_none,
	}
}

// Returns true if there is no value.
func (me *ClassName) IsNone() bool {
	return me.tag == _ClassNameTag_none
}

// Creates a new [ClassName] with an [ATOM] value.
func ClassNameAtom(atom ATOM) ClassName {
	return ClassName{
		tag:  _ClassNameTag_atom,
		atom: atom,
	}
}

// If the value is an [ATOM], returns it and true.
func (me *ClassName) Atom() (ATOM, bool) {
	return me.atom, me.tag == _ClassNameTag_atom
}

// Creates a new [ClassName] with a string value.
func ClassNameStr(str string) ClassName {
	return ClassName{
		tag: _ClassNameTag_str,
		str: str,
	}
}

// If the value is a string, returns it and true.
func (me *ClassName) Str() (string, bool) {
	return me.str, me.tag == _ClassNameTag_str
}

// Converts the internal value to uintptr.
func (me *ClassName) raw(wideBuf *wstr.Buf[wstr.Stack20]) uintptr {
	switch me.tag {
	case _ClassNameTag_none:
		return 0
	case _ClassNameTag_atom:
		return uintptr(me.atom)
	case _ClassNameTag_str:
		wideBuf.Set(me.str, wstr.EMPTY_IS_NIL)
		return uintptr(wideBuf.UnsafePtr())
	default:
		panic("Invalid ClassName value.")
	}
}

// Tagged union for a [cursor resource] identifier, which can be:
//   - [co.IDC]
//   - uint16
//   - string
//
// # Example
//
//	curId := win.CursorResIdc(co.IDC_ARROW)
//
//	if idc, ok := curId.Idc(); ok {
//		println(idc)
//	}
//
// [cursor resource]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
type CursorRes struct {
	tag  _CursorResTag
	data uintptr
	str  string
}

type _CursorResTag uint8

const (
	_CursorTag_idc _CursorResTag = 0x1
	_CursorTag_id  _CursorResTag = 0x2
	_CursorTag_str _CursorResTag = 0x3
)

// Creates a new [CursorRes] with a [co.IDC] value.
func CursorResIdc(idc co.IDC) CursorRes {
	return CursorRes{
		tag:  _CursorTag_idc,
		data: uintptr(idc),
	}
}

// If the value is a [co.IDC], returns it and true.
func (me *CursorRes) Idc() (co.IDC, bool) {
	if me.tag == _CursorTag_idc {
		return co.IDC(me.data), true
	}
	return co.IDC(0), false
}

// Creates a new [CursorRes] with an ID value.
func CursorResId(id uint16) CursorRes {
	return CursorRes{
		tag:  _CursorTag_id,
		data: uintptr(id),
	}
}

// If the value is an ID, returns it and true.
func (me *CursorRes) Id() (uint16, bool) {
	if me.tag == _CursorTag_id {
		return uint16(me.data), true
	}
	return 0, false
}

// Creates a new [CursorRes] with a string value.
func CursorResStr(str string) CursorRes {
	return CursorRes{
		tag: _CursorTag_str,
		str: str,
	}
}

// If the value is a string, returns it and true.
func (me *CursorRes) Str() (string, bool) {
	return me.str, me.tag == _CursorTag_str
}

// Converts the internal value to uintptr.
func (me *CursorRes) raw(wideBuf *wstr.Buf[wstr.Stack20]) uintptr {
	switch me.tag {
	case _CursorTag_idc, _CursorTag_id:
		return me.data
	case _CursorTag_str:
		wideBuf.Set(me.str, wstr.EMPTY_IS_NIL)
		return uintptr(wideBuf.UnsafePtr())
	default:
		panic("Invalid CursorRes value.")
	}
}

// Tagged union for an [icon resource] identifier, which can be:
//   - [co.IDI]
//   - uint16
//   - string
//
// # Example:
//
//	icoId := win.IconResIdi(co.IDI_HAND)
//
//	if idi, ok := icoId.Idi(); ok {
//		println(idi)
//	}
//
// [icon resource]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
type IconRes struct {
	tag  _IconResTag
	data uintptr
	str  string
}

type _IconResTag uint8

const (
	_IconResTag_idi _IconResTag = 0x1
	_IconResTag_id  _IconResTag = 0x2
	_IconResTag_str _IconResTag = 0x3
)

// Creates a new [IconRes] with a [co.IDI] value.
func IconResIdi(idi co.IDI) IconRes {
	return IconRes{
		tag:  _IconResTag_idi,
		data: uintptr(idi),
	}
}

// If the value is a [co.IDI], returns it and true.
func (me *IconRes) Idi() (co.IDI, bool) {
	if me.tag == _IconResTag_idi {
		return co.IDI(me.data), true
	}
	return co.IDI(0), false
}

// Creates a new [IconRes] with an ID value.
func IconResId(id uint16) IconRes {
	return IconRes{
		tag:  _IconResTag_id,
		data: uintptr(id),
	}
}

// If the value is an ID, returns it and true.
func (me *IconRes) Id() (uint16, bool) {
	if me.tag == _IconResTag_id {
		return uint16(me.data), true
	}
	return 0, false
}

// Creates a new [IconRes] with a string value.
func IconResStr(str string) IconRes {
	return IconRes{
		tag: _IconResTag_str,
		str: str,
	}
}

// If the value is a string, returns it and true.
func (me *IconRes) Str() (string, bool) {
	return me.str, me.tag == _IconResTag_str
}

// Converts the internal value to uintptr.
func (me *IconRes) raw(wideBuf *wstr.Buf[wstr.Stack20]) uintptr {
	switch me.tag {
	case _IconResTag_idi, _IconResTag_id:
		return me.data
	case _IconResTag_str:
		wideBuf.Set(me.str, wstr.EMPTY_IS_NIL)
		return uintptr(wideBuf.UnsafePtr())
	default:
		panic("Invalid IconRes value.")
	}
}

// Tagged union for a resource identifier, which can be:
//   - uint16
//   - string
//
// # Example
//
//	resId := win.ResIdInt(0x400)
//
//	if id, ok := resId.Int(); ok {
//		println(id)
//	}
type ResId struct {
	tag _ResIdTag
	id  uint16
	str string
}

type _ResIdTag uint8

const (
	_ResIdTag_id  _ResIdTag = 0x1
	_ResIdTag_str _ResIdTag = 0x2
)

// Creates a new [ResId] with an integer value.
func ResIdInt(id uint16) ResId {
	return ResId{
		tag: _ResIdTag_id,
		id:  id,
	}
}

// If the value is an integer, returns it and true.
func (me *ResId) Int() (uint16, bool) {
	return me.id, me.tag == _ResIdTag_id
}

// Creates a new [ResId] with a string value.
func ResIdStr(str string) ResId {
	return ResId{
		tag: _ResIdTag_str,
		str: str,
	}
}

// If the value is a string, returns it and true.
func (me *ResId) Str() (string, bool) {
	return me.str, me.tag == _ResIdTag_str
}

// Converts the internal value to uintptr.
func (me *ResId) raw(wideBuf *wstr.Buf[wstr.Stack20]) uintptr {
	switch me.tag {
	case _ResIdTag_id:
		return uintptr(me.id)
	case _ResIdTag_str:
		wideBuf.Set(me.str, wstr.EMPTY_IS_NIL)
		return uintptr(wideBuf.UnsafePtr())
	default:
		panic("Invalid ResId value.")
	}
}
