package api

import (
	"encoding/json"
	"endpointlab/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResInspection struct {}

func NewResInspection() *ResInspection {
	return &ResInspection{}
}

//	@Summary		Returns a 304 if an If-Modified-Since header or If-None-Match is present. Returns the same as a GET otherwise.
//	@Tags			Response Inspection
//	@Param			If-Modified-Since	header	string	false	"Optional header to check if the resource has been modified since the specified date"
//	@Param			If-None-Match		header	string	false	"Optional header to check if the resource matches the given ETag"
//	@Produce		json
//	@Success		200		"Cached response"
//	@Success		304		"Modified"
//	@Router			/cache [get]
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

//	@Summary		Sets a Cache-Control header for n seconds.
//	@Tags			Response Inspection
//	@Param			value	path	int	true	"Number of seconds for the Cache-Control max-age directive"
//	@Produce		json
//	@Success		200		"Cache control set"
//	@Router			/cache/{value} [get]
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

//	@Summary		Assumes the resource has the given etag and responds to If-None-Match and If-Match headers appropriately.
//	@Tags			Response Inspection
//	@Param			If-None-Match	header	string	false	"Optional header to check if the resource does not match the given ETag"
//	@Param			If-Match		header	string	false	"Optional header to check if the resource matches the given ETag"
//	@Produce		json
//	@Success		200		"Normal response"
//	@Failure		412	"match"
//	@Router			/etag/{etag} [get]
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

//	@Summary		Returns response headers.
//	@Description	Returns all response headers including freeform values.
//	@Tags			Response Inspection
//	@Accept			json
//	@Produce		json
//	@Param			freeform	query	string	false	"Freeform query parameter"
//	@Success		200		"Response headers"
//	@Router			/response-headers [get]
//	@Router			/response-headers [post]
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
			headers[strings.ToLower(key)] = values[0]
		} else {
			headers[strings.ToLower(key)] = values
		}
	}

	// Convert response to JSON bytes to get content length
	jsonData, _ := json.Marshal(headers)
	headers["Content-Length"] = len(jsonData)

	ctx.JSON(http.StatusOK, headers)
}
