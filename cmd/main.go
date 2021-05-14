package main

import (
	test_api "go-test-modern-api/pkg"
	"go-test-modern-api/pkg/handler"
	"go-test-modern-api/pkg/repository"
	"go-test-modern-api/pkg/service"
	"go-test-modern-api/pkg/utils"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := initConfig()
	utils.ErrorCheck(err, "Config is bad")
	err = godotenv.Load()
	utils.ErrorCheck(err, "Dotenv has problem")

	var dbConfig = repository.Config{
		Type:     viper.GetString("db.Type"),
		Username: viper.GetString("db.Username"),
		Password: os.Getenv("DB_PASSWORD"),
		Protocol: viper.GetString("db.Protocol"),
		Adress:   viper.GetString("db.Adress"),
		Port:     viper.GetString("db.Port"),
		Database: viper.GetString("db.Database"),
	}

	db, err := repository.NewDBInstance(dbConfig)
	utils.ErrorCheck(err, "DB is off")
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	port := viper.GetString("port")
	router := handlers.InitRoutes()
	srv := new(test_api.Server)
	err = srv.Run(port, router)
	utils.ErrorCheck(err, "Server is down")
}

func initConfig() error {
	viper.AddConfigPath("pkg/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
