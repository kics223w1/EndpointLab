package utils

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertHeaders(header http.Header) map[string]string {
	headers := make(map[string]string)
	for key, values := range header {
		if len(values) > 0 && len(values) == 1 {
			headers[key] = values[0]
		}
	}
	return headers
}

func ConvertQuery(query url.Values) map[string]string {
	queries := make(map[string]string)
	for key, values := range query {    
		if len(values) > 0 && len(values) == 1 {
			queries[key] = values[0]
            }
      }
      return queries
}


func GetQueryInt(c *gin.Context, key string, defaultValue int) int {
	value, err := strconv.Atoi(c.Query(key))
	if err != nil {
		return defaultValue
	}
	return value
}