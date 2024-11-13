package api

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ContentLengthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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