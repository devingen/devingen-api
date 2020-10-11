package database_service

import (
	"github.com/devingen/devingen-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (service DatabaseService) FindWorkspaceOwnershipsOfUser(base, userId string) ([]*model.WorkspaceOwnership, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	result := make([]*model.WorkspaceOwnership, 0)
	query := []bson.M{
		{"$match": bson.M{"$and": bson.A{bson.M{"user._id": id}}}},
		{"$lookup": bson.M{
			"from":         model.CollectionWorkspaces,
			"localField":   "workspace._id",
			"foreignField": "_id",
			"as":           "workspace",
		}},
		{"$unwind": bson.M{"path": "$workspace", "preserveNullAndEmptyArrays": true}},
	}

	err = service.Database.Aggregate(base, model.CollectionWorkspaceOwnerships, query, func(cur *mongo.Cursor) error {
		var data model.WorkspaceOwnership
		err := cur.Decode(&data)
		if err != nil {
			return err
		}
		result = append(result, &data)
		return nil
	})
	return result, err
}
