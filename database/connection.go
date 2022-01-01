package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(HelloTable{})
	if err != nil {
		return err
	}
	fmt.Println("DB's Migrated!")
	return nil
}

func ConnectDB(DB *gorm.DB, DBName string) (*gorm.DB, error) {
	var err error

	DB, err = gorm.Open(sqlite.Open(DBName),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	err = autoMigrate(DB)
	if err != nil {
		return nil, err
	}
	return DB, nil
}
