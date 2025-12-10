# go-api-demo

Minimal Go HTTP server that exposes a single JSON endpoint for quick API demos or testing.

## Prerequisites
- Go 1.21+ installed locally

## Run locally
```bash
go run main.go
```
The server listens on `http://localhost:8080`.

## API
- `GET /hello` → `200 OK` with JSON body `{"hello":"world"}`

Example request:
```bash
curl http://localhost:8080/hello
```

## Project layout
- `main.go` — contains the HTTP handler and server setup

## License
MIT — see `LICENSE` for details.
