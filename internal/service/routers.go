package service

import (
	"github.com/gin-gonic/gin"

	"github.com/wengchaoxi/one-suno-api/internal/service/middleware"
)

func (s *Service) Init(engine *gin.Engine) {
	engine.Use(middleware.Cors())
	engine.Use(middleware.Auth(s.opts.ApiKey, "/"))

	engine.GET("/", func(ctx *gin.Context) {
		ctx.File("./static/index.html")
	})
	engine.Static("/assets", "./static/assets")

	v1 := engine.Group("/v1")

	// manage users
	users := v1.Group("/users")
	users.POST("/token", s.createUserTokenHandler)
	users.PATCH("/:id", s.updateUserHandler)

	// manage providers
	providers := v1.Group("/providers")
	providers.POST("/", s.createProviderHandler)
	providers.GET("/", s.getProvidersHandler)
	providers.PATCH("/:id", s.updateProviderHandler)
	providers.DELETE("/:id", s.deleteProviderHandler)

	// manage api keys
	keys := v1.Group("/keys")
	keys.POST("/", s.createApiKeyHandler)
	keys.GET("/", s.getApiKeysHandler)
	keys.PATCH("/:id", s.updateApiKeyHandler)
	keys.DELETE("/:id", s.deleteApiKeyHandler)

	// audios
	v1.POST("/audios", s.createAudioHandler)
	// Deprecated: use /v1/audios instead
	engine.POST("/api/ai_music", s.createAudioHandler)
}
