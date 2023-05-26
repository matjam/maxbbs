// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/matjam/maxbbs/ent/boards"
)

// Boards is the model entity for the Boards schema.
type Boards struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Boards) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case boards.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Boards fields.
func (b *Boards) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case boards.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Boards.
// This includes values selected through modifiers, order, etc.
func (b *Boards) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Boards.
// Note that you need to call Boards.Unwrap() before calling this method if this Boards
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Boards) Update() *BoardsUpdateOne {
	return NewBoardsClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Boards entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Boards) Unwrap() *Boards {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Boards is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Boards) String() string {
	var builder strings.Builder
	builder.WriteString("Boards(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteByte(')')
	return builder.String()
}

// BoardsSlice is a parsable slice of Boards.
type BoardsSlice []*Boards
