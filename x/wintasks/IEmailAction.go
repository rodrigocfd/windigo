//go:build windows

package wintasks

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cotasks"
)

// [IEmailAction] COM interface.
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
func (*IEmailAction) IID() *co.IID {
	return &cotasks.IID_IEmailAction
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IEmailAction) AddRef(releaser *win.OleReleaser) *IEmailAction {
	return utl.OleNewFromAddRef[*IEmailAction](me, releaser)
}

// [get_Bcc] method.
//
// [get_Bcc]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_bcc
func (me *IEmailAction) GetBcc() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEmailActionVt](me.Ppvt()).Get_Bcc)
}

// [get_Body] method.
//
// [get_Body]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_body
func (me *IEmailAction) GetBody() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEmailActionVt](me.Ppvt()).Get_Body)
}

// [get_Cc] method.
//
// [get_Cc]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_cc
func (me *IEmailAction) GetCc() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEmailActionVt](me.Ppvt()).Get_Cc)
}

// [get_From] method.
//
// [get_From]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_from
func (me *IEmailAction) GetFrom() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEmailActionVt](me.Ppvt()).Get_From)
}

// [get_ReplyTo] method.
//
// [get_ReplyTo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_replyto
func (me *IEmailAction) GetReplyTo() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEmailActionVt](me.Ppvt()).Get_ReplyTo)
}

// [get_Server] method.
//
// [get_Server]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_server
func (me *IEmailAction) GetServer() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEmailActionVt](me.Ppvt()).Get_Server)
}

// [get_Subject] method.
//
// [get_Subject]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_subject
func (me *IEmailAction) GetSubject() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEmailActionVt](me.Ppvt()).Get_Subject)
}

// [get_To] method.
//
// [get_To]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-get_to
func (me *IEmailAction) GetTo() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEmailActionVt](me.Ppvt()).Get_To)
}

// [put_Bcc] method.
//
// [put_Bcc]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_bcc
func (me *IEmailAction) PutBcc(bcc string) error {
	return oleCallSetBstr(me, bcc, utl.Vt[_IEmailActionVt](me.Ppvt()).Put_Bcc)
}

// [put_Body] method.
//
// [put_Body]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_body
func (me *IEmailAction) PutBody(body string) error {
	return oleCallSetBstr(me, body, utl.Vt[_IEmailActionVt](me.Ppvt()).Put_Body)
}

// [put_Cc] method.
//
// [put_Cc]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_cc
func (me *IEmailAction) PutCc(cc string) error {
	return oleCallSetBstr(me, cc, utl.Vt[_IEmailActionVt](me.Ppvt()).Put_Cc)
}

// [put_From] method.
//
// [put_From]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_from
func (me *IEmailAction) PutFrom(from string) error {
	return oleCallSetBstr(me, from, utl.Vt[_IEmailActionVt](me.Ppvt()).Put_From)
}

// [put_ReplyTo] method.
//
// [put_ReplyTo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_replyto
func (me *IEmailAction) PutReplyTo(replyTo string) error {
	return oleCallSetBstr(me, replyTo, utl.Vt[_IEmailActionVt](me.Ppvt()).Put_ReplyTo)
}

// [put_Server] method.
//
// [put_Server]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_server
func (me *IEmailAction) PutServer(server string) error {
	return oleCallSetBstr(me, server, utl.Vt[_IEmailActionVt](me.Ppvt()).Put_Server)
}

// [put_Subject] method.
//
// [put_Subject]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_subject
func (me *IEmailAction) PutSubject(subject string) error {
	return oleCallSetBstr(me, subject, utl.Vt[_IEmailActionVt](me.Ppvt()).Put_Subject)
}

// [put_To] method.
//
// [put_To]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iemailaction-put_to
func (me *IEmailAction) PutTo(to string) error {
	return oleCallSetBstr(me, to, utl.Vt[_IEmailActionVt](me.Ppvt()).Put_To)
}
