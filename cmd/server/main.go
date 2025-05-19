package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sickodev/dasherby-backend/internal/api"
	"github.com/sickodev/dasherby-backend/internal/config"
	"github.com/sickodev/dasherby-backend/internal/supabase"
)

func main() {
	cfg := config.Load()

	supabase.Init(cfg)

	router := gin.Default()
	api.RegisterRoutes(router)

	log.Fatal(router.Run(":" + cfg.Port))
}
