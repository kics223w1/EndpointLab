package api

import (
	_ "endpointlab/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	router := gin.Default()
	
	server := &Server{
		router: router,
	}

	server.registerRoutes()	
	
	return server, nil
}

//	@title			EndpointLab
//	@version		1.0
//	@description	 An alternative to httpbin.org in Gin. 
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	https://github.com/kics223w1/EndpointLab

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
func (s *Server) registerRoutes() {
	// Basic HTTP method endpoints
	httpMethod := NewHttpMethod()

	s.router.GET("/swagger/*any", func(c *gin.Context) {
		if c.Request.RequestURI == "/swagger/" {
			c.Redirect(302, "/swagger/index.html")
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})

	s.router.GET("/get", httpMethod.HandleGet)
	s.router.POST("/post", httpMethod.HandlePost)
	s.router.PUT("/put", httpMethod.HandlePut)
	s.router.DELETE("/delete", httpMethod.HandleDelete)
	s.router.PATCH("/patch", httpMethod.HandlePatch)

	
	// Auth endpoint
	httpAuth := NewHttpAuth()
	//	@Summary		Bearer auth
	//	@Description	Protected endpoint that requires bearer token
	//	@Tags			auth
	//	@Accept			json
	//	@Produce		json
	//	@Security		ApiKeyAuth
	//	@Success		200	{object}	object
	//	@Failure		401	{object}	object
	//	@Router			/bearer [get]
	s.router.GET("/bearer", httpAuth.HandleBearer)

	//	@Summary		Basic auth
	//	@Description	Protected endpoint that requires basic auth
	//	@Tags			auth
	//	@Accept			json
	//	@Produce		json
	//	@Param			user	path	string	true	"Username"
	//	@Param			passwd	path	string	true	"Password"
	//	@Security		BasicAuth
	//	@Success		200	{object}	object
	//	@Failure		401	{object}	object
	//	@Router			/basic-auth/{user}/{passwd} [get]
	s.router.GET("/basic-auth/:user/:passwd", httpAuth.HandleBasicAuth)

	//	@Summary		Digest auth
	//	@Description	Protected endpoint that requires digest auth
	//	@Tags			auth
	//	@Accept			json
	//	@Produce		json
	//	@Param			qop		path		string	true	"Quality of Protection"
	//	@Param			user	path		string	true	"Username"
	//	@Param			passwd	path		string	true	"Password"
	//	@Success		200		{object}	object
	//	@Failure		401		{object}	object
	//	@Router			/digest-auth/{qop}/{user}/{passwd} [get]
	s.router.GET("/digest-auth/:qop/:user/:passwd", httpAuth.HandleDigestAuth)

	//	@Summary		Digest auth with algorithm
	//	@Description	Protected endpoint that requires digest auth with specific algorithm
	//	@Tags			auth
	//	@Accept			json
	//	@Produce		json
	//	@Param			qop			path		string	true	"Quality of Protection"
	//	@Param			user		path		string	true	"Username"
	//	@Param			passwd		path		string	true	"Password"
	//	@Param			algorithm	path		string	true	"Algorithm"
	//	@Success		200			{object}	object
	//	@Failure		401			{object}	object
	//	@Router			/digest-auth/{qop}/{user}/{passwd}/{algorithm} [get]
	s.router.GET("/digest-auth/:qop/:user/:passwd/:algorithm", httpAuth.HandleDigestAuthAlgorithm)

	// Status endpoint
	httpStatus := NewHttpStatus()
	s.router.Any("/status/:code", httpStatus.HandleStatus)

	// Request inspection
	httpReqInspection := NewReqInspection()
	s.router.GET("/headers", httpReqInspection.HandleHeaders)
	s.router.GET("/ip", httpReqInspection.HandleIP)
	s.router.GET("/user-agent", httpReqInspection.HandleUserAgent)

	// Response inspection
	httpResInspection := NewResInspection()
	s.router.GET("/cache", httpResInspection.HandleCache)
	s.router.GET("/cache/:value", httpResInspection.HandleCacheValue)
	s.router.GET("/etag/:etag", httpResInspection.HandleETag)
	s.router.GET("/response-headers", httpResInspection.HandleResponseHeaders)
	s.router.POST("/response-headers", httpResInspection.HandleResponseHeaders)

	httpFormat := NewHttpFormat()
	s.router.GET("/brotli", httpFormat.HandleBrotli)
	s.router.GET("/deflate", httpFormat.HandleDeflate)
	s.router.GET("/deny", httpFormat.HandleDeny)
	s.router.GET("/encoding/utf8", httpFormat.handleUTF8)
	s.router.GET("/gzip", httpFormat.HandleGzip)
	s.router.GET("/html", httpFormat.HandleHtml)
	s.router.GET("/json", httpFormat.HandleJson)
	s.router.GET("/robots.txt", httpFormat.HandleRobotTxt)
	s.router.GET("/xml", httpFormat.HandleXML)

	// Images
	httpImage := NewHttpImage()
	s.router.GET("/image", httpImage.HandleImage)
	s.router.GET("/image/png", httpImage.HandleImagePNG)
	s.router.GET("/image/jpeg", httpImage.HandleImageJPEG)
	s.router.GET("/image/webp", httpImage.HandleImageWebp)
	s.router.GET("/image/svg", httpImage.HandleImageSVG)

	// Anything
	httpAnything := NewHttpAnything()
	s.router.Any("/anything", httpAnything.HandleAnything)

	// Dynamic
	httpDynamic := NewHttpDynamic()
	s.router.GET("/base64/:value", httpDynamic.HandleBase64)
	s.router.GET("/bytes/:n", httpDynamic.HandleBytes)
	s.router.Any("/delay/:delay", httpDynamic.HandleDeplay)
	s.router.Any("/drip", httpDynamic.HandleDrip)
	s.router.GET("/links/:n/:offset", httpDynamic.HandleLinks)
	s.router.GET("/range/:numbytes", httpDynamic.HandleRange)
	s.router.GET("/stream-bytes/:n", httpDynamic.HandleStreamBytes)
	s.router.GET("/stream/:n", httpDynamic.HandleStream)
	s.router.GET("/uuid", httpDynamic.HandleUuid)

	// Cookies
	httpCookies := NewHttpCookies()
	s.router.GET("/cookies", httpCookies.HandleCookies)
	s.router.GET("/cookies/set", httpCookies.HandleSetCookie)
	s.router.GET("/cookies/set/:name/:value", httpCookies.HandleSetCookieWithParams)
	s.router.GET("/cookies/delete", httpCookies.HandleDeleteCookie)

	// Redirects
	httpRedirects := NewHttpRedirect()
	//	@Summary		Absolute redirect
	//	@Description	Performs an absolute URL redirect the specified number of times
	//	@Tags			redirects
	//	@Accept			json
	//	@Produce		json
	//	@Param			n	path		int		true	"Number of redirects"
	//	@Success		302	{object}	object	"Redirect response"
	//	@Failure		400	{object}	object	"Invalid number of redirects"
	//	@Router			/absolute-redirect/{n} [get]
	s.router.GET("/absolute-redirect/:n", httpRedirects.HandleAbsoluteRedirect)

	//	@Summary		Redirect to URL
	//	@Description	Redirects to the URL specified in the 'url' query parameter
	//	@Tags			redirects
	//	@Accept			json
	//	@Produce		json
	//	@Param			url			query		string	true	"URL to redirect to"
	//	@Param			status_code	query		int		false	"HTTP status code for redirect (default 302)"
	//	@Success		302			{object}	object	"Redirect response"
	//	@Success		301			{object}	object	"Permanent redirect response"
	//	@Success		307			{object}	object	"Temporary redirect response"
	//	@Success		308			{object}	object	"Permanent redirect response"
	//	@Failure		400			{object}	object	"Invalid URL or status code"
	//	@Router			/redirect-to [get]
	s.router.Any("/redirect-to", httpRedirects.HandleRedirectTo)

	//	@Summary		Redirect n times
	//	@Description	Redirects n times before returning a response
	//	@Tags			redirects
	//	@Accept			json
	//	@Produce		json
	//	@Param			n	path		int		true	"Number of redirects"
	//	@Success		302	{object}	object	"Redirect response"
	//	@Failure		400	{object}	object	"Invalid number of redirects"
	//	@Router			/redirect/{n} [get]
	s.router.GET("/redirect/:n", httpRedirects.HandleRedirect)

	//	@Summary		Relative redirect
	//	@Description	Performs a relative redirect the specified number of times
	//	@Tags			redirects
	//	@Accept			json
	//	@Produce		json
	//	@Param			n	path		int		true	"Number of redirects"
	//	@Success		302	{object}	object	"Redirect response"
	//	@Failure		400	{object}	object	"Invalid number of redirects"
	//	@Router			/relative-redirect/{n} [get]
	s.router.GET("/relative-redirect/:n", httpRedirects.HandleRelativeRedirect)
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}



