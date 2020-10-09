// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/participant"
	"github.com/empiricaly/recruitment/internal/ent/providerid"
	"github.com/facebook/ent/dialect/sql"
)

// ProviderID is the model entity for the ProviderID schema.
type ProviderID struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// MturkWorkerID holds the value of the "mturkWorkerID" field.
	MturkWorkerID string `json:"mturkWorkerID,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProviderIDQuery when eager-loading is set.
	Edges                    ProviderIDEdges `json:"edges"`
	participant_provider_ids *string
}

// ProviderIDEdges holds the relations/edges for other nodes in the graph.
type ProviderIDEdges struct {
	// Particpant holds the value of the particpant edge.
	Particpant *Participant
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParticpantOrErr returns the Particpant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProviderIDEdges) ParticpantOrErr() (*Participant, error) {
	if e.loadedTypes[0] {
		if e.Particpant == nil {
			// The edge particpant was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: participant.Label}
		}
		return e.Particpant, nil
	}
	return nil, &NotLoadedError{edge: "particpant"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProviderID) scanValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullString{}, // mturkWorkerID
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*ProviderID) fkValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // participant_provider_ids
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProviderID fields.
func (pi *ProviderID) assignValues(values ...interface{}) error {
	if m, n := len(values), len(providerid.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		pi.ID = value.String
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		pi.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[1])
	} else if value.Valid {
		pi.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field mturkWorkerID", values[2])
	} else if value.Valid {
		pi.MturkWorkerID = value.String
	}
	values = values[3:]
	if len(values) == len(providerid.ForeignKeys) {
		if value, ok := values[0].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field participant_provider_ids", values[0])
		} else if value.Valid {
			pi.participant_provider_ids = new(string)
			*pi.participant_provider_ids = value.String
		}
	}
	return nil
}

// QueryParticpant queries the particpant edge of the ProviderID.
func (pi *ProviderID) QueryParticpant() *ParticipantQuery {
	return (&ProviderIDClient{config: pi.config}).QueryParticpant(pi)
}

// Update returns a builder for updating this ProviderID.
// Note that, you need to call ProviderID.Unwrap() before calling this method, if this ProviderID
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *ProviderID) Update() *ProviderIDUpdateOne {
	return (&ProviderIDClient{config: pi.config}).UpdateOne(pi)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pi *ProviderID) Unwrap() *ProviderID {
	tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProviderID is not a transactional entity")
	}
	pi.config.driver = tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *ProviderID) String() string {
	var builder strings.Builder
	builder.WriteString("ProviderID(")
	builder.WriteString(fmt.Sprintf("id=%v", pi.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", mturkWorkerID=")
	builder.WriteString(pi.MturkWorkerID)
	builder.WriteByte(')')
	return builder.String()
}

// ProviderIDs is a parsable slice of ProviderID.
type ProviderIDs []*ProviderID

func (pi ProviderIDs) config(cfg config) {
	for _i := range pi {
		pi[_i].config = cfg
	}
}