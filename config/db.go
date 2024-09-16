package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() (*gorm.DB, error) {
	logger := GetLogger("Database")

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres connection error: %s", err.Error())
		return db, err
	}
	logger.Info("Connection established")

	logger.Info("Migrating database...")
	err = db.AutoMigrate()
	if err != nil {
		logger.Errorf("postgres migration error: %s", err.Error())
		return db, err
	}
	logger.Info("Database migrated")

	return db, nil
}
