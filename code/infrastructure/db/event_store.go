package mongoDb

import (
	"context"
	"crud/domain/model"
	"crud/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrderRepository struct {
	collection *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) repository.OrderRepository {
	return &MongoOrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *MongoOrderRepository) Save(ctx context.Context, order *model.Order) error {
	_, err := r.collection.InsertOne(ctx, order)
	return err
}
