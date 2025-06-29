# Development Guidelines

## Development Methodology

### Test-Driven Development (TDD)

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

- Test the complete flow from CLI to Hue scene activation
- Mock Philips Hue Bridge API responses
- Verify error handling scenarios
- Test bridge discovery and authentication

### Test Structure

```
keylight/
├── internal/
│   ├── hue/
│   │   ├── client.go
│   │   ├── client_test.go      # Unit tests for client
│   │   ├── config.go
│   │   ├── config_test.go      # Unit tests for config
│   │   ├── scenes.go
│   │   └── scenes_test.go      # Unit tests for scenes
│   └── cli/
│       ├── commands.go
│       └── commands_test.go    # Unit tests for CLI
└── integration/
    └── keylight_test.go        # End-to-end integration tests
```

## Git Conventions

This project follows Conventional Commits specification for commit messages.

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
- **docs**: Documentation changes
- **build**: Build configuration and scripts
- **test**: Test-related changes

### Examples

```
feat(cli): add success scene activation
fix(config): handle Hue Bridge discovery failure
docs(design): add TDD guidelines and Git conventions
test(client): add unit tests for Hue Bridge client
refactor(scenes): simplify scene state management
perf(client): optimize HTTP request handling
```

## Development Environment Setup

### Cross-Platform Development Setup

The Hue Bridge is accessible from both Linux and Windows environments via local network:

```bash
# Discover Hue Bridge on local network
curl -X GET "https://discovery.meethue.com/"

# Test connection to bridge (replace <bridge_ip> with actual IP)
curl -X GET "http://<bridge_ip>/api/<username>/scenes"

# Test scene activation
curl -X PUT "http://<bridge_ip>/api/<username>/groups/0/action" \
  -H "Content-Type: application/json" \
  -d '{"scene": "Success"}'
```

This allows development and testing from any environment with network access to the Hue Bridge.

### Cross-Platform Development

- **Development OS**: Linux (WSL2)
- **Target OS**: Windows (x64)
- **Go Cross-Compilation**: Enabled via GOOS/GOARCH environment variables

### Testing Commands

```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./internal/hue

# Run tests with coverage
go test -cover ./...

# Run integration tests
go test ./integration
```

### Build Commands

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

## Code Quality Standards

### Code Style

- Follow standard Go formatting with `gofmt`
- Use meaningful variable and function names
- Write clear, concise comments for public APIs
- Avoid deep nesting and complex functions

### Error Handling

- Handle all errors explicitly
- Provide meaningful error messages to users
- Log detailed errors for debugging
- Use appropriate error types and wrapping

### Performance Guidelines

- Minimize memory allocations in hot paths
- Reuse HTTP connections
- Set appropriate timeouts
- Profile performance-critical code

### Security Practices

- Validate all inputs
- Use secure defaults
- Avoid logging sensitive information
- Follow principle of least privilege