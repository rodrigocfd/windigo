//go:build windows

// This package contains high-level abstractions for GUI windows and controls.
// They are built on top of win and co packages, and attempt to provide a more
// ergonomic way to build GUI applications.
//
// The windows themselves can be built programmatically, or by loading dialog
// resources, which can be manipulated with a WYSIWYG editor like
// [Visual Studio] or [Resource Hacker].
//
// [Visual Studio]: https://visualstudio.microsoft.com/vs
// [Resource Hacker]: https://www.angusj.com/resourcehacker
package ui
