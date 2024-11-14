package utils

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

func ParseMultiValueHeader(header string) []string {
    if header == "" {
        return nil
    }
    
    // Split by comma and trim whitespace and quotes
    values := strings.Split(header, ",")
    cleaned := make([]string, 0, len(values))
    
    for _, value := range values {
        value = strings.TrimSpace(value)
        value = strings.Trim(value, "\"")
        if value != "" {
            cleaned = append(cleaned, value)
        }
    }
    
    return cleaned
}