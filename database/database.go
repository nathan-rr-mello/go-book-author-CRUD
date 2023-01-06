package database

import (
	"fmt"

	"github.com/nathan-rr-mello/go-book-author-CRUD/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to establish connection with database")
	}
	fmt.Println("Database connected...")
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.Book{}, &models.Author{})
	Db = db
}
