package api

import "github.com/gin-gonic/gin"

type HttpCookies struct {}

func NewHttpCookies() *HttpCookies {
	return &HttpCookies{}
}

//	@Summary		Get cookies
//	@Description	Returns the cookies sent by the client
//	@Tags			Cookies
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/cookies [get]
func (c *HttpCookies) HandleCookies(ctx *gin.Context) {
	cookies := ctx.Request.Cookies()
	ctx.JSON(200, gin.H{
		"cookies": cookies,
	})
}

//	@Summary		Sets cookie(s) as provided by the query string and redirects to cookie list.
//	@Tags			Cookies
//	@Accept			json
//	@Produce		json
//	@Param			freeform	query	string	false	"Freeform cookie value"
//	@Success		302		{string}	string	"Redirects to cookie list"
//	@Router			/cookies/set [get]
func (c *HttpCookies) HandleSetCookie(ctx *gin.Context) {
	// Get the freeform value from query parameters
	freeform := ctx.Query("freeform")
	if freeform != "" {
		ctx.SetCookie("freeform", freeform, 3600, "/", "", false, false)
	}
	
	// Redirect to cookie list
	ctx.Redirect(302, "/cookies")
}

//	@Summary		Set a cookie with specified name and value
//	@Tags			Cookies
//	@Accept			json
//	@Produce		json
//	@Param			name	path	string	true	"Cookie name"
//	@Param			value	path	string	true	"Cookie value"
//	@Success		302		{string}	string	"Redirects to cookie list"
//	@Router			/cookies/set/{name}/{value} [get]
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

//	@Summary		Deletes cookie(s) as provided by the query string and redirects to cookie list.
//	@Tags			Cookies
//	@Accept			json
//	@Produce		json
//	@Param			freeform	query	string	false	"Freeform cookie value"
//	@Success		302		{string}	string	"Redirects to cookie list"
//	@Router			/cookies/delete [get]
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


