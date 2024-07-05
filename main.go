package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/json", helloJSONHandler)
	fmt.Println("Starting server with address: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func helloHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hello, Go World!")
}

func helloJSONHandler(responseWriter http.ResponseWriter, request *http.Request) {
	msg := Message{
		Greeting: "Hello, Go World in JSON!",
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(msg)
}

type Message struct {
	Greeting string `json:"greeting"`
}
