// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArtcilesColumns holds the columns for the "artciles" table.
	ArtcilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Default: ""},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// ArtcilesTable holds the schema information for the "artciles" table.
	ArtcilesTable = &schema.Table{
		Name:       "artciles",
		Columns:    ArtcilesColumns,
		PrimaryKey: []*schema.Column{ArtcilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "artciles_users_articles",
				Columns:    []*schema.Column{ArtcilesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "age", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArtcilesTable,
		UsersTable,
	}
)

func init() {
	ArtcilesTable.ForeignKeys[0].RefTable = UsersTable
}