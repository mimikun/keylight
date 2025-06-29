package hue

import (
	"io/ioutil"
	"os"
	"testing"
)

// Red: Failing test for JSON config loading
func TestLoadConfig_JSON(t *testing.T) {
	// Create temporary JSON config file
	jsonContent := `{
		"bridge_ip": "192.168.1.100",
		"username": "test-api-key",
		"scenes": {
			"default_scene": "Default_State",
			"success_scene": "Success_Notification",
			"failure_scene": "Failure_Notification"
		},
		"auto_create_scenes": true
	}`
	
	tmpfile, err := ioutil.TempFile("", "keylight-test-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	
	if _, err := tmpfile.Write([]byte(jsonContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	
	config, err := LoadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	if config.BridgeIP != "192.168.1.100" {
		t.Errorf("Expected BridgeIP '192.168.1.100', got '%s'", config.BridgeIP)
	}
	
	if config.Username != "test-api-key" {
		t.Errorf("Expected Username 'test-api-key', got '%s'", config.Username)
	}
	
	if config.Scenes.DefaultScene != "Default_State" {
		t.Errorf("Expected DefaultScene 'Default_State', got '%s'", config.Scenes.DefaultScene)
	}
	
	if !config.AutoCreateScenes {
		t.Error("Expected AutoCreateScenes to be true")
	}
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	_, err := LoadConfig("nonexistent.json")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestLoadConfig_InvalidJSON(t *testing.T) {
	// Create temporary invalid JSON file
	invalidContent := `{
		"bridge_ip": "192.168.1.100",
		"username": "test-api-key"
		// missing comma
		"auto_create_scenes": true
	}`
	
	tmpfile, err := ioutil.TempFile("", "keylight-test-invalid-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	
	if _, err := tmpfile.Write([]byte(invalidContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	
	_, err = LoadConfig(tmpfile.Name())
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}