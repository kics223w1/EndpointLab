package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type HttpImage struct {
	imagePath string
}

func NewHttpImage() *HttpImage {
	basePath := os.Getenv("IMAGE_PATH")
	if basePath == "" {
		basePath = "templates/images"
	}
	return &HttpImage{imagePath: basePath}
}

func (h *HttpImage) HandleImage(c *gin.Context) {
	imagePath := fmt.Sprintf("%s/wolf_1.webp", h.imagePath)
	
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Image not found at path: %s", imagePath)})
		return
	}
	
	image, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read image file: %v", err)})
		return
	}

	c.Header("Content-Type", "image/webp")
	c.Header("Content-Length", fmt.Sprint(len(image)))
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Data(http.StatusOK, "image/webp", image)
}

func (h *HttpImage) HandleImagePNG(c *gin.Context) {
	imagePath := fmt.Sprintf("%s/pig_icon.png", h.imagePath)
	
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Image not found at path: %s", imagePath)})
		return
	}
	
	image, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read image file: %v", err)})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", fmt.Sprint(len(image)))
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Data(http.StatusOK, "image/png", image)
}

func (h *HttpImage) HandleImageJPEG(c *gin.Context) {
	imagePath := fmt.Sprintf("%s/jackal.jpg", h.imagePath)
	
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Image not found at path: %s", imagePath)})
		return
	}
	
	image, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read image file: %v", err)})
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Length", fmt.Sprint(len(image)))
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Data(http.StatusOK, "image/jpeg", image)
}

func (h *HttpImage) HandleImageSVG(c *gin.Context) {
	imagePath := fmt.Sprintf("%s/svg_logo.svg", h.imagePath)
	
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Image not found at path: %s", imagePath)})
		return
	}
	
	image, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read image file: %v", err)})
		return
	}

	c.Header("Content-Type", "image/svg+xml")
	c.Header("Content-Length", fmt.Sprint(len(image)))
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Data(http.StatusOK, "image/svg+xml", image)
}

func (h *HttpImage) HandleImageWebp(c *gin.Context) {
	imagePath := fmt.Sprintf("%s/wolf_1.webp", h.imagePath)
	
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Image not found at path: %s", imagePath)})
		return
	}
	
	image, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read image file: %v", err)})
		return
	}

	c.Header("Content-Type", "image/webp")
	c.Header("Content-Length", fmt.Sprint(len(image)))
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Data(http.StatusOK, "image/webp", image)
}

