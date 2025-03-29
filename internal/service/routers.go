package service

import (
	"github.com/gin-gonic/gin"

	"github.com/wengchaoxi/one-suno-api/internal/service/middleware"
)

func (s *Service) Init(engine *gin.Engine) {
	engine.Use(middleware.Cors())
	engine.Use(middleware.Auth(s.opts.ApiKey, "/"))

	engine.GET("/", s.whoamiHandler())

	// Deprecated: use /v1/create instead
	engine.POST("/api/ai_music", s.createAudioHandler)

	// v1
	v1 := engine.Group("/v1")

	// create
	v1.POST("/create", s.createAudioHandler)
}
