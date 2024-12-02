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

// @Summary      Return status code or random status code if more than one are given
// @Description  Returns a status code based on the path
// @Tags         Status
// @Param        code  path  int  true  "HTTP Status Code"
// @Produce      json
// @Success      200  {object}  object
// @Router       /status/{code} [get]
// @Router       /status/{code} [post]
// @Router       /status/{code} [put]
// @Router       /status/{code} [delete]
// @Router       /status/{code} [patch]
// @Router       /status/{code} [options]
// @Router       /status/{code} [head]
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


