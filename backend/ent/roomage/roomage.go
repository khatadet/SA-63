// Code generated by entc, DO NOT EDIT.

package roomage

const (
	// Label holds the string label denoting the roomage type in the database.
	Label = "roomage"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoomAge holds the string denoting the roomage field in the database.
	FieldRoomAge = "room_age"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"

	// EdgeRoomageRent holds the string denoting the roomagerent edge name in mutations.
	EdgeRoomageRent = "RoomageRent"

	// Table holds the table name of the roomage in the database.
	Table = "roomages"
	// RoomageRentTable is the table the holds the RoomageRent relation/edge.
	RoomageRentTable = "rents"
	// RoomageRentInverseTable is the table name for the Rent entity.
	// It exists in this package in order to avoid circular dependency with the "rent" package.
	RoomageRentInverseTable = "rents"
	// RoomageRentColumn is the table column denoting the RoomageRent relation/edge.
	RoomageRentColumn = "roomage_roomage_rent"
)

// Columns holds all SQL columns for roomage fields.
var Columns = []string{
	FieldID,
	FieldRoomAge,
	FieldText,
}
