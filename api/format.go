package api

import (
	"bytes"
	"encoding/json"
	"endpointlab/utils"
	"net/http"
	"os"

	"compress/gzip"

	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/zlib"

	"github.com/gin-gonic/gin"
)

type HttpFormat struct {
}

type FormatResponse struct {
	Brotli bool `json:"brotli,omitempty"`
	Gzip   string `json:"gzip,omitempty"`
	Deflate bool `json:"deflate,omitempty"`
	Headers map[string]string `json:"headers"`
	Method  string           `json:"method"`
	Origin  string           `json:"origin"`
}

func NewHttpFormat() *HttpFormat {
	return &HttpFormat{}
}

//	@Summary		Handle Brotli compression.
//	@Description	Returns a Brotli compressed response.
//	@Tags			Response Formats
//	@Accept			json
//	@Produce		json
//	@Success		200	 "Brotli-encoded data."
//	@Router			/brotli [get]
func (h *HttpFormat) HandleBrotli(c *gin.Context) {
	response := FormatResponse{
		Brotli:  true,
		Headers: utils.ConvertHeaders(c.Request.Header),
		Origin:  c.ClientIP(),
		Method:  c.Request.Method,
	}
	
	// Convert response to JSON bytes
	jsonData, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal response"})
		return
	}

	// Create a buffer for Brotli compressed data
	var compressed bytes.Buffer
	writer := brotli.NewWriter(&compressed)
	
	// Write and flush the compressed data
	if _, err := writer.Write(jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Compression failed"})
		return
	}
	if err := writer.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Compression failed"})
		return
	}

	// Set appropriate headers
	c.Header("Content-Encoding", "br")
	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "application/json", compressed.Bytes())
}

//	@Summary		Handle Deflate compression.
//	@Description	Returns a Deflate compressed response.
//	@Tags			Response Formats
//	@Accept			json
//	@Produce		json
//	@Success		200	 "Deflate-encoded data."
//	@Router			/deflate [get]
func (h *HttpFormat) HandleDeflate(c *gin.Context) {
	response := FormatResponse{
		Deflate: true,
		Headers: utils.ConvertHeaders(c.Request.Header),
		Origin:  c.ClientIP(),
		Method:  c.Request.Method,
	}
	
	// Convert response to JSON bytes
	jsonData, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal response"})
		return
	}

	// Create a buffer for Deflate compressed data
	var compressed bytes.Buffer
	writer := zlib.NewWriter(&compressed)
	
	// Write and flush the compressed data
	if _, err := writer.Write(jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Compression failed"})
		return
	}
	if err := writer.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Compression failed"})
		return
	}

	// Set appropriate headers
	c.Header("Content-Encoding", "deflate")
	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "application/json", compressed.Bytes())
}

//	@Summary		Handle Deny response.
//	@Description	Returns a plain text response with denial content.
//	@Tags			Response Formats
//	@Accept			plain
//	@Produce		plain
//	@Success		200	{string} string "Deny response."
//	@Router			/deny [get]
func (h *HttpFormat) HandleDeny(c *gin.Context) {

	// Set content type to HTML and return the content
	c.Header("Content-Type", "text/plain")
	c.Data(http.StatusOK, "text/plain", []byte(utils.AngryASCII))
}

//	@Summary		Serve UTF-8 text content.
//	@Description	Returns UTF-8 encoded text content from a file.
//	@Tags			Response Formats
//	@Accept			plain
//	@Produce		plain
//	@Success		200	"UTF-8 encoded data."
//	@Router			/encoding/utf8 [get]
func (h *HttpFormat) handleUTF8(c *gin.Context) {
	html, err := os.ReadFile("templates/UTF-8-demo.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read HTML file"})
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}

//	@Summary		Handle Gzip compression.
//	@Description	Returns a Gzip compressed response.
//	@Tags			Response Formats
//	@Accept			json
//	@Produce		json
//	@Success		200 "Gzip-encoded data."
//	@Router			/gzip [get]
func (h *HttpFormat) HandleGzip(c *gin.Context) {
	response := FormatResponse{
		Gzip:    "true",
		Headers: utils.ConvertHeaders(c.Request.Header),
		Origin:  c.ClientIP(),
		Method:  c.Request.Method,
	}
	
	// Convert response to JSON bytes
	jsonData, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal response"})
		return
	}

	// Create a buffer for Gzip compressed data
	var compressed bytes.Buffer
	writer := gzip.NewWriter(&compressed)
	
	// Write and flush the compressed data
	if _, err := writer.Write(jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Compression failed"})
		return
	}
	if err := writer.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Compression failed"})
		return
	}

	// Set appropriate headers
	c.Header("Content-Encoding", "gzip")
	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "application/json", compressed.Bytes())
}

//	@Summary		Serve HTML content.
//	@Description	Returns HTML content from a file.
//	@Tags			Response Formats
//	@Accept			html
//	@Produce		html
//	@Success		200	 "HTML content."
//	@Router			/html [get]
func (h *HttpFormat) HandleHtml(c *gin.Context) {
	html, err := os.ReadFile("templates/demo.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read HTML file"})
		return
	}

	c.Header("Content-Type", "text/html")
	c.Data(http.StatusOK, "text/html", html)
}	


//	@Summary		Serve JSON content.
//	@Description	Returns JSON content from a file.
//	@Tags			Response Formats
//	@Accept			json
//	@Produce		json
//	@Success		200 "JSON content."
//	@Router			/json [get]
func (h *HttpFormat) HandleJson(c *gin.Context) {
	json, err := os.ReadFile("templates/demo.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "application/json", json)
}

//	@Summary		Serve robots.txt content.
//	@Description	Returns the content of robots.txt.
//	@Tags			Response Formats
//	@Accept			plain
//	@Produce		plain
//	@Success		200 "robots.txt content."
//	@Router			/robots.txt [get]
func (h *HttpFormat) HandleRobotTxt(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.Data(http.StatusOK, "text/plain", []byte(utils.RobotTXT))
}

//	@Summary		Serve XML content.
//	@Description	Returns XML content from a file.
//	@Tags			Response Formats
//	@Accept			xml
//	@Produce		xml
//	@Success		200 "XML content."
//	@Router			/xml [get]
func (h *HttpFormat) HandleXML(c *gin.Context) {
	xml, err := os.ReadFile("templates/demo.xml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read HTML file"})
		return
	}

	c.Header("Content-Type", "application/xml")
	c.Data(http.StatusOK, "application/xml", xml)
}	