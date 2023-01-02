package public

import (
	"edulx/web/qbotdb"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (u User) Group() {
	cq := Parse(u.RawMessage)
	s := SendData{}
	s = s.GroupInit(u)
	mp := cq.GetMap()
	if mp["at"]["qq"] == "2911868466" {
		GroupGPT(u)
	} else if u.RawMessage[0:3] == "fy " {
		u.RawMessage = strings.ReplaceAll(u.RawMessage, "\n", "")
		u.RawMessage = strings.ReplaceAll(u.RawMessage, "\r", "")
		if len(u.RawMessage) > 1024 {
			s.Requests("文章过长")
		}
		s.Requests("原文:" + u.RawMessage[3:] + "\n\r" + "翻译:" + Translation(strings.ReplaceAll(u.RawMessage, "fy ", "")))
	} else if u.RawMessage[0:5] == "music" {
		music := CqCode{
			CQ:    "record",
			Name:  []string{"file"},
			Value: []string{"http://aqqmusic.tc.qq.com/amobile.music.tc.qq.com/M500000HjG8v1DTWRO.mp3?guid=B69D8BC956E47C2B65440380380B7E9A&vkey=1D5EB1935E0E4AC5C075BD6CE71D5F4003BF6D0E0F25F9731BE7566D90EFCBFABFCA671FB8F13FAEE332F275642ED84F537B863932E9399F&uin=1828222534&fromtag=119045"},
		}
		s.Requests(music.Get())
	} else {
		if cq.msg[0][0] == '*' {
			GroupGPT(u)
		}
		if cq.msg[0][0] == '#' {
			s.Requests("直接@机器人或者输入前缀带有*的消息即可调用GPT,\n\r例如：*你好\n\r或者@机器人 你好\n当上下文长度超过1024后自动清除上下文\n其余功能等待开发")
		}
	}

}
func (s SendData) GroupInit(u User) (S SendData) {
	S.GroupId = fmt.Sprintf("%d", u.GroupID)
	S.Message = u.RawMessage
	S.AutoEscape = "false"
	return S
}
func GroupGPT(u User) {
	var s SendData
	s = s.GroupInit(u)
	var S SendGPT
	S.Prompt = ""
	uid := strconv.FormatInt(u.Sender.UserID, 10)
	var cq CqCode
	cq.CQ = "at"
	cq.Name = append(cq.Name, "qq")
	cq.Value = append(cq.Value, strconv.FormatInt(u.Sender.UserID, 10))
	S.Stop = `[" ` + uid + `:"," AI:"]`
	gpt, err := qbotdb.Get("GPT" + uid + "GID" + strconv.FormatInt(u.GroupID, 10))
	gpts := ""
	if strings.Contains(gpt, strconv.FormatInt(u.Sender.UserID, 10)+":"+u.RawMessage[1:]) {
		s.Requests(cq.Get() + "请不要重复发送")
		return
	}
	if err != nil {
		S.Prompt = uid + ":" + u.RawMessage[1:] + "\n"
		gpts = "AI:" + S.Start()
		err1 := qbotdb.Set("GPT"+uid+"GID"+strconv.FormatInt(u.GroupID, 10), S.Prompt+gpts+"\n", time.Minute*5)
		if err1 != nil {
			cq.Value[0] = "2508339002"
			s.Requests(cq.Get() + "出错了")
			return
		}
	} else {
		S.Prompt = gpt + uid + ":" + u.RawMessage[1:] + "\n"
		if len(S.Prompt) > 1024 {
			S.Prompt = uid + ":" + u.RawMessage[1:] + "\n"
			s.Requests(cq.Get() + "对话过长，已重置")
		}
		gpts = S.Start()
		err1 := qbotdb.Set("GPT"+uid+"GID"+strconv.FormatInt(u.GroupID, 10), S.Prompt+gpts+"\n", time.Minute)
		if err1 != nil {
			cq.Value[0] = "2508339002"
			s.Requests(cq.Get() + "出错了")
			return
		}
	}
	gpts = strings.Replace(gpts, "AI:", "", 1)
	gpts = cq.Get() + gpts
	s.Requests(gpts)

}
