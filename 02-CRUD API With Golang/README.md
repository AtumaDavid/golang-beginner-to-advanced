# Movies CRUD API (Golang)

Simple RESTful CRUD API for managing movies, implemented in Go using `github.com/gorilla/mux`.

**Project layout**

- `main.go`: application entrypoint with handlers and routes
- `go.mod`: module and dependency configuration

**Prerequisites**

- Go (tested with the version declared in `go.mod`)

**Install dependencies**
Run from the project root.

If this is a fresh project (no `go.mod` present), initialize the module:

```bash
go mod init movies-crud
```

If `go.mod` already exists (you've already initialized the module), download dependencies or tidy the module:

```bash
go mod download
# or
go mod tidy
```

**Build & Run**

Run directly:

```bash
go run main.go
```

Build an executable:

```bash
go build -o movies-api ./
./movies-api
```

The server listens on port `8000` by default.

**API Endpoints**

Base URL: `http://localhost:8000`

- GET `/movies` — List all movies
- GET `/movies/{id}` — Get a single movie by `id`
- POST `/movies` — Create a new movie
- PUT `/movies/{id}` — Update an existing movie
- DELETE `/movies/{id}` — Delete a movie

**Example requests**

List movies:

```bash
curl -s http://localhost:8000/movies | jq
```

Get single movie (replace `<id>`):

```bash
curl -s http://localhost:8000/movies/1 | jq
```

Create movie:

```bash
curl -s -X POST http://localhost:8000/movies \
  -H "Content-Type: application/json" \
  -d '{"isbn":"123456","title":"New Movie","director":{"firstname":"Jane","lastname":"Roe"}}' | jq
```

Update movie (replace `<id>`):

```bash
curl -s -X PUT http://localhost:8000/movies/<id> \
  -H "Content-Type: application/json" \
  -d '{"isbn":"987654","title":"Updated Movie","director":{"firstname":"Jane","lastname":"Roe"}}' | jq
```

Delete movie (replace `<id>`):

```bash
curl -s -X DELETE http://localhost:8000/movies/<id>
```

**Request / Response notes**

- Request and response bodies use JSON.
- The server assigns a random `id` when creating a movie. IDs are strings.
- Error handling is minimal: this example returns simple responses and does not provide comprehensive error messages or validation. For production, add detailed error handling, validation, and proper HTTP status codes.

**Known constraints & next steps**

- Port is currently hardcoded to `8000` in `main.go`. Consider reading `PORT` from an environment variable for flexibility.
- There is no persistent storage — movies are stored in-memory and will reset when the server restarts. Add a database (SQLite/Postgres) to persist data.
- Add input validation and better error handling for production readiness.
- Consider adding tests, a `Dockerfile`, and CI workflow.

**Contact / License**
This is a small example project
