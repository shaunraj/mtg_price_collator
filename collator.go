package main

import (
	"collator/entities"
	"collator/entities/objects/card"
	"collator/parsers/scryfall"
	"os"
)

func main() {
	file, _ := os.Open("all-cards-20240303101547.json")
	channel := make(chan []scryfall.ScryfallCard)
	entities.InitializeDb()
	go scryfall.DecodeStream(file, channel, 1000)
	for card_data := range channel {
		card.Insert(card_data)
	}
}
