// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent/applicantprofile"
	"github.com/marianozunino/cc-backend-go/ent/application"
	"github.com/marianozunino/cc-backend-go/ent/joboffer"
)

// Application is the model entity for the Application schema.
type Application struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// JobOfferID holds the value of the "job_offer_id" field.
	JobOfferID uuid.UUID `json:"job_offer_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// ApplicantProfileID holds the value of the "applicant_profile_id" field.
	ApplicantProfileID uuid.UUID `json:"applicant_profile_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApplicationQuery when eager-loading is set.
	Edges        ApplicationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ApplicationEdges holds the relations/edges for other nodes in the graph.
type ApplicationEdges struct {
	// ApplicantProfile holds the value of the applicant_profile edge.
	ApplicantProfile *ApplicantProfile `json:"applicant_profile,omitempty"`
	// JobOffer holds the value of the job_offer edge.
	JobOffer *JobOffer `json:"job_offer,omitempty"`
	// Interviews holds the value of the interviews edge.
	Interviews []*Interview `json:"interviews,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ApplicantProfileOrErr returns the ApplicantProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApplicationEdges) ApplicantProfileOrErr() (*ApplicantProfile, error) {
	if e.loadedTypes[0] {
		if e.ApplicantProfile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: applicantprofile.Label}
		}
		return e.ApplicantProfile, nil
	}
	return nil, &NotLoadedError{edge: "applicant_profile"}
}

// JobOfferOrErr returns the JobOffer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApplicationEdges) JobOfferOrErr() (*JobOffer, error) {
	if e.loadedTypes[1] {
		if e.JobOffer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: joboffer.Label}
		}
		return e.JobOffer, nil
	}
	return nil, &NotLoadedError{edge: "job_offer"}
}

// InterviewsOrErr returns the Interviews value or an error if the edge
// was not loaded in eager-loading.
func (e ApplicationEdges) InterviewsOrErr() ([]*Interview, error) {
	if e.loadedTypes[2] {
		return e.Interviews, nil
	}
	return nil, &NotLoadedError{edge: "interviews"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Application) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case application.FieldCreatedAt, application.FieldUpdatedAt, application.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case application.FieldID, application.FieldJobOfferID, application.FieldApplicantProfileID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Application fields.
func (a *Application) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case application.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case application.FieldJobOfferID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field job_offer_id", values[i])
			} else if value != nil {
				a.JobOfferID = *value
			}
		case application.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case application.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case application.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		case application.FieldApplicantProfileID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field applicant_profile_id", values[i])
			} else if value != nil {
				a.ApplicantProfileID = *value
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Application.
// This includes values selected through modifiers, order, etc.
func (a *Application) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryApplicantProfile queries the "applicant_profile" edge of the Application entity.
func (a *Application) QueryApplicantProfile() *ApplicantProfileQuery {
	return NewApplicationClient(a.config).QueryApplicantProfile(a)
}

// QueryJobOffer queries the "job_offer" edge of the Application entity.
func (a *Application) QueryJobOffer() *JobOfferQuery {
	return NewApplicationClient(a.config).QueryJobOffer(a)
}

// QueryInterviews queries the "interviews" edge of the Application entity.
func (a *Application) QueryInterviews() *InterviewQuery {
	return NewApplicationClient(a.config).QueryInterviews(a)
}

// Update returns a builder for updating this Application.
// Note that you need to call Application.Unwrap() before calling this method if this Application
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Application) Update() *ApplicationUpdateOne {
	return NewApplicationClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Application entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Application) Unwrap() *Application {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Application is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Application) String() string {
	var builder strings.Builder
	builder.WriteString("Application(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("job_offer_id=")
	builder.WriteString(fmt.Sprintf("%v", a.JobOfferID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("applicant_profile_id=")
	builder.WriteString(fmt.Sprintf("%v", a.ApplicantProfileID))
	builder.WriteByte(')')
	return builder.String()
}

// Applications is a parsable slice of Application.
type Applications []*Application