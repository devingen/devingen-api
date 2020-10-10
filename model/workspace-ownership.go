package model

import (
	coremodel "github.com/devingen/api-core/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkspaceOwnership struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User      *coremodel.DBRef    `json:"user"`
	Workspace *coremodel.DBRef    `json:"workspace"`
}
