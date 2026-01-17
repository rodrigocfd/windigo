//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [IAction] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IAction]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iaction
type IAction struct{ IDispatch }

type _IActionVt struct {
	_IDispatchVt
	Get_Id   uintptr
	Put_Id   uintptr
	Get_Type uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IAction) IID() co.IID {
	return co.IID_IAction
}

// [get_Id] method.
//
// [get_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iaction-get_id
func (me *IAction) GetId() (string, error) {
	return com_callBstrGet(me,
		(*_IActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Id)
}

// [get_Type] method.
//
// [get_Type]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iaction-get_type
func (me *IAction) GetType() (co.TASK_ACTION, error) {
	var tat uint32
	ret, _, _ := syscall.SyscallN(
		(*_IActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Type,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&tat)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.TASK_ACTION(tat), nil
	} else {
		return co.TASK_ACTION(0), hr
	}
}

// [put_Id] method.
//
// [put_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iaction-put_id
func (me *IAction) PutId(id string) error {
	return com_callBstrSet(me, id,
		(*_IActionVt)(unsafe.Pointer(*me.Ppvt())).Put_Id)
}

// [IActionCollection] method.
//
// Implements [OleObj] and [OleResource].
//
// [IActionCollection]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iactioncollection
type IActionCollection struct{ IDispatch }

type _IActionCollectionVt struct {
	_IDispatchVt
	Get_Count    uintptr
	Get_Item     uintptr
	Get__NewEnum uintptr
	Get_XmlText  uintptr
	Put_XmlText  uintptr
	Create       uintptr
	Remove       uintptr
	Clear        uintptr
	Get_Context  uintptr
	Put_Context  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IActionCollection) IID() co.IID {
	return co.IID_IAction
}

// [Clear] method.
//
// [Clear]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-clear
func (me *IActionCollection) Clear() error {
	return com_callErr(me,
		(*_IActionCollectionVt)(unsafe.Pointer(*me.Ppvt())).Clear)
}

// [Create] method.
//
// [Create]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-create
func (me *IActionCollection) Create(
	releaser *OleReleaser,
	actionType co.TASK_ACTION,
) (*IAction, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IActionCollectionVt)(unsafe.Pointer(*me.Ppvt())).Create,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(actionType),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*IAction](ret, ppvtQueried, releaser)
}

// Returns all [IAction] objects by calling [IActionCollection.GetCount] and
// [IActionCollection.GetItem].
func (me *IActionCollection) Enum(releaser *OleReleaser) ([]*IAction, error) {
	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	actions := make([]*IAction, count)
	for i := 0; i < count; i++ {
		action, err := me.GetItem(releaser, i)
		if err != nil {
			return nil, err // stop immediately
		}
		actions = append(actions, action)
	}
	return actions, nil
}

// [get_Context] method.
//
// [get_Context]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-get_context
func (me *IActionCollection) GetContext() (string, error) {
	return com_callBstrGet(me,
		(*_IActionCollectionVt)(unsafe.Pointer(*me.Ppvt())).Get_Context)
}

// [get_Count] method.
//
// [get_Count]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-get_count
func (me *IActionCollection) GetCount() (int, error) {
	var count int32
	ret, _, _ := syscall.SyscallN(
		(*_IActionCollectionVt)(unsafe.Pointer(*me.Ppvt())).Get_Count,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&count)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(count), nil
	} else {
		return 0, hr
	}
}

// [get_Item] method.
//
// [get_Item]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-get_item
func (me *IActionCollection) GetItem(releaser *OleReleaser, index int) (*IAction, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IActionCollectionVt)(unsafe.Pointer(*me.Ppvt())).Get_Item,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(int32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*IAction](ret, ppvtQueried, releaser)
}

// [get_XmlText] method.
//
// [get_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-get_xmltext
func (me *IActionCollection) GetXmlText() (string, error) {
	return com_callBstrGet(me,
		(*_IActionCollectionVt)(unsafe.Pointer(*me.Ppvt())).Get_XmlText)
}

// [put_Context] method.
//
// [put_Context]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-put_context
func (me *IActionCollection) PutContext(context string) error {
	return com_callBstrSet(me, context,
		(*_IActionCollectionVt)(unsafe.Pointer(*me.Ppvt())).Put_Context)
}

// [put_XmlText] method.
//
// [put_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-put_xmltext
func (me *IActionCollection) PutXmlText(text string) error {
	return com_callBstrSet(me, text,
		(*_IActionCollectionVt)(unsafe.Pointer(*me.Ppvt())).Put_XmlText)
}

// [IBootTrigger] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IBootTrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iboottrigger
type IBootTrigger struct{ ITrigger }

type _IBootTriggerVt struct {
	_ITriggerVt
	Get_Delay uintptr
	Put_Delay uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IBootTrigger) IID() co.IID {
	return co.IID_IBootTrigger
}

