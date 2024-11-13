package main

import (
	"endpointlab/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Basic HTTP method endpoints
	router.GET("/get", handlers.HandleGet)
	router.POST("/post", handlers.HandlePost)
	router.PUT("/put", handlers.HandlePut)
	router.DELETE("/delete", handlers.HandleDelete)
	router.PATCH("/patch", handlers.HandlePatch)
	
	// Response format endpoints
	// router.GET("/ip", handlers.HandleIP)
	// router.GET("/headers", handlers.HandleHeaders)
	// router.GET("/user-agent", handlers.HandleUserAgent)
	
	// // Form handling
	// router.GET("/forms/post", handleFormPage)
	// router.POST("/forms/post", handleFormSubmit)

	router.Run(":8080")
}


