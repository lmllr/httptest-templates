package main

import (
	"net/http"
	"test-handlers/handler"
)

func main() {
	http.HandleFunc("/", handler.TestHandler)

	http.ListenAndServe(":8080", nil)
}
