package service

import "github.com/gin-gonic/gin"

func (s *Service) getUserHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *Service) updateUserHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
