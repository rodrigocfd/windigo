//go:build windows

package proc

var (
	GetWindowLongPtr = user32.NewProc("GetWindowLongW")
	SetWindowLongPtr = user32.NewProc("SetWindowLongW")
)
