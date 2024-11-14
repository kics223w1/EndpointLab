package api

import "github.com/gin-gonic/gin"

type Cookies struct {}

func NewCookies() *Cookies {
	return &Cookies{}
}

func (c *Cookies) HandleCookies(ctx *gin.Context) {
	cookies := ctx.Request.Cookies()
	ctx.JSON(200, gin.H{
		"cookies": cookies,
	})
}

func (c *Cookies) HandleSetCookie(ctx *gin.Context) {
	// Get the freeform value from query parameters
	freeform := ctx.Query("freeform")
	if freeform != "" {
		ctx.SetCookie("freeform", freeform, 3600, "/", "", false, false)
	}
	
	// Redirect to cookie list
	ctx.Redirect(302, "/cookies")
}

func (c *Cookies) HandleSetCookieWithParams(ctx *gin.Context) {
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

func (c *Cookies) HandleDeleteCookie(ctx *gin.Context) {

}


