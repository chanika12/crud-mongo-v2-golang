package injecter

import (
	"crud/domain/repository"
	mongoDb "crud/infrastructure/db"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideOrderRepository(db *mongo.Database) repository.IOrderRepository {
	return mongoDb.NewMongoOrderRepository(db)
}

var RepositoryProviderSet = wire.NewSet(
	SetupMongoDB,
	ProvideOrderRepository,
)
