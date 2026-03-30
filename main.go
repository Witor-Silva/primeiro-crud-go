package main

import (
	"context"
	"log"

	mongodb "github.com/Witor-Silva/primeiro-crud-go/infra/database"
	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/logger"
	"github.com/Witor-Silva/primeiro-crud-go/src/controlller/routes/controller"
	"github.com/Witor-Silva/primeiro-crud-go/src/controlller/routes/controller/routes"
	"github.com/Witor-Silva/primeiro-crud-go/src/model/repository"
	"github.com/Witor-Silva/primeiro-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Aboiut to start user application")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())

	if err != nil {
		log.Fatalf("Error connecting to database, error=%s \n", err.Error())
		return
	}

	//Init Dependecies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
