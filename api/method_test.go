package api

import (
	"bytes"
	"encoding/json"
	"io"
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
	req, _ := http.NewRequest("POST", "/post?id=123", nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Header-Name-1", "Header-Value-1")
	
	req.Body = io.NopCloser(bytes.NewBufferString(`{"name":"John", "age":30, "city":"New York"}`))

	router.ServeHTTP(w, req)

	var response HTTPMethodResponse

	err := json.Unmarshal(w.Body.Bytes(), &response)
	expectedData := `{"name":"John", "age":30, "city":"New York"}`
	expectedJSON := map[string]interface{}{
        "name": "John",
        "age":  float64(30),
        "city": "New York",
    }

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, "123", response.Args["id"])	
	assert.Equal(t, "POST", response.Method)

	assert.Equal(t, "Header-Value-1", response.Headers["Header-Name-1"])

	assert.Equal(t, expectedData, response.Data)
	assert.Equal(t, expectedJSON, response.JSON)
}

func TestHandlePut(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/put?id=123", nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Header-Name-1", "Header-Value-1")

	req.Body = io.NopCloser(bytes.NewBufferString(`{"name":"John", "age":30, "city":"New York"}`))

	router.ServeHTTP(w, req)

	var response HTTPMethodResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	expectedData := `{"name":"John", "age":30, "city":"New York"}`
	expectedJSON := map[string]interface{}{
        "name": "John",
        "age":  float64(30),
        "city": "New York",
    }

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "123", response.Args["id"])
	assert.Equal(t, "PUT", response.Method)
	assert.Equal(t, "Header-Value-1", response.Headers["Header-Name-1"])
	assert.Equal(t, expectedData, response.Data)
	assert.Equal(t, expectedJSON, response.JSON)
}

func TestHandleDelete(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/delete", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Header-Name-1", "Header-Value-1")

	req.Body = io.NopCloser(bytes.NewBufferString(`{"name":"John", "age":30, "city":"New York"}`))

	router.ServeHTTP(w, req)

	var response HTTPMethodResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	expectedData := `{"name":"John", "age":30, "city":"New York"}`
	expectedJSON := map[string]interface{}{
        "name": "John",
        "age":  float64(30),
        "city": "New York",
    }

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "DELETE", response.Method)
	assert.Equal(t, "Header-Value-1", response.Headers["Header-Name-1"])
	assert.Equal(t, expectedData, response.Data)
	assert.Equal(t, expectedJSON, response.JSON)
}

func TestHandlePatch(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/patch", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions to check the response body
}

