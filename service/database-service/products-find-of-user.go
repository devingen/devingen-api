package database_service

import (
	"github.com/devingen/devingen-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (service DatabaseService) FindProductsOfUser(base, userId string) ([]*model.Product, error) {
	workspaceOwnerships, err := service.FindWorkspaceOwnershipsOfUser(base, userId)
	if err != nil {
		return nil, err
	}

	workspaceIds := []primitive.ObjectID{}
	for _, v := range workspaceOwnerships {
		workspaceIds = append(workspaceIds, v.Workspace.ID)
	}

	result := make([]*model.Product, 0)
	query := bson.M{"workspace._id": bson.M{"$in": workspaceIds}}

	err = service.Database.Query(base, model.CollectionProducts, query, 0, func(cur *mongo.Cursor) error {
		var data model.Product
		err := cur.Decode(&data)
		if err != nil {
			return err
		}
		result = append(result, &data)
		return nil
	})
	return result, err
}
