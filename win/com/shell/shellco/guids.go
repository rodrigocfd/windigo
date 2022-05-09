//go:build windows

package shellco

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// Shell COM CLSIDs.
const (
	CLSID_DesktopWallpaper co.CLSID = "c2cf3110-460e-4fc1-b9d0-8a1c0c9cc4bd"
	CLSID_FileOpenDialog   co.CLSID = "dc1c5a9c-e88a-4dde-a5a1-60f82a20aef7"
	CLSID_FileSaveDialog   co.CLSID = "c0b4e2f3-ba21-4773-8dba-335ec946eb8b"
	CLSID_ShellLink        co.CLSID = "00021401-0000-0000-c000-000000000046"
	CLSID_TaskbarList      co.CLSID = "56fdf344-fd6d-11d0-958a-006097c9a090"
)

// Shell COM IIDs.
const (
	IID_IDesktopWallpaper co.IID = "b92b56a9-8b55-4e14-9a89-0199bbb6f93b"
	IID_IFileDialog       co.IID = "42f85136-db7e-439c-85f1-e4075d135fc8"
	IID_IFileOpenDialog   co.IID = "d57c7288-d4ad-4768-be02-9d969532d960"
	IID_IFileSaveDialog   co.IID = "84bccd23-5fde-4cdb-aea4-af64b83d78ab"
	IID_IModalWindow      co.IID = "b4db1657-70d7-485e-8e3e-6fcb5a5c1802"
	IID_IShellItem        co.IID = "43826d1e-e718-42ee-bc55-a1e261c37bfe"
	IID_IShellItemArray   co.IID = "b63ea76d-1f85-456f-a19c-48159efa858b"
	IID_IShellLink        co.IID = "000214f9-0000-0000-c000-000000000046"
	IID_ITaskbarList      co.IID = "56fdf342-fd6d-11d0-958a-006097c9a090"
	IID_ITaskbarList2     co.IID = "602d4995-b13a-429b-a66e-1935e44f4317"
	IID_ITaskbarList3     co.IID = "ea1afb91-9e28-4b86-90e9-9e9f8a5eefaf"
	IID_ITaskbarList4     co.IID = "c43dc798-95d1-4bea-9030-bb99e2983a1a"
)
