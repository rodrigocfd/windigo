//go:build windows

package win

import (
	"fmt"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Tagged union for [DWMWINDOWATTRIBUTE] values.
//
// [DWMWINDOWATTRIBUTE]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmwindowattribute
type DwmAttr struct {
	tag co.DWMWA
	dw  uint32 // stores BOOL, COLORREF and const values
	rc  RECT
}

// Creates a [DwmAttr] with a [co.DWMWA_NCRENDERING_ENABLED] value.
func DwmAttrNcRenderingEnabled(enabled bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_NCRENDERING_ENABLED,
		dw:  utl.BoolToUint32(enabled),
	}
}

// If the value is [co.DWMWA_NCRENDERING_ENABLED], returns it and true.
func (me *DwmAttr) NcRenderingEnabled() (actualValue, ok bool) {
	if me.tag == co.DWMWA_NCRENDERING_ENABLED {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_NCRENDERING_POLICY] value.
func DwmAttrNcRenderingPolicy(policy co.DWMNCRP) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_NCRENDERING_POLICY,
		dw:  uint32(policy),
	}
}

// If the value is [co.DWMWA_NCRENDERING_POLICY], returns it and true.
func (me *DwmAttr) NcRenderingPolicy() (co.DWMNCRP, bool) {
	if me.tag == co.DWMWA_NCRENDERING_POLICY {
		return co.DWMNCRP(me.dw), true
	}
	return co.DWMNCRP(0), false
}

// Creates a [DwmAttr] with a [co.DWMWA_TRANSITIONS_FORCEDISABLED] value.
func DwmAttrTransitionsForceDisabled(force bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_TRANSITIONS_FORCEDISABLED,
		dw:  utl.BoolToUint32(force),
	}
}

// If the value is [co.DWMWA_TRANSITIONS_FORCEDISABLED], returns it and true.
func (me *DwmAttr) TransitionsForceDisabled() (actualValue, ok bool) {
	if me.tag == co.DWMWA_TRANSITIONS_FORCEDISABLED {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_ALLOW_NCPAINT] value.
func DwmAttrAllowNcPaint(allow bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_ALLOW_NCPAINT,
		dw:  utl.BoolToUint32(allow),
	}
}

// If the value is [co.DWMWA_ALLOW_NCPAINT], returns it and true.
func (me *DwmAttr) AllowNcPaint() (actualValue, ok bool) {
	if me.tag == co.DWMWA_ALLOW_NCPAINT {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_CAPTION_BUTTON_BOUNDS] value.
func DwmAttrCaptionButtonBounds(rc RECT) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_CAPTION_BUTTON_BOUNDS,
		rc:  rc,
	}
}

// If the value is [co.DWMWA_CAPTION_BUTTON_BOUNDS], returns it and true.
func (me *DwmAttr) CaptionButtonBounds() (RECT, bool) {
	if me.tag == co.DWMWA_CAPTION_BUTTON_BOUNDS {
		return me.rc, true
	}
	return RECT{}, false
}

// Creates a [DwmAttr] with a [co.DWMWA_NONCLIENT_RTL_LAYOUT] value.
func DwmAttrNonClientRtlLayout(ncRtl bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_NONCLIENT_RTL_LAYOUT,
		dw:  utl.BoolToUint32(ncRtl),
	}
}

// If the value is [co.DWMWA_NONCLIENT_RTL_LAYOUT], returns it and true.
func (me *DwmAttr) NonClientRtlLayout() (actualValue, ok bool) {
	if me.tag == co.DWMWA_NONCLIENT_RTL_LAYOUT {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_FORCE_ICONIC_REPRESENTATION] value.
func DwmAttrForceIconicRepresentation(force bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_FORCE_ICONIC_REPRESENTATION,
		dw:  utl.BoolToUint32(force),
	}
}

// If the value is [co.DWMWA_FORCE_ICONIC_REPRESENTATION], returns it and true.
func (me *DwmAttr) ForceIconicRepresentation() (actualValue, ok bool) {
	if me.tag == co.DWMWA_FORCE_ICONIC_REPRESENTATION {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_FLIP3D_POLICY] value.
func DwmAttrFlip3dPolicy(policy co.DWMFLIP3D) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_FLIP3D_POLICY,
		dw:  uint32(policy),
	}
}

