// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/matjam/maxbbs/ent/message"
	"github.com/matjam/maxbbs/ent/schema"
	"github.com/matjam/maxbbs/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreated is the schema descriptor for created field.
	messageDescCreated := messageFields[0].Descriptor()
	// message.DefaultCreated holds the default value on creation for the created field.
	message.DefaultCreated = messageDescCreated.Default.(func() time.Time)
	// messageDescUpdated is the schema descriptor for updated field.
	messageDescUpdated := messageFields[1].Descriptor()
	// message.DefaultUpdated holds the default value on creation for the updated field.
	message.DefaultUpdated = messageDescUpdated.Default.(func() time.Time)
	// messageDescDeleted is the schema descriptor for deleted field.
	messageDescDeleted := messageFields[2].Descriptor()
	// message.DefaultDeleted holds the default value on creation for the deleted field.
	message.DefaultDeleted = messageDescDeleted.Default.(func() time.Time)
	// messageDescSubject is the schema descriptor for subject field.
	messageDescSubject := messageFields[3].Descriptor()
	// message.SubjectValidator is a validator for the "subject" field. It is called by the builders before save.
	message.SubjectValidator = messageDescSubject.Validators[0].(func(string) error)
	// messageDescBody is the schema descriptor for body field.
	messageDescBody := messageFields[4].Descriptor()
	// message.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	message.BodyValidator = messageDescBody.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreated is the schema descriptor for created field.
	userDescCreated := userFields[0].Descriptor()
	// user.DefaultCreated holds the default value on creation for the created field.
	user.DefaultCreated = userDescCreated.Default.(func() time.Time)
	// userDescUpdated is the schema descriptor for updated field.
	userDescUpdated := userFields[1].Descriptor()
	// user.DefaultUpdated holds the default value on creation for the updated field.
	user.DefaultUpdated = userDescUpdated.Default.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[2].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[4].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[5].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[6].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
}