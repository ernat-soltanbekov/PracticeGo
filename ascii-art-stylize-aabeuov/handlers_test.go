package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandleASCIIArtRejectsUnsupportedInput(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader("text=こんにちは&banner=standard"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	HandleASCIIArt(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	if strings.TrimSpace(recorder.Body.String()) != "bad request" {
		t.Fatalf("expected generic bad request message, got %q", recorder.Body.String())
	}
}

func TestHandleASCIIArtAllowsNewlines(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader("text=Hello%0D%0AWorld&banner=standard"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	HandleASCIIArt(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	if body := strings.TrimSpace(recorder.Body.String()); body == "" {
		t.Fatalf("expected rendered ascii art output, got %q", body)
	}
}

func TestHandleHomeReturnsInternalServerErrorWhenTemplateMissing(t *testing.T) {
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}

	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("failed to chdir to temp dir: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(originalDir)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	HandleHome(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, recorder.Code)
	}

	if !strings.Contains(recorder.Body.String(), "Internal server error") {
		t.Fatalf("expected internal server error response, got %q", recorder.Body.String())
	}
}
