package models

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Log struct {
	Id      uint      `json:"id" gorm:"primaryKey"`
	Title   string    `json:"title"`
	Time    time.Time `json:"created_at"`
	Content string    `json:"content"`
}

func ConnectDatabase() {
	dsn := "host=localhost user=jahleeljackson dbname=jahleeljackson sslmode=disable port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Log{})

	DB = database
}
