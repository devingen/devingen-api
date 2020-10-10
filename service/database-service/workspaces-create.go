package database_service

import (
	"context"
	"github.com/devingen/devingen-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service DatabaseService) WorkspacesCreate(base, name string) (*model.Workspace, error) {
	collection, err := service.Database.ConnectToCollection(base, model.CollectionWorkspaces)
	if err != nil {
		return nil, err
	}

	item := &model.Workspace{
		Name: name,
	}

	result, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}

	item.ID = result.InsertedID.(primitive.ObjectID)
	return item, nil
}
