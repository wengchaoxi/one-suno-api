package acedata

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/wengchaoxi/one-suno-api/internal/provider"
)

type AcedataProviderOptions struct {
	BaseUrl       string
	PlatformToken string
	AppId         string
	AppToken      string
}

type AcedataProvider struct {
	opts    *AcedataProviderOptions
	httpCli *http.Client
}

func NewAcedataProvider(opts *AcedataProviderOptions) *AcedataProvider {
	return &AcedataProvider{
		opts: opts,
		httpCli: provider.NewHTTPClient(
			map[string]string{
				"Authorization": "Bearer " + opts.AppToken,
				"Accept":        "application/json",
				"Content-Type":  "application/json",
			},
		),
	}
}

//	{
//		"action": "generate",
//		"prompt": "A song for Christmas",
//		"model": "chirp-v4",
//		"lyric": "lyricstring",
//		"custom": true,
//		"instrumental": true,
//		"title": "titlestring",
//		"style": "stylestring",
//		"style_negative": "style_negativestring",
//		"audio_id": "audio_idstring",
//		"persona_id": "persona_idstring",
//		"continue_at": 1,
//		"replace_section_end": 2,
//		"replace_section_start": 3,
//		"callback_url": "http://callback.com"
//	  }
type AcedataCreateSunoAudioRequest struct {
	Action              string `json:"action"`
	Prompt              string `json:"prompt"`
	Model               string `json:"model"`
	Lyric               string `json:"lyric"`
	Custom              bool   `json:"custom"`
	Instrument          bool   `json:"instrumental"`
	Title               string `json:"title"`
	Style               string `json:"style"`
	StyleNegative       string `json:"style_negative"`
	AudioId             string `json:"audio_id"`
	PersonaId           string `json:"persona_id"`
	ContinueAt          int    `json:"continue_at"`
	ReplaceSectionEnd   int    `json:"replace_section_end"`
	ReplaceSectionStart int    `json:"replace_section_start"`
	CallbackUrl         string `json:"callback_url"`
}

//	{
//		"data": [
//		  {
//			"id": "75d8e08f-b25f-450e-9496-7b52e393098b",
//			"lyric": "[Verse]\nSleigh bells ringin', choirs singin'\nSnowflakes fallin', presents glistenin' (glistenin')\nIn the air, there's a feeling of joy\nSpreadin' love to every girl and boy\n[Verse 2]\nCandles glowin', fire cracklin'\nStockings hangin', children wrappin'\nWith a smile, they unwrap their surprise\nIn their hearts, the magic never dies\n[Chorus]\nJingle all the way (jingle all the way)\nIn the winter wonderland, we play (oh-oh)\nHear the carols echo through the night (echo through the night)\nMerry Christmas, oh what a delight (oh-oh-oh)",
//			"model": "chirp-v3",
//			"style": "pop upbeat",
//			"title": "Jingle All the Way",
//			"prompt": "a christmas song",
//			"audio_url": "https://audiopipe.suno.ai/?item_id=75d8e08f-b25f-450e-9496-7b52e393098b",
//			"image_url": "https://cdn1.suno.ai/image_75d8e08f-b25f-450e-9496-7b52e393098b.png",
//			"video_url": "",
//			"created_at": "2024-04-03T11:54:30.424Z"
//		  },
//		  {
//			"id": "e639fefd-bbd3-4858-b16d-45e7d4aa9313",
//			"lyric": "[Verse]\nSleigh bells ringin', choirs singin'\nSnowflakes fallin', presents glistenin' (glistenin')\nIn the air, there's a feeling of joy\nSpreadin' love to every girl and boy\n[Verse 2]\nCandles glowin', fire cracklin'\nStockings hangin', children wrappin'\nWith a smile, they unwrap their surprise\nIn their hearts, the magic never dies\n[Chorus]\nJingle all the way (jingle all the way)\nIn the winter wonderland, we play (oh-oh)\nHear the carols echo through the night (echo through the night)\nMerry Christmas, oh what a delight (oh-oh-oh)",
//			"model": "chirp-v3",
//			"style": "pop upbeat",
//			"title": "Jingle All the Way",
//			"prompt": "a christmas song",
//			"audio_url": "https://audiopipe.suno.ai/?item_id=e639fefd-bbd3-4858-b16d-45e7d4aa9313",
//			"image_url": "https://cdn1.suno.ai/image_e639fefd-bbd3-4858-b16d-45e7d4aa9313.png",
//			"video_url": "",
//			"created_at": "2024-04-03T11:54:30.424Z"
//		  }
//		],
//		"success": true
//	  }
type AcedataCreateSunoAudioResponse struct {
	Data []struct {
		Id        string `json:"id"`
		Lyric     string `json:"lyric"`
		Model     string `json:"model"`
		Style     string `json:"style"`
		Title     string `json:"title"`
		Prompt    string `json:"prompt"`
		AudioUrl  string `json:"audio_url"`
		ImageUrl  string `json:"image_url"`
		VideoUrl  string `json:"video_url"`
		CreatedAt string `json:"created_at"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (a *AcedataProvider) CreateSunoAudio(req *AcedataCreateSunoAudioRequest) (*AcedataCreateSunoAudioResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return &AcedataCreateSunoAudioResponse{}, err
	}

	resp, err := a.httpCli.Post(a.opts.BaseUrl+"/suno/audios", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &AcedataCreateSunoAudioResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return &AcedataCreateSunoAudioResponse{}, errors.New(string(body))
	}

	var data AcedataCreateSunoAudioResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return &AcedataCreateSunoAudioResponse{}, err
	}

	return &data, nil
}
