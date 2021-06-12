package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "../db.sqlite3")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Todo{})

	DB = database
}
