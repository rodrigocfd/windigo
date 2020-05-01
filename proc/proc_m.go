package proc

var (
	MessageBox = dllUser32.NewProc("MessageBoxW")
)
