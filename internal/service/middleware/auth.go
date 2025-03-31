package middleware

import (
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/wengchaoxi/one-suno-api/internal/dto"
)

func Auth(apiKey string, excludePath ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(apiKey) == 0 || slices.Contains(excludePath, ctx.Request.URL.Path) {
			ctx.Next()
			return
		}

		// http header x-api-key: xxx
		if apiKey == ctx.Request.Header.Get("x-api-key") {
			ctx.Next()
			return
		}

		// http header Authorization: Bearer xxx
		if len(ctx.Request.Header.Get("Authorization")) > 7 {
			if apiKey == ctx.Request.Header.Get("Authorization")[7:] {
				ctx.Next()
				return
			}
		}

		dto.SendResponse(ctx, dto.STATUS_UNAUTHORIZED, nil)
		ctx.Abort()
	}
}
