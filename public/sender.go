package public

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type SendData struct {
	MessageType string `json:"message_type"`
	UserId      string `json:"user_id"`
	GroupId     string `json:"group_id"`
	Message     string `json:"message"`
	AutoEscape  string `json:"auto_escape"`
}
type User struct {
	MessageType string `json:"message_type"`
	Sender      Sender `json:"sender"`
	MessageId   int64  `json:"message_id"`
	RawMessage  string `json:"raw_message"`
	GroupID     int64  `json:"group_id"`
}
type Sender struct {
	Age      int    `json:"age"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"`
	UserID   int64  `json:"user_id"`
}
type Context struct {
}

var data string

func space(i interface{}) {

}
func (u User) Start() {
	switch u.MessageType {
	case "private":
		u.Private()
	case "group":
		u.Group()
	default:
		return
	}
}

func (s SendData) Request() {
	url := "http://127.0.0.1:5700/send_msg"
	method := "POST"
	p := make(map[string]string)
	p["message_type"] = s.MessageType
	p["user_id"] = s.UserId
	p["group_id"] = s.GroupId
	p["message"] = s.Message
	p["auto_escape"] = s.AutoEscape
	payload, _ := json.Marshal(p)
	fmt.Println(p)
	client := &http.Client{}
	fmt.Println(strings.NewReader(string(payload)))
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func (s SendData) Init(u User) (S SendData) {
	if u.MessageType == "group" {
		S.GroupId = fmt.Sprintf("%d", u.GroupID)
	} else {
		S.GroupId = ""
	}
	S.UserId = fmt.Sprintf("%d", u.Sender.UserID)
	S.Message = u.RawMessage
	S.AutoEscape = "false"
	return S
}
func (s SendData) Requests(msg string) {
	msg = strings.Replace(msg, "\n", "", 2)
	s.Message = msg
	s.Request()
}
