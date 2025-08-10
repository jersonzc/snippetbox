package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"snippetbox/internal/assert"
	"testing"
)

const healthMsg = "OK"

func TestSecureHeaders(t *testing.T) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(healthMsg))
	})

	secureHeaders(next).ServeHTTP(res, req)

	ans := res.Result()

	// Check Content-Security-Policy
	expectedValue := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, ans.Header.Get("Content-Security-Policy"), expectedValue)

	// Check Referrer-Policy
	expectedValue = "origin-when-cross-origin"
	assert.Equal(t, ans.Header.Get("Referrer-Policy"), expectedValue)

	// Check X-Content-Type-Options
	expectedValue = "nosniff"
	assert.Equal(t, ans.Header.Get("X-Content-Type-Options"), expectedValue)

	// Check X-Frame-Options
	expectedValue = "deny"
	assert.Equal(t, ans.Header.Get("X-Frame-Options"), expectedValue)

	// Check X-XSS-Protection
	expectedValue = "0"
	assert.Equal(t, ans.Header.Get("X-XSS-Protection"), expectedValue)

	// Status code
	assert.Equal(t, ans.StatusCode, http.StatusOK)

	// Body
	defer ans.Body.Close()
	body, err := io.ReadAll(ans.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), healthMsg)
}
