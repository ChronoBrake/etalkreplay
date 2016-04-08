package main

import (
	"encoding/json"
	"github.com/1lann/lol-replay/record"
	"log"
	"os"
)

type configuration struct {
	Players             []player `json:"players"`
	RecordingsDirectory string   `json:"recordings_directory"`
	BindAddress         string   `json:"bind_address"`
	RiotAPIKey          string   `json:"riot_api_key"`
	RefreshRate         int      `json:"refresh_rate_seconds"`
	KeepNumRecordings   int      `json:"keep_num_recordings"`
	ShowPerPage         int      `json:"show_per_page"`
	ShowReplayPortAs    int      `json:"show_replay_port_as"`
}

var config configuration

func readConfiguration(location string) {
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	dec := json.NewDecoder(file)

	err = dec.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	for _, player := range config.Players {
		if !record.IsValidPlatform(player.Platform) {
			log.Fatal(player.ID + "'s platform " + player.Platform +
				" is not a valid platform.")
		}
	}
}