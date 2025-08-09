//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type (
	_StorageMsgLib struct { // ordinary WM messages
		id  co.WM
		fun func(p Wm)
	}
	_StorageNfyLib struct { // WM_NOTIFY
		idFrom uint16
		code   co.NM
		fun    func(p unsafe.Pointer)
	}
)

// Stores messages added internally by the library.
type _EventsWindowLib struct {
	inits []func()         // WM_CREATE and WM_INITDIALOG
	msgs  []_StorageMsgLib // ordinary WM messages
	nfys  []_StorageNfyLib // WM_NOTIFY
}

// Constructor.
func newEventsWindowLib() _EventsWindowLib {
	return _EventsWindowLib{
		inits: make([]func(), 0),
		msgs:  make([]_StorageMsgLib, 0),
		nfys:  make([]_StorageNfyLib, 0),
	}
}

// To be called after the first WM_CREATE/INITDIALOG processing. Releases the
// memory in all these closures, which are never called again.
func (me *_EventsWindowLib) removeWmCreateInitdialog() {
	me.inits = nil
}

// Releases the memory of all closures.
func (me *_EventsWindowLib) clear() {
	me.removeWmCreateInitdialog()
	me.msgs = nil
	me.nfys = nil
}

// Runs all the internal closures for the given message, discarding the results.
func (me *_EventsWindowLib) processAll(p Wm) (atLeastOne bool) {
	switch p.Msg {
	case co.WM_CREATE, co.WM_INITDIALOG:
		for _, fun := range me.inits {
			fun()
			atLeastOne = true
		}
	case co.WM_NOTIFY:
		pHdr := unsafe.Pointer(p.LParam)
		hdr := (*win.NMHDR)(pHdr)
		for _, obj := range me.nfys {
			if obj.idFrom == uint16(hdr.IdFrom) && obj.code == co.NM(hdr.Code) {
				obj.fun(pHdr)
				atLeastOne = true
			}
		}
	default:
		for _, obj := range me.msgs {
			if obj.id == p.Msg {
				obj.fun(p)
				atLeastOne = true
			}
		}
	}
	return
}

func (me *_EventsWindowLib) wmCreateOrInitdialog(fun func()) {
	me.inits = append(me.inits, fun)
}

func (me *_EventsWindowLib) wm(id co.WM, fun func(p Wm)) {
	me.msgs = append(me.msgs, _StorageMsgLib{id, fun})
}

func (me *_EventsWindowLib) wmNotify(idFrom uint16, code co.NM, fun func(p unsafe.Pointer)) {
	me.nfys = append(me.nfys, _StorageNfyLib{idFrom, code, fun})
}
