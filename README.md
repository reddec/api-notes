# API-Notes

Dead simple service for publishing notes via API. Markdown supported.

The service exposes minimal API for upload markdown (with attachments) and render it as HTML.

The generated link is 16-bytes randomly generated and can be shared relatively safely.

The API-Notes is NOT serving notes - any reverse proxy must do it. With authorization if needed.
See [docker-compose.yaml](docker-compose.yaml).

Supported markdown extensions:

- GFM (GitHub Flavored Markdown)
- Footnotes
- Basic syntax highlighting
- Mermaid
- MathJax
- Embedded youtube

## Installation

- Docker: `ghcr.io/reddec/api-notes:latest`
- Go: `go install github.com/reddec/api-notes/cmd/...@latest`

## Usage

```
Usage:
  api-notes [OPTIONS] serve [serve-OPTIONS]

Help Options:
  -h, --help            Show this help message

[serve command options]
      -u, --public-url= Public URL for redirects (default: http://127.0.0.1:8080) [$API_NOTES_PUBLIC_URL]
      -b, --bind=       API binding address (default: 127.0.0.1:8080) [$API_NOTES_BIND]
      -d, --dir=        Directory to store notes (default: notes) [$API_NOTES_DIR]
      -t, --token=      Authorization token, empty means any token can be usedauth is disabled [$API_NOTES_TOKEN]

```

Differences in docker version

| Environment variable | Default        | Description         |
|----------------------|----------------|---------------------|
| `API_NOTES_BIND`     | `0.0.0.0:8080` | Binding address     |
| `API_NOTES_DIR`      | `/data`        | Directory for notes |

## API

- [OpenAPI spec](openapi.yaml)
- [Live docs](https://elements-demo.stoplight.io/?spec=https://raw.githubusercontent.com/reddec/api-notes/master/openapi.yaml)

### Example

Create note

    curl -v -F author=reddec -F title=hello -F text=world -F attachment=@somefile.pdf http://127.0.0.1:8080/notes?token=deadbeaf


## Sample note

![Screenshot 2023-06-07 230208](https://github.com/reddec/api-notes/assets/6597086/c5b7b999-ea25-4fb2-996d-bf8f6a6e9c0d)
