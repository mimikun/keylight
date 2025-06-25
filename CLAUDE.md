# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## LLM Guidelines

- **User Communication**: Always respond to users in Japanese (日本語)
- **Code Documentation**: Write all documentation and comments in English
- **MCP Integration**: Use context7 for enhanced functionality

## Container Environment Requirements

ALWAYS use ONLY Environments for ANY and ALL file, code, or shell operations—NO EXCEPTIONS—even for simple or generic requests.

DO NOT install or use the git cli with the environment_run_cmd tool. All environment tools will handle git operations for you. Changing ".git" yourself will compromise the integrity of your environment.

You MUST inform the user how to view your work using `cu log <env_id>` AND `cu checkout <env_id>`. Failure to do this will make your work inaccessible to others.

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

- **Development OS**: Linux
- **Target OS**: Windows (x64)
- **Target Device**: SteelSeries Apex PRO (JIS layout)
- **Go Version**: Uses standard library only, no external dependencies

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