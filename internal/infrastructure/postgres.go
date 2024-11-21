package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"task-plan/internal/domain"
)

const (
	tasksTable    = "tasks"
	usersTable    = "users"
	commandsTable = "commands"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.Task{}, &domain.User{}, &domain.Group{}, &domain.Token{})
	return db, err
}
