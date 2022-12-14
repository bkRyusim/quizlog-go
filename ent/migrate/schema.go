// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlogsColumns holds the columns for the "blogs" table.
	BlogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "blog_name", Type: field.TypeString},
		{Name: "blog_title", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_blogs", Type: field.TypeInt, Nullable: true},
	}
	// BlogsTable holds the schema information for the "blogs" table.
	BlogsTable = &schema.Table{
		Name:       "blogs",
		Columns:    BlogsColumns,
		PrimaryKey: []*schema.Column{BlogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blogs_users_blogs",
				Columns:    []*schema.Column{BlogsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// QuizsColumns holds the columns for the "quizs" table.
	QuizsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "post_url", Type: field.TypeString},
		{Name: "question", Type: field.TypeString},
		{Name: "answer", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_quiz", Type: field.TypeInt, Nullable: true},
	}
	// QuizsTable holds the schema information for the "quizs" table.
	QuizsTable = &schema.Table{
		Name:       "quizs",
		Columns:    QuizsColumns,
		PrimaryKey: []*schema.Column{QuizsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "quizs_users_quiz",
				Columns:    []*schema.Column{QuizsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "user_name", Type: field.TypeString},
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
		BlogsTable,
		PostsTable,
		QuizsTable,
		UsersTable,
	}
)

func init() {
	BlogsTable.ForeignKeys[0].RefTable = UsersTable
	QuizsTable.ForeignKeys[0].RefTable = UsersTable
}
