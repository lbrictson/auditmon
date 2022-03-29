// Code generated by entc, DO NOT EDIT.

package eventnameautofill

const (
	// Label holds the string label denoting the eventnameautofill type in the database.
	Label = "event_name_autofill"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEventName holds the string denoting the eventname field in the database.
	FieldEventName = "event_name"
	// Table holds the table name of the eventnameautofill in the database.
	Table = "event_name_autofills"
)

// Columns holds all SQL columns for eventnameautofill fields.
var Columns = []string{
	FieldID,
	FieldEventName,
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