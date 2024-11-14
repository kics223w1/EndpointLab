package api

import (
	"github.com/gin-gonic/gin"
)

type ResInspection struct {}

func NewResInspection() *ResInspection {
	return &ResInspection{}
}


func (r *ResInspection) HandleCache(ctx *gin.Context) {
	
}

func (r *ResInspection) HandleCacheValue(ctx *gin.Context) {

}

func (r *ResInspection) HandleETag(ctx *gin.Context) {

}

func (r *ResInspection) HandleResponseHeaders(ctx *gin.Context) {

}
