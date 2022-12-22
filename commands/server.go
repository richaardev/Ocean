package commands

import (
	"github.com/disgoorg/disgo/discord"
	"ocean/utils/command"
)

func init() {
	command.RegisterCommand(command.Command{
		Name:        "server",
		Description: "_",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionSubCommand{
				Name:        "info",
				Description: "Get info about the bot",
			},
		},

		Runner: handleserver,
	})
}

func handleserver(ctx *command.Context) error {
	switch ctx.GetSubCommandName() {
	case "info":
		return ctx.Event.CreateMessage(
			discord.NewMessageCreateBuilder().
				SetContent("OK").
				Build(),
		)
	default:
		return nil
	}
}
