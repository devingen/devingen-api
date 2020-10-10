package model

import (
	coremodel "github.com/devingen/api-core/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workspace struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `json:"name"`
}

func (w *Workspace) DBRef(database string) *coremodel.DBRef {
	return &coremodel.DBRef{
		Ref:      CollectionWorkspaces,
		ID:       w.ID,
		Database: database,
	}
}
