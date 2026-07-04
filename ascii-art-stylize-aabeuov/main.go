package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Set up HTTP routes
	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/ascii-art", HandleASCIIArt)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start server
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