// If the value is [co.DWMWA_FLIP3D_POLICY], returns it and true.
func (me *DwmAttr) Flip3dPolicy() (co.DWMFLIP3D, bool) {
	if me.tag == co.DWMWA_FLIP3D_POLICY {
		return co.DWMFLIP3D(me.dw), true
	}
	return co.DWMFLIP3D(0), false
}

// Creates a [DwmAttr] with a [co.DWMWA_EXTENDED_FRAME_BOUNDS] value.
func DwmAttrExtendedFrameBounds(rc RECT) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_EXTENDED_FRAME_BOUNDS,
		rc:  rc,
	}
}

// If the value is [co.DWMWA_EXTENDED_FRAME_BOUNDS], returns it and true.
func (me *DwmAttr) ExtendedFrameBounds() (RECT, bool) {
	if me.tag == co.DWMWA_EXTENDED_FRAME_BOUNDS {
		return me.rc, true
	}
	return RECT{}, false
}

// Creates a [DwmAttr] with a [co.DWMWA_HAS_ICONIC_BITMAP] value.
func DwmAttrHasIconicBitmap(has bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_HAS_ICONIC_BITMAP,
		dw:  utl.BoolToUint32(has),
	}
}

// If the value is [co.DWMWA_HAS_ICONIC_BITMAP], returns it and true.
func (me *DwmAttr) HasIconicBitmap() (actualValue, ok bool) {
	if me.tag == co.DWMWA_HAS_ICONIC_BITMAP {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_DISALLOW_PEEK] value.
func DwmAttrDisallowPeek(disallow bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_DISALLOW_PEEK,
		dw:  utl.BoolToUint32(disallow),
	}
}

// If the value is [co.DWMWA_DISALLOW_PEEK], returns it and true.
func (me *DwmAttr) DisallowPeek() (actualValue, ok bool) {
	if me.tag == co.DWMWA_DISALLOW_PEEK {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_EXCLUDED_FROM_PEEK] value.
func DwmAttrExcludedFromPeek(excluded bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_EXCLUDED_FROM_PEEK,
		dw:  utl.BoolToUint32(excluded),
	}
}

// If the value is [co.DWMWA_EXCLUDED_FROM_PEEK], returns it and true.
func (me *DwmAttr) ExcludedFromPeek() (actualValue, ok bool) {
	if me.tag == co.DWMWA_EXCLUDED_FROM_PEEK {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_CLOAK] value.
func DwmAttrCloak(cloak bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_CLOAK,
		dw:  utl.BoolToUint32(cloak),
	}
}

// If the value is [co.DWMWA_CLOAK], returns it and true.
func (me *DwmAttr) Cloak() (actualValue, ok bool) {
	if me.tag == co.DWMWA_CLOAK {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_CLOAKED] value.
func DwmAttrCloaked(cloaked co.DWM_CLOAKED) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_CLOAKED,
		dw:  uint32(cloaked),
	}
}

// If the value is [co.DWMWA_CLOAKED], returns it and true.
func (me *DwmAttr) Cloaked() (co.DWM_CLOAKED, bool) {
	if me.tag == co.DWMWA_CLOAKED {
		return co.DWM_CLOAKED(me.dw), true
	}
	return co.DWM_CLOAKED(0), false
}

// Creates a [DwmAttr] with a [co.DWMWA_FREEZE_REPRESENTATION] value.
func DwmAttrFreezeRepresentation(freeze bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_FREEZE_REPRESENTATION,
		dw:  utl.BoolToUint32(freeze),
	}
}

// If the value is [co.DWMWA_FREEZE_REPRESENTATION], returns it and true.
func (me *DwmAttr) FreezeRepresentation() (actualValue, ok bool) {
	if me.tag == co.DWMWA_FREEZE_REPRESENTATION {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_PASSIVE_UPDATE_MODE] value.
func DwmAttrPassiveUpdateMode(passive bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_PASSIVE_UPDATE_MODE,
		dw:  utl.BoolToUint32(passive),
	}
}

// If the value is [co.DWMWA_PASSIVE_UPDATE_MODE], returns it and true.
func (me *DwmAttr) PassiveUpdateMode() (actualValue, ok bool) {
	if me.tag == co.DWMWA_PASSIVE_UPDATE_MODE {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_USE_HOSTBACKDROPBRUSH] value.
func DwmAttrUseHostBackdropBrush(use bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_USE_HOSTBACKDROPBRUSH,
		dw:  utl.BoolToUint32(use),
	}
}

