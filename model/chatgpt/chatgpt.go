package chatgpt

import "edulx/qbot/model/common"

var KEY = "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"

type ChatGpt struct {
	Url    string
	Header common.Head
	Method string
	Body   ChatGptBody
}

type ChatGptBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      int     `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	Stop             string  `json:"stop"`
}
