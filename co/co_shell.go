//go:build windows

package co

// Shell CLSID identifier.
var (
	CLSID_FileOpenDialog = CLSID(GUID{0xdc1c5a9c, 0xe88a, 0x4dde, [8]byte{0xa5, 0xa1, 0x60, 0xf8, 0x2a, 0x20, 0xae, 0xf7}})
	CLSID_FileOperation  = CLSID(GUID{0x3ad05575, 0x8857, 0x4850, [8]byte{0x92, 0x77, 0x11, 0xb8, 0x5b, 0xdb, 0x8e, 0x09}})
	CLSID_FileSaveDialog = CLSID(GUID{0xc0b4e2f3, 0xba21, 0x4773, [8]byte{0x8d, 0xba, 0x33, 0x5e, 0xc9, 0x46, 0xeb, 0x8b}})
	CLSID_ShellLink      = CLSID(GUID{0x00021401, 0x0000, 0x0000, [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}})
	CLSID_TaskbarList    = CLSID(GUID{0x56fdf344, 0xfd6d, 0x11d0, [8]byte{0x95, 0x8a, 0x00, 0x60, 0x97, 0xc9, 0xa0, 0x90}})
)

// [FDAP] enumeration.
//
// [FDAP]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-fdap
type FDAP uint32

const (
	FDAP_BOTTOM FDAP = 0
	FDAP_TOP    FDAP = 1
)

// [FDE_OVERWRITE_RESPONSE] enumeration.
//
// [FDE_OVERWRITE_RESPONSE]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-fde_overwrite_response
type FDEOR uint32

const (
	FDEOR_DEFAULT FDEOR = iota
	FDEOR_ACCEPT
	FDEOR_REFUSE
)

// [FDE_SHAREVIOLATION_RESPONSE] enumeration.
//
// [FDE_SHAREVIOLATION_RESPONSE]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-fde_shareviolation_response
type FDESVR uint32

const (
	FDESVR_DEFAULT FDESVR = iota
	FDESVR_ACCEPT
	FDESVR_REFUSE
)

// [SHFILEOPSTRUCT] flags.
//
// [SHFILEOPSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileopstructw
type FOF uint16

const (
	FOF_MULTIDESTFILES        FOF = 0x0001
	FOF_CONFIRMMOUSE          FOF = 0x0002
	FOF_SILENT                FOF = 0x0004
	FOF_RENAMEONCOLLISION     FOF = 0x0008
	FOF_NOCONFIRMATION        FOF = 0x0010
	FOF_WANTMAPPINGHANDLE     FOF = 0x0020
	FOF_ALLOWUNDO             FOF = 0x0040
	FOF_FILESONLY             FOF = 0x0080
	FOF_SIMPLEPROGRESS        FOF = 0x0100
	FOF_NOCONFIRMMKDIR        FOF = 0x0200
	FOF_NOERRORUI             FOF = 0x0400
	FOF_NOCOPYSECURITYATTRIBS FOF = 0x0800
	FOF_NORECURSION           FOF = 0x1000
	FOF_NO_CONNECTED_ELEMENTS FOF = 0x2000
	FOF_WANTNUKEWARNING       FOF = 0x4000
	FOF_NORECURSEREPARSE      FOF = 0x8000
	FOF_NO_UI                     = FOF_SILENT | FOF_NOCONFIRMATION | FOF_NOERRORUI | FOF_NOCONFIRMMKDIR
)

// [_FILEOPENDIALOGOPTIONS] enumeration.
//
// [_FILEOPENDIALOGOPTIONS]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-_fileopendialogoptions
type FOS uint32

const (
	FOS_OVERWRITEPROMPT          FOS = 0x2
	FOS_STRICTFILETYPES          FOS = 0x4
	FOS_NOCHANGEDIR              FOS = 0x8
	FOS_PICKFOLDERS              FOS = 0x20
	FOS_FORCEFILESYSTEM          FOS = 0x40
	FOS_ALLNONSTORAGEITEMS       FOS = 0x80
	FOS_NOVALIDATE               FOS = 0x100
	FOS_ALLOWMULTISELECT         FOS = 0x200
	FOS_PATHMUSTEXIST            FOS = 0x800
	FOS_FILEMUSTEXIST            FOS = 0x1000
	FOS_CREATEPROMPT             FOS = 0x2000
	FOS_SHAREAWARE               FOS = 0x4000
	FOS_NOREADONLYRETURN         FOS = 0x8000
	FOS_NOTESTFILECREATE         FOS = 0x1_0000
	FOS_HIDEMRUPLACES            FOS = 0x2_0000
	FOS_HIDEPINNEDPLACES         FOS = 0x4_0000
	FOS_NODEREFERENCELINKS       FOS = 0x10_0000
	FOS_OKBUTTONNEEDSINTERACTION FOS = 0x20_0000
	FOS_DONTADDTORECENT          FOS = 0x200_0000
	FOS_FORCESHOWHIDDEN          FOS = 0x1000_0000
	FOS_DEFAULTNOMINIMODE        FOS = 0x2000_0000
	FOS_FORCEPREVIEWPANEON       FOS = 0x4000_0000
	FOS_SUPPORTSTREAMABLEITEMS   FOS = 0x8000_0000
)

// [KNOWNFOLDERID] constants.
//
// [KNOWNFOLDERID]: https://learn.microsoft.com/en-us/windows/win32/shell/knownfolderid
type FOLDERID GUID

var (
	FOLDERID_NetworkFolder          = FOLDERID(GUID{0xd20beec4, 0x5ca8, 0x4905, [8]byte{0xae, 0x3b, 0xbf, 0x25, 0x1e, 0xa0, 0x9b, 0x53}})
	FOLDERID_ComputerFolder         = FOLDERID(GUID{0x0ac0837c, 0xbbf8, 0x452a, [8]byte{0x85, 0x0d, 0x79, 0xd0, 0x8e, 0x66, 0x7c, 0xa7}})
	FOLDERID_InternetFolder         = FOLDERID(GUID{0x4d9f7874, 0x4e0c, 0x4904, [8]byte{0x96, 0x7b, 0x40, 0xb0, 0xd2, 0x0c, 0x3e, 0x4b}})
	FOLDERID_ControlPanelFolder     = FOLDERID(GUID{0x82a74aeb, 0xaeb4, 0x465c, [8]byte{0xa0, 0x14, 0xd0, 0x97, 0xee, 0x34, 0x6d, 0x63}})
	FOLDERID_PrintersFolder         = FOLDERID(GUID{0x76fc4e2d, 0xd6ad, 0x4519, [8]byte{0xa6, 0x63, 0x37, 0xbd, 0x56, 0x06, 0x81, 0x85}})
	FOLDERID_SyncManagerFolder      = FOLDERID(GUID{0x43668bf8, 0xc14e, 0x49b2, [8]byte{0x97, 0xc9, 0x74, 0x77, 0x84, 0xd7, 0x84, 0xb7}})
	FOLDERID_SyncSetupFolder        = FOLDERID(GUID{0x0f214138, 0xb1d3, 0x4a90, [8]byte{0xbb, 0xa9, 0x27, 0xcb, 0xc0, 0xc5, 0x38, 0x9a}})
	FOLDERID_ConflictFolder         = FOLDERID(GUID{0x4bfefb45, 0x347d, 0x4006, [8]byte{0xa5, 0xbe, 0xac, 0x0c, 0xb0, 0x56, 0x71, 0x92}})
	FOLDERID_SyncResultsFolder      = FOLDERID(GUID{0x289a9a43, 0xbe44, 0x4057, [8]byte{0xa4, 0x1b, 0x58, 0x7a, 0x76, 0xd7, 0xe7, 0xf9}})
	FOLDERID_RecycleBinFolder       = FOLDERID(GUID{0xb7534046, 0x3ecb, 0x4c18, [8]byte{0xbe, 0x4e, 0x64, 0xcd, 0x4c, 0xb7, 0xd6, 0xac}})
	FOLDERID_ConnectionsFolder      = FOLDERID(GUID{0x6f0cd92b, 0x2e97, 0x45d1, [8]byte{0x88, 0xff, 0xb0, 0xd1, 0x86, 0xb8, 0xde, 0xdd}})
	FOLDERID_Fonts                  = FOLDERID(GUID{0xfd228cb7, 0xae11, 0x4ae3, [8]byte{0x86, 0x4c, 0x16, 0xf3, 0x91, 0x0a, 0xb8, 0xfe}})
	FOLDERID_Desktop                = FOLDERID(GUID{0xb4bfcc3a, 0xdb2c, 0x424c, [8]byte{0xb0, 0x29, 0x7f, 0xe9, 0x9a, 0x87, 0xc6, 0x41}})
	FOLDERID_Startup                = FOLDERID(GUID{0xb97d20bb, 0xf46a, 0x4c97, [8]byte{0xba, 0x10, 0x5e, 0x36, 0x08, 0x43, 0x08, 0x54}})
	FOLDERID_Programs               = FOLDERID(GUID{0xa77f5d77, 0x2e2b, 0x44c3, [8]byte{0xa6, 0xa2, 0xab, 0xa6, 0x01, 0x05, 0x4a, 0x51}})
	FOLDERID_StartMenu              = FOLDERID(GUID{0x625b53c3, 0xab48, 0x4ec1, [8]byte{0xba, 0x1f, 0xa1, 0xef, 0x41, 0x46, 0xfc, 0x19}})
	FOLDERID_Recent                 = FOLDERID(GUID{0xae50c081, 0xebd2, 0x438a, [8]byte{0x86, 0x55, 0x8a, 0x09, 0x2e, 0x34, 0x98, 0x7a}})
	FOLDERID_SendTo                 = FOLDERID(GUID{0x8983036c, 0x27c0, 0x404b, [8]byte{0x8f, 0x08, 0x10, 0x2d, 0x10, 0xdc, 0xfd, 0x74}})
	FOLDERID_Documents              = FOLDERID(GUID{0xfdd39ad0, 0x238f, 0x46af, [8]byte{0xad, 0xb4, 0x6c, 0x85, 0x48, 0x03, 0x69, 0xc7}})
	FOLDERID_Favorites              = FOLDERID(GUID{0x1777f761, 0x68ad, 0x4d8a, [8]byte{0x87, 0xbd, 0x30, 0xb7, 0x59, 0xfa, 0x33, 0xdd}})
	FOLDERID_NetHood                = FOLDERID(GUID{0xc5abbf53, 0xe17f, 0x4121, [8]byte{0x89, 0x00, 0x86, 0x62, 0x6f, 0xc2, 0xc9, 0x73}})
	FOLDERID_PrintHood              = FOLDERID(GUID{0x9274bd8d, 0xcfd1, 0x41c3, [8]byte{0xb3, 0x5e, 0xb1, 0x3f, 0x55, 0xa7, 0x58, 0xf4}})
	FOLDERID_Templates              = FOLDERID(GUID{0xa63293e8, 0x664e, 0x48db, [8]byte{0xa0, 0x79, 0xdf, 0x75, 0x9e, 0x05, 0x09, 0xf7}})
	FOLDERID_CommonStartup          = FOLDERID(GUID{0x82a5ea35, 0xd9cd, 0x47c5, [8]byte{0x96, 0x29, 0xe1, 0x5d, 0x2f, 0x71, 0x4e, 0x6e}})
	FOLDERID_CommonPrograms         = FOLDERID(GUID{0x0139d44e, 0x6afe, 0x49f2, [8]byte{0x86, 0x90, 0x3d, 0xaf, 0xca, 0xe6, 0xff, 0xb8}})
	FOLDERID_CommonStartMenu        = FOLDERID(GUID{0xa4115719, 0xd62e, 0x491d, [8]byte{0xaa, 0x7c, 0xe7, 0x4b, 0x8b, 0xe3, 0xb0, 0x67}})
	FOLDERID_PublicDesktop          = FOLDERID(GUID{0xc4aa340d, 0xf20f, 0x4863, [8]byte{0xaf, 0xef, 0xf8, 0x7e, 0xf2, 0xe6, 0xba, 0x25}})
	FOLDERID_ProgramData            = FOLDERID(GUID{0x62ab5d82, 0xfdc1, 0x4dc3, [8]byte{0xa9, 0xdd, 0x07, 0x0d, 0x1d, 0x49, 0x5d, 0x97}})
	FOLDERID_CommonTemplates        = FOLDERID(GUID{0xb94237e7, 0x57ac, 0x4347, [8]byte{0x91, 0x51, 0xb0, 0x8c, 0x6c, 0x32, 0xd1, 0xf7}})
	FOLDERID_PublicDocuments        = FOLDERID(GUID{0xed4824af, 0xdce4, 0x45a8, [8]byte{0x81, 0xe2, 0xfc, 0x79, 0x65, 0x08, 0x36, 0x34}})
	FOLDERID_RoamingAppData         = FOLDERID(GUID{0x3eb685db, 0x65f9, 0x4cf6, [8]byte{0xa0, 0x3a, 0xe3, 0xef, 0x65, 0x72, 0x9f, 0x3d}})
	FOLDERID_LocalAppData           = FOLDERID(GUID{0xf1b32785, 0x6fba, 0x4fcf, [8]byte{0x9d, 0x55, 0x7b, 0x8e, 0x7f, 0x15, 0x70, 0x91}})
	FOLDERID_LocalAppDataLow        = FOLDERID(GUID{0xa520a1a4, 0x1780, 0x4ff6, [8]byte{0xbd, 0x18, 0x16, 0x73, 0x43, 0xc5, 0xaf, 0x16}})
	FOLDERID_InternetCache          = FOLDERID(GUID{0x352481e8, 0x33be, 0x4251, [8]byte{0xba, 0x85, 0x60, 0x07, 0xca, 0xed, 0xcf, 0x9d}})
	FOLDERID_Cookies                = FOLDERID(GUID{0x2b0f765d, 0xc0e9, 0x4171, [8]byte{0x90, 0x8e, 0x08, 0xa6, 0x11, 0xb8, 0x4f, 0xf6}})
	FOLDERID_History                = FOLDERID(GUID{0xd9dc8a3b, 0xb784, 0x432e, [8]byte{0xa7, 0x81, 0x5a, 0x11, 0x30, 0xa7, 0x59, 0x63}})
	FOLDERID_System                 = FOLDERID(GUID{0x1ac14e77, 0x02e7, 0x4e5d, [8]byte{0xb7, 0x44, 0x2e, 0xb1, 0xae, 0x51, 0x98, 0xb7}})
	FOLDERID_SystemX86              = FOLDERID(GUID{0xd65231b0, 0xb2f1, 0x4857, [8]byte{0xa4, 0xce, 0xa8, 0xe7, 0xc6, 0xea, 0x7d, 0x27}})
	FOLDERID_Windows                = FOLDERID(GUID{0xf38bf404, 0x1d43, 0x42f2, [8]byte{0x93, 0x05, 0x67, 0xde, 0x0b, 0x28, 0xfc, 0x23}})
	FOLDERID_Profile                = FOLDERID(GUID{0x5e6c858f, 0x0e22, 0x4760, [8]byte{0x9a, 0xfe, 0xea, 0x33, 0x17, 0xb6, 0x71, 0x73}})
	FOLDERID_Pictures               = FOLDERID(GUID{0x33e28130, 0x4e1e, 0x4676, [8]byte{0x83, 0x5a, 0x98, 0x39, 0x5c, 0x3b, 0xc3, 0xbb}})
	FOLDERID_ProgramFilesX86        = FOLDERID(GUID{0x7c5a40ef, 0xa0fb, 0x4bfc, [8]byte{0x87, 0x4a, 0xc0, 0xf2, 0xe0, 0xb9, 0xfa, 0x8e}})
	FOLDERID_ProgramFilesCommonX86  = FOLDERID(GUID{0xde974d24, 0xd9c6, 0x4d3e, [8]byte{0xbf, 0x91, 0xf4, 0x45, 0x51, 0x20, 0xb9, 0x17}})
	FOLDERID_ProgramFilesX64        = FOLDERID(GUID{0x6d809377, 0x6af0, 0x444b, [8]byte{0x89, 0x57, 0xa3, 0x77, 0x3f, 0x02, 0x20, 0x0e}})
	FOLDERID_ProgramFilesCommonX64  = FOLDERID(GUID{0x6365d5a7, 0x0f0d, 0x45e5, [8]byte{0x87, 0xf6, 0x0d, 0xa5, 0x6b, 0x6a, 0x4f, 0x7d}})
	FOLDERID_ProgramFiles           = FOLDERID(GUID{0x905e63b6, 0xc1bf, 0x494e, [8]byte{0xb2, 0x9c, 0x65, 0xb7, 0x32, 0xd3, 0xd2, 0x1a}})
	FOLDERID_ProgramFilesCommon     = FOLDERID(GUID{0xf7f1ed05, 0x9f6d, 0x47a2, [8]byte{0xaa, 0xae, 0x29, 0xd3, 0x17, 0xc6, 0xf0, 0x66}})
	FOLDERID_UserProgramFiles       = FOLDERID(GUID{0x5cd7aee2, 0x2219, 0x4a67, [8]byte{0xb8, 0x5d, 0x6c, 0x9c, 0xe1, 0x56, 0x60, 0xcb}})
	FOLDERID_UserProgramFilesCommon = FOLDERID(GUID{0xbcbd3057, 0xca5c, 0x4622, [8]byte{0xb4, 0x2d, 0xbc, 0x56, 0xdb, 0x0a, 0xe5, 0x16}})
	FOLDERID_AdminTools             = FOLDERID(GUID{0x724ef170, 0xa42d, 0x4fef, [8]byte{0x9f, 0x26, 0xb6, 0x0e, 0x84, 0x6f, 0xba, 0x4f}})
	FOLDERID_CommonAdminTools       = FOLDERID(GUID{0xd0384e7d, 0xbac3, 0x4797, [8]byte{0x8f, 0x14, 0xcb, 0xa2, 0x29, 0xb3, 0x92, 0xb5}})
	FOLDERID_Music                  = FOLDERID(GUID{0x4bd8d571, 0x6d19, 0x48d3, [8]byte{0xbe, 0x97, 0x42, 0x22, 0x20, 0x08, 0x0e, 0x43}})
	FOLDERID_Videos                 = FOLDERID(GUID{0x18989b1d, 0x99b5, 0x455b, [8]byte{0x84, 0x1c, 0xab, 0x7c, 0x74, 0xe4, 0xdd, 0xfc}})
	FOLDERID_Ringtones              = FOLDERID(GUID{0xc870044b, 0xf49e, 0x4126, [8]byte{0xa9, 0xc3, 0xb5, 0x2a, 0x1f, 0xf4, 0x11, 0xe8}})
	FOLDERID_PublicPictures         = FOLDERID(GUID{0xb6ebfb86, 0x6907, 0x413c, [8]byte{0x9a, 0xf7, 0x4f, 0xc2, 0xab, 0xf0, 0x7c, 0xc5}})
	FOLDERID_PublicMusic            = FOLDERID(GUID{0x3214fab5, 0x9757, 0x4298, [8]byte{0xbb, 0x61, 0x92, 0xa9, 0xde, 0xaa, 0x44, 0xff}})
	FOLDERID_PublicVideos           = FOLDERID(GUID{0x2400183a, 0x6185, 0x49fb, [8]byte{0xa2, 0xd8, 0x4a, 0x39, 0x2a, 0x60, 0x2b, 0xa3}})
	FOLDERID_PublicRingtones        = FOLDERID(GUID{0xe555ab60, 0x153b, 0x4d17, [8]byte{0x9f, 0x04, 0xa5, 0xfe, 0x99, 0xfc, 0x15, 0xec}})
	FOLDERID_ResourceDir            = FOLDERID(GUID{0x8ad10c31, 0x2adb, 0x4296, [8]byte{0xa8, 0xf7, 0xe4, 0x70, 0x12, 0x32, 0xc9, 0x72}})
	FOLDERID_LocalizedResourcesDir  = FOLDERID(GUID{0x2a00375e, 0x224c, 0x49de, [8]byte{0xb8, 0xd1, 0x44, 0x0d, 0xf7, 0xef, 0x3d, 0xdc}})
	FOLDERID_CommonOEMLinks         = FOLDERID(GUID{0xc1bae2d0, 0x10df, 0x4334, [8]byte{0xbe, 0xdd, 0x7a, 0xa2, 0x0b, 0x22, 0x7a, 0x9d}})
	FOLDERID_CDBurning              = FOLDERID(GUID{0x9e52ab10, 0xf80d, 0x49df, [8]byte{0xac, 0xb8, 0x43, 0x30, 0xf5, 0x68, 0x78, 0x55}})
	FOLDERID_UserProfiles           = FOLDERID(GUID{0x0762d272, 0xc50a, 0x4bb0, [8]byte{0xa3, 0x82, 0x69, 0x7d, 0xcd, 0x72, 0x9b, 0x80}})
	FOLDERID_Playlists              = FOLDERID(GUID{0xde92c1c7, 0x837f, 0x4f69, [8]byte{0xa3, 0xbb, 0x86, 0xe6, 0x31, 0x20, 0x4a, 0x23}})
	FOLDERID_SamplePlaylists        = FOLDERID(GUID{0x15ca69b3, 0x30ee, 0x49c1, [8]byte{0xac, 0xe1, 0x6b, 0x5e, 0xc3, 0x72, 0xaf, 0xb5}})
	FOLDERID_SampleMusic            = FOLDERID(GUID{0xb250c668, 0xf57d, 0x4ee1, [8]byte{0xa6, 0x3c, 0x29, 0x0e, 0xe7, 0xd1, 0xaa, 0x1f}})
	FOLDERID_SamplePictures         = FOLDERID(GUID{0xc4900540, 0x2379, 0x4c75, [8]byte{0x84, 0x4b, 0x64, 0xe6, 0xfa, 0xf8, 0x71, 0x6b}})
	FOLDERID_SampleVideos           = FOLDERID(GUID{0x859ead94, 0x2e85, 0x48ad, [8]byte{0xa7, 0x1a, 0x09, 0x69, 0xcb, 0x56, 0xa6, 0xcd}})
	FOLDERID_PhotoAlbums            = FOLDERID(GUID{0x69d2cf90, 0xfc33, 0x4fb7, [8]byte{0x9a, 0x0c, 0xeb, 0xb0, 0xf0, 0xfc, 0xb4, 0x3c}})
	FOLDERID_Public                 = FOLDERID(GUID{0xdfdf76a2, 0xc82a, 0x4d63, [8]byte{0x90, 0x6a, 0x56, 0x44, 0xac, 0x45, 0x73, 0x85}})
	FOLDERID_ChangeRemovePrograms   = FOLDERID(GUID{0xdf7266ac, 0x9274, 0x4867, [8]byte{0x8d, 0x55, 0x3b, 0xd6, 0x61, 0xde, 0x87, 0x2d}})
	FOLDERID_AppUpdates             = FOLDERID(GUID{0xa305ce99, 0xf527, 0x492b, [8]byte{0x8b, 0x1a, 0x7e, 0x76, 0xfa, 0x98, 0xd6, 0xe4}})
	FOLDERID_AddNewPrograms         = FOLDERID(GUID{0xde61d971, 0x5ebc, 0x4f02, [8]byte{0xa3, 0xa9, 0x6c, 0x82, 0x89, 0x5e, 0x5c, 0x04}})
	FOLDERID_Downloads              = FOLDERID(GUID{0x374de290, 0x123f, 0x4565, [8]byte{0x91, 0x64, 0x39, 0xc4, 0x92, 0x5e, 0x46, 0x7b}})
	FOLDERID_PublicDownloads        = FOLDERID(GUID{0x3d644c9b, 0x1fb8, 0x4f30, [8]byte{0x9b, 0x45, 0xf6, 0x70, 0x23, 0x5f, 0x79, 0xc0}})
	FOLDERID_SavedSearches          = FOLDERID(GUID{0x7d1d3a04, 0xdebb, 0x4115, [8]byte{0x95, 0xcf, 0x2f, 0x29, 0xda, 0x29, 0x20, 0xda}})
	FOLDERID_QuickLaunch            = FOLDERID(GUID{0x52a4f021, 0x7b75, 0x48a9, [8]byte{0x9f, 0x6b, 0x4b, 0x87, 0xa2, 0x10, 0xbc, 0x8f}})
	FOLDERID_Contacts               = FOLDERID(GUID{0x56784854, 0xc6cb, 0x462b, [8]byte{0x81, 0x69, 0x88, 0xe3, 0x50, 0xac, 0xb8, 0x82}})
	FOLDERID_SidebarParts           = FOLDERID(GUID{0xa75d362e, 0x50fc, 0x4fb7, [8]byte{0xac, 0x2c, 0xa8, 0xbe, 0xaa, 0x31, 0x44, 0x93}})
	FOLDERID_SidebarDefaultParts    = FOLDERID(GUID{0x7b396e54, 0x9ec5, 0x4300, [8]byte{0xbe, 0x0a, 0x24, 0x82, 0xeb, 0xae, 0x1a, 0x26}})
	FOLDERID_PublicGameTasks        = FOLDERID(GUID{0xdebf2536, 0xe1a8, 0x4c59, [8]byte{0xb6, 0xa2, 0x41, 0x45, 0x86, 0x47, 0x6a, 0xea}})
	FOLDERID_GameTasks              = FOLDERID(GUID{0x054fae61, 0x4dd8, 0x4787, [8]byte{0x80, 0xb6, 0x09, 0x02, 0x20, 0xc4, 0xb7, 0x00}})
	FOLDERID_SavedGames             = FOLDERID(GUID{0x4c5c32ff, 0xbb9d, 0x43b0, [8]byte{0xb5, 0xb4, 0x2d, 0x72, 0xe5, 0x4e, 0xaa, 0xa4}})
	FOLDERID_Games                  = FOLDERID(GUID{0xcac52c1a, 0xb53d, 0x4edc, [8]byte{0x92, 0xd7, 0x6b, 0x2e, 0x8a, 0xc1, 0x94, 0x34}})
	FOLDERID_SEARCH_MAPI            = FOLDERID(GUID{0x98ec0e18, 0x2098, 0x4d44, [8]byte{0x86, 0x44, 0x66, 0x97, 0x93, 0x15, 0xa2, 0x81}})
	FOLDERID_SEARCH_CSC             = FOLDERID(GUID{0xee32e446, 0x31ca, 0x4aba, [8]byte{0x81, 0x4f, 0xa5, 0xeb, 0xd2, 0xfd, 0x6d, 0x5e}})
	FOLDERID_Links                  = FOLDERID(GUID{0xbfb9d5e0, 0xc6a9, 0x404c, [8]byte{0xb2, 0xb2, 0xae, 0x6d, 0xb6, 0xaf, 0x49, 0x68}})
	FOLDERID_UsersFiles             = FOLDERID(GUID{0xf3ce0f7c, 0x4901, 0x4acc, [8]byte{0x86, 0x48, 0xd5, 0xd4, 0x4b, 0x04, 0xef, 0x8f}})
	FOLDERID_UsersLibraries         = FOLDERID(GUID{0xa302545d, 0xdeff, 0x464b, [8]byte{0xab, 0xe8, 0x61, 0xc8, 0x64, 0x8d, 0x93, 0x9b}})
	FOLDERID_SearchHome             = FOLDERID(GUID{0x190337d1, 0xb8ca, 0x4121, [8]byte{0xa6, 0x39, 0x6d, 0x47, 0x2d, 0x16, 0x97, 0x2a}})
	FOLDERID_OriginalImages         = FOLDERID(GUID{0x2c36c0aa, 0x5812, 0x4b87, [8]byte{0xbf, 0xd0, 0x4c, 0xd0, 0xdf, 0xb1, 0x9b, 0x39}})
	FOLDERID_DocumentsLibrary       = FOLDERID(GUID{0x7b0db17d, 0x9cd2, 0x4a93, [8]byte{0x97, 0x33, 0x46, 0xcc, 0x89, 0x02, 0x2e, 0x7c}})
	FOLDERID_MusicLibrary           = FOLDERID(GUID{0x2112ab0a, 0xc86a, 0x4ffe, [8]byte{0xa3, 0x68, 0x0d, 0xe9, 0x6e, 0x47, 0x01, 0x2e}})
	FOLDERID_PicturesLibrary        = FOLDERID(GUID{0xa990ae9f, 0xa03b, 0x4e80, [8]byte{0x94, 0xbc, 0x99, 0x12, 0xd7, 0x50, 0x41, 0x04}})
	FOLDERID_VideosLibrary          = FOLDERID(GUID{0x491e922f, 0x5643, 0x4af4, [8]byte{0xa7, 0xeb, 0x4e, 0x7a, 0x13, 0x8d, 0x81, 0x74}})
	FOLDERID_RecordedTVLibrary      = FOLDERID(GUID{0x1a6fdba2, 0xf42d, 0x4358, [8]byte{0xa7, 0x98, 0xb7, 0x4d, 0x74, 0x59, 0x26, 0xc5}})
	FOLDERID_HomeGroup              = FOLDERID(GUID{0x52528a6b, 0xb9e3, 0x4add, [8]byte{0xb6, 0x0d, 0x58, 0x8c, 0x2d, 0xba, 0x84, 0x2d}})
	FOLDERID_HomeGroupCurrentUser   = FOLDERID(GUID{0x9b74b6a3, 0x0dfd, 0x4f11, [8]byte{0x9e, 0x78, 0x5f, 0x78, 0x00, 0xf2, 0xe7, 0x72}})
	FOLDERID_DeviceMetadataStore    = FOLDERID(GUID{0x5ce4a5e9, 0xe4eb, 0x479d, [8]byte{0xb8, 0x9f, 0x13, 0x0c, 0x02, 0x88, 0x61, 0x55}})
	FOLDERID_Libraries              = FOLDERID(GUID{0x1b3ea5dc, 0xb587, 0x4786, [8]byte{0xb4, 0xef, 0xbd, 0x1d, 0xc3, 0x32, 0xae, 0xae}})
	FOLDERID_PublicLibraries        = FOLDERID(GUID{0x48daf80b, 0xe6cf, 0x4f4e, [8]byte{0xb8, 0x00, 0x0e, 0x69, 0xd8, 0x4e, 0xe3, 0x84}})
	FOLDERID_UserPinned             = FOLDERID(GUID{0x9e3995ab, 0x1f9c, 0x4f13, [8]byte{0xb8, 0x27, 0x48, 0xb2, 0x4b, 0x6c, 0x71, 0x74}})
	FOLDERID_ImplicitAppShortcuts   = FOLDERID(GUID{0xbcb5256f, 0x79f6, 0x4cee, [8]byte{0xb7, 0x25, 0xdc, 0x34, 0xe4, 0x02, 0xfd, 0x46}})
	FOLDERID_AccountPictures        = FOLDERID(GUID{0x008ca0b1, 0x55b4, 0x4c56, [8]byte{0xb8, 0xa8, 0x4d, 0xe4, 0xb2, 0x99, 0xd3, 0xbe}})
	FOLDERID_PublicUserTiles        = FOLDERID(GUID{0x0482af6c, 0x08f1, 0x4c34, [8]byte{0x8c, 0x90, 0xe1, 0x7e, 0xc9, 0x8b, 0x1e, 0x17}})
	FOLDERID_AppsFolder             = FOLDERID(GUID{0x1e87508d, 0x89c2, 0x42f0, [8]byte{0x8a, 0x7e, 0x64, 0x5a, 0x0f, 0x50, 0xca, 0x58}})
	FOLDERID_StartMenuAllPrograms   = FOLDERID(GUID{0xf26305ef, 0x6948, 0x40b9, [8]byte{0xb2, 0x55, 0x81, 0x45, 0x3d, 0x09, 0xc7, 0x85}})
	FOLDERID_CommonStartMenuPlaces  = FOLDERID(GUID{0xa440879f, 0x87a0, 0x4f7d, [8]byte{0xb7, 0x00, 0x02, 0x07, 0xb9, 0x66, 0x19, 0x4a}})
	FOLDERID_ApplicationShortcuts   = FOLDERID(GUID{0xa3918781, 0xe5f2, 0x4890, [8]byte{0xb3, 0xd9, 0xa7, 0xe5, 0x43, 0x32, 0x32, 0x8c}})
	FOLDERID_RoamingTiles           = FOLDERID(GUID{0x00bcfc5a, 0xed94, 0x4e48, [8]byte{0x96, 0xa1, 0x3f, 0x62, 0x17, 0xf2, 0x19, 0x90}})
	FOLDERID_RoamedTileImages       = FOLDERID(GUID{0xaaa8d5a5, 0xf1d6, 0x4259, [8]byte{0xba, 0xa8, 0x78, 0xe7, 0xef, 0x60, 0x83, 0x5e}})
	FOLDERID_Screenshots            = FOLDERID(GUID{0xb7bede81, 0xdf94, 0x4682, [8]byte{0xa7, 0xd8, 0x57, 0xa5, 0x26, 0x20, 0xb8, 0x6f}})
	FOLDERID_CameraRoll             = FOLDERID(GUID{0xab5fb87b, 0x7ce2, 0x4f83, [8]byte{0x91, 0x5d, 0x55, 0x08, 0x46, 0xc9, 0x53, 0x7b}})
	FOLDERID_SkyDrive               = FOLDERID(GUID{0xa52bba46, 0xe9e1, 0x435f, [8]byte{0xb3, 0xd9, 0x28, 0xda, 0xa6, 0x48, 0xc0, 0xf6}})
	FOLDERID_OneDrive               = FOLDERID(GUID{0xa52bba46, 0xe9e1, 0x435f, [8]byte{0xb3, 0xd9, 0x28, 0xda, 0xa6, 0x48, 0xc0, 0xf6}})
	FOLDERID_SkyDriveDocuments      = FOLDERID(GUID{0x24d89e24, 0x2f19, 0x4534, [8]byte{0x9d, 0xde, 0x6a, 0x66, 0x71, 0xfb, 0xb8, 0xfe}})
	FOLDERID_SkyDrivePictures       = FOLDERID(GUID{0x339719b5, 0x8c47, 0x4894, [8]byte{0x94, 0xc2, 0xd8, 0xf7, 0x7a, 0xdd, 0x44, 0xa6}})
	FOLDERID_SkyDriveMusic          = FOLDERID(GUID{0xc3f2459e, 0x80d6, 0x45dc, [8]byte{0xbf, 0xef, 0x1f, 0x76, 0x9f, 0x2b, 0xe7, 0x30}})
	FOLDERID_SkyDriveCameraRoll     = FOLDERID(GUID{0x767e6811, 0x49cb, 0x4273, [8]byte{0x87, 0xc2, 0x20, 0xf3, 0x55, 0xe1, 0x08, 0x5b}})
	FOLDERID_SearchHistory          = FOLDERID(GUID{0x0d4c3db6, 0x03a3, 0x462f, [8]byte{0xa0, 0xe6, 0x08, 0x92, 0x4c, 0x41, 0xb5, 0xd4}})
	FOLDERID_SearchTemplates        = FOLDERID(GUID{0x7e636bfe, 0xdfa9, 0x4d5e, [8]byte{0xb4, 0x56, 0xd7, 0xb3, 0x98, 0x51, 0xd8, 0xa9}})
	FOLDERID_CameraRollLibrary      = FOLDERID(GUID{0x2b20df75, 0x1eda, 0x4039, [8]byte{0x80, 0x97, 0x38, 0x79, 0x82, 0x27, 0xd5, 0xb7}})
	FOLDERID_SavedPictures          = FOLDERID(GUID{0x3b193882, 0xd3ad, 0x4eab, [8]byte{0x96, 0x5a, 0x69, 0x82, 0x9d, 0x1f, 0xb5, 0x9f}})
	FOLDERID_SavedPicturesLibrary   = FOLDERID(GUID{0xe25b5812, 0xbe88, 0x4bd9, [8]byte{0x94, 0xb0, 0x29, 0x23, 0x34, 0x77, 0xb6, 0xc3}})
	FOLDERID_RetailDemo             = FOLDERID(GUID{0x12d4c69e, 0x24ad, 0x4923, [8]byte{0xbe, 0x19, 0x31, 0x32, 0x1c, 0x43, 0xa7, 0x67}})
	FOLDERID_Device                 = FOLDERID(GUID{0x1c2ac1dc, 0x4358, 0x4b6c, [8]byte{0x97, 0x33, 0xaf, 0x21, 0x15, 0x65, 0x76, 0xf0}})
	FOLDERID_DevelopmentFiles       = FOLDERID(GUID{0xdbe8e08e, 0x3053, 0x4bbc, [8]byte{0xb1, 0x83, 0x2a, 0x7b, 0x2b, 0x19, 0x1e, 0x59}})
	FOLDERID_Objects3D              = FOLDERID(GUID{0x31c0dd25, 0x9439, 0x4f12, [8]byte{0xbf, 0x41, 0x7f, 0xf4, 0xed, 0xa3, 0x87, 0x22}})
	FOLDERID_AppCaptures            = FOLDERID(GUID{0xedc0fe71, 0x98d8, 0x4f4a, [8]byte{0xb9, 0x20, 0xc8, 0xdc, 0x13, 0x3c, 0xb1, 0x65}})
	FOLDERID_LocalDocuments         = FOLDERID(GUID{0xf42ee2d3, 0x909f, 0x4907, [8]byte{0x88, 0x71, 0x4c, 0x22, 0xfc, 0x0b, 0xf7, 0x56}})
	FOLDERID_LocalPictures          = FOLDERID(GUID{0x0ddd015d, 0xb06c, 0x45d5, [8]byte{0x8c, 0x4c, 0xf5, 0x97, 0x13, 0x85, 0x46, 0x39}})
	FOLDERID_LocalVideos            = FOLDERID(GUID{0x35286a68, 0x3c57, 0x41a1, [8]byte{0xbb, 0xb1, 0x0e, 0xae, 0x73, 0xd7, 0x6c, 0x95}})
	FOLDERID_LocalMusic             = FOLDERID(GUID{0xa0c69a99, 0x21c8, 0x4671, [8]byte{0x87, 0x03, 0x79, 0x34, 0x16, 0x2f, 0xcf, 0x1d}})
	FOLDERID_LocalDownloads         = FOLDERID(GUID{0x7d83ee9b, 0x2244, 0x4e70, [8]byte{0xb1, 0xf5, 0x53, 0x93, 0x04, 0x2a, 0xf1, 0xe4}})
	FOLDERID_RecordedCalls          = FOLDERID(GUID{0x2f8b40c2, 0x83ed, 0x48ee, [8]byte{0xb3, 0x83, 0xa1, 0xf1, 0x57, 0xec, 0x6f, 0x9a}})
	FOLDERID_AllAppMods             = FOLDERID(GUID{0x7ad67899, 0x66af, 0x43ba, [8]byte{0x91, 0x56, 0x6a, 0xad, 0x42, 0xe6, 0xc5, 0x96}})
	FOLDERID_CurrentAppMods         = FOLDERID(GUID{0x3db40b20, 0x2a30, 0x4dbe, [8]byte{0x91, 0x7e, 0x77, 0x1d, 0xd2, 0x1d, 0xd0, 0x99}})
	FOLDERID_AppDataDesktop         = FOLDERID(GUID{0xb2c5e279, 0x7add, 0x439f, [8]byte{0xb2, 0x8c, 0xc4, 0x1f, 0xe1, 0xbb, 0xf6, 0x72}})
	FOLDERID_AppDataDocuments       = FOLDERID(GUID{0x7be16610, 0x1f7f, 0x44ac, [8]byte{0xbf, 0xf0, 0x83, 0xe1, 0x5f, 0x2f, 0xfc, 0xa1}})
	FOLDERID_AppDataFavorites       = FOLDERID(GUID{0x7cfbefbc, 0xde1f, 0x45aa, [8]byte{0xb8, 0x43, 0xa5, 0x42, 0xac, 0x53, 0x6c, 0xc9}})
	FOLDERID_AppDataProgramData     = FOLDERID(GUID{0x559d40a3, 0xa036, 0x40fa, [8]byte{0xaf, 0x61, 0x84, 0xcb, 0x43, 0x0a, 0x4d, 0x34}})
	FOLDERID_LocalStorage           = FOLDERID(GUID{0xb3eb08d3, 0xa1f3, 0x496b, [8]byte{0x86, 0x5a, 0x42, 0xb5, 0x36, 0xcd, 0xa0, 0xec}})
)

