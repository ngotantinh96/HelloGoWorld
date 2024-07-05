package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", helloHandler).Methods("GET")
	r.HandleFunc("/json", helloJSONHandler).Methods("GET")
	fmt.Println("Starting server with address: http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go World!")
}

func helloJSONHandler(w http.ResponseWriter, r *http.Request) {
	msg := Message{
		Greeting: "Hello, Go World in JSON!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received: ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

type Message struct {
	Greeting string `json:"greeting"`
}
