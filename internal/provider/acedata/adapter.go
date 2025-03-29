package acedata

import (
	"errors"

	"github.com/wengchaoxi/one-suno-api/internal/dto"
)

func (a *AcedataProvider) CreateAudio(req *dto.CreateAudioRequest) (*dto.CreateAudioResponse, error) {
	acedataReq := &AcedataCreateSunoAudioRequest{
		Action:              "generate",
		Prompt:              req.Prompt,
		Model:               req.Model,
		Lyric:               req.Prompt,
		Custom:              req.IsCustom,
		Instrument:          req.IsInstrument,
		Title:               req.Title,
		Style:               req.Tags,
		StyleNegative:       req.TagsNegative,
		PersonaId:           req.PersonaId,
		ContinueAt:          req.ContinueAt,
		ReplaceSectionEnd:   req.ReplaceSectionEnd,
		ReplaceSectionStart: req.ReplaceSectionStart,
		CallbackUrl:         req.CallbackUrl,
	}

	if req.AudioId != "" {
		acedataReq.AudioId = req.AudioId
	} else {
		acedataReq.AudioId = req.ContinueClipId
	}

	if acedataReq.AudioId != "" {
		acedataReq.Action = "extend"
	}

	acedataResp, err := a.CreateSunoAudio(acedataReq)
	if err != nil {
		return &dto.CreateAudioResponse{}, err
	}
	if !acedataResp.Success {
		return &dto.CreateAudioResponse{}, errors.New("acedata create audio failed")
	}

	var resp dto.CreateAudioResponse
	for _, d := range acedataResp.Data {
		audio := dto.AudioData{
			Id:                d.Id,
			ModelName:         d.Model,
			MajorModelVersion: d.Model,
			Title:             d.Title,
			AudioUrl:          d.AudioUrl,
			ImageUrl:          d.ImageUrl,
			ImageLargeUrl:     d.ImageUrl,
			VideoUrl:          d.VideoUrl,
			Metadata: dto.AudioMetadata{
				Prompt:               d.Prompt,
				Tags:                 d.Style,
				GPTDescriptionPrompt: d.Prompt,
			},
		}
		resp.Data = append(resp.Data, audio)
	}
	return &resp, nil
}
