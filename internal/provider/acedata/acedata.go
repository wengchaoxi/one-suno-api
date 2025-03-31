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
//		"callback_url": "https://xxx.com/callback"
//		}
type AcedataCreateSunoAudioRequest struct {
	Action              string `json:"action"` // generate, extend, contact, artist_consistency, upload_extend, cover, upload_cover, replace_section, stems
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
//		"success": true,
//		"task_id": "04067c0f-6655-44b6-8c4f-f443af573ebf",
//		"trace_id": "c0a574b6-d5b1-4ddc-9529-b1a3d502a7e9",
//		"data": [
//		  {
//			"id": "8bc267d2-9794-412c-a504-974ce3cba254",
//			"title": "Jingle All the Way",
//			"image_url": "https://cdn2.suno.ai/image_8bc267d2-9794-412c-a504-974ce3cba254.jpeg",
//			"lyric": "[Verse]\nSnowflakes dancing in the street\nWarm cocoa can't be beat\nNeighbours laugh as carolers sing\nJoyful bells begin to ring\n[Verse 2]\nCandy canes and mistletoe\nSparkling lights that gleam and glow\nWrinkled paper ribbons flying\nSmiling faces never lying\n[Chorus]\nJingle all the way tonight\nStars above are shining bright\nGifts of laughter love and cheer\nMaking memories so dear\n[Verse 3]\nStockings hanging by the fire\nEvery heart is full of desire\nSilent wishes whispered low\nUnderneath the mistletoe\n[Bridge]\nChildren dreaming of delight\nSanta's sleigh takes flight tonight\nMagic sparkles in the air\nChristmas wonder everywhere\n[Verse 4]\nWrapping presents with a bow\nFeeling love in every show\nFamily gathered round the tree\nHoliday spirit wild and free",
//			"audio_url": "https://cdn1.suno.ai/8bc267d2-9794-412c-a504-974ce3cba254.mp3",
//			"video_url": "https://cdn1.suno.ai/8bc267d2-9794-412c-a504-974ce3cba254.mp4",
//			"created_at": "2025-01-04T03:26:59.710Z",
//			"model": "chirp-v4",
//			"state": "succeeded",
//			"prompt": "A song for Christmas",
//			"style": "pop",
//			"duration": 181.56
//		  },
//		  {
//			"id": "8ebdcf48-1d4d-4b8a-94bc-ec46a793f590",
//			"title": "Jingle All the Way",
//			"image_url": "https://cdn2.suno.ai/image_8ebdcf48-1d4d-4b8a-94bc-ec46a793f590.jpeg",
//			"lyric": "[Verse]\nSnowflakes dancing in the street\nWarm cocoa can't be beat\nNeighbours laugh as carolers sing\nJoyful bells begin to ring\n[Verse 2]\nCandy canes and mistletoe\nSparkling lights that gleam and glow\nWrinkled paper ribbons flying\nSmiling faces never lying\n[Chorus]\nJingle all the way tonight\nStars above are shining bright\nGifts of laughter love and cheer\nMaking memories so dear\n[Verse 3]\nStockings hanging by the fire\nEvery heart is full of desire\nSilent wishes whispered low\nUnderneath the mistletoe\n[Bridge]\nChildren dreaming of delight\nSanta's sleigh takes flight tonight\nMagic sparkles in the air\nChristmas wonder everywhere\n[Verse 4]\nWrapping presents with a bow\nFeeling love in every show\nFamily gathered round the tree\nHoliday spirit wild and free",
//			"audio_url": "https://cdn1.suno.ai/8ebdcf48-1d4d-4b8a-94bc-ec46a793f590.mp3",
//			"video_url": "https://cdn1.suno.ai/8ebdcf48-1d4d-4b8a-94bc-ec46a793f590.mp4",
//			"created_at": "2025-01-04T03:26:59.710Z",
//			"model": "chirp-v4",
//			"state": "succeeded",
//			"prompt": "A song for Christmas",
//			"style": "pop",
//			"duration": 203.8
//		  }
//		]
//	}
type AcedataSunoAudioData struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	ImageUrl  string  `json:"image_url"`
	Lyric     string  `json:"lyric"`
	AudioUrl  string  `json:"audio_url"`
	VideoUrl  string  `json:"video_url"`
	CreatedAt string  `json:"created_at"`
	Model     string  `json:"model"`
	State     string  `json:"state"` // succeeded, pending, running, error
	Prompt    string  `json:"prompt"`
	Style     string  `json:"style"`
	Duration  float64 `json:"duration"`
}

//	"error": {
//	    "code": "forbidden",
//	    "message": "Song Description contained artist name: eminem"
//	}
type AcedataError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AcedataCreateSunoAudioResponse struct {
	Success bool                   `json:"success"`
	TaskId  string                 `json:"task_id"`
	TraceId string                 `json:"trace_id"`
	Data    []AcedataSunoAudioData `json:"data,omitempty"`
	Error   AcedataError           `json:"error,omitempty"`
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
