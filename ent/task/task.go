// Code generated by entc, DO NOT EDIT.

package task

import (
	"time"

	"github.com/kcarretto/paragon/ent/schema"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQueueTime holds the string denoting the queuetime vertex property in the database.
	FieldQueueTime = "queue_time"
	// FieldLastChangedTime holds the string denoting the lastchangedtime vertex property in the database.
	FieldLastChangedTime = "last_changed_time"
	// FieldClaimTime holds the string denoting the claimtime vertex property in the database.
	FieldClaimTime = "claim_time"
	// FieldExecStartTime holds the string denoting the execstarttime vertex property in the database.
	FieldExecStartTime = "exec_start_time"
	// FieldExecStopTime holds the string denoting the execstoptime vertex property in the database.
	FieldExecStopTime = "exec_stop_time"
	// FieldContent holds the string denoting the content vertex property in the database.
	FieldContent = "content"
	// FieldOutput holds the string denoting the output vertex property in the database.
	FieldOutput = "output"
	// FieldError holds the string denoting the error vertex property in the database.
	FieldError = "error"
	// FieldSessionID holds the string denoting the sessionid vertex property in the database.
	FieldSessionID = "session_id"

	// Table holds the table name of the task in the database.
	Table = "tasks"
	// TagsTable is the table the holds the tags relation/edge. The primary key declared below.
	TagsTable = "task_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// JobTable is the table the holds the job relation/edge.
	JobTable = "tasks"
	// JobInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobInverseTable = "jobs"
	// JobColumn is the table column denoting the job relation/edge.
	JobColumn = "job_id"
	// TargetTable is the table the holds the target relation/edge.
	TargetTable = "tasks"
	// TargetInverseTable is the table name for the Target entity.
	// It exists in this package in order to avoid circular dependency with the "target" package.
	TargetInverseTable = "targets"
	// TargetColumn is the table column denoting the target relation/edge.
	TargetColumn = "target_id"
)

// Columns holds all SQL columns are task fields.
var Columns = []string{
	FieldID,
	FieldQueueTime,
	FieldLastChangedTime,
	FieldClaimTime,
	FieldExecStartTime,
	FieldExecStopTime,
	FieldContent,
	FieldOutput,
	FieldError,
	FieldSessionID,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"task_id", "tag_id"}
)

var (
	fields = schema.Task{}.Fields()

	// descQueueTime is the schema descriptor for QueueTime field.
	descQueueTime = fields[0].Descriptor()
	// DefaultQueueTime holds the default value on creation for the QueueTime field.
	DefaultQueueTime = descQueueTime.Default.(func() time.Time)

	// descContent is the schema descriptor for Content field.
	descContent = fields[5].Descriptor()
	// ContentValidator is a validator for the "Content" field. It is called by the builders before save.
	ContentValidator = descContent.Validators[0].(func(string) error)

	// descSessionID is the schema descriptor for SessionID field.
	descSessionID = fields[8].Descriptor()
	// SessionIDValidator is a validator for the "SessionID" field. It is called by the builders before save.
	SessionIDValidator = descSessionID.Validators[0].(func(string) error)
)
