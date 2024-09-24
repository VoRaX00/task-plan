package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
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

	server := new(Server)

	go func() {
		if err = server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error starting server: %v", err)
		}
	}()

	logrus.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("Shutting down server...")
	if err = server.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("Error shutting down server: %v", err)
	}

	if err = db.Close(); err != nil {
		logrus.Fatalf("Error closing db: %v", err)
	}

	logrus.Info("Server stopped")
}
