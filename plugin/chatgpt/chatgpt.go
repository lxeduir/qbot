package chatgpt

import (
	"edulx/qbot/model/chatgpt"
	json2 "edulx/qbot/model/chatgpt/json"
	"edulx/qbot/service"
	"encoding/json"
)

type ChatGpt struct {
}

var ChatGptService = service.ServiceGroupApp.ChatGptPlugGroup

func (c *ChatGpt) PrivateChat(body chatgpt.ChatGpt) string {
	Body, err := ChatGptService.Request(body)
	if err != nil {
		return "出现错误"
	}
	var m map[string]interface{}
	_ = json.Unmarshal(Body, &m)
	if m["error"] != nil {
		M := m["error"].(map[string]interface{})
		return M["message"].(string)
	} else {
		var c json2.ChatGpt
		_ = json.Unmarshal(Body, &c)
		return c.Choices[0].Text
	}
}
