package main

import (
	"fmt"
	"net/http"
	//	"mime"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	message := "This is a test call"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
