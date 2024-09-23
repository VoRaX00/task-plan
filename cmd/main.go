package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"task-plan/internal/application"
	"task-plan/internal/handler"
	"task-plan/internal/infrastructure"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDbConfig() infrastructure.Config {
	return infrastructure.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}

	cfg := initDbConfig()
	db, err := infrastructure.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("Error initializing db: %v", err)
	}

	defer func() {
		_ = db.Close()
	}()

	repos := infrastructure.NewRepository()
	services := application.NewService(repos)
	handlers := handler.NewHandler(services)

}
