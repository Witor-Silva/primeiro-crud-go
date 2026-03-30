package repository

import (
	"context"
	"os"

	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/logger"
	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/rest_err"
	"github.com/Witor-Silva/primeiro-crud-go/src/model"
	"github.com/Witor-Silva/primeiro-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init CreateUser repository")
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error(
			"Error trying to create user", err, zap.String("Journey", "CreateUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info(
		"User repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("Journey", "CreateUser"))

	return converter.ConvertEntityToDomain(*value), nil

}
