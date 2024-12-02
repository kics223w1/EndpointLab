package api

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpAuth struct {
}

func NewHttpAuth() *HttpAuth {
	return &HttpAuth{}
}

//	@Summary		Handle Bearer Token Authentication
//	@Description	Authenticates requests using a Bearer token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Success		200				{object}	map[string]interface{}
//	@Failure		401				{string}	string	"Unauthorized"
//	@Router			/bearer [get]
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

//	@Summary		Handle Basic Authentication
//	@Description	Authenticates requests using Basic Auth
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user	path		string	true	"Username"
//	@Param			passwd	path		string	true	"Password"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		401		{string}	string	"Unauthorized"
//	@Router			/basic-auth/{user}/{passwd} [get]
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

//	@Summary		Handle Digest Authentication
//	@Description	Authenticates requests using Digest Auth
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			qop		path		string	true	"Quality of Protection"	default(auth)
//	@Param			user	path		string	true	"Username"				default(user)
//	@Param			passwd	path		string	true	"Password"				default(passwd)
//	@Success		200		{object}	map[string]interface{}
//	@Failure		401		{string}	string	"Unauthorized"
//	@Router			/auth/digest/{qop}/{user}/{passwd} [get]
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

//	@Summary		Handle Digest Authentication with Algorithm
//	@Description	Authenticates requests using Digest Auth with specified algorithm
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			qop			path		string	true	"Quality of Protection"				default(auth)
//	@Param			user		path		string	true	"Username"							default(user)
//	@Param			passwd		path		string	true	"Password"							default(passwd)
//	@Param			algorithm	path		string	true	"Algorithm (MD5, SHA-256, SHA-512)"	default(MD5)
//	@Success		200			{object}	map[string]interface{}
//	@Failure		401			{string}	string	"Unauthorized"
//	@Router			/auth/digest-algorithm/{qop}/{user}/{passwd}/{algorithm} [get]
func (h *HttpAuth) HandleDigestAuthAlgorithm(ctx *gin.Context) {
	// Get parameters from URL
	qop := ctx.Param("qop")
	user := ctx.Param("user")
	passwd := ctx.Param("passwd")
	algorithm := ctx.Param("algorithm")

	// Set defaults if not provided
	if user == "" {
		user = "user"
	}
	if passwd == "" {
		passwd = "passwd"
	}

	// Validate qop parameter
	if qop != "auth" && qop != "auth-int" {
		ctx.AbortWithStatus(400)
		return
	}

	// Validate and normalize algorithm
	algorithm = strings.ToUpper(algorithm)
	switch algorithm {
	case "MD5", "SHA-256", "SHA-512":
		// Valid algorithms
	case "":
		algorithm = "MD5" // Set default
	default:
		ctx.AbortWithStatus(400)
		return
	}

	// Generate a static nonce to match Python implementation
	nonce := "dcd98b7102dd2f0e8b11d0f600bfb0c093"
	realm := "Authentication Required"
	opaque := "5ccc069c403ebaf9f0171e9517f40e41"

	// Get the Authorization header
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		// If no Authorization header, send WWW-Authenticate header
		ctx.Header("WWW-Authenticate", fmt.Sprintf(
			`Digest realm="%s", qop="%s", nonce="%s", opaque="%s", algorithm="%s", stale=FALSE`,
			realm, qop, nonce, opaque, algorithm))
		ctx.AbortWithStatus(401)
		return
	}

	// Parse the Digest Authorization header
	params := parseDigestHeader(authHeader)
	if params == nil {
		ctx.AbortWithStatus(401)
		return
	}

	// Verify the provided username matches
	if params["username"] != user {
		ctx.AbortWithStatus(401)
		return
	}

	// Calculate HA1 based on algorithm
	var ha1 string
	baseString := fmt.Sprintf("%s:%s:%s", user, realm, passwd)
	switch algorithm {
	case "SHA-256":
		ha1 = sha256hex(baseString)
	case "SHA-512":
		ha1 = sha512hex(baseString)
	default: // MD5
		ha1 = md5hex(baseString)
	}

	// Calculate HA2
	ha2String := fmt.Sprintf("%s:%s", ctx.Request.Method, params["uri"])
	var ha2 string
	switch algorithm {
	case "SHA-256":
		ha2 = sha256hex(ha2String)
	case "SHA-512":
		ha2 = sha512hex(ha2String)
	default: // MD5
		ha2 = md5hex(ha2String)
	}

	// Calculate expected response
	var expectedResponse string
	if qop == "auth" {
		responseString := fmt.Sprintf("%s:%s:%s:%s:%s:%s",
			ha1, nonce, params["nc"], params["cnonce"], qop, ha2)
		switch algorithm {
		case "SHA-256":
			expectedResponse = sha256hex(responseString)
		case "SHA-512":
			expectedResponse = sha512hex(responseString)
		default: // MD5
			expectedResponse = md5hex(responseString)
		}
	} else {
		responseString := fmt.Sprintf("%s:%s:%s", ha1, nonce, ha2)
		switch algorithm {
		case "SHA-256":
			expectedResponse = sha256hex(responseString)
		case "SHA-512":
			expectedResponse = sha512hex(responseString)
		default: // MD5
			expectedResponse = md5hex(responseString)
		}
	}

	// Compare the response
	if expectedResponse != params["response"] {
		ctx.AbortWithStatus(401)
		return
	}

	// If authentication is successful, return success response
	ctx.JSON(200, gin.H{
		"authenticated": true,
		"user": user,
	})
}