// If the value is [co.DWMWA_USE_HOSTBACKDROPBRUSH], returns it and true.
func (me *DwmAttr) UseHostBackdropBrush() (actualValue, ok bool) {
	if me.tag == co.DWMWA_USE_HOSTBACKDROPBRUSH {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_USE_IMMERSIVE_DARK_MODE] value.
func DwmAttrUseImmersiveDarkMode(use bool) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_USE_IMMERSIVE_DARK_MODE,
		dw:  utl.BoolToUint32(use),
	}
}

// If the value is [co.DWMWA_USE_IMMERSIVE_DARK_MODE], returns it and true.
func (me *DwmAttr) UseImmersiveDarkMode() (actualValue, ok bool) {
	if me.tag == co.DWMWA_USE_IMMERSIVE_DARK_MODE {
		return me.dw != 0, true
	}
	return false, false
}

// Creates a [DwmAttr] with a [co.DWMWA_WINDOW_CORNER_PREFERENCE] value.
func DwmAttrWindowCornerPreference(corner co.DWMWCP) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_WINDOW_CORNER_PREFERENCE,
		dw:  uint32(corner),
	}
}

// If the value is [co.DWMWA_WINDOW_CORNER_PREFERENCE], returns it and true.
func (me *DwmAttr) WindowCornerPreference() (co.DWMWCP, bool) {
	if me.tag == co.DWMWA_WINDOW_CORNER_PREFERENCE {
		return co.DWMWCP(me.dw), true
	}
	return co.DWMWCP(0), false
}

// Creates a [DwmAttr] with a [co.DWMWA_BORDER_COLOR] value.
//
// You may want to specify:
//   - [COLORREF_DWMA_NONE]
//   - [COLORREF_DWMA_DEFAULT]
func DwmAttrBorderColor(color COLORREF) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_BORDER_COLOR,
		dw:  uint32(color),
	}
}

// If the value is [co.DWMWA_BORDER_COLOR], returns it and true.
//
// You may also get:
//   - [COLORREF_DWMA_NONE]
//   - [COLORREF_DWMA_DEFAULT]
func (me *DwmAttr) BorderColor() (COLORREF, bool) {
	if me.tag == co.DWMWA_BORDER_COLOR {
		return COLORREF(me.dw), true
	}
	return COLORREF(0), false
}

// Creates a [DwmAttr] with a [co.DWMWA_CAPTION_COLOR] value.
//
// You may want to specify:
//   - [COLORREF_DWMA_DEFAULT]
func DwmAttrCaptionColor(color COLORREF) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_CAPTION_COLOR,
		dw:  uint32(color),
	}
}

// If the value is [co.DWMWA_CAPTION_COLOR], returns it and true.
//
// You may also get:
//   - [COLORREF_DWMA_DEFAULT]
func (me *DwmAttr) CaptionColor() (COLORREF, bool) {
	if me.tag == co.DWMWA_CAPTION_COLOR {
		return COLORREF(me.dw), true
	}
	return COLORREF(0), false
}

// Creates a [DwmAttr] with a [co.DWMWA_TEXT_COLOR] value.
//
// You may want to specify:
//   - [COLORREF_DWMA_DEFAULT]
func DwmAttrTextColor(color COLORREF) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_TEXT_COLOR,
		dw:  uint32(color),
	}
}

// If the value is [co.DWMWA_TEXT_COLOR], returns it and true.
//
// You may also get:
//   - [COLORREF_DWMA_DEFAULT]
func (me *DwmAttr) TextColor() (COLORREF, bool) {
	if me.tag == co.DWMWA_TEXT_COLOR {
		return COLORREF(me.dw), true
	}
	return COLORREF(0), false
}

// Creates a [DwmAttr] with a [co.DWMWA_VISIBLE_FRAME_BORDER_THICKNESS] value.
func DwmAttrVisibleFrameBorderThickness(width uint) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_VISIBLE_FRAME_BORDER_THICKNESS,
		dw:  uint32(width),
	}
}

// If the value is [co.DWMWA_VISIBLE_FRAME_BORDER_THICKNESS], returns it and true.
func (me *DwmAttr) VisibleFrameBorderThickness() (uint, bool) {
	if me.tag == co.DWMWA_VISIBLE_FRAME_BORDER_THICKNESS {
		return uint(me.dw), true
	}
	return uint(0), false
}