// [get_Delay] method.
//
// [get_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iboottrigger-get_delay
func (me *IBootTrigger) GetDelay() (string, error) {
	return com_callBstrGet(me,
		(*_IBootTriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_Delay)
}

// [put_Delay] method.
//
// [put_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iboottrigger-put_delay
func (me *IBootTrigger) PutDelay(delay string) error {
	return com_callBstrSet(me, delay,
		(*_IBootTriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_Delay)
}

// [IComHandlerAction] COM interface.
//
// Implements [OleObj] and [OleResource].
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
func (*IComHandlerAction) IID() co.IID {
	return co.IID_IComHandlerAction
}

// [get_ClassId] method.
//
// [get_ClassId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-icomhandleraction-get_classid
func (me *IComHandlerAction) GetClassId() (string, error) {
	return com_callBstrGet(me,
		(*_IComHandlerActionVt)(unsafe.Pointer(*me.Ppvt())).Get_ClassId)
}

// [get_Data] method.
//
// [get_Data]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-icomhandleraction-get_data
func (me *IComHandlerAction) GetData() (string, error) {
	return com_callBstrGet(me,
		(*_IComHandlerActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Data)
}

// [put_ClassId] method.
//
// [put_ClassId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-icomhandleraction-put_classid
func (me *IComHandlerAction) PutClassId(clsId string) error {
	return com_callBstrSet(me, clsId,
		(*_IComHandlerActionVt)(unsafe.Pointer(*me.Ppvt())).Put_ClassId)
}

// [put_Data] method.
//
// [put_Data]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-icomhandleraction-put_data
func (me *IComHandlerAction) PutData(data string) error {
	return com_callBstrSet(me, data,
		(*_IComHandlerActionVt)(unsafe.Pointer(*me.Ppvt())).Put_Data)
}

// [IDailyTrigger] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IDailyTrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-idailytrigger
type IDailyTrigger struct{ ITrigger }

type _IDailyTriggerVt struct {
	_ITriggerVt
	Get_DaysInterval uintptr
	Put_DaysInterval uintptr
	Get_RandomDelay  uintptr
	Put_RandomDelay  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDailyTrigger) IID() co.IID {
	return co.IID_IDailyTrigger
}

// [get_DaysInterval] method.
//
// [get_DaysInterval]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-idailytrigger-get_daysinterval
func (me *IDailyTrigger) GetDaysInterval() (int, error) {
	var days int16
	ret, _, _ := syscall.SyscallN(
		(*_IDailyTriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_DaysInterval,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&days)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(days), nil
	} else {
		return 0, hr
	}
}

// [get_RandomDelay] method.
//
// [get_RandomDelay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-idailytrigger-get_randomdelay
func (me *IDailyTrigger) GetRandomDelay() (string, error) {
	return com_callBstrGet(me,
		(*_IDailyTriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_RandomDelay)
}

// [put_DaysInterval] method.
//
// [put_DaysInterval]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-idailytrigger-put_daysinterval
func (me *IDailyTrigger) PutDaysInterval(days int) error {
	ret, _, _ := syscall.SyscallN(
		(*_IDailyTriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_DaysInterval,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(int16(days)))
	return utl.HresultToError(ret)
}

// [put_RandomDelay] method.
//
// [put_RandomDelay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-idailytrigger-put_randomdelay
func (me *IDailyTrigger) PutRandomDelay(randomDelay string) error {
	return com_callBstrSet(me, randomDelay,
		(*_IDailyTriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_RandomDelay)
}

// [IEmailAction] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IEmailAction]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iemailaction
type IEmailAction struct{ IAction }

type _IEmailActionVt struct {
	_IActionVt
	Get_Server       uintptr
	Put_Server       uintptr
	Get_Subject      uintptr
	Put_Subject      uintptr
	Get_To           uintptr
	Put_To           uintptr
	Get_Cc           uintptr
	Put_Cc           uintptr
	Get_Bcc          uintptr
	Put_Bcc          uintptr
	Get_ReplyTo      uintptr
	Put_ReplyTo      uintptr
	Get_From         uintptr
	Put_From         uintptr
	Get_HeaderFields uintptr
	Put_HeaderFields uintptr
	Get_Body         uintptr
	Put_Body         uintptr
	Get_Attachments  uintptr
	Put_Attachments  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEmailAction) IID() co.IID {
	return co.IID_IEmailAction
}

// [get_Bcc] method.
//
// [get_Bcc]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_bcc
func (me *IEmailAction) GetBcc() (string, error) {
	return com_callBstrGet(me,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Bcc)
}

// [get_Body] method.
//
// [get_Body]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_body
func (me *IEmailAction) GetBody() (string, error) {
	return com_callBstrGet(me,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Body)
}

// [get_Cc] method.
//
// [get_Cc]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_cc
func (me *IEmailAction) GetCc() (string, error) {
	return com_callBstrGet(me,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Cc)
}

// [get_From] method.
//
// [get_From]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_from
func (me *IEmailAction) GetFrom() (string, error) {
	return com_callBstrGet(me,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Get_From)
}

// [get_ReplyTo] method.
//
// [get_ReplyTo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_replyto
func (me *IEmailAction) GetReplyTo() (string, error) {
	return com_callBstrGet(me,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Get_ReplyTo)
}

// [get_Server] method.
//
// [get_Server]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_server
func (me *IEmailAction) GetServer() (string, error) {
	return com_callBstrGet(me,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Server)
}

// [get_Subject] method.
//
// [get_Subject]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_subject
func (me *IEmailAction) GetSubject() (string, error) {
	return com_callBstrGet(me,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Subject)
}

// [get_To] method.
//
// [get_To]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_to
func (me *IEmailAction) GetTo() (string, error) {
	return com_callBstrGet(me,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Get_To)
}

// [put_Bcc] method.
//
// [put_Bcc]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_bcc
func (me *IEmailAction) PutBcc(bcc string) error {
	return com_callBstrSet(me, bcc,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Put_Bcc)
}

// [put_Body] method.
//
// [put_Body]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_body
func (me *IEmailAction) PutBody(bcc string) error {
	return com_callBstrSet(me, bcc,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Put_Body)
}

// [put_Cc] method.
//
// [put_Cc]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_cc
func (me *IEmailAction) PutCc(cc string) error {
	return com_callBstrSet(me, cc,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Put_Cc)
}

// [put_From] method.
//
// [put_From]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_from
func (me *IEmailAction) PutFrom(from string) error {
	return com_callBstrSet(me, from,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Put_From)
}

// [put_ReplyTo] method.
//
// [put_ReplyTo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_replyto
func (me *IEmailAction) PutReplyTo(replyTo string) error {
	return com_callBstrSet(me, replyTo,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Put_ReplyTo)
}

// [put_Server] method.
//
// [put_Server]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_server
func (me *IEmailAction) PutServer(server string) error {
	return com_callBstrSet(me, server,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Put_Server)
}

// [put_Subject] method.
//
// [put_Subject]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_subject
func (me *IEmailAction) PutSubject(subject string) error {
	return com_callBstrSet(me, subject,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Put_Subject)
}

// [put_To] method.
//
// [put_To]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_to
func (me *IEmailAction) PutTo(to string) error {
	return com_callBstrSet(me, to,
		(*_IEmailActionVt)(unsafe.Pointer(*me.Ppvt())).Put_To)
}

// [IEventTrigger] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IEventTrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-ieventtrigger
type IEventTrigger struct{ ITrigger }

type _IEventTriggerVt struct {
	_ITriggerVt
	Get_Subscription uintptr
	Put_Subscription uintptr
	Get_Delay        uintptr
	Put_Delay        uintptr
	Get_ValueQueries uintptr
	Put_ValueQueries uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEventTrigger) IID() co.IID {
	return co.IID_IEventTrigger
}

// [get_Delay] method.
//
// [get_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-get_delay
func (me *IEventTrigger) GetDelay() (string, error) {
	return com_callBstrGet(me,
		(*_IEventTriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_Delay)
}

// [get_Subscription] method.
//
// [get_Subscription]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-get_subscription
func (me *IEventTrigger) GetSubscription() (string, error) {
	return com_callBstrGet(me,
		(*_IEventTriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_Subscription)
}

// [get_ValueQueries] method.
//
// [get_ValueQueries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-get_valuequeries
func (me *IEventTrigger) GetValueQueries(releaser *OleReleaser) (*ITaskNamedValueCollection, error) {
	return com_callObj[*ITaskNamedValueCollection](me, releaser,
		(*_IEventTriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_ValueQueries)
}

// [put_Delay] method.
//
// [put_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-put_delay
func (me *IEventTrigger) PutDelay(delay string) error {
	return com_callBstrSet(me, delay,
		(*_IEventTriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_Delay)
}

// [put_Subscription] method.
//
// [put_Subscription]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-put_subscription
func (me *IEventTrigger) PutSubscription(subscription string) error {
	return com_callBstrSet(me, subscription,
		(*_IEventTriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_Subscription)
}

// [put_ValueQueries] method.
//
// [put_ValueQueries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-put_valuequeries
func (me *IEventTrigger) PutValueQueries(namedXPaths *ITaskNamedValueCollection) error {
	ret, _, _ := syscall.SyscallN(
		(*_IEventTriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_ValueQueries,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(namedXPaths.Ppvt())))
	return utl.HresultToError(ret)
}

// [IExecAction] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IExecAction]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iexecaction
type IExecAction struct{ IAction }

type _IExecActionVt struct {
	_IActionVt
	Get_Path             uintptr
	Put_Path             uintptr
	Get_Arguments        uintptr
	Put_Arguments        uintptr
	Get_WorkingDirectory uintptr
	Put_WorkingDirectory uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IExecAction) IID() co.IID {
	return co.IID_IExecAction
}

// [get_Arguments] method.
//
// [get_Arguments]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-get_arguments
func (me *IExecAction) GetArguments() (string, error) {
	return com_callBstrGet(me,
		(*_IExecActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Arguments)
}

// [get_Path] method.
//
// [get_Path]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-get_path
func (me *IExecAction) GetPath() (string, error) {
	return com_callBstrGet(me,
		(*_IExecActionVt)(unsafe.Pointer(*me.Ppvt())).Get_Path)
}

// [get_WorkingDirectory] method.
//
// [get_WorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-get_workingdirectory
func (me *IExecAction) GetWorkingDirectory() (string, error) {
	return com_callBstrGet(me,
		(*_IExecActionVt)(unsafe.Pointer(*me.Ppvt())).Get_WorkingDirectory)
}

// [put_Arguments] method.
//
// [put_Arguments]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-put_arguments
func (me *IExecAction) PutArguments(arguments string) error {
	return com_callBstrSet(me, arguments,
		(*_IExecActionVt)(unsafe.Pointer(*me.Ppvt())).Put_Arguments)
}

// [put_Path] method.
//
// [put_Path]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-put_path
func (me *IExecAction) PutPath(path string) error {
	return com_callBstrSet(me, path,
		(*_IExecActionVt)(unsafe.Pointer(*me.Ppvt())).Put_Path)
}

// [put_WorkingDirectory] method.
//
// [put_WorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-put_workingdirectory
func (me *IExecAction) PutWorkingDirectory(workingDirectory string) error {
	return com_callBstrSet(me, workingDirectory,
		(*_IExecActionVt)(unsafe.Pointer(*me.Ppvt())).Put_WorkingDirectory)
}

// [ILogonTrigger] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ILogonTrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iLogontrigger
type ILogonTrigger struct{ ITrigger }

type _ILogonTriggerVt struct {
	_ITriggerVt
	Get_Delay  uintptr
	Put_Delay  uintptr
	Get_UserId uintptr
	Put_UserId uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ILogonTrigger) IID() co.IID {
	return co.IID_ILogonTrigger
}

// [get_Delay] method.
//
// [get_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ilogontrigger-get_delay
func (me *ILogonTrigger) GetDelay() (string, error) {
	return com_callBstrGet(me,
		(*_ILogonTriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_Delay)
}

// [get_UserId] method.
//
// [get_UserId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ilogontrigger-get_userid
func (me *ILogonTrigger) GetUserId() (string, error) {
	return com_callBstrGet(me,
		(*_ILogonTriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_UserId)
}

// [put_Delay] method.
//
// [put_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ilogontrigger-put_delay
func (me *ILogonTrigger) PutDelay(delay string) error {
	return com_callBstrSet(me, delay,
		(*_ILogonTriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_Delay)
}

// [put_UserId] method.
//
// [put_UserId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ilogontrigger-put_userid
func (me *ILogonTrigger) PutUserId(userId string) error {
	return com_callBstrSet(me, userId,
		(*_ILogonTriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_UserId)
}

// [IPrincipal] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IPrincipal]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iprincipal
type IPrincipal struct{ IDispatch }

type _IPrincipalVt struct {
	_IDispatchVt
	Get_Id          uintptr
	Put_Id          uintptr
	Get_DisplayName uintptr
	Put_DisplayName uintptr
	Get_UserId      uintptr
	Put_UserId      uintptr
	Get_LogonType   uintptr
	Put_LogonType   uintptr
	Get_GroupId     uintptr
	Put_GroupId     uintptr
	Get_RunLevel    uintptr
	Put_RunLevel    uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IPrincipal) IID() co.IID {
	return co.IID_IPrincipal
}

// [get_DisplayName] method.
//
// [get_DisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_displayname
func (me *IPrincipal) GetDisplayName() (string, error) {
	return com_callBstrGet(me,
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Get_DisplayName)
}

// [get_GroupId] method.
//
// [get_GroupId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_groupid
func (me *IPrincipal) GetGroupId() (string, error) {
	return com_callBstrGet(me,
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Get_GroupId)
}

// [get_Id] method.
//
// [get_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_id
func (me *IPrincipal) GetId() (string, error) {
	return com_callBstrGet(me,
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Get_Id)
}

// [get_LogonType] method.
//
// [get_LogonType]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_logontype
func (me *IPrincipal) GetLogonType() (co.TASK_LOGON, error) {
	var taskLogon uint32
	ret, _, _ := syscall.SyscallN(
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Get_LogonType,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&taskLogon)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.TASK_LOGON(taskLogon), nil
	} else {
		return co.TASK_LOGON(0), hr
	}
}

// [get_RunLevel] method.
//
// [get_RunLevel]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_runlevel
func (me *IPrincipal) GetRunLevel() (co.TASK_RUNLEVEL, error) {
	var rl uint32
	ret, _, _ := syscall.SyscallN(
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Get_RunLevel,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&rl)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.TASK_RUNLEVEL(rl), nil
	} else {
		return co.TASK_RUNLEVEL(0), hr
	}
}

// [get_UserId] method.
//
// [get_UserId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_userid
func (me *IPrincipal) GetUserId() (string, error) {
	return com_callBstrGet(me,
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Get_UserId)
}

// [put_DisplayName] method.
//
// [put_DisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_displayname
func (me *IPrincipal) PutDisplayName(name string) error {
	return com_callBstrSet(me, name,
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Put_DisplayName)
}

// [put_GroupId] method.
//
// [put_GroupId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_groupid
func (me *IPrincipal) PutGroupId(group string) error {
	return com_callBstrSet(me, group,
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Put_GroupId)
}

// [put_Id] method.
//
// [put_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_id
func (me *IPrincipal) PutId(id string) error {
	return com_callBstrSet(me, id,
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Put_Id)
}

// [put_LogonType] method.
//
// [put_LogonType]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_logontype
func (me *IPrincipal) PutLogonType(logon co.TASK_LOGON) error {
	ret, _, _ := syscall.SyscallN(
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Put_LogonType,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(logon))
	return utl.HresultToError(ret)
}

// [put_RunLevel] method.
//
// [put_RunLevel]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_runlevel
func (me *IPrincipal) PutRunLevel(runLevel co.TASK_RUNLEVEL) error {
	ret, _, _ := syscall.SyscallN(
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Put_RunLevel,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&runLevel)))
	return utl.HresultToError(ret)
}

// [put_UserId] method.
//
// [put_UserId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_userid
func (me *IPrincipal) PutUserId(user string) error {
	return com_callBstrSet(me, user,
		(*_IPrincipalVt)(unsafe.Pointer(*me.Ppvt())).Put_UserId)
}

// [IRegisteredTask] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IRegisteredTask]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iregisteredtask
type IRegisteredTask struct{ IDispatch }

type _IRegisteredTaskVt struct {
	_IDispatchVt
	Get_Name               uintptr
	Get_Path               uintptr
	Get_State              uintptr
	Get_Enabled            uintptr
	Put_Enabled            uintptr
	Run                    uintptr
	RunEx                  uintptr
	GetInstances           uintptr
	Get_LastRunTime        uintptr
	Get_LastTaskResult     uintptr
	Get_NumberOfMissedRuns uintptr
	Get_NextRunTime        uintptr
	Get_Definition         uintptr
	Get_Xml                uintptr
	GetSecurityDescriptor  uintptr
	SetSecurityDescriptor  uintptr
	Stop                   uintptr
	GetRunTimes            uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IRegisteredTask) IID() co.IID {
	return co.IID_IRegisteredTask
}

// [get_Definition] method.
//
// [get_Definition]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_definition
func (me *IRegisteredTask) GetDefinition(releaser *OleReleaser) (*ITaskDefinition, error) {
	return com_callObj[*ITaskDefinition](me, releaser,
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Get_Definition)
}

// [get_Enabled] method.
//
// [get_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_enabled
func (me *IRegisteredTask) GetEnabled() (bool, error) {
	var enabled int16
	ret, _, _ := syscall.SyscallN(
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Get_Enabled,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&enabled)))
	return utl.HresultToBoolError(int32(enabled), ret)
}

// [get_LastTaskResult] method.
//
// [get_LastTaskResult]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_lasttaskresult
func (me *IRegisteredTask) GetLastTaskResult() (int, error) {
	var last int32
	ret, _, _ := syscall.SyscallN(
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Get_LastTaskResult,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&last)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(last), nil
	} else {
		return 0, hr
	}
}

// [get_Name] method.
//
// [get_Name]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_name
func (me *IRegisteredTask) GetName() (string, error) {
	return com_callBstrGet(me,
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Get_Name)
}

// [get_NumberOfMissedRuns] method.
//
// [get_NumberOfMissedRuns]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_numberofmissedruns
func (me *IRegisteredTask) GetNumberOfMissedRuns() (int, error) {
	var nmr int32
	ret, _, _ := syscall.SyscallN(
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Get_NumberOfMissedRuns,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&nmr)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(nmr), nil
	} else {
		return 0, hr
	}
}

// [get_Path] method.
//
// [get_Path]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_path
func (me *IRegisteredTask) GetPath() (string, error) {
	return com_callBstrGet(me,
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Get_Path)
}

// [GetRunTimes] method.
//
// [GetRunTimes]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-getruntimes
func (me *IRegisteredTask) GetRunTimes(pStart, pEnd *SYSTEMTIME, count int) ([]SYSTEMTIME, error) {
	count32 := uint32(count)
	var pRunTimes HTASKMEM
	defer pRunTimes.CoTaskMemFree()

	ret, _, _ := syscall.SyscallN(
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).GetRunTimes,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(pStart)),
		uintptr(unsafe.Pointer(pEnd)),
		uintptr(unsafe.Pointer(&count32)),
		uintptr(unsafe.Pointer(&pRunTimes)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		hTaskST := unsafe.Slice((*SYSTEMTIME)(unsafe.Pointer(pRunTimes)), count32)
		buf := make([]SYSTEMTIME, 0, count32)
		for _, st := range hTaskST {
			buf = append(buf, st)
		}
		return buf, nil
	} else {
		return nil, hr
	}
}

// [get_State] method.
//
// [get_State]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_state
func (me *IRegisteredTask) GetState() (co.TASK_STATE, error) {
	var state uint32
	ret, _, _ := syscall.SyscallN(
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Get_State,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&state)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.TASK_STATE(state), nil
	} else {
		return co.TASK_STATE(0), hr
	}
}

// [get_Xml] method.
//
// [get_Xml]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_xml
func (me *IRegisteredTask) GetXml() (string, error) {
	return com_callBstrGet(me,
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Get_Xml)
}

// [put_Enabled] method.
//
// [put_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-put_enabled
func (me *IRegisteredTask) PutEnabled(enabled bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Put_Enabled,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(enabled))
	return utl.HresultToError(ret)
}

// [Stop] method.
//
// [Stop]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-stop
func (me *IRegisteredTask) Stop() error {
	ret, _, _ := syscall.SyscallN(
		(*_IRegisteredTaskVt)(unsafe.Pointer(*me.Ppvt())).Stop,
		uintptr(unsafe.Pointer(me.Ppvt())),
		0)
	return utl.HresultToError(ret)
}

// [IRegistrationInfo] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IRegistrationInfo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iregistrationinfo
type IRegistrationInfo struct{ IDispatch }

type _IRegistrationInfoVt struct {
	_IDispatchVt
	Get_Description        uintptr
	Put_Description        uintptr
	Get_Author             uintptr
	Put_Author             uintptr
	Get_Version            uintptr
	Put_Version            uintptr
	Get_Date               uintptr
	Put_Date               uintptr
	Get_Documentation      uintptr
	Put_Documentation      uintptr
	Get_XmlText            uintptr
	Put_XmlText            uintptr
	Get_URI                uintptr
	Put_URI                uintptr
	Get_SecurityDescriptor uintptr
	Put_SecurityDescriptor uintptr
	Get_Source             uintptr
	Put_Source             uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IRegistrationInfo) IID() co.IID {
	return co.IID_IRegistrationInfo
}

// [get_Author] method.
//
// [get_Author]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_author
func (me *IRegistrationInfo) GetAuthor() (string, error) {
	return com_callBstrGet(me,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Get_Author)
}

// [get_Date] method.
//
// [get_Date]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_date
func (me *IRegistrationInfo) GetDate() (string, error) {
	return com_callBstrGet(me,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Get_Date)
}

// [get_Description] method.
//
// [get_Description]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_description
func (me *IRegistrationInfo) GetDescription() (string, error) {
	return com_callBstrGet(me,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Get_Description)
}

// [get_Documentation] method.
//
// [get_Documentation]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_documentation
func (me *IRegistrationInfo) GetDocumentation() (string, error) {
	return com_callBstrGet(me,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Get_Documentation)
}

// [get_Source] method.
//
// [get_Source]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_source
func (me *IRegistrationInfo) GetSource() (string, error) {
	return com_callBstrGet(me,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Get_Source)
}

// [get_URI] method.
//
// [get_URI]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_uri
func (me *IRegistrationInfo) GetURI() (string, error) {
	return com_callBstrGet(me,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Get_URI)
}

// [get_Version] method.
//
// [get_Version]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_version
func (me *IRegistrationInfo) GetVersion() (string, error) {
	return com_callBstrGet(me,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Get_Version)
}

// [get_XmlText] method.
//
// [get_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_xmltext
func (me *IRegistrationInfo) GetXmlText() (string, error) {
	return com_callBstrGet(me,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Get_XmlText)
}

// [put_Author] method.
//
// [put_Author]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_author
func (me *IRegistrationInfo) PutAuthor(author string) error {
	return com_callBstrSet(me, author,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Put_Author)
}

// [put_Date] method.
//
// [put_Date]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_date
func (me *IRegistrationInfo) PutDate(date string) error {
	return com_callBstrSet(me, date,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Put_Date)
}

// [put_Description] method.
//
// [put_Description]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_description
func (me *IRegistrationInfo) PutDescription(description string) error {
	return com_callBstrSet(me, description,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Put_Description)
}

// [put_Documentation] method.
//
// [put_Documentation]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_documentation
func (me *IRegistrationInfo) PutDocumentation(documentation string) error {
	return com_callBstrSet(me, documentation,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Put_Documentation)
}

// [put_Source] method.
//
// [put_Source]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_source
func (me *IRegistrationInfo) PutSource(source string) error {
	return com_callBstrSet(me, source,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Put_Source)
}

// [put_URI] method.
//
// [put_URI]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_uri
func (me *IRegistrationInfo) PutURI(uri string) error {
	return com_callBstrSet(me, uri,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Put_URI)
}

// [put_Version] method.
//
// [put_Version]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_version
func (me *IRegistrationInfo) PutVersion(version string) error {
	return com_callBstrSet(me, version,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Put_Version)
}

// [put_XmlText] method.
//
// [put_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_xmltext
func (me *IRegistrationInfo) PutXmlText(text string) error {
	return com_callBstrSet(me, text,
		(*_IRegistrationInfoVt)(unsafe.Pointer(*me.Ppvt())).Put_XmlText)
}

// [ITaskDefinition] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITaskDefinition]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itaskdefinition
type ITaskDefinition struct{ IDispatch }

type _ITaskDefinitionVt struct {
	_IDispatchVt
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
func (*ITaskDefinition) IID() co.IID {
	return co.IID_ITaskDefinition
}

// [get_Actions] method.
//
// [get_Actions]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_actions
func (me *ITaskDefinition) GetActions(releaser *OleReleaser) (*IActionCollection, error) {
	return com_callObj[*IActionCollection](me, releaser,
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Get_Actions)
}

// [get_Data] method.
//
// [get_Data]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_data
func (me *ITaskDefinition) GetData() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Get_Data)
}

// [get_Principal] method.
//
// [get_Principal]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_principal
func (me *ITaskDefinition) GetPrincipal(releaser *OleReleaser) (*IPrincipal, error) {
	return com_callObj[*IPrincipal](me, releaser,
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Get_Principal)
}

// [get_RegistrationInfo] method.
//
// [get_RegistrationInfo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_registrationinfo
func (me *ITaskDefinition) GetRegistrationInfo(releaser *OleReleaser) (*IRegistrationInfo, error) {
	return com_callObj[*IRegistrationInfo](me, releaser,
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Get_RegistrationInfo)
}

// [get_Settings] method.
//
// [get_Settings]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_settings
func (me *ITaskDefinition) GetSettings(releaser *OleReleaser) (*ITaskSettings, error) {
	return com_callObj[*ITaskSettings](me, releaser,
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Get_Settings)
}

// [get_Triggers] method.
//
// [get_Triggers]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_triggers
func (me *ITaskDefinition) GetTriggers(releaser *OleReleaser) (*ITriggerCollection, error) {
	return com_callObj[*ITriggerCollection](me, releaser,
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Get_Triggers)
}

// [get_XmlText] method.
//
// [get_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-get_xmltext
func (me *ITaskDefinition) GetXmlText() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Get_XmlText)
}

// [put_Data] method.
//
// [put_Data]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_data
func (me *ITaskDefinition) PutData(data string) error {
	return com_callBstrSet(me, data,
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Put_Data)
}

// [put_Principal] method.
//
// [put_Principal]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_principal
func (me *ITaskDefinition) PutPrincipal(principal *IPrincipal) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Put_Principal,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(principal.Ppvt())))
	return utl.HresultToError(ret)
}

