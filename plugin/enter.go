package plugin

import "edulx/qbot/plugin/chatgpt"

type PlugGroup struct {
	ChatGptPlugGroup chatgpt.PlugGroup
}

var PluginApp = new(PlugGroup)
