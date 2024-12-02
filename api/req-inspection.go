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

//	@Summary		Returns the request headers.
//	@Description	Returns all headers of the request
//	@Tags			Request inspection
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/headers [get]
func (r *ReqInspection) HandleHeaders(ctx *gin.Context) {
    headers := utils.ConvertHeaders(ctx.Request.Header)
    ctx.JSON(http.StatusOK, gin.H{"headers": headers})
}

//	@Summary		Returns the client's IP address.
//	@Description	Returns the IP address of the client making the request
//	@Tags			Request inspection
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/ip [get]
func (r *ReqInspection) HandleIP(ctx *gin.Context) {
	ip := ctx.ClientIP()
	ctx.JSON(http.StatusOK, gin.H{"origin": ip})
}

//	@Summary		Returns the User-Agent string.
//	@Description	Returns the User-Agent string of the request
//	@Tags			Request inspection
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Router			/user-agent [get]
func (r *ReqInspection) HandleUserAgent(ctx *gin.Context) {
	userAgent := ctx.Request.UserAgent()
	ctx.JSON(http.StatusOK, gin.H{"user_agent": userAgent})
}
