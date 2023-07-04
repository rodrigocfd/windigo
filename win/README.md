# Windigo / win

This package contains the bindings to the Win32 API functions and structs.

Most files are named following this pattern:

    subject_windowslib.go

Where:

* `subject` can be a handle, like `HWND`, or a group of features, like functions or structs;

* `windowslib` is the native Win32 library where the bindings are coming from, like user32.dll or uxtheme.dll libraries.

The names of structs and functions are the same found in the Win32 API.
