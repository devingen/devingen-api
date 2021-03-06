package database_service

import (
	"context"
	"github.com/devingen/devingen-api/model"
	kimlik_model "github.com/devingen/kimlik-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service DatabaseService) CreateWorkspace(base, name string, createdBy *kimlik_model.User) (*model.Workspace, error) {
	collection, err := service.Database.ConnectToCollection(base, model.CollectionWorkspaces)
	if err != nil {
		return nil, err
	}

	item := &model.Workspace{
		CreatedBy: createdBy.DBRef(base),
		Name:      name,
	}
	item.AddCreationFields()

	result, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}

	item.ID = result.InsertedID.(primitive.ObjectID)
	return item, nil
}
