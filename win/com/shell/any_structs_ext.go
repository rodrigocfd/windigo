//go:build windows

package shell

// COMDLG_FILTERSPEC syntactic sugar.
type FilterSpec struct {
	Name string
	Spec string
}
