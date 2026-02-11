package tests

import (
	"ascii-art-web/internals/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"net/url"
)

func TestHandleAsciiArt_MissingContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.HandleAsciiArt)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}


func TestServeTemplate_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.ServeTemplate)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rr.Code)
	}
}

func TestServeTemplate_ContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.ServeTemplate)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	ct := rr.Header().Get("Content-Type")
	if ct == "" {
		t.Fatal("expected Content-Type header to be set")
	}
}

func TestHandleAsciiArt_ValidContentType(t *testing.T) {
	form := url.Values{}
	form.Set("artstyle", "standard")
	form.Set("text", "Hello")

	req := httptest.NewRequest(
		http.MethodPost,
		"/ascii-art",
		strings.NewReader(form.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.HandleAsciiArt)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}