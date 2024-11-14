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

}

func (r *ResInspection) HandleETag(ctx *gin.Context) {

}

func (r *ResInspection) HandleResponseHeaders(ctx *gin.Context) {

}
