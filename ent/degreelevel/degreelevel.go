// Code generated by ent, DO NOT EDIT.

package degreelevel

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the degreelevel type in the database.
	Label = "degree_level"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeEducations holds the string denoting the educations edge name in mutations.
	EdgeEducations = "educations"
	// Table holds the table name of the degreelevel in the database.
	Table = "degree_levels"
	// EducationsTable is the table that holds the educations relation/edge.
	EducationsTable = "education"
	// EducationsInverseTable is the table name for the Education entity.
	// It exists in this package in order to avoid circular dependency with the "education" package.
	EducationsInverseTable = "education"
	// EducationsColumn is the table column denoting the educations relation/edge.
	EducationsColumn = "degree_level_id"
)

// Columns holds all SQL columns for degreelevel fields.
var Columns = []string{
	FieldID,
	FieldName,
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

// OrderOption defines the ordering options for the DegreeLevel queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
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

// ByEducationsCount orders the results by educations count.
func ByEducationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEducationsStep(), opts...)
	}
}

// ByEducations orders the results by educations terms.
func ByEducations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEducationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEducationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EducationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EducationsTable, EducationsColumn),
	)
}
