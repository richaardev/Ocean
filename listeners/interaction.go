package listeners

import (
    "ocean/utils/command"

	"github.com/disgoorg/disgo/events"
)

func InteractionListener(e *events.ApplicationCommandInteractionCreate) {
	command.InteractionEvent <- e
}