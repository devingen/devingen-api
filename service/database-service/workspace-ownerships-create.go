package database_service

import (
	"context"
	"github.com/devingen/devingen-api/model"
	kimlik_model "github.com/devingen/kimlik-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service DatabaseService) WorkspaceOwnershipsCreate(base string, user *kimlik_model.User, workspace *model.Workspace) (*model.WorkspaceOwnership, error) {
	collection, err := service.Database.ConnectToCollection(base, model.CollectionWorkspaceOwnerships)
	if err != nil {
		return nil, err
	}

	item := &model.WorkspaceOwnership{
		User:      user.DBRef(base),
		Workspace: workspace.DBRef(base),
	}

	result, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}

	item.ID = result.InsertedID.(primitive.ObjectID)
	return item, nil
}
