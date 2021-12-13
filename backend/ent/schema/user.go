package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("active").
			Default(false),
		field.String("email").
			Unique().
			Annotations(
				entgql.OrderField("EMAIL"),
			),
		field.String("email_authentication_key").
			Unique(),
		field.Time("email_authentication_key_created_at"),
		field.Bool("email_authentication_status").
			Default(false),
		field.String("name").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("password").
			Default(""),
		field.Int("role").
			Default(0),
		field.String("reset_password_key").
			Unique(),
		field.Time("reset_password_key_created_at"),
		field.Int("height"),
		field.Enum("height_display").
			Values("PRIVATE", "FRIENDS", "PUBLIC"),
		field.Int("weight"),
		field.Enum("weight_display").
			Values("PRIVATE", "FRIENDS", "PUBLIC"),
		field.Int("wingspan"),
		field.Enum("wingspan_display").
			Values("PRIVATE", "FRIENDS", "PUBLIC"),
		field.Time("birthday"),
		field.Enum("birthday_display").
			Values("PRIVATE", "FRIENDS", "PUBLIC"),
		field.Enum("gender").
			Values("MALE", "FEMALE"),
		field.Enum("gender_display").
			Values("PRIVATE", "FRIENDS", "PUBLIC"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
