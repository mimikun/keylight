# keylight - SteelSeries LED Notification Tool Design Document

## 1. System Overview

### 1.1 Purpose

A CLI tool to control SteelSeries Apex PRO gaming keyboard LEDs for visual notification of long-running task completion status such as Linux builds.

### 1.2 Key Features

- Build success: Display green checkmark (circle) pattern on keyboard
- Build failure: Display red cross pattern on keyboard
- Simple CLI interface

### 1.3 Technology Stack

- **Language**: Go (standard library only)
- **Development Environment**: Linux
- **Runtime Environment**: Windows (x64)
- **Target Device**: SteelSeries Apex PRO (JIS layout)
- **Communication Protocol**: HTTP REST (GameSense API)

## 2. System Architecture

### 2.1 Layer Structure

```
┌─────────────────────────────────────┐
│         CLI Layer (main.go)         │
├─────────────────────────────────────┤
│      Command Layer (cli/)           │
├─────────────────────────────────────┤
│   SteelSeries Client (steelseries/) │
├─────────────────────────────────────┤
│      GameSense HTTP API             │
└─────────────────────────────────────┘
```

### 2.2 Module Structure

```
keylight/
├── main.go                   # Entry point
├── internal/
│   ├── steelseries/
│   │   ├── client.go        # GameSense API client
│   │   ├── config.go        # Configuration file handler
│   │   ├── patterns.go      # LED pattern definitions
│   │   └── types.go         # Type definitions
│   └── cli/
│       └── commands.go      # Command line processing
├── go.mod
├── go.sum
├── README.md
└── docs/
    └── design/
        └── DESIGN.md
```

## 3. GameSense API Specification

### 3.1 API Overview

Uses local HTTP API provided by SteelSeries Engine. No authentication required.

### 3.2 Configuration File

- **Windows**: `%PROGRAMDATA%/SteelSeries/SteelSeries Engine 3/coreProps.json`
- **Format**:

```json
{
  "address": "127.0.0.1:12345"
}
```

### 3.3 Main Endpoints

1. **Game Registration**: `POST http://{address}/game_metadata`
2. **Event Binding**: `POST http://{address}/bind_game_event`
3. **Event Sending**: `POST http://{address}/game_event`
4. **Heartbeat**: `POST http://{address}/game_heartbeat`

### 3.4 Communication Flow

```
1. Read coreProps.json → Get API address
2. Register game (KEYLIGHT)
3. Bind event handler
4. Send LED control event
5. Auto-off after 3-5 seconds
```

## 4. LED Pattern Design

### 4.1 Keyboard Layout (JIS)

Each key corresponds to an index in the bitmap array. Managed with 132-key array.

### 4.2 Pattern Definitions

#### Success Pattern (Green Circle)

- **Lit Keys**: 5, 6, 7, 8, R, I, D, K, C, M, SPACE
- **Color**: RGB(0, 255, 0)
- **Shape**: Circle (○) representation

#### Failure Pattern (Red Cross)

- **Lit Keys**: 5, 8, T, Y, H, N, C
- **Color**: RGB(255, 0, 0)
- **Shape**: Cross (×) representation

### 4.3 Bitmap Implementation

```go
// 132-element array (all keys)
type KeyboardBitmap [132][3]uint8  // [R, G, B]

// Key name to index mapping
var keyIndexMap = map[string]int{
    "5": 5,
    "6": 6,
    // ...
}
```

## 5. Detailed Design

### 5.1 main.go

```go
// Command line argument processing
// --success: Display success pattern
// --failure: Display failure pattern
// Error handling and logging
```

### 5.2 internal/steelseries/config.go

```go
type CoreProps struct {
    Address string `json:"address"`
}

// LoadConfig(): Read coreProps.json
// GetAPIEndpoint(): Generate endpoint URL
```

### 5.3 internal/steelseries/client.go

```go
type Client struct {
    baseURL string
    httpClient *http.Client
}

// RegisterGame(): Register game
// BindEvent(): Bind event handler
// SendEvent(): Send LED pattern
// StartHeartbeat(): Send heartbeat
```

### 5.4 internal/steelseries/patterns.go

```go
// GetSuccessPattern(): Generate success pattern
// GetFailurePattern(): Generate failure pattern
// CreateBitmap(): Create bitmap array
```

### 5.5 internal/cli/commands.go

```go
// ParseArgs(): Parse command line arguments
// Execute(): Execute command
```

## 6. API Request/Response Specification

### 6.1 Game Registration

```json
// Request
{
  "game": "KEYLIGHT",
  "game_display_name": "Keylight LED Notifier",
  "developer": "mimikun"
}


// Response: 200 OK
```

### 6.2 Event Binding

```json
// Request
{
  "game": "KEYLIGHT",
  "event": "BUILD_STATUS",
  "handlers": [
    {
      "device-type": "rgb-per-key-zones",
      "zone": "all",
      "mode": "bitmap"
    }
  ]
}
```

### 6.3 Event Sending (Bitmap Mode)

