package controllers

import "github.com/gin-gonic/gin"

func NewController() {
	r := gin.Default()

	DreamService(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
