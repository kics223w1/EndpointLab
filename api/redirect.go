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

//	@Summary		Handle absolute redirects.
//	@Description	Redirects the request to a new URL a specified number of times.
//	@Tags			Redirects
//	@Accept			json
//	@Produce		json
//	@Param			n	path		int	true	"Number of redirects"
//	@Success		302	{string}	string	"Redirects to the next URL"
//	@Failure		400	{object}	object	"Invalid redirect count"
//	@Router			/absolute-redirect/{n} [get]
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

//	@Summary		Handle redirect to a specific URL.
//	@Description	Redirects the request to a specified URL with an optional status code.
//	@Tags			Redirects
//	@Accept			json
//	@Produce		json
//	@Param			url			formData	string	true	"URL to redirect to"
//	@Param			status_code	formData	int		false	"HTTP status code for the redirect"
//	@Success		302	{string}	string	"Redirects to the specified URL"
//	@Failure		400	{object}	object	"URL is required"
//	@Router			/redirect-to [post]
//	@Router			/redirect-to [delete]
//	@Router			/redirect-to [put]
//	@Router			/redirect-to [patch]
//	@Router			/redirect-to [get]
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

//	@Summary		Handle relative redirects.
//	@Description	Redirects the request to a relative URL a specified number of times.
//	@Tags			Redirects
//	@Accept			json
//	@Produce		json
//	@Param			n	path		int	true	"Number of redirects"
//	@Success		302	{string}	string	"Redirects to the next URL"
//	@Failure		400	{object}	object	"Invalid redirect count"
//	@Router			/redirect/{n} [get]
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

//	@Summary		Handle relative redirects with a different path.
//	@Description	Redirects the request to a relative URL with a different path a specified number of times.
//	@Tags			Redirects
//	@Accept			json
//	@Produce		json
//	@Param			n	path		int	true	"Number of redirects"
//	@Success		302	{string}	string	"Redirects to the next URL"
//	@Failure		400	{object}	object	"Invalid redirect count"
//	@Router			/relative-redirect/{n} [get]
func (h *HttpRedirect) HandleRelativeRedirect(ctx *gin.Context) {
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
	nextURL := fmt.Sprintf("/relative-redirect/%d", redirectCount-1)
	ctx.Redirect(302, nextURL)
}