//go:build windows

package co

// IBindCtx [BindToHandler] bhid, represented as a string.
//
// [BindToHandler]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-bindtohandler
type BHID GUID

const (
	BHID_SFObject           BHID = "3981e224-f559-11d3-8e3a-00c04f6837d5"
	BHID_SFUIObject         BHID = "3981e225-f559-11d3-8e3a-00c04f6837d5"
	BHID_SFViewObject       BHID = "3981e226-f559-11d3-8e3a-00c04f6837d5"
	BHID_Storage            BHID = "3981e227-f559-11d3-8e3a-00c04f6837d5"
	BHID_Stream             BHID = "1cebb3ab-7c10-499a-a417-92ca16c4cb83"
	BHID_RandomAccessStream BHID = "f16fc93b-77ae-4cfe-bda7-a866eea6878d"
	BHID_LinkTargetItem     BHID = "3981e228-f559-11d3-8e3a-00c04f6837d5"
	BHID_StorageEnum        BHID = "4621a4e3-f0d6-4773-8a9c-46e77b174840"
	BHID_Transfer           BHID = "5d080304-fe2c-48fc-84ce-cf620b0f3c53"
	BHID_PropertyStore      BHID = "0384e1a4-1523-439c-a4c8-ab911052f586"
	BHID_ThumbnailHandler   BHID = "7b2e650a-8e20-4f4a-b09e-6597afc72fb0"
	BHID_EnumItems          BHID = "94f60519-2850-4924-aa5a-d15e84868039"
	BHID_DataObject         BHID = "b8c0bd9f-ed24-455c-83e6-d5390c4fe8c4"
	BHID_AssociationArray   BHID = "bea9ef17-82f1-4f60-9284-4f8db75c3be9"
	BHID_Filter             BHID = "38d08778-f557-4690-9ebf-ba54706ad8f7"
	BHID_EnumAssocHandlers  BHID = "b8ab0b9c-c2ec-4f7a-918d-314900e6280a"
	BHID_StorageItem        BHID = "404e2109-77d2-4699-a5a0-4fdf10db9837"
	BHID_FilePlaceholder    BHID = "8677dceb-aae0-4005-8d3d-547fa852f825"
)

// [BIND_FLAGS] enumeration.
//
// [BIND_FLAGS]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ne-objidl-bind_flags
type BIND uint32

const (
	BIND_MAYBOTHERUSER     BIND = 1
	BIND_JUSTTESTEXISTENCE BIND = 2
)

// [CoCreateInstance] dwClsContext.
//
// [CoCreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
type CLSCTX uint32

const (
	CLSCTX_INPROC_SERVER          CLSCTX = 0x1
	CLSCTX_INPROC_HANDLER         CLSCTX = 0x2
	CLSCTX_LOCAL_SERVER           CLSCTX = 0x4
	CLSCTX_INPROC_SERVER16        CLSCTX = 0x8
	CLSCTX_REMOTE_SERVER          CLSCTX = 0x10
	CLSCTX_INPROC_HANDLER16       CLSCTX = 0x20
	CLSCTX_NO_CODE_DOWNLOAD       CLSCTX = 0x400
	CLSCTX_NO_CUSTOM_MARSHAL      CLSCTX = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD   CLSCTX = 0x2000
	CLSCTX_NO_FAILURE_LOG         CLSCTX = 0x4000
	CLSCTX_DISABLE_AAA            CLSCTX = 0x8000
	CLSCTX_ENABLE_AAA             CLSCTX = 0x1_0000
	CLSCTX_FROM_DEFAULT_CONTEXT   CLSCTX = 0x2_0000
	CLSCTX_ACTIVATE_X86_SERVER    CLSCTX = 0x4_0000
	CLSCTX_ACTIVATE_32_BIT_SERVER        = CLSCTX_ACTIVATE_X86_SERVER
	CLSCTX_ACTIVATE_64_BIT_SERVER CLSCTX = 0x8_0000
	CLSCTX_ENABLE_CLOAKING        CLSCTX = 0x10_0000
	CLSCTX_APPCONTAINER           CLSCTX = 0x40_0000
	CLSCTX_ACTIVATE_AAA_AS_IU     CLSCTX = 0x80_0000
	CLSCTX_ACTIVATE_ARM32_SERVER  CLSCTX = 0x200_0000
	CLSCTX_PS_DLL                 CLSCTX = 0x8000_0000
	CLSCTX_ALL                           = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
	CLSCTX_SERVER                        = CLSCTX_INPROC_SERVER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
)

