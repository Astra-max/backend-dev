# ASCII Art Web


This repository contains a Go web application that renders ASCII art from user-provided text using selectable banners (standard, shadow, thinkertoy). The app serves HTML templates and assets exclusively from the templates directory and exposes a minimal set of routes for UI and form submission.


## Features
- Generate ASCII art from input text with selectable banner styles
- HTML templating (layout + pages) with dark theme UI
- Serves assets from the templates directory only (no static/)
- Robust server setup with timeouts and logging middleware (requests, panics)
- Clean routing and handlers
- Unit tests for core HTTP handlers


## Requirements
- Go 1.22+


## Project Structure
```
.
├── ascii-art/                 # Core ASCII art logic and banner files
│   ├── AsciiArt.go
│   └── artstyles/
│       ├── standard.txt
│       ├── shadow.txt
│       └── thinkertoy.txt
├── cmd/
│   └── main.go                # Entry point (go run cmd/main.go)
├── internals/
│   ├── handlers/              # JSON helpers, template routing, ascii-art API
│   │   ├── ascii-art.go
│   │   ├── response.go
│   │   └── serve-static.go
│   ├── routes/
│   │   └── api.go             # HTTP routes mapped to handlers
│   └── server/
│       └── server.go          # HTTP server, middleware, StartServer()
├── handlers/                  # Template-aware handlers used by routes
│   └── handlers.go
├── templates/                 # All UI templates and assets (only served directory)
│   ├── layout.html
│   ├── index.html
│   ├── index.css
│   ├── index.js
│   └── errors/
│       ├── 404.html
│       └── 500.html
├── tests/
│   └── main_test.go
├── go.mod
└── README.md
```


## Running the Application
- Development server:
  - go run cmd/main.go
- The server listens on:
  - http://localhost:3000/

If the port is already taken, change the Addr in internals/server/server.go.


## Routes
- GET /
  - Renders templates/layout.html with templates/index.html content.
- POST /ascii-art
  - Accepts form data (application/x-www-form-urlencoded):
    - artstyle: one of standard, shadow, thinkertoy
    - text: user input to convert
  - Re-renders index.html with the generated ASCII art in the Output panel.
- GET /templates/*
  - Serves assets directly from the templates directory (e.g. /templates/index.css).


## Templates and Assets
- Layout: templates/layout.html defines the base layout and links to /templates/index.css for the dark theme.
- Page: templates/index.html defines the main form and output area.
- Error pages: templates/errors/404.html and templates/errors/500.html are rendered via the shared layout when errors occur.


## Logging and Error Handling
- Console logs include:
  - Request method, path, and duration for each request
  - Panic recovery with 500 responses and error output
  - Server startup messages and fatal errors


## Tests
- Run all tests:
  - go test ./...
- Verbose output:
  - go test -v ./...
- Specific package:
  - go test -v ./tests
- Specific test by name:
  - go test -v ./tests -run TestServeTemplate
  - go test -v ./tests -run TestHandleAsciiArt


## Development Notes
- Only the templates directory is exposed for assets under the path prefix /templates/.
- The UI adheres to a dark theme; update templates/index.css to tweak the styling.
- Handler flow for ASCII generation:
  - handlers.HandleAsciiArt parses form values, invokes ascii_art.AsciiArt, then renders the index template with the result.


## Branching and Remote
- Active branch for recent changes: update
- Remote repository: https://learn.zone01kisumu.ke/git/maodongo/ascii-art-web.git
- Push current branch:
  - git push -u origin update


## License
See LICENSE for details.


## Maintainer
