package config

import (
	"fmt"

	"festech.de/rmm/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = fmt.Sprintf("root:%s@tcp(%s:3306)/rmm?charset=utf8mb4&parseTime=True&loc=Local", Getenv("PW", ""), Getenv("HOST", "localhost"))

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&models.Device{}, &models.SystemInfo{})

	return nil
}
