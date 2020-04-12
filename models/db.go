package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
}

func Sync() {
	// Sync models to the database
	DB.DropTable(&Author{}, &Book{})
	DB.AutoMigrate(&Author{}, &Book{})
}