// A COM [class ID], represented as a string.
//
// [class ID]: https://learn.microsoft.com/en-us/windows/win32/com/clsid-key-hklm
type CLSID GUID

// [CoInitializeEx] dwCoInit.
//
// [CoInitializeEx]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
type COINIT uint32

const (
	COINIT_APARTMENTTHREADED COINIT = 0x2
	COINIT_MULTITHREADED     COINIT = 0x0
	COINIT_DISABLE_OLE1DDE   COINIT = 0x4
	COINIT_SPEED_OVER_MEMORY COINIT = 0x8
)

// [DROPEFFECT] constants.
//
// [DROPEFFECT]: https://learn.microsoft.com/en-us/windows/win32/com/dropeffect-constants
type DROPEFFECT uint32

const (
	DROPEFFECT_NONE   DROPEFFECT = 0
	DROPEFFECT_COPY   DROPEFFECT = 1
	DROPEFFECT_MOVE   DROPEFFECT = 2
	DROPEFFECT_LINK   DROPEFFECT = 4
	DROPEFFECT_SCROLL DROPEFFECT = 0x8000_0000
)

// [DVASPECT] enumeration.
//
// [DVASPECT]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ne-wtypes-dvaspect
type DVASPECT uint32

const (
	DVASPECT_CONTENT   DVASPECT = 1
	DVASPECT_THUMBNAIL DVASPECT = 2
	DVASPECT_ICON      DVASPECT = 4
	DVASPECT_DOCPRINT  DVASPECT = 8
)

// [EOLE_AUTHENTICATION_CAPABILITIES] enumeration.
//
// [EOLE_AUTHENTICATION_CAPABILITIES]: https://learn.microsoft.com/en-us/windows/win32/api/objidlbase/ne-objidlbase-eole_authentication_capabilities
type EOAC uint32

const (
	EOAC_NONE              EOAC = 0
	EOAC_MUTUAL_AUTH       EOAC = 0x1
	EOAC_STATIC_CLOAKING   EOAC = 0x20
	EOAC_DYNAMIC_CLOAKING  EOAC = 0x40
	EOAC_ANY_AUTHORITY     EOAC = 0x80
	EOAC_MAKE_FULLSIC      EOAC = 0x100
	EOAC_DEFAULT           EOAC = 0x800
	EOAC_SECURE_REFS       EOAC = 0x2
	EOAC_ACCESS_CONTROL    EOAC = 0x4
	EOAC_APPID             EOAC = 0x8
	EOAC_DYNAMIC           EOAC = 0x10
	EOAC_REQUIRE_FULLSIC   EOAC = 0x200
	EOAC_AUTO_IMPERSONATE  EOAC = 0x400
	EOAC_DISABLE_AAA       EOAC = 0x1000
	EOAC_NO_CUSTOM_MARSHAL EOAC = 0x2000
	EOAC_RESERVED1         EOAC = 0x4000
)

// [COAUTHINFO] dwCapabilities, with flags from [EOAC] and
// [RPC_C_QOS_CAPABILITIES].
//
// [COAUTHINFO]: https://learn.microsoft.com/en-us/windows/win32/api/wtypesbase/ns-wtypesbase-coauthinfo
type EOAC_QOS uint32

const (
	EOAC_QOS_NONE        = EOAC_QOS(EOAC_NONE)
	EOAC_QOS_MUTUAL_AUTH = EOAC_QOS(RPC_C_QOS_CAPABILITIES_MUTUAL_AUTH)
)

// [HRESULT] facility.
//
// [HRESULT]: https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref/0642cb2f-2075-4469-918c-4441e69c548a
type FACILITY uint32

