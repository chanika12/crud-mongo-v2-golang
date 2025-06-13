package mongoDb

import (
	"context"
	"crud/domain/model"
	"crud/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoOrderRepository struct {
	collection *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) repository.IOrderRepository {
	return &MongoOrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *MongoOrderRepository) Save(order model.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, &order)
	return err
}
