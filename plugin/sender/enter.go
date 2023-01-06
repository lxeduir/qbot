package sender

import (
	"edulx/qbot/plugin/sender/group"
	"edulx/qbot/plugin/sender/private"
)

type PlugGroup struct {
	Sender
	private.PlugGroup
	group.Group
}
