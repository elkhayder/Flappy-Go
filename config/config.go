package config

import (
	"encoding/json"
	"log"
	"os"
)

const ConfigFilename = "config.json"

type SoundConfig struct {
	Enabled bool    `json:"enabled"`
	Volume  float64 `json:"volume"`
}

var Config = struct {
	Fx    SoundConfig `json:"fx"`
	Music SoundConfig `json:"music"`
}{
	Fx: SoundConfig{
		Enabled: true,
		Volume:  .7,
	},
	Music: SoundConfig{
		Enabled: true,
		Volume:  .3,
	},
}

func LoadConfig() {
	configFile, err := os.Open(ConfigFilename)

	var (
		jsonParser *json.Decoder
	)

	if err != nil {
		return
	}

	defer configFile.Close()

	jsonParser = json.NewDecoder(configFile)

	if err := jsonParser.Decode(&Config); err != nil {
		log.Fatal(err)
	}
}

func SaveConfig() error {
	configFile, err := os.Create(ConfigFilename)
	if err != nil {
		return err
	}
	defer configFile.Close()

	jsonEncoder := json.NewEncoder(configFile)

	if err := jsonEncoder.Encode(&Config); err != nil {
		return err
	}

	return nil
}
