package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Books struct {
	ID        int    `json:"id" gorm:"primary key;autoincrement"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	if err != nil {
		fmt.Println("failed to migrate db", err)
	}
}
