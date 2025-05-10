//go:build windows

package vt

type IFileDialog struct {
	IModalWindow
	SetFileTypes        uintptr
	SetFileTypeIndex    uintptr
	GetFileTypeIndex    uintptr
	Advise              uintptr
	Unadvise            uintptr
	SetOptions          uintptr
	GetOptions          uintptr
	SetDefaultFolder    uintptr
	SetFolder           uintptr
	GetFolder           uintptr
	GetCurrentSelection uintptr
	SetFileName         uintptr
	GetFileName         uintptr
	SetTitle            uintptr
	SetOkButtonLabel    uintptr
	SetFileNameLabel    uintptr
	GetResult           uintptr
	AddPlace            uintptr
	SetDefaultExtension uintptr
	Close               uintptr
	SetClientGuid       uintptr
	ClearClientData     uintptr
	SetFilter           uintptr
}

type IFileDialogEvents struct {
	IUnknown
	OnFileOk          uintptr
	OnFolderChanging  uintptr
	OnFolderChange    uintptr
	OnSelectionChange uintptr
	OnShareViolation  uintptr
	OnTypeChange      uintptr
	OnOverwrite       uintptr
}

type IFileOpenDialog struct {
	IFileDialog
	GetResults       uintptr
	GetSelectedItems uintptr
}

type IFileSaveDialog struct {
	IFileDialog
	SetSaveAsItem          uintptr
	SetProperties          uintptr
	SetCollectedProperties uintptr
	GetProperties          uintptr
	ApplyProperties        uintptr
}

type IModalWindow struct {
	IUnknown
	Show uintptr
}

type IShellItem struct {
	IUnknown
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr
	GetAttributes  uintptr
	Compare        uintptr
}

type IShellItem2 struct {
	IShellItem
	GetPropertyStore                 uintptr
	GetPropertyStoreWithCreateObject uintptr
	GetPropertyStoreForKeys          uintptr
	GetPropertyDescriptionList       uintptr
	Update                           uintptr
	GetProperty                      uintptr
	GetCLSID                         uintptr
	GetFileTime                      uintptr
	GetInt32                         uintptr
	GetString                        uintptr
	GetUInt32                        uintptr
	GetUInt64                        uintptr
	GetBool                          uintptr
}

type IShellItemArray struct {
	IUnknown
	BindToHandler              uintptr
	GetPropertyStore           uintptr
	GetPropertyDescriptionList uintptr
	GetAttributes              uintptr
	GetCount                   uintptr
	GetItemAt                  uintptr
	EnumItems                  uintptr
}

type IShellItemFilter struct {
	IUnknown
	IncludeItem         uintptr
	GetEnumFlagsForItem uintptr
}

type IShellLink struct {
	IUnknown
	GetPath             uintptr
	GetIDList           uintptr
	SetIDList           uintptr
	GetDescription      uintptr
	SetDescription      uintptr
	GetWorkingDirectory uintptr
	SetWorkingDirectory uintptr
	GetArguments        uintptr
	SetArguments        uintptr
	GetHotkey           uintptr
	SetHotkey           uintptr
	GetShowCmd          uintptr
	SetShowCmd          uintptr
	GetIconLocation     uintptr
	SetIconLocation     uintptr
	SetRelativePath     uintptr
	Resolve             uintptr
	SetPath             uintptr
}

type ITaskbarList struct {
	IUnknown
	HrInit       uintptr
	AddTab       uintptr
	DeleteTab    uintptr
	ActivateTab  uintptr
	SetActiveAlt uintptr
}

type ITaskbarList2 struct {
	ITaskbarList
	MarkFullscreenWindow uintptr
}

type ITaskbarList3 struct {
	ITaskbarList2
	SetProgressValue      uintptr
	SetProgressState      uintptr
	RegisterTab           uintptr
	UnregisterTab         uintptr
	SetTabOrder           uintptr
	SetTabActive          uintptr
	ThumbBarAddButtons    uintptr
	ThumbBarUpdateButtons uintptr
	ThumbBarSetImageList  uintptr
	SetOverlayIcon        uintptr
	SetThumbnailTooltip   uintptr
	SetThumbnailClip      uintptr
}

type ITaskbarList4 struct {
	ITaskbarList3
	SetTabProperties uintptr
}
