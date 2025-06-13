package injecter

import (
	"context"
	"crud/infrastructure/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//	func NewMongoClient(lc fx.Lifecycle) (*mongo.Client, error) {
//		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Env.Mongo.Host))
//		if err != nil {
//			return nil, err
//		}
//
//		lc.Append(fx.Hook{
//			OnStop: func(ctx context.Context) error {
//				return client.Disconnect(ctx)
//			},
//		})
//
//		return client, nil
//	}
func SetupMongoDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	host, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Env.Mongo.Host))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("start mongo.........")
	fmt.Println("-------------------")
	return host.Database(config.Env.Mongo.DbName)
}
