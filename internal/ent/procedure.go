// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/ent/procedure"
	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/facebook/ent/dialect/sql"
)

// Procedure is the model entity for the Procedure schema.
type Procedure struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// SelectionType holds the value of the "selectionType" field.
	SelectionType string `json:"selectionType,omitempty"`
	// ParticipantCount holds the value of the "participantCount" field.
	ParticipantCount int `json:"participantCount,omitempty"`
	// InternalCriteria holds the value of the "internalCriteria" field.
	InternalCriteria []byte `json:"internalCriteria,omitempty"`
	// MturkCriteria holds the value of the "mturkCriteria" field.
	MturkCriteria []byte `json:"mturkCriteria,omitempty"`
	// Adult holds the value of the "adult" field.
	Adult bool `json:"adult,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProcedureQuery when eager-loading is set.
	Edges              ProcedureEdges `json:"edges"`
	admin_procedures   *string
	project_procedures *string
	run_procedure      *string
}

// ProcedureEdges holds the relations/edges for other nodes in the graph.
type ProcedureEdges struct {
	// Steps holds the value of the steps edge.
	Steps []*Step
	// Project holds the value of the project edge.
	Project *Project
	// Creator holds the value of the creator edge.
	Creator *Admin
	// Run holds the value of the run edge.
	Run *Run
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// StepsOrErr returns the Steps value or an error if the edge
// was not loaded in eager-loading.
func (e ProcedureEdges) StepsOrErr() ([]*Step, error) {
	if e.loadedTypes[0] {
		return e.Steps, nil
	}
	return nil, &NotLoadedError{edge: "steps"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProcedureEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[1] {
		if e.Project == nil {
			// The edge project was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProcedureEdges) CreatorOrErr() (*Admin, error) {
	if e.loadedTypes[2] {
		if e.Creator == nil {
			// The edge creator was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// RunOrErr returns the Run value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProcedureEdges) RunOrErr() (*Run, error) {
	if e.loadedTypes[3] {
		if e.Run == nil {
			// The edge run was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: run.Label}
		}
		return e.Run, nil
	}
	return nil, &NotLoadedError{edge: "run"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Procedure) scanValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // id
		&sql.NullTime{},   // createdAt
		&sql.NullTime{},   // updatedAt
		&sql.NullString{}, // name
		&sql.NullString{}, // selectionType
		&sql.NullInt64{},  // participantCount
		&[]byte{},         // internalCriteria
		&[]byte{},         // mturkCriteria
		&sql.NullBool{},   // adult
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Procedure) fkValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // admin_procedures
		&sql.NullString{}, // project_procedures
		&sql.NullString{}, // run_procedure
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Procedure fields.
func (pr *Procedure) assignValues(values ...interface{}) error {
	if m, n := len(values), len(procedure.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		pr.ID = value.String
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[0])
	} else if value.Valid {
		pr.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[1])
	} else if value.Valid {
		pr.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		pr.Name = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field selectionType", values[3])
	} else if value.Valid {
		pr.SelectionType = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field participantCount", values[4])
	} else if value.Valid {
		pr.ParticipantCount = int(value.Int64)
	}
	if value, ok := values[5].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field internalCriteria", values[5])
	} else if value != nil {
		pr.InternalCriteria = *value
	}
	if value, ok := values[6].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field mturkCriteria", values[6])
	} else if value != nil {
		pr.MturkCriteria = *value
	}
	if value, ok := values[7].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field adult", values[7])
	} else if value.Valid {
		pr.Adult = value.Bool
	}
	values = values[8:]
	if len(values) == len(procedure.ForeignKeys) {
		if value, ok := values[0].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field admin_procedures", values[0])
		} else if value.Valid {
			pr.admin_procedures = new(string)
			*pr.admin_procedures = value.String
		}
		if value, ok := values[1].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field project_procedures", values[1])
		} else if value.Valid {
			pr.project_procedures = new(string)
			*pr.project_procedures = value.String
		}
		if value, ok := values[2].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field run_procedure", values[2])
		} else if value.Valid {
			pr.run_procedure = new(string)
			*pr.run_procedure = value.String
		}
	}
	return nil
}

// QuerySteps queries the steps edge of the Procedure.
func (pr *Procedure) QuerySteps() *StepQuery {
	return (&ProcedureClient{config: pr.config}).QuerySteps(pr)
}

// QueryProject queries the project edge of the Procedure.
func (pr *Procedure) QueryProject() *ProjectQuery {
	return (&ProcedureClient{config: pr.config}).QueryProject(pr)
}

// QueryCreator queries the creator edge of the Procedure.
func (pr *Procedure) QueryCreator() *AdminQuery {
	return (&ProcedureClient{config: pr.config}).QueryCreator(pr)
}

// QueryRun queries the run edge of the Procedure.
func (pr *Procedure) QueryRun() *RunQuery {
	return (&ProcedureClient{config: pr.config}).QueryRun(pr)
}

// Update returns a builder for updating this Procedure.
// Note that, you need to call Procedure.Unwrap() before calling this method, if this Procedure
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Procedure) Update() *ProcedureUpdateOne {
	return (&ProcedureClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Procedure) Unwrap() *Procedure {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Procedure is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Procedure) String() string {
	var builder strings.Builder
	builder.WriteString("Procedure(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", createdAt=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", selectionType=")
	builder.WriteString(pr.SelectionType)
	builder.WriteString(", participantCount=")
	builder.WriteString(fmt.Sprintf("%v", pr.ParticipantCount))
	builder.WriteString(", internalCriteria=")
	builder.WriteString(fmt.Sprintf("%v", pr.InternalCriteria))
	builder.WriteString(", mturkCriteria=")
	builder.WriteString(fmt.Sprintf("%v", pr.MturkCriteria))
	builder.WriteString(", adult=")
	builder.WriteString(fmt.Sprintf("%v", pr.Adult))
	builder.WriteByte(')')
	return builder.String()
}

// Procedures is a parsable slice of Procedure.
type Procedures []*Procedure

func (pr Procedures) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}