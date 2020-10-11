package database_service

import (
	"context"
	"github.com/devingen/devingen-api/model"
	kimlik_model "github.com/devingen/kimlik-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service DatabaseService) CreateProduct(base, name string, productType model.ProductType, workspace *model.Workspace, createdBy *kimlik_model.User) (*model.Product, error) {
	collection, err := service.Database.ConnectToCollection(base, model.CollectionProducts)
	if err != nil {
		return nil, err
	}

	item := &model.Product{
		CreatedBy: createdBy.DBRef(base),
		Name:      name,
		Type:      productType,
		Workspace: workspace.DBRef(base),
	}
	item.AddCreationFields()

	result, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}

	item.ID = result.InsertedID.(primitive.ObjectID)
	return item, nil
}
