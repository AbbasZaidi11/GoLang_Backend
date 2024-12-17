# Simple HTTP Server in Go

This project is a basic HTTP server written in Go, designed to help beginners understand and learn the fundamentals of web server development. The server supports:
- Serving static files from the `./static` directory.
- Handling form submissions via the `/form` endpoint.
- Responding to GET requests at the `/hello` endpoint.

---

## Features

### 1. **Static File Serving**
- Serves files from the `./static` directory.
- Example: If `./static/index.html` exists, it will be accessible at `http://localhost:8080/`.

### 2. **Form Handling**
- Endpoint: `/form`
- Accepts `POST` requests with form data.
- Extracts and displays the `name` and `address` fields from the form.

### 3. **Custom Endpoint**
- Endpoint: `/hello`
- Responds to `GET` requests with a friendly "hello!" message.
- Returns a `404` error if the path is not `/hello` or the method is not `GET`.

---

## Prerequisites

- Go installed (version 1.19 or above recommended).
- Basic understanding of Go's `net/http` package.

---

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/go-http-server.git
cd go-http-server
