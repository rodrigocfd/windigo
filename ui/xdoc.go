//go:build windows

// This package contains high-level abstractions for GUI windows and controls.
// They are built on top of win and co packages, and attempt to provide a more
// ergonomic way to build GUI applications.
//
// The windows themselves can be built programmatically, or by loading dialog
// resources, which can be manipulated with a WYSIWYG editor like
// [Visual Studio] or [Resource Hacker].
//
// The following windows and controls are implemented:
//
//   - [Main] – main application window
//   - [Modal] – modal window
//   - [Control] – custom control
//   - [Button]
//   - [CheckBox]
//   - [ComboBox]
//   - [DateTimePicker]
//   - [Edit]
//   - [Header]
//   - [ListView]
//   - [MonthCalendar]
//   - [ProgressBar]
//   - [RadioGroup] and [RadioButton]
//   - [Static]
//   - [StatusBar]
//   - [SysLink]
//   - [Tab]
//   - [Toolbar]
//   - [Trackbar]
//   - [TreeView]
//   - [UpDown]
//
// The following interfaces are declared:
//
//   - [Window] – any window
//   - [ChildControl] – a child control window
//   - [Parent] – a parent window
//
// [Visual Studio]: https://visualstudio.microsoft.com/vs
// [Resource Hacker]: https://www.angusj.com/resourcehacker
package ui
