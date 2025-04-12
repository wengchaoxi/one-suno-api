package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wengchaoxi/one-suno-api/internal/dto"
	"github.com/wengchaoxi/one-suno-api/pkg/logger"
)

func Hash(x string) string {
	return x
	// hash := sha256.Sum256([]byte(x))
	// return hex.EncodeToString(hash[:])
}

func (s *Service) createUserTokenHandler(ctx *gin.Context) {
	var req dto.CreateUserTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.ErrorWithRequest(ctx, "get user", err, req)
		dto.SendResponse(ctx, dto.STATUS_SERVER_ERROR, nil)
		return
	}
	user, err := s.opts.UserRepository.GetByUsername(req.Username)
	if err != nil {
		logger.ErrorWithRequest(ctx, "get user", err, req)
		dto.SendResponse(ctx, dto.STATUS_UNAUTHORIZED, nil)
		return
	}
	if user.Password != Hash(req.Password) {
		dto.SendResponse(ctx, dto.STATUS_UNAUTHORIZED, nil)
		return
	}
	var resp dto.CreateUserTokenResponse
	resp.Token = "tokenxxxxxxxxxxxxxxx"
	dto.SendResponse(ctx, dto.STATUS_SUCCESS, resp)
}

func (s *Service) updateUserHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
