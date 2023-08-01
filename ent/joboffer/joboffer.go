// Code generated by ent, DO NOT EDIT.

package joboffer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the joboffer type in the database.
	Label = "job_offer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldReference holds the string denoting the reference field in the database.
	FieldReference = "reference"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldAddress1 holds the string denoting the address1 field in the database.
	FieldAddress1 = "address1"
	// FieldAddress2 holds the string denoting the address2 field in the database.
	FieldAddress2 = "address2"
	// FieldDepartment holds the string denoting the department field in the database.
	FieldDepartment = "department"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldWorkingHours holds the string denoting the working_hours field in the database.
	FieldWorkingHours = "working_hours"
	// FieldSalary holds the string denoting the salary field in the database.
	FieldSalary = "salary"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldIsFeatured holds the string denoting the is_featured field in the database.
	FieldIsFeatured = "is_featured"
	// FieldHasBeenEmailed holds the string denoting the has_been_emailed field in the database.
	FieldHasBeenEmailed = "has_been_emailed"
	// FieldStatusID holds the string denoting the status_id field in the database.
	FieldStatusID = "status_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeApplications holds the string denoting the applications edge name in mutations.
	EdgeApplications = "applications"
	// EdgeJobOfferCategories holds the string denoting the job_offer_categories edge name in mutations.
	EdgeJobOfferCategories = "job_offer_categories"
	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// Table holds the table name of the joboffer in the database.
	Table = "job_offers"
	// ApplicationsTable is the table that holds the applications relation/edge.
	ApplicationsTable = "applications"
	// ApplicationsInverseTable is the table name for the Application entity.
	// It exists in this package in order to avoid circular dependency with the "application" package.
	ApplicationsInverseTable = "applications"
	// ApplicationsColumn is the table column denoting the applications relation/edge.
	ApplicationsColumn = "job_offer_id"
	// JobOfferCategoriesTable is the table that holds the job_offer_categories relation/edge.
	JobOfferCategoriesTable = "job_offer_categories"
	// JobOfferCategoriesInverseTable is the table name for the JobOfferCategory entity.
	// It exists in this package in order to avoid circular dependency with the "joboffercategory" package.
	JobOfferCategoriesInverseTable = "job_offer_categories"
	// JobOfferCategoriesColumn is the table column denoting the job_offer_categories relation/edge.
	JobOfferCategoriesColumn = "job_offer_id"
	// StatusTable is the table that holds the status relation/edge.
	StatusTable = "job_offers"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "status_id"
)

// Columns holds all SQL columns for joboffer fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldReference,
	FieldStartDate,
	FieldEndDate,
	FieldAddress1,
	FieldAddress2,
	FieldDepartment,
	FieldDescription,
	FieldWorkingHours,
	FieldSalary,
	FieldSlug,
	FieldIsFeatured,
	FieldHasBeenEmailed,
	FieldStatusID,
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

// OrderOption defines the ordering options for the JobOffer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByReference orders the results by the reference field.
func ByReference(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReference, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByAddress1 orders the results by the address1 field.
func ByAddress1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress1, opts...).ToFunc()
}

// ByAddress2 orders the results by the address2 field.
func ByAddress2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress2, opts...).ToFunc()
}

// ByDepartment orders the results by the department field.
func ByDepartment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartment, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByWorkingHours orders the results by the working_hours field.
func ByWorkingHours(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkingHours, opts...).ToFunc()
}

// BySalary orders the results by the salary field.
func BySalary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSalary, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByIsFeatured orders the results by the is_featured field.
func ByIsFeatured(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFeatured, opts...).ToFunc()
}

// ByHasBeenEmailed orders the results by the has_been_emailed field.
func ByHasBeenEmailed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasBeenEmailed, opts...).ToFunc()
}

// ByStatusID orders the results by the status_id field.
func ByStatusID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatusID, opts...).ToFunc()
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

// ByApplicationsCount orders the results by applications count.
func ByApplicationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newApplicationsStep(), opts...)
	}
}

// ByApplications orders the results by applications terms.
func ByApplications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApplicationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByJobOfferCategoriesCount orders the results by job_offer_categories count.
func ByJobOfferCategoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newJobOfferCategoriesStep(), opts...)
	}
}

// ByJobOfferCategories orders the results by job_offer_categories terms.
func ByJobOfferCategories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newJobOfferCategoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStatusField orders the results by status field.
func ByStatusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatusStep(), sql.OrderByField(field, opts...))
	}
}
func newApplicationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApplicationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ApplicationsTable, ApplicationsColumn),
	)
}
func newJobOfferCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(JobOfferCategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, JobOfferCategoriesTable, JobOfferCategoriesColumn),
	)
}
func newStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, StatusTable, StatusColumn),
	)
}