const (
	FACILITY_NULL                                     FACILITY = 0
	FACILITY_RPC                                      FACILITY = 1
	FACILITY_DISPATCH                                 FACILITY = 2
	FACILITY_STORAGE                                  FACILITY = 3
	FACILITY_ITF                                      FACILITY = 4
	FACILITY_WIN32                                    FACILITY = 7
	FACILITY_WINDOWS                                  FACILITY = 8
	FACILITY_SSPI                                     FACILITY = 9
	FACILITY_SECURITY                                 FACILITY = 9
	FACILITY_CONTROL                                  FACILITY = 10
	FACILITY_CERT                                     FACILITY = 11
	FACILITY_INTERNET                                 FACILITY = 12
	FACILITY_MEDIASERVER                              FACILITY = 13
	FACILITY_MSMQ                                     FACILITY = 14
	FACILITY_SETUPAPI                                 FACILITY = 15
	FACILITY_SCARD                                    FACILITY = 16
	FACILITY_COMPLUS                                  FACILITY = 17
	FACILITY_AAF                                      FACILITY = 18
	FACILITY_URT                                      FACILITY = 19
	FACILITY_ACS                                      FACILITY = 20
	FACILITY_DPLAY                                    FACILITY = 21
	FACILITY_UMI                                      FACILITY = 22
	FACILITY_SXS                                      FACILITY = 23
	FACILITY_WINDOWS_CE                               FACILITY = 24
	FACILITY_HTTP                                     FACILITY = 25
	FACILITY_USERMODE_COMMONLOG                       FACILITY = 26
	FACILITY_WER                                      FACILITY = 27
	FACILITY_USERMODE_FILTER_MANAGER                  FACILITY = 31
	FACILITY_BACKGROUNDCOPY                           FACILITY = 32
	FACILITY_CONFIGURATION                            FACILITY = 33
	FACILITY_WIA                                      FACILITY = 33
	FACILITY_STATE_MANAGEMENT                         FACILITY = 34
	FACILITY_METADIRECTORY                            FACILITY = 35
	FACILITY_WINDOWSUPDATE                            FACILITY = 36
	FACILITY_DIRECTORYSERVICE                         FACILITY = 37
	FACILITY_GRAPHICS                                 FACILITY = 38
	FACILITY_SHELL                                    FACILITY = 39
	FACILITY_NAP                                      FACILITY = 39
	FACILITY_TPM_SERVICES                             FACILITY = 40
	FACILITY_TPM_SOFTWARE                             FACILITY = 41
	FACILITY_UI                                       FACILITY = 42
	FACILITY_XAML                                     FACILITY = 43
	FACILITY_ACTION_QUEUE                             FACILITY = 44
	FACILITY_PLA                                      FACILITY = 48
	FACILITY_WINDOWS_SETUP                            FACILITY = 48
	FACILITY_FVE                                      FACILITY = 49
	FACILITY_FWP                                      FACILITY = 50
	FACILITY_WINRM                                    FACILITY = 51
	FACILITY_NDIS                                     FACILITY = 52
	FACILITY_USERMODE_HYPERVISOR                      FACILITY = 53
	FACILITY_CMI                                      FACILITY = 54
	FACILITY_USERMODE_VIRTUALIZATION                  FACILITY = 55
	FACILITY_USERMODE_VOLMGR                          FACILITY = 56
	FACILITY_BCD                                      FACILITY = 57
	FACILITY_USERMODE_VHD                             FACILITY = 58
	FACILITY_USERMODE_HNS                             FACILITY = 59
	FACILITY_SDIAG                                    FACILITY = 60
	FACILITY_WEBSERVICES                              FACILITY = 61
	FACILITY_WINPE                                    FACILITY = 61
	FACILITY_WPN                                      FACILITY = 62
	FACILITY_WINDOWS_STORE                            FACILITY = 63
	FACILITY_INPUT                                    FACILITY = 64
	FACILITY_EAP                                      FACILITY = 66
	FACILITY_WINDOWS_DEFENDER                         FACILITY = 80
	FACILITY_OPC                                      FACILITY = 81
	FACILITY_XPS                                      FACILITY = 82
	FACILITY_MBN                                      FACILITY = 84
	FACILITY_POWERSHELL                               FACILITY = 84
	FACILITY_RAS                                      FACILITY = 83
	FACILITY_P2P_INT                                  FACILITY = 98
	FACILITY_P2P                                      FACILITY = 99
	FACILITY_DAF                                      FACILITY = 100
	FACILITY_BLUETOOTH_ATT                            FACILITY = 101
	FACILITY_AUDIO                                    FACILITY = 102
	FACILITY_STATEREPOSITORY                          FACILITY = 103
	FACILITY_VISUALCPP                                FACILITY = 109
	FACILITY_SCRIPT                                   FACILITY = 112
	FACILITY_PARSE                                    FACILITY = 113
	FACILITY_BLB                                      FACILITY = 120
	FACILITY_BLB_CLI                                  FACILITY = 121
	FACILITY_WSBAPP                                   FACILITY = 122
	FACILITY_BLBUI                                    FACILITY = 128
	FACILITY_USN                                      FACILITY = 129
	FACILITY_USERMODE_VOLSNAP                         FACILITY = 130
	FACILITY_TIERING                                  FACILITY = 131
	FACILITY_WSB_ONLINE                               FACILITY = 133
	FACILITY_ONLINE_ID                                FACILITY = 134
	FACILITY_DEVICE_UPDATE_AGENT                      FACILITY = 135
	FACILITY_DRVSERVICING                             FACILITY = 136
	FACILITY_DLS                                      FACILITY = 153
	FACILITY_DELIVERY_OPTIMIZATION                    FACILITY = 208
	FACILITY_USERMODE_SPACES                          FACILITY = 231
	FACILITY_USER_MODE_SECURITY_CORE                  FACILITY = 232
	FACILITY_USERMODE_LICENSING                       FACILITY = 234
	FACILITY_SOS                                      FACILITY = 160
	FACILITY_DEBUGGERS                                FACILITY = 176
	FACILITY_SPP                                      FACILITY = 256
	FACILITY_RESTORE                                  FACILITY = 256
	FACILITY_DMSERVER                                 FACILITY = 256
	FACILITY_DEPLOYMENT_SERVICES_SERVER               FACILITY = 257
	FACILITY_DEPLOYMENT_SERVICES_IMAGING              FACILITY = 258
	FACILITY_DEPLOYMENT_SERVICES_MANAGEMENT           FACILITY = 259
	FACILITY_DEPLOYMENT_SERVICES_UTIL                 FACILITY = 260
	FACILITY_DEPLOYMENT_SERVICES_BINLSVC              FACILITY = 261
	FACILITY_DEPLOYMENT_SERVICES_PXE                  FACILITY = 263
	FACILITY_DEPLOYMENT_SERVICES_TFTP                 FACILITY = 264
	FACILITY_DEPLOYMENT_SERVICES_TRANSPORT_MANAGEMENT FACILITY = 272
	FACILITY_DEPLOYMENT_SERVICES_DRIVER_PROVISIONING  FACILITY = 278
	FACILITY_DEPLOYMENT_SERVICES_MULTICAST_SERVER     FACILITY = 289
	FACILITY_DEPLOYMENT_SERVICES_MULTICAST_CLIENT     FACILITY = 290
	FACILITY_DEPLOYMENT_SERVICES_CONTENT_PROVIDER     FACILITY = 293
	FACILITY_LINGUISTIC_SERVICES                      FACILITY = 305
	FACILITY_AUDIOSTREAMING                           FACILITY = 1094
	FACILITY_ACCELERATOR                              FACILITY = 1536
	FACILITY_WMAAECMA                                 FACILITY = 1996
	FACILITY_DIRECTMUSIC                              FACILITY = 2168
	FACILITY_DIRECT3D10                               FACILITY = 2169
	FACILITY_DXGI                                     FACILITY = 2170
	FACILITY_DXGI_DDI                                 FACILITY = 2171
	FACILITY_DIRECT3D11                               FACILITY = 2172
	FACILITY_DIRECT3D11_DEBUG                         FACILITY = 2173
	FACILITY_DIRECT3D12                               FACILITY = 2174
	FACILITY_DIRECT3D12_DEBUG                         FACILITY = 2175
	FACILITY_LEAP                                     FACILITY = 2184
	FACILITY_AUDCLNT                                  FACILITY = 2185
	FACILITY_WINCODEC_DWRITE_DWM                      FACILITY = 2200
	FACILITY_WINML                                    FACILITY = 2192
	FACILITY_DIRECT2D                                 FACILITY = 2201
	FACILITY_DEFRAG                                   FACILITY = 2304
	FACILITY_USERMODE_SDBUS                           FACILITY = 2305
	FACILITY_JSCRIPT                                  FACILITY = 2306
	FACILITY_PIDGENX                                  FACILITY = 2561
	FACILITY_EAS                                      FACILITY = 85
	FACILITY_WEB                                      FACILITY = 885
	FACILITY_WEB_SOCKET                               FACILITY = 886
	FACILITY_MOBILE                                   FACILITY = 1793
	FACILITY_SQLITE                                   FACILITY = 1967
	FACILITY_UTC                                      FACILITY = 1989
	FACILITY_WEP                                      FACILITY = 2049
	FACILITY_SYNCENGINE                               FACILITY = 2050
	FACILITY_XBOX                                     FACILITY = 2339
	FACILITY_GAME                                     FACILITY = 2340
	FACILITY_PIX                                      FACILITY = 2748
)

