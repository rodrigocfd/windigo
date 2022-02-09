# Windigo / win

This package contains the bindings to the Win32 API functions and structs.

Most files are named following this pattern:

    windowslib_subject.go

Where:

* `windowslib` is the native Win32 library where the bindings are coming from, like user32.dll or uxtheme.dll libraries;

* `subject` can be a handle, like `HWND`, or a group of features, like functions or structs.

The names of structs and functions are the same found in the Win32 API.
