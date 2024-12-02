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
//	@description	An alternative to httpbin.org in Gin. <br/> <br/> <b>Run locally: </b> <code>$ docker run -p 8080:8080 viethuy/endpointlab</code>
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	John Cao
//	@contact.url	https://github.com/kics223w1/EndpointLab

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080
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
	s.router.GET("/bearer", httpAuth.HandleBearer)
	s.router.GET("/basic-auth/:user/:passwd", httpAuth.HandleBasicAuth)
	s.router.GET("/digest-auth/:qop/:user/:passwd", httpAuth.HandleDigestAuth)
	s.router.GET("/hidden-basic-auth/:user/:passwd", httpAuth.HandleHiddenBasicAuth)
	s.router.GET("/digest-auth/:qop/:user/:passwd/:algorithm", httpAuth.HandleDigestAuthAlgorithm)
	s.router.GET("/digest-auth/:qop/:user/:passwd/:algorithm/:stale_after", httpAuth.HandleDigestAuthStaleAfter)

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
	s.router.GET("/absolute-redirect/:n", httpRedirects.HandleAbsoluteRedirect)
	s.router.Any("/redirect-to", httpRedirects.HandleRedirectTo)
	s.router.GET("/redirect/:n", httpRedirects.HandleRedirect)
	s.router.GET("/relative-redirect/:n", httpRedirects.HandleRelativeRedirect)
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}



