package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Price struct {
	gorm.Model
	UUID      uuid.UUID
	Usd       float64 `json:"usd,string"`
	UsdFoil   float64 `json:"usd_foil,string"`
	UsdEtched float64 `json:"usd_etched,string"`
	Eur       float64 `json:"eur,string"`
	EurFoil   float64 `json:"eur_foil,string"`
	Tix       float64 `json:"tix,string"`
}
