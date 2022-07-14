package main

import (
	"log"
	"net/http"
)

// create a hander struct
type HttpHandler struct{}

// define and initialize const with a value
const PORT = ":8080"

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// create response data
	data := []byte("Hello, World!\n")

	// write `data` to response
	res.Write(data)
}

func main() {

	// create a new handler
	handler := HttpHandler{}

	// log the message on starting server
	log.Printf("Server starting on port %v\n", 8080)

	// listen and serve
	http.ListenAndServe(PORT, handler)
}
