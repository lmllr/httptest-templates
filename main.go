package main

import (
	"httptest-templates/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.TestHandler)

	http.ListenAndServe(":8080", nil)
}
