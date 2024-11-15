package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HttpRedirect struct {}

func NewHttpRedirect() *HttpRedirect {
	return &HttpRedirect{}
}

func (h *HttpRedirect) HandleAbsoluteRedirect(ctx *gin.Context) {
	// Get the number of redirects from path parameter
	n := ctx.Param("n")
	redirectCount, err := strconv.Atoi(n)
	if err != nil || redirectCount <= 0 {
		ctx.JSON(400, gin.H{"error": "Invalid redirect count"})
		return
	}

	// If this is the final redirect, redirect to /get endpoint
	if redirectCount == 1 {
		ctx.Redirect(302, "/get")
		return
	}

	// Get current host and scheme
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	host := ctx.Request.Host

	// Construct next redirect URL
	nextURL := fmt.Sprintf("%s://%s/absolute-redirect/%d", scheme, host, redirectCount-1)
	ctx.Redirect(302, nextURL)
}

func (h *HttpRedirect) HandleRedirectTo(ctx *gin.Context) {
	// Get URL from form body
	url := ctx.Request.FormValue("url")
	if url == "" {
		ctx.JSON(400, gin.H{"error": "URL is required"})
		return
	}

	// Parse status code, default to 302 if not provided or invalid
	statusCode := 302
	if code := ctx.Request.FormValue("status_code"); code != "" {
		if parsed, err := strconv.Atoi(code); err == nil && parsed >= 300 && parsed < 400 {
				statusCode = parsed
		}
	}

	ctx.Redirect(statusCode, url)
}

func (h *HttpRedirect) HandleRedirect(ctx *gin.Context) {
	// Get the number of redirects from path parameter
	n := ctx.Param("n")
	redirectCount, err := strconv.Atoi(n)
	if err != nil || redirectCount <= 0 {
		ctx.JSON(400, gin.H{"error": "Invalid redirect count"})
		return
	}

	// If this is the final redirect, redirect to /get endpoint
	if redirectCount == 1 {
		ctx.Redirect(302, "/get")
		return
	}

	// Construct next redirect URL with relative path
	nextURL := fmt.Sprintf("/redirect/%d", redirectCount-1)
	ctx.Redirect(302, nextURL)
}