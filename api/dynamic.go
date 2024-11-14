package api

import (
	"encoding/base64"
	"net/http"

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
