package api

import (
	"github.com/gin-gonic/gin"
)

type HttpAuth struct {
}

func NewHttpAuth() *HttpAuth {
	return &HttpAuth{}
}

func (h *HttpAuth) HandleBearer(ctx *gin.Context) {
	// Get the Authorization header
	authHeader := ctx.GetHeader("Authorization")
	
	// If no Authorization header is present, return 401 Unauthorized
	if authHeader == "" {
		ctx.Header("WWW-Authenticate", "Bearer")
		ctx.AbortWithStatus(401)
		return
	}
	
	// Check if it starts with "Bearer "
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		ctx.Header("WWW-Authenticate", "Bearer error=\"invalid_token\"")
		ctx.AbortWithStatus(401)
		return
	}
	
	// Extract the token (everything after "Bearer ")
	token := authHeader[7:]
	
	// Store the token in the context for later use
	ctx.Set("bearer_token", token)
	ctx.JSON(200, gin.H{"token": token, "authenticated": true})
}

func (h *HttpAuth) HandleBasicAuth(ctx *gin.Context) {
	// Get user and password from URL parameters
	user := ctx.Param("user")
	passwd := ctx.Param("passwd")

	// Get the Authorization header
	authHeader := ctx.GetHeader("Authorization")

	// If no Authorization header is present, return 401 Unauthorized
	if authHeader == "" {
		ctx.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
		ctx.AbortWithStatus(401)
		return
	}

	// Get the credentials from the request's Basic Auth
	providedUser, providedPass, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
		ctx.AbortWithStatus(401)
		return
	}

	// Compare the provided credentials with expected values
	if providedUser != user || providedPass != passwd {
		ctx.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
		ctx.AbortWithStatus(401)
		return
	}

	// If authentication is successful, return success response
	ctx.JSON(200, gin.H{
		"authenticated": true,
		"user": user,
	})
}

