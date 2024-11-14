package api

import (
	"github.com/gin-gonic/gin"
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

func (s *Server) registerRoutes() {
	// Register middleware
	// s.router.Use(ContentLengthMiddleware())

	// Basic HTTP method endpoints
	httpMethod := NewHttpMethod()
	s.router.GET("/get", httpMethod.HandleGet)
	s.router.POST("/post", httpMethod.HandlePost)
	s.router.PUT("/put", httpMethod.HandlePut)
	s.router.DELETE("/delete", httpMethod.HandleDelete)
	s.router.PATCH("/patch", httpMethod.HandlePatch)


	// Status endpoint
	httpStatus := NewHttpStatus()
	s.router.Any("/status/:code", httpStatus.HandleStatus)

	// Format endpoint
	httpFormat := NewHttpFormat()
	s.router.GET("/brotli", httpFormat.HandleBrotli)
	s.router.GET("/deflate", httpFormat.HandleDeflate)
	s.router.GET("/deny", httpFormat.HandleDeny)
	s.router.GET("/gzip", httpFormat.HandleGzip)
	s.router.GET("/html", httpFormat.HandleHtml)
	s.router.GET("/json", httpFormat.HandleJson)
	s.router.GET("/encoding/utf8", httpFormat.handleUTF8)
	s.router.GET("/robots.txt", httpFormat.HandleRobotTxt)
	s.router.GET("/xml", httpFormat.HandleXML)

	// Images
	httpImage := NewHttpImage()
	s.router.GET("/image", httpImage.HandleImage)
	s.router.GET("/image/png", httpImage.HandleImagePNG)
	s.router.GET("/image/jpeg", httpImage.HandleImageJPEG)
	s.router.GET("/image/svg", httpImage.HandleImageSVG)
	s.router.GET("/image/webp", httpImage.HandleImageWebp)

	// Anything
	httpAnything := NewHttpAnything()
	s.router.Any("/anything", httpAnything.HandleAnything)

	// Dynamic
	httpDynamic := NewDynamic()
	s.router.GET("/base64/:value", httpDynamic.HandleBase64)
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}