```json
// Request
{
  "game": "KEYLIGHT",
  "event": "BUILD_STATUS",
  "data": {
    "value": 1,
    "frame": {
      "bitmap": [
        [0, 0, 0], // Key 0
        [0, 255, 0] // Key 1 (lit)
        // ... 132 keys total
      ]
    }
  }
}
```

## 7. Error Handling

### 7.1 Expected Errors

1. **SteelSeries Engine Not Running**

   - Error Message: "SteelSeries Engine is not running"
   - Action: Prompt user to start

2. **Configuration File Not Found**

   - Error Message: "Configuration file not found"
   - Action: Prompt installation check

3. **Network Error**
   - Error Message: "Failed to connect to GameSense API"
   - Action: Retry or show detailed error

### 7.2 Error Handling Policy

- User-friendly messages
- Detailed logs only in debug mode
- Return status via exit code (0: success, 1: error)

## 8. Implementation Considerations

### 8.1 Performance

- HTTP connection reuse
- Minimize unnecessary API calls
- Timeout setting (5 seconds)

### 8.2 Compatibility

- Windows 10/11 support
- SteelSeries Engine 3/GG compatibility
- 32bit/64bit environment consideration

### 8.3 Maintainability

- Loose coupling between modules
- Interface abstraction
- Unit testable design

## 9. Build and Deploy

### 9.1 Build Command

```bash
# Build for Windows x64
GOOS=windows GOARCH=amd64 go build -o keylight.exe
```

### 9.2 Distribution Format

- Single executable (keylight.exe)
- No dependency libraries
- ZIP archive with README.md

## 10. Future Extensibility

### 10.1 Potential Features

- Custom pattern support
- Animation effects
- Other SteelSeries device support
- Configuration file customization

### 10.2 Architecture Considerations

- Pattern definition externalization
- Plugin mechanism preparation
- Internationalization foundation

## 11. Security Considerations

### 11.1 Communication Security

- Localhost only communication
- No external network connections

### 11.2 File Access

- Read-only file access
- No administrator privileges required

## 12. Development Methodology

### 12.1 Test-Driven Development (TDD)

This project follows TDD practices to ensure code quality and reliability.

#### TDD Workflow

1. **Red**: Write a failing test that describes the desired behavior
2. **Green**: Write the minimal code to make the test pass
3. **Refactor**: Clean up the code while keeping tests passing

#### Testing Guidelines

- Write tests before implementing functionality
- Each function should have corresponding unit tests
- Use table-driven tests for multiple scenarios
- Mock external dependencies (HTTP API calls, file system operations)
- Aim for high test coverage, especially for core business logic

#### Test Organization

- Test files follow Go convention: `*_test.go`
- Place tests in the same package as the code under test
- Use `testify` assertions if needed, but prefer standard Go testing
- Group related tests using subtests (`t.Run()`)

#### Integration Testing

- Test the complete flow from CLI to LED pattern display
- Mock SteelSeries GameSense API responses
- Verify error handling scenarios
- Test configuration file parsing

### 12.2 Test Structure

```
keylight/
├── internal/
│   ├── steelseries/
│   │   ├── client.go
│   │   ├── client_test.go      # Unit tests for client
│   │   ├── config.go
│   │   ├── config_test.go      # Unit tests for config
│   │   ├── patterns.go
│   │   └── patterns_test.go    # Unit tests for patterns
│   └── cli/
│       ├── commands.go
│       └── commands_test.go    # Unit tests for CLI
└── integration/
    └── keylight_test.go        # End-to-end integration tests
```

## 13. Git Conventions

This project follows Conventional Commits specification for commit messages.

### 13.1 Commit Message Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### 13.2 Commit Types

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

### 13.3 Scopes

Use scopes to indicate the area of the codebase being modified:

- **cli**: Command line interface and argument parsing
- **client**: SteelSeries GameSense API client
- **config**: Configuration file handling
- **patterns**: LED pattern definitions
- **docs**: Documentation changes
- **build**: Build configuration and scripts
- **test**: Test-related changes

### 13.4 Examples

```
feat(cli): add success LED pattern display
fix(config): handle missing coreProps.json file
docs(design): add TDD guidelines and Git conventions
test(client): add unit tests for GameSense client
refactor(patterns): simplify LED bitmap generation
perf(client): optimize HTTP request handling
```

## 14. Development Environment Setup

### 14.1 WSL Development Setup

When developing on WSL2, you can directly access Windows files to verify SteelSeries Engine configuration:

```bash
# Check SteelSeries Engine configuration
cat /mnt/c/ProgramData/SteelSeries/SteelSeries\ Engine\ 3/coreProps.json

# Alternative path if user profile is needed
cat /mnt/c/Users/mimikun/AppData/Local/SteelSeries/SteelSeries\ Engine\ 3/coreProps.json
```

This allows real-time verification of GameSense API endpoint without running on Windows.

### 14.2 Cross-Platform Development

- **Development OS**: Linux (WSL2)
- **Target OS**: Windows (x64)
- **Go Cross-Compilation**: Enabled via GOOS/GOARCH environment variables

### 14.3 Testing Commands

```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./internal/steelseries

# Run tests with coverage
go test -cover ./...

# Run integration tests
go test ./integration
```

