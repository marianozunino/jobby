// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent/joboffer"
	"github.com/marianozunino/cc-backend-go/ent/status"
)

// JobOffer is the model entity for the JobOffer schema.
type JobOffer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Reference holds the value of the "reference" field.
	Reference int32 `json:"reference,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate time.Time `json:"end_date,omitempty"`
	// Address1 holds the value of the "address1" field.
	Address1 string `json:"address1,omitempty"`
	// Address2 holds the value of the "address2" field.
	Address2 string `json:"address2,omitempty"`
	// Department holds the value of the "department" field.
	Department string `json:"department,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// WorkingHours holds the value of the "working_hours" field.
	WorkingHours string `json:"working_hours,omitempty"`
	// Salary holds the value of the "salary" field.
	Salary string `json:"salary,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// StatusID holds the value of the "status_id" field.
	StatusID *uuid.UUID `json:"status_id,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// IsFeatured holds the value of the "is_featured" field.
	IsFeatured bool `json:"is_featured,omitempty"`
	// HasBeenEmailed holds the value of the "has_been_emailed" field.
	HasBeenEmailed bool `json:"has_been_emailed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobOfferQuery when eager-loading is set.
	Edges        JobOfferEdges `json:"edges"`
	selectValues sql.SelectValues
}

// JobOfferEdges holds the relations/edges for other nodes in the graph.
type JobOfferEdges struct {
	// Applications holds the value of the applications edge.
	Applications []*Application `json:"applications,omitempty"`
	// JobOfferCategories holds the value of the job_offer_categories edge.
	JobOfferCategories []*JobOfferCategory `json:"job_offer_categories,omitempty"`
	// Status holds the value of the status edge.
	Status *Status `json:"status,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ApplicationsOrErr returns the Applications value or an error if the edge
// was not loaded in eager-loading.
func (e JobOfferEdges) ApplicationsOrErr() ([]*Application, error) {
	if e.loadedTypes[0] {
		return e.Applications, nil
	}
	return nil, &NotLoadedError{edge: "applications"}
}

// JobOfferCategoriesOrErr returns the JobOfferCategories value or an error if the edge
// was not loaded in eager-loading.
func (e JobOfferEdges) JobOfferCategoriesOrErr() ([]*JobOfferCategory, error) {
	if e.loadedTypes[1] {
		return e.JobOfferCategories, nil
	}
	return nil, &NotLoadedError{edge: "job_offer_categories"}
}

// StatusOrErr returns the Status value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobOfferEdges) StatusOrErr() (*Status, error) {
	if e.loadedTypes[2] {
		if e.Status == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: status.Label}
		}
		return e.Status, nil
	}
	return nil, &NotLoadedError{edge: "status"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JobOffer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case joboffer.FieldStatusID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case joboffer.FieldIsFeatured, joboffer.FieldHasBeenEmailed:
			values[i] = new(sql.NullBool)
		case joboffer.FieldReference:
			values[i] = new(sql.NullInt64)
		case joboffer.FieldTitle, joboffer.FieldAddress1, joboffer.FieldAddress2, joboffer.FieldDepartment, joboffer.FieldDescription, joboffer.FieldWorkingHours, joboffer.FieldSalary, joboffer.FieldSlug:
			values[i] = new(sql.NullString)
		case joboffer.FieldStartDate, joboffer.FieldEndDate, joboffer.FieldCreatedAt, joboffer.FieldUpdatedAt, joboffer.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case joboffer.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JobOffer fields.
func (jo *JobOffer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case joboffer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				jo.ID = *value
			}
		case joboffer.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				jo.Title = value.String
			}
		case joboffer.FieldReference:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reference", values[i])
			} else if value.Valid {
				jo.Reference = int32(value.Int64)
			}
		case joboffer.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				jo.StartDate = value.Time
			}
		case joboffer.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				jo.EndDate = value.Time
			}
		case joboffer.FieldAddress1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address1", values[i])
			} else if value.Valid {
				jo.Address1 = value.String
			}
		case joboffer.FieldAddress2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address2", values[i])
			} else if value.Valid {
				jo.Address2 = value.String
			}
		case joboffer.FieldDepartment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field department", values[i])
			} else if value.Valid {
				jo.Department = value.String
			}
		case joboffer.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				jo.Description = value.String
			}
		case joboffer.FieldWorkingHours:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field working_hours", values[i])
			} else if value.Valid {
				jo.WorkingHours = value.String
			}
		case joboffer.FieldSalary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salary", values[i])
			} else if value.Valid {
				jo.Salary = value.String
			}
		case joboffer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				jo.CreatedAt = value.Time
			}
		case joboffer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				jo.UpdatedAt = value.Time
			}
		case joboffer.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				jo.DeletedAt = value.Time
			}
		case joboffer.FieldStatusID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field status_id", values[i])
			} else if value.Valid {
				jo.StatusID = new(uuid.UUID)
				*jo.StatusID = *value.S.(*uuid.UUID)
			}
		case joboffer.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				jo.Slug = value.String
			}
		case joboffer.FieldIsFeatured:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_featured", values[i])
			} else if value.Valid {
				jo.IsFeatured = value.Bool
			}
		case joboffer.FieldHasBeenEmailed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_been_emailed", values[i])
			} else if value.Valid {
				jo.HasBeenEmailed = value.Bool
			}
		default:
			jo.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the JobOffer.
