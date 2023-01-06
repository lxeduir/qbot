package chatgpt

import (
	"edulx/qbot/model/chatgpt"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type ChatGptService struct {
}

func (c *ChatGptService) Request(body chatgpt.ChatGpt) ([]byte, error) {
	payload, _ := json.Marshal(body)
	client := &http.Client{}
	req, err := http.NewRequest(body.Method, body.Url, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", body.Header.ContentTyper)
	req.Header.Add("Authorization", body.Header.Authorization)
	req.Header.Add("Accept-Encoding", body.Header.AcceptEncoding)
	req.Header.Add("Content-Length", body.Header.ContentLength)
	req.Header.Add("Transfer-Encoding", body.Header.TransferEncoding)
	res, err2 := client.Do(req)
	if err2 != nil {
		return nil, err2
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	Body, err := io.ReadAll(res.Body)
	return Body, nil
}