// [put_RegistrationInfo] method.
//
// [put_RegistrationInfo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_registrationinfo
func (me *ITaskDefinition) PutRegistrationInfo(registrationInfo *IRegistrationInfo) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Put_RegistrationInfo,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(registrationInfo.Ppvt())))
	return utl.HresultToError(ret)
}

// [put_Settings] method.
//
// [put_Settings]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_settings
func (me *ITaskDefinition) PutSettings(settings *ITaskSettings) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Put_Settings,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(settings.Ppvt())))
	return utl.HresultToError(ret)
}

// [put_Triggers] method.
//
// [put_Triggers]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_triggers
func (me *ITaskDefinition) PutTriggers(triggers *ITriggerCollection) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Put_Triggers,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(triggers.Ppvt())))
	return utl.HresultToError(ret)
}

// [put_XmlText] method.
//
// [put_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskdefinition-put_xmltext
func (me *ITaskDefinition) PutXmlText(xml string) error {
	return com_callBstrSet(me, xml,
		(*_ITaskDefinitionVt)(unsafe.Pointer(*me.Ppvt())).Put_XmlText)
}

// [ITaskFolder] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITaskFolder]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itaskfolder
type ITaskFolder struct{ IDispatch }

type _ITaskFolderVt struct {
	_IDispatchVt
	Get_Name               uintptr
	Get_Path               uintptr
	GetFolder              uintptr
	GetFolders             uintptr
	CreateFolder           uintptr
	DeleteFolder           uintptr
	GetTask                uintptr
	GetTasks               uintptr
	DeleteTask             uintptr
	RegisterTask           uintptr
	RegisterTaskDefinition uintptr
	GetSecurityDescriptor  uintptr
	SetSecurityDescriptor  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskFolder) IID() co.IID {
	return co.IID_ITaskFolder
}

