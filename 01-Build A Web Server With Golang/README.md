# Go Web Server

A simple web server built with Go that serves static files and handles form submissions.

## Features

- Serves static HTML files from the `static/` directory
- Form submission handling with POST requests
- Simple "Hello" endpoint

## Project Structure

```
.
├── main.go           # Main server code
├── go.mod            # Go module file
└── static/
    ├── index.html    # Home page
    └── form.html     # Form page
```

## Getting Started

### Prerequisites

- Go 1.25.4 or higher

### Installation

1. Clone or download this project
2. Navigate to the project directory

### Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Building the Project

To build an executable:

```bash
go build -o webserver
```

Then run it:

```bash
./webserver
```

## Available Endpoints

- `http://localhost:8080/` - Serves static files (index.html)
- `http://localhost:8080/form.html` - Form page
- `http://localhost:8080/form` - Form submission handler (POST)
- `http://localhost:8080/hello` - Simple hello endpoint (GET)

## Common Go Commands

```bash
# Initialize a new Go module
go mod init webserver

# Run the application
go run main.go

# Build the application
go build

# Build with custom output name
go build -o webserver

# Download dependencies
go mod download

# Clean up dependencies
go mod tidy

# Format your code
go fmt

# Run tests
go test
```

## Usage Example

1. Start the server: `go run main.go`
2. Open your browser and navigate to `http://localhost:8080`
3. Visit `http://localhost:8080/form.html` to test the form submission
4. Fill out the form and submit to see the POST request response
