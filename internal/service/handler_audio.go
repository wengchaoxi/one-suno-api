package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/wengchaoxi/one-suno-api/pkg/logger"

	"github.com/wengchaoxi/one-suno-api/internal/dto"
)

func (s *Service) createAudioHandler(ctx *gin.Context) {
	var req dto.CreateAudioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.ErrorWithRequest(ctx, "create audio", err, req)
		dto.SendResponse(ctx, dto.STATUS_SERVER_ERROR, nil)
		return
	}

	meta := s.opts.ProviderManager.Select()
	if meta == nil {
		logger.ErrorWithRequest(ctx, "create audio", errors.New("no provider"), req)
		dto.SendResponse(ctx, dto.STATUS_SERVER_ERROR, nil)
		return
	}

	resp, err := meta.Provider.CreateAudio(&req)
	if err != nil {
		dto.SendResponse(ctx, dto.STATUS_SERVER_ERROR, nil)
		return
	}

	dto.SendResponse(ctx, dto.STATUS_SUCCESS, resp)
}
