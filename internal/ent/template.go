// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/facebook/ent/dialect/sql"
)

// Template is the model entity for the Template schema.
type Template struct {
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
	// The values are being populated by the TemplateQuery when eager-loading is set.
	Edges             TemplateEdges `json:"edges"`
	admin_templates   *string
	project_templates *string
	run_template      *string
}

// TemplateEdges holds the relations/edges for other nodes in the graph.
type TemplateEdges struct {
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
func (e TemplateEdges) StepsOrErr() ([]*Step, error) {
	if e.loadedTypes[0] {
		return e.Steps, nil
	}
	return nil, &NotLoadedError{edge: "steps"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TemplateEdges) ProjectOrErr() (*Project, error) {
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
func (e TemplateEdges) CreatorOrErr() (*Admin, error) {
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
func (e TemplateEdges) RunOrErr() (*Run, error) {
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
func (*Template) scanValues() []interface{} {
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
func (*Template) fkValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // admin_templates
		&sql.NullString{}, // project_templates
		&sql.NullString{}, // run_template
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Template fields.
func (t *Template) assignValues(values ...interface{}) error {
	if m, n := len(values), len(template.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		t.ID = value.String
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[0])
	} else if value.Valid {
		t.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[1])
	} else if value.Valid {
		t.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		t.Name = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field selectionType", values[3])
	} else if value.Valid {
		t.SelectionType = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field participantCount", values[4])
	} else if value.Valid {
		t.ParticipantCount = int(value.Int64)
	}
	if value, ok := values[5].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field internalCriteria", values[5])
	} else if value != nil {
		t.InternalCriteria = *value
	}
	if value, ok := values[6].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field mturkCriteria", values[6])
	} else if value != nil {
		t.MturkCriteria = *value
	}
	if value, ok := values[7].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field adult", values[7])
	} else if value.Valid {
		t.Adult = value.Bool
	}
	values = values[8:]
	if len(values) == len(template.ForeignKeys) {
		if value, ok := values[0].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field admin_templates", values[0])
		} else if value.Valid {
			t.admin_templates = new(string)
			*t.admin_templates = value.String
		}
		if value, ok := values[1].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field project_templates", values[1])
		} else if value.Valid {
			t.project_templates = new(string)
			*t.project_templates = value.String
		}
		if value, ok := values[2].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field run_template", values[2])
		} else if value.Valid {
			t.run_template = new(string)
			*t.run_template = value.String
		}
	}
	return nil
}

// QuerySteps queries the steps edge of the Template.
func (t *Template) QuerySteps() *StepQuery {
	return (&TemplateClient{config: t.config}).QuerySteps(t)
}

// QueryProject queries the project edge of the Template.
func (t *Template) QueryProject() *ProjectQuery {
	return (&TemplateClient{config: t.config}).QueryProject(t)
}

// QueryCreator queries the creator edge of the Template.
func (t *Template) QueryCreator() *AdminQuery {
	return (&TemplateClient{config: t.config}).QueryCreator(t)
}

// QueryRun queries the run edge of the Template.
func (t *Template) QueryRun() *RunQuery {
	return (&TemplateClient{config: t.config}).QueryRun(t)
}

// Update returns a builder for updating this Template.
// Note that, you need to call Template.Unwrap() before calling this method, if this Template
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Template) Update() *TemplateUpdateOne {
	return (&TemplateClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Template) Unwrap() *Template {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Template is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Template) String() string {
	var builder strings.Builder
	builder.WriteString("Template(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", createdAt=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", selectionType=")
	builder.WriteString(t.SelectionType)
	builder.WriteString(", participantCount=")
	builder.WriteString(fmt.Sprintf("%v", t.ParticipantCount))
	builder.WriteString(", internalCriteria=")
	builder.WriteString(fmt.Sprintf("%v", t.InternalCriteria))
	builder.WriteString(", mturkCriteria=")
	builder.WriteString(fmt.Sprintf("%v", t.MturkCriteria))
	builder.WriteString(", adult=")
	builder.WriteString(fmt.Sprintf("%v", t.Adult))
	builder.WriteByte(')')
	return builder.String()
}

// Templates is a parsable slice of Template.
type Templates []*Template

func (t Templates) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}