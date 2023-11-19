// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "user_posts", Type: field.TypeInt, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PostsTable,
		UsersTable,
	}
)

func init() {
	PostsTable.ForeignKeys[0].RefTable = UsersTable
}
