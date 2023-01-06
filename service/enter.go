package service

import (
	"edulx/qbot/plugin/sender"
	"edulx/qbot/service/chatgpt"
)

type ServiceGroup struct {
	ChatGptPlugGroup chatgpt.PlugGroup
	SenderPlugGroup  sender.PlugGroup
} //存放插件

var ServiceGroupApp = new(ServiceGroup)

//service 是存放对数据库存在操作的插件如若没有就不必创建对应的接口
