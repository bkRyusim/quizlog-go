// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/bkRyusim/quizlog-go/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserId holds the value of the "userId" field.
	UserId string `json:"userId,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Blogs holds the value of the blogs edge.
	Blogs []*Blog `json:"blogs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BlogsOrErr returns the Blogs value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) BlogsOrErr() ([]*Blog, error) {
	if e.loadedTypes[0] {
		return e.Blogs, nil
	}
	return nil, &NotLoadedError{edge: "blogs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUserId, user.FieldName:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUserId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				u.UserId = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryBlogs queries the "blogs" edge of the User entity.
func (u *User) QueryBlogs() *BlogQuery {
	return (&UserClient{config: u.config}).QueryBlogs(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("userId=")
	builder.WriteString(u.UserId)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
