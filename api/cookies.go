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
