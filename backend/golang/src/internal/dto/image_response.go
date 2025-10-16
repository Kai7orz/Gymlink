package dto

type ImageData struct {
	B64Json string `json:"b64_json"`
}

type ImageUsage struct {
	TotalTokens       int                `json:"total_tokens"`
	InputTokens       int                `json:"input_tokens"`
	OutputTokens      int                `json:"output_tokens"`
	InputTokensDetail InputTokensDetails `json:"input_tokens_details"`
}

type InputTokensDetails struct {
	TextTokens  int `json:"text_tokens"`
	ImageTokens int `json:"image_tokens"`
}

type ImageResponseType struct {
	Created      int64       `json:"created"`
	Data         []ImageData `json:"data"`
	Background   string      `json:"background"`
	OutputFormat string      `json:"output_format"`
	Size         string      `json:"size"`
	Quality      string      `json:"quality"`
	Usage        ImageUsage  `json:"usage"`
}
