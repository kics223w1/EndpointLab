package api

import (
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

}
