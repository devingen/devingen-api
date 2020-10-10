package model

import (
	"github.com/devingen/kimlik-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type WorkspaceOwnership struct {
	// DBRef fields
	Ref      string             `bson:"_ref,omitempty" json:"ref,omitempty"`
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Database string             `bson:"_db,omitempty" json:"db,omitempty"`

	// common model fields
	CreatedAt *time.Time `json:"_created,omitempty" bson:"_created,omitempty"`
	UpdatedAt *time.Time `json:"_updated,omitempty" bson:"_updated,omitempty"`
	Revision  int        `json:"_revision,omitempty" bson:"_revision,omitempty"`

	User      *model.User `json:"user"`
	Workspace *Workspace  `json:"workspace"`
}

func (w *WorkspaceOwnership) AddCreationFields() {
	w.ID = primitive.NewObjectID()
	now := time.Now()
	w.CreatedAt = &now
	w.UpdatedAt = &now
	w.Revision = 1
}

// PrepareUpdateFields sets the UpdatedAt and deletes the Revision. Giving 0 value to Revision results bson
// ignoring the revision field in $set function. It's incremented by the $inc command
func (w *WorkspaceOwnership) PrepareUpdateFields() {
	w.Revision = 0
	now := time.Now()
	w.UpdatedAt = &now
}
