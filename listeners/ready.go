package listeners

import (
	"github.com/richaardev/Ocean/utils/telemetry"

	"github.com/disgoorg/disgo/events"
)

func ReadyListener(e *events.Ready) {
	telemetry.Infof("Logged in as %s#%s", e.User.Username, e.User.Discriminator)
}
