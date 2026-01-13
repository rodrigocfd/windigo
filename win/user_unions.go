//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/wstr"
)

// Tagged union for a class name identifier, which can be:
//   - none
//   - [ATOM]
//   - string
//
// Example:
//
//	clsName := win.ClassNameStr("FOO")
//
//	if s, ok := clsName.Str(); ok {
//		println(s)
//	}
type ClassName struct {
	tag  _TagClassName
	atom ATOM
	str  string
}

type _TagClassName uint8

const (
	_TagClassName_none _TagClassName = iota
	_TagClassName_atom
	_TagClassName_str
)

// Constructs a new [ClassName] with an empty value.
func ClassNameNone() ClassName {
	return ClassName{
		tag: _TagClassName_none,
	}
}

// Returns true if there is no value.
func (me *ClassName) IsNone() bool {
	return me.tag == _TagClassName_none
}

// Constructs a new [ClassName] with an [ATOM] value.
func ClassNameAtom(atom ATOM) ClassName {
	return ClassName{
		tag:  _TagClassName_atom,
		atom: atom,
	}
}

// If the value is an [ATOM], returns it and true.
func (me *ClassName) Atom() (ATOM, bool) {
	return me.atom, me.tag == _TagClassName_atom
}

// Constructs a new [ClassName] with a string value.
func ClassNameStr(s string) ClassName {
	return ClassName{
		tag: _TagClassName_str,
		str: s,
	}
}

// If the value is a string, returns it and true.
func (me *ClassName) Str() (string, bool) {
	return me.str, me.tag == _TagClassName_str
}

// Converts the internal value to uintptr.
func (me *ClassName) raw(wBuf *wstr.BufEncoder) uintptr {
	switch me.tag {
	case _TagClassName_none:
		return 0
	case _TagClassName_atom:
		return uintptr(me.atom)
	case _TagClassName_str:
		return uintptr(wBuf.EmptyIsNil(me.str))
	default:
		panic("Invalid ClassName value.")
	}
}

// Tagged union for a [cursor resource] identifier, which can be:
//   - [co.IDC]
//   - uint16
//   - string
//
// Example:
//
//	curId := win.CursorResIdc(co.IDC_ARROW)
//
//	if idc, ok := curId.Idc(); ok {
//		println(idc)
//	}
//
// [cursor resource]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
type CursorRes struct {
	tag  _TagCursorRes
	data uintptr
	str  string
}

type _TagCursorRes uint8

const (
	_TagCursorRes_idc _TagCursorRes = 1 + iota
	_TagCursorRes_id
	_TagCursorRes_str
)

// Constructs a new [CursorRes] with a [co.IDC] value.
func CursorResIdc(idc co.IDC) CursorRes {
	return CursorRes{
		tag:  _TagCursorRes_idc,
		data: uintptr(idc),
	}
}

// If the value is a [co.IDC], returns it and true.
func (me *CursorRes) Idc() (co.IDC, bool) {
	if me.tag == _TagCursorRes_idc {
		return co.IDC(me.data), true
	}
	return co.IDC(0), false
}

// Constructs a new [CursorRes] with an ID value.
func CursorResId(id uint16) CursorRes {
	return CursorRes{
		tag:  _TagCursorRes_id,
		data: uintptr(id),
	}
}

// If the value is an ID, returns it and true.
func (me *CursorRes) Id() (uint16, bool) {
	if me.tag == _TagCursorRes_id {
		return uint16(me.data), true
	}
	return 0, false
}

// Constructs a new [CursorRes] with a string value.
func CursorResStr(s string) CursorRes {
	return CursorRes{
		tag: _TagCursorRes_str,
		str: s,
	}
}

// If the value is a string, returns it and true.
func (me *CursorRes) Str() (string, bool) {
	return me.str, me.tag == _TagCursorRes_str
}

