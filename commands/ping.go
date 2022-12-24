package commands

import (
  "github.com/disgoorg/disgo/discord"
  handler "github.com/richaardev/Ocean/utils/command"
)

func init() {
	handler.RegisterCommand(handler.Command{
		Name:        "ping",
		Description: "Pings the bot",
    
		Runner:      pinghandle,
	})
}

func pinghandle(ctx *handler.Context) (err error) {
	return ctx.Event.CreateMessage(
		discord.NewMessageCreateBuilder().SetContentf(ctx.T("commands:ping.response"), ctx.Client().Gateway().Latency().Milliseconds()).
			Build(),
	)
}
