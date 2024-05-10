package main

import (
	"log"

	"github.com/isido5ik/StoryPublishingPlatform/configs"
	"github.com/isido5ik/StoryPublishingPlatform/internal/delivery/http"
	"github.com/isido5ik/StoryPublishingPlatform/internal/repository"
	"github.com/isido5ik/StoryPublishingPlatform/internal/usecase"
	"github.com/isido5ik/StoryPublishingPlatform/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

//@title Story Publishing Platform
//@version 1.0
//@description Rest Api Server

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := configs.InitConfigs(); err != nil {
		logrus.WithError(err).Fatal("error initializing configs")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db := server.InitDB()
	repo := repository.NewRepository(db)
	usecases := usecase.NewUsecase(repo)
	handlers := http.NewHandler(usecases)

	server := server.NewServer()
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.WithError(err).Fatal("error occured while running http server")
	}
}
