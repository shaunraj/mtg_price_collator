package scryfall

import (
	"encoding/json"
	"io"
)

func DecodeStream(reader io.Reader, channel chan []ScryfallCard, batch_size int) error {
	if batch_size == 0 {
		batch_size = 1000
	}

	var entries []ScryfallCard
	decoder := json.NewDecoder(reader)
	decoder.Token()
	for decoder.More() {
		var entry ScryfallCard
		decoder.Decode(&entry)
		entries = append(entries, entry)
		if len(entries) == batch_size {
			sendToChannel(channel, &entries)
		}
	}
	//This is here because we want to send any remaining messages after file processing is done to channel
	sendToChannel(channel, &entries)
	decoder.Token()
	close(channel)
	return nil
}

func sendToChannel(channel chan []ScryfallCard, entries *[]ScryfallCard) {
	channel <- *entries
	*entries = nil
}
