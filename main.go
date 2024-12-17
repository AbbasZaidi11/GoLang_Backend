package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler handles form submissions sent via POST requests to the /form endpoint.
// It parses form data and extracts the "name" and "address" fields.

// w is an instance of http.ResponseWriter.
// It's an interface provided by the Go net/http package.

// It allows you to construct and send the HTTP response back
// to the client (e.g., a browser or API consumer).

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request body
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Acknowledge successful POST request
	fmt.Fprintf(w, "POST request successful\n")

	// Extract form values for "name" and "address"
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Respond with the extracted values
	fmt.Fprintf(w, "Name=%s\n", name)
	fmt.Fprintf(w, "Address=%s\n", address)
	//You use w to:
	//Write the response body (e.g., fmt.Fprintf(w, "some response")).
	//Set HTTP headers (e.g., w.Header().Set("Content-Type", "text/html")).
	//Set the HTTP status code (e.g., w.WriteHeader(http.StatusOK)).

}

// helloHandler responds to GET requests to the /hello endpoint with a greeting message.
// If the path is not /hello or the method is not GET, it returns an appropriate error.

// r is a pointer to an http.Request struct.

//It represents the incoming HTTP request from the client.
// Information about the request, such as:

// HTTP method (r.Method: GET, POST, etc.).
// URL path (r.URL.Path).
// Query parameters (r.URL.Query()).
// Form data (r.FormValue("key") after parsing with r.ParseForm()).
// Headers (r.Header.Get("Header-Name")).
// Request body (r.Body).

// Customizable?

// The name r is not reserved. You can name it anything (e.g., req, request, etc.), but r is conventional.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request path is /hello
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Ensure that only GET requests are allowed
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// Respond with a greeting message
	fmt.Fprintf(w, "hello!")
}

func main() {
	// Set up a file server to serve static files from the ./static directory.
	// http.FileServer returns a handler that serves files based on the requested URL.
	fileServer := http.FileServer(http.Dir("./static"))

	// Map the root path (/) to the file server.
	// Example:
	// - If ./static contains index.html, visiting http://localhost:8080/ will serve index.html.
	// - If ./static contains other files like about.html, they can be accessed directly,
	//   e.g., http://localhost:8080/about.html.
	http.Handle("/", fileServer)

	// Map the /form endpoint to the formHandler function to handle form submissions.
	http.HandleFunc("/form", formHandler)

	// Map the /hello endpoint to the helloHandler function for GET requests.
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server on port 8080.
	fmt.Println("Starting Server at port 8080")

	// Listen and serve incoming requests. If an error occurs, log it and exit.
	err := http.ListenAndServe(":8080", nil)
	// 	This function starts an HTTP server:
	// ":8080": Specifies the server should listen on port 8080.
	// nil: Means the DefaultServeMux (default multiplexer) will handle incoming requests.

	if err != nil {
		log.Fatal(err)
	}
	//nil if the server starts successfully.
	// An error if there’s a failure (e.g., port 8080 is already in use).

	//err is a variable that holds an error returned by a function.
	//In your code, it captures the error returned by r.ParseForm().

}
