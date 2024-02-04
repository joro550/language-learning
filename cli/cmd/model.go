package cmd

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

type Card struct {
	gorm.Model
	English string `gorm:"column:engilsh"`
	German  string `gorm:"column:german"`
	Id      uint   `gorm:"primaryKey"`
}

func NewCard(english, german string) Card {
	return Card{English: english, German: german}
}

func ConnectToDatabase() (*gorm.DB, error) {
	DbConnection, err := gorm.Open(sqlite.Open("cards.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DbConnection.AutoMigrate(&Card{})

	return DbConnection, nil
}
