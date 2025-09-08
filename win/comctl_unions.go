//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/co"
)

// Tagged union for an icon identifier for [TASKDIALOGCONFIG], which can be:
//   - none
//   - [HICON]
//   - uint16
//   - [co.TDICON]
//
// Example:
//
//	ico := TdcIconTdi(co.TD_ICON_ERROR)
//
//	if tdi, ok := ico.Tdi(); ok {
//		println(tdi)
//	}
type TdcIcon struct {
	tag  _TdcIconTag
	data uint64
}

type _TdcIconTag uint8

const (
	_TdcIconTag_none _TdcIconTag = iota
	_TdcIconTag_hIcon
	_TdcIconTag_id
	_TdcIconTag_tdIcon
)

// Creates a new [TdcIcon] with an empty value.
func TdcIconNone() TdcIcon {
	return TdcIcon{
		tag: _TdcIconTag_none,
	}
}

// Returns true if there is no value.
func (me *TdcIcon) IsNone() bool {
	return me.tag == _TdcIconTag_none
}

// Creates a new [TdcIcon] with a [HICON] value.
func TdcIconHicon(hIcon HICON) TdcIcon {
	return TdcIcon{
		tag:  _TdcIconTag_hIcon,
		data: uint64(hIcon),
	}
}

// If the value is a [HICON], returns it and true.
func (me *TdcIcon) HIcon() (HICON, bool) {
	if me.tag == _TdcIconTag_hIcon {
		return HICON(me.data), true
	}
	return HICON(0), false
}

// Creates a new [TdcIcon] with an ID value.
func TdcIconId(id uint16) TdcIcon {
	return TdcIcon{
		tag:  _TdcIconTag_id,
		data: uint64(uint16(id)),
	}
}

// If the value is an ID, returns it and true.
func (me *TdcIcon) Id() (uint16, bool) {
	if me.tag == _TdcIconTag_id {
		return uint16(me.data), true
	}
	return 0, false
}

// Creates a new [TdcIcon] with a [co.TDICON] value.
func TdcIconTdi(tdIcon co.TDICON) TdcIcon {
	return TdcIcon{
		tag:  _TdcIconTag_tdIcon,
		data: uint64(tdIcon),
	}
}

// If the value is a [co.TDICON], returns it and true.
func (me *TdcIcon) Tdi() (co.TDICON, bool) {
	if me.tag == _TdcIconTag_tdIcon {
		return co.TDICON(me.data), true
	}
	return co.TDICON(0), false
}

// Returns the internal value as uint64.
func (me *TdcIcon) raw() uint64 {
	return me.data
}
