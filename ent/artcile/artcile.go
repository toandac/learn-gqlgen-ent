// Code generated by entc, DO NOT EDIT.

package artcile

const (
	// Label holds the string label denoting the artcile type in the database.
	Label = "artcile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the artcile in the database.
	Table = "artciles"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "artciles"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for artcile fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "artciles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
)
