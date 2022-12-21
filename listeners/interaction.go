package listeners

import (
	"bot/utils/command"

	"github.com/disgoorg/disgo/events"
)

func InteractionListener(e *events.ApplicationCommandInteractionCreate) {
	command.InteractionEvent <- e
}