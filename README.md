# API-Notes

Dead simple service for publishing notes via API. Markdown supported.

The service exposes minimal API for upload markdown (with attachments) and render it as HTML.

The generated link is 16-bytes randomly generated and can be shared relatively safely.

The API-Notes is NOT serving notes - any reverse proxy must do it. With authorization if needed.
See [docker-compose.yaml].

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

## Configuration

| Environment variable | Default                 | Description                  |
|----------------------|-------------------------|------------------------------|
| `PUBLIC_URL`         | `http://127.0.0.1:8080` | Public URL to where redirect |
| `BIND`               | `127.0.0.1:8080`        | Binding address              |
| `DIR`                | `notes`                 | Directory for notes          |

Differences in docker version

| Environment variable | Default        | Description         |
|----------------------|----------------|---------------------|
| `BIND`               | `0.0.0.0:8080` | Binding address     |
| `DIR`                | `/data`        | Directory for notes |

## API

Supported content type:

- `application/json`: only `title` (string) and `text` (string) supported
- `application/x-www-form-urlencoded`: only `title` and `text` field supported
- `multipart/form-data`: `title`, `text` and any file(s) supported

> Note: note ID should not be escaped

### Create note

    POST /note

Returns `303 See Other` with public link in header `Location`, and ID in `X-Correlation-Id`.

**Example**

Create with basic form

    curl -v -d 'title=hello&text=this is text' http://127.0.0.1:8080/note

Create with JSON

    curl -v -H 'Content-Type: application/json' --data-binary '{"title": "hello", "text": "this is text"}' http://127.0.0.1:8080/note

Create with attachments

    curl -v -F title=hello -F text="this is text" -F file1=@file1 -F file2=@file2 http://127.0.0.1:8080/note

### Update note

    PUT /note/{note ID}

Returns `303 See Other` with public link (old) in header `Location`.

**Example**

Update note (id: `b6/22/7f/2ce9505bef907fa98075f852b6`) with basic form

    curl -X PUT -v -d 'title=hello&text=this is another text' http://127.0.0.1:8080/note/b6/22/7f/2ce9505bef907fa98075f852b6

### Delete note

    DELETE /note/{note ID}

Returns `204 No Content`

**Example**

Delete note (id: `b6/22/7f/2ce9505bef907fa98075f852b6`)

    curl -X DELETE -v http://127.0.0.1:8080/note/b6/22/7f/2ce9505bef907fa98075f852b6


## Sample note

![Screenshot 2023-06-07 230208](https://github.com/reddec/api-notes/assets/6597086/c5b7b999-ea25-4fb2-996d-bf8f6a6e9c0d)
