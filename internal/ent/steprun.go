// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/facebook/ent/dialect/sql"
)

// StepRun is the model entity for the StepRun schema.
type StepRun struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status holds the value of the "status" field.
	Status steprun.Status `json:"status,omitempty"`
	// StartedAt holds the value of the "startedAt" field.
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// EndedAt holds the value of the "endedAt" field.
	EndedAt *time.Time `json:"endedAt,omitempty"`
	// ParticipantsCount holds the value of the "participantsCount" field.
	ParticipantsCount int `json:"participantsCount,omitempty"`
	// HitID holds the value of the "hitID" field.
	HitID *string `json:"hitID,omitempty"`
	// UrlToken holds the value of the "urlToken" field.
	UrlToken string `json:"urlToken,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StepRunQuery when eager-loading is set.
	Edges     StepRunEdges `json:"edges"`
	run_steps *string
}

// StepRunEdges holds the relations/edges for other nodes in the graph.
type StepRunEdges struct {
	// CreatedParticipants holds the value of the createdParticipants edge.
	CreatedParticipants []*Participant
	// Participants holds the value of the participants edge.
	Participants []*Participant
	// Participations holds the value of the participations edge.
	Participations []*Participation
	// Step holds the value of the step edge.
	Step *Step
	// Run holds the value of the run edge.
	Run *Run
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// CreatedParticipantsOrErr returns the CreatedParticipants value or an error if the edge
// was not loaded in eager-loading.
func (e StepRunEdges) CreatedParticipantsOrErr() ([]*Participant, error) {
	if e.loadedTypes[0] {
		return e.CreatedParticipants, nil
	}
	return nil, &NotLoadedError{edge: "createdParticipants"}
}

// ParticipantsOrErr returns the Participants value or an error if the edge
// was not loaded in eager-loading.
func (e StepRunEdges) ParticipantsOrErr() ([]*Participant, error) {
	if e.loadedTypes[1] {
		return e.Participants, nil
	}
	return nil, &NotLoadedError{edge: "participants"}
}

// ParticipationsOrErr returns the Participations value or an error if the edge
// was not loaded in eager-loading.
func (e StepRunEdges) ParticipationsOrErr() ([]*Participation, error) {
	if e.loadedTypes[2] {
		return e.Participations, nil
	}
	return nil, &NotLoadedError{edge: "participations"}
}

// StepOrErr returns the Step value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StepRunEdges) StepOrErr() (*Step, error) {
	if e.loadedTypes[3] {
		if e.Step == nil {
			// The edge step was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: step.Label}
		}
		return e.Step, nil
	}
	return nil, &NotLoadedError{edge: "step"}
}

// RunOrErr returns the Run value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StepRunEdges) RunOrErr() (*Run, error) {
	if e.loadedTypes[4] {
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
func (*StepRun) scanValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullString{}, // status
		&sql.NullTime{},   // startedAt
		&sql.NullTime{},   // endedAt
		&sql.NullInt64{},  // participantsCount
		&sql.NullString{}, // hitID
		&sql.NullString{}, // urlToken
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*StepRun) fkValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // run_steps
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StepRun fields.
func (sr *StepRun) assignValues(values ...interface{}) error {
	if m, n := len(values), len(steprun.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		sr.ID = value.String
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		sr.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[1])
	} else if value.Valid {
		sr.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[2])
	} else if value.Valid {
		sr.Status = steprun.Status(value.String)
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field startedAt", values[3])
	} else if value.Valid {
		sr.StartedAt = new(time.Time)
		*sr.StartedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field endedAt", values[4])
	} else if value.Valid {
		sr.EndedAt = new(time.Time)
		*sr.EndedAt = value.Time
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field participantsCount", values[5])
	} else if value.Valid {
		sr.ParticipantsCount = int(value.Int64)
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field hitID", values[6])
	} else if value.Valid {
		sr.HitID = new(string)
		*sr.HitID = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field urlToken", values[7])
	} else if value.Valid {
		sr.UrlToken = value.String
	}
	values = values[8:]
	if len(values) == len(steprun.ForeignKeys) {
		if value, ok := values[0].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field run_steps", values[0])
		} else if value.Valid {
			sr.run_steps = new(string)
			*sr.run_steps = value.String
		}
	}
	return nil
}

// QueryCreatedParticipants queries the createdParticipants edge of the StepRun.
func (sr *StepRun) QueryCreatedParticipants() *ParticipantQuery {
	return (&StepRunClient{config: sr.config}).QueryCreatedParticipants(sr)
}

// QueryParticipants queries the participants edge of the StepRun.
func (sr *StepRun) QueryParticipants() *ParticipantQuery {
	return (&StepRunClient{config: sr.config}).QueryParticipants(sr)
}

// QueryParticipations queries the participations edge of the StepRun.
func (sr *StepRun) QueryParticipations() *ParticipationQuery {
	return (&StepRunClient{config: sr.config}).QueryParticipations(sr)
}

// QueryStep queries the step edge of the StepRun.
func (sr *StepRun) QueryStep() *StepQuery {
	return (&StepRunClient{config: sr.config}).QueryStep(sr)
}

// QueryRun queries the run edge of the StepRun.
func (sr *StepRun) QueryRun() *RunQuery {
	return (&StepRunClient{config: sr.config}).QueryRun(sr)
}

// Update returns a builder for updating this StepRun.
// Note that, you need to call StepRun.Unwrap() before calling this method, if this StepRun
// was returned from a transaction, and the transaction was committed or rolled back.
func (sr *StepRun) Update() *StepRunUpdateOne {
	return (&StepRunClient{config: sr.config}).UpdateOne(sr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (sr *StepRun) Unwrap() *StepRun {
	tx, ok := sr.config.driver.(*txDriver)
	if !ok {
		panic("ent: StepRun is not a transactional entity")
	}
	sr.config.driver = tx.drv
	return sr
}

// String implements the fmt.Stringer.
func (sr *StepRun) String() string {
	var builder strings.Builder
	builder.WriteString("StepRun(")
	builder.WriteString(fmt.Sprintf("id=%v", sr.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(sr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(sr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", sr.Status))
	if v := sr.StartedAt; v != nil {
		builder.WriteString(", startedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := sr.EndedAt; v != nil {
		builder.WriteString(", endedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", participantsCount=")
	builder.WriteString(fmt.Sprintf("%v", sr.ParticipantsCount))
	if v := sr.HitID; v != nil {
		builder.WriteString(", hitID=")
		builder.WriteString(*v)
	}
	builder.WriteString(", urlToken=")
	builder.WriteString(sr.UrlToken)
	builder.WriteByte(')')
	return builder.String()
}

// StepRuns is a parsable slice of StepRun.
type StepRuns []*StepRun

func (sr StepRuns) config(cfg config) {
	for _i := range sr {
		sr[_i].config = cfg
	}
}