//	@Summary		Handle Digest Authentication with Stale After
//	@Description	Authenticates requests using Digest Auth with stale after parameter
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			qop			path		string	true	"Quality of Protection"				default(auth)
//	@Param			user		path		string	true	"Username"							default(user)
//	@Param			passwd		path		string	true	"Password"							default(passwd)
//	@Param			algorithm	path		string	true	"Algorithm (MD5, SHA-256, SHA-512)"	default(MD5)
//	@Param			stale_after	path		string	true	"Stale After"						default(never)
//	@Success		200			{object}	map[string]interface{}
//	@Failure		401			{string}	string	"Unauthorized"
//	@Router			/auth/digest-stale/{qop}/{user}/{passwd}/{algorithm}/{stale_after} [get]
func (h *HttpAuth) HandleDigestAuthStaleAfter(ctx *gin.Context) {
	// Get parameters from URL
	qop := ctx.Param("qop")
	user := ctx.Param("user")
	passwd := ctx.Param("passwd")
	algorithm := ctx.Param("algorithm")
	staleAfter := ctx.Param("stale_after")

	// Get cookie requirement from query params
	requireCookie := strings.ToLower(ctx.Query("require-cookie"))
	cookieRequired := requireCookie == "1" || requireCookie == "t" || requireCookie == "true"

	// Set defaults and validate parameters
	if user == "" {
		user = "user"
	}
	if passwd == "" {
		passwd = "passwd"
	}
	if algorithm == "" {
		algorithm = "MD5"
	}
	if staleAfter == "" {
		staleAfter = "never"
	}

	// Validate qop parameter
	if qop != "auth" && qop != "auth-int" {
		qop = ""
	}

	// Validate algorithm
	algorithm = strings.ToUpper(algorithm)
	switch algorithm {
	case "MD5", "SHA-256", "SHA-512":
		// Valid algorithms
	default:
		algorithm = "MD5"
	}

	// Get the Authorization header
	authHeader := ctx.GetHeader("Authorization")

	// Check cookie requirement
	if cookieRequired && ctx.GetHeader("Cookie") == "" {
		sendChallengeResponse(ctx, qop, algorithm, false, staleAfter)
		return
	}

	// If no valid authorization, send challenge
	if authHeader == "" {
		sendChallengeResponse(ctx, qop, algorithm, false, staleAfter)
		return
	}

	// Parse the Digest Authorization header
	params := parseDigestHeader(authHeader)
	if params == nil {
		sendChallengeResponse(ctx, qop, algorithm, false, staleAfter)
		return
	}

	// Check for required cookie
	fakeValue, _ := ctx.Cookie("fake")
	if cookieRequired && fakeValue != "fake_value" {
		ctx.SetCookie("fake", "fake_value", 3600, "/", "", false, true)
		ctx.JSON(403, gin.H{"errors": []string{"missing cookie set on challenge"}})
		return
	}

	currentNonce := params["nonce"]
	lastNonce, _ := ctx.Cookie("last_nonce")
	storedStaleAfter, _ := ctx.Cookie("stale_after")

	// Check if nonce is stale
	if (lastNonce != "" && currentNonce == lastNonce) || storedStaleAfter == "0" {
		sendChallengeResponse(ctx, qop, algorithm, true, staleAfter)
		ctx.SetCookie("last_nonce", currentNonce, 3600, "/", "", false, true)
		return
	}

	// Verify the credentials
	if !verifyDigestAuth(ctx, params, user, passwd, algorithm, qop) {
		sendChallengeResponse(ctx, qop, algorithm, false, staleAfter)
		ctx.SetCookie("last_nonce", currentNonce, 3600, "/", "", false, true)
		return
	}

	// Successful authentication
	ctx.SetCookie("fake", "fake_value", 3600, "/", "", false, true)
	if storedStaleAfter != "" {
		nextStaleAfter := calculateNextStaleAfter(storedStaleAfter)
		ctx.SetCookie("stale_after", nextStaleAfter, 3600, "/", "", false, true)
	}

	ctx.JSON(200, gin.H{
		"authenticated": true,
		"user":         user,
	})
}

