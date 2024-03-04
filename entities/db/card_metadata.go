package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CardMetaData struct {
	gorm.Model
	Name      string `json:name`
	Lang      string `json:lang`
	Prices    Price  `gorm:"foreignKey:PriceUUID;references:uuid"`
	PriceUUID uuid.UUID
	UUID      uuid.UUID
}
