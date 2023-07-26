// Code generated by ent, DO NOT EDIT.

package interview

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the interview type in the database.
	Label = "interview"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldInterviewDate holds the string denoting the interview_date field in the database.
	FieldInterviewDate = "interview_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldApplicationID holds the string denoting the application_id field in the database.
	FieldApplicationID = "application_id"
	// FieldInterviewerID holds the string denoting the interviewer_id field in the database.
	FieldInterviewerID = "interviewer_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeApplication holds the string denoting the application edge name in mutations.
	EdgeApplication = "application"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the interview in the database.
	Table = "interviews"
	// ApplicationTable is the table that holds the application relation/edge.
	ApplicationTable = "interviews"
	// ApplicationInverseTable is the table name for the Application entity.
	// It exists in this package in order to avoid circular dependency with the "application" package.
	ApplicationInverseTable = "applications"
	// ApplicationColumn is the table column denoting the application relation/edge.
	ApplicationColumn = "application_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "interviews"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "interviewer_id"
)

// Columns holds all SQL columns for interview fields.
var Columns = []string{
	FieldID,
	FieldComment,
	FieldInterviewDate,
	FieldStatus,
	FieldApplicationID,
	FieldInterviewerID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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

// OrderOption defines the ordering options for the Interview queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByInterviewDate orders the results by the interview_date field.
func ByInterviewDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInterviewDate, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByApplicationID orders the results by the application_id field.
func ByApplicationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplicationID, opts...).ToFunc()
}

// ByInterviewerID orders the results by the interviewer_id field.
func ByInterviewerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInterviewerID, opts...).ToFunc()
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

// ByApplicationField orders the results by application field.
func ByApplicationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApplicationStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newApplicationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApplicationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ApplicationTable, ApplicationColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
