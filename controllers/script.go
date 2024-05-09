package controllers

import (
	"RAG/services"
	"github.com/gin-gonic/gin"
)

func DreamService(r *gin.Engine) {
	r.POST("/model", func(c *gin.Context) {
		services.DreamDream(c)
	})
}