// A [GUID] struct, represented as a string.
//
// [GUID]: https://learn.microsoft.com/en-us/windows/win32/api/guiddef/ns-guiddef-guid
type GUID string

const (
	GUID_NULL GUID = "00000000-0000-0000-0000-000000000000"
)

// A COM [interface ID], represented as a string.
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
type IID GUID

// Ole IID identifier.
const (
	IID_IBindCtx          IID = "0000000e-0000-0000-c000-000000000046"
	IID_IDataObject       IID = "0000010e-0000-0000-c000-000000000046"
	IID_IDropTarget       IID = "00000122-0000-0000-c000-000000000046"
	IID_IEnumString       IID = "00000101-0000-0000-c000-000000000046"
	IID_ISequentialStream IID = "0c733a30-2a1c-11ce-ade5-00aa0044773d"
	IID_IStream           IID = "0000000c-0000-0000-c000-000000000046"
	IID_IUnknown          IID = "00000000-0000-0000-c000-000000000046"
	IID_NULL                  = IID(GUID_NULL)
)

// [LOCKTYPE] enumeration.
//
// [LOCKTYPE]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ne-objidl-locktype
type LOCKTYPE uint32

const (
	LOCKTYPE_WRITE     LOCKTYPE = 1
	LOCKTYPE_EXCLUSIVE LOCKTYPE = 2
	LOCKTYPE_ONLYONCE  LOCKTYPE = 4
)

