package database

import (
	"fmt"
	"go-fiber-todos/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err.Error())
	}

	DBConn = db

	db.AutoMigrate(&models.Todo{})
}
