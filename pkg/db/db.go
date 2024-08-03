package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Connection struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"unique;not null"`
	Address string `gorm:"not null"`
}

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("connections.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&Connection{})
}
