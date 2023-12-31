// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldResetPasswordToken holds the string denoting the reset_password_token field in the database.
	FieldResetPasswordToken = "reset_password_token"
	// FieldResetPasswordExpires holds the string denoting the reset_password_expires field in the database.
	FieldResetPasswordExpires = "reset_password_expires"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldExternalID holds the string denoting the external_id field in the database.
	FieldExternalID = "external_id"
	// EdgeApplicantProfiles holds the string denoting the applicant_profiles edge name in mutations.
	EdgeApplicantProfiles = "applicant_profiles"
	// EdgeInterviews holds the string denoting the interviews edge name in mutations.
	EdgeInterviews = "interviews"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ApplicantProfilesTable is the table that holds the applicant_profiles relation/edge.
	ApplicantProfilesTable = "applicant_profiles"
	// ApplicantProfilesInverseTable is the table name for the ApplicantProfile entity.
	// It exists in this package in order to avoid circular dependency with the "applicantprofile" package.
	ApplicantProfilesInverseTable = "applicant_profiles"
	// ApplicantProfilesColumn is the table column denoting the applicant_profiles relation/edge.
	ApplicantProfilesColumn = "user_id"
	// InterviewsTable is the table that holds the interviews relation/edge.
	InterviewsTable = "interviews"
	// InterviewsInverseTable is the table name for the Interview entity.
	// It exists in this package in order to avoid circular dependency with the "interview" package.
	InterviewsInverseTable = "interviews"
	// InterviewsColumn is the table column denoting the interviews relation/edge.
	InterviewsColumn = "interviewer_id"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "author_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPassword,
	FieldRole,
	FieldResetPasswordToken,
	FieldResetPasswordExpires,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldFullName,
	FieldSalt,
	FieldExternalID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Role defines the type for the "role" enum field.
type Role string

// Role values.
const (
	RoleADMIN     Role = "ADMIN"
	RoleAPPLICANT Role = "APPLICANT"
	RoleRECRUITER Role = "RECRUITER"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleADMIN, RoleAPPLICANT, RoleRECRUITER:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByResetPasswordToken orders the results by the reset_password_token field.
func ByResetPasswordToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResetPasswordToken, opts...).ToFunc()
}

// ByResetPasswordExpires orders the results by the reset_password_expires field.
func ByResetPasswordExpires(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResetPasswordExpires, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// BySalt orders the results by the salt field.
func BySalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSalt, opts...).ToFunc()
}

// ByExternalID orders the results by the external_id field.
func ByExternalID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExternalID, opts...).ToFunc()
}

// ByApplicantProfilesCount orders the results by applicant_profiles count.
func ByApplicantProfilesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newApplicantProfilesStep(), opts...)
	}
}

// ByApplicantProfiles orders the results by applicant_profiles terms.
func ByApplicantProfiles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApplicantProfilesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInterviewsCount orders the results by interviews count.
func ByInterviewsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInterviewsStep(), opts...)
	}
}

// ByInterviews orders the results by interviews terms.
func ByInterviews(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInterviewsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostsCount orders the results by posts count.
func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostsStep(), opts...)
	}
}

// ByPosts orders the results by posts terms.
func ByPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newApplicantProfilesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApplicantProfilesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ApplicantProfilesTable, ApplicantProfilesColumn),
	)
}
func newInterviewsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InterviewsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InterviewsTable, InterviewsColumn),
	)
}
func newPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
	)
}