// [PICTUREATTRIBUTES] enumeration.
//
// [PICTUREATTRIBUTES]: https://learn.microsoft.com/en-us/windows/win32/api/ocidl/ne-ocidl-pictureattributes
type PICATTR uint32

const (
	PICATTR_SCALABLE    PICATTR = 0x01
	PICATTR_TRANSPARENT PICATTR = 0x02
)

// [PICTYPE] constants.
//
// [PICTYPE]: https://learn.microsoft.com/en-us/windows/win32/com/pictype-constants
type PICTYPE int16

const (
	PICTYPE_UNINITIALIZED PICTYPE = -1
	PICTYPE_NONE          PICTYPE = 0
	PICTYPE_BITMAP        PICTYPE = 1
	PICTYPE_METAFILE      PICTYPE = 2
	PICTYPE_ICON          PICTYPE = 3
	PICTYPE_ENHMETAFILE   PICTYPE = 4
)

// Authentication service [constants].
//
// [constants]: https://learn.microsoft.com/en-us/windows/win32/com/com-authentication-service-constants
type RPC_C_AUTHN uint32

const (
	RPC_C_AUTHN_NONE          RPC_C_AUTHN = 0
	RPC_C_AUTHN_DCE_PRIVATE   RPC_C_AUTHN = 1
	RPC_C_AUTHN_DCE_PUBLIC    RPC_C_AUTHN = 2
	RPC_C_AUTHN_DEC_PUBLIC    RPC_C_AUTHN = 4
	RPC_C_AUTHN_GSS_NEGOTIATE RPC_C_AUTHN = 9
	RPC_C_AUTHN_WINNT         RPC_C_AUTHN = 10
	RPC_C_AUTHN_GSS_SCHANNEL  RPC_C_AUTHN = 14
	RPC_C_AUTHN_GSS_KERBEROS  RPC_C_AUTHN = 16
	RPC_C_AUTHN_DPA           RPC_C_AUTHN = 17
	RPC_C_AUTHN_MSN           RPC_C_AUTHN = 18
	RPC_C_AUTHN_KERNEL        RPC_C_AUTHN = 20
	RPC_C_AUTHN_DIGEST        RPC_C_AUTHN = 21
	RPC_C_AUTHN_NEGO_EXTENDER RPC_C_AUTHN = 30
	RPC_C_AUTHN_PKU2U         RPC_C_AUTHN = 31
	RPC_C_AUTHN_MQ            RPC_C_AUTHN = 100
	RPC_C_AUTHN_DEFAULT       RPC_C_AUTHN = 0xffff_ffff
)