// [DeleteFolder] method.
//
// [DeleteFolder]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-deletefolder
func (me *ITaskFolder) DeleteFolder(subFolderName string) error {
	bstrSubFolderName, err := SysAllocString(subFolderName)
	if err != nil {
		return err
	}
	defer bstrSubFolderName.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		(*_ITaskFolderVt)(unsafe.Pointer(*me.Ppvt())).DeleteFolder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(bstrSubFolderName),
		0)
	return utl.HresultToError(ret)
}

// [DeleteTask] method.
//
// [DeleteTask]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-deletetask
func (me *ITaskFolder) DeleteTask(name string) error {
	bstrName, err := SysAllocString(name)
	if err != nil {
		return err
	}
	defer bstrName.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		(*_ITaskFolderVt)(unsafe.Pointer(*me.Ppvt())).DeleteTask,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(bstrName),
		0)
	return utl.HresultToError(ret)
}

// [get_Name] method.
//
// [get_Name]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-get_name
func (me *ITaskFolder) GetName() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskFolderVt)(unsafe.Pointer(*me.Ppvt())).Get_Name)
}

// [get_Path] method.
//
// [get_Path]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-get_path
func (me *ITaskFolder) GetPath() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskFolderVt)(unsafe.Pointer(*me.Ppvt())).Get_Path)
}

