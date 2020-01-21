package config

import (
	"fmt"
	"os"
	"strconv"
)

type Settings struct {
	DataFilePath string
	Host string
	Port int
}

func NewSettings() (*Settings, error) {
	dataFilePath := os.Getenv("DATA_FILE_PATH")
	if dataFilePath == "" {
		dataFilePath = "./data/sources/"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "888"
	}

	nport, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("unable to parse PORT: %v", err)
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	settings := &Settings{dataFilePath, host, nport}

	return settings, nil
}
