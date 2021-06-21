package ui

// Manages a group of native radio button controls.
type RadioGroup interface {
	// Exposes all the Button messages the can be handled by all radios in the group.
	// Cannot be called after the control was created.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
	On() *_RadioButtonEvents

	AsCtrls() []AnyControl         // Returns all radios buttons as AnyControl.
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
// RadioButtonOpts() to define the options of each RadioButton to be passed to
// the underlying CreateWindowEx().
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
func NewRadioGroupDlg(parent AnyParent, ctrlIds ...int) RadioGroup {
	radios := make([]RadioButton, 0, len(ctrlIds))
	for _, ctrlId := range ctrlIds {
		radios = append(radios, NewRadioButtonDlg(parent, ctrlId))
	}

	me := &_RadioGroup{}
	me.radios = radios
	me.events.new(&me.radios)
	return me
}

func (me *_RadioGroup) AsCtrls() []AnyControl {
	ctrls := make([]AnyControl, 0, len(me.radios))
	for _, rad := range me.radios {
		ctrls = append(ctrls, rad)
	}
	return ctrls
}

func (me *_RadioGroup) Item(index int) RadioButton {
	return me.radios[index]
}

func (me *_RadioGroup) Parent() AnyParent {
	return me.radios[0].Parent()
}

func (me *_RadioGroup) Selected() (RadioButton, bool) {
	for _, radio := range me.radios {
		if radio.IsChecked() {
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