// This includes values selected through modifiers, order, etc.
func (jo *JobOffer) Value(name string) (ent.Value, error) {
	return jo.selectValues.Get(name)
}

// QueryApplications queries the "applications" edge of the JobOffer entity.
func (jo *JobOffer) QueryApplications() *ApplicationQuery {
	return NewJobOfferClient(jo.config).QueryApplications(jo)
}

// QueryJobOfferCategories queries the "job_offer_categories" edge of the JobOffer entity.
func (jo *JobOffer) QueryJobOfferCategories() *JobOfferCategoryQuery {
	return NewJobOfferClient(jo.config).QueryJobOfferCategories(jo)
}

// QueryStatus queries the "status" edge of the JobOffer entity.
func (jo *JobOffer) QueryStatus() *StatusQuery {
	return NewJobOfferClient(jo.config).QueryStatus(jo)
}

// Update returns a builder for updating this JobOffer.
// Note that you need to call JobOffer.Unwrap() before calling this method if this JobOffer
// was returned from a transaction, and the transaction was committed or rolled back.
func (jo *JobOffer) Update() *JobOfferUpdateOne {
	return NewJobOfferClient(jo.config).UpdateOne(jo)
}

// Unwrap unwraps the JobOffer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (jo *JobOffer) Unwrap() *JobOffer {
	_tx, ok := jo.config.driver.(*txDriver)
	if !ok {
		panic("ent: JobOffer is not a transactional entity")
	}
	jo.config.driver = _tx.drv
	return jo
}

// String implements the fmt.Stringer.
func (jo *JobOffer) String() string {
	var builder strings.Builder
	builder.WriteString("JobOffer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", jo.ID))
	builder.WriteString("title=")
	builder.WriteString(jo.Title)
	builder.WriteString(", ")
	builder.WriteString("reference=")
	builder.WriteString(fmt.Sprintf("%v", jo.Reference))
	builder.WriteString(", ")
	builder.WriteString("start_date=")
	builder.WriteString(jo.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_date=")
	builder.WriteString(jo.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("address1=")
	builder.WriteString(jo.Address1)
	builder.WriteString(", ")
	builder.WriteString("address2=")
	builder.WriteString(jo.Address2)
	builder.WriteString(", ")
	builder.WriteString("department=")
	builder.WriteString(jo.Department)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(jo.Description)
	builder.WriteString(", ")
	builder.WriteString("working_hours=")
	builder.WriteString(jo.WorkingHours)
	builder.WriteString(", ")
	builder.WriteString("salary=")
	builder.WriteString(jo.Salary)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(jo.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(jo.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(jo.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := jo.StatusID; v != nil {
		builder.WriteString("status_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(jo.Slug)
	builder.WriteString(", ")
	builder.WriteString("is_featured=")
	builder.WriteString(fmt.Sprintf("%v", jo.IsFeatured))
	builder.WriteString(", ")
	builder.WriteString("has_been_emailed=")
	builder.WriteString(fmt.Sprintf("%v", jo.HasBeenEmailed))
	builder.WriteByte(')')
	return builder.String()
}

// JobOffers is a parsable slice of JobOffer.
type JobOffers []*JobOffer
