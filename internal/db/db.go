package db

import (
	"go-grpc/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func GetDB() *gorm.DB {
	if database == nil{
		RefreshDB()
	}

	return database
}

func SetDB(db *gorm.DB)  {
	database = db
}

func RefreshDB()  {
	db, err := gorm.Open(sqlite.Open(config.Database.Name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database = db
}