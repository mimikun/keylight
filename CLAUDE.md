# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## LLM Guidelines

- **User Communication**: Always respond to users in Japanese (日本語)
- **Code Documentation**: Write all documentation and comments in English
- **MCP Integration**: Use context7 for enhanced functionality

## Project Overview

keylight is a CLI tool that controls SteelSeries Apex PRO gaming keyboard LEDs to provide visual notifications for build completion status. The tool communicates with the SteelSeries GameSense API via HTTP REST calls.

## Build Commands

```bash
# Build for Windows x64 (primary target platform)
GOOS=windows GOARCH=amd64 go build -o keylight.exe

# Build for current platform (development)
go build -o keylight

# Run tests
go test ./...

# Run tests for specific package
go test ./internal/steelseries
```

## Architecture

The project follows a layered architecture:

- **CLI Layer** (`main.go`): Entry point and command line argument processing
- **Command Layer** (`internal/cli/`): Command parsing and execution logic
- **SteelSeries Client** (`internal/steelseries/`): GameSense API client implementation
- **GameSense HTTP API**: Local HTTP API provided by SteelSeries Engine

### Key Components

- `internal/steelseries/client.go`: Core GameSense API client with methods for game registration, event binding, and LED control
- `internal/steelseries/config.go`: Handles reading SteelSeries Engine configuration from `%PROGRAMDATA%/SteelSeries/SteelSeries Engine 3/coreProps.json`
- `internal/steelseries/patterns.go`: LED pattern definitions for success (green circle) and failure (red cross) indicators
- `internal/steelseries/types.go`: Type definitions for GameSense API requests/responses

### Communication Flow

1. Read `coreProps.json` to get GameSense API address
2. Register game "KEYLIGHT" with GameSense API
3. Bind event handler for "BUILD_STATUS" event
4. Send LED control event with bitmap pattern
5. Auto-off after 3-5 seconds

## LED Pattern System

The keyboard uses a 132-key bitmap array where each key is represented by RGB values `[R, G, B]`. Patterns are defined for:

- **Success Pattern**: Green circle using keys 5, 6, 7, 8, R, I, D, K, C, M, SPACE
- **Failure Pattern**: Red cross using keys 5, 8, T, Y, H, N, C

## Development Environment

- **Development OS**: Linux (WSL2)
- **Target OS**: Windows (x64)
- **Target Device**: SteelSeries Apex PRO (JIS layout)
- **Go Version**: Uses standard library only, no external dependencies

### WSL Development Setup

When developing on WSL, you can directly access Windows files:

```bash
# Check SteelSeries Engine configuration
cat /mnt/c/ProgramData/SteelSeries/SteelSeries\ Engine\ 3/coreProps.json

# Alternative path if user profile is needed
cat /mnt/c/Users/mimikun/AppData/Local/SteelSeries/SteelSeries\ Engine\ 3/coreProps.json
```

This allows real-time verification of GameSense API endpoint without running on Windows.

## CLI Usage

```bash
# Display success pattern
keylight --success

# Display failure pattern  
keylight --failure
```

## Error Handling

The tool handles common scenarios:
- SteelSeries Engine not running
- Configuration file not found
- Network connectivity issues with GameSense API

Exit codes: 0 for success, 1 for error.

## Development Methodology

This project follows Test-Driven Development (TDD) practices:

### TDD Workflow

1. **Red**: Write a failing test that describes the desired behavior
2. **Green**: Write the minimal code to make the test pass
3. **Refactor**: Clean up the code while keeping tests passing

### Testing Guidelines

- Write tests before implementing functionality
- Each function should have corresponding unit tests
- Use table-driven tests for multiple scenarios
- Mock external dependencies (HTTP API calls, file system operations)
- Aim for high test coverage, especially for core business logic

### Test Organization

- Test files follow Go convention: `*_test.go`
- Place tests in the same package as the code under test
- Use `testify` assertions if needed, but prefer standard Go testing
- Group related tests using subtests (`t.Run()`)

### Integration Testing

- Test the complete flow from CLI to LED pattern display
- Mock SteelSeries GameSense API responses
- Verify error handling scenarios
- Test configuration file parsing
