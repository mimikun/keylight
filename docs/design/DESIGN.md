# keylight - Philips Hue Notification Tool Design Document

## Overview

### Purpose

A CLI tool to control Philips Hue desk lights for visual notification of long-running task completion status such as Linux builds.

### Key Features

- Build success: Activate green success scene on Hue lights
- Build failure: Activate red failure scene on Hue lights
- Scene state restoration: Capture and restore original lighting state
- Simple CLI interface

### Technology Stack

- **Language**: Go (standard library only)
- **Development Environment**: Linux
- **Runtime Environment**: Windows (x64)
- **Target Device**: Philips Hue Bridge and Hue desk lights
- **Communication Protocol**: HTTP REST (Hue Bridge API)

## Documentation Structure

This design document is organized into several focused documents:

### [ARCHITECTURE.md](./ARCHITECTURE.md)
- System layer structure and module organization
- Scene management and state restoration design
- Detailed component design specifications
- Error handling strategies
- Implementation considerations (performance, compatibility, maintainability)
- Build and deployment procedures
- Future extensibility planning
- Security considerations

### [API.md](./API.md)
- Philips Hue Bridge API specification and endpoints
- Configuration and discovery mechanisms
- Communication flow and protocols
- Request/response specifications with examples
- Error handling and HTTP status codes
- Authentication and security

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
# Activate success scene
keylight --success

# Activate failure scene
keylight --failure

# Initialize/verify required scenes
keylight --init-scenes

# Migrate configuration format
keylight --migrate-config source.json target.yaml
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
- **Dependencies**: Philips Hue Bridge on local network
- **Hardware**: Philips Hue Bridge and Hue desk lights
- **Configuration**: Supports JSON, YAML, TOML formats

## Project Structure

```
keylight/
├── main.go                   # Entry point
├── internal/
│   ├── hue/                 # Hue Bridge API client
│   │   ├── client.go        # Bridge communication
│   │   ├── config.go        # Multi-format configuration
│   │   ├── scenes.go        # Scene management & auto-creation
│   │   └── types.go         # Type definitions
│   └── cli/                 # Command line processing
├── docs/
│   └── design/             # Design documentation
├── integration/            # Integration tests
├── examples/               # Configuration examples
│   ├── keylight.json       # JSON format example
│   ├── keylight.yaml       # YAML format example
│   └── keylight.toml       # TOML format example
└── CLAUDE.md              # Development guidelines
```

## Configuration Examples

### JSON Format (Default)
```json
{
  "bridge_ip": "192.168.1.100",
  "username": "your-api-key",
  "scenes": {
    "default_scene": "Default_State",
    "success_scene": "Success_Notification",
    "failure_scene": "Failure_Notification"
  },
  "auto_create_scenes": true
}
```

### YAML Format
```yaml
bridge_ip: 192.168.1.100
username: your-api-key
scenes:
  default_scene: Default_State
  success_scene: Success_Notification
  failure_scene: Failure_Notification
auto_create_scenes: true
```

### TOML Format
```toml
bridge_ip = "192.168.1.100"
username = "your-api-key"
auto_create_scenes = true

[scenes]
default_scene = "Default_State"
success_scene = "Success_Notification"
failure_scene = "Failure_Notification"
```

## Configuration Management Features

- **Format Auto-Detection**: Based on file extension
- **Configuration Migration**: Convert between formats seamlessly
- **Scene Auto-Creation**: Automatically creates required scenes
- **Fallback Strategy**: Default scene restoration when state capture fails

For detailed information about each aspect of the system, please refer to the appropriate specialized document above.