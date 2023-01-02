package public

import (
	"edulx/web/db"
	"strconv"
	"strings"
	"time"
)

func (u User) Private() {
	s := SendData{}
	s = s.Init(u)
	Types, err := db.Get("Type" + strconv.FormatInt(u.Sender.UserID, 10))
	if err != nil {
		switch u.RawMessage {
		case "GPT":
			s.Requests("GPT模式已开启")
			err := db.Set("Type"+strconv.FormatInt(u.Sender.UserID, 10), "GPT", time.Minute*5)
			if err != nil {
				s.Requests("出错了")
			}
		case "JD查询":
			var jt JdToken
			jt.Login()
			ck := jt.GetCK()
			i := 0
			for _, v := range ck {
				if v.Name == "JD_COOKIE" && v.Remarks == strconv.FormatInt(u.Sender.UserID, 10) {
					//s.Requests(v.Value)
					i++
				}
			}
			if i == 0 {
				s.Requests("你还没有绑定京东账号")
			} else {
				s.Requests("你已经绑定了" + strconv.Itoa(i) + "个京东账号")
			}
		case "JD绑定":
			var jt JdToken
			jt.Login()
			ck := jt.GetCK()
			i := 0
			for _, v := range ck {
				if v.Name == "JD_COOKIE" && v.Remarks == strconv.FormatInt(u.Sender.UserID, 10) {
					//s.Requests(v.Value)
					i++
				}
			}
			if i == 0 {
				s.Requests("你还没有绑定京东账号")
				err := db.Set("Type"+strconv.FormatInt(u.Sender.UserID, 10), "JD绑定", time.Minute)
				if err != nil {
					s.Requests("出错了")
				} else {
					s.Requests("请发送你的京东cookie(一分钟内有效)")
					s.Requests("获取方法:\n\r" +
						"https://www.kistom.com/?p=541\n\r" +
						"只提取cookie中的pt_key和pt_pin\n\r" +
						"格式为pt_key=xxx;pt_pin=xxx;")
				}
			} else {
				s.Requests("你已经绑定了京东账号,请勿重复绑定")
			}
		case "JD查询ALL":
			if u.Sender.UserID == 2508339002 {
				var jt JdToken
				jt.Login()
				ck := jt.GetCK()
				for _, v := range ck {
					if v.Name == "JD_COOKIE" {
						s.Requests("QQ号:" + v.Remarks + "\n\rCookie:\n\r" + v.Value)
					}
				}
			} else {
				s.Requests("你没有权限")
			}
		case "jd日志查询":
			if u.Sender.UserID == 2508339002 {
				var jt JdToken
				jt.Login()
				names := jt.FindLogs()
				if names != nil {
					s.Requests("查询成功")
					s.Requests(jt.FindLog("772"))
				} else {
					s.Requests("查询失败")
				}
			} else {
				s.Requests("你没有权限")
			}
		default:
			s.Requests("未知指令")
			s.Requests("目前支持的指令有：\n\rGPT\n\rJD查询\rJD绑定")
		}
	} else {
		switch Types {
		case "GPT":
			privateGPT(u)
		case "JD绑定":
			var jt JdToken
			jt.Login()
			jt.Add(u.RawMessage, strconv.FormatInt(u.Sender.UserID, 10))
			s.Requests("绑定完成")
			db.Del("Type" + strconv.FormatInt(u.Sender.UserID, 10))
			s.Requests("请查询是否绑定成功")
		}
	}

}
func privateGPT(u User) {
	s := SendData{}
	s = s.Init(u)
	var S SendGPT
	S.Prompt = ""
	uid := strconv.FormatInt(u.Sender.UserID, 10)
	S.Stop = `[" ` + uid + `:"," AI:"]`
	gpt, err := db.Get("GPT" + uid)
	gpts := ""
	if err != nil {
		S.Prompt = uid + ":" + u.RawMessage[1:] + "\n"
		gpts = "AI:" + S.Start()
		err1 := db.Set("GPT"+uid, S.Prompt+gpts+"\n", time.Minute*5)

		if err1 != nil {
			s.Requests("出错了")
		}
	} else {
		S.Prompt = gpt + uid + ":" + u.RawMessage[1:] + "\n"
		gpts = S.Start()
		err1 := db.Set("GPT"+uid, S.Prompt+gpts+"\n", time.Minute*5)
		if err1 != nil {
			s.Requests("出错了")
		}
	}
	gpts = strings.Replace(gpts, "AI:", "", 1)
	s.Requests(gpts)
}
