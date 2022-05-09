//go:build windows

package ui

// Manages a group of native radio button controls.
type RadioGroup interface {
	implRadioGroup() // prevent public implementation

	// Exposes all the Button messages the can be handled by all radios in the group.
	//
	// Panics if called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_RadioButtonEvents

	AsCtrls() []AnyControl         // Returns all radios buttons as AnyControl.
	EnableAll(enable bool)         // Calls EnableWindow on all radio buttons.
	FocusSelected()                // Puts the focus on the currently selected radio button, if any.
	Item(index int) RadioButton    // Returns the RadioButton at the given index.
	Parent() AnyParent             // Returns the parent of the radio buttons.
	Selected() (RadioButton, bool) // Returns the currently selected RadioButton, if any.
}

//------------------------------------------------------------------------------

type _RadioGroup struct {
	radios []RadioButton
	events _RadioButtonEvents
}

// Creates a new RadioGroup, with one or more RadioButton controls. Call
// ui.RadioButtonOpts() to define the options of each RadioButton to be passed
// to the underlying CreateWindowEx().
//
// Example:
//
//		var owner ui.AnyParent // initialized somewhere
//
//		myRadios := ui.NewRadioGroup(
//			owner,
//			ui.RadioButtonOpts(
//				Text("First option").
//				Position(win.POINT{X: 10, Y: 40}).
//				WndStyles(co.WS_VISIBLE|co.WS_CHILD|co.WS_TABSTOP|co.WS_GROUP),
//			),
//			ui.RadioButtonOpts(
//				Text("Second option").
//				Position(win.POINT{X: 10, Y: 80}).
//				Select(true),
//			),
//			ui.RadioButtonOpts(
//				Text("Third option").
//				Position(win.POINT{X: 10, Y: 120}),
//			),
//		)
func NewRadioGroup(parent AnyParent, opts ...*_RadioButtonO) RadioGroup {
	if len(opts) == 0 {
		panic("A RadioGroup must have at least 1 RadioButton.")
	}

	radios := make([]RadioButton, 0, len(opts))
	for i := range opts {
		radios = append(radios, NewRadioButton(parent, opts[i]))
	}

	me := &_RadioGroup{}
	me.radios = radios
	me.events.new(&me.radios)
	return me
}

// Creates a new RadioGroup, where each RadioButton is from a dialog resource.
func NewRadioGroupDlg(
	parent AnyParent, horz HORZ, vert VERT, ctrlIds ...int) RadioGroup {

	radios := make([]RadioButton, 0, len(ctrlIds))
	for _, ctrlId := range ctrlIds {
		radios = append(radios, NewRadioButtonDlg(parent, ctrlId, horz, vert))
	}

	me := &_RadioGroup{}
	me.radios = radios
	me.events.new(&me.radios)
	return me
}

// Implements RadioGroup.
func (*_RadioGroup) implRadioGroup() {}

func (me *_RadioGroup) AsCtrls() []AnyControl {
	ctrls := make([]AnyControl, 0, len(me.radios))
	for _, rad := range me.radios {
		ctrls = append(ctrls, rad)
	}
	return ctrls
}

func (me *_RadioGroup) EnableAll(enable bool) {
	for _, radio := range me.radios {
		radio.Hwnd().EnableWindow(enable)
	}
}

func (me *_RadioGroup) FocusSelected() {
	if rad, ok := me.Selected(); ok {
		rad.Focus()
	}
}

func (me *_RadioGroup) Item(index int) RadioButton {
	return me.radios[index]
}

func (me *_RadioGroup) Parent() AnyParent {
	return me.radios[0].Parent()
}

func (me *_RadioGroup) Selected() (RadioButton, bool) {
	for _, radio := range me.radios {
		if radio.IsSelected() {
			return radio, true
		}
	}
	return nil, false
}

func (me *_RadioGroup) On() *_RadioButtonEvents {
	return &me.events
}

//------------------------------------------------------------------------------

// Button control notifications for multiple RadioButton controls.
type _RadioButtonEvents struct {
	radios *[]RadioButton
}

func (me *_RadioButtonEvents) new(radios *[]RadioButton) {
	me.radios = radios
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bn-clicked
func (me *_RadioButtonEvents) BnClicked(userFunc func(radio RadioButton)) {
	for _, radio := range *me.radios {
		curRadio := radio
		radio.On().BnClicked(func() {
			userFunc(curRadio)
		})
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bn-dblclk
func (me *_RadioButtonEvents) BnDblClk(userFunc func(radio RadioButton)) {
	for _, radio := range *me.radios {
		curRadio := radio
		radio.On().BnDblClk(func() {
			userFunc(curRadio)
		})
	}
}
