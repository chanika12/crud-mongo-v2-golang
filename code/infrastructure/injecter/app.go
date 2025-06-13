package injecter

import (
	"crud/app/commands"
	"crud/domain/repository"
	mongoDb "crud/infrastructure/db"
	"crud/infrastructure/http"
	"crud/interfaces/api"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"log"
)

func NewMongoDatabase(client *mongo.Client) *mongo.Database {
	return client.Database("orderservice")
}

func ProvideOrderRepository(db *mongo.Database) repository.OrderRepository {
	return mongoDb.NewMongoOrderRepository(db)
}

func ProvideCreateOrderHandler(repo repository.OrderRepository) *commands.CreateOrderHandler {
	return commands.NewCreateOrderHandler(repo)
}

func ProvideOrderHandler(createHandler *commands.CreateOrderHandler) *api.OrderHandler {
	return api.NewOrderHandler(createHandler)
}

func NewFxApp() *fx.App {
	return fx.New(
		fx.Provide(
			NewMongoClient,
			NewMongoDatabase,
			ProvideOrderRepository,
			ProvideCreateOrderHandler,
			ProvideOrderHandler,
			http.SetupRouter,
		),
		fx.Invoke(func(r *gin.Engine) {
			log.Println("Starting server on :8080")
			if err := r.Run(":8080"); err != nil {
				log.Fatal(err)
			}
		}),
	)
}