// Authentication level [constants].
//
// [constants]: https://learn.microsoft.com/en-us/windows/win32/com/com-authentication-level-constants
type RPC_C_AUTHN_LEVEL uint32

const (
	RPC_C_AUTHN_LEVEL_DEFAULT       RPC_C_AUTHN_LEVEL = 0
	RPC_C_AUTHN_LEVEL_NONE          RPC_C_AUTHN_LEVEL = 1
	RPC_C_AUTHN_LEVEL_CONNECT       RPC_C_AUTHN_LEVEL = 2
	RPC_C_AUTHN_LEVEL_CALL          RPC_C_AUTHN_LEVEL = 3
	RPC_C_AUTHN_LEVEL_PKT           RPC_C_AUTHN_LEVEL = 4
	RPC_C_AUTHN_LEVEL_PKT_INTEGRITY RPC_C_AUTHN_LEVEL = 5
	RPC_C_AUTHN_LEVEL_PKT_PRIVACY   RPC_C_AUTHN_LEVEL = 6
)

// Authorization [constants].
//
// [constants]: https://learn.microsoft.com/en-us/windows/win32/com/com-authorization-constants
type RPC_C_AUTHZ uint32

const (
	RPC_C_AUTHZ_NONE    RPC_C_AUTHZ = 0
	RPC_C_AUTHZ_NAME    RPC_C_AUTHZ = 1
	RPC_C_AUTHZ_DCE     RPC_C_AUTHZ = 2
	RPC_C_AUTHZ_DEFAULT RPC_C_AUTHZ = 0xffff_ffff
)

// Impersonation level [constants].
//
// [constants]: https://learn.microsoft.com/en-us/windows/win32/com/com-impersonation-level-constants
type RPC_C_IMP_LEVEL uint32

const (
	RPC_C_IMP_LEVEL_DEFAULT     RPC_C_IMP_LEVEL = 0
	RPC_C_IMP_LEVEL_ANONYMOUS   RPC_C_IMP_LEVEL = 1
	RPC_C_IMP_LEVEL_IDENTIFY    RPC_C_IMP_LEVEL = 2
	RPC_C_IMP_LEVEL_IMPERSONATE RPC_C_IMP_LEVEL = 3
	RPC_C_IMP_LEVEL_DELEGATE    RPC_C_IMP_LEVEL = 4
)

// Quality of service [capabilities].
//
// [capabilities]: https://learn.microsoft.com/en-us/windows/win32/rpc/quality-of-service
type RPC_C_QOS_CAPABILITIES uint32

const (
	RPC_C_QOS_CAPABILITIES_DEFAULT                     RPC_C_QOS_CAPABILITIES = 0x0
	RPC_C_QOS_CAPABILITIES_MUTUAL_AUTH                 RPC_C_QOS_CAPABILITIES = 0x1
	RPC_C_QOS_CAPABILITIES_MAKE_FULLSIC                RPC_C_QOS_CAPABILITIES = 0x2
	RPC_C_QOS_CAPABILITIES_ANY_AUTHORITY               RPC_C_QOS_CAPABILITIES = 0x4
	RPC_C_QOS_CAPABILITIES_IGNORE_DELEGATE_FAILURE     RPC_C_QOS_CAPABILITIES = 0x8
	RPC_C_QOS_CAPABILITIES_LOCAL_MA_HINT               RPC_C_QOS_CAPABILITIES = 0x10
	RPC_C_QOS_CAPABILITIES_SCHANNEL_FULL_AUTH_IDENTITY RPC_C_QOS_CAPABILITIES = 0x20
)

