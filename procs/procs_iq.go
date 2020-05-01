package procs

var (
	MessageBox = dllUser32.NewProc("MessageBoxW")
)