// [GetFolder] method.
//
// [GetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-getfolder
func (me *ITaskFolder) GetFolder(releaser *OleReleaser, path string) (*ITaskFolder, error) {
	var ppvtQueried **_IUnknownVt

	bstrPath, err := SysAllocString(path)
	if err != nil {
		return nil, err
	}
	defer bstrPath.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		(*_ITaskFolderVt)(unsafe.Pointer(*me.Ppvt())).GetFolder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(bstrPath),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*ITaskFolder](ret, ppvtQueried, releaser)
}

// [GetTask] method.
//
// [GetTask]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-gettask
func (me *ITaskFolder) GetTask(releaser *OleReleaser, path string) (*IRegisteredTask, error) {
	var ppvtQueried **_IUnknownVt

	bstrPath, err := SysAllocString(path)
	if err != nil {
		return nil, err
	}
	defer bstrPath.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		(*_ITaskFolderVt)(unsafe.Pointer(*me.Ppvt())).GetTask,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(bstrPath),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*IRegisteredTask](ret, ppvtQueried, releaser)
}

// [ITaskNamedValueCollection] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITaskNamedValueCollection]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itasknamedvaluecollection
type ITaskNamedValueCollection struct{ IDispatch }

type _ITaskNamedValueCollectionVt struct {
	_IDispatchVt
	Get_Count    uintptr
	Get_Item     uintptr
	Get__NewEnum uintptr
	Create       uintptr
	Remove       uintptr
	Clear        uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskNamedValueCollection) IID() co.IID {
	return co.IID_ITaskNamedValueCollection
}

// [Clear] method;
//
// [Clear]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-clear
func (me *ITaskNamedValueCollection) Clear() error {
	return com_callErr(me,
		(*_ITaskNamedValueCollectionVt)(unsafe.Pointer(*me.Ppvt())).Clear)
}

// [Create] method.
//
// [Create]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-create
func (me *ITaskNamedValueCollection) Create(
	releaser *OleReleaser,
	name, value string,
) (*ITaskNamedValuePair, error) {
	var ppvtQueried **_IUnknownVt

	bstrName, err := SysAllocString(name)
	if err != nil {
		return nil, err
	}
	defer bstrName.SysFreeString()

	bstrValue, err := SysAllocString(value)
	if err != nil {
		return nil, err
	}
	defer bstrValue.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		(*_ITaskNamedValueCollectionVt)(unsafe.Pointer(*me.Ppvt())).Create,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(bstrName),
		uintptr(bstrValue),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*ITaskNamedValuePair](ret, ppvtQueried, releaser)
}

// Returns all [ITaskNamedValuePair] objects by calling
// [ITaskNamedValueCollection.GetCount] and [ITaskNamedValueCollection.GetItem].
func (me *ITaskNamedValueCollection) Enum(releaser *OleReleaser) ([]*ITaskNamedValuePair, error) {
	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	pairs := make([]*ITaskNamedValuePair, count)
	for i := 0; i < count; i++ {
		pair, err := me.GetItem(releaser, i)
		if err != nil {
			return nil, err // stop immediately
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}

// [get_Count] method.
//
// [get_Count]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-get_count
func (me *ITaskNamedValueCollection) GetCount() (int, error) {
	var count int32
	ret, _, _ := syscall.SyscallN(
		(*_ITaskNamedValueCollectionVt)(unsafe.Pointer(*me.Ppvt())).Get_Count,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&count)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(count), nil
	} else {
		return 0, hr
	}
}

// [get_Item] method.
//
// [get_Item]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-get_item
func (me *ITaskNamedValueCollection) GetItem(releaser *OleReleaser, index int) (*ITaskNamedValuePair, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_ITaskNamedValueCollectionVt)(unsafe.Pointer(*me.Ppvt())).Get_Item,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(int32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*ITaskNamedValuePair](ret, ppvtQueried, releaser)
}

