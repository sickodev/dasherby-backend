package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Upload Endpoint hit",
	})
}