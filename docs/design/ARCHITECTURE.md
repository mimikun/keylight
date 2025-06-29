# System Architecture

## Layer Structure

```
┌─────────────────────────────────────┐
│         CLI Layer (main.go)         │
├─────────────────────────────────────┤
│      Command Layer (cli/)           │
├─────────────────────────────────────┤
│     Hue Client (hue/)               │
├─────────────────────────────────────┤
│    Philips Hue Bridge API           │
└─────────────────────────────────────┘
```

## Module Structure

```
keylight/
├── main.go                   # Entry point
├── internal/
│   ├── hue/
│   │   ├── client.go        # Hue Bridge API client
│   │   ├── config.go        # Configuration and discovery
│   │   ├── scenes.go        # Scene management
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

## Scene Management Design

### Target Device

Philips Hue Bridge and Hue desk lights connected to the local network.
Communication via Hue Bridge API to control lighting scenes and individual light states.

### Scene Definitions

#### Success Scene

- **Name**: "Success"
- **Color**: Green lighting pattern
- **Duration**: 10 seconds
- **Behavior**: Activate scene, wait, then restore original state

#### Failure Scene

- **Name**: "Failure"
- **Color**: Red lighting pattern
- **Duration**: 10 seconds
- **Behavior**: Activate scene, wait, then restore original state

### State Management Implementation

```go
// Scene state for restoration
type SceneState struct {
    LightStates map[string]LightState `json:"light_states"`
    GroupState  GroupState            `json:"group_state"`
}

// Light state structure
type LightState struct {
    On         bool    `json:"on"`
    Brightness uint8   `json:"bri"`
    Hue        uint16  `json:"hue"`
    Saturation uint8   `json:"sat"`
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

### internal/hue/config.go

```go
// Flexible configuration system supporting multiple formats
type ConfigProvider interface {
    Load(filePath string) (*Config, error)
    Save(filePath string, config *Config) error
    GetExtension() string
    GetDefaultFilename() string
}

type Config struct {
    BridgeIP         string      `json:"bridge_ip" yaml:"bridge_ip" toml:"bridge_ip"`
    Username         string      `json:"username" yaml:"username" toml:"username"`
    Scenes           SceneConfig `json:"scenes" yaml:"scenes" toml:"scenes"`
    AutoCreateScenes bool        `json:"auto_create_scenes" yaml:"auto_create_scenes" toml:"auto_create_scenes"`
}

type SceneConfig struct {
    DefaultScene string `json:"default_scene" yaml:"default_scene" toml:"default_scene"`
    SuccessScene string `json:"success_scene" yaml:"success_scene" toml:"success_scene"`
    FailureScene string `json:"failure_scene" yaml:"failure_scene" toml:"failure_scene"`
}

// DiscoverBridge(): Discover bridge on network
// LoadConfig(): Read configuration file with format auto-detection
// SaveConfig(): Save configuration in specified format
// MigrateFormat(): Convert between configuration formats
// GetAPIEndpoint(): Generate endpoint URL
```

### internal/hue/client.go

```go
type Client struct {
    baseURL    string
    username   string
    httpClient *http.Client
}

// Authenticate(): Authenticate with bridge
// GetScenes(): List available scenes
// EnsureScenes(): Auto-create required scenes if missing
// CaptureState(): Capture current light state
// ActivateScene(): Activate specified scene
// RestoreState(): Restore captured state or fallback to default scene
```

### internal/hue/scenes.go

```go
type SceneDefinition struct {
    Name       string `json:"name"`
    Color      string `json:"color"`
    Brightness uint8  `json:"brightness"`
    Duration   int    `json:"duration,omitempty"`
}

type SceneManager struct {
    client           *Client
    config           *Config
    originalStates   map[string]LightState
    fallbackMode     bool
}

// GetDefaultScenes(): Get predefined scene configurations
// CreateScene(): Create scene via API v2
// EnsureRequiredScenes(): Auto-create missing scenes
// GetSuccessScene(): Get success scene configuration
// GetFailureScene(): Get failure scene configuration
// ValidateScenes(): Validate scenes exist on bridge
```

### internal/cli/commands.go

```go
// ParseArgs(): Parse command line arguments
// Execute(): Execute command
```

## Error Handling

### Expected Errors

1. **Hue Bridge Not Found**
   - Error Message: "Hue Bridge not found on network"
   - Action: Prompt user to check network connection

2. **Authentication Failed**
   - Error Message: "Authentication with Hue Bridge failed"
   - Action: Prompt user to press bridge button and retry

3. **Scene Not Found**
   - Error Message: "Required scene not found on bridge"
   - Action: Prompt user to create scenes

4. **Network Error**
   - Error Message: "Failed to connect to Hue Bridge"
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
- Philips Hue Bridge v2 compatibility
- Local network accessibility

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

## Configuration System Architecture

### Multi-Format Support

```go
type ConfigManager struct {
    provider ConfigProvider
    filePath string
}

// Supported formats: JSON, YAML, TOML
// Format auto-detection by file extension
// Migration between formats without data loss
```

### Format Providers

- **JSONConfigProvider**: Default format, compact and widely supported
- **YAMLConfigProvider**: Human-readable format, good for manual editing
- **TOMLConfigProvider**: Configuration-focused format, good for structured data

### Migration Strategy

```go
// Example: Migrate from JSON to YAML
configManager.MigrateFormat("keylight.yaml")

// Preserves all configuration values
// Automatic backup of original file
// Validation of migrated configuration
```

## Future Extensibility

### Potential Features

- Custom scene support
- Animation effects
- Multiple room support
- Plugin-based configuration providers
- Remote configuration management

### Architecture Considerations

- Scene definition externalization
- Plugin mechanism preparation
- Configuration format extensibility
- Internationalization foundation

## Security Considerations

### Communication Security

- Localhost only communication
- No external network connections

### File Access

- Read-only file access
- No administrator privileges required