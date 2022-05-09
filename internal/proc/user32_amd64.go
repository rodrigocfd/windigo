//go:build windows

package proc

var (
	GetWindowLongPtr = user32.NewProc("GetWindowLongPtrW")
	SetWindowLongPtr = user32.NewProc("SetWindowLongPtrW")
)
