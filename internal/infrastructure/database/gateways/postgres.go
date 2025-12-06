package gateways

import (
	"fmt"

	"crud/config"
	"crud/internal/infrastructure/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresGateway управляет подключением к PostgreSQL через GORM
type PostgresGateway struct {
	db *gorm.DB
}

// NewPostgresGateway создает новое подключение к PostgreSQL и выполняет миграции
func NewPostgresGateway(cfg *config.Config) (*PostgresGateway, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		cfg.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Выполняем миграции
	if err := db.AutoMigrate(
		&models.User{},
		&models.Task{},
	); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return &PostgresGateway{db: db}, nil
}

// DB возвращает экземпляр *gorm.DB
func (g *PostgresGateway) DB() *gorm.DB {
	return g.db
}

// Close закрывает подключение к базе данных
func (g *PostgresGateway) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
