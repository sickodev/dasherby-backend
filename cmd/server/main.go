package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sickodev/dasherby-backend/internal/config"
)

func main() {
	cfg := config.Load()

	router := gin.Default()

	// Health Check
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Running on port %s", cfg.Port)
	})

	log.Fatal(router.Run(":" + cfg.Port))
}
