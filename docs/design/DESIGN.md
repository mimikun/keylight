# keylight - SteelSeries LED Notification Tool Design Document

## Overview

### Purpose

A CLI tool to control SteelSeries Apex PRO gaming keyboard LEDs for visual notification of long-running task completion status such as Linux builds.

### Key Features

- Build success: Display green checkmark (circle) pattern on keyboard
- Build failure: Display red cross pattern on keyboard
- Simple CLI interface

### Technology Stack

- **Language**: Go (standard library only)
- **Development Environment**: Linux
- **Runtime Environment**: Windows (x64)
- **Target Device**: SteelSeries Apex Pro JP (Product No. 64629) - Full-size JIS layout
- **Communication Protocol**: HTTP REST (GameSense API)

## Documentation Structure

This design document is organized into several focused documents:

### [ARCHITECTURE.md](./ARCHITECTURE.md)
- System layer structure and module organization
- LED pattern design and bitmap implementation
- Detailed component design specifications
- Error handling strategies
- Implementation considerations (performance, compatibility, maintainability)
- Build and deployment procedures
- Future extensibility planning
- Security considerations

### [API.md](./API.md)
- GameSense API specification and endpoints
- Configuration file format and location
- Communication flow and protocols
- Request/response specifications with examples
- Error handling and HTTP status codes
- Rate limiting and API versioning

### [DEVELOPMENT.md](./DEVELOPMENT.md)
- Test-Driven Development (TDD) methodology
- Testing guidelines and structure
- Git conventions (Conventional Commits)
- Development environment setup (WSL, cross-platform)
- Code quality standards and best practices
- Build commands and testing procedures

## Quick Start

### CLI Usage

```bash
# Display success pattern
keylight --success

# Display failure pattern  
keylight --failure
```

### Build Commands

```bash
# Build for Windows x64 (primary target platform)
GOOS=windows GOARCH=amd64 go build -o keylight.exe

# Build for current platform (development)
go build -o keylight
```

### Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...
```

## System Requirements

- **Development**: Linux (WSL2 recommended)
- **Runtime**: Windows 10/11 (x64)
- **Dependencies**: SteelSeries Engine 3 or SteelSeries GG
- **Hardware**: SteelSeries Apex Pro JP keyboard (Product No. 64629) - Full-size JIS layout with numpad

## Project Structure

```
keylight/
├── main.go                   # Entry point
├── internal/
│   ├── steelseries/         # GameSense API client
│   └── cli/                 # Command line processing
├── docs/
│   └── design/             # Design documentation
├── integration/            # Integration tests
└── CLAUDE.md              # Development guidelines
```

For detailed information about each aspect of the system, please refer to the appropriate specialized document above.