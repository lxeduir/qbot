package public

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ChatGpt struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created string   `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}
type Choice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
type SendGPT struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      int     `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	Stop             string  `json:"stop"`
}

func (s SendGPT) Start() string {
	url := "https://api.openai.com/v1/completions"
	method := "POST"
	GPTs := "Marv 是一个聊天机器人，它不情愿地用讽刺的回答来回答问题："
	p := SendGPT{
		Model:            "text-davinci-003",
		Prompt:           GPTs + s.Prompt,
		MaxTokens:        2048,
		Temperature:      0,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0.6,
		Stop:             s.Stop,
	}

	payload, _ := json.Marshal(p)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+"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")
	req.Header.Add("Accept-Encoding", "gzip,deflate")
	req.Header.Add("Content-Length", "1024")
	req.Header.Add("Transfer-Encoding", "chunked")
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var chat ChatGpt
	err = json.Unmarshal(body, &chat)
	var mp map[string]interface{}
	err = json.Unmarshal(body, &mp)
	if mp["error"] != nil {
		mps := mp["error"].(map[string]interface{})
		return mps["message"].(string)
	}
	chats := chat.Choices[0].Text
	return chats
}
