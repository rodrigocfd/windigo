//go:build windows

package cotasks

import (
	"github.com/rodrigocfd/windigo/co"
)

const (
	HRESULT_SCHED_S_TASK_READY             co.HRESULT = 0x0004_1300 // The task is ready to run at its next scheduled time.
	HRESULT_SCHED_S_TASK_RUNNING           co.HRESULT = 0x0004_1301 // The task is currently running.
	HRESULT_SCHED_S_TASK_DISABLED          co.HRESULT = 0x0004_1302 // The task will not run at the scheduled times because it has been disabled.
	HRESULT_SCHED_S_TASK_HAS_NOT_RUN       co.HRESULT = 0x0004_1303 // The task has not yet run.
	HRESULT_SCHED_S_TASK_NO_MORE_RUNS      co.HRESULT = 0x0004_1304 // There are no more runs scheduled for this task.
	HRESULT_SCHED_S_TASK_NOT_SCHEDULED     co.HRESULT = 0x0004_1305 // One or more of the properties that are needed to run this task on a schedule have not been set.
	HRESULT_SCHED_S_TASK_TERMINATED        co.HRESULT = 0x0004_1306 // The last run of the task was terminated by the user.
	HRESULT_SCHED_S_TASK_NO_VALID_TRIGGERS co.HRESULT = 0x0004_1307 // Either the task has no triggers or the existing triggers are disabled or not set.
	HRESULT_SCHED_S_EVENT_TRIGGER          co.HRESULT = 0x0004_1308 // Event triggers don't have set run times.

	HRESULT_SCHED_E_TRIGGER_NOT_FOUND           co.HRESULT = 0x8004_1309 // Trigger not found.
	HRESULT_SCHED_E_TASK_NOT_READY              co.HRESULT = 0x8004_130a // One or more of the properties that are needed to run this task have not been set.
	HRESULT_SCHED_E_TASK_NOT_RUNNING            co.HRESULT = 0x8004_130b // There is no running instance of the task.
	HRESULT_SCHED_E_SERVICE_NOT_INSTALLED       co.HRESULT = 0x8004_130c // The Task Scheduler Service is not installed on this computer.
	HRESULT_SCHED_E_CANNOT_OPEN_TASK            co.HRESULT = 0x8004_130d // The task object could not be opened.
	HRESULT_SCHED_E_INVALID_TASK                co.HRESULT = 0x8004_130e // The object is either an invalid task object or is not a task object.
	HRESULT_SCHED_E_ACCOUNT_INFORMATION_NOT_SET co.HRESULT = 0x8004_130f // No account information could be found in the Task Scheduler security database for the task indicated.
	HRESULT_SCHED_E_ACCOUNT_NAME_NOT_FOUND      co.HRESULT = 0x8004_1310 // Unable to establish existence of the account specified.
	HRESULT_SCHED_E_ACCOUNT_DBASE_CORRUPT       co.HRESULT = 0x8004_1311 // Corruption was detected in the Task Scheduler security database; the database has been reset.
	HRESULT_SCHED_E_NO_SECURITY_SERVICES        co.HRESULT = 0x8004_1312 // Task Scheduler security services are available only on Windows NT.
	HRESULT_SCHED_E_UNKNOWN_OBJECT_VERSION      co.HRESULT = 0x8004_1313 // The task object version is either unsupported or invalid.
	HRESULT_SCHED_E_UNSUPPORTED_ACCOUNT_OPTION  co.HRESULT = 0x8004_1314 // The task has been configured with an unsupported combination of account settings and run time options.
	HRESULT_SCHED_E_SERVICE_NOT_RUNNING         co.HRESULT = 0x8004_1315 // The Task Scheduler Service is not running.
	HRESULT_SCHED_E_UNEXPECTEDNODE              co.HRESULT = 0x8004_1316 // The task XML contains an unexpected node.
	HRESULT_SCHED_E_NAMESPACE                   co.HRESULT = 0x8004_1317 // The task XML contains an element or attribute from an unexpected namespace.
	HRESULT_SCHED_E_INVALIDVALUE                co.HRESULT = 0x8004_1318 // The task XML contains a value which is incorrectly formatted or out of range.
	HRESULT_SCHED_E_MISSINGNODE                 co.HRESULT = 0x8004_1319 // The task XML is missing a required element or attribute.
	HRESULT_SCHED_E_MALFORMEDXML                co.HRESULT = 0x8004_131a // The task XML is malformed.
	HRESULT_SCHED_S_SOME_TRIGGERS_FAILED        co.HRESULT = 0x0004_131b // The task is registered, but not all specified triggers will start the task, check task scheduler event log for detailed information.
	HRESULT_SCHED_S_BATCH_LOGON_PROBLEM         co.HRESULT = 0x0004_131c // The task is registered, but may fail to start. Batch logon privilege needs to be enabled for the task principal.
	HRESULT_SCHED_E_TOO_MANY_NODES              co.HRESULT = 0x8004_131d // The task XML contains too many nodes of the same type.
	HRESULT_SCHED_E_PAST_END_BOUNDARY           co.HRESULT = 0x8004_131e // The task cannot be started after the trigger's end boundary.
	HRESULT_SCHED_E_ALREADY_RUNNING             co.HRESULT = 0x8004_131f // An instance of this task is already running.
	HRESULT_SCHED_E_USER_NOT_LOGGED_ON          co.HRESULT = 0x8004_1320 // The task will not run because the user is not logged on.
	HRESULT_SCHED_E_INVALID_TASK_HASH           co.HRESULT = 0x8004_1321 // The task image is corrupt or has been tampered with.
	HRESULT_SCHED_E_SERVICE_NOT_AVAILABLE       co.HRESULT = 0x8004_1322 // The Task Scheduler service is not available.
	HRESULT_SCHED_E_SERVICE_TOO_BUSY            co.HRESULT = 0x8004_1323 // The Task Scheduler service is too busy to handle your request. Please try again later.
	HRESULT_SCHED_E_TASK_ATTEMPTED              co.HRESULT = 0x8004_1324 // The Task Scheduler service attempted to run the task, but the task did not run due to one of the constraints in the task definition.
	HRESULT_SCHED_S_TASK_QUEUED                 co.HRESULT = 0x0004_1325 // The Task Scheduler service has asked the task to run.
	HRESULT_SCHED_E_TASK_DISABLED               co.HRESULT = 0x8004_1326 // The task is disabled.
	HRESULT_SCHED_E_TASK_NOT_V1_COMPAT          co.HRESULT = 0x8004_1327 // The task has properties that are not compatible with previous versions of Windows.
	HRESULT_SCHED_E_START_ON_DEMAND             co.HRESULT = 0x8004_1328 // The task settings do not allow the task to start on demand.
	HRESULT_SCHED_E_TASK_NOT_UBPM_COMPAT        co.HRESULT = 0x8004_1329 // The combination of properties that task is using is not compatible with the scheduling engine.
	HRESULT_SCHED_E_DEPRECATED_FEATURE_USED     co.HRESULT = 0x8004_1330 // The task definition uses a deprecated feature.
)