// [Remove] method.
//
// [Remove]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-remove
func (me *ITaskNamedValueCollection) Remove(index int) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskNamedValueCollectionVt)(unsafe.Pointer(*me.Ppvt())).Remove,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(int32(index)))
	return utl.HresultToError(ret)
}

// [ITaskNamedValuePair] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITaskNamedValuePair]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itasknamedvaluepair
type ITaskNamedValuePair struct{ IDispatch }

type _ITaskNamedValuePairVt struct {
	_IDispatchVt
	Get_Name  uintptr
	Put_Name  uintptr
	Get_Value uintptr
	Put_Value uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskNamedValuePair) IID() co.IID {
	return co.IID_ITaskNamedValuePair
}

// [get_Name] method.
//
// [get_Name]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluepair-get_name
func (me *ITaskNamedValuePair) GetName() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskNamedValuePairVt)(unsafe.Pointer(*me.Ppvt())).Get_Name)
}

// [get_Value] method.
//
// [get_Value]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluepair-get_value
func (me *ITaskNamedValuePair) GetValue() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskNamedValuePairVt)(unsafe.Pointer(*me.Ppvt())).Get_Value)
}

// [put_Name] method.
//
// [put_Name]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluepair-put_name
func (me *ITaskNamedValuePair) PutName(name string) error {
	return com_callBstrSet(me, name,
		(*_ITaskNamedValuePairVt)(unsafe.Pointer(*me.Ppvt())).Put_Name)
}

// [put_Value] method.
//
// [put_Value]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluepair-put_value
func (me *ITaskNamedValuePair) PutValue(value string) error {
	return com_callBstrSet(me, value,
		(*_ITaskNamedValuePairVt)(unsafe.Pointer(*me.Ppvt())).Put_Value)
}

// [ITaskService] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITaskService]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itaskservice
type ITaskService struct{ IDispatch }

type _ITaskServiceVt struct {
	_IDispatchVt
	GetFolder           uintptr
	GetRunningTasks     uintptr
	NewTask             uintptr
	Connect             uintptr
	Get_Connected       uintptr
	Get_TargetServer    uintptr
	Get_ConnectedUser   uintptr
	Get_ConnectedDomain uintptr
	Get_HighestVersion  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskService) IID() co.IID {
	return co.IID_ITaskService
}

// [Connect] method.
//
// [Connect]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-connect
func (me *ITaskService) Connect(serverName, user, domain, password string) error {
	localRel := NewOleReleaser()
	defer localRel.Release()

	vServerName := NewVariant(localRel, serverName)
	vUser := NewVariant(localRel, user)
	vDomain := NewVariant(localRel, domain)
	vPassword := NewVariant(localRel, password)

	ret, _, _ := syscall.SyscallN(
		(*_ITaskServiceVt)(unsafe.Pointer(*me.Ppvt())).Connect,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(vServerName)),
		uintptr(unsafe.Pointer(vUser)),
		uintptr(unsafe.Pointer(vDomain)),
		uintptr(unsafe.Pointer(vPassword)))
	return utl.HresultToError(ret)
}

// [get_Connected] method.
//
// [get_Connected]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_connected
func (me *ITaskService) GetConnected() (bool, error) {
	var connected int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskServiceVt)(unsafe.Pointer(*me.Ppvt())).Get_Connected,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&connected)))
	return utl.HresultToBoolError(int32(connected), ret)
}

// [get_ConnectedDomain] method.
//
// [get_ConnectedDomain]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_connecteddomain
func (me *ITaskService) GetConnectedDomain() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskServiceVt)(unsafe.Pointer(*me.Ppvt())).Get_ConnectedDomain)
}

// [get_ConnectedUser] method.
//
// [get_ConnectedUser]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_connecteduser
func (me *ITaskService) GetConnectedUser() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskServiceVt)(unsafe.Pointer(*me.Ppvt())).Get_ConnectedUser)
}

// [get_HighestVersion] method.
//
// [get_HighestVersion]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_highestversion
func (me *ITaskService) GetHighestVersion() (int, error) {
	var version uint32
	ret, _, _ := syscall.SyscallN(
		(*_ITaskServiceVt)(unsafe.Pointer(*me.Ppvt())).Get_HighestVersion,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&version)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(version), nil
	} else {
		return 0, hr
	}
}

// [get_TargetServer] method.
//
// [get_TargetServer]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_targetserver
func (me *ITaskService) GetTargetServer() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskServiceVt)(unsafe.Pointer(*me.Ppvt())).Get_TargetServer)
}

// [GetFolder] method.
//
// [GetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-getfolder
func (me *ITaskService) GetFolder(releaser *OleReleaser, path string) (*ITaskFolder, error) {
	var ppvtQueried **_IUnknownVt

	bstrPath, err := SysAllocString(path)
	if err != nil {
		return nil, err
	}
	defer bstrPath.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		(*_ITaskServiceVt)(unsafe.Pointer(*me.Ppvt())).GetFolder,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(bstrPath),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*ITaskFolder](ret, ppvtQueried, releaser)
}

// [NewTask] method.
//
// [NewTask]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-newtask
func (me *ITaskService) NewTask(releaser *OleReleaser) (*ITaskDefinition, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_ITaskServiceVt)(unsafe.Pointer(*me.Ppvt())).NewTask,
		uintptr(unsafe.Pointer(me.Ppvt())),
		0,
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*ITaskDefinition](ret, ppvtQueried, releaser)
}

// [ITaskSettings] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITaskSettings]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itasksettings
type ITaskSettings struct{ IDispatch }

type _ITaskSettingsVt struct {
	_IDispatchVt
	Get_AllowDemandStart           uintptr
	Put_AllowDemandStart           uintptr
	Get_RestartInterval            uintptr
	Put_RestartInterval            uintptr
	Get_RestartCount               uintptr
	Put_RestartCount               uintptr
	Get_MultipleInstances          uintptr
	Put_MultipleInstances          uintptr
	Get_StopIfGoingOnBatteries     uintptr
	Put_StopIfGoingOnBatteries     uintptr
	Get_DisallowStartIfOnBatteries uintptr
	Put_DisallowStartIfOnBatteries uintptr
	Get_AllowHardTerminate         uintptr
	Put_AllowHardTerminate         uintptr
	Get_StartWhenAvailable         uintptr
	Put_StartWhenAvailable         uintptr
	Get_XmlText                    uintptr
	Put_XmlText                    uintptr
	Get_RunOnlyIfNetworkAvailable  uintptr
	Put_RunOnlyIfNetworkAvailable  uintptr
	Get_ExecutionTimeLimit         uintptr
	Put_ExecutionTimeLimit         uintptr
	Get_Enabled                    uintptr
	Put_Enabled                    uintptr
	Get_DeleteExpiredTaskAfter     uintptr
	Put_DeleteExpiredTaskAfter     uintptr
	Get_Priority                   uintptr
	Put_Priority                   uintptr
	Get_Compatibility              uintptr
	Put_Compatibility              uintptr
	Get_Hidden                     uintptr
	Put_Hidden                     uintptr
	Get_IdleSettings               uintptr
	Put_IdleSettings               uintptr
	Get_RunOnlyIfIdle              uintptr
	Put_RunOnlyIfIdle              uintptr
	Get_WakeToRun                  uintptr
	Put_WakeToRun                  uintptr
	Get_NetworkSettings            uintptr
	Put_NetworkSettings            uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskSettings) IID() co.IID {
	return co.IID_ITaskSettings
}

// [get_AllowDemandStart] method.
//
// [get_AllowDemandStart]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_allowdemandstart
func (me *ITaskSettings) GetAllowDemandStart() (bool, error) {
	var allow int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_AllowDemandStart,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&allow)))
	return utl.HresultToBoolError(int32(allow), ret)
}

// [get_AllowHardTerminate] method.
//
// [get_AllowHardTerminate]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_allowhardterminate
func (me *ITaskSettings) GetAllowHardTerminate() (bool, error) {
	var allow int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_AllowHardTerminate,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&allow)))
	return utl.HresultToBoolError(int32(allow), ret)
}