// Converts the internal value to uintptr.
func (me *CursorRes) raw(wBuf *wstr.BufEncoder) uintptr {
	switch me.tag {
	case _TagCursorRes_idc, _TagCursorRes_id:
		return me.data
	case _TagCursorRes_str:
		return uintptr(wBuf.EmptyIsNil(me.str))
	default:
		panic("Invalid CursorRes value.")
	}
}

// Tagged union for an [icon resource] identifier, which can be:
//   - [co.IDI]
//   - uint16
//   - string
//
// Example:
//
//	icoId := win.IconResIdi(co.IDI_HAND)
//
//	if idi, ok := icoId.Idi(); ok {
//		println(idi)
//	}
//
// [icon resource]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
type IconRes struct {
	tag  _TagIconRes
	data uintptr
	str  string
}

type _TagIconRes uint8

const (
	_TagIconRes_idi _TagIconRes = 1 + iota
	_TagIconRes_id
	_TagIconRes_str
)

// Constructs a new [IconRes] with a [co.IDI] value.
func IconResIdi(idi co.IDI) IconRes {
	return IconRes{
		tag:  _TagIconRes_idi,
		data: uintptr(idi),
	}
}

// If the value is a [co.IDI], returns it and true.
func (me *IconRes) Idi() (co.IDI, bool) {
	if me.tag == _TagIconRes_idi {
		return co.IDI(me.data), true
	}
	return co.IDI(0), false
}

// Constructs a new [IconRes] with an ID value.
func IconResId(id uint16) IconRes {
	return IconRes{
		tag:  _TagIconRes_id,
		data: uintptr(id),
	}
}

// If the value is an ID, returns it and true.
func (me *IconRes) Id() (uint16, bool) {
	if me.tag == _TagIconRes_id {
		return uint16(me.data), true
	}
	return 0, false
}

// Constructs a new [IconRes] with a string value.
func IconResStr(s string) IconRes {
	return IconRes{
		tag: _TagIconRes_str,
		str: s,
	}
}

// If the value is a string, returns it and true.
func (me *IconRes) Str() (string, bool) {
	return me.str, me.tag == _TagIconRes_str
}

// Converts the internal value to uintptr.
func (me *IconRes) raw(wBuf *wstr.BufEncoder) uintptr {
	switch me.tag {
	case _TagIconRes_idi, _TagIconRes_id:
		return me.data
	case _TagIconRes_str:
		return uintptr(wBuf.EmptyIsNil(me.str))
	default:
		panic("Invalid IconRes value.")
	}
}

// Tagged union for a resource identifier, which can be:
//   - uint16
//   - string
//
// Example:
//
//	resId := win.ResIdInt(0x400)
//
//	if id, ok := resId.Int(); ok {
//		println(id)
//	}
type ResId struct {
	tag _TagResId
	id  uint16
	str string
}

type _TagResId uint8

const (
	_TagResId_id _TagResId = 1 + iota
	_TagResId_str
)

// Constructs a new [ResId] with an integer value.
func ResIdInt(id uint16) ResId {
	return ResId{
		tag: _TagResId_id,
		id:  id,
	}
}

// If the value is an integer, returns it and true.
func (me *ResId) Int() (uint16, bool) {
	return me.id, me.tag == _TagResId_id
}

// Constructs a new [ResId] with a string value.
func ResIdStr(s string) ResId {
	return ResId{
		tag: _TagResId_str,
		str: s,
	}
}

// If the value is a string, returns it and true.
func (me *ResId) Str() (string, bool) {
	return me.str, me.tag == _TagResId_str
}

// Converts the internal value to uintptr.
func (me *ResId) raw(wBuf *wstr.BufEncoder) uintptr {
	switch me.tag {
	case _TagResId_id:
		return uintptr(me.id)
	case _TagResId_str:
		return uintptr(wBuf.EmptyIsNil(me.str))
	default:
		panic("Invalid ResId value.")
	}
}
