package services

import (
	"RAG/scripts"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Msg struct {
	Msg string `json:"msg"`
}

func DreamDream(c *gin.Context) {
	var text Msg

	err := c.ShouldBindJSON(&text)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	msg := scripts.RunDream(text.Msg + "이 꿈을 문장으로 해몽 해줘")
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
