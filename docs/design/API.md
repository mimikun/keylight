# Philips Hue Bridge API Specification

## API Overview

Uses local HTTP API provided by Philips Hue Bridge. Requires authentication with username token.

## Bridge Discovery

### Discovery Methods

1. **Hue Discovery Service**: `GET https://discovery.meethue.com/`
2. **UPnP Discovery**: Search for device type `urn:schemas-upnp-org:device:basic:1`
3. **IP Scan**: Scan local network for bridge (fallback method)

### Discovery Response

```json
[
  {
    "id": "001788fffe09abcd",
    "internalipaddress": "192.168.1.100"
  }
]
```

## Authentication

### Initial Setup (Button Press Required)

```http
POST http://{bridge_ip}/api
Content-Type: application/json

{
  "devicetype": "keylight#development"
}
```

**Response (Success)**:
```json
[
  {
    "success": {
      "username": "abcdef1234567890"
    }
  }
]
```

**Response (Button Not Pressed)**:
```json
[
  {
    "error": {
      "type": 101,
      "address": "",
      "description": "link button not pressed"
    }
  }
]
```

## Main Endpoints

1. **Bridge Info**: `GET http://{bridge_ip}/api/{username}`
2. **Scene Management**: `GET/PUT http://{bridge_ip}/api/{username}/scenes`
3. **Light Control**: `GET/PUT http://{bridge_ip}/api/{username}/lights`
4. **Group Control**: `GET/PUT http://{bridge_ip}/api/{username}/groups`

## Communication Flow

```
1. Discover bridge on network
2. Authenticate with bridge (username)
3. Capture current scene state
4. Activate notification scene
5. Wait 10 seconds
6. Restore original state
```

## API Request/Response Specification

### Get Current State

```json
// Request: GET /api/{username}/lights
// Response
{
  "1": {
    "state": {
      "on": true,
      "bri": 144,
      "hue": 13088,
      "sat": 212,
      "ct": 343,
      "xy": [0.3460, 0.3568]
    },
    "name": "Desk Light"
  }
}
```

### Get Scenes

```json
// Request: GET /api/{username}/scenes
// Response
{
  "Success": {
    "name": "Success",
    "lights": ["1", "2"],
    "owner": "abcdef1234567890",
    "recycle": false,
    "locked": false,
    "appdata": {},
    "picture": "",
    "lastupdated": "2023-01-01T00:00:00"
  },
  "Failure": {
    "name": "Failure",
    "lights": ["1", "2"],
    "owner": "abcdef1234567890",
    "recycle": false,
    "locked": false,
    "appdata": {},
    "picture": "",
    "lastupdated": "2023-01-01T00:00:00"
  }
}
```

### Activate Scene

```json
// Request: PUT /api/{username}/groups/0/action
{
  "scene": "Success"
}

// Response
[
  {
    "success": {
      "/groups/0/action/scene": "Success"
    }
  }
]
```

### Control Individual Light

```json
// Request: PUT /api/{username}/lights/1/state
{
  "on": true,
  "bri": 144,
  "hue": 13088,
  "sat": 212
}

// Response
[
  {
    "success": {
      "/lights/1/state/on": true
    }
  },
  {
    "success": {
      "/lights/1/state/bri": 144
    }
  }
]
```

## Scene Definitions

### Success Scene
- **Color**: Green (hue: ~25500, sat: 254)
- **Brightness**: High (bri: 254)
- **Transition**: Smooth (transitiontime: 4)

### Failure Scene
- **Color**: Red (hue: 0, sat: 254)
- **Brightness**: High (bri: 254)
- **Transition**: Smooth (transitiontime: 4)

## Error Responses

### Common HTTP Status Codes

- `200 OK`: Request successful
- `400 Bad Request`: Invalid request format
- `401 Unauthorized`: Invalid username
- `404 Not Found`: Resource not found
- `503 Service Unavailable`: Bridge busy

### Error Response Format

```json
[
  {
    "error": {
      "type": 1,
      "address": "/lights/99/state",
      "description": "resource, /lights/99/state, not available"
    }
  }
]
```

### Common Error Types

- **Type 1**: Unauthorized user
- **Type 3**: Resource not available
- **Type 101**: Link button not pressed
- **Type 201**: Parameter not available

## Rate Limiting

- **Limit**: 10 commands per second per IP
- **Burst**: Up to 20 commands in short burst
- **Recommendation**: Add delays between commands for reliability

## Security Considerations

### Network Security

- **Local Network Only**: Bridge accessible only from local network
- **HTTPS**: Use HTTPS for discovery service calls
- **Username Storage**: Store username securely in configuration file

### Authentication

- **Username Generation**: Each application should have unique username
- **Token Expiration**: Usernames don't expire but can be revoked
- **Button Press**: Physical access required for initial setup

## API Versioning

- **API Version**: v1 (current stable version)
- **Bridge Version**: Support for Bridge API v1.19+
- **Backward Compatibility**: Maintained across bridge firmware updates