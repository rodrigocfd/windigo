//go:build windows

package co

const (
	CLSID_FileOpenDialog CLSID = "dc1c5a9c-e88a-4dde-a5a1-60f82a20aef7"
	CLSID_FileOperation  CLSID = "3ad05575-8857-4850-9277-11b85bdb8e09"
	CLSID_FileSaveDialog CLSID = "c0b4e2f3-ba21-4773-8dba-335ec946eb8b"
	CLSID_ShellLink      CLSID = "00021401-0000-0000-c000-000000000046"
	CLSID_TaskbarList    CLSID = "56fdf344-fd6d-11d0-958a-006097c9a090"

	IID_IEnumShellItems            IID = "70629033-e363-4a28-a567-0db78006e6d7"
	IID_IFileDialog                IID = "42f85136-db7e-439c-85f1-e4075d135fc8"
	IID_IFileDialogEvents          IID = "973510db-7d7f-452b-8975-74a85828d354"
	IID_IFileOpenDialog            IID = "d57c7288-d4ad-4768-be02-9d969532d960"
	IID_IFileOperation             IID = "947aab5f-0a5c-4c13-b4d6-4bf7836fc9f8"
	IID_IFileOperationProgressSink IID = "04b0f1a7-9490-44bc-96e1-4296a31252e2"
	IID_IFileSaveDialog            IID = "84bccd23-5fde-4cdb-aea4-af64b83d78ab"
	IID_IModalWindow               IID = "b4db1657-70d7-485e-8e3e-6fcb5a5c1802"
	IID_IOleWindow                 IID = "00000114-0000-0000-c000-000000000046"
	IID_IShellFolder               IID = "000214e6-0000-0000-c000-000000000046"
	IID_IShellItem                 IID = "43826d1e-e718-42ee-bc55-a1e261c37bfe"
	IID_IShellItem2                IID = "7e9fb0d3-919f-4307-ab2e-9b1860310c93"
	IID_IShellItemArray            IID = "b63ea76d-1f85-456f-a19c-48159efa858b"
	IID_IShellItemFilter           IID = "2659b475-eeb8-48b7-8f07-b378810f48cf"
	IID_IShellLink                 IID = "000214f9-0000-0000-c000-000000000046"
	IID_IShellView                 IID = "000214e3-0000-0000-c000-000000000046"
	IID_ITaskbarList               IID = "56fdf342-fd6d-11d0-958a-006097c9a090"
	IID_ITaskbarList2              IID = "602d4995-b13a-429b-a66e-1935e44f4317"
	IID_ITaskbarList3              IID = "ea1afb91-9e28-4b86-90e9-9e9f8a5eefaf"
	IID_ITaskbarList4              IID = "c43dc798-95d1-4bea-9030-bb99e2983a1a"
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

// [KNOWNFOLDERID] constants, represented as a string.
//
// [KNOWNFOLDERID]: https://learn.microsoft.com/en-us/windows/win32/shell/knownfolderid
type FOLDERID string

const (
	FOLDERID_NetworkFolder          FOLDERID = "d20beec4-5ca8-4905-ae3b-bf251ea09b53"
	FOLDERID_ComputerFolder         FOLDERID = "0ac0837c-bbf8-452a-850d-79d08e667ca7"
	FOLDERID_InternetFolder         FOLDERID = "4d9f7874-4e0c-4904-967b-40b0d20c3e4b"
	FOLDERID_ControlPanelFolder     FOLDERID = "82a74aeb-aeb4-465c-a014-d097ee346d63"
	FOLDERID_PrintersFolder         FOLDERID = "76fc4e2d-d6ad-4519-a663-37bd56068185"
	FOLDERID_SyncManagerFolder      FOLDERID = "43668bf8-c14e-49b2-97c9-747784d784b7"
	FOLDERID_SyncSetupFolder        FOLDERID = "0f214138-b1d3-4a90-bba9-27cbc0c5389a"
	FOLDERID_ConflictFolder         FOLDERID = "4bfefb45-347d-4006-a5be-ac0cb0567192"
	FOLDERID_SyncResultsFolder      FOLDERID = "289a9a43-be44-4057-a41b-587a76d7e7f9"
	FOLDERID_RecycleBinFolder       FOLDERID = "b7534046-3ecb-4c18-be4e-64cd4cb7d6ac"
	FOLDERID_ConnectionsFolder      FOLDERID = "6f0cd92b-2e97-45d1-88ff-b0d186b8dedd"
	FOLDERID_Fonts                  FOLDERID = "fd228cb7-ae11-4ae3-864c-16f3910ab8fe"
	FOLDERID_Desktop                FOLDERID = "b4bfcc3a-db2c-424c-b029-7fe99a87c641"
	FOLDERID_Startup                FOLDERID = "b97d20bb-f46a-4c97-ba10-5e3608430854"
	FOLDERID_Programs               FOLDERID = "a77f5d77-2e2b-44c3-a6a2-aba601054a51"
	FOLDERID_StartMenu              FOLDERID = "625b53c3-ab48-4ec1-ba1f-a1ef4146fc19"
	FOLDERID_Recent                 FOLDERID = "ae50c081-ebd2-438a-8655-8a092e34987a"
	FOLDERID_SendTo                 FOLDERID = "8983036c-27c0-404b-8f08-102d10dcfd74"
	FOLDERID_Documents              FOLDERID = "fdd39ad0-238f-46af-adb4-6c85480369c7"
	FOLDERID_Favorites              FOLDERID = "1777f761-68ad-4d8a-87bd-30b759fa33dd"
	FOLDERID_NetHood                FOLDERID = "c5abbf53-e17f-4121-8900-86626fc2c973"
	FOLDERID_PrintHood              FOLDERID = "9274bd8d-cfd1-41c3-b35e-b13f55a758f4"
	FOLDERID_Templates              FOLDERID = "a63293e8-664e-48db-a079-df759e0509f7"
	FOLDERID_CommonStartup          FOLDERID = "82a5ea35-d9cd-47c5-9629-e15d2f714e6e"
	FOLDERID_CommonPrograms         FOLDERID = "0139d44e-6afe-49f2-8690-3dafcae6ffb8"
	FOLDERID_CommonStartMenu        FOLDERID = "a4115719-d62e-491d-aa7c-e74b8be3b067"
	FOLDERID_PublicDesktop          FOLDERID = "c4aa340d-f20f-4863-afef-f87ef2e6ba25"
	FOLDERID_ProgramData            FOLDERID = "62ab5d82-fdc1-4dc3-a9dd-070d1d495d97"
	FOLDERID_CommonTemplates        FOLDERID = "b94237e7-57ac-4347-9151-b08c6c32d1f7"
	FOLDERID_PublicDocuments        FOLDERID = "ed4824af-dce4-45a8-81e2-fc7965083634"
	FOLDERID_RoamingAppData         FOLDERID = "3eb685db-65f9-4cf6-a03a-e3ef65729f3d"
	FOLDERID_LocalAppData           FOLDERID = "f1b32785-6fba-4fcf-9d55-7b8e7f157091"
	FOLDERID_LocalAppDataLow        FOLDERID = "a520a1a4-1780-4ff6-bd18-167343c5af16"
	FOLDERID_InternetCache          FOLDERID = "352481e8-33be-4251-ba85-6007caedcf9d"
	FOLDERID_Cookies                FOLDERID = "2b0f765d-c0e9-4171-908e-08a611b84ff6"
	FOLDERID_History                FOLDERID = "d9dc8a3b-b784-432e-a781-5a1130a75963"
	FOLDERID_System                 FOLDERID = "1ac14e77-02e7-4e5d-b744-2eb1ae5198b7"
	FOLDERID_SystemX86              FOLDERID = "d65231b0-b2f1-4857-a4ce-a8e7c6ea7d27"
	FOLDERID_Windows                FOLDERID = "f38bf404-1d43-42f2-9305-67de0b28fc23"
	FOLDERID_Profile                FOLDERID = "5e6c858f-0e22-4760-9afe-ea3317b67173"
	FOLDERID_Pictures               FOLDERID = "33e28130-4e1e-4676-835a-98395c3bc3bb"
	FOLDERID_ProgramFilesX86        FOLDERID = "7c5a40ef-a0fb-4bfc-874a-c0f2e0b9fa8e"
	FOLDERID_ProgramFilesCommonX86  FOLDERID = "de974d24-d9c6-4d3e-bf91-f4455120b917"
	FOLDERID_ProgramFilesX64        FOLDERID = "6d809377-6af0-444b-8957-a3773f02200e"
	FOLDERID_ProgramFilesCommonX64  FOLDERID = "6365d5a7-0f0d-45e5-87f6-0da56b6a4f7d"
	FOLDERID_ProgramFiles           FOLDERID = "905e63b6-c1bf-494e-b29c-65b732d3d21a"
	FOLDERID_ProgramFilesCommon     FOLDERID = "f7f1ed05-9f6d-47a2-aaae-29d317c6f066"
	FOLDERID_UserProgramFiles       FOLDERID = "5cd7aee2-2219-4a67-b85d-6c9ce15660cb"
	FOLDERID_UserProgramFilesCommon FOLDERID = "bcbd3057-ca5c-4622-b42d-bc56db0ae516"
	FOLDERID_AdminTools             FOLDERID = "724ef170-a42d-4fef-9f26-b60e846fba4f"
	FOLDERID_CommonAdminTools       FOLDERID = "d0384e7d-bac3-4797-8f14-cba229b392b5"
	FOLDERID_Music                  FOLDERID = "4bd8d571-6d19-48d3-be97-422220080e43"
	FOLDERID_Videos                 FOLDERID = "18989b1d-99b5-455b-841c-ab7c74e4ddfc"
	FOLDERID_Ringtones              FOLDERID = "c870044b-f49e-4126-a9c3-b52a1ff411e8"
	FOLDERID_PublicPictures         FOLDERID = "b6ebfb86-6907-413c-9af7-4fc2abf07cc5"
	FOLDERID_PublicMusic            FOLDERID = "3214fab5-9757-4298-bb61-92a9deaa44ff"
	FOLDERID_PublicVideos           FOLDERID = "2400183a-6185-49fb-a2d8-4a392a602ba3"
	FOLDERID_PublicRingtones        FOLDERID = "e555ab60-153b-4d17-9f04-a5fe99fc15ec"
	FOLDERID_ResourceDir            FOLDERID = "8ad10c31-2adb-4296-a8f7-e4701232c972"
	FOLDERID_LocalizedResourcesDir  FOLDERID = "2a00375e-224c-49de-b8d1-440df7ef3ddc"
	FOLDERID_CommonOEMLinks         FOLDERID = "c1bae2d0-10df-4334-bedd-7aa20b227a9d"
	FOLDERID_CDBurning              FOLDERID = "9e52ab10-f80d-49df-acb8-4330f5687855"
	FOLDERID_UserProfiles           FOLDERID = "0762d272-c50a-4bb0-a382-697dcd729b80"
	FOLDERID_Playlists              FOLDERID = "de92c1c7-837f-4f69-a3bb-86e631204a23"
	FOLDERID_SamplePlaylists        FOLDERID = "15ca69b3-30ee-49c1-ace1-6b5ec372afb5"
	FOLDERID_SampleMusic            FOLDERID = "b250c668-f57d-4ee1-a63c-290ee7d1aa1f"
	FOLDERID_SamplePictures         FOLDERID = "c4900540-2379-4c75-844b-64e6faf8716b"
	FOLDERID_SampleVideos           FOLDERID = "859ead94-2e85-48ad-a71a-0969cb56a6cd"
	FOLDERID_PhotoAlbums            FOLDERID = "69d2cf90-fc33-4fb7-9a0c-ebb0f0fcb43c"
	FOLDERID_Public                 FOLDERID = "dfdf76a2-c82a-4d63-906a-5644ac457385"
	FOLDERID_ChangeRemovePrograms   FOLDERID = "df7266ac-9274-4867-8d55-3bd661de872d"
	FOLDERID_AppUpdates             FOLDERID = "a305ce99-f527-492b-8b1a-7e76fa98d6e4"
	FOLDERID_AddNewPrograms         FOLDERID = "de61d971-5ebc-4f02-a3a9-6c82895e5c04"
	FOLDERID_Downloads              FOLDERID = "374de290-123f-4565-9164-39c4925e467b"
	FOLDERID_PublicDownloads        FOLDERID = "3d644c9b-1fb8-4f30-9b45-f670235f79c0"
	FOLDERID_SavedSearches          FOLDERID = "7d1d3a04-debb-4115-95cf-2f29da2920da"
	FOLDERID_QuickLaunch            FOLDERID = "52a4f021-7b75-48a9-9f6b-4b87a210bc8f"
	FOLDERID_Contacts               FOLDERID = "56784854-c6cb-462b-8169-88e350acb882"
	FOLDERID_SidebarParts           FOLDERID = "a75d362e-50fc-4fb7-ac2c-a8beaa314493"
	FOLDERID_SidebarDefaultParts    FOLDERID = "7b396e54-9ec5-4300-be0a-2482ebae1a26"
	FOLDERID_PublicGameTasks        FOLDERID = "debf2536-e1a8-4c59-b6a2-414586476aea"
	FOLDERID_GameTasks              FOLDERID = "054fae61-4dd8-4787-80b6-090220c4b700"
	FOLDERID_SavedGames             FOLDERID = "4c5c32ff-bb9d-43b0-b5b4-2d72e54eaaa4"
	FOLDERID_Games                  FOLDERID = "cac52c1a-b53d-4edc-92d7-6b2e8ac19434"
	FOLDERID_SEARCH_MAPI            FOLDERID = "98ec0e18-2098-4d44-8644-66979315a281"
	FOLDERID_SEARCH_CSC             FOLDERID = "ee32e446-31ca-4aba-814f-a5ebd2fd6d5e"
	FOLDERID_Links                  FOLDERID = "bfb9d5e0-c6a9-404c-b2b2-ae6db6af4968"
	FOLDERID_UsersFiles             FOLDERID = "f3ce0f7c-4901-4acc-8648-d5d44b04ef8f"
	FOLDERID_UsersLibraries         FOLDERID = "a302545d-deff-464b-abe8-61c8648d939b"
	FOLDERID_SearchHome             FOLDERID = "190337d1-b8ca-4121-a639-6d472d16972a"
	FOLDERID_OriginalImages         FOLDERID = "2c36c0aa-5812-4b87-bfd0-4cd0dfb19b39"
	FOLDERID_DocumentsLibrary       FOLDERID = "7b0db17d-9cd2-4a93-9733-46cc89022e7c"
	FOLDERID_MusicLibrary           FOLDERID = "2112ab0a-c86a-4ffe-a368-0de96e47012e"
	FOLDERID_PicturesLibrary        FOLDERID = "a990ae9f-a03b-4e80-94bc-9912d7504104"
	FOLDERID_VideosLibrary          FOLDERID = "491e922f-5643-4af4-a7eb-4e7a138d8174"
	FOLDERID_RecordedTVLibrary      FOLDERID = "1a6fdba2-f42d-4358-a798-b74d745926c5"
	FOLDERID_HomeGroup              FOLDERID = "52528a6b-b9e3-4add-b60d-588c2dba842d"
	FOLDERID_HomeGroupCurrentUser   FOLDERID = "9b74b6a3-0dfd-4f11-9e78-5f7800f2e772"
	FOLDERID_DeviceMetadataStore    FOLDERID = "5ce4a5e9-e4eb-479d-b89f-130c02886155"
	FOLDERID_Libraries              FOLDERID = "1b3ea5dc-b587-4786-b4ef-bd1dc332aeae"
	FOLDERID_PublicLibraries        FOLDERID = "48daf80b-e6cf-4f4e-b800-0e69d84ee384"
	FOLDERID_UserPinned             FOLDERID = "9e3995ab-1f9c-4f13-b827-48b24b6c7174"
	FOLDERID_ImplicitAppShortcuts   FOLDERID = "bcb5256f-79f6-4cee-b725-dc34e402fd46"
	FOLDERID_AccountPictures        FOLDERID = "008ca0b1-55b4-4c56-b8a8-4de4b299d3be"
	FOLDERID_PublicUserTiles        FOLDERID = "0482af6c-08f1-4c34-8c90-e17ec98b1e17"
	FOLDERID_AppsFolder             FOLDERID = "1e87508d-89c2-42f0-8a7e-645a0f50ca58"
	FOLDERID_StartMenuAllPrograms   FOLDERID = "f26305ef-6948-40b9-b255-81453d09c785"
	FOLDERID_CommonStartMenuPlaces  FOLDERID = "a440879f-87a0-4f7d-b700-0207b966194a"
	FOLDERID_ApplicationShortcuts   FOLDERID = "a3918781-e5f2-4890-b3d9-a7e54332328c"
	FOLDERID_RoamingTiles           FOLDERID = "00bcfc5a-ed94-4e48-96a1-3f6217f21990"
	FOLDERID_RoamedTileImages       FOLDERID = "aaa8d5a5-f1d6-4259-baa8-78e7ef60835e"
	FOLDERID_Screenshots            FOLDERID = "b7bede81-df94-4682-a7d8-57a52620b86f"
	FOLDERID_CameraRoll             FOLDERID = "ab5fb87b-7ce2-4f83-915d-550846c9537b"
	FOLDERID_SkyDrive               FOLDERID = "a52bba46-e9e1-435f-b3d9-28daa648c0f6"
	FOLDERID_OneDrive               FOLDERID = "a52bba46-e9e1-435f-b3d9-28daa648c0f6"
	FOLDERID_SkyDriveDocuments      FOLDERID = "24d89e24-2f19-4534-9dde-6a6671fbb8fe"
	FOLDERID_SkyDrivePictures       FOLDERID = "339719b5-8c47-4894-94c2-d8f77add44a6"
	FOLDERID_SkyDriveMusic          FOLDERID = "c3f2459e-80d6-45dc-bfef-1f769f2be730"
	FOLDERID_SkyDriveCameraRoll     FOLDERID = "767e6811-49cb-4273-87c2-20f355e1085b"
	FOLDERID_SearchHistory          FOLDERID = "0d4c3db6-03a3-462f-a0e6-08924c41b5d4"
	FOLDERID_SearchTemplates        FOLDERID = "7e636bfe-dfa9-4d5e-b456-d7b39851d8a9"
	FOLDERID_CameraRollLibrary      FOLDERID = "2b20df75-1eda-4039-8097-38798227d5b7"
	FOLDERID_SavedPictures          FOLDERID = "3b193882-d3ad-4eab-965a-69829d1fb59f"
	FOLDERID_SavedPicturesLibrary   FOLDERID = "e25b5812-be88-4bd9-94b0-29233477b6c3"
	FOLDERID_RetailDemo             FOLDERID = "12d4c69e-24ad-4923-be19-31321c43a767"
	FOLDERID_Device                 FOLDERID = "1c2ac1dc-4358-4b6c-9733-af21156576f0"
	FOLDERID_DevelopmentFiles       FOLDERID = "dbe8e08e-3053-4bbc-b183-2a7b2b191e59"
	FOLDERID_Objects3D              FOLDERID = "31c0dd25-9439-4f12-bf41-7ff4eda38722"
	FOLDERID_AppCaptures            FOLDERID = "edc0fe71-98d8-4f4a-b920-c8dc133cb165"
	FOLDERID_LocalDocuments         FOLDERID = "f42ee2d3-909f-4907-8871-4c22fc0bf756"
	FOLDERID_LocalPictures          FOLDERID = "0ddd015d-b06c-45d5-8c4c-f59713854639"
	FOLDERID_LocalVideos            FOLDERID = "35286a68-3c57-41a1-bbb1-0eae73d76c95"
	FOLDERID_LocalMusic             FOLDERID = "a0c69a99-21c8-4671-8703-7934162fcf1d"
	FOLDERID_LocalDownloads         FOLDERID = "7d83ee9b-2244-4e70-b1f5-5393042af1e4"
	FOLDERID_RecordedCalls          FOLDERID = "2f8b40c2-83ed-48ee-b383-a1f157ec6f9a"
	FOLDERID_AllAppMods             FOLDERID = "7ad67899-66af-43ba-9156-6aad42e6c596"
	FOLDERID_CurrentAppMods         FOLDERID = "3db40b20-2a30-4dbe-917e-771dd21dd099"
	FOLDERID_AppDataDesktop         FOLDERID = "b2c5e279-7add-439f-b28c-c41fe1bbf672"
	FOLDERID_AppDataDocuments       FOLDERID = "7be16610-1f7f-44ac-bff0-83e15f2ffca1"
	FOLDERID_AppDataFavorites       FOLDERID = "7cfbefbc-de1f-45aa-b843-a542ac536cc9"
	FOLDERID_AppDataProgramData     FOLDERID = "559d40a3-a036-40fa-af61-84cb430a4d34"
	FOLDERID_LocalStorage           FOLDERID = "b3eb08d3-a1f3-496b-865a-42b536cda0ec"
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

// [IShellFolder.CompareIDs] lParam.
//
// [IShellFolder.CompareIDs]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellfolder-compareids
type SHCIDS uint32

const (
	SHCIDS_ALLFIELDS     SHCIDS = 0x8000_0000
	SHCIDS_CANONICALONLY SHCIDS = 0x1000_0000
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

// [IShellLink.Resolve] flags.
//
// [IShellLink.Resolve]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-resolve
type SLR uint32

const (
	SLR_NONE                      SLR = 0
	SLR_NO_UI                     SLR = 0x1
	SLR_ANY_MATCH                 SLR = 0x2
	SLR_UPDATE                    SLR = 0x4
	SLR_NOUPDATE                  SLR = 0x8
	SLR_NOSEARCH                  SLR = 0x10
	SLR_NOTRACK                   SLR = 0x20
	SLR_NOLINKINFO                SLR = 0x40
	SLR_INVOKE_MSI                SLR = 0x80
	SLR_NO_UI_WITH_MSG_PUMP       SLR = 0x101
	SLR_OFFER_DELETE_WITHOUT_FILE SLR = 0x200
	SLR_KNOWNFOLDER               SLR = 0x400
	SLR_MACHINE_IN_LOCAL_TARGET   SLR = 0x800
	SLR_UPDATE_MACHINE_AND_SID    SLR = 0x1000
	SLR_NO_OBJECT_ID              SLR = 0x2000
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
