package mongodb

import (
	"context"
	"os"

	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongo_uri := os.Getenv(MONGODB_URL)
	mongo_database := os.Getenv(MONGODB_USER_DB)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Conseguiu conectar ao banco de dados")
	return client.Database(mongo_database), nil
}
