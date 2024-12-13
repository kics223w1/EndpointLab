package api

import (
	"github.com/gin-gonic/gin"
)

func setupRouterStatus() *gin.Engine {
	r := gin.Default()
	httpStatus := NewHttpStatus()

	r.Any("/status/:code", httpStatus.HandleStatus)

	return r
}
