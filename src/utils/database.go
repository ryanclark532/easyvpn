package utils

import (
	"easyvpn/src/dtos"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDB(databaseURL string) error {
	var err error
	db, err = gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return HandleError(err.Error(), "InitializeDB")
	}

	err = db.AutoMigrate(&dtos.User{})
	if err != nil {
		return HandleError(err.Error(), "migrate")
	}

	return nil
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, HandleError("DB not initialized", "GetDB")
	}
	return db, nil
}
