//go:build windows

package co

// Taskschd IID identifier.
var (
	IID_IAction                   = IID(GUID{0xbae54997, 0x48b1, 0x4cbe, [8]byte{0x99, 0x65, 0xd6, 0xbe, 0x26, 0x3e, 0xbe, 0xa4}})
	IID_IActionCollection         = IID(GUID{0x02820e19, 0x7b98, 0x4ed2, [8]byte{0xb2, 0xe8, 0xfd, 0xcc, 0xce, 0xff, 0x61, 0x9b}})
	IID_IBootTrigger              = IID(GUID{0x2a9c35da, 0xd357, 0x41f4, [8]byte{0xbb, 0xc1, 0x20, 0x7a, 0xc1, 0xb1, 0xf3, 0xcb}})
	IID_IComHandlerAction         = IID(GUID{0x6d2fd252, 0x75c5, 0x4f66, [8]byte{0x90, 0xba, 0x2a, 0x7d, 0x8c, 0xc3, 0x03, 0x9f}})
	IID_IDailyTrigger             = IID(GUID{0x126c5cd8, 0xb288, 0x41d5, [8]byte{0x8d, 0xbf, 0xe4, 0x91, 0x44, 0x6a, 0xdc, 0x5c}})
	IID_IEmailAction              = IID(GUID{0x10f62c64, 0x7e16, 0x4314, [8]byte{0xa0, 0xc2, 0x0c, 0x36, 0x83, 0xf9, 0x9d, 0x40}})
	IID_IEventTrigger             = IID(GUID{0xd45b0167, 0x9653, 0x4eef, [8]byte{0xb9, 0x4f, 0x07, 0x32, 0xca, 0x7a, 0xf2, 0x51}})
	IID_IExecAction               = IID(GUID{0x4c3d624d, 0xfd6b, 0x49a3, [8]byte{0xb9, 0xb7, 0x09, 0xcb, 0x3c, 0xd3, 0xf0, 0x47}})
	IID_ILogonTrigger             = IID(GUID{0x72dade38, 0xfae4, 0x4b3e, [8]byte{0xba, 0xf4, 0x5d, 0x00, 0x9a, 0xf0, 0x2b, 0x1c}})
	IID_IPrincipal                = IID(GUID{0xd98d51e5, 0xc9b4, 0x496a, [8]byte{0xa9, 0xc1, 0x18, 0x98, 0x02, 0x61, 0xcf, 0x0f}})
	IID_IRegisteredTask           = IID(GUID{0x9c86f320, 0xdee3, 0x4dd1, [8]byte{0xb9, 0x72, 0xa3, 0x03, 0xf2, 0x6b, 0x06, 0x1e}})
	IID_IRegistrationInfo         = IID(GUID{0x416d8b73, 0xcb41, 0x4ea1, [8]byte{0x80, 0x5c, 0x9b, 0xe9, 0xa5, 0xac, 0x4a, 0x74}})
	IID_ITaskDefinition           = IID(GUID{0xf5bc8fc5, 0x536d, 0x4f77, [8]byte{0xb8, 0x52, 0xfb, 0xc1, 0x35, 0x6f, 0xde, 0xb6}})
	IID_ITaskFolder               = IID(GUID{0x8cfac062, 0xa080, 0x4c15, [8]byte{0x9a, 0x88, 0xaa, 0x7c, 0x2a, 0xf8, 0x0d, 0xfc}})
	IID_ITaskNamedValueCollection = IID(GUID{0xb4ef826b, 0x63c3, 0x46e4, [8]byte{0xa5, 0x04, 0xef, 0x69, 0xe4, 0xf7, 0xea, 0x4d}})
	IID_ITaskNamedValuePair       = IID(GUID{0x39038068, 0x2b46, 0x4afd, [8]byte{0x86, 0x62, 0x7b, 0xb6, 0xf8, 0x68, 0xd2, 0x21}})
	IID_ITaskService              = IID(GUID{0x2faba4c7, 0x4da9, 0x4013, [8]byte{0x96, 0x97, 0x20, 0xcc, 0x3f, 0xd4, 0x0f, 0x85}})
	IID_ITaskSettings             = IID(GUID{0x8fd4711d, 0x2d02, 0x4c8c, [8]byte{0x87, 0xe3, 0xef, 0xf6, 0x99, 0xde, 0x12, 0x7e}})
	IID_ITrigger                  = IID(GUID{0x09941815, 0xea89, 0x4b5b, [8]byte{0x89, 0xe0, 0x2a, 0x77, 0x38, 0x01, 0xfa, 0xc3}})
	IID_ITriggerCollection        = IID(GUID{0x85df5081, 0x1b24, 0x4f32, [8]byte{0x87, 0x8a, 0xd9, 0xd1, 0x4d, 0xf4, 0xcb, 0x77}})
)

