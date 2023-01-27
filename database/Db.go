package database

import (
	"app/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var err error
var (
	DBConn *gorm.DB
)

func DbConnection() {

	DBConn, err = gorm.Open(sqlite.Open("todo.db"))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func SyncDb() {
	DBConn.AutoMigrate(&models.ToDo{})
}
