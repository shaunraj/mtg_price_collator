package card

import (
	entity "collator/entities/db"
	"collator/parsers/scryfall"

	"github.com/google/uuid"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Insert(cards []scryfall.ScryfallCard) {
	db, err := gorm.Open(sqlite.Open("mtg_cards.db"))
	var metadatas []entity.CardMetaData

	for _, card := range cards {
		price_uuid := uuid.New()
		metadata := entity.CardMetaData{
			UUID:      uuid.New(),
			PriceUUID: price_uuid,
			Name:      card.Name,
			Lang:      card.Lang,
			Prices: entity.Price{
				Usd:       card.Prices.Usd,
				UsdFoil:   card.Prices.UsdFoil,
				UsdEtched: card.Prices.UsdEtched,
				Eur:       card.Prices.Eur,
				EurFoil:   card.Prices.EurFoil,
				Tix:       card.Prices.Tix,
				UUID:      price_uuid,
			},
		}
		metadatas = append(metadatas, metadata)
	}

	if err != nil {
		panic("failed to connect to database")
	}
	db.Create(metadatas)
}