// [COAUTHIDENTITY] flags.
//
// [COAUTHIDENTITY]: https://learn.microsoft.com/en-us/windows/win32/api/wtypesbase/ns-wtypesbase-coauthidentity
type SEC_WINNT_AUTH_IDENTITY uint32

const (
	SEC_WINNT_AUTH_IDENTITY_ANSI    SEC_WINNT_AUTH_IDENTITY = 0x1
	SEC_WINNT_AUTH_IDENTITY_UNICODE SEC_WINNT_AUTH_IDENTITY = 0x2
)

// [HRESULT] severity.
//
// [HRESULT]: https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref/0642cb2f-2075-4469-918c-4441e69c548a
type SEVERITY uint8

const (
	SEVERITY_SUCCESS SEVERITY = 0
	SEVERITY_ERROR   SEVERITY = 1
)

// [BIND_OPTS3] and [IShellLink.Resolve] flags.
//
// [BIND_OPTS3]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ns-objidl-bind_opts3-r1
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

// [STATFLAG] enumeration.
//
// [STATFLAG]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ne-wtypes-statflag
type STATFLAG uint32

const (
	STATFLAG_DEFAULT STATFLAG = 0
	STATFLAG_NONAME  STATFLAG = 1
	STATFLAG_NOOPEN  STATFLAG = 2
)

// [STGC] enumeration.
//
// [STGC]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ne-wtypes-stgc
type STGC uint32

const (
	STGC_DEFAULT                            STGC = 0
	STGC_OVERWRITE                          STGC = 1
	STGC_ONLYIFCURRENT                      STGC = 2
	STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE STGC = 4
	STGC_CONSOLIDATE                        STGC = 8
)

// [STGM] constants.
//
// [STGM]: https://learn.microsoft.com/en-us/windows/win32/stg/stgm-constants
type STGM uint32

const (
	STGM_READ             STGM = 0x0000_0000
	STGM_WRITE            STGM = 0x0000_0001
	STGM_READWRITE        STGM = 0x0000_0002
	STGM_SHARE_DENY_NONE  STGM = 0x0000_0040
	STGM_SHARE_DENY_READ  STGM = 0x0000_0030
	STGM_SHARE_DENY_WRITE STGM = 0x0000_0020
	STGM_SHARE_EXCLUSIVE  STGM = 0x0000_0010
	STGM_PRIORITY         STGM = 0x0004_0000
	STGM_CREATE           STGM = 0x0000_1000
	STGM_CONVERT          STGM = 0x0002_0000
	STGM_FAILIFTHERE      STGM = 0x0000_0000
	STGM_DIRECT           STGM = 0x0000_0000
	STGM_TRANSACTED       STGM = 0x0001_0000
	STGM_NOSCRATCH        STGM = 0x0010_0000
	STGM_NOSNAPSHOT       STGM = 0x0020_0000
	STGM_SIMPLE           STGM = 0x0800_0000
	STGM_DIRECT_SWMR      STGM = 0x0040_0000
	STGM_DELETEONRELEASE  STGM = 0x0400_0000
)

// [STGTY] enumeration.
//
// [STGTY]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ne-objidl-stgty
type STGTY uint32

const (
	STGTY_STORAGE   STGC = 1
	STGTY_STREAM    STGC = 2
	STGTY_LOCKBYTES STGC = 3
	STGTY_PROPERTY  STGC = 4
)

// [STREAM_SEEK] enumeration.
//
// [STREAM_SEEK]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ne-objidl-stream_seek
type STREAM_SEEK uint32

const (
	STREAM_SEEK_SET STREAM_SEEK = 0
	STREAM_SEEK_CUR STREAM_SEEK = 1
	STREAM_SEEK_END STREAM_SEEK = 2
)

// [TYMED] enumeration.
//
// [TYMED]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/ne-objidl-tymed
type TYMED uint32

const (
	TYMED_HGLOBAL  TYMED = 1
	TYMED_FILE     TYMED = 2
	TYMED_ISTREAM  TYMED = 4
	TYMED_ISTORAGE TYMED = 8
	TYMED_GDI      TYMED = 16
	TYMED_MFPICT   TYMED = 32
	TYMED_ENHMF    TYMED = 64
	TYMED_NULL     TYMED = 0
)
