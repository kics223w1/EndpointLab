package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

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

func (h *HttpAuth) HandleDigestAuth(ctx *gin.Context) {
	// Get parameters from URL
	qop := ctx.Param("qop")
	user := ctx.Param("user")
	passwd := ctx.Param("passwd")

	// Validate qop parameter
	if qop != "auth" && qop != "auth-int" {
		ctx.AbortWithStatus(400)
		return
	}

	// Get the Authorization header
	authHeader := ctx.GetHeader("Authorization")

	// Generate a static nonce to match Python implementation
	nonce := "dcd98b7102dd2f0e8b11d0f600bfb0c093"
	realm := "Authentication Required"
	opaque := "5ccc069c403ebaf9f0171e9517f40e41"

	// If no Authorization header is present, send WWW-Authenticate header
	if authHeader == "" {
		ctx.Header("WWW-Authenticate", fmt.Sprintf(
			`Digest realm="%s", qop="%s", nonce="%s", opaque="%s", algorithm="MD5", stale=FALSE`,
			realm, qop, nonce, opaque))
		ctx.AbortWithStatus(401)
		return
	}

	// Parse the Digest Authorization header
	params := parseDigestHeader(authHeader)
	if params == nil {
		ctx.AbortWithStatus(401)
		return
	}

	// Verify the response
	ha1 := md5hex(fmt.Sprintf("%s:%s:%s", user, realm, passwd))
	ha2 := md5hex(fmt.Sprintf("%s:%s", ctx.Request.Method, params["uri"]))
	
	expectedResponse := ""
	if qop == "auth" {
		expectedResponse = md5hex(fmt.Sprintf("%s:%s:%s:%s:%s:%s",
			ha1, nonce, params["nc"], params["cnonce"], qop, ha2))
	} else {
		expectedResponse = md5hex(fmt.Sprintf("%s:%s:%s", ha1, nonce, ha2))
	}

	if expectedResponse != params["response"] {
		ctx.AbortWithStatus(401)
		return
	}

	// If authentication is successful, return success response
	ctx.JSON(200, gin.H{
		"authenticated": true,
		"user":         user,
	})
}

// Helper function to parse digest authorization header
func parseDigestHeader(header string) map[string]string {
	if !strings.HasPrefix(header, "Digest ") {
		return nil
	}
	
	parts := strings.Split(header[7:], ",")
	params := make(map[string]string)
	
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if eqIndex := strings.Index(part, "="); eqIndex > 0 {
			key := strings.TrimSpace(part[:eqIndex])
			value := strings.Trim(strings.TrimSpace(part[eqIndex+1:]), "\"")
			params[key] = value
		}
	}
	
	return params
}

// Helper function to calculate MD5 hex
func md5hex(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

