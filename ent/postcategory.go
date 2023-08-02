// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent/postcategory"
)

// PostCategory is the model entity for the PostCategory schema.
type PostCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostCategoryQuery when eager-loading is set.
	Edges        PostCategoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PostCategoryEdges holds the relations/edges for other nodes in the graph.
type PostCategoryEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e PostCategoryEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PostCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case postcategory.FieldName:
			values[i] = new(sql.NullString)
		case postcategory.FieldCreatedAt, postcategory.FieldUpdatedAt, postcategory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case postcategory.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PostCategory fields.
func (pc *PostCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case postcategory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pc.ID = *value
			}
		case postcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pc.Name = value.String
			}
		case postcategory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pc.CreatedAt = value.Time
			}
		case postcategory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pc.UpdatedAt = value.Time
			}
		case postcategory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pc.DeletedAt = new(time.Time)
				*pc.DeletedAt = value.Time
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PostCategory.
// This includes values selected through modifiers, order, etc.
func (pc *PostCategory) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// QueryPosts queries the "posts" edge of the PostCategory entity.
func (pc *PostCategory) QueryPosts() *PostQuery {
	return NewPostCategoryClient(pc.config).QueryPosts(pc)
}

// Update returns a builder for updating this PostCategory.
// Note that you need to call PostCategory.Unwrap() before calling this method if this PostCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PostCategory) Update() *PostCategoryUpdateOne {
	return NewPostCategoryClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the PostCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PostCategory) Unwrap() *PostCategory {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PostCategory is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PostCategory) String() string {
	var builder strings.Builder
	builder.WriteString("PostCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("name=")
	builder.WriteString(pc.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := pc.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// PostCategories is a parsable slice of PostCategory.
type PostCategories []*PostCategory
