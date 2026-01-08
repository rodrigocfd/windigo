//go:build windows

package co

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
type FOLDERID GUID

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

// [PROPERTYKEY] struct predefined values, represented as a string.
//
// [PROPERTYKEY]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ns-wtypes-propertykey
type PKEY GUID

const (
	// Address properties

	PKEY_Address_Country     PKEY = "c07b4199-e1df-4493-b1e1-de5946fb58f8 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Address_CountryCode PKEY = "c07b4199-e1df-4493-b1e1-de5946fb58f8 101" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Address_Region      PKEY = "c07b4199-e1df-4493-b1e1-de5946fb58f8 102" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Address_RegionCode  PKEY = "c07b4199-e1df-4493-b1e1-de5946fb58f8 103" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Address_Town        PKEY = "c07b4199-e1df-4493-b1e1-de5946fb58f8 104" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Audio properties

	PKEY_Audio_ChannelCount      PKEY = "64440490-4c8b-11d1-8b70-080036b11a03 7"   // UInt32 -- VT_UI4
	PKEY_Audio_Compression       PKEY = "64440490-4c8b-11d1-8b70-080036b11a03 10"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Audio_EncodingBitrate   PKEY = "64440490-4c8b-11d1-8b70-080036b11a03 4"   // UInt32 -- VT_UI4
	PKEY_Audio_Format            PKEY = "64440490-4c8b-11d1-8b70-080036b11a03 2"   // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_BSTR.
	PKEY_Audio_IsVariableBitRate PKEY = "e6822fee-8c17-4d62-823c-8e9cfcbd1d5c 100" // Boolean -- VT_BOOL
	PKEY_Audio_PeakValue         PKEY = "2579e5d0-1116-4084-bd9a-9b4f7cb4df5e 100" // UInt32 -- VT_UI4
	PKEY_Audio_SampleRate        PKEY = "64440490-4c8b-11d1-8b70-080036b11a03 5"   // UInt32 -- VT_UI4
	PKEY_Audio_SampleSize        PKEY = "64440490-4c8b-11d1-8b70-080036b11a03 6"   // UInt32 -- VT_UI4
	PKEY_Audio_StreamName        PKEY = "64440490-4c8b-11d1-8b70-080036b11a03 9"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Audio_StreamNumber      PKEY = "64440490-4c8b-11d1-8b70-080036b11a03 8"   // UInt16 -- VT_UI2

	// Calendar properties

	PKEY_Calendar_Duration                  PKEY = "293ca35a-09aa-4dd2-b180-1fe245728a52 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Calendar_IsOnline                  PKEY = "bfee9149-e3e2-49a7-a862-c05988145cec 100" // Boolean -- VT_BOOL
	PKEY_Calendar_IsRecurring               PKEY = "315b9c8d-80a9-4ef9-ae16-8e746da51d70 100" // Boolean -- VT_BOOL
	PKEY_Calendar_Location                  PKEY = "f6272d18-cecc-40b1-b26a-3911717aa7bd 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Calendar_OptionalAttendeeAddresses PKEY = "d55bae5a-3892-417a-a649-c6ac5aaaeab3 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_OptionalAttendeeNames     PKEY = "09429607-582d-437f-84c3-de93a2b24c3c 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_OrganizerAddress          PKEY = "744c8242-4df5-456c-ab9e-014efb9021e3 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Calendar_OrganizerName             PKEY = "aaa660f9-9865-458e-b484-01bc7fe3973e 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Calendar_ReminderTime              PKEY = "72fc5ba4-24f9-4011-9f3f-add27afad818 100" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Calendar_RequiredAttendeeAddresses PKEY = "0ba7d6c3-568d-4159-ab91-781a91fb71e5 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_RequiredAttendeeNames     PKEY = "b33af30b-f552-4584-936c-cb93e5cda29f 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_Resources                 PKEY = "00f58a38-c54b-4c40-8696-97235980eae1 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Calendar_ResponseStatus            PKEY = "188c1f91-3c40-4132-9ec5-d8b03b72a8a2 100" // UInt16 -- VT_UI2
	PKEY_Calendar_ShowTimeAs                PKEY = "5bf396d4-5eb2-466f-bde9-2fb3f2361d6e 100" // UInt16 -- VT_UI2
	PKEY_Calendar_ShowTimeAsText            PKEY = "53da57cf-62c0-45c4-81de-7610bcefd7f5 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Communication properties

	PKEY_Communication_AccountName       PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 9"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Communication_DateItemExpires   PKEY = "428040ac-a177-4c8a-9760-f6f761227f9a 100" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Communication_Direction         PKEY = "8e531030-b960-4346-ae0d-66bc9a86fb94 100" // UInt16 -- VT_UI2
	PKEY_Communication_FollowupIconIndex PKEY = "83a6347e-6fe4-4f40-ba9c-c4865240d1f4 100" // Int32 -- VT_I4
	PKEY_Communication_HeaderItem        PKEY = "c9c34f84-2241-4401-b607-bd20ed75ae7f 100" // Boolean -- VT_BOOL
	PKEY_Communication_PolicyTag         PKEY = "ec0b4191-ab0b-4c66-90b6-c6637cdebbab 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Communication_SecurityFlags     PKEY = "8619a4b6-9f4d-4429-8c0f-b996ca59e335 100" // Int32 -- VT_I4
	PKEY_Communication_Suffix            PKEY = "807b653a-9e91-43ef-8f97-11ce04ee20c5 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Communication_TaskStatus        PKEY = "be1a72c6-9a1d-46b7-afe7-afaf8cef4999 100" // UInt16 -- VT_UI2
	PKEY_Communication_TaskStatusText    PKEY = "a6744477-c237-475b-a075-54f34498292a 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Computer properties

	PKEY_Computer_DecoratedFreeSpace PKEY = "9b174b35-40ff-11d2-a27e-00c04fc30871 7" // Multivalue UInt64 -- VT_VECTOR | VT_UI8  (For variants: VT_ARRAY | VT_UI8)

	// Contact properties

	PKEY_Contact_AccountPictureDynamicVideo       PKEY = "0b8bb018-2725-4b44-92ba-7933aeb2dde7 2"   // Stream -- VT_STREAM
	PKEY_Contact_AccountPictureLarge              PKEY = "0b8bb018-2725-4b44-92ba-7933aeb2dde7 3"   // Stream -- VT_STREAM
	PKEY_Contact_AccountPictureSmall              PKEY = "0b8bb018-2725-4b44-92ba-7933aeb2dde7 4"   // Stream -- VT_STREAM
	PKEY_Contact_Anniversary                      PKEY = "9ad5badb-cea7-4470-a03d-b84e51b9949e 100" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Contact_AssistantName                    PKEY = "cd102c9c-5540-4a88-a6f6-64e4981c8cd1 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_AssistantTelephone               PKEY = "9a93244d-a7ad-4ff8-9b99-45ee4cc09af6 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Birthday                         PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 47"  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Contact_BusinessAddress                  PKEY = "730fb6dd-cf7c-426b-a03f-bd166cc9ee24 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1Country          PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 119" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1Locality         PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 117" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1PostalCode       PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 120" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1Region           PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 118" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress1Street           PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 116" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2Country          PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 124" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2Locality         PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 122" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2PostalCode       PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 125" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2Region           PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 123" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress2Street           PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 121" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3Country          PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 129" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3Locality         PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 127" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3PostalCode       PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 130" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3Region           PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 128" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddress3Street           PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 126" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressCity              PKEY = "402b5934-ec5a-48c3-93e6-85e86a2d934e 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressCountry           PKEY = "b0b87314-fcf6-4feb-8dff-a50da6af561c 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressPostalCode        PKEY = "e1d4a09e-d758-4cd1-b6ec-34a8b5a73f80 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressPostOfficeBox     PKEY = "bc4e71ce-17f9-48d5-bee9-021df0ea5409 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressState             PKEY = "446f787f-10c4-41cb-a6c4-4d0343551597 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessAddressStreet            PKEY = "ddd1460f-c0bf-4553-8ce4-10433c908fb0 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessEmailAddresses           PKEY = "f271c659-7e5e-471f-ba25-7f77b286f836 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_BusinessFaxNumber                PKEY = "91eff6f3-2e27-42ca-933e-7c999fbe310b 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessHomePage                 PKEY = "56310920-2491-4919-99ce-eadb06fafdb2 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_BusinessTelephone                PKEY = "6a15e5a0-0a1e-4cd7-bb8c-d2f1b0c929bc 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_CallbackTelephone                PKEY = "bf53d1c3-49e0-4f7f-8567-5a821d8ac542 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_CarTelephone                     PKEY = "8fdc6dea-b929-412b-ba90-397a257465fe 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Children                         PKEY = "d4729704-8ef1-43ef-9024-2bd381187fd5 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_CompanyMainTelephone             PKEY = "8589e481-6040-473d-b171-7fa89c2708ed 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_ConnectedServiceDisplayName      PKEY = "39b77f4f-a104-4863-b395-2db2ad8f7bc1 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_ConnectedServiceIdentities       PKEY = "80f41eb8-afc4-4208-aa5f-cce21a627281 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_ConnectedServiceName             PKEY = "b5c84c9e-5927-46b5-a3cc-933c21b78469 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_ConnectedServiceSupportedActions PKEY = "a19fb7a9-024b-4371-a8bf-4d29c3e4e9c9 100" // UInt32 -- VT_UI4
	PKEY_Contact_DataSuppliers                    PKEY = "9660c283-fc3a-4a08-a096-eed3aac46da2 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_Department                       PKEY = "fc9f7306-ff8f-4d49-9fb6-3ffe5c0951ec 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_DisplayBusinessPhoneNumbers      PKEY = "364028da-d895-41fe-a584-302b1bb70a76 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_DisplayHomePhoneNumbers          PKEY = "5068bcdf-d697-4d85-8c53-1f1cdab01763 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_DisplayMobilePhoneNumbers        PKEY = "9cb0c358-9d7a-46b1-b466-dcc6f1a3d93d 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_DisplayOtherPhoneNumbers         PKEY = "03089873-8ee8-4191-bd60-d31f72b7900b 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_EmailAddress                     PKEY = "f8fa7fa3-d12b-4785-8a4e-691a94f7a3e7 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_EmailAddress2                    PKEY = "38965063-edc8-4268-8491-b7723172cf29 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_EmailAddress3                    PKEY = "644d37b4-e1b3-4bad-b099-7e7c04966aca 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_EmailAddresses                   PKEY = "84d8f337-981d-44b3-9615-c7596dba17e3 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_EmailName                        PKEY = "cc6f4f24-6083-4bd4-8754-674d0de87ab8 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_FileAsName                       PKEY = "f1a24aa7-9ca7-40f6-89ec-97def9ffe8db 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_FirstName                        PKEY = "14977844-6b49-4aad-a714-a4513bf60460 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_FullName                         PKEY = "635e9051-50a5-4ba2-b9db-4ed056c77296 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Gender                           PKEY = "3c8cee58-d4f0-4cf9-b756-4e5d24447bcd 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_GenderValue                      PKEY = "3c8cee58-d4f0-4cf9-b756-4e5d24447bcd 101" // UInt16 -- VT_UI2
	PKEY_Contact_Hobbies                          PKEY = "5dc2253f-5e11-4adf-9cfe-910dd01e3e70 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_HomeAddress                      PKEY = "98f98354-617a-46b8-8560-5b1b64bf1f89 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1Country              PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 104" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1Locality             PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 102" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1PostalCode           PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 105" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1Region               PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 103" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress1Street               PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 101" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2Country              PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 109" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2Locality             PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 107" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2PostalCode           PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 110" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2Region               PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 108" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress2Street               PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 106" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3Country              PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 114" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3Locality             PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 112" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3PostalCode           PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 115" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3Region               PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 113" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddress3Street               PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 111" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressCity                  PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 65"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressCountry               PKEY = "08a65aa1-f4c9-43dd-9ddf-a33d8e7ead85 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressPostalCode            PKEY = "8afcc170-8a46-4b53-9eee-90bae7151e62 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressPostOfficeBox         PKEY = "7b9f6399-0a3f-4b12-89bd-4adc51c918af 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressState                 PKEY = "c89a23d0-7d6d-4eb8-87d4-776a82d493e5 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeAddressStreet                PKEY = "0adef160-db3f-4308-9a21-06237b16fa2a 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeEmailAddresses               PKEY = "56c90e9d-9d46-4963-886f-2e1cd9a694ef 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_HomeFaxNumber                    PKEY = "660e04d6-81ab-4977-a09f-82313113ab26 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_HomeTelephone                    PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 20"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_IMAddress                        PKEY = "d68dbd8a-3374-4b81-9972-3ec30682db3d 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_Initials                         PKEY = "f3d8f40d-50cb-44a2-9718-40cb9119495d 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JA_CompanyNamePhonetic           PKEY = "897b3694-fe9e-43e6-8066-260f590c0100 2"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JA_FirstNamePhonetic             PKEY = "897b3694-fe9e-43e6-8066-260f590c0100 3"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JA_LastNamePhonetic              PKEY = "897b3694-fe9e-43e6-8066-260f590c0100 4"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1CompanyAddress           PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 120" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1CompanyName              PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 102" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1Department               PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 106" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1Manager                  PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 105" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1OfficeLocation           PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 104" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1Title                    PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 103" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo1YomiCompanyName          PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 101" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2CompanyAddress           PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 121" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2CompanyName              PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 108" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2Department               PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 113" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2Manager                  PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 112" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2OfficeLocation           PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 110" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2Title                    PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 109" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo2YomiCompanyName          PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 107" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3CompanyAddress           PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 123" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3CompanyName              PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 115" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3Department               PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 119" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3Manager                  PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 118" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3OfficeLocation           PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 117" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3Title                    PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 116" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobInfo3YomiCompanyName          PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 114" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_JobTitle                         PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 6"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Label                            PKEY = "97b0ad89-df49-49cc-834e-660974fd755b 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_LastName                         PKEY = "8f367200-c270-457c-b1d4-e07c5bcd90c7 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_MailingAddress                   PKEY = "c0ac206a-827e-4650-95ae-77e2bb74fcc9 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_MiddleName                       PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 71"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_MobileTelephone                  PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 35"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_NickName                         PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 74"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OfficeLocation                   PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 7"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress                     PKEY = "508161fa-313b-43d5-83a1-c1accf68622c 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1Country             PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 134" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1Locality            PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 132" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1PostalCode          PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 135" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1Region              PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 133" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress1Street              PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 131" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2Country             PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 139" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2Locality            PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 137" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2PostalCode          PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 140" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2Region              PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 138" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress2Street              PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 136" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3Country             PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 144" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3Locality            PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 142" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3PostalCode          PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 145" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3Region              PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 143" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddress3Street              PKEY = "a7b6f596-d678-4bc1-b05f-0203d27e8aa1 141" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressCity                 PKEY = "6e682923-7f7b-4f0c-a337-cfca296687bf 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressCountry              PKEY = "8f167568-0aae-4322-8ed9-6055b7b0e398 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressPostalCode           PKEY = "95c656c1-2abf-4148-9ed3-9ec602e3b7cd 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressPostOfficeBox        PKEY = "8b26ea41-058f-43f6-aecc-4035681ce977 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressState                PKEY = "71b377d6-e570-425f-a170-809fae73e54e 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherAddressStreet               PKEY = "ff962609-b7d6-4999-862d-95180d529aea 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_OtherEmailAddresses              PKEY = "11d6336b-38c4-4ec9-84d6-eb38d0b150af 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_PagerTelephone                   PKEY = "d6304e01-f8f5-4f45-8b15-d024a6296789 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PersonalTitle                    PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 69"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PhoneNumbersCanonical            PKEY = "d042d2a1-927e-40b5-a503-6edbd42a517e 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Contact_Prefix                           PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 75"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressCity               PKEY = "c8ea94f0-a9e3-4969-a94b-9c62a95324e0 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressCountry            PKEY = "e53d799d-0f3f-466e-b2ff-74634a3cb7a4 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressPostalCode         PKEY = "18bbd425-ecfd-46ef-b612-7b4a6034eda0 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressPostOfficeBox      PKEY = "de5ef3c7-46e1-484e-9999-62c5308394c1 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressState              PKEY = "f1176dfe-7138-4640-8b4c-ae375dc70a6d 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryAddressStreet             PKEY = "63c25b20-96be-488f-8788-c09c407ad812 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryEmailAddress              PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 48"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_PrimaryTelephone                 PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 25"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Profession                       PKEY = "7268af55-1ce4-4f6e-a41f-b6e4ef10e4a9 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_SpouseName                       PKEY = "9d2408b6-3167-422b-82b0-f583b7a7cfe3 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Suffix                           PKEY = "176dc63c-2688-4e89-8143-a347800f25e9 73"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_TelexNumber                      PKEY = "c554493c-c1f7-40c1-a76c-ef8c0614003e 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_TTYTDDTelephone                  PKEY = "aaf16bac-2b55-45e6-9f6d-415eb94910df 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_WebPage                          PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 18"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Webpage2                         PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 124" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Contact_Webpage3                         PKEY = "00f63dd8-22bd-4a5d-ba34-5cb0b9bdcb03 125" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Core properties

	PKEY_AcquisitionID                                      PKEY = "65a98875-3c80-40ab-abbc-efdaf77dbee2 100"   // Int32 -- VT_I4
	PKEY_ApplicationDefinedProperties                       PKEY = "cdbfc167-337e-41d8-af7c-8c09205429c7 100"   // Any -- VT_NULL  Legacy code may treat this as VT_UNKNOWN.
	PKEY_ApplicationName                                    PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 18"    // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_AppZoneIdentifier                                  PKEY = "502cfeab-47eb-459c-b960-e6d8728f7701 102"   // UInt32 -- VT_UI4
	PKEY_Author                                             PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 4"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_CachedFileUpdaterContentIdForConflictResolution    PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 114"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_CachedFileUpdaterContentIdForStream                PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 113"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Capacity                                           PKEY = "9b174b35-40ff-11d2-a27e-00c04fc30871 3"     // UInt64 -- VT_UI8
	PKEY_Category                                           PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 2"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Comment                                            PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 6"     // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_Company                                            PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 15"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ComputerName                                       PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 5"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ContainedItems                                     PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 29"    // Multivalue Guid -- VT_VECTOR | VT_CLSID  (For variants: VT_ARRAY | VT_CLSID)
	PKEY_ContentId                                          PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 132"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ContentStatus                                      PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 27"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ContentType                                        PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 26"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ContentUri                                         PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 131"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Copyright                                          PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 11"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_CreatorAppId                                       PKEY = "c2ea046e-033c-4e91-bd5b-d4942f6bbe49 2"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_CreatorOpenWithUIOptions                           PKEY = "c2ea046e-033c-4e91-bd5b-d4942f6bbe49 3"     // UInt32 -- VT_UI4
	PKEY_DataObjectFormat                                   PKEY = "1e81a3f8-a30f-4247-b9ee-1d0368a9425c 2"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_DateAccessed                                       PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 16"    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateAcquired                                       PKEY = "2cbaa8f5-d81f-47ca-b17a-f8d822300131 100"   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateArchived                                       PKEY = "43f8d7b7-a444-4f87-9383-52271c9b915c 100"   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateCompleted                                      PKEY = "72fab781-acda-43e5-b155-b2434f85e678 100"   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateCreated                                        PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 15"    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateImported                                       PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 18258" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DateModified                                       PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 14"    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DefaultSaveLocationDisplay                         PKEY = "5d76b67f-9b3d-44bb-b6ae-25da4f638a67 10"    // UInt32 -- VT_UI4
	PKEY_DueDate                                            PKEY = "3f8472b5-e0af-4db2-8071-c53fe76ae7ce 100"   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_EndDate                                            PKEY = "c75faa05-96fd-49e7-9cb4-9f601082d553 100"   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_ExpandoProperties                                  PKEY = "6fa20de6-d11c-4d9d-a154-64317628c12d 100"   // Any -- VT_NULL  Legacy code may treat this as VT_UNKNOWN.
	PKEY_FileAllocationSize                                 PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 18"    // UInt64 -- VT_UI8
	PKEY_FileAttributes                                     PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 13"    // UInt32 -- VT_UI4
	PKEY_FileCount                                          PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 12"    // UInt64 -- VT_UI8
	PKEY_FileDescription                                    PKEY = "0cef7d53-fa64-11d1-a203-0000f81fedee 3"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FileExtension                                      PKEY = "e4f10a3c-49e6-405d-8288-a23bd4eeaa6c 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FileFRN                                            PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 21"    // UInt64 -- VT_UI8
	PKEY_FileName                                           PKEY = "41cf5ae0-f75a-4806-bd87-59c7d9248eb9 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FileOfflineAvailabilityStatus                      PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 100"   // UInt32 -- VT_UI4
	PKEY_FileOwner                                          PKEY = "9b174b34-40ff-11d2-a27e-00c04fc30871 4"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FilePlaceholderStatus                              PKEY = "b2f9b9d6-fec4-4dd5-94d7-8957488c807b 2"     // UInt32 -- VT_UI4
	PKEY_FileVersion                                        PKEY = "0cef7d53-fa64-11d1-a203-0000f81fedee 4"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FindData                                           PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 0"     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_FlagColor                                          PKEY = "67df94de-0ca7-4d6f-b792-053a3e4f03cf 100"   // UInt16 -- VT_UI2
	PKEY_FlagColorText                                      PKEY = "45eae747-8e2a-40ae-8cbf-ca52aba6152a 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FlagStatus                                         PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 12"    // Int32 -- VT_I4
	PKEY_FlagStatusText                                     PKEY = "dc54fd2e-189d-4871-aa01-08c2f57a4abc 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FolderKind                                         PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 101"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_FolderNameDisplay                                  PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 25"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_FreeSpace                                          PKEY = "9b174b35-40ff-11d2-a27e-00c04fc30871 2"     // UInt64 -- VT_UI8
	PKEY_FullText                                           PKEY = "1e3ee840-bc2b-476c-8237-2acd1a839b22 6"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_HighKeywords                                       PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 24"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Identity                                           PKEY = "a26f4afc-7346-4299-be47-eb1ae613139f 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_Blob                                      PKEY = "8c3b93a4-baed-1a83-9a32-102ee313f6eb 100"   // Blob -- VT_BLOB
	PKEY_Identity_DisplayName                               PKEY = "7d683fc9-d155-45a8-bb1f-89d19bcb792f 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_InternetSid                               PKEY = "6d6d5d49-265d-4688-9f4e-1fdd33e7cc83 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_IsMeIdentity                              PKEY = "a4108708-09df-4377-9dfc-6d99986d5a67 100"   // Boolean -- VT_BOOL
	PKEY_Identity_KeyProviderContext                        PKEY = "a26f4afc-7346-4299-be47-eb1ae613139f 17"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_KeyProviderName                           PKEY = "a26f4afc-7346-4299-be47-eb1ae613139f 16"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_LogonStatusString                         PKEY = "f18dedf3-337f-42c0-9e03-cee08708a8c3 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_PrimaryEmailAddress                       PKEY = "fcc16823-baed-4f24-9b32-a0982117f7fa 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_PrimarySid                                PKEY = "2b1b801e-c0c1-4987-9ec5-72fa89814787 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_ProviderData                              PKEY = "a8a74b92-361b-4e9a-b722-7c4a7330a312 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_ProviderID                                PKEY = "74a7de49-fa11-4d3d-a006-db7e08675916 100"   // Guid -- VT_CLSID
	PKEY_Identity_QualifiedUserName                         PKEY = "da520e51-f4e9-4739-ac82-02e0a95c9030 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_UniqueID                                  PKEY = "e55fc3b0-2b60-4220-918e-b21e8bf16016 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Identity_UserName                                  PKEY = "c4322503-78ca-49c6-9acc-a68e2afd7b6b 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_IdentityProvider_Name                              PKEY = "b96eff7b-35ca-4a35-8607-29e3a54c46ea 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_IdentityProvider_Picture                           PKEY = "2425166f-5642-4864-992f-98fd98f294c3 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ImageParsingName                                   PKEY = "d7750ee0-c6a4-48ec-b53e-b87b52e6d073 100"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Importance                                         PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 11"    // Int32 -- VT_I4
	PKEY_ImportanceText                                     PKEY = "a3b29791-7713-4e1d-bb40-17db85f01831 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_IsAttachment                                       PKEY = "f23f425c-71a1-4fa8-922f-678ea4a60408 100"   // Boolean -- VT_BOOL
	PKEY_IsDefaultNonOwnerSaveLocation                      PKEY = "5d76b67f-9b3d-44bb-b6ae-25da4f638a67 5"     // Boolean -- VT_BOOL
	PKEY_IsDefaultSaveLocation                              PKEY = "5d76b67f-9b3d-44bb-b6ae-25da4f638a67 3"     // Boolean -- VT_BOOL
	PKEY_IsDeleted                                          PKEY = "5cda5fc8-33ee-4ff3-9094-ae7bd8868c4d 100"   // Boolean -- VT_BOOL
	PKEY_IsEncrypted                                        PKEY = "90e5e14e-648b-4826-b2aa-acaf790e3513 10"    // Boolean -- VT_BOOL
	PKEY_IsFlagged                                          PKEY = "5da84765-e3ff-4278-86b0-a27967fbdd03 100"   // Boolean -- VT_BOOL
	PKEY_IsFlaggedComplete                                  PKEY = "a6f360d2-55f9-48de-b909-620e090a647c 100"   // Boolean -- VT_BOOL
	PKEY_IsIncomplete                                       PKEY = "346c8bd1-2e6a-4c45-89a4-61b78e8e700f 100"   // Boolean -- VT_BOOL
	PKEY_IsLocationSupported                                PKEY = "5d76b67f-9b3d-44bb-b6ae-25da4f638a67 8"     // Boolean -- VT_BOOL
	PKEY_IsPinnedToNameSpaceTree                            PKEY = "5d76b67f-9b3d-44bb-b6ae-25da4f638a67 2"     // Boolean -- VT_BOOL
	PKEY_IsRead                                             PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 10"    // Boolean -- VT_BOOL
	PKEY_IsSearchOnlyItem                                   PKEY = "5d76b67f-9b3d-44bb-b6ae-25da4f638a67 4"     // Boolean -- VT_BOOL
	PKEY_IsSendToTarget                                     PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 33"    // Boolean -- VT_BOOL
	PKEY_IsShared                                           PKEY = "ef884c5b-2bfe-41bb-aae5-76eedf4f9902 100"   // Boolean -- VT_BOOL
	PKEY_ItemAuthors                                        PKEY = "d0a04f0a-462a-48a4-bb2f-3706e88dbd7d 100"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_ItemClassType                                      PKEY = "048658ad-2db8-41a4-bbb6-ac1ef1207eb1 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemDate                                           PKEY = "f7db74b4-4287-4103-afba-f1b13dcd75cf 100"   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_ItemFolderNameDisplay                              PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 2"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemFolderPathDisplay                              PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 6"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemFolderPathDisplayNarrow                        PKEY = "dabd30ed-0043-4789-a7f8-d013a4736622 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemName                                           PKEY = "6b8da074-3b5c-43bc-886f-0a2cdce00b6f 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemNameDisplay                                    PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 10"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemNameDisplayWithoutExtension                    PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 24"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemNamePrefix                                     PKEY = "d7313ff1-a77a-401c-8c99-3dbdd68add36 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemNameSortOverride                               PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 23"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemParticipants                                   PKEY = "d4d0aa16-9948-41a4-aa85-d97ff9646993 100"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_ItemPathDisplay                                    PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 7"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemPathDisplayNarrow                              PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 8"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemSubType                                        PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 37"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemType                                           PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 11"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemTypeText                                       PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 4"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ItemUrl                                            PKEY = "49691c90-7e17-101a-a91c-08002b2ecda9 9"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Keywords                                           PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 5"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_Kind                                               PKEY = "1e3ee840-bc2b-476c-8237-2acd1a839b22 3"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_KindText                                           PKEY = "f04bef95-c585-4197-a2b7-df46fdc9ee6d 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Language                                           PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 28"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_LastSyncError                                      PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 107"   // UInt32 -- VT_UI4
	PKEY_LastSyncWarning                                    PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 128"   // UInt32 -- VT_UI4
	PKEY_LastWriterPackageFamilyName                        PKEY = "502cfeab-47eb-459c-b960-e6d8728f7701 101"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_LowKeywords                                        PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 25"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_MediumKeywords                                     PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 26"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_MileageInformation                                 PKEY = "fdf84370-031a-4add-9e91-0d775f1c6605 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_MIMEType                                           PKEY = "0b63e350-9ccc-11d0-bcdb-00805fccce04 5"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Null                                               PKEY = "00000000-0000-0000-0000-000000000000 0"     // Null -- VT_NULL
	PKEY_OfflineAvailability                                PKEY = "a94688b6-7d9f-4570-a648-e3dfc0ab2b3f 100"   // UInt32 -- VT_UI4
	PKEY_OfflineStatus                                      PKEY = "6d24888f-4718-4bda-afed-ea0fb4386cd8 100"   // UInt32 -- VT_UI4
	PKEY_OriginalFileName                                   PKEY = "0cef7d53-fa64-11d1-a203-0000f81fedee 6"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_OwnerSID                                           PKEY = "5d76b67f-9b3d-44bb-b6ae-25da4f638a67 6"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParentalRating                                     PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 21"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParentalRatingReason                               PKEY = "10984e0a-f9f2-4321-b7ef-baf195af4319 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParentalRatingsOrganization                        PKEY = "a7fe0840-1344-46f0-8d37-52ed712a4bf9 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParsingBindContext                                 PKEY = "dfb9a04d-362f-4ca3-b30b-0254b17b5b84 100"   // Any -- VT_NULL  Legacy code may treat this as VT_UNKNOWN.
	PKEY_ParsingName                                        PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 24"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ParsingPath                                        PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 30"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PerceivedType                                      PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 9"     // Int32 -- VT_I4
	PKEY_PercentFull                                        PKEY = "9b174b35-40ff-11d2-a27e-00c04fc30871 5"     // UInt32 -- VT_UI4
	PKEY_Priority                                           PKEY = "9c1fcf74-2d97-41ba-b4ae-cb2e3661a6e4 5"     // UInt16 -- VT_UI2
	PKEY_PriorityText                                       PKEY = "d98be98b-b86b-4095-bf52-9d23b2e0a752 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Project                                            PKEY = "39a7f922-477c-48de-8bc8-b28441e342e3 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_ProviderItemID                                     PKEY = "f21d9941-81f0-471a-adee-4e74b49217ed 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Rating                                             PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 9"     // UInt32 -- VT_UI4
	PKEY_RatingText                                         PKEY = "90197ca7-fd8f-4e8c-9da3-b57e1e609295 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RemoteConflictingFile                              PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 115"   // Object -- VT_UNKNOWN
	PKEY_Security_AllowedEnterpriseDataProtectionIdentities PKEY = "38d43380-d418-4830-84d5-46935a81c5c6 32"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Security_EncryptionOwners                          PKEY = "5f5aff6a-37e5-4780-97ea-80c7565cf535 34"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Security_EncryptionOwnersDisplay                   PKEY = "de621b8f-e125-43a3-a32d-5665446d632a 25"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Sensitivity                                        PKEY = "f8d3f6ac-4874-42cb-be59-ab454b30716a 100"   // UInt16 -- VT_UI2
	PKEY_SensitivityText                                    PKEY = "d0c7f054-3f72-4725-8527-129a577cb269 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SFGAOFlags                                         PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 25"    // UInt32 -- VT_UI4
	PKEY_SharedWith                                         PKEY = "ef884c5b-2bfe-41bb-aae5-76eedf4f9902 200"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_ShareUserRating                                    PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 12"    // UInt32 -- VT_UI4
	PKEY_SharingStatus                                      PKEY = "ef884c5b-2bfe-41bb-aae5-76eedf4f9902 300"   // UInt32 -- VT_UI4
	PKEY_Shell_OmitFromView                                 PKEY = "de35258c-c695-4cbc-b982-38b0ad24ced0 2"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SimpleRating                                       PKEY = "a09f084e-ad41-489f-8076-aa5be3082bca 100"   // UInt32 -- VT_UI4
	PKEY_Size                                               PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 12"    // UInt64 -- VT_UI8
	PKEY_SoftwareUsed                                       PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 305"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SourceItem                                         PKEY = "668cdfa5-7a1b-4323-ae4b-e527393a1d81 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SourcePackageFamilyName                            PKEY = "ffae9db7-1c8d-43ff-818c-84403aa3732d 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StartDate                                          PKEY = "48fd6ec8-8a12-4cdf-a03e-4ec5a511edde 100"   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Status                                             PKEY = "000214a1-0000-0000-c000-000000000046 9"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderCallerVersionInformation            PKEY = "b2f9b9d6-fec4-4dd5-94d7-8957488c807b 7"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderError                               PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 109"   // UInt32 -- VT_UI4
	PKEY_StorageProviderFileChecksum                        PKEY = "b2f9b9d6-fec4-4dd5-94d7-8957488c807b 5"     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_StorageProviderFileFlags                           PKEY = "b2f9b9d6-fec4-4dd5-94d7-8957488c807b 8"     // UInt32 -- VT_UI4
	PKEY_StorageProviderFileHasConflict                     PKEY = "b2f9b9d6-fec4-4dd5-94d7-8957488c807b 9"     // Boolean -- VT_BOOL
	PKEY_StorageProviderFileIdentifier                      PKEY = "b2f9b9d6-fec4-4dd5-94d7-8957488c807b 3"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderFileRemoteUri                       PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 112"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderFileVersion                         PKEY = "b2f9b9d6-fec4-4dd5-94d7-8957488c807b 4"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderFileVersionWaterline                PKEY = "b2f9b9d6-fec4-4dd5-94d7-8957488c807b 6"     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_StorageProviderId                                  PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 108"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StorageProviderShareStatuses                       PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 111"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_StorageProviderSharingStatus                       PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 117"   // UInt32 -- VT_UI4
	PKEY_StorageProviderStatus                              PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 110"   // UInt64 -- VT_UI8
	PKEY_Subject                                            PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 3"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_SyncTransferStatus                                 PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 103"   // UInt32 -- VT_UI4
	PKEY_Thumbnail                                          PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 17"    // Clipboard -- VT_CF
	PKEY_ThumbnailCacheId                                   PKEY = "446d16b1-8dad-4870-a748-402ea43d788c 100"   // UInt64 -- VT_UI8
	PKEY_ThumbnailStream                                    PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 27"    // Stream -- VT_STREAM
	PKEY_Title                                              PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 2"     // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_TitleSortOverride                                  PKEY = "f0f7984d-222e-4ad2-82ab-1dd8ea40e57e 300"   // String -- VT_LPWSTR  (For variants: VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_TotalFileSize                                      PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 14"    // UInt64 -- VT_UI8
	PKEY_Trademarks                                         PKEY = "0cef7d53-fa64-11d1-a203-0000f81fedee 9"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_TransferOrder                                      PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 106"   // UInt64 -- VT_UI8
	PKEY_TransferPosition                                   PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 104"   // UInt64 -- VT_UI8
	PKEY_TransferSize                                       PKEY = "fceff153-e839-4cf3-a9e7-ea22832094b8 105"   // UInt64 -- VT_UI8
	PKEY_VolumeId                                           PKEY = "446d16b1-8dad-4870-a748-402ea43d788c 104"   // Guid -- VT_CLSID
	PKEY_ZoneIdentifier                                     PKEY = "502cfeab-47eb-459c-b960-e6d8728f7701 100"   // UInt32 -- VT_UI4

	// Devices properties

	PKEY_Device_PrinterURL                                       PKEY = "0b48f35a-be6e-4f17-b108-3c4073d1669a 15"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Bluetooth_DeviceAddress                 PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 1"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Bluetooth_Flags                         PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 3"     // UInt32 -- VT_UI4
	PKEY_DeviceInterface_Bluetooth_LastConnectedTime             PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 11"    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DeviceInterface_Bluetooth_Manufacturer                  PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 4"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Bluetooth_ModelNumber                   PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 5"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Bluetooth_ProductId                     PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 8"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Bluetooth_ProductVersion                PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 9"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Bluetooth_ServiceGuid                   PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 2"     // Guid -- VT_CLSID
	PKEY_DeviceInterface_Bluetooth_VendorId                      PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 7"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Bluetooth_VendorIdSource                PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 6"     // Byte -- VT_UI1
	PKEY_DeviceInterface_Hid_IsReadOnly                          PKEY = "cbf38310-4a17-4310-a1eb-247f0b67593b 4"     // Boolean -- VT_BOOL
	PKEY_DeviceInterface_Hid_ProductId                           PKEY = "cbf38310-4a17-4310-a1eb-247f0b67593b 6"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Hid_UsageId                             PKEY = "cbf38310-4a17-4310-a1eb-247f0b67593b 3"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Hid_UsagePage                           PKEY = "cbf38310-4a17-4310-a1eb-247f0b67593b 2"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Hid_VendorId                            PKEY = "cbf38310-4a17-4310-a1eb-247f0b67593b 5"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Hid_VersionNumber                       PKEY = "cbf38310-4a17-4310-a1eb-247f0b67593b 7"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_PrinterDriverDirectory                  PKEY = "847c66de-b8d6-4af9-abc3-6f4f926bc039 14"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_PrinterDriverName                       PKEY = "afc47170-14f5-498c-8f30-b0d19be449c6 11"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_PrinterEnumerationFlag                  PKEY = "a00742a1-cd8c-4b37-95ab-70755587767a 3"     // UInt32 -- VT_UI4
	PKEY_DeviceInterface_PrinterName                             PKEY = "0a7b84ef-0c27-463f-84ef-06c5070001be 10"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_PrinterPortName                         PKEY = "eec7b761-6f94-41b1-949f-c729720dd13c 12"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Proximity_SupportsNfc                   PKEY = "fb3842cd-9e2a-4f83-8fcc-4b0761139ae9 2"     // Boolean -- VT_BOOL
	PKEY_DeviceInterface_Serial_PortName                         PKEY = "4c6bf15c-4c03-4aac-91f5-64c0f852bcf4 4"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DeviceInterface_Serial_UsbProductId                     PKEY = "4c6bf15c-4c03-4aac-91f5-64c0f852bcf4 3"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_Serial_UsbVendorId                      PKEY = "4c6bf15c-4c03-4aac-91f5-64c0f852bcf4 2"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_WinUsb_DeviceInterfaceClasses           PKEY = "95e127b5-79cc-4e83-9c9e-8422187b3e0e 7"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_DeviceInterface_WinUsb_UsbClass                         PKEY = "95e127b5-79cc-4e83-9c9e-8422187b3e0e 4"     // Byte -- VT_UI1
	PKEY_DeviceInterface_WinUsb_UsbProductId                     PKEY = "95e127b5-79cc-4e83-9c9e-8422187b3e0e 3"     // UInt16 -- VT_UI2
	PKEY_DeviceInterface_WinUsb_UsbProtocol                      PKEY = "95e127b5-79cc-4e83-9c9e-8422187b3e0e 6"     // Byte -- VT_UI1
	PKEY_DeviceInterface_WinUsb_UsbSubClass                      PKEY = "95e127b5-79cc-4e83-9c9e-8422187b3e0e 5"     // Byte -- VT_UI1
	PKEY_DeviceInterface_WinUsb_UsbVendorId                      PKEY = "95e127b5-79cc-4e83-9c9e-8422187b3e0e 2"     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_AepId                                       PKEY = "3b2ce006-5e61-4fde-bab8-9b8aac9b26df 8"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Aep_Bluetooth_Cod_Major                         PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 2"     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Cod_Minor                         PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 3"     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Audio                PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 10"    // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Capturing            PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 8"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Information          PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 12"    // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_LimitedDiscovery     PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 4"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Networking           PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 6"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_ObjectXfer           PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 9"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Positioning          PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 5"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Rendering            PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 7"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_Cod_Services_Telephony            PKEY = "5fbd34cd-561a-412e-ba98-478a6b0fef1d 11"    // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Bluetooth_LastSeenTime                      PKEY = "2bd67d8b-8beb-48d5-87e0-6cda3428040a 12"    // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Devices_Aep_Bluetooth_Le_AddressType                    PKEY = "995ef0b0-7eb3-4a8b-b9ce-068bb3f4af69 4"     // Byte -- VT_UI1
	PKEY_Devices_Aep_Bluetooth_Le_Appearance                     PKEY = "995ef0b0-7eb3-4a8b-b9ce-068bb3f4af69 1"     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Le_Appearance_Category            PKEY = "995ef0b0-7eb3-4a8b-b9ce-068bb3f4af69 5"     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Le_Appearance_Subcategory         PKEY = "995ef0b0-7eb3-4a8b-b9ce-068bb3f4af69 6"     // UInt16 -- VT_UI2
	PKEY_Devices_Aep_Bluetooth_Le_IsConnectable                  PKEY = "995ef0b0-7eb3-4a8b-b9ce-068bb3f4af69 8"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_CanPair                                     PKEY = "e7c3fb29-caa7-4f47-8c8b-be59b330d4c5 3"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Category                                    PKEY = "a35996ab-11cf-4935-8b61-a6761081ecdf 17"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Aep_ContainerId                                 PKEY = "e7c3fb29-caa7-4f47-8c8b-be59b330d4c5 2"     // Guid -- VT_CLSID
	PKEY_Devices_Aep_DeviceAddress                               PKEY = "a35996ab-11cf-4935-8b61-a6761081ecdf 12"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Aep_IsConnected                                 PKEY = "a35996ab-11cf-4935-8b61-a6761081ecdf 7"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_IsPaired                                    PKEY = "a35996ab-11cf-4935-8b61-a6761081ecdf 16"    // Boolean -- VT_BOOL
	PKEY_Devices_Aep_IsPresent                                   PKEY = "a35996ab-11cf-4935-8b61-a6761081ecdf 9"     // Boolean -- VT_BOOL
	PKEY_Devices_Aep_Manufacturer                                PKEY = "a35996ab-11cf-4935-8b61-a6761081ecdf 5"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Aep_ModelId                                     PKEY = "a35996ab-11cf-4935-8b61-a6761081ecdf 4"     // Guid -- VT_CLSID
	PKEY_Devices_Aep_ModelName                                   PKEY = "a35996ab-11cf-4935-8b61-a6761081ecdf 3"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Aep_PointOfService_ConnectionTypes              PKEY = "d4bf61b3-442e-4ada-882d-fa7b70c832d9 6"     // Int32 -- VT_I4
	PKEY_Devices_Aep_ProtocolId                                  PKEY = "3b2ce006-5e61-4fde-bab8-9b8aac9b26df 5"     // Guid -- VT_CLSID
	PKEY_Devices_Aep_SignalStrength                              PKEY = "a35996ab-11cf-4935-8b61-a6761081ecdf 6"     // Int32 -- VT_I4
	PKEY_Devices_AepContainer_CanPair                            PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 3"     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_Categories                         PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 9"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_Children                           PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 2"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_ContainerId                        PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 12"    // Guid -- VT_CLSID
	PKEY_Devices_AepContainer_DialProtocol_InstalledApplications PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 6"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_IsPaired                           PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 4"     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_IsPresent                          PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 11"    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_Manufacturer                       PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 6"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AepContainer_ModelIds                           PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 8"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_ModelName                          PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 7"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AepContainer_ProtocolIds                        PKEY = "0bba1ede-7566-4f47-90ec-25fc567ced2a 13"    // Multivalue Guid -- VT_VECTOR | VT_CLSID  (For variants: VT_ARRAY | VT_CLSID)
	PKEY_Devices_AepContainer_SupportedUriSchemes                PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 5"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepContainer_SupportsAudio                      PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 2"     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsCapturing                  PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 11"    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsImages                     PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 4"     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsInformation                PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 14"    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsLimitedDiscovery           PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 7"     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsNetworking                 PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 9"     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsObjectTransfer             PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 12"    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsPositioning                PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 8"     // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsRendering                  PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 10"    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsTelephony                  PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 13"    // Boolean -- VT_BOOL
	PKEY_Devices_AepContainer_SupportsVideo                      PKEY = "6af55d45-38db-4495-acb0-d4728a3b8314 3"     // Boolean -- VT_BOOL
	PKEY_Devices_AepService_AepId                                PKEY = "c9c141a9-1b4c-4f17-a9d1-f298538cadb8 6"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AepService_Bluetooth_CacheMode                  PKEY = "9744311e-7951-4b2e-b6f0-ecb293cac119 5"     // Byte -- VT_UI1
	PKEY_Devices_AepService_Bluetooth_ServiceGuid                PKEY = "a399aac7-c265-474e-b073-ffce57721716 2"     // Guid -- VT_CLSID
	PKEY_Devices_AepService_Bluetooth_TargetDevice               PKEY = "9744311e-7951-4b2e-b6f0-ecb293cac119 6"     // UInt64 -- VT_UI8
	PKEY_Devices_AepService_ContainerId                          PKEY = "71724756-3e74-4432-9b59-e7b2f668a593 4"     // Guid -- VT_CLSID
	PKEY_Devices_AepService_FriendlyName                         PKEY = "71724756-3e74-4432-9b59-e7b2f668a593 2"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AepService_IoT_ServiceInterfaces                PKEY = "79d94e82-4d79-45aa-821a-74858b4e4ca6 2"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AepService_ParentAepIsPaired                    PKEY = "c9c141a9-1b4c-4f17-a9d1-f298538cadb8 7"     // Boolean -- VT_BOOL
	PKEY_Devices_AepService_ProtocolId                           PKEY = "c9c141a9-1b4c-4f17-a9d1-f298538cadb8 5"     // Guid -- VT_CLSID
	PKEY_Devices_AepService_ServiceClassId                       PKEY = "71724756-3e74-4432-9b59-e7b2f668a593 3"     // Guid -- VT_CLSID
	PKEY_Devices_AepService_ServiceId                            PKEY = "c9c141a9-1b4c-4f17-a9d1-f298538cadb8 2"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_AppPackageFamilyName                            PKEY = "51236583-0c4a-4fe8-b81f-166aec13f510 100"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_AudioDevice_Microphone_IsFarField               PKEY = "8943b373-388c-4395-b557-bc6dbaffafdb 6"     // Boolean -- VT_BOOL
	PKEY_Devices_AudioDevice_Microphone_SensitivityInDbfs        PKEY = "8943b373-388c-4395-b557-bc6dbaffafdb 3"     // Double -- VT_R8
	PKEY_Devices_AudioDevice_Microphone_SensitivityInDbfs2       PKEY = "8943b373-388c-4395-b557-bc6dbaffafdb 5"     // Double -- VT_R8
	PKEY_Devices_AudioDevice_Microphone_SignalToNoiseRatioInDb   PKEY = "8943b373-388c-4395-b557-bc6dbaffafdb 4"     // Double -- VT_R8
	PKEY_Devices_AudioDevice_RawProcessingSupported              PKEY = "8943b373-388c-4395-b557-bc6dbaffafdb 2"     // Boolean -- VT_BOOL
	PKEY_Devices_AudioDevice_SpeechProcessingSupported           PKEY = "fb1de864-e06d-47f4-82a6-8a0aef44493c 2"     // Boolean -- VT_BOOL
	PKEY_Devices_BatteryLife                                     PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 10"    // Byte -- VT_UI1
	PKEY_Devices_BatteryPlusCharging                             PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 22"    // Byte -- VT_UI1
	PKEY_Devices_BatteryPlusChargingText                         PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 23"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Category                                        PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 91"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_CategoryGroup                                   PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 94"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_CategoryIds                                     PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 90"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_CategoryPlural                                  PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 92"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_ChallengeAep                                    PKEY = "0774315e-b714-48ec-8de8-8125c077ac11 2"     // Boolean -- VT_BOOL
	PKEY_Devices_ChargingState                                   PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 11"    // Byte -- VT_UI1
	PKEY_Devices_Children                                        PKEY = "4340a6c5-93fa-4706-972c-7b648008a5a7 9"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_ClassGuid                                       PKEY = "a45c254e-df1c-4efd-8020-67d146a850e0 10"    // Guid -- VT_CLSID
	PKEY_Devices_CompatibleIds                                   PKEY = "a45c254e-df1c-4efd-8020-67d146a850e0 4"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Connected                                       PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 55"    // Boolean -- VT_BOOL
	PKEY_Devices_ContainerId                                     PKEY = "8c7ed206-3f8a-4827-b3ab-ae9e1faefc6c 2"     // Guid -- VT_CLSID
	PKEY_Devices_DefaultTooltip                                  PKEY = "880f70a2-6082-47ac-8aab-a739d1a300c3 153"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DeviceCapabilities                              PKEY = "a45c254e-df1c-4efd-8020-67d146a850e0 17"    // UInt32 -- VT_UI4
	PKEY_Devices_DeviceCharacteristics                           PKEY = "a45c254e-df1c-4efd-8020-67d146a850e0 29"    // UInt32 -- VT_UI4
	PKEY_Devices_DeviceDescription1                              PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 81"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DeviceDescription2                              PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 82"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DeviceHasProblem                                PKEY = "540b947e-8b40-45bc-a8a2-6a0b894cbda2 6"     // Boolean -- VT_BOOL
	PKEY_Devices_DeviceInstanceId                                PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 256"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DeviceManufacturer                              PKEY = "a45c254e-df1c-4efd-8020-67d146a850e0 13"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_DevObjectType                                   PKEY = "13673f42-a3d6-49f6-b4da-ae46e0c5237c 2"     // UInt32 -- VT_UI4
	PKEY_Devices_DialProtocol_InstalledApplications              PKEY = "6845cc72-1b71-48c3-af86-b09171a19b14 3"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_DiscoveryMethod                                 PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 52"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Dnssd_Domain                                    PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 3"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_FullName                                  PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 5"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_HostName                                  PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 7"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_InstanceName                              PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 4"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_NetworkAdapterId                          PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 11"    // Guid -- VT_CLSID
	PKEY_Devices_Dnssd_PortNumber                                PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 12"    // UInt16 -- VT_UI2
	PKEY_Devices_Dnssd_Priority                                  PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 9"     // UInt16 -- VT_UI2
	PKEY_Devices_Dnssd_ServiceName                               PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 2"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Dnssd_TextAttributes                            PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 6"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Dnssd_Ttl                                       PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 10"    // UInt32 -- VT_UI4
	PKEY_Devices_Dnssd_Weight                                    PKEY = "bf79c0ab-bb74-4cee-b070-470b5ae202ea 8"     // UInt16 -- VT_UI2
	PKEY_Devices_FriendlyName                                    PKEY = "656a3bb3-ecc0-43fd-8477-4ae0404a96cd 12288" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_FunctionPaths                                   PKEY = "d08dd4c0-3a9e-462e-8290-7b636b2576b9 3"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_GlyphIcon                                       PKEY = "51236583-0c4a-4fe8-b81f-166aec13f510 123"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_HardwareIds                                     PKEY = "a45c254e-df1c-4efd-8020-67d146a850e0 3"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Icon                                            PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 57"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_InLocalMachineContainer                         PKEY = "8c7ed206-3f8a-4827-b3ab-ae9e1faefc6c 4"     // Boolean -- VT_BOOL
	PKEY_Devices_InterfaceClassGuid                              PKEY = "026e516e-b814-414b-83cd-856d6fef4822 4"     // Guid -- VT_CLSID
	PKEY_Devices_InterfaceEnabled                                PKEY = "026e516e-b814-414b-83cd-856d6fef4822 3"     // Boolean -- VT_BOOL
	PKEY_Devices_InterfacePaths                                  PKEY = "d08dd4c0-3a9e-462e-8290-7b636b2576b9 2"     // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_IpAddress                                       PKEY = "656a3bb3-ecc0-43fd-8477-4ae0404a96cd 12297" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_IsDefault                                       PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 86"    // Boolean -- VT_BOOL
	PKEY_Devices_IsNetworkConnected                              PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 85"    // Boolean -- VT_BOOL
	PKEY_Devices_IsShared                                        PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 84"    // Boolean -- VT_BOOL
	PKEY_Devices_IsSoftwareInstalling                            PKEY = "83da6326-97a6-4088-9453-a1923f573b29 9"     // Boolean -- VT_BOOL
	PKEY_Devices_LaunchDeviceStageFromExplorer                   PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 77"    // Boolean -- VT_BOOL
	PKEY_Devices_LocalMachine                                    PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 70"    // Boolean -- VT_BOOL
	PKEY_Devices_LocationPaths                                   PKEY = "a45c254e-df1c-4efd-8020-67d146a850e0 37"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Manufacturer                                    PKEY = "656a3bb3-ecc0-43fd-8477-4ae0404a96cd 8192"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_MetadataPath                                    PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 71"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_MicrophoneArray_Geometry                        PKEY = "a1829ea2-27eb-459e-935d-b2fad7b07762 2"     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_MissedCalls                                     PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 5"     // Byte -- VT_UI1
	PKEY_Devices_ModelId                                         PKEY = "80d81ea6-7473-4b0c-8216-efc11a2c4c8b 2"     // Guid -- VT_CLSID
	PKEY_Devices_ModelName                                       PKEY = "656a3bb3-ecc0-43fd-8477-4ae0404a96cd 8194"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_ModelNumber                                     PKEY = "656a3bb3-ecc0-43fd-8477-4ae0404a96cd 8195"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_NetworkedTooltip                                PKEY = "880f70a2-6082-47ac-8aab-a739d1a300c3 152"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_NetworkName                                     PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 7"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_NetworkType                                     PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 8"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_NewPictures                                     PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 4"     // UInt16 -- VT_UI2
	PKEY_Devices_Notification                                    PKEY = "06704b0c-e830-4c81-9178-91e4e95a80a0 3"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Notifications_LowBattery                        PKEY = "c4c07f2b-8524-4e66-ae3a-a6235f103beb 2"     // Byte -- VT_UI1
	PKEY_Devices_Notifications_MissedCall                        PKEY = "6614ef48-4efe-4424-9eda-c79f404edf3e 2"     // Byte -- VT_UI1
	PKEY_Devices_Notifications_NewMessage                        PKEY = "2be9260a-2012-4742-a555-f41b638b7dcb 2"     // Byte -- VT_UI1
	PKEY_Devices_Notifications_NewVoicemail                      PKEY = "59569556-0a08-4212-95b9-fae2ad6413db 2"     // Byte -- VT_UI1
	PKEY_Devices_Notifications_StorageFull                       PKEY = "a0e00ee1-f0c7-4d41-b8e7-26a7bd8d38b0 2"     // UInt64 -- VT_UI8
	PKEY_Devices_Notifications_StorageFullLinkText               PKEY = "a0e00ee1-f0c7-4d41-b8e7-26a7bd8d38b0 3"     // UInt64 -- VT_UI8
	PKEY_Devices_NotificationStore                               PKEY = "06704b0c-e830-4c81-9178-91e4e95a80a0 2"     // Object -- VT_UNKNOWN
	PKEY_Devices_NotWorkingProperly                              PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 83"    // Boolean -- VT_BOOL
	PKEY_Devices_Paired                                          PKEY = "78c34fc8-104a-4aca-9ea4-524d52996e57 56"    // Boolean -- VT_BOOL
	PKEY_Devices_Panel_PanelGroup                                PKEY = "8dbc9c86-97a9-4bff-9bc6-bfe95d3e6dad 3"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Panel_PanelId                                   PKEY = "8dbc9c86-97a9-4bff-9bc6-bfe95d3e6dad 2"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Parent                                          PKEY = "4340a6c5-93fa-4706-972c-7b648008a5a7 8"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_PhoneLineTransportDevice_Connected              PKEY = "aecf2fe8-1d00-4fee-8a6d-a70d719b772b 2"     // Boolean -- VT_BOOL
	PKEY_Devices_PhysicalDeviceLocation                          PKEY = "540b947e-8b40-45bc-a8a2-6a0b894cbda2 9"     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_PlaybackPositionPercent                         PKEY = "3633de59-6825-4381-a49b-9f6ba13a1471 5"     // UInt32 -- VT_UI4
	PKEY_Devices_PlaybackState                                   PKEY = "3633de59-6825-4381-a49b-9f6ba13a1471 2"     // Byte -- VT_UI1
	PKEY_Devices_PlaybackTitle                                   PKEY = "3633de59-6825-4381-a49b-9f6ba13a1471 3"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Present                                         PKEY = "540b947e-8b40-45bc-a8a2-6a0b894cbda2 5"     // Boolean -- VT_BOOL
	PKEY_Devices_PresentationUrl                                 PKEY = "656a3bb3-ecc0-43fd-8477-4ae0404a96cd 8198"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_PrimaryCategory                                 PKEY = "d08dd4c0-3a9e-462e-8290-7b636b2576b9 10"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_RemainingDuration                               PKEY = "3633de59-6825-4381-a49b-9f6ba13a1471 4"     // UInt64 -- VT_UI8
	PKEY_Devices_RestrictedInterface                             PKEY = "026e516e-b814-414b-83cd-856d6fef4822 6"     // Boolean -- VT_BOOL
	PKEY_Devices_Roaming                                         PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 9"     // Byte -- VT_UI1
	PKEY_Devices_SafeRemovalRequired                             PKEY = "afd97640-86a3-4210-b67c-289c41aabe55 2"     // Boolean -- VT_BOOL
	PKEY_Devices_SchematicName                                   PKEY = "026e516e-b814-414b-83cd-856d6fef4822 9"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_ServiceAddress                                  PKEY = "656a3bb3-ecc0-43fd-8477-4ae0404a96cd 16384" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_ServiceId                                       PKEY = "656a3bb3-ecc0-43fd-8477-4ae0404a96cd 16385" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_SharedTooltip                                   PKEY = "880f70a2-6082-47ac-8aab-a739d1a300c3 151"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_SignalStrength                                  PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 2"     // Byte -- VT_UI1
	PKEY_Devices_SmartCards_ReaderKind                           PKEY = "d6b5b883-18bd-4b4d-b2ec-9e38affeda82 2"     // Byte -- VT_UI1
	PKEY_Devices_Status                                          PKEY = "d08dd4c0-3a9e-462e-8290-7b636b2576b9 259"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_Status1                                         PKEY = "d08dd4c0-3a9e-462e-8290-7b636b2576b9 257"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_Status2                                         PKEY = "d08dd4c0-3a9e-462e-8290-7b636b2576b9 258"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_StorageCapacity                                 PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 12"    // UInt64 -- VT_UI8
	PKEY_Devices_StorageFreeSpace                                PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 13"    // UInt64 -- VT_UI8
	PKEY_Devices_StorageFreeSpacePercent                         PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 14"    // UInt32 -- VT_UI4
	PKEY_Devices_TextMessages                                    PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 3"     // Byte -- VT_UI1
	PKEY_Devices_Voicemail                                       PKEY = "49cd1f76-5626-4b17-a4e8-18b4aa1a2213 6"     // Byte -- VT_UI1
	PKEY_Devices_WiaDeviceType                                   PKEY = "6bdd1fc6-810f-11d0-bec7-08002be2092f 2"     // UInt32 -- VT_UI4
	PKEY_Devices_WiFi_InterfaceGuid                              PKEY = "ef1167eb-cbfc-4341-a568-a7c91a68982c 2"     // Guid -- VT_CLSID
	PKEY_Devices_WiFiDirect_DeviceAddress                        PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 13"    // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirect_GroupId                              PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 4"     // Guid -- VT_CLSID
	PKEY_Devices_WiFiDirect_InformationElements                  PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 12"    // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirect_InterfaceAddress                     PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 2"     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirect_InterfaceGuid                        PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 3"     // Guid -- VT_CLSID
	PKEY_Devices_WiFiDirect_IsConnected                          PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 5"     // Boolean -- VT_BOOL
	PKEY_Devices_WiFiDirect_IsLegacyDevice                       PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 7"     // Boolean -- VT_BOOL
	PKEY_Devices_WiFiDirect_IsMiracastLcpSupported               PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 9"     // Boolean -- VT_BOOL
	PKEY_Devices_WiFiDirect_IsVisible                            PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 6"     // Boolean -- VT_BOOL
	PKEY_Devices_WiFiDirect_MiracastVersion                      PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 8"     // UInt32 -- VT_UI4
	PKEY_Devices_WiFiDirect_Services                             PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 10"    // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Devices_WiFiDirect_SupportedChannelList                 PKEY = "1506935d-e3e7-450f-8637-82233ebe5f6e 11"    // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirectServices_AdvertisementId              PKEY = "31b37743-7c5e-4005-93e6-e953f92b82e9 5"     // UInt32 -- VT_UI4
	PKEY_Devices_WiFiDirectServices_RequestServiceInformation    PKEY = "31b37743-7c5e-4005-93e6-e953f92b82e9 7"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_WiFiDirectServices_ServiceAddress               PKEY = "31b37743-7c5e-4005-93e6-e953f92b82e9 2"     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirectServices_ServiceConfigMethods         PKEY = "31b37743-7c5e-4005-93e6-e953f92b82e9 6"     // UInt32 -- VT_UI4
	PKEY_Devices_WiFiDirectServices_ServiceInformation           PKEY = "31b37743-7c5e-4005-93e6-e953f92b82e9 4"     // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Devices_WiFiDirectServices_ServiceName                  PKEY = "31b37743-7c5e-4005-93e6-e953f92b82e9 3"     // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Devices_WinPhone8CameraFlags                            PKEY = "b7b4d61c-5a64-4187-a52e-b1539f359099 2"     // UInt32 -- VT_UI4
	PKEY_Devices_Wwan_InterfaceGuid                              PKEY = "ff1167eb-cbfc-4341-a568-a7c91a68982c 2"     // Guid -- VT_CLSID
	PKEY_Storage_Portable                                        PKEY = "4d1ebee8-0803-4774-9842-b77db50265e9 2"     // Boolean -- VT_BOOL
	PKEY_Storage_RemovableMedia                                  PKEY = "4d1ebee8-0803-4774-9842-b77db50265e9 3"     // Boolean -- VT_BOOL
	PKEY_Storage_SystemCritical                                  PKEY = "4d1ebee8-0803-4774-9842-b77db50265e9 4"     // Boolean -- VT_BOOL

	// Document properties

	PKEY_Document_ByteCount           PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 4"   // Int32 -- VT_I4
	PKEY_Document_CharacterCount      PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 16"  // Int32 -- VT_I4
	PKEY_Document_ClientID            PKEY = "276d7bb0-5b34-4fb0-aa4b-158ed12a1809 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_Contributor         PKEY = "f334115e-da1b-4509-9b3d-119504dc7abb 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Document_DateCreated         PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 12"  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Document_DatePrinted         PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 11"  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Document_DateSaved           PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 13"  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Document_Division            PKEY = "1e005ee6-bf27-428b-b01c-79676acd2870 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_DocumentID          PKEY = "e08805c8-e395-40df-80d2-54f0d6c43154 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_HiddenSlideCount    PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 9"   // Int32 -- VT_I4
	PKEY_Document_LastAuthor          PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 8"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_LineCount           PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 5"   // Int32 -- VT_I4
	PKEY_Document_Manager             PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 14"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_MultimediaClipCount PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 10"  // Int32 -- VT_I4
	PKEY_Document_NoteCount           PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 8"   // Int32 -- VT_I4
	PKEY_Document_PageCount           PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 14"  // Int32 -- VT_I4
	PKEY_Document_ParagraphCount      PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 6"   // Int32 -- VT_I4
	PKEY_Document_PresentationFormat  PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 3"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_RevisionNumber      PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 9"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_Security            PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 19"  // Int32 -- VT_I4
	PKEY_Document_SlideCount          PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 7"   // Int32 -- VT_I4
	PKEY_Document_Template            PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 7"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_TotalEditingTime    PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 10"  // UInt64 -- VT_UI8
	PKEY_Document_Version             PKEY = "d5cdd502-2e9c-101b-9397-08002b2cf9ae 29"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Document_WordCount           PKEY = "f29f85e0-4ff9-1068-ab91-08002b27b3d9 15"  // Int32 -- VT_I4

	// DRM properties

	PKEY_DRM_DatePlayExpires PKEY = "aeac19e4-89ae-4508-b9b7-bb867abee2ed 6" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DRM_DatePlayStarts  PKEY = "aeac19e4-89ae-4508-b9b7-bb867abee2ed 5" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_DRM_Description     PKEY = "aeac19e4-89ae-4508-b9b7-bb867abee2ed 3" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_DRM_IsDisabled      PKEY = "aeac19e4-89ae-4508-b9b7-bb867abee2ed 7" // Boolean -- VT_BOOL
	PKEY_DRM_IsProtected     PKEY = "aeac19e4-89ae-4508-b9b7-bb867abee2ed 2" // Boolean -- VT_BOOL
	PKEY_DRM_PlayCount       PKEY = "aeac19e4-89ae-4508-b9b7-bb867abee2ed 4" // UInt32 -- VT_UI4

	// GPS properties

	PKEY_GPS_Altitude                 PKEY = "827edb4f-5b73-44a7-891d-fdffabea35ca 100" // Double -- VT_R8
	PKEY_GPS_AltitudeDenominator      PKEY = "78342dcb-e358-4145-ae9a-6bfe4e0f9f51 100" // UInt32 -- VT_UI4
	PKEY_GPS_AltitudeNumerator        PKEY = "2dad1eb7-816d-40d3-9ec3-c9773be2aade 100" // UInt32 -- VT_UI4
	PKEY_GPS_AltitudeRef              PKEY = "46ac629d-75ea-4515-867f-6dc4321c5844 100" // Byte -- VT_UI1
	PKEY_GPS_AreaInformation          PKEY = "972e333e-ac7e-49f1-8adf-a70d07a9bcab 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Date                     PKEY = "3602c812-0f3b-45f0-85ad-603468d69423 100" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_GPS_DestBearing              PKEY = "c66d4b3c-e888-47cc-b99f-9dca3ee34dea 100" // Double -- VT_R8
	PKEY_GPS_DestBearingDenominator   PKEY = "7abcf4f8-7c3f-4988-ac91-8d2c2e97eca5 100" // UInt32 -- VT_UI4
	PKEY_GPS_DestBearingNumerator     PKEY = "ba3b1da9-86ee-4b5d-a2a4-a271a429f0cf 100" // UInt32 -- VT_UI4
	PKEY_GPS_DestBearingRef           PKEY = "9ab84393-2a0f-4b75-bb22-7279786977cb 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_DestDistance             PKEY = "a93eae04-6804-4f24-ac81-09b266452118 100" // Double -- VT_R8
	PKEY_GPS_DestDistanceDenominator  PKEY = "9bc2c99b-ac71-4127-9d1c-2596d0d7dcb7 100" // UInt32 -- VT_UI4
	PKEY_GPS_DestDistanceNumerator    PKEY = "2bda47da-08c6-4fe1-80bc-a72fc517c5d0 100" // UInt32 -- VT_UI4
	PKEY_GPS_DestDistanceRef          PKEY = "ed4df2d3-8695-450b-856f-f5c1c53acb66 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_DestLatitude             PKEY = "9d1d7cc5-5c39-451c-86b3-928e2d18cc47 100" // Multivalue Double -- VT_VECTOR | VT_R8  (For variants: VT_ARRAY | VT_R8)
	PKEY_GPS_DestLatitudeDenominator  PKEY = "3a372292-7fca-49a7-99d5-e47bb2d4e7ab 100" // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_DestLatitudeNumerator    PKEY = "ecf4b6f6-d5a6-433c-bb92-4076650fc890 100" // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_DestLatitudeRef          PKEY = "cea820b9-ce61-4885-a128-005d9087c192 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_DestLongitude            PKEY = "47a96261-cb4c-4807-8ad3-40b9d9dbc6bc 100" // Multivalue Double -- VT_VECTOR | VT_R8  (For variants: VT_ARRAY | VT_R8)
	PKEY_GPS_DestLongitudeDenominator PKEY = "425d69e5-48ad-4900-8d80-6eb6b8d0ac86 100" // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_DestLongitudeNumerator   PKEY = "a3250282-fb6d-48d5-9a89-dbcace75cccf 100" // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_DestLongitudeRef         PKEY = "182c1ea6-7c1c-4083-ab4b-ac6c9f4ed128 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Differential             PKEY = "aaf4ee25-bd3b-4dd7-bfc4-47f77bb00f6d 100" // UInt16 -- VT_UI2
	PKEY_GPS_DOP                      PKEY = "0cf8fb02-1837-42f1-a697-a7017aa289b9 100" // Double -- VT_R8
	PKEY_GPS_DOPDenominator           PKEY = "a0be94c5-50ba-487b-bd35-0654be8881ed 100" // UInt32 -- VT_UI4
	PKEY_GPS_DOPNumerator             PKEY = "47166b16-364f-4aa0-9f31-e2ab3df449c3 100" // UInt32 -- VT_UI4
	PKEY_GPS_ImgDirection             PKEY = "16473c91-d017-4ed9-ba4d-b6baa55dbcf8 100" // Double -- VT_R8
	PKEY_GPS_ImgDirectionDenominator  PKEY = "10b24595-41a2-4e20-93c2-5761c1395f32 100" // UInt32 -- VT_UI4
	PKEY_GPS_ImgDirectionNumerator    PKEY = "dc5877c7-225f-45f7-bac7-e81334b6130a 100" // UInt32 -- VT_UI4
	PKEY_GPS_ImgDirectionRef          PKEY = "a4aaa5b7-1ad0-445f-811a-0f8f6e67f6b5 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Latitude                 PKEY = "8727cfff-4868-4ec6-ad5b-81b98521d1ab 100" // Multivalue Double -- VT_VECTOR | VT_R8  (For variants: VT_ARRAY | VT_R8)
	PKEY_GPS_LatitudeDecimal          PKEY = "0f55cde2-4f49-450d-92c1-dcd16301b1b7 100" // Double -- VT_R8
	PKEY_GPS_LatitudeDenominator      PKEY = "16e634ee-2bff-497b-bd8a-4341ad39eeb9 100" // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_LatitudeNumerator        PKEY = "7ddaaad1-ccc8-41ae-b750-b2cb8031aea2 100" // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_LatitudeRef              PKEY = "029c0252-5b86-46c7-aca0-2769ffc8e3d4 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Longitude                PKEY = "c4c4dbb2-b593-466b-bbda-d03d27d5e43a 100" // Multivalue Double -- VT_VECTOR | VT_R8  (For variants: VT_ARRAY | VT_R8)
	PKEY_GPS_LongitudeDecimal         PKEY = "4679c1b5-844d-4590-baf5-f322231f1b81 100" // Double -- VT_R8
	PKEY_GPS_LongitudeDenominator     PKEY = "be6e176c-4534-4d2c-ace5-31dedac1606b 100" // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_LongitudeNumerator       PKEY = "02b0f689-a914-4e45-821d-1dda452ed2c4 100" // Multivalue UInt32 -- VT_VECTOR | VT_UI4  (For variants: VT_ARRAY | VT_UI4)
	PKEY_GPS_LongitudeRef             PKEY = "33dcf22b-28d5-464c-8035-1ee9efd25278 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_MapDatum                 PKEY = "2ca2dae6-eddc-407d-bef1-773942abfa95 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_MeasureMode              PKEY = "a015ed5d-aaea-4d58-8a86-3c586920ea0b 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_ProcessingMethod         PKEY = "59d49e61-840f-4aa9-a939-e2099b7f6399 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Satellites               PKEY = "467ee575-1f25-4557-ad4e-b8b58b0d9c15 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Speed                    PKEY = "da5d0862-6e76-4e1b-babd-70021bd25494 100" // Double -- VT_R8
	PKEY_GPS_SpeedDenominator         PKEY = "7d122d5a-ae5e-4335-8841-d71e7ce72f53 100" // UInt32 -- VT_UI4
	PKEY_GPS_SpeedNumerator           PKEY = "acc9ce3d-c213-4942-8b48-6d0820f21c6d 100" // UInt32 -- VT_UI4
	PKEY_GPS_SpeedRef                 PKEY = "ecf7f4c9-544f-4d6d-9d98-8ad79adaf453 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Status                   PKEY = "125491f4-818f-46b2-91b5-d537753617b2 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_Track                    PKEY = "76c09943-7c33-49e3-9e7e-cdba872cfada 100" // Double -- VT_R8
	PKEY_GPS_TrackDenominator         PKEY = "c8d1920c-01f6-40c0-ac86-2f3a4ad00770 100" // UInt32 -- VT_UI4
	PKEY_GPS_TrackNumerator           PKEY = "702926f4-44a6-43e1-ae71-45627116893b 100" // UInt32 -- VT_UI4
	PKEY_GPS_TrackRef                 PKEY = "35dbe6fe-44c3-4400-aaae-d2c799c407e8 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_GPS_VersionID                PKEY = "22704da4-c6b2-4a99-8e56-f16df8c92599 100" // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)

	// History properties

	PKEY_History_VisitCount PKEY = "5cbf2787-48cf-4208-b90e-ee5e5d420294 7" // Int32 -- VT_I4

	// Image properties

	PKEY_Image_BitDepth                          PKEY = "6444048f-4c8b-11d1-8b70-080036b11a03 7"     // UInt32 -- VT_UI4
	PKEY_Image_ColorSpace                        PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 40961" // UInt16 -- VT_UI2
	PKEY_Image_CompressedBitsPerPixel            PKEY = "364b6fa9-37ab-482a-be2b-ae02f60d4318 100"   // Double -- VT_R8
	PKEY_Image_CompressedBitsPerPixelDenominator PKEY = "1f8844e1-24ad-4508-9dfd-5326a415ce02 100"   // UInt32 -- VT_UI4
	PKEY_Image_CompressedBitsPerPixelNumerator   PKEY = "d21a7148-d32c-4624-8900-277210f79c0f 100"   // UInt32 -- VT_UI4
	PKEY_Image_Compression                       PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 259"   // UInt16 -- VT_UI2
	PKEY_Image_CompressionText                   PKEY = "3f08e66f-2f44-4bb9-a682-ac35d2562322 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Image_Dimensions                        PKEY = "6444048f-4c8b-11d1-8b70-080036b11a03 13"    // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Image_HorizontalResolution              PKEY = "6444048f-4c8b-11d1-8b70-080036b11a03 5"     // Double -- VT_R8
	PKEY_Image_HorizontalSize                    PKEY = "6444048f-4c8b-11d1-8b70-080036b11a03 3"     // UInt32 -- VT_UI4
	PKEY_Image_ImageID                           PKEY = "10dabe05-32aa-4c29-bf1a-63e2d220587f 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Image_ResolutionUnit                    PKEY = "19b51fa6-1f92-4a5c-ab48-7df0abd67444 100"   // Int16 -- VT_I2
	PKEY_Image_VerticalResolution                PKEY = "6444048f-4c8b-11d1-8b70-080036b11a03 6"     // Double -- VT_R8
	PKEY_Image_VerticalSize                      PKEY = "6444048f-4c8b-11d1-8b70-080036b11a03 4"     // UInt32 -- VT_UI4

	// Journal properties

	PKEY_Journal_Contacts  PKEY = "dea7c82c-1d89-4a66-9427-a4e3debabcb1 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Journal_EntryType PKEY = "95beb1fc-326d-4644-b396-cd3ed90e6ddf 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// LayoutPattern properties

	PKEY_LayoutPattern_ContentViewModeForBrowse PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 500" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_LayoutPattern_ContentViewModeForSearch PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 501" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Link properties

	PKEY_History_SelectionCount    PKEY = "1ce0d6bc-536c-4600-b0dd-7e0c66b350d5 8"   // UInt32 -- VT_UI4
	PKEY_History_TargetUrlHostName PKEY = "1ce0d6bc-536c-4600-b0dd-7e0c66b350d5 9"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_Arguments            PKEY = "436f2667-14e2-4feb-b30a-146c53b5b674 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_Comment              PKEY = "b9b4b3fc-2b51-4a42-b5d8-324146afcf25 5"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_DateVisited          PKEY = "5cbf2787-48cf-4208-b90e-ee5e5d420294 23"  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Link_Description          PKEY = "5cbf2787-48cf-4208-b90e-ee5e5d420294 21"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_FeedItemLocalId      PKEY = "8a2f99f9-3c37-465d-a8d7-69777a246d0c 2"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_Status               PKEY = "b9b4b3fc-2b51-4a42-b5d8-324146afcf25 3"   // Int32 -- VT_I4
	PKEY_Link_TargetExtension      PKEY = "7a7d76f4-b630-4bd7-95ff-37cc51a975c9 2"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Link_TargetParsingPath    PKEY = "b9b4b3fc-2b51-4a42-b5d8-324146afcf25 2"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_TargetSFGAOFlags     PKEY = "b9b4b3fc-2b51-4a42-b5d8-324146afcf25 8"   // UInt32 -- VT_UI4
	PKEY_Link_TargetUrlHostName    PKEY = "8a2f99f9-3c37-465d-a8d7-69777a246d0c 5"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_TargetUrlPath        PKEY = "8a2f99f9-3c37-465d-a8d7-69777a246d0c 6"   // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Media properties

	PKEY_Media_AuthorUrl                 PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 32"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_AverageLevel              PKEY = "09edd5b6-b301-43c5-9990-d00302effd46 100" // UInt32 -- VT_UI4
	PKEY_Media_ClassPrimaryID            PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 13"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ClassSecondaryID          PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 14"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_CollectionGroupID         PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 24"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_CollectionID              PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 25"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ContentDistributor        PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 18"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ContentID                 PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 26"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_CreatorApplication        PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 27"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_CreatorApplicationVersion PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 28"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_DateEncoded               PKEY = "2e4b640d-5019-46d8-8881-55414cc5caa0 100" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Media_DateReleased              PKEY = "de41cc29-6971-4290-b472-f59f2e2f31e2 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_DlnaProfileID             PKEY = "cfa31b45-525d-4998-bb44-3f7d81542fa4 100" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Media_Duration                  PKEY = "64440490-4c8b-11d1-8b70-080036b11a03 3"   // UInt64 -- VT_UI8
	PKEY_Media_DVDID                     PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 15"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_EncodedBy                 PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 36"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_EncodingSettings          PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 37"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_EpisodeNumber             PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 100" // UInt32 -- VT_UI4
	PKEY_Media_FrameCount                PKEY = "6444048f-4c8b-11d1-8b70-080036b11a03 12"  // UInt32 -- VT_UI4
	PKEY_Media_MCDI                      PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 16"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_MetadataContentProvider   PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 17"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_Producer                  PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 22"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Media_PromotionUrl              PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 33"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ProtectionType            PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 38"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ProviderRating            PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 39"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ProviderStyle             PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 40"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_Publisher                 PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 30"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_SeasonNumber              PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 101" // UInt32 -- VT_UI4
	PKEY_Media_SeriesName                PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 42"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_SubscriptionContentId     PKEY = "9aebae7a-9644-487d-a92c-657585ed751a 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_SubTitle                  PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 38"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ThumbnailLargePath        PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 47"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ThumbnailLargeUri         PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 48"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ThumbnailSmallPath        PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 49"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_ThumbnailSmallUri         PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 50"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_UniqueFileIdentifier      PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 35"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_UserNoAutoInfo            PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 41"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_UserWebUrl                PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 34"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Media_Writer                    PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 23"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Media_Year                      PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 5"   // UInt32 -- VT_UI4

	// Message properties

	PKEY_Message_AttachmentContents PKEY = "3143bf7c-80a8-4854-8880-e2e40189bdd0 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_AttachmentNames    PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 21"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_BccAddress         PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 2"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_BccName            PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 3"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_CcAddress          PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 4"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_CcName             PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 5"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_ConversationID     PKEY = "dc8f80bd-af1e-4289-85b6-3dfc1b493992 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_ConversationIndex  PKEY = "dc8f80bd-af1e-4289-85b6-3dfc1b493992 101" // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Message_DateReceived       PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 20"  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Message_DateSent           PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 19"  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Message_Flags              PKEY = "a82d9ee7-ca67-4312-965e-226bcea85023 100" // Int32 -- VT_I4
	PKEY_Message_FromAddress        PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 13"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_FromName           PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 14"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_HasAttachments     PKEY = "9c1fcf74-2d97-41ba-b4ae-cb2e3661a6e4 8"   // Boolean -- VT_BOOL
	PKEY_Message_IsFwdOrReply       PKEY = "9a9bc088-4f6d-469e-9919-e705412040f9 100" // Int32 -- VT_I4
	PKEY_Message_MessageClass       PKEY = "cd9ed458-08ce-418f-a70e-f912c7bb9c5c 103" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_Participants       PKEY = "1a9ba605-8e7c-4d11-ad7d-a50ada18ba1b 2"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_ProofInProgress    PKEY = "9098f33c-9a7d-48a8-8de5-2e1227a64e91 100" // Boolean -- VT_BOOL
	PKEY_Message_SenderAddress      PKEY = "0be1c8e7-1981-4676-ae14-fdd78f05a6e7 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_SenderName         PKEY = "0da41cfa-d224-4a18-ae2f-596158db4b3a 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_Store              PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 15"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_ToAddress          PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 16"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Message_ToDoFlags          PKEY = "1f856a9f-6900-4aba-9505-2d5f1b4d66cb 100" // Int32 -- VT_I4
	PKEY_Message_ToDoTitle          PKEY = "bccc8a3c-8cef-42e5-9b1c-c69079398bc7 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Message_ToName             PKEY = "e3e0584c-b788-4a5a-bb20-7f5a44c9acdd 17"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)

	// Music properties

	PKEY_Music_AlbumArtist             PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 13"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_AlbumArtistSortOverride PKEY = "f1fdb4af-f78c-466c-bb05-56e92db0b8ec 103" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_AlbumID                 PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_AlbumTitle              PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 4"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_AlbumTitleSortOverride  PKEY = "13eb7ffc-ec89-4346-b19d-ccc6f1784223 101" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Artist                  PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 2"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_ArtistSortOverride      PKEY = "deeb2db5-0696-4ce0-94fe-a01f77a45fb5 102" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_BeatsPerMinute          PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 35"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Composer                PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 19"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_ComposerSortOverride    PKEY = "00bc20a3-bd48-4085-872c-a88d77f5097e 105" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_Conductor               PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 36"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_ContentGroupDescription PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 33"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_DiscNumber              PKEY = "6afe7437-9bcd-49c7-80fe-4a5c65fa5874 104" // UInt32 -- VT_UI4
	PKEY_Music_DisplayArtist           PKEY = "fd122953-fa93-4ef7-92c3-04c946b2f7c8 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Genre                   PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 11"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Music_InitialKey              PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 34"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_IsCompilation           PKEY = "c449d5cb-9ea4-4809-82e8-af9d59ded6d1 100" // Boolean -- VT_BOOL
	PKEY_Music_Lyrics                  PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 12"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Mood                    PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 39"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_PartOfSet               PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 37"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_Period                  PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 31"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Music_SynchronizedLyrics      PKEY = "6b223b6a-162e-4aa9-b39f-05d678fc6d77 100" // Blob -- VT_BLOB
	PKEY_Music_TrackNumber             PKEY = "56a3372e-ce9c-11d2-9f0e-006097c686f6 7"   // UInt32 -- VT_UI4

	// Note properties

	PKEY_Note_Color     PKEY = "4776cafa-bce4-4cb1-a23e-265e76d8eb11 100" // UInt16 -- VT_UI2
	PKEY_Note_ColorText PKEY = "46b4e8de-cdb2-440d-885c-1658eb65b914 100" // String -- VT_LPWSTR  (For variants: VT_BSTR

	// Photo properties

	PKEY_Photo_Aperture                         PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 37378" // Double -- VT_R8
	PKEY_Photo_ApertureDenominator              PKEY = "e1a9a38b-6685-46bd-875e-570dc7ad7320 100"   // UInt32 -- VT_UI4
	PKEY_Photo_ApertureNumerator                PKEY = "0337ecec-39fb-4581-a0bd-4c4cc51e9914 100"   // UInt32 -- VT_UI4
	PKEY_Photo_Brightness                       PKEY = "1a701bf6-478c-4361-83ab-3701bb053c58 100"   // Double -- VT_R8
	PKEY_Photo_BrightnessDenominator            PKEY = "6ebe6946-2321-440a-90f0-c043efd32476 100"   // UInt32 -- VT_UI4
	PKEY_Photo_BrightnessNumerator              PKEY = "9e7d118f-b314-45a0-8cfb-d654b917c9e9 100"   // UInt32 -- VT_UI4
	PKEY_Photo_CameraManufacturer               PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 271"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_CameraModel                      PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 272"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_CameraSerialNumber               PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 273"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_Contrast                         PKEY = "2a785ba9-8d23-4ded-82e6-60a350c86a10 100"   // UInt32 -- VT_UI4
	PKEY_Photo_ContrastText                     PKEY = "59dde9f2-5253-40ea-9a8b-479e96c6249a 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_DateTaken                        PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 36867" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Photo_DigitalZoom                      PKEY = "f85bf840-a925-4bc2-b0c4-8e36b598679e 100"   // Double -- VT_R8
	PKEY_Photo_DigitalZoomDenominator           PKEY = "745baf0e-e5c1-4cfb-8a1b-d031a0a52393 100"   // UInt32 -- VT_UI4
	PKEY_Photo_DigitalZoomNumerator             PKEY = "16cbb924-6500-473b-a5be-f1599bcbe413 100"   // UInt32 -- VT_UI4
	PKEY_Photo_Event                            PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 18248" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Photo_EXIFVersion                      PKEY = "d35f743a-eb2e-47f2-a286-844132cb1427 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ExposureBias                     PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 37380" // Double -- VT_R8
	PKEY_Photo_ExposureBiasDenominator          PKEY = "ab205e50-04b7-461c-a18c-2f233836e627 100"   // Int32 -- VT_I4
	PKEY_Photo_ExposureBiasNumerator            PKEY = "738bf284-1d87-420b-92cf-5834bf6ef9ed 100"   // Int32 -- VT_I4
	PKEY_Photo_ExposureIndex                    PKEY = "967b5af8-995a-46ed-9e11-35b3c5b9782d 100"   // Double -- VT_R8
	PKEY_Photo_ExposureIndexDenominator         PKEY = "93112f89-c28b-492f-8a9d-4be2062cee8a 100"   // UInt32 -- VT_UI4
	PKEY_Photo_ExposureIndexNumerator           PKEY = "cdedcf30-8919-44df-8f4c-4eb2ffdb8d89 100"   // UInt32 -- VT_UI4
	PKEY_Photo_ExposureProgram                  PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 34850" // UInt32 -- VT_UI4
	PKEY_Photo_ExposureProgramText              PKEY = "fec690b7-5f30-4646-ae47-4caafba884a3 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ExposureTime                     PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 33434" // Double -- VT_R8
	PKEY_Photo_ExposureTimeDenominator          PKEY = "55e98597-ad16-42e0-b624-21599a199838 100"   // UInt32 -- VT_UI4
	PKEY_Photo_ExposureTimeNumerator            PKEY = "257e44e2-9031-4323-ac38-85c552871b2e 100"   // UInt32 -- VT_UI4
	PKEY_Photo_Flash                            PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 37385" // Byte -- VT_UI1
	PKEY_Photo_FlashEnergy                      PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 41483" // Double -- VT_R8
	PKEY_Photo_FlashEnergyDenominator           PKEY = "d7b61c70-6323-49cd-a5fc-c84277162c97 100"   // UInt32 -- VT_UI4
	PKEY_Photo_FlashEnergyNumerator             PKEY = "fcad3d3d-0858-400f-aaa3-2f66cce2a6bc 100"   // UInt32 -- VT_UI4
	PKEY_Photo_FlashManufacturer                PKEY = "aabaf6c9-e0c5-4719-8585-57b103e584fe 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_FlashModel                       PKEY = "fe83bb35-4d1a-42e2-916b-06f3e1af719e 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_FlashText                        PKEY = "6b8b68f6-200b-47ea-8d25-d8050f57339f 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_FNumber                          PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 33437" // Double -- VT_R8
	PKEY_Photo_FNumberDenominator               PKEY = "e92a2496-223b-4463-a4e3-30eabba79d80 100"   // UInt32 -- VT_UI4
	PKEY_Photo_FNumberNumerator                 PKEY = "1b97738a-fdfc-462f-9d93-1957e08be90c 100"   // UInt32 -- VT_UI4
	PKEY_Photo_FocalLength                      PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 37386" // Double -- VT_R8
	PKEY_Photo_FocalLengthDenominator           PKEY = "305bc615-dca1-44a5-9fd4-10c0ba79412e 100"   // UInt32 -- VT_UI4
	PKEY_Photo_FocalLengthInFilm                PKEY = "a0e74609-b84d-4f49-b860-462bd9971f98 100"   // UInt16 -- VT_UI2
	PKEY_Photo_FocalLengthNumerator             PKEY = "776b6b3b-1e3d-4b0c-9a0e-8fbaf2a8492a 100"   // UInt32 -- VT_UI4
	PKEY_Photo_FocalPlaneXResolution            PKEY = "cfc08d97-c6f7-4484-89dd-ebef4356fe76 100"   // Double -- VT_R8
	PKEY_Photo_FocalPlaneXResolutionDenominator PKEY = "0933f3f5-4786-4f46-a8e8-d64dd37fa521 100"   // UInt32 -- VT_UI4
	PKEY_Photo_FocalPlaneXResolutionNumerator   PKEY = "dccb10af-b4e2-4b88-95f9-031b4d5ab490 100"   // UInt32 -- VT_UI4
	PKEY_Photo_FocalPlaneYResolution            PKEY = "4fffe4d0-914f-4ac4-8d6f-c9c61de169b1 100"   // Double -- VT_R8
	PKEY_Photo_FocalPlaneYResolutionDenominator PKEY = "1d6179a6-a876-4031-b013-3347b2b64dc8 100"   // UInt32 -- VT_UI4
	PKEY_Photo_FocalPlaneYResolutionNumerator   PKEY = "a2e541c5-4440-4ba8-867e-75cfc06828cd 100"   // UInt32 -- VT_UI4
	PKEY_Photo_GainControl                      PKEY = "fa304789-00c7-4d80-904a-1e4dcc7265aa 100"   // Double -- VT_R8
	PKEY_Photo_GainControlDenominator           PKEY = "42864dfd-9da4-4f77-bded-4aad7b256735 100"   // UInt32 -- VT_UI4
	PKEY_Photo_GainControlNumerator             PKEY = "8e8ecf7c-b7b8-4eb8-a63f-0ee715c96f9e 100"   // UInt32 -- VT_UI4
	PKEY_Photo_GainControlText                  PKEY = "c06238b2-0bf9-4279-a723-25856715cb9d 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ISOSpeed                         PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 34855" // UInt16 -- VT_UI2
	PKEY_Photo_LensManufacturer                 PKEY = "e6ddcaf7-29c5-4f0a-9a68-d19412ec7090 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_LensModel                        PKEY = "e1277516-2b5f-4869-89b1-2e585bd38b7a 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_LightSource                      PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 37384" // UInt32 -- VT_UI4
	PKEY_Photo_MakerNote                        PKEY = "fa303353-b659-4052-85e9-bcac79549b84 100"   // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_Photo_MakerNoteOffset                  PKEY = "813f4124-34e6-4d17-ab3e-6b1f3c2247a1 100"   // UInt64 -- VT_UI8
	PKEY_Photo_MaxAperture                      PKEY = "08f6d7c2-e3f2-44fc-af1e-5aa5c81a2d3e 100"   // Double -- VT_R8
	PKEY_Photo_MaxApertureDenominator           PKEY = "c77724d4-601f-46c5-9b89-c53f93bceb77 100"   // UInt32 -- VT_UI4
	PKEY_Photo_MaxApertureNumerator             PKEY = "c107e191-a459-44c5-9ae6-b952ad4b906d 100"   // UInt32 -- VT_UI4
	PKEY_Photo_MeteringMode                     PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 37383" // UInt16 -- VT_UI2
	PKEY_Photo_MeteringModeText                 PKEY = "f628fd8c-7ba8-465a-a65b-c5aa79263a9e 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_Orientation                      PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 274"   // UInt16 -- VT_UI2
	PKEY_Photo_OrientationText                  PKEY = "a9ea193c-c511-498a-a06b-58e2776dcc28 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_PeopleNames                      PKEY = "e8309b6e-084c-49b4-b1fc-90a80331b638 100"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_Photo_PhotometricInterpretation        PKEY = "341796f1-1df9-4b1c-a564-91bdefa43877 100"   // UInt16 -- VT_UI2
	PKEY_Photo_PhotometricInterpretationText    PKEY = "821437d6-9eab-4765-a589-3b1cbbd22a61 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ProgramMode                      PKEY = "6d217f6d-3f6a-4825-b470-5f03ca2fbe9b 100"   // UInt32 -- VT_UI4
	PKEY_Photo_ProgramModeText                  PKEY = "7fe3aa27-2648-42f3-89b0-454e5cb150c3 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_RelatedSoundFile                 PKEY = "318a6b45-087f-4dc2-b8cc-05359551fc9e 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_Saturation                       PKEY = "49237325-a95a-4f67-b211-816b2d45d2e0 100"   // UInt32 -- VT_UI4
	PKEY_Photo_SaturationText                   PKEY = "61478c08-b600-4a84-bbe4-e99c45f0a072 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_Sharpness                        PKEY = "fc6976db-8349-4970-ae97-b3c5316a08f0 100"   // UInt32 -- VT_UI4
	PKEY_Photo_SharpnessText                    PKEY = "51ec3f47-dd50-421d-8769-334f50424b1e 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Photo_ShutterSpeed                     PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 37377" // Double -- VT_R8
	PKEY_Photo_ShutterSpeedDenominator          PKEY = "e13d8975-81c7-4948-ae3f-37cae11e8ff7 100"   // Int32 -- VT_I4
	PKEY_Photo_ShutterSpeedNumerator            PKEY = "16ea4042-d6f4-4bca-8349-7c78d30fb333 100"   // Int32 -- VT_I4
	PKEY_Photo_SubjectDistance                  PKEY = "14b81da1-0135-4d31-96d9-6cbfc9671a99 37382" // Double -- VT_R8
	PKEY_Photo_SubjectDistanceDenominator       PKEY = "0c840a88-b043-466d-9766-d4b26da3fa77 100"   // UInt32 -- VT_UI4
	PKEY_Photo_SubjectDistanceNumerator         PKEY = "8af4961c-f526-43e5-aa81-db768219178d 100"   // UInt32 -- VT_UI4
	PKEY_Photo_TagViewAggregate                 PKEY = "b812f15d-c2d8-4bbf-bacd-79744346113f 100"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)  Legacy code may treat this as VT_LPSTR.
	PKEY_Photo_TranscodedForSync                PKEY = "9a8ebb75-6458-4e82-bacb-35c0095b03bb 100"   // Boolean -- VT_BOOL
	PKEY_Photo_WhiteBalance                     PKEY = "ee3d3d8a-5381-4cfa-b13b-aaf66b5f4ec9 100"   // UInt32 -- VT_UI4
	PKEY_Photo_WhiteBalanceText                 PKEY = "6336b95e-c7a7-426d-86fd-7ae3d39c84b4 100"   // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// PropGroup properties

	PKEY_PropGroup_Advanced      PKEY = "900a403b-097b-4b95-8ae2-071fdaeeb118 100" // Null -- VT_NULL
	PKEY_PropGroup_Audio         PKEY = "2804d469-788f-48aa-8570-71b9c187e138 100" // Null -- VT_NULL
	PKEY_PropGroup_Calendar      PKEY = "9973d2b5-bfd8-438a-ba94-5349b293181a 100" // Null -- VT_NULL
	PKEY_PropGroup_Camera        PKEY = "de00de32-547e-4981-ad4b-542f2e9007d8 100" // Null -- VT_NULL
	PKEY_PropGroup_Contact       PKEY = "df975fd3-250a-4004-858f-34e29a3e37aa 100" // Null -- VT_NULL
	PKEY_PropGroup_Content       PKEY = "d0dab0ba-368a-4050-a882-6c010fd19a4f 100" // Null -- VT_NULL
	PKEY_PropGroup_Description   PKEY = "8969b275-9475-4e00-a887-ff93b8b41e44 100" // Null -- VT_NULL
	PKEY_PropGroup_FileSystem    PKEY = "e3a7d2c1-80fc-4b40-8f34-30ea111bdc2e 100" // Null -- VT_NULL
	PKEY_PropGroup_General       PKEY = "cc301630-b192-4c22-b372-9f4c6d338e07 100" // Null -- VT_NULL
	PKEY_PropGroup_GPS           PKEY = "f3713ada-90e3-4e11-aae5-fdc17685b9be 100" // Null -- VT_NULL
	PKEY_PropGroup_Image         PKEY = "e3690a87-0fa8-4a2a-9a9f-fce8827055ac 100" // Null -- VT_NULL
	PKEY_PropGroup_Media         PKEY = "61872cf7-6b5e-4b4b-ac2d-59da84459248 100" // Null -- VT_NULL
	PKEY_PropGroup_MediaAdvanced PKEY = "8859a284-de7e-4642-99ba-d431d044b1ec 100" // Null -- VT_NULL
	PKEY_PropGroup_Message       PKEY = "7fd7259d-16b4-4135-9f97-7c96ecd2fa9e 100" // Null -- VT_NULL
	PKEY_PropGroup_Music         PKEY = "68dd6094-7216-40f1-a029-43fe7127043f 100" // Null -- VT_NULL
	PKEY_PropGroup_Origin        PKEY = "2598d2fb-5569-4367-95df-5cd3a177e1a5 100" // Null -- VT_NULL
	PKEY_PropGroup_PhotoAdvanced PKEY = "0cb2bf5a-9ee7-4a86-8222-f01e07fdadaf 100" // Null -- VT_NULL
	PKEY_PropGroup_RecordedTV    PKEY = "e7b33238-6584-4170-a5c0-ac25efd9da56 100" // Null -- VT_NULL
	PKEY_PropGroup_Video         PKEY = "bebe0920-7671-4c54-a3eb-49fddfc191ee 100" // Null -- VT_NULL

	// PropList properties

	PKEY_InfoTipText                       PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 17"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_ConflictPrompt           PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 11"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_ContentViewModeForBrowse PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 13"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_ContentViewModeForSearch PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 14"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_ExtendedTileInfo         PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 9"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_FileOperationPrompt      PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 10"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_FullDetails              PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 2"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_InfoTip                  PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 4"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_NonPersonal              PKEY = "49d1091f-082e-493f-b23f-d2308aa9668c 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_PreviewDetails           PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 8"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_PreviewTitle             PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 6"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_QuickTip                 PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 5"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_TileInfo                 PKEY = "c9944a21-a406-48fe-8225-aec7e24c211b 3"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_PropList_XPDetailsPanel           PKEY = "f2275480-f782-4291-bd94-f13693513aec 0"   // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// RecordedTV properties

	PKEY_RecordedTV_ChannelNumber               PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 7"   // UInt32 -- VT_UI4
	PKEY_RecordedTV_Credits                     PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 4"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_DateContentExpires          PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 15"  // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_RecordedTV_EpisodeName                 PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 2"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_IsATSCContent               PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 16"  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsClosedCaptioningAvailable PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 12"  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsDTVContent                PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 17"  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsHDContent                 PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 18"  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsRepeatBroadcast           PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 13"  // Boolean -- VT_BOOL
	PKEY_RecordedTV_IsSAP                       PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 14"  // Boolean -- VT_BOOL
	PKEY_RecordedTV_NetworkAffiliation          PKEY = "2c53c813-fb63-4e22-a1ab-0b331ca1e273 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_OriginalBroadcastDate       PKEY = "4684fe97-8765-4842-9c13-f006447b178c 100" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_RecordedTV_ProgramDescription          PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 3"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_RecordingTime               PKEY = "a5477f61-7a82-4eca-9dde-98b69b2479b3 100" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_RecordedTV_StationCallSign             PKEY = "6d748de2-8d38-4cc3-ac60-f009b057c557 5"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_RecordedTV_StationName                 PKEY = "1b5439e7-eba1-4af8-bdd7-7af1d4549493 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Search properties

	PKEY_Search_AutoSummary                     PKEY = "560c36c0-503a-11cf-baa1-00004c752a9a 2"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_ContainerHash                   PKEY = "bceee283-35df-4d53-826a-f36a3eefc6be 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_Contents                        PKEY = "b725f130-47ef-101a-a5f1-02608c9eebac 19"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_EntryID                         PKEY = "49691c90-7e17-101a-a91c-08002b2ecda9 5"   // Int32 -- VT_I4
	PKEY_Search_ExtendedProperties              PKEY = "7b03b546-fa4f-4a52-a2fe-03d5311e5865 100" // Blob -- VT_BLOB
	PKEY_Search_GatherTime                      PKEY = "0b63e350-9ccc-11d0-bcdb-00805fccce04 8"   // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Search_HitCount                        PKEY = "49691c90-7e17-101a-a91c-08002b2ecda9 4"   // Int32 -- VT_I4
	PKEY_Search_IsClosedDirectory               PKEY = "0b63e343-9ccc-11d0-bcdb-00805fccce04 23"  // Boolean -- VT_BOOL
	PKEY_Search_IsFullyContained                PKEY = "0b63e343-9ccc-11d0-bcdb-00805fccce04 24"  // Boolean -- VT_BOOL
	PKEY_Search_QueryFocusedSummary             PKEY = "560c36c0-503a-11cf-baa1-00004c752a9a 3"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_QueryFocusedSummaryWithFallback PKEY = "560c36c0-503a-11cf-baa1-00004c752a9a 4"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_QueryPropertyHits               PKEY = "49691c90-7e17-101a-a91c-08002b2ecda9 21"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Search_Rank                            PKEY = "49691c90-7e17-101a-a91c-08002b2ecda9 3"   // Int32 -- VT_I4
	PKEY_Search_Store                           PKEY = "a06992b3-8caf-4ed7-a547-b259e32ac9fc 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_UrlToIndex                      PKEY = "0b63e343-9ccc-11d0-bcdb-00805fccce04 2"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Search_UrlToIndexWithModificationTime  PKEY = "0b63e343-9ccc-11d0-bcdb-00805fccce04 12"  // Multivalue Any -- VT_VECTOR | VT_NULL  (For variants: VT_ARRAY | VT_NULL)
	PKEY_Supplemental_Album                     PKEY = "0c73b141-39d6-4653-a683-cab291eaf95b 6"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Supplemental_AlbumID                   PKEY = "0c73b141-39d6-4653-a683-cab291eaf95b 2"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Supplemental_Location                  PKEY = "0c73b141-39d6-4653-a683-cab291eaf95b 5"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Supplemental_Person                    PKEY = "0c73b141-39d6-4653-a683-cab291eaf95b 7"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Supplemental_ResourceId                PKEY = "0c73b141-39d6-4653-a683-cab291eaf95b 3"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Supplemental_Tag                       PKEY = "0c73b141-39d6-4653-a683-cab291eaf95b 4"   // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)

	// Shell properties

	PKEY_DescriptionID                PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 2" // Buffer -- VT_VECTOR | VT_UI1  (For variants: VT_ARRAY | VT_UI1)
	PKEY_InternalName                 PKEY = "0cef7d53-fa64-11d1-a203-0000f81fedee 5" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_LibraryLocationsCount        PKEY = "908696c7-8f87-44f2-80ed-a8c1c6894575 2" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Link_TargetSFGAOFlagsStrings PKEY = "d6942081-d53b-443d-ad47-5e059d9cd27a 3" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Link_TargetUrl               PKEY = "5cbf2787-48cf-4208-b90e-ee5e5d420294 2" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_NamespaceCLSID               PKEY = "28636aa6-953d-11d2-b5d6-00c04fd918d0 6" // Guid -- VT_CLSID
	PKEY_Shell_SFGAOFlagsStrings      PKEY = "d6942081-d53b-443d-ad47-5e059d9cd27a 2" // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_StatusBarSelectedItemCount   PKEY = "26dc287c-6e3d-4bd3-b2b0-6a26ba2e346d 3" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_StatusBarViewItemCount       PKEY = "26dc287c-6e3d-4bd3-b2b0-6a26ba2e346d 2" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Software properties

	PKEY_AppUserModel_ExcludeFromShowInNewInstall    PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 8"  // Boolean -- VT_BOOL
	PKEY_AppUserModel_ID                             PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 5"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_AppUserModel_IsDestListSeparator            PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 6"  // Boolean -- VT_BOOL
	PKEY_AppUserModel_IsDualMode                     PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 11" // Boolean -- VT_BOOL
	PKEY_AppUserModel_PreventPinning                 PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 9"  // Boolean -- VT_BOOL
	PKEY_AppUserModel_RelaunchCommand                PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 2"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_AppUserModel_RelaunchDisplayNameResource    PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 4"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_AppUserModel_RelaunchIconResource           PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 3"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_AppUserModel_StartPinOption                 PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 12" // UInt32 -- VT_UI4
	PKEY_AppUserModel_ToastActivatorCLSID            PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 26" // Guid -- VT_CLSID
	PKEY_AppUserModel_VisualElementsManifestHintPath PKEY = "9f4c2855-9f79-4b39-a8d0-e1d42de1d5f3 31" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_EdgeGesture_DisableTouchWhenFullscreen      PKEY = "32ce38b2-2c9a-41b1-9bc5-b3784394aa44 2"  // Boolean -- VT_BOOL
	PKEY_Software_DateLastUsed                       PKEY = "841e4f90-ff59-4d16-8947-e81bbffab36d 16" // DateTime -- VT_FILETIME  (For variants: VT_DATE)
	PKEY_Software_ProductName                        PKEY = "0cef7d53-fa64-11d1-a203-0000f81fedee 7"  // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Sync properties

	PKEY_Sync_Comments               PKEY = "7bd5533e-af15-44db-b8c8-bd6624e1d032 13" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ConflictDescription    PKEY = "ce50c159-2fb8-41fd-be68-d3e042e274bc 4"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ConflictFirstLocation  PKEY = "ce50c159-2fb8-41fd-be68-d3e042e274bc 6"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ConflictSecondLocation PKEY = "ce50c159-2fb8-41fd-be68-d3e042e274bc 7"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_HandlerCollectionID    PKEY = "7bd5533e-af15-44db-b8c8-bd6624e1d032 2"  // Guid -- VT_CLSID
	PKEY_Sync_HandlerID              PKEY = "7bd5533e-af15-44db-b8c8-bd6624e1d032 3"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_HandlerName            PKEY = "ce50c159-2fb8-41fd-be68-d3e042e274bc 2"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_HandlerType            PKEY = "7bd5533e-af15-44db-b8c8-bd6624e1d032 8"  // UInt32 -- VT_UI4
	PKEY_Sync_HandlerTypeLabel       PKEY = "7bd5533e-af15-44db-b8c8-bd6624e1d032 9"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ItemID                 PKEY = "7bd5533e-af15-44db-b8c8-bd6624e1d032 6"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ItemName               PKEY = "ce50c159-2fb8-41fd-be68-d3e042e274bc 3"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Sync_ProgressPercentage     PKEY = "7bd5533e-af15-44db-b8c8-bd6624e1d032 23" // UInt32 -- VT_UI4
	PKEY_Sync_State                  PKEY = "7bd5533e-af15-44db-b8c8-bd6624e1d032 24" // UInt32 -- VT_UI4
	PKEY_Sync_Status                 PKEY = "7bd5533e-af15-44db-b8c8-bd6624e1d032 10" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Task properties

	PKEY_Task_BillingInformation PKEY = "d37d52c6-261c-4303-82b3-08b926ac6f12 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Task_CompletionStatus   PKEY = "084d8a0a-e6d5-40de-bf1f-c8820e7c877c 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Task_Owner              PKEY = "08c7cc5f-60f2-4494-ad75-55e3e0b5add0 100" // String -- VT_LPWSTR  (For variants: VT_BSTR)

	// Video properties

	PKEY_Video_Compression           PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 10"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Video_Director              PKEY = "64440492-4c8b-11d1-8b70-080036b11a03 20"  // Multivalue String -- VT_VECTOR | VT_LPWSTR  (For variants: VT_ARRAY | VT_BSTR)
	PKEY_Video_EncodingBitrate       PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 8"   // UInt32 -- VT_UI4
	PKEY_Video_FourCC                PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 44"  // UInt32 -- VT_UI4
	PKEY_Video_FrameHeight           PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 4"   // UInt32 -- VT_UI4
	PKEY_Video_FrameRate             PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 6"   // UInt32 -- VT_UI4
	PKEY_Video_FrameWidth            PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 3"   // UInt32 -- VT_UI4
	PKEY_Video_HorizontalAspectRatio PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 42"  // UInt32 -- VT_UI4
	PKEY_Video_IsSpherical           PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 100" // Boolean -- VT_BOOL
	PKEY_Video_IsStereo              PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 98"  // Boolean -- VT_BOOL
	PKEY_Video_Orientation           PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 99"  // UInt32 -- VT_UI4
	PKEY_Video_SampleSize            PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 9"   // UInt32 -- VT_UI4
	PKEY_Video_StreamName            PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 2"   // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Video_StreamNumber          PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 11"  // UInt16 -- VT_UI2
	PKEY_Video_TotalBitrate          PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 43"  // UInt32 -- VT_UI4
	PKEY_Video_TranscodedForSync     PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 46"  // Boolean -- VT_BOOL
	PKEY_Video_VerticalAspectRatio   PKEY = "64440491-4c8b-11d1-8b70-080036b11a03 45"  // UInt32 -- VT_UI4

	// Volume properties

	PKEY_Volume_FileSystem    PKEY = "9b174b35-40ff-11d2-a27e-00c04fc30871 4"  // String -- VT_LPWSTR  (For variants: VT_BSTR)
	PKEY_Volume_IsMappedDrive PKEY = "149c0b69-2c2d-48fc-808f-d318d78c4636 2"  // Boolean -- VT_BOOL
	PKEY_Volume_IsRoot        PKEY = "9b174b35-40ff-11d2-a27e-00c04fc30871 10" // Boolean -- VT_BOOL
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
