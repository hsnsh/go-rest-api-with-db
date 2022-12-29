package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DbName   string
	SslMode  string
}

func PostgresNewConnectionWithConfig(config *PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Europe/Istanbul",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DbName,
		config.SslMode,
	)

	dial := postgres.Open(dsn)

	db, conErr := newConnection(dial)
	if conErr != nil {
		return db, conErr
	}

	return db, nil
}