// [GETPROPERTYSTOREFLAGS] enumeration.
//
// [GETPROPERTYSTOREFLAGS]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/ne-propsys-getpropertystoreflags
type GPS uint32

const (
	GPS_DEFAULT                 GPS = 0
	GPS_HANDLERPROPERTIESONLY   GPS = 0x1
	GPS_READWRITE               GPS = 0x2
	GPS_TEMPORARY               GPS = 0x4
	GPS_FASTPROPERTIESONLY      GPS = 0x8
	GPS_OPENSLOWITEM            GPS = 0x10
	GPS_DELAYCREATION           GPS = 0x20
	GPS_BESTEFFORT              GPS = 0x40
	GPS_NO_OPLOCK               GPS = 0x80
	GPS_PREFERQUERYPROPERTIES   GPS = 0x100
	GPS_EXTRINSICPROPERTIES     GPS = 0x200
	GPS_EXTRINSICPROPERTIESONLY GPS = 0x400
	GPS_VOLATILEPROPERTIES      GPS = 0x800
	GPS_VOLATILEPROPERTIESONLY  GPS = 0x1000
	GPS_MASK_VALID              GPS = 0x1fff
)

// [IShellLink.GetHotkey] returned value.
//
// [IShellLink.GetHotkey]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-gethotkey
type HOTKEYF uint16

const (
	HOTKEYF_SHIFT   HOTKEYF = 0x01
	HOTKEYF_CONTROL HOTKEYF = 0x02
	HOTKEYF_ALT     HOTKEYF = 0x04
	HOTKEYF_EXT     HOTKEYF = 0x08
)

// Shell IID identifier.
var (
	IID_IEnumIDList                = IID(GUID{0x000214f2, 0x0000, 0x0000, [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}})
	IID_IEnumShellItems            = IID(GUID{0x70629033, 0xe363, 0x4a28, [8]byte{0xa5, 0x67, 0x0d, 0xb7, 0x80, 0x06, 0xe6, 0xd7}})
	IID_IFileDialog                = IID(GUID{0x42f85136, 0xdb7e, 0x439c, [8]byte{0x85, 0xf1, 0xe4, 0x07, 0x5d, 0x13, 0x5f, 0xc8}})
	IID_IFileDialogEvents          = IID(GUID{0x973510db, 0x7d7f, 0x452b, [8]byte{0x89, 0x75, 0x74, 0xa8, 0x58, 0x28, 0xd3, 0x54}})
	IID_IFileOpenDialog            = IID(GUID{0xd57c7288, 0xd4ad, 0x4768, [8]byte{0xbe, 0x02, 0x9d, 0x96, 0x95, 0x32, 0xd9, 0x60}})
	IID_IFileOperation             = IID(GUID{0x947aab5f, 0x0a5c, 0x4c13, [8]byte{0xb4, 0xd6, 0x4b, 0xf7, 0x83, 0x6f, 0xc9, 0xf8}})
	IID_IFileOperationProgressSink = IID(GUID{0x04b0f1a7, 0x9490, 0x44bc, [8]byte{0x96, 0xe1, 0x42, 0x96, 0xa3, 0x12, 0x52, 0xe2}})
	IID_IFileSaveDialog            = IID(GUID{0x84bccd23, 0x5fde, 0x4cdb, [8]byte{0xae, 0xa4, 0xaf, 0x64, 0xb8, 0x3d, 0x78, 0xab}})
	IID_IModalWindow               = IID(GUID{0xb4db1657, 0x70d7, 0x485e, [8]byte{0x8e, 0x3e, 0x6f, 0xcb, 0x5a, 0x5c, 0x18, 0x02}})
	IID_IOleWindow                 = IID(GUID{0x00000114, 0x0000, 0x0000, [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}})
	IID_IShellFolder               = IID(GUID{0x000214e6, 0x0000, 0x0000, [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}})
	IID_IShellItem                 = IID(GUID{0x43826d1e, 0xe718, 0x42ee, [8]byte{0xbc, 0x55, 0xa1, 0xe2, 0x61, 0xc3, 0x7b, 0xfe}})
	IID_IShellItem2                = IID(GUID{0x7e9fb0d3, 0x919f, 0x4307, [8]byte{0xab, 0x2e, 0x9b, 0x18, 0x60, 0x31, 0x0c, 0x93}})
	IID_IShellItemArray            = IID(GUID{0xb63ea76d, 0x1f85, 0x456f, [8]byte{0xa1, 0x9c, 0x48, 0x15, 0x9e, 0xfa, 0x85, 0x8b}})
	IID_IShellItemFilter           = IID(GUID{0x2659b475, 0xeeb8, 0x48b7, [8]byte{0x8f, 0x07, 0xb3, 0x78, 0x81, 0x0f, 0x48, 0xcf}})
	IID_IShellLink                 = IID(GUID{0x000214f9, 0x0000, 0x0000, [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}})
	IID_IShellView                 = IID(GUID{0x000214e3, 0x0000, 0x0000, [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}})
	IID_ITaskbarList               = IID(GUID{0x56fdf342, 0xfd6d, 0x11d0, [8]byte{0x95, 0x8a, 0x00, 0x60, 0x97, 0xc9, 0xa0, 0x90}})
	IID_ITaskbarList2              = IID(GUID{0x602d4995, 0xb13a, 0x429b, [8]byte{0xa6, 0x6e, 0x19, 0x35, 0xe4, 0x4f, 0x43, 0x17}})
	IID_ITaskbarList3              = IID(GUID{0xea1afb91, 0x9e28, 0x4b86, [8]byte{0x90, 0xe9, 0x9e, 0x9f, 0x8a, 0x5e, 0xef, 0xaf}})
	IID_ITaskbarList4              = IID(GUID{0xc43dc798, 0x95d1, 0x4bea, [8]byte{0x90, 0x30, 0xbb, 0x99, 0xe2, 0x98, 0x3a, 0x1a}})
)

// [KNOWN_FOLDER_FLAG] enumeration.
//
// [KNOWN_FOLDER_FLAG]: https://learn.microsoft.com/en-us/windows/win32/api/shlobj_core/ne-shlobj_core-known_folder_flag
type KF uint32

const (
	KF_DEFAULT                          KF = 0x0000_0000
	KF_FORCE_APP_DATA_REDIRECTION       KF = 0x0008_0000
	KF_RETURN_FILTER_REDIRECTION_TARGET KF = 0x0004_0000
	KF_FORCE_PACKAGE_REDIRECTION        KF = 0x0002_0000
	KF_NO_PACKAGE_REDIRECTION           KF = 0x0001_0000
	KF_FORCE_APPCONTAINER_REDIRECTION   KF = 0x0002_0000
	KF_NO_APPCONTAINER_REDIRECTION      KF = 0x0001_0000
	KF_CREATE                           KF = 0x0000_8000
	KF_DONT_VERIFY                      KF = 0x0000_4000
	KF_DONT_UNEXPAND                    KF = 0x0000_2000
	KF_NO_ALIAS                         KF = 0x0000_1000
	KF_INIT                             KF = 0x0000_0800
	KF_DEFAULT_PATH                     KF = 0x0000_0400
	KF_NOT_PARENT_RELATIVE              KF = 0x0000_0200
	KF_SIMPLE_IDLIST                    KF = 0x0000_0100
	KF_ALIAS_ONLY                       KF = 0x8000_0000
)

// [NOTIFYICONDATA] uFlags.
//
// [NOTIFYICONDATA]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NIF uint32

const (
	NIF_MESSAGE  NIF = 0x0000_0001 // The UCallbackMessage member is valid.
	NIF_ICON     NIF = 0x0000_0002 // The HIcon member is valid.
	NIF_TIP      NIF = 0x0000_0004 // The szTip member is valid.
	NIF_STATE    NIF = 0x0000_0008 // The DwState and DwStateMask members are valid.
	NIF_INFO     NIF = 0x0000_0010 // Display a balloon notification. The szInfo, szInfoTitle and DwInfoFlags are valid.
	NIF_GUID     NIF = 0x0000_0020 // The GuidItem member is valid.
	NIF_REALTIME NIF = 0x0000_0040 // If the balloon notification cannot be displayed immediately, discard it.
	NIF_SHOWTIP  NIF = 0x0000_0080 // Use the standard tooltip instead of an application-drawn, pop-up UI.
)

// [NOTIFYICONDATA] dwInfoFlags.
//
// [NOTIFYICONDATA]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NIIF uint32

const (
	NIIF_NONE               NIIF = 0x0000_0000
	NIIF_INFO               NIIF = 0x0000_0001
	NIIF_WARNING            NIIF = 0x0000_0002
	NIIF_ERROR              NIIF = 0x0000_0003
	NIIF_USER               NIIF = 0x0000_0004
	NIIF_NOSOUND            NIIF = 0x0000_0010
	NIIF_LARGE_ICON         NIIF = 0x0000_0020
	NIIF_RESPECT_QUIET_TIME NIIF = 0x0000_0080
	NIIF_ICON_MASK          NIIF = 0x0000_000f
)

// [Shell_NotifyIcon] dwMessage.
//
// [Shell_NotifyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
type NIM uint32

const (
	NIM_ADD        NIM = 0x0000_0000
	NIM_MODIFY     NIM = 0x0000_0001
	NIM_DELETE     NIM = 0x0000_0002
	NIM_SETFOCUS   NIM = 0x0000_0003
	NIM_SETVERSION NIM = 0x0000_0004
)

// [Shell_NotifyIcon] notifications, which are sent in the low-order bits of
// WPARAM.
//
// [Shell_NotifyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shell_notifyiconw
type NIN uint16

const (
	_NINF_KEY NIN = 0x1

	NIN_SELECT    = NIN(WM_USER + 0)
	NIN_KEYSELECT = NIN_SELECT | _NINF_KEY

	NIN_BALLOONSHOW      = NIN(WM_USER + 2)
	NIN_BALLOONHIDE      = NIN(WM_USER + 3)
	NIN_BALLOONTIMEOUT   = NIN(WM_USER + 4)
	NIN_BALLOONUSERCLICK = NIN(WM_USER + 5)
	NIN_POPUPOPEN        = NIN(WM_USER + 6)
	NIN_POPUPCLOSE       = NIN(WM_USER + 7)
)

// [NOTIFYICONDATA] dwState.
//
// [NOTIFYICONDATA]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NIS uint32

const (
	NIS_HIDDEN     NIS = 0x0000_0001
	NIS_SHAREDICON NIS = 0x0000_0002
)

// [PROPERTYKEY] struct.
//
// [PROPERTYKEY]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ns-wtypes-propertykey
type PROPERTYKEY struct {
	Fmtid GUID
	Pid   uint32
}