// Creates a [DwmAttr] with a [co.DWMWA_SYSTEMBACKDROP_TYPE] value.
func DwmAttrSystemBackdropType(sb co.DWMSBT) DwmAttr {
	return DwmAttr{
		tag: co.DWMWA_SYSTEMBACKDROP_TYPE,
		dw:  uint32(sb),
	}
}

// If the value is [co.DWMWA_SYSTEMBACKDROP_TYPE], returns it and true.
func (me *DwmAttr) SystemBackdropType() (co.DWMSBT, bool) {
	if me.tag == co.DWMWA_SYSTEMBACKDROP_TYPE {
		return co.DWMSBT(me.dw), true
	}
	return co.DWMSBT(0), false
}

func dwmAttrFromRaw(attr co.DWMWA, dwBuf uint32, rcBuf RECT) DwmAttr {
	switch attr {
	case co.DWMWA_NCRENDERING_ENABLED:
		return DwmAttrNcRenderingEnabled(dwBuf != 0)
	case co.DWMWA_NCRENDERING_POLICY:
		return DwmAttrNcRenderingPolicy(co.DWMNCRP(dwBuf))
	case co.DWMWA_TRANSITIONS_FORCEDISABLED:
		return DwmAttrTransitionsForceDisabled(dwBuf != 0)
	case co.DWMWA_ALLOW_NCPAINT:
		return DwmAttrAllowNcPaint(dwBuf != 0)
	case co.DWMWA_CAPTION_BUTTON_BOUNDS:
		return DwmAttrCaptionButtonBounds(rcBuf)
	case co.DWMWA_NONCLIENT_RTL_LAYOUT:
		return DwmAttrNonClientRtlLayout(dwBuf != 0)
	case co.DWMWA_FORCE_ICONIC_REPRESENTATION:
		return DwmAttrForceIconicRepresentation(dwBuf != 0)
	case co.DWMWA_FLIP3D_POLICY:
		return DwmAttrFlip3dPolicy(co.DWMFLIP3D(dwBuf))
	case co.DWMWA_EXTENDED_FRAME_BOUNDS:
		return DwmAttrExtendedFrameBounds(rcBuf)
	case co.DWMWA_HAS_ICONIC_BITMAP:
		return DwmAttrHasIconicBitmap(dwBuf != 0)
	case co.DWMWA_DISALLOW_PEEK:
		return DwmAttrDisallowPeek(dwBuf != 0)
	case co.DWMWA_EXCLUDED_FROM_PEEK:
		return DwmAttrExcludedFromPeek(dwBuf != 0)
	case co.DWMWA_CLOAK:
		return DwmAttrCloak(dwBuf != 0)
	case co.DWMWA_CLOAKED:
		return DwmAttrCloaked(co.DWM_CLOAKED(dwBuf))
	case co.DWMWA_FREEZE_REPRESENTATION:
		return DwmAttrFreezeRepresentation(dwBuf != 0)
	case co.DWMWA_PASSIVE_UPDATE_MODE:
		return DwmAttrPassiveUpdateMode(dwBuf != 0)
	case co.DWMWA_USE_HOSTBACKDROPBRUSH:
		return DwmAttrUseHostBackdropBrush(dwBuf != 0)
	case co.DWMWA_USE_IMMERSIVE_DARK_MODE:
		return DwmAttrUseImmersiveDarkMode(dwBuf != 0)
	case co.DWMWA_WINDOW_CORNER_PREFERENCE:
		return DwmAttrWindowCornerPreference(co.DWMWCP(dwBuf))
	case co.DWMWA_BORDER_COLOR:
		return DwmAttrBorderColor(COLORREF(dwBuf))
	case co.DWMWA_CAPTION_COLOR:
		return DwmAttrCaptionColor(COLORREF(dwBuf))
	case co.DWMWA_TEXT_COLOR:
		return DwmAttrTextColor(COLORREF(dwBuf))
	case co.DWMWA_VISIBLE_FRAME_BORDER_THICKNESS:
		return DwmAttrVisibleFrameBorderThickness(uint(dwBuf))
	case co.DWMWA_SYSTEMBACKDROP_TYPE:
		return DwmAttrSystemBackdropType(co.DWMSBT(dwBuf))
	default:
		panic(fmt.Sprintf("Invalid co.DWMWA value: %d.", attr))
	}
}
