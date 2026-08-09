//go:build windows

package wintasks

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/cotasks"
	"github.com/rodrigocfd/windigo/x/winaut"
)

// [ITaskNamedValuePair] COM interface.
//
// [ITaskNamedValuePair]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itasknamedvaluepair
type ITaskNamedValuePair struct{ winaut.IDispatch }

type _ITaskNamedValuePairVt struct {
	utl.IDispatchVt
	Get_Name  uintptr
	Put_Name  uintptr
	Get_Value uintptr
	Put_Value uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskNamedValuePair) IID() *co.IID {
	return &cotasks.IID_ITaskNamedValuePair
}

// [get_Name] method.
//
// [get_Name]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluepair-get_name
func (me *ITaskNamedValuePair) GetName() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskNamedValuePairVt](me.Ppvt()).Get_Name)
}

// [get_Value] method.
//
// [get_Value]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluepair-get_value
func (me *ITaskNamedValuePair) GetValue() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskNamedValuePairVt](me.Ppvt()).Get_Value)
}

// [put_Name] method.
//
// [put_Name]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluepair-put_name
func (me *ITaskNamedValuePair) PutName(name string) error {
	return oleCallSetBstr(me, name, utl.Vt[_ITaskNamedValuePairVt](me.Ppvt()).Put_Name)
}

// [put_Value] method.
//
// [put_Value]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluepair-put_value
func (me *ITaskNamedValuePair) PutValue(value string) error {
	return oleCallSetBstr(me, value, utl.Vt[_ITaskNamedValuePairVt](me.Ppvt()).Put_Value)
}
