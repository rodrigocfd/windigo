# Windigo / co

This package contains the types and definitions of Win32 constants.

Notice that each constant has a type, which is used as the prefix for its values. For example, the type:

    type MB uint32

will have all its constants prefixed with `MB`:

    const (
        MB_OK          MB = 0x0000_0000
        MB_OKCANCEL    MB = 0x0000_0001
        MB_YESNOCANCEL MB = 0x0000_0003
    )

The names try to follow the Win32 name as much as possible. However, the Win32 API has many name clashes, and in such cases, the types will have slightly different names. These cases are documented in the constants themselves.

The files are named after the native Win32 library where they come from.
