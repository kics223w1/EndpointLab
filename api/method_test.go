package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	httpMethod := NewHttpMethod()

	r.GET("/get", httpMethod.HandleGet)
	r.POST("/post", httpMethod.HandlePost)
	r.PUT("/put", httpMethod.HandlePut)
	r.DELETE("/delete", httpMethod.HandleDelete)
	r.PATCH("/patch", httpMethod.HandlePatch)

	return r
}

func TestHandleGet(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get?id=123", nil)
	req.Header.Set("Header-Name-1", "Header-Value-1")
	req.Header.Set("Header-Name-2", "Header-Value-2")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response HTTPMethodResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)
	assert.Equal(t, "123", response.Args["id"])
	assert.Equal(t, "Header-Value-1", response.Headers["Header-Name-1"])
	assert.Equal(t, "Header-Value-2", response.Headers["Header-Name-2"])
	assert.Equal(t, "GET", response.Method)
}

func TestHandlePost(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/post", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions to check the response body
}

func TestHandlePut(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/put", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions to check the response body
}

func TestHandleDelete(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/delete", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions to check the response body
}

func TestHandlePatch(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/patch", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions to check the response body
}

