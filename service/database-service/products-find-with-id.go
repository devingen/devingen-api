package database_service

import (
	"github.com/devingen/devingen-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (service DatabaseService) FindProductWithId(base, id string) (*model.Product, error) {
	result := make([]*model.Product, 0)

	mId, mErr := primitive.ObjectIDFromHex(id)
	if mErr != nil {
		return nil, mErr
	}
	query := bson.M{"_id": mId}

	err := service.Database.Query(base, model.CollectionProducts, query, 0, func(cur *mongo.Cursor) error {
		var data model.Product
		err := cur.Decode(&data)
		if err != nil {
			return err
		}
		result = append(result, &data)
		return nil
	})
	if len(result) > 0 {
		return result[0], err
	}
	return nil, err
}
