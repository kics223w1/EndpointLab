package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HttpStatus struct {
}

func NewHttpStatus() *HttpStatus {
	return &HttpStatus{}
}

func (h *HttpStatus) HandleStatus(c *gin.Context) {
	// Get the full path
	path := c.Request.URL.Path
	
	// Extract the status code from the path
	code := path[len("/status/"):]
	
	// Convert string to integer
	statusCode := 0
	if _, err := fmt.Sscanf(code, "%d", &statusCode); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid status code format",
		})
		return
	}
	
	// Return the requested status code with an empty body
	c.Status(statusCode)
}


