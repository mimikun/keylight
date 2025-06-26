# GameSense API Specification

## API Overview

Uses local HTTP API provided by SteelSeries Engine. No authentication required.

## Configuration File

- **Windows**: `%PROGRAMDATA%/SteelSeries/SteelSeries Engine 3/coreProps.json`
- **Format**:

```json
{
  "address": "127.0.0.1:12345"
}
```

## Main Endpoints

1. **Game Registration**: `POST http://{address}/game_metadata`
2. **Event Binding**: `POST http://{address}/bind_game_event`
3. **Event Sending**: `POST http://{address}/game_event`
4. **Heartbeat**: `POST http://{address}/game_heartbeat`

## Communication Flow

```
1. Read coreProps.json → Get API address
2. Register game (KEYLIGHT)
3. Bind event handler
4. Send LED control event
5. Auto-off after 3-5 seconds
```

## API Request/Response Specification

### Game Registration

```json
// Request
{
  "game": "KEYLIGHT",
  "game_display_name": "Keylight LED Notifier",
  "developer": "mimikun"
}


// Response: 200 OK
```

### Event Binding

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

### Event Sending (Bitmap Mode)

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

## Error Responses

### Common HTTP Status Codes

- `200 OK`: Request successful
- `400 Bad Request`: Invalid request format
- `404 Not Found`: Endpoint not found
- `500 Internal Server Error`: SteelSeries Engine error

### Error Handling

- **Connection Refused**: SteelSeries Engine not running
- **Timeout**: Network connectivity issues
- **Invalid JSON**: Malformed request payload

## Rate Limiting

- No explicit rate limits documented
- Recommended to avoid excessive API calls
- Use heartbeat mechanism for persistent connections

## API Versioning

- SteelSeries Engine 3: Base API version
- SteelSeries GG: Enhanced compatibility
- Backward compatibility maintained