// [get_Compatibility] method.
//
// [get_Compatibility]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_compatibility
func (me *ITaskSettings) GetCompatibility() (co.TASK_COMPATIBILITY, error) {
	var compat uint32
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_Compatibility,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&compat)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.TASK_COMPATIBILITY(compat), nil
	} else {
		return co.TASK_COMPATIBILITY(0), hr
	}
}

// [get_DeleteExpiredTaskAfter] method.
//
// [get_DeleteExpiredTaskAfter]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_deleteexpiredtaskafter
func (me *ITaskSettings) GetDeleteExpiredTaskAfter() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_DeleteExpiredTaskAfter)
}

// [get_DisallowStartIfOnBatteries] method.
//
// [get_DisallowStartIfOnBatteries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_disallowstartifonbatteries
func (me *ITaskSettings) GetDisallowStartIfOnBatteries() (bool, error) {
	var disallow int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_DisallowStartIfOnBatteries,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&disallow)))
	return utl.HresultToBoolError(int32(disallow), ret)
}

// [get_Enabled] method.
//
// [get_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_enabled
func (me *ITaskSettings) GetEnabled() (bool, error) {
	var enabled int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_Enabled,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&enabled)))
	return utl.HresultToBoolError(int32(enabled), ret)
}

// [get_ExecutionTimeLimit] method.
//
// [get_ExecutionTimeLimit]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_executiontimelimit
func (me *ITaskSettings) GetExecutionTimeLimit() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_ExecutionTimeLimit)
}

// [get_Hidden] method.
//
// [get_Hidden]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_hidden
func (me *ITaskSettings) GetHidden() (bool, error) {
	var hidden int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_Hidden,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&hidden)))
	return utl.HresultToBoolError(int32(hidden), ret)
}

// [get_MultipleInstances] method.
//
// [get_MultipleInstances]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_multipleinstances
func (me *ITaskSettings) GetMultipleInstances() (co.TASK_INSTANCES, error) {
	var compat uint32
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_MultipleInstances,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&compat)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.TASK_INSTANCES(compat), nil
	} else {
		return co.TASK_INSTANCES(0), hr
	}
}

// [get_Priority] method.
//
// [get_Priority]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_priority
func (me *ITaskSettings) GetPriority() (int, error) {
	var priority int32
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_Priority,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&priority)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(priority), nil
	} else {
		return 0, hr
	}
}

// [get_RestartCount] method.
//
// [get_RestartCount]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_restartcount
func (me *ITaskSettings) GetRestartCount() (int, error) {
	var priority int32
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_RestartCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&priority)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(priority), nil
	} else {
		return 0, hr
	}
}

// [get_RestartInterval] method.
//
// [get_RestartInterval]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_restartinterval
func (me *ITaskSettings) GetRestartInterval() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_RestartInterval)
}

// [get_RunOnlyIfIdle] method.
//
// [get_RunOnlyIfIdle]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_runonlyifidle
func (me *ITaskSettings) GetRunOnlyIfIdle() (bool, error) {
	var runOnlyIfIdle int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_RunOnlyIfIdle,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&runOnlyIfIdle)))
	return utl.HresultToBoolError(int32(runOnlyIfIdle), ret)
}

// [get_RunOnlyIfNetworkAvailable] method.
//
// [get_RunOnlyIfNetworkAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_runonlyifnetworkavailable
func (me *ITaskSettings) GetRunOnlyIfNetworkAvailable() (bool, error) {
	var runOnlyIfNetworkAvailable int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_RunOnlyIfNetworkAvailable,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&runOnlyIfNetworkAvailable)))
	return utl.HresultToBoolError(int32(runOnlyIfNetworkAvailable), ret)
}

// [get_StartWhenAvailable] method.
//
// [get_StartWhenAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_startwhenavailable
func (me *ITaskSettings) GetStartWhenAvailable() (bool, error) {
	var startWhenAvailable int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_StartWhenAvailable,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&startWhenAvailable)))
	return utl.HresultToBoolError(int32(startWhenAvailable), ret)
}

// [get_StopIfGoingOnBatteries] method.
//
// [get_StopIfGoingOnBatteries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_stopifgoingonbatteries
func (me *ITaskSettings) GetStopIfGoingOnBatteries() (bool, error) {
	var stopIfGoingOnBatt int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_StopIfGoingOnBatteries,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&stopIfGoingOnBatt)))
	return utl.HresultToBoolError(int32(stopIfGoingOnBatt), ret)
}

// [get_WakeToRun] method.
//
// [get_WakeToRun]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_waketorun
func (me *ITaskSettings) GetWakeToRun() (bool, error) {
	var wakeToRun int16
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_WakeToRun,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&wakeToRun)))
	return utl.HresultToBoolError(int32(wakeToRun), ret)
}

// [get_XmlText] method.
//
// [get_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_xmltext
func (me *ITaskSettings) GetXmlText() (string, error) {
	return com_callBstrGet(me,
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Get_XmlText)
}

// [put_AllowDemandStart] method.
//
// [put_AllowDemandStart]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_allowdemandstart
func (me *ITaskSettings) PutAllowDemandStart(allowDemandStart bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_AllowDemandStart,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(allowDemandStart))
	return utl.HresultToError(ret)
}

// [put_AllowHardTerminate] method.
//
// [put_AllowHardTerminate]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_allowhardterminate
func (me *ITaskSettings) PutAllowHardTerminate(allowHardTerminate bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_AllowHardTerminate,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(allowHardTerminate))
	return utl.HresultToError(ret)
}

// [put_Compatibility] method.
//
// [put_Compatibility]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_compatibility
func (me *ITaskSettings) PutCompatibility(compatLevel co.TASK_COMPATIBILITY) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_Compatibility,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(compatLevel))
	return utl.HresultToError(ret)
}

// [put_DeleteExpiredTaskAfter] method.
//
// [put_DeleteExpiredTaskAfter]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_deleteexpiredtaskafter
func (me *ITaskSettings) PutDeleteExpiredTaskAfter(expirationDelay string) error {
	return com_callBstrSet(me, expirationDelay,
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_DeleteExpiredTaskAfter)
}

// [put_DisallowStartIfOnBatteries] method.
//
// [put_DisallowStartIfOnBatteries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_disallowstartifonbatteries
func (me *ITaskSettings) PutDisallowStartIfOnBatteries(disallowStart bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_DisallowStartIfOnBatteries,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(disallowStart))
	return utl.HresultToError(ret)
}

// [put_ExecutionTimeLimit] method.
//
// [put_ExecutionTimeLimit]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_executiontimelimit
func (me *ITaskSettings) PutExecutionTimeLimit(executionTimeLimit string) error {
	return com_callBstrSet(me, executionTimeLimit,
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_ExecutionTimeLimit)
}

// [put_Hidden] method.
//
// [put_Hidden]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_hidden
func (me *ITaskSettings) PutHidden(hidden bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_Hidden,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(hidden))
	return utl.HresultToError(ret)
}

// [put_MultipleInstances] method.
//
// [put_MultipleInstances]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_multipleinstances
func (me *ITaskSettings) PutMultipleInstances(policy co.TASK_INSTANCES) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_MultipleInstances,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(policy))
	return utl.HresultToError(ret)
}

// [put_Priority] method.
//
// [put_Priority]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_priority
func (me *ITaskSettings) PutPriority(priority int) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_Priority,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(int32(priority)))
	return utl.HresultToError(ret)
}

// [put_RestartCount] method.
//
// [put_RestartCount]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_restartcount
func (me *ITaskSettings) PutRestartCount(restartCount int) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_RestartCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(int32(restartCount)))
	return utl.HresultToError(ret)
}

