// Code generated by ent, DO NOT EDIT.

package tower

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tower type in the database.
	Label = "tower"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumberOfFloors holds the string denoting the number_of_floors field in the database.
	FieldNumberOfFloors = "number_of_floors"
	// FieldNumberOfApartmentsPerFloor holds the string denoting the number_of_apartments_per_floor field in the database.
	FieldNumberOfApartmentsPerFloor = "number_of_apartments_per_floor"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRegisteredAt holds the string denoting the registered_at field in the database.
	FieldRegisteredAt = "registered_at"
	// EdgeApartments holds the string denoting the apartments edge name in mutations.
	EdgeApartments = "apartments"
	// Table holds the table name of the tower in the database.
	Table = "towers"
	// ApartmentsTable is the table that holds the apartments relation/edge.
	ApartmentsTable = "apartments"
	// ApartmentsInverseTable is the table name for the Apartment entity.
	// It exists in this package in order to avoid circular dependency with the "apartment" package.
	ApartmentsInverseTable = "apartments"
	// ApartmentsColumn is the table column denoting the apartments relation/edge.
	ApartmentsColumn = "tower_apartments"
)

// Columns holds all SQL columns for tower fields.
var Columns = []string{
	FieldID,
	FieldNumberOfFloors,
	FieldNumberOfApartmentsPerFloor,
	FieldName,
	FieldRegisteredAt,
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

var (
	// NumberOfFloorsValidator is a validator for the "number_of_floors" field. It is called by the builders before save.
	NumberOfFloorsValidator func(int) error
	// NumberOfApartmentsPerFloorValidator is a validator for the "number_of_apartments_per_floor" field. It is called by the builders before save.
	NumberOfApartmentsPerFloorValidator func(int) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)

// OrderOption defines the ordering options for the Tower queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNumberOfFloors orders the results by the number_of_floors field.
func ByNumberOfFloors(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumberOfFloors, opts...).ToFunc()
}

// ByNumberOfApartmentsPerFloor orders the results by the number_of_apartments_per_floor field.
func ByNumberOfApartmentsPerFloor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumberOfApartmentsPerFloor, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByRegisteredAt orders the results by the registered_at field.
func ByRegisteredAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegisteredAt, opts...).ToFunc()
}

// ByApartmentsCount orders the results by apartments count.
func ByApartmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newApartmentsStep(), opts...)
	}
}

// ByApartments orders the results by apartments terms.
func ByApartments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApartmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newApartmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApartmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ApartmentsTable, ApartmentsColumn),
	)
}
