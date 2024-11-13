package handlers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type HTTPResponse struct {
	Args    map[string]string `json:"args"`
	Data    string             `json:"data,omitempty"`
	Files   map[string]string  `json:"files,omitempty"`
	Form    map[string]string  `json:"form,omitempty"`
	Headers map[string]string `json:"headers"`
	JSON    interface{}        `json:"json,omitempty"`
	Origin  string             `json:"origin"`
	URL     string             `json:"url"`
	Method  string             `json:"method"`
}

func convertHeaders(header http.Header) map[string]string {
	headers := make(map[string]string)
	for key, values := range header {
		if len(values) > 0 && len(values) == 1 {
			headers[key] = values[0]
		}
	}
	return headers
}

func convertQuery(query url.Values) map[string]string {
	queries := make(map[string]string)
	for key, values := range query {    
		if len(values) > 0 && len(values) == 1 {
			queries[key] = values[0]
            }
      }
      return queries
}

func HandleGet(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	fullURL := scheme + "://" + c.Request.Host + c.Request.URL.String()
	
	response := HTTPResponse{
		Args:    convertQuery(c.Request.URL.Query()),
		Headers: convertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     fullURL,
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

func HandlePost(c *gin.Context) {
	response := HTTPResponse{
		Args:   convertQuery(c.Request.URL.Query()),
		Data:    "",
		Files:   make(map[string]string),
		Form:    make(map[string]string),
		Headers: convertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

func HandlePut(c *gin.Context) {
	response := HTTPResponse{
		Args:    convertQuery(c.Request.URL.Query()),
		Data:    "",
		Files:   make(map[string]string),
		Form:    make(map[string]string),
		Headers: convertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

func HandleDelete(c *gin.Context) {
	response := HTTPResponse{
		Args:   convertQuery(c.Request.URL.Query()),
		Data:    "",
		Files:   make(map[string]string),
		Form:    make(map[string]string),
		Headers: convertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

func HandlePatch(c *gin.Context) {
	response := HTTPResponse{
		Args:   convertQuery(c.Request.URL.Query()),
		Data:    "",
		Files:   make(map[string]string),
		Form:    make(map[string]string),
		Headers: convertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}