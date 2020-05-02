package winffi

type ACCEL struct {
	virt uint8
	key  uint16
	cmd  uint16
}

type CREATESTRUCT struct {
	CreateParams    uintptr
	Instance        HINSTANCE
	Menu            HMENU
	Parent          HWND
	Cy, Cx, Y, X    int32
	Style           int32
	Name, ClassName uintptr
	ExStyle         uint32
}