//	@Summary		Handle Hidden Basic Authentication
//	@Description	Authenticates requests using Basic Auth but hides the requirement
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			user	path		string	true	"Username"
//	@Param			passwd	path		string	true	"Password"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		404		{string}	string	"Not Found"
//	@Router			/hidden-basic-auth/{user}/{passwd} [get]
func (h *HttpAuth) HandleHiddenBasicAuth(ctx *gin.Context) {
	// Get user and password from URL parameters
	user := ctx.Param("user")
	passwd := ctx.Param("passwd")

	// Get the credentials from the request's Basic Auth
	providedUser, providedPass, ok := ctx.Request.BasicAuth()
	
	// If no auth provided or credentials don't match, return 404
	// Note: We don't send WWW-Authenticate header to hide that auth is required
	if !ok || providedUser != user || providedPass != passwd {
		ctx.AbortWithStatus(404)
		return
	}

	// If authentication is successful, return success response
	ctx.JSON(200, gin.H{
		"authenticated": true,
		"user": user,
	})
}

// Helper function to send challenge response
func sendChallengeResponse(ctx *gin.Context, qop, algorithm string, stale bool, staleAfter string) {
	nonce := "dcd98b7102dd2f0e8b11d0f600bfb0c093"
	realm := "Authentication Required"
	opaque := "5ccc069c403ebaf9f0171e9517f40e41"

	ctx.Header("WWW-Authenticate", fmt.Sprintf(
		`Digest realm="%s", qop="%s", nonce="%s", opaque="%s", algorithm="%s", stale=%t`,
		realm, qop, nonce, opaque, algorithm, stale))
	
	ctx.SetCookie("stale_after", staleAfter, 3600, "/", "", false, true)
	ctx.SetCookie("fake", "fake_value", 3600, "/", "", false, true)
	
	ctx.AbortWithStatus(401)
}

// Helper function to calculate next stale_after value
func calculateNextStaleAfter(current string) string {
	if current == "never" {
		return "never"
	}
	
	val, err := strconv.Atoi(current)
	if err != nil {
		return current
	}
	
	if val <= 0 {
		return "0"
	}
	
	return strconv.Itoa(val - 1)
}

// Helper function to verify digest auth credentials
func verifyDigestAuth(ctx *gin.Context, params map[string]string, user, passwd, algorithm, qop string) bool {
	realm := "Authentication Required"
	nonce := "dcd98b7102dd2f0e8b11d0f600bfb0c093"

	// Verify username
	if params["username"] != user {
		return false
	}

	// Calculate HA1
	var ha1 string
	baseString := fmt.Sprintf("%s:%s:%s", user, realm, passwd)
	switch algorithm {
	case "SHA-256":
		ha1 = sha256hex(baseString)
	case "SHA-512":
		ha1 = sha512hex(baseString)
	default:
		ha1 = md5hex(baseString)
	}

	// Calculate HA2
	ha2String := fmt.Sprintf("%s:%s", ctx.Request.Method, params["uri"])
	var ha2 string
	switch algorithm {
	case "SHA-256":
		ha2 = sha256hex(ha2String)
	case "SHA-512":
		ha2 = sha512hex(ha2String)
	default:
		ha2 = md5hex(ha2String)
	}

	// Calculate expected response
	var expectedResponse string
	if qop == "auth" {
		responseString := fmt.Sprintf("%s:%s:%s:%s:%s:%s",
			ha1, nonce, params["nc"], params["cnonce"], qop, ha2)
		switch algorithm {
		case "SHA-256":
			expectedResponse = sha256hex(responseString)
		case "SHA-512":
			expectedResponse = sha512hex(responseString)
		default:
			expectedResponse = md5hex(responseString)
		}
	} else {
		responseString := fmt.Sprintf("%s:%s:%s", ha1, nonce, ha2)
		switch algorithm {
		case "SHA-256":
			expectedResponse = sha256hex(responseString)
		case "SHA-512":
			expectedResponse = sha512hex(responseString)
		default:
			expectedResponse = md5hex(responseString)
		}
	}

	return expectedResponse == params["response"]
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

func sha256hex(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func sha512hex(data string) string {
	hash := sha512.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

