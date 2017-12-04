package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config information to load from config
type Config struct {
	ListenAddress string `json:"listen_address"`
	Database      struct {
		Host      string `json:"host"`
		Name      string `json:"name"`
		User      string `json:"user"`
		Pasasword string `json:"password"`
	}
}

// LoadConfig read config from json file
func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
