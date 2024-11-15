package api

import (
	"encoding/json"
	"endpointlab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResInspection struct {}

func NewResInspection() *ResInspection {
	return &ResInspection{}
}

func (r *ResInspection) HandleCache(ctx *gin.Context) {
	// Check for conditional request headers
	ifModifiedSince := ctx.GetHeader("If-Modified-Since")
	ifNoneMatch := ctx.GetHeader("If-None-Match")

	// If either conditional header is present, return 304 Not Modified
	if ifModifiedSince != "" || ifNoneMatch != "" {
		ctx.Status(304) // Not Modified
		return
	}

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	fullURL := scheme + "://" + ctx.Request.Host + ctx.Request.URL.String()

	response := HTTPMethodResponse{
		Args:    utils.ConvertQuery(ctx.Request.URL.Query()),
		Headers: utils.ConvertHeaders(ctx.Request.Header),
		JSON:    nil,
		Origin:  ctx.ClientIP(),
		URL:     fullURL,
		Method:  ctx.Request.Method,
	}
	ctx.JSON(http.StatusOK, response)
}

func (r *ResInspection) HandleCacheValue(ctx *gin.Context) {
	// Get the value parameter from the URL
	value := ctx.Param("value")
	
	// Set Cache-Control header
	ctx.Header("Cache-Control", "public, max-age="+value)
	
	// Build the response as in HandleCache
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	fullURL := scheme + "://" + ctx.Request.Host + ctx.Request.URL.String()

	response := HTTPMethodResponse{
		Args:    utils.ConvertQuery(ctx.Request.URL.Query()),
		Headers: utils.ConvertHeaders(ctx.Request.Header),
		JSON:    nil,
		Origin:  ctx.ClientIP(),
		URL:     fullURL,
		Method:  ctx.Request.Method,
	}
	ctx.JSON(http.StatusOK, response)
}

func (r *ResInspection) HandleETag(ctx *gin.Context) {
	etag := ctx.Param("etag")
	
	// Get If-None-Match and If-Match headers
	ifNoneMatch := utils.ParseMultiValueHeader(ctx.GetHeader("If-None-Match"))
	ifMatch := utils.ParseMultiValueHeader(ctx.GetHeader("If-Match"))
	
	// Handle If-None-Match first
	if len(ifNoneMatch) > 0 {
		for _, value := range ifNoneMatch {
			if value == etag || value == "*" {
				ctx.Header("ETag", etag)
				ctx.Status(http.StatusNotModified)
				return
			}
		}
	} else if len(ifMatch) > 0 { // Only check If-Match if If-None-Match is not present
		matches := false
		for _, value := range ifMatch {
			if value == etag || value == "*" {
				matches = true
				break
			}
		}
		if !matches {
			ctx.Status(http.StatusPreconditionFailed)
			return
		}
	}
	
	// Normal response with ETag header
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	fullURL := scheme + "://" + ctx.Request.Host + ctx.Request.URL.String()

	response := HTTPMethodResponse{
		Args:    utils.ConvertQuery(ctx.Request.URL.Query()),
		Headers: utils.ConvertHeaders(ctx.Request.Header),
		JSON:    nil,
		Origin:  ctx.ClientIP(),
		URL:     fullURL,
		Method:  ctx.Request.Method,
	}
	
	ctx.Header("ETag", etag)
	ctx.JSON(http.StatusOK, response)
}

func (r *ResInspection) HandleResponseHeaders(ctx *gin.Context) {
	// Get only the "freeform" query parameters
	freeformValues := ctx.Request.URL.Query()["freeform"]
	
	// Set response headers from freeform values
	for _, value := range freeformValues {
		ctx.Header("freeform", value)
	}

	// Set Content-Type header
	ctx.Header("Content-Type", "application/json")

	// Build response containing all response headers
	headers := make(map[string]interface{})
	for key := range ctx.Writer.Header() {
		values := ctx.Writer.Header()[key]
		if len(values) == 1 {
			headers[key] = values[0]
		} else {
			headers[key] = values
		}
	}


	// Convert response to JSON bytes to get content length
	jsonData, _ := json.Marshal(headers)
	headers["Content-Length"] = len(jsonData)

	ctx.JSON(http.StatusOK, headers)
}