package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:")
	log.Printf("Method: %s", r.Method)
	log.Printf("URL: %s", r.URL.String())
	log.Println("Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			log.Printf("  %s: %s", name, value)
		}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body:", err)
	} else {
		log.Println("Body:")
		log.Println(string(body))
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "Request received")
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
