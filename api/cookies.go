package api

import "github.com/gin-gonic/gin"

type HttpCookies struct {}

func NewHttpCookies() *HttpCookies {
	return &HttpCookies{}
}

func (c *HttpCookies) HandleCookies(ctx *gin.Context) {
	cookies := ctx.Request.Cookies()
	ctx.JSON(200, gin.H{
		"cookies": cookies,
	})
}

func (c *HttpCookies) HandleSetCookie(ctx *gin.Context) {
	// Get the freeform value from query parameters
	freeform := ctx.Query("freeform")
	if freeform != "" {
		ctx.SetCookie("freeform", freeform, 3600, "/", "", false, false)
	}
	
	// Redirect to cookie list
	ctx.Redirect(302, "/cookies")
}

func (c *HttpCookies) HandleSetCookieWithParams(ctx *gin.Context) {
	name := ctx.Param("name")
	value := ctx.Param("value")

	if name == "" || value == "" {
		ctx.JSON(400, gin.H{
			"error": "name and value are required",
		})
		return
	}

	ctx.SetCookie(name, value, 3600, "/", "", false, false)
	ctx.Redirect(302, "/cookies")
}

func (c *HttpCookies) HandleDeleteCookie(ctx *gin.Context) {
	// Get the freeform value from query parameters
	freeform := ctx.Query("freeform")
	if freeform != "" {
		// Set cookie with MaxAge = -1 to delete it
		ctx.SetCookie("freeform", "", -1, "/", "", false, false)
	}
	
	// Redirect to cookie list
	ctx.Redirect(302, "/cookies")
}


