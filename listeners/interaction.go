package listeners

import (
	"github.com/richaardev/Ocean/utils/command"

	"github.com/disgoorg/disgo/events"
)

func InteractionListener(e *events.ApplicationCommandInteractionCreate) {
	command.InteractionEvent <- e
}
