//go:build windows

package co

// Taskschd IID identifier.
const (
	IID_IAction                   IID = "bae54997-48b1-4cbe-9965-d6be263ebea4"
	IID_IActionCollection         IID = "02820e19-7b98-4ed2-b2e8-fdccceff619b"
	IID_IBootTrigger              IID = "2a9c35da-d357-41f4-bbc1-207ac1b1f3cb"
	IID_IComHandlerAction         IID = "6d2fd252-75c5-4f66-90ba-2a7d8cc3039f"
	IID_IDailyTrigger             IID = "126c5cd8-b288-41d5-8dbf-e491446adc5c"
	IID_IEmailAction              IID = "10f62c64-7e16-4314-a0c2-0c3683f99d40"
	IID_IEventTrigger             IID = "d45b0167-9653-4eef-b94f-0732ca7af251"
	IID_IExecAction               IID = "4c3d624d-fd6b-49a3-b9b7-09cb3cd3f047"
	IID_ILogonTrigger             IID = "72dade38-fae4-4b3e-baf4-5d009af02b1c"
	IID_IPrincipal                IID = "d98d51e5-c9b4-496a-a9c1-18980261cf0f"
	IID_IRegisteredTask           IID = "9c86f320-dee3-4dd1-b972-a303f26b061e"
	IID_IRegistrationInfo         IID = "416d8b73-cb41-4ea1-805c-9be9a5ac4a74"
	IID_ITaskDefinition           IID = "f5bc8fc5-536d-4f77-b852-fbc1356fdeb6"
	IID_ITaskFolder               IID = "8cfac062-a080-4c15-9a88-aa7c2af80dfc"
	IID_ITaskNamedValueCollection IID = "b4ef826b-63c3-46e4-a504-ef69e4f7ea4d"
	IID_ITaskNamedValuePair       IID = "39038068-2b46-4afd-8662-7bb6f868d221"
	IID_ITaskService              IID = "2faba4c7-4da9-4013-9697-20cc3fd40f85"
	IID_ITaskSettings             IID = "8fd4711d-2d02-4c8c-87e3-eff699de127e"
	IID_ITrigger                  IID = "09941815-ea89-4b5b-89e0-2a773801fac3"
	IID_ITriggerCollection        IID = "85df5081-1b24-4f32-878a-d9d14df4cb77"
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
