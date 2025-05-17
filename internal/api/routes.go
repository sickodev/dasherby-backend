package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sickodev/dasherby-backend/internal/api/handlers"
)

func RegisterRoutes(router *gin.Engine){
	// Health Check
	router.GET("/", func(c *gin.Context){
		c.String(200, "Dashboard backend running")
	})

	// Error Check
	router.GET("/error", func(c *gin.Context){
		c.String(500, "Ughh!! Another boo boo!")
	})

	//Upload Route
	router.POST("/upload", handlers.UploadHandler)
}