package api

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Dynamic struct {
}

func NewDynamic() *Dynamic {
	return &Dynamic{}
}


func (d *Dynamic) HandleBase64(c *gin.Context) {
	// Get the base64 value from URL parameter
	value := c.Param("value")
	
	// Decode base64 string
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		c.String(http.StatusBadRequest, "Incorrect Base64 data please try: RW5kcG9pbnRMYWIgaXMgYXdlc29tZQ==")
		return
	}
	
	// Return decoded value in response body
	c.String(http.StatusOK, string(decoded))
}

func (d *Dynamic) HandleBytes(c *gin.Context) {
	// Get the number of bytes from URL parameter
	n := c.Param("n")
	numBytes, err := strconv.Atoi(n)
	if err != nil || numBytes < 0 {
		c.String(http.StatusBadRequest, "Invalid number of bytes requested")
		return
	}

	// Generate random bytes
	randomBytes := make([]byte, numBytes)
	_, err = rand.Read(randomBytes)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate random bytes")
		return
	}

	// Set content type and send response
	c.Header("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/octet-stream", randomBytes)
}

func (d *Dynamic) HandleDeplay(c *gin.Context) {

}

func (d *Dynamic) HandleDrip(c *gin.Context) {

}

func (d *Dynamic) HandleLinks(c *gin.Context) {

}

func (d *Dynamic) HandleRange(c *gin.Context) {

}

func (d *Dynamic) HandleStreamBytes(c *gin.Context) {

}

func (d *Dynamic) HandleStream(c *gin.Context) {

}

func (d *Dynamic) HandleUuid(c *gin.Context) {

}
