package api

import (
	cryptorand "crypto/rand"
	"encoding/base64"
	"endpointlab/utils" // Add this import
	"fmt"
	mathrand "math/rand" // Update alias to match usage
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpDynamic struct {
}

type streamJSONResponse struct {
	utils.HTTPResponse
	ID int `json:"id"`
}


func NewHttpDynamic() *HttpDynamic {
	return &HttpDynamic{}
}


func (d *HttpDynamic) HandleBase64(c *gin.Context) {
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

func (d *HttpDynamic) HandleBytes(c *gin.Context) {
	// Get the number of bytes from URL parameter
	n := c.Param("n")
	numBytes, err := strconv.Atoi(n)
	if err != nil || numBytes < 0 {
		c.String(http.StatusBadRequest, "Invalid number of bytes requested")
		return
	}

	// Generate random bytes
	randomBytes := make([]byte, numBytes)
	_, err = cryptorand.Read(randomBytes)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate random bytes")
		return
	}

	// Set content type and send response
	c.Header("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/octet-stream", randomBytes)
}

func (d *HttpDynamic) HandleDeplay(c *gin.Context) {
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

func (d *HttpDynamic) HandleDrip(c *gin.Context) {
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
		cryptorand.Read(chunk)

		// Write chunk
		c.Writer.Write(chunk)
		c.Writer.Flush()

		// Wait before next chunk (except for last iteration)
		if i < chunks-1 {
			time.Sleep(time.Duration(delayPerChunk) * time.Second)
		}
	}
}

func (d *HttpDynamic) HandleLinks(c *gin.Context) {
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

func (d *HttpDynamic) HandleRange(c *gin.Context) {
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
		_, err := cryptorand.Read(chunk)
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

func (d *HttpDynamic) HandleStreamBytes(c *gin.Context) {
	// Get number of bytes from URL parameter
	n := c.Param("n")
	numBytes, err := strconv.Atoi(n)
	if err != nil || numBytes <= 0 {
		c.String(http.StatusBadRequest, "Invalid number of bytes")
		return
	}

	// Limit to 100KB like the Python version
	if numBytes > 100*1024 {
		numBytes = 100 * 1024
	}

	// Get optional query parameters
	chunkSize := utils.GetQueryInt(c, "chunk_size", 10*1024) // Default 10KB chunks like Python
	if chunkSize < 1 {
		chunkSize = 1
	}

	// Set up random source
	var randSource mathrand.Rand
	if seed := utils.GetQueryInt(c, "seed", -1); seed != -1 {
		randSource = *mathrand.New(mathrand.NewSource(int64(seed)))
	} else {
		randSource = *mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	}

	// Set up streaming response
	c.Header("Content-Type", "application/octet-stream")
	c.Status(http.StatusOK)

	// Generate and stream bytes in chunks
	chunks := make([]byte, 0, chunkSize)
	for i := 0; i < numBytes; i++ {
		chunks = append(chunks, byte(randSource.Intn(256)))
		if len(chunks) == chunkSize {
			c.Writer.Write(chunks)
			c.Writer.Flush()
			chunks = make([]byte, 0, chunkSize)
		}
	}

	// Send any remaining bytes
	if len(chunks) > 0 {
		c.Writer.Write(chunks)
		c.Writer.Flush()
	}
}

func (d *HttpDynamic) HandleStream(c *gin.Context) {
	// Parse number of responses from URL parameter
	n := c.Param("n")
	count, err := strconv.Atoi(n)
	if err != nil || count <= 0 {
		c.String(http.StatusBadRequest, "Invalid number of responses")
		return
	}

	// Set headers for streaming JSON
	c.Header("Content-Type", "application/json")
	c.Header("Transfer-Encoding", "chunked")

	

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	fullURL := scheme + "://" + c.Request.Host + c.Request.URL.String()

	// Stream n responses
	for i := 0; i < count; i++ {
		response := streamJSONResponse{
			HTTPResponse: utils.HTTPResponse{
                Args:    utils.ConvertQuery(c.Request.URL.Query()),
                Data:    "",
                Files:   map[string]string{},
                Form:    map[string]string{},
                Headers: utils.ConvertHeaders(c.Request.Header),
                JSON:    nil,
                Origin:  c.ClientIP(),
                URL:     fullURL,
				Method:  c.Request.Method,
			},
			ID: i,
		}

		// Write each response as a separate JSON object
		c.JSON(http.StatusOK, response)
		c.Writer.Write([]byte("\n"))
		c.Writer.Flush()
	}
}


func (d *HttpDynamic) HandleUuid(c *gin.Context) {
	// Generate 16 random bytes
	uuid := make([]byte, 16)
	_, err := cryptorand.Read(uuid)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate UUID")
		return
	}

	// Set version (4) and variant (RFC 4122) bits
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant RFC 4122

	// Format as UUID string
	uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

	c.JSON(http.StatusOK, gin.H{"uuid": uuidStr})
}
