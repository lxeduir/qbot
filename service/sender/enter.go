package sender

import (
	"edulx/qbot/service/sender/group"
	"edulx/qbot/service/sender/private"
)

type PlugGroup struct {
	SenderService
	private.PlugGroupPrivate
	group.PlugGroupGroup
}
