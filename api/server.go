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
	//	@Summary		Status code
	//	@Description	Returns response with specified status code
	//	@Tags			status
	//	@Accept			json
	//	@Produce		json
	//	@Param			code	path		int	true	"HTTP Status Code"
	//	@Success		200		{object}	object
	//	@Router			/status/{code} [get]
	s.router.Any("/status/:code", httpStatus.HandleStatus)

	// Request inspection
	httpReqInspection := NewReqInspection()
	//	@Summary		Request headers
	//	@Description	Returns the request's headers
	//	@Tags			request-inspection
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{object}	object
	//	@Router			/headers [get]
	s.router.GET("/headers", httpReqInspection.HandleHeaders)

	//	@Summary		Client IP
	//	@Description	Returns the client's IP address
	//	@Tags			request-inspection
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{object}	object
	//	@Router			/ip [get]
	s.router.GET("/ip", httpReqInspection.HandleIP)

	//	@Summary		User agent
	//	@Description	Returns the user agent string
	//	@Tags			request-inspection
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{object}	object
	//	@Router			/user-agent [get]
	s.router.GET("/user-agent", httpReqInspection.HandleUserAgent)

	// Response inspection
	httpResInspection := NewResInspection()
	//	@Summary		Cache control
	//	@Description	Returns cache control headers
	//	@Tags			response-inspection
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{object}	object
	//	@Router			/cache [get]
	s.router.GET("/cache", httpResInspection.HandleCache)

	// Format endpoint
	httpFormat := NewHttpFormat()
	//	@Summary		Brotli compressed response
	//	@Description	Returns brotli-encoded data
	//	@Tags			response-format
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{object}	object
	//	@Router			/brotli [get]
	s.router.GET("/brotli", httpFormat.HandleBrotli)

	//	@Summary		JSON response
	//	@Description	Returns JSON-encoded data
	//	@Tags			response-format
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{object}	object
	//	@Router			/json [get]
	s.router.GET("/json", httpFormat.HandleJson)

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
	//	@Summary		Base64 encoded data
	//	@Description	Returns base64 encoded data
	//	@Tags			dynamic-data
	//	@Accept			json
	//	@Produce		json
	//	@Param			value	path		string	true	"Value to encode"
	//	@Success		200		{object}	object
	//	@Router			/base64/{value} [get]
	s.router.GET("/base64/:value", httpDynamic.HandleBase64)

	//	@Summary		Random bytes
	//	@Description	Returns n random bytes generated with specified seed
	//	@Tags			dynamic-data
	//	@Accept			json
	//	@Produce		application/octet-stream
	//	@Param			n	path		int		true	"Number of bytes to generate (max 100KB)"
	//	@Success		200	{file}		binary	"Random binary data"
	//	@Failure		400	{object}	object	"Invalid number of bytes requested"
	//	@Router			/bytes/{n} [get]
	s.router.GET("/bytes/:n", httpDynamic.HandleBytes)

	//	@Summary		Delayed response
	//	@Description	Returns a delayed response (max: 10 seconds)
	//	@Tags			dynamic-data
	//	@Accept			json
	//	@Produce		json
	//	@Param			delay	path		number	true	"Delay in seconds (float value, max 10)"
	//	@Success		200		{object}	object
	//	@Failure		400		{object}	object	"Invalid delay value"
	//	@Router			/delay/{delay} [get]
	s.router.Any("/delay/:delay", httpDynamic.HandleDeplay)

	//	@Summary		Dripped response
	//	@Description	Drips data over a duration after an optional initial delay
	//	@Tags			dynamic-data
	//	@Accept			json
	//	@Produce		application/octet-stream
	//	@Param			duration	query		number	false	"Duration in seconds to drip data"
	//	@Param			numbytes	query		int		false	"Number of bytes to drip (default: 10)"
	//	@Param			code		query		int		false	"Status code (default: 200)"
	//	@Param			delay		query		number	false	"Initial delay in seconds"
	//	@Success		200			{file}		binary	"Dripped data"
	//	@Failure		400			{object}	object	"Invalid parameters"
	//	@Router			/drip [get]
	s.router.Any("/drip", httpDynamic.HandleDrip)

	//	@Summary		Generate links
	//	@Description	Returns a page containing n links to other pages offset by the specified amount
	//	@Tags			dynamic-data
	//	@Accept			json
	//	@Produce		json
	//	@Param			n		path		int	true	"Number of links to generate"
	//	@Param			offset	path		int	true	"Offset for link numbering"
	//	@Success		200		{object}	object
	//	@Failure		400		{object}	object	"Invalid parameters"
	//	@Router			/links/{n}/{offset} [get]
	s.router.GET("/links/:n/:offset", httpDynamic.HandleLinks)

	//	@Summary		Byte range
	//	@Description	Returns a range of bytes (Content-Range header)
	//	@Tags			dynamic-data
	//	@Accept			json
	//	@Produce		application/octet-stream
	//	@Param			numbytes	path		int		true	"Number of bytes to return"
	//	@Param			offset		query		int		false	"Range offset"
	//	@Success		200			{file}		binary	"Requested range of bytes"
	//	@Success		206			{file}		binary	"Partial content when using Range header"
	//	@Failure		400			{object}	object	"Invalid range requested"
	//	@Router			/range/{numbytes} [get]
	s.router.GET("/range/:numbytes", httpDynamic.HandleRange)

	//	@Summary		Streamed random bytes
	//	@Description	Streams n random bytes generated with given seed, streamed in chunks
	//	@Tags			dynamic-data
	//	@Accept			json
	//	@Produce		application/octet-stream
	//	@Param			n		path		int		true	"Number of bytes to stream"
	//	@Param			chunk	query		int		false	"Chunk size (default: 10240)"
	//	@Success		200		{file}		binary	"Streamed random bytes"
	//	@Failure		400		{object}	object	"Invalid parameters"
	//	@Router			/stream-bytes/{n} [get]
	s.router.GET("/stream-bytes/:n", httpDynamic.HandleStreamBytes)

	//	@Summary		Chunked transfer encoding
	//	@Description	Streams n chunks of JSON data with delay
	//	@Tags			dynamic-data
	//	@Accept			json
	//	@Produce		json
	//	@Param			n		path		int		true	"Number of chunks to stream"
	//	@Param			delay	query		number	false	"Delay between chunks in seconds (default: 0)"
	//	@Success		200		{array}		object	"Array of streamed data chunks"
	//	@Failure		400		{object}	object	"Invalid parameters"
	//	@Router			/stream/{n} [get]
	s.router.GET("/stream/:n", httpDynamic.HandleStream)

	//	@Summary		UUID generation
	//	@Description	Returns a randomly generated UUID v4
	//	@Tags			dynamic-data
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{object}	object	"Contains generated UUID"
	//	@Router			/uuid [get]
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



