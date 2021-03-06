// Code generated by entc, DO NOT EDIT.

package userstatus

const (
	// Label holds the string label denoting the userstatus type in the database.
	Label = "user_status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserStatus holds the string denoting the userstatus field in the database.
	FieldUserStatus = "user_status"

	// EdgeUserstatusUser holds the string denoting the userstatususer edge name in mutations.
	EdgeUserstatusUser = "UserstatusUser"

	// Table holds the table name of the userstatus in the database.
	Table = "user_status"
	// UserstatusUserTable is the table the holds the UserstatusUser relation/edge.
	UserstatusUserTable = "users"
	// UserstatusUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserstatusUserInverseTable = "users"
	// UserstatusUserColumn is the table column denoting the UserstatusUser relation/edge.
	UserstatusUserColumn = "user_status_userstatus_user"
)

// Columns holds all SQL columns for userstatus fields.
var Columns = []string{
	FieldID,
	FieldUserStatus,
}
