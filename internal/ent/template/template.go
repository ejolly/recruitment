// Code generated by entc, DO NOT EDIT.

package template

import (
	"time"
)

const (
	// Label holds the string label denoting the template type in the database.
	Label = "template"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSelectionType holds the string denoting the selectiontype field in the database.
	FieldSelectionType = "selection_type"
	// FieldParticipantCount holds the string denoting the participantcount field in the database.
	FieldParticipantCount = "participant_count"
	// FieldInternalCriteria holds the string denoting the internalcriteria field in the database.
	FieldInternalCriteria = "internal_criteria"
	// FieldMturkCriteria holds the string denoting the mturkcriteria field in the database.
	FieldMturkCriteria = "mturk_criteria"
	// FieldAdult holds the string denoting the adult field in the database.
	FieldAdult = "adult"

	// EdgeSteps holds the string denoting the steps edge name in mutations.
	EdgeSteps = "steps"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeRun holds the string denoting the run edge name in mutations.
	EdgeRun = "run"

	// Table holds the table name of the template in the database.
	Table = "templates"
	// StepsTable is the table the holds the steps relation/edge.
	StepsTable = "steps"
	// StepsInverseTable is the table name for the Step entity.
	// It exists in this package in order to avoid circular dependency with the "step" package.
	StepsInverseTable = "steps"
	// StepsColumn is the table column denoting the steps relation/edge.
	StepsColumn = "template_steps"
	// ProjectTable is the table the holds the project relation/edge.
	ProjectTable = "templates"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_templates"
	// CreatorTable is the table the holds the creator relation/edge.
	CreatorTable = "templates"
	// CreatorInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	CreatorInverseTable = "admins"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "admin_templates"
	// RunTable is the table the holds the run relation/edge.
	RunTable = "templates"
	// RunInverseTable is the table name for the Run entity.
	// It exists in this package in order to avoid circular dependency with the "run" package.
	RunInverseTable = "runs"
	// RunColumn is the table column denoting the run relation/edge.
	RunColumn = "run_template"
)

// Columns holds all SQL columns for template fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldSelectionType,
	FieldParticipantCount,
	FieldInternalCriteria,
	FieldMturkCriteria,
	FieldAdult,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Template type.
var ForeignKeys = []string{
	"admin_templates",
	"project_templates",
	"run_template",
}

var (
	// DefaultCreatedAt holds the default value on creation for the createdAt field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultParticipantCount holds the default value on creation for the participantCount field.
	DefaultParticipantCount int
	// ParticipantCountValidator is a validator for the "participantCount" field. It is called by the builders before save.
	ParticipantCountValidator func(int) error
	// DefaultAdult holds the default value on creation for the adult field.
	DefaultAdult bool
)