var (
	// Address properties

	PKEY_Address_Country     = PROPERTYKEY{GUID{Data1: 0xc07b4199, Data2: 0xe1df, Data3: 0x4493, Data4: [8]byte{0xb1, 0xe1, 0xde, 0x59, 0x46, 0xfb, 0x58, 0xf8}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Address_CountryCode = PROPERTYKEY{GUID{Data1: 0xc07b4199, Data2: 0xe1df, Data3: 0x4493, Data4: [8]byte{0xb1, 0xe1, 0xde, 0x59, 0x46, 0xfb, 0x58, 0xf8}}, 101} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Address_Region      = PROPERTYKEY{GUID{Data1: 0xc07b4199, Data2: 0xe1df, Data3: 0x4493, Data4: [8]byte{0xb1, 0xe1, 0xde, 0x59, 0x46, 0xfb, 0x58, 0xf8}}, 102} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Address_RegionCode  = PROPERTYKEY{GUID{Data1: 0xc07b4199, Data2: 0xe1df, Data3: 0x4493, Data4: [8]byte{0xb1, 0xe1, 0xde, 0x59, 0x46, 0xfb, 0x58, 0xf8}}, 103} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Address_Town        = PROPERTYKEY{GUID{Data1: 0xc07b4199, Data2: 0xe1df, Data3: 0x4493, Data4: [8]byte{0xb1, 0xe1, 0xde, 0x59, 0x46, 0xfb, 0x58, 0xf8}}, 104} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Audio properties

	PKEY_Audio_ChannelCount      = PROPERTYKEY{GUID{Data1: 0x64440490, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 7}   // UInt32 -- VT_UI4
	PKEY_Audio_Compression       = PROPERTYKEY{GUID{Data1: 0x64440490, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 10}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Audio_EncodingBitrate   = PROPERTYKEY{GUID{Data1: 0x64440490, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 4}   // UInt32 -- VT_UI4
	PKEY_Audio_Format            = PROPERTYKEY{GUID{Data1: 0x64440490, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 2}   // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_BSTR.
	PKEY_Audio_IsVariableBitRate = PROPERTYKEY{GUID{Data1: 0xe6822fee, Data2: 0x8c17, Data3: 0x4d62, Data4: [8]byte{0x82, 0x3c, 0x8e, 0x9c, 0xfc, 0xbd, 0x1d, 0x5c}}, 100} // Boolean -- VT_BOOL
	PKEY_Audio_PeakValue         = PROPERTYKEY{GUID{Data1: 0x2579e5d0, Data2: 0x1116, Data3: 0x4084, Data4: [8]byte{0xbd, 0x9a, 0x9b, 0x4f, 0x7c, 0xb4, 0xdf, 0x5e}}, 100} // UInt32 -- VT_UI4
	PKEY_Audio_SampleRate        = PROPERTYKEY{GUID{Data1: 0x64440490, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 5}   // UInt32 -- VT_UI4
	PKEY_Audio_SampleSize        = PROPERTYKEY{GUID{Data1: 0x64440490, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 6}   // UInt32 -- VT_UI4
	PKEY_Audio_StreamName        = PROPERTYKEY{GUID{Data1: 0x64440490, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 9}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Audio_StreamNumber      = PROPERTYKEY{GUID{Data1: 0x64440490, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 8}   // UInt16 -- VT_UI2

	// Calendar properties

	PKEY_Calendar_Duration                  = PROPERTYKEY{GUID{Data1: 0x293ca35a, Data2: 0x09aa, Data3: 0x4dd2, Data4: [8]byte{0xb1, 0x80, 0x1f, 0xe2, 0x45, 0x72, 0x8a, 0x52}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Calendar_IsOnline                  = PROPERTYKEY{GUID{Data1: 0xbfee9149, Data2: 0xe3e2, Data3: 0x49a7, Data4: [8]byte{0xa8, 0x62, 0xc0, 0x59, 0x88, 0x14, 0x5c, 0xec}}, 100} // Boolean -- VT_BOOL
	PKEY_Calendar_IsRecurring               = PROPERTYKEY{GUID{Data1: 0x315b9c8d, Data2: 0x80a9, Data3: 0x4ef9, Data4: [8]byte{0xae, 0x16, 0x8e, 0x74, 0x6d, 0xa5, 0x1d, 0x70}}, 100} // Boolean -- VT_BOOL
	PKEY_Calendar_Location                  = PROPERTYKEY{GUID{Data1: 0xf6272d18, Data2: 0xcecc, Data3: 0x40b1, Data4: [8]byte{0xb2, 0x6a, 0x39, 0x11, 0x71, 0x7a, 0xa7, 0xbd}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Calendar_OptionalAttendeeAddresses = PROPERTYKEY{GUID{Data1: 0xd55bae5a, Data2: 0x3892, Data3: 0x417a, Data4: [8]byte{0xa6, 0x49, 0xc6, 0xac, 0x5a, 0xaa, 0xea, 0xb3}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_OptionalAttendeeNames     = PROPERTYKEY{GUID{Data1: 0x09429607, Data2: 0x582d, Data3: 0x437f, Data4: [8]byte{0x84, 0xc3, 0xde, 0x93, 0xa2, 0xb2, 0x4c, 0x3c}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_OrganizerAddress          = PROPERTYKEY{GUID{Data1: 0x744c8242, Data2: 0x4df5, Data3: 0x456c, Data4: [8]byte{0xab, 0x9e, 0x01, 0x4e, 0xfb, 0x90, 0x21, 0xe3}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Calendar_OrganizerName             = PROPERTYKEY{GUID{Data1: 0xaaa660f9, Data2: 0x9865, Data3: 0x458e, Data4: [8]byte{0xb4, 0x84, 0x01, 0xbc, 0x7f, 0xe3, 0x97, 0x3e}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Calendar_ReminderTime              = PROPERTYKEY{GUID{Data1: 0x72fc5ba4, Data2: 0x24f9, Data3: 0x4011, Data4: [8]byte{0x9f, 0x3f, 0xad, 0xd2, 0x7a, 0xfa, 0xd8, 0x18}}, 100} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Calendar_RequiredAttendeeAddresses = PROPERTYKEY{GUID{Data1: 0x0ba7d6c3, Data2: 0x568d, Data3: 0x4159, Data4: [8]byte{0xab, 0x91, 0x78, 0x1a, 0x91, 0xfb, 0x71, 0xe5}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_RequiredAttendeeNames     = PROPERTYKEY{GUID{Data1: 0xb33af30b, Data2: 0xf552, Data3: 0x4584, Data4: [8]byte{0x93, 0x6c, 0xcb, 0x93, 0xe5, 0xcd, 0xa2, 0x9f}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_Resources                 = PROPERTYKEY{GUID{Data1: 0x00f58a38, Data2: 0xc54b, Data3: 0x4c40, Data4: [8]byte{0x86, 0x96, 0x97, 0x23, 0x59, 0x80, 0xea, 0xe1}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_ResponseStatus            = PROPERTYKEY{GUID{Data1: 0x188c1f91, Data2: 0x3c40, Data3: 0x4132, Data4: [8]byte{0x9e, 0xc5, 0xd8, 0xb0, 0x3b, 0x72, 0xa8, 0xa2}}, 100} // UInt16 -- VT_UI2
	PKEY_Calendar_ShowTimeAs                = PROPERTYKEY{GUID{Data1: 0x5bf396d4, Data2: 0x5eb2, Data3: 0x466f, Data4: [8]byte{0xbd, 0xe9, 0x2f, 0xb3, 0xf2, 0x36, 0x1d, 0x6e}}, 100} // UInt16 -- VT_UI2
	PKEY_Calendar_ShowTimeAsText            = PROPERTYKEY{GUID{Data1: 0x53da57cf, Data2: 0x62c0, Data3: 0x45c4, Data4: [8]byte{0x81, 0xde, 0x76, 0x10, 0xbc, 0xef, 0xd7, 0xf5}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Communication properties

	PKEY_Communication_AccountName       = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 9}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Communication_DateItemExpires   = PROPERTYKEY{GUID{Data1: 0x428040ac, Data2: 0xa177, Data3: 0x4c8a, Data4: [8]byte{0x97, 0x60, 0xf6, 0xf7, 0x61, 0x22, 0x7f, 0x9a}}, 100} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Communication_Direction         = PROPERTYKEY{GUID{Data1: 0x8e531030, Data2: 0xb960, Data3: 0x4346, Data4: [8]byte{0xae, 0x0d, 0x66, 0xbc, 0x9a, 0x86, 0xfb, 0x94}}, 100} // UInt16 -- VT_UI2
	PKEY_Communication_FollowupIconIndex = PROPERTYKEY{GUID{Data1: 0x83a6347e, Data2: 0x6fe4, Data3: 0x4f40, Data4: [8]byte{0xba, 0x9c, 0xc4, 0x86, 0x52, 0x40, 0xd1, 0xf4}}, 100} // Int32 -- VT_I4
	PKEY_Communication_HeaderItem        = PROPERTYKEY{GUID{Data1: 0xc9c34f84, Data2: 0x2241, Data3: 0x4401, Data4: [8]byte{0xb6, 0x07, 0xbd, 0x20, 0xed, 0x75, 0xae, 0x7f}}, 100} // Boolean -- VT_BOOL
	PKEY_Communication_PolicyTag         = PROPERTYKEY{GUID{Data1: 0xec0b4191, Data2: 0xab0b, Data3: 0x4c66, Data4: [8]byte{0x90, 0xb6, 0xc6, 0x63, 0x7c, 0xde, 0xbb, 0xab}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Communication_SecurityFlags     = PROPERTYKEY{GUID{Data1: 0x8619a4b6, Data2: 0x9f4d, Data3: 0x4429, Data4: [8]byte{0x8c, 0x0f, 0xb9, 0x96, 0xca, 0x59, 0xe3, 0x35}}, 100} // Int32 -- VT_I4
	PKEY_Communication_Suffix            = PROPERTYKEY{GUID{Data1: 0x807b653a, Data2: 0x9e91, Data3: 0x43ef, Data4: [8]byte{0x8f, 0x97, 0x11, 0xce, 0x04, 0xee, 0x20, 0xc5}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Communication_TaskStatus        = PROPERTYKEY{GUID{Data1: 0xbe1a72c6, Data2: 0x9a1d, Data3: 0x46b7, Data4: [8]byte{0xaf, 0xe7, 0xaf, 0xaf, 0x8c, 0xef, 0x49, 0x99}}, 100} // UInt16 -- VT_UI2
	PKEY_Communication_TaskStatusText    = PROPERTYKEY{GUID{Data1: 0xa6744477, Data2: 0xc237, Data3: 0x475b, Data4: [8]byte{0xa0, 0x75, 0x54, 0xf3, 0x44, 0x98, 0x29, 0x2a}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Computer properties

	PKEY_Computer_DecoratedFreeSpace = PROPERTYKEY{GUID{Data1: 0x9b174b35, Data2: 0x40ff, Data3: 0x11d2, Data4: [8]byte{0xa2, 0x7e, 0x00, 0xc0, 0x4f, 0xc3, 0x08, 0x71}}, 7} // Multivalue UInt64 -- VT_VECTOR | VT_UI8  (For variants: VT_ARRAY | VT_UI8)

	// Contact properties

	PKEY_Contact_AccountPictureDynamicVideo       = PROPERTYKEY{GUID{Data1: 0x0b8bb018, Data2: 0x2725, Data3: 0x4b44, Data4: [8]byte{0x92, 0xba, 0x79, 0x33, 0xae, 0xb2, 0xdd, 0xe7}}, 2}   // Stream -- VT_STREAM
	PKEY_Contact_AccountPictureLarge              = PROPERTYKEY{GUID{Data1: 0x0b8bb018, Data2: 0x2725, Data3: 0x4b44, Data4: [8]byte{0x92, 0xba, 0x79, 0x33, 0xae, 0xb2, 0xdd, 0xe7}}, 3}   // Stream -- VT_STREAM
	PKEY_Contact_AccountPictureSmall              = PROPERTYKEY{GUID{Data1: 0x0b8bb018, Data2: 0x2725, Data3: 0x4b44, Data4: [8]byte{0x92, 0xba, 0x79, 0x33, 0xae, 0xb2, 0xdd, 0xe7}}, 4}   // Stream -- VT_STREAM
	PKEY_Contact_Anniversary                      = PROPERTYKEY{GUID{Data1: 0x9ad5badb, Data2: 0xcea7, Data3: 0x4470, Data4: [8]byte{0xa0, 0x3d, 0xb8, 0x4e, 0x51, 0xb9, 0x94, 0x9e}}, 100} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Contact_AssistantName                    = PROPERTYKEY{GUID{Data1: 0xcd102c9c, Data2: 0x5540, Data3: 0x4a88, Data4: [8]byte{0xa6, 0xf6, 0x64, 0xe4, 0x98, 0x1c, 0x8c, 0xd1}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_AssistantTelephone               = PROPERTYKEY{GUID{Data1: 0x9a93244d, Data2: 0xa7ad, Data3: 0x4ff8, Data4: [8]byte{0x9b, 0x99, 0x45, 0xee, 0x4c, 0xc0, 0x9a, 0xf6}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Birthday                         = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 47}  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Contact_BusinessAddress                  = PROPERTYKEY{GUID{Data1: 0x730fb6dd, Data2: 0xcf7c, Data3: 0x426b, Data4: [8]byte{0xa0, 0x3f, 0xbd, 0x16, 0x6c, 0xc9, 0xee, 0x24}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1Country          = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 119} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1Locality         = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 117} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1PostalCode       = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 120} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1Region           = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 118} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1Street           = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 116} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2Country          = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 124} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2Locality         = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 122} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2PostalCode       = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 125} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2Region           = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 123} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2Street           = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 121} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3Country          = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 129} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3Locality         = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 127} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3PostalCode       = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 130} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3Region           = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 128} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3Street           = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 126} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressCity              = PROPERTYKEY{GUID{Data1: 0x402b5934, Data2: 0xec5a, Data3: 0x48c3, Data4: [8]byte{0x93, 0xe6, 0x85, 0xe8, 0x6a, 0x2d, 0x93, 0x4e}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressCountry           = PROPERTYKEY{GUID{Data1: 0xb0b87314, Data2: 0xfcf6, Data3: 0x4feb, Data4: [8]byte{0x8d, 0xff, 0xa5, 0x0d, 0xa6, 0xaf, 0x56, 0x1c}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressPostalCode        = PROPERTYKEY{GUID{Data1: 0xe1d4a09e, Data2: 0xd758, Data3: 0x4cd1, Data4: [8]byte{0xb6, 0xec, 0x34, 0xa8, 0xb5, 0xa7, 0x3f, 0x80}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressPostOfficeBox     = PROPERTYKEY{GUID{Data1: 0xbc4e71ce, Data2: 0x17f9, Data3: 0x48d5, Data4: [8]byte{0xbe, 0xe9, 0x02, 0x1d, 0xf0, 0xea, 0x54, 0x09}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressState             = PROPERTYKEY{GUID{Data1: 0x446f787f, Data2: 0x10c4, Data3: 0x41cb, Data4: [8]byte{0xa6, 0xc4, 0x4d, 0x03, 0x43, 0x55, 0x15, 0x97}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressStreet            = PROPERTYKEY{GUID{Data1: 0xddd1460f, Data2: 0xc0bf, Data3: 0x4553, Data4: [8]byte{0x8c, 0xe4, 0x10, 0x43, 0x3c, 0x90, 0x8f, 0xb0}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessEmailAddresses           = PROPERTYKEY{GUID{Data1: 0xf271c659, Data2: 0x7e5e, Data3: 0x471f, Data4: [8]byte{0xba, 0x25, 0x7f, 0x77, 0xb2, 0x86, 0xf8, 0x36}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_BusinessFaxNumber                = PROPERTYKEY{GUID{Data1: 0x91eff6f3, Data2: 0x2e27, Data3: 0x42ca, Data4: [8]byte{0x93, 0x3e, 0x7c, 0x99, 0x9f, 0xbe, 0x31, 0x0b}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessHomePage                 = PROPERTYKEY{GUID{Data1: 0x56310920, Data2: 0x2491, Data3: 0x4919, Data4: [8]byte{0x99, 0xce, 0xea, 0xdb, 0x06, 0xfa, 0xfd, 0xb2}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessTelephone                = PROPERTYKEY{GUID{Data1: 0x6a15e5a0, Data2: 0x0a1e, Data3: 0x4cd7, Data4: [8]byte{0xbb, 0x8c, 0xd2, 0xf1, 0xb0, 0xc9, 0x29, 0xbc}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_CallbackTelephone                = PROPERTYKEY{GUID{Data1: 0xbf53d1c3, Data2: 0x49e0, Data3: 0x4f7f, Data4: [8]byte{0x85, 0x67, 0x5a, 0x82, 0x1d, 0x8a, 0xc5, 0x42}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_CarTelephone                     = PROPERTYKEY{GUID{Data1: 0x8fdc6dea, Data2: 0xb929, Data3: 0x412b, Data4: [8]byte{0xba, 0x90, 0x39, 0x7a, 0x25, 0x74, 0x65, 0xfe}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Children                         = PROPERTYKEY{GUID{Data1: 0xd4729704, Data2: 0x8ef1, Data3: 0x43ef, Data4: [8]byte{0x90, 0x24, 0x2b, 0xd3, 0x81, 0x18, 0x7f, 0xd5}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_CompanyMainTelephone             = PROPERTYKEY{GUID{Data1: 0x8589e481, Data2: 0x6040, Data3: 0x473d, Data4: [8]byte{0xb1, 0x71, 0x7f, 0xa8, 0x9c, 0x27, 0x08, 0xed}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_ConnectedServiceDisplayName      = PROPERTYKEY{GUID{Data1: 0x39b77f4f, Data2: 0xa104, Data3: 0x4863, Data4: [8]byte{0xb3, 0x95, 0x2d, 0xb2, 0xad, 0x8f, 0x7b, 0xc1}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_ConnectedServiceIdentities       = PROPERTYKEY{GUID{Data1: 0x80f41eb8, Data2: 0xafc4, Data3: 0x4208, Data4: [8]byte{0xaa, 0x5f, 0xcc, 0xe2, 0x1a, 0x62, 0x72, 0x81}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_ConnectedServiceName             = PROPERTYKEY{GUID{Data1: 0xb5c84c9e, Data2: 0x5927, Data3: 0x46b5, Data4: [8]byte{0xa3, 0xcc, 0x93, 0x3c, 0x21, 0xb7, 0x84, 0x69}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_ConnectedServiceSupportedActions = PROPERTYKEY{GUID{Data1: 0xa19fb7a9, Data2: 0x024b, Data3: 0x4371, Data4: [8]byte{0xa8, 0xbf, 0x4d, 0x29, 0xc3, 0xe4, 0xe9, 0xc9}}, 100} // UInt32 -- VT_UI4
	PKEY_Contact_DataSuppliers                    = PROPERTYKEY{GUID{Data1: 0x9660c283, Data2: 0xfc3a, Data3: 0x4a08, Data4: [8]byte{0xa0, 0x96, 0xee, 0xd3, 0xaa, 0xc4, 0x6d, 0xa2}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_Department                       = PROPERTYKEY{GUID{Data1: 0xfc9f7306, Data2: 0xff8f, Data3: 0x4d49, Data4: [8]byte{0x9f, 0xb6, 0x3f, 0xfe, 0x5c, 0x09, 0x51, 0xec}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_DisplayBusinessPhoneNumbers      = PROPERTYKEY{GUID{Data1: 0x364028da, Data2: 0xd895, Data3: 0x41fe, Data4: [8]byte{0xa5, 0x84, 0x30, 0x2b, 0x1b, 0xb7, 0x0a, 0x76}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_DisplayHomePhoneNumbers          = PROPERTYKEY{GUID{Data1: 0x5068bcdf, Data2: 0xd697, Data3: 0x4d85, Data4: [8]byte{0x8c, 0x53, 0x1f, 0x1c, 0xda, 0xb0, 0x17, 0x63}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_DisplayMobilePhoneNumbers        = PROPERTYKEY{GUID{Data1: 0x9cb0c358, Data2: 0x9d7a, Data3: 0x46b1, Data4: [8]byte{0xb4, 0x66, 0xdc, 0xc6, 0xf1, 0xa3, 0xd9, 0x3d}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_DisplayOtherPhoneNumbers         = PROPERTYKEY{GUID{Data1: 0x03089873, Data2: 0x8ee8, Data3: 0x4191, Data4: [8]byte{0xbd, 0x60, 0xd3, 0x1f, 0x72, 0xb7, 0x90, 0x0b}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_EmailAddress                     = PROPERTYKEY{GUID{Data1: 0xf8fa7fa3, Data2: 0xd12b, Data3: 0x4785, Data4: [8]byte{0x8a, 0x4e, 0x69, 0x1a, 0x94, 0xf7, 0xa3, 0xe7}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_EmailAddress2                    = PROPERTYKEY{GUID{Data1: 0x38965063, Data2: 0xedc8, Data3: 0x4268, Data4: [8]byte{0x84, 0x91, 0xb7, 0x72, 0x31, 0x72, 0xcf, 0x29}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_EmailAddress3                    = PROPERTYKEY{GUID{Data1: 0x644d37b4, Data2: 0xe1b3, Data3: 0x4bad, Data4: [8]byte{0xb0, 0x99, 0x7e, 0x7c, 0x04, 0x96, 0x6a, 0xca}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_EmailAddresses                   = PROPERTYKEY{GUID{Data1: 0x84d8f337, Data2: 0x981d, Data3: 0x44b3, Data4: [8]byte{0x96, 0x15, 0xc7, 0x59, 0x6d, 0xba, 0x17, 0xe3}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_EmailName                        = PROPERTYKEY{GUID{Data1: 0xcc6f4f24, Data2: 0x6083, Data3: 0x4bd4, Data4: [8]byte{0x87, 0x54, 0x67, 0x4d, 0x0d, 0xe8, 0x7a, 0xb8}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_FileAsName                       = PROPERTYKEY{GUID{Data1: 0xf1a24aa7, Data2: 0x9ca7, Data3: 0x40f6, Data4: [8]byte{0x89, 0xec, 0x97, 0xde, 0xf9, 0xff, 0xe8, 0xdb}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_FirstName                        = PROPERTYKEY{GUID{Data1: 0x14977844, Data2: 0x6b49, Data3: 0x4aad, Data4: [8]byte{0xa7, 0x14, 0xa4, 0x51, 0x3b, 0xf6, 0x04, 0x60}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_FullName                         = PROPERTYKEY{GUID{Data1: 0x635e9051, Data2: 0x50a5, Data3: 0x4ba2, Data4: [8]byte{0xb9, 0xdb, 0x4e, 0xd0, 0x56, 0xc7, 0x72, 0x96}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Gender                           = PROPERTYKEY{GUID{Data1: 0x3c8cee58, Data2: 0xd4f0, Data3: 0x4cf9, Data4: [8]byte{0xb7, 0x56, 0x4e, 0x5d, 0x24, 0x44, 0x7b, 0xcd}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_GenderValue                      = PROPERTYKEY{GUID{Data1: 0x3c8cee58, Data2: 0xd4f0, Data3: 0x4cf9, Data4: [8]byte{0xb7, 0x56, 0x4e, 0x5d, 0x24, 0x44, 0x7b, 0xcd}}, 101} // UInt16 -- VT_UI2
	PKEY_Contact_Hobbies                          = PROPERTYKEY{GUID{Data1: 0x5dc2253f, Data2: 0x5e11, Data3: 0x4adf, Data4: [8]byte{0x9c, 0xfe, 0x91, 0x0d, 0xd0, 0x1e, 0x3e, 0x70}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_HomeAddress                      = PROPERTYKEY{GUID{Data1: 0x98f98354, Data2: 0x617a, Data3: 0x46b8, Data4: [8]byte{0x85, 0x60, 0x5b, 0x1b, 0x64, 0xbf, 0x1f, 0x89}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1Country              = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 104} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1Locality             = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 102} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1PostalCode           = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 105} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1Region               = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 103} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1Street               = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 101} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2Country              = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 109} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2Locality             = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 107} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2PostalCode           = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 110} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2Region               = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 108} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2Street               = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 106} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3Country              = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 114} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3Locality             = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 112} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3PostalCode           = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 115} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3Region               = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 113} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3Street               = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 111} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressCity                  = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 65}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressCountry               = PROPERTYKEY{GUID{Data1: 0x08a65aa1, Data2: 0xf4c9, Data3: 0x43dd, Data4: [8]byte{0x9d, 0xdf, 0xa3, 0x3d, 0x8e, 0x7e, 0xad, 0x85}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressPostalCode            = PROPERTYKEY{GUID{Data1: 0x8afcc170, Data2: 0x8a46, Data3: 0x4b53, Data4: [8]byte{0x9e, 0xee, 0x90, 0xba, 0xe7, 0x15, 0x1e, 0x62}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressPostOfficeBox         = PROPERTYKEY{GUID{Data1: 0x7b9f6399, Data2: 0x0a3f, Data3: 0x4b12, Data4: [8]byte{0x89, 0xbd, 0x4a, 0xdc, 0x51, 0xc9, 0x18, 0xaf}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressState                 = PROPERTYKEY{GUID{Data1: 0xc89a23d0, Data2: 0x7d6d, Data3: 0x4eb8, Data4: [8]byte{0x87, 0xd4, 0x77, 0x6a, 0x82, 0xd4, 0x93, 0xe5}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressStreet                = PROPERTYKEY{GUID{Data1: 0x0adef160, Data2: 0xdb3f, Data3: 0x4308, Data4: [8]byte{0x9a, 0x21, 0x06, 0x23, 0x7b, 0x16, 0xfa, 0x2a}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeEmailAddresses               = PROPERTYKEY{GUID{Data1: 0x56c90e9d, Data2: 0x9d46, Data3: 0x4963, Data4: [8]byte{0x88, 0x6f, 0x2e, 0x1c, 0xd9, 0xa6, 0x94, 0xef}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_HomeFaxNumber                    = PROPERTYKEY{GUID{Data1: 0x660e04d6, Data2: 0x81ab, Data3: 0x4977, Data4: [8]byte{0xa0, 0x9f, 0x82, 0x31, 0x31, 0x13, 0xab, 0x26}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeTelephone                    = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 20}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_IMAddress                        = PROPERTYKEY{GUID{Data1: 0xd68dbd8a, Data2: 0x3374, Data3: 0x4b81, Data4: [8]byte{0x99, 0x72, 0x3e, 0xc3, 0x06, 0x82, 0xdb, 0x3d}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_Initials                         = PROPERTYKEY{GUID{Data1: 0xf3d8f40d, Data2: 0x50cb, Data3: 0x44a2, Data4: [8]byte{0x97, 0x18, 0x40, 0xcb, 0x91, 0x19, 0x49, 0x5d}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JA_CompanyNamePhonetic           = PROPERTYKEY{GUID{Data1: 0x897b3694, Data2: 0xfe9e, Data3: 0x43e6, Data4: [8]byte{0x80, 0x66, 0x26, 0x0f, 0x59, 0x0c, 0x01, 0x00}}, 2}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JA_FirstNamePhonetic             = PROPERTYKEY{GUID{Data1: 0x897b3694, Data2: 0xfe9e, Data3: 0x43e6, Data4: [8]byte{0x80, 0x66, 0x26, 0x0f, 0x59, 0x0c, 0x01, 0x00}}, 3}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JA_LastNamePhonetic              = PROPERTYKEY{GUID{Data1: 0x897b3694, Data2: 0xfe9e, Data3: 0x43e6, Data4: [8]byte{0x80, 0x66, 0x26, 0x0f, 0x59, 0x0c, 0x01, 0x00}}, 4}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1CompanyAddress           = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 120} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1CompanyName              = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 102} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1Department               = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 106} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1Manager                  = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 105} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1OfficeLocation           = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 104} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1Title                    = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 103} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1YomiCompanyName          = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 101} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2CompanyAddress           = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 121} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2CompanyName              = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 108} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2Department               = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 113} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2Manager                  = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 112} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2OfficeLocation           = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 110} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2Title                    = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 109} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2YomiCompanyName          = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 107} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3CompanyAddress           = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 123} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3CompanyName              = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 115} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3Department               = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 119} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3Manager                  = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 118} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3OfficeLocation           = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 117} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3Title                    = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 116} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3YomiCompanyName          = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 114} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobTitle                         = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 6}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Label                            = PROPERTYKEY{GUID{Data1: 0x97b0ad89, Data2: 0xdf49, Data3: 0x49cc, Data4: [8]byte{0x83, 0x4e, 0x66, 0x09, 0x74, 0xfd, 0x75, 0x5b}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_LastName                         = PROPERTYKEY{GUID{Data1: 0x8f367200, Data2: 0xc270, Data3: 0x457c, Data4: [8]byte{0xb1, 0xd4, 0xe0, 0x7c, 0x5b, 0xcd, 0x90, 0xc7}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_MailingAddress                   = PROPERTYKEY{GUID{Data1: 0xc0ac206a, Data2: 0x827e, Data3: 0x4650, Data4: [8]byte{0x95, 0xae, 0x77, 0xe2, 0xbb, 0x74, 0xfc, 0xc9}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_MiddleName                       = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 71}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_MobileTelephone                  = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 35}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_NickName                         = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 74}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OfficeLocation                   = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 7}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress                     = PROPERTYKEY{GUID{Data1: 0x508161fa, Data2: 0x313b, Data3: 0x43d5, Data4: [8]byte{0x83, 0xa1, 0xc1, 0xac, 0xcf, 0x68, 0x62, 0x2c}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1Country             = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 134} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1Locality            = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 132} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1PostalCode          = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 135} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1Region              = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 133} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1Street              = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 131} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2Country             = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 139} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2Locality            = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 137} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2PostalCode          = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 140} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2Region              = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 138} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2Street              = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 136} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3Country             = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 144} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3Locality            = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 142} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3PostalCode          = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 145} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3Region              = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 143} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3Street              = PROPERTYKEY{GUID{Data1: 0xa7b6f596, Data2: 0xd678, Data3: 0x4bc1, Data4: [8]byte{0xb0, 0x5f, 0x02, 0x03, 0xd2, 0x7e, 0x8a, 0xa1}}, 141} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressCity                 = PROPERTYKEY{GUID{Data1: 0x6e682923, Data2: 0x7f7b, Data3: 0x4f0c, Data4: [8]byte{0xa3, 0x37, 0xcf, 0xca, 0x29, 0x66, 0x87, 0xbf}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressCountry              = PROPERTYKEY{GUID{Data1: 0x8f167568, Data2: 0x0aae, Data3: 0x4322, Data4: [8]byte{0x8e, 0xd9, 0x60, 0x55, 0xb7, 0xb0, 0xe3, 0x98}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressPostalCode           = PROPERTYKEY{GUID{Data1: 0x95c656c1, Data2: 0x2abf, Data3: 0x4148, Data4: [8]byte{0x9e, 0xd3, 0x9e, 0xc6, 0x02, 0xe3, 0xb7, 0xcd}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressPostOfficeBox        = PROPERTYKEY{GUID{Data1: 0x8b26ea41, Data2: 0x058f, Data3: 0x43f6, Data4: [8]byte{0xae, 0xcc, 0x40, 0x35, 0x68, 0x1c, 0xe9, 0x77}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressState                = PROPERTYKEY{GUID{Data1: 0x71b377d6, Data2: 0xe570, Data3: 0x425f, Data4: [8]byte{0xa1, 0x70, 0x80, 0x9f, 0xae, 0x73, 0xe5, 0x4e}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressStreet               = PROPERTYKEY{GUID{Data1: 0xff962609, Data2: 0xb7d6, Data3: 0x4999, Data4: [8]byte{0x86, 0x2d, 0x95, 0x18, 0x0d, 0x52, 0x9a, 0xea}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherEmailAddresses              = PROPERTYKEY{GUID{Data1: 0x11d6336b, Data2: 0x38c4, Data3: 0x4ec9, Data4: [8]byte{0x84, 0xd6, 0xeb, 0x38, 0xd0, 0xb1, 0x50, 0xaf}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_PagerTelephone                   = PROPERTYKEY{GUID{Data1: 0xd6304e01, Data2: 0xf8f5, Data3: 0x4f45, Data4: [8]byte{0x8b, 0x15, 0xd0, 0x24, 0xa6, 0x29, 0x67, 0x89}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PersonalTitle                    = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 69}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PhoneNumbersCanonical            = PROPERTYKEY{GUID{Data1: 0xd042d2a1, Data2: 0x927e, Data3: 0x40b5, Data4: [8]byte{0xa5, 0x03, 0x6e, 0xdb, 0xd4, 0x2a, 0x51, 0x7e}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_Prefix                           = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 75}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressCity               = PROPERTYKEY{GUID{Data1: 0xc8ea94f0, Data2: 0xa9e3, Data3: 0x4969, Data4: [8]byte{0xa9, 0x4b, 0x9c, 0x62, 0xa9, 0x53, 0x24, 0xe0}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressCountry            = PROPERTYKEY{GUID{Data1: 0xe53d799d, Data2: 0x0f3f, Data3: 0x466e, Data4: [8]byte{0xb2, 0xff, 0x74, 0x63, 0x4a, 0x3c, 0xb7, 0xa4}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressPostalCode         = PROPERTYKEY{GUID{Data1: 0x18bbd425, Data2: 0xecfd, Data3: 0x46ef, Data4: [8]byte{0xb6, 0x12, 0x7b, 0x4a, 0x60, 0x34, 0xed, 0xa0}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressPostOfficeBox      = PROPERTYKEY{GUID{Data1: 0xde5ef3c7, Data2: 0x46e1, Data3: 0x484e, Data4: [8]byte{0x99, 0x99, 0x62, 0xc5, 0x30, 0x83, 0x94, 0xc1}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressState              = PROPERTYKEY{GUID{Data1: 0xf1176dfe, Data2: 0x7138, Data3: 0x4640, Data4: [8]byte{0x8b, 0x4c, 0xae, 0x37, 0x5d, 0xc7, 0x0a, 0x6d}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressStreet             = PROPERTYKEY{GUID{Data1: 0x63c25b20, Data2: 0x96be, Data3: 0x488f, Data4: [8]byte{0x87, 0x88, 0xc0, 0x9c, 0x40, 0x7a, 0xd8, 0x12}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryEmailAddress              = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 48}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryTelephone                 = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 25}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Profession                       = PROPERTYKEY{GUID{Data1: 0x7268af55, Data2: 0x1ce4, Data3: 0x4f6e, Data4: [8]byte{0xa4, 0x1f, 0xb6, 0xe4, 0xef, 0x10, 0xe4, 0xa9}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_SpouseName                       = PROPERTYKEY{GUID{Data1: 0x9d2408b6, Data2: 0x3167, Data3: 0x422b, Data4: [8]byte{0x82, 0xb0, 0xf5, 0x83, 0xb7, 0xa7, 0xcf, 0xe3}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Suffix                           = PROPERTYKEY{GUID{Data1: 0x176dc63c, Data2: 0x2688, Data3: 0x4e89, Data4: [8]byte{0x81, 0x43, 0xa3, 0x47, 0x80, 0x0f, 0x25, 0xe9}}, 73}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_TelexNumber                      = PROPERTYKEY{GUID{Data1: 0xc554493c, Data2: 0xc1f7, Data3: 0x40c1, Data4: [8]byte{0xa7, 0x6c, 0xef, 0x8c, 0x06, 0x14, 0x00, 0x3e}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_TTYTDDTelephone                  = PROPERTYKEY{GUID{Data1: 0xaaf16bac, Data2: 0x2b55, Data3: 0x45e6, Data4: [8]byte{0x9f, 0x6d, 0x41, 0x5e, 0xb9, 0x49, 0x10, 0xdf}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_WebPage                          = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 18}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Webpage2                         = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 124} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Webpage3                         = PROPERTYKEY{GUID{Data1: 0x00f63dd8, Data2: 0x22bd, Data3: 0x4a5d, Data4: [8]byte{0xba, 0x34, 0x5c, 0xb0, 0xb9, 0xbd, 0xcb, 0x03}}, 125} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Core properties

	PKEY_AcquisitionID                                      = PROPERTYKEY{GUID{Data1: 0x65a98875, Data2: 0x3c80, Data3: 0x40ab, Data4: [8]byte{0xab, 0xbc, 0xef, 0xda, 0xf7, 0x7d, 0xbe, 0xe2}}, 100}   // Int32 -- VT_I4
	PKEY_ApplicationDefinedProperties                       = PROPERTYKEY{GUID{Data1: 0xcdbfc167, Data2: 0x337e, Data3: 0x41d8, Data4: [8]byte{0xaf, 0x7c, 0x8c, 0x09, 0x20, 0x54, 0x29, 0xc7}}, 100}   // Any -- VT_NULL  Legacy code may treat this as VT_UNKNOWN.
	PKEY_ApplicationName                                    = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 18}    // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_AppZoneIdentifier                                  = PROPERTYKEY{GUID{Data1: 0x502cfeab, Data2: 0x47eb, Data3: 0x459c, Data4: [8]byte{0xb9, 0x60, 0xe6, 0xd8, 0x72, 0x8f, 0x77, 0x01}}, 102}   // UInt32 -- VT_UI4
	PKEY_Author                                             = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 4}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_CachedFileUpdaterContentIdForConflictResolution    = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 114}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_CachedFileUpdaterContentIdForStream                = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 113}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Capacity                                           = PROPERTYKEY{GUID{Data1: 0x9b174b35, Data2: 0x40ff, Data3: 0x11d2, Data4: [8]byte{0xa2, 0x7e, 0x00, 0xc0, 0x4f, 0xc3, 0x08, 0x71}}, 3}     // UInt64 -- VT_UI8
	PKEY_Category                                           = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 2}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Comment                                            = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 6}     // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_Company                                            = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 15}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ComputerName                                       = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 5}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ContainedItems                                     = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 29}    // Multivalue Guid -- VT_VECTOR | VT_CLSID  (For variants: VT_ARRAY | VT_CLSID)
	PKEY_ContentId                                          = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 132}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ContentStatus                                      = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 27}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ContentType                                        = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 26}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ContentUri                                         = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 131}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Copyright                                          = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 11}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_CreatorAppId                                       = PROPERTYKEY{GUID{Data1: 0xc2ea046e, Data2: 0x033c, Data3: 0x4e91, Data4: [8]byte{0xbd, 0x5b, 0xd4, 0x94, 0x2f, 0x6b, 0xbe, 0x49}}, 2}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_CreatorOpenWithUIOptions                           = PROPERTYKEY{GUID{Data1: 0xc2ea046e, Data2: 0x033c, Data3: 0x4e91, Data4: [8]byte{0xbd, 0x5b, 0xd4, 0x94, 0x2f, 0x6b, 0xbe, 0x49}}, 3}     // UInt32 -- VT_UI4
	PKEY_DataObjectFormat                                   = PROPERTYKEY{GUID{Data1: 0x1e81a3f8, Data2: 0xa30f, Data3: 0x4247, Data4: [8]byte{0xb9, 0xee, 0x1d, 0x03, 0x68, 0xa9, 0x42, 0x5c}}, 2}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_DateAccessed                                       = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 16}    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateAcquired                                       = PROPERTYKEY{GUID{Data1: 0x2cbaa8f5, Data2: 0xd81f, Data3: 0x47ca, Data4: [8]byte{0xb1, 0x7a, 0xf8, 0xd8, 0x22, 0x30, 0x01, 0x31}}, 100}   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateArchived                                       = PROPERTYKEY{GUID{Data1: 0x43f8d7b7, Data2: 0xa444, Data3: 0x4f87, Data4: [8]byte{0x93, 0x83, 0x52, 0x27, 0x1c, 0x9b, 0x91, 0x5c}}, 100}   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateCompleted                                      = PROPERTYKEY{GUID{Data1: 0x72fab781, Data2: 0xacda, Data3: 0x43e5, Data4: [8]byte{0xb1, 0x55, 0xb2, 0x43, 0x4f, 0x85, 0xe6, 0x78}}, 100}   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateCreated                                        = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 15}    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateImported                                       = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 18258} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateModified                                       = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 14}    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DefaultSaveLocationDisplay                         = PROPERTYKEY{GUID{Data1: 0x5d76b67f, Data2: 0x9b3d, Data3: 0x44bb, Data4: [8]byte{0xb6, 0xae, 0x25, 0xda, 0x4f, 0x63, 0x8a, 0x67}}, 10}    // UInt32 -- VT_UI4
	PKEY_DueDate                                            = PROPERTYKEY{GUID{Data1: 0x3f8472b5, Data2: 0xe0af, Data3: 0x4db2, Data4: [8]byte{0x80, 0x71, 0xc5, 0x3f, 0xe7, 0x6a, 0xe7, 0xce}}, 100}   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_EndDate                                            = PROPERTYKEY{GUID{Data1: 0xc75faa05, Data2: 0x96fd, Data3: 0x49e7, Data4: [8]byte{0x9c, 0xb4, 0x9f, 0x60, 0x10, 0x82, 0xd5, 0x53}}, 100}   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_ExpandoProperties                                  = PROPERTYKEY{GUID{Data1: 0x6fa20de6, Data2: 0xd11c, Data3: 0x4d9d, Data4: [8]byte{0xa1, 0x54, 0x64, 0x31, 0x76, 0x28, 0xc1, 0x2d}}, 100}   // Any -- VT_NULL  Legacy code may treat this as VT_UNKNOWN.
	PKEY_FileAllocationSize                                 = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 18}    // UInt64 -- VT_UI8
	PKEY_FileAttributes                                     = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 13}    // UInt32 -- VT_UI4
	PKEY_FileCount                                          = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 12}    // UInt64 -- VT_UI8
	PKEY_FileDescription                                    = PROPERTYKEY{GUID{Data1: 0x0cef7d53, Data2: 0xfa64, Data3: 0x11d1, Data4: [8]byte{0xa2, 0x03, 0x00, 0x00, 0xf8, 0x1f, 0xed, 0xee}}, 3}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FileExtension                                      = PROPERTYKEY{GUID{Data1: 0xe4f10a3c, Data2: 0x49e6, Data3: 0x405d, Data4: [8]byte{0x82, 0x88, 0xa2, 0x3b, 0xd4, 0xee, 0xaa, 0x6c}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FileFRN                                            = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 21}    // UInt64 -- VT_UI8
	PKEY_FileName                                           = PROPERTYKEY{GUID{Data1: 0x41cf5ae0, Data2: 0xf75a, Data3: 0x4806, Data4: [8]byte{0xbd, 0x87, 0x59, 0xc7, 0xd9, 0x24, 0x8e, 0xb9}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FileOfflineAvailabilityStatus                      = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 100}   // UInt32 -- VT_UI4
	PKEY_FileOwner                                          = PROPERTYKEY{GUID{Data1: 0x9b174b34, Data2: 0x40ff, Data3: 0x11d2, Data4: [8]byte{0xa2, 0x7e, 0x00, 0xc0, 0x4f, 0xc3, 0x08, 0x71}}, 4}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FilePlaceholderStatus                              = PROPERTYKEY{GUID{Data1: 0xb2f9b9d6, Data2: 0xfec4, Data3: 0x4dd5, Data4: [8]byte{0x94, 0xd7, 0x89, 0x57, 0x48, 0x8c, 0x80, 0x7b}}, 2}     // UInt32 -- VT_UI4
	PKEY_FileVersion                                        = PROPERTYKEY{GUID{Data1: 0x0cef7d53, Data2: 0xfa64, Data3: 0x11d1, Data4: [8]byte{0xa2, 0x03, 0x00, 0x00, 0xf8, 0x1f, 0xed, 0xee}}, 4}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FindData                                           = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 0}     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_FlagColor                                          = PROPERTYKEY{GUID{Data1: 0x67df94de, Data2: 0x0ca7, Data3: 0x4d6f, Data4: [8]byte{0xb7, 0x92, 0x05, 0x3a, 0x3e, 0x4f, 0x03, 0xcf}}, 100}   // UInt16 -- VT_UI2
	PKEY_FlagColorText                                      = PROPERTYKEY{GUID{Data1: 0x45eae747, Data2: 0x8e2a, Data3: 0x40ae, Data4: [8]byte{0x8c, 0xbf, 0xca, 0x52, 0xab, 0xa6, 0x15, 0x2a}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FlagStatus                                         = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 12}    // Int32 -- VT_I4
	PKEY_FlagStatusText                                     = PROPERTYKEY{GUID{Data1: 0xdc54fd2e, Data2: 0x189d, Data3: 0x4871, Data4: [8]byte{0xaa, 0x01, 0x08, 0xc2, 0xf5, 0x7a, 0x4a, 0xbc}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FolderKind                                         = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 101}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_FolderNameDisplay                                  = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 25}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FreeSpace                                          = PROPERTYKEY{GUID{Data1: 0x9b174b35, Data2: 0x40ff, Data3: 0x11d2, Data4: [8]byte{0xa2, 0x7e, 0x00, 0xc0, 0x4f, 0xc3, 0x08, 0x71}}, 2}     // UInt64 -- VT_UI8
	PKEY_FullText                                           = PROPERTYKEY{GUID{Data1: 0x1e3ee840, Data2: 0xbc2b, Data3: 0x476c, Data4: [8]byte{0x82, 0x37, 0x2a, 0xcd, 0x1a, 0x83, 0x9b, 0x22}}, 6}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_HighKeywords                                       = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 24}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Identity                                           = PROPERTYKEY{GUID{Data1: 0xa26f4afc, Data2: 0x7346, Data3: 0x4299, Data4: [8]byte{0xbe, 0x47, 0xeb, 0x1a, 0xe6, 0x13, 0x13, 0x9f}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_Blob                                      = PROPERTYKEY{GUID{Data1: 0x8c3b93a4, Data2: 0xbaed, Data3: 0x1a83, Data4: [8]byte{0x9a, 0x32, 0x10, 0x2e, 0xe3, 0x13, 0xf6, 0xeb}}, 100}   // Blob -- VT_BLOB
	PKEY_Identity_DisplayName                               = PROPERTYKEY{GUID{Data1: 0x7d683fc9, Data2: 0xd155, Data3: 0x45a8, Data4: [8]byte{0xbb, 0x1f, 0x89, 0xd1, 0x9b, 0xcb, 0x79, 0x2f}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_InternetSid                               = PROPERTYKEY{GUID{Data1: 0x6d6d5d49, Data2: 0x265d, Data3: 0x4688, Data4: [8]byte{0x9f, 0x4e, 0x1f, 0xdd, 0x33, 0xe7, 0xcc, 0x83}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_IsMeIdentity                              = PROPERTYKEY{GUID{Data1: 0xa4108708, Data2: 0x09df, Data3: 0x4377, Data4: [8]byte{0x9d, 0xfc, 0x6d, 0x99, 0x98, 0x6d, 0x5a, 0x67}}, 100}   // Boolean -- VT_BOOL
	PKEY_Identity_KeyProviderContext                        = PROPERTYKEY{GUID{Data1: 0xa26f4afc, Data2: 0x7346, Data3: 0x4299, Data4: [8]byte{0xbe, 0x47, 0xeb, 0x1a, 0xe6, 0x13, 0x13, 0x9f}}, 17}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_KeyProviderName                           = PROPERTYKEY{GUID{Data1: 0xa26f4afc, Data2: 0x7346, Data3: 0x4299, Data4: [8]byte{0xbe, 0x47, 0xeb, 0x1a, 0xe6, 0x13, 0x13, 0x9f}}, 16}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_LogonStatusString                         = PROPERTYKEY{GUID{Data1: 0xf18dedf3, Data2: 0x337f, Data3: 0x42c0, Data4: [8]byte{0x9e, 0x03, 0xce, 0xe0, 0x87, 0x08, 0xa8, 0xc3}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_PrimaryEmailAddress                       = PROPERTYKEY{GUID{Data1: 0xfcc16823, Data2: 0xbaed, Data3: 0x4f24, Data4: [8]byte{0x9b, 0x32, 0xa0, 0x98, 0x21, 0x17, 0xf7, 0xfa}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_PrimarySid                                = PROPERTYKEY{GUID{Data1: 0x2b1b801e, Data2: 0xc0c1, Data3: 0x4987, Data4: [8]byte{0x9e, 0xc5, 0x72, 0xfa, 0x89, 0x81, 0x47, 0x87}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_ProviderData                              = PROPERTYKEY{GUID{Data1: 0xa8a74b92, Data2: 0x361b, Data3: 0x4e9a, Data4: [8]byte{0xb7, 0x22, 0x7c, 0x4a, 0x73, 0x30, 0xa3, 0x12}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_ProviderID                                = PROPERTYKEY{GUID{Data1: 0x74a7de49, Data2: 0xfa11, Data3: 0x4d3d, Data4: [8]byte{0xa0, 0x06, 0xdb, 0x7e, 0x08, 0x67, 0x59, 0x16}}, 100}   // Guid -- VT_CLSID
	PKEY_Identity_QualifiedUserName                         = PROPERTYKEY{GUID{Data1: 0xda520e51, Data2: 0xf4e9, Data3: 0x4739, Data4: [8]byte{0xac, 0x82, 0x02, 0xe0, 0xa9, 0x5c, 0x90, 0x30}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_UniqueID                                  = PROPERTYKEY{GUID{Data1: 0xe55fc3b0, Data2: 0x2b60, Data3: 0x4220, Data4: [8]byte{0x91, 0x8e, 0xb2, 0x1e, 0x8b, 0xf1, 0x60, 0x16}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_UserName                                  = PROPERTYKEY{GUID{Data1: 0xc4322503, Data2: 0x78ca, Data3: 0x49c6, Data4: [8]byte{0x9a, 0xcc, 0xa6, 0x8e, 0x2a, 0xfd, 0x7b, 0x6b}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_IdentityProvider_Name                              = PROPERTYKEY{GUID{Data1: 0xb96eff7b, Data2: 0x35ca, Data3: 0x4a35, Data4: [8]byte{0x86, 0x07, 0x29, 0xe3, 0xa5, 0x4c, 0x46, 0xea}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_IdentityProvider_Picture                           = PROPERTYKEY{GUID{Data1: 0x2425166f, Data2: 0x5642, Data3: 0x4864, Data4: [8]byte{0x99, 0x2f, 0x98, 0xfd, 0x98, 0xf2, 0x94, 0xc3}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ImageParsingName                                   = PROPERTYKEY{GUID{Data1: 0xd7750ee0, Data2: 0xc6a4, Data3: 0x48ec, Data4: [8]byte{0xb5, 0x3e, 0xb8, 0x7b, 0x52, 0xe6, 0xd0, 0x73}}, 100}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Importance                                         = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 11}    // Int32 -- VT_I4
	PKEY_ImportanceText                                     = PROPERTYKEY{GUID{Data1: 0xa3b29791, Data2: 0x7713, Data3: 0x4e1d, Data4: [8]byte{0xbb, 0x40, 0x17, 0xdb, 0x85, 0xf0, 0x18, 0x31}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_IsAttachment                                       = PROPERTYKEY{GUID{Data1: 0xf23f425c, Data2: 0x71a1, Data3: 0x4fa8, Data4: [8]byte{0x92, 0x2f, 0x67, 0x8e, 0xa4, 0xa6, 0x04, 0x08}}, 100}   // Boolean -- VT_BOOL
	PKEY_IsDefaultNonOwnerSaveLocation                      = PROPERTYKEY{GUID{Data1: 0x5d76b67f, Data2: 0x9b3d, Data3: 0x44bb, Data4: [8]byte{0xb6, 0xae, 0x25, 0xda, 0x4f, 0x63, 0x8a, 0x67}}, 5}     // Boolean -- VT_BOOL
	PKEY_IsDefaultSaveLocation                              = PROPERTYKEY{GUID{Data1: 0x5d76b67f, Data2: 0x9b3d, Data3: 0x44bb, Data4: [8]byte{0xb6, 0xae, 0x25, 0xda, 0x4f, 0x63, 0x8a, 0x67}}, 3}     // Boolean -- VT_BOOL
	PKEY_IsDeleted                                          = PROPERTYKEY{GUID{Data1: 0x5cda5fc8, Data2: 0x33ee, Data3: 0x4ff3, Data4: [8]byte{0x90, 0x94, 0xae, 0x7b, 0xd8, 0x86, 0x8c, 0x4d}}, 100}   // Boolean -- VT_BOOL
	PKEY_IsEncrypted                                        = PROPERTYKEY{GUID{Data1: 0x90e5e14e, Data2: 0x648b, Data3: 0x4826, Data4: [8]byte{0xb2, 0xaa, 0xac, 0xaf, 0x79, 0x0e, 0x35, 0x13}}, 10}    // Boolean -- VT_BOOL
	PKEY_IsFlagged                                          = PROPERTYKEY{GUID{Data1: 0x5da84765, Data2: 0xe3ff, Data3: 0x4278, Data4: [8]byte{0x86, 0xb0, 0xa2, 0x79, 0x67, 0xfb, 0xdd, 0x03}}, 100}   // Boolean -- VT_BOOL
	PKEY_IsFlaggedComplete                                  = PROPERTYKEY{GUID{Data1: 0xa6f360d2, Data2: 0x55f9, Data3: 0x48de, Data4: [8]byte{0xb9, 0x09, 0x62, 0x0e, 0x09, 0x0a, 0x64, 0x7c}}, 100}   // Boolean -- VT_BOOL
	PKEY_IsIncomplete                                       = PROPERTYKEY{GUID{Data1: 0x346c8bd1, Data2: 0x2e6a, Data3: 0x4c45, Data4: [8]byte{0x89, 0xa4, 0x61, 0xb7, 0x8e, 0x8e, 0x70, 0x0f}}, 100}   // Boolean -- VT_BOOL
	PKEY_IsLocationSupported                                = PROPERTYKEY{GUID{Data1: 0x5d76b67f, Data2: 0x9b3d, Data3: 0x44bb, Data4: [8]byte{0xb6, 0xae, 0x25, 0xda, 0x4f, 0x63, 0x8a, 0x67}}, 8}     // Boolean -- VT_BOOL
	PKEY_IsPinnedToNameSpaceTree                            = PROPERTYKEY{GUID{Data1: 0x5d76b67f, Data2: 0x9b3d, Data3: 0x44bb, Data4: [8]byte{0xb6, 0xae, 0x25, 0xda, 0x4f, 0x63, 0x8a, 0x67}}, 2}     // Boolean -- VT_BOOL
	PKEY_IsRead                                             = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 10}    // Boolean -- VT_BOOL
	PKEY_IsSearchOnlyItem                                   = PROPERTYKEY{GUID{Data1: 0x5d76b67f, Data2: 0x9b3d, Data3: 0x44bb, Data4: [8]byte{0xb6, 0xae, 0x25, 0xda, 0x4f, 0x63, 0x8a, 0x67}}, 4}     // Boolean -- VT_BOOL
	PKEY_IsSendToTarget                                     = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 33}    // Boolean -- VT_BOOL
	PKEY_IsShared                                           = PROPERTYKEY{GUID{Data1: 0xef884c5b, Data2: 0x2bfe, Data3: 0x41bb, Data4: [8]byte{0xaa, 0xe5, 0x76, 0xee, 0xdf, 0x4f, 0x99, 0x02}}, 100}   // Boolean -- VT_BOOL
	PKEY_ItemAuthors                                        = PROPERTYKEY{GUID{Data1: 0xd0a04f0a, Data2: 0x462a, Data3: 0x48a4, Data4: [8]byte{0xbb, 0x2f, 0x37, 0x06, 0xe8, 0x8d, 0xbd, 0x7d}}, 100}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_ItemClassType                                      = PROPERTYKEY{GUID{Data1: 0x048658ad, Data2: 0x2db8, Data3: 0x41a4, Data4: [8]byte{0xbb, 0xb6, 0xac, 0x1e, 0xf1, 0x20, 0x7e, 0xb1}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemDate                                           = PROPERTYKEY{GUID{Data1: 0xf7db74b4, Data2: 0x4287, Data3: 0x4103, Data4: [8]byte{0xaf, 0xba, 0xf1, 0xb1, 0x3d, 0xcd, 0x75, 0xcf}}, 100}   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_ItemFolderNameDisplay                              = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 2}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemFolderPathDisplay                              = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 6}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemFolderPathDisplayNarrow                        = PROPERTYKEY{GUID{Data1: 0xdabd30ed, Data2: 0x0043, Data3: 0x4789, Data4: [8]byte{0xa7, 0xf8, 0xd0, 0x13, 0xa4, 0x73, 0x66, 0x22}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemName                                           = PROPERTYKEY{GUID{Data1: 0x6b8da074, Data2: 0x3b5c, Data3: 0x43bc, Data4: [8]byte{0x88, 0x6f, 0x0a, 0x2c, 0xdc, 0xe0, 0x0b, 0x6f}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemNameDisplay                                    = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 10}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemNameDisplayWithoutExtension                    = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 24}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemNamePrefix                                     = PROPERTYKEY{GUID{Data1: 0xd7313ff1, Data2: 0xa77a, Data3: 0x401c, Data4: [8]byte{0x8c, 0x99, 0x3d, 0xbd, 0xd6, 0x8a, 0xdd, 0x36}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemNameSortOverride                               = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 23}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemParticipants                                   = PROPERTYKEY{GUID{Data1: 0xd4d0aa16, Data2: 0x9948, Data3: 0x41a4, Data4: [8]byte{0xaa, 0x85, 0xd9, 0x7f, 0xf9, 0x64, 0x69, 0x93}}, 100}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_ItemPathDisplay                                    = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 7}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemPathDisplayNarrow                              = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 8}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemSubType                                        = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 37}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemType                                           = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 11}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemTypeText                                       = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 4}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemUrl                                            = PROPERTYKEY{GUID{Data1: 0x49691c90, Data2: 0x7e17, Data3: 0x101a, Data4: [8]byte{0xa9, 0x1c, 0x08, 0x00, 0x2b, 0x2e, 0xcd, 0xa9}}, 9}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Keywords                                           = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 5}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_Kind                                               = PROPERTYKEY{GUID{Data1: 0x1e3ee840, Data2: 0xbc2b, Data3: 0x476c, Data4: [8]byte{0x82, 0x37, 0x2a, 0xcd, 0x1a, 0x83, 0x9b, 0x22}}, 3}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_KindText                                           = PROPERTYKEY{GUID{Data1: 0xf04bef95, Data2: 0xc585, Data3: 0x4197, Data4: [8]byte{0xa2, 0xb7, 0xdf, 0x46, 0xfd, 0xc9, 0xee, 0x6d}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Language                                           = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 28}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_LastSyncError                                      = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 107}   // UInt32 -- VT_UI4
	PKEY_LastSyncWarning                                    = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 128}   // UInt32 -- VT_UI4
	PKEY_LastWriterPackageFamilyName                        = PROPERTYKEY{GUID{Data1: 0x502cfeab, Data2: 0x47eb, Data3: 0x459c, Data4: [8]byte{0xb9, 0x60, 0xe6, 0xd8, 0x72, 0x8f, 0x77, 0x01}}, 101}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_LowKeywords                                        = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 25}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_MediumKeywords                                     = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 26}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_MileageInformation                                 = PROPERTYKEY{GUID{Data1: 0xfdf84370, Data2: 0x031a, Data3: 0x4add, Data4: [8]byte{0x9e, 0x91, 0x0d, 0x77, 0x5f, 0x1c, 0x66, 0x05}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_MIMEType                                           = PROPERTYKEY{GUID{Data1: 0x0b63e350, Data2: 0x9ccc, Data3: 0x11d0, Data4: [8]byte{0xbc, 0xdb, 0x00, 0x80, 0x5f, 0xcc, 0xce, 0x04}}, 5}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Null                                               = PROPERTYKEY{GUID{Data1: 0x00000000, Data2: 0x0000, Data3: 0x0000, Data4: [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}, 0}     // Null -- VT_NULL
	PKEY_OfflineAvailability                                = PROPERTYKEY{GUID{Data1: 0xa94688b6, Data2: 0x7d9f, Data3: 0x4570, Data4: [8]byte{0xa6, 0x48, 0xe3, 0xdf, 0xc0, 0xab, 0x2b, 0x3f}}, 100}   // UInt32 -- VT_UI4
	PKEY_OfflineStatus                                      = PROPERTYKEY{GUID{Data1: 0x6d24888f, Data2: 0x4718, Data3: 0x4bda, Data4: [8]byte{0xaf, 0xed, 0xea, 0x0f, 0xb4, 0x38, 0x6c, 0xd8}}, 100}   // UInt32 -- VT_UI4
	PKEY_OriginalFileName                                   = PROPERTYKEY{GUID{Data1: 0x0cef7d53, Data2: 0xfa64, Data3: 0x11d1, Data4: [8]byte{0xa2, 0x03, 0x00, 0x00, 0xf8, 0x1f, 0xed, 0xee}}, 6}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_OwnerSID                                           = PROPERTYKEY{GUID{Data1: 0x5d76b67f, Data2: 0x9b3d, Data3: 0x44bb, Data4: [8]byte{0xb6, 0xae, 0x25, 0xda, 0x4f, 0x63, 0x8a, 0x67}}, 6}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParentalRating                                     = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 21}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParentalRatingReason                               = PROPERTYKEY{GUID{Data1: 0x10984e0a, Data2: 0xf9f2, Data3: 0x4321, Data4: [8]byte{0xb7, 0xef, 0xba, 0xf1, 0x95, 0xaf, 0x43, 0x19}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParentalRatingsOrganization                        = PROPERTYKEY{GUID{Data1: 0xa7fe0840, Data2: 0x1344, Data3: 0x46f0, Data4: [8]byte{0x8d, 0x37, 0x52, 0xed, 0x71, 0x2a, 0x4b, 0xf9}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParsingBindContext                                 = PROPERTYKEY{GUID{Data1: 0xdfb9a04d, Data2: 0x362f, Data3: 0x4ca3, Data4: [8]byte{0xb3, 0x0b, 0x02, 0x54, 0xb1, 0x7b, 0x5b, 0x84}}, 100}   // Any -- VT_NULL  Legacy code may treat this as VT_UNKNOWN.
	PKEY_ParsingName                                        = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 24}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParsingPath                                        = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 30}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PerceivedType                                      = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 9}     // Int32 -- VT_I4
	PKEY_PercentFull                                        = PROPERTYKEY{GUID{Data1: 0x9b174b35, Data2: 0x40ff, Data3: 0x11d2, Data4: [8]byte{0xa2, 0x7e, 0x00, 0xc0, 0x4f, 0xc3, 0x08, 0x71}}, 5}     // UInt32 -- VT_UI4
	PKEY_Priority                                           = PROPERTYKEY{GUID{Data1: 0x9c1fcf74, Data2: 0x2d97, Data3: 0x41ba, Data4: [8]byte{0xb4, 0xae, 0xcb, 0x2e, 0x36, 0x61, 0xa6, 0xe4}}, 5}     // UInt16 -- VT_UI2
	PKEY_PriorityText                                       = PROPERTYKEY{GUID{Data1: 0xd98be98b, Data2: 0xb86b, Data3: 0x4095, Data4: [8]byte{0xbf, 0x52, 0x9d, 0x23, 0xb2, 0xe0, 0xa7, 0x52}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Project                                            = PROPERTYKEY{GUID{Data1: 0x39a7f922, Data2: 0x477c, Data3: 0x48de, Data4: [8]byte{0x8b, 0xc8, 0xb2, 0x84, 0x41, 0xe3, 0x42, 0xe3}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ProviderItemID                                     = PROPERTYKEY{GUID{Data1: 0xf21d9941, Data2: 0x81f0, Data3: 0x471a, Data4: [8]byte{0xad, 0xee, 0x4e, 0x74, 0xb4, 0x92, 0x17, 0xed}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Rating                                             = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 9}     // UInt32 -- VT_UI4
	PKEY_RatingText                                         = PROPERTYKEY{GUID{Data1: 0x90197ca7, Data2: 0xfd8f, Data3: 0x4e8c, Data4: [8]byte{0x9d, 0xa3, 0xb5, 0x7e, 0x1e, 0x60, 0x92, 0x95}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RemoteConflictingFile                              = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 115}   // Object -- VT_UNKNOWN
	PKEY_Security_AllowedEnterpriseDataProtectionIdentities = PROPERTYKEY{GUID{Data1: 0x38d43380, Data2: 0xd418, Data3: 0x4830, Data4: [8]byte{0x84, 0xd5, 0x46, 0x93, 0x5a, 0x81, 0xc5, 0xc6}}, 32}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Security_EncryptionOwners                          = PROPERTYKEY{GUID{Data1: 0x5f5aff6a, Data2: 0x37e5, Data3: 0x4780, Data4: [8]byte{0x97, 0xea, 0x80, 0xc7, 0x56, 0x5c, 0xf5, 0x35}}, 34}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Security_EncryptionOwnersDisplay                   = PROPERTYKEY{GUID{Data1: 0xde621b8f, Data2: 0xe125, Data3: 0x43a3, Data4: [8]byte{0xa3, 0x2d, 0x56, 0x65, 0x44, 0x6d, 0x63, 0x2a}}, 25}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Sensitivity                                        = PROPERTYKEY{GUID{Data1: 0xf8d3f6ac, Data2: 0x4874, Data3: 0x42cb, Data4: [8]byte{0xbe, 0x59, 0xab, 0x45, 0x4b, 0x30, 0x71, 0x6a}}, 100}   // UInt16 -- VT_UI2
	PKEY_SensitivityText                                    = PROPERTYKEY{GUID{Data1: 0xd0c7f054, Data2: 0x3f72, Data3: 0x4725, Data4: [8]byte{0x85, 0x27, 0x12, 0x9a, 0x57, 0x7c, 0xb2, 0x69}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SFGAOFlags                                         = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 25}    // UInt32 -- VT_UI4
	PKEY_SharedWith                                         = PROPERTYKEY{GUID{Data1: 0xef884c5b, Data2: 0x2bfe, Data3: 0x41bb, Data4: [8]byte{0xaa, 0xe5, 0x76, 0xee, 0xdf, 0x4f, 0x99, 0x02}}, 200}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_ShareUserRating                                    = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 12}    // UInt32 -- VT_UI4
	PKEY_SharingStatus                                      = PROPERTYKEY{GUID{Data1: 0xef884c5b, Data2: 0x2bfe, Data3: 0x41bb, Data4: [8]byte{0xaa, 0xe5, 0x76, 0xee, 0xdf, 0x4f, 0x99, 0x02}}, 300}   // UInt32 -- VT_UI4
	PKEY_Shell_OmitFromView                                 = PROPERTYKEY{GUID{Data1: 0xde35258c, Data2: 0xc695, Data3: 0x4cbc, Data4: [8]byte{0xb9, 0x82, 0x38, 0xb0, 0xad, 0x24, 0xce, 0xd0}}, 2}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SimpleRating                                       = PROPERTYKEY{GUID{Data1: 0xa09f084e, Data2: 0xad41, Data3: 0x489f, Data4: [8]byte{0x80, 0x76, 0xaa, 0x5b, 0xe3, 0x08, 0x2b, 0xca}}, 100}   // UInt32 -- VT_UI4
	PKEY_Size                                               = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 12}    // UInt64 -- VT_UI8
	PKEY_SoftwareUsed                                       = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 305}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SourceItem                                         = PROPERTYKEY{GUID{Data1: 0x668cdfa5, Data2: 0x7a1b, Data3: 0x4323, Data4: [8]byte{0xae, 0x4b, 0xe5, 0x27, 0x39, 0x3a, 0x1d, 0x81}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SourcePackageFamilyName                            = PROPERTYKEY{GUID{Data1: 0xffae9db7, Data2: 0x1c8d, Data3: 0x43ff, Data4: [8]byte{0x81, 0x8c, 0x84, 0x40, 0x3a, 0xa3, 0x73, 0x2d}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StartDate                                          = PROPERTYKEY{GUID{Data1: 0x48fd6ec8, Data2: 0x8a12, Data3: 0x4cdf, Data4: [8]byte{0xa0, 0x3e, 0x4e, 0xc5, 0xa5, 0x11, 0xed, 0xde}}, 100}   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Status                                             = PROPERTYKEY{GUID{Data1: 0x000214a1, Data2: 0x0000, Data3: 0x0000, Data4: [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}, 9}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderCallerVersionInformation            = PROPERTYKEY{GUID{Data1: 0xb2f9b9d6, Data2: 0xfec4, Data3: 0x4dd5, Data4: [8]byte{0x94, 0xd7, 0x89, 0x57, 0x48, 0x8c, 0x80, 0x7b}}, 7}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderError                               = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 109}   // UInt32 -- VT_UI4
	PKEY_StorageProviderFileChecksum                        = PROPERTYKEY{GUID{Data1: 0xb2f9b9d6, Data2: 0xfec4, Data3: 0x4dd5, Data4: [8]byte{0x94, 0xd7, 0x89, 0x57, 0x48, 0x8c, 0x80, 0x7b}}, 5}     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_StorageProviderFileFlags                           = PROPERTYKEY{GUID{Data1: 0xb2f9b9d6, Data2: 0xfec4, Data3: 0x4dd5, Data4: [8]byte{0x94, 0xd7, 0x89, 0x57, 0x48, 0x8c, 0x80, 0x7b}}, 8}     // UInt32 -- VT_UI4
	PKEY_StorageProviderFileHasConflict                     = PROPERTYKEY{GUID{Data1: 0xb2f9b9d6, Data2: 0xfec4, Data3: 0x4dd5, Data4: [8]byte{0x94, 0xd7, 0x89, 0x57, 0x48, 0x8c, 0x80, 0x7b}}, 9}     // Boolean -- VT_BOOL
	PKEY_StorageProviderFileIdentifier                      = PROPERTYKEY{GUID{Data1: 0xb2f9b9d6, Data2: 0xfec4, Data3: 0x4dd5, Data4: [8]byte{0x94, 0xd7, 0x89, 0x57, 0x48, 0x8c, 0x80, 0x7b}}, 3}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderFileRemoteUri                       = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 112}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderFileVersion                         = PROPERTYKEY{GUID{Data1: 0xb2f9b9d6, Data2: 0xfec4, Data3: 0x4dd5, Data4: [8]byte{0x94, 0xd7, 0x89, 0x57, 0x48, 0x8c, 0x80, 0x7b}}, 4}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderFileVersionWaterline                = PROPERTYKEY{GUID{Data1: 0xb2f9b9d6, Data2: 0xfec4, Data3: 0x4dd5, Data4: [8]byte{0x94, 0xd7, 0x89, 0x57, 0x48, 0x8c, 0x80, 0x7b}}, 6}     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_StorageProviderId                                  = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 108}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderShareStatuses                       = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 111}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_StorageProviderSharingStatus                       = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 117}   // UInt32 -- VT_UI4
	PKEY_StorageProviderStatus                              = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 110}   // UInt64 -- VT_UI8
	PKEY_Subject                                            = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 3}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SyncTransferStatus                                 = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 103}   // UInt32 -- VT_UI4
	PKEY_Thumbnail                                          = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 17}    // Clipboard -- VT_CF
	PKEY_ThumbnailCacheId                                   = PROPERTYKEY{GUID{Data1: 0x446d16b1, Data2: 0x8dad, Data3: 0x4870, Data4: [8]byte{0xa7, 0x48, 0x40, 0x2e, 0xa4, 0x3d, 0x78, 0x8c}}, 100}   // UInt64 -- VT_UI8
	PKEY_ThumbnailStream                                    = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 27}    // Stream -- VT_STREAM
	PKEY_Title                                              = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 2}     // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_TitleSortOverride                                  = PROPERTYKEY{GUID{Data1: 0xf0f7984d, Data2: 0x222e, Data3: 0x4ad2, Data4: [8]byte{0x82, 0xab, 0x1d, 0xd8, 0xea, 0x40, 0xe5, 0x7e}}, 300}   // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_TotalFileSize                                      = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 14}    // UInt64 -- VT_UI8
	PKEY_Trademarks                                         = PROPERTYKEY{GUID{Data1: 0x0cef7d53, Data2: 0xfa64, Data3: 0x11d1, Data4: [8]byte{0xa2, 0x03, 0x00, 0x00, 0xf8, 0x1f, 0xed, 0xee}}, 9}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_TransferOrder                                      = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 106}   // UInt64 -- VT_UI8
	PKEY_TransferPosition                                   = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 104}   // UInt64 -- VT_UI8
	PKEY_TransferSize                                       = PROPERTYKEY{GUID{Data1: 0xfceff153, Data2: 0xe839, Data3: 0x4cf3, Data4: [8]byte{0xa9, 0xe7, 0xea, 0x22, 0x83, 0x20, 0x94, 0xb8}}, 105}   // UInt64 -- VT_UI8
	PKEY_VolumeId                                           = PROPERTYKEY{GUID{Data1: 0x446d16b1, Data2: 0x8dad, Data3: 0x4870, Data4: [8]byte{0xa7, 0x48, 0x40, 0x2e, 0xa4, 0x3d, 0x78, 0x8c}}, 104}   // Guid -- VT_CLSID
	PKEY_ZoneIdentifier                                     = PROPERTYKEY{GUID{Data1: 0x502cfeab, Data2: 0x47eb, Data3: 0x459c, Data4: [8]byte{0xb9, 0x60, 0xe6, 0xd8, 0x72, 0x8f, 0x77, 0x01}}, 100}   // UInt32 -- VT_UI4

	// Devices properties

	PKEY_Device_PrinterURL                                       = PROPERTYKEY{GUID{Data1: 0x0b48f35a, Data2: 0xbe6e, Data3: 0x4f17, Data4: [8]byte{0xb1, 0x08, 0x3c, 0x40, 0x73, 0xd1, 0x66, 0x9a}}, 15}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Bluetooth_DeviceAddress                 = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 1}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Bluetooth_Flags                         = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 3}     // UInt32 -- VT_UI4
	PKEY_DeviceInterface_Bluetooth_LastConnectedTime             = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 11}    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DeviceInterface_Bluetooth_Manufacturer                  = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 4}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Bluetooth_ModelNumber                   = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 5}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Bluetooth_ProductId                     = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 8}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Bluetooth_ProductVersion                = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 9}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Bluetooth_ServiceGuid                   = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 2}     // Guid -- VT_CLSID
	PKEY_DeviceInterface_Bluetooth_VendorId                      = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 7}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Bluetooth_VendorIdSource                = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 6}     // Byte -- VT_UI1
	PKEY_DeviceInterface_Hid_IsReadOnly                          = PROPERTYKEY{GUID{Data1: 0xcbf38310, Data2: 0x4a17, Data3: 0x4310, Data4: [8]byte{0xa1, 0xeb, 0x24, 0x7f, 0x0b, 0x67, 0x59, 0x3b}}, 4}     // Boolean -- VT_BOOL
	PKEY_DeviceInterface_Hid_ProductId                           = PROPERTYKEY{GUID{Data1: 0xcbf38310, Data2: 0x4a17, Data3: 0x4310, Data4: [8]byte{0xa1, 0xeb, 0x24, 0x7f, 0x0b, 0x67, 0x59, 0x3b}}, 6}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Hid_UsageId                             = PROPERTYKEY{GUID{Data1: 0xcbf38310, Data2: 0x4a17, Data3: 0x4310, Data4: [8]byte{0xa1, 0xeb, 0x24, 0x7f, 0x0b, 0x67, 0x59, 0x3b}}, 3}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Hid_UsagePage                           = PROPERTYKEY{GUID{Data1: 0xcbf38310, Data2: 0x4a17, Data3: 0x4310, Data4: [8]byte{0xa1, 0xeb, 0x24, 0x7f, 0x0b, 0x67, 0x59, 0x3b}}, 2}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Hid_VendorId                            = PROPERTYKEY{GUID{Data1: 0xcbf38310, Data2: 0x4a17, Data3: 0x4310, Data4: [8]byte{0xa1, 0xeb, 0x24, 0x7f, 0x0b, 0x67, 0x59, 0x3b}}, 5}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Hid_VersionNumber                       = PROPERTYKEY{GUID{Data1: 0xcbf38310, Data2: 0x4a17, Data3: 0x4310, Data4: [8]byte{0xa1, 0xeb, 0x24, 0x7f, 0x0b, 0x67, 0x59, 0x3b}}, 7}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_PrinterDriverDirectory                  = PROPERTYKEY{GUID{Data1: 0x847c66de, Data2: 0xb8d6, Data3: 0x4af9, Data4: [8]byte{0xab, 0xc3, 0x6f, 0x4f, 0x92, 0x6b, 0xc0, 0x39}}, 14}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_PrinterDriverName                       = PROPERTYKEY{GUID{Data1: 0xafc47170, Data2: 0x14f5, Data3: 0x498c, Data4: [8]byte{0x8f, 0x30, 0xb0, 0xd1, 0x9b, 0xe4, 0x49, 0xc6}}, 11}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_PrinterEnumerationFlag                  = PROPERTYKEY{GUID{Data1: 0xa00742a1, Data2: 0xcd8c, Data3: 0x4b37, Data4: [8]byte{0x95, 0xab, 0x70, 0x75, 0x55, 0x87, 0x76, 0x7a}}, 3}     // UInt32 -- VT_UI4
	PKEY_DeviceInterface_PrinterName                             = PROPERTYKEY{GUID{Data1: 0x0a7b84ef, Data2: 0x0c27, Data3: 0x463f, Data4: [8]byte{0x84, 0xef, 0x06, 0xc5, 0x07, 0x00, 0x01, 0xbe}}, 10}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_PrinterPortName                         = PROPERTYKEY{GUID{Data1: 0xeec7b761, Data2: 0x6f94, Data3: 0x41b1, Data4: [8]byte{0x94, 0x9f, 0xc7, 0x29, 0x72, 0x0d, 0xd1, 0x3c}}, 12}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Proximity_SupportsNfc                   = PROPERTYKEY{GUID{Data1: 0xfb3842cd, Data2: 0x9e2a, Data3: 0x4f83, Data4: [8]byte{0x8f, 0xcc, 0x4b, 0x07, 0x61, 0x13, 0x9a, 0xe9}}, 2}     // Boolean -- VT_BOOL
	PKEY_DeviceInterface_Serial_PortName                         = PROPERTYKEY{GUID{Data1: 0x4c6bf15c, Data2: 0x4c03, Data3: 0x4aac, Data4: [8]byte{0x91, 0xf5, 0x64, 0xc0, 0xf8, 0x52, 0xbc, 0xf4}}, 4}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Serial_UsbProductId                     = PROPERTYKEY{GUID{Data1: 0x4c6bf15c, Data2: 0x4c03, Data3: 0x4aac, Data4: [8]byte{0x91, 0xf5, 0x64, 0xc0, 0xf8, 0x52, 0xbc, 0xf4}}, 3}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Serial_UsbVendorId                      = PROPERTYKEY{GUID{Data1: 0x4c6bf15c, Data2: 0x4c03, Data3: 0x4aac, Data4: [8]byte{0x91, 0xf5, 0x64, 0xc0, 0xf8, 0x52, 0xbc, 0xf4}}, 2}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_WinUsb_DeviceInterfaceClasses           = PROPERTYKEY{GUID{Data1: 0x95e127b5, Data2: 0x79cc, Data3: 0x4e83, Data4: [8]byte{0x9c, 0x9e, 0x84, 0x22, 0x18, 0x7b, 0x3e, 0x0e}}, 7}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_DeviceInterface_WinUsb_UsbClass                         = PROPERTYKEY{GUID{Data1: 0x95e127b5, Data2: 0x79cc, Data3: 0x4e83, Data4: [8]byte{0x9c, 0x9e, 0x84, 0x22, 0x18, 0x7b, 0x3e, 0x0e}}, 4}     // Byte -- VT_UI1
	PKEY_DeviceInterface_WinUsb_UsbProductId                     = PROPERTYKEY{GUID{Data1: 0x95e127b5, Data2: 0x79cc, Data3: 0x4e83, Data4: [8]byte{0x9c, 0x9e, 0x84, 0x22, 0x18, 0x7b, 0x3e, 0x0e}}, 3}     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_WinUsb_UsbProtocol                      = PROPERTYKEY{GUID{Data1: 0x95e127b5, Data2: 0x79cc, Data3: 0x4e83, Data4: [8]byte{0x9c, 0x9e, 0x84, 0x22, 0x18, 0x7b, 0x3e, 0x0e}}, 6}     // Byte -- VT_UI1
	PKEY_DeviceInterface_WinUsb_UsbSubClass                      = PROPERTYKEY{GUID{Data1: 0x95e127b5, Data2: 0x79cc, Data3: 0x4e83, Data4: [8]byte{0x9c, 0x9e, 0x84, 0x22, 0x18, 0x7b, 0x3e, 0x0e}}, 5}     // Byte -- VT_UI1
	PKEY_DeviceInterface_WinUsb_UsbVendorId                      = PROPERTYKEY{GUID{Data1: 0x95e127b5, Data2: 0x79cc, Data3: 0x4e83, Data4: [8]byte{0x9c, 0x9e, 0x84, 0x22, 0x18, 0x7b, 0x3e, 0x0e}}, 2}     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_AepId                                       = PROPERTYKEY{GUID{Data1: 0x3b2ce006, Data2: 0x5e61, Data3: 0x4fde, Data4: [8]byte{0xba, 0xb8, 0x9b, 0x8a, 0xac, 0x9b, 0x26, 0xdf}}, 8}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Aep_Bluetooth_Cod_Major                         = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 2}     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Cod_Minor                         = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 3}     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Audio                = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 10}    // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Capturing            = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 8}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Information          = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 12}    // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_LimitedDiscovery     = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 4}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Networking           = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 6}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_ObjectXfer           = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 9}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Positioning          = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 5}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Rendering            = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 7}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Telephony            = PROPERTYKEY{GUID{Data1: 0x5fbd34cd, Data2: 0x561a, Data3: 0x412e, Data4: [8]byte{0xba, 0x98, 0x47, 0x8a, 0x6b, 0x0f, 0xef, 0x1d}}, 11}    // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_LastSeenTime                      = PROPERTYKEY{GUID{Data1: 0x2bd67d8b, Data2: 0x8beb, Data3: 0x48d5, Data4: [8]byte{0x87, 0xe0, 0x6c, 0xda, 0x34, 0x28, 0x04, 0x0a}}, 12}    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Devices_Aep_Bluetooth_Le_AddressType                    = PROPERTYKEY{GUID{Data1: 0x995ef0b0, Data2: 0x7eb3, Data3: 0x4a8b, Data4: [8]byte{0xb9, 0xce, 0x06, 0x8b, 0xb3, 0xf4, 0xaf, 0x69}}, 4}     // Byte -- VT_UI1
	PKEY_Devices_Aep_Bluetooth_Le_Appearance                     = PROPERTYKEY{GUID{Data1: 0x995ef0b0, Data2: 0x7eb3, Data3: 0x4a8b, Data4: [8]byte{0xb9, 0xce, 0x06, 0x8b, 0xb3, 0xf4, 0xaf, 0x69}}, 1}     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Le_Appearance_Category            = PROPERTYKEY{GUID{Data1: 0x995ef0b0, Data2: 0x7eb3, Data3: 0x4a8b, Data4: [8]byte{0xb9, 0xce, 0x06, 0x8b, 0xb3, 0xf4, 0xaf, 0x69}}, 5}     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Le_Appearance_Subcategory         = PROPERTYKEY{GUID{Data1: 0x995ef0b0, Data2: 0x7eb3, Data3: 0x4a8b, Data4: [8]byte{0xb9, 0xce, 0x06, 0x8b, 0xb3, 0xf4, 0xaf, 0x69}}, 6}     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Le_IsConnectable                  = PROPERTYKEY{GUID{Data1: 0x995ef0b0, Data2: 0x7eb3, Data3: 0x4a8b, Data4: [8]byte{0xb9, 0xce, 0x06, 0x8b, 0xb3, 0xf4, 0xaf, 0x69}}, 8}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_CanPair                                     = PROPERTYKEY{GUID{Data1: 0xe7c3fb29, Data2: 0xcaa7, Data3: 0x4f47, Data4: [8]byte{0x8c, 0x8b, 0xbe, 0x59, 0xb3, 0x30, 0xd4, 0xc5}}, 3}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Category                                    = PROPERTYKEY{GUID{Data1: 0xa35996ab, Data2: 0x11cf, Data3: 0x4935, Data4: [8]byte{0x8b, 0x61, 0xa6, 0x76, 0x10, 0x81, 0xec, 0xdf}}, 17}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Aep_ContainerId                                 = PROPERTYKEY{GUID{Data1: 0xe7c3fb29, Data2: 0xcaa7, Data3: 0x4f47, Data4: [8]byte{0x8c, 0x8b, 0xbe, 0x59, 0xb3, 0x30, 0xd4, 0xc5}}, 2}     // Guid -- VT_CLSID
	PKEY_Devices_Aep_DeviceAddress                               = PROPERTYKEY{GUID{Data1: 0xa35996ab, Data2: 0x11cf, Data3: 0x4935, Data4: [8]byte{0x8b, 0x61, 0xa6, 0x76, 0x10, 0x81, 0xec, 0xdf}}, 12}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Aep_IsConnected                                 = PROPERTYKEY{GUID{Data1: 0xa35996ab, Data2: 0x11cf, Data3: 0x4935, Data4: [8]byte{0x8b, 0x61, 0xa6, 0x76, 0x10, 0x81, 0xec, 0xdf}}, 7}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_IsPaired                                    = PROPERTYKEY{GUID{Data1: 0xa35996ab, Data2: 0x11cf, Data3: 0x4935, Data4: [8]byte{0x8b, 0x61, 0xa6, 0x76, 0x10, 0x81, 0xec, 0xdf}}, 16}    // Boolean -- VT_BOOL
	PKEY_Devices_Aep_IsPresent                                   = PROPERTYKEY{GUID{Data1: 0xa35996ab, Data2: 0x11cf, Data3: 0x4935, Data4: [8]byte{0x8b, 0x61, 0xa6, 0x76, 0x10, 0x81, 0xec, 0xdf}}, 9}     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Manufacturer                                = PROPERTYKEY{GUID{Data1: 0xa35996ab, Data2: 0x11cf, Data3: 0x4935, Data4: [8]byte{0x8b, 0x61, 0xa6, 0x76, 0x10, 0x81, 0xec, 0xdf}}, 5}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Aep_ModelId                                     = PROPERTYKEY{GUID{Data1: 0xa35996ab, Data2: 0x11cf, Data3: 0x4935, Data4: [8]byte{0x8b, 0x61, 0xa6, 0x76, 0x10, 0x81, 0xec, 0xdf}}, 4}     // Guid -- VT_CLSID
	PKEY_Devices_Aep_ModelName                                   = PROPERTYKEY{GUID{Data1: 0xa35996ab, Data2: 0x11cf, Data3: 0x4935, Data4: [8]byte{0x8b, 0x61, 0xa6, 0x76, 0x10, 0x81, 0xec, 0xdf}}, 3}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Aep_PointOfService_ConnectionTypes              = PROPERTYKEY{GUID{Data1: 0xd4bf61b3, Data2: 0x442e, Data3: 0x4ada, Data4: [8]byte{0x88, 0x2d, 0xfa, 0x7b, 0x70, 0xc8, 0x32, 0xd9}}, 6}     // Int32 -- VT_I4
	PKEY_Devices_Aep_ProtocolId                                  = PROPERTYKEY{GUID{Data1: 0x3b2ce006, Data2: 0x5e61, Data3: 0x4fde, Data4: [8]byte{0xba, 0xb8, 0x9b, 0x8a, 0xac, 0x9b, 0x26, 0xdf}}, 5}     // Guid -- VT_CLSID
	PKEY_Devices_Aep_SignalStrength                              = PROPERTYKEY{GUID{Data1: 0xa35996ab, Data2: 0x11cf, Data3: 0x4935, Data4: [8]byte{0x8b, 0x61, 0xa6, 0x76, 0x10, 0x81, 0xec, 0xdf}}, 6}     // Int32 -- VT_I4
	PKEY_Devices_AepContainer_CanPair                            = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 3}     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_Categories                         = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 9}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_Children                           = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 2}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_ContainerId                        = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 12}    // Guid -- VT_CLSID
	PKEY_Devices_AepContainer_DialProtocol_InstalledApplications = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 6}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_IsPaired                           = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 4}     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_IsPresent                          = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 11}    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_Manufacturer                       = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 6}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AepContainer_ModelIds                           = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 8}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_ModelName                          = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 7}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AepContainer_ProtocolIds                        = PROPERTYKEY{GUID{Data1: 0x0bba1ede, Data2: 0x7566, Data3: 0x4f47, Data4: [8]byte{0x90, 0xec, 0x25, 0xfc, 0x56, 0x7c, 0xed, 0x2a}}, 13}    // Multivalue Guid -- VT_VECTOR | VT_CLSID  (For variants: VT_ARRAY | VT_CLSID)
	PKEY_Devices_AepContainer_SupportedUriSchemes                = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 5}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_SupportsAudio                      = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 2}     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsCapturing                  = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 11}    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsImages                     = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 4}     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsInformation                = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 14}    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsLimitedDiscovery           = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 7}     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsNetworking                 = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 9}     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsObjectTransfer             = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 12}    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsPositioning                = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 8}     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsRendering                  = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 10}    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsTelephony                  = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 13}    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsVideo                      = PROPERTYKEY{GUID{Data1: 0x6af55d45, Data2: 0x38db, Data3: 0x4495, Data4: [8]byte{0xac, 0xb0, 0xd4, 0x72, 0x8a, 0x3b, 0x83, 0x14}}, 3}     // Boolean -- VT_BOOL
	PKEY_Devices_AepService_AepId                                = PROPERTYKEY{GUID{Data1: 0xc9c141a9, Data2: 0x1b4c, Data3: 0x4f17, Data4: [8]byte{0xa9, 0xd1, 0xf2, 0x98, 0x53, 0x8c, 0xad, 0xb8}}, 6}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AepService_Bluetooth_CacheMode                  = PROPERTYKEY{GUID{Data1: 0x9744311e, Data2: 0x7951, Data3: 0x4b2e, Data4: [8]byte{0xb6, 0xf0, 0xec, 0xb2, 0x93, 0xca, 0xc1, 0x19}}, 5}     // Byte -- VT_UI1
	PKEY_Devices_AepService_Bluetooth_ServiceGuid                = PROPERTYKEY{GUID{Data1: 0xa399aac7, Data2: 0xc265, Data3: 0x474e, Data4: [8]byte{0xb0, 0x73, 0xff, 0xce, 0x57, 0x72, 0x17, 0x16}}, 2}     // Guid -- VT_CLSID
	PKEY_Devices_AepService_Bluetooth_TargetDevice               = PROPERTYKEY{GUID{Data1: 0x9744311e, Data2: 0x7951, Data3: 0x4b2e, Data4: [8]byte{0xb6, 0xf0, 0xec, 0xb2, 0x93, 0xca, 0xc1, 0x19}}, 6}     // UInt64 -- VT_UI8
	PKEY_Devices_AepService_ContainerId                          = PROPERTYKEY{GUID{Data1: 0x71724756, Data2: 0x3e74, Data3: 0x4432, Data4: [8]byte{0x9b, 0x59, 0xe7, 0xb2, 0xf6, 0x68, 0xa5, 0x93}}, 4}     // Guid -- VT_CLSID
	PKEY_Devices_AepService_FriendlyName                         = PROPERTYKEY{GUID{Data1: 0x71724756, Data2: 0x3e74, Data3: 0x4432, Data4: [8]byte{0x9b, 0x59, 0xe7, 0xb2, 0xf6, 0x68, 0xa5, 0x93}}, 2}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AepService_IoT_ServiceInterfaces                = PROPERTYKEY{GUID{Data1: 0x79d94e82, Data2: 0x4d79, Data3: 0x45aa, Data4: [8]byte{0x82, 0x1a, 0x74, 0x85, 0x8b, 0x4e, 0x4c, 0xa6}}, 2}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepService_ParentAepIsPaired                    = PROPERTYKEY{GUID{Data1: 0xc9c141a9, Data2: 0x1b4c, Data3: 0x4f17, Data4: [8]byte{0xa9, 0xd1, 0xf2, 0x98, 0x53, 0x8c, 0xad, 0xb8}}, 7}     // Boolean -- VT_BOOL
	PKEY_Devices_AepService_ProtocolId                           = PROPERTYKEY{GUID{Data1: 0xc9c141a9, Data2: 0x1b4c, Data3: 0x4f17, Data4: [8]byte{0xa9, 0xd1, 0xf2, 0x98, 0x53, 0x8c, 0xad, 0xb8}}, 5}     // Guid -- VT_CLSID
	PKEY_Devices_AepService_ServiceClassId                       = PROPERTYKEY{GUID{Data1: 0x71724756, Data2: 0x3e74, Data3: 0x4432, Data4: [8]byte{0x9b, 0x59, 0xe7, 0xb2, 0xf6, 0x68, 0xa5, 0x93}}, 3}     // Guid -- VT_CLSID
	PKEY_Devices_AepService_ServiceId                            = PROPERTYKEY{GUID{Data1: 0xc9c141a9, Data2: 0x1b4c, Data3: 0x4f17, Data4: [8]byte{0xa9, 0xd1, 0xf2, 0x98, 0x53, 0x8c, 0xad, 0xb8}}, 2}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AppPackageFamilyName                            = PROPERTYKEY{GUID{Data1: 0x51236583, Data2: 0x0c4a, Data3: 0x4fe8, Data4: [8]byte{0xb8, 0x1f, 0x16, 0x6a, 0xec, 0x13, 0xf5, 0x10}}, 100}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AudioDevice_Microphone_IsFarField               = PROPERTYKEY{GUID{Data1: 0x8943b373, Data2: 0x388c, Data3: 0x4395, Data4: [8]byte{0xb5, 0x57, 0xbc, 0x6d, 0xba, 0xff, 0xaf, 0xdb}}, 6}     // Boolean -- VT_BOOL
	PKEY_Devices_AudioDevice_Microphone_SensitivityInDbfs        = PROPERTYKEY{GUID{Data1: 0x8943b373, Data2: 0x388c, Data3: 0x4395, Data4: [8]byte{0xb5, 0x57, 0xbc, 0x6d, 0xba, 0xff, 0xaf, 0xdb}}, 3}     // Double -- VT_R8
	PKEY_Devices_AudioDevice_Microphone_SensitivityInDbfs2       = PROPERTYKEY{GUID{Data1: 0x8943b373, Data2: 0x388c, Data3: 0x4395, Data4: [8]byte{0xb5, 0x57, 0xbc, 0x6d, 0xba, 0xff, 0xaf, 0xdb}}, 5}     // Double -- VT_R8
	PKEY_Devices_AudioDevice_Microphone_SignalToNoiseRatioInDb   = PROPERTYKEY{GUID{Data1: 0x8943b373, Data2: 0x388c, Data3: 0x4395, Data4: [8]byte{0xb5, 0x57, 0xbc, 0x6d, 0xba, 0xff, 0xaf, 0xdb}}, 4}     // Double -- VT_R8
	PKEY_Devices_AudioDevice_RawProcessingSupported              = PROPERTYKEY{GUID{Data1: 0x8943b373, Data2: 0x388c, Data3: 0x4395, Data4: [8]byte{0xb5, 0x57, 0xbc, 0x6d, 0xba, 0xff, 0xaf, 0xdb}}, 2}     // Boolean -- VT_BOOL
	PKEY_Devices_AudioDevice_SpeechProcessingSupported           = PROPERTYKEY{GUID{Data1: 0xfb1de864, Data2: 0xe06d, Data3: 0x47f4, Data4: [8]byte{0x82, 0xa6, 0x8a, 0x0a, 0xef, 0x44, 0x49, 0x3c}}, 2}     // Boolean -- VT_BOOL
	PKEY_Devices_BatteryLife                                     = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 10}    // Byte -- VT_UI1
	PKEY_Devices_BatteryPlusCharging                             = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 22}    // Byte -- VT_UI1
	PKEY_Devices_BatteryPlusChargingText                         = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 23}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Category                                        = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 91}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_CategoryGroup                                   = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 94}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_CategoryIds                                     = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 90}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_CategoryPlural                                  = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 92}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_ChallengeAep                                    = PROPERTYKEY{GUID{Data1: 0x0774315e, Data2: 0xb714, Data3: 0x48ec, Data4: [8]byte{0x8d, 0xe8, 0x81, 0x25, 0xc0, 0x77, 0xac, 0x11}}, 2}     // Boolean -- VT_BOOL
	PKEY_Devices_ChargingState                                   = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 11}    // Byte -- VT_UI1
	PKEY_Devices_Children                                        = PROPERTYKEY{GUID{Data1: 0x4340a6c5, Data2: 0x93fa, Data3: 0x4706, Data4: [8]byte{0x97, 0x2c, 0x7b, 0x64, 0x80, 0x08, 0xa5, 0xa7}}, 9}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_ClassGuid                                       = PROPERTYKEY{GUID{Data1: 0xa45c254e, Data2: 0xdf1c, Data3: 0x4efd, Data4: [8]byte{0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0}}, 10}    // Guid -- VT_CLSID
	PKEY_Devices_CompatibleIds                                   = PROPERTYKEY{GUID{Data1: 0xa45c254e, Data2: 0xdf1c, Data3: 0x4efd, Data4: [8]byte{0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0}}, 4}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Connected                                       = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 55}    // Boolean -- VT_BOOL
	PKEY_Devices_ContainerId                                     = PROPERTYKEY{GUID{Data1: 0x8c7ed206, Data2: 0x3f8a, Data3: 0x4827, Data4: [8]byte{0xb3, 0xab, 0xae, 0x9e, 0x1f, 0xae, 0xfc, 0x6c}}, 2}     // Guid -- VT_CLSID
	PKEY_Devices_DefaultTooltip                                  = PROPERTYKEY{GUID{Data1: 0x880f70a2, Data2: 0x6082, Data3: 0x47ac, Data4: [8]byte{0x8a, 0xab, 0xa7, 0x39, 0xd1, 0xa3, 0x00, 0xc3}}, 153}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DeviceCapabilities                              = PROPERTYKEY{GUID{Data1: 0xa45c254e, Data2: 0xdf1c, Data3: 0x4efd, Data4: [8]byte{0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0}}, 17}    // UInt32 -- VT_UI4
	PKEY_Devices_DeviceCharacteristics                           = PROPERTYKEY{GUID{Data1: 0xa45c254e, Data2: 0xdf1c, Data3: 0x4efd, Data4: [8]byte{0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0}}, 29}    // UInt32 -- VT_UI4
	PKEY_Devices_DeviceDescription1                              = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 81}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DeviceDescription2                              = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 82}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DeviceHasProblem                                = PROPERTYKEY{GUID{Data1: 0x540b947e, Data2: 0x8b40, Data3: 0x45bc, Data4: [8]byte{0xa8, 0xa2, 0x6a, 0x0b, 0x89, 0x4c, 0xbd, 0xa2}}, 6}     // Boolean -- VT_BOOL
	PKEY_Devices_DeviceInstanceId                                = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 256}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DeviceManufacturer                              = PROPERTYKEY{GUID{Data1: 0xa45c254e, Data2: 0xdf1c, Data3: 0x4efd, Data4: [8]byte{0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0}}, 13}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DevObjectType                                   = PROPERTYKEY{GUID{Data1: 0x13673f42, Data2: 0xa3d6, Data3: 0x49f6, Data4: [8]byte{0xb4, 0xda, 0xae, 0x46, 0xe0, 0xc5, 0x23, 0x7c}}, 2}     // UInt32 -- VT_UI4
	PKEY_Devices_DialProtocol_InstalledApplications              = PROPERTYKEY{GUID{Data1: 0x6845cc72, Data2: 0x1b71, Data3: 0x48c3, Data4: [8]byte{0xaf, 0x86, 0xb0, 0x91, 0x71, 0xa1, 0x9b, 0x14}}, 3}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_DiscoveryMethod                                 = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 52}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Dnssd_Domain                                    = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 3}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_FullName                                  = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 5}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_HostName                                  = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 7}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_InstanceName                              = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 4}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_NetworkAdapterId                          = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 11}    // Guid -- VT_CLSID
	PKEY_Devices_Dnssd_PortNumber                                = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 12}    // UInt16 -- VT_UI2
	PKEY_Devices_Dnssd_Priority                                  = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 9}     // UInt16 -- VT_UI2
	PKEY_Devices_Dnssd_ServiceName                               = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 2}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_TextAttributes                            = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 6}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Dnssd_Ttl                                       = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 10}    // UInt32 -- VT_UI4
	PKEY_Devices_Dnssd_Weight                                    = PROPERTYKEY{GUID{Data1: 0xbf79c0ab, Data2: 0xbb74, Data3: 0x4cee, Data4: [8]byte{0xb0, 0x70, 0x47, 0x0b, 0x5a, 0xe2, 0x02, 0xea}}, 8}     // UInt16 -- VT_UI2
	PKEY_Devices_FriendlyName                                    = PROPERTYKEY{GUID{Data1: 0x656a3bb3, Data2: 0xecc0, Data3: 0x43fd, Data4: [8]byte{0x84, 0x77, 0x4a, 0xe0, 0x40, 0x4a, 0x96, 0xcd}}, 12288} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_FunctionPaths                                   = PROPERTYKEY{GUID{Data1: 0xd08dd4c0, Data2: 0x3a9e, Data3: 0x462e, Data4: [8]byte{0x82, 0x90, 0x7b, 0x63, 0x6b, 0x25, 0x76, 0xb9}}, 3}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_GlyphIcon                                       = PROPERTYKEY{GUID{Data1: 0x51236583, Data2: 0x0c4a, Data3: 0x4fe8, Data4: [8]byte{0xb8, 0x1f, 0x16, 0x6a, 0xec, 0x13, 0xf5, 0x10}}, 123}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_HardwareIds                                     = PROPERTYKEY{GUID{Data1: 0xa45c254e, Data2: 0xdf1c, Data3: 0x4efd, Data4: [8]byte{0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0}}, 3}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Icon                                            = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 57}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_InLocalMachineContainer                         = PROPERTYKEY{GUID{Data1: 0x8c7ed206, Data2: 0x3f8a, Data3: 0x4827, Data4: [8]byte{0xb3, 0xab, 0xae, 0x9e, 0x1f, 0xae, 0xfc, 0x6c}}, 4}     // Boolean -- VT_BOOL
	PKEY_Devices_InterfaceClassGuid                              = PROPERTYKEY{GUID{Data1: 0x026e516e, Data2: 0xb814, Data3: 0x414b, Data4: [8]byte{0x83, 0xcd, 0x85, 0x6d, 0x6f, 0xef, 0x48, 0x22}}, 4}     // Guid -- VT_CLSID
	PKEY_Devices_InterfaceEnabled                                = PROPERTYKEY{GUID{Data1: 0x026e516e, Data2: 0xb814, Data3: 0x414b, Data4: [8]byte{0x83, 0xcd, 0x85, 0x6d, 0x6f, 0xef, 0x48, 0x22}}, 3}     // Boolean -- VT_BOOL
	PKEY_Devices_InterfacePaths                                  = PROPERTYKEY{GUID{Data1: 0xd08dd4c0, Data2: 0x3a9e, Data3: 0x462e, Data4: [8]byte{0x82, 0x90, 0x7b, 0x63, 0x6b, 0x25, 0x76, 0xb9}}, 2}     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_IpAddress                                       = PROPERTYKEY{GUID{Data1: 0x656a3bb3, Data2: 0xecc0, Data3: 0x43fd, Data4: [8]byte{0x84, 0x77, 0x4a, 0xe0, 0x40, 0x4a, 0x96, 0xcd}}, 12297} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_IsDefault                                       = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 86}    // Boolean -- VT_BOOL
	PKEY_Devices_IsNetworkConnected                              = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 85}    // Boolean -- VT_BOOL
	PKEY_Devices_IsShared                                        = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 84}    // Boolean -- VT_BOOL
	PKEY_Devices_IsSoftwareInstalling                            = PROPERTYKEY{GUID{Data1: 0x83da6326, Data2: 0x97a6, Data3: 0x4088, Data4: [8]byte{0x94, 0x53, 0xa1, 0x92, 0x3f, 0x57, 0x3b, 0x29}}, 9}     // Boolean -- VT_BOOL
	PKEY_Devices_LaunchDeviceStageFromExplorer                   = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 77}    // Boolean -- VT_BOOL
	PKEY_Devices_LocalMachine                                    = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 70}    // Boolean -- VT_BOOL
	PKEY_Devices_LocationPaths                                   = PROPERTYKEY{GUID{Data1: 0xa45c254e, Data2: 0xdf1c, Data3: 0x4efd, Data4: [8]byte{0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0}}, 37}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Manufacturer                                    = PROPERTYKEY{GUID{Data1: 0x656a3bb3, Data2: 0xecc0, Data3: 0x43fd, Data4: [8]byte{0x84, 0x77, 0x4a, 0xe0, 0x40, 0x4a, 0x96, 0xcd}}, 8192}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_MetadataPath                                    = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 71}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_MicrophoneArray_Geometry                        = PROPERTYKEY{GUID{Data1: 0xa1829ea2, Data2: 0x27eb, Data3: 0x459e, Data4: [8]byte{0x93, 0x5d, 0xb2, 0xfa, 0xd7, 0xb0, 0x77, 0x62}}, 2}     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_MissedCalls                                     = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 5}     // Byte -- VT_UI1
	PKEY_Devices_ModelId                                         = PROPERTYKEY{GUID{Data1: 0x80d81ea6, Data2: 0x7473, Data3: 0x4b0c, Data4: [8]byte{0x82, 0x16, 0xef, 0xc1, 0x1a, 0x2c, 0x4c, 0x8b}}, 2}     // Guid -- VT_CLSID
	PKEY_Devices_ModelName                                       = PROPERTYKEY{GUID{Data1: 0x656a3bb3, Data2: 0xecc0, Data3: 0x43fd, Data4: [8]byte{0x84, 0x77, 0x4a, 0xe0, 0x40, 0x4a, 0x96, 0xcd}}, 8194}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_ModelNumber                                     = PROPERTYKEY{GUID{Data1: 0x656a3bb3, Data2: 0xecc0, Data3: 0x43fd, Data4: [8]byte{0x84, 0x77, 0x4a, 0xe0, 0x40, 0x4a, 0x96, 0xcd}}, 8195}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_NetworkedTooltip                                = PROPERTYKEY{GUID{Data1: 0x880f70a2, Data2: 0x6082, Data3: 0x47ac, Data4: [8]byte{0x8a, 0xab, 0xa7, 0x39, 0xd1, 0xa3, 0x00, 0xc3}}, 152}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_NetworkName                                     = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 7}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_NetworkType                                     = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 8}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_NewPictures                                     = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 4}     // UInt16 -- VT_UI2
	PKEY_Devices_Notification                                    = PROPERTYKEY{GUID{Data1: 0x06704b0c, Data2: 0xe830, Data3: 0x4c81, Data4: [8]byte{0x91, 0x78, 0x91, 0xe4, 0xe9, 0x5a, 0x80, 0xa0}}, 3}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Notifications_LowBattery                        = PROPERTYKEY{GUID{Data1: 0xc4c07f2b, Data2: 0x8524, Data3: 0x4e66, Data4: [8]byte{0xae, 0x3a, 0xa6, 0x23, 0x5f, 0x10, 0x3b, 0xeb}}, 2}     // Byte -- VT_UI1
	PKEY_Devices_Notifications_MissedCall                        = PROPERTYKEY{GUID{Data1: 0x6614ef48, Data2: 0x4efe, Data3: 0x4424, Data4: [8]byte{0x9e, 0xda, 0xc7, 0x9f, 0x40, 0x4e, 0xdf, 0x3e}}, 2}     // Byte -- VT_UI1
	PKEY_Devices_Notifications_NewMessage                        = PROPERTYKEY{GUID{Data1: 0x2be9260a, Data2: 0x2012, Data3: 0x4742, Data4: [8]byte{0xa5, 0x55, 0xf4, 0x1b, 0x63, 0x8b, 0x7d, 0xcb}}, 2}     // Byte -- VT_UI1
	PKEY_Devices_Notifications_NewVoicemail                      = PROPERTYKEY{GUID{Data1: 0x59569556, Data2: 0x0a08, Data3: 0x4212, Data4: [8]byte{0x95, 0xb9, 0xfa, 0xe2, 0xad, 0x64, 0x13, 0xdb}}, 2}     // Byte -- VT_UI1
	PKEY_Devices_Notifications_StorageFull                       = PROPERTYKEY{GUID{Data1: 0xa0e00ee1, Data2: 0xf0c7, Data3: 0x4d41, Data4: [8]byte{0xb8, 0xe7, 0x26, 0xa7, 0xbd, 0x8d, 0x38, 0xb0}}, 2}     // UInt64 -- VT_UI8
	PKEY_Devices_Notifications_StorageFullLinkText               = PROPERTYKEY{GUID{Data1: 0xa0e00ee1, Data2: 0xf0c7, Data3: 0x4d41, Data4: [8]byte{0xb8, 0xe7, 0x26, 0xa7, 0xbd, 0x8d, 0x38, 0xb0}}, 3}     // UInt64 -- VT_UI8
	PKEY_Devices_NotificationStore                               = PROPERTYKEY{GUID{Data1: 0x06704b0c, Data2: 0xe830, Data3: 0x4c81, Data4: [8]byte{0x91, 0x78, 0x91, 0xe4, 0xe9, 0x5a, 0x80, 0xa0}}, 2}     // Object -- VT_UNKNOWN
	PKEY_Devices_NotWorkingProperly                              = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 83}    // Boolean -- VT_BOOL
	PKEY_Devices_Paired                                          = PROPERTYKEY{GUID{Data1: 0x78c34fc8, Data2: 0x104a, Data3: 0x4aca, Data4: [8]byte{0x9e, 0xa4, 0x52, 0x4d, 0x52, 0x99, 0x6e, 0x57}}, 56}    // Boolean -- VT_BOOL
	PKEY_Devices_Panel_PanelGroup                                = PROPERTYKEY{GUID{Data1: 0x8dbc9c86, Data2: 0x97a9, Data3: 0x4bff, Data4: [8]byte{0x9b, 0xc6, 0xbf, 0xe9, 0x5d, 0x3e, 0x6d, 0xad}}, 3}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Panel_PanelId                                   = PROPERTYKEY{GUID{Data1: 0x8dbc9c86, Data2: 0x97a9, Data3: 0x4bff, Data4: [8]byte{0x9b, 0xc6, 0xbf, 0xe9, 0x5d, 0x3e, 0x6d, 0xad}}, 2}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Parent                                          = PROPERTYKEY{GUID{Data1: 0x4340a6c5, Data2: 0x93fa, Data3: 0x4706, Data4: [8]byte{0x97, 0x2c, 0x7b, 0x64, 0x80, 0x08, 0xa5, 0xa7}}, 8}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_PhoneLineTransportDevice_Connected              = PROPERTYKEY{GUID{Data1: 0xaecf2fe8, Data2: 0x1d00, Data3: 0x4fee, Data4: [8]byte{0x8a, 0x6d, 0xa7, 0x0d, 0x71, 0x9b, 0x77, 0x2b}}, 2}     // Boolean -- VT_BOOL
	PKEY_Devices_PhysicalDeviceLocation                          = PROPERTYKEY{GUID{Data1: 0x540b947e, Data2: 0x8b40, Data3: 0x45bc, Data4: [8]byte{0xa8, 0xa2, 0x6a, 0x0b, 0x89, 0x4c, 0xbd, 0xa2}}, 9}     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_PlaybackPositionPercent                         = PROPERTYKEY{GUID{Data1: 0x3633de59, Data2: 0x6825, Data3: 0x4381, Data4: [8]byte{0xa4, 0x9b, 0x9f, 0x6b, 0xa1, 0x3a, 0x14, 0x71}}, 5}     // UInt32 -- VT_UI4
	PKEY_Devices_PlaybackState                                   = PROPERTYKEY{GUID{Data1: 0x3633de59, Data2: 0x6825, Data3: 0x4381, Data4: [8]byte{0xa4, 0x9b, 0x9f, 0x6b, 0xa1, 0x3a, 0x14, 0x71}}, 2}     // Byte -- VT_UI1
	PKEY_Devices_PlaybackTitle                                   = PROPERTYKEY{GUID{Data1: 0x3633de59, Data2: 0x6825, Data3: 0x4381, Data4: [8]byte{0xa4, 0x9b, 0x9f, 0x6b, 0xa1, 0x3a, 0x14, 0x71}}, 3}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Present                                         = PROPERTYKEY{GUID{Data1: 0x540b947e, Data2: 0x8b40, Data3: 0x45bc, Data4: [8]byte{0xa8, 0xa2, 0x6a, 0x0b, 0x89, 0x4c, 0xbd, 0xa2}}, 5}     // Boolean -- VT_BOOL
	PKEY_Devices_PresentationUrl                                 = PROPERTYKEY{GUID{Data1: 0x656a3bb3, Data2: 0xecc0, Data3: 0x43fd, Data4: [8]byte{0x84, 0x77, 0x4a, 0xe0, 0x40, 0x4a, 0x96, 0xcd}}, 8198}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_PrimaryCategory                                 = PROPERTYKEY{GUID{Data1: 0xd08dd4c0, Data2: 0x3a9e, Data3: 0x462e, Data4: [8]byte{0x82, 0x90, 0x7b, 0x63, 0x6b, 0x25, 0x76, 0xb9}}, 10}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_RemainingDuration                               = PROPERTYKEY{GUID{Data1: 0x3633de59, Data2: 0x6825, Data3: 0x4381, Data4: [8]byte{0xa4, 0x9b, 0x9f, 0x6b, 0xa1, 0x3a, 0x14, 0x71}}, 4}     // UInt64 -- VT_UI8
	PKEY_Devices_RestrictedInterface                             = PROPERTYKEY{GUID{Data1: 0x026e516e, Data2: 0xb814, Data3: 0x414b, Data4: [8]byte{0x83, 0xcd, 0x85, 0x6d, 0x6f, 0xef, 0x48, 0x22}}, 6}     // Boolean -- VT_BOOL
	PKEY_Devices_Roaming                                         = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 9}     // Byte -- VT_UI1
	PKEY_Devices_SafeRemovalRequired                             = PROPERTYKEY{GUID{Data1: 0xafd97640, Data2: 0x86a3, Data3: 0x4210, Data4: [8]byte{0xb6, 0x7c, 0x28, 0x9c, 0x41, 0xaa, 0xbe, 0x55}}, 2}     // Boolean -- VT_BOOL
	PKEY_Devices_SchematicName                                   = PROPERTYKEY{GUID{Data1: 0x026e516e, Data2: 0xb814, Data3: 0x414b, Data4: [8]byte{0x83, 0xcd, 0x85, 0x6d, 0x6f, 0xef, 0x48, 0x22}}, 9}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_ServiceAddress                                  = PROPERTYKEY{GUID{Data1: 0x656a3bb3, Data2: 0xecc0, Data3: 0x43fd, Data4: [8]byte{0x84, 0x77, 0x4a, 0xe0, 0x40, 0x4a, 0x96, 0xcd}}, 16384} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_ServiceId                                       = PROPERTYKEY{GUID{Data1: 0x656a3bb3, Data2: 0xecc0, Data3: 0x43fd, Data4: [8]byte{0x84, 0x77, 0x4a, 0xe0, 0x40, 0x4a, 0x96, 0xcd}}, 16385} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_SharedTooltip                                   = PROPERTYKEY{GUID{Data1: 0x880f70a2, Data2: 0x6082, Data3: 0x47ac, Data4: [8]byte{0x8a, 0xab, 0xa7, 0x39, 0xd1, 0xa3, 0x00, 0xc3}}, 151}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_SignalStrength                                  = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 2}     // Byte -- VT_UI1
	PKEY_Devices_SmartCards_ReaderKind                           = PROPERTYKEY{GUID{Data1: 0xd6b5b883, Data2: 0x18bd, Data3: 0x4b4d, Data4: [8]byte{0xb2, 0xec, 0x9e, 0x38, 0xaf, 0xfe, 0xda, 0x82}}, 2}     // Byte -- VT_UI1
	PKEY_Devices_Status                                          = PROPERTYKEY{GUID{Data1: 0xd08dd4c0, Data2: 0x3a9e, Data3: 0x462e, Data4: [8]byte{0x82, 0x90, 0x7b, 0x63, 0x6b, 0x25, 0x76, 0xb9}}, 259}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Status1                                         = PROPERTYKEY{GUID{Data1: 0xd08dd4c0, Data2: 0x3a9e, Data3: 0x462e, Data4: [8]byte{0x82, 0x90, 0x7b, 0x63, 0x6b, 0x25, 0x76, 0xb9}}, 257}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Status2                                         = PROPERTYKEY{GUID{Data1: 0xd08dd4c0, Data2: 0x3a9e, Data3: 0x462e, Data4: [8]byte{0x82, 0x90, 0x7b, 0x63, 0x6b, 0x25, 0x76, 0xb9}}, 258}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_StorageCapacity                                 = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 12}    // UInt64 -- VT_UI8
	PKEY_Devices_StorageFreeSpace                                = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 13}    // UInt64 -- VT_UI8
	PKEY_Devices_StorageFreeSpacePercent                         = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 14}    // UInt32 -- VT_UI4
	PKEY_Devices_TextMessages                                    = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 3}     // Byte -- VT_UI1
	PKEY_Devices_Voicemail                                       = PROPERTYKEY{GUID{Data1: 0x49cd1f76, Data2: 0x5626, Data3: 0x4b17, Data4: [8]byte{0xa4, 0xe8, 0x18, 0xb4, 0xaa, 0x1a, 0x22, 0x13}}, 6}     // Byte -- VT_UI1
	PKEY_Devices_WiaDeviceType                                   = PROPERTYKEY{GUID{Data1: 0x6bdd1fc6, Data2: 0x810f, Data3: 0x11d0, Data4: [8]byte{0xbe, 0xc7, 0x08, 0x00, 0x2b, 0xe2, 0x09, 0x2f}}, 2}     // UInt32 -- VT_UI4
	PKEY_Devices_WiFi_InterfaceGuid                              = PROPERTYKEY{GUID{Data1: 0xef1167eb, Data2: 0xcbfc, Data3: 0x4341, Data4: [8]byte{0xa5, 0x68, 0xa7, 0xc9, 0x1a, 0x68, 0x98, 0x2c}}, 2}     // Guid -- VT_CLSID
	PKEY_Devices_WiFiDirect_DeviceAddress                        = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 13}    // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirect_GroupId                              = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 4}     // Guid -- VT_CLSID
	PKEY_Devices_WiFiDirect_InformationElements                  = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 12}    // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirect_InterfaceAddress                     = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 2}     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirect_InterfaceGuid                        = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 3}     // Guid -- VT_CLSID
	PKEY_Devices_WiFiDirect_IsConnected                          = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 5}     // Boolean -- VT_BOOL
	PKEY_Devices_WiFiDirect_IsLegacyDevice                       = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 7}     // Boolean -- VT_BOOL
	PKEY_Devices_WiFiDirect_IsMiracastLcpSupported               = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 9}     // Boolean -- VT_BOOL
	PKEY_Devices_WiFiDirect_IsVisible                            = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 6}     // Boolean -- VT_BOOL
	PKEY_Devices_WiFiDirect_MiracastVersion                      = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 8}     // UInt32 -- VT_UI4
	PKEY_Devices_WiFiDirect_Services                             = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 10}    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_WiFiDirect_SupportedChannelList                 = PROPERTYKEY{GUID{Data1: 0x1506935d, Data2: 0xe3e7, Data3: 0x450f, Data4: [8]byte{0x86, 0x37, 0x82, 0x23, 0x3e, 0xbe, 0x5f, 0x6e}}, 11}    // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirectServices_AdvertisementId              = PROPERTYKEY{GUID{Data1: 0x31b37743, Data2: 0x7c5e, Data3: 0x4005, Data4: [8]byte{0x93, 0xe6, 0xe9, 0x53, 0xf9, 0x2b, 0x82, 0xe9}}, 5}     // UInt32 -- VT_UI4
	PKEY_Devices_WiFiDirectServices_RequestServiceInformation    = PROPERTYKEY{GUID{Data1: 0x31b37743, Data2: 0x7c5e, Data3: 0x4005, Data4: [8]byte{0x93, 0xe6, 0xe9, 0x53, 0xf9, 0x2b, 0x82, 0xe9}}, 7}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_WiFiDirectServices_ServiceAddress               = PROPERTYKEY{GUID{Data1: 0x31b37743, Data2: 0x7c5e, Data3: 0x4005, Data4: [8]byte{0x93, 0xe6, 0xe9, 0x53, 0xf9, 0x2b, 0x82, 0xe9}}, 2}     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirectServices_ServiceConfigMethods         = PROPERTYKEY{GUID{Data1: 0x31b37743, Data2: 0x7c5e, Data3: 0x4005, Data4: [8]byte{0x93, 0xe6, 0xe9, 0x53, 0xf9, 0x2b, 0x82, 0xe9}}, 6}     // UInt32 -- VT_UI4
	PKEY_Devices_WiFiDirectServices_ServiceInformation           = PROPERTYKEY{GUID{Data1: 0x31b37743, Data2: 0x7c5e, Data3: 0x4005, Data4: [8]byte{0x93, 0xe6, 0xe9, 0x53, 0xf9, 0x2b, 0x82, 0xe9}}, 4}     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirectServices_ServiceName                  = PROPERTYKEY{GUID{Data1: 0x31b37743, Data2: 0x7c5e, Data3: 0x4005, Data4: [8]byte{0x93, 0xe6, 0xe9, 0x53, 0xf9, 0x2b, 0x82, 0xe9}}, 3}     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_WinPhone8CameraFlags                            = PROPERTYKEY{GUID{Data1: 0xb7b4d61c, Data2: 0x5a64, Data3: 0x4187, Data4: [8]byte{0xa5, 0x2e, 0xb1, 0x53, 0x9f, 0x35, 0x90, 0x99}}, 2}     // UInt32 -- VT_UI4
	PKEY_Devices_Wwan_InterfaceGuid                              = PROPERTYKEY{GUID{Data1: 0xff1167eb, Data2: 0xcbfc, Data3: 0x4341, Data4: [8]byte{0xa5, 0x68, 0xa7, 0xc9, 0x1a, 0x68, 0x98, 0x2c}}, 2}     // Guid -- VT_CLSID
	PKEY_Storage_Portable                                        = PROPERTYKEY{GUID{Data1: 0x4d1ebee8, Data2: 0x0803, Data3: 0x4774, Data4: [8]byte{0x98, 0x42, 0xb7, 0x7d, 0xb5, 0x02, 0x65, 0xe9}}, 2}     // Boolean -- VT_BOOL
	PKEY_Storage_RemovableMedia                                  = PROPERTYKEY{GUID{Data1: 0x4d1ebee8, Data2: 0x0803, Data3: 0x4774, Data4: [8]byte{0x98, 0x42, 0xb7, 0x7d, 0xb5, 0x02, 0x65, 0xe9}}, 3}     // Boolean -- VT_BOOL
	PKEY_Storage_SystemCritical                                  = PROPERTYKEY{GUID{Data1: 0x4d1ebee8, Data2: 0x0803, Data3: 0x4774, Data4: [8]byte{0x98, 0x42, 0xb7, 0x7d, 0xb5, 0x02, 0x65, 0xe9}}, 4}     // Boolean -- VT_BOOL

	// Document properties

	PKEY_Document_ByteCount           = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 4}   // Int32 -- VT_I4
	PKEY_Document_CharacterCount      = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 16}  // Int32 -- VT_I4
	PKEY_Document_ClientID            = PROPERTYKEY{GUID{Data1: 0x276d7bb0, Data2: 0x5b34, Data3: 0x4fb0, Data4: [8]byte{0xaa, 0x4b, 0x15, 0x8e, 0xd1, 0x2a, 0x18, 0x09}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_Contributor         = PROPERTYKEY{GUID{Data1: 0xf334115e, Data2: 0xda1b, Data3: 0x4509, Data4: [8]byte{0x9b, 0x3d, 0x11, 0x95, 0x04, 0xdc, 0x7a, 0xbb}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Document_DateCreated         = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 12}  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Document_DatePrinted         = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 11}  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Document_DateSaved           = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 13}  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Document_Division            = PROPERTYKEY{GUID{Data1: 0x1e005ee6, Data2: 0xbf27, Data3: 0x428b, Data4: [8]byte{0xb0, 0x1c, 0x79, 0x67, 0x6a, 0xcd, 0x28, 0x70}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_DocumentID          = PROPERTYKEY{GUID{Data1: 0xe08805c8, Data2: 0xe395, Data3: 0x40df, Data4: [8]byte{0x80, 0xd2, 0x54, 0xf0, 0xd6, 0xc4, 0x31, 0x54}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_HiddenSlideCount    = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 9}   // Int32 -- VT_I4
	PKEY_Document_LastAuthor          = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 8}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_LineCount           = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 5}   // Int32 -- VT_I4
	PKEY_Document_Manager             = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 14}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_MultimediaClipCount = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 10}  // Int32 -- VT_I4
	PKEY_Document_NoteCount           = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 8}   // Int32 -- VT_I4
	PKEY_Document_PageCount           = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 14}  // Int32 -- VT_I4
	PKEY_Document_ParagraphCount      = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 6}   // Int32 -- VT_I4
	PKEY_Document_PresentationFormat  = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 3}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_RevisionNumber      = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 9}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_Security            = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 19}  // Int32 -- VT_I4
	PKEY_Document_SlideCount          = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 7}   // Int32 -- VT_I4
	PKEY_Document_Template            = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 7}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_TotalEditingTime    = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 10}  // UInt64 -- VT_UI8
	PKEY_Document_Version             = PROPERTYKEY{GUID{Data1: 0xd5cdd502, Data2: 0x2e9c, Data3: 0x101b, Data4: [8]byte{0x93, 0x97, 0x08, 0x00, 0x2b, 0x2c, 0xf9, 0xae}}, 29}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_WordCount           = PROPERTYKEY{GUID{Data1: 0xf29f85e0, Data2: 0x4ff9, Data3: 0x1068, Data4: [8]byte{0xab, 0x91, 0x08, 0x00, 0x2b, 0x27, 0xb3, 0xd9}}, 15}  // Int32 -- VT_I4

	// DRM properties

	PKEY_DRM_DatePlayExpires = PROPERTYKEY{GUID{Data1: 0xaeac19e4, Data2: 0x89ae, Data3: 0x4508, Data4: [8]byte{0xb9, 0xb7, 0xbb, 0x86, 0x7a, 0xbe, 0xe2, 0xed}}, 6} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DRM_DatePlayStarts  = PROPERTYKEY{GUID{Data1: 0xaeac19e4, Data2: 0x89ae, Data3: 0x4508, Data4: [8]byte{0xb9, 0xb7, 0xbb, 0x86, 0x7a, 0xbe, 0xe2, 0xed}}, 5} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DRM_Description     = PROPERTYKEY{GUID{Data1: 0xaeac19e4, Data2: 0x89ae, Data3: 0x4508, Data4: [8]byte{0xb9, 0xb7, 0xbb, 0x86, 0x7a, 0xbe, 0xe2, 0xed}}, 3} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DRM_IsDisabled      = PROPERTYKEY{GUID{Data1: 0xaeac19e4, Data2: 0x89ae, Data3: 0x4508, Data4: [8]byte{0xb9, 0xb7, 0xbb, 0x86, 0x7a, 0xbe, 0xe2, 0xed}}, 7} // Boolean -- VT_BOOL
	PKEY_DRM_IsProtected     = PROPERTYKEY{GUID{Data1: 0xaeac19e4, Data2: 0x89ae, Data3: 0x4508, Data4: [8]byte{0xb9, 0xb7, 0xbb, 0x86, 0x7a, 0xbe, 0xe2, 0xed}}, 2} // Boolean -- VT_BOOL
	PKEY_DRM_PlayCount       = PROPERTYKEY{GUID{Data1: 0xaeac19e4, Data2: 0x89ae, Data3: 0x4508, Data4: [8]byte{0xb9, 0xb7, 0xbb, 0x86, 0x7a, 0xbe, 0xe2, 0xed}}, 4} // UInt32 -- VT_UI4

	// GPS properties

	PKEY_GPS_Altitude                 = PROPERTYKEY{GUID{Data1: 0x827edb4f, Data2: 0x5b73, Data3: 0x44a7, Data4: [8]byte{0x89, 0x1d, 0xfd, 0xff, 0xab, 0xea, 0x35, 0xca}}, 100} // Double -- VT_R8
	PKEY_GPS_AltitudeDenominator      = PROPERTYKEY{GUID{Data1: 0x78342dcb, Data2: 0xe358, Data3: 0x4145, Data4: [8]byte{0xae, 0x9a, 0x6b, 0xfe, 0x4e, 0x0f, 0x9f, 0x51}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_AltitudeNumerator        = PROPERTYKEY{GUID{Data1: 0x2dad1eb7, Data2: 0x816d, Data3: 0x40d3, Data4: [8]byte{0x9e, 0xc3, 0xc9, 0x77, 0x3b, 0xe2, 0xaa, 0xde}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_AltitudeRef              = PROPERTYKEY{GUID{Data1: 0x46ac629d, Data2: 0x75ea, Data3: 0x4515, Data4: [8]byte{0x86, 0x7f, 0x6d, 0xc4, 0x32, 0x1c, 0x58, 0x44}}, 100} // Byte -- VT_UI1
	PKEY_GPS_AreaInformation          = PROPERTYKEY{GUID{Data1: 0x972e333e, Data2: 0xac7e, Data3: 0x49f1, Data4: [8]byte{0x8a, 0xdf, 0xa7, 0x0d, 0x07, 0xa9, 0xbc, 0xab}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Date                     = PROPERTYKEY{GUID{Data1: 0x3602c812, Data2: 0x0f3b, Data3: 0x45f0, Data4: [8]byte{0x85, 0xad, 0x60, 0x34, 0x68, 0xd6, 0x94, 0x23}}, 100} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_GPS_DestBearing              = PROPERTYKEY{GUID{Data1: 0xc66d4b3c, Data2: 0xe888, Data3: 0x47cc, Data4: [8]byte{0xb9, 0x9f, 0x9d, 0xca, 0x3e, 0xe3, 0x4d, 0xea}}, 100} // Double -- VT_R8
	PKEY_GPS_DestBearingDenominator   = PROPERTYKEY{GUID{Data1: 0x7abcf4f8, Data2: 0x7c3f, Data3: 0x4988, Data4: [8]byte{0xac, 0x91, 0x8d, 0x2c, 0x2e, 0x97, 0xec, 0xa5}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_DestBearingNumerator     = PROPERTYKEY{GUID{Data1: 0xba3b1da9, Data2: 0x86ee, Data3: 0x4b5d, Data4: [8]byte{0xa2, 0xa4, 0xa2, 0x71, 0xa4, 0x29, 0xf0, 0xcf}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_DestBearingRef           = PROPERTYKEY{GUID{Data1: 0x9ab84393, Data2: 0x2a0f, Data3: 0x4b75, Data4: [8]byte{0xbb, 0x22, 0x72, 0x79, 0x78, 0x69, 0x77, 0xcb}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_DestDistance             = PROPERTYKEY{GUID{Data1: 0xa93eae04, Data2: 0x6804, Data3: 0x4f24, Data4: [8]byte{0xac, 0x81, 0x09, 0xb2, 0x66, 0x45, 0x21, 0x18}}, 100} // Double -- VT_R8
	PKEY_GPS_DestDistanceDenominator  = PROPERTYKEY{GUID{Data1: 0x9bc2c99b, Data2: 0xac71, Data3: 0x4127, Data4: [8]byte{0x9d, 0x1c, 0x25, 0x96, 0xd0, 0xd7, 0xdc, 0xb7}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_DestDistanceNumerator    = PROPERTYKEY{GUID{Data1: 0x2bda47da, Data2: 0x08c6, Data3: 0x4fe1, Data4: [8]byte{0x80, 0xbc, 0xa7, 0x2f, 0xc5, 0x17, 0xc5, 0xd0}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_DestDistanceRef          = PROPERTYKEY{GUID{Data1: 0xed4df2d3, Data2: 0x8695, Data3: 0x450b, Data4: [8]byte{0x85, 0x6f, 0xf5, 0xc1, 0xc5, 0x3a, 0xcb, 0x66}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_DestLatitude             = PROPERTYKEY{GUID{Data1: 0x9d1d7cc5, Data2: 0x5c39, Data3: 0x451c, Data4: [8]byte{0x86, 0xb3, 0x92, 0x8e, 0x2d, 0x18, 0xcc, 0x47}}, 100} // Multivalue Double -- VT_VECTOR | VT_R8  (For variants: VT_ARRAY | VT_R8)
	PKEY_GPS_DestLatitudeDenominator  = PROPERTYKEY{GUID{Data1: 0x3a372292, Data2: 0x7fca, Data3: 0x49a7, Data4: [8]byte{0x99, 0xd5, 0xe4, 0x7b, 0xb2, 0xd4, 0xe7, 0xab}}, 100} // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_DestLatitudeNumerator    = PROPERTYKEY{GUID{Data1: 0xecf4b6f6, Data2: 0xd5a6, Data3: 0x433c, Data4: [8]byte{0xbb, 0x92, 0x40, 0x76, 0x65, 0x0f, 0xc8, 0x90}}, 100} // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_DestLatitudeRef          = PROPERTYKEY{GUID{Data1: 0xcea820b9, Data2: 0xce61, Data3: 0x4885, Data4: [8]byte{0xa1, 0x28, 0x00, 0x5d, 0x90, 0x87, 0xc1, 0x92}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_DestLongitude            = PROPERTYKEY{GUID{Data1: 0x47a96261, Data2: 0xcb4c, Data3: 0x4807, Data4: [8]byte{0x8a, 0xd3, 0x40, 0xb9, 0xd9, 0xdb, 0xc6, 0xbc}}, 100} // Multivalue Double -- VT_VECTOR | VT_R8  (For variants: VT_ARRAY | VT_R8)
	PKEY_GPS_DestLongitudeDenominator = PROPERTYKEY{GUID{Data1: 0x425d69e5, Data2: 0x48ad, Data3: 0x4900, Data4: [8]byte{0x8d, 0x80, 0x6e, 0xb6, 0xb8, 0xd0, 0xac, 0x86}}, 100} // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_DestLongitudeNumerator   = PROPERTYKEY{GUID{Data1: 0xa3250282, Data2: 0xfb6d, Data3: 0x48d5, Data4: [8]byte{0x9a, 0x89, 0xdb, 0xca, 0xce, 0x75, 0xcc, 0xcf}}, 100} // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_DestLongitudeRef         = PROPERTYKEY{GUID{Data1: 0x182c1ea6, Data2: 0x7c1c, Data3: 0x4083, Data4: [8]byte{0xab, 0x4b, 0xac, 0x6c, 0x9f, 0x4e, 0xd1, 0x28}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Differential             = PROPERTYKEY{GUID{Data1: 0xaaf4ee25, Data2: 0xbd3b, Data3: 0x4dd7, Data4: [8]byte{0xbf, 0xc4, 0x47, 0xf7, 0x7b, 0xb0, 0x0f, 0x6d}}, 100} // UInt16 -- VT_UI2
	PKEY_GPS_DOP                      = PROPERTYKEY{GUID{Data1: 0x0cf8fb02, Data2: 0x1837, Data3: 0x42f1, Data4: [8]byte{0xa6, 0x97, 0xa7, 0x01, 0x7a, 0xa2, 0x89, 0xb9}}, 100} // Double -- VT_R8
	PKEY_GPS_DOPDenominator           = PROPERTYKEY{GUID{Data1: 0xa0be94c5, Data2: 0x50ba, Data3: 0x487b, Data4: [8]byte{0xbd, 0x35, 0x06, 0x54, 0xbe, 0x88, 0x81, 0xed}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_DOPNumerator             = PROPERTYKEY{GUID{Data1: 0x47166b16, Data2: 0x364f, Data3: 0x4aa0, Data4: [8]byte{0x9f, 0x31, 0xe2, 0xab, 0x3d, 0xf4, 0x49, 0xc3}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_ImgDirection             = PROPERTYKEY{GUID{Data1: 0x16473c91, Data2: 0xd017, Data3: 0x4ed9, Data4: [8]byte{0xba, 0x4d, 0xb6, 0xba, 0xa5, 0x5d, 0xbc, 0xf8}}, 100} // Double -- VT_R8
	PKEY_GPS_ImgDirectionDenominator  = PROPERTYKEY{GUID{Data1: 0x10b24595, Data2: 0x41a2, Data3: 0x4e20, Data4: [8]byte{0x93, 0xc2, 0x57, 0x61, 0xc1, 0x39, 0x5f, 0x32}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_ImgDirectionNumerator    = PROPERTYKEY{GUID{Data1: 0xdc5877c7, Data2: 0x225f, Data3: 0x45f7, Data4: [8]byte{0xba, 0xc7, 0xe8, 0x13, 0x34, 0xb6, 0x13, 0x0a}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_ImgDirectionRef          = PROPERTYKEY{GUID{Data1: 0xa4aaa5b7, Data2: 0x1ad0, Data3: 0x445f, Data4: [8]byte{0x81, 0x1a, 0x0f, 0x8f, 0x6e, 0x67, 0xf6, 0xb5}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Latitude                 = PROPERTYKEY{GUID{Data1: 0x8727cfff, Data2: 0x4868, Data3: 0x4ec6, Data4: [8]byte{0xad, 0x5b, 0x81, 0xb9, 0x85, 0x21, 0xd1, 0xab}}, 100} // Multivalue Double -- VT_VECTOR | VT_R8  (For variants: VT_ARRAY | VT_R8)
	PKEY_GPS_LatitudeDecimal          = PROPERTYKEY{GUID{Data1: 0x0f55cde2, Data2: 0x4f49, Data3: 0x450d, Data4: [8]byte{0x92, 0xc1, 0xdc, 0xd1, 0x63, 0x01, 0xb1, 0xb7}}, 100} // Double -- VT_R8
	PKEY_GPS_LatitudeDenominator      = PROPERTYKEY{GUID{Data1: 0x16e634ee, Data2: 0x2bff, Data3: 0x497b, Data4: [8]byte{0xbd, 0x8a, 0x43, 0x41, 0xad, 0x39, 0xee, 0xb9}}, 100} // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_LatitudeNumerator        = PROPERTYKEY{GUID{Data1: 0x7ddaaad1, Data2: 0xccc8, Data3: 0x41ae, Data4: [8]byte{0xb7, 0x50, 0xb2, 0xcb, 0x80, 0x31, 0xae, 0xa2}}, 100} // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_LatitudeRef              = PROPERTYKEY{GUID{Data1: 0x029c0252, Data2: 0x5b86, Data3: 0x46c7, Data4: [8]byte{0xac, 0xa0, 0x27, 0x69, 0xff, 0xc8, 0xe3, 0xd4}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Longitude                = PROPERTYKEY{GUID{Data1: 0xc4c4dbb2, Data2: 0xb593, Data3: 0x466b, Data4: [8]byte{0xbb, 0xda, 0xd0, 0x3d, 0x27, 0xd5, 0xe4, 0x3a}}, 100} // Multivalue Double -- VT_VECTOR | VT_R8  (For variants: VT_ARRAY | VT_R8)
	PKEY_GPS_LongitudeDecimal         = PROPERTYKEY{GUID{Data1: 0x4679c1b5, Data2: 0x844d, Data3: 0x4590, Data4: [8]byte{0xba, 0xf5, 0xf3, 0x22, 0x23, 0x1f, 0x1b, 0x81}}, 100} // Double -- VT_R8
	PKEY_GPS_LongitudeDenominator     = PROPERTYKEY{GUID{Data1: 0xbe6e176c, Data2: 0x4534, Data3: 0x4d2c, Data4: [8]byte{0xac, 0xe5, 0x31, 0xde, 0xda, 0xc1, 0x60, 0x6b}}, 100} // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_LongitudeNumerator       = PROPERTYKEY{GUID{Data1: 0x02b0f689, Data2: 0xa914, Data3: 0x4e45, Data4: [8]byte{0x82, 0x1d, 0x1d, 0xda, 0x45, 0x2e, 0xd2, 0xc4}}, 100} // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_LongitudeRef             = PROPERTYKEY{GUID{Data1: 0x33dcf22b, Data2: 0x28d5, Data3: 0x464c, Data4: [8]byte{0x80, 0x35, 0x1e, 0xe9, 0xef, 0xd2, 0x52, 0x78}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_MapDatum                 = PROPERTYKEY{GUID{Data1: 0x2ca2dae6, Data2: 0xeddc, Data3: 0x407d, Data4: [8]byte{0xbe, 0xf1, 0x77, 0x39, 0x42, 0xab, 0xfa, 0x95}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_MeasureMode              = PROPERTYKEY{GUID{Data1: 0xa015ed5d, Data2: 0xaaea, Data3: 0x4d58, Data4: [8]byte{0x8a, 0x86, 0x3c, 0x58, 0x69, 0x20, 0xea, 0x0b}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_ProcessingMethod         = PROPERTYKEY{GUID{Data1: 0x59d49e61, Data2: 0x840f, Data3: 0x4aa9, Data4: [8]byte{0xa9, 0x39, 0xe2, 0x09, 0x9b, 0x7f, 0x63, 0x99}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Satellites               = PROPERTYKEY{GUID{Data1: 0x467ee575, Data2: 0x1f25, Data3: 0x4557, Data4: [8]byte{0xad, 0x4e, 0xb8, 0xb5, 0x8b, 0x0d, 0x9c, 0x15}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Speed                    = PROPERTYKEY{GUID{Data1: 0xda5d0862, Data2: 0x6e76, Data3: 0x4e1b, Data4: [8]byte{0xba, 0xbd, 0x70, 0x02, 0x1b, 0xd2, 0x54, 0x94}}, 100} // Double -- VT_R8
	PKEY_GPS_SpeedDenominator         = PROPERTYKEY{GUID{Data1: 0x7d122d5a, Data2: 0xae5e, Data3: 0x4335, Data4: [8]byte{0x88, 0x41, 0xd7, 0x1e, 0x7c, 0xe7, 0x2f, 0x53}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_SpeedNumerator           = PROPERTYKEY{GUID{Data1: 0xacc9ce3d, Data2: 0xc213, Data3: 0x4942, Data4: [8]byte{0x8b, 0x48, 0x6d, 0x08, 0x20, 0xf2, 0x1c, 0x6d}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_SpeedRef                 = PROPERTYKEY{GUID{Data1: 0xecf7f4c9, Data2: 0x544f, Data3: 0x4d6d, Data4: [8]byte{0x9d, 0x98, 0x8a, 0xd7, 0x9a, 0xda, 0xf4, 0x53}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Status                   = PROPERTYKEY{GUID{Data1: 0x125491f4, Data2: 0x818f, Data3: 0x46b2, Data4: [8]byte{0x91, 0xb5, 0xd5, 0x37, 0x75, 0x36, 0x17, 0xb2}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Track                    = PROPERTYKEY{GUID{Data1: 0x76c09943, Data2: 0x7c33, Data3: 0x49e3, Data4: [8]byte{0x9e, 0x7e, 0xcd, 0xba, 0x87, 0x2c, 0xfa, 0xda}}, 100} // Double -- VT_R8
	PKEY_GPS_TrackDenominator         = PROPERTYKEY{GUID{Data1: 0xc8d1920c, Data2: 0x01f6, Data3: 0x40c0, Data4: [8]byte{0xac, 0x86, 0x2f, 0x3a, 0x4a, 0xd0, 0x07, 0x70}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_TrackNumerator           = PROPERTYKEY{GUID{Data1: 0x702926f4, Data2: 0x44a6, Data3: 0x43e1, Data4: [8]byte{0xae, 0x71, 0x45, 0x62, 0x71, 0x16, 0x89, 0x3b}}, 100} // UInt32 -- VT_UI4
	PKEY_GPS_TrackRef                 = PROPERTYKEY{GUID{Data1: 0x35dbe6fe, Data2: 0x44c3, Data3: 0x4400, Data4: [8]byte{0xaa, 0xae, 0xd2, 0xc7, 0x99, 0xc4, 0x07, 0xe8}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_VersionID                = PROPERTYKEY{GUID{Data1: 0x22704da4, Data2: 0xc6b2, Data3: 0x4a99, Data4: [8]byte{0x8e, 0x56, 0xf1, 0x6d, 0xf8, 0xc9, 0x25, 0x99}}, 100} // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)

	// History properties

	PKEY_History_VisitCount = PROPERTYKEY{GUID{Data1: 0x5cbf2787, Data2: 0x48cf, Data3: 0x4208, Data4: [8]byte{0xb9, 0x0e, 0xee, 0x5e, 0x5d, 0x42, 0x02, 0x94}}, 7} // Int32 -- VT_I4

	// Image properties

	PKEY_Image_BitDepth                          = PROPERTYKEY{GUID{Data1: 0x6444048f, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 7}     // UInt32 -- VT_UI4
	PKEY_Image_ColorSpace                        = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 40961} // UInt16 -- VT_UI2
	PKEY_Image_CompressedBitsPerPixel            = PROPERTYKEY{GUID{Data1: 0x364b6fa9, Data2: 0x37ab, Data3: 0x482a, Data4: [8]byte{0xbe, 0x2b, 0xae, 0x02, 0xf6, 0x0d, 0x43, 0x18}}, 100}   // Double -- VT_R8
	PKEY_Image_CompressedBitsPerPixelDenominator = PROPERTYKEY{GUID{Data1: 0x1f8844e1, Data2: 0x24ad, Data3: 0x4508, Data4: [8]byte{0x9d, 0xfd, 0x53, 0x26, 0xa4, 0x15, 0xce, 0x02}}, 100}   // UInt32 -- VT_UI4
	PKEY_Image_CompressedBitsPerPixelNumerator   = PROPERTYKEY{GUID{Data1: 0xd21a7148, Data2: 0xd32c, Data3: 0x4624, Data4: [8]byte{0x89, 0x00, 0x27, 0x72, 0x10, 0xf7, 0x9c, 0x0f}}, 100}   // UInt32 -- VT_UI4
	PKEY_Image_Compression                       = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 259}   // UInt16 -- VT_UI2
	PKEY_Image_CompressionText                   = PROPERTYKEY{GUID{Data1: 0x3f08e66f, Data2: 0x2f44, Data3: 0x4bb9, Data4: [8]byte{0xa6, 0x82, 0xac, 0x35, 0xd2, 0x56, 0x23, 0x22}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Image_Dimensions                        = PROPERTYKEY{GUID{Data1: 0x6444048f, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 13}    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Image_HorizontalResolution              = PROPERTYKEY{GUID{Data1: 0x6444048f, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 5}     // Double -- VT_R8
	PKEY_Image_HorizontalSize                    = PROPERTYKEY{GUID{Data1: 0x6444048f, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 3}     // UInt32 -- VT_UI4
	PKEY_Image_ImageID                           = PROPERTYKEY{GUID{Data1: 0x10dabe05, Data2: 0x32aa, Data3: 0x4c29, Data4: [8]byte{0xbf, 0x1a, 0x63, 0xe2, 0xd2, 0x20, 0x58, 0x7f}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Image_ResolutionUnit                    = PROPERTYKEY{GUID{Data1: 0x19b51fa6, Data2: 0x1f92, Data3: 0x4a5c, Data4: [8]byte{0xab, 0x48, 0x7d, 0xf0, 0xab, 0xd6, 0x74, 0x44}}, 100}   // Int16 -- VT_I2
	PKEY_Image_VerticalResolution                = PROPERTYKEY{GUID{Data1: 0x6444048f, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 6}     // Double -- VT_R8
	PKEY_Image_VerticalSize                      = PROPERTYKEY{GUID{Data1: 0x6444048f, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 4}     // UInt32 -- VT_UI4

	// Journal properties

	PKEY_Journal_Contacts  = PROPERTYKEY{GUID{Data1: 0xdea7c82c, Data2: 0x1d89, Data3: 0x4a66, Data4: [8]byte{0x94, 0x27, 0xa4, 0xe3, 0xde, 0xba, 0xbc, 0xb1}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Journal_EntryType = PROPERTYKEY{GUID{Data1: 0x95beb1fc, Data2: 0x326d, Data3: 0x4644, Data4: [8]byte{0xb3, 0x96, 0xcd, 0x3e, 0xd9, 0x0e, 0x6d, 0xdf}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// LayoutPattern properties

	PKEY_LayoutPattern_ContentViewModeForBrowse = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 500} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_LayoutPattern_ContentViewModeForSearch = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 501} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Link properties

	PKEY_History_SelectionCount    = PROPERTYKEY{GUID{Data1: 0x1ce0d6bc, Data2: 0x536c, Data3: 0x4600, Data4: [8]byte{0xb0, 0xdd, 0x7e, 0x0c, 0x66, 0xb3, 0x50, 0xd5}}, 8}   // UInt32 -- VT_UI4
	PKEY_History_TargetUrlHostName = PROPERTYKEY{GUID{Data1: 0x1ce0d6bc, Data2: 0x536c, Data3: 0x4600, Data4: [8]byte{0xb0, 0xdd, 0x7e, 0x0c, 0x66, 0xb3, 0x50, 0xd5}}, 9}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_Arguments            = PROPERTYKEY{GUID{Data1: 0x436f2667, Data2: 0x14e2, Data3: 0x4feb, Data4: [8]byte{0xb3, 0x0a, 0x14, 0x6c, 0x53, 0xb5, 0xb6, 0x74}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_Comment              = PROPERTYKEY{GUID{Data1: 0xb9b4b3fc, Data2: 0x2b51, Data3: 0x4a42, Data4: [8]byte{0xb5, 0xd8, 0x32, 0x41, 0x46, 0xaf, 0xcf, 0x25}}, 5}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_DateVisited          = PROPERTYKEY{GUID{Data1: 0x5cbf2787, Data2: 0x48cf, Data3: 0x4208, Data4: [8]byte{0xb9, 0x0e, 0xee, 0x5e, 0x5d, 0x42, 0x02, 0x94}}, 23}  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Link_Description          = PROPERTYKEY{GUID{Data1: 0x5cbf2787, Data2: 0x48cf, Data3: 0x4208, Data4: [8]byte{0xb9, 0x0e, 0xee, 0x5e, 0x5d, 0x42, 0x02, 0x94}}, 21}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_FeedItemLocalId      = PROPERTYKEY{GUID{Data1: 0x8a2f99f9, Data2: 0x3c37, Data3: 0x465d, Data4: [8]byte{0xa8, 0xd7, 0x69, 0x77, 0x7a, 0x24, 0x6d, 0x0c}}, 2}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_Status               = PROPERTYKEY{GUID{Data1: 0xb9b4b3fc, Data2: 0x2b51, Data3: 0x4a42, Data4: [8]byte{0xb5, 0xd8, 0x32, 0x41, 0x46, 0xaf, 0xcf, 0x25}}, 3}   // Int32 -- VT_I4
	PKEY_Link_TargetExtension      = PROPERTYKEY{GUID{Data1: 0x7a7d76f4, Data2: 0xb630, Data3: 0x4bd7, Data4: [8]byte{0x95, 0xff, 0x37, 0xcc, 0x51, 0xa9, 0x75, 0xc9}}, 2}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Link_TargetParsingPath    = PROPERTYKEY{GUID{Data1: 0xb9b4b3fc, Data2: 0x2b51, Data3: 0x4a42, Data4: [8]byte{0xb5, 0xd8, 0x32, 0x41, 0x46, 0xaf, 0xcf, 0x25}}, 2}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_TargetSFGAOFlags     = PROPERTYKEY{GUID{Data1: 0xb9b4b3fc, Data2: 0x2b51, Data3: 0x4a42, Data4: [8]byte{0xb5, 0xd8, 0x32, 0x41, 0x46, 0xaf, 0xcf, 0x25}}, 8}   // UInt32 -- VT_UI4
	PKEY_Link_TargetUrlHostName    = PROPERTYKEY{GUID{Data1: 0x8a2f99f9, Data2: 0x3c37, Data3: 0x465d, Data4: [8]byte{0xa8, 0xd7, 0x69, 0x77, 0x7a, 0x24, 0x6d, 0x0c}}, 5}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_TargetUrlPath        = PROPERTYKEY{GUID{Data1: 0x8a2f99f9, Data2: 0x3c37, Data3: 0x465d, Data4: [8]byte{0xa8, 0xd7, 0x69, 0x77, 0x7a, 0x24, 0x6d, 0x0c}}, 6}   // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Media properties

	PKEY_Media_AuthorUrl                 = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 32}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_AverageLevel              = PROPERTYKEY{GUID{Data1: 0x09edd5b6, Data2: 0xb301, Data3: 0x43c5, Data4: [8]byte{0x99, 0x90, 0xd0, 0x03, 0x02, 0xef, 0xfd, 0x46}}, 100} // UInt32 -- VT_UI4
	PKEY_Media_ClassPrimaryID            = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 13}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ClassSecondaryID          = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 14}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_CollectionGroupID         = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 24}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_CollectionID              = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 25}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ContentDistributor        = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 18}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ContentID                 = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 26}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_CreatorApplication        = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 27}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_CreatorApplicationVersion = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 28}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_DateEncoded               = PROPERTYKEY{GUID{Data1: 0x2e4b640d, Data2: 0x5019, Data3: 0x46d8, Data4: [8]byte{0x88, 0x81, 0x55, 0x41, 0x4c, 0xc5, 0xca, 0xa0}}, 100} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Media_DateReleased              = PROPERTYKEY{GUID{Data1: 0xde41cc29, Data2: 0x6971, Data3: 0x4290, Data4: [8]byte{0xb4, 0x72, 0xf5, 0x9f, 0x2e, 0x2f, 0x31, 0xe2}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_DlnaProfileID             = PROPERTYKEY{GUID{Data1: 0xcfa31b45, Data2: 0x525d, Data3: 0x4998, Data4: [8]byte{0xbb, 0x44, 0x3f, 0x7d, 0x81, 0x54, 0x2f, 0xa4}}, 100} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Media_Duration                  = PROPERTYKEY{GUID{Data1: 0x64440490, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 3}   // UInt64 -- VT_UI8
	PKEY_Media_DVDID                     = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 15}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_EncodedBy                 = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 36}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_EncodingSettings          = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 37}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_EpisodeNumber             = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 100} // UInt32 -- VT_UI4
	PKEY_Media_FrameCount                = PROPERTYKEY{GUID{Data1: 0x6444048f, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 12}  // UInt32 -- VT_UI4
	PKEY_Media_MCDI                      = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 16}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_MetadataContentProvider   = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 17}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_Producer                  = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 22}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Media_PromotionUrl              = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 33}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ProtectionType            = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 38}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ProviderRating            = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 39}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ProviderStyle             = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 40}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_Publisher                 = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 30}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_SeasonNumber              = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 101} // UInt32 -- VT_UI4
	PKEY_Media_SeriesName                = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 42}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_SubscriptionContentId     = PROPERTYKEY{GUID{Data1: 0x9aebae7a, Data2: 0x9644, Data3: 0x487d, Data4: [8]byte{0xa9, 0x2c, 0x65, 0x75, 0x85, 0xed, 0x75, 0x1a}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_SubTitle                  = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 38}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ThumbnailLargePath        = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 47}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ThumbnailLargeUri         = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 48}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ThumbnailSmallPath        = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 49}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ThumbnailSmallUri         = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 50}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_UniqueFileIdentifier      = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 35}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_UserNoAutoInfo            = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 41}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_UserWebUrl                = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 34}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_Writer                    = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 23}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Media_Year                      = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 5}   // UInt32 -- VT_UI4

	// Message properties

	PKEY_Message_AttachmentContents = PROPERTYKEY{GUID{Data1: 0x3143bf7c, Data2: 0x80a8, Data3: 0x4854, Data4: [8]byte{0x88, 0x80, 0xe2, 0xe4, 0x01, 0x89, 0xbd, 0xd0}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_AttachmentNames    = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 21}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_BccAddress         = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 2}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_BccName            = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 3}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_CcAddress          = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 4}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_CcName             = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 5}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_ConversationID     = PROPERTYKEY{GUID{Data1: 0xdc8f80bd, Data2: 0xaf1e, Data3: 0x4289, Data4: [8]byte{0x85, 0xb6, 0x3d, 0xfc, 0x1b, 0x49, 0x39, 0x92}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_ConversationIndex  = PROPERTYKEY{GUID{Data1: 0xdc8f80bd, Data2: 0xaf1e, Data3: 0x4289, Data4: [8]byte{0x85, 0xb6, 0x3d, 0xfc, 0x1b, 0x49, 0x39, 0x92}}, 101} // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Message_DateReceived       = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 20}  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Message_DateSent           = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 19}  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Message_Flags              = PROPERTYKEY{GUID{Data1: 0xa82d9ee7, Data2: 0xca67, Data3: 0x4312, Data4: [8]byte{0x96, 0x5e, 0x22, 0x6b, 0xce, 0xa8, 0x50, 0x23}}, 100} // Int32 -- VT_I4
	PKEY_Message_FromAddress        = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 13}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_FromName           = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 14}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_HasAttachments     = PROPERTYKEY{GUID{Data1: 0x9c1fcf74, Data2: 0x2d97, Data3: 0x41ba, Data4: [8]byte{0xb4, 0xae, 0xcb, 0x2e, 0x36, 0x61, 0xa6, 0xe4}}, 8}   // Boolean -- VT_BOOL
	PKEY_Message_IsFwdOrReply       = PROPERTYKEY{GUID{Data1: 0x9a9bc088, Data2: 0x4f6d, Data3: 0x469e, Data4: [8]byte{0x99, 0x19, 0xe7, 0x05, 0x41, 0x20, 0x40, 0xf9}}, 100} // Int32 -- VT_I4
	PKEY_Message_MessageClass       = PROPERTYKEY{GUID{Data1: 0xcd9ed458, Data2: 0x08ce, Data3: 0x418f, Data4: [8]byte{0xa7, 0x0e, 0xf9, 0x12, 0xc7, 0xbb, 0x9c, 0x5c}}, 103} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_Participants       = PROPERTYKEY{GUID{Data1: 0x1a9ba605, Data2: 0x8e7c, Data3: 0x4d11, Data4: [8]byte{0xad, 0x7d, 0xa5, 0x0a, 0xda, 0x18, 0xba, 0x1b}}, 2}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_ProofInProgress    = PROPERTYKEY{GUID{Data1: 0x9098f33c, Data2: 0x9a7d, Data3: 0x48a8, Data4: [8]byte{0x8d, 0xe5, 0x2e, 0x12, 0x27, 0xa6, 0x4e, 0x91}}, 100} // Boolean -- VT_BOOL
	PKEY_Message_SenderAddress      = PROPERTYKEY{GUID{Data1: 0x0be1c8e7, Data2: 0x1981, Data3: 0x4676, Data4: [8]byte{0xae, 0x14, 0xfd, 0xd7, 0x8f, 0x05, 0xa6, 0xe7}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_SenderName         = PROPERTYKEY{GUID{Data1: 0x0da41cfa, Data2: 0xd224, Data3: 0x4a18, Data4: [8]byte{0xae, 0x2f, 0x59, 0x61, 0x58, 0xdb, 0x4b, 0x3a}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_Store              = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 15}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_ToAddress          = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 16}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_ToDoFlags          = PROPERTYKEY{GUID{Data1: 0x1f856a9f, Data2: 0x6900, Data3: 0x4aba, Data4: [8]byte{0x95, 0x05, 0x2d, 0x5f, 0x1b, 0x4d, 0x66, 0xcb}}, 100} // Int32 -- VT_I4
	PKEY_Message_ToDoTitle          = PROPERTYKEY{GUID{Data1: 0xbccc8a3c, Data2: 0x8cef, Data3: 0x42e5, Data4: [8]byte{0x9b, 0x1c, 0xc6, 0x90, 0x79, 0x39, 0x8b, 0xc7}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_ToName             = PROPERTYKEY{GUID{Data1: 0xe3e0584c, Data2: 0xb788, Data3: 0x4a5a, Data4: [8]byte{0xbb, 0x20, 0x7f, 0x5a, 0x44, 0xc9, 0xac, 0xdd}}, 17}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)

	// Music properties

	PKEY_Music_AlbumArtist             = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 13}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_AlbumArtistSortOverride = PROPERTYKEY{GUID{Data1: 0xf1fdb4af, Data2: 0xf78c, Data3: 0x466c, Data4: [8]byte{0xbb, 0x05, 0x56, 0xe9, 0x2d, 0xb0, 0xb8, 0xec}}, 103} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_AlbumID                 = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_AlbumTitle              = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 4}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_AlbumTitleSortOverride  = PROPERTYKEY{GUID{Data1: 0x13eb7ffc, Data2: 0xec89, Data3: 0x4346, Data4: [8]byte{0xb1, 0x9d, 0xcc, 0xc6, 0xf1, 0x78, 0x42, 0x23}}, 101} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Artist                  = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 2}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_ArtistSortOverride      = PROPERTYKEY{GUID{Data1: 0xdeeb2db5, Data2: 0x0696, Data3: 0x4ce0, Data4: [8]byte{0x94, 0xfe, 0xa0, 0x1f, 0x77, 0xa4, 0x5f, 0xb5}}, 102} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_BeatsPerMinute          = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 35}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Composer                = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 19}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_ComposerSortOverride    = PROPERTYKEY{GUID{Data1: 0x00bc20a3, Data2: 0xbd48, Data3: 0x4085, Data4: [8]byte{0x87, 0x2c, 0xa8, 0x8d, 0x77, 0xf5, 0x09, 0x7e}}, 105} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_Conductor               = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 36}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_ContentGroupDescription = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 33}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_DiscNumber              = PROPERTYKEY{GUID{Data1: 0x6afe7437, Data2: 0x9bcd, Data3: 0x49c7, Data4: [8]byte{0x80, 0xfe, 0x4a, 0x5c, 0x65, 0xfa, 0x58, 0x74}}, 104} // UInt32 -- VT_UI4
	PKEY_Music_DisplayArtist           = PROPERTYKEY{GUID{Data1: 0xfd122953, Data2: 0xfa93, Data3: 0x4ef7, Data4: [8]byte{0x92, 0xc3, 0x04, 0xc9, 0x46, 0xb2, 0xf7, 0xc8}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Genre                   = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 11}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_InitialKey              = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 34}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_IsCompilation           = PROPERTYKEY{GUID{Data1: 0xc449d5cb, Data2: 0x9ea4, Data3: 0x4809, Data4: [8]byte{0x82, 0xe8, 0xaf, 0x9d, 0x59, 0xde, 0xd6, 0xd1}}, 100} // Boolean -- VT_BOOL
	PKEY_Music_Lyrics                  = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 12}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Mood                    = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 39}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_PartOfSet               = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 37}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Period                  = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 31}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_SynchronizedLyrics      = PROPERTYKEY{GUID{Data1: 0x6b223b6a, Data2: 0x162e, Data3: 0x4aa9, Data4: [8]byte{0xb3, 0x9f, 0x05, 0xd6, 0x78, 0xfc, 0x6d, 0x77}}, 100} // Blob -- VT_BLOB
	PKEY_Music_TrackNumber             = PROPERTYKEY{GUID{Data1: 0x56a3372e, Data2: 0xce9c, Data3: 0x11d2, Data4: [8]byte{0x9f, 0x0e, 0x00, 0x60, 0x97, 0xc6, 0x86, 0xf6}}, 7}   // UInt32 -- VT_UI4

	// Note properties

	PKEY_Note_Color     = PROPERTYKEY{GUID{Data1: 0x4776cafa, Data2: 0xbce4, Data3: 0x4cb1, Data4: [8]byte{0xa2, 0x3e, 0x26, 0x5e, 0x76, 0xd8, 0xeb, 0x11}}, 100} // UInt16 -- VT_UI2
	PKEY_Note_ColorText = PROPERTYKEY{GUID{Data1: 0x46b4e8de, Data2: 0xcdb2, Data3: 0x440d, Data4: [8]byte{0x88, 0x5c, 0x16, 0x58, 0xeb, 0x65, 0xb9, 0x14}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR

	// Photo properties

	PKEY_Photo_Aperture                         = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 37378} // Double -- VT_R8
	PKEY_Photo_ApertureDenominator              = PROPERTYKEY{GUID{Data1: 0xe1a9a38b, Data2: 0x6685, Data3: 0x46bd, Data4: [8]byte{0x87, 0x5e, 0x57, 0x0d, 0xc7, 0xad, 0x73, 0x20}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_ApertureNumerator                = PROPERTYKEY{GUID{Data1: 0x0337ecec, Data2: 0x39fb, Data3: 0x4581, Data4: [8]byte{0xa0, 0xbd, 0x4c, 0x4c, 0xc5, 0x1e, 0x99, 0x14}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_Brightness                       = PROPERTYKEY{GUID{Data1: 0x1a701bf6, Data2: 0x478c, Data3: 0x4361, Data4: [8]byte{0x83, 0xab, 0x37, 0x01, 0xbb, 0x05, 0x3c, 0x58}}, 100}   // Double -- VT_R8
	PKEY_Photo_BrightnessDenominator            = PROPERTYKEY{GUID{Data1: 0x6ebe6946, Data2: 0x2321, Data3: 0x440a, Data4: [8]byte{0x90, 0xf0, 0xc0, 0x43, 0xef, 0xd3, 0x24, 0x76}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_BrightnessNumerator              = PROPERTYKEY{GUID{Data1: 0x9e7d118f, Data2: 0xb314, Data3: 0x45a0, Data4: [8]byte{0x8c, 0xfb, 0xd6, 0x54, 0xb9, 0x17, 0xc9, 0xe9}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_CameraManufacturer               = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 271}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_CameraModel                      = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 272}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_CameraSerialNumber               = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 273}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_Contrast                         = PROPERTYKEY{GUID{Data1: 0x2a785ba9, Data2: 0x8d23, Data3: 0x4ded, Data4: [8]byte{0x82, 0xe6, 0x60, 0xa3, 0x50, 0xc8, 0x6a, 0x10}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_ContrastText                     = PROPERTYKEY{GUID{Data1: 0x59dde9f2, Data2: 0x5253, Data3: 0x40ea, Data4: [8]byte{0x9a, 0x8b, 0x47, 0x9e, 0x96, 0xc6, 0x24, 0x9a}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_DateTaken                        = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 36867} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Photo_DigitalZoom                      = PROPERTYKEY{GUID{Data1: 0xf85bf840, Data2: 0xa925, Data3: 0x4bc2, Data4: [8]byte{0xb0, 0xc4, 0x8e, 0x36, 0xb5, 0x98, 0x67, 0x9e}}, 100}   // Double -- VT_R8
	PKEY_Photo_DigitalZoomDenominator           = PROPERTYKEY{GUID{Data1: 0x745baf0e, Data2: 0xe5c1, Data3: 0x4cfb, Data4: [8]byte{0x8a, 0x1b, 0xd0, 0x31, 0xa0, 0xa5, 0x23, 0x93}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_DigitalZoomNumerator             = PROPERTYKEY{GUID{Data1: 0x16cbb924, Data2: 0x6500, Data3: 0x473b, Data4: [8]byte{0xa5, 0xbe, 0xf1, 0x59, 0x9b, 0xcb, 0xe4, 0x13}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_Event                            = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 18248} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Photo_EXIFVersion                      = PROPERTYKEY{GUID{Data1: 0xd35f743a, Data2: 0xeb2e, Data3: 0x47f2, Data4: [8]byte{0xa2, 0x86, 0x84, 0x41, 0x32, 0xcb, 0x14, 0x27}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ExposureBias                     = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 37380} // Double -- VT_R8
	PKEY_Photo_ExposureBiasDenominator          = PROPERTYKEY{GUID{Data1: 0xab205e50, Data2: 0x04b7, Data3: 0x461c, Data4: [8]byte{0xa1, 0x8c, 0x2f, 0x23, 0x38, 0x36, 0xe6, 0x27}}, 100}   // Int32 -- VT_I4
	PKEY_Photo_ExposureBiasNumerator            = PROPERTYKEY{GUID{Data1: 0x738bf284, Data2: 0x1d87, Data3: 0x420b, Data4: [8]byte{0x92, 0xcf, 0x58, 0x34, 0xbf, 0x6e, 0xf9, 0xed}}, 100}   // Int32 -- VT_I4
	PKEY_Photo_ExposureIndex                    = PROPERTYKEY{GUID{Data1: 0x967b5af8, Data2: 0x995a, Data3: 0x46ed, Data4: [8]byte{0x9e, 0x11, 0x35, 0xb3, 0xc5, 0xb9, 0x78, 0x2d}}, 100}   // Double -- VT_R8
	PKEY_Photo_ExposureIndexDenominator         = PROPERTYKEY{GUID{Data1: 0x93112f89, Data2: 0xc28b, Data3: 0x492f, Data4: [8]byte{0x8a, 0x9d, 0x4b, 0xe2, 0x06, 0x2c, 0xee, 0x8a}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_ExposureIndexNumerator           = PROPERTYKEY{GUID{Data1: 0xcdedcf30, Data2: 0x8919, Data3: 0x44df, Data4: [8]byte{0x8f, 0x4c, 0x4e, 0xb2, 0xff, 0xdb, 0x8d, 0x89}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_ExposureProgram                  = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 34850} // UInt32 -- VT_UI4
	PKEY_Photo_ExposureProgramText              = PROPERTYKEY{GUID{Data1: 0xfec690b7, Data2: 0x5f30, Data3: 0x4646, Data4: [8]byte{0xae, 0x47, 0x4c, 0xaa, 0xfb, 0xa8, 0x84, 0xa3}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ExposureTime                     = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 33434} // Double -- VT_R8
	PKEY_Photo_ExposureTimeDenominator          = PROPERTYKEY{GUID{Data1: 0x55e98597, Data2: 0xad16, Data3: 0x42e0, Data4: [8]byte{0xb6, 0x24, 0x21, 0x59, 0x9a, 0x19, 0x98, 0x38}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_ExposureTimeNumerator            = PROPERTYKEY{GUID{Data1: 0x257e44e2, Data2: 0x9031, Data3: 0x4323, Data4: [8]byte{0xac, 0x38, 0x85, 0xc5, 0x52, 0x87, 0x1b, 0x2e}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_Flash                            = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 37385} // Byte -- VT_UI1
	PKEY_Photo_FlashEnergy                      = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 41483} // Double -- VT_R8
	PKEY_Photo_FlashEnergyDenominator           = PROPERTYKEY{GUID{Data1: 0xd7b61c70, Data2: 0x6323, Data3: 0x49cd, Data4: [8]byte{0xa5, 0xfc, 0xc8, 0x42, 0x77, 0x16, 0x2c, 0x97}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_FlashEnergyNumerator             = PROPERTYKEY{GUID{Data1: 0xfcad3d3d, Data2: 0x0858, Data3: 0x400f, Data4: [8]byte{0xaa, 0xa3, 0x2f, 0x66, 0xcc, 0xe2, 0xa6, 0xbc}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_FlashManufacturer                = PROPERTYKEY{GUID{Data1: 0xaabaf6c9, Data2: 0xe0c5, Data3: 0x4719, Data4: [8]byte{0x85, 0x85, 0x57, 0xb1, 0x03, 0xe5, 0x84, 0xfe}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_FlashModel                       = PROPERTYKEY{GUID{Data1: 0xfe83bb35, Data2: 0x4d1a, Data3: 0x42e2, Data4: [8]byte{0x91, 0x6b, 0x06, 0xf3, 0xe1, 0xaf, 0x71, 0x9e}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_FlashText                        = PROPERTYKEY{GUID{Data1: 0x6b8b68f6, Data2: 0x200b, Data3: 0x47ea, Data4: [8]byte{0x8d, 0x25, 0xd8, 0x05, 0x0f, 0x57, 0x33, 0x9f}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_FNumber                          = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 33437} // Double -- VT_R8
	PKEY_Photo_FNumberDenominator               = PROPERTYKEY{GUID{Data1: 0xe92a2496, Data2: 0x223b, Data3: 0x4463, Data4: [8]byte{0xa4, 0xe3, 0x30, 0xea, 0xbb, 0xa7, 0x9d, 0x80}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_FNumberNumerator                 = PROPERTYKEY{GUID{Data1: 0x1b97738a, Data2: 0xfdfc, Data3: 0x462f, Data4: [8]byte{0x9d, 0x93, 0x19, 0x57, 0xe0, 0x8b, 0xe9, 0x0c}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_FocalLength                      = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 37386} // Double -- VT_R8
	PKEY_Photo_FocalLengthDenominator           = PROPERTYKEY{GUID{Data1: 0x305bc615, Data2: 0xdca1, Data3: 0x44a5, Data4: [8]byte{0x9f, 0xd4, 0x10, 0xc0, 0xba, 0x79, 0x41, 0x2e}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_FocalLengthInFilm                = PROPERTYKEY{GUID{Data1: 0xa0e74609, Data2: 0xb84d, Data3: 0x4f49, Data4: [8]byte{0xb8, 0x60, 0x46, 0x2b, 0xd9, 0x97, 0x1f, 0x98}}, 100}   // UInt16 -- VT_UI2
	PKEY_Photo_FocalLengthNumerator             = PROPERTYKEY{GUID{Data1: 0x776b6b3b, Data2: 0x1e3d, Data3: 0x4b0c, Data4: [8]byte{0x9a, 0x0e, 0x8f, 0xba, 0xf2, 0xa8, 0x49, 0x2a}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_FocalPlaneXResolution            = PROPERTYKEY{GUID{Data1: 0xcfc08d97, Data2: 0xc6f7, Data3: 0x4484, Data4: [8]byte{0x89, 0xdd, 0xeb, 0xef, 0x43, 0x56, 0xfe, 0x76}}, 100}   // Double -- VT_R8
	PKEY_Photo_FocalPlaneXResolutionDenominator = PROPERTYKEY{GUID{Data1: 0x0933f3f5, Data2: 0x4786, Data3: 0x4f46, Data4: [8]byte{0xa8, 0xe8, 0xd6, 0x4d, 0xd3, 0x7f, 0xa5, 0x21}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_FocalPlaneXResolutionNumerator   = PROPERTYKEY{GUID{Data1: 0xdccb10af, Data2: 0xb4e2, Data3: 0x4b88, Data4: [8]byte{0x95, 0xf9, 0x03, 0x1b, 0x4d, 0x5a, 0xb4, 0x90}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_FocalPlaneYResolution            = PROPERTYKEY{GUID{Data1: 0x4fffe4d0, Data2: 0x914f, Data3: 0x4ac4, Data4: [8]byte{0x8d, 0x6f, 0xc9, 0xc6, 0x1d, 0xe1, 0x69, 0xb1}}, 100}   // Double -- VT_R8
	PKEY_Photo_FocalPlaneYResolutionDenominator = PROPERTYKEY{GUID{Data1: 0x1d6179a6, Data2: 0xa876, Data3: 0x4031, Data4: [8]byte{0xb0, 0x13, 0x33, 0x47, 0xb2, 0xb6, 0x4d, 0xc8}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_FocalPlaneYResolutionNumerator   = PROPERTYKEY{GUID{Data1: 0xa2e541c5, Data2: 0x4440, Data3: 0x4ba8, Data4: [8]byte{0x86, 0x7e, 0x75, 0xcf, 0xc0, 0x68, 0x28, 0xcd}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_GainControl                      = PROPERTYKEY{GUID{Data1: 0xfa304789, Data2: 0x00c7, Data3: 0x4d80, Data4: [8]byte{0x90, 0x4a, 0x1e, 0x4d, 0xcc, 0x72, 0x65, 0xaa}}, 100}   // Double -- VT_R8
	PKEY_Photo_GainControlDenominator           = PROPERTYKEY{GUID{Data1: 0x42864dfd, Data2: 0x9da4, Data3: 0x4f77, Data4: [8]byte{0xbd, 0xed, 0x4a, 0xad, 0x7b, 0x25, 0x67, 0x35}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_GainControlNumerator             = PROPERTYKEY{GUID{Data1: 0x8e8ecf7c, Data2: 0xb7b8, Data3: 0x4eb8, Data4: [8]byte{0xa6, 0x3f, 0x0e, 0xe7, 0x15, 0xc9, 0x6f, 0x9e}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_GainControlText                  = PROPERTYKEY{GUID{Data1: 0xc06238b2, Data2: 0x0bf9, Data3: 0x4279, Data4: [8]byte{0xa7, 0x23, 0x25, 0x85, 0x67, 0x15, 0xcb, 0x9d}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ISOSpeed                         = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 34855} // UInt16 -- VT_UI2
	PKEY_Photo_LensManufacturer                 = PROPERTYKEY{GUID{Data1: 0xe6ddcaf7, Data2: 0x29c5, Data3: 0x4f0a, Data4: [8]byte{0x9a, 0x68, 0xd1, 0x94, 0x12, 0xec, 0x70, 0x90}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_LensModel                        = PROPERTYKEY{GUID{Data1: 0xe1277516, Data2: 0x2b5f, Data3: 0x4869, Data4: [8]byte{0x89, 0xb1, 0x2e, 0x58, 0x5b, 0xd3, 0x8b, 0x7a}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_LightSource                      = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 37384} // UInt32 -- VT_UI4
	PKEY_Photo_MakerNote                        = PROPERTYKEY{GUID{Data1: 0xfa303353, Data2: 0xb659, Data3: 0x4052, Data4: [8]byte{0x85, 0xe9, 0xbc, 0xac, 0x79, 0x54, 0x9b, 0x84}}, 100}   // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Photo_MakerNoteOffset                  = PROPERTYKEY{GUID{Data1: 0x813f4124, Data2: 0x34e6, Data3: 0x4d17, Data4: [8]byte{0xab, 0x3e, 0x6b, 0x1f, 0x3c, 0x22, 0x47, 0xa1}}, 100}   // UInt64 -- VT_UI8
	PKEY_Photo_MaxAperture                      = PROPERTYKEY{GUID{Data1: 0x08f6d7c2, Data2: 0xe3f2, Data3: 0x44fc, Data4: [8]byte{0xaf, 0x1e, 0x5a, 0xa5, 0xc8, 0x1a, 0x2d, 0x3e}}, 100}   // Double -- VT_R8
	PKEY_Photo_MaxApertureDenominator           = PROPERTYKEY{GUID{Data1: 0xc77724d4, Data2: 0x601f, Data3: 0x46c5, Data4: [8]byte{0x9b, 0x89, 0xc5, 0x3f, 0x93, 0xbc, 0xeb, 0x77}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_MaxApertureNumerator             = PROPERTYKEY{GUID{Data1: 0xc107e191, Data2: 0xa459, Data3: 0x44c5, Data4: [8]byte{0x9a, 0xe6, 0xb9, 0x52, 0xad, 0x4b, 0x90, 0x6d}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_MeteringMode                     = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 37383} // UInt16 -- VT_UI2
	PKEY_Photo_MeteringModeText                 = PROPERTYKEY{GUID{Data1: 0xf628fd8c, Data2: 0x7ba8, Data3: 0x465a, Data4: [8]byte{0xa6, 0x5b, 0xc5, 0xaa, 0x79, 0x26, 0x3a, 0x9e}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_Orientation                      = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 274}   // UInt16 -- VT_UI2
	PKEY_Photo_OrientationText                  = PROPERTYKEY{GUID{Data1: 0xa9ea193c, Data2: 0xc511, Data3: 0x498a, Data4: [8]byte{0xa0, 0x6b, 0x58, 0xe2, 0x77, 0x6d, 0xcc, 0x28}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_PeopleNames                      = PROPERTYKEY{GUID{Data1: 0xe8309b6e, Data2: 0x084c, Data3: 0x49b4, Data4: [8]byte{0xb1, 0xfc, 0x90, 0xa8, 0x03, 0x31, 0xb6, 0x38}}, 100}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_Photo_PhotometricInterpretation        = PROPERTYKEY{GUID{Data1: 0x341796f1, Data2: 0x1df9, Data3: 0x4b1c, Data4: [8]byte{0xa5, 0x64, 0x91, 0xbd, 0xef, 0xa4, 0x38, 0x77}}, 100}   // UInt16 -- VT_UI2
	PKEY_Photo_PhotometricInterpretationText    = PROPERTYKEY{GUID{Data1: 0x821437d6, Data2: 0x9eab, Data3: 0x4765, Data4: [8]byte{0xa5, 0x89, 0x3b, 0x1c, 0xbb, 0xd2, 0x2a, 0x61}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ProgramMode                      = PROPERTYKEY{GUID{Data1: 0x6d217f6d, Data2: 0x3f6a, Data3: 0x4825, Data4: [8]byte{0xb4, 0x70, 0x5f, 0x03, 0xca, 0x2f, 0xbe, 0x9b}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_ProgramModeText                  = PROPERTYKEY{GUID{Data1: 0x7fe3aa27, Data2: 0x2648, Data3: 0x42f3, Data4: [8]byte{0x89, 0xb0, 0x45, 0x4e, 0x5c, 0xb1, 0x50, 0xc3}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_RelatedSoundFile                 = PROPERTYKEY{GUID{Data1: 0x318a6b45, Data2: 0x087f, Data3: 0x4dc2, Data4: [8]byte{0xb8, 0xcc, 0x05, 0x35, 0x95, 0x51, 0xfc, 0x9e}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_Saturation                       = PROPERTYKEY{GUID{Data1: 0x49237325, Data2: 0xa95a, Data3: 0x4f67, Data4: [8]byte{0xb2, 0x11, 0x81, 0x6b, 0x2d, 0x45, 0xd2, 0xe0}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_SaturationText                   = PROPERTYKEY{GUID{Data1: 0x61478c08, Data2: 0xb600, Data3: 0x4a84, Data4: [8]byte{0xbb, 0xe4, 0xe9, 0x9c, 0x45, 0xf0, 0xa0, 0x72}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_Sharpness                        = PROPERTYKEY{GUID{Data1: 0xfc6976db, Data2: 0x8349, Data3: 0x4970, Data4: [8]byte{0xae, 0x97, 0xb3, 0xc5, 0x31, 0x6a, 0x08, 0xf0}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_SharpnessText                    = PROPERTYKEY{GUID{Data1: 0x51ec3f47, Data2: 0xdd50, Data3: 0x421d, Data4: [8]byte{0x87, 0x69, 0x33, 0x4f, 0x50, 0x42, 0x4b, 0x1e}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ShutterSpeed                     = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 37377} // Double -- VT_R8
	PKEY_Photo_ShutterSpeedDenominator          = PROPERTYKEY{GUID{Data1: 0xe13d8975, Data2: 0x81c7, Data3: 0x4948, Data4: [8]byte{0xae, 0x3f, 0x37, 0xca, 0xe1, 0x1e, 0x8f, 0xf7}}, 100}   // Int32 -- VT_I4
	PKEY_Photo_ShutterSpeedNumerator            = PROPERTYKEY{GUID{Data1: 0x16ea4042, Data2: 0xd6f4, Data3: 0x4bca, Data4: [8]byte{0x83, 0x49, 0x7c, 0x78, 0xd3, 0x0f, 0xb3, 0x33}}, 100}   // Int32 -- VT_I4
	PKEY_Photo_SubjectDistance                  = PROPERTYKEY{GUID{Data1: 0x14b81da1, Data2: 0x0135, Data3: 0x4d31, Data4: [8]byte{0x96, 0xd9, 0x6c, 0xbf, 0xc9, 0x67, 0x1a, 0x99}}, 37382} // Double -- VT_R8
	PKEY_Photo_SubjectDistanceDenominator       = PROPERTYKEY{GUID{Data1: 0x0c840a88, Data2: 0xb043, Data3: 0x466d, Data4: [8]byte{0x97, 0x66, 0xd4, 0xb2, 0x6d, 0xa3, 0xfa, 0x77}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_SubjectDistanceNumerator         = PROPERTYKEY{GUID{Data1: 0x8af4961c, Data2: 0xf526, Data3: 0x43e5, Data4: [8]byte{0xaa, 0x81, 0xdb, 0x76, 0x82, 0x19, 0x17, 0x8d}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_TagViewAggregate                 = PROPERTYKEY{GUID{Data1: 0xb812f15d, Data2: 0xc2d8, Data3: 0x4bbf, Data4: [8]byte{0xba, 0xcd, 0x79, 0x74, 0x43, 0x46, 0x11, 0x3f}}, 100}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_Photo_TranscodedForSync                = PROPERTYKEY{GUID{Data1: 0x9a8ebb75, Data2: 0x6458, Data3: 0x4e82, Data4: [8]byte{0xba, 0xcb, 0x35, 0xc0, 0x09, 0x5b, 0x03, 0xbb}}, 100}   // Boolean -- VT_BOOL
	PKEY_Photo_WhiteBalance                     = PROPERTYKEY{GUID{Data1: 0xee3d3d8a, Data2: 0x5381, Data3: 0x4cfa, Data4: [8]byte{0xb1, 0x3b, 0xaa, 0xf6, 0x6b, 0x5f, 0x4e, 0xc9}}, 100}   // UInt32 -- VT_UI4
	PKEY_Photo_WhiteBalanceText                 = PROPERTYKEY{GUID{Data1: 0x6336b95e, Data2: 0xc7a7, Data3: 0x426d, Data4: [8]byte{0x86, 0xfd, 0x7a, 0xe3, 0xd3, 0x9c, 0x84, 0xb4}}, 100}   // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// PropGroup properties

	PKEY_PropGroup_Advanced      = PROPERTYKEY{GUID{Data1: 0x900a403b, Data2: 0x097b, Data3: 0x4b95, Data4: [8]byte{0x8a, 0xe2, 0x07, 0x1f, 0xda, 0xee, 0xb1, 0x18}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Audio         = PROPERTYKEY{GUID{Data1: 0x2804d469, Data2: 0x788f, Data3: 0x48aa, Data4: [8]byte{0x85, 0x70, 0x71, 0xb9, 0xc1, 0x87, 0xe1, 0x38}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Calendar      = PROPERTYKEY{GUID{Data1: 0x9973d2b5, Data2: 0xbfd8, Data3: 0x438a, Data4: [8]byte{0xba, 0x94, 0x53, 0x49, 0xb2, 0x93, 0x18, 0x1a}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Camera        = PROPERTYKEY{GUID{Data1: 0xde00de32, Data2: 0x547e, Data3: 0x4981, Data4: [8]byte{0xad, 0x4b, 0x54, 0x2f, 0x2e, 0x90, 0x07, 0xd8}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Contact       = PROPERTYKEY{GUID{Data1: 0xdf975fd3, Data2: 0x250a, Data3: 0x4004, Data4: [8]byte{0x85, 0x8f, 0x34, 0xe2, 0x9a, 0x3e, 0x37, 0xaa}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Content       = PROPERTYKEY{GUID{Data1: 0xd0dab0ba, Data2: 0x368a, Data3: 0x4050, Data4: [8]byte{0xa8, 0x82, 0x6c, 0x01, 0x0f, 0xd1, 0x9a, 0x4f}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Description   = PROPERTYKEY{GUID{Data1: 0x8969b275, Data2: 0x9475, Data3: 0x4e00, Data4: [8]byte{0xa8, 0x87, 0xff, 0x93, 0xb8, 0xb4, 0x1e, 0x44}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_FileSystem    = PROPERTYKEY{GUID{Data1: 0xe3a7d2c1, Data2: 0x80fc, Data3: 0x4b40, Data4: [8]byte{0x8f, 0x34, 0x30, 0xea, 0x11, 0x1b, 0xdc, 0x2e}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_General       = PROPERTYKEY{GUID{Data1: 0xcc301630, Data2: 0xb192, Data3: 0x4c22, Data4: [8]byte{0xb3, 0x72, 0x9f, 0x4c, 0x6d, 0x33, 0x8e, 0x07}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_GPS           = PROPERTYKEY{GUID{Data1: 0xf3713ada, Data2: 0x90e3, Data3: 0x4e11, Data4: [8]byte{0xaa, 0xe5, 0xfd, 0xc1, 0x76, 0x85, 0xb9, 0xbe}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Image         = PROPERTYKEY{GUID{Data1: 0xe3690a87, Data2: 0x0fa8, Data3: 0x4a2a, Data4: [8]byte{0x9a, 0x9f, 0xfc, 0xe8, 0x82, 0x70, 0x55, 0xac}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Media         = PROPERTYKEY{GUID{Data1: 0x61872cf7, Data2: 0x6b5e, Data3: 0x4b4b, Data4: [8]byte{0xac, 0x2d, 0x59, 0xda, 0x84, 0x45, 0x92, 0x48}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_MediaAdvanced = PROPERTYKEY{GUID{Data1: 0x8859a284, Data2: 0xde7e, Data3: 0x4642, Data4: [8]byte{0x99, 0xba, 0xd4, 0x31, 0xd0, 0x44, 0xb1, 0xec}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Message       = PROPERTYKEY{GUID{Data1: 0x7fd7259d, Data2: 0x16b4, Data3: 0x4135, Data4: [8]byte{0x9f, 0x97, 0x7c, 0x96, 0xec, 0xd2, 0xfa, 0x9e}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Music         = PROPERTYKEY{GUID{Data1: 0x68dd6094, Data2: 0x7216, Data3: 0x40f1, Data4: [8]byte{0xa0, 0x29, 0x43, 0xfe, 0x71, 0x27, 0x04, 0x3f}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Origin        = PROPERTYKEY{GUID{Data1: 0x2598d2fb, Data2: 0x5569, Data3: 0x4367, Data4: [8]byte{0x95, 0xdf, 0x5c, 0xd3, 0xa1, 0x77, 0xe1, 0xa5}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_PhotoAdvanced = PROPERTYKEY{GUID{Data1: 0x0cb2bf5a, Data2: 0x9ee7, Data3: 0x4a86, Data4: [8]byte{0x82, 0x22, 0xf0, 0x1e, 0x07, 0xfd, 0xad, 0xaf}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_RecordedTV    = PROPERTYKEY{GUID{Data1: 0xe7b33238, Data2: 0x6584, Data3: 0x4170, Data4: [8]byte{0xa5, 0xc0, 0xac, 0x25, 0xef, 0xd9, 0xda, 0x56}}, 100} // Null -- VT_NULL
	PKEY_PropGroup_Video         = PROPERTYKEY{GUID{Data1: 0xbebe0920, Data2: 0x7671, Data3: 0x4c54, Data4: [8]byte{0xa3, 0xeb, 0x49, 0xfd, 0xdf, 0xc1, 0x91, 0xee}}, 100} // Null -- VT_NULL

	// PropList properties

	PKEY_InfoTipText                       = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 17}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_ConflictPrompt           = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 11}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_ContentViewModeForBrowse = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 13}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_ContentViewModeForSearch = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 14}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_ExtendedTileInfo         = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 9}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_FileOperationPrompt      = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 10}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_FullDetails              = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 2}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_InfoTip                  = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 4}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_NonPersonal              = PROPERTYKEY{GUID{Data1: 0x49d1091f, Data2: 0x082e, Data3: 0x493f, Data4: [8]byte{0xb2, 0x3f, 0xd2, 0x30, 0x8a, 0xa9, 0x66, 0x8c}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_PreviewDetails           = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 8}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_PreviewTitle             = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 6}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_QuickTip                 = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 5}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_TileInfo                 = PROPERTYKEY{GUID{Data1: 0xc9944a21, Data2: 0xa406, Data3: 0x48fe, Data4: [8]byte{0x82, 0x25, 0xae, 0xc7, 0xe2, 0x4c, 0x21, 0x1b}}, 3}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_XPDetailsPanel           = PROPERTYKEY{GUID{Data1: 0xf2275480, Data2: 0xf782, Data3: 0x4291, Data4: [8]byte{0xbd, 0x94, 0xf1, 0x36, 0x93, 0x51, 0x3a, 0xec}}, 0}   // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// RecordedTV properties

	PKEY_RecordedTV_ChannelNumber               = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 7}   // UInt32 -- VT_UI4
	PKEY_RecordedTV_Credits                     = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 4}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_DateContentExpires          = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 15}  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_RecordedTV_EpisodeName                 = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 2}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_IsATSCContent               = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 16}  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsClosedCaptioningAvailable = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 12}  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsDTVContent                = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 17}  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsHDContent                 = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 18}  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsRepeatBroadcast           = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 13}  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsSAP                       = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 14}  // Boolean -- VT_BOOL
	PKEY_RecordedTV_NetworkAffiliation          = PROPERTYKEY{GUID{Data1: 0x2c53c813, Data2: 0xfb63, Data3: 0x4e22, Data4: [8]byte{0xa1, 0xab, 0x0b, 0x33, 0x1c, 0xa1, 0xe2, 0x73}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_OriginalBroadcastDate       = PROPERTYKEY{GUID{Data1: 0x4684fe97, Data2: 0x8765, Data3: 0x4842, Data4: [8]byte{0x9c, 0x13, 0xf0, 0x06, 0x44, 0x7b, 0x17, 0x8c}}, 100} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_RecordedTV_ProgramDescription          = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 3}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_RecordingTime               = PROPERTYKEY{GUID{Data1: 0xa5477f61, Data2: 0x7a82, Data3: 0x4eca, Data4: [8]byte{0x9d, 0xde, 0x98, 0xb6, 0x9b, 0x24, 0x79, 0xb3}}, 100} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_RecordedTV_StationCallSign             = PROPERTYKEY{GUID{Data1: 0x6d748de2, Data2: 0x8d38, Data3: 0x4cc3, Data4: [8]byte{0xac, 0x60, 0xf0, 0x09, 0xb0, 0x57, 0xc5, 0x57}}, 5}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_StationName                 = PROPERTYKEY{GUID{Data1: 0x1b5439e7, Data2: 0xeba1, Data3: 0x4af8, Data4: [8]byte{0xbd, 0xd7, 0x7a, 0xf1, 0xd4, 0x54, 0x94, 0x93}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Search properties

	PKEY_Search_AutoSummary                     = PROPERTYKEY{GUID{Data1: 0x560c36c0, Data2: 0x503a, Data3: 0x11cf, Data4: [8]byte{0xba, 0xa1, 0x00, 0x00, 0x4c, 0x75, 0x2a, 0x9a}}, 2}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_ContainerHash                   = PROPERTYKEY{GUID{Data1: 0xbceee283, Data2: 0x35df, Data3: 0x4d53, Data4: [8]byte{0x82, 0x6a, 0xf3, 0x6a, 0x3e, 0xef, 0xc6, 0xbe}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_Contents                        = PROPERTYKEY{GUID{Data1: 0xb725f130, Data2: 0x47ef, Data3: 0x101a, Data4: [8]byte{0xa5, 0xf1, 0x02, 0x60, 0x8c, 0x9e, 0xeb, 0xac}}, 19}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_EntryID                         = PROPERTYKEY{GUID{Data1: 0x49691c90, Data2: 0x7e17, Data3: 0x101a, Data4: [8]byte{0xa9, 0x1c, 0x08, 0x00, 0x2b, 0x2e, 0xcd, 0xa9}}, 5}   // Int32 -- VT_I4
	PKEY_Search_ExtendedProperties              = PROPERTYKEY{GUID{Data1: 0x7b03b546, Data2: 0xfa4f, Data3: 0x4a52, Data4: [8]byte{0xa2, 0xfe, 0x03, 0xd5, 0x31, 0x1e, 0x58, 0x65}}, 100} // Blob -- VT_BLOB
	PKEY_Search_GatherTime                      = PROPERTYKEY{GUID{Data1: 0x0b63e350, Data2: 0x9ccc, Data3: 0x11d0, Data4: [8]byte{0xbc, 0xdb, 0x00, 0x80, 0x5f, 0xcc, 0xce, 0x04}}, 8}   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Search_HitCount                        = PROPERTYKEY{GUID{Data1: 0x49691c90, Data2: 0x7e17, Data3: 0x101a, Data4: [8]byte{0xa9, 0x1c, 0x08, 0x00, 0x2b, 0x2e, 0xcd, 0xa9}}, 4}   // Int32 -- VT_I4
	PKEY_Search_IsClosedDirectory               = PROPERTYKEY{GUID{Data1: 0x0b63e343, Data2: 0x9ccc, Data3: 0x11d0, Data4: [8]byte{0xbc, 0xdb, 0x00, 0x80, 0x5f, 0xcc, 0xce, 0x04}}, 23}  // Boolean -- VT_BOOL
	PKEY_Search_IsFullyContained                = PROPERTYKEY{GUID{Data1: 0x0b63e343, Data2: 0x9ccc, Data3: 0x11d0, Data4: [8]byte{0xbc, 0xdb, 0x00, 0x80, 0x5f, 0xcc, 0xce, 0x04}}, 24}  // Boolean -- VT_BOOL
	PKEY_Search_QueryFocusedSummary             = PROPERTYKEY{GUID{Data1: 0x560c36c0, Data2: 0x503a, Data3: 0x11cf, Data4: [8]byte{0xba, 0xa1, 0x00, 0x00, 0x4c, 0x75, 0x2a, 0x9a}}, 3}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_QueryFocusedSummaryWithFallback = PROPERTYKEY{GUID{Data1: 0x560c36c0, Data2: 0x503a, Data3: 0x11cf, Data4: [8]byte{0xba, 0xa1, 0x00, 0x00, 0x4c, 0x75, 0x2a, 0x9a}}, 4}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_QueryPropertyHits               = PROPERTYKEY{GUID{Data1: 0x49691c90, Data2: 0x7e17, Data3: 0x101a, Data4: [8]byte{0xa9, 0x1c, 0x08, 0x00, 0x2b, 0x2e, 0xcd, 0xa9}}, 21}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Search_Rank                            = PROPERTYKEY{GUID{Data1: 0x49691c90, Data2: 0x7e17, Data3: 0x101a, Data4: [8]byte{0xa9, 0x1c, 0x08, 0x00, 0x2b, 0x2e, 0xcd, 0xa9}}, 3}   // Int32 -- VT_I4
	PKEY_Search_Store                           = PROPERTYKEY{GUID{Data1: 0xa06992b3, Data2: 0x8caf, Data3: 0x4ed7, Data4: [8]byte{0xa5, 0x47, 0xb2, 0x59, 0xe3, 0x2a, 0xc9, 0xfc}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_UrlToIndex                      = PROPERTYKEY{GUID{Data1: 0x0b63e343, Data2: 0x9ccc, Data3: 0x11d0, Data4: [8]byte{0xbc, 0xdb, 0x00, 0x80, 0x5f, 0xcc, 0xce, 0x04}}, 2}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_UrlToIndexWithModificationTime  = PROPERTYKEY{GUID{Data1: 0x0b63e343, Data2: 0x9ccc, Data3: 0x11d0, Data4: [8]byte{0xbc, 0xdb, 0x00, 0x80, 0x5f, 0xcc, 0xce, 0x04}}, 12}  // Multivalue Any -- VT_VECTOR | VT_NULL  (For variants: VT_ARRAY | VT_NULL)
	PKEY_Supplemental_Album                     = PROPERTYKEY{GUID{Data1: 0x0c73b141, Data2: 0x39d6, Data3: 0x4653, Data4: [8]byte{0xa6, 0x83, 0xca, 0xb2, 0x91, 0xea, 0xf9, 0x5b}}, 6}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Supplemental_AlbumID                   = PROPERTYKEY{GUID{Data1: 0x0c73b141, Data2: 0x39d6, Data3: 0x4653, Data4: [8]byte{0xa6, 0x83, 0xca, 0xb2, 0x91, 0xea, 0xf9, 0x5b}}, 2}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Supplemental_Location                  = PROPERTYKEY{GUID{Data1: 0x0c73b141, Data2: 0x39d6, Data3: 0x4653, Data4: [8]byte{0xa6, 0x83, 0xca, 0xb2, 0x91, 0xea, 0xf9, 0x5b}}, 5}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Supplemental_Person                    = PROPERTYKEY{GUID{Data1: 0x0c73b141, Data2: 0x39d6, Data3: 0x4653, Data4: [8]byte{0xa6, 0x83, 0xca, 0xb2, 0x91, 0xea, 0xf9, 0x5b}}, 7}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Supplemental_ResourceId                = PROPERTYKEY{GUID{Data1: 0x0c73b141, Data2: 0x39d6, Data3: 0x4653, Data4: [8]byte{0xa6, 0x83, 0xca, 0xb2, 0x91, 0xea, 0xf9, 0x5b}}, 3}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Supplemental_Tag                       = PROPERTYKEY{GUID{Data1: 0x0c73b141, Data2: 0x39d6, Data3: 0x4653, Data4: [8]byte{0xa6, 0x83, 0xca, 0xb2, 0x91, 0xea, 0xf9, 0x5b}}, 4}   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)

	// Shell properties

	PKEY_DescriptionID                = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 2} // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_InternalName                 = PROPERTYKEY{GUID{Data1: 0x0cef7d53, Data2: 0xfa64, Data3: 0x11d1, Data4: [8]byte{0xa2, 0x03, 0x00, 0x00, 0xf8, 0x1f, 0xed, 0xee}}, 5} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_LibraryLocationsCount        = PROPERTYKEY{GUID{Data1: 0x908696c7, Data2: 0x8f87, Data3: 0x44f2, Data4: [8]byte{0x80, 0xed, 0xa8, 0xc1, 0xc6, 0x89, 0x45, 0x75}}, 2} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_TargetSFGAOFlagsStrings = PROPERTYKEY{GUID{Data1: 0xd6942081, Data2: 0xd53b, Data3: 0x443d, Data4: [8]byte{0xad, 0x47, 0x5e, 0x05, 0x9d, 0x9c, 0xd2, 0x7a}}, 3} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Link_TargetUrl               = PROPERTYKEY{GUID{Data1: 0x5cbf2787, Data2: 0x48cf, Data3: 0x4208, Data4: [8]byte{0xb9, 0x0e, 0xee, 0x5e, 0x5d, 0x42, 0x02, 0x94}}, 2} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_NamespaceCLSID               = PROPERTYKEY{GUID{Data1: 0x28636aa6, Data2: 0x953d, Data3: 0x11d2, Data4: [8]byte{0xb5, 0xd6, 0x00, 0xc0, 0x4f, 0xd9, 0x18, 0xd0}}, 6} // Guid -- VT_CLSID
	PKEY_Shell_SFGAOFlagsStrings      = PROPERTYKEY{GUID{Data1: 0xd6942081, Data2: 0xd53b, Data3: 0x443d, Data4: [8]byte{0xad, 0x47, 0x5e, 0x05, 0x9d, 0x9c, 0xd2, 0x7a}}, 2} // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_StatusBarSelectedItemCount   = PROPERTYKEY{GUID{Data1: 0x26dc287c, Data2: 0x6e3d, Data3: 0x4bd3, Data4: [8]byte{0xb2, 0xb0, 0x6a, 0x26, 0xba, 0x2e, 0x34, 0x6d}}, 3} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StatusBarViewItemCount       = PROPERTYKEY{GUID{Data1: 0x26dc287c, Data2: 0x6e3d, Data3: 0x4bd3, Data4: [8]byte{0xb2, 0xb0, 0x6a, 0x26, 0xba, 0x2e, 0x34, 0x6d}}, 2} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Software properties

	PKEY_AppUserModel_ExcludeFromShowInNewInstall    = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 8}  // Boolean -- VT_BOOL
	PKEY_AppUserModel_ID                             = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 5}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_AppUserModel_IsDestListSeparator            = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 6}  // Boolean -- VT_BOOL
	PKEY_AppUserModel_IsDualMode                     = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 11} // Boolean -- VT_BOOL
	PKEY_AppUserModel_PreventPinning                 = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 9}  // Boolean -- VT_BOOL
	PKEY_AppUserModel_RelaunchCommand                = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 2}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_AppUserModel_RelaunchDisplayNameResource    = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 4}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_AppUserModel_RelaunchIconResource           = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 3}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_AppUserModel_StartPinOption                 = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 12} // UInt32 -- VT_UI4
	PKEY_AppUserModel_ToastActivatorCLSID            = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 26} // Guid -- VT_CLSID
	PKEY_AppUserModel_VisualElementsManifestHintPath = PROPERTYKEY{GUID{Data1: 0x9f4c2855, Data2: 0x9f79, Data3: 0x4b39, Data4: [8]byte{0xa8, 0xd0, 0xe1, 0xd4, 0x2d, 0xe1, 0xd5, 0xf3}}, 31} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_EdgeGesture_DisableTouchWhenFullscreen      = PROPERTYKEY{GUID{Data1: 0x32ce38b2, Data2: 0x2c9a, Data3: 0x41b1, Data4: [8]byte{0x9b, 0xc5, 0xb3, 0x78, 0x43, 0x94, 0xaa, 0x44}}, 2}  // Boolean -- VT_BOOL
	PKEY_Software_DateLastUsed                       = PROPERTYKEY{GUID{Data1: 0x841e4f90, Data2: 0xff59, Data3: 0x4d16, Data4: [8]byte{0x89, 0x47, 0xe8, 0x1b, 0xbf, 0xfa, 0xb3, 0x6d}}, 16} // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Software_ProductName                        = PROPERTYKEY{GUID{Data1: 0x0cef7d53, Data2: 0xfa64, Data3: 0x11d1, Data4: [8]byte{0xa2, 0x03, 0x00, 0x00, 0xf8, 0x1f, 0xed, 0xee}}, 7}  // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Sync properties

	PKEY_Sync_Comments               = PROPERTYKEY{GUID{Data1: 0x7bd5533e, Data2: 0xaf15, Data3: 0x44db, Data4: [8]byte{0xb8, 0xc8, 0xbd, 0x66, 0x24, 0xe1, 0xd0, 0x32}}, 13} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ConflictDescription    = PROPERTYKEY{GUID{Data1: 0xce50c159, Data2: 0x2fb8, Data3: 0x41fd, Data4: [8]byte{0xbe, 0x68, 0xd3, 0xe0, 0x42, 0xe2, 0x74, 0xbc}}, 4}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ConflictFirstLocation  = PROPERTYKEY{GUID{Data1: 0xce50c159, Data2: 0x2fb8, Data3: 0x41fd, Data4: [8]byte{0xbe, 0x68, 0xd3, 0xe0, 0x42, 0xe2, 0x74, 0xbc}}, 6}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ConflictSecondLocation = PROPERTYKEY{GUID{Data1: 0xce50c159, Data2: 0x2fb8, Data3: 0x41fd, Data4: [8]byte{0xbe, 0x68, 0xd3, 0xe0, 0x42, 0xe2, 0x74, 0xbc}}, 7}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_HandlerCollectionID    = PROPERTYKEY{GUID{Data1: 0x7bd5533e, Data2: 0xaf15, Data3: 0x44db, Data4: [8]byte{0xb8, 0xc8, 0xbd, 0x66, 0x24, 0xe1, 0xd0, 0x32}}, 2}  // Guid -- VT_CLSID
	PKEY_Sync_HandlerID              = PROPERTYKEY{GUID{Data1: 0x7bd5533e, Data2: 0xaf15, Data3: 0x44db, Data4: [8]byte{0xb8, 0xc8, 0xbd, 0x66, 0x24, 0xe1, 0xd0, 0x32}}, 3}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_HandlerName            = PROPERTYKEY{GUID{Data1: 0xce50c159, Data2: 0x2fb8, Data3: 0x41fd, Data4: [8]byte{0xbe, 0x68, 0xd3, 0xe0, 0x42, 0xe2, 0x74, 0xbc}}, 2}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_HandlerType            = PROPERTYKEY{GUID{Data1: 0x7bd5533e, Data2: 0xaf15, Data3: 0x44db, Data4: [8]byte{0xb8, 0xc8, 0xbd, 0x66, 0x24, 0xe1, 0xd0, 0x32}}, 8}  // UInt32 -- VT_UI4
	PKEY_Sync_HandlerTypeLabel       = PROPERTYKEY{GUID{Data1: 0x7bd5533e, Data2: 0xaf15, Data3: 0x44db, Data4: [8]byte{0xb8, 0xc8, 0xbd, 0x66, 0x24, 0xe1, 0xd0, 0x32}}, 9}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ItemID                 = PROPERTYKEY{GUID{Data1: 0x7bd5533e, Data2: 0xaf15, Data3: 0x44db, Data4: [8]byte{0xb8, 0xc8, 0xbd, 0x66, 0x24, 0xe1, 0xd0, 0x32}}, 6}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ItemName               = PROPERTYKEY{GUID{Data1: 0xce50c159, Data2: 0x2fb8, Data3: 0x41fd, Data4: [8]byte{0xbe, 0x68, 0xd3, 0xe0, 0x42, 0xe2, 0x74, 0xbc}}, 3}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ProgressPercentage     = PROPERTYKEY{GUID{Data1: 0x7bd5533e, Data2: 0xaf15, Data3: 0x44db, Data4: [8]byte{0xb8, 0xc8, 0xbd, 0x66, 0x24, 0xe1, 0xd0, 0x32}}, 23} // UInt32 -- VT_UI4
	PKEY_Sync_State                  = PROPERTYKEY{GUID{Data1: 0x7bd5533e, Data2: 0xaf15, Data3: 0x44db, Data4: [8]byte{0xb8, 0xc8, 0xbd, 0x66, 0x24, 0xe1, 0xd0, 0x32}}, 24} // UInt32 -- VT_UI4
	PKEY_Sync_Status                 = PROPERTYKEY{GUID{Data1: 0x7bd5533e, Data2: 0xaf15, Data3: 0x44db, Data4: [8]byte{0xb8, 0xc8, 0xbd, 0x66, 0x24, 0xe1, 0xd0, 0x32}}, 10} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Task properties

	PKEY_Task_BillingInformation = PROPERTYKEY{GUID{Data1: 0xd37d52c6, Data2: 0x261c, Data3: 0x4303, Data4: [8]byte{0x82, 0xb3, 0x08, 0xb9, 0x26, 0xac, 0x6f, 0x12}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Task_CompletionStatus   = PROPERTYKEY{GUID{Data1: 0x084d8a0a, Data2: 0xe6d5, Data3: 0x40de, Data4: [8]byte{0xbf, 0x1f, 0xc8, 0x82, 0x0e, 0x7c, 0x87, 0x7c}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Task_Owner              = PROPERTYKEY{GUID{Data1: 0x08c7cc5f, Data2: 0x60f2, Data3: 0x4494, Data4: [8]byte{0xad, 0x75, 0x55, 0xe3, 0xe0, 0xb5, 0xad, 0xd0}}, 100} // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Video properties

	PKEY_Video_Compression           = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 10}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Video_Director              = PROPERTYKEY{GUID{Data1: 0x64440492, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 20}  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Video_EncodingBitrate       = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 8}   // UInt32 -- VT_UI4
	PKEY_Video_FourCC                = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 44}  // UInt32 -- VT_UI4
	PKEY_Video_FrameHeight           = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 4}   // UInt32 -- VT_UI4
	PKEY_Video_FrameRate             = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 6}   // UInt32 -- VT_UI4
	PKEY_Video_FrameWidth            = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 3}   // UInt32 -- VT_UI4
	PKEY_Video_HorizontalAspectRatio = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 42}  // UInt32 -- VT_UI4
	PKEY_Video_IsSpherical           = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 100} // Boolean -- VT_BOOL
	PKEY_Video_IsStereo              = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 98}  // Boolean -- VT_BOOL
	PKEY_Video_Orientation           = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 99}  // UInt32 -- VT_UI4
	PKEY_Video_SampleSize            = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 9}   // UInt32 -- VT_UI4
	PKEY_Video_StreamName            = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 2}   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Video_StreamNumber          = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 11}  // UInt16 -- VT_UI2
	PKEY_Video_TotalBitrate          = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 43}  // UInt32 -- VT_UI4
	PKEY_Video_TranscodedForSync     = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 46}  // Boolean -- VT_BOOL
	PKEY_Video_VerticalAspectRatio   = PROPERTYKEY{GUID{Data1: 0x64440491, Data2: 0x4c8b, Data3: 0x11d1, Data4: [8]byte{0x8b, 0x70, 0x08, 0x00, 0x36, 0xb1, 0x1a, 0x03}}, 45}  // UInt32 -- VT_UI4

	// Volume properties

	PKEY_Volume_FileSystem    = PROPERTYKEY{GUID{Data1: 0x9b174b35, Data2: 0x40ff, Data3: 0x11d2, Data4: [8]byte{0xa2, 0x7e, 0x00, 0xc0, 0x4f, 0xc3, 0x08, 0x71}}, 4}  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Volume_IsMappedDrive = PROPERTYKEY{GUID{Data1: 0x149c0b69, Data2: 0x2c2d, Data3: 0x48fc, Data4: [8]byte{0x80, 0x8f, 0xd3, 0x18, 0xd7, 0x8c, 0x46, 0x36}}, 2}  // Boolean -- VT_BOOL
	PKEY_Volume_IsRoot        = PROPERTYKEY{GUID{Data1: 0x9b174b35, Data2: 0x40ff, Data3: 0x11d2, Data4: [8]byte{0xa2, 0x7e, 0x00, 0xc0, 0x4f, 0xc3, 0x08, 0x71}}, 10} // Boolean -- VT_BOOL
)

// [SHFILEINFO] dwAttributes.
//
// [SHFILEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileinfow
type SFGAO uint32

const (
	_DROPEFFECT_NONE   SFGAO = 0
	_DROPEFFECT_COPY   SFGAO = 1
	_DROPEFFECT_MOVE   SFGAO = 2
	_DROPEFFECT_LINK   SFGAO = 4
	_DROPEFFECT_SCROLL SFGAO = 0x8000_0000

	SFGAO_CANCOPY               = _DROPEFFECT_COPY
	SFGAO_CANMOVE               = _DROPEFFECT_MOVE
	SFGAO_CANLINK               = _DROPEFFECT_LINK
	SFGAO_STORAGE         SFGAO = 0x0000_0008
	SFGAO_CANRENAME       SFGAO = 0x0000_0010
	SFGAO_CANDELETE       SFGAO = 0x0000_0020
	SFGAO_HASPROPSHEET    SFGAO = 0x0000_0040
	SFGAO_DROPTARGET      SFGAO = 0x0000_0100
	SFGAO_CAPABILITYMASK  SFGAO = 0x0000_0177
	SFGAO_PLACEHOLDER     SFGAO = 0x0000_0800
	SFGAO_SYSTEM          SFGAO = 0x0000_1000
	SFGAO_ENCRYPTED       SFGAO = 0x0000_2000
	SFGAO_ISSLOW          SFGAO = 0x0000_4000
	SFGAO_GHOSTED         SFGAO = 0x0000_8000
	SFGAO_LINK            SFGAO = 0x0001_0000
	SFGAO_SHARE           SFGAO = 0x0002_0000
	SFGAO_READONLY        SFGAO = 0x0004_0000
	SFGAO_HIDDEN          SFGAO = 0x0008_0000
	SFGAO_DISPLAYATTRMASK SFGAO = 0x000f_c000
	SFGAO_FILESYSANCESTOR SFGAO = 0x1000_0000
	SFGAO_FOLDER          SFGAO = 0x2000_0000
	SFGAO_FILESYSTEM      SFGAO = 0x4000_0000
	SFGAO_HASSUBFOLDER    SFGAO = 0x8000_0000
	SFGAO_CONTENTSMASK    SFGAO = 0x8000_0000
	SFGAO_VALIDATE        SFGAO = 0x0100_0000
	SFGAO_REMOVABLE       SFGAO = 0x0200_0000
	SFGAO_COMPRESSED      SFGAO = 0x0400_0000
	SFGAO_BROWSABLE       SFGAO = 0x0800_0000
	SFGAO_NONENUMERATED   SFGAO = 0x0010_0000
	SFGAO_NEWCONTENT      SFGAO = 0x0020_0000
	SFGAO_CANMONIKER      SFGAO = 0x0040_0000
	SFGAO_HASSTORAGE      SFGAO = 0x0040_0000
	SFGAO_STREAM          SFGAO = 0x0040_0000
	SFGAO_STORAGEANCESTOR SFGAO = 0x0080_0000
	SFGAO_STORAGECAPMASK  SFGAO = 0x70c5_0008
	SFGAO_PKEYSFGAOMASK   SFGAO = 0x8104_4000
)

// [IShellFolder.CompareIDs] lParam.
//
// [IShellFolder.CompareIDs]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-compareids
type SHCIDS uint32

const (
	SHCIDS_ALLFIELDS     SHCIDS = 0x8000_0000
	SHCIDS_CANONICALONLY SHCIDS = 0x1000_0000
)

// [SHCONTF] enumeration.
//
// [SHCONTF]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-_shcontf
type SHCONTF uint32

const (
	SHCONTF_CHECKING_FOR_CHILDREN SHCONTF = 0x10
	SHCONTF_FOLDERS               SHCONTF = 0x20
	SHCONTF_NONFOLDERS            SHCONTF = 0x40
	SHCONTF_INCLUDEHIDDEN         SHCONTF = 0x80
	SHCONTF_INIT_ON_FIRST_NEXT    SHCONTF = 0x100
	SHCONTF_NETPRINTERSRCH        SHCONTF = 0x200
	SHCONTF_SHAREABLE             SHCONTF = 0x400
	SHCONTF_STORAGE               SHCONTF = 0x800
	SHCONTF_NAVIGATION_ENUM       SHCONTF = 0x1000
	SHCONTF_FASTITEMS             SHCONTF = 0x2000
	SHCONTF_FLATLIST              SHCONTF = 0x4000
	SHCONTF_ENABLE_ASYNC          SHCONTF = 0x8000
	SHCONTF_INCLUDESUPERHIDDEN    SHCONTF = 0x1_0000
)

// [_SHGDNF] enumeration.
//
// [_SHGDNF]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-_shgdnf
type SHGDN uint32

const (
	SHGDN_NORMAL        SHGDN = 0
	SHGDN_INFOLDER      SHGDN = 0x1
	SHGDN_FOREDITING    SHGDN = 0x1000
	SHGDN_FORADDRESSBAR SHGDN = 0x4000
	SHGDN_FORPARSING    SHGDN = 0x8000
)

// [SHGetFileInfo] uFlags.
//
// [SHGetFileInfo]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
type SHGFI uint32

const (
	SHGFI_NONE              SHGFI = 0
	SHGFI_ICON              SHGFI = 0x0000_0100
	SHGFI_DISPLAYNAME       SHGFI = 0x0000_0200
	SHGFI_TYPENAME          SHGFI = 0x0000_0400
	SHGFI_ATTRIBUTES        SHGFI = 0x0000_0800
	SHGFI_ICONLOCATION      SHGFI = 0x0000_1000
	SHGFI_EXETYPE           SHGFI = 0x0000_2000
	SHGFI_SYSICONINDEX      SHGFI = 0x0000_4000
	SHGFI_LINKOVERLAY       SHGFI = 0x0000_8000
	SHGFI_SELECTED          SHGFI = 0x0001_0000
	SHGFI_ATTR_SPECIFIED    SHGFI = 0x0002_0000
	SHGFI_LARGEICON         SHGFI = 0x0000_0000
	SHGFI_SMALLICON         SHGFI = 0x0000_0001
	SHGFI_OPENICON          SHGFI = 0x0000_0002
	SHGFI_SHELLICONSIZE     SHGFI = 0x0000_0004
	SHGFI_PIDL              SHGFI = 0x0000_0008
	SHGFI_USEFILEATTRIBUTES SHGFI = 0x0000_0010
	SHGFI_ADDOVERLAYS       SHGFI = 0x0000_0020
	SHGFI_OVERLAYINDEX      SHGFI = 0x0000_0040
)

// [_SICHINTF] enumeration.
//
// [_SICHINTF]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-_sichintf
type SICHINT uint32

const (
	SICHINT_DISPLAY                       SICHINT = 0
	SICHINT_ALLFIELDS                     SICHINT = 0x8000_0000
	SICHINT_CANONICAL                     SICHINT = 0x1000_0000
	SICHINT_TEST_FILESYSPATH_IF_NOT_EQUAL SICHINT = 0x2000_0000
)

// [SIGDN] enumeration.
//
// [SIGDN]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-sigdn
type SIGDN uint32

const (
	SIGDN_NORMALDISPLAY               SIGDN = 0
	SIGDN_PARENTRELATIVEPARSING       SIGDN = 0x8001_8001
	SIGDN_DESKTOPABSOLUTEPARSING      SIGDN = 0x8002_8000
	SIGDN_PARENTRELATIVEEDITING       SIGDN = 0x8003_1001
	SIGDN_DESKTOPABSOLUTEEDITING      SIGDN = 0x8004_c000
	SIGDN_FILESYSPATH                 SIGDN = 0x8005_8000
	SIGDN_URL                         SIGDN = 0x8006_8000
	SIGDN_PARENTRELATIVEFORADDRESSBAR SIGDN = 0x8007_c001
	SIGDN_PARENTRELATIVE              SIGDN = 0x8008_0001
	SIGDN_PARENTRELATIVEFORUI         SIGDN = 0x8009_4001
)

// [IShellLink.GetPath] flags.
//
// [IShellLink.GetPath]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getpath
type SLGP uint32

const (
	SLGP_SHORTPATH        SLGP = 0x1
	SLGP_UNCPRIORITY      SLGP = 0x2
	SLGP_RAWPATH          SLGP = 0x4
	SLGP_RELATIVEPRIORITY SLGP = 0x8
)

// [STPFLAG] enumeration.
//
// [STPFLAG]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-stpflag
type STPFLAG uint32

const (
	STPFLAG_NONE                      STPFLAG = 0
	STPFLAG_USEAPPTHUMBNAILALWAYS     STPFLAG = 0x1
	STPFLAG_USEAPPTHUMBNAILWHENACTIVE STPFLAG = 0x2
	STPFLAG_USEAPPPEEKALWAYS          STPFLAG = 0x4
	STPFLAG_USEAPPPEEKWHENACTIVE      STPFLAG = 0x8
)

// [IShellView.UIActivate] state.
//
// [IShellView.UIActivate]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellview-uiactivate
type SVUIA uint32

const (
	SVUIA_DEACTIVATE SVUIA = iota
	SVUIA_ACTIVATE_NOFOCUS
	SVUIA_ACTIVATE_FOCUS
	SVUIA_INPLACEACTIVATE
)

// [ITaskbarList3.SetProgressState] tbpFlags.
//
// [ITaskbarList3.SetProgressState]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
type TBPF uint32

const (
	// Stops displaying progress and returns the button to its normal state.
	// Call the method with this flag to dismiss the progress bar when the
	// operation is complete or canceled.
	TBPF_NOPROGRESS TBPF = 0
	// The progress indicator does not grow in size, but cycles repeatedly along
	// the length of the taskbar button. This indicates activity without
	// specifying what proportion of the progress is complete. Progress is
	// taking place, but there is no prediction as to how long the operation
	// will take.
	TBPF_INDETERMINATE TBPF = 0x1
	// The progress indicator grows in size from left to right in proportion to
	// the estimated amount of the operation completed. This is a determinate
	// progress indicator; a prediction is being made as to the duration of the
	// operation.
	TBPF_NORMAL TBPF = 0x2
	// The progress indicator turns red to show that an error has occurred in
	// one of the windows that is broadcasting progress. This is a determinate
	// state. If the progress indicator is in the indeterminate state, it
	// switches to a red determinate display of a generic percentage not
	// indicative of actual progress.
	TBPF_ERROR TBPF = 0x4
	// The progress indicator turns yellow to show that progress is currently
	// stopped in one of the windows but can be resumed by the user. No error
	// condition exists and nothing is preventing the progress from continuing.
	// This is a determinate state. If the progress indicator is in the
	// indeterminate state, it switches to a yellow determinate display of a
	// generic percentage not indicative of actual progress.
	TBPF_PAUSED TBPF = 0x8
)

// [THUMBBUTTONMASK] enumeration.
//
// [THUMBBUTTONMASK]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-thumbbuttonmask
type THB uint32

const (
	THB_BITMAP  THB = 0x1
	THB_ICON    THB = 0x2
	THB_TOOLTIP THB = 0x4
	THB_FLAGS   THB = 0x8
)

// [THUMBBUTTONFLAGS] enumeration.
//
// [THUMBBUTTONFLAGS]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-thumbbuttonflags
type THBF uint32

const (
	THBF_ENABLED        THBF = 0
	THBF_DISABLED       THBF = 0x1
	THBF_DISMISSONCLICK THBF = 0x2
	THBF_NOBACKGROUND   THBF = 0x4
	THBF_HIDDEN         THBF = 0x8
	THBF_NONINTERACTIVE THBF = 0x10
)

// [_TRANSFER_SOURCE_FLAGS] enumeration.
//
// [_TRANSFER_SOURCE_FLAGS]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-_transfer_source_flags
type TSF uint32

const (
	TSF_NORMAL                     TSF = 0
	TSF_FAIL_EXIST                 TSF = 0
	TSF_RENAME_EXIST               TSF = 0x1
	TSF_OVERWRITE_EXIST            TSF = 0x2
	TSF_ALLOW_DECRYPTION           TSF = 0x4
	TSF_NO_SECURITY                TSF = 0x8
	TSF_COPY_CREATION_TIME         TSF = 0x10
	TSF_COPY_WRITE_TIME            TSF = 0x20
	TSF_USE_FULL_ACCESS            TSF = 0x40
	TSF_DELETE_RECYCLE_IF_POSSIBLE TSF = 0x80
	TSF_COPY_HARD_LINK             TSF = 0x100
	TSF_COPY_LOCALIZED_NAME        TSF = 0x200
	TSF_MOVE_AS_COPY_DELETE        TSF = 0x400
	TSF_SUSPEND_SHELLEVENTS        TSF = 0x800
)
