// Code generated by entc, DO NOT EDIT.

package run

import (
	"time"
)

const (
	// Label holds the string label denoting the run type in the database.
	Label = "run"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStartAt holds the string denoting the startat field in the database.
	FieldStartAt = "start_at"
	// FieldStartedAt holds the string denoting the startedat field in the database.
	FieldStartedAt = "started_at"
	// FieldEndedAt holds the string denoting the endedat field in the database.
	FieldEndedAt = "ended_at"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"

	// Table holds the table name of the run in the database.
	Table = "runs"
)

// Columns holds all SQL columns for run fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStartAt,
	FieldStartedAt,
	FieldEndedAt,
	FieldError,
}

var (
	// DefaultCreatedAt holds the default value on creation for the createdAt field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	UpdateDefaultUpdatedAt func() time.Time
)
