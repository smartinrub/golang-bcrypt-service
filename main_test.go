package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestBCryptEndpointToReturnCreatedStatusAndBCryptPassword(t *testing.T) {
	router := SetupRouter()

	passwordJSON := "{\"password\":\"Password1\"}"

	var reader io.Reader
	reader = strings.NewReader(passwordJSON)

	w := performRequest(router, "POST", "/bcrypt", reader)

	response := w.Body.String()
	err := bcrypt.CompareHashAndPassword([]byte(response), []byte("Password1"))

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Nil(t, err)
}
