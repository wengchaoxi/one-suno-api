package service

import "github.com/gin-gonic/gin"

func (s *Service) whoamiHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	}
}