// [TASK_COMPATIBILITY] enumeration.
//
// [TASK_COMPATIBILITY]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/ne-taskschd-task_compatibility
type TASK_COMPATIBILITY uint32

const (
	TASK_COMPATIBILITY_AT   TASK_COMPATIBILITY = 0
	TASK_COMPATIBILITY_V1   TASK_COMPATIBILITY = 1
	TASK_COMPATIBILITY_V2   TASK_COMPATIBILITY = 2
	TASK_COMPATIBILITY_V2_1 TASK_COMPATIBILITY = 3
	TASK_COMPATIBILITY_V2_2 TASK_COMPATIBILITY = 4
	TASK_COMPATIBILITY_V2_3 TASK_COMPATIBILITY = 5
	TASK_COMPATIBILITY_V2_4 TASK_COMPATIBILITY = 6
)

// [TASK_ACTION] enumeration.
//
// [TASK_ACTION]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/ne-taskschd-task_action_type
type TASK_ACTION uint32

const (
	TASK_ACTION_EXEC         TASK_ACTION = 0
	TASK_ACTION_COM_HANDLER  TASK_ACTION = 5
	TASK_ACTION_SEND_EMAIL   TASK_ACTION = 6
	TASK_ACTION_SHOW_MESSAGE TASK_ACTION = 7
)

// [TASK_INSTANCES] enumeration.
//
// [TASK_INSTANCES]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/ne-taskschd-task_instances_policy
type TASK_INSTANCES uint32

const (
	TASK_INSTANCES_PARALLEL      TASK_INSTANCES = 0
	TASK_INSTANCES_QUEUE         TASK_INSTANCES = 1
	TASK_INSTANCES_IGNORE_NEW    TASK_INSTANCES = 2
	TASK_INSTANCES_STOP_EXISTING TASK_INSTANCES = 3
)

// [TASK_LOGON_TYPE] enumeration.
//
// [TASK_LOGON_TYPE]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/ne-taskschd-task_logon_type
type TASK_LOGON uint32

const (
	TASK_LOGON_NONE                          TASK_LOGON = 0
	TASK_LOGON_PASSWORD                      TASK_LOGON = 1
	TASK_LOGON_S4U                           TASK_LOGON = 2
	TASK_LOGON_INTERACTIVE_TOKEN             TASK_LOGON = 3
	TASK_LOGON_GROUP                         TASK_LOGON = 4
	TASK_LOGON_SERVICE_ACCOUNT               TASK_LOGON = 5
	TASK_LOGON_INTERACTIVE_TOKEN_OR_PASSWORD TASK_LOGON = 6
)

// [TASK_RUNLEVEL_TYPE] enumeration.
//
// [TASK_RUNLEVEL_TYPE]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/ne-taskschd-task_runlevel_type
type TASK_RUNLEVEL uint32

const (
	TASK_RUNLEVEL_LUA     TASK_RUNLEVEL = 0
	TASK_RUNLEVEL_HIGHEST TASK_RUNLEVEL = 1
)

// [TASK_STATE] enumeration.
//
// [TASK_STATE]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/ne-taskschd-task_state
type TASK_STATE uint32

const (
	TASK_STATE_UNKNOWN  TASK_STATE = 0
	TASK_STATE_DISABLED TASK_STATE = 1
	TASK_STATE_QUEUED   TASK_STATE = 2
	TASK_STATE_READY    TASK_STATE = 3
	TASK_STATE_RUNNING  TASK_STATE = 4
)

// [TASK_TRIGGER2] enumeration.
//
// [TASK_TRIGGER2]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/ne-taskschd-task_trigger_type2
type TASK_TRIGGER2 uint32

const (
	TASK_TRIGGER2_EVENT                TASK_TRIGGER2 = 0
	TASK_TRIGGER2_TIME                 TASK_TRIGGER2 = 1
	TASK_TRIGGER2_DAILY                TASK_TRIGGER2 = 2
	TASK_TRIGGER2_WEEKLY               TASK_TRIGGER2 = 3
	TASK_TRIGGER2_MONTHLY              TASK_TRIGGER2 = 4
	TASK_TRIGGER2_MONTHLYDOW           TASK_TRIGGER2 = 5
	TASK_TRIGGER2_IDLE                 TASK_TRIGGER2 = 6
	TASK_TRIGGER2_REGISTRATION         TASK_TRIGGER2 = 7
	TASK_TRIGGER2_BOOT                 TASK_TRIGGER2 = 8
	TASK_TRIGGER2_LOGON                TASK_TRIGGER2 = 9
	TASK_TRIGGER2_SESSION_STATE_CHANGE TASK_TRIGGER2 = 11
	TASK_TRIGGER2_CUSTOM_TRIGGER_01    TASK_TRIGGER2 = 12
)
