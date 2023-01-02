package public

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (u User) Group() {
	cq := Parse(u.RawMessage)
	s := SendData{}
	s = s.GroupInit(u)
	s.Message = cq.Get()
	s.GroupRequest()
}
func (s SendData) GroupInit(u User) (S SendData) {
	S.GroupId = fmt.Sprintf("%d", u.GroupID)
	S.Message = u.RawMessage
	S.AutoEscape = "false"
	return S
}
func (s SendData) GroupGPT(u User) {

}
func (s SendData) GroupRequest() {
	url := "http://127.0.0.1:5700/send_group_msg"
	method := "POST"
	p := make(map[string]string)
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
