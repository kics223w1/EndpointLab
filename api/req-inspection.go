package api

import (
	"endpointlab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqInspection struct {

}

func NewReqInspection() *ReqInspection {
	return &ReqInspection{}
}

func (r *ReqInspection) HandleHeaders(ctx *gin.Context) {
    headers := utils.ConvertHeaders(ctx.Request.Header)
    ctx.JSON(http.StatusOK, gin.H{"headers": headers})
}

func (r *ReqInspection) HandleIP(ctx *gin.Context) {
	ip := ctx.ClientIP()
	ctx.JSON(http.StatusOK, gin.H{"origin": ip})
}

func (r *ReqInspection) HandleUserAgent(ctx *gin.Context) {
	userAgent := ctx.Request.UserAgent()
	ctx.JSON(http.StatusOK, gin.H{"user_agent": userAgent})
}
