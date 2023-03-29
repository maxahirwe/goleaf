package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxahirwe/goleaf/initializer"
	"github.com/maxahirwe/goleaf/models"
	"github.com/stretchr/testify/assert"
)

const baseUrl = "/api/v1/"
const basicAuthUser = "idt"
const basicAuthPass = "leaf"

// generate base 64 basicAuth key
func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// add http headers to req (basic auth & content type)
func addHeadersToRequests(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(basicAuthUser, basicAuthPass))
}

// Test if server can handle non authenticated requests
func TestNoAuth(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", baseUrl, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// Test if server allows authenticated requests
func TestAuth(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", baseUrl, nil)
	addHeadersToRequests(req)
	router.ServeHTTP(w, req)
	mockResponse := `{"message":"users endpoints"}`
	assert.Equal(t, mockResponse, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

// Test user creation endpoint
func TestCreateUser(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	values := map[string]any{"Name": "Sample Name", "SignupTime": 1679991082}
	jsonValue, _ := json.Marshal(values)
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonValue))
	addHeadersToRequests(req)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

// Test user creation endpoint: No creation of user with 'SignupTime' lesser than 1850
func TestNotCreateUserSignupTimeCheck(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	values := map[string]any{"Name": "Sample Name", "SignupTime": -3810880865000}
	jsonValue, _ := json.Marshal(values)
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonValue))
	addHeadersToRequests(req)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test user creation endpoint: No creation of user with 'Name' lesser than 2 characters
func TestNotCreateUserNameCheck(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	values := map[string]any{"Name": "S", "SignupTime": 1679991082}
	jsonValue, _ := json.Marshal(values)
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonValue))
	addHeadersToRequests(req)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test get user endpoint
func TestGetUser(t *testing.T) {
	// under assunption that at least one use exists in the db
	var user models.User
	initializer.DATABASE.Take(&user)
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%d", baseUrl, user.ID), nil)
	addHeadersToRequests(req)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// Test get user endpoint of non existing user id
func TestGetUserNotFound(t *testing.T) {
	var user models.User
	initializer.DATABASE.Take(&user)
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%d", baseUrl, -1), nil)
	addHeadersToRequests(req)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Test get users endpoint
func TestGetUsers(t *testing.T) {
	var user models.User
	initializer.DATABASE.Take(&user)
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s%s", baseUrl, "all"), nil)
	addHeadersToRequests(req)
	router.ServeHTTP(w, req)
	assert.Contains(t, w.Body.String(), "users")
	assert.Equal(t, http.StatusOK, w.Code)

}
