//go:build windows

package wintasks

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cotasks"
	"github.com/rodrigocfd/windigo/x/winaut"
)

// [ITaskDefinition] COM interface.
//
// [ITaskDefinition]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itaskdefinition
type ITaskDefinition struct{ winaut.IDispatch }

type _ITaskDefinitionVt struct {
	utl.IDispatchVt
	Get_RegistrationInfo uintptr
	Put_RegistrationInfo uintptr
	Get_Triggers         uintptr
	Put_Triggers         uintptr
	Get_Settings         uintptr
	Put_Settings         uintptr
	Get_Data             uintptr
	Put_Data             uintptr
	Get_Principal        uintptr
	Put_Principal        uintptr
	Get_Actions          uintptr
	Put_Actions          uintptr
	Get_XmlText          uintptr
	Put_XmlText          uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskDefinition) IID() *co.IID {
	return &cotasks.IID_ITaskDefinition
}

// [get_Actions] method.
//
// [get_Actions]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_actions
func (me *ITaskDefinition) GetActions(releaser *win.OleReleaser) (*IActionCollection, error) {
	return utl.OleNewFromCallWithoutParms[*IActionCollection](me, releaser,
		utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Get_Actions)
}

// [get_Data] method.
//
// [get_Data]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_data
func (me *ITaskDefinition) GetData() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Get_Data)
}

// [get_Principal] method.
//
// [get_Principal]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_principal
func (me *ITaskDefinition) GetPrincipal(releaser *win.OleReleaser) (*IPrincipal, error) {
	return utl.OleNewFromCallWithoutParms[*IPrincipal](me, releaser,
		utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Get_Principal)
}

// [get_RegistrationInfo] method.
//
// [get_RegistrationInfo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_registrationinfo
func (me *ITaskDefinition) GetRegistrationInfo(releaser *win.OleReleaser) (*IRegistrationInfo, error) {
	return utl.OleNewFromCallWithoutParms[*IRegistrationInfo](me, releaser,
		utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Get_RegistrationInfo)
}

// [get_Settings] method.
//
// [get_Settings]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_settings
func (me *ITaskDefinition) GetSettings(releaser *win.OleReleaser) (*ITaskSettings, error) {
	return utl.OleNewFromCallWithoutParms[*ITaskSettings](me, releaser,
		utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Get_Settings)
}

// [get_Triggers] method.
//
// [get_Triggers]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_triggers
func (me *ITaskDefinition) GetTriggers(releaser *win.OleReleaser) (*ITriggerCollection, error) {
	return utl.OleNewFromCallWithoutParms[*ITriggerCollection](me, releaser,
		utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Get_Triggers)
}

// [get_XmlText] method.
//
// [get_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_xmltext
func (me *ITaskDefinition) GetXmlText() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Get_XmlText)
}

// [put_Data] method.
//
// [put_Data]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_data
func (me *ITaskDefinition) PutData(data string) error {
	return oleCallSetBstr(me, data, utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Put_Data)
}

// [put_Principal] method.
//
// [put_Principal]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_principal
func (me *ITaskDefinition) PutPrincipal(principal *IPrincipal) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Put_Principal,
		me.Ppvt(),
		principal.Ppvt())
	return utl.HresultToError(ret)
}

// [put_RegistrationInfo] method.
//
// [put_RegistrationInfo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_registrationinfo
func (me *ITaskDefinition) PutRegistrationInfo(registrationInfo *IRegistrationInfo) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Put_RegistrationInfo,
		me.Ppvt(),
		registrationInfo.Ppvt())
	return utl.HresultToError(ret)
}

// [put_Settings] method.
//
// [put_Settings]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_settings
func (me *ITaskDefinition) PutSettings(settings *ITaskSettings) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Put_Settings,
		me.Ppvt(),
		settings.Ppvt())
	return utl.HresultToError(ret)
}

// [put_Triggers] method.
//
// [put_Triggers]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_triggers
func (me *ITaskDefinition) PutTriggers(triggers *ITriggerCollection) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Put_Triggers,
		me.Ppvt(),
		triggers.Ppvt())
	return utl.HresultToError(ret)
}

// [put_XmlText] method.
//
// [put_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_xmltext
func (me *ITaskDefinition) PutXmlText(xmlText string) error {
	return oleCallSetBstr(me, xmlText, utl.Vt[_ITaskDefinitionVt](me.Ppvt()).Put_XmlText)
}
