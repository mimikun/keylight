package hue

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	BridgeIP         string      `json:"bridge_ip"`
	Username         string      `json:"username"`
	Scenes           SceneConfig `json:"scenes"`
	AutoCreateScenes bool        `json:"auto_create_scenes"`
}

type SceneConfig struct {
	DefaultScene string `json:"default_scene"`
	SuccessScene string `json:"success_scene"`
	FailureScene string `json:"failure_scene"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()
	
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode JSON config: %w", err)
	}
	
	return &config, nil
}