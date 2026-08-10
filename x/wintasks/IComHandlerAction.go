//go:build windows

package wintasks

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cotasks"
)

// [IComHandlerAction] COM interface.
//
// [IComHandlerAction]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-icomhandleraction
type IComHandlerAction struct{ IAction }

type _IComHandlerActionVt struct {
	_IActionVt
	Get_ClassId uintptr
	Put_ClassId uintptr
	Get_Data    uintptr
	Put_Data    uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IComHandlerAction) IID() *co.IID {
	return &cotasks.IID_IComHandlerAction
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IComHandlerAction) AddRef(releaser *win.OleReleaser) *IComHandlerAction {
	return utl.OleNewFromAddRef[*IComHandlerAction](me, releaser)
}

// [get_ClassId] method.
//
// [get_ClassId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-icomhandleraction-get_classid
func (me *IComHandlerAction) GetClassId() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IComHandlerActionVt](me.Ppvt()).Get_ClassId)
}

// [get_Data] method.
//
// [get_Data]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-icomhandleraction-get_data
func (me *IComHandlerAction) GetData() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IComHandlerActionVt](me.Ppvt()).Get_Data)
}

// [put_ClassId] method.
//
// [put_ClassId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-icomhandleraction-put_classid
func (me *IComHandlerAction) PutClassId(clsId string) error {
	return oleCallSetBstr(me, clsId, utl.Vt[_IComHandlerActionVt](me.Ppvt()).Put_ClassId)
}

// [put_Data] method.
//
// [put_Data]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-icomhandleraction-put_data
func (me *IComHandlerAction) PutData(data string) error {
	return oleCallSetBstr(me, data, utl.Vt[_IComHandlerActionVt](me.Ppvt()).Put_Data)
}
