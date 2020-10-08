// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:        "admins",
		Columns:     AdminsColumns,
		PrimaryKey:  []*schema.Column{AdminsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProceduresColumns holds the columns for the "procedures" table.
	ProceduresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "selection_type", Type: field.TypeString},
		{Name: "participant_count", Type: field.TypeInt},
		{Name: "internal_criteria", Type: field.TypeBytes},
		{Name: "mturk_criteria", Type: field.TypeBytes},
		{Name: "adult", Type: field.TypeBool},
		{Name: "admin_procedures", Type: field.TypeString, Nullable: true},
		{Name: "project_procedures", Type: field.TypeString, Nullable: true},
		{Name: "run_procedure", Type: field.TypeString, Unique: true, Nullable: true},
	}
	// ProceduresTable holds the schema information for the "procedures" table.
	ProceduresTable = &schema.Table{
		Name:       "procedures",
		Columns:    ProceduresColumns,
		PrimaryKey: []*schema.Column{ProceduresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "procedures_admins_procedures",
				Columns: []*schema.Column{ProceduresColumns[9]},

				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "procedures_projects_procedures",
				Columns: []*schema.Column{ProceduresColumns[10]},

				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "procedures_runs_procedure",
				Columns: []*schema.Column{ProceduresColumns[11]},

				RefColumns: []*schema.Column{RunsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "project_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "admin_projects", Type: field.TypeString, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "projects_admins_projects",
				Columns: []*schema.Column{ProjectsColumns[5]},

				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RunsColumns holds the columns for the "runs" table.
	RunsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"CREATED", "RUNNING", "PAUSED", "DONE", "TERMINATED", "FAILED"}},
		{Name: "start_at", Type: field.TypeTime, Nullable: true},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "ended_at", Type: field.TypeTime, Nullable: true},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "project_runs", Type: field.TypeString, Nullable: true},
	}
	// RunsTable holds the schema information for the "runs" table.
	RunsTable = &schema.Table{
		Name:       "runs",
		Columns:    RunsColumns,
		PrimaryKey: []*schema.Column{RunsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "runs_projects_runs",
				Columns: []*schema.Column{RunsColumns[9]},

				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StepsColumns holds the columns for the "steps" table.
	StepsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"MTURK_HIT", "MTURK_MESSAGE", "PARTICIPANT_FILTER"}},
		{Name: "index", Type: field.TypeInt},
		{Name: "duration", Type: field.TypeInt},
		{Name: "msg_args", Type: field.TypeBytes, Nullable: true},
		{Name: "hit_args", Type: field.TypeBytes, Nullable: true},
		{Name: "filter_args", Type: field.TypeBytes, Nullable: true},
		{Name: "procedure_steps", Type: field.TypeString, Nullable: true},
	}
	// StepsTable holds the schema information for the "steps" table.
	StepsTable = &schema.Table{
		Name:       "steps",
		Columns:    StepsColumns,
		PrimaryKey: []*schema.Column{StepsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "steps_procedures_steps",
				Columns: []*schema.Column{StepsColumns[9]},

				RefColumns: []*schema.Column{ProceduresColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StepRunsColumns holds the columns for the "step_runs" table.
	StepRunsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "ended_at", Type: field.TypeTime},
		{Name: "participants_count", Type: field.TypeInt},
		{Name: "run_steps", Type: field.TypeString, Nullable: true},
	}
	// StepRunsTable holds the schema information for the "step_runs" table.
	StepRunsTable = &schema.Table{
		Name:       "step_runs",
		Columns:    StepRunsColumns,
		PrimaryKey: []*schema.Column{StepRunsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "step_runs_runs_steps",
				Columns: []*schema.Column{StepRunsColumns[6]},

				RefColumns: []*schema.Column{RunsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		ProceduresTable,
		ProjectsTable,
		RunsTable,
		StepsTable,
		StepRunsTable,
	}
)

func init() {
	ProceduresTable.ForeignKeys[0].RefTable = AdminsTable
	ProceduresTable.ForeignKeys[1].RefTable = ProjectsTable
	ProceduresTable.ForeignKeys[2].RefTable = RunsTable
	ProjectsTable.ForeignKeys[0].RefTable = AdminsTable
	RunsTable.ForeignKeys[0].RefTable = ProjectsTable
	StepsTable.ForeignKeys[0].RefTable = ProceduresTable
	StepRunsTable.ForeignKeys[0].RefTable = RunsTable
}