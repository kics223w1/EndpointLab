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