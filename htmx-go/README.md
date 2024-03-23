# HTMX-GO

## A very simple htmx-go 1 page example (non-canon)

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, `
            <!DOCTYPE html>
            <html>
            <head>
                <script src="https://unpkg.com/htmx.org@1.6.1"></script>
            </head>
            <body>
                <button hx-get="/clicked" hx-swap="outerHTML">Click me</button>
            </body>
            </html>
        `)
    })

    http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Button clicked!")
    })

    http.ListenAndServe(":3000", nil)
}
```

## Project Structure

```bash
.
├── cmd
│   └── main.go 
├── pkg
│   ├── handlers
│   │   └── home.go
│   └── models
├── go.mod
├── go.sum
└── web
    ├── htmx
    ├── static
    │   ├── css
    │   │   └── style.css
    │   └── js
    │       └── htmx.min.js
    └── templates
        └── index.html
```

This is of course configurable, but this is a basic structure.

## Setup

### Install HTMX

Copy the script from [here](https://unpkg.com/htmx.org@1.9.11/dist/htmx.min.js) and it to `/web/static/js`.


### Install `Go`lang

[umm](https://go.dev/doc/install)

### Setup Go

create `go.mod` and add a unique _string identifier_ for the project and specify Go version.

I'm using my github repo in this case.

```mod
module github.com/kinxyo/knowledge-box/htmx-go

go 1.22.1
```

### Install Air (for hot-reloading)

It's a torture to restart the server everytime there's a small change, hence, we need this module.

#### install the module

```bash
go install github.com/cosmtrek/air@latest
```

#### create its configuration file

```bash
air init
```

#### replace the following lines

```toml
tmp_dir = "temp"

[build]
bin = "temp\\main.exe"
cmd = "go build -o temp/main.exe cmd/main.go"
```

run the hot-reload module from the root directory

```bash
air
```
## Coding

ask chatgpt

## Remarks

`Go`lang bad; HTMX good.
