package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPResponse struct {
	Args    map[string][]string `json:"args"`
	Headers map[string][]string `json:"headers"`
	Origin  string             `json:"origin"`
	URL     string             `json:"url"`
	Method  string             `json:"method"`
}

func HandleGet(c *gin.Context) {
	response := HTTPResponse{
	Args:    c.Request.URL.Query(),
		Headers: c.Request.Header,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

func HandlePost(c *gin.Context) {
	response := HTTPResponse{
		Args:    c.Request.URL.Query(),
		Headers: c.Request.Header,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

func HandlePut(c *gin.Context) {
	response := HTTPResponse{
		Args:    c.Request.URL.Query(),
		Headers: c.Request.Header,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
}

func HandleDelete(c *gin.Context) {
	response := HTTPResponse{
		Args:    c.Request.URL.Query(),
		Headers: c.Request.Header,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}
	c.JSON(http.StatusOK, response)
} 