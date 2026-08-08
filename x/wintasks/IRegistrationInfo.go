//go:build windows

package wintasks

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/cotasks"
	"github.com/rodrigocfd/windigo/x/winaut"
)

// [IRegistrationInfo] COM interface.
//
// Implements [OleResource].
//
// [IRegistrationInfo]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iregistrationinfo
type IRegistrationInfo struct{ winaut.IDispatch }

type _IRegistrationInfoVt struct {
	utl.IDispatchVt
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
func (*IRegistrationInfo) IID() *co.IID {
	return &cotasks.IID_IRegistrationInfo
}

// [get_Author] method.
//
// [get_Author]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_author
func (me *IRegistrationInfo) GetAuthor() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Get_Author)
}

// [get_Date] method.
//
// [get_Date]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_date
func (me *IRegistrationInfo) GetDate() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Get_Date)
}

// [get_Description] method.
//
// [get_Description]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_description
func (me *IRegistrationInfo) GetDescription() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Get_Description)
}

// [get_Documentation] method.
//
// [get_Documentation]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_documentation
func (me *IRegistrationInfo) GetDocumentation() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Get_Documentation)
}

// [get_Source] method.
//
// [get_Source]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_source
func (me *IRegistrationInfo) GetSource() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Get_Source)
}

// [get_URI] method.
//
// [get_URI]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_uri
func (me *IRegistrationInfo) GetURI() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Get_URI)
}

// [get_Version] method.
//
// [get_Version]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_version
func (me *IRegistrationInfo) GetVersion() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Get_Version)
}

// [get_XmlText] method.
//
// [get_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-get_xmltext
func (me *IRegistrationInfo) GetXmlText() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Get_XmlText)
}

// [put_Author] method.
//
// [put_Author]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_author
func (me *IRegistrationInfo) PutAuthor(author string) error {
	return oleCallSetBstr(me, author, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Put_Author)
}

// [put_Date] method.
//
// [put_Date]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_date
func (me *IRegistrationInfo) PutDate(date string) error {
	return oleCallSetBstr(me, date, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Put_Date)
}

// [put_Description] method.
//
// [put_Description]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_description
func (me *IRegistrationInfo) PutDescription(description string) error {
	return oleCallSetBstr(me, description, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Put_Description)
}

// [put_Documentation] method.
//
// [put_Documentation]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_documentation
func (me *IRegistrationInfo) PutDocumentation(documentation string) error {
	return oleCallSetBstr(me, documentation, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Put_Documentation)
}

// [put_Source] method.
//
// [put_Source]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_source
func (me *IRegistrationInfo) PutSource(source string) error {
	return oleCallSetBstr(me, source, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Put_Source)
}

// [put_URI] method.
//
// [put_URI]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_uri
func (me *IRegistrationInfo) PutURI(uri string) error {
	return oleCallSetBstr(me, uri, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Put_URI)
}

// [put_Version] method.
//
// [put_Version]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_version
func (me *IRegistrationInfo) PutVersion(version string) error {
	return oleCallSetBstr(me, version, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Put_Version)
}

// [put_XmlText] method.
//
// [put_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregistrationinfo-put_xmltext
func (me *IRegistrationInfo) PutXmlText(xmlText string) error {
	return oleCallSetBstr(me, xmlText, utl.Vt[_IRegistrationInfoVt](me.Ppvt()).Put_XmlText)
}
