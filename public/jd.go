package public

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type JdLogin struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
type JdToken struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}
type JdCK struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Remarks string `json:"remarks"`
	Id      string `json:"_id"`
}
type FindLog struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var JdLoginData = JdLogin{
	ClientId:     "A5rpz3-l-oVM",
	ClientSecret: "2gz99mj_4YfDClRzSI0e8Ttt",
}

func (j *JdToken) Login() {
	url := "http://123.249.92.218:5700/open/auth/token"
	method := "GET"
	client := &http.Client{}
	url = url + "?client_id=" + JdLoginData.ClientId + "&client_secret=" + JdLoginData.ClientSecret
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	res, err2 := client.Do(req)
	if err2 != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	mp := make(map[string]interface{})
	err = json.Unmarshal(body, &mp)
	if err != nil {
		return
	}
	data := mp["data"].(map[string]interface{})
	j.Token = data["token"].(string)
	j.TokenType = data["token_type"].(string)
	return
}
func (j JdToken) GetCK() []JdCK {
	url := "http://123.249.92.218:5700/open/envs"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}
	req.Header.Add("Authorization", j.TokenType+" "+j.Token)
	res, err2 := client.Do(req)
	if err2 != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	mp := make(map[string]interface{})
	err = json.Unmarshal(body, &mp)
	if err != nil {
		return nil
	}
	data := mp["data"]
	datas := data.([]interface{})
	var ck []JdCK
	for _, v := range datas {
		var c JdCK
		d := v.(map[string]interface{})
		c.Name = d["name"].(string)
		c.Value = d["value"].(string)
		c.Remarks = d["remarks"].(string)
		c.Id = strconv.Itoa(int(d["id"].(float64)))
		ck = append(ck, c)
	}
	return ck
}
func (j JdToken) Add(Value string, Remarks string) {
	url := "http://123.249.92.218:5700/open/envs"
	method := "POST"
	client := &http.Client{}
	payload := strings.NewReader(`[{
		"name": "` + "JD_COOKIE" + `",
		"value": "` + Value + `",
		"remarks": "` + Remarks + `" }]`)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", j.TokenType+" "+j.Token)
	req.Header.Add("Content-Type", "application/json")
	res, err2 := client.Do(req)
	if err2 != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	println(string(body))

}
func (j JdToken) FindLogs() (names []string) {
	url := "http://123.249.92.218:5700/open/crons"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}
	req.Header.Add("Authorization", j.TokenType+" "+j.Token)
	res, err2 := client.Do(req)
	if err2 != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	var mp map[string]interface{}
	err = json.Unmarshal(body, &mp)
	if err != nil {
		return nil
	}
	datas := mp["data"].([]interface{})
	for _, v := range datas {
		d := v.(map[string]interface{})
		Name := d["name"].(string)
		names = append(names, Name)
	}
	return names

}
func (j JdToken) FindLog(id string) string {
	url := "http://123.249.92.218:5700/open/crons" + "/" + id + "/log"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "查询失败"
	}
	req.Header.Add("Authorization", j.TokenType+" "+j.Token)
	res, err2 := client.Do(req)
	if err2 != nil {
		return "查询失败"
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "查询失败"
	}
	var mp map[string]interface{}
	err = json.Unmarshal(body, &mp)
	if err != nil {
		return "查询失败"
	}
	s := mp["data"].(string)
	var strs []string
	s1 := strings.Index(s, "--f--")
	s2 := strings.Index(s[s1+4:], "--f--")
	space(strs)
	space(s2)
	return s
}
func (j JdToken) RunLog(id string) string {
	url := "http://123.249.92.218:5700/open/crons/run"
	method := "PUT"
	client := &http.Client{}
	p := strings.NewReader(`[` + id + `]`)
	req, err := http.NewRequest(method, url, p)
	if err != nil {
		return "查询失败"
	}
	req.Header.Add("Authorization", j.TokenType+" "+j.Token)
	res, err2 := client.Do(req)
	if err2 != nil {
		return "查询失败"
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	fmt.Println(body)
	mp := make(map[string]interface{})
	err = json.Unmarshal(body, &mp)
	fmt.Println(mp)
	return ""
}
