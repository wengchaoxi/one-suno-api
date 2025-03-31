package dto

type CreateAudioRequest struct {
	ProviderID          string `json:"provider_id"`
	Model               string `json:"mv"`
	Title               string `json:"title"`
	Prompt              string `json:"prompt"`
	Tags                string `json:"tags"`
	NegativeTags        string `json:"negative_tags"`
	IsCustom            bool   `json:"custom"`
	IsInstrument        bool   `json:"make_instrumental"`
	ContinueAt          int    `json:"continue_at"`
	ContinueClipId      string `json:"continue_clip_id"` // Deprecated: use AudioId instead
	AudioId             string `json:"audio_id"`
	PersonaId           string `json:"persona_id,omitempty"`
	ReplaceSectionEnd   int    `json:"replace_section_end,omitempty"`
	ReplaceSectionStart int    `json:"replace_section_start,omitempty"`
	CallbackUrl         string `json:"callback_url,omitempty"`
}

type CreateAudioResponse struct {
	ProviderID string      `json:"provider_id"`
	Data       []AudioData `json:"data"`
}

type AudioData struct {
	Id                string        `json:"id"`
	ModelName         string        `json:"model_name"`
	MajorModelVersion string        `json:"major_model_version"`
	Title             string        `json:"title"`
	AudioUrl          string        `json:"audio_url"`
	ImageUrl          string        `json:"image_url"`
	ImageLargeUrl     string        `json:"image_large_url"`
	VideoUrl          string        `json:"video_url"`
	Metadata          AudioMetadata `json:"metadata,omitempty"`
	CreatedAt         string        `json:"created_at"`
}

type AudioMetadata struct {
	Tags                 string  `json:"tags"`
	Prompt               string  `json:"prompt"`
	GPTDescriptionPrompt string  `json:"gpt_description_prompt"` // Deprecated: use Prompt instead
	Lyric                string  `json:"lyric"`
	Duration             float64 `json:"duration"`
}
