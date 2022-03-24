// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gqlgen-ent/ent/artcile"
	"gqlgen-ent/ent/user"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Artcile is the model entity for the Artcile schema.
type Artcile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArtcileQuery when eager-loading is set.
	Edges   ArtcileEdges `json:"edges"`
	user_id *int
}

// ArtcileEdges holds the relations/edges for other nodes in the graph.
type ArtcileEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtcileEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Artcile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case artcile.FieldID:
			values[i] = new(sql.NullInt64)
		case artcile.FieldTitle, artcile.FieldDescription:
			values[i] = new(sql.NullString)
		case artcile.ForeignKeys[0]: // user_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Artcile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Artcile fields.
func (a *Artcile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artcile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case artcile.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case artcile.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case artcile.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_id", value)
			} else if value.Valid {
				a.user_id = new(int)
				*a.user_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Artcile entity.
func (a *Artcile) QueryUser() *UserQuery {
	return (&ArtcileClient{config: a.config}).QueryUser(a)
}

// Update returns a builder for updating this Artcile.
// Note that you need to call Artcile.Unwrap() before calling this method if this Artcile
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Artcile) Update() *ArtcileUpdateOne {
	return (&ArtcileClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Artcile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Artcile) Unwrap() *Artcile {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Artcile is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Artcile) String() string {
	var builder strings.Builder
	builder.WriteString("Artcile(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", title=")
	builder.WriteString(a.Title)
	builder.WriteString(", description=")
	builder.WriteString(a.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Artciles is a parsable slice of Artcile.
type Artciles []*Artcile

func (a Artciles) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
