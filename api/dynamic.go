package api

import (
	"crypto/rand"
	"encoding/base64"
	"endpointlab/utils"
	"net/http"
	"strconv"
	"time"

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
	delay := c.Param("delay")
	delayInt, err := strconv.Atoi(delay)
	if err != nil || delayInt < 0 {
		c.String(http.StatusBadRequest, "Invalid delay value")
		return
	}

	maxDelay := 10000
	timeDelay := delayInt
	if timeDelay > maxDelay {
		timeDelay = maxDelay
	}
	
	response := utils.HTTPResponse{
		Args:    utils.ConvertQuery(c.Request.URL.Query()),
		Data:    "",
		Files:   map[string]string{},
		Form:    map[string]string{},
		Headers: utils.ConvertHeaders(c.Request.Header),
		JSON:    nil,
		Origin:  c.ClientIP(),
		URL:     c.Request.URL.String(),
		Method:  c.Request.Method,
	}


	time.Sleep(time.Duration(timeDelay) * time.Millisecond)
	c.JSON(http.StatusOK, response)
}

func (d *Dynamic) HandleDrip(c *gin.Context) {
	// Parse query parameters with defaults
	duration := utils.GetQueryInt(c, "duration", 2)
	numbytes := utils.GetQueryInt(c, "numbytes", 10)
	code := utils.GetQueryInt(c, "code", 200)
	delay := utils.GetQueryInt(c, "delay", 2)

	// Validate parameters
	if duration <= 0 || numbytes <= 0 || delay < 0 {
		c.String(http.StatusBadRequest, "Invalid parameters")
		return
	}

	// Apply initial delay
	time.Sleep(time.Duration(delay) * time.Second)

	// Calculate delay between drips
	chunks := 2
	bytesPerChunk := numbytes / chunks
	delayPerChunk := duration / chunks

	// Set headers for chunked transfer
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", strconv.Itoa(numbytes))

	// Start chunked response
	c.Status(code)

	for i := 0; i < chunks; i++ {
		// Generate random bytes for this chunk
		chunk := make([]byte, bytesPerChunk)
		rand.Read(chunk)

		// Write chunk
		c.Writer.Write(chunk)
		c.Writer.Flush()

		// Wait before next chunk (except for last iteration)
		if i < chunks-1 {
			time.Sleep(time.Duration(delayPerChunk) * time.Second)
		}
	}
}

func (d *Dynamic) HandleLinks(c *gin.Context) {
	// Parse path parameters
	n := c.Param("n")
	offset := c.Param("offset")
	
	numLinks, err := strconv.Atoi(n)
	if err != nil || numLinks < 0 {
		c.String(http.StatusBadRequest, "Invalid number of links")
		return
	}
	
	offsetNum, err := strconv.Atoi(offset)
	if err != nil || offsetNum < 0 {
		c.String(http.StatusBadRequest, "Invalid offset")
		return
	}

	// Build HTML response
	html := "<html><head><title>Links</title></head><body>"
	html += "<h1>Page " + offset + "</h1>"
	
	// Generate n links to next pages
	for i := 1; i <= numLinks; i++ {
		nextOffset := offsetNum + i
		html += "<p><a href='/links/" + n + "/" + strconv.Itoa(nextOffset) + 
			   "'>Link to page " + strconv.Itoa(nextOffset) + "</a></p>"
	}
	
	html += "</body></html>"

	// Send response with HTML content type
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}

func (d *Dynamic) HandleRange(c *gin.Context) {
	// Get number of bytes from URL parameter
	numbytes := c.Param("numbytes")
	n, err := strconv.Atoi(numbytes)
	if err != nil || n <= 0 {
		c.String(http.StatusBadRequest, "Invalid number of bytes")
		return
	}

	// Get optional query parameters
	chunkSize := utils.GetQueryInt(c, "chunk_size", 10240) // Default 10KB chunks
	duration := utils.GetQueryInt(c, "duration", 1)        // Default 1 second total duration

	// Validate parameters
	if chunkSize <= 0 || duration <= 0 {
		c.String(http.StatusBadRequest, "Invalid chunk_size or duration")
		return
	}

	// Calculate delay between chunks
	totalChunks := (n + chunkSize - 1) / chunkSize // Round up division
	delayBetweenChunks := time.Duration(duration) * time.Second / time.Duration(totalChunks)

	// Set up streaming response
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", strconv.Itoa(n))
	c.Status(http.StatusOK)

	bytesRemaining := n
	for bytesRemaining > 0 {
		// Calculate size of next chunk
		currentChunkSize := chunkSize
		if bytesRemaining < chunkSize {
			currentChunkSize = bytesRemaining
		}

		// Generate and write random bytes
		chunk := make([]byte, currentChunkSize)
		_, err := rand.Read(chunk)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Writer.Write(chunk)
		c.Writer.Flush()

		bytesRemaining -= currentChunkSize

		// Delay before next chunk (skip delay for last chunk)
		if bytesRemaining > 0 {
			time.Sleep(delayBetweenChunks)
		}
	}
}

func (d *Dynamic) HandleStreamBytes(c *gin.Context) {

}

func (d *Dynamic) HandleStream(c *gin.Context) {

}

func (d *Dynamic) HandleUuid(c *gin.Context) {

}
