# System Architecture

## Layer Structure

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

## Module Structure

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

## LED Pattern Design

### Keyboard Layout (JIS)

Target device: SteelSeries Apex Pro JP (Product No. 64629) - Full-size Japanese JIS layout with numpad.
Each key corresponds to an index in the bitmap array. Managed with 132-key array covering all keys including function row, main typing area, and numeric keypad.

### Pattern Definitions

#### Success Pattern (Green Circle)

- **Lit Keys**: 5, 6, 7, 8, R, I, D, K, C, M, SPACE
- **Color**: RGB(0, 255, 0)
- **Shape**: Circle (○) representation
- **Additional Options**: Can utilize numpad keys (1-9) for enhanced patterns on full-size layout

#### Failure Pattern (Red Cross)

- **Lit Keys**: 5, 8, T, Y, H, N, C
- **Color**: RGB(255, 0, 0)
- **Shape**: Cross (×) representation
- **Additional Options**: Can utilize function keys (F1-F12) for status indication

### Bitmap Implementation

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

## Detailed Design

### main.go

```go
// Command line argument processing
// --success: Display success pattern
// --failure: Display failure pattern
// Error handling and logging
```

### internal/steelseries/config.go

```go
type CoreProps struct {
    Address string `json:"address"`
}

// LoadConfig(): Read coreProps.json
// GetAPIEndpoint(): Generate endpoint URL
```

### internal/steelseries/client.go

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

### internal/steelseries/patterns.go

```go
// GetSuccessPattern(): Generate success pattern
// GetFailurePattern(): Generate failure pattern
// CreateBitmap(): Create bitmap array
```

### internal/cli/commands.go

```go
// ParseArgs(): Parse command line arguments
// Execute(): Execute command
```

## Error Handling

### Expected Errors

1. **SteelSeries Engine Not Running**
   - Error Message: "SteelSeries Engine is not running"
   - Action: Prompt user to start

2. **Configuration File Not Found**
   - Error Message: "Configuration file not found"
   - Action: Prompt installation check

3. **Network Error**
   - Error Message: "Failed to connect to GameSense API"
   - Action: Retry or show detailed error

### Error Handling Policy

- User-friendly messages
- Detailed logs only in debug mode
- Return status via exit code (0: success, 1: error)

## Implementation Considerations

### Performance

- HTTP connection reuse
- Minimize unnecessary API calls
- Timeout setting (5 seconds)

### Compatibility

- Windows 10/11 support
- SteelSeries Engine 3/GG compatibility
- 32bit/64bit environment consideration

### Maintainability

- Loose coupling between modules
- Interface abstraction
- Unit testable design

## Build and Deploy

### Build Command

```bash
# Build for Windows x64
GOOS=windows GOARCH=amd64 go build -o keylight.exe
```

### Distribution Format

- Single executable (keylight.exe)
- No dependency libraries
- ZIP archive with README.md

## Future Extensibility

### Potential Features

- Custom pattern support
- Animation effects
- Other SteelSeries device support
- Configuration file customization

### Architecture Considerations

- Pattern definition externalization
- Plugin mechanism preparation
- Internationalization foundation

## Security Considerations

### Communication Security

- Localhost only communication
- No external network connections

### File Access

- Read-only file access
- No administrator privileges required