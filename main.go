package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	address, err := getGameSenseAddress()
	if err != nil {
		fmt.Printf("Failed to get GameSense address: %v\n", err)
		os.Exit(1)
	}
	
	// Register game
	if err := registerGame(address); err != nil {
		fmt.Printf("Failed to register game: %v\n", err)
		os.Exit(1)
	}
	
	// Bind event handler
	if err := bindEvent(address); err != nil {
		fmt.Printf("Failed to bind event: %v\n", err)
		os.Exit(1)
	}
	
	// Turn off all keys first
	if err := turnOffAllKeys(address); err != nil {
		fmt.Printf("Failed to turn off keys: %v\n", err)
		os.Exit(1)
	}
	
	// Light up HJKL keys
	if err := lightHJKLKeys(address); err != nil {
		fmt.Printf("Failed to light HJKL keys: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("HJKL keys lit for 10 seconds...")
	time.Sleep(10 * time.Second)
	
	// Turn off all keys again
	if err := turnOffAllKeys(address); err != nil {
		fmt.Printf("Failed to turn off keys: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("Done!")
}

func registerGame(address string) error {
	gameData := map[string]interface{}{
		"game":         "KEYLIGHT",
		"game_display_name": "Keylight",
		"developer":    "mimikun",
	}
	
	return sendPostRequest(address, "/game_metadata", gameData)
}

func bindEvent(address string) error {
	bindData := map[string]interface{}{
		"game": "KEYLIGHT",
		"event": "KEYBOARD_CONTROL",
		"handlers": []map[string]interface{}{
			{
				"device-type": "keyboard",
				"zone": "all",
				"mode": "color",
			},
		},
	}
	
	return sendPostRequest(address, "/bind_game_event", bindData)
}

func turnOffAllKeys(address string) error {
	// Create bitmap with all keys off (black)
	bitmap := make([][]int, 132)
	for i := range bitmap {
		bitmap[i] = []int{0, 0, 0} // RGB black
	}
	
	eventData := map[string]interface{}{
		"game": "KEYLIGHT",
		"event": "KEYBOARD_CONTROL",
		"data": map[string]interface{}{
			"bitmap": bitmap,
		},
	}
	
	return sendPostRequest(address, "/game_event", eventData)
}

func lightHJKLKeys(address string) error {
	// Create bitmap with all keys lit up in red for testing
	bitmap := make([][]int, 132)
	for i := range bitmap {
		bitmap[i] = []int{255, 0, 0} // RGB red - light up all keys to test
	}
	
	eventData := map[string]interface{}{
		"game": "KEYLIGHT",
		"event": "KEYBOARD_CONTROL",
		"data": map[string]interface{}{
			"bitmap": bitmap,
		},
	}
	
	return sendPostRequest(address, "/game_event", eventData)
}

func sendPostRequest(address, endpoint string, data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}
	
	url := fmt.Sprintf("http://%s%s", address, endpoint)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}
	
	return nil
}

func getGameSenseAddress() (string, error) {
	configPath := "C:\\ProgramData\\SteelSeries\\SteelSeries Engine 3\\coreProps.json"
	
	file, err := os.Open(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to open coreProps.json: %v", err)
	}
	defer file.Close()
	
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read coreProps.json: %v", err)
	}
	
	var config struct {
		Address           string `json:"address"`
		EncryptedAddress  string `json:"encryptedAddress"`
		GGEncryptedAddress string `json:"ggEncryptedAddress"`
	}
	
	if err := json.Unmarshal(data, &config); err != nil {
		return "", fmt.Errorf("failed to parse coreProps.json: %v", err)
	}
	
	if config.Address == "" {
		return "", fmt.Errorf("address not found in coreProps.json")
	}
	
	return config.Address, nil
}
