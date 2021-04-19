package ui

import (
	"github.com/rodrigocfd/windigo/win"
)

// Any window.
type AnyWindow interface {
	// Returns the underlying HWND handle of this window.
	//
	// Note that this handle is initially zero, existing only after window creation.
	Hwnd() win.HWND
}

// Any window that can have child controls.
type AnyParent interface {
	AnyWindow
	internalOn() *_EventsInternal
	isDialog() bool

	// Exposes all the window notifications the can be handled.
	//
	// Cannot be called after the window was created.
	On() *_EventsNfy

	// Runs a closure synchronously in the window original UI thread.
	//
	// When in a goroutine, you *MUST* use this method to update the UI,
	// otherwise your application may deadlock.
	RunUiThread(userFunc func())
}

// Any child window control.
type AnyControl interface {
	AnyWindow

	// Returns the ID of this control.
	CtrlId() int

	// Returns the parent of this control.
	Parent() AnyParent
}

// Any native child window control.
type AnyNativeControl interface {
	AnyControl

	// Exposes all the window messages that can be handled with subclassing.
	//
	// Warning: Subclassing is a potentially slow technique, try to use the
	// regular events first.
	//
	// Cannot be called after the control was created.
	OnSubclass() *_EventsWm
}

// User-custom main application window.
type WindowMain interface {
	AnyParent

	// Creates the main window and runs the main application loop.
	//
	// Will block until the window is closed.
	RunAsMain() int
}

// User-custom modal window.
type WindowModal interface {
	AnyParent

	// Creates and shows the modal window.
	//
	// Will block until the window is closed.
	ShowModal(parent AnyParent)
}

// User-custom child control.
type WindowControl interface {
	AnyParent

	// Returns the ID of this control.
	CtrlId() int

	// Returns the parent of this control.
	Parent() AnyParent
}
