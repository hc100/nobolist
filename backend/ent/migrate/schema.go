// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "email_authentication_key", Type: field.TypeString, Unique: true},
		{Name: "email_authentication_key_created_at", Type: field.TypeTime},
		{Name: "email_authentication_status", Type: field.TypeBool, Default: false},
		{Name: "name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString, Default: ""},
		{Name: "role", Type: field.TypeInt, Default: 0},
		{Name: "reset_password_key", Type: field.TypeString, Unique: true},
		{Name: "reset_password_key_created_at", Type: field.TypeTime},
		{Name: "height", Type: field.TypeInt},
		{Name: "height_display", Type: field.TypeEnum, Enums: []string{"PRIVATE", "FRIENDS", "PUBLIC"}},
		{Name: "weight", Type: field.TypeInt},
		{Name: "weight_display", Type: field.TypeEnum, Enums: []string{"PRIVATE", "FRIENDS", "PUBLIC"}},
		{Name: "wingspan", Type: field.TypeInt},
		{Name: "wingspan_display", Type: field.TypeEnum, Enums: []string{"PRIVATE", "FRIENDS", "PUBLIC"}},
		{Name: "birthday", Type: field.TypeTime},
		{Name: "birthday_display", Type: field.TypeEnum, Enums: []string{"PRIVATE", "FRIENDS", "PUBLIC"}},
		{Name: "gender", Type: field.TypeEnum, Enums: []string{"MALE", "FEMALE"}},
		{Name: "gender_display", Type: field.TypeEnum, Enums: []string{"PRIVATE", "FRIENDS", "PUBLIC"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}
