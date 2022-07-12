package main

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

type Config struct {
    AuthToken *string `json:"token,omitempty"`
}

func NewConfig(config_file_path string) (*Config, error) {
    home, _ := os.UserHomeDir()
    file_uri := path.Join(home, ".config", config_file_path)
    content, err := os.ReadFile(file_uri)
    if err != nil {
        return nil, errors.New("Config file not found")
    }
    config := new(Config)
    json.Unmarshal([]byte(content), config)

    return config, nil
}
