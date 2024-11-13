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

func (h *HttpFormat) HandleDeny(c *gin.Context) {

	// Set content type to HTML and return the content
	c.Header("Content-Type", "text/plain")
	c.Data(http.StatusOK, "text/plain", []byte(utils.AngryASCII))
}


func (h *HttpFormat) HandleHtml(c *gin.Context) {
	html, err := os.ReadFile("templates/demo.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read HTML file"})
		return
	}

	c.Header("Content-Type", "text/html")
	c.Data(http.StatusOK, "text/html", html)
}	

func (h *HttpFormat) HandleXML(c *gin.Context) {
	xml, err := os.ReadFile("templates/demo.xml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read HTML file"})
		return
	}

	c.Header("Content-Type", "application/xml")
	c.Data(http.StatusOK, "application/xml", xml)
}	

func (h *HttpFormat) HandleJson(c *gin.Context) {
	json, err := os.ReadFile("templates/demo.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "application/json", json)
}

func (h *HttpFormat) handleUTF8(c *gin.Context) {
	html, err := os.ReadFile("templates/UTF-8-demo.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read HTML file"})
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}


func (h *HttpFormat) HandleRobotTxt(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.Data(http.StatusOK, "text/plain", []byte(utils.RobotTXT))
}



