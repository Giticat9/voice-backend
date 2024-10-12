package handlers

import "github.com/gin-gonic/gin"

func LoginHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
