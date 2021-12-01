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
			Default(false).
			Annotations(
				entgql.OrderField("ACTIVE"),
			),
		field.String("email").
			Unique().
			Annotations(
				entgql.OrderField("EMAIL"),
			),
		field.String("email_authentication_key").
			Unique().
			Annotations(
				entgql.OrderField("EMAIL_AUTHENTICATION_KEY"),
			),
		field.Time("email_authentication_key_created_at").
			Annotations(
				entgql.OrderField("EMAIL_AUTHENTICATION_KEY_CREATED_AT"),
			),
		field.Bool("email_authentication_status").
			Default(false).
			Annotations(
				entgql.OrderField("EMAIL_AUTHENTICATION_STATUS"),
			),
		field.String("name").
			Default("").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
