package config

import (
	"festech.de/rmm/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = Getenv("DATABASE_URI", "")

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	err = Database.AutoMigrate(&models.Device{}, &models.SystemInfo{}, &models.User{}, &models.DeviceToken{})

	if err != nil {
		panic(err)
	}

	return nil
}
