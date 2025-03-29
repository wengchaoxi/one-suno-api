package middleware

import (
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/wengchaoxi/one-suno-api/internal/dto"
)

func Auth(apiKey string, excludePath ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if slices.Contains(excludePath, ctx.Request.URL.Path) {
			ctx.Next()
			return
		}

		k := ctx.Request.Header.Get("x-api-key")
		if len(apiKey) > 0 && apiKey == k {
			ctx.Next()
			return
		}

		dto.SendResponse(ctx, dto.STATUS_UNAUTHORIZED, nil)
		ctx.Abort()
	}
}
