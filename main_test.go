package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestCreateBCryptEndpoint(t *testing.T) {
	router := SetupRouter()

	passwordJSON := "{\"password\":\"Password1\"}"

	var reader io.Reader
	reader = strings.NewReader(passwordJSON)

	w := performRequest(router, "POST", "/bcrypt", reader)

	assert.Equal(t, http.StatusCreated, w.Code)
}