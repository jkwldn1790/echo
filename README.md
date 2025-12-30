# Echo

A lightweight HTTP service written in Go that displays network and system information about its host.

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `/` | Displays the server's hostname and local IP address |
| `/hello` | Returns a greeting. Accepts optional `name` query parameter |

### Examples

```bash
curl http://localhost:3333/
curl http://localhost:3333/hello
curl "http://localhost:3333/hello?name=World"
```

## Running with Docker

```bash
# Build the image
make build

# Run the container
make run

# Remove the image
make clean
```

The service listens on port 3333.

## Running locally

```bash
go run main.go
```
