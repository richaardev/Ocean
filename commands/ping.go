package commands

import (
	handler "bot/utils/command"

	"github.com/disgoorg/disgo/discord"
)

func init() {
	handler.RegisterCommand(handler.Command{
		Name: "ping",
		Description: "Pings the bot",
		Runner: pinghandle,
	})
}

func pinghandle(ctx *handler.CommandContext) {
	ctx.Event.CreateMessage(
		discord.NewMessageCreateBuilder().SetContent("Pong!").
		Build(),
	)
}