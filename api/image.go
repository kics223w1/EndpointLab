package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type HttpImage struct {
}

func NewHttpImage() *HttpImage {
	return &HttpImage{}
}

func (h *HttpImage) HandleImage(c *gin.Context) {
	imagePath := "templates/images/wolf_1.webp"
	image, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image file"})
		return
	}

	c.Header("Content-Type", "image/webp")
	c.Header("Content-Length", fmt.Sprint(len(image)))
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Data(http.StatusOK, "image/webp", image)
}


func (h *HttpImage) HandleImagePNG(c *gin.Context) {
	imagePath := "templates/images/pig_icon.png"
	image, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image file"})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", fmt.Sprint(len(image)))
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Data(http.StatusOK, "image/png", image)
}

func (h *HttpImage) HandleImageJPEG(c *gin.Context) {
	imagePath := "templates/images/jackal.jpg"
	image, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image file"})
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Length", fmt.Sprint(len(image)))
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Data(http.StatusOK, "image/jpeg", image)
}



