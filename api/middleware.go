package api

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ContentLengthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip buffering for binary content types
		contentType := c.Writer.Header().Get("Content-Type")
		if isStreamingContentType(contentType) {
			c.Next()
			return
		}

		rw := &responseWriter{
			ResponseWriter: c.Writer,
			buffer:        &bytes.Buffer{},
		}
		c.Writer = rw
		c.Next()
		
		// Only set Content-Length if there's a body
		if rw.buffer.Len() > 0 {
			c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", rw.buffer.Len()))
			c.Writer.Write(rw.buffer.Bytes())
		}
	}
}

type responseWriter struct {
	gin.ResponseWriter
	buffer *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.buffer.Write(b)
}

func isStreamingContentType(contentType string) bool {
	streamingTypes := []string{
		"image/",
		"video/",
		"audio/",
		"application/octet-stream",
	}

	for _, t := range streamingTypes {
		if strings.HasPrefix(contentType, t) {
			return true
		}
	}
	return false
} 