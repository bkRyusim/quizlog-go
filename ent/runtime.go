// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/bkRyusim/quizlog-go/ent/blog"
	"github.com/bkRyusim/quizlog-go/ent/quiz"
	"github.com/bkRyusim/quizlog-go/ent/schema"
	"github.com/bkRyusim/quizlog-go/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	blogFields := schema.Blog{}.Fields()
	_ = blogFields
	// blogDescCreatedAt is the schema descriptor for createdAt field.
	blogDescCreatedAt := blogFields[2].Descriptor()
	// blog.DefaultCreatedAt holds the default value on creation for the createdAt field.
	blog.DefaultCreatedAt = blogDescCreatedAt.Default.(func() time.Time)
	// blogDescUpdatedAt is the schema descriptor for updatedAt field.
	blogDescUpdatedAt := blogFields[3].Descriptor()
	// blog.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	blog.DefaultUpdatedAt = blogDescUpdatedAt.Default.(func() time.Time)
	// blog.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	blog.UpdateDefaultUpdatedAt = blogDescUpdatedAt.UpdateDefault.(func() time.Time)
	quizFields := schema.Quiz{}.Fields()
	_ = quizFields
	// quizDescCreatedAt is the schema descriptor for createdAt field.
	quizDescCreatedAt := quizFields[3].Descriptor()
	// quiz.DefaultCreatedAt holds the default value on creation for the createdAt field.
	quiz.DefaultCreatedAt = quizDescCreatedAt.Default.(func() time.Time)
	// quizDescUpdatedAt is the schema descriptor for updatedAt field.
	quizDescUpdatedAt := quizFields[4].Descriptor()
	// quiz.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	quiz.DefaultUpdatedAt = quizDescUpdatedAt.Default.(func() time.Time)
	// quiz.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	quiz.UpdateDefaultUpdatedAt = quizDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the createdAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updatedAt field.
	userDescUpdatedAt := userFields[3].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}