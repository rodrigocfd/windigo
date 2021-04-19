package shell

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// Shell COM CLSIDs.
var CLSID = struct {
	FileOpenDialog co.CLSID
	FileSaveDialog co.CLSID
	TaskbarList    co.CLSID
}{
	FileOpenDialog: "dc1c5a9c-e88a-4dde-a5a1-60f82a20aef7",
	FileSaveDialog: "c0b4e2f3-ba21-4773-8dba-335ec946eb8b",
	TaskbarList:    "56fdf344-fd6d-11d0-958a-006097c9a090",
}

// Shell COM IIDs.
var IID = struct {
	IFileDialog     co.IID
	IFileOpenDialog co.IID
	IFileSaveDialog co.IID
	IModalWindow    co.IID
	IShellItem      co.IID
	IShellItemArray co.IID
	ITaskbarList    co.IID
	ITaskbarList2   co.IID
	ITaskbarList3   co.IID
}{
	IFileDialog:     "42f85136-db7e-439c-85f1-e4075d135fc8",
	IFileOpenDialog: "d57c7288-d4ad-4768-be02-9d969532d960",
	IFileSaveDialog: "84bccd23-5fde-4cdb-aea4-af64b83d78ab",
	IModalWindow:    "b4db1657-70d7-485e-8e3e-6fcb5a5c1802",
	IShellItem:      "43826d1e-e718-42ee-bc55-a1e261c37bfe",
	IShellItemArray: "b63ea76d-1f85-456f-a19c-48159efa858b",
	ITaskbarList:    "56fdf342-fd6d-11d0-958a-006097c9a090",
	ITaskbarList2:   "602d4995-b13a-429b-a66e-1935e44f4317",
	ITaskbarList3:   "ea1afb91-9e28-4b86-90e9-9e9f8a5eefaf",
}
