// Code generated by entc, DO NOT EDIT.

package project

import (
	"time"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldProjectID holds the string denoting the projectid field in the database.
	FieldProjectID = "project_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeRuns holds the string denoting the runs edge name in mutations.
	EdgeRuns = "runs"
	// EdgeProcedures holds the string denoting the procedures edge name in mutations.
	EdgeProcedures = "procedures"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the project in the database.
	Table = "projects"
	// RunsTable is the table the holds the runs relation/edge.
	RunsTable = "runs"
	// RunsInverseTable is the table name for the Run entity.
	// It exists in this package in order to avoid circular dependency with the "run" package.
	RunsInverseTable = "runs"
	// RunsColumn is the table column denoting the runs relation/edge.
	RunsColumn = "project_runs"
	// ProceduresTable is the table the holds the procedures relation/edge.
	ProceduresTable = "procedures"
	// ProceduresInverseTable is the table name for the Procedure entity.
	// It exists in this package in order to avoid circular dependency with the "procedure" package.
	ProceduresInverseTable = "procedures"
	// ProceduresColumn is the table column denoting the procedures relation/edge.
	ProceduresColumn = "project_procedures"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "projects"
	// OwnerInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	OwnerInverseTable = "admins"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "admin_projects"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldProjectID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Project type.
var ForeignKeys = []string{
	"admin_projects",
}

var (
	// DefaultCreatedAt holds the default value on creation for the createdAt field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	UpdateDefaultUpdatedAt func() time.Time
)
