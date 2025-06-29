package cli

import (
	"testing"
)

// Red: First failing test for CLI argument parsing
func TestParseArgs_Success(t *testing.T) {
	args := []string{"--success"}
	cmd, err := ParseArgs(args)
	
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	if cmd.Action != "success" {
		t.Errorf("Expected action 'success', got '%s'", cmd.Action)
	}
}

func TestParseArgs_Failure(t *testing.T) {
	args := []string{"--failure"}
	cmd, err := ParseArgs(args)
	
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	if cmd.Action != "failure" {
		t.Errorf("Expected action 'failure', got '%s'", cmd.Action)
	}
}

func TestParseArgs_InitScenes(t *testing.T) {
	args := []string{"--init-scenes"}
	cmd, err := ParseArgs(args)
	
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	if cmd.Action != "init-scenes" {
		t.Errorf("Expected action 'init-scenes', got '%s'", cmd.Action)
	}
}

func TestParseArgs_NoArgs(t *testing.T) {
	args := []string{}
	_, err := ParseArgs(args)
	
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}

func TestParseArgs_InvalidArg(t *testing.T) {
	args := []string{"--invalid"}
	_, err := ParseArgs(args)
	
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}