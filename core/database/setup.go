package database

import (
	"discord-sentinel/core/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// buildDSN generates a string with corresponding connection
// data provided by configuration.
func buildDSN(cfg *config.Database) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)
}

// SetupDatabaseConnection starts database connection handshake and ORM parametrization.
func SetupDatabaseConnection(cfg *config.Database) (*gorm.DB, error) {

	dsn := buildDSN(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	hDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxLifetime := time.Duration(cfg.MaxLifetime) * time.Second

	hDb.SetConnMaxLifetime(maxLifetime)
	hDb.SetMaxIdleConns(cfg.MaxIdle)
	hDb.SetMaxOpenConns(cfg.MaxConnections)

	return db, nil

}
