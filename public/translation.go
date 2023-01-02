package public

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var APP_ID = "20230104001519150"
var SECURITY = "K4P9_8g7qoe9YHsDMwc5"

type TranslationData struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func Translation(q string) string {
	url := "http://api.fanyi.baidu.com/api/trans/vip/translate"
	method := "POST"
	p := make(map[string]string)
	p["q"] = q
	p["appid"] = APP_ID
	p["salt"] = strconv.FormatInt(time.Now().Unix(), 10)
	p["sign"] = MD5(APP_ID + q + p["salt"] + SECURITY)
	payload := strings.NewReader("q=" + q + "&from=auto&to=zh&appid=" + APP_ID + "&salt=" + p["salt"] + "&sign=" + p["sign"])
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "翻译失败"
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err2 := client.Do(req)
	if err2 != nil {

		return "翻译失败"
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "翻译失败"
	}
	data := TranslationData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "翻译失败"
	}
	return data.TransResult[0].Dst
}
func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
