# ascii-art-web

## Description

`ascii-art-web` is a web-based version of the `ascii-art` project. It runs an HTTP server, written in Go, that lets users generate ASCII art banners from text directly in their browser. Users can enter any text and choose between three banner styles — **standard**, **shadow**, and **thinkertoy** — and the server renders the corresponding ASCII art and displays it on the page.

The project is built entirely using Go's standard library, with no external dependencies.

## Authors

- Shady Elmougy
-Abdullah Essam

## Usage

### Prerequisites
- [Go](https://go.dev/dl/) installed (version 1.18 or later recommended)

### Running the server

1. Clone or download this repository.
2. Open a terminal in the project's root directory (where `go.mod` is located).
3. Run the server:

```bash
   go run .
```

4. You should see the following output in your terminal:
Server starting on http://localhost:8080
5. Open your browser and go to:
http://localhost:8080
6. Enter some text, select a banner style (Standard, Shadow, or Thinkertoy), and click **Generate** to see the ASCII art result.

## Implementation details

### HTTP Server & Routing

The server is built using Go's `net/http` standard library package. It exposes two endpoints:

- **`GET /`** — serves the main page containing the input form.
- **`POST /ascii-art`** — receives the submitted text and banner choice, generates the ASCII art, and returns the page with the result included.

Since the root path `"/"` acts as a catch-all in Go's default router, the home handler explicitly checks that the requested path is exactly `"/"`, returning a `404 Not Found` for any other unmatched route.

### Templates

HTML is served using Go's `html/template` package, which automatically escapes dynamic content for safety and separates presentation (HTML) from logic (Go code). The template is parsed once at server startup for efficiency, and reused across every request. The single template file (`templates/index.html`) is reused for both the initial empty form and the page showing a result, using a conditional (`{{if .Result}}`) so the result section only appears once ASCII art has actually been generated.

### ASCII Art Conversion Algorithm

The core conversion logic works as follows:

1. **Input validation** — the submitted text is checked to ensure it only contains printable ASCII characters (space through `~`). Non-printable or invalid characters are rejected with a `400 Bad Request`.
2. **Banner selection** — the chosen banner name is validated against the three allowed options (`standard`, `shadow`, `thinkertoy`) before being mapped to its corresponding `.txt` file.
3. **Banner parsing** — the selected banner file is read from disk and parsed into a map, associating each printable ASCII character with its 8-line ASCII-art representation.
4. **Line splitting** — the input text is split on literal `\n` sequences, allowing multiple lines of output from a single input.
5. **Rendering** — for each line of input, the corresponding 8-row ASCII-art representation of each character is assembled row by row into the final output string.
6. **Response** — the rendered result is passed into the HTML template and returned to the browser inside a `<pre>` tag, which preserves the exact spacing and line breaks required for the ASCII art to display correctly.

### Error Handling

The server returns appropriate HTTP status codes throughout:

- **200 OK** — successful rendering of the home page or ASCII art result.
- **400 Bad Request** — invalid input text or invalid banner selection.
- **404 Not Found** — unknown routes, or a missing banner file.
- **500 Internal Server Error** — unexpected failures, such as a malformed banner file.