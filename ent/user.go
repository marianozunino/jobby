// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
	// ResetPasswordToken holds the value of the "reset_password_token" field.
	ResetPasswordToken string `json:"reset_password_token,omitempty"`
	// ResetPasswordExpires holds the value of the "reset_password_expires" field.
	ResetPasswordExpires time.Time `json:"reset_password_expires,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// Salt holds the value of the "salt" field.
	Salt string `json:"salt,omitempty"`
	// ExternalID holds the value of the "external_id" field.
	ExternalID int32 `json:"external_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// ApplicantProfiles holds the value of the applicant_profiles edge.
	ApplicantProfiles []*ApplicantProfile `json:"applicant_profiles,omitempty"`
	// Interviews holds the value of the interviews edge.
	Interviews []*Interview `json:"interviews,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ApplicantProfilesOrErr returns the ApplicantProfiles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ApplicantProfilesOrErr() ([]*ApplicantProfile, error) {
	if e.loadedTypes[0] {
		return e.ApplicantProfiles, nil
	}
	return nil, &NotLoadedError{edge: "applicant_profiles"}
}

// InterviewsOrErr returns the Interviews value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) InterviewsOrErr() ([]*Interview, error) {
	if e.loadedTypes[1] {
		return e.Interviews, nil
	}
	return nil, &NotLoadedError{edge: "interviews"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[2] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldExternalID:
			values[i] = new(sql.NullInt64)
		case user.FieldEmail, user.FieldPassword, user.FieldRole, user.FieldResetPasswordToken, user.FieldFullName, user.FieldSalt:
			values[i] = new(sql.NullString)
		case user.FieldResetPasswordExpires, user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = user.Role(value.String)
			}
		case user.FieldResetPasswordToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reset_password_token", values[i])
			} else if value.Valid {
				u.ResetPasswordToken = value.String
			}
		case user.FieldResetPasswordExpires:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field reset_password_expires", values[i])
			} else if value.Valid {
				u.ResetPasswordExpires = value.Time
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Time
			}
		case user.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				u.FullName = value.String
			}
		case user.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				u.Salt = value.String
			}
		case user.FieldExternalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field external_id", values[i])
			} else if value.Valid {
				u.ExternalID = int32(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryApplicantProfiles queries the "applicant_profiles" edge of the User entity.
func (u *User) QueryApplicantProfiles() *ApplicantProfileQuery {
	return NewUserClient(u.config).QueryApplicantProfiles(u)
}

// QueryInterviews queries the "interviews" edge of the User entity.
func (u *User) QueryInterviews() *InterviewQuery {
	return NewUserClient(u.config).QueryInterviews(u)
}

// QueryPosts queries the "posts" edge of the User entity.
func (u *User) QueryPosts() *PostQuery {
	return NewUserClient(u.config).QueryPosts(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", ")
	builder.WriteString("reset_password_token=")
	builder.WriteString(u.ResetPasswordToken)
	builder.WriteString(", ")
	builder.WriteString("reset_password_expires=")
	builder.WriteString(u.ResetPasswordExpires.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(u.FullName)
	builder.WriteString(", ")
	builder.WriteString("salt=")
	builder.WriteString(u.Salt)
	builder.WriteString(", ")
	builder.WriteString("external_id=")
	builder.WriteString(fmt.Sprintf("%v", u.ExternalID))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
