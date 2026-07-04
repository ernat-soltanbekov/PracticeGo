package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// HandleHome serves the main page
func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	templatePath := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// HandleASCIIArt processes the form submission and returns ASCII art
func HandleASCIIArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data with a limit of 10MB
	// For multipart forms, we need to explicitly parse multipart
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil && err.Error() != "request Content-Type isn't multipart/form-data" {
		// If it's not a multipart form error, try regular ParseForm
		if err = r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	text = normalizeInput(text)

	// Validate input
	if text == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if containsUnsupportedCharacters(text) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Validate banner
	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	if banner == "" {
		banner = "standard"
	}

	if !validBanners[banner] {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Generate ASCII art
	result, err := RenderASCII(text, banner)
	if err != nil {
		// Check if it's a file not found error (banner file missing)
		if os.IsNotExist(err) || strings.Contains(err.Error(), "cannot read banner file") {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		log.Printf("RenderASCII error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, result)
}

func containsUnsupportedCharacters(text string) bool {
	for _, r := range text {
		if r == '\n' || r == '\r' {
			continue
		}

		if r < 32 || r > 126 {
			return true
		}
	}

	return false
}

func normalizeInput(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	return text
}
