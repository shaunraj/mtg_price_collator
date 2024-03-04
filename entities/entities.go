package entities

import (
	entity "collator/entities/db"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDb() {
	db, err := gorm.Open(sqlite.Open("mtg_cards.db"))
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&entity.CardMetaData{}, &entity.Price{})
}
