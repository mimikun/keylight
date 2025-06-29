# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## LLM Guidelines

- **User Communication**: Always respond to users in Japanese (日本語)
- **Code Documentation**: Write all documentation and comments in English
- **MCP Integration**: Use context7 for enhanced functionality

## Project Overview

keylight is a CLI tool that controls Philips Hue desk lights to provide visual notifications for build completion status. The tool communicates with the Philips Hue Bridge API via HTTP REST calls to trigger predefined scenes.

## Build Commands

```bash
# Build for Windows x64 (primary target platform)
GOOS=windows GOARCH=amd64 go build -o keylight.exe

# Build for current platform (development)
go build -o keylight

# Run tests
go test ./...

# Run tests for specific package
go test ./internal/hue
```

## Architecture

The project follows a layered architecture:

- **CLI Layer** (`main.go`): Entry point and command line argument processing
- **Command Layer** (`internal/cli/`): Command parsing and execution logic
- **Hue Client** (`internal/hue/`): Philips Hue Bridge API client implementation
- **Philips Hue Bridge API**: Local HTTP API provided by Hue Bridge

### Key Components

- `internal/hue/client.go`: Core Hue Bridge API client with methods for scene management and light control
- `internal/hue/config.go`: Handles Hue Bridge discovery and configuration management
- `internal/hue/scenes.go`: Scene definitions for success and failure notifications
- `internal/hue/types.go`: Type definitions for Hue API requests/responses

### Communication Flow

1. Discover Hue Bridge on local network or use configured IP
2. Authenticate with Hue Bridge using stored API key
3. Capture current scene state for restoration
4. Activate predefined scene ("Success" or "Failure")
5. Wait for 10 seconds
6. Restore original scene state

## Scene System

The tool uses predefined Philips Hue scenes for visual notifications:

- **Success Scene**: Green lighting pattern for successful build completion
- **Failure Scene**: Red lighting pattern for failed build completion

## Development Environment

- **Development OS**: Linux (WSL2)
- **Target OS**: Windows (x64)
- **Target Device**: Philips Hue Bridge and Hue desk lights
- **Go Version**: Uses standard library only, no external dependencies

### Development Setup

The Hue Bridge is accessible from both Linux and Windows environments via local network:

```bash
# Discover Hue Bridge on local network
curl -X GET "https://discovery.meethue.com/"

# Test connection to bridge (replace <bridge_ip> with actual IP)
curl -X GET "http://<bridge_ip>/api/<username>/scenes"
```

This allows development and testing from any environment with network access to the Hue Bridge.

## CLI Usage

```bash
# Display success scene
keylight --success

# Display failure scene
keylight --failure
```

## Error Handling

The tool handles common scenarios:
- Hue Bridge not accessible
- Authentication failures
- Scene not found
- Network connectivity issues with Hue Bridge API

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

- Test the complete flow from CLI to Hue scene activation
- Mock Philips Hue Bridge API responses
- Verify error handling scenarios
- Test scene state capture and restoration

## Git Conventions

This project follows Conventional Commits specification for commit messages:

### Commit Message Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Commit Types

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect the meaning of the code (white-space, formatting, etc)
- **refactor**: A code change that neither fixes a bug nor adds a feature
- **perf**: A code change that improves performance
- **test**: Adding missing tests or correcting existing tests
- **build**: Changes that affect the build system or external dependencies
- **ci**: Changes to CI configuration files and scripts
- **chore**: Other changes that don't modify src or test files

### Scopes

Use scopes to indicate the area of the codebase being modified:

- **cli**: Command line interface and argument parsing
- **client**: Philips Hue Bridge API client
- **config**: Configuration file handling
- **scenes**: Scene definitions and management
- **claude**: CLAUDE.md file changes
- **docs**: Other documentation changes (README, etc.)
- **build**: Build configuration and scripts
- **test**: Test-related changes

### Examples

```
feat(cli): add success scene activation
fix(config): handle Hue Bridge discovery failure
docs(claude): add TDD guidelines and development setup
docs(readme): update installation instructions
test(client): add unit tests for Hue Bridge client
refactor(scenes): simplify scene state management
perf(client): optimize HTTP request handling
```
