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

// @Summary      Returns an image based on the Accept header
// @Description  Returns a simple image of the type suggested by the Accept header
// @Tags         Images
// @Produce      image/webp, image/png, image/jpeg, image/svg+xml
// @Success      200  {file}  file
// @Failure      406  {object}  object
// @Failure      500  {object}  object
// @Router       /image [get]
func (h *HttpImage) HandleImage(c *gin.Context) {
	acceptHeader := c.GetHeader("Accept")

	var imagePath, contentType string

	switch {
	case acceptHeader == "image/webp":
		imagePath = fmt.Sprintf("%s/wolf_1.webp", h.imagePath)
		contentType = "image/webp"
	case acceptHeader == "image/png":
		imagePath = fmt.Sprintf("%s/pig_icon.png", h.imagePath)
		contentType = "image/png"
	case acceptHeader == "image/jpeg":
		imagePath = fmt.Sprintf("%s/jackal.jpg", h.imagePath)
		contentType = "image/jpeg"
	case acceptHeader == "image/svg+xml":
		imagePath = fmt.Sprintf("%s/svg_logo.svg", h.imagePath)
		contentType = "image/svg+xml"
	default:
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Unsupported media type"})
		return
	}

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Image not found at path: %s", imagePath)})
		return
	}

	image, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read image file: %v", err)})
		return
	}

	c.Header("Content-Type", contentType)
	c.Header("Content-Length", fmt.Sprint(len(image)))
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Data(http.StatusOK, contentType, image)
}

// @Summary      Returns a WEBP image
// @Description  Returns a simple WEBP image
// @Tags         Images
// @Produce      image/webp
// @Success      200  {file}  file
// @Router       /image/webp [get]
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

// @Summary      Returns a PNG image
// @Description  Returns a simple PNG image
// @Tags         Images
// @Produce      image/png
// @Success      200  {file}  file
// @Router       /image/png [get]
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

// @Summary      Returns a JPEG image
// @Description  Returns a simple JPEG image
// @Tags         Images
// @Produce      image/jpeg
// @Success      200  {file}  file
// @Router       /image/jpeg [get]
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

// @Summary      Returns an SVG image
// @Description  Returns a simple SVG image
// @Tags         Images
// @Produce      image/svg+xml
// @Success      200  {file}  file
// @Router       /image/svg [get]
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