// [put_RestartInterval] method.
//
// [put_RestartInterval]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_restartinterval
func (me *ITaskSettings) PutRestartInterval(restartInterval string) error {
	return com_callBstrSet(me, restartInterval,
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_RestartInterval)
}

// [put_RunOnlyIfIdle] method.
//
// [put_RunOnlyIfIdle]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_runonlyifidle
func (me *ITaskSettings) PutRunOnlyIfIdle(runOnlyIfIdle bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_RunOnlyIfIdle,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(runOnlyIfIdle))
	return utl.HresultToError(ret)
}

// [put_RunOnlyIfNetworkAvailable] method.
//
// [put_RunOnlyIfNetworkAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_runonlyifnetworkavailable
func (me *ITaskSettings) PutRunOnlyIfNetworkAvailable(runOnlyIfNetworkAvailable bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_RunOnlyIfNetworkAvailable,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(runOnlyIfNetworkAvailable))
	return utl.HresultToError(ret)
}

// [put_StartWhenAvailable] method.
//
// [put_StartWhenAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_startwhenavailable
func (me *ITaskSettings) PutStartWhenAvailable(startWhenAvailable bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_StartWhenAvailable,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(startWhenAvailable))
	return utl.HresultToError(ret)
}

// [put_StopIfGoingOnBatteries] method.
//
// [put_StopIfGoingOnBatteries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_stopifgoingonbatteries
func (me *ITaskSettings) PutStopIfGoingOnBatteries(stopIfOnBatteries bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_StartWhenAvailable,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(stopIfOnBatteries))
	return utl.HresultToError(ret)
}

// [put_WakeToRun] method.
//
// [put_WakeToRun]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_waketorun
func (me *ITaskSettings) PutWakeToRun(wake bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_WakeToRun,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(wake))
	return utl.HresultToError(ret)
}

// [put_XmlText] method.
//
// [put_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_xmltext
func (me *ITaskSettings) PutXmlText(text string) error {
	return com_callBstrSet(me, text,
		(*_ITaskSettingsVt)(unsafe.Pointer(*me.Ppvt())).Put_XmlText)
}

// [ITrigger] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itrigger
type ITrigger struct{ IDispatch }

type _ITriggerVt struct {
	_IDispatchVt
	Get_Type               uintptr
	Get_Id                 uintptr
	Put_Id                 uintptr
	Get_Repetition         uintptr
	Put_Repetition         uintptr
	Get_ExecutionTimeLimit uintptr
	Put_ExecutionTimeLimit uintptr
	Get_StartBoundary      uintptr
	Put_StartBoundary      uintptr
	Get_EndBoundary        uintptr
	Put_EndBoundary        uintptr
	Get_Enabled            uintptr
	Put_Enabled            uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITrigger) IID() co.IID {
	return co.IID_ITrigger
}

// [get_Enabled] method.
//
// [get_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_enabled
func (me *ITrigger) GetEnabled() (bool, error) {
	var enabled int16
	ret, _, _ := syscall.SyscallN(
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_Enabled,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&enabled)))
	return utl.HresultToBoolError(int32(enabled), ret)
}

// [get_EndBoundary] method.
//
// [get_EndBoundary]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_endboundary
func (me *ITrigger) GetEndBoundary() (string, error) {
	return com_callBstrGet(me,
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_EndBoundary)
}

// [get_ExecutionTimeLimit] method.
//
// [get_ExecutionTimeLimit]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_executiontimelimit
func (me *ITrigger) GetExecutionTimeLimit() (string, error) {
	return com_callBstrGet(me,
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_ExecutionTimeLimit)
}

// [get_Id] method.
//
// [get_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_id
func (me *ITrigger) GetId() (string, error) {
	return com_callBstrGet(me,
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_Id)
}

// [get_StartBoundary] method.
//
// [get_StartBoundary]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_startboundary
func (me *ITrigger) GetStartBoundary() (string, error) {
	return com_callBstrGet(me,
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_StartBoundary)
}

// [get_Type] method.
//
// [get_Type]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_type
func (me *ITrigger) GetType() (co.TASK_TRIGGER2, error) {
	var ty uint32
	ret, _, _ := syscall.SyscallN(
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Get_Type,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ty)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.TASK_TRIGGER2(ty), nil
	} else {
		return co.TASK_TRIGGER2(0), hr
	}
}

// [put_Enabled] method.
//
// [put_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_enabled
func (me *ITrigger) PutEnabled(enabled bool) error {
	ret, _, _ := syscall.SyscallN(
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_Enabled,
		uintptr(unsafe.Pointer(me.Ppvt())),
		utl.BoolToUintptr(enabled))
	return utl.HresultToError(ret)
}

// [put_EndBoundary] method.
//
// [put_EndBoundary]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_endboundary
func (me *ITrigger) PutEndBoundary(end string) error {
	return com_callBstrSet(me, end,
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_EndBoundary)
}

// [put_ExecutionTimeLimit] method.
//
// [put_ExecutionTimeLimit]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_executiontimelimit
func (me *ITrigger) PutExecutionTimeLimit(timeLimit string) error {
	return com_callBstrSet(me, timeLimit,
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_ExecutionTimeLimit)
}

// [put_Id] method.
//
// [put_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_id
func (me *ITrigger) PutId(id string) error {
	return com_callBstrSet(me, id,
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_Id)
}

// [put_StartBoundary] method.
//
// [put_StartBoundary]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_startboundary
func (me *ITrigger) PutStartBoundary(start string) error {
	return com_callBstrSet(me, start,
		(*_ITriggerVt)(unsafe.Pointer(*me.Ppvt())).Put_StartBoundary)
}

// [ITriggerCollection] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITriggerCollection]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itriggercollection
type ITriggerCollection struct{ IDispatch }

type _ITriggerCollectionVt struct {
	_IDispatchVt
	Get_Count    uintptr
	Get_Item     uintptr
	Get__NewEnum uintptr
	Create       uintptr
	Remove       uintptr
	Clear        uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITriggerCollection) IID() co.IID {
	return co.IID_ITriggerCollection
}

// [Clear] method.
//
// [Clear]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itriggercollection-clear
func (me *ITriggerCollection) Clear() error {
	return com_callErr(me,
		(*_ITriggerCollectionVt)(unsafe.Pointer(*me.Ppvt())).Clear)
}

// [Create] method.
//
// [Create]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itriggercollection-create
func (me *ITriggerCollection) Create(
	releaser *OleReleaser,
	triggerType co.TASK_TRIGGER2,
) (*ITrigger, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_ITriggerCollectionVt)(unsafe.Pointer(*me.Ppvt())).Create,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&triggerType)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*ITrigger](ret, ppvtQueried, releaser)
}

// Returns all [ITrigger] objects by calling [ITriggerCollection.GetCount] and
// [ITriggerCollection.GetItem].
func (me *ITriggerCollection) Enum(releaser *OleReleaser) ([]*ITrigger, error) {
	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	triggers := make([]*ITrigger, count)
	for i := 0; i < count; i++ {
		trigger, err := me.GetItem(releaser, i)
		if err != nil {
			return nil, err // stop immediately
		}
		triggers = append(triggers, trigger)
	}
	return triggers, nil
}

// [get_Count] method.
//
// [get_Count]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itriggercollection-get_count
func (me *ITriggerCollection) GetCount() (int, error) {
	var count int32
	ret, _, _ := syscall.SyscallN(
		(*_ITriggerCollectionVt)(unsafe.Pointer(*me.Ppvt())).Get_Count,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&count)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(count), nil
	} else {
		return 0, hr
	}
}

// [get_Item] method.
//
// [get_Item]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itriggercollection-get_item
func (me *ITriggerCollection) GetItem(releaser *OleReleaser, index int) (*ITrigger, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IActionCollectionVt)(unsafe.Pointer(*me.Ppvt())).Create,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(int32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*ITrigger](ret, ppvtQueried, releaser)
